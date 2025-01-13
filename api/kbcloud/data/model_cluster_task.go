// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package data

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// ClusterTask task is the information of the operation 
type ClusterTask struct {
	// ID of the task
	Id *string `json:"id,omitempty"`
	// Name of the task
	Name string `json:"name"`
	// Namespace of the task
	Namespace string `json:"namespace"`
	// OrgName of the task
	OrgName *string `json:"orgName,omitempty"`
	// ClusterName of the task
	ClusterName *string `json:"clusterName,omitempty"`
	// status of the task
	Status string `json:"status"`
	// task type
	TaskType string `json:"taskType"`
	// progress of the task
	Progress string `json:"progress"`
	// clusterTaskProgresses is a list of task progress detail
	TaskProgresses *ClusterTaskProgresses `json:"taskProgresses,omitempty"`
	// taskConditions is a list of task condition
	TaskDetails *ClusterTaskDetails `json:"taskDetails,omitempty"`
	// pod log of the custom task
	OpsLog *string `json:"opsLog,omitempty"`
	// description of the custom task
	Description *string `json:"description,omitempty"`
	// Time when the task started
	StartTime *time.Time `json:"startTime,omitempty"`
	// Time when the task completed or failed
	CompletionTime common.NullableTime `json:"completionTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewClusterTask instantiates a new ClusterTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterTask(name string, namespace string, status string, taskType string, progress string) *ClusterTask {
	this := ClusterTask{}
	this.Name = name
	this.Namespace = namespace
	this.Status = status
	this.TaskType = taskType
	this.Progress = progress
	return &this
}

// NewClusterTaskWithDefaults instantiates a new ClusterTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterTaskWithDefaults() *ClusterTask {
	this := ClusterTask{}
	return &this
}
// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterTask) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterTask) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterTask) SetId(v string) {
	o.Id = &v
}


// GetName returns the Name field value.
func (o *ClusterTask) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ClusterTask) SetName(v string) {
	o.Name = v
}


// GetNamespace returns the Namespace field value.
func (o *ClusterTask) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *ClusterTask) SetNamespace(v string) {
	o.Namespace = v
}


// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *ClusterTask) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *ClusterTask) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *ClusterTask) SetOrgName(v string) {
	o.OrgName = &v
}


// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *ClusterTask) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ClusterTask) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ClusterTask) SetClusterName(v string) {
	o.ClusterName = &v
}


// GetStatus returns the Status field value.
func (o *ClusterTask) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ClusterTask) SetStatus(v string) {
	o.Status = v
}


// GetTaskType returns the TaskType field value.
func (o *ClusterTask) GetTaskType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetTaskTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskType, true
}

// SetTaskType sets field value.
func (o *ClusterTask) SetTaskType(v string) {
	o.TaskType = v
}


// GetProgress returns the Progress field value.
func (o *ClusterTask) GetProgress() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetProgressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value.
func (o *ClusterTask) SetProgress(v string) {
	o.Progress = v
}


// GetTaskProgresses returns the TaskProgresses field value if set, zero value otherwise.
func (o *ClusterTask) GetTaskProgresses() ClusterTaskProgresses {
	if o == nil || o.TaskProgresses == nil {
		var ret ClusterTaskProgresses
		return ret
	}
	return *o.TaskProgresses
}

// GetTaskProgressesOk returns a tuple with the TaskProgresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetTaskProgressesOk() (*ClusterTaskProgresses, bool) {
	if o == nil || o.TaskProgresses == nil {
		return nil, false
	}
	return o.TaskProgresses, true
}

// HasTaskProgresses returns a boolean if a field has been set.
func (o *ClusterTask) HasTaskProgresses() bool {
	return o != nil && o.TaskProgresses != nil
}

// SetTaskProgresses gets a reference to the given ClusterTaskProgresses and assigns it to the TaskProgresses field.
func (o *ClusterTask) SetTaskProgresses(v ClusterTaskProgresses) {
	o.TaskProgresses = &v
}


// GetTaskDetails returns the TaskDetails field value if set, zero value otherwise.
func (o *ClusterTask) GetTaskDetails() ClusterTaskDetails {
	if o == nil || o.TaskDetails == nil {
		var ret ClusterTaskDetails
		return ret
	}
	return *o.TaskDetails
}

// GetTaskDetailsOk returns a tuple with the TaskDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetTaskDetailsOk() (*ClusterTaskDetails, bool) {
	if o == nil || o.TaskDetails == nil {
		return nil, false
	}
	return o.TaskDetails, true
}

// HasTaskDetails returns a boolean if a field has been set.
func (o *ClusterTask) HasTaskDetails() bool {
	return o != nil && o.TaskDetails != nil
}

// SetTaskDetails gets a reference to the given ClusterTaskDetails and assigns it to the TaskDetails field.
func (o *ClusterTask) SetTaskDetails(v ClusterTaskDetails) {
	o.TaskDetails = &v
}


// GetOpsLog returns the OpsLog field value if set, zero value otherwise.
func (o *ClusterTask) GetOpsLog() string {
	if o == nil || o.OpsLog == nil {
		var ret string
		return ret
	}
	return *o.OpsLog
}

// GetOpsLogOk returns a tuple with the OpsLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetOpsLogOk() (*string, bool) {
	if o == nil || o.OpsLog == nil {
		return nil, false
	}
	return o.OpsLog, true
}

// HasOpsLog returns a boolean if a field has been set.
func (o *ClusterTask) HasOpsLog() bool {
	return o != nil && o.OpsLog != nil
}

// SetOpsLog gets a reference to the given string and assigns it to the OpsLog field.
func (o *ClusterTask) SetOpsLog(v string) {
	o.OpsLog = &v
}


// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterTask) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterTask) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClusterTask) SetDescription(v string) {
	o.Description = &v
}


// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ClusterTask) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTask) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ClusterTask) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ClusterTask) SetStartTime(v time.Time) {
	o.StartTime = &v
}


// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterTask) GetCompletionTime() time.Time {
	if o == nil || o.CompletionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletionTime.Get()
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterTask) GetCompletionTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompletionTime.Get(), o.CompletionTime.IsSet()
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *ClusterTask) HasCompletionTime() bool {
	return o != nil && o.CompletionTime.IsSet()
}

// SetCompletionTime gets a reference to the given common.NullableTime and assigns it to the CompletionTime field.
func (o *ClusterTask) SetCompletionTime(v time.Time) {
	o.CompletionTime.Set(&v)
}
// SetCompletionTimeNil sets the value for CompletionTime to be an explicit nil.
func (o *ClusterTask) SetCompletionTimeNil() {
	o.CompletionTime.Set(nil)
}

// UnsetCompletionTime ensures that no value is present for CompletionTime, not even an explicit nil.
func (o *ClusterTask) UnsetCompletionTime() {
	o.CompletionTime.Unset()
}



// MarshalJSON serializes the struct using spec logic.
func (o ClusterTask) MarshalJSON() ([]byte, error) {
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
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *string `json:"id,omitempty"`
		Name *string `json:"name"`
		Namespace *string `json:"namespace"`
		OrgName *string `json:"orgName,omitempty"`
		ClusterName *string `json:"clusterName,omitempty"`
		Status *string `json:"status"`
		TaskType *string `json:"taskType"`
		Progress *string `json:"progress"`
		TaskProgresses *ClusterTaskProgresses `json:"taskProgresses,omitempty"`
		TaskDetails *ClusterTaskDetails `json:"taskDetails,omitempty"`
		OpsLog *string `json:"opsLog,omitempty"`
		Description *string `json:"description,omitempty"`
		StartTime *time.Time `json:"startTime,omitempty"`
		CompletionTime common.NullableTime `json:"completionTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
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
		common.DeleteKeys(additionalProperties, &[]string{ "id", "name", "namespace", "orgName", "clusterName", "status", "taskType", "progress", "taskProgresses", "taskDetails", "opsLog", "description", "startTime", "completionTime",  })
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = *all.Name
	o.Namespace = *all.Namespace
	o.OrgName = all.OrgName
	o.ClusterName = all.ClusterName
	o.Status = *all.Status
	o.TaskType = *all.TaskType
	o.Progress = *all.Progress
	if  all.TaskProgresses != nil && all.TaskProgresses.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TaskProgresses = all.TaskProgresses
	if  all.TaskDetails != nil && all.TaskDetails.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TaskDetails = all.TaskDetails
	o.OpsLog = all.OpsLog
	o.Description = all.Description
	o.StartTime = all.StartTime
	o.CompletionTime = all.CompletionTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
