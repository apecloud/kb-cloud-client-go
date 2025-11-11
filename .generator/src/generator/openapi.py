import hashlib
import json
import pathlib
import random
import uuid
import warnings
import yaml

from jsonref import JsonRef
from urllib.parse import urlparse
try:
    from yaml import CSafeLoader as SafeLoader
except ImportError:
    from yaml import SafeLoader

from . import formatter
from . import utils

def load(filename):
    path = pathlib.Path(filename)
    with path.open() as fp:
        return JsonRef.replace_refs(yaml.load(fp, Loader=SafeLoader))


def get_name(schema):
    name = None
    if hasattr(schema, "__reference__"):
        name = schema.__reference__["$ref"].split("/")[-1]
        name = utils.upperfirst(name)

    return name


def type_to_go(schema, alternative_name=None, render_nullable=False, render_new=False):
    if schema is None:
        print(f"schema is None for {alternative_name}")
        raise ValueError(f"Unknown type {schema}")

    """Return Go type name for the type."""
    if render_nullable and schema.get("nullable", False):
        prefix = "Nullable"
    else:
        prefix = ""

    # special case for additionalProperties: true
    if schema is True:
        return "interface{}"

    if "enum" not in schema:
        name = formatter.simple_type(schema, render_nullable=render_nullable, render_new=render_new)
        if name is not None:
            return name

    name = get_name(schema)
    if name:
        if "enum" in schema:
            if render_new and schema.get("nullable", False):
                return f"New{prefix}{name}"
            return prefix + name
        if not (schema.get("additionalProperties") and not schema.get("properties")) and schema.get("type", "object") == "object":
            return prefix + name

    type_ = schema.get("type")
    if type_ is None:
        if "items" in schema:
            type_ = "array"
        elif "properties" in schema:
            type_ = "object"
        else:
            type_ = "object"
            warnings.warn(f"Unknown type for schema: {schema} ({name or alternative_name})")

    if type_ == "array":
        if name and schema.get("x-generate-alias-as-model", False):
            return prefix + name
        if name or alternative_name:
            alternative_name = (name or alternative_name) + "Item"
        name = type_to_go(schema["items"], alternative_name=alternative_name)
        # handle nullable arrays
        if formatter.simple_type(schema["items"]) and schema["items"].get("nullable"):
            name = "*" + name
        if schema.get("nullable") and formatter.is_primitive(schema["items"]):
            name = formatter.simple_type(schema["items"], render_nullable=render_nullable, render_new=render_new)
            if render_nullable:
                return f"common.{prefix}List[{name}]"
        return "[]{}".format(name)
    elif type_ == "object":
        if "additionalProperties" in schema:
            return "map[string]{}".format(type_to_go(schema["additionalProperties"]))
        return (
            prefix + alternative_name
            if alternative_name
            and ("properties" in schema or "oneOf" in schema or "anyOf" in schema or "allOf" in schema)
            else "interface{}"
        )

    raise ValueError(f"Unknown type {type_}")


def get_type_for_attribute(schema, attribute, current_name=None):
    """Return Go type name for the attribute."""
    child_schema = schema.get("properties", {}).get(attribute)
    alternative_name = current_name + utils.upperfirst(attribute) if current_name else None
    return type_to_go(child_schema, alternative_name=alternative_name)


def get_type_for_parameter(parameter):
    """Return Go type name for the parameter."""
    if "content" in parameter:
        assert "in" not in parameter
        for content in parameter["content"].values():
            return type_to_go(content["schema"])
    return type_to_go(parameter.get("schema"))


def get_type_for_response(response):
    """Return Go type name for the response."""
    if "content" in response:
        for content in response["content"].values():
            if "schema" in content:
                return type_to_go(content["schema"])


def responses_by_types(operation):
    result = {}
    for response_code, response in operation["responses"].items():
        if int(response_code) < 300:
            continue
        response_type = get_type_for_response(response)
        if response_type in result:
            result[response_type][1].append(response_code)
        else:
            result[response_type] = [response, [response_code]]
    return result.items()


