// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type SecurityRoleTemplate struct {
	Template map[string]interface{}      `json:"template,omitempty"`
	Format   *SecurityRoleTemplateFormat `json:"format,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityRoleTemplate instantiates a new SecurityRoleTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityRoleTemplate() *SecurityRoleTemplate {
	this := SecurityRoleTemplate{}
	return &this
}

// NewSecurityRoleTemplateWithDefaults instantiates a new SecurityRoleTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityRoleTemplateWithDefaults() *SecurityRoleTemplate {
	this := SecurityRoleTemplate{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *SecurityRoleTemplate) GetTemplate() map[string]interface{} {
	if o == nil || o.Template == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleTemplate) GetTemplateOk() (*map[string]interface{}, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return &o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *SecurityRoleTemplate) HasTemplate() bool {
	return o != nil && o.Template != nil
}

// SetTemplate gets a reference to the given map[string]interface{} and assigns it to the Template field.
func (o *SecurityRoleTemplate) SetTemplate(v map[string]interface{}) {
	o.Template = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *SecurityRoleTemplate) GetFormat() SecurityRoleTemplateFormat {
	if o == nil || o.Format == nil {
		var ret SecurityRoleTemplateFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRoleTemplate) GetFormatOk() (*SecurityRoleTemplateFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *SecurityRoleTemplate) HasFormat() bool {
	return o != nil && o.Format != nil
}

// SetFormat gets a reference to the given SecurityRoleTemplateFormat and assigns it to the Format field.
func (o *SecurityRoleTemplate) SetFormat(v SecurityRoleTemplateFormat) {
	o.Format = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityRoleTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityRoleTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Template map[string]interface{}      `json:"template,omitempty"`
		Format   *SecurityRoleTemplateFormat `json:"format,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"template", "format"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Template = all.Template
	if all.Format != nil && !all.Format.IsValid() {
		hasInvalidField = true
	} else {
		o.Format = all.Format
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
