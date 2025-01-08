// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsRestart OpsRestart is the payload to restart a KubeBlocks cluster
type OpsRestart struct {
	// component type
	Component string `json:"component"`
	// Specifies the maximum time in seconds that the OpsRequest will wait for its pre-conditions to be met before it aborts the operation
	PreConditionDeadlineSeconds common.NullableInt32 `json:"preConditionDeadlineSeconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsRestart instantiates a new OpsRestart object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsRestart(component string) *OpsRestart {
	this := OpsRestart{}
	this.Component = component
	return &this
}

// NewOpsRestartWithDefaults instantiates a new OpsRestart object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsRestartWithDefaults() *OpsRestart {
	this := OpsRestart{}
	return &this
}

// GetComponent returns the Component field value.
func (o *OpsRestart) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OpsRestart) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *OpsRestart) SetComponent(v string) {
	o.Component = v
}

// GetPreConditionDeadlineSeconds returns the PreConditionDeadlineSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpsRestart) GetPreConditionDeadlineSeconds() int32 {
	if o == nil || o.PreConditionDeadlineSeconds.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PreConditionDeadlineSeconds.Get()
}

// GetPreConditionDeadlineSecondsOk returns a tuple with the PreConditionDeadlineSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OpsRestart) GetPreConditionDeadlineSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreConditionDeadlineSeconds.Get(), o.PreConditionDeadlineSeconds.IsSet()
}

// HasPreConditionDeadlineSeconds returns a boolean if a field has been set.
func (o *OpsRestart) HasPreConditionDeadlineSeconds() bool {
	return o != nil && o.PreConditionDeadlineSeconds.IsSet()
}

// SetPreConditionDeadlineSeconds gets a reference to the given common.NullableInt32 and assigns it to the PreConditionDeadlineSeconds field.
func (o *OpsRestart) SetPreConditionDeadlineSeconds(v int32) {
	o.PreConditionDeadlineSeconds.Set(&v)
}

// SetPreConditionDeadlineSecondsNil sets the value for PreConditionDeadlineSeconds to be an explicit nil.
func (o *OpsRestart) SetPreConditionDeadlineSecondsNil() {
	o.PreConditionDeadlineSeconds.Set(nil)
}

// UnsetPreConditionDeadlineSeconds ensures that no value is present for PreConditionDeadlineSeconds, not even an explicit nil.
func (o *OpsRestart) UnsetPreConditionDeadlineSeconds() {
	o.PreConditionDeadlineSeconds.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsRestart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	if o.PreConditionDeadlineSeconds.IsSet() {
		toSerialize["preConditionDeadlineSeconds"] = o.PreConditionDeadlineSeconds.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsRestart) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component                   *string              `json:"component"`
		PreConditionDeadlineSeconds common.NullableInt32 `json:"preConditionDeadlineSeconds,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "preConditionDeadlineSeconds"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.PreConditionDeadlineSeconds = all.PreConditionDeadlineSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
