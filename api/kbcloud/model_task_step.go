// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TaskStep struct {
	// Unique identifier for the step
	StepId string `json:"stepId"`
	// Name of the step
	StepName string `json:"stepName"`
	// Method to be executed in this step
	MethodName string `json:"methodName"`
	// Input parameters for the step
	Inputs map[string]interface{} `json:"inputs"`
	// Output parameters from the step
	Outputs map[string]interface{} `json:"outputs,omitempty"`
	// Number of times to retry the step in case of failure
	RetryLimit *int32 `json:"retryLimit,omitempty"`
	// Current retry count for the step
	CurrRetryCount *int32 `json:"currRetryCount,omitempty"`
	// Timeout duration for the step in seconds
	TimeoutSecond *int32 `json:"timeoutSecond,omitempty"`
	// Current status of the task
	Status TaskStatus `json:"status"`
	// Detailed message about the step status
	Message *string `json:"message,omitempty"`
	// Timestamp when the step was created
	CreatedAt time.Time `json:"createdAt"`
	// Timestamp when the step was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// Timestamp when the task was deleted
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Time when the task completed or failed
	CompletionTime *time.Time `json:"completionTime,omitempty"`
	// Detailed information about the step
	Detail map[string]interface{} `json:"detail,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTaskStep instantiates a new TaskStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTaskStep(stepId string, stepName string, methodName string, inputs map[string]interface{}, status TaskStatus, createdAt time.Time, updatedAt time.Time) *TaskStep {
	this := TaskStep{}
	this.StepId = stepId
	this.StepName = stepName
	this.MethodName = methodName
	this.Inputs = inputs
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewTaskStepWithDefaults instantiates a new TaskStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTaskStepWithDefaults() *TaskStep {
	this := TaskStep{}
	return &this
}

// GetStepId returns the StepId field value.
func (o *TaskStep) GetStepId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value
// and a boolean to check if the value has been set.
func (o *TaskStep) GetStepIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepId, true
}

// SetStepId sets field value.
func (o *TaskStep) SetStepId(v string) {
	o.StepId = v
}

// GetStepName returns the StepName field value.
func (o *TaskStep) GetStepName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StepName
}

// GetStepNameOk returns a tuple with the StepName field value
// and a boolean to check if the value has been set.
func (o *TaskStep) GetStepNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepName, true
}

// SetStepName sets field value.
func (o *TaskStep) SetStepName(v string) {
	o.StepName = v
}

// GetMethodName returns the MethodName field value.
func (o *TaskStep) GetMethodName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MethodName
}

// GetMethodNameOk returns a tuple with the MethodName field value
// and a boolean to check if the value has been set.
func (o *TaskStep) GetMethodNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MethodName, true
}

// SetMethodName sets field value.
func (o *TaskStep) SetMethodName(v string) {
	o.MethodName = v
}

// GetInputs returns the Inputs field value.
func (o *TaskStep) GetInputs() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *TaskStep) GetInputsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *TaskStep) SetInputs(v map[string]interface{}) {
	o.Inputs = v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *TaskStep) GetOutputs() map[string]interface{} {
	if o == nil || o.Outputs == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStep) GetOutputsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Outputs == nil {
		return nil, false
	}
	return &o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *TaskStep) HasOutputs() bool {
	return o != nil && o.Outputs != nil
}

// SetOutputs gets a reference to the given map[string]interface{} and assigns it to the Outputs field.
func (o *TaskStep) SetOutputs(v map[string]interface{}) {
	o.Outputs = v
}

// GetRetryLimit returns the RetryLimit field value if set, zero value otherwise.
func (o *TaskStep) GetRetryLimit() int32 {
	if o == nil || o.RetryLimit == nil {
		var ret int32
		return ret
	}
	return *o.RetryLimit
}

// GetRetryLimitOk returns a tuple with the RetryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStep) GetRetryLimitOk() (*int32, bool) {
	if o == nil || o.RetryLimit == nil {
		return nil, false
	}
	return o.RetryLimit, true
}

// HasRetryLimit returns a boolean if a field has been set.
func (o *TaskStep) HasRetryLimit() bool {
	return o != nil && o.RetryLimit != nil
}

// SetRetryLimit gets a reference to the given int32 and assigns it to the RetryLimit field.
func (o *TaskStep) SetRetryLimit(v int32) {
	o.RetryLimit = &v
}

// GetCurrRetryCount returns the CurrRetryCount field value if set, zero value otherwise.
func (o *TaskStep) GetCurrRetryCount() int32 {
	if o == nil || o.CurrRetryCount == nil {
		var ret int32
		return ret
	}
	return *o.CurrRetryCount
}

// GetCurrRetryCountOk returns a tuple with the CurrRetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStep) GetCurrRetryCountOk() (*int32, bool) {
	if o == nil || o.CurrRetryCount == nil {
		return nil, false
	}
	return o.CurrRetryCount, true
}

// HasCurrRetryCount returns a boolean if a field has been set.
func (o *TaskStep) HasCurrRetryCount() bool {
	return o != nil && o.CurrRetryCount != nil
}

// SetCurrRetryCount gets a reference to the given int32 and assigns it to the CurrRetryCount field.
func (o *TaskStep) SetCurrRetryCount(v int32) {
	o.CurrRetryCount = &v
}

// GetTimeoutSecond returns the TimeoutSecond field value if set, zero value otherwise.
func (o *TaskStep) GetTimeoutSecond() int32 {
	if o == nil || o.TimeoutSecond == nil {
		var ret int32
		return ret
	}
	return *o.TimeoutSecond
}

// GetTimeoutSecondOk returns a tuple with the TimeoutSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStep) GetTimeoutSecondOk() (*int32, bool) {
	if o == nil || o.TimeoutSecond == nil {
		return nil, false
	}
	return o.TimeoutSecond, true
}

// HasTimeoutSecond returns a boolean if a field has been set.
func (o *TaskStep) HasTimeoutSecond() bool {
	return o != nil && o.TimeoutSecond != nil
}

// SetTimeoutSecond gets a reference to the given int32 and assigns it to the TimeoutSecond field.
func (o *TaskStep) SetTimeoutSecond(v int32) {
	o.TimeoutSecond = &v
}

// GetStatus returns the Status field value.
func (o *TaskStep) GetStatus() TaskStatus {
	if o == nil {
		var ret TaskStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TaskStep) GetStatusOk() (*TaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *TaskStep) SetStatus(v TaskStatus) {
	o.Status = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TaskStep) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStep) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TaskStep) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TaskStep) SetMessage(v string) {
	o.Message = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *TaskStep) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskStep) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *TaskStep) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *TaskStep) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskStep) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *TaskStep) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *TaskStep) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStep) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *TaskStep) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *TaskStep) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *TaskStep) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStep) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *TaskStep) HasCompletionTime() bool {
	return o != nil && o.CompletionTime != nil
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *TaskStep) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *TaskStep) GetDetail() map[string]interface{} {
	if o == nil || o.Detail == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStep) GetDetailOk() (*map[string]interface{}, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return &o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *TaskStep) HasDetail() bool {
	return o != nil && o.Detail != nil
}

// SetDetail gets a reference to the given map[string]interface{} and assigns it to the Detail field.
func (o *TaskStep) SetDetail(v map[string]interface{}) {
	o.Detail = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TaskStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["stepId"] = o.StepId
	toSerialize["stepName"] = o.StepName
	toSerialize["methodName"] = o.MethodName
	toSerialize["inputs"] = o.Inputs
	if o.Outputs != nil {
		toSerialize["outputs"] = o.Outputs
	}
	if o.RetryLimit != nil {
		toSerialize["retryLimit"] = o.RetryLimit
	}
	if o.CurrRetryCount != nil {
		toSerialize["currRetryCount"] = o.CurrRetryCount
	}
	if o.TimeoutSecond != nil {
		toSerialize["timeoutSecond"] = o.TimeoutSecond
	}
	toSerialize["status"] = o.Status
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.StartedAt != nil {
		if o.StartedAt.Nanosecond() == 0 {
			toSerialize["startedAt"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startedAt"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CompletionTime != nil {
		if o.CompletionTime.Nanosecond() == 0 {
			toSerialize["completionTime"] = o.CompletionTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["completionTime"] = o.CompletionTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TaskStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StepId         *string                 `json:"stepId"`
		StepName       *string                 `json:"stepName"`
		MethodName     *string                 `json:"methodName"`
		Inputs         *map[string]interface{} `json:"inputs"`
		Outputs        map[string]interface{}  `json:"outputs,omitempty"`
		RetryLimit     *int32                  `json:"retryLimit,omitempty"`
		CurrRetryCount *int32                  `json:"currRetryCount,omitempty"`
		TimeoutSecond  *int32                  `json:"timeoutSecond,omitempty"`
		Status         *TaskStatus             `json:"status"`
		Message        *string                 `json:"message,omitempty"`
		CreatedAt      *time.Time              `json:"createdAt"`
		UpdatedAt      *time.Time              `json:"updatedAt"`
		StartedAt      *time.Time              `json:"startedAt,omitempty"`
		CompletionTime *time.Time              `json:"completionTime,omitempty"`
		Detail         map[string]interface{}  `json:"detail,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.StepId == nil {
		return fmt.Errorf("required field stepId missing")
	}
	if all.StepName == nil {
		return fmt.Errorf("required field stepName missing")
	}
	if all.MethodName == nil {
		return fmt.Errorf("required field methodName missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"stepId", "stepName", "methodName", "inputs", "outputs", "retryLimit", "currRetryCount", "timeoutSecond", "status", "message", "createdAt", "updatedAt", "startedAt", "completionTime", "detail"})
	} else {
		return err
	}

	hasInvalidField := false
	o.StepId = *all.StepId
	o.StepName = *all.StepName
	o.MethodName = *all.MethodName
	o.Inputs = *all.Inputs
	o.Outputs = all.Outputs
	o.RetryLimit = all.RetryLimit
	o.CurrRetryCount = all.CurrRetryCount
	o.TimeoutSecond = all.TimeoutSecond
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Message = all.Message
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt
	o.StartedAt = all.StartedAt
	o.CompletionTime = all.CompletionTime
	o.Detail = all.Detail

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
