// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DatabaseParameterItem A database parameter item
type DatabaseParameterItem struct {
	// Name of database engine
	EngineName string `json:"engineName"`
	// Name of component
	Component string `json:"component"`
	// The name of the parameter
	Name string `json:"name"`
	// The description of the parameter
	Description string `json:"description"`
	// The type of the parameter value
	Type DatabaseParameterType `json:"type"`
	// Since which version the parameter is supported
	Since *string `json:"since,omitempty"`
	// which version the parameter is deprecated
	Deprecated *string `json:"deprecated,omitempty"`
	// The section of the parameter
	Section *string `json:"section,omitempty"`
	// The value of the parameter, if parameter is not set in tpl, it's value equal to cue default value.
	Default interface{} `json:"default,omitempty"`
	// Whether the parameter requires a restart to take effect
	NeedRestart bool `json:"needRestart"`
	// Whether the parameter is an immutable parameter, immutable parameters cannot be modified
	Immutable bool `json:"immutable"`
	// The level of the parameter
	Level DatabaseParameterLevel `json:"level"`
	// The maximum value of the parameter
	Maximum *string `json:"maximum,omitempty"`
	// The minimum value of the parameter
	Minimum *string `json:"minimum,omitempty"`
	// The value options of the parameter
	Enum []interface{} `json:"enum,omitempty"`
	// The unit of the parameter
	Unit *string `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabaseParameterItem instantiates a new DatabaseParameterItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabaseParameterItem(engineName string, component string, name string, description string, typeVar DatabaseParameterType, needRestart bool, immutable bool, level DatabaseParameterLevel) *DatabaseParameterItem {
	this := DatabaseParameterItem{}
	this.EngineName = engineName
	this.Component = component
	this.Name = name
	this.Description = description
	this.Type = typeVar
	this.NeedRestart = needRestart
	this.Immutable = immutable
	this.Level = level
	return &this
}

// NewDatabaseParameterItemWithDefaults instantiates a new DatabaseParameterItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseParameterItemWithDefaults() *DatabaseParameterItem {
	this := DatabaseParameterItem{}
	return &this
}

// GetEngineName returns the EngineName field value.
func (o *DatabaseParameterItem) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *DatabaseParameterItem) SetEngineName(v string) {
	o.EngineName = v
}

// GetComponent returns the Component field value.
func (o *DatabaseParameterItem) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *DatabaseParameterItem) SetComponent(v string) {
	o.Component = v
}

// GetName returns the Name field value.
func (o *DatabaseParameterItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DatabaseParameterItem) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value.
func (o *DatabaseParameterItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *DatabaseParameterItem) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value.
func (o *DatabaseParameterItem) GetType() DatabaseParameterType {
	if o == nil {
		var ret DatabaseParameterType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetTypeOk() (*DatabaseParameterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DatabaseParameterItem) SetType(v DatabaseParameterType) {
	o.Type = v
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *DatabaseParameterItem) GetSince() string {
	if o == nil || o.Since == nil {
		var ret string
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetSinceOk() (*string, bool) {
	if o == nil || o.Since == nil {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *DatabaseParameterItem) HasSince() bool {
	return o != nil && o.Since != nil
}

// SetSince gets a reference to the given string and assigns it to the Since field.
func (o *DatabaseParameterItem) SetSince(v string) {
	o.Since = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *DatabaseParameterItem) GetDeprecated() string {
	if o == nil || o.Deprecated == nil {
		var ret string
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetDeprecatedOk() (*string, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *DatabaseParameterItem) HasDeprecated() bool {
	return o != nil && o.Deprecated != nil
}

// SetDeprecated gets a reference to the given string and assigns it to the Deprecated field.
func (o *DatabaseParameterItem) SetDeprecated(v string) {
	o.Deprecated = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *DatabaseParameterItem) GetSection() string {
	if o == nil || o.Section == nil {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetSectionOk() (*string, bool) {
	if o == nil || o.Section == nil {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *DatabaseParameterItem) HasSection() bool {
	return o != nil && o.Section != nil
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *DatabaseParameterItem) SetSection(v string) {
	o.Section = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *DatabaseParameterItem) GetDefault() interface{} {
	if o == nil || o.Default == nil {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *DatabaseParameterItem) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *DatabaseParameterItem) SetDefault(v interface{}) {
	o.Default = v
}

// GetNeedRestart returns the NeedRestart field value.
func (o *DatabaseParameterItem) GetNeedRestart() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NeedRestart
}

// GetNeedRestartOk returns a tuple with the NeedRestart field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetNeedRestartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedRestart, true
}

// SetNeedRestart sets field value.
func (o *DatabaseParameterItem) SetNeedRestart(v bool) {
	o.NeedRestart = v
}

// GetImmutable returns the Immutable field value.
func (o *DatabaseParameterItem) GetImmutable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Immutable
}

// GetImmutableOk returns a tuple with the Immutable field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetImmutableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Immutable, true
}

// SetImmutable sets field value.
func (o *DatabaseParameterItem) SetImmutable(v bool) {
	o.Immutable = v
}

// GetLevel returns the Level field value.
func (o *DatabaseParameterItem) GetLevel() DatabaseParameterLevel {
	if o == nil {
		var ret DatabaseParameterLevel
		return ret
	}
	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetLevelOk() (*DatabaseParameterLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value.
func (o *DatabaseParameterItem) SetLevel(v DatabaseParameterLevel) {
	o.Level = v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *DatabaseParameterItem) GetMaximum() string {
	if o == nil || o.Maximum == nil {
		var ret string
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetMaximumOk() (*string, bool) {
	if o == nil || o.Maximum == nil {
		return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *DatabaseParameterItem) HasMaximum() bool {
	return o != nil && o.Maximum != nil
}

// SetMaximum gets a reference to the given string and assigns it to the Maximum field.
func (o *DatabaseParameterItem) SetMaximum(v string) {
	o.Maximum = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *DatabaseParameterItem) GetMinimum() string {
	if o == nil || o.Minimum == nil {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetMinimumOk() (*string, bool) {
	if o == nil || o.Minimum == nil {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *DatabaseParameterItem) HasMinimum() bool {
	return o != nil && o.Minimum != nil
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *DatabaseParameterItem) SetMinimum(v string) {
	o.Minimum = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *DatabaseParameterItem) GetEnum() []interface{} {
	if o == nil || o.Enum == nil {
		var ret []interface{}
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetEnumOk() (*[]interface{}, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return &o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *DatabaseParameterItem) HasEnum() bool {
	return o != nil && o.Enum != nil
}

// SetEnum gets a reference to the given []interface{} and assigns it to the Enum field.
func (o *DatabaseParameterItem) SetEnum(v []interface{}) {
	o.Enum = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *DatabaseParameterItem) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterItem) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *DatabaseParameterItem) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *DatabaseParameterItem) SetUnit(v string) {
	o.Unit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabaseParameterItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineName"] = o.EngineName
	toSerialize["component"] = o.Component
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["type"] = o.Type
	if o.Since != nil {
		toSerialize["since"] = o.Since
	}
	if o.Deprecated != nil {
		toSerialize["deprecated"] = o.Deprecated
	}
	if o.Section != nil {
		toSerialize["section"] = o.Section
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	toSerialize["needRestart"] = o.NeedRestart
	toSerialize["immutable"] = o.Immutable
	toSerialize["level"] = o.Level
	if o.Maximum != nil {
		toSerialize["maximum"] = o.Maximum
	}
	if o.Minimum != nil {
		toSerialize["minimum"] = o.Minimum
	}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
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
func (o *DatabaseParameterItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineName  *string                 `json:"engineName"`
		Component   *string                 `json:"component"`
		Name        *string                 `json:"name"`
		Description *string                 `json:"description"`
		Type        *DatabaseParameterType  `json:"type"`
		Since       *string                 `json:"since,omitempty"`
		Deprecated  *string                 `json:"deprecated,omitempty"`
		Section     *string                 `json:"section,omitempty"`
		Default     interface{}             `json:"default,omitempty"`
		NeedRestart *bool                   `json:"needRestart"`
		Immutable   *bool                   `json:"immutable"`
		Level       *DatabaseParameterLevel `json:"level"`
		Maximum     *string                 `json:"maximum,omitempty"`
		Minimum     *string                 `json:"minimum,omitempty"`
		Enum        []interface{}           `json:"enum,omitempty"`
		Unit        *string                 `json:"unit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.NeedRestart == nil {
		return fmt.Errorf("required field needRestart missing")
	}
	if all.Immutable == nil {
		return fmt.Errorf("required field immutable missing")
	}
	if all.Level == nil {
		return fmt.Errorf("required field level missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineName", "component", "name", "description", "type", "since", "deprecated", "section", "default", "needRestart", "immutable", "level", "maximum", "minimum", "enum", "unit"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EngineName = *all.EngineName
	o.Component = *all.Component
	o.Name = *all.Name
	o.Description = *all.Description
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Since = all.Since
	o.Deprecated = all.Deprecated
	o.Section = all.Section
	o.Default = all.Default
	o.NeedRestart = *all.NeedRestart
	o.Immutable = *all.Immutable
	if !all.Level.IsValid() {
		hasInvalidField = true
	} else {
		o.Level = *all.Level
	}
	o.Maximum = all.Maximum
	o.Minimum = all.Minimum
	o.Enum = all.Enum
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
