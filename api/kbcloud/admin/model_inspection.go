// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Inspection inspection
type Inspection struct {
	Id             *int32     `json:"id,omitempty"`
	OrgName        *string    `json:"orgName,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	Creator        *string    `json:"creator,omitempty"`
	ClusterId      *string    `json:"clusterID,omitempty"`
	ClusterName    *string    `json:"clusterName,omitempty"`
	ClusterEngine  *string    `json:"clusterEngine,omitempty"`
	Result         *string    `json:"result,omitempty"`
	Severity       *string    `json:"severity,omitempty"`
	Status         *string    `json:"status,omitempty"`
	IsAuto         *bool      `json:"isAuto,omitempty"`
	ScriptId       *int32     `json:"scriptID,omitempty"`
	ScriptName     *string    `json:"scriptName,omitempty"`
	ScriptCategory *string    `json:"scriptCategory,omitempty"`
	Reason         *string    `json:"reason,omitempty"`
	Suggestion     *string    `json:"suggestion,omitempty"`
	Unit           *string    `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInspection instantiates a new Inspection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInspection() *Inspection {
	this := Inspection{}
	return &this
}

// NewInspectionWithDefaults instantiates a new Inspection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInspectionWithDefaults() *Inspection {
	this := Inspection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Inspection) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Inspection) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Inspection) SetId(v int32) {
	o.Id = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Inspection) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Inspection) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Inspection) SetOrgName(v string) {
	o.OrgName = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Inspection) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Inspection) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Inspection) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Inspection) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Inspection) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Inspection) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Inspection) GetCreator() string {
	if o == nil || o.Creator == nil {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetCreatorOk() (*string, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Inspection) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *Inspection) SetCreator(v string) {
	o.Creator = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *Inspection) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *Inspection) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *Inspection) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *Inspection) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *Inspection) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *Inspection) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterEngine returns the ClusterEngine field value if set, zero value otherwise.
func (o *Inspection) GetClusterEngine() string {
	if o == nil || o.ClusterEngine == nil {
		var ret string
		return ret
	}
	return *o.ClusterEngine
}

// GetClusterEngineOk returns a tuple with the ClusterEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetClusterEngineOk() (*string, bool) {
	if o == nil || o.ClusterEngine == nil {
		return nil, false
	}
	return o.ClusterEngine, true
}

// HasClusterEngine returns a boolean if a field has been set.
func (o *Inspection) HasClusterEngine() bool {
	return o != nil && o.ClusterEngine != nil
}

// SetClusterEngine gets a reference to the given string and assigns it to the ClusterEngine field.
func (o *Inspection) SetClusterEngine(v string) {
	o.ClusterEngine = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *Inspection) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *Inspection) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *Inspection) SetResult(v string) {
	o.Result = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *Inspection) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *Inspection) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *Inspection) SetSeverity(v string) {
	o.Severity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Inspection) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Inspection) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Inspection) SetStatus(v string) {
	o.Status = &v
}

// GetIsAuto returns the IsAuto field value if set, zero value otherwise.
func (o *Inspection) GetIsAuto() bool {
	if o == nil || o.IsAuto == nil {
		var ret bool
		return ret
	}
	return *o.IsAuto
}

// GetIsAutoOk returns a tuple with the IsAuto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetIsAutoOk() (*bool, bool) {
	if o == nil || o.IsAuto == nil {
		return nil, false
	}
	return o.IsAuto, true
}

// HasIsAuto returns a boolean if a field has been set.
func (o *Inspection) HasIsAuto() bool {
	return o != nil && o.IsAuto != nil
}

// SetIsAuto gets a reference to the given bool and assigns it to the IsAuto field.
func (o *Inspection) SetIsAuto(v bool) {
	o.IsAuto = &v
}

// GetScriptId returns the ScriptId field value if set, zero value otherwise.
func (o *Inspection) GetScriptId() int32 {
	if o == nil || o.ScriptId == nil {
		var ret int32
		return ret
	}
	return *o.ScriptId
}

// GetScriptIdOk returns a tuple with the ScriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetScriptIdOk() (*int32, bool) {
	if o == nil || o.ScriptId == nil {
		return nil, false
	}
	return o.ScriptId, true
}

// HasScriptId returns a boolean if a field has been set.
func (o *Inspection) HasScriptId() bool {
	return o != nil && o.ScriptId != nil
}

// SetScriptId gets a reference to the given int32 and assigns it to the ScriptId field.
func (o *Inspection) SetScriptId(v int32) {
	o.ScriptId = &v
}

// GetScriptName returns the ScriptName field value if set, zero value otherwise.
func (o *Inspection) GetScriptName() string {
	if o == nil || o.ScriptName == nil {
		var ret string
		return ret
	}
	return *o.ScriptName
}

