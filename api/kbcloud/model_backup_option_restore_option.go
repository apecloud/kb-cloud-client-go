// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type BackupOptionRestoreOption struct {
	// cross mode recovery options
	CrossModeRecovery []string `json:"crossModeRecovery,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupOptionRestoreOption instantiates a new BackupOptionRestoreOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupOptionRestoreOption() *BackupOptionRestoreOption {
	this := BackupOptionRestoreOption{}
	return &this
}

// NewBackupOptionRestoreOptionWithDefaults instantiates a new BackupOptionRestoreOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupOptionRestoreOptionWithDefaults() *BackupOptionRestoreOption {
	this := BackupOptionRestoreOption{}
	return &this
}

// GetCrossModeRecovery returns the CrossModeRecovery field value if set, zero value otherwise.
func (o *BackupOptionRestoreOption) GetCrossModeRecovery() []string {
	if o == nil || o.CrossModeRecovery == nil {
		var ret []string
		return ret
	}
	return o.CrossModeRecovery
}

// GetCrossModeRecoveryOk returns a tuple with the CrossModeRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupOptionRestoreOption) GetCrossModeRecoveryOk() (*[]string, bool) {
	if o == nil || o.CrossModeRecovery == nil {
		return nil, false
	}
	return &o.CrossModeRecovery, true
}

// HasCrossModeRecovery returns a boolean if a field has been set.
func (o *BackupOptionRestoreOption) HasCrossModeRecovery() bool {
	return o != nil && o.CrossModeRecovery != nil
}

// SetCrossModeRecovery gets a reference to the given []string and assigns it to the CrossModeRecovery field.
func (o *BackupOptionRestoreOption) SetCrossModeRecovery(v []string) {
	o.CrossModeRecovery = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupOptionRestoreOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CrossModeRecovery != nil {
		toSerialize["crossModeRecovery"] = o.CrossModeRecovery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupOptionRestoreOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CrossModeRecovery []string `json:"crossModeRecovery,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"crossModeRecovery"})
	} else {
		return err
	}
	o.CrossModeRecovery = all.CrossModeRecovery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
