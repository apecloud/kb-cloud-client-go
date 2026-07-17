// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InspectionScript struct {
	Id          *string               `json:"id,omitempty"`
	Name        string                `json:"name"`
	DisplayName *LocalizedDescription `json:"displayName,omitempty"`
	// The engine type this script is applicable to, such as cluster, mysql, node
	Engine string `json:"engine"`
	// The type of the script, such as "promQL"
	Type string `json:"type"`
	// Current inspection category vocabulary. Existing legacy categories are normalized by the backend before they are returned.
	Category    InspectionScriptCategory `json:"category"`
	Description *LocalizedDescription    `json:"description,omitempty"`
	ScriptExpr  string                   `json:"scriptExpr"`
	CheckExpr   *string                  `json:"checkExpr,omitempty"`
	// First-version criticality assumption for score weighting and red-item veto behavior. Missing legacy values are treated as medium.
	Criticality *InspectionCriticality `json:"criticality,omitempty"`
	// First-version warning threshold assumption. It documents the script contract and is not a permanent product conclusion.
	WarnThreshold *float64 `json:"warnThreshold,omitempty"`
	// First-version critical threshold assumption. It documents the script contract and is not a permanent product conclusion.
	CritThreshold *float64 `json:"critThreshold,omitempty"`
	// Direction used to interpret warnThreshold and critThreshold. asc means larger values are worse, desc means smaller values are worse, boolean means the check expression decides.
	Direction   *InspectionThresholdDirection `json:"direction,omitempty"`
	Remediation *LocalizedDescription         `json:"remediation,omitempty"`
	DocLink     *string                       `json:"docLink,omitempty"`
	// scope type, such as "system"/"global"/"org"
	ScopeType string `json:"scopeType"`
	// The identifier of the scope, such as org_id
	ScopeId *string `json:"scopeID,omitempty"`
	// The identifier of the scope, such as org_name
	ScopeName *string `json:"scopeName,omitempty"`
	Enabled   bool    `json:"enabled"`
	Unit      *string `json:"unit,omitempty"`
	// Timestamp of creation time
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// Timestamp of update time
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInspectionScript instantiates a new InspectionScript object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInspectionScript(name string, engine string, typeVar string, category InspectionScriptCategory, scriptExpr string, scopeType string, enabled bool) *InspectionScript {
	this := InspectionScript{}
	this.Name = name
	this.Engine = engine
	this.Type = typeVar
	this.Category = category
	this.ScriptExpr = scriptExpr
	this.ScopeType = scopeType
	this.Enabled = enabled
	return &this
}

// NewInspectionScriptWithDefaults instantiates a new InspectionScript object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInspectionScriptWithDefaults() *InspectionScript {
	this := InspectionScript{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InspectionScript) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InspectionScript) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InspectionScript) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value.
func (o *InspectionScript) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InspectionScript) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InspectionScript) GetDisplayName() LocalizedDescription {
	if o == nil || o.DisplayName == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetDisplayNameOk() (*LocalizedDescription, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InspectionScript) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given LocalizedDescription and assigns it to the DisplayName field.
func (o *InspectionScript) SetDisplayName(v LocalizedDescription) {
	o.DisplayName = &v
}

// GetEngine returns the Engine field value.
func (o *InspectionScript) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *InspectionScript) SetEngine(v string) {
	o.Engine = v
}

// GetType returns the Type field value.
func (o *InspectionScript) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *InspectionScript) SetType(v string) {
	o.Type = v
}

// GetCategory returns the Category field value.
func (o *InspectionScript) GetCategory() InspectionScriptCategory {
	if o == nil {
		var ret InspectionScriptCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetCategoryOk() (*InspectionScriptCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *InspectionScript) SetCategory(v InspectionScriptCategory) {
	o.Category = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InspectionScript) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InspectionScript) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *InspectionScript) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetScriptExpr returns the ScriptExpr field value.
