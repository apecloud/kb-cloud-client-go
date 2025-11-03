// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportEnumField Configuration for an enum-type field.
type ImportEnumField struct {
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
	// List of available options for the enum.
	Options []string              `json:"options,omitempty"`
	Default common.NullableString `json:"default,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportEnumField instantiates a new ImportEnumField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportEnumField() *ImportEnumField {
	this := ImportEnumField{}
	var required bool = false
	this.Required = &required
	var sensitive bool = false
	this.Sensitive = &sensitive
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImportEnumField) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEnumField) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImportEnumField) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImportEnumField) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ImportEnumField) GetLabel() LocalizedDescription {
	if o == nil || o.Label == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEnumField) GetLabelOk() (*LocalizedDescription, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ImportEnumField) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given LocalizedDescription and assigns it to the Label field.
func (o *ImportEnumField) SetLabel(v LocalizedDescription) {
	o.Label = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ImportEnumField) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEnumField) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ImportEnumField) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ImportEnumField) SetRequired(v bool) {
	o.Required = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *ImportEnumField) GetSensitive() bool {
	if o == nil || o.Sensitive == nil {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEnumField) GetSensitiveOk() (*bool, bool) {
	if o == nil || o.Sensitive == nil {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *ImportEnumField) HasSensitive() bool {
	return o != nil && o.Sensitive != nil
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *ImportEnumField) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImportEnumField) GetDescription() LocalizedDescription {
	if o == nil || o.Description == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEnumField) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImportEnumField) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given LocalizedDescription and assigns it to the Description field.
func (o *ImportEnumField) SetDescription(v LocalizedDescription) {
	o.Description = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *ImportEnumField) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEnumField) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *ImportEnumField) HasPlaceholder() bool {
	return o != nil && o.Placeholder != nil
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *ImportEnumField) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ImportEnumField) GetType() ImportFieldType {
	if o == nil || o.Type == nil {
		var ret ImportFieldType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEnumField) GetTypeOk() (*ImportFieldType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ImportEnumField) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ImportFieldType and assigns it to the Type field.
func (o *ImportEnumField) SetType(v ImportFieldType) {
	o.Type = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ImportEnumField) GetOptions() []string {
	if o == nil || o.Options == nil {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEnumField) GetOptionsOk() (*[]string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ImportEnumField) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *ImportEnumField) SetOptions(v []string) {
	o.Options = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ImportEnumField) GetDefault() string {
	if o == nil || o.Default.Get() == nil {
		var ret string
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportEnumField) GetDefaultOk() (*string, bool) {
	if o == nil || o.Default.Get() == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *ImportEnumField) HasDefault() bool {
	return o != nil && o.Default.IsSet()
}

// SetDefault gets a reference to the given common.NullableString and assigns it to the Default field.
func (o *ImportEnumField) SetDefault(v string) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil.
func (o *ImportEnumField) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil.
func (o *ImportEnumField) UnsetDefault() {
	o.Default.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportEnumField) MarshalJSON() ([]byte, error) {
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
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportEnumField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string               `json:"name,omitempty"`
		Label       *LocalizedDescription `json:"label,omitempty"`
		Required    *bool                 `json:"required,omitempty"`
		Sensitive   *bool                 `json:"sensitive,omitempty"`
		Description *LocalizedDescription `json:"description,omitempty"`
		Placeholder *string               `json:"placeholder,omitempty"`
		Type        *ImportFieldType      `json:"type,omitempty"`
		Options     []string              `json:"options,omitempty"`
		Default     common.NullableString `json:"default,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "label", "required", "sensitive", "description", "placeholder", "type", "options", "default"})
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
	o.Options = all.Options
	o.Default = all.Default
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}
	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
