// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

// ClusterBackup clusterBackup is the payload for cluster backup
type ClusterBackup struct {
	// PITREnabled or not
	PitrEnabled *bool `json:"pitrEnabled,omitempty"`
	// autoBackup or not
	AutoBackup *bool `json:"autoBackup,omitempty"`
	// name of the backup method
	AutoBackupMethod *string `json:"autoBackupMethod,omitempty"`
	// backupRepoName is the name of backupRepo and it is used to store the backup data
	BackupRepo *string `json:"backupRepo,omitempty"`
	// cronExpression specifies the cron expression
	CronExpression *string `json:"cronExpression,omitempty"`
	// retentionPeriod specifies the retention period
	RetentionPeriod *string `json:"retentionPeriod,omitempty"`
	// backup retention policy when cluster is deleted
	RetentionPolicy *BackupRetentionPolicy `json:"retentionPolicy,omitempty"`
	// snapshotVolumes specifies whether to take snapshots of persistent volumes to back up
	SnapshotVolumes common.NullableBool `json:"snapshotVolumes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterBackup instantiates a new ClusterBackup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterBackup() *ClusterBackup {
	this := ClusterBackup{}
	var pitrEnabled bool = false
	this.PitrEnabled = &pitrEnabled
	var autoBackup bool = false
	this.AutoBackup = &autoBackup
	return &this
}

// NewClusterBackupWithDefaults instantiates a new ClusterBackup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterBackupWithDefaults() *ClusterBackup {
	this := ClusterBackup{}
	var pitrEnabled bool = false
	this.PitrEnabled = &pitrEnabled
	var autoBackup bool = false
	this.AutoBackup = &autoBackup
	return &this
}

// GetPitrEnabled returns the PitrEnabled field value if set, zero value otherwise.
func (o *ClusterBackup) GetPitrEnabled() bool {
	if o == nil || o.PitrEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PitrEnabled
}

// GetPitrEnabledOk returns a tuple with the PitrEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBackup) GetPitrEnabledOk() (*bool, bool) {
	if o == nil || o.PitrEnabled == nil {
		return nil, false
	}
	return o.PitrEnabled, true
}

// HasPitrEnabled returns a boolean if a field has been set.
func (o *ClusterBackup) HasPitrEnabled() bool {
	return o != nil && o.PitrEnabled != nil
}

// SetPitrEnabled gets a reference to the given bool and assigns it to the PitrEnabled field.
func (o *ClusterBackup) SetPitrEnabled(v bool) {
	o.PitrEnabled = &v
}

// GetAutoBackup returns the AutoBackup field value if set, zero value otherwise.
func (o *ClusterBackup) GetAutoBackup() bool {
	if o == nil || o.AutoBackup == nil {
		var ret bool
		return ret
	}
	return *o.AutoBackup
}

// GetAutoBackupOk returns a tuple with the AutoBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBackup) GetAutoBackupOk() (*bool, bool) {
	if o == nil || o.AutoBackup == nil {
		return nil, false
	}
	return o.AutoBackup, true
}

// HasAutoBackup returns a boolean if a field has been set.
func (o *ClusterBackup) HasAutoBackup() bool {
	return o != nil && o.AutoBackup != nil
}

// SetAutoBackup gets a reference to the given bool and assigns it to the AutoBackup field.
func (o *ClusterBackup) SetAutoBackup(v bool) {
	o.AutoBackup = &v
}

// GetAutoBackupMethod returns the AutoBackupMethod field value if set, zero value otherwise.
func (o *ClusterBackup) GetAutoBackupMethod() string {
	if o == nil || o.AutoBackupMethod == nil {
		var ret string
		return ret
	}
	return *o.AutoBackupMethod
}

// GetAutoBackupMethodOk returns a tuple with the AutoBackupMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBackup) GetAutoBackupMethodOk() (*string, bool) {
	if o == nil || o.AutoBackupMethod == nil {
		return nil, false
	}
	return o.AutoBackupMethod, true
}

// HasAutoBackupMethod returns a boolean if a field has been set.
func (o *ClusterBackup) HasAutoBackupMethod() bool {
	return o != nil && o.AutoBackupMethod != nil
}

// SetAutoBackupMethod gets a reference to the given string and assigns it to the AutoBackupMethod field.
func (o *ClusterBackup) SetAutoBackupMethod(v string) {
	o.AutoBackupMethod = &v
}

// GetBackupRepo returns the BackupRepo field value if set, zero value otherwise.
func (o *ClusterBackup) GetBackupRepo() string {
	if o == nil || o.BackupRepo == nil {
		var ret string
		return ret
	}
	return *o.BackupRepo
}

// GetBackupRepoOk returns a tuple with the BackupRepo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBackup) GetBackupRepoOk() (*string, bool) {
	if o == nil || o.BackupRepo == nil {
		return nil, false
	}
	return o.BackupRepo, true
}

// HasBackupRepo returns a boolean if a field has been set.
func (o *ClusterBackup) HasBackupRepo() bool {
	return o != nil && o.BackupRepo != nil
}

// SetBackupRepo gets a reference to the given string and assigns it to the BackupRepo field.
func (o *ClusterBackup) SetBackupRepo(v string) {
	o.BackupRepo = &v
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *ClusterBackup) GetCronExpression() string {
	if o == nil || o.CronExpression == nil {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBackup) GetCronExpressionOk() (*string, bool) {
	if o == nil || o.CronExpression == nil {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *ClusterBackup) HasCronExpression() bool {
	return o != nil && o.CronExpression != nil
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *ClusterBackup) SetCronExpression(v string) {
	o.CronExpression = &v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *ClusterBackup) GetRetentionPeriod() string {
	if o == nil || o.RetentionPeriod == nil {
		var ret string
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBackup) GetRetentionPeriodOk() (*string, bool) {
	if o == nil || o.RetentionPeriod == nil {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *ClusterBackup) HasRetentionPeriod() bool {
	return o != nil && o.RetentionPeriod != nil
}

// SetRetentionPeriod gets a reference to the given string and assigns it to the RetentionPeriod field.
func (o *ClusterBackup) SetRetentionPeriod(v string) {
	o.RetentionPeriod = &v
}

// GetRetentionPolicy returns the RetentionPolicy field value if set, zero value otherwise.
func (o *ClusterBackup) GetRetentionPolicy() BackupRetentionPolicy {
	if o == nil || o.RetentionPolicy == nil {
		var ret BackupRetentionPolicy
		return ret
	}
	return *o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBackup) GetRetentionPolicyOk() (*BackupRetentionPolicy, bool) {
	if o == nil || o.RetentionPolicy == nil {
		return nil, false
	}
	return o.RetentionPolicy, true
}

// HasRetentionPolicy returns a boolean if a field has been set.
func (o *ClusterBackup) HasRetentionPolicy() bool {
	return o != nil && o.RetentionPolicy != nil
}

// SetRetentionPolicy gets a reference to the given BackupRetentionPolicy and assigns it to the RetentionPolicy field.
func (o *ClusterBackup) SetRetentionPolicy(v BackupRetentionPolicy) {
	o.RetentionPolicy = &v
}

// GetSnapshotVolumes returns the SnapshotVolumes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterBackup) GetSnapshotVolumes() bool {
	if o == nil || o.SnapshotVolumes.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SnapshotVolumes.Get()
}

// GetSnapshotVolumesOk returns a tuple with the SnapshotVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterBackup) GetSnapshotVolumesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnapshotVolumes.Get(), o.SnapshotVolumes.IsSet()
}

// HasSnapshotVolumes returns a boolean if a field has been set.
func (o *ClusterBackup) HasSnapshotVolumes() bool {
	return o != nil && o.SnapshotVolumes.IsSet()
}

// SetSnapshotVolumes gets a reference to the given common.NullableBool and assigns it to the SnapshotVolumes field.
func (o *ClusterBackup) SetSnapshotVolumes(v bool) {
	o.SnapshotVolumes.Set(&v)
}

// SetSnapshotVolumesNil sets the value for SnapshotVolumes to be an explicit nil.
func (o *ClusterBackup) SetSnapshotVolumesNil() {
	o.SnapshotVolumes.Set(nil)
}

// UnsetSnapshotVolumes ensures that no value is present for SnapshotVolumes, not even an explicit nil.
func (o *ClusterBackup) UnsetSnapshotVolumes() {
	o.SnapshotVolumes.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterBackup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.PitrEnabled != nil {
		toSerialize["pitrEnabled"] = o.PitrEnabled
	}
	if o.AutoBackup != nil {
		toSerialize["autoBackup"] = o.AutoBackup
	}
	if o.AutoBackupMethod != nil {
		toSerialize["autoBackupMethod"] = o.AutoBackupMethod
	}
	if o.BackupRepo != nil {
		toSerialize["backupRepo"] = o.BackupRepo
	}
	if o.CronExpression != nil {
		toSerialize["cronExpression"] = o.CronExpression
	}
	if o.RetentionPeriod != nil {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if o.RetentionPolicy != nil {
		toSerialize["retentionPolicy"] = o.RetentionPolicy
	}
	if o.SnapshotVolumes.IsSet() {
		toSerialize["snapshotVolumes"] = o.SnapshotVolumes.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterBackup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PitrEnabled      *bool                  `json:"pitrEnabled,omitempty"`
		AutoBackup       *bool                  `json:"autoBackup,omitempty"`
		AutoBackupMethod *string                `json:"autoBackupMethod,omitempty"`
		BackupRepo       *string                `json:"backupRepo,omitempty"`
		CronExpression   *string                `json:"cronExpression,omitempty"`
		RetentionPeriod  *string                `json:"retentionPeriod,omitempty"`
		RetentionPolicy  *BackupRetentionPolicy `json:"retentionPolicy,omitempty"`
		SnapshotVolumes  common.NullableBool    `json:"snapshotVolumes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pitrEnabled", "autoBackup", "autoBackupMethod", "backupRepo", "cronExpression", "retentionPeriod", "retentionPolicy", "snapshotVolumes"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PitrEnabled = all.PitrEnabled
	o.AutoBackup = all.AutoBackup
	o.AutoBackupMethod = all.AutoBackupMethod
	o.BackupRepo = all.BackupRepo
	o.CronExpression = all.CronExpression
	o.RetentionPeriod = all.RetentionPeriod
	if all.RetentionPolicy != nil && !all.RetentionPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.RetentionPolicy = all.RetentionPolicy
	}
	o.SnapshotVolumes = all.SnapshotVolumes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
