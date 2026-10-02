from generator.openapi import responses_by_types


def test_responses_by_types_skips_status_only_errors():
    operation = {
        "responses": {
            "200": {
                "description": "ok",
                "content": {"application/json": {"schema": {"type": "string"}}},
            },
            "400": {"description": "Invalid request"},
            "401": {
                "description": "Unauthorized",
                "content": {
                    "application/json": {
                        "schema": {
                            "type": "object",
                            "properties": {"code": {"type": "integer"}},
                        }
                    }
                },
            },
        }
    }

    grouped = dict(responses_by_types(operation))

    assert None not in grouped
    assert "interface{}" in grouped
    assert grouped["interface{}"][1] == ["401"]


def test_responses_by_types_skips_empty_content_object():
    operation = {
        "responses": {
            "413": {
                "description": "too large",
                "content": {},
            }
        }
    }

    assert list(responses_by_types(operation)) == []
