// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

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
	// Numeric form of result when the result is parseable.
	ValueNum      *float64 `json:"valueNum,omitempty"`
	WarnThreshold *float64 `json:"warnThreshold,omitempty"`
	CritThreshold *float64 `json:"critThreshold,omitempty"`
	// Direction used to interpret warnThreshold and critThreshold. asc means larger values are worse, desc means smaller values are worse, boolean means the check expression decides.
	Direction *InspectionThresholdDirection `json:"direction,omitempty"`
	// First-version criticality assumption for score weighting and red-item veto behavior. Missing legacy values are treated as medium.
	Criticality *InspectionCriticality `json:"criticality,omitempty"`
	// Structured evidence used to explain the item result. Prometheus/exporter no-data is reported here as no data or unknown instead of being converted into a healthy result.
	Evidence map[string]interface{} `json:"evidence,omitempty"`
	// Item-level timestamp for when this item's status was last evaluated or changed.
	StatusChangedAt *time.Time            `json:"statusChangedAt,omitempty"`
	Remediation     *LocalizedDescription `json:"remediation,omitempty"`
	DocLink         *string               `json:"docLink,omitempty"`
	Severity        *string               `json:"severity,omitempty"`
	Unit            *string               `json:"unit,omitempty"`
	CreatedAt       *time.Time            `json:"createdAt,omitempty"`
	UpdatedAt       *time.Time            `json:"updatedAt,omitempty"`
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

// GetValueNum returns the ValueNum field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetValueNum() float64 {
	if o == nil || o.ValueNum == nil {
		var ret float64
		return ret
	}
	return *o.ValueNum
}

// GetValueNumOk returns a tuple with the ValueNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetValueNumOk() (*float64, bool) {
	if o == nil || o.ValueNum == nil {
		return nil, false
	}
	return o.ValueNum, true
}

// HasValueNum returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasValueNum() bool {
	return o != nil && o.ValueNum != nil
}

// SetValueNum gets a reference to the given float64 and assigns it to the ValueNum field.
func (o *InspectionTaskItem) SetValueNum(v float64) {
	o.ValueNum = &v
}

// GetWarnThreshold returns the WarnThreshold field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetWarnThreshold() float64 {
	if o == nil || o.WarnThreshold == nil {
		var ret float64
		return ret
	}
	return *o.WarnThreshold
}

// GetWarnThresholdOk returns a tuple with the WarnThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetWarnThresholdOk() (*float64, bool) {
	if o == nil || o.WarnThreshold == nil {
		return nil, false
	}
	return o.WarnThreshold, true
}

// HasWarnThreshold returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasWarnThreshold() bool {
	return o != nil && o.WarnThreshold != nil
}

// SetWarnThreshold gets a reference to the given float64 and assigns it to the WarnThreshold field.
func (o *InspectionTaskItem) SetWarnThreshold(v float64) {
	o.WarnThreshold = &v
}

// GetCritThreshold returns the CritThreshold field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetCritThreshold() float64 {
	if o == nil || o.CritThreshold == nil {
		var ret float64
		return ret
	}
	return *o.CritThreshold
}

// GetCritThresholdOk returns a tuple with the CritThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetCritThresholdOk() (*float64, bool) {
	if o == nil || o.CritThreshold == nil {
		return nil, false
	}
	return o.CritThreshold, true
}

// HasCritThreshold returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasCritThreshold() bool {
	return o != nil && o.CritThreshold != nil
}

// SetCritThreshold gets a reference to the given float64 and assigns it to the CritThreshold field.
func (o *InspectionTaskItem) SetCritThreshold(v float64) {
	o.CritThreshold = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetDirection() InspectionThresholdDirection {
	if o == nil || o.Direction == nil {
		var ret InspectionThresholdDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetDirectionOk() (*InspectionThresholdDirection, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasDirection() bool {
	return o != nil && o.Direction != nil
}

// SetDirection gets a reference to the given InspectionThresholdDirection and assigns it to the Direction field.
func (o *InspectionTaskItem) SetDirection(v InspectionThresholdDirection) {
	o.Direction = &v
}

// GetCriticality returns the Criticality field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetCriticality() InspectionCriticality {
	if o == nil || o.Criticality == nil {
		var ret InspectionCriticality
		return ret
	}
	return *o.Criticality
}

// GetCriticalityOk returns a tuple with the Criticality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetCriticalityOk() (*InspectionCriticality, bool) {
	if o == nil || o.Criticality == nil {
		return nil, false
	}
	return o.Criticality, true
}

// HasCriticality returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasCriticality() bool {
	return o != nil && o.Criticality != nil
}

// SetCriticality gets a reference to the given InspectionCriticality and assigns it to the Criticality field.
func (o *InspectionTaskItem) SetCriticality(v InspectionCriticality) {
	o.Criticality = &v
}

// GetEvidence returns the Evidence field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetEvidence() map[string]interface{} {
	if o == nil || o.Evidence == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Evidence
}

// GetEvidenceOk returns a tuple with the Evidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetEvidenceOk() (*map[string]interface{}, bool) {
	if o == nil || o.Evidence == nil {
		return nil, false
	}
	return &o.Evidence, true
}

// HasEvidence returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasEvidence() bool {
	return o != nil && o.Evidence != nil
}

// SetEvidence gets a reference to the given map[string]interface{} and assigns it to the Evidence field.
func (o *InspectionTaskItem) SetEvidence(v map[string]interface{}) {
	o.Evidence = v
}

// GetStatusChangedAt returns the StatusChangedAt field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetStatusChangedAt() time.Time {
	if o == nil || o.StatusChangedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StatusChangedAt
}