def child_models(schema, alternative_name=None, seen=None, parent=None):
    if schema is None:
        return

    seen = seen or set()
    current_name = get_name(schema)
    name = current_name or alternative_name

    # schema["name"] = name

    if parent is not None:
        schema["parent"] = parent

    has_sub_models = False
    has_oneof_or_anyof = False
    if "allOf" in schema:
        has_sub_models = True
        for index in range(len(schema["allOf"])):
            yield from child_models(schema["allOf"][index], seen=seen, parent=schema)
    if "oneOf" in schema:
        has_sub_models = True
        has_oneof_or_anyof = True
        for index in range(len(schema["oneOf"])):
            yield from child_models(schema["oneOf"][index], seen=seen, parent=schema)
    if "anyOf" in schema:
        has_sub_models = True
        has_oneof_or_anyof = True
        for index in range(len(schema["anyOf"])):
            yield from child_models(schema["anyOf"][index], seen=seen, parent=schema)

    if "items" in schema:
        if current_name is not None and schema.get("x-generate-alias-as-model", False):
            if name in seen:
                return
            seen.add(name)
            yield name, schema

        yield from child_models(
            schema["items"],
            alternative_name=name + "Item" if name is not None else None,
            seen=seen,
            parent=schema,
        )

    # Check if this is a pure oneOf/anyOf schema without type or properties
    # (e.g., default field with oneOf: [string, integer, number, boolean])
    is_pure_oneof_anyof = (
        has_oneof_or_anyof
        and schema.get("type") is None
        and "properties" not in schema
    )

    # For pure oneOf/anyOf schemas, yield them as separate models but DON'T return early
    # We still need to process sub-schemas
    if is_pure_oneof_anyof and name is not None and name not in seen:
        seen.add(name)
        yield name, schema
        # Continue to process oneOf/anyOf sub-schemas (already done above)
        return  # But we're done with this schema after yielding it

    if (schema.get("type") == "object" or "properties" in schema or has_sub_models) and (
        not (schema.get("additionalProperties") and not schema.get("properties"))
    ):
        if not has_sub_models and name is None:
            # this is a basic map object so we don't need a type
            return

        if name is None:
            raise ValueError(f"Schema {schema} has no name")

        if name in seen:
            return

        if "properties" in schema or has_sub_models:
            seen.add(name)
            yield name, schema

        # Process properties from the schema itself
        for key in schema.get("properties", {}):
            yield from child_models(
                schema["properties"][key],
                alternative_name=name + utils.upperfirst(key),
                seen=seen,
                # parent=schema,
            )

        # Also process properties from allOf sub-schemas to catch nested types like oneOf
        if "allOf" in schema:
            for subschema in schema["allOf"]:
                for key in subschema.get("properties", {}):
                    yield from child_models(
                        subschema["properties"][key],
                        alternative_name=name + utils.upperfirst(key),
                        seen=seen,
                        # parent=schema,
                    )

    if "enum" in schema:
        if name is None:
            raise ValueError(f"Schema {schema} has no name")

        if name in seen:
            return

        seen.add(name)
        yield name, schema

    if "additionalProperties" in schema:
        nested_name = get_name(schema["additionalProperties"])
        if nested_name:
            yield from child_models(
                schema["additionalProperties"],
                seen=seen,
                # parent=schema,
            )


def models(spec):
    name_to_schema = {}

    # First collect models from API paths
    for path in spec["paths"]:
        if path.startswith("x-"):
            continue
        for method in spec["paths"][path]:
            operation = spec["paths"][path][method]

            for content in operation.get("parameters", []):
                if "schema" in content:
                    name_to_schema.update(dict(child_models(content["schema"])))

            for content in operation.get("requestBody", {}).get("content", {}).values():
                if "schema" in content:
                    name_to_schema.update(dict(child_models(content["schema"])))

            for response in operation.get("responses", {}).values():
                for content in response.get("content", {}).values():
                    if "schema" in content:
                        name_to_schema.update(dict(child_models(content["schema"])))

    # Also collect all schemas from components/schemas to ensure
    # we don't miss any referenced types like ImportFieldType, ImportStringValidation, etc.
    if "components" in spec and "schemas" in spec["components"]:
        for schema_name, schema_def in spec["components"]["schemas"].items():
            # Add the schema itself
            name_to_schema[schema_name] = schema_def
            # Also collect any child models from this schema
            name_to_schema.update(dict(child_models(schema_def, alternative_name=schema_name)))

    return name_to_schema


