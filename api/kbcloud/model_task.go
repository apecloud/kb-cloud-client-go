// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Task struct {
	// ID of the scaling task
	TaskId string `json:"taskId"`
	// Name of the task
	TaskName string `json:"taskName"`
	// Type of task operation
	TaskType string `json:"taskType"`
	// Current status of the task
	Status TaskStatus `json:"status"`
	// Timestamp when the task was created
	CreatedAt time.Time `json:"createdAt"`
	// Timestamp when the task was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// Timestamp when the task was deleted
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// Timestamp when the task was deleted
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Time when the task completed or failed
	CompletionTime *time.Time `json:"completionTime,omitempty"`
	// Detailed message about the task status
	Message *string `json:"message,omitempty"`
	// Progress percentage of the task
	Progress *int32     `json:"progress,omitempty"`
	Steps    []TaskStep `json:"steps,omitempty"`
	// Degree of parallelism for the task
	Parallelism *int32 `json:"parallelism,omitempty"`
	// Policy to handle failures
	FailurePolicy *TaskFailurePolicy `json:"failurePolicy,omitempty"`
	// Maximum number of retries for the task
	RetryLimit *int32 `json:"retryLimit,omitempty"`
	// Timeout duration for the task in seconds
	TimeoutSecond *int32 `json:"timeoutSecond,omitempty"`
	// The user created the task
	Operator *string `json:"operator,omitempty"`
	// Resource type of the task, e.g. "cluster", "environment", etc.
	ResourceType *string `json:"resourceType,omitempty"`
	// Resource ID of the task
	ResourceId *string `json:"resourceId,omitempty"`
	// Resource name of the task
	ResourceName *string `json:"resourceName,omitempty"`
	// Organization name of the task
	OrgName   *string       `json:"orgName,omitempty"`
	Scheduled *TaskSchedule `json:"scheduled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTask instantiates a new Task object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTask(taskId string, taskName string, taskType string, status TaskStatus, createdAt time.Time, updatedAt time.Time) *Task {
	this := Task{}
	this.TaskId = taskId
	this.TaskName = taskName
	this.TaskType = taskType
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
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

// GetTaskName returns the TaskName field value.
func (o *Task) GetTaskName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value
// and a boolean to check if the value has been set.
func (o *Task) GetTaskNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskName, true
}

// SetTaskName sets field value.
func (o *Task) SetTaskName(v string) {
	o.TaskName = v
}

// GetTaskType returns the TaskType field value.
func (o *Task) GetTaskType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value
// and a boolean to check if the value has been set.
func (o *Task) GetTaskTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskType, true
}

// SetTaskType sets field value.
func (o *Task) SetTaskType(v string) {
	o.TaskType = v
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

// GetCreatedAt returns the CreatedAt field value.
func (o *Task) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Task) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Task) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *Task) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Task) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *Task) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Task) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Task) HasDeletedAt() bool {
	return o != nil && o.DeletedAt != nil
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Task) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *Task) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *Task) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *Task) SetStartedAt(v time.Time) {
	o.StartedAt = &v
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

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *Task) GetSteps() []TaskStep {
	if o == nil || o.Steps == nil {
		var ret []TaskStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetStepsOk() (*[]TaskStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *Task) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []TaskStep and assigns it to the Steps field.
func (o *Task) SetSteps(v []TaskStep) {
	o.Steps = v
}

// GetParallelism returns the Parallelism field value if set, zero value otherwise.
func (o *Task) GetParallelism() int32 {
	if o == nil || o.Parallelism == nil {
		var ret int32
		return ret
	}
	return *o.Parallelism
}

// GetParallelismOk returns a tuple with the Parallelism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetParallelismOk() (*int32, bool) {
	if o == nil || o.Parallelism == nil {
		return nil, false
	}
	return o.Parallelism, true
}

// HasParallelism returns a boolean if a field has been set.
func (o *Task) HasParallelism() bool {
	return o != nil && o.Parallelism != nil
}

// SetParallelism gets a reference to the given int32 and assigns it to the Parallelism field.
func (o *Task) SetParallelism(v int32) {
	o.Parallelism = &v
}

// GetFailurePolicy returns the FailurePolicy field value if set, zero value otherwise.
func (o *Task) GetFailurePolicy() TaskFailurePolicy {
	if o == nil || o.FailurePolicy == nil {
		var ret TaskFailurePolicy
		return ret
	}
	return *o.FailurePolicy
}

// GetFailurePolicyOk returns a tuple with the FailurePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetFailurePolicyOk() (*TaskFailurePolicy, bool) {
	if o == nil || o.FailurePolicy == nil {
		return nil, false
	}
	return o.FailurePolicy, true
}

// HasFailurePolicy returns a boolean if a field has been set.
func (o *Task) HasFailurePolicy() bool {
	return o != nil && o.FailurePolicy != nil
}

// SetFailurePolicy gets a reference to the given TaskFailurePolicy and assigns it to the FailurePolicy field.
func (o *Task) SetFailurePolicy(v TaskFailurePolicy) {
	o.FailurePolicy = &v
}

// GetRetryLimit returns the RetryLimit field value if set, zero value otherwise.
func (o *Task) GetRetryLimit() int32 {
	if o == nil || o.RetryLimit == nil {
		var ret int32
		return ret
	}
	return *o.RetryLimit
}

// GetRetryLimitOk returns a tuple with the RetryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetRetryLimitOk() (*int32, bool) {
	if o == nil || o.RetryLimit == nil {
		return nil, false
	}
	return o.RetryLimit, true
}

// HasRetryLimit returns a boolean if a field has been set.
func (o *Task) HasRetryLimit() bool {
	return o != nil && o.RetryLimit != nil
}

// SetRetryLimit gets a reference to the given int32 and assigns it to the RetryLimit field.
func (o *Task) SetRetryLimit(v int32) {
	o.RetryLimit = &v
}

// GetTimeoutSecond returns the TimeoutSecond field value if set, zero value otherwise.
func (o *Task) GetTimeoutSecond() int32 {
	if o == nil || o.TimeoutSecond == nil {
		var ret int32
		return ret
	}
	return *o.TimeoutSecond
}

// GetTimeoutSecondOk returns a tuple with the TimeoutSecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTimeoutSecondOk() (*int32, bool) {
	if o == nil || o.TimeoutSecond == nil {
		return nil, false
	}
	return o.TimeoutSecond, true
}

// HasTimeoutSecond returns a boolean if a field has been set.
func (o *Task) HasTimeoutSecond() bool {
	return o != nil && o.TimeoutSecond != nil
}

// SetTimeoutSecond gets a reference to the given int32 and assigns it to the TimeoutSecond field.
func (o *Task) SetTimeoutSecond(v int32) {
	o.TimeoutSecond = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *Task) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *Task) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *Task) SetOperator(v string) {
	o.Operator = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *Task) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *Task) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *Task) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *Task) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *Task) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *Task) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *Task) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *Task) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *Task) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Task) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Task) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Task) SetOrgName(v string) {
	o.OrgName = &v
}

// GetScheduled returns the Scheduled field value if set, zero value otherwise.
func (o *Task) GetScheduled() TaskSchedule {
	if o == nil || o.Scheduled == nil {
		var ret TaskSchedule
		return ret
	}
	return *o.Scheduled
}

// GetScheduledOk returns a tuple with the Scheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetScheduledOk() (*TaskSchedule, bool) {
	if o == nil || o.Scheduled == nil {
		return nil, false
	}
	return o.Scheduled, true
}

// HasScheduled returns a boolean if a field has been set.
func (o *Task) HasScheduled() bool {
	return o != nil && o.Scheduled != nil
}

// SetScheduled gets a reference to the given TaskSchedule and assigns it to the Scheduled field.
func (o *Task) SetScheduled(v TaskSchedule) {
	o.Scheduled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId
	toSerialize["taskName"] = o.TaskName
	toSerialize["taskType"] = o.TaskType
	toSerialize["status"] = o.Status
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
	if o.DeletedAt != nil {
		if o.DeletedAt.Nanosecond() == 0 {
			toSerialize["deletedAt"] = o.DeletedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["deletedAt"] = o.DeletedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.Parallelism != nil {
		toSerialize["parallelism"] = o.Parallelism
	}
	if o.FailurePolicy != nil {
		toSerialize["failurePolicy"] = o.FailurePolicy
	}
	if o.RetryLimit != nil {
		toSerialize["retryLimit"] = o.RetryLimit
	}
	if o.TimeoutSecond != nil {
		toSerialize["timeoutSecond"] = o.TimeoutSecond
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.ResourceName != nil {
		toSerialize["resourceName"] = o.ResourceName
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.Scheduled != nil {
		toSerialize["scheduled"] = o.Scheduled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Task) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId         *string            `json:"taskId"`
		TaskName       *string            `json:"taskName"`
		TaskType       *string            `json:"taskType"`
		Status         *TaskStatus        `json:"status"`
		CreatedAt      *time.Time         `json:"createdAt"`
		UpdatedAt      *time.Time         `json:"updatedAt"`
		DeletedAt      *time.Time         `json:"deletedAt,omitempty"`
		StartedAt      *time.Time         `json:"startedAt,omitempty"`
		CompletionTime *time.Time         `json:"completionTime,omitempty"`
		Message        *string            `json:"message,omitempty"`
		Progress       *int32             `json:"progress,omitempty"`
		Steps          []TaskStep         `json:"steps,omitempty"`
		Parallelism    *int32             `json:"parallelism,omitempty"`
		FailurePolicy  *TaskFailurePolicy `json:"failurePolicy,omitempty"`
		RetryLimit     *int32             `json:"retryLimit,omitempty"`
		TimeoutSecond  *int32             `json:"timeoutSecond,omitempty"`
		Operator       *string            `json:"operator,omitempty"`
		ResourceType   *string            `json:"resourceType,omitempty"`
		ResourceId     *string            `json:"resourceId,omitempty"`
		ResourceName   *string            `json:"resourceName,omitempty"`
		OrgName        *string            `json:"orgName,omitempty"`
		Scheduled      *TaskSchedule      `json:"scheduled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	if all.TaskName == nil {
		return fmt.Errorf("required field taskName missing")
	}
	if all.TaskType == nil {
		return fmt.Errorf("required field taskType missing")
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
		common.DeleteKeys(additionalProperties, &[]string{"taskId", "taskName", "taskType", "status", "createdAt", "updatedAt", "deletedAt", "startedAt", "completionTime", "message", "progress", "steps", "parallelism", "failurePolicy", "retryLimit", "timeoutSecond", "operator", "resourceType", "resourceId", "resourceName", "orgName", "scheduled"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TaskId = *all.TaskId
	o.TaskName = *all.TaskName
	o.TaskType = *all.TaskType
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt
	o.DeletedAt = all.DeletedAt
	o.StartedAt = all.StartedAt
	o.CompletionTime = all.CompletionTime
	o.Message = all.Message
	o.Progress = all.Progress
	o.Steps = all.Steps
	o.Parallelism = all.Parallelism
	if all.FailurePolicy != nil && !all.FailurePolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.FailurePolicy = all.FailurePolicy
	}
	o.RetryLimit = all.RetryLimit
	o.TimeoutSecond = all.TimeoutSecond
	o.Operator = all.Operator
	o.ResourceType = all.ResourceType
	o.ResourceId = all.ResourceId
	o.ResourceName = all.ResourceName
	o.OrgName = all.OrgName
	if all.Scheduled != nil && all.Scheduled.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scheduled = all.Scheduled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
