// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupBatchDelete BackupBatchDelete is the payload for deleting multiple backups
type BackupBatchDelete struct {
	// IDs of the backups to delete
	BackupIds []string `json:"backupIds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupBatchDelete instantiates a new BackupBatchDelete object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupBatchDelete(backupIds []string) *BackupBatchDelete {
	this := BackupBatchDelete{}
	this.BackupIds = backupIds
	return &this
}

// NewBackupBatchDeleteWithDefaults instantiates a new BackupBatchDelete object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupBatchDeleteWithDefaults() *BackupBatchDelete {
	this := BackupBatchDelete{}
	return &this
}

// GetBackupIds returns the BackupIds field value.
func (o *BackupBatchDelete) GetBackupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BackupIds
}

// GetBackupIdsOk returns a tuple with the BackupIds field value
// and a boolean to check if the value has been set.
func (o *BackupBatchDelete) GetBackupIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupIds, true
}

// SetBackupIds sets field value.
func (o *BackupBatchDelete) SetBackupIds(v []string) {
	o.BackupIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupBatchDelete) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["backupIds"] = o.BackupIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupBatchDelete) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BackupIds *[]string `json:"backupIds"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.BackupIds == nil {
		return fmt.Errorf("required field backupIds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"backupIds"})
	} else {
		return err
	}
	o.BackupIds = *all.BackupIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
