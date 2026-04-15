// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineModeTransition Parameters for switching the engine mode of a cluster.
type EngineModeTransition struct {
	// Target engine mode.
	Mode string `json:"mode"`
	// Target replicas for a component during mode change.
	Components EngineModeComponentReplicas `json:"components"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineModeTransition instantiates a new EngineModeTransition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineModeTransition(mode string, components EngineModeComponentReplicas) *EngineModeTransition {
	this := EngineModeTransition{}
	this.Mode = mode
	this.Components = components
	return &this
}

// NewEngineModeTransitionWithDefaults instantiates a new EngineModeTransition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineModeTransitionWithDefaults() *EngineModeTransition {
	this := EngineModeTransition{}
	return &this
}

// GetMode returns the Mode field value.
func (o *EngineModeTransition) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *EngineModeTransition) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *EngineModeTransition) SetMode(v string) {
	o.Mode = v
}

// GetComponents returns the Components field value.
func (o *EngineModeTransition) GetComponents() EngineModeComponentReplicas {
	if o == nil {
		var ret EngineModeComponentReplicas
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *EngineModeTransition) GetComponentsOk() (*EngineModeComponentReplicas, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value.
func (o *EngineModeTransition) SetComponents(v EngineModeComponentReplicas) {
	o.Components = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineModeTransition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["mode"] = o.Mode
	toSerialize["components"] = o.Components

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineModeTransition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Mode       *string                      `json:"mode"`
		Components *EngineModeComponentReplicas `json:"components"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Components == nil {
		return fmt.Errorf("required field components missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"mode", "components"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Mode = *all.Mode
	if all.Components.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Components = *all.Components

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
