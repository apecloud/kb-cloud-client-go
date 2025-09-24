// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TaskSummary struct {
	// Type of task operation
	TaskType *string `json:"taskType,omitempty"`
	// Progress percentage of the task
	Progress *int32 `json:"progress,omitempty"`
	// Detailed message about the task status
	Message *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTaskSummary instantiates a new TaskSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTaskSummary() *TaskSummary {
	this := TaskSummary{}
	return &this
}

// NewTaskSummaryWithDefaults instantiates a new TaskSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTaskSummaryWithDefaults() *TaskSummary {
	this := TaskSummary{}
	return &this
}

// GetTaskType returns the TaskType field value if set, zero value otherwise.
func (o *TaskSummary) GetTaskType() string {
	if o == nil || o.TaskType == nil {
		var ret string
		return ret
	}
	return *o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSummary) GetTaskTypeOk() (*string, bool) {
	if o == nil || o.TaskType == nil {
		return nil, false
	}
	return o.TaskType, true
}

// HasTaskType returns a boolean if a field has been set.
func (o *TaskSummary) HasTaskType() bool {
	return o != nil && o.TaskType != nil
}

// SetTaskType gets a reference to the given string and assigns it to the TaskType field.
func (o *TaskSummary) SetTaskType(v string) {
	o.TaskType = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *TaskSummary) GetProgress() int32 {
	if o == nil || o.Progress == nil {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSummary) GetProgressOk() (*int32, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *TaskSummary) HasProgress() bool {
	return o != nil && o.Progress != nil
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *TaskSummary) SetProgress(v int32) {
	o.Progress = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TaskSummary) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSummary) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TaskSummary) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TaskSummary) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TaskSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TaskType != nil {
		toSerialize["taskType"] = o.TaskType
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TaskSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskType *string `json:"taskType,omitempty"`
		Progress *int32  `json:"progress,omitempty"`
		Message  *string `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskType", "progress", "message"})
	} else {
		return err
	}
	o.TaskType = all.TaskType
	o.Progress = all.Progress
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
