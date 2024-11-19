// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// DisasterRecoveryTask Task information of disasterRecovery job
type DisasterRecoveryTask struct {
	// Task ID
	TaskId common.NullableInt32 `json:"taskId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDisasterRecoveryTask instantiates a new DisasterRecoveryTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDisasterRecoveryTask(taskId common.NullableInt32) *DisasterRecoveryTask {
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

// GetTaskId returns the TaskId field value.
// If the value is explicit nil, the zero value for int32 will be returned.
func (o *DisasterRecoveryTask) GetTaskId() int32 {
	if o == nil || o.TaskId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryTask) GetTaskIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// SetTaskId sets field value.
func (o *DisasterRecoveryTask) SetTaskId(v int32) {
	o.TaskId.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o DisasterRecoveryTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DisasterRecoveryTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId common.NullableInt32 `json:"taskId"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.TaskId.IsSet() {
		return fmt.Errorf("required field taskId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskId"})
	} else {
		return err
	}
	o.TaskId = all.TaskId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
