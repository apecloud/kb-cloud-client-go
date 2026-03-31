// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BlueGreenDeploymentTask Task information of a blue-green deployment job.
type BlueGreenDeploymentTask struct {
	// Task ID.
	TaskId string `json:"taskId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBlueGreenDeploymentTask instantiates a new BlueGreenDeploymentTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBlueGreenDeploymentTask(taskId string) *BlueGreenDeploymentTask {
	this := BlueGreenDeploymentTask{}
	this.TaskId = taskId
	return &this
}

// NewBlueGreenDeploymentTaskWithDefaults instantiates a new BlueGreenDeploymentTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBlueGreenDeploymentTaskWithDefaults() *BlueGreenDeploymentTask {
	this := BlueGreenDeploymentTask{}
	return &this
}

// GetTaskId returns the TaskId field value.
func (o *BlueGreenDeploymentTask) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *BlueGreenDeploymentTask) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *BlueGreenDeploymentTask) SetTaskId(v string) {
	o.TaskId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BlueGreenDeploymentTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BlueGreenDeploymentTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId *string `json:"taskId"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskId"})
	} else {
		return err
	}
	o.TaskId = *all.TaskId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