def apis(spec):
    operations = {}

    for path in spec["paths"]:
        if path.startswith("x-"):
            continue
        for method in spec["paths"][path]:
            operation = spec["paths"][path][method]
            tag = operation.get("tags", [None])[0]
            operations.setdefault(tag, []).append((path, method, operation))

    return operations


def operation(spec, operation_id):
    for path in spec["paths"]:
        for method in spec["paths"][path]:
            operation = spec["paths"][path][method]
            if operation["operationId"] == operation_id:
                return operation
    return None


def parameters(operation):
    for content in operation.get("parameters", []):
        if "schema" in content and content.get("required"):
            yield content["name"], content

    if "requestBody" in operation:
        for content_type in operation["requestBody"]["content"]:
            if content_type == "multipart/form-data":
                parent = operation["requestBody"]["content"]["multipart/form-data"]["schema"]
                for name, schema in parent["properties"].items():
                    yield name, {
                        "in": "form",
                        "schema": schema,
                        "name": name,
                        "description": schema.get("description"),
                        "required": name in parent.get("required", []),
                    }
            else:
                name = operation.get("x-codegen-request-body-name", "body")
                yield name, operation["requestBody"]

    for content in operation.get("parameters", []):
        if "schema" in content and not content.get("required"):
            yield content["name"], content


def form_parameter(operation):
    if "requestBody" in operation and "multipart/form-data" in operation["requestBody"]["content"]:
        parent = operation["requestBody"]["content"]["multipart/form-data"]["schema"]
        [(name, schema)] = list(parent["properties"].items())
        return {
            "schema": schema,
            "name": name,
            "description": schema.get("description"),
            "required": name in parent.get("required", []),
        }


def need_body_parameter(operation):
    if "requestBody" in operation:
        for content_type in operation["requestBody"]["content"]:
            if content_type != "multipart/form-data":
                return True
    return False


def parameter_schema(parameter):
    if "schema" in parameter:
        return parameter["schema"]
    if "content" in parameter:
        for content in parameter.get("content", {}).values():
            if "schema" in content:
                return content["schema"]
    raise ValueError(f"Unknown schema for parameter {parameter}")


def return_type(operation):
    for response in operation.get("responses", {}).values():
        for content in response.get("content", {}).values():
            if content["schema"] is None:
                print(f"schema is None for {operation.get('operationId')}")
            if "schema" in content:
                return type_to_go(content["schema"])
        return


def accept_headers(operation):
    any_type = "*/*"
    seen = []
    for response in operation.get("responses", {}).values():
        if "content" in response:
            for media_type in response["content"].keys():
                if media_type not in seen:
                    seen.append(media_type)
        else:
            return [any_type]
    return seen


def collection_format(parameter):
    in_to_style = {
        "query": "form",
        "path": "simple",
        "header": "simple",
        "cookie": "form",
    }
    schema = parameter_schema(parameter)
    matrix = {
        ("form", False): "csv",
        ("form", True): "multi",
        # TODO add more cases from https://swagger.io/specification/#parameter-style
    }
    if schema.get("type") == "array" or "items" in schema:
        in_ = parameter.get("in", "query")
        style = parameter.get("style", in_to_style[in_])
        explode = parameter.get("explode", True if style == "form" else False)
        return matrix.get((style, explode), "multi")
    return ""


def format_server(server, server_variables=None, path=""):
    url = server["url"] + path
    # replace potential path variables
    for variable, value in (server_variables or {}).items():
        url = url.replace("{" + variable + "}", value)
    # replace server variables if they were not replace before
    for variable in server["variables"]:
        if server_variables and variable in server_variables:
            continue
        url = url.replace("{" + variable + "}", server["variables"][variable]["default"])
    return urlparse(url)