// GetScriptNameOk returns a tuple with the ScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetScriptNameOk() (*string, bool) {
	if o == nil || o.ScriptName == nil {
		return nil, false
	}
	return o.ScriptName, true
}

// HasScriptName returns a boolean if a field has been set.
func (o *Inspection) HasScriptName() bool {
	return o != nil && o.ScriptName != nil
}

// SetScriptName gets a reference to the given string and assigns it to the ScriptName field.
func (o *Inspection) SetScriptName(v string) {
	o.ScriptName = &v
}

// GetScriptCategory returns the ScriptCategory field value if set, zero value otherwise.
func (o *Inspection) GetScriptCategory() string {
	if o == nil || o.ScriptCategory == nil {
		var ret string
		return ret
	}
	return *o.ScriptCategory
}

// GetScriptCategoryOk returns a tuple with the ScriptCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetScriptCategoryOk() (*string, bool) {
	if o == nil || o.ScriptCategory == nil {
		return nil, false
	}
	return o.ScriptCategory, true
}

// HasScriptCategory returns a boolean if a field has been set.
func (o *Inspection) HasScriptCategory() bool {
	return o != nil && o.ScriptCategory != nil
}

// SetScriptCategory gets a reference to the given string and assigns it to the ScriptCategory field.
func (o *Inspection) SetScriptCategory(v string) {
	o.ScriptCategory = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Inspection) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Inspection) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Inspection) SetReason(v string) {
	o.Reason = &v
}

// GetSuggestion returns the Suggestion field value if set, zero value otherwise.
func (o *Inspection) GetSuggestion() string {
	if o == nil || o.Suggestion == nil {
		var ret string
		return ret
	}
	return *o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetSuggestionOk() (*string, bool) {
	if o == nil || o.Suggestion == nil {
		return nil, false
	}
	return o.Suggestion, true
}

// HasSuggestion returns a boolean if a field has been set.
func (o *Inspection) HasSuggestion() bool {
	return o != nil && o.Suggestion != nil
}

// SetSuggestion gets a reference to the given string and assigns it to the Suggestion field.
func (o *Inspection) SetSuggestion(v string) {
	o.Suggestion = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Inspection) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inspection) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Inspection) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Inspection) SetUnit(v string) {
	o.Unit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Inspection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.ClusterId != nil {
		toSerialize["clusterID"] = o.ClusterId
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.ClusterEngine != nil {
		toSerialize["clusterEngine"] = o.ClusterEngine
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.IsAuto != nil {
		toSerialize["isAuto"] = o.IsAuto
	}
	if o.ScriptId != nil {
		toSerialize["scriptID"] = o.ScriptId
	}
	if o.ScriptName != nil {
		toSerialize["scriptName"] = o.ScriptName
	}
	if o.ScriptCategory != nil {
		toSerialize["scriptCategory"] = o.ScriptCategory
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Suggestion != nil {
		toSerialize["suggestion"] = o.Suggestion
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Inspection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id             *int32     `json:"id,omitempty"`
		OrgName        *string    `json:"orgName,omitempty"`
		UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
		CreatedAt      *time.Time `json:"createdAt,omitempty"`
		Creator        *string    `json:"creator,omitempty"`
		ClusterId      *string    `json:"clusterID,omitempty"`
		ClusterName    *string    `json:"clusterName,omitempty"`
		ClusterEngine  *string    `json:"clusterEngine,omitempty"`
		Result         *string    `json:"result,omitempty"`
		Severity       *string    `json:"severity,omitempty"`
		Status         *string    `json:"status,omitempty"`
		IsAuto         *bool      `json:"isAuto,omitempty"`
		ScriptId       *int32     `json:"scriptID,omitempty"`
		ScriptName     *string    `json:"scriptName,omitempty"`
		ScriptCategory *string    `json:"scriptCategory,omitempty"`
		Reason         *string    `json:"reason,omitempty"`
		Suggestion     *string    `json:"suggestion,omitempty"`
		Unit           *string    `json:"unit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "orgName", "updatedAt", "createdAt", "creator", "clusterID", "clusterName", "clusterEngine", "result", "severity", "status", "isAuto", "scriptID", "scriptName", "scriptCategory", "reason", "suggestion", "unit"})
	} else {
		return err
	}
	o.Id = all.Id
	o.OrgName = all.OrgName
	o.UpdatedAt = all.UpdatedAt
	o.CreatedAt = all.CreatedAt
	o.Creator = all.Creator
	o.ClusterId = all.ClusterId
	o.ClusterName = all.ClusterName
	o.ClusterEngine = all.ClusterEngine
	o.Result = all.Result
	o.Severity = all.Severity
	o.Status = all.Status
	o.IsAuto = all.IsAuto
	o.ScriptId = all.ScriptId
	o.ScriptName = all.ScriptName
	o.ScriptCategory = all.ScriptCategory
	o.Reason = all.Reason
	o.Suggestion = all.Suggestion
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
