// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

type BackupConfig struct {
	Provider        string  `json:"provider"`
	Schedule        *string `json:"schedule,omitempty"`
	AccessKeyId     string  `json:"accessKeyId"`
	SecretAccessKey string  `json:"secretAccessKey"`
	Endpoint        string  `json:"endpoint"`
	Region          *string `json:"region,omitempty"`
	Bucket          string  `json:"bucket"`
	// enable or disable auto backup
	AutoBackup *bool `json:"autoBackup,omitempty"`
	// time for next backup
	NextBackupTime *string `json:"nextBackupTime,omitempty"`
	// specify the retention policy
	RetentionPolicy *string `json:"retentionPolicy,omitempty"`
	// the time of last backup
	LastBackupTime *string `json:"lastBackupTime,omitempty"`
	// Whether the server can connect to the backup storage
	Connected *bool `json:"connected,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupConfig instantiates a new BackupConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupConfig(provider string, accessKeyId string, secretAccessKey string, endpoint string, bucket string) *BackupConfig {
	this := BackupConfig{}
	this.Provider = provider
	this.AccessKeyId = accessKeyId
	this.SecretAccessKey = secretAccessKey
	this.Endpoint = endpoint
	this.Bucket = bucket
	var autoBackup bool = false
	this.AutoBackup = &autoBackup
	return &this
}

// NewBackupConfigWithDefaults instantiates a new BackupConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupConfigWithDefaults() *BackupConfig {
	this := BackupConfig{}
	var autoBackup bool = false
	this.AutoBackup = &autoBackup
	return &this
}

// GetProvider returns the Provider field value.
func (o *BackupConfig) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *BackupConfig) SetProvider(v string) {
	o.Provider = v
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

// GetAccessKeyId returns the AccessKeyId field value.
func (o *BackupConfig) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value.
func (o *BackupConfig) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetSecretAccessKey returns the SecretAccessKey field value.
func (o *BackupConfig) GetSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretAccessKey, true
}

// SetSecretAccessKey sets field value.
func (o *BackupConfig) SetSecretAccessKey(v string) {
	o.SecretAccessKey = v
}

// GetEndpoint returns the Endpoint field value.
func (o *BackupConfig) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value.
func (o *BackupConfig) SetEndpoint(v string) {
	o.Endpoint = v
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

// GetBucket returns the Bucket field value.
func (o *BackupConfig) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value.
func (o *BackupConfig) SetBucket(v string) {
	o.Bucket = v
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

// GetRetentionPolicy returns the RetentionPolicy field value if set, zero value otherwise.
func (o *BackupConfig) GetRetentionPolicy() string {
	if o == nil || o.RetentionPolicy == nil {
		var ret string
		return ret
	}
	return *o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetRetentionPolicyOk() (*string, bool) {
	if o == nil || o.RetentionPolicy == nil {
		return nil, false
	}
	return o.RetentionPolicy, true
}

// HasRetentionPolicy returns a boolean if a field has been set.
func (o *BackupConfig) HasRetentionPolicy() bool {
	return o != nil && o.RetentionPolicy != nil
}

// SetRetentionPolicy gets a reference to the given string and assigns it to the RetentionPolicy field.
func (o *BackupConfig) SetRetentionPolicy(v string) {
	o.RetentionPolicy = &v
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

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *BackupConfig) GetConnected() bool {
	if o == nil || o.Connected == nil {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupConfig) GetConnectedOk() (*bool, bool) {
	if o == nil || o.Connected == nil {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *BackupConfig) HasConnected() bool {
	return o != nil && o.Connected != nil
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *BackupConfig) SetConnected(v bool) {
	o.Connected = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["provider"] = o.Provider
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	toSerialize["accessKeyId"] = o.AccessKeyId
	toSerialize["secretAccessKey"] = o.SecretAccessKey
	toSerialize["endpoint"] = o.Endpoint
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	toSerialize["bucket"] = o.Bucket
	if o.AutoBackup != nil {
		toSerialize["autoBackup"] = o.AutoBackup
	}
	if o.NextBackupTime != nil {
		toSerialize["nextBackupTime"] = o.NextBackupTime
	}
	if o.RetentionPolicy != nil {
		toSerialize["retentionPolicy"] = o.RetentionPolicy
	}
	if o.LastBackupTime != nil {
		toSerialize["lastBackupTime"] = o.LastBackupTime
	}
	if o.Connected != nil {
		toSerialize["connected"] = o.Connected
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Provider        *string `json:"provider"`
		Schedule        *string `json:"schedule,omitempty"`
		AccessKeyId     *string `json:"accessKeyId"`
		SecretAccessKey *string `json:"secretAccessKey"`
		Endpoint        *string `json:"endpoint"`
		Region          *string `json:"region,omitempty"`
		Bucket          *string `json:"bucket"`
		AutoBackup      *bool   `json:"autoBackup,omitempty"`
		NextBackupTime  *string `json:"nextBackupTime,omitempty"`
		RetentionPolicy *string `json:"retentionPolicy,omitempty"`
		LastBackupTime  *string `json:"lastBackupTime,omitempty"`
		Connected       *bool   `json:"connected,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.AccessKeyId == nil {
		return fmt.Errorf("required field accessKeyId missing")
	}
	if all.SecretAccessKey == nil {
		return fmt.Errorf("required field secretAccessKey missing")
	}
	if all.Endpoint == nil {
		return fmt.Errorf("required field endpoint missing")
	}
	if all.Bucket == nil {
		return fmt.Errorf("required field bucket missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"provider", "schedule", "accessKeyId", "secretAccessKey", "endpoint", "region", "bucket", "autoBackup", "nextBackupTime", "retentionPolicy", "lastBackupTime", "connected"})
	} else {
		return err
	}
	o.Provider = *all.Provider
	o.Schedule = all.Schedule
	o.AccessKeyId = *all.AccessKeyId
	o.SecretAccessKey = *all.SecretAccessKey
	o.Endpoint = *all.Endpoint
	o.Region = all.Region
	o.Bucket = *all.Bucket
	o.AutoBackup = all.AutoBackup
	o.NextBackupTime = all.NextBackupTime
	o.RetentionPolicy = all.RetentionPolicy
	o.LastBackupTime = all.LastBackupTime
	o.Connected = all.Connected

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
