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
	// If set to true, the current parameters of the component can be exported and saved as a parameter template.
	ExportTpl bool `json:"exportTpl"`
	// If set to true, the parameter templates of the component can be used. Mainly used by frontend.
	EnableTemplate bool `json:"enableTemplate"`
	// all parameter configuration specs of this component
	ConfigSpecs []string `json:"configSpecs"`
	DisableHa   *bool    `json:"disableHA,omitempty"`
	// parameter templates, including default parameter templates for different major versions or default parameter templates for different purposes.
	Templates []ParameterTemplate `json:"templates,omitempty"`
	// parameter constraints, mainly used to verify the correctness of the parameter value.
	Constraints []ParameterConstraint `json:"constraints,omitempty"`
	// Parameters to be initialized when cluster is created
	InitOptions []InitParam `json:"initOptions,omitempty"`
	// Parameters to be calculated based on the instance specifications
	ExprParams []ExprParam `json:"exprParams,omitempty"`
	// occupationParams means that the parameter value is different from the dedicated specification under the shared specification
	OccupationParams []OccupationParam `json:"occupationParams,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterOption instantiates a new ParameterOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterOption(component string, exportTpl bool, enableTemplate bool, configSpecs []string) *ParameterOption {
	this := ParameterOption{}
	this.Component = component
	this.ExportTpl = exportTpl
	this.EnableTemplate = enableTemplate
	this.ConfigSpecs = configSpecs
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

// GetConfigSpecs returns the ConfigSpecs field value.
func (o *ParameterOption) GetConfigSpecs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ConfigSpecs
}

// GetConfigSpecsOk returns a tuple with the ConfigSpecs field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetConfigSpecsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpecs, true
}

// SetConfigSpecs sets field value.
func (o *ParameterOption) SetConfigSpecs(v []string) {
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

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *ParameterOption) GetConstraints() []ParameterConstraint {
	if o == nil || o.Constraints == nil {
		var ret []ParameterConstraint
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetConstraintsOk() (*[]ParameterConstraint, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return &o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *ParameterOption) HasConstraints() bool {
	return o != nil && o.Constraints != nil
}

// SetConstraints gets a reference to the given []ParameterConstraint and assigns it to the Constraints field.
func (o *ParameterOption) SetConstraints(v []ParameterConstraint) {
	o.Constraints = v
}

// GetInitOptions returns the InitOptions field value if set, zero value otherwise.
func (o *ParameterOption) GetInitOptions() []InitParam {
	if o == nil || o.InitOptions == nil {
		var ret []InitParam
		return ret
	}
	return o.InitOptions
}

// GetInitOptionsOk returns a tuple with the InitOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetInitOptionsOk() (*[]InitParam, bool) {
	if o == nil || o.InitOptions == nil {
		return nil, false
	}
	return &o.InitOptions, true
}

// HasInitOptions returns a boolean if a field has been set.
func (o *ParameterOption) HasInitOptions() bool {
	return o != nil && o.InitOptions != nil
}

// SetInitOptions gets a reference to the given []InitParam and assigns it to the InitOptions field.
func (o *ParameterOption) SetInitOptions(v []InitParam) {
	o.InitOptions = v
}

// GetExprParams returns the ExprParams field value if set, zero value otherwise.
func (o *ParameterOption) GetExprParams() []ExprParam {
	if o == nil || o.ExprParams == nil {
		var ret []ExprParam
		return ret
	}
	return o.ExprParams
}

// GetExprParamsOk returns a tuple with the ExprParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetExprParamsOk() (*[]ExprParam, bool) {
	if o == nil || o.ExprParams == nil {
		return nil, false
	}
	return &o.ExprParams, true
}

// HasExprParams returns a boolean if a field has been set.
func (o *ParameterOption) HasExprParams() bool {
	return o != nil && o.ExprParams != nil
}

// SetExprParams gets a reference to the given []ExprParam and assigns it to the ExprParams field.
func (o *ParameterOption) SetExprParams(v []ExprParam) {
	o.ExprParams = v
}

// GetOccupationParams returns the OccupationParams field value if set, zero value otherwise.
func (o *ParameterOption) GetOccupationParams() []OccupationParam {
	if o == nil || o.OccupationParams == nil {
		var ret []OccupationParam
		return ret
	}
	return o.OccupationParams
}

// GetOccupationParamsOk returns a tuple with the OccupationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetOccupationParamsOk() (*[]OccupationParam, bool) {
	if o == nil || o.OccupationParams == nil {
		return nil, false
	}
	return &o.OccupationParams, true
}

// HasOccupationParams returns a boolean if a field has been set.
func (o *ParameterOption) HasOccupationParams() bool {
	return o != nil && o.OccupationParams != nil
}

// SetOccupationParams gets a reference to the given []OccupationParam and assigns it to the OccupationParams field.
func (o *ParameterOption) SetOccupationParams(v []OccupationParam) {
	o.OccupationParams = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["exportTpl"] = o.ExportTpl
	toSerialize["enableTemplate"] = o.EnableTemplate
	toSerialize["configSpecs"] = o.ConfigSpecs
	if o.DisableHa != nil {
		toSerialize["disableHA"] = o.DisableHa
	}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if o.InitOptions != nil {
		toSerialize["initOptions"] = o.InitOptions
	}
	if o.ExprParams != nil {
		toSerialize["exprParams"] = o.ExprParams
	}
	if o.OccupationParams != nil {
		toSerialize["occupationParams"] = o.OccupationParams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component        *string               `json:"component"`
		ExportTpl        *bool                 `json:"exportTpl"`
		EnableTemplate   *bool                 `json:"enableTemplate"`
		ConfigSpecs      *[]string             `json:"configSpecs"`
		DisableHa        *bool                 `json:"disableHA,omitempty"`
		Templates        []ParameterTemplate   `json:"templates,omitempty"`
		Constraints      []ParameterConstraint `json:"constraints,omitempty"`
		InitOptions      []InitParam           `json:"initOptions,omitempty"`
		ExprParams       []ExprParam           `json:"exprParams,omitempty"`
		OccupationParams []OccupationParam     `json:"occupationParams,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.ExportTpl == nil {
		return fmt.Errorf("required field exportTpl missing")
	}
	if all.EnableTemplate == nil {
		return fmt.Errorf("required field enableTemplate missing")
	}
	if all.ConfigSpecs == nil {
		return fmt.Errorf("required field configSpecs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "exportTpl", "enableTemplate", "configSpecs", "disableHA", "templates", "constraints", "initOptions", "exprParams", "occupationParams"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.ExportTpl = *all.ExportTpl
	o.EnableTemplate = *all.EnableTemplate
	o.ConfigSpecs = *all.ConfigSpecs
	o.DisableHa = all.DisableHa
	o.Templates = all.Templates
	o.Constraints = all.Constraints
	o.InitOptions = all.InitOptions
	o.ExprParams = all.ExprParams
	o.OccupationParams = all.OccupationParams

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
