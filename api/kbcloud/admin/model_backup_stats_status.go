// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// BackupStatsStatus Number of backups for the status
// NODESCRIPTION BackupStatsStatus
//
// Deprecated: This model is deprecated.
type BackupStatsStatus struct {
	// Backup status
	Status *string `json:"status,omitempty"`
	// Number of backups for each status
	BackupNum *int64 `json:"backupNum,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupStatsStatus instantiates a new BackupStatsStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupStatsStatus() *BackupStatsStatus {
	this := BackupStatsStatus{}
	return &this
}

// NewBackupStatsStatusWithDefaults instantiates a new BackupStatsStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupStatsStatusWithDefaults() *BackupStatsStatus {
	this := BackupStatsStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BackupStatsStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BackupStatsStatus) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BackupStatsStatus) SetStatus(v string) {
	o.Status = &v
}

// GetBackupNum returns the BackupNum field value if set, zero value otherwise.
func (o *BackupStatsStatus) GetBackupNum() int64 {
	if o == nil || o.BackupNum == nil {
		var ret int64
		return ret
	}
	return *o.BackupNum
}

// GetBackupNumOk returns a tuple with the BackupNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsStatus) GetBackupNumOk() (*int64, bool) {
	if o == nil || o.BackupNum == nil {
		return nil, false
	}
	return o.BackupNum, true
}

// HasBackupNum returns a boolean if a field has been set.
func (o *BackupStatsStatus) HasBackupNum() bool {
	return o != nil && o.BackupNum != nil
}

// SetBackupNum gets a reference to the given int64 and assigns it to the BackupNum field.
func (o *BackupStatsStatus) SetBackupNum(v int64) {
	o.BackupNum = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupStatsStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.BackupNum != nil {
		toSerialize["backupNum"] = o.BackupNum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupStatsStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status    *string `json:"status,omitempty"`
		BackupNum *int64  `json:"backupNum,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "backupNum"})
	} else {
		return err
	}
	o.Status = all.Status
	o.BackupNum = all.BackupNum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
