// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DatabaseParameterListItem database parameter item
type DatabaseParameterListItem struct {
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

// NewDatabaseParameterListItem instantiates a new DatabaseParameterListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabaseParameterListItem(name string, description string, typeVar DatabaseParameterType, needRestart bool, immutable bool, level DatabaseParameterLevel) *DatabaseParameterListItem {
	this := DatabaseParameterListItem{}
	this.Name = name
	this.Description = description
	this.Type = typeVar
	this.NeedRestart = needRestart
	this.Immutable = immutable
	this.Level = level
	return &this
}

// NewDatabaseParameterListItemWithDefaults instantiates a new DatabaseParameterListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseParameterListItemWithDefaults() *DatabaseParameterListItem {
	this := DatabaseParameterListItem{}
	return &this
}

// GetName returns the Name field value.
func (o *DatabaseParameterListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DatabaseParameterListItem) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value.
func (o *DatabaseParameterListItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *DatabaseParameterListItem) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value.
func (o *DatabaseParameterListItem) GetType() DatabaseParameterType {
	if o == nil {
		var ret DatabaseParameterType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetTypeOk() (*DatabaseParameterType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DatabaseParameterListItem) SetType(v DatabaseParameterType) {
	o.Type = v
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *DatabaseParameterListItem) GetSince() string {
	if o == nil || o.Since == nil {
		var ret string
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetSinceOk() (*string, bool) {
	if o == nil || o.Since == nil {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *DatabaseParameterListItem) HasSince() bool {
	return o != nil && o.Since != nil
}

// SetSince gets a reference to the given string and assigns it to the Since field.
func (o *DatabaseParameterListItem) SetSince(v string) {
	o.Since = &v
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *DatabaseParameterListItem) GetDeprecated() string {
	if o == nil || o.Deprecated == nil {
		var ret string
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetDeprecatedOk() (*string, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *DatabaseParameterListItem) HasDeprecated() bool {
	return o != nil && o.Deprecated != nil
}

// SetDeprecated gets a reference to the given string and assigns it to the Deprecated field.
func (o *DatabaseParameterListItem) SetDeprecated(v string) {
	o.Deprecated = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *DatabaseParameterListItem) GetSection() string {
	if o == nil || o.Section == nil {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetSectionOk() (*string, bool) {
	if o == nil || o.Section == nil {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *DatabaseParameterListItem) HasSection() bool {
	return o != nil && o.Section != nil
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *DatabaseParameterListItem) SetSection(v string) {
	o.Section = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *DatabaseParameterListItem) GetDefault() interface{} {
	if o == nil || o.Default == nil {
		var ret interface{}
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetDefaultOk() (*interface{}, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return &o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *DatabaseParameterListItem) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given interface{} and assigns it to the Default field.
func (o *DatabaseParameterListItem) SetDefault(v interface{}) {
	o.Default = v
}

// GetNeedRestart returns the NeedRestart field value.
func (o *DatabaseParameterListItem) GetNeedRestart() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NeedRestart
}

// GetNeedRestartOk returns a tuple with the NeedRestart field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetNeedRestartOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedRestart, true
}

// SetNeedRestart sets field value.
func (o *DatabaseParameterListItem) SetNeedRestart(v bool) {
	o.NeedRestart = v
}

// GetImmutable returns the Immutable field value.
func (o *DatabaseParameterListItem) GetImmutable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Immutable
}

// GetImmutableOk returns a tuple with the Immutable field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetImmutableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Immutable, true
}

// SetImmutable sets field value.
func (o *DatabaseParameterListItem) SetImmutable(v bool) {
	o.Immutable = v
}

// GetLevel returns the Level field value.
func (o *DatabaseParameterListItem) GetLevel() DatabaseParameterLevel {
	if o == nil {
		var ret DatabaseParameterLevel
		return ret
	}
	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetLevelOk() (*DatabaseParameterLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value.
func (o *DatabaseParameterListItem) SetLevel(v DatabaseParameterLevel) {
	o.Level = v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *DatabaseParameterListItem) GetMaximum() string {
	if o == nil || o.Maximum == nil {
		var ret string
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetMaximumOk() (*string, bool) {
	if o == nil || o.Maximum == nil {
		return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *DatabaseParameterListItem) HasMaximum() bool {
	return o != nil && o.Maximum != nil
}

// SetMaximum gets a reference to the given string and assigns it to the Maximum field.
func (o *DatabaseParameterListItem) SetMaximum(v string) {
	o.Maximum = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *DatabaseParameterListItem) GetMinimum() string {
	if o == nil || o.Minimum == nil {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetMinimumOk() (*string, bool) {
	if o == nil || o.Minimum == nil {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *DatabaseParameterListItem) HasMinimum() bool {
	return o != nil && o.Minimum != nil
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *DatabaseParameterListItem) SetMinimum(v string) {
	o.Minimum = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *DatabaseParameterListItem) GetEnum() []interface{} {
	if o == nil || o.Enum == nil {
		var ret []interface{}
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetEnumOk() (*[]interface{}, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return &o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *DatabaseParameterListItem) HasEnum() bool {
	return o != nil && o.Enum != nil
}

// SetEnum gets a reference to the given []interface{} and assigns it to the Enum field.
func (o *DatabaseParameterListItem) SetEnum(v []interface{}) {
	o.Enum = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *DatabaseParameterListItem) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseParameterListItem) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *DatabaseParameterListItem) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *DatabaseParameterListItem) SetUnit(v string) {
	o.Unit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabaseParameterListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
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
func (o *DatabaseParameterListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
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
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "type", "since", "deprecated", "section", "default", "needRestart", "immutable", "level", "maximum", "minimum", "enum", "unit"})
	} else {
		return err
	}

	hasInvalidField := false
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
