// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBPrivilegeCapability struct {
	Privilege string `json:"privilege"`
	// global, object, optionalObject or quota
	Scope    string   `json:"scope"`
	Actions  []string `json:"actions"`
	Subjects []string `json:"subjects"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBPrivilegeCapability instantiates a new DolphinDBPrivilegeCapability object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBPrivilegeCapability(privilege string, scope string, actions []string, subjects []string) *DolphinDBPrivilegeCapability {
	this := DolphinDBPrivilegeCapability{}
	this.Privilege = privilege
	this.Scope = scope
	this.Actions = actions
	this.Subjects = subjects
	return &this
}

// NewDolphinDBPrivilegeCapabilityWithDefaults instantiates a new DolphinDBPrivilegeCapability object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBPrivilegeCapabilityWithDefaults() *DolphinDBPrivilegeCapability {
	this := DolphinDBPrivilegeCapability{}
	return &this
}

// GetPrivilege returns the Privilege field value.
func (o *DolphinDBPrivilegeCapability) GetPrivilege() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPrivilegeCapability) GetPrivilegeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privilege, true
}

// SetPrivilege sets field value.
func (o *DolphinDBPrivilegeCapability) SetPrivilege(v string) {
	o.Privilege = v
}

// GetScope returns the Scope field value.
func (o *DolphinDBPrivilegeCapability) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPrivilegeCapability) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *DolphinDBPrivilegeCapability) SetScope(v string) {
	o.Scope = v
}

// GetActions returns the Actions field value.
func (o *DolphinDBPrivilegeCapability) GetActions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPrivilegeCapability) GetActionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actions, true
}

// SetActions sets field value.
func (o *DolphinDBPrivilegeCapability) SetActions(v []string) {
	o.Actions = v
}

// GetSubjects returns the Subjects field value.
func (o *DolphinDBPrivilegeCapability) GetSubjects() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPrivilegeCapability) GetSubjectsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subjects, true
}

// SetSubjects sets field value.
func (o *DolphinDBPrivilegeCapability) SetSubjects(v []string) {
	o.Subjects = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBPrivilegeCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["privilege"] = o.Privilege
	toSerialize["scope"] = o.Scope
	toSerialize["actions"] = o.Actions
	toSerialize["subjects"] = o.Subjects

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBPrivilegeCapability) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Privilege *string   `json:"privilege"`
		Scope     *string   `json:"scope"`
		Actions   *[]string `json:"actions"`
		Subjects  *[]string `json:"subjects"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Privilege == nil {
		return fmt.Errorf("required field privilege missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	if all.Actions == nil {
		return fmt.Errorf("required field actions missing")
	}
	if all.Subjects == nil {
		return fmt.Errorf("required field subjects missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"privilege", "scope", "actions", "subjects"})
	} else {
		return err
	}
	o.Privilege = *all.Privilege
	o.Scope = *all.Scope
	o.Actions = *all.Actions
	o.Subjects = *all.Subjects

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