def server_url_and_method(spec, operation_id, server_index=0, server_variables=None):
    for path in spec["paths"]:
        for method in spec["paths"][path]:
            operation = spec["paths"][path][method]
            if operation["operationId"] == operation_id:
                if "servers" in operation:
                    server = operation["servers"][server_index]
                else:
                    server = spec["servers"][server_index]
                return (
                    format_server(server, server_variables=server_variables, path=path).geturl(),
                    method,
                )

    raise ValueError(f"Operation {operation_id} not found")


def response_code_and_accept_type(operation, status_code=None):
    for response in operation["responses"]:
        if status_code is None:
            return int(response), next(iter(operation["responses"][response].get("content", {None: None})))
        if response == str(status_code):
            return status_code, next(iter(operation["responses"][response].get("content", {None: None})))
    return status_code, None


def request_content_type(operation, status_code=None):
    return next(iter(operation.get("requestBody", {}).get("content", {None: None})))


def response(operation, status_code=None):
    for response in operation["responses"]:
        if status_code is None or response == str(status_code):
            return list(operation["responses"][response]["content"].values())[0]["schema"]
    return None


def get_default(operation, attribute_path):
    attrs = attribute_path.split(".")
    for name, parameter in parameters(operation):
        if name == attrs[0]:
            break
    if name == attribute_path:
        # We found a top level attribute matching the full path, let's use the default
        return parameter["schema"]["default"]

    if name == "body":
        parameter = next(iter(parameter["content"].values()))["schema"]
    for attr in attrs[1:]:
        parameter = parameter["properties"][attr]
    return parameter["default"]


def get_container(operation, attribute_path, container_name="o[0]"):
    attribute_name = attribute_path.split(".")[0]
    for name, parameter in parameters(operation):
        if name == attribute_name and parameter["required"]:
            return '{}.{}'.format(name, ".".join(formatter.attribute_name(a) for a in attribute_path.split(".")[1:]))
    return f'{container_name}.{formatter.attribute_path(attribute_path)}'


def get_container_type(operation, attribute_path, stop=None):
    attrs = attribute_path.split(".")[:stop]
    for name, parameter in parameters(operation):
        if name == attrs[0]:
            break

    if attrs[0] == "body":
        parameter = next(iter(parameter["content"].values()))

    if name == attrs[0] and len(attrs) == 1:
        return type_to_go(parameter["schema"])

    parameter = parameter["schema"]
    for attr in attrs[1:]:
        parameter = parameter["properties"][attr]
    return type_to_go(parameter)


def get_type_at_path(operation, attribute_path):
    content = None
    for code, response in operation.get("responses", {}).items():
        if int(code) >= 300:
            continue
        for content in response.get("content", {}).values():
            if "schema" in content:
                break
    if content is None:
        raise RuntimeError("Default response not found")
    content = content["schema"]
    if not attribute_path:
        return get_name(content.get("items"))
    for attr in attribute_path.split("."):
        content = content["properties"][attr]
    return get_name(content.get("items"))


def get_model_imports(model, current_package="kbcloud"):
    """Get the list of imports required for a model."""
    imports = set()
    visited = set()  # Track visited schemas to prevent infinite recursion

    def check_schema_for_imports(schema, depth=0):
        if schema is None or depth > 10:  # Limit recursion depth
            return

        # Create a schema identifier to track visited schemas
        schema_id = id(schema)
        if schema_id in visited:
            return
        visited.add(schema_id)

        try:
            # Check if this schema uses common types (nullable types)
            if hasattr(schema, 'get') and schema.get("nullable", False):
                imports.add("github.com/apecloud/kb-cloud-client-go/api/common")

            # Check for array items
            if hasattr(schema, 'get') and "items" in schema:
                check_schema_for_imports(schema["items"], depth + 1)

            # Check for object properties
            if hasattr(schema, 'get') and "properties" in schema:
                for prop_schema in schema["properties"].values():
                    check_schema_for_imports(prop_schema, depth + 1)

            # Check for allOf, oneOf, anyOf
            for key in ["allOf", "oneOf", "anyOf"]:
                if hasattr(schema, 'get') and key in schema:
                    for sub_schema in schema[key]:
                        check_schema_for_imports(sub_schema, depth + 1)

            # Check for additionalProperties
            if hasattr(schema, 'get') and "additionalProperties" in schema and schema["additionalProperties"] is not True:
                check_schema_for_imports(schema["additionalProperties"], depth + 1)
        except (AttributeError, RecursionError):
            # Skip problematic schemas
            pass

    # Check the main model
    check_schema_for_imports(model)

    # For allOf models, check if we need to import from the parent package
    if hasattr(model, 'get') and "allOf" in model and current_package == "admin":
        for sub_schema in model["allOf"]:
            if hasattr(sub_schema, "__reference__"):
                # This is a reference to another schema, likely in the parent package
                imports.add("github.com/apecloud/kb-cloud-client-go/api/kbcloud")

    return list(imports)


