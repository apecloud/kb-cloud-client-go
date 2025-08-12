// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InspectionScript struct {
	Id              *int32  `json:"id,omitempty"`
	Engine          *string `json:"engine,omitempty"`
	Name            *string `json:"name,omitempty"`
	Category        *string `json:"category,omitempty"`
	Type            *string `json:"type,omitempty"`
	ScriptType      *string `json:"scriptType,omitempty"`
	Reason          *string `json:"reason,omitempty"`
	Suggestion      *string `json:"suggestion,omitempty"`
	Enabled         bool    `json:"enabled"`
	ScriptName      *string `json:"scriptName,omitempty"`
	Script          *string `json:"script,omitempty"`
	StatusCheckName *string `json:"statusCheckName,omitempty"`
	StatusCheck     *string `json:"statusCheck,omitempty"`
	Unit            *string `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInspectionScript instantiates a new InspectionScript object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInspectionScript(enabled bool) *InspectionScript {
	this := InspectionScript{}
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
func (o *InspectionScript) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InspectionScript) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InspectionScript) SetId(v int32) {
	o.Id = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *InspectionScript) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *InspectionScript) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *InspectionScript) SetEngine(v string) {
	o.Engine = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InspectionScript) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InspectionScript) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InspectionScript) SetName(v string) {
	o.Name = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InspectionScript) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InspectionScript) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InspectionScript) SetCategory(v string) {
	o.Category = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InspectionScript) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InspectionScript) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InspectionScript) SetType(v string) {
	o.Type = &v
}

// GetScriptType returns the ScriptType field value if set, zero value otherwise.
func (o *InspectionScript) GetScriptType() string {
	if o == nil || o.ScriptType == nil {
		var ret string
		return ret
	}
	return *o.ScriptType
}

// GetScriptTypeOk returns a tuple with the ScriptType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetScriptTypeOk() (*string, bool) {
	if o == nil || o.ScriptType == nil {
		return nil, false
	}
	return o.ScriptType, true
}

// HasScriptType returns a boolean if a field has been set.
func (o *InspectionScript) HasScriptType() bool {
	return o != nil && o.ScriptType != nil
}

// SetScriptType gets a reference to the given string and assigns it to the ScriptType field.
func (o *InspectionScript) SetScriptType(v string) {
	o.ScriptType = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InspectionScript) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InspectionScript) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InspectionScript) SetReason(v string) {
	o.Reason = &v
}

// GetSuggestion returns the Suggestion field value if set, zero value otherwise.
func (o *InspectionScript) GetSuggestion() string {
	if o == nil || o.Suggestion == nil {
		var ret string
		return ret
	}
	return *o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetSuggestionOk() (*string, bool) {
	if o == nil || o.Suggestion == nil {
		return nil, false
	}
	return o.Suggestion, true
}

// HasSuggestion returns a boolean if a field has been set.
func (o *InspectionScript) HasSuggestion() bool {
	return o != nil && o.Suggestion != nil
}

// SetSuggestion gets a reference to the given string and assigns it to the Suggestion field.
func (o *InspectionScript) SetSuggestion(v string) {
	o.Suggestion = &v
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

// GetScriptName returns the ScriptName field value if set, zero value otherwise.
func (o *InspectionScript) GetScriptName() string {
	if o == nil || o.ScriptName == nil {
		var ret string
		return ret
	}
	return *o.ScriptName
}

// GetScriptNameOk returns a tuple with the ScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetScriptNameOk() (*string, bool) {
	if o == nil || o.ScriptName == nil {
		return nil, false
	}
	return o.ScriptName, true
}

// HasScriptName returns a boolean if a field has been set.
func (o *InspectionScript) HasScriptName() bool {
	return o != nil && o.ScriptName != nil
}

// SetScriptName gets a reference to the given string and assigns it to the ScriptName field.
func (o *InspectionScript) SetScriptName(v string) {
	o.ScriptName = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *InspectionScript) GetScript() string {
	if o == nil || o.Script == nil {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetScriptOk() (*string, bool) {
	if o == nil || o.Script == nil {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *InspectionScript) HasScript() bool {
	return o != nil && o.Script != nil
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *InspectionScript) SetScript(v string) {
	o.Script = &v
}

// GetStatusCheckName returns the StatusCheckName field value if set, zero value otherwise.
func (o *InspectionScript) GetStatusCheckName() string {
	if o == nil || o.StatusCheckName == nil {
		var ret string
		return ret
	}
	return *o.StatusCheckName
}

// GetStatusCheckNameOk returns a tuple with the StatusCheckName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetStatusCheckNameOk() (*string, bool) {
	if o == nil || o.StatusCheckName == nil {
		return nil, false
	}
	return o.StatusCheckName, true
}

// HasStatusCheckName returns a boolean if a field has been set.
func (o *InspectionScript) HasStatusCheckName() bool {
	return o != nil && o.StatusCheckName != nil
}

// SetStatusCheckName gets a reference to the given string and assigns it to the StatusCheckName field.
func (o *InspectionScript) SetStatusCheckName(v string) {
	o.StatusCheckName = &v
}

// GetStatusCheck returns the StatusCheck field value if set, zero value otherwise.
func (o *InspectionScript) GetStatusCheck() string {
	if o == nil || o.StatusCheck == nil {
		var ret string
		return ret
	}
	return *o.StatusCheck
}

// GetStatusCheckOk returns a tuple with the StatusCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionScript) GetStatusCheckOk() (*string, bool) {
	if o == nil || o.StatusCheck == nil {
		return nil, false
	}
	return o.StatusCheck, true
}

// HasStatusCheck returns a boolean if a field has been set.
func (o *InspectionScript) HasStatusCheck() bool {
	return o != nil && o.StatusCheck != nil
}

// SetStatusCheck gets a reference to the given string and assigns it to the StatusCheck field.
func (o *InspectionScript) SetStatusCheck(v string) {
	o.StatusCheck = &v
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

// MarshalJSON serializes the struct using spec logic.
func (o InspectionScript) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ScriptType != nil {
		toSerialize["scriptType"] = o.ScriptType
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Suggestion != nil {
		toSerialize["suggestion"] = o.Suggestion
	}
	toSerialize["enabled"] = o.Enabled
	if o.ScriptName != nil {
		toSerialize["scriptName"] = o.ScriptName
	}
	if o.Script != nil {
		toSerialize["script"] = o.Script
	}
	if o.StatusCheckName != nil {
		toSerialize["statusCheckName"] = o.StatusCheckName
	}
	if o.StatusCheck != nil {
		toSerialize["statusCheck"] = o.StatusCheck
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
func (o *InspectionScript) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id              *int32  `json:"id,omitempty"`
		Engine          *string `json:"engine,omitempty"`
		Name            *string `json:"name,omitempty"`
		Category        *string `json:"category,omitempty"`
		Type            *string `json:"type,omitempty"`
		ScriptType      *string `json:"scriptType,omitempty"`
		Reason          *string `json:"reason,omitempty"`
		Suggestion      *string `json:"suggestion,omitempty"`
		Enabled         *bool   `json:"enabled"`
		ScriptName      *string `json:"scriptName,omitempty"`
		Script          *string `json:"script,omitempty"`
		StatusCheckName *string `json:"statusCheckName,omitempty"`
		StatusCheck     *string `json:"statusCheck,omitempty"`
		Unit            *string `json:"unit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "engine", "name", "category", "type", "scriptType", "reason", "suggestion", "enabled", "scriptName", "script", "statusCheckName", "statusCheck", "unit"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Engine = all.Engine
	o.Name = all.Name
	o.Category = all.Category
	o.Type = all.Type
	o.ScriptType = all.ScriptType
	o.Reason = all.Reason
	o.Suggestion = all.Suggestion
	o.Enabled = *all.Enabled
	o.ScriptName = all.ScriptName
	o.Script = all.Script
	o.StatusCheckName = all.StatusCheckName
	o.StatusCheck = all.StatusCheck
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
