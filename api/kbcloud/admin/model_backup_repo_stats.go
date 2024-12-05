// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type BackupRepoStats struct {
	// the total backup in backup repo
	TotalBackup *int64 `json:"totalBackup,omitempty"`
	// the total backup size in backup repo
	TotalSize *string `json:"totalSize,omitempty"`
	// Number of backups for each defferent status
	BackupStatsStatus []BackupStatsStatus `json:"backupStatsStatus,omitempty"`
	// Totalsize and number of backups for each engine
	BackupStatsEngine []BackupStatsEngine `json:"backupStatsEngine,omitempty"`
	// Totalsize and number of backups for each engine
	BackupStatsType []BackupStatsType `json:"backupStatsType,omitempty"`
	// the time that the stats is created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupRepoStats instantiates a new BackupRepoStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupRepoStats() *BackupRepoStats {
	this := BackupRepoStats{}
	return &this
}

// NewBackupRepoStatsWithDefaults instantiates a new BackupRepoStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupRepoStatsWithDefaults() *BackupRepoStats {
	this := BackupRepoStats{}
	return &this
}

// GetTotalBackup returns the TotalBackup field value if set, zero value otherwise.
func (o *BackupRepoStats) GetTotalBackup() int64 {
	if o == nil || o.TotalBackup == nil {
		var ret int64
		return ret
	}
	return *o.TotalBackup
}

// GetTotalBackupOk returns a tuple with the TotalBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoStats) GetTotalBackupOk() (*int64, bool) {
	if o == nil || o.TotalBackup == nil {
		return nil, false
	}
	return o.TotalBackup, true
}

// HasTotalBackup returns a boolean if a field has been set.
func (o *BackupRepoStats) HasTotalBackup() bool {
	return o != nil && o.TotalBackup != nil
}

// SetTotalBackup gets a reference to the given int64 and assigns it to the TotalBackup field.
func (o *BackupRepoStats) SetTotalBackup(v int64) {
	o.TotalBackup = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *BackupRepoStats) GetTotalSize() string {
	if o == nil || o.TotalSize == nil {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoStats) GetTotalSizeOk() (*string, bool) {
	if o == nil || o.TotalSize == nil {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *BackupRepoStats) HasTotalSize() bool {
	return o != nil && o.TotalSize != nil
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *BackupRepoStats) SetTotalSize(v string) {
	o.TotalSize = &v
}

// GetBackupStatsStatus returns the BackupStatsStatus field value if set, zero value otherwise.
func (o *BackupRepoStats) GetBackupStatsStatus() []BackupStatsStatus {
	if o == nil || o.BackupStatsStatus == nil {
		var ret []BackupStatsStatus
		return ret
	}
	return o.BackupStatsStatus
}

// GetBackupStatsStatusOk returns a tuple with the BackupStatsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoStats) GetBackupStatsStatusOk() (*[]BackupStatsStatus, bool) {
	if o == nil || o.BackupStatsStatus == nil {
		return nil, false
	}
	return &o.BackupStatsStatus, true
}

// HasBackupStatsStatus returns a boolean if a field has been set.
func (o *BackupRepoStats) HasBackupStatsStatus() bool {
	return o != nil && o.BackupStatsStatus != nil
}

// SetBackupStatsStatus gets a reference to the given []BackupStatsStatus and assigns it to the BackupStatsStatus field.
func (o *BackupRepoStats) SetBackupStatsStatus(v []BackupStatsStatus) {
	o.BackupStatsStatus = v
}

// GetBackupStatsEngine returns the BackupStatsEngine field value if set, zero value otherwise.
func (o *BackupRepoStats) GetBackupStatsEngine() []BackupStatsEngine {
	if o == nil || o.BackupStatsEngine == nil {
		var ret []BackupStatsEngine
		return ret
	}
	return o.BackupStatsEngine
}

// GetBackupStatsEngineOk returns a tuple with the BackupStatsEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoStats) GetBackupStatsEngineOk() (*[]BackupStatsEngine, bool) {
	if o == nil || o.BackupStatsEngine == nil {
		return nil, false
	}
	return &o.BackupStatsEngine, true
}

// HasBackupStatsEngine returns a boolean if a field has been set.
func (o *BackupRepoStats) HasBackupStatsEngine() bool {
	return o != nil && o.BackupStatsEngine != nil
}

// SetBackupStatsEngine gets a reference to the given []BackupStatsEngine and assigns it to the BackupStatsEngine field.
func (o *BackupRepoStats) SetBackupStatsEngine(v []BackupStatsEngine) {
	o.BackupStatsEngine = v
}

// GetBackupStatsType returns the BackupStatsType field value if set, zero value otherwise.
func (o *BackupRepoStats) GetBackupStatsType() []BackupStatsType {
	if o == nil || o.BackupStatsType == nil {
		var ret []BackupStatsType
		return ret
	}
	return o.BackupStatsType
}

// GetBackupStatsTypeOk returns a tuple with the BackupStatsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoStats) GetBackupStatsTypeOk() (*[]BackupStatsType, bool) {
	if o == nil || o.BackupStatsType == nil {
		return nil, false
	}
	return &o.BackupStatsType, true
}

// HasBackupStatsType returns a boolean if a field has been set.
func (o *BackupRepoStats) HasBackupStatsType() bool {
	return o != nil && o.BackupStatsType != nil
}

// SetBackupStatsType gets a reference to the given []BackupStatsType and assigns it to the BackupStatsType field.
func (o *BackupRepoStats) SetBackupStatsType(v []BackupStatsType) {
	o.BackupStatsType = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BackupRepoStats) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoStats) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BackupRepoStats) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BackupRepoStats) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupRepoStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TotalBackup != nil {
		toSerialize["totalBackup"] = o.TotalBackup
	}
	if o.TotalSize != nil {
		toSerialize["totalSize"] = o.TotalSize
	}
	if o.BackupStatsStatus != nil {
		toSerialize["backupStatsStatus"] = o.BackupStatsStatus
	}
	if o.BackupStatsEngine != nil {
		toSerialize["backupStatsEngine"] = o.BackupStatsEngine
	}
	if o.BackupStatsType != nil {
		toSerialize["backupStatsType"] = o.BackupStatsType
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupRepoStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalBackup       *int64              `json:"totalBackup,omitempty"`
		TotalSize         *string             `json:"totalSize,omitempty"`
		BackupStatsStatus []BackupStatsStatus `json:"backupStatsStatus,omitempty"`
		BackupStatsEngine []BackupStatsEngine `json:"backupStatsEngine,omitempty"`
		BackupStatsType   []BackupStatsType   `json:"backupStatsType,omitempty"`
		CreatedAt         *time.Time          `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"totalBackup", "totalSize", "backupStatsStatus", "backupStatsEngine", "backupStatsType", "createdAt"})
	} else {
		return err
	}
	o.TotalBackup = all.TotalBackup
	o.TotalSize = all.TotalSize
	o.BackupStatsStatus = all.BackupStatsStatus
	o.BackupStatsEngine = all.BackupStatsEngine
	o.BackupStatsType = all.BackupStatsType
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
