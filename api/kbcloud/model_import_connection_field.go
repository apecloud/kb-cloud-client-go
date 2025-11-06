// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportConnectionField Connection field configuration. Use optional attributes to describe type-specific properties.
type ImportConnectionField struct {
	// Field's programmatic name (e.g., 'db_host')
	Name  *string               `json:"name,omitempty"`
	Label *LocalizedDescription `json:"label,omitempty"`
	// Whether the field is required
	Required *bool `json:"required,omitempty"`
	// Whether it contains sensitive information (e.g., password)
	Sensitive   *bool                 `json:"sensitive,omitempty"`
	Description *LocalizedDescription `json:"description,omitempty"`
	// Placeholder text for the input
	Placeholder *string `json:"placeholder,omitempty"`
	// Import field type
	Type *ImportFieldType `json:"type,omitempty"`
	// Default value applied when the field is omitted. Coerced according to `type`.
	Default NullableImportConnectionFieldDefault `json:"default,omitempty"`
	// Enumerated options when the field type is enum.
	Options []string `json:"options,omitempty"`
	// Additional validation constraints (e.g., `min`, `max`, `pattern`).
	Validation map[string]interface{} `json:"validation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportConnectionField instantiates a new ImportConnectionField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportConnectionField() *ImportConnectionField {
	this := ImportConnectionField{}
	var required bool = false
	this.Required = &required
	var sensitive bool = false
	this.Sensitive = &sensitive
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImportConnectionField) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportConnectionField) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImportConnectionField) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImportConnectionField) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ImportConnectionField) GetLabel() LocalizedDescription {
	if o == nil || o.Label == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportConnectionField) GetLabelOk() (*LocalizedDescription, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ImportConnectionField) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given LocalizedDescription and assigns it to the Label field.
func (o *ImportConnectionField) SetLabel(v LocalizedDescription) {
	o.Label = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ImportConnectionField) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportConnectionField) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ImportConnectionField) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ImportConnectionField) SetRequired(v bool) {
	o.Required = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *ImportConnectionField) GetSensitive() bool {
	if o == nil || o.Sensitive == nil {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportConnectionField) GetSensitiveOk() (*bool, bool) {
	if o == nil || o.Sensitive == nil {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *ImportConnectionField) HasSensitive() bool {
	return o != nil && o.Sensitive != nil
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *ImportConnectionField) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImportConnectionField) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportConnectionField) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImportConnectionField) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *ImportConnectionField) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *ImportConnectionField) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportConnectionField) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *ImportConnectionField) HasPlaceholder() bool {
	return o != nil && o.Placeholder != nil
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *ImportConnectionField) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ImportConnectionField) GetType() ImportFieldType {
	if o == nil || o.Type == nil {
		var ret ImportFieldType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportConnectionField) GetTypeOk() (*ImportFieldType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ImportConnectionField) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ImportFieldType and assigns it to the Type field.
func (o *ImportConnectionField) SetType(v ImportFieldType) {
	o.Type = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ImportConnectionField) GetDefault() ImportConnectionFieldDefault {
	if o == nil || o.Default.Get() == nil {
		var ret ImportConnectionFieldDefault
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportConnectionField) GetDefaultOk() (*ImportConnectionFieldDefault, bool) {
	if o == nil || o.Default.Get() == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *ImportConnectionField) HasDefault() bool {
	return o != nil && o.Default.IsSet()
}

// SetDefault gets a reference to the given NullableImportConnectionFieldDefault and assigns it to the Default field.
func (o *ImportConnectionField) SetDefault(v ImportConnectionFieldDefault) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil.
func (o *ImportConnectionField) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil.
func (o *ImportConnectionField) UnsetDefault() {
	o.Default.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ImportConnectionField) GetOptions() []string {
	if o == nil || o.Options == nil {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportConnectionField) GetOptionsOk() (*[]string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ImportConnectionField) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *ImportConnectionField) SetOptions(v []string) {
	o.Options = v
}

// GetValidation returns the Validation field value if set, zero value otherwise.
func (o *ImportConnectionField) GetValidation() map[string]interface{} {
	if o == nil || o.Validation == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Validation
}

// GetValidationOk returns a tuple with the Validation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportConnectionField) GetValidationOk() (*map[string]interface{}, bool) {
	if o == nil || o.Validation == nil {
		return nil, false
	}
	return &o.Validation, true
}

// HasValidation returns a boolean if a field has been set.
func (o *ImportConnectionField) HasValidation() bool {
	return o != nil && o.Validation != nil
}

// SetValidation gets a reference to the given map[string]interface{} and assigns it to the Validation field.
func (o *ImportConnectionField) SetValidation(v map[string]interface{}) {
	o.Validation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportConnectionField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Sensitive != nil {
		toSerialize["sensitive"] = o.Sensitive
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Placeholder != nil {
		toSerialize["placeholder"] = o.Placeholder
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Validation != nil {
		toSerialize["validation"] = o.Validation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportConnectionField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string                              `json:"name,omitempty"`
		Label       *LocalizedDescription                `json:"label,omitempty"`
		Required    *bool                                `json:"required,omitempty"`
		Sensitive   *bool                                `json:"sensitive,omitempty"`
		Description *LocalizedDescription                `json:"description,omitempty"`
		Placeholder *string                              `json:"placeholder,omitempty"`
		Type        *ImportFieldType                     `json:"type,omitempty"`
		Default     NullableImportConnectionFieldDefault `json:"default,omitempty"`
		Options     []string                             `json:"options,omitempty"`
		Validation  map[string]interface{}               `json:"validation,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "label", "required", "sensitive", "description", "placeholder", "type", "default", "options", "validation"})
	} else {
		return err
	}
	hasInvalidField := false
	o.Name = all.Name
	if all.Label != nil && all.Label.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Label = all.Label
	o.Required = all.Required
	o.Sensitive = all.Sensitive
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	o.Placeholder = all.Placeholder
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Default = all.Default
	o.Options = all.Options
	o.Validation = all.Validation
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}
	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
