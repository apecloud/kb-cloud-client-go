// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// DmsTableColumnGenerated Generated column information
type DmsTableColumnGenerated struct {
	// The type of generation (e.g., VIRTUAL or STORED)
	Type *string `json:"type,omitempty"`
	// The expression used to generate the column's value
	Expression *string `json:"expression,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsTableColumnGenerated instantiates a new DmsTableColumnGenerated object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsTableColumnGenerated() *DmsTableColumnGenerated {
	this := DmsTableColumnGenerated{}
	return &this
}

// NewDmsTableColumnGeneratedWithDefaults instantiates a new DmsTableColumnGenerated object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsTableColumnGeneratedWithDefaults() *DmsTableColumnGenerated {
	this := DmsTableColumnGenerated{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DmsTableColumnGenerated) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableColumnGenerated) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DmsTableColumnGenerated) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DmsTableColumnGenerated) SetType(v string) {
	o.Type = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *DmsTableColumnGenerated) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsTableColumnGenerated) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *DmsTableColumnGenerated) HasExpression() bool {
	return o != nil && o.Expression != nil
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *DmsTableColumnGenerated) SetExpression(v string) {
	o.Expression = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsTableColumnGenerated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsTableColumnGenerated) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type       *string `json:"type,omitempty"`
		Expression *string `json:"expression,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "expression"})
	} else {
		return err
	}
	o.Type = all.Type
	o.Expression = all.Expression

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableDmsTableColumnGenerated handles when a null is used for DmsTableColumnGenerated.
type NullableDmsTableColumnGenerated struct {
	value *DmsTableColumnGenerated
	isSet bool
}

// Get returns the associated value.
func (v NullableDmsTableColumnGenerated) Get() *DmsTableColumnGenerated {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDmsTableColumnGenerated) Set(val *DmsTableColumnGenerated) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDmsTableColumnGenerated) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDmsTableColumnGenerated) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDmsTableColumnGenerated initializes the struct as if Set has been called.
func NewNullableDmsTableColumnGenerated(val *DmsTableColumnGenerated) *NullableDmsTableColumnGenerated {
	return &NullableDmsTableColumnGenerated{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDmsTableColumnGenerated) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDmsTableColumnGenerated) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return common.Unmarshal(src, &v.value)
}
