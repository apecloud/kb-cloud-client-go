// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

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
	// Specifies the category of the inspection script.
	Category    InspectionScriptCategory `json:"category"`
	Description *LocalizedDescription    `json:"description,omitempty"`
	ScriptExpr  string                   `json:"scriptExpr"`
	CheckExpr   *string                  `json:"checkExpr,omitempty"`
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
		Id          *string                   `json:"id,omitempty"`
		Name        *string                   `json:"name"`
		DisplayName *LocalizedDescription     `json:"displayName,omitempty"`
		Engine      *string                   `json:"engine"`
		Type        *string                   `json:"type"`
		Category    *InspectionScriptCategory `json:"category"`
		Description *LocalizedDescription     `json:"description,omitempty"`
		ScriptExpr  *string                   `json:"scriptExpr"`
		CheckExpr   *string                   `json:"checkExpr,omitempty"`
		ScopeType   *string                   `json:"scopeType"`
		ScopeId     *string                   `json:"scopeID,omitempty"`
		ScopeName   *string                   `json:"scopeName,omitempty"`
		Enabled     *bool                     `json:"enabled"`
		Unit        *string                   `json:"unit,omitempty"`
		CreatedAt   *int32                    `json:"createdAt,omitempty"`
		UpdatedAt   *int32                    `json:"updatedAt,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "displayName", "engine", "type", "category", "description", "scriptExpr", "checkExpr", "scopeType", "scopeID", "scopeName", "enabled", "unit", "createdAt", "updatedAt"})
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
