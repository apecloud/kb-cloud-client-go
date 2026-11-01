// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBPermissionState struct {
	// Native snapshot field. ACCESS_READ/INSERT/UPDATE/DELETE represent global table/database access.
	Privilege string `json:"privilege"`
	// allow, deny or none; absent when not applicable.
	GlobalEffect   *string  `json:"globalEffect,omitempty"`
	AllowedObjects []string `json:"allowedObjects,omitempty"`
	DeniedObjects  []string `json:"deniedObjects,omitempty"`
	// Native quota or account limit, including default sentinel values.
	Value *float64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBPermissionState instantiates a new DolphinDBPermissionState object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBPermissionState(privilege string) *DolphinDBPermissionState {
	this := DolphinDBPermissionState{}
	this.Privilege = privilege
	return &this
}

// NewDolphinDBPermissionStateWithDefaults instantiates a new DolphinDBPermissionState object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBPermissionStateWithDefaults() *DolphinDBPermissionState {
	this := DolphinDBPermissionState{}
	return &this
}

// GetPrivilege returns the Privilege field value.
func (o *DolphinDBPermissionState) GetPrivilege() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Privilege
}

// GetPrivilegeOk returns a tuple with the Privilege field value
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionState) GetPrivilegeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privilege, true
}

// SetPrivilege sets field value.
func (o *DolphinDBPermissionState) SetPrivilege(v string) {
	o.Privilege = v
}

// GetGlobalEffect returns the GlobalEffect field value if set, zero value otherwise.
func (o *DolphinDBPermissionState) GetGlobalEffect() string {
	if o == nil || o.GlobalEffect == nil {
		var ret string
		return ret
	}
	return *o.GlobalEffect
}

// GetGlobalEffectOk returns a tuple with the GlobalEffect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionState) GetGlobalEffectOk() (*string, bool) {
	if o == nil || o.GlobalEffect == nil {
		return nil, false
	}
	return o.GlobalEffect, true
}

// HasGlobalEffect returns a boolean if a field has been set.
func (o *DolphinDBPermissionState) HasGlobalEffect() bool {
	return o != nil && o.GlobalEffect != nil
}

// SetGlobalEffect gets a reference to the given string and assigns it to the GlobalEffect field.
func (o *DolphinDBPermissionState) SetGlobalEffect(v string) {
	o.GlobalEffect = &v
}

// GetAllowedObjects returns the AllowedObjects field value if set, zero value otherwise.
func (o *DolphinDBPermissionState) GetAllowedObjects() []string {
	if o == nil || o.AllowedObjects == nil {
		var ret []string
		return ret
	}
	return o.AllowedObjects
}

// GetAllowedObjectsOk returns a tuple with the AllowedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionState) GetAllowedObjectsOk() (*[]string, bool) {
	if o == nil || o.AllowedObjects == nil {
		return nil, false
	}
	return &o.AllowedObjects, true
}

// HasAllowedObjects returns a boolean if a field has been set.
func (o *DolphinDBPermissionState) HasAllowedObjects() bool {
	return o != nil && o.AllowedObjects != nil
}

// SetAllowedObjects gets a reference to the given []string and assigns it to the AllowedObjects field.
func (o *DolphinDBPermissionState) SetAllowedObjects(v []string) {
	o.AllowedObjects = v
}

// GetDeniedObjects returns the DeniedObjects field value if set, zero value otherwise.
func (o *DolphinDBPermissionState) GetDeniedObjects() []string {
	if o == nil || o.DeniedObjects == nil {
		var ret []string
		return ret
	}
	return o.DeniedObjects
}

// GetDeniedObjectsOk returns a tuple with the DeniedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionState) GetDeniedObjectsOk() (*[]string, bool) {
	if o == nil || o.DeniedObjects == nil {
		return nil, false
	}
	return &o.DeniedObjects, true
}

// HasDeniedObjects returns a boolean if a field has been set.
func (o *DolphinDBPermissionState) HasDeniedObjects() bool {
	return o != nil && o.DeniedObjects != nil
}

// SetDeniedObjects gets a reference to the given []string and assigns it to the DeniedObjects field.
func (o *DolphinDBPermissionState) SetDeniedObjects(v []string) {
	o.DeniedObjects = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DolphinDBPermissionState) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBPermissionState) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DolphinDBPermissionState) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *DolphinDBPermissionState) SetValue(v float64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBPermissionState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["privilege"] = o.Privilege
	if o.GlobalEffect != nil {
		toSerialize["globalEffect"] = o.GlobalEffect
	}
	if o.AllowedObjects != nil {
		toSerialize["allowedObjects"] = o.AllowedObjects
	}
	if o.DeniedObjects != nil {
		toSerialize["deniedObjects"] = o.DeniedObjects
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
func (o *DolphinDBPermissionState) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Privilege      *string  `json:"privilege"`
		GlobalEffect   *string  `json:"globalEffect,omitempty"`
		AllowedObjects []string `json:"allowedObjects,omitempty"`
		DeniedObjects  []string `json:"deniedObjects,omitempty"`
		Value          *float64 `json:"value,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Privilege == nil {
		return fmt.Errorf("required field privilege missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"privilege", "globalEffect", "allowedObjects", "deniedObjects", "value"})
	} else {
		return err
	}
	o.Privilege = *all.Privilege
	o.GlobalEffect = all.GlobalEffect
	o.AllowedObjects = all.AllowedObjects
	o.DeniedObjects = all.DeniedObjects
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
