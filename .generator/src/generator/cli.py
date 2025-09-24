import pathlib

import click
from jinja2 import Environment, FileSystemLoader

from . import openapi
from . import formatter
from . import utils

MODULE = "github.com/apecloud/kb-cloud-client-go"
COMMON_PACKAGE_NAME = "common"


def _model_uses_common_types(model):
    if not model:
        return False
    
    for attr, schema in model.get("properties", {}).items():
        if _schema_uses_common_types(schema):
            return True
    
    for subschema in model.get("allOf", []):
        if _model_uses_common_types(subschema):
            return True
    
    return False


def _schema_uses_common_types(schema):
    if not schema:
        return False
    
    if schema.get("nullable", False):
        return True
    
    for attr, subschema in schema.get("properties", {}).items():
        if _schema_uses_common_types(subschema):
            return True
    
    if "items" in schema:
        if _schema_uses_common_types(schema["items"]):
            return True
    
    return False


def _schema_references_external_types(schema, models):
    if not schema:
        return False
    
    if hasattr(schema, "__reference__"):
        ref_name = schema.__reference__["$ref"].split("/")[-1]
        if ref_name in models:
            return True
    
    for attr, subschema in schema.get("properties", {}).items():
        if _schema_references_external_types(subschema, models):
            return True
    
    # Check iems (for arrays)
    if "items" in schema:
        if _schema_references_external_types(schema["items"], models):
            return True
    
    return False


@click.command()
@click.argument(
    "specs",
    nargs=-1,
    type=click.Path(exists=True, file_okay=True, dir_okay=False, path_type=pathlib.Path),
)
@click.option(
    "-o",
    "--output",
    type=click.Path(path_type=pathlib.Path),
)
def cli(specs, output):
    """
    Generate a Go code snippet from OpenAPI specification.
    """
    env = Environment(loader=FileSystemLoader(str(pathlib.Path(__file__).parent / "templates")))

    env.filters["accept_headers"] = openapi.accept_headers
    env.filters["attribute_name"] = formatter.attribute_name
    env.filters["block_comment"] = formatter.block_comment
    env.filters["camel_case"] = formatter.camel_case
    env.filters["collection_format"] = openapi.collection_format
    env.filters["format_server"] = openapi.format_server
    env.filters["format_value"] = formatter.format_value
    env.filters["is_reference"] = formatter.is_reference
    env.filters["is_primitive"] = formatter.is_primitive
    env.filters["parameter_schema"] = openapi.parameter_schema
    env.filters["parameters"] = openapi.parameters
    env.filters["form_parameter"] = openapi.form_parameter
    env.filters["need_body_parameter"] = openapi.need_body_parameter
    env.filters["response_type"] = openapi.get_type_for_response
    env.filters["responses_by_types"] = openapi.responses_by_types
    env.filters["return_type"] = openapi.return_type
    env.filters["simple_type"] = formatter.simple_type
    env.filters["snake_case"] = formatter.snake_case
    env.filters["untitle_case"] = formatter.untitle_case
    env.filters["upperfirst"] = utils.upperfirst
    env.filters["variable_name"] = formatter.variable_name

    env.globals["enumerate"] = enumerate
    env.globals["get_name"] = openapi.get_name
    env.globals["get_type_for_attribute"] = openapi.get_type_for_attribute
    env.globals["get_type_for_parameter"] = openapi.get_type_for_parameter
    env.globals["get_type"] = openapi.type_to_go
    env.globals["get_default"] = openapi.get_default
    env.globals["get_container"] = openapi.get_container
    env.globals["get_container_type"] = openapi.get_container_type
    env.globals["get_type_at_path"] = openapi.get_type_at_path
    env.globals["common_package_name"] = COMMON_PACKAGE_NAME
    env.globals["module"] = MODULE

    api_j2 = env.get_template("api.j2")
    model_j2 = env.get_template("model.j2")
    doc_j2 = env.get_template("doc.j2")

    extra_files = {
        "client.go": env.get_template("client.j2"),
        "configuration.go": env.get_template("configuration.j2"),
        "utils.go": env.get_template("utils.j2"),
        "encoding_json.go": env.get_template("encoding_json.j2"),
        "goccy_gojson.go": env.get_template("goccy_gojson.j2"),
    }

    test_scenarios_files = {
        "api_mappings.go": env.get_template("scenarios_api_mappings.j2"),
    }

    output.mkdir(parents=True, exist_ok=True)

    all_specs = {}
    all_apis = {}
    all_models = {}
    
    for spec_path in specs:
        spec = openapi.load(spec_path)
        spec_name = spec_path.stem
        models = openapi.models(spec)
        all_models.update(models)
        all_specs[spec_name] = spec

    for spec_path in specs:
        spec = openapi.load(spec_path)
        spec_name = spec_path.stem

        print(f"Generating {spec_name}")
        apis = openapi.apis(spec)
        models = openapi.models(spec)

        env.globals["openapi"] = spec
        
        env.globals["package_name"] = f"kbcloud"
        resources_dir = output / env.globals["package_name"]
        if spec_name == "adminapi":
            env.globals["package_name"] = f"admin"
            resources_dir = output / f"kbcloud" / env.globals["package_name"]
        resources_dir.mkdir(parents=True, exist_ok=True)

        for name, model in models.items():
            filename = "model_" + formatter.model_filename(name) + ".go"
            model_path = resources_dir / filename
            model_path.parent.mkdir(parents=True, exist_ok=True)
            
            imports = set()
            
            if _model_uses_common_types(model):
                imports.add("github.com/apecloud/kb-cloud-client-go/api/common")
            
            for attr, schema in model.get("properties", {}).items():
                if _schema_references_external_types(schema, all_models):
                    imports.add("github.com/apecloud/kb-cloud-client-go/api/common")
            
            with model_path.open("w") as fp:
                print(f"Generating model {name} to {model_path}")
                fp.write(model_j2.render(
                    name=utils.upperfirst(name), 
                    model=model, 
                    models=all_models,
                    imports=list(imports),
                    all_models=all_models
                ))

        all_operations = []

        for name, operations in apis.items():
            print(f"Generating API {name}")
            if name == None:
                print(f"missing tag for {operations}")
                exit(1)
            operations = utils.removeWebSocketOP(operations)
            if len(operations) == 0:
                continue
            filename = "api_" + formatter.snake_case(name) + ".go"
            api_path = resources_dir / filename
            api_path.parent.mkdir(parents=True, exist_ok=True)
            with api_path.open("w") as fp:
                fp.write(api_j2.render(name=utils.upperfirst(name), operations=operations))
            all_operations.append((name, operations))

        doc_path = resources_dir / "doc.go"
        with doc_path.open("w") as fp:
            fp.write(doc_j2.render(all_operations=all_operations))

    common_package_output = pathlib.Path(f"../api/{COMMON_PACKAGE_NAME}")
    common_package_output.mkdir(parents=True, exist_ok=True)
    for name, template in extra_files.items():
        filename = common_package_output / name
        with filename.open("w") as fp:
            fp.write(template.render(apis=all_apis, all_specs=all_specs))

    # scenarios_test_output = pathlib.Path("../tests/scenarios/")
    # for name, template in test_scenarios_files.items():
    #     filename = scenarios_test_output / name
    #     with filename.open("w") as fp:
    #         fp.write(template.render(apis=all_apis, all_specs=all_specs))
