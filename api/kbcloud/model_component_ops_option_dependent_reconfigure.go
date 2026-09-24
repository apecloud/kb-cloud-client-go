// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ComponentOpsOptionDependentReconfigure Reconfigure ops submitted after the current ops succeeds on KubeBlocks 1.0.
// Use this to update engine parameters that must follow a resource change.
// KubeBlocks 0.9 ignores this field.
// when and parameter values are Go templates with the same built-in objects as dependentCustomOps.
// Read the parent OpsRequest with a spec expression, for example
// {{ (index $.ops.spec.verticalScaling 0).limits.memory }}.
// Arithmetic comes from sprig. quantity parses that expression's Kubernetes quantity string into bytes.
type ComponentOpsOptionDependentReconfigure struct {
	// Target component type in the same cluster, resolved through engineOption.components.
	// Empty means the component type of the current ops. All matching components and shardings are reconfigured.
	// This allows an operation on component A to update parameters on component B.
	// An unknown type or a type with no matching cluster components fails before the parent ops is submitted.
	//
	Component *string `json:"component,omitempty"`
	// Go template expression. The reconfigure ops is submitted only when the expression evaluates to "true".
	// Same built-in objects as dependentCustomOps.
	//
	When *string `json:"when,omitempty"`
	// Parameter key-value pairs applied by the reconfigure ops.
	// value is a Go template over $.ops.spec. Example:
	// {{ mulf (quantity (index $.ops.spec.verticalScaling 0).limits.memory) 0.9 | int }}
	//
	Parameters []ComponentOpsOptionDependentReconfigureParametersItem `json:"parameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOpsOptionDependentReconfigure instantiates a new ComponentOpsOptionDependentReconfigure object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOpsOptionDependentReconfigure() *ComponentOpsOptionDependentReconfigure {
	this := ComponentOpsOptionDependentReconfigure{}
	return &this
}

// NewComponentOpsOptionDependentReconfigureWithDefaults instantiates a new ComponentOpsOptionDependentReconfigure object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOpsOptionDependentReconfigureWithDefaults() *ComponentOpsOptionDependentReconfigure {
	this := ComponentOpsOptionDependentReconfigure{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ComponentOpsOptionDependentReconfigure) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentReconfigure) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ComponentOpsOptionDependentReconfigure) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ComponentOpsOptionDependentReconfigure) SetComponent(v string) {
	o.Component = &v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *ComponentOpsOptionDependentReconfigure) GetWhen() string {
	if o == nil || o.When == nil {
		var ret string
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentReconfigure) GetWhenOk() (*string, bool) {
	if o == nil || o.When == nil {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *ComponentOpsOptionDependentReconfigure) HasWhen() bool {
	return o != nil && o.When != nil
}

// SetWhen gets a reference to the given string and assigns it to the When field.
func (o *ComponentOpsOptionDependentReconfigure) SetWhen(v string) {
	o.When = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ComponentOpsOptionDependentReconfigure) GetParameters() []ComponentOpsOptionDependentReconfigureParametersItem {
	if o == nil || o.Parameters == nil {
		var ret []ComponentOpsOptionDependentReconfigureParametersItem
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentReconfigure) GetParametersOk() (*[]ComponentOpsOptionDependentReconfigureParametersItem, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ComponentOpsOptionDependentReconfigure) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []ComponentOpsOptionDependentReconfigureParametersItem and assigns it to the Parameters field.
func (o *ComponentOpsOptionDependentReconfigure) SetParameters(v []ComponentOpsOptionDependentReconfigureParametersItem) {
	o.Parameters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOpsOptionDependentReconfigure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.When != nil {
		toSerialize["when"] = o.When
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOpsOptionDependentReconfigure) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component  *string                                                `json:"component,omitempty"`
		When       *string                                                `json:"when,omitempty"`
		Parameters []ComponentOpsOptionDependentReconfigureParametersItem `json:"parameters,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "when", "parameters"})
	} else {
		return err
	}
	o.Component = all.Component
	o.When = all.When
	o.Parameters = all.Parameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
