// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// BackupPolicy BackupPolicy is the payload for KubeBlocks cluster backup policy 
type BackupPolicy struct {
	// specify whether to use auto backup
	AutoBackup *bool `json:"autoBackup,omitempty"`
	// the auto backup method
	AutoBackupMethod *string `json:"autoBackupMethod,omitempty"`
	// specify whether to enable point-in-time recovery
	PitrEnabled *bool `json:"pitrEnabled,omitempty"`
	// the crop expression for schedule
	CronExpression *string `json:"cronExpression,omitempty"`
	// specify the retention period
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
	// the backup repo name, which is used to store backup data
	BackupRepo *string `json:"backupRepo,omitempty"`
	// backup retention policy when cluster is deleted
	RetentionPolicy *BackupRetentionPolicy `json:"retentionPolicy,omitempty"`
	// the time to do next backup
	NextBackupTime *time.Time `json:"nextBackupTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewBackupPolicy instantiates a new BackupPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupPolicy() *BackupPolicy {
	this := BackupPolicy{}
	var autoBackup bool = false
	this.AutoBackup = &autoBackup
	var pitrEnabled bool = false
	this.PitrEnabled = &pitrEnabled
	return &this
}

// NewBackupPolicyWithDefaults instantiates a new BackupPolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupPolicyWithDefaults() *BackupPolicy {
	this := BackupPolicy{}
	var autoBackup bool = false
	this.AutoBackup = &autoBackup
	var pitrEnabled bool = false
	this.PitrEnabled = &pitrEnabled
	return &this
}
// GetAutoBackup returns the AutoBackup field value if set, zero value otherwise.
func (o *BackupPolicy) GetAutoBackup() bool {
	if o == nil || o.AutoBackup == nil {
		var ret bool
		return ret
	}
	return *o.AutoBackup
}

// GetAutoBackupOk returns a tuple with the AutoBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetAutoBackupOk() (*bool, bool) {
	if o == nil || o.AutoBackup == nil {
		return nil, false
	}
	return o.AutoBackup, true
}

// HasAutoBackup returns a boolean if a field has been set.
func (o *BackupPolicy) HasAutoBackup() bool {
	return o != nil && o.AutoBackup != nil
}

// SetAutoBackup gets a reference to the given bool and assigns it to the AutoBackup field.
func (o *BackupPolicy) SetAutoBackup(v bool) {
	o.AutoBackup = &v
}


// GetAutoBackupMethod returns the AutoBackupMethod field value if set, zero value otherwise.
func (o *BackupPolicy) GetAutoBackupMethod() string {
	if o == nil || o.AutoBackupMethod == nil {
		var ret string
		return ret
	}
	return *o.AutoBackupMethod
}

// GetAutoBackupMethodOk returns a tuple with the AutoBackupMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetAutoBackupMethodOk() (*string, bool) {
	if o == nil || o.AutoBackupMethod == nil {
		return nil, false
	}
	return o.AutoBackupMethod, true
}

// HasAutoBackupMethod returns a boolean if a field has been set.
func (o *BackupPolicy) HasAutoBackupMethod() bool {
	return o != nil && o.AutoBackupMethod != nil
}

// SetAutoBackupMethod gets a reference to the given string and assigns it to the AutoBackupMethod field.
func (o *BackupPolicy) SetAutoBackupMethod(v string) {
	o.AutoBackupMethod = &v
}


// GetPitrEnabled returns the PitrEnabled field value if set, zero value otherwise.
func (o *BackupPolicy) GetPitrEnabled() bool {
	if o == nil || o.PitrEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PitrEnabled
}

// GetPitrEnabledOk returns a tuple with the PitrEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetPitrEnabledOk() (*bool, bool) {
	if o == nil || o.PitrEnabled == nil {
		return nil, false
	}
	return o.PitrEnabled, true
}

// HasPitrEnabled returns a boolean if a field has been set.
func (o *BackupPolicy) HasPitrEnabled() bool {
	return o != nil && o.PitrEnabled != nil
}

// SetPitrEnabled gets a reference to the given bool and assigns it to the PitrEnabled field.
func (o *BackupPolicy) SetPitrEnabled(v bool) {
	o.PitrEnabled = &v
}


// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *BackupPolicy) GetCronExpression() string {
	if o == nil || o.CronExpression == nil {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetCronExpressionOk() (*string, bool) {
	if o == nil || o.CronExpression == nil {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *BackupPolicy) HasCronExpression() bool {
	return o != nil && o.CronExpression != nil
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *BackupPolicy) SetCronExpression(v string) {
	o.CronExpression = &v
}


// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *BackupPolicy) GetRetentionPeriod() string {
	if o == nil || o.RetentionPeriod == nil {
		var ret string
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetRetentionPeriodOk() (*string, bool) {
	if o == nil || o.RetentionPeriod == nil {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *BackupPolicy) HasRetentionPeriod() bool {
	return o != nil && o.RetentionPeriod != nil
}

// SetRetentionPeriod gets a reference to the given string and assigns it to the RetentionPeriod field.
func (o *BackupPolicy) SetRetentionPeriod(v string) {
	o.RetentionPeriod = &v
}


// GetBackupRepo returns the BackupRepo field value if set, zero value otherwise.
func (o *BackupPolicy) GetBackupRepo() string {
	if o == nil || o.BackupRepo == nil {
		var ret string
		return ret
	}
	return *o.BackupRepo
}

// GetBackupRepoOk returns a tuple with the BackupRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetBackupRepoOk() (*string, bool) {
	if o == nil || o.BackupRepo == nil {
		return nil, false
	}
	return o.BackupRepo, true
}

// HasBackupRepo returns a boolean if a field has been set.
func (o *BackupPolicy) HasBackupRepo() bool {
	return o != nil && o.BackupRepo != nil
}

// SetBackupRepo gets a reference to the given string and assigns it to the BackupRepo field.
func (o *BackupPolicy) SetBackupRepo(v string) {
	o.BackupRepo = &v
}


// GetRetentionPolicy returns the RetentionPolicy field value if set, zero value otherwise.
func (o *BackupPolicy) GetRetentionPolicy() BackupRetentionPolicy {
	if o == nil || o.RetentionPolicy == nil {
		var ret BackupRetentionPolicy
		return ret
	}
	return *o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetRetentionPolicyOk() (*BackupRetentionPolicy, bool) {
	if o == nil || o.RetentionPolicy == nil {
		return nil, false
	}
	return o.RetentionPolicy, true
}

// HasRetentionPolicy returns a boolean if a field has been set.
func (o *BackupPolicy) HasRetentionPolicy() bool {
	return o != nil && o.RetentionPolicy != nil
}

// SetRetentionPolicy gets a reference to the given BackupRetentionPolicy and assigns it to the RetentionPolicy field.
func (o *BackupPolicy) SetRetentionPolicy(v BackupRetentionPolicy) {
	o.RetentionPolicy = &v
}


// GetNextBackupTime returns the NextBackupTime field value if set, zero value otherwise.
func (o *BackupPolicy) GetNextBackupTime() time.Time {
	if o == nil || o.NextBackupTime == nil {
		var ret time.Time
		return ret
	}
	return *o.NextBackupTime
}

// GetNextBackupTimeOk returns a tuple with the NextBackupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetNextBackupTimeOk() (*time.Time, bool) {
	if o == nil || o.NextBackupTime == nil {
		return nil, false
	}
	return o.NextBackupTime, true
}

// HasNextBackupTime returns a boolean if a field has been set.
func (o *BackupPolicy) HasNextBackupTime() bool {
	return o != nil && o.NextBackupTime != nil
}

// SetNextBackupTime gets a reference to the given time.Time and assigns it to the NextBackupTime field.
func (o *BackupPolicy) SetNextBackupTime(v time.Time) {
	o.NextBackupTime = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o BackupPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.AutoBackup != nil {
		toSerialize["autoBackup"] = o.AutoBackup
	}
	if o.AutoBackupMethod != nil {
		toSerialize["autoBackupMethod"] = o.AutoBackupMethod
	}
	if o.PitrEnabled != nil {
		toSerialize["pitrEnabled"] = o.PitrEnabled
	}
	if o.CronExpression != nil {
		toSerialize["cronExpression"] = o.CronExpression
	}
	if o.RetentionPeriod != nil {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if o.BackupRepo != nil {
		toSerialize["backupRepo"] = o.BackupRepo
	}
	if o.RetentionPolicy != nil {
		toSerialize["retentionPolicy"] = o.RetentionPolicy
	}
	if o.NextBackupTime != nil {
		if o.NextBackupTime.Nanosecond() == 0 {
			toSerialize["nextBackupTime"] = o.NextBackupTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["nextBackupTime"] = o.NextBackupTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoBackup *bool `json:"autoBackup,omitempty"`
		AutoBackupMethod *string `json:"autoBackupMethod,omitempty"`
		PitrEnabled *bool `json:"pitrEnabled,omitempty"`
		CronExpression *string `json:"cronExpression,omitempty"`
		RetentionPeriod *string `json:"retentionPeriod,omitempty"`
		BackupRepo *string `json:"backupRepo,omitempty"`
		RetentionPolicy *BackupRetentionPolicy `json:"retentionPolicy,omitempty"`
		NextBackupTime *time.Time `json:"nextBackupTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "autoBackup", "autoBackupMethod", "pitrEnabled", "cronExpression", "retentionPeriod", "backupRepo", "retentionPolicy", "nextBackupTime",  })
	} else {
		return err
	}

	hasInvalidField := false
	o.AutoBackup = all.AutoBackup
	o.AutoBackupMethod = all.AutoBackupMethod
	o.PitrEnabled = all.PitrEnabled
	o.CronExpression = all.CronExpression
	o.RetentionPeriod = all.RetentionPeriod
	o.BackupRepo = all.BackupRepo
	if all.RetentionPolicy != nil &&!all.RetentionPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.RetentionPolicy = all.RetentionPolicy
	}
	o.NextBackupTime = all.NextBackupTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
