// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SelectiveMethodOption struct {
	Name                 string                `json:"name"`
	Description          *LocalizedDescription `json:"description,omitempty"`
	RestoreTargetRole    *string               `json:"restoreTargetRole,omitempty"`
	SelectedObjectOption *SelectedObjectOption `json:"selectedObjectOption,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSelectiveMethodOption instantiates a new SelectiveMethodOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSelectiveMethodOption(name string) *SelectiveMethodOption {
	this := SelectiveMethodOption{}
	this.Name = name
	return &this
}

// NewSelectiveMethodOptionWithDefaults instantiates a new SelectiveMethodOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSelectiveMethodOptionWithDefaults() *SelectiveMethodOption {
	this := SelectiveMethodOption{}
	return &this
}

// GetName returns the Name field value.
func (o *SelectiveMethodOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SelectiveMethodOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SelectiveMethodOption) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SelectiveMethodOption) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectiveMethodOption) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SelectiveMethodOption) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *SelectiveMethodOption) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetRestoreTargetRole returns the RestoreTargetRole field value if set, zero value otherwise.
func (o *SelectiveMethodOption) GetRestoreTargetRole() string {
	if o == nil || o.RestoreTargetRole == nil {
		var ret string
		return ret
	}
	return *o.RestoreTargetRole
}

// GetRestoreTargetRoleOk returns a tuple with the RestoreTargetRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectiveMethodOption) GetRestoreTargetRoleOk() (*string, bool) {
	if o == nil || o.RestoreTargetRole == nil {
		return nil, false
	}
	return o.RestoreTargetRole, true
}

// HasRestoreTargetRole returns a boolean if a field has been set.
func (o *SelectiveMethodOption) HasRestoreTargetRole() bool {
	return o != nil && o.RestoreTargetRole != nil
}

// SetRestoreTargetRole gets a reference to the given string and assigns it to the RestoreTargetRole field.
func (o *SelectiveMethodOption) SetRestoreTargetRole(v string) {
	o.RestoreTargetRole = &v
}

// GetSelectedObjectOption returns the SelectedObjectOption field value if set, zero value otherwise.
func (o *SelectiveMethodOption) GetSelectedObjectOption() SelectedObjectOption {
	if o == nil || o.SelectedObjectOption == nil {
		var ret SelectedObjectOption
		return ret
	}
	return *o.SelectedObjectOption
}

// GetSelectedObjectOptionOk returns a tuple with the SelectedObjectOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectiveMethodOption) GetSelectedObjectOptionOk() (*SelectedObjectOption, bool) {
	if o == nil || o.SelectedObjectOption == nil {
		return nil, false
	}
	return o.SelectedObjectOption, true
}

// HasSelectedObjectOption returns a boolean if a field has been set.
func (o *SelectiveMethodOption) HasSelectedObjectOption() bool {
	return o != nil && o.SelectedObjectOption != nil
}

// SetSelectedObjectOption gets a reference to the given SelectedObjectOption and assigns it to the SelectedObjectOption field.
func (o *SelectiveMethodOption) SetSelectedObjectOption(v SelectedObjectOption) {
	o.SelectedObjectOption = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SelectiveMethodOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.RestoreTargetRole != nil {
		toSerialize["restoreTargetRole"] = o.RestoreTargetRole
	}
	if o.SelectedObjectOption != nil {
		toSerialize["selectedObjectOption"] = o.SelectedObjectOption
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SelectiveMethodOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                 *string               `json:"name"`
		Description          *LocalizedDescription `json:"description,omitempty"`
		RestoreTargetRole    *string               `json:"restoreTargetRole,omitempty"`
		SelectedObjectOption *SelectedObjectOption `json:"selectedObjectOption,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "restoreTargetRole", "selectedObjectOption"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	o.RestoreTargetRole = all.RestoreTargetRole
	if all.SelectedObjectOption != nil && all.SelectedObjectOption.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SelectedObjectOption = all.SelectedObjectOption

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
