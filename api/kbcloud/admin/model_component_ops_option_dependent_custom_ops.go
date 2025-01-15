// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ComponentOpsOptionDependentCustomOps struct {
	// opsDefinition name
	OpsDefName *string `json:"opsDefName,omitempty"`
	// component type name
	Component *string `json:"component,omitempty"`
	// go template conditional judgment expression, such as $.root.inPlace == true
	// available built-in objects that can be referenced in the expression include:
	// - cluster: cluster record
	// - params: ops input parameters
	// - component: current component
	// - root: current object
	//
	When *string `json:"when,omitempty"`
	// custom ops parameters
	Params []ComponentOpsOptionDependentCustomOpsParamsItem `json:"params,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOpsOptionDependentCustomOps instantiates a new ComponentOpsOptionDependentCustomOps object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOpsOptionDependentCustomOps() *ComponentOpsOptionDependentCustomOps {
	this := ComponentOpsOptionDependentCustomOps{}
	return &this
}

// NewComponentOpsOptionDependentCustomOpsWithDefaults instantiates a new ComponentOpsOptionDependentCustomOps object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOpsOptionDependentCustomOpsWithDefaults() *ComponentOpsOptionDependentCustomOps {
	this := ComponentOpsOptionDependentCustomOps{}
	return &this
}

// GetOpsDefName returns the OpsDefName field value if set, zero value otherwise.
func (o *ComponentOpsOptionDependentCustomOps) GetOpsDefName() string {
	if o == nil || o.OpsDefName == nil {
		var ret string
		return ret
	}
	return *o.OpsDefName
}

// GetOpsDefNameOk returns a tuple with the OpsDefName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentCustomOps) GetOpsDefNameOk() (*string, bool) {
	if o == nil || o.OpsDefName == nil {
		return nil, false
	}
	return o.OpsDefName, true
}

// HasOpsDefName returns a boolean if a field has been set.
func (o *ComponentOpsOptionDependentCustomOps) HasOpsDefName() bool {
	return o != nil && o.OpsDefName != nil
}

// SetOpsDefName gets a reference to the given string and assigns it to the OpsDefName field.
func (o *ComponentOpsOptionDependentCustomOps) SetOpsDefName(v string) {
	o.OpsDefName = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ComponentOpsOptionDependentCustomOps) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentCustomOps) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ComponentOpsOptionDependentCustomOps) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ComponentOpsOptionDependentCustomOps) SetComponent(v string) {
	o.Component = &v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *ComponentOpsOptionDependentCustomOps) GetWhen() string {
	if o == nil || o.When == nil {
		var ret string
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentCustomOps) GetWhenOk() (*string, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *ComponentOpsOptionDependentCustomOps) HasWhen() bool {
	return o != nil && o.When != nil
}

// SetWhen gets a reference to the given string and assigns it to the When field.
func (o *ComponentOpsOptionDependentCustomOps) SetWhen(v string) {
	o.When = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ComponentOpsOptionDependentCustomOps) GetParams() []ComponentOpsOptionDependentCustomOpsParamsItem {
	if o == nil || o.Params == nil {
		var ret []ComponentOpsOptionDependentCustomOpsParamsItem
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentCustomOps) GetParamsOk() (*[]ComponentOpsOptionDependentCustomOpsParamsItem, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return &o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ComponentOpsOptionDependentCustomOps) HasParams() bool {
	return o != nil && o.Params != nil
}

// SetParams gets a reference to the given []ComponentOpsOptionDependentCustomOpsParamsItem and assigns it to the Params field.
func (o *ComponentOpsOptionDependentCustomOps) SetParams(v []ComponentOpsOptionDependentCustomOpsParamsItem) {
	o.Params = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOpsOptionDependentCustomOps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OpsDefName != nil {
		toSerialize["opsDefName"] = o.OpsDefName
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOpsOptionDependentCustomOps) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OpsDefName *string                                          `json:"opsDefName,omitempty"`
		Component  *string                                          `json:"component,omitempty"`
		When       *string                                          `json:"when,omitempty"`
		Params     []ComponentOpsOptionDependentCustomOpsParamsItem `json:"params,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"opsDefName", "component", "when", "params"})
	} else {
		return err
	}
	o.OpsDefName = all.OpsDefName
	o.Component = all.Component
	o.When = all.When
	o.Params = all.Params

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