def generate_value(schema, use_random=False, prefix=None):
    spec = schema.spec
    if not use_random:
        if "example" in spec:
            return spec["example"]
        if "default" in spec:
            return spec["default"]

    if spec["type"] == "string":
        if use_random:
            return str(
                uuid.UUID(
                    bytes=hashlib.sha256(
                        str(prefix or schema.keys).encode("utf-8"),
                    ).digest()[:16]
                )
            )
        return "string"
    elif spec["type"] == "integer":
        return random.randint(0, 32000) if use_random else len(str(prefix or schema.keys))
    elif spec["type"] == "number":
        return random.random() if use_random else 1.0 / len(str(prefix or schema.keys))
    elif spec["type"] == "boolean":
        return True
    elif spec["type"] == "array":
        return [generate_value(schema[0], use_random=use_random)]
    elif spec["type"] == "object":
        return {key: generate_value(schema[key], use_random=use_random) for key in spec["properties"]}
    else:
        raise TypeError(f"Unknown type: {spec['type']}")


class Schema:
    def __init__(self, spec, value=None, keys=None):
        self.spec = spec
        self.value = value if value is not None else generate_value
        self.keys = keys or tuple()

    def __getattr__(self, key):
        return self[key]

    def __getitem__(self, key):
        type_ = self.spec.get("type", "object")
        if type_ == "object":
            try:
                return self.__class__(
                    self.spec["properties"][key],
                    value=self.value,
                    keys=self.keys + (key,),
                )
            except KeyError:
                if "oneOf" in self.spec:
                    for schema in self.spec["oneOf"]:
                        if schema.get("type", "object") == "object":
                            try:
                                return self.__class__(
                                    schema["properties"][key],
                                    value=self.value,
                                    keys=self.keys + (key,),
                                )
                            except KeyError:
                                pass
            raise KeyError(f"{key} not found in {self.spec.get('properties', {}).keys()}: {self.spec}")
        if type_ == "array":
            return self.__class__(self.spec["items"], value=self.value, keys=self.keys + (key,))

        raise KeyError(f"{key} not found in {self.spec}")

    def __repr__(self):
        value = self.value(self)
        if isinstance(value, (dict, list)):
            return json.dumps(value, indent=2)
        return str(value)


class Operation:
    def __init__(self, name, spec, method, path):
        self.name = name
        self.spec = spec
        self.method = method
        self.path = path

    def server_url_and_method(self, spec, server_index=0, server_variables=None):
        def format_server(server, path):
            url = server["url"] + path
            # replace potential path variables
            for variable, value in server_variables.items():
                url = url.replace("{" + variable + "}", value)
            # replace server variables if they were not replace before
            for variable in server["variables"]:
                if variable in server_variables:
                    continue
                url = url.replace(
                    "{" + variable + "}",
                    server["variables"][variable]["default"],
                )
            return url

        server_variables = server_variables or {}
        if "servers" in self.spec:
            server = self.spec["servers"][server_index]
        else:
            server = spec["servers"][server_index]
        return format_server(server, self.path), self.method

    def response_code_and_accept_type(self):
        for response in self.spec["responses"]:
            return int(response), next(iter(self.spec["responses"][response].get("content", {None: None})))
        return None, None

    def request_content_type(self):
        return next(iter(self.spec.get("requestBody", {}).get("content", {None: None})))

    def response(self):
        for response in self.spec["responses"]:
            return Schema(next(iter((self.spec["responses"][response]["content"].values())))["schema"])

    def request(self):
        return Schema(next(iter(self.spec["requestBody"]["content"].values()))["schema"])
