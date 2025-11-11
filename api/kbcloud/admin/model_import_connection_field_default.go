// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportConnectionFieldDefault - Default value applied when the field is omitted. Coerced according to `type`.
type ImportConnectionFieldDefault struct {
	String  *string
	Int32   *int32
	Float64 *float64
	Bool    *bool

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// StringAsImportConnectionFieldDefault is a convenience function that returns string wrapped in ImportConnectionFieldDefault.
func StringAsImportConnectionFieldDefault(v *string) ImportConnectionFieldDefault {
	return ImportConnectionFieldDefault{String: v}
}

// Int32AsImportConnectionFieldDefault is a convenience function that returns int32 wrapped in ImportConnectionFieldDefault.
func Int32AsImportConnectionFieldDefault(v *int32) ImportConnectionFieldDefault {
	return ImportConnectionFieldDefault{Int32: v}
}

// Float64AsImportConnectionFieldDefault is a convenience function that returns float64 wrapped in ImportConnectionFieldDefault.
func Float64AsImportConnectionFieldDefault(v *float64) ImportConnectionFieldDefault {
	return ImportConnectionFieldDefault{Float64: v}
}

// BoolAsImportConnectionFieldDefault is a convenience function that returns bool wrapped in ImportConnectionFieldDefault.
func BoolAsImportConnectionFieldDefault(v *bool) ImportConnectionFieldDefault {
	return ImportConnectionFieldDefault{Bool: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ImportConnectionFieldDefault) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = common.Unmarshal(data, &obj.String)
	if err == nil {
		if obj.String != nil {
			jsonString, _ := common.Marshal(obj.String)
			if string(jsonString) == "{}" { // empty struct
				obj.String = nil
			} else {
				match++
			}
		} else {
			obj.String = nil
		}
	} else {
		obj.String = nil
	}

	// try to unmarshal data into Int32
	err = common.Unmarshal(data, &obj.Int32)
	if err == nil {
		if obj.Int32 != nil {
			jsonInt32, _ := common.Marshal(obj.Int32)
			if string(jsonInt32) == "{}" { // empty struct
				obj.Int32 = nil
			} else {
				match++
			}
		} else {
			obj.Int32 = nil
		}
	} else {
		obj.Int32 = nil
	}

	// try to unmarshal data into Float64
	err = common.Unmarshal(data, &obj.Float64)
	if err == nil {
		if obj.Float64 != nil {
			jsonFloat64, _ := common.Marshal(obj.Float64)
			if string(jsonFloat64) == "{}" { // empty struct
				obj.Float64 = nil
			} else {
				match++
			}
		} else {
			obj.Float64 = nil
		}
	} else {
		obj.Float64 = nil
	}

	// try to unmarshal data into Bool
	err = common.Unmarshal(data, &obj.Bool)
	if err == nil {
		if obj.Bool != nil {
			jsonBool, _ := common.Marshal(obj.Bool)
			if string(jsonBool) == "{}" { // empty struct
				obj.Bool = nil
			} else {
				match++
			}
		} else {
			obj.Bool = nil
		}
	} else {
		obj.Bool = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.String = nil
		obj.Int32 = nil
		obj.Float64 = nil
		obj.Bool = nil
		return common.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ImportConnectionFieldDefault) MarshalJSON() ([]byte, error) {
	if obj.String != nil {
		return common.Marshal(&obj.String)
	}

	if obj.Int32 != nil {
		return common.Marshal(&obj.Int32)
	}

	if obj.Float64 != nil {
		return common.Marshal(&obj.Float64)
	}

	if obj.Bool != nil {
		return common.Marshal(&obj.Bool)
	}

	if obj.UnparsedObject != nil {
		return common.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ImportConnectionFieldDefault) GetActualInstance() interface{} {
	if obj.String != nil {
		return obj.String
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.Float64 != nil {
		return obj.Float64
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	// all schemas are nil
	return nil
}

// NullableImportConnectionFieldDefault handles when a null is used for ImportConnectionFieldDefault.
type NullableImportConnectionFieldDefault struct {
	value *ImportConnectionFieldDefault
	isSet bool
}

// Get returns the associated value.
func (v NullableImportConnectionFieldDefault) Get() *ImportConnectionFieldDefault {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableImportConnectionFieldDefault) Set(val *ImportConnectionFieldDefault) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableImportConnectionFieldDefault) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableImportConnectionFieldDefault) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableImportConnectionFieldDefault initializes the struct as if Set has been called.
func NewNullableImportConnectionFieldDefault(val *ImportConnectionFieldDefault) *NullableImportConnectionFieldDefault {
	return &NullableImportConnectionFieldDefault{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableImportConnectionFieldDefault) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableImportConnectionFieldDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return common.Unmarshal(src, &v.value)
}
