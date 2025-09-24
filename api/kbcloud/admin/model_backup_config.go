// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type BackupConfig struct {
	Provider        *string `json:"provider,omitempty"`
	Schedule        *string `json:"schedule,omitempty"`
	AccessKeyId     *string `json:"accessKeyId,omitempty"`
	SecretAccessKey *string `json:"secretAccessKey,omitempty"`
	Endpoint        *string `json:"endpoint,omitempty"`
	Region          *string `json:"region,omitempty"`
	Bucket          *string `json:"bucket,omitempty"`
	// enable or disable auto backup
	AutoBackup *bool `json:"autoBackup,omitempty"`
	// time for next backup
	NextBackupTime *string `json:"nextBackupTime,omitempty"`
	// specify the retention days
	RetentionDay *int32 `json:"retentionDay,omitempty"`
	// the time of last backup
	LastBackupTime *string `json:"lastBackupTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupConfig instantiates a new BackupConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupConfig() *BackupConfig {
	this := BackupConfig{}
	var autoBackup bool = false
	this.AutoBackup = &autoBackup
	var retentionDay int32 = 5
	this.RetentionDay = &retentionDay
	return &this
}

// NewBackupConfigWithDefaults instantiates a new BackupConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupConfigWithDefaults() *BackupConfig {
	this := BackupConfig{}
	var autoBackup bool = false
	this.AutoBackup = &autoBackup
	var retentionDay int32 = 5
	this.RetentionDay = &retentionDay
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *BackupConfig) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *BackupConfig) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *BackupConfig) SetProvider(v string) {
	o.Provider = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *BackupConfig) GetSchedule() string {
	if o == nil || o.Schedule == nil {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetScheduleOk() (*string, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *BackupConfig) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *BackupConfig) SetSchedule(v string) {
	o.Schedule = &v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *BackupConfig) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *BackupConfig) HasAccessKeyId() bool {
	return o != nil && o.AccessKeyId != nil
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *BackupConfig) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetSecretAccessKey returns the SecretAccessKey field value if set, zero value otherwise.
func (o *BackupConfig) GetSecretAccessKey() string {
	if o == nil || o.SecretAccessKey == nil {
		var ret string
		return ret
	}
	return *o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil || o.SecretAccessKey == nil {
		return nil, false
	}
	return o.SecretAccessKey, true
}

// HasSecretAccessKey returns a boolean if a field has been set.
func (o *BackupConfig) HasSecretAccessKey() bool {
	return o != nil && o.SecretAccessKey != nil
}

// SetSecretAccessKey gets a reference to the given string and assigns it to the SecretAccessKey field.
func (o *BackupConfig) SetSecretAccessKey(v string) {
	o.SecretAccessKey = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *BackupConfig) GetEndpoint() string {
	if o == nil || o.Endpoint == nil {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetEndpointOk() (*string, bool) {
	if o == nil || o.Endpoint == nil {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *BackupConfig) HasEndpoint() bool {
	return o != nil && o.Endpoint != nil
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *BackupConfig) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *BackupConfig) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *BackupConfig) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *BackupConfig) SetRegion(v string) {
	o.Region = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *BackupConfig) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *BackupConfig) HasBucket() bool {
	return o != nil && o.Bucket != nil
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *BackupConfig) SetBucket(v string) {
	o.Bucket = &v
}

// GetAutoBackup returns the AutoBackup field value if set, zero value otherwise.
func (o *BackupConfig) GetAutoBackup() bool {
	if o == nil || o.AutoBackup == nil {
		var ret bool
		return ret
	}
	return *o.AutoBackup
}

// GetAutoBackupOk returns a tuple with the AutoBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetAutoBackupOk() (*bool, bool) {
	if o == nil || o.AutoBackup == nil {
		return nil, false
	}
	return o.AutoBackup, true
}

// HasAutoBackup returns a boolean if a field has been set.
func (o *BackupConfig) HasAutoBackup() bool {
	return o != nil && o.AutoBackup != nil
}

// SetAutoBackup gets a reference to the given bool and assigns it to the AutoBackup field.
func (o *BackupConfig) SetAutoBackup(v bool) {
	o.AutoBackup = &v
}

// GetNextBackupTime returns the NextBackupTime field value if set, zero value otherwise.
func (o *BackupConfig) GetNextBackupTime() string {
	if o == nil || o.NextBackupTime == nil {
		var ret string
		return ret
	}
	return *o.NextBackupTime
}

// GetNextBackupTimeOk returns a tuple with the NextBackupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetNextBackupTimeOk() (*string, bool) {
	if o == nil || o.NextBackupTime == nil {
		return nil, false
	}
	return o.NextBackupTime, true
}

// HasNextBackupTime returns a boolean if a field has been set.
func (o *BackupConfig) HasNextBackupTime() bool {
	return o != nil && o.NextBackupTime != nil
}

// SetNextBackupTime gets a reference to the given string and assigns it to the NextBackupTime field.
func (o *BackupConfig) SetNextBackupTime(v string) {
	o.NextBackupTime = &v
}

// GetRetentionDay returns the RetentionDay field value if set, zero value otherwise.
func (o *BackupConfig) GetRetentionDay() int32 {
	if o == nil || o.RetentionDay == nil {
		var ret int32
		return ret
	}
	return *o.RetentionDay
}

// GetRetentionDayOk returns a tuple with the RetentionDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetRetentionDayOk() (*int32, bool) {
	if o == nil || o.RetentionDay == nil {
		return nil, false
	}
	return o.RetentionDay, true
}

// HasRetentionDay returns a boolean if a field has been set.
func (o *BackupConfig) HasRetentionDay() bool {
	return o != nil && o.RetentionDay != nil
}

// SetRetentionDay gets a reference to the given int32 and assigns it to the RetentionDay field.
func (o *BackupConfig) SetRetentionDay(v int32) {
	o.RetentionDay = &v
}

// GetLastBackupTime returns the LastBackupTime field value if set, zero value otherwise.
func (o *BackupConfig) GetLastBackupTime() string {
	if o == nil || o.LastBackupTime == nil {
		var ret string
		return ret
	}
	return *o.LastBackupTime
}

// GetLastBackupTimeOk returns a tuple with the LastBackupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetLastBackupTimeOk() (*string, bool) {
	if o == nil || o.LastBackupTime == nil {
		return nil, false
	}
	return o.LastBackupTime, true
}

// HasLastBackupTime returns a boolean if a field has been set.
func (o *BackupConfig) HasLastBackupTime() bool {
	return o != nil && o.LastBackupTime != nil
}

// SetLastBackupTime gets a reference to the given string and assigns it to the LastBackupTime field.
func (o *BackupConfig) SetLastBackupTime(v string) {
	o.LastBackupTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.AccessKeyId != nil {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if o.SecretAccessKey != nil {
		toSerialize["secretAccessKey"] = o.SecretAccessKey
	}
	if o.Endpoint != nil {
		toSerialize["endpoint"] = o.Endpoint
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.AutoBackup != nil {
		toSerialize["autoBackup"] = o.AutoBackup
	}
	if o.NextBackupTime != nil {
		toSerialize["nextBackupTime"] = o.NextBackupTime
	}
	if o.RetentionDay != nil {
		toSerialize["retentionDay"] = o.RetentionDay
	}
	if o.LastBackupTime != nil {
		toSerialize["lastBackupTime"] = o.LastBackupTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Provider        *string `json:"provider,omitempty"`
		Schedule        *string `json:"schedule,omitempty"`
		AccessKeyId     *string `json:"accessKeyId,omitempty"`
		SecretAccessKey *string `json:"secretAccessKey,omitempty"`
		Endpoint        *string `json:"endpoint,omitempty"`
		Region          *string `json:"region,omitempty"`
		Bucket          *string `json:"bucket,omitempty"`
		AutoBackup      *bool   `json:"autoBackup,omitempty"`
		NextBackupTime  *string `json:"nextBackupTime,omitempty"`
		RetentionDay    *int32  `json:"retentionDay,omitempty"`
		LastBackupTime  *string `json:"lastBackupTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"provider", "schedule", "accessKeyId", "secretAccessKey", "endpoint", "region", "bucket", "autoBackup", "nextBackupTime", "retentionDay", "lastBackupTime"})
	} else {
		return err
	}
	o.Provider = all.Provider
	o.Schedule = all.Schedule
	o.AccessKeyId = all.AccessKeyId
	o.SecretAccessKey = all.SecretAccessKey
	o.Endpoint = all.Endpoint
	o.Region = all.Region
	o.Bucket = all.Bucket
	o.AutoBackup = all.AutoBackup
	o.NextBackupTime = all.NextBackupTime
	o.RetentionDay = all.RetentionDay
	o.LastBackupTime = all.LastBackupTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
