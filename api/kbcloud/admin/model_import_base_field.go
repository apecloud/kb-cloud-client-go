// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ImportBaseField struct {
	// Field's programmatic name (e.g., 'db_host')
	Name string `json:"name"`
	// User-facing label for UI (e.g., 'Database Host')
	Label string `json:"label"`
	// Whether the field is required
	Required *bool `json:"required,omitempty"`
	// Whether it contains sensitive information (e.g., password)
	Sensitive *bool `json:"sensitive,omitempty"`
	// Detailed usage instructions for the field
	Description *string `json:"description,omitempty"`
	// Placeholder text for the input
	Placeholder *string `json:"placeholder,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportBaseField instantiates a new ImportBaseField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportBaseField(name string, label string) *ImportBaseField {
	this := ImportBaseField{}
	this.Name = name
	this.Label = label
	var required bool = false
	this.Required = &required
	var sensitive bool = false
	this.Sensitive = &sensitive
	return &this
}

// NewImportBaseFieldWithDefaults instantiates a new ImportBaseField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportBaseFieldWithDefaults() *ImportBaseField {
	this := ImportBaseField{}
	var required bool = false
	this.Required = &required
	var sensitive bool = false
	this.Sensitive = &sensitive
	return &this
}

// GetName returns the Name field value.
func (o *ImportBaseField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImportBaseField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ImportBaseField) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value.
func (o *ImportBaseField) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ImportBaseField) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *ImportBaseField) SetLabel(v string) {
	o.Label = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ImportBaseField) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBaseField) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ImportBaseField) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ImportBaseField) SetRequired(v bool) {
	o.Required = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *ImportBaseField) GetSensitive() bool {
	if o == nil || o.Sensitive == nil {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBaseField) GetSensitiveOk() (*bool, bool) {
	if o == nil || o.Sensitive == nil {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *ImportBaseField) HasSensitive() bool {
	return o != nil && o.Sensitive != nil
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *ImportBaseField) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImportBaseField) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBaseField) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImportBaseField) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ImportBaseField) SetDescription(v string) {
	o.Description = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *ImportBaseField) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBaseField) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *ImportBaseField) HasPlaceholder() bool {
	return o != nil && o.Placeholder != nil
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *ImportBaseField) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportBaseField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["label"] = o.Label
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportBaseField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string `json:"name"`
		Label       *string `json:"label"`
		Required    *bool   `json:"required,omitempty"`
		Sensitive   *bool   `json:"sensitive,omitempty"`
		Description *string `json:"description,omitempty"`
		Placeholder *string `json:"placeholder,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "label", "required", "sensitive", "description", "placeholder"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Label = *all.Label
	o.Required = all.Required
	o.Sensitive = all.Sensitive
	o.Description = all.Description
	o.Placeholder = all.Placeholder

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
