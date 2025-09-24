// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportConnectionField - Connection field configuration. Use `oneOf` to enforce strict type-specific properties.
type ImportConnectionField struct {
	ImportStringField  *ImportStringField
	ImportIntegerField *ImportIntegerField
	ImportNumberField  *ImportNumberField
	ImportBooleanField *ImportBooleanField
	ImportEnumField    *ImportEnumField

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ImportStringFieldAsImportConnectionField is a convenience function that returns ImportStringField wrapped in ImportConnectionField.
func ImportStringFieldAsImportConnectionField(v *ImportStringField) ImportConnectionField {
	return ImportConnectionField{ImportStringField: v}
}

// ImportIntegerFieldAsImportConnectionField is a convenience function that returns ImportIntegerField wrapped in ImportConnectionField.
func ImportIntegerFieldAsImportConnectionField(v *ImportIntegerField) ImportConnectionField {
	return ImportConnectionField{ImportIntegerField: v}
}

// ImportNumberFieldAsImportConnectionField is a convenience function that returns ImportNumberField wrapped in ImportConnectionField.
func ImportNumberFieldAsImportConnectionField(v *ImportNumberField) ImportConnectionField {
	return ImportConnectionField{ImportNumberField: v}
}

// ImportBooleanFieldAsImportConnectionField is a convenience function that returns ImportBooleanField wrapped in ImportConnectionField.
func ImportBooleanFieldAsImportConnectionField(v *ImportBooleanField) ImportConnectionField {
	return ImportConnectionField{ImportBooleanField: v}
}

// ImportEnumFieldAsImportConnectionField is a convenience function that returns ImportEnumField wrapped in ImportConnectionField.
func ImportEnumFieldAsImportConnectionField(v *ImportEnumField) ImportConnectionField {
	return ImportConnectionField{ImportEnumField: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ImportConnectionField) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ImportStringField
	err = common.Unmarshal(data, &obj.ImportStringField)
	if err == nil {
		if obj.ImportStringField != nil && obj.ImportStringField.UnparsedObject == nil {
			jsonImportStringField, _ := common.Marshal(obj.ImportStringField)
			if string(jsonImportStringField) == "{}" && string(data) != "{}" { // empty struct
				obj.ImportStringField = nil
			} else {
				match++
			}
		} else {
			obj.ImportStringField = nil
		}
	} else {
		obj.ImportStringField = nil
	}

	// try to unmarshal data into ImportIntegerField
	err = common.Unmarshal(data, &obj.ImportIntegerField)
	if err == nil {
		if obj.ImportIntegerField != nil && obj.ImportIntegerField.UnparsedObject == nil {
			jsonImportIntegerField, _ := common.Marshal(obj.ImportIntegerField)
			if string(jsonImportIntegerField) == "{}" && string(data) != "{}" { // empty struct
				obj.ImportIntegerField = nil
			} else {
				match++
			}
		} else {
			obj.ImportIntegerField = nil
		}
	} else {
		obj.ImportIntegerField = nil
	}

	// try to unmarshal data into ImportNumberField
	err = common.Unmarshal(data, &obj.ImportNumberField)
	if err == nil {
		if obj.ImportNumberField != nil && obj.ImportNumberField.UnparsedObject == nil {
			jsonImportNumberField, _ := common.Marshal(obj.ImportNumberField)
			if string(jsonImportNumberField) == "{}" && string(data) != "{}" { // empty struct
				obj.ImportNumberField = nil
			} else {
				match++
			}
		} else {
			obj.ImportNumberField = nil
		}
	} else {
		obj.ImportNumberField = nil
	}

	// try to unmarshal data into ImportBooleanField
	err = common.Unmarshal(data, &obj.ImportBooleanField)
	if err == nil {
		if obj.ImportBooleanField != nil && obj.ImportBooleanField.UnparsedObject == nil {
			jsonImportBooleanField, _ := common.Marshal(obj.ImportBooleanField)
			if string(jsonImportBooleanField) == "{}" && string(data) != "{}" { // empty struct
				obj.ImportBooleanField = nil
			} else {
				match++
			}
		} else {
			obj.ImportBooleanField = nil
		}
	} else {
		obj.ImportBooleanField = nil
	}

	// try to unmarshal data into ImportEnumField
	err = common.Unmarshal(data, &obj.ImportEnumField)
	if err == nil {
		if obj.ImportEnumField != nil && obj.ImportEnumField.UnparsedObject == nil {
			jsonImportEnumField, _ := common.Marshal(obj.ImportEnumField)
			if string(jsonImportEnumField) == "{}" && string(data) != "{}" { // empty struct
				obj.ImportEnumField = nil
			} else {
				match++
			}
		} else {
			obj.ImportEnumField = nil
		}
	} else {
		obj.ImportEnumField = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ImportStringField = nil
		obj.ImportIntegerField = nil
		obj.ImportNumberField = nil
		obj.ImportBooleanField = nil
		obj.ImportEnumField = nil
		return common.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ImportConnectionField) MarshalJSON() ([]byte, error) {
	if obj.ImportStringField != nil {
		return common.Marshal(&obj.ImportStringField)
	}

	if obj.ImportIntegerField != nil {
		return common.Marshal(&obj.ImportIntegerField)
	}

	if obj.ImportNumberField != nil {
		return common.Marshal(&obj.ImportNumberField)
	}

	if obj.ImportBooleanField != nil {
		return common.Marshal(&obj.ImportBooleanField)
	}

	if obj.ImportEnumField != nil {
		return common.Marshal(&obj.ImportEnumField)
	}

	if obj.UnparsedObject != nil {
		return common.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ImportConnectionField) GetActualInstance() interface{} {
	if obj.ImportStringField != nil {
		return obj.ImportStringField
	}

	if obj.ImportIntegerField != nil {
		return obj.ImportIntegerField
	}

	if obj.ImportNumberField != nil {
		return obj.ImportNumberField
	}

	if obj.ImportBooleanField != nil {
		return obj.ImportBooleanField
	}

	if obj.ImportEnumField != nil {
		return obj.ImportEnumField
	}

	// all schemas are nil
	return nil
}
