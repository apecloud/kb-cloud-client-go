from generator.formatter import presence_method_name


def test_presence_method_name_keeps_default_when_no_property_conflict():
    model = {
        "properties": {
            "detail": {},
            "title": {},
        }
    }

    assert presence_method_name(model, "detail") == "HasDetail"


def test_presence_method_name_avoids_conflicting_property_name():
    model = {
        "properties": {
            "hasDetail": {},
            "detail": {},
        }
    }

    assert presence_method_name(model, "detail") == "HasDetailSet"


def test_presence_method_name_checks_allof_properties():
    model = {
        "allOf": [
            {
                "properties": {
                    "hasDetail": {},
                }
            },
            {
                "properties": {
                    "detail": {},
                }
            },
        ]
    }

    assert presence_method_name(model, "detail") == "HasDetailSet"
