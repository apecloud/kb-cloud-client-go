// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineModeComponentReplicas Target replicas for a component during mode change.
type EngineModeComponentReplicas struct {
	// Component name.
	Component string `json:"component"`
	// Target number of replicas.
	Replicas int64 `json:"replicas"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineModeComponentReplicas instantiates a new EngineModeComponentReplicas object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineModeComponentReplicas(component string, replicas int64) *EngineModeComponentReplicas {
	this := EngineModeComponentReplicas{}
	this.Component = component
	this.Replicas = replicas
	return &this
}

// NewEngineModeComponentReplicasWithDefaults instantiates a new EngineModeComponentReplicas object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineModeComponentReplicasWithDefaults() *EngineModeComponentReplicas {
	this := EngineModeComponentReplicas{}
	return &this
}

// GetComponent returns the Component field value.
func (o *EngineModeComponentReplicas) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *EngineModeComponentReplicas) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *EngineModeComponentReplicas) SetComponent(v string) {
	o.Component = v
}

// GetReplicas returns the Replicas field value.
func (o *EngineModeComponentReplicas) GetReplicas() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *EngineModeComponentReplicas) GetReplicasOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value.
func (o *EngineModeComponentReplicas) SetReplicas(v int64) {
	o.Replicas = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineModeComponentReplicas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["replicas"] = o.Replicas

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineModeComponentReplicas) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string `json:"component"`
		Replicas  *int64  `json:"replicas"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Replicas == nil {
		return fmt.Errorf("required field replicas missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "replicas"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.Replicas = *all.Replicas

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
