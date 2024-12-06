// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

type ParameterConfigSharedSpecParamsItem struct {
	// name of the parameter
	Name *string `json:"name,omitempty"`
	// expression of the parameter value
	Expression *string `json:"expression,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterConfigSharedSpecParamsItem instantiates a new ParameterConfigSharedSpecParamsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterConfigSharedSpecParamsItem() *ParameterConfigSharedSpecParamsItem {
	this := ParameterConfigSharedSpecParamsItem{}
	return &this
}

// NewParameterConfigSharedSpecParamsItemWithDefaults instantiates a new ParameterConfigSharedSpecParamsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterConfigSharedSpecParamsItemWithDefaults() *ParameterConfigSharedSpecParamsItem {
	this := ParameterConfigSharedSpecParamsItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ParameterConfigSharedSpecParamsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterConfigSharedSpecParamsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ParameterConfigSharedSpecParamsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ParameterConfigSharedSpecParamsItem) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *ParameterConfigSharedSpecParamsItem) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterConfigSharedSpecParamsItem) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *ParameterConfigSharedSpecParamsItem) HasExpression() bool {
	return o != nil && o.Expression != nil
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *ParameterConfigSharedSpecParamsItem) SetExpression(v string) {
	o.Expression = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterConfigSharedSpecParamsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
func (o *ParameterConfigSharedSpecParamsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string `json:"name,omitempty"`
		Expression *string `json:"expression,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "expression"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Expression = all.Expression

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
