// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// BasicTask Task information
type BasicTask struct {
	// Task ID
	TaskId common.NullableString `json:"taskId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBasicTask instantiates a new BasicTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBasicTask(taskId common.NullableString) *BasicTask {
	this := BasicTask{}
	this.TaskId = taskId
	return &this
}

// NewBasicTaskWithDefaults instantiates a new BasicTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBasicTaskWithDefaults() *BasicTask {
	this := BasicTask{}
	return &this
}

// GetTaskId returns the TaskId field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *BasicTask) GetTaskId() string {
	if o == nil || o.TaskId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BasicTask) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// SetTaskId sets field value.
func (o *BasicTask) SetTaskId(v string) {
	o.TaskId.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o BasicTask) MarshalJSON() ([]byte, error) {
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
func (o *BasicTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId common.NullableString `json:"taskId"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
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
