// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"
)

type Task struct {
	// ID of the scaling task
	TaskId string `json:"taskId"`
	// Type of scaling operation
	Operation TaskOperation `json:"operation"`
	// Current status of the task
	Status TaskStatus `json:"status"`
	// Progress percentage of the task
	Progress *int32 `json:"progress,omitempty"`
	// Time when the task started
	StartTime *time.Time `json:"startTime,omitempty"`
	// Time when the task completed or failed
	CompletionTime *time.Time `json:"completionTime,omitempty"`
	// Detailed message about the task status
	Message *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTask instantiates a new Task object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTask(taskId string, operation TaskOperation, status TaskStatus) *Task {
	this := Task{}
	this.TaskId = taskId
	this.Operation = operation
	this.Status = status
	return &this
}

// NewTaskWithDefaults instantiates a new Task object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetTaskId returns the TaskId field value.
func (o *Task) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *Task) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *Task) SetTaskId(v string) {
	o.TaskId = v
}

// GetOperation returns the Operation field value.
func (o *Task) GetOperation() TaskOperation {
	if o == nil {
		var ret TaskOperation
		return ret
	}
	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *Task) GetOperationOk() (*TaskOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value.
func (o *Task) SetOperation(v TaskOperation) {
	o.Operation = v
}

// GetStatus returns the Status field value.
func (o *Task) GetStatus() TaskStatus {
	if o == nil {
		var ret TaskStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Task) GetStatusOk() (*TaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *Task) SetStatus(v TaskStatus) {
	o.Status = v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *Task) GetProgress() int32 {
	if o == nil || o.Progress == nil {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetProgressOk() (*int32, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *Task) HasProgress() bool {
	return o != nil && o.Progress != nil
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *Task) SetProgress(v int32) {
	o.Progress = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Task) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Task) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *Task) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *Task) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *Task) HasCompletionTime() bool {
	return o != nil && o.CompletionTime != nil
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *Task) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Task) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Task) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Task) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId
	toSerialize["operation"] = o.Operation
	toSerialize["status"] = o.Status
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.StartTime != nil {
		if o.StartTime.Nanosecond() == 0 {
			toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CompletionTime != nil {
		if o.CompletionTime.Nanosecond() == 0 {
			toSerialize["completionTime"] = o.CompletionTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["completionTime"] = o.CompletionTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *Task) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId         *string        `json:"taskId"`
		Operation      *TaskOperation `json:"operation"`
		Status         *TaskStatus    `json:"status"`
		Progress       *int32         `json:"progress,omitempty"`
		StartTime      *time.Time     `json:"startTime,omitempty"`
		CompletionTime *time.Time     `json:"completionTime,omitempty"`
		Message        *string        `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	if all.Operation == nil {
		return fmt.Errorf("required field operation missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskId", "operation", "status", "progress", "startTime", "completionTime", "message"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TaskId = *all.TaskId
	if !all.Operation.IsValid() {
		hasInvalidField = true
	} else {
		o.Operation = *all.Operation
	}
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Progress = all.Progress
	o.StartTime = all.StartTime
	o.CompletionTime = all.CompletionTime
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
