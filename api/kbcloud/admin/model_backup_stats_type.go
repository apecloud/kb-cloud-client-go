// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// BackupStatsType Totalsize and number of backups for the backup type
type BackupStatsType struct {
	// backup type
	BackupType *string `json:"backupType,omitempty"`
	// totalsize of backups for each engine, A string with capacity units in the form of "1Gi", "1Mi", "1Ki".
	BackupSize *string `json:"backupSize,omitempty"`
	// The number of backups for each engine
	BackupNum *int64 `json:"backupNum,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupStatsType instantiates a new BackupStatsType object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupStatsType() *BackupStatsType {
	this := BackupStatsType{}
	return &this
}

// NewBackupStatsTypeWithDefaults instantiates a new BackupStatsType object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupStatsTypeWithDefaults() *BackupStatsType {
	this := BackupStatsType{}
	return &this
}

// GetBackupType returns the BackupType field value if set, zero value otherwise.
func (o *BackupStatsType) GetBackupType() string {
	if o == nil || o.BackupType == nil {
		var ret string
		return ret
	}
	return *o.BackupType
}

// GetBackupTypeOk returns a tuple with the BackupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsType) GetBackupTypeOk() (*string, bool) {
	if o == nil || o.BackupType == nil {
		return nil, false
	}
	return o.BackupType, true
}

// HasBackupType returns a boolean if a field has been set.
func (o *BackupStatsType) HasBackupType() bool {
	return o != nil && o.BackupType != nil
}

// SetBackupType gets a reference to the given string and assigns it to the BackupType field.
func (o *BackupStatsType) SetBackupType(v string) {
	o.BackupType = &v
}

// GetBackupSize returns the BackupSize field value if set, zero value otherwise.
func (o *BackupStatsType) GetBackupSize() string {
	if o == nil || o.BackupSize == nil {
		var ret string
		return ret
	}
	return *o.BackupSize
}

// GetBackupSizeOk returns a tuple with the BackupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsType) GetBackupSizeOk() (*string, bool) {
	if o == nil || o.BackupSize == nil {
		return nil, false
	}
	return o.BackupSize, true
}

// HasBackupSize returns a boolean if a field has been set.
func (o *BackupStatsType) HasBackupSize() bool {
	return o != nil && o.BackupSize != nil
}

// SetBackupSize gets a reference to the given string and assigns it to the BackupSize field.
func (o *BackupStatsType) SetBackupSize(v string) {
	o.BackupSize = &v
}

// GetBackupNum returns the BackupNum field value if set, zero value otherwise.
func (o *BackupStatsType) GetBackupNum() int64 {
	if o == nil || o.BackupNum == nil {
		var ret int64
		return ret
	}
	return *o.BackupNum
}

// GetBackupNumOk returns a tuple with the BackupNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStatsType) GetBackupNumOk() (*int64, bool) {
	if o == nil || o.BackupNum == nil {
		return nil, false
	}
	return o.BackupNum, true
}

// HasBackupNum returns a boolean if a field has been set.
func (o *BackupStatsType) HasBackupNum() bool {
	return o != nil && o.BackupNum != nil
}

// SetBackupNum gets a reference to the given int64 and assigns it to the BackupNum field.
func (o *BackupStatsType) SetBackupNum(v int64) {
	o.BackupNum = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupStatsType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.BackupType != nil {
		toSerialize["backupType"] = o.BackupType
	}
	if o.BackupSize != nil {
		toSerialize["backupSize"] = o.BackupSize
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
func (o *BackupStatsType) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BackupType *string `json:"backupType,omitempty"`
		BackupSize *string `json:"backupSize,omitempty"`
		BackupNum  *int64  `json:"backupNum,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"backupType", "backupSize", "backupNum"})
	} else {
		return err
	}
	o.BackupType = all.BackupType
	o.BackupSize = all.BackupSize
	o.BackupNum = all.BackupNum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
