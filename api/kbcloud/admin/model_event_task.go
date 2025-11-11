// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventTask event task is the information of the operation event
type EventTask struct {
	// ID of the task
	Id *string `json:"id,omitempty"`
	// Name of the task
	Name string `json:"name"`
	// Namespace of the task
	Namespace string `json:"namespace"`
	// OrgName of the task
	OrgName *string `json:"orgName,omitempty"`
	// ResourceType of the task
	ResourceType *string `json:"resourceType,omitempty"`
	// ResourceId of the task
	ResourceId *string `json:"resourceId,omitempty"`
	// ResourceName of the task
	ResourceName *string `json:"resourceName,omitempty"`
	// status of the task
	Status string `json:"status"`
	// task type
	TaskType string `json:"taskType"`
	// progress of the task
	Progress string `json:"progress"`
	// eventTaskProgresses is a list of event task progress detail
	TaskProgresses *EventTaskProgresses `json:"taskProgresses,omitempty"`
	// eventTaskDetails is a list of event task detail
	TaskDetails *EventTaskDetails `json:"taskDetails,omitempty"`
	// pod log of the custom task
	OpsLog *string `json:"opsLog,omitempty"`
	// description of the custom task
	Description *string `json:"description,omitempty"`
	// Time when the task started
	StartTime *time.Time `json:"startTime,omitempty"`
	// Time when the task completed or failed
	CompletionTime common.NullableTime `json:"completionTime,omitempty"`
	// Same as `id`. Frontend can ignore this field.
	EventTaskId *string `json:"eventTaskId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventTask instantiates a new EventTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventTask(name string, namespace string, status string, taskType string, progress string) *EventTask {
	this := EventTask{}
	this.Name = name
	this.Namespace = namespace
	this.Status = status
	this.TaskType = taskType
	this.Progress = progress
	return &this
}

// NewEventTaskWithDefaults instantiates a new EventTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventTaskWithDefaults() *EventTask {
	this := EventTask{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventTask) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventTask) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventTask) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value.
func (o *EventTask) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventTask) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EventTask) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value.
func (o *EventTask) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *EventTask) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *EventTask) SetNamespace(v string) {
	o.Namespace = v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *EventTask) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *EventTask) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *EventTask) SetOrgName(v string) {
	o.OrgName = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *EventTask) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *EventTask) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *EventTask) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *EventTask) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *EventTask) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *EventTask) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *EventTask) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *EventTask) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *EventTask) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetStatus returns the Status field value.
func (o *EventTask) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EventTask) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *EventTask) SetStatus(v string) {
	o.Status = v
}

// GetTaskType returns the TaskType field value.
func (o *EventTask) GetTaskType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value
// and a boolean to check if the value has been set.
func (o *EventTask) GetTaskTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskType, true
}

// SetTaskType sets field value.
func (o *EventTask) SetTaskType(v string) {
	o.TaskType = v
}

// GetProgress returns the Progress field value.
func (o *EventTask) GetProgress() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *EventTask) GetProgressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value.
func (o *EventTask) SetProgress(v string) {
	o.Progress = v
}

// GetTaskProgresses returns the TaskProgresses field value if set, zero value otherwise.
func (o *EventTask) GetTaskProgresses() EventTaskProgresses {
	if o == nil || o.TaskProgresses == nil {
		var ret EventTaskProgresses
		return ret
	}
	return *o.TaskProgresses
}

// GetTaskProgressesOk returns a tuple with the TaskProgresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetTaskProgressesOk() (*EventTaskProgresses, bool) {
	if o == nil || o.TaskProgresses == nil {
		return nil, false
	}
	return o.TaskProgresses, true
}

// HasTaskProgresses returns a boolean if a field has been set.
func (o *EventTask) HasTaskProgresses() bool {
	return o != nil && o.TaskProgresses != nil
}

// SetTaskProgresses gets a reference to the given EventTaskProgresses and assigns it to the TaskProgresses field.
func (o *EventTask) SetTaskProgresses(v EventTaskProgresses) {
	o.TaskProgresses = &v
}

// GetTaskDetails returns the TaskDetails field value if set, zero value otherwise.
func (o *EventTask) GetTaskDetails() EventTaskDetails {
	if o == nil || o.TaskDetails == nil {
		var ret EventTaskDetails
		return ret
	}
	return *o.TaskDetails
}

// GetTaskDetailsOk returns a tuple with the TaskDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetTaskDetailsOk() (*EventTaskDetails, bool) {
	if o == nil || o.TaskDetails == nil {
		return nil, false
	}
	return o.TaskDetails, true
}

// HasTaskDetails returns a boolean if a field has been set.
func (o *EventTask) HasTaskDetails() bool {
	return o != nil && o.TaskDetails != nil
}

// SetTaskDetails gets a reference to the given EventTaskDetails and assigns it to the TaskDetails field.
func (o *EventTask) SetTaskDetails(v EventTaskDetails) {
	o.TaskDetails = &v
}

// GetOpsLog returns the OpsLog field value if set, zero value otherwise.
func (o *EventTask) GetOpsLog() string {
	if o == nil || o.OpsLog == nil {
		var ret string
		return ret
	}
	return *o.OpsLog
}

// GetOpsLogOk returns a tuple with the OpsLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetOpsLogOk() (*string, bool) {
	if o == nil || o.OpsLog == nil {
		return nil, false
	}
	return o.OpsLog, true
}

// HasOpsLog returns a boolean if a field has been set.
func (o *EventTask) HasOpsLog() bool {
	return o != nil && o.OpsLog != nil
}

// SetOpsLog gets a reference to the given string and assigns it to the OpsLog field.
func (o *EventTask) SetOpsLog(v string) {
	o.OpsLog = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EventTask) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EventTask) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EventTask) SetDescription(v string) {
	o.Description = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *EventTask) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *EventTask) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *EventTask) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTask) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime.Get()
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventTask) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletionTime.Get(), o.CompletionTime.IsSet()
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *EventTask) HasCompletionTime() bool {
	return o != nil && o.CompletionTime.IsSet()
}

// SetCompletionTime gets a reference to the given common.NullableTime and assigns it to the CompletionTime field.
func (o *EventTask) SetCompletionTime(v time.Time) {
	o.CompletionTime.Set(&v)
}

// SetCompletionTimeNil sets the value for CompletionTime to be an explicit nil.
func (o *EventTask) SetCompletionTimeNil() {
	o.CompletionTime.Set(nil)
}

// UnsetCompletionTime ensures that no value is present for CompletionTime, not even an explicit nil.
func (o *EventTask) UnsetCompletionTime() {
	o.CompletionTime.Unset()
}

// GetEventTaskId returns the EventTaskId field value if set, zero value otherwise.
func (o *EventTask) GetEventTaskId() string {
	if o == nil || o.EventTaskId == nil {
		var ret string
		return ret
	}
	return *o.EventTaskId
}

// GetEventTaskIdOk returns a tuple with the EventTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTask) GetEventTaskIdOk() (*string, bool) {
	if o == nil || o.EventTaskId == nil {
		return nil, false
	}
	return o.EventTaskId, true
}

// HasEventTaskId returns a boolean if a field has been set.
func (o *EventTask) HasEventTaskId() bool {
	return o != nil && o.EventTaskId != nil
}

// SetEventTaskId gets a reference to the given string and assigns it to the EventTaskId field.
func (o *EventTask) SetEventTaskId(v string) {
	o.EventTaskId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
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
	toSerialize["status"] = o.Status
	toSerialize["taskType"] = o.TaskType
	toSerialize["progress"] = o.Progress
	if o.TaskProgresses != nil {
		toSerialize["taskProgresses"] = o.TaskProgresses
	}
	if o.TaskDetails != nil {
		toSerialize["taskDetails"] = o.TaskDetails
	}
	if o.OpsLog != nil {
		toSerialize["opsLog"] = o.OpsLog
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.StartTime != nil {
		if o.StartTime.Nanosecond() == 0 {
			toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CompletionTime.IsSet() {
		toSerialize["completionTime"] = o.CompletionTime.Get()
	}
	if o.EventTaskId != nil {
		toSerialize["eventTaskId"] = o.EventTaskId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id             *string              `json:"id,omitempty"`
		Name           *string              `json:"name"`
		Namespace      *string              `json:"namespace"`
		OrgName        *string              `json:"orgName,omitempty"`
		ResourceType   *string              `json:"resourceType,omitempty"`
		ResourceId     *string              `json:"resourceId,omitempty"`
		ResourceName   *string              `json:"resourceName,omitempty"`
		Status         *string              `json:"status"`
		TaskType       *string              `json:"taskType"`
		Progress       *string              `json:"progress"`
		TaskProgresses *EventTaskProgresses `json:"taskProgresses,omitempty"`
		TaskDetails    *EventTaskDetails    `json:"taskDetails,omitempty"`
		OpsLog         *string              `json:"opsLog,omitempty"`
		Description    *string              `json:"description,omitempty"`
		StartTime      *time.Time           `json:"startTime,omitempty"`
		CompletionTime common.NullableTime  `json:"completionTime,omitempty"`
		EventTaskId    *string              `json:"eventTaskId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.TaskType == nil {
		return fmt.Errorf("required field taskType missing")
	}
	if all.Progress == nil {
		return fmt.Errorf("required field progress missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "namespace", "orgName", "resourceType", "resourceId", "resourceName", "status", "taskType", "progress", "taskProgresses", "taskDetails", "opsLog", "description", "startTime", "completionTime", "eventTaskId"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = *all.Name
	o.Namespace = *all.Namespace
	o.OrgName = all.OrgName
	o.ResourceType = all.ResourceType
	o.ResourceId = all.ResourceId
	o.ResourceName = all.ResourceName
	o.Status = *all.Status
	o.TaskType = *all.TaskType
	o.Progress = *all.Progress
	if all.TaskProgresses != nil && all.TaskProgresses.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TaskProgresses = all.TaskProgresses
	if all.TaskDetails != nil && all.TaskDetails.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TaskDetails = all.TaskDetails
	o.OpsLog = all.OpsLog
	o.Description = all.Description
	o.StartTime = all.StartTime
	o.CompletionTime = all.CompletionTime
	o.EventTaskId = all.EventTaskId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
