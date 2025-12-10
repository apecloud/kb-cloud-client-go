// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TaskListItem struct {
	// ID of the scaling task
	TaskId string `json:"taskId"`
	// Name of the task
	TaskName string `json:"taskName"`
	// Type of task operation
	TaskType string `json:"taskType"`
	// Current status of the task
	Status TaskStatus `json:"status"`
	// Resource type of the task
	ResourceType *string `json:"resourceType,omitempty"`
	// Resource ID of the task
	ResourceId *string `json:"resourceId,omitempty"`
	// Organization name of the task
	OrgName *string `json:"orgName,omitempty"`
	// Timestamp when the task was created
	CreatedAt time.Time `json:"createdAt"`
	// Timestamp when the task was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// Timestamp when the task was deleted
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Time when the task completed or failed
	CompletionTime *time.Time `json:"completionTime,omitempty"`
	// Detailed message about the task status
	Message *string `json:"message,omitempty"`
	// Progress percentage of the task
	Progress *int32 `json:"progress,omitempty"`
	// The user created the task
	Operator *string `json:"operator,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTaskListItem instantiates a new TaskListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTaskListItem(taskId string, taskName string, taskType string, status TaskStatus, createdAt time.Time, updatedAt time.Time) *TaskListItem {
	this := TaskListItem{}
	this.TaskId = taskId
	this.TaskName = taskName
	this.TaskType = taskType
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewTaskListItemWithDefaults instantiates a new TaskListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTaskListItemWithDefaults() *TaskListItem {
	this := TaskListItem{}
	return &this
}

// GetTaskId returns the TaskId field value.
func (o *TaskListItem) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *TaskListItem) SetTaskId(v string) {
	o.TaskId = v
}

// GetTaskName returns the TaskName field value.
func (o *TaskListItem) GetTaskName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetTaskNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskName, true
}

// SetTaskName sets field value.
func (o *TaskListItem) SetTaskName(v string) {
	o.TaskName = v
}

// GetTaskType returns the TaskType field value.
func (o *TaskListItem) GetTaskType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetTaskTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskType, true
}

// SetTaskType sets field value.
func (o *TaskListItem) SetTaskType(v string) {
	o.TaskType = v
}

// GetStatus returns the Status field value.
func (o *TaskListItem) GetStatus() TaskStatus {
	if o == nil {
		var ret TaskStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetStatusOk() (*TaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *TaskListItem) SetStatus(v TaskStatus) {
	o.Status = v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *TaskListItem) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *TaskListItem) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *TaskListItem) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *TaskListItem) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *TaskListItem) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *TaskListItem) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *TaskListItem) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *TaskListItem) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *TaskListItem) SetOrgName(v string) {
	o.OrgName = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *TaskListItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *TaskListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *TaskListItem) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *TaskListItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *TaskListItem) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *TaskListItem) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *TaskListItem) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *TaskListItem) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *TaskListItem) HasCompletionTime() bool {
	return o != nil && o.CompletionTime != nil
}

// SetCompletionTime gets a reference to the given time.Time and assigns it to the CompletionTime field.
func (o *TaskListItem) SetCompletionTime(v time.Time) {
	o.CompletionTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TaskListItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TaskListItem) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TaskListItem) SetMessage(v string) {
	o.Message = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *TaskListItem) GetProgress() int32 {
	if o == nil || o.Progress == nil {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetProgressOk() (*int32, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *TaskListItem) HasProgress() bool {
	return o != nil && o.Progress != nil
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *TaskListItem) SetProgress(v int32) {
	o.Progress = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *TaskListItem) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskListItem) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *TaskListItem) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *TaskListItem) SetOperator(v string) {
	o.Operator = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TaskListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId
	toSerialize["taskName"] = o.TaskName
	toSerialize["taskType"] = o.TaskType
	toSerialize["status"] = o.Status
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
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
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TaskListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId         *string     `json:"taskId"`
		TaskName       *string     `json:"taskName"`
		TaskType       *string     `json:"taskType"`
		Status         *TaskStatus `json:"status"`
		ResourceType   *string     `json:"resourceType,omitempty"`
		ResourceId     *string     `json:"resourceId,omitempty"`
		OrgName        *string     `json:"orgName,omitempty"`
		CreatedAt      *time.Time  `json:"createdAt"`
		UpdatedAt      *time.Time  `json:"updatedAt"`
		StartedAt      *time.Time  `json:"startedAt,omitempty"`
		CompletionTime *time.Time  `json:"completionTime,omitempty"`
		Message        *string     `json:"message,omitempty"`
		Progress       *int32      `json:"progress,omitempty"`
		Operator       *string     `json:"operator,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"taskId", "taskName", "taskType", "status", "resourceType", "resourceId", "orgName", "createdAt", "updatedAt", "startedAt", "completionTime", "message", "progress", "operator"})
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
	o.ResourceType = all.ResourceType
	o.ResourceId = all.ResourceId
	o.OrgName = all.OrgName
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt
	o.StartedAt = all.StartedAt
	o.CompletionTime = all.CompletionTime
	o.Message = all.Message
	o.Progress = all.Progress
	o.Operator = all.Operator

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