func (o *InspectionScript) GetScriptExpr() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ScriptExpr
}

// GetScriptExprOk returns a tuple with the ScriptExpr field value
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetScriptExprOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScriptExpr, true
}

// SetScriptExpr sets field value.
func (o *InspectionScript) SetScriptExpr(v string) {
	o.ScriptExpr = v
}

// GetCheckExpr returns the CheckExpr field value if set, zero value otherwise.
func (o *InspectionScript) GetCheckExpr() string {
	if o == nil || o.CheckExpr == nil {
		var ret string
		return ret
	}
	return *o.CheckExpr
}

// GetCheckExprOk returns a tuple with the CheckExpr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetCheckExprOk() (*string, bool) {
	if o == nil || o.CheckExpr == nil {
		return nil, false
	}
	return o.CheckExpr, true
}

// HasCheckExpr returns a boolean if a field has been set.
func (o *InspectionScript) HasCheckExpr() bool {
	return o != nil && o.CheckExpr != nil
}

// SetCheckExpr gets a reference to the given string and assigns it to the CheckExpr field.
func (o *InspectionScript) SetCheckExpr(v string) {
	o.CheckExpr = &v
}

// GetCriticality returns the Criticality field value if set, zero value otherwise.
func (o *InspectionScript) GetCriticality() InspectionCriticality {
	if o == nil || o.Criticality == nil {
		var ret InspectionCriticality
		return ret
	}
	return *o.Criticality
}

// GetCriticalityOk returns a tuple with the Criticality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetCriticalityOk() (*InspectionCriticality, bool) {
	if o == nil || o.Criticality == nil {
		return nil, false
	}
	return o.Criticality, true
}

// HasCriticality returns a boolean if a field has been set.
func (o *InspectionScript) HasCriticality() bool {
	return o != nil && o.Criticality != nil
}

// SetCriticality gets a reference to the given InspectionCriticality and assigns it to the Criticality field.
func (o *InspectionScript) SetCriticality(v InspectionCriticality) {
	o.Criticality = &v
}

// GetWarnThreshold returns the WarnThreshold field value if set, zero value otherwise.
func (o *InspectionScript) GetWarnThreshold() float64 {
	if o == nil || o.WarnThreshold == nil {
		var ret float64
		return ret
	}
	return *o.WarnThreshold
}

// GetWarnThresholdOk returns a tuple with the WarnThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetWarnThresholdOk() (*float64, bool) {
	if o == nil || o.WarnThreshold == nil {
		return nil, false
	}
	return o.WarnThreshold, true
}

// HasWarnThreshold returns a boolean if a field has been set.
func (o *InspectionScript) HasWarnThreshold() bool {
	return o != nil && o.WarnThreshold != nil
}

// SetWarnThreshold gets a reference to the given float64 and assigns it to the WarnThreshold field.
func (o *InspectionScript) SetWarnThreshold(v float64) {
	o.WarnThreshold = &v
}

// GetCritThreshold returns the CritThreshold field value if set, zero value otherwise.
func (o *InspectionScript) GetCritThreshold() float64 {
	if o == nil || o.CritThreshold == nil {
		var ret float64
		return ret
	}
	return *o.CritThreshold
}

// GetCritThresholdOk returns a tuple with the CritThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetCritThresholdOk() (*float64, bool) {
	if o == nil || o.CritThreshold == nil {
		return nil, false
	}
	return o.CritThreshold, true
}

// HasCritThreshold returns a boolean if a field has been set.
func (o *InspectionScript) HasCritThreshold() bool {
	return o != nil && o.CritThreshold != nil
}

