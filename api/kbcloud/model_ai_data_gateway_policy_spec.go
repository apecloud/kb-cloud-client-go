// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGatewayPolicySpec struct {
	Name           string                 `json:"name"`
	Description    *string                `json:"description,omitempty"`
	Status         *string                `json:"status,omitempty"`
	Priority       *int32                 `json:"priority,omitempty"`
	Decision       string                 `json:"decision"`
	Scope          map[string]interface{} `json:"scope,omitempty"`
	SqlTypes       []string               `json:"sqlTypes,omitempty"`
	RiskRules      map[string]interface{} `json:"riskRules,omitempty"`
	MaxRows        *int32                 `json:"maxRows,omitempty"`
	TimeoutSeconds *int32                 `json:"timeoutSeconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewayPolicySpec instantiates a new AiDataGatewayPolicySpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewayPolicySpec(name string, decision string) *AiDataGatewayPolicySpec {
	this := AiDataGatewayPolicySpec{}
	this.Name = name
	this.Decision = decision
	return &this
}

// NewAiDataGatewayPolicySpecWithDefaults instantiates a new AiDataGatewayPolicySpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiDataGatewayPolicySpecWithDefaults() *AiDataGatewayPolicySpec {
	this := AiDataGatewayPolicySpec{}
	return &this
}

// GetName returns the Name field value.
func (o *AiDataGatewayPolicySpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicySpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiDataGatewayPolicySpec) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AiDataGatewayPolicySpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicySpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AiDataGatewayPolicySpec) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AiDataGatewayPolicySpec) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiDataGatewayPolicySpec) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicySpec) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiDataGatewayPolicySpec) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiDataGatewayPolicySpec) SetStatus(v string) {
	o.Status = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AiDataGatewayPolicySpec) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicySpec) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AiDataGatewayPolicySpec) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *AiDataGatewayPolicySpec) SetPriority(v int32) {
	o.Priority = &v
}

// GetDecision returns the Decision field value.
func (o *AiDataGatewayPolicySpec) GetDecision() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicySpec) GetDecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value.
func (o *AiDataGatewayPolicySpec) SetDecision(v string) {
	o.Decision = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AiDataGatewayPolicySpec) GetScope() map[string]interface{} {
	if o == nil || o.Scope == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicySpec) GetScopeOk() (*map[string]interface{}, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AiDataGatewayPolicySpec) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given map[string]interface{} and assigns it to the Scope field.
func (o *AiDataGatewayPolicySpec) SetScope(v map[string]interface{}) {
	o.Scope = v
}

// GetSqlTypes returns the SqlTypes field value if set, zero value otherwise.
func (o *AiDataGatewayPolicySpec) GetSqlTypes() []string {
	if o == nil || o.SqlTypes == nil {
		var ret []string
		return ret
	}
	return o.SqlTypes
}

// GetSqlTypesOk returns a tuple with the SqlTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicySpec) GetSqlTypesOk() (*[]string, bool) {
	if o == nil || o.SqlTypes == nil {
		return nil, false
	}
	return &o.SqlTypes, true
}

// HasSqlTypes returns a boolean if a field has been set.
func (o *AiDataGatewayPolicySpec) HasSqlTypes() bool {
	return o != nil && o.SqlTypes != nil
}

// SetSqlTypes gets a reference to the given []string and assigns it to the SqlTypes field.
func (o *AiDataGatewayPolicySpec) SetSqlTypes(v []string) {
	o.SqlTypes = v
}

// GetRiskRules returns the RiskRules field value if set, zero value otherwise.
func (o *AiDataGatewayPolicySpec) GetRiskRules() map[string]interface{} {
	if o == nil || o.RiskRules == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RiskRules
}

// GetRiskRulesOk returns a tuple with the RiskRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicySpec) GetRiskRulesOk() (*map[string]interface{}, bool) {
	if o == nil || o.RiskRules == nil {
		return nil, false
	}
	return &o.RiskRules, true
}

// HasRiskRules returns a boolean if a field has been set.
func (o *AiDataGatewayPolicySpec) HasRiskRules() bool {
	return o != nil && o.RiskRules != nil
}

// SetRiskRules gets a reference to the given map[string]interface{} and assigns it to the RiskRules field.
func (o *AiDataGatewayPolicySpec) SetRiskRules(v map[string]interface{}) {
	o.RiskRules = v
}

// GetMaxRows returns the MaxRows field value if set, zero value otherwise.
func (o *AiDataGatewayPolicySpec) GetMaxRows() int32 {
	if o == nil || o.MaxRows == nil {
		var ret int32
		return ret
	}
	return *o.MaxRows
}

// GetMaxRowsOk returns a tuple with the MaxRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicySpec) GetMaxRowsOk() (*int32, bool) {
	if o == nil || o.MaxRows == nil {
		return nil, false
	}
	return o.MaxRows, true
}

// HasMaxRows returns a boolean if a field has been set.
func (o *AiDataGatewayPolicySpec) HasMaxRows() bool {
	return o != nil && o.MaxRows != nil
}

// SetMaxRows gets a reference to the given int32 and assigns it to the MaxRows field.
func (o *AiDataGatewayPolicySpec) SetMaxRows(v int32) {
	o.MaxRows = &v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *AiDataGatewayPolicySpec) GetTimeoutSeconds() int32 {
	if o == nil || o.TimeoutSeconds == nil {
		var ret int32
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayPolicySpec) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil || o.TimeoutSeconds == nil {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *AiDataGatewayPolicySpec) HasTimeoutSeconds() bool {
	return o != nil && o.TimeoutSeconds != nil
}

// SetTimeoutSeconds gets a reference to the given int32 and assigns it to the TimeoutSeconds field.
func (o *AiDataGatewayPolicySpec) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewayPolicySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	toSerialize["decision"] = o.Decision
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewayPolicySpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string                `json:"name"`
		Description    *string                `json:"description,omitempty"`
		Status         *string                `json:"status,omitempty"`
		Priority       *int32                 `json:"priority,omitempty"`
		Decision       *string                `json:"decision"`
		Scope          map[string]interface{} `json:"scope,omitempty"`
		SqlTypes       []string               `json:"sqlTypes,omitempty"`
		RiskRules      map[string]interface{} `json:"riskRules,omitempty"`
		MaxRows        *int32                 `json:"maxRows,omitempty"`
		TimeoutSeconds *int32                 `json:"timeoutSeconds,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Decision == nil {
		return fmt.Errorf("required field decision missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "status", "priority", "decision", "scope", "sqlTypes", "riskRules", "maxRows", "timeoutSeconds"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Description = all.Description
	o.Status = all.Status
	o.Priority = all.Priority
	o.Decision = *all.Decision
	o.Scope = all.Scope
	o.SqlTypes = all.SqlTypes
	o.RiskRules = all.RiskRules
	o.MaxRows = all.MaxRows
	o.TimeoutSeconds = all.TimeoutSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