// GetStatusChangedAtOk returns a tuple with the StatusChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetStatusChangedAtOk() (*time.Time, bool) {
	if o == nil || o.StatusChangedAt == nil {
		return nil, false
	}
	return o.StatusChangedAt, true
}

// HasStatusChangedAt returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasStatusChangedAt() bool {
	return o != nil && o.StatusChangedAt != nil
}

// SetStatusChangedAt gets a reference to the given time.Time and assigns it to the StatusChangedAt field.
func (o *InspectionTaskItem) SetStatusChangedAt(v time.Time) {
	o.StatusChangedAt = &v
}

// GetRemediation returns the Remediation field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetRemediation() LocalizedDescription {
	if o == nil || o.Remediation == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Remediation
}

// GetRemediationOk returns a tuple with the Remediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetRemediationOk() (*LocalizedDescription, bool) {
	if o == nil || o.Remediation == nil {
		return nil, false
	}
	return o.Remediation, true
}

// HasRemediation returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasRemediation() bool {
	return o != nil && o.Remediation != nil
}

// SetRemediation gets a reference to the given LocalizedDescription and assigns it to the Remediation field.
func (o *InspectionTaskItem) SetRemediation(v LocalizedDescription) {
	o.Remediation = &v
}

// GetDocLink returns the DocLink field value if set, zero value otherwise.
func (o *InspectionTaskItem) GetDocLink() string {
	if o == nil || o.DocLink == nil {
		var ret string
		return ret
	}
	return *o.DocLink
}

// GetDocLinkOk returns a tuple with the DocLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskItem) GetDocLinkOk() (*string, bool) {
	if o == nil || o.DocLink == nil {
		return nil, false
	}
	return o.DocLink, true
}

// HasDocLink returns a boolean if a field has been set.
func (o *InspectionTaskItem) HasDocLink() bool {
	return o != nil && o.DocLink != nil
}

// SetDocLink gets a reference to the given string and assigns it to the DocLink field.
func (o *InspectionTaskItem) SetDocLink(v string) {
	o.DocLink = &v
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
	if o.ValueNum != nil {
		toSerialize["valueNum"] = o.ValueNum
	}
	if o.WarnThreshold != nil {
		toSerialize["warnThreshold"] = o.WarnThreshold
	}
	if o.CritThreshold != nil {
		toSerialize["critThreshold"] = o.CritThreshold
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Criticality != nil {
		toSerialize["criticality"] = o.Criticality
	}
	if o.Evidence != nil {
		toSerialize["evidence"] = o.Evidence
	}
	if o.StatusChangedAt != nil {
		if o.StatusChangedAt.Nanosecond() == 0 {
			toSerialize["statusChangedAt"] = o.StatusChangedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["statusChangedAt"] = o.StatusChangedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Remediation != nil {
		toSerialize["remediation"] = o.Remediation
	}
	if o.DocLink != nil {
		toSerialize["docLink"] = o.DocLink
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
		Id                *string                       `json:"id,omitempty"`
		TaskId            *string                       `json:"taskID,omitempty"`
		ScriptId          *string                       `json:"scriptID,omitempty"`
		ScriptName        *LocalizedDescription         `json:"scriptName,omitempty"`
		ScriptDescription *LocalizedDescription         `json:"scriptDescription,omitempty"`
		ScriptCategory    *string                       `json:"scriptCategory,omitempty"`
		ResourceType      *string                       `json:"resourceType,omitempty"`
		ResourceId        *string                       `json:"resourceID,omitempty"`
		ResourceName      *string                       `json:"resourceName,omitempty"`
		Status            *string                       `json:"status,omitempty"`
		Result            *string                       `json:"result,omitempty"`
		ValueNum          *float64                      `json:"valueNum,omitempty"`
		WarnThreshold     *float64                      `json:"warnThreshold,omitempty"`
		CritThreshold     *float64                      `json:"critThreshold,omitempty"`
		Direction         *InspectionThresholdDirection `json:"direction,omitempty"`
		Criticality       *InspectionCriticality        `json:"criticality,omitempty"`
		Evidence          map[string]interface{}        `json:"evidence,omitempty"`
		StatusChangedAt   *time.Time                    `json:"statusChangedAt,omitempty"`
		Remediation       *LocalizedDescription         `json:"remediation,omitempty"`
		DocLink           *string                       `json:"docLink,omitempty"`
		Severity          *string                       `json:"severity,omitempty"`
		Unit              *string                       `json:"unit,omitempty"`
		CreatedAt         *time.Time                    `json:"createdAt,omitempty"`
		UpdatedAt         *time.Time                    `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "taskID", "scriptID", "scriptName", "scriptDescription", "scriptCategory", "resourceType", "resourceID", "resourceName", "status", "result", "valueNum", "warnThreshold", "critThreshold", "direction", "criticality", "evidence", "statusChangedAt", "remediation", "docLink", "severity", "unit", "createdAt", "updatedAt"})
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
	o.ValueNum = all.ValueNum
	o.WarnThreshold = all.WarnThreshold
	o.CritThreshold = all.CritThreshold
	if all.Direction != nil && !all.Direction.IsValid() {
		hasInvalidField = true
	} else {
		o.Direction = all.Direction
	}
	if all.Criticality != nil && !all.Criticality.IsValid() {
		hasInvalidField = true
	} else {
		o.Criticality = all.Criticality
	}
	o.Evidence = all.Evidence
	o.StatusChangedAt = all.StatusChangedAt
	if all.Remediation != nil && all.Remediation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Remediation = all.Remediation
	o.DocLink = all.DocLink
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
