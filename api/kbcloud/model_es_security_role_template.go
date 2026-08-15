// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ESSecurityRoleTemplate struct {
	Template map[string]interface{}        `json:"template,omitempty"`
	Format   *ESSecurityRoleTemplateFormat `json:"format,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewESSecurityRoleTemplate instantiates a new ESSecurityRoleTemplate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewESSecurityRoleTemplate() *ESSecurityRoleTemplate {
	this := ESSecurityRoleTemplate{}
	return &this
}

// NewESSecurityRoleTemplateWithDefaults instantiates a new ESSecurityRoleTemplate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewESSecurityRoleTemplateWithDefaults() *ESSecurityRoleTemplate {
	this := ESSecurityRoleTemplate{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ESSecurityRoleTemplate) GetTemplate() map[string]interface{} {
	if o == nil || o.Template == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleTemplate) GetTemplateOk() (*map[string]interface{}, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return &o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ESSecurityRoleTemplate) HasTemplate() bool {
	return o != nil && o.Template != nil
}

// SetTemplate gets a reference to the given map[string]interface{} and assigns it to the Template field.
func (o *ESSecurityRoleTemplate) SetTemplate(v map[string]interface{}) {
	o.Template = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ESSecurityRoleTemplate) GetFormat() ESSecurityRoleTemplateFormat {
	if o == nil || o.Format == nil {
		var ret ESSecurityRoleTemplateFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ESSecurityRoleTemplate) GetFormatOk() (*ESSecurityRoleTemplateFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ESSecurityRoleTemplate) HasFormat() bool {
	return o != nil && o.Format != nil
}

// SetFormat gets a reference to the given ESSecurityRoleTemplateFormat and assigns it to the Format field.
func (o *ESSecurityRoleTemplate) SetFormat(v ESSecurityRoleTemplateFormat) {
	o.Format = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ESSecurityRoleTemplate) MarshalJSON() ([]byte, error) {
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
func (o *ESSecurityRoleTemplate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Template map[string]interface{}        `json:"template,omitempty"`
		Format   *ESSecurityRoleTemplateFormat `json:"format,omitempty"`
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
