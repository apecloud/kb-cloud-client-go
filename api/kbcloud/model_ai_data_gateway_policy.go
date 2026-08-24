// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGatewayPolicy struct {
	Name           *string                `json:"name,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Status         *string                `json:"status,omitempty"`
	Priority       *int32                 `json:"priority,omitempty"`
	Decision       *string                `json:"decision,omitempty"`
	Scope          map[string]interface{} `json:"scope,omitempty"`
	SqlTypes       []string               `json:"sqlTypes,omitempty"`
	RiskRules      map[string]interface{} `json:"riskRules,omitempty"`
	MaxRows        *int32                 `json:"maxRows,omitempty"`
	TimeoutSeconds *int32                 `json:"timeoutSeconds,omitempty"`
	PolicyId       *string                `json:"policyId,omitempty"`
	GatewayId      *string                `json:"gatewayId,omitempty"`
	OrgName        *string                `json:"orgName,omitempty"`
	CreatedBy      *string                `json:"createdBy,omitempty"`
	CreatedAt      *time.Time             `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time             `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewayPolicy instantiates a new AiDataGatewayPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewayPolicy() *AiDataGatewayPolicy {
	this := AiDataGatewayPolicy{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AiDataGatewayPolicy) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AiDataGatewayPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiDataGatewayPolicy) SetStatus(v string) {
	o.Status = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *AiDataGatewayPolicy) SetPriority(v int32) {
	o.Priority = &v
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetDecision() string {
	if o == nil || o.Decision == nil {
		var ret string
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetDecisionOk() (*string, bool) {
	if o == nil || o.Decision == nil {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasDecision() bool {
	return o != nil && o.Decision != nil
}

// SetDecision gets a reference to the given string and assigns it to the Decision field.
func (o *AiDataGatewayPolicy) SetDecision(v string) {
	o.Decision = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetScope() map[string]interface{} {
	if o == nil || o.Scope == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetScopeOk() (*map[string]interface{}, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given map[string]interface{} and assigns it to the Scope field.
func (o *AiDataGatewayPolicy) SetScope(v map[string]interface{}) {
	o.Scope = v
}

// GetSqlTypes returns the SqlTypes field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetSqlTypes() []string {
	if o == nil || o.SqlTypes == nil {
		var ret []string
		return ret
	}
	return o.SqlTypes
}

// GetSqlTypesOk returns a tuple with the SqlTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetSqlTypesOk() (*[]string, bool) {
	if o == nil || o.SqlTypes == nil {
		return nil, false
	}
	return &o.SqlTypes, true
}

// HasSqlTypes returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasSqlTypes() bool {
	return o != nil && o.SqlTypes != nil
}

// SetSqlTypes gets a reference to the given []string and assigns it to the SqlTypes field.
func (o *AiDataGatewayPolicy) SetSqlTypes(v []string) {
	o.SqlTypes = v
}

// GetRiskRules returns the RiskRules field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetRiskRules() map[string]interface{} {
	if o == nil || o.RiskRules == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RiskRules
}

// GetRiskRulesOk returns a tuple with the RiskRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetRiskRulesOk() (*map[string]interface{}, bool) {
	if o == nil || o.RiskRules == nil {
		return nil, false
	}
	return &o.RiskRules, true
}

// HasRiskRules returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasRiskRules() bool {
	return o != nil && o.RiskRules != nil
}

// SetRiskRules gets a reference to the given map[string]interface{} and assigns it to the RiskRules field.
func (o *AiDataGatewayPolicy) SetRiskRules(v map[string]interface{}) {
	o.RiskRules = v
}

// GetMaxRows returns the MaxRows field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetMaxRows() int32 {
	if o == nil || o.MaxRows == nil {
		var ret int32
		return ret
	}
	return *o.MaxRows
}

// GetMaxRowsOk returns a tuple with the MaxRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetMaxRowsOk() (*int32, bool) {
	if o == nil || o.MaxRows == nil {
		return nil, false
	}
	return o.MaxRows, true
}

// HasMaxRows returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasMaxRows() bool {
	return o != nil && o.MaxRows != nil
}

// SetMaxRows gets a reference to the given int32 and assigns it to the MaxRows field.
func (o *AiDataGatewayPolicy) SetMaxRows(v int32) {
	o.MaxRows = &v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetTimeoutSeconds() int32 {
	if o == nil || o.TimeoutSeconds == nil {
		var ret int32
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil || o.TimeoutSeconds == nil {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasTimeoutSeconds() bool {
	return o != nil && o.TimeoutSeconds != nil
}

// SetTimeoutSeconds gets a reference to the given int32 and assigns it to the TimeoutSeconds field.
func (o *AiDataGatewayPolicy) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasPolicyId() bool {
	return o != nil && o.PolicyId != nil
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *AiDataGatewayPolicy) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasGatewayId() bool {
	return o != nil && o.GatewayId != nil
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *AiDataGatewayPolicy) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AiDataGatewayPolicy) SetOrgName(v string) {
	o.OrgName = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AiDataGatewayPolicy) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AiDataGatewayPolicy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayPolicy) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayPolicy) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AiDataGatewayPolicy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewayPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Decision != nil {
		toSerialize["decision"] = o.Decision
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.SqlTypes != nil {
		toSerialize["sqlTypes"] = o.SqlTypes
	}
	if o.RiskRules != nil {
		toSerialize["riskRules"] = o.RiskRules
	}
	if o.MaxRows != nil {
		toSerialize["maxRows"] = o.MaxRows
	}
	if o.TimeoutSeconds != nil {
		toSerialize["timeoutSeconds"] = o.TimeoutSeconds
	}
	if o.PolicyId != nil {
		toSerialize["policyId"] = o.PolicyId
	}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
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
func (o *AiDataGatewayPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string                `json:"name,omitempty"`
		Description    *string                `json:"description,omitempty"`
		Status         *string                `json:"status,omitempty"`
		Priority       *int32                 `json:"priority,omitempty"`
		Decision       *string                `json:"decision,omitempty"`
		Scope          map[string]interface{} `json:"scope,omitempty"`
		SqlTypes       []string               `json:"sqlTypes,omitempty"`
		RiskRules      map[string]interface{} `json:"riskRules,omitempty"`
		MaxRows        *int32                 `json:"maxRows,omitempty"`
		TimeoutSeconds *int32                 `json:"timeoutSeconds,omitempty"`
		PolicyId       *string                `json:"policyId,omitempty"`
		GatewayId      *string                `json:"gatewayId,omitempty"`
		OrgName        *string                `json:"orgName,omitempty"`
		CreatedBy      *string                `json:"createdBy,omitempty"`
		CreatedAt      *time.Time             `json:"createdAt,omitempty"`
		UpdatedAt      *time.Time             `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "status", "priority", "decision", "scope", "sqlTypes", "riskRules", "maxRows", "timeoutSeconds", "policyId", "gatewayId", "orgName", "createdBy", "createdAt", "updatedAt"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Description = all.Description
	o.Status = all.Status
	o.Priority = all.Priority
	o.Decision = all.Decision
	o.Scope = all.Scope
	o.SqlTypes = all.SqlTypes
	o.RiskRules = all.RiskRules
	o.MaxRows = all.MaxRows
	o.TimeoutSeconds = all.TimeoutSeconds
	o.PolicyId = all.PolicyId
	o.GatewayId = all.GatewayId
	o.OrgName = all.OrgName
	o.CreatedBy = all.CreatedBy
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
