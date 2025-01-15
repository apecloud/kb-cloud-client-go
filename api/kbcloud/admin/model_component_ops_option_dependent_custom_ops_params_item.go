// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ComponentOpsOptionDependentCustomOpsParamsItem struct {
	// parameter name.
	Name *string `json:"name,omitempty"`
	// parameter value, you can define a go template expression to refer the variable of the current ops.
	// available built-in objects that can be referenced in the expression include:
	// - cluster: cluster record
	// - params: ops input parameters
	// - component: current component
	// - root: current object
	//
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOpsOptionDependentCustomOpsParamsItem instantiates a new ComponentOpsOptionDependentCustomOpsParamsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOpsOptionDependentCustomOpsParamsItem() *ComponentOpsOptionDependentCustomOpsParamsItem {
	this := ComponentOpsOptionDependentCustomOpsParamsItem{}
	return &this
}

// NewComponentOpsOptionDependentCustomOpsParamsItemWithDefaults instantiates a new ComponentOpsOptionDependentCustomOpsParamsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOpsOptionDependentCustomOpsParamsItemWithDefaults() *ComponentOpsOptionDependentCustomOpsParamsItem {
	this := ComponentOpsOptionDependentCustomOpsParamsItem{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComponentOpsOptionDependentCustomOpsParamsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentCustomOpsParamsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComponentOpsOptionDependentCustomOpsParamsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComponentOpsOptionDependentCustomOpsParamsItem) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ComponentOpsOptionDependentCustomOpsParamsItem) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentCustomOpsParamsItem) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ComponentOpsOptionDependentCustomOpsParamsItem) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ComponentOpsOptionDependentCustomOpsParamsItem) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOpsOptionDependentCustomOpsParamsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOpsOptionDependentCustomOpsParamsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string `json:"name,omitempty"`
		Value *string `json:"value,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "value"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
