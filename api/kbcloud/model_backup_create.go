// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// BackupCreate BackupCreate is the payload to create a KubeBlocks cluster backup
type BackupCreate struct {
	// name of the backup, if not specified, a name will be generated automatically
	Name *string `json:"name,omitempty"`
	// specified the backup method
	BackupMethod string `json:"backupMethod"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupCreate instantiates a new BackupCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupCreate(backupMethod string) *BackupCreate {
	this := BackupCreate{}
	this.BackupMethod = backupMethod
	return &this
}

// NewBackupCreateWithDefaults instantiates a new BackupCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupCreateWithDefaults() *BackupCreate {
	this := BackupCreate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BackupCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BackupCreate) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BackupCreate) SetName(v string) {
	o.Name = &v
}

// GetBackupMethod returns the BackupMethod field value.
func (o *BackupCreate) GetBackupMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BackupMethod
}

// GetBackupMethodOk returns a tuple with the BackupMethod field value
// and a boolean to check if the value has been set.
func (o *BackupCreate) GetBackupMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupMethod, true
}

// SetBackupMethod sets field value.
func (o *BackupCreate) SetBackupMethod(v string) {
	o.BackupMethod = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["backupMethod"] = o.BackupMethod

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string `json:"name,omitempty"`
		BackupMethod *string `json:"backupMethod"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BackupMethod == nil {
		return fmt.Errorf("required field backupMethod missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "backupMethod"})
	} else {
		return err
	}
	o.Name = all.Name
	o.BackupMethod = *all.BackupMethod

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
