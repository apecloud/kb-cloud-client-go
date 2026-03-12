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
	// Reference another engine's parameter option to reuse its parameter configuration.
	// When set on a parameterOption, the referenced engine's parameter option (identified by
	// engineName + component) is used as the base configuration. Fields explicitly set in
	// the local parameterOption will override those from the referenced one.
	//
	ReferenceEngine *ParameterOptionRef `json:"referenceEngine,omitempty"`
	// If set to true, the current parameters of the component can be exported and saved as a parameter template.
	ExportTpl *bool `json:"exportTpl,omitempty"`
	// If set to true, the parameter templates of the component can be used. Mainly used by frontend.
	EnableTemplate *bool `json:"enableTemplate,omitempty"`
	// If set to true, the component can display the raw content of the parameter configuration file.
	EnableRawContent *bool `json:"enableRawContent,omitempty"`
	// all parameter configuration specs of this component
	ConfigSpecs []ConfigSpecOption `json:"configSpecs,omitempty"`
	DisableHa   *bool              `json:"disableHA,omitempty"`
	// parameter templates, including default parameter templates for different major versions or default parameter templates for different purposes.
	Templates []ParameterTemplate `json:"templates,omitempty"`
	// Parameters to be initialized when cluster is created
	InitOptions map[string]interface{} `json:"initOptions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterOption instantiates a new ParameterOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterOption(component string) *ParameterOption {
	this := ParameterOption{}
	this.Component = component
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

// GetReferenceEngine returns the ReferenceEngine field value if set, zero value otherwise.
func (o *ParameterOption) GetReferenceEngine() ParameterOptionRef {
	if o == nil || o.ReferenceEngine == nil {
		var ret ParameterOptionRef
		return ret
	}
	return *o.ReferenceEngine
}

// GetReferenceEngineOk returns a tuple with the ReferenceEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetReferenceEngineOk() (*ParameterOptionRef, bool) {
	if o == nil || o.ReferenceEngine == nil {
		return nil, false
	}
	return o.ReferenceEngine, true
}

// HasReferenceEngine returns a boolean if a field has been set.
func (o *ParameterOption) HasReferenceEngine() bool {
	return o != nil && o.ReferenceEngine != nil
}

// SetReferenceEngine gets a reference to the given ParameterOptionRef and assigns it to the ReferenceEngine field.
func (o *ParameterOption) SetReferenceEngine(v ParameterOptionRef) {
	o.ReferenceEngine = &v
}

// GetExportTpl returns the ExportTpl field value if set, zero value otherwise.
func (o *ParameterOption) GetExportTpl() bool {
	if o == nil || o.ExportTpl == nil {
		var ret bool
		return ret
	}
	return *o.ExportTpl
}

// GetExportTplOk returns a tuple with the ExportTpl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetExportTplOk() (*bool, bool) {
	if o == nil || o.ExportTpl == nil {
		return nil, false
	}
	return o.ExportTpl, true
}

// HasExportTpl returns a boolean if a field has been set.
func (o *ParameterOption) HasExportTpl() bool {
	return o != nil && o.ExportTpl != nil
}

// SetExportTpl gets a reference to the given bool and assigns it to the ExportTpl field.
func (o *ParameterOption) SetExportTpl(v bool) {
	o.ExportTpl = &v
}

// GetEnableTemplate returns the EnableTemplate field value if set, zero value otherwise.
func (o *ParameterOption) GetEnableTemplate() bool {
	if o == nil || o.EnableTemplate == nil {
		var ret bool
		return ret
	}
	return *o.EnableTemplate
}

// GetEnableTemplateOk returns a tuple with the EnableTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetEnableTemplateOk() (*bool, bool) {
	if o == nil || o.EnableTemplate == nil {
		return nil, false
	}
	return o.EnableTemplate, true
}

// HasEnableTemplate returns a boolean if a field has been set.
func (o *ParameterOption) HasEnableTemplate() bool {
	return o != nil && o.EnableTemplate != nil
}

// SetEnableTemplate gets a reference to the given bool and assigns it to the EnableTemplate field.
func (o *ParameterOption) SetEnableTemplate(v bool) {
	o.EnableTemplate = &v
}

// GetEnableRawContent returns the EnableRawContent field value if set, zero value otherwise.
func (o *ParameterOption) GetEnableRawContent() bool {
	if o == nil || o.EnableRawContent == nil {
		var ret bool
		return ret
	}
	return *o.EnableRawContent
}

// GetEnableRawContentOk returns a tuple with the EnableRawContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetEnableRawContentOk() (*bool, bool) {
	if o == nil || o.EnableRawContent == nil {
		return nil, false
	}
	return o.EnableRawContent, true
}

// HasEnableRawContent returns a boolean if a field has been set.
func (o *ParameterOption) HasEnableRawContent() bool {
	return o != nil && o.EnableRawContent != nil
}

// SetEnableRawContent gets a reference to the given bool and assigns it to the EnableRawContent field.
func (o *ParameterOption) SetEnableRawContent(v bool) {
	o.EnableRawContent = &v
}

// GetConfigSpecs returns the ConfigSpecs field value if set, zero value otherwise.
func (o *ParameterOption) GetConfigSpecs() []ConfigSpecOption {
	if o == nil || o.ConfigSpecs == nil {
		var ret []ConfigSpecOption
		return ret
	}
	return o.ConfigSpecs
}

// GetConfigSpecsOk returns a tuple with the ConfigSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetConfigSpecsOk() (*[]ConfigSpecOption, bool) {
	if o == nil || o.ConfigSpecs == nil {
		return nil, false
	}
	return &o.ConfigSpecs, true
}

// HasConfigSpecs returns a boolean if a field has been set.
func (o *ParameterOption) HasConfigSpecs() bool {
	return o != nil && o.ConfigSpecs != nil
}

// SetConfigSpecs gets a reference to the given []ConfigSpecOption and assigns it to the ConfigSpecs field.
func (o *ParameterOption) SetConfigSpecs(v []ConfigSpecOption) {
	o.ConfigSpecs = v
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

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *ParameterOption) GetTemplates() []ParameterTemplate {
	if o == nil || o.Templates == nil {
		var ret []ParameterTemplate
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetTemplatesOk() (*[]ParameterTemplate, bool) {
	if o == nil || o.Templates == nil {
		return nil, false
	}
	return &o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *ParameterOption) HasTemplates() bool {
	return o != nil && o.Templates != nil
}

// SetTemplates gets a reference to the given []ParameterTemplate and assigns it to the Templates field.
func (o *ParameterOption) SetTemplates(v []ParameterTemplate) {
	o.Templates = v
}

// GetInitOptions returns the InitOptions field value if set, zero value otherwise.
func (o *ParameterOption) GetInitOptions() map[string]interface{} {
	if o == nil || o.InitOptions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.InitOptions
}

// GetInitOptionsOk returns a tuple with the InitOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetInitOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.InitOptions == nil {
		return nil, false
	}
	return &o.InitOptions, true
}

// HasInitOptions returns a boolean if a field has been set.
func (o *ParameterOption) HasInitOptions() bool {
	return o != nil && o.InitOptions != nil
}

// SetInitOptions gets a reference to the given map[string]interface{} and assigns it to the InitOptions field.
func (o *ParameterOption) SetInitOptions(v map[string]interface{}) {
	o.InitOptions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	if o.ReferenceEngine != nil {
		toSerialize["referenceEngine"] = o.ReferenceEngine
	}
	if o.ExportTpl != nil {
		toSerialize["exportTpl"] = o.ExportTpl
	}
	if o.EnableTemplate != nil {
		toSerialize["enableTemplate"] = o.EnableTemplate
	}
	if o.EnableRawContent != nil {
		toSerialize["enableRawContent"] = o.EnableRawContent
	}
	if o.ConfigSpecs != nil {
		toSerialize["configSpecs"] = o.ConfigSpecs
	}
	if o.DisableHa != nil {
		toSerialize["disableHA"] = o.DisableHa
	}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	if o.InitOptions != nil {
		toSerialize["initOptions"] = o.InitOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component        *string                `json:"component"`
		ReferenceEngine  *ParameterOptionRef    `json:"referenceEngine,omitempty"`
		ExportTpl        *bool                  `json:"exportTpl,omitempty"`
		EnableTemplate   *bool                  `json:"enableTemplate,omitempty"`
		EnableRawContent *bool                  `json:"enableRawContent,omitempty"`
		ConfigSpecs      []ConfigSpecOption     `json:"configSpecs,omitempty"`
		DisableHa        *bool                  `json:"disableHA,omitempty"`
		Templates        []ParameterTemplate    `json:"templates,omitempty"`
		InitOptions      map[string]interface{} `json:"initOptions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "referenceEngine", "exportTpl", "enableTemplate", "enableRawContent", "configSpecs", "disableHA", "templates", "initOptions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Component = *all.Component
	if all.ReferenceEngine != nil && all.ReferenceEngine.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReferenceEngine = all.ReferenceEngine
	o.ExportTpl = all.ExportTpl
	o.EnableTemplate = all.EnableTemplate
	o.EnableRawContent = all.EnableRawContent
	o.ConfigSpecs = all.ConfigSpecs
	o.DisableHa = all.DisableHa
	o.Templates = all.Templates
	o.InitOptions = all.InitOptions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
