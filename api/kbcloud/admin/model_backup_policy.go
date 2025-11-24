// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BackupPolicy BackupPolicy is the payload for KubeBlocks cluster backup policy
type BackupPolicy struct {
	// specify whether to use auto backup
	AutoBackup *bool `json:"autoBackup,omitempty"`
	// the auto full backup method
	AutoBackupMethod *string `json:"autoBackupMethod,omitempty"`
	// specify whether to enable point-in-time recovery
	PitrEnabled *bool `json:"pitrEnabled,omitempty"`
	// continuous backup method for pitr
	ContinuousBackupMethod *string `json:"continuousBackupMethod,omitempty"`
	// the crop expression for schedule
	CronExpression *string `json:"cronExpression,omitempty"`
	// specify whether to enable incremental backup
	IncrementalBackupEnabled *bool `json:"incrementalBackupEnabled,omitempty"`
	// the crop expression for incremental backup schedule
	IncrementalCronExpression *string `json:"incrementalCronExpression,omitempty"`
	// specify the retention period
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
	// the backup repo name, which is used to store backup data
	BackupRepo *string `json:"backupRepo,omitempty"`
	// backup retention policy when cluster is deleted
	RetentionPolicy *BackupRetentionPolicy `json:"retentionPolicy,omitempty"`
	// the time to do next backup
	NextBackupTime common.NullableTime `json:"nextBackupTime,omitempty"`
	// encryption config for cluster
	EncryptionConfig *EncryptionConfig            `json:"encryptionConfig,omitempty"`
	BackupParameters map[string]map[string]string `json:"backupParameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
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
	var incrementalBackupEnabled bool = false
	this.IncrementalBackupEnabled = &incrementalBackupEnabled
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
	var incrementalBackupEnabled bool = false
	this.IncrementalBackupEnabled = &incrementalBackupEnabled
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

// GetContinuousBackupMethod returns the ContinuousBackupMethod field value if set, zero value otherwise.
func (o *BackupPolicy) GetContinuousBackupMethod() string {
	if o == nil || o.ContinuousBackupMethod == nil {
		var ret string
		return ret
	}
	return *o.ContinuousBackupMethod
}

// GetContinuousBackupMethodOk returns a tuple with the ContinuousBackupMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetContinuousBackupMethodOk() (*string, bool) {
	if o == nil || o.ContinuousBackupMethod == nil {
		return nil, false
	}
	return o.ContinuousBackupMethod, true
}

// HasContinuousBackupMethod returns a boolean if a field has been set.
func (o *BackupPolicy) HasContinuousBackupMethod() bool {
	return o != nil && o.ContinuousBackupMethod != nil
}

// SetContinuousBackupMethod gets a reference to the given string and assigns it to the ContinuousBackupMethod field.
func (o *BackupPolicy) SetContinuousBackupMethod(v string) {
	o.ContinuousBackupMethod = &v
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

// GetIncrementalBackupEnabled returns the IncrementalBackupEnabled field value if set, zero value otherwise.
func (o *BackupPolicy) GetIncrementalBackupEnabled() bool {
	if o == nil || o.IncrementalBackupEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IncrementalBackupEnabled
}

// GetIncrementalBackupEnabledOk returns a tuple with the IncrementalBackupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetIncrementalBackupEnabledOk() (*bool, bool) {
	if o == nil || o.IncrementalBackupEnabled == nil {
		return nil, false
	}
	return o.IncrementalBackupEnabled, true
}

// HasIncrementalBackupEnabled returns a boolean if a field has been set.
func (o *BackupPolicy) HasIncrementalBackupEnabled() bool {
	return o != nil && o.IncrementalBackupEnabled != nil
}

// SetIncrementalBackupEnabled gets a reference to the given bool and assigns it to the IncrementalBackupEnabled field.
func (o *BackupPolicy) SetIncrementalBackupEnabled(v bool) {
	o.IncrementalBackupEnabled = &v
}

// GetIncrementalCronExpression returns the IncrementalCronExpression field value if set, zero value otherwise.
func (o *BackupPolicy) GetIncrementalCronExpression() string {
	if o == nil || o.IncrementalCronExpression == nil {
		var ret string
		return ret
	}
	return *o.IncrementalCronExpression
}

// GetIncrementalCronExpressionOk returns a tuple with the IncrementalCronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetIncrementalCronExpressionOk() (*string, bool) {
	if o == nil || o.IncrementalCronExpression == nil {
		return nil, false
	}
	return o.IncrementalCronExpression, true
}

// HasIncrementalCronExpression returns a boolean if a field has been set.
func (o *BackupPolicy) HasIncrementalCronExpression() bool {
	return o != nil && o.IncrementalCronExpression != nil
}

// SetIncrementalCronExpression gets a reference to the given string and assigns it to the IncrementalCronExpression field.
func (o *BackupPolicy) SetIncrementalCronExpression(v string) {
	o.IncrementalCronExpression = &v
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

// GetNextBackupTime returns the NextBackupTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BackupPolicy) GetNextBackupTime() time.Time {
	if o == nil || o.NextBackupTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.NextBackupTime.Get()
}

// GetNextBackupTimeOk returns a tuple with the NextBackupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BackupPolicy) GetNextBackupTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextBackupTime.Get(), o.NextBackupTime.IsSet()
}

// HasNextBackupTime returns a boolean if a field has been set.
func (o *BackupPolicy) HasNextBackupTime() bool {
	return o != nil && o.NextBackupTime.IsSet()
}

// SetNextBackupTime gets a reference to the given common.NullableTime and assigns it to the NextBackupTime field.
func (o *BackupPolicy) SetNextBackupTime(v time.Time) {
	o.NextBackupTime.Set(&v)
}

// SetNextBackupTimeNil sets the value for NextBackupTime to be an explicit nil.
func (o *BackupPolicy) SetNextBackupTimeNil() {
	o.NextBackupTime.Set(nil)
}

// UnsetNextBackupTime ensures that no value is present for NextBackupTime, not even an explicit nil.
func (o *BackupPolicy) UnsetNextBackupTime() {
	o.NextBackupTime.Unset()
}

// GetEncryptionConfig returns the EncryptionConfig field value if set, zero value otherwise.
func (o *BackupPolicy) GetEncryptionConfig() EncryptionConfig {
	if o == nil || o.EncryptionConfig == nil {
		var ret EncryptionConfig
		return ret
	}
	return *o.EncryptionConfig
}

// GetEncryptionConfigOk returns a tuple with the EncryptionConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetEncryptionConfigOk() (*EncryptionConfig, bool) {
	if o == nil || o.EncryptionConfig == nil {
		return nil, false
	}
	return o.EncryptionConfig, true
}

// HasEncryptionConfig returns a boolean if a field has been set.
func (o *BackupPolicy) HasEncryptionConfig() bool {
	return o != nil && o.EncryptionConfig != nil
}

// SetEncryptionConfig gets a reference to the given EncryptionConfig and assigns it to the EncryptionConfig field.
func (o *BackupPolicy) SetEncryptionConfig(v EncryptionConfig) {
	o.EncryptionConfig = &v
}

// GetBackupParameters returns the BackupParameters field value if set, zero value otherwise.
func (o *BackupPolicy) GetBackupParameters() map[string]map[string]string {
	if o == nil || o.BackupParameters == nil {
		var ret map[string]map[string]string
		return ret
	}
	return o.BackupParameters
}

// GetBackupParametersOk returns a tuple with the BackupParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupPolicy) GetBackupParametersOk() (*map[string]map[string]string, bool) {
	if o == nil || o.BackupParameters == nil {
		return nil, false
	}
	return &o.BackupParameters, true
}

// HasBackupParameters returns a boolean if a field has been set.
func (o *BackupPolicy) HasBackupParameters() bool {
	return o != nil && o.BackupParameters != nil
}

// SetBackupParameters gets a reference to the given map[string]map[string]string and assigns it to the BackupParameters field.
func (o *BackupPolicy) SetBackupParameters(v map[string]map[string]string) {
	o.BackupParameters = v
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
	if o.ContinuousBackupMethod != nil {
		toSerialize["continuousBackupMethod"] = o.ContinuousBackupMethod
	}
	if o.CronExpression != nil {
		toSerialize["cronExpression"] = o.CronExpression
	}
	if o.IncrementalBackupEnabled != nil {
		toSerialize["incrementalBackupEnabled"] = o.IncrementalBackupEnabled
	}
	if o.IncrementalCronExpression != nil {
		toSerialize["incrementalCronExpression"] = o.IncrementalCronExpression
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
	if o.NextBackupTime.IsSet() {
		toSerialize["nextBackupTime"] = o.NextBackupTime.Get()
	}
	if o.EncryptionConfig != nil {
		toSerialize["encryptionConfig"] = o.EncryptionConfig
	}
	if o.BackupParameters != nil {
		toSerialize["backupParameters"] = o.BackupParameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoBackup                *bool                        `json:"autoBackup,omitempty"`
		AutoBackupMethod          *string                      `json:"autoBackupMethod,omitempty"`
		PitrEnabled               *bool                        `json:"pitrEnabled,omitempty"`
		ContinuousBackupMethod    *string                      `json:"continuousBackupMethod,omitempty"`
		CronExpression            *string                      `json:"cronExpression,omitempty"`
		IncrementalBackupEnabled  *bool                        `json:"incrementalBackupEnabled,omitempty"`
		IncrementalCronExpression *string                      `json:"incrementalCronExpression,omitempty"`
		RetentionPeriod           *string                      `json:"retentionPeriod,omitempty"`
		BackupRepo                *string                      `json:"backupRepo,omitempty"`
		RetentionPolicy           *BackupRetentionPolicy       `json:"retentionPolicy,omitempty"`
		NextBackupTime            common.NullableTime          `json:"nextBackupTime,omitempty"`
		EncryptionConfig          *EncryptionConfig            `json:"encryptionConfig,omitempty"`
		BackupParameters          map[string]map[string]string `json:"backupParameters,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"autoBackup", "autoBackupMethod", "pitrEnabled", "continuousBackupMethod", "cronExpression", "incrementalBackupEnabled", "incrementalCronExpression", "retentionPeriod", "backupRepo", "retentionPolicy", "nextBackupTime", "encryptionConfig", "backupParameters"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AutoBackup = all.AutoBackup
	o.AutoBackupMethod = all.AutoBackupMethod
	o.PitrEnabled = all.PitrEnabled
	o.ContinuousBackupMethod = all.ContinuousBackupMethod
	o.CronExpression = all.CronExpression
	o.IncrementalBackupEnabled = all.IncrementalBackupEnabled
	o.IncrementalCronExpression = all.IncrementalCronExpression
	o.RetentionPeriod = all.RetentionPeriod
	o.BackupRepo = all.BackupRepo
	if all.RetentionPolicy != nil && !all.RetentionPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.RetentionPolicy = all.RetentionPolicy
	}
	o.NextBackupTime = all.NextBackupTime
	if all.EncryptionConfig != nil && all.EncryptionConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EncryptionConfig = all.EncryptionConfig
	o.BackupParameters = all.BackupParameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
