// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModuleUpdate Update information for an environment module
type EnvironmentModuleUpdate struct {
	// Name of the environment module to update
	Name string `json:"name"`
	// Action to perform on the environment module:
	// - Enable: Enable the module
	// - Disable: Disable the module
	// - Upgrade: Upgrade the module to specified version
	//
	Action EnvironmentModuleAction `json:"action"`
	// Version of the environment module to upgrade to
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModuleUpdate instantiates a new EnvironmentModuleUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModuleUpdate(name string, action EnvironmentModuleAction) *EnvironmentModuleUpdate {
	this := EnvironmentModuleUpdate{}
	this.Name = name
	this.Action = action
	return &this
}

// NewEnvironmentModuleUpdateWithDefaults instantiates a new EnvironmentModuleUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModuleUpdateWithDefaults() *EnvironmentModuleUpdate {
	this := EnvironmentModuleUpdate{}
	return &this
}

// GetName returns the Name field value.
func (o *EnvironmentModuleUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EnvironmentModuleUpdate) SetName(v string) {
	o.Name = v
}

// GetAction returns the Action field value.
func (o *EnvironmentModuleUpdate) GetAction() EnvironmentModuleAction {
	if o == nil {
		var ret EnvironmentModuleAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleUpdate) GetActionOk() (*EnvironmentModuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *EnvironmentModuleUpdate) SetAction(v EnvironmentModuleAction) {
	o.Action = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EnvironmentModuleUpdate) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleUpdate) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EnvironmentModuleUpdate) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EnvironmentModuleUpdate) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModuleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["action"] = o.Action
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentModuleUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string                  `json:"name"`
		Action  *EnvironmentModuleAction `json:"action"`
		Version *string                  `json:"version,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "action", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
