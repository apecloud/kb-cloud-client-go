// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ParameterTemplateImportResult struct {
	// parameter template information
	Template ParamTplListItem `json:"template"`
	Warnings []string         `json:"warnings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterTemplateImportResult instantiates a new ParameterTemplateImportResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterTemplateImportResult(template ParamTplListItem, warnings []string) *ParameterTemplateImportResult {
	this := ParameterTemplateImportResult{}
	this.Template = template
	this.Warnings = warnings
	return &this
}

// NewParameterTemplateImportResultWithDefaults instantiates a new ParameterTemplateImportResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterTemplateImportResultWithDefaults() *ParameterTemplateImportResult {
	this := ParameterTemplateImportResult{}
	return &this
}

// GetTemplate returns the Template field value.
func (o *ParameterTemplateImportResult) GetTemplate() ParamTplListItem {
	if o == nil {
		var ret ParamTplListItem
		return ret
	}
	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplateImportResult) GetTemplateOk() (*ParamTplListItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value.
func (o *ParameterTemplateImportResult) SetTemplate(v ParamTplListItem) {
	o.Template = v
}

// GetWarnings returns the Warnings field value.
func (o *ParameterTemplateImportResult) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *ParameterTemplateImportResult) GetWarningsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value.
func (o *ParameterTemplateImportResult) SetWarnings(v []string) {
	o.Warnings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterTemplateImportResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["template"] = o.Template
	toSerialize["warnings"] = o.Warnings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterTemplateImportResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Template *ParamTplListItem `json:"template"`
		Warnings *[]string         `json:"warnings"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Template == nil {
		return fmt.Errorf("required field template missing")
	}
	if all.Warnings == nil {
		return fmt.Errorf("required field warnings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"template", "warnings"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Template.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Template = *all.Template
	o.Warnings = *all.Warnings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
