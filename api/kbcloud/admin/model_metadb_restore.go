// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Metadb_restore struct {
	// the namespace of new cluster, default value is original namespace
	Namespace *string `json:"namespace,omitempty"`
	// the name of backup
	BackupName *string `json:"backupName,omitempty"`
	// the name of new cluster
	ClusterName string `json:"clusterName"`
	// determine backup files based on timestamp
	Timestamp *string `json:"timestamp,omitempty"`
	// the duration time of restore execution
	Duration *string `json:"duration,omitempty"`
	// completion time
	CompletionTimestamp *time.Time `json:"completionTimestamp,omitempty"`
	// start time
	StartTimestamp *time.Time `json:"startTimestamp,omitempty"`
	// NODESCRIPTION Cpu
	Cpu *float64 `json:"cpu,omitempty"`
	// NODESCRIPTION Memory
	Memory *float64 `json:"memory,omitempty"`
	// NODESCRIPTION Storage
	Storage *float64 `json:"storage,omitempty"`
	// the number of postgresql pods
	Replicas *int32 `json:"replicas,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetadb_restore instantiates a new Metadb_restore object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetadb_restore(clusterName string) *Metadb_restore {
	this := Metadb_restore{}
	this.ClusterName = clusterName
	return &this
}

// NewMetadb_restoreWithDefaults instantiates a new Metadb_restore object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetadb_restoreWithDefaults() *Metadb_restore {
	this := Metadb_restore{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Metadb_restore) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Metadb_restore) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Metadb_restore) SetNamespace(v string) {
	o.Namespace = &v
}

// GetBackupName returns the BackupName field value if set, zero value otherwise.
func (o *Metadb_restore) GetBackupName() string {
	if o == nil || o.BackupName == nil {
		var ret string
		return ret
	}
	return *o.BackupName
}

// GetBackupNameOk returns a tuple with the BackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetBackupNameOk() (*string, bool) {
	if o == nil || o.BackupName == nil {
		return nil, false
	}
	return o.BackupName, true
}

// HasBackupName returns a boolean if a field has been set.
func (o *Metadb_restore) HasBackupName() bool {
	return o != nil && o.BackupName != nil
}

// SetBackupName gets a reference to the given string and assigns it to the BackupName field.
func (o *Metadb_restore) SetBackupName(v string) {
	o.BackupName = &v
}

// GetClusterName returns the ClusterName field value.
func (o *Metadb_restore) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value.
func (o *Metadb_restore) SetClusterName(v string) {
	o.ClusterName = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Metadb_restore) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Metadb_restore) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *Metadb_restore) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Metadb_restore) GetDuration() string {
	if o == nil || o.Duration == nil {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetDurationOk() (*string, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Metadb_restore) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *Metadb_restore) SetDuration(v string) {
	o.Duration = &v
}

// GetCompletionTimestamp returns the CompletionTimestamp field value if set, zero value otherwise.
func (o *Metadb_restore) GetCompletionTimestamp() time.Time {
	if o == nil || o.CompletionTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTimestamp
}

// GetCompletionTimestampOk returns a tuple with the CompletionTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetCompletionTimestampOk() (*time.Time, bool) {
	if o == nil || o.CompletionTimestamp == nil {
		return nil, false
	}
	return o.CompletionTimestamp, true
}

// HasCompletionTimestamp returns a boolean if a field has been set.
func (o *Metadb_restore) HasCompletionTimestamp() bool {
	return o != nil && o.CompletionTimestamp != nil
}

// SetCompletionTimestamp gets a reference to the given time.Time and assigns it to the CompletionTimestamp field.
func (o *Metadb_restore) SetCompletionTimestamp(v time.Time) {
	o.CompletionTimestamp = &v
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *Metadb_restore) GetStartTimestamp() time.Time {
	if o == nil || o.StartTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetStartTimestampOk() (*time.Time, bool) {
	if o == nil || o.StartTimestamp == nil {
		return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *Metadb_restore) HasStartTimestamp() bool {
	return o != nil && o.StartTimestamp != nil
}

// SetStartTimestamp gets a reference to the given time.Time and assigns it to the StartTimestamp field.
func (o *Metadb_restore) SetStartTimestamp(v time.Time) {
	o.StartTimestamp = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Metadb_restore) GetCpu() float64 {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetCpuOk() (*float64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Metadb_restore) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given float64 and assigns it to the Cpu field.
func (o *Metadb_restore) SetCpu(v float64) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Metadb_restore) GetMemory() float64 {
	if o == nil || o.Memory == nil {
		var ret float64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetMemoryOk() (*float64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Metadb_restore) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given float64 and assigns it to the Memory field.
func (o *Metadb_restore) SetMemory(v float64) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *Metadb_restore) GetStorage() float64 {
	if o == nil || o.Storage == nil {
		var ret float64
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetStorageOk() (*float64, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *Metadb_restore) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given float64 and assigns it to the Storage field.
func (o *Metadb_restore) SetStorage(v float64) {
	o.Storage = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *Metadb_restore) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Metadb_restore) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *Metadb_restore) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *Metadb_restore) SetReplicas(v int32) {
	o.Replicas = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Metadb_restore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.BackupName != nil {
		toSerialize["backupName"] = o.BackupName
	}
	toSerialize["clusterName"] = o.ClusterName
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.CompletionTimestamp != nil {
		if o.CompletionTimestamp.Nanosecond() == 0 {
			toSerialize["completionTimestamp"] = o.CompletionTimestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["completionTimestamp"] = o.CompletionTimestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.StartTimestamp != nil {
		if o.StartTimestamp.Nanosecond() == 0 {
			toSerialize["startTimestamp"] = o.StartTimestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startTimestamp"] = o.StartTimestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Metadb_restore) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Namespace           *string    `json:"namespace,omitempty"`
		BackupName          *string    `json:"backupName,omitempty"`
		ClusterName         *string    `json:"clusterName"`
		Timestamp           *string    `json:"timestamp,omitempty"`
		Duration            *string    `json:"duration,omitempty"`
		CompletionTimestamp *time.Time `json:"completionTimestamp,omitempty"`
		StartTimestamp      *time.Time `json:"startTimestamp,omitempty"`
		Cpu                 *float64   `json:"cpu,omitempty"`
		Memory              *float64   `json:"memory,omitempty"`
		Storage             *float64   `json:"storage,omitempty"`
		Replicas            *int32     `json:"replicas,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClusterName == nil {
		return fmt.Errorf("required field clusterName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"namespace", "backupName", "clusterName", "timestamp", "duration", "completionTimestamp", "startTimestamp", "cpu", "memory", "storage", "replicas"})
	} else {
		return err
	}
	o.Namespace = all.Namespace
	o.BackupName = all.BackupName
	o.ClusterName = *all.ClusterName
	o.Timestamp = all.Timestamp
	o.Duration = all.Duration
	o.CompletionTimestamp = all.CompletionTimestamp
	o.StartTimestamp = all.StartTimestamp
	o.Cpu = all.Cpu
	o.Memory = all.Memory
	o.Storage = all.Storage
	o.Replicas = all.Replicas

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
