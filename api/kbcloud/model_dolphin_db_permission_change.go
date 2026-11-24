// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBPermissionChange struct {
	// grant, deny or revoke. Revoke clears the direct setting; it does not deny inherited access.
	Action string `json:"action"`
	// Native privilege name from capabilities.
	Privilege string   `json:"privilege"`
	Objects   []string `json:"objects,omitempty"`
	// Numeric quota value; only for grant of quota privileges.
	Value *float64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBPermissionChange instantiates a new DolphinDBPermissionChange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBPermissionChange(action string, privilege string) *DolphinDBPermissionChange {
	this := DolphinDBPermissionChange{}
	this.Action = action
	this.Privilege = privilege
	return &this
}

// NewDolphinDBPermissionChangeWithDefaults instantiates a new DolphinDBPermissionChange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBPermissionChangeWithDefaults() *DolphinDBPermissionChange {
	this := DolphinDBPermissionChange{}
	return &this
}

// GetAction returns the Action field value.
func (o *DolphinDBPermissionChange) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionChange) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *DolphinDBPermissionChange) SetAction(v string) {
	o.Action = v
}

// GetPrivilege returns the Privilege field value.
func (o *DolphinDBPermissionChange) GetPrivilege() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionChange) GetPrivilegeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privilege, true
}

// SetPrivilege sets field value.
func (o *DolphinDBPermissionChange) SetPrivilege(v string) {
	o.Privilege = v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *DolphinDBPermissionChange) GetObjects() []string {
	if o == nil || o.Objects == nil {
		var ret []string
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionChange) GetObjectsOk() (*[]string, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return &o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *DolphinDBPermissionChange) HasObjects() bool {
	return o != nil && o.Objects != nil
}

// SetObjects gets a reference to the given []string and assigns it to the Objects field.
func (o *DolphinDBPermissionChange) SetObjects(v []string) {
	o.Objects = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DolphinDBPermissionChange) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionChange) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DolphinDBPermissionChange) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *DolphinDBPermissionChange) SetValue(v float64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBPermissionChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["privilege"] = o.Privilege
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBPermissionChange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action    *string  `json:"action"`
		Privilege *string  `json:"privilege"`
		Objects   []string `json:"objects,omitempty"`
		Value     *float64 `json:"value,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Privilege == nil {
		return fmt.Errorf("required field privilege missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"action", "privilege", "objects", "value"})
	} else {
		return err
	}
	o.Action = *all.Action
	o.Privilege = *all.Privilege
	o.Objects = all.Objects
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
