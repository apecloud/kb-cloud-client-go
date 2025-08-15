// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DisasterRecoveryTask Task information of disasterRecovery job
type DisasterRecoveryTask struct {
	// Logical instance ID
	LogicalInstanceId common.NullableString `json:"logicalInstanceID,omitempty"`
	// Task ID
	TaskId common.NullableString `json:"taskId"`
	// Cluster task ID
	ClusterTaskId common.NullableString `json:"clusterTaskId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDisasterRecoveryTask instantiates a new DisasterRecoveryTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDisasterRecoveryTask(taskId common.NullableString) *DisasterRecoveryTask {
	this := DisasterRecoveryTask{}
	this.TaskId = taskId
	return &this
}

// NewDisasterRecoveryTaskWithDefaults instantiates a new DisasterRecoveryTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDisasterRecoveryTaskWithDefaults() *DisasterRecoveryTask {
	this := DisasterRecoveryTask{}
	return &this
}

// GetLogicalInstanceId returns the LogicalInstanceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryTask) GetLogicalInstanceId() string {
	if o == nil || o.LogicalInstanceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogicalInstanceId.Get()
}

// GetLogicalInstanceIdOk returns a tuple with the LogicalInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryTask) GetLogicalInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogicalInstanceId.Get(), o.LogicalInstanceId.IsSet()
}

// HasLogicalInstanceId returns a boolean if a field has been set.
func (o *DisasterRecoveryTask) HasLogicalInstanceId() bool {
	return o != nil && o.LogicalInstanceId.IsSet()
}

// SetLogicalInstanceId gets a reference to the given common.NullableString and assigns it to the LogicalInstanceId field.
func (o *DisasterRecoveryTask) SetLogicalInstanceId(v string) {
	o.LogicalInstanceId.Set(&v)
}

// SetLogicalInstanceIdNil sets the value for LogicalInstanceId to be an explicit nil.
func (o *DisasterRecoveryTask) SetLogicalInstanceIdNil() {
	o.LogicalInstanceId.Set(nil)
}

// UnsetLogicalInstanceId ensures that no value is present for LogicalInstanceId, not even an explicit nil.
func (o *DisasterRecoveryTask) UnsetLogicalInstanceId() {
	o.LogicalInstanceId.Unset()
}

// GetTaskId returns the TaskId field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *DisasterRecoveryTask) GetTaskId() string {
	if o == nil || o.TaskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryTask) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// SetTaskId sets field value.
func (o *DisasterRecoveryTask) SetTaskId(v string) {
	o.TaskId.Set(&v)
}

// GetClusterTaskId returns the ClusterTaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryTask) GetClusterTaskId() string {
	if o == nil || o.ClusterTaskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterTaskId.Get()
}

// GetClusterTaskIdOk returns a tuple with the ClusterTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryTask) GetClusterTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterTaskId.Get(), o.ClusterTaskId.IsSet()
}

// HasClusterTaskId returns a boolean if a field has been set.
func (o *DisasterRecoveryTask) HasClusterTaskId() bool {
	return o != nil && o.ClusterTaskId.IsSet()
}

// SetClusterTaskId gets a reference to the given common.NullableString and assigns it to the ClusterTaskId field.
func (o *DisasterRecoveryTask) SetClusterTaskId(v string) {
	o.ClusterTaskId.Set(&v)
}

// SetClusterTaskIdNil sets the value for ClusterTaskId to be an explicit nil.
func (o *DisasterRecoveryTask) SetClusterTaskIdNil() {
	o.ClusterTaskId.Set(nil)
}

// UnsetClusterTaskId ensures that no value is present for ClusterTaskId, not even an explicit nil.
func (o *DisasterRecoveryTask) UnsetClusterTaskId() {
	o.ClusterTaskId.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DisasterRecoveryTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.LogicalInstanceId.IsSet() {
		toSerialize["logicalInstanceID"] = o.LogicalInstanceId.Get()
	}
	toSerialize["taskId"] = o.TaskId.Get()
	if o.ClusterTaskId.IsSet() {
		toSerialize["clusterTaskId"] = o.ClusterTaskId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DisasterRecoveryTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LogicalInstanceId common.NullableString `json:"logicalInstanceID,omitempty"`
		TaskId            common.NullableString `json:"taskId"`
		ClusterTaskId     common.NullableString `json:"clusterTaskId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if !all.TaskId.IsSet() {
		return fmt.Errorf("required field taskId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"logicalInstanceID", "taskId", "clusterTaskId"})
	} else {
		return err
	}
	o.LogicalInstanceId = all.LogicalInstanceId
	o.TaskId = all.TaskId
	o.ClusterTaskId = all.ClusterTaskId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
