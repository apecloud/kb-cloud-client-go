// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TaskType struct {
	// Name of the task type
	Name        string                `json:"name"`
	DisplayName *LocalizedDescription `json:"displayName,omitempty"`
	Description *LocalizedDescription `json:"description,omitempty"`
	// Category of the task type
	Category *string `json:"category,omitempty"`
	// Workflow engine of the task type
	WorkflowEngine string `json:"workflowEngine"`
	// Resource type of the task type
	ResourceType *string `json:"resourceType,omitempty"`
	// Creation time of the task type
	CreatedAt time.Time `json:"createdAt"`
	// Update time of the task type
	UpdatedAt time.Time `json:"updatedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTaskType instantiates a new TaskType object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTaskType(name string, workflowEngine string, createdAt time.Time, updatedAt time.Time) *TaskType {
	this := TaskType{}
	this.Name = name
	this.WorkflowEngine = workflowEngine
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewTaskTypeWithDefaults instantiates a new TaskType object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTaskTypeWithDefaults() *TaskType {
	this := TaskType{}
	return &this
}

// GetName returns the Name field value.
func (o *TaskType) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TaskType) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TaskType) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TaskType) GetDisplayName() LocalizedDescription {
	if o == nil || o.DisplayName == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskType) GetDisplayNameOk() (*LocalizedDescription, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TaskType) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given LocalizedDescription and assigns it to the DisplayName field.
func (o *TaskType) SetDisplayName(v LocalizedDescription) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TaskType) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskType) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TaskType) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *TaskType) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *TaskType) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskType) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *TaskType) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *TaskType) SetCategory(v string) {
	o.Category = &v
}

// GetWorkflowEngine returns the WorkflowEngine field value.
func (o *TaskType) GetWorkflowEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WorkflowEngine
}

// GetWorkflowEngineOk returns a tuple with the WorkflowEngine field value
// and a boolean to check if the value has been set.
func (o *TaskType) GetWorkflowEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowEngine, true
}

// SetWorkflowEngine sets field value.
func (o *TaskType) SetWorkflowEngine(v string) {
	o.WorkflowEngine = v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *TaskType) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskType) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *TaskType) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *TaskType) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *TaskType) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskType) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *TaskType) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *TaskType) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskType) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *TaskType) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TaskType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	toSerialize["workflowEngine"] = o.WorkflowEngine
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TaskType) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string               `json:"name"`
		DisplayName    *LocalizedDescription `json:"displayName,omitempty"`
		Description    *LocalizedDescription `json:"description,omitempty"`
		Category       *string               `json:"category,omitempty"`
		WorkflowEngine *string               `json:"workflowEngine"`
		ResourceType   *string               `json:"resourceType,omitempty"`
		CreatedAt      *time.Time            `json:"createdAt"`
		UpdatedAt      *time.Time            `json:"updatedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.WorkflowEngine == nil {
		return fmt.Errorf("required field workflowEngine missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "displayName", "description", "category", "workflowEngine", "resourceType", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.DisplayName != nil && all.DisplayName.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisplayName = all.DisplayName
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	o.Category = all.Category
	o.WorkflowEngine = *all.WorkflowEngine
	o.ResourceType = all.ResourceType
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