// SetCritThreshold gets a reference to the given float64 and assigns it to the CritThreshold field.
func (o *InspectionScript) SetCritThreshold(v float64) {
	o.CritThreshold = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *InspectionScript) GetDirection() InspectionThresholdDirection {
	if o == nil || o.Direction == nil {
		var ret InspectionThresholdDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetDirectionOk() (*InspectionThresholdDirection, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *InspectionScript) HasDirection() bool {
	return o != nil && o.Direction != nil
}

// SetDirection gets a reference to the given InspectionThresholdDirection and assigns it to the Direction field.
func (o *InspectionScript) SetDirection(v InspectionThresholdDirection) {
	o.Direction = &v
}

// GetRemediation returns the Remediation field value if set, zero value otherwise.
func (o *InspectionScript) GetRemediation() LocalizedDescription {
	if o == nil || o.Remediation == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Remediation
}

// GetRemediationOk returns a tuple with the Remediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetRemediationOk() (*LocalizedDescription, bool) {
	if o == nil || o.Remediation == nil {
		return nil, false
	}
	return o.Remediation, true
}

// HasRemediation returns a boolean if a field has been set.
func (o *InspectionScript) HasRemediation() bool {
	return o != nil && o.Remediation != nil
}

// SetRemediation gets a reference to the given LocalizedDescription and assigns it to the Remediation field.
func (o *InspectionScript) SetRemediation(v LocalizedDescription) {
	o.Remediation = &v
}

// GetDocLink returns the DocLink field value if set, zero value otherwise.
func (o *InspectionScript) GetDocLink() string {
	if o == nil || o.DocLink == nil {
		var ret string
		return ret
	}
	return *o.DocLink
}

// GetDocLinkOk returns a tuple with the DocLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetDocLinkOk() (*string, bool) {
	if o == nil || o.DocLink == nil {
		return nil, false
	}
	return o.DocLink, true
}

// HasDocLink returns a boolean if a field has been set.
func (o *InspectionScript) HasDocLink() bool {
	return o != nil && o.DocLink != nil
}

// SetDocLink gets a reference to the given string and assigns it to the DocLink field.
func (o *InspectionScript) SetDocLink(v string) {
	o.DocLink = &v
}

// GetScopeType returns the ScopeType field value.
func (o *InspectionScript) GetScopeType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ScopeType
}

// GetScopeTypeOk returns a tuple with the ScopeType field value
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetScopeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeType, true
}

