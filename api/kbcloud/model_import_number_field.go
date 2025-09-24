// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportNumberField Configuration for a number-type field.
type ImportNumberField struct {
	// Field's programmatic name (e.g., 'db_host')
	Name string `json:"name,omitempty"`
	// User-facing label for UI (e.g., 'Database Host')
	Label string `json:"label,omitempty"`
	// Whether the field is required
	Required *bool `json:"required,omitempty"`
	// Whether it contains sensitive information (e.g., password)
	Sensitive *bool `json:"sensitive,omitempty"`
	// Detailed usage instructions for the field
	Description *string `json:"description,omitempty"`
	// Placeholder text for the input
	Placeholder *string `json:"placeholder,omitempty"`
	// Import field type
	Type    ImportFieldType        `json:"type,omitempty"`
	Default common.NullableFloat64 `json:"default,omitempty"`
	// Validation rules for numeric field type
	Validation *ImportNumericValidation `json:"validation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportNumberField instantiates a new ImportNumberField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportNumberField() *ImportNumberField {
	this := ImportNumberField{}
	var required bool = false
	this.Required = &required
	var sensitive bool = false
	this.Sensitive = &sensitive
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImportNumberField) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumberField) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImportNumberField) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImportNumberField) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ImportNumberField) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumberField) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return &o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ImportNumberField) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ImportNumberField) SetLabel(v string) {
	o.Label = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ImportNumberField) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumberField) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ImportNumberField) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ImportNumberField) SetRequired(v bool) {
	o.Required = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *ImportNumberField) GetSensitive() bool {
	if o == nil || o.Sensitive == nil {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumberField) GetSensitiveOk() (*bool, bool) {
	if o == nil || o.Sensitive == nil {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *ImportNumberField) HasSensitive() bool {
	return o != nil && o.Sensitive != nil
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *ImportNumberField) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImportNumberField) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumberField) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImportNumberField) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ImportNumberField) SetDescription(v string) {
	o.Description = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *ImportNumberField) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumberField) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *ImportNumberField) HasPlaceholder() bool {
	return o != nil && o.Placeholder != nil
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *ImportNumberField) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ImportNumberField) GetType() ImportFieldType {
	if o == nil || o.Type == nil {
		var ret ImportFieldType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumberField) GetTypeOk() (*ImportFieldType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ImportNumberField) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ImportFieldType and assigns it to the Type field.
func (o *ImportNumberField) SetType(v ImportFieldType) {
	o.Type = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ImportNumberField) GetDefault() float64 {
	if o == nil || o.Default.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumberField) GetDefaultOk() (*float64, bool) {
	if o == nil || o.Default.Get() == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *ImportNumberField) HasDefault() bool {
	return o != nil && o.Default.IsSet()
}

// SetDefault gets a reference to the given common.NullableFloat64 and assigns it to the Default field.
func (o *ImportNumberField) SetDefault(v float64) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil.
func (o *ImportNumberField) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil.
func (o *ImportNumberField) UnsetDefault() {
	o.Default.Unset()
}

// GetValidation returns the Validation field value if set, zero value otherwise.
func (o *ImportNumberField) GetValidation() ImportNumericValidation {
	if o == nil || o.Validation == nil {
		var ret ImportNumericValidation
		return ret
	}
	return *o.Validation
}

// GetValidationOk returns a tuple with the Validation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumberField) GetValidationOk() (*ImportNumericValidation, bool) {
	if o == nil || o.Validation == nil {
		return nil, false
	}
	return o.Validation, true
}

// HasValidation returns a boolean if a field has been set.
func (o *ImportNumberField) HasValidation() bool {
	return o != nil && o.Validation != nil
}

// SetValidation gets a reference to the given ImportNumericValidation and assigns it to the Validation field.
func (o *ImportNumberField) SetValidation(v ImportNumericValidation) {
	o.Validation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportNumberField) MarshalJSON() ([]byte, error) {
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
	if o.Validation != nil {
		toSerialize["validation"] = o.Validation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportNumberField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        string                   `json:"name,omitempty"`
		Label       string                   `json:"label,omitempty"`
		Required    *bool                    `json:"required,omitempty"`
		Sensitive   *bool                    `json:"sensitive,omitempty"`
		Description *string                  `json:"description,omitempty"`
		Placeholder *string                  `json:"placeholder,omitempty"`
		Type        ImportFieldType          `json:"type,omitempty"`
		Default     common.NullableFloat64   `json:"default,omitempty"`
		Validation  *ImportNumericValidation `json:"validation,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "label", "required", "sensitive", "description", "placeholder", "type", "default", "validation"})
	} else {
		return err
	}
	hasInvalidField := false
	o.Name = all.Name
	o.Label = all.Label
	o.Required = all.Required
	o.Sensitive = all.Sensitive
	o.Description = all.Description
	o.Placeholder = all.Placeholder
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Default = all.Default
	if all.Validation != nil && all.Validation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Validation = all.Validation
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}
	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
