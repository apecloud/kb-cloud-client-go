// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParameterOptionRef Reference another engine's parameter option to reuse its parameter configuration.
// When set on a parameterOption, the referenced engine's parameter option (identified by
// engineName + component) is used as the base configuration. Fields explicitly set in
// the local parameterOption will override those from the referenced one.
type ParameterOptionRef struct {
	// The engine name to reference parameter configuration from.
	EngineName string `json:"engineName"`
	// The component type in the referenced engine whose parameter option to inherit.
	Component string `json:"component"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterOptionRef instantiates a new ParameterOptionRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterOptionRef(engineName string, component string) *ParameterOptionRef {
	this := ParameterOptionRef{}
	this.EngineName = engineName
	this.Component = component
	return &this
}

// NewParameterOptionRefWithDefaults instantiates a new ParameterOptionRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterOptionRefWithDefaults() *ParameterOptionRef {
	this := ParameterOptionRef{}
	return &this
}

// GetEngineName returns the EngineName field value.
func (o *ParameterOptionRef) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *ParameterOptionRef) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *ParameterOptionRef) SetEngineName(v string) {
	o.EngineName = v
}

// GetComponent returns the Component field value.
func (o *ParameterOptionRef) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ParameterOptionRef) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ParameterOptionRef) SetComponent(v string) {
	o.Component = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterOptionRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["component"] = o.Component

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterOptionRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName *string `json:"engineName"`
		Component  *string `json:"component"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineName", "component"})
	} else {
		return err
	}
	o.EngineName = *all.EngineName
	o.Component = *all.Component

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