// SetScopeType sets field value.
func (o *InspectionScript) SetScopeType(v string) {
	o.ScopeType = v
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise.
func (o *InspectionScript) GetScopeId() string {
	if o == nil || o.ScopeId == nil {
		var ret string
		return ret
	}
	return *o.ScopeId
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetScopeIdOk() (*string, bool) {
	if o == nil || o.ScopeId == nil {
		return nil, false
	}
	return o.ScopeId, true
}

// HasScopeId returns a boolean if a field has been set.
func (o *InspectionScript) HasScopeId() bool {
	return o != nil && o.ScopeId != nil
}

// SetScopeId gets a reference to the given string and assigns it to the ScopeId field.
func (o *InspectionScript) SetScopeId(v string) {
	o.ScopeId = &v
}

// GetScopeName returns the ScopeName field value if set, zero value otherwise.
func (o *InspectionScript) GetScopeName() string {
	if o == nil || o.ScopeName == nil {
		var ret string
		return ret
	}
	return *o.ScopeName
}

// GetScopeNameOk returns a tuple with the ScopeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetScopeNameOk() (*string, bool) {
	if o == nil || o.ScopeName == nil {
		return nil, false
	}
	return o.ScopeName, true
}

// HasScopeName returns a boolean if a field has been set.
func (o *InspectionScript) HasScopeName() bool {
	return o != nil && o.ScopeName != nil
}

// SetScopeName gets a reference to the given string and assigns it to the ScopeName field.
func (o *InspectionScript) SetScopeName(v string) {
	o.ScopeName = &v
}

// GetEnabled returns the Enabled field value.
func (o *InspectionScript) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *InspectionScript) SetEnabled(v bool) {
	o.Enabled = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *InspectionScript) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *InspectionScript) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *InspectionScript) SetUnit(v string) {
	o.Unit = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InspectionScript) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InspectionScript) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *InspectionScript) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InspectionScript) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InspectionScript) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *InspectionScript) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InspectionScript) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["engine"] = o.Engine
	toSerialize["type"] = o.Type
	toSerialize["category"] = o.Category
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["scriptExpr"] = o.ScriptExpr
	if o.CheckExpr != nil {
		toSerialize["checkExpr"] = o.CheckExpr
	}
	if o.Criticality != nil {
		toSerialize["criticality"] = o.Criticality
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
	if o.Remediation != nil {
		toSerialize["remediation"] = o.Remediation
	}
	if o.DocLink != nil {
		toSerialize["docLink"] = o.DocLink
	}
	toSerialize["scopeType"] = o.ScopeType
	if o.ScopeId != nil {
		toSerialize["scopeID"] = o.ScopeId
	}
	if o.ScopeName != nil {
		toSerialize["scopeName"] = o.ScopeName
	}
	toSerialize["enabled"] = o.Enabled
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InspectionScript) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *string                       `json:"id,omitempty"`
		Name          *string                       `json:"name"`
		DisplayName   *LocalizedDescription         `json:"displayName,omitempty"`
		Engine        *string                       `json:"engine"`
		Type          *string                       `json:"type"`
		Category      *InspectionScriptCategory     `json:"category"`
		Description   *LocalizedDescription         `json:"description,omitempty"`
		ScriptExpr    *string                       `json:"scriptExpr"`
		CheckExpr     *string                       `json:"checkExpr,omitempty"`
		Criticality   *InspectionCriticality        `json:"criticality,omitempty"`
		WarnThreshold *float64                      `json:"warnThreshold,omitempty"`
		CritThreshold *float64                      `json:"critThreshold,omitempty"`
		Direction     *InspectionThresholdDirection `json:"direction,omitempty"`
		Remediation   *LocalizedDescription         `json:"remediation,omitempty"`
		DocLink       *string                       `json:"docLink,omitempty"`
		ScopeType     *string                       `json:"scopeType"`
		ScopeId       *string                       `json:"scopeID,omitempty"`
		ScopeName     *string                       `json:"scopeName,omitempty"`
		Enabled       *bool                         `json:"enabled"`
		Unit          *string                       `json:"unit,omitempty"`
		CreatedAt     *int32                        `json:"createdAt,omitempty"`
		UpdatedAt     *int32                        `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.ScriptExpr == nil {
		return fmt.Errorf("required field scriptExpr missing")
	}
	if all.ScopeType == nil {
		return fmt.Errorf("required field scopeType missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "displayName", "engine", "type", "category", "description", "scriptExpr", "checkExpr", "criticality", "warnThreshold", "critThreshold", "direction", "remediation", "docLink", "scopeType", "scopeID", "scopeName", "enabled", "unit", "createdAt", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = *all.Name
	if all.DisplayName != nil && all.DisplayName.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisplayName = all.DisplayName
	o.Engine = *all.Engine
	o.Type = *all.Type
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	o.ScriptExpr = *all.ScriptExpr
	o.CheckExpr = all.CheckExpr
	if all.Criticality != nil && !all.Criticality.IsValid() {
		hasInvalidField = true
	} else {
		o.Criticality = all.Criticality
	}
	o.WarnThreshold = all.WarnThreshold
	o.CritThreshold = all.CritThreshold
	if all.Direction != nil && !all.Direction.IsValid() {
		hasInvalidField = true
	} else {
		o.Direction = all.Direction
	}
	if all.Remediation != nil && all.Remediation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Remediation = all.Remediation
	o.DocLink = all.DocLink
	o.ScopeType = *all.ScopeType
	o.ScopeId = all.ScopeId
	o.ScopeName = all.ScopeName
	o.Enabled = *all.Enabled
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
