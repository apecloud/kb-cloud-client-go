// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ParameterOption struct {
	// component type
	Component string `json:"component"`
	// default parameter templates, Including default parameter templates for different major versions or default parameter templates for different purposes.
	DefaultTemplates []ParameterTemplate `json:"defaultTemplates"`
	// If set to true, the current parameters of the component can be exported and saved as a parameter template.
	ExportTpl bool `json:"exportTpl"`
	// If set to true, the parameter templates of the component can be used. Mainly used by frontend.
	EnableTemplate bool  `json:"enableTemplate"`
	DisableHa      *bool `json:"disableHA,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterOption instantiates a new ParameterOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterOption(component string, defaultTemplates []ParameterTemplate, exportTpl bool, enableTemplate bool) *ParameterOption {
	this := ParameterOption{}
	this.Component = component
	this.DefaultTemplates = defaultTemplates
	this.ExportTpl = exportTpl
	this.EnableTemplate = enableTemplate
	var disableHa bool = false
	this.DisableHa = &disableHa
	return &this
}

// NewParameterOptionWithDefaults instantiates a new ParameterOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterOptionWithDefaults() *ParameterOption {
	this := ParameterOption{}
	var disableHa bool = false
	this.DisableHa = &disableHa
	return &this
}

// GetComponent returns the Component field value.
func (o *ParameterOption) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ParameterOption) SetComponent(v string) {
	o.Component = v
}

// GetDefaultTemplates returns the DefaultTemplates field value.
func (o *ParameterOption) GetDefaultTemplates() []ParameterTemplate {
	if o == nil {
		var ret []ParameterTemplate
		return ret
	}
	return o.DefaultTemplates
}

// GetDefaultTemplatesOk returns a tuple with the DefaultTemplates field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetDefaultTemplatesOk() (*[]ParameterTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultTemplates, true
}

// SetDefaultTemplates sets field value.
func (o *ParameterOption) SetDefaultTemplates(v []ParameterTemplate) {
	o.DefaultTemplates = v
}

// GetExportTpl returns the ExportTpl field value.
func (o *ParameterOption) GetExportTpl() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ExportTpl
}

// GetExportTplOk returns a tuple with the ExportTpl field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetExportTplOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportTpl, true
}

// SetExportTpl sets field value.
func (o *ParameterOption) SetExportTpl(v bool) {
	o.ExportTpl = v
}

// GetEnableTemplate returns the EnableTemplate field value.
func (o *ParameterOption) GetEnableTemplate() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.EnableTemplate
}

// GetEnableTemplateOk returns a tuple with the EnableTemplate field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetEnableTemplateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableTemplate, true
}

// SetEnableTemplate sets field value.
func (o *ParameterOption) SetEnableTemplate(v bool) {
	o.EnableTemplate = v
}

// GetDisableHa returns the DisableHa field value if set, zero value otherwise.
func (o *ParameterOption) GetDisableHa() bool {
	if o == nil || o.DisableHa == nil {
		var ret bool
		return ret
	}
	return *o.DisableHa
}

// GetDisableHaOk returns a tuple with the DisableHa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetDisableHaOk() (*bool, bool) {
	if o == nil || o.DisableHa == nil {
		return nil, false
	}
	return o.DisableHa, true
}

// HasDisableHa returns a boolean if a field has been set.
func (o *ParameterOption) HasDisableHa() bool {
	return o != nil && o.DisableHa != nil
}

// SetDisableHa gets a reference to the given bool and assigns it to the DisableHa field.
func (o *ParameterOption) SetDisableHa(v bool) {
	o.DisableHa = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["defaultTemplates"] = o.DefaultTemplates
	toSerialize["exportTpl"] = o.ExportTpl
	toSerialize["enableTemplate"] = o.EnableTemplate
	if o.DisableHa != nil {
		toSerialize["disableHA"] = o.DisableHa
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component        *string              `json:"component"`
		DefaultTemplates *[]ParameterTemplate `json:"defaultTemplates"`
		ExportTpl        *bool                `json:"exportTpl"`
		EnableTemplate   *bool                `json:"enableTemplate"`
		DisableHa        *bool                `json:"disableHA,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.DefaultTemplates == nil {
		return fmt.Errorf("required field defaultTemplates missing")
	}
	if all.ExportTpl == nil {
		return fmt.Errorf("required field exportTpl missing")
	}
	if all.EnableTemplate == nil {
		return fmt.Errorf("required field enableTemplate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "defaultTemplates", "exportTpl", "enableTemplate", "disableHA"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.DefaultTemplates = *all.DefaultTemplates
	o.ExportTpl = *all.ExportTpl
	o.EnableTemplate = *all.EnableTemplate
	o.DisableHa = all.DisableHa

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
