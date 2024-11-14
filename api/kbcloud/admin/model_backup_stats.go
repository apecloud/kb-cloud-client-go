// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupStats Backup statistic info
type BackupStats struct {
	// Number of backups for each different status
	BackupStatsStatus []BackupStatsStatus `json:"backupStatsStatus,omitempty"`
	// TotalSize and number of backups for each engine
	BackupStatsEngine []BackupStatsEngine `json:"backupStatsEngine,omitempty"`
	// TotalSize and number of backups for each type
	BackupStatsType []BackupStatsType `json:"backupStatsType,omitempty"`
	// create time of the latest backup
	LatestBackupTime common.NullableTime `json:"latestBackupTime,omitempty"`
	// backup status of the latest backup
	LatestBackupStatus common.NullableString `json:"latestBackupStatus,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupStats instantiates a new BackupStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupStats() *BackupStats {
	this := BackupStats{}
	return &this
}

// NewBackupStatsWithDefaults instantiates a new BackupStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupStatsWithDefaults() *BackupStats {
	this := BackupStats{}
	return &this
}

// GetBackupStatsStatus returns the BackupStatsStatus field value if set, zero value otherwise.
func (o *BackupStats) GetBackupStatsStatus() []BackupStatsStatus {
	if o == nil || o.BackupStatsStatus == nil {
		var ret []BackupStatsStatus
		return ret
	}
	return o.BackupStatsStatus
}

// GetBackupStatsStatusOk returns a tuple with the BackupStatsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStats) GetBackupStatsStatusOk() (*[]BackupStatsStatus, bool) {
	if o == nil || o.BackupStatsStatus == nil {
		return nil, false
	}
	return &o.BackupStatsStatus, true
}

// HasBackupStatsStatus returns a boolean if a field has been set.
func (o *BackupStats) HasBackupStatsStatus() bool {
	return o != nil && o.BackupStatsStatus != nil
}

// SetBackupStatsStatus gets a reference to the given []BackupStatsStatus and assigns it to the BackupStatsStatus field.
func (o *BackupStats) SetBackupStatsStatus(v []BackupStatsStatus) {
	o.BackupStatsStatus = v
}

// GetBackupStatsEngine returns the BackupStatsEngine field value if set, zero value otherwise.
func (o *BackupStats) GetBackupStatsEngine() []BackupStatsEngine {
	if o == nil || o.BackupStatsEngine == nil {
		var ret []BackupStatsEngine
		return ret
	}
	return o.BackupStatsEngine
}

// GetBackupStatsEngineOk returns a tuple with the BackupStatsEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStats) GetBackupStatsEngineOk() (*[]BackupStatsEngine, bool) {
	if o == nil || o.BackupStatsEngine == nil {
		return nil, false
	}
	return &o.BackupStatsEngine, true
}

// HasBackupStatsEngine returns a boolean if a field has been set.
func (o *BackupStats) HasBackupStatsEngine() bool {
	return o != nil && o.BackupStatsEngine != nil
}

// SetBackupStatsEngine gets a reference to the given []BackupStatsEngine and assigns it to the BackupStatsEngine field.
func (o *BackupStats) SetBackupStatsEngine(v []BackupStatsEngine) {
	o.BackupStatsEngine = v
}

// GetBackupStatsType returns the BackupStatsType field value if set, zero value otherwise.
func (o *BackupStats) GetBackupStatsType() []BackupStatsType {
	if o == nil || o.BackupStatsType == nil {
		var ret []BackupStatsType
		return ret
	}
	return o.BackupStatsType
}

// GetBackupStatsTypeOk returns a tuple with the BackupStatsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupStats) GetBackupStatsTypeOk() (*[]BackupStatsType, bool) {
	if o == nil || o.BackupStatsType == nil {
		return nil, false
	}
	return &o.BackupStatsType, true
}

// HasBackupStatsType returns a boolean if a field has been set.
func (o *BackupStats) HasBackupStatsType() bool {
	return o != nil && o.BackupStatsType != nil
}

// SetBackupStatsType gets a reference to the given []BackupStatsType and assigns it to the BackupStatsType field.
func (o *BackupStats) SetBackupStatsType(v []BackupStatsType) {
	o.BackupStatsType = v
}

// GetLatestBackupTime returns the LatestBackupTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupStats) GetLatestBackupTime() time.Time {
	if o == nil || o.LatestBackupTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LatestBackupTime.Get()
}

// GetLatestBackupTimeOk returns a tuple with the LatestBackupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BackupStats) GetLatestBackupTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestBackupTime.Get(), o.LatestBackupTime.IsSet()
}

// HasLatestBackupTime returns a boolean if a field has been set.
func (o *BackupStats) HasLatestBackupTime() bool {
	return o != nil && o.LatestBackupTime.IsSet()
}

// SetLatestBackupTime gets a reference to the given common.NullableTime and assigns it to the LatestBackupTime field.
func (o *BackupStats) SetLatestBackupTime(v time.Time) {
	o.LatestBackupTime.Set(&v)
}

// SetLatestBackupTimeNil sets the value for LatestBackupTime to be an explicit nil.
func (o *BackupStats) SetLatestBackupTimeNil() {
	o.LatestBackupTime.Set(nil)
}

// UnsetLatestBackupTime ensures that no value is present for LatestBackupTime, not even an explicit nil.
func (o *BackupStats) UnsetLatestBackupTime() {
	o.LatestBackupTime.Unset()
}

// GetLatestBackupStatus returns the LatestBackupStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupStats) GetLatestBackupStatus() string {
	if o == nil || o.LatestBackupStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.LatestBackupStatus.Get()
}

// GetLatestBackupStatusOk returns a tuple with the LatestBackupStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BackupStats) GetLatestBackupStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestBackupStatus.Get(), o.LatestBackupStatus.IsSet()
}

// HasLatestBackupStatus returns a boolean if a field has been set.
func (o *BackupStats) HasLatestBackupStatus() bool {
	return o != nil && o.LatestBackupStatus.IsSet()
}

// SetLatestBackupStatus gets a reference to the given common.NullableString and assigns it to the LatestBackupStatus field.
func (o *BackupStats) SetLatestBackupStatus(v string) {
	o.LatestBackupStatus.Set(&v)
}

// SetLatestBackupStatusNil sets the value for LatestBackupStatus to be an explicit nil.
func (o *BackupStats) SetLatestBackupStatusNil() {
	o.LatestBackupStatus.Set(nil)
}

// UnsetLatestBackupStatus ensures that no value is present for LatestBackupStatus, not even an explicit nil.
func (o *BackupStats) UnsetLatestBackupStatus() {
	o.LatestBackupStatus.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
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
	if o.LatestBackupTime.IsSet() {
		toSerialize["latestBackupTime"] = o.LatestBackupTime.Get()
	}
	if o.LatestBackupStatus.IsSet() {
		toSerialize["latestBackupStatus"] = o.LatestBackupStatus.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BackupStatsStatus  []BackupStatsStatus   `json:"backupStatsStatus,omitempty"`
		BackupStatsEngine  []BackupStatsEngine   `json:"backupStatsEngine,omitempty"`
		BackupStatsType    []BackupStatsType     `json:"backupStatsType,omitempty"`
		LatestBackupTime   common.NullableTime   `json:"latestBackupTime,omitempty"`
		LatestBackupStatus common.NullableString `json:"latestBackupStatus,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"backupStatsStatus", "backupStatsEngine", "backupStatsType", "latestBackupTime", "latestBackupStatus"})
	} else {
		return err
	}
	o.BackupStatsStatus = all.BackupStatsStatus
	o.BackupStatsEngine = all.BackupStatsEngine
	o.BackupStatsType = all.BackupStatsType
	o.LatestBackupTime = all.LatestBackupTime
	o.LatestBackupStatus = all.LatestBackupStatus

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
