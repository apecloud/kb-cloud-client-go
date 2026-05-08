// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParameterTemplateDiff A list of parameter template differences against the matching default template
type ParameterTemplateDiff struct {
	// A parameter template reference used by the diff response
	SourceTemplate ParameterTemplateDiffTemplateRef `json:"sourceTemplate"`
	// A parameter template reference used by the diff response
	DefaultTemplate ParameterTemplateDiffTemplateRef `json:"defaultTemplate"`
	// Items is the list of parameter template differences
	Items []ParameterTemplateDiffItem `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterTemplateDiff instantiates a new ParameterTemplateDiff object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterTemplateDiff(sourceTemplate ParameterTemplateDiffTemplateRef, defaultTemplate ParameterTemplateDiffTemplateRef, items []ParameterTemplateDiffItem) *ParameterTemplateDiff {
	this := ParameterTemplateDiff{}
	this.SourceTemplate = sourceTemplate
	this.DefaultTemplate = defaultTemplate
	this.Items = items
	return &this
}

// NewParameterTemplateDiffWithDefaults instantiates a new ParameterTemplateDiff object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterTemplateDiffWithDefaults() *ParameterTemplateDiff {
	this := ParameterTemplateDiff{}
	return &this
}

// GetSourceTemplate returns the SourceTemplate field value.
func (o *ParameterTemplateDiff) GetSourceTemplate() ParameterTemplateDiffTemplateRef {
	if o == nil {
		var ret ParameterTemplateDiffTemplateRef
		return ret
	}
	return o.SourceTemplate
}

// GetSourceTemplateOk returns a tuple with the SourceTemplate field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiff) GetSourceTemplateOk() (*ParameterTemplateDiffTemplateRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceTemplate, true
}

// SetSourceTemplate sets field value.
func (o *ParameterTemplateDiff) SetSourceTemplate(v ParameterTemplateDiffTemplateRef) {
	o.SourceTemplate = v
}

// GetDefaultTemplate returns the DefaultTemplate field value.
func (o *ParameterTemplateDiff) GetDefaultTemplate() ParameterTemplateDiffTemplateRef {
	if o == nil {
		var ret ParameterTemplateDiffTemplateRef
		return ret
	}
	return o.DefaultTemplate
}

// GetDefaultTemplateOk returns a tuple with the DefaultTemplate field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiff) GetDefaultTemplateOk() (*ParameterTemplateDiffTemplateRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultTemplate, true
}

// SetDefaultTemplate sets field value.
func (o *ParameterTemplateDiff) SetDefaultTemplate(v ParameterTemplateDiffTemplateRef) {
	o.DefaultTemplate = v
}

// GetItems returns the Items field value.
func (o *ParameterTemplateDiff) GetItems() []ParameterTemplateDiffItem {
	if o == nil {
		var ret []ParameterTemplateDiffItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplateDiff) GetItemsOk() (*[]ParameterTemplateDiffItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *ParameterTemplateDiff) SetItems(v []ParameterTemplateDiffItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterTemplateDiff) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["sourceTemplate"] = o.SourceTemplate
	toSerialize["defaultTemplate"] = o.DefaultTemplate
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterTemplateDiff) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SourceTemplate  *ParameterTemplateDiffTemplateRef `json:"sourceTemplate"`
		DefaultTemplate *ParameterTemplateDiffTemplateRef `json:"defaultTemplate"`
		Items           *[]ParameterTemplateDiffItem      `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SourceTemplate == nil {
		return fmt.Errorf("required field sourceTemplate missing")
	}
	if all.DefaultTemplate == nil {
		return fmt.Errorf("required field defaultTemplate missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sourceTemplate", "defaultTemplate", "items"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.SourceTemplate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SourceTemplate = *all.SourceTemplate
	if all.DefaultTemplate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DefaultTemplate = *all.DefaultTemplate
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
