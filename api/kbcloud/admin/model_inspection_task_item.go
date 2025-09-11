// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InspectionTaskItem struct {
	Id                *string               `json:"id,omitempty"`
	TaskId            *string               `json:"taskID,omitempty"`
	ScriptId          *string               `json:"scriptID,omitempty"`
	ScriptName        *LocalizedDescription `json:"scriptName,omitempty"`
	ScriptDescription *LocalizedDescription `json:"scriptDescription,omitempty"`
	ScriptCategory    *string               `json:"scriptCategory,omitempty"`
	ResourceType      *string               `json:"resourceType,omitempty"`
	ResourceId        *string               `json:"resourceID,omitempty"`
	ResourceName      *string               `json:"resourceName,omitempty"`
	Status            *string               `json:"status,omitempty"`
	Result            *string               `json:"result,omitempty"`
	Severity          *string               `json:"severity,omitempty"`
	Unit              *string               `json:"unit,omitempty"`
	CreatedAt         *time.Time            `json:"createdAt,omitempty"`
	UpdatedAt         *time.Time            `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInspectionTaskItem instantiates a new InspectionTaskItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInspectionTaskItem() *InspectionTaskItem {
	this := InspectionTaskItem{}
	return &this
}

// NewInspectionTaskItemWithDefaults instantiates a new InspectionTaskItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInspectionTaskItemWithDefaults() *InspectionTaskItem {
	this := InspectionTaskItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InspectionTaskItem) SetId(v string) {
	o.Id = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasTaskId() bool {
	return o != nil && o.TaskId != nil
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *InspectionTaskItem) SetTaskId(v string) {
	o.TaskId = &v
}

// GetScriptId returns the ScriptId field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetScriptId() string {
	if o == nil || o.ScriptId == nil {
		var ret string
		return ret
	}
	return *o.ScriptId
}

// GetScriptIdOk returns a tuple with the ScriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetScriptIdOk() (*string, bool) {
	if o == nil || o.ScriptId == nil {
		return nil, false
	}
	return o.ScriptId, true
}

// HasScriptId returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasScriptId() bool {
	return o != nil && o.ScriptId != nil
}

// SetScriptId gets a reference to the given string and assigns it to the ScriptId field.
func (o *InspectionTaskItem) SetScriptId(v string) {
	o.ScriptId = &v
}

// GetScriptName returns the ScriptName field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetScriptName() LocalizedDescription {
	if o == nil || o.ScriptName == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.ScriptName
}

// GetScriptNameOk returns a tuple with the ScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetScriptNameOk() (*LocalizedDescription, bool) {
	if o == nil || o.ScriptName == nil {
		return nil, false
	}
	return o.ScriptName, true
}

// HasScriptName returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasScriptName() bool {
	return o != nil && o.ScriptName != nil
}

// SetScriptName gets a reference to the given LocalizedDescription and assigns it to the ScriptName field.
func (o *InspectionTaskItem) SetScriptName(v LocalizedDescription) {
	o.ScriptName = &v
}

// GetScriptDescription returns the ScriptDescription field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetScriptDescription() LocalizedDescription {
	if o == nil || o.ScriptDescription == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.ScriptDescription
}

// GetScriptDescriptionOk returns a tuple with the ScriptDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetScriptDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.ScriptDescription == nil {
		return nil, false
	}
	return o.ScriptDescription, true
}

// HasScriptDescription returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasScriptDescription() bool {
	return o != nil && o.ScriptDescription != nil
}

// SetScriptDescription gets a reference to the given LocalizedDescription and assigns it to the ScriptDescription field.
func (o *InspectionTaskItem) SetScriptDescription(v LocalizedDescription) {
	o.ScriptDescription = &v
}

// GetScriptCategory returns the ScriptCategory field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetScriptCategory() string {
	if o == nil || o.ScriptCategory == nil {
		var ret string
		return ret
	}
	return *o.ScriptCategory
}

// GetScriptCategoryOk returns a tuple with the ScriptCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetScriptCategoryOk() (*string, bool) {
	if o == nil || o.ScriptCategory == nil {
		return nil, false
	}
	return o.ScriptCategory, true
}

// HasScriptCategory returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasScriptCategory() bool {
	return o != nil && o.ScriptCategory != nil
}

// SetScriptCategory gets a reference to the given string and assigns it to the ScriptCategory field.
func (o *InspectionTaskItem) SetScriptCategory(v string) {
	o.ScriptCategory = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasResourceType() bool {
	return o != nil && o.ResourceType != nil
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *InspectionTaskItem) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *InspectionTaskItem) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasResourceName() bool {
	return o != nil && o.ResourceName != nil
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *InspectionTaskItem) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InspectionTaskItem) SetStatus(v string) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *InspectionTaskItem) SetResult(v string) {
	o.Result = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *InspectionTaskItem) SetSeverity(v string) {
	o.Severity = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *InspectionTaskItem) SetUnit(v string) {
	o.Unit = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InspectionTaskItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InspectionTaskItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InspectionTaskItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TaskId != nil {
		toSerialize["taskID"] = o.TaskId
	}
	if o.ScriptId != nil {
		toSerialize["scriptID"] = o.ScriptId
	}
	if o.ScriptName != nil {
		toSerialize["scriptName"] = o.ScriptName
	}
	if o.ScriptDescription != nil {
		toSerialize["scriptDescription"] = o.ScriptDescription
	}
	if o.ScriptCategory != nil {
		toSerialize["scriptCategory"] = o.ScriptCategory
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.ResourceId != nil {
		toSerialize["resourceID"] = o.ResourceId
	}
	if o.ResourceName != nil {
		toSerialize["resourceName"] = o.ResourceName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InspectionTaskItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                *string               `json:"id,omitempty"`
		TaskId            *string               `json:"taskID,omitempty"`
		ScriptId          *string               `json:"scriptID,omitempty"`
		ScriptName        *LocalizedDescription `json:"scriptName,omitempty"`
		ScriptDescription *LocalizedDescription `json:"scriptDescription,omitempty"`
		ScriptCategory    *string               `json:"scriptCategory,omitempty"`
		ResourceType      *string               `json:"resourceType,omitempty"`
		ResourceId        *string               `json:"resourceID,omitempty"`
		ResourceName      *string               `json:"resourceName,omitempty"`
		Status            *string               `json:"status,omitempty"`
		Result            *string               `json:"result,omitempty"`
		Severity          *string               `json:"severity,omitempty"`
		Unit              *string               `json:"unit,omitempty"`
		CreatedAt         *time.Time            `json:"createdAt,omitempty"`
		UpdatedAt         *time.Time            `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "taskID", "scriptID", "scriptName", "scriptDescription", "scriptCategory", "resourceType", "resourceID", "resourceName", "status", "result", "severity", "unit", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.TaskId = all.TaskId
	o.ScriptId = all.ScriptId
	if all.ScriptName != nil && all.ScriptName.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScriptName = all.ScriptName
	if all.ScriptDescription != nil && all.ScriptDescription.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScriptDescription = all.ScriptDescription
	o.ScriptCategory = all.ScriptCategory
	o.ResourceType = all.ResourceType
	o.ResourceId = all.ResourceId
	o.ResourceName = all.ResourceName
	o.Status = all.Status
	o.Result = all.Result
	o.Severity = all.Severity
	o.Unit = all.Unit
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
