// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineModeTransition struct {
	// Rule telling whether the env of an existing component is re-rendered when the mode transitions.
	EnvRerender *EnvRerender `json:"envRerender,omitempty"`
	// Target mode name.
	Mode string `json:"mode"`
	// Components that must be running once the mode change is done, in addition to the ones the
	// transition creates. Declare a component here when the transition re-renders it instead of
	// creating it, so it must settle before the transition is considered finished.
	//
	WaitComponentsRunningAfterModeChange []string `json:"waitComponentsRunningAfterModeChange,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineModeTransition instantiates a new EngineModeTransition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineModeTransition(mode string) *EngineModeTransition {
	this := EngineModeTransition{}
	this.Mode = mode
	return &this
}

// NewEngineModeTransitionWithDefaults instantiates a new EngineModeTransition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineModeTransitionWithDefaults() *EngineModeTransition {
	this := EngineModeTransition{}
	return &this
}

// GetEnvRerender returns the EnvRerender field value if set, zero value otherwise.
func (o *EngineModeTransition) GetEnvRerender() EnvRerender {
	if o == nil || o.EnvRerender == nil {
		var ret EnvRerender
		return ret
	}
	return *o.EnvRerender
}

// GetEnvRerenderOk returns a tuple with the EnvRerender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineModeTransition) GetEnvRerenderOk() (*EnvRerender, bool) {
	if o == nil || o.EnvRerender == nil {
		return nil, false
	}
	return o.EnvRerender, true
}

// HasEnvRerender returns a boolean if a field has been set.
func (o *EngineModeTransition) HasEnvRerender() bool {
	return o != nil && o.EnvRerender != nil
}

// SetEnvRerender gets a reference to the given EnvRerender and assigns it to the EnvRerender field.
func (o *EngineModeTransition) SetEnvRerender(v EnvRerender) {
	o.EnvRerender = &v
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

// GetWaitComponentsRunningAfterModeChange returns the WaitComponentsRunningAfterModeChange field value if set, zero value otherwise.
func (o *EngineModeTransition) GetWaitComponentsRunningAfterModeChange() []string {
	if o == nil || o.WaitComponentsRunningAfterModeChange == nil {
		var ret []string
		return ret
	}
	return o.WaitComponentsRunningAfterModeChange
}

// GetWaitComponentsRunningAfterModeChangeOk returns a tuple with the WaitComponentsRunningAfterModeChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineModeTransition) GetWaitComponentsRunningAfterModeChangeOk() (*[]string, bool) {
	if o == nil || o.WaitComponentsRunningAfterModeChange == nil {
		return nil, false
	}
	return &o.WaitComponentsRunningAfterModeChange, true
}

// HasWaitComponentsRunningAfterModeChange returns a boolean if a field has been set.
func (o *EngineModeTransition) HasWaitComponentsRunningAfterModeChange() bool {
	return o != nil && o.WaitComponentsRunningAfterModeChange != nil
}

// SetWaitComponentsRunningAfterModeChange gets a reference to the given []string and assigns it to the WaitComponentsRunningAfterModeChange field.
func (o *EngineModeTransition) SetWaitComponentsRunningAfterModeChange(v []string) {
	o.WaitComponentsRunningAfterModeChange = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineModeTransition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EnvRerender != nil {
		toSerialize["envRerender"] = o.EnvRerender
	}
	toSerialize["mode"] = o.Mode
	if o.WaitComponentsRunningAfterModeChange != nil {
		toSerialize["waitComponentsRunningAfterModeChange"] = o.WaitComponentsRunningAfterModeChange
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineModeTransition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvRerender                          *EnvRerender `json:"envRerender,omitempty"`
		Mode                                 *string      `json:"mode"`
		WaitComponentsRunningAfterModeChange []string     `json:"waitComponentsRunningAfterModeChange,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"envRerender", "mode", "waitComponentsRunningAfterModeChange"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.EnvRerender != nil && all.EnvRerender.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EnvRerender = all.EnvRerender
	o.Mode = *all.Mode
	o.WaitComponentsRunningAfterModeChange = all.WaitComponentsRunningAfterModeChange

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
