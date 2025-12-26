// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsHScale OpsHScale is the payload to horizontally scale a KubeBlocks cluster. It requires specifying either the number of replicas or the number of shards.
type OpsHScale struct {
	// component type
	Component string `json:"component"`
	// name of the backup to restore from
	BackupName common.NullableString `json:"backupName,omitempty"`
	// number of replicas
	Replicas common.NullableInt32 `json:"replicas,omitempty"`
	// number of shards, mutually exclusive with replicas.
	Shards common.NullableInt32 `json:"shards,omitempty"`
	// Specifies the maximum time in seconds that the OpsRequest will wait for its pre-conditions to be met before it aborts the operation
	PreConditionDeadlineSeconds common.NullableInt32 `json:"preConditionDeadlineSeconds,omitempty"`
	Schedule                    *TaskSchedule        `json:"schedule,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsHScale instantiates a new OpsHScale object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsHScale(component string) *OpsHScale {
	this := OpsHScale{}
	this.Component = component
	return &this
}

// NewOpsHScaleWithDefaults instantiates a new OpsHScale object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsHScaleWithDefaults() *OpsHScale {
	this := OpsHScale{}
	return &this
}

// GetComponent returns the Component field value.
func (o *OpsHScale) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OpsHScale) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *OpsHScale) SetComponent(v string) {
	o.Component = v
}

// GetBackupName returns the BackupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpsHScale) GetBackupName() string {
	if o == nil || o.BackupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BackupName.Get()
}

// GetBackupNameOk returns a tuple with the BackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OpsHScale) GetBackupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupName.Get(), o.BackupName.IsSet()
}

// HasBackupName returns a boolean if a field has been set.
func (o *OpsHScale) HasBackupName() bool {
	return o != nil && o.BackupName.IsSet()
}

// SetBackupName gets a reference to the given common.NullableString and assigns it to the BackupName field.
func (o *OpsHScale) SetBackupName(v string) {
	o.BackupName.Set(&v)
}

// SetBackupNameNil sets the value for BackupName to be an explicit nil.
func (o *OpsHScale) SetBackupNameNil() {
	o.BackupName.Set(nil)
}

// UnsetBackupName ensures that no value is present for BackupName, not even an explicit nil.
func (o *OpsHScale) UnsetBackupName() {
	o.BackupName.Unset()
}

// GetReplicas returns the Replicas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpsHScale) GetReplicas() int32 {
	if o == nil || o.Replicas.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Replicas.Get()
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OpsHScale) GetReplicasOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Replicas.Get(), o.Replicas.IsSet()
}

// HasReplicas returns a boolean if a field has been set.
func (o *OpsHScale) HasReplicas() bool {
	return o != nil && o.Replicas.IsSet()
}

// SetReplicas gets a reference to the given common.NullableInt32 and assigns it to the Replicas field.
func (o *OpsHScale) SetReplicas(v int32) {
	o.Replicas.Set(&v)
}

// SetReplicasNil sets the value for Replicas to be an explicit nil.
func (o *OpsHScale) SetReplicasNil() {
	o.Replicas.Set(nil)
}

// UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil.
func (o *OpsHScale) UnsetReplicas() {
	o.Replicas.Unset()
}

// GetShards returns the Shards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpsHScale) GetShards() int32 {
	if o == nil || o.Shards.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Shards.Get()
}

// GetShardsOk returns a tuple with the Shards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OpsHScale) GetShardsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shards.Get(), o.Shards.IsSet()
}

// HasShards returns a boolean if a field has been set.
func (o *OpsHScale) HasShards() bool {
	return o != nil && o.Shards.IsSet()
}

// SetShards gets a reference to the given common.NullableInt32 and assigns it to the Shards field.
func (o *OpsHScale) SetShards(v int32) {
	o.Shards.Set(&v)
}

// SetShardsNil sets the value for Shards to be an explicit nil.
func (o *OpsHScale) SetShardsNil() {
	o.Shards.Set(nil)
}

// UnsetShards ensures that no value is present for Shards, not even an explicit nil.
func (o *OpsHScale) UnsetShards() {
	o.Shards.Unset()
}

// GetPreConditionDeadlineSeconds returns the PreConditionDeadlineSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpsHScale) GetPreConditionDeadlineSeconds() int32 {
	if o == nil || o.PreConditionDeadlineSeconds.Get() == nil {
		var ret int32
		return ret
	}
	return *o.PreConditionDeadlineSeconds.Get()
}

// GetPreConditionDeadlineSecondsOk returns a tuple with the PreConditionDeadlineSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OpsHScale) GetPreConditionDeadlineSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreConditionDeadlineSeconds.Get(), o.PreConditionDeadlineSeconds.IsSet()
}

// HasPreConditionDeadlineSeconds returns a boolean if a field has been set.
func (o *OpsHScale) HasPreConditionDeadlineSeconds() bool {
	return o != nil && o.PreConditionDeadlineSeconds.IsSet()
}

// SetPreConditionDeadlineSeconds gets a reference to the given common.NullableInt32 and assigns it to the PreConditionDeadlineSeconds field.
func (o *OpsHScale) SetPreConditionDeadlineSeconds(v int32) {
	o.PreConditionDeadlineSeconds.Set(&v)
}

// SetPreConditionDeadlineSecondsNil sets the value for PreConditionDeadlineSeconds to be an explicit nil.
func (o *OpsHScale) SetPreConditionDeadlineSecondsNil() {
	o.PreConditionDeadlineSeconds.Set(nil)
}

// UnsetPreConditionDeadlineSeconds ensures that no value is present for PreConditionDeadlineSeconds, not even an explicit nil.
func (o *OpsHScale) UnsetPreConditionDeadlineSeconds() {
	o.PreConditionDeadlineSeconds.Unset()
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *OpsHScale) GetSchedule() TaskSchedule {
	if o == nil || o.Schedule == nil {
		var ret TaskSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsHScale) GetScheduleOk() (*TaskSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *OpsHScale) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given TaskSchedule and assigns it to the Schedule field.
func (o *OpsHScale) SetSchedule(v TaskSchedule) {
	o.Schedule = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsHScale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	if o.BackupName.IsSet() {
		toSerialize["backupName"] = o.BackupName.Get()
	}
	if o.Replicas.IsSet() {
		toSerialize["replicas"] = o.Replicas.Get()
	}
	if o.Shards.IsSet() {
		toSerialize["shards"] = o.Shards.Get()
	}
	if o.PreConditionDeadlineSeconds.IsSet() {
		toSerialize["preConditionDeadlineSeconds"] = o.PreConditionDeadlineSeconds.Get()
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsHScale) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component                   *string               `json:"component"`
		BackupName                  common.NullableString `json:"backupName,omitempty"`
		Replicas                    common.NullableInt32  `json:"replicas,omitempty"`
		Shards                      common.NullableInt32  `json:"shards,omitempty"`
		PreConditionDeadlineSeconds common.NullableInt32  `json:"preConditionDeadlineSeconds,omitempty"`
		Schedule                    *TaskSchedule         `json:"schedule,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "backupName", "replicas", "shards", "preConditionDeadlineSeconds", "schedule"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Component = *all.Component
	o.BackupName = all.BackupName
	o.Replicas = all.Replicas
	o.Shards = all.Shards
	o.PreConditionDeadlineSeconds = all.PreConditionDeadlineSeconds
	if all.Schedule != nil && all.Schedule.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schedule = all.Schedule

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
