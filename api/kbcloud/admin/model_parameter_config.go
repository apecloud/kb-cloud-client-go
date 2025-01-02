// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParameterConfig parameterConfig contains specific configuration templates for each configuration file, primarily consisting of parameter templates and parameter constraints, mainly used by initializing the default template from addon.
type ParameterConfig struct {
	// name of configuration file
	ConfigFileName string `json:"configFileName"`
	// name of the configuration template, equal to componentDefinition.configs[x].name
	Name string `json:"name"`
	// constraint of the configConstraint, equal to componentDefinition.configs[x].constraintRef
	Constraint string `json:"constraint"`
	// regular expression of the parameters, mainly used by the frontend for parameter parsing.
	Regex string `json:"regex"`
	// name of the referenced configuration template ConfigMap object, equal to componentDefinition.configs[x].templateRef
	ConfigTplName string                 `json:"configTplName"`
	InitOptions   map[string]interface{} `json:"initOptions,omitempty"`
	// Parameters to be calculated based on the instance specifications
	CalculationParams []CalculationParam `json:"calculationParams,omitempty"`
	// sharedSpecParams means that the parameter value is different from the dedicated specification under the shared specification
	SharedSpecParams []SharedSpecParams `json:"sharedSpecParams,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterConfig instantiates a new ParameterConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterConfig(configFileName string, name string, constraint string, regex string, configTplName string) *ParameterConfig {
	this := ParameterConfig{}
	this.ConfigFileName = configFileName
	this.Name = name
	this.Constraint = constraint
	this.Regex = regex
	this.ConfigTplName = configTplName
	return &this
}

// NewParameterConfigWithDefaults instantiates a new ParameterConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterConfigWithDefaults() *ParameterConfig {
	this := ParameterConfig{}
	return &this
}

// GetConfigFileName returns the ConfigFileName field value.
func (o *ParameterConfig) GetConfigFileName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigFileName
}

// GetConfigFileNameOk returns a tuple with the ConfigFileName field value
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetConfigFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigFileName, true
}

// SetConfigFileName sets field value.
func (o *ParameterConfig) SetConfigFileName(v string) {
	o.ConfigFileName = v
}

// GetName returns the Name field value.
func (o *ParameterConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ParameterConfig) SetName(v string) {
	o.Name = v
}

// GetConstraint returns the Constraint field value.
func (o *ParameterConfig) GetConstraint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Constraint
}

// GetConstraintOk returns a tuple with the Constraint field value
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetConstraintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Constraint, true
}

// SetConstraint sets field value.
func (o *ParameterConfig) SetConstraint(v string) {
	o.Constraint = v
}

// GetRegex returns the Regex field value.
func (o *ParameterConfig) GetRegex() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Regex
}

// GetRegexOk returns a tuple with the Regex field value
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Regex, true
}

// SetRegex sets field value.
func (o *ParameterConfig) SetRegex(v string) {
	o.Regex = v
}

// GetConfigTplName returns the ConfigTplName field value.
func (o *ParameterConfig) GetConfigTplName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigTplName
}

// GetConfigTplNameOk returns a tuple with the ConfigTplName field value
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetConfigTplNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigTplName, true
}

// SetConfigTplName sets field value.
func (o *ParameterConfig) SetConfigTplName(v string) {
	o.ConfigTplName = v
}

// GetInitOptions returns the InitOptions field value if set, zero value otherwise.
func (o *ParameterConfig) GetInitOptions() map[string]interface{} {
	if o == nil || o.InitOptions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.InitOptions
}

// GetInitOptionsOk returns a tuple with the InitOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetInitOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.InitOptions == nil {
		return nil, false
	}
	return &o.InitOptions, true
}

// HasInitOptions returns a boolean if a field has been set.
func (o *ParameterConfig) HasInitOptions() bool {
	return o != nil && o.InitOptions != nil
}

// SetInitOptions gets a reference to the given map[string]interface{} and assigns it to the InitOptions field.
func (o *ParameterConfig) SetInitOptions(v map[string]interface{}) {
	o.InitOptions = v
}

// GetCalculationParams returns the CalculationParams field value if set, zero value otherwise.
func (o *ParameterConfig) GetCalculationParams() []CalculationParam {
	if o == nil || o.CalculationParams == nil {
		var ret []CalculationParam
		return ret
	}
	return o.CalculationParams
}

// GetCalculationParamsOk returns a tuple with the CalculationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetCalculationParamsOk() (*[]CalculationParam, bool) {
	if o == nil || o.CalculationParams == nil {
		return nil, false
	}
	return &o.CalculationParams, true
}

// HasCalculationParams returns a boolean if a field has been set.
func (o *ParameterConfig) HasCalculationParams() bool {
	return o != nil && o.CalculationParams != nil
}

// SetCalculationParams gets a reference to the given []CalculationParam and assigns it to the CalculationParams field.
func (o *ParameterConfig) SetCalculationParams(v []CalculationParam) {
	o.CalculationParams = v
}

// GetSharedSpecParams returns the SharedSpecParams field value if set, zero value otherwise.
func (o *ParameterConfig) GetSharedSpecParams() []SharedSpecParams {
	if o == nil || o.SharedSpecParams == nil {
		var ret []SharedSpecParams
		return ret
	}
	return o.SharedSpecParams
}

// GetSharedSpecParamsOk returns a tuple with the SharedSpecParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetSharedSpecParamsOk() (*[]SharedSpecParams, bool) {
	if o == nil || o.SharedSpecParams == nil {
		return nil, false
	}
	return &o.SharedSpecParams, true
}

// HasSharedSpecParams returns a boolean if a field has been set.
func (o *ParameterConfig) HasSharedSpecParams() bool {
	return o != nil && o.SharedSpecParams != nil
}

// SetSharedSpecParams gets a reference to the given []SharedSpecParams and assigns it to the SharedSpecParams field.
func (o *ParameterConfig) SetSharedSpecParams(v []SharedSpecParams) {
	o.SharedSpecParams = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["configFileName"] = o.ConfigFileName
	toSerialize["name"] = o.Name
	toSerialize["constraint"] = o.Constraint
	toSerialize["regex"] = o.Regex
	toSerialize["configTplName"] = o.ConfigTplName
	if o.InitOptions != nil {
		toSerialize["initOptions"] = o.InitOptions
	}
	if o.CalculationParams != nil {
		toSerialize["calculationParams"] = o.CalculationParams
	}
	if o.SharedSpecParams != nil {
		toSerialize["sharedSpecParams"] = o.SharedSpecParams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigFileName    *string                `json:"configFileName"`
		Name              *string                `json:"name"`
		Constraint        *string                `json:"constraint"`
		Regex             *string                `json:"regex"`
		ConfigTplName     *string                `json:"configTplName"`
		InitOptions       map[string]interface{} `json:"initOptions,omitempty"`
		CalculationParams []CalculationParam     `json:"calculationParams,omitempty"`
		SharedSpecParams  []SharedSpecParams     `json:"sharedSpecParams,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigFileName == nil {
		return fmt.Errorf("required field configFileName missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Constraint == nil {
		return fmt.Errorf("required field constraint missing")
	}
	if all.Regex == nil {
		return fmt.Errorf("required field regex missing")
	}
	if all.ConfigTplName == nil {
		return fmt.Errorf("required field configTplName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"configFileName", "name", "constraint", "regex", "configTplName", "initOptions", "calculationParams", "sharedSpecParams"})
	} else {
		return err
	}
	o.ConfigFileName = *all.ConfigFileName
	o.Name = *all.Name
	o.Constraint = *all.Constraint
	o.Regex = *all.Regex
	o.ConfigTplName = *all.ConfigTplName
	o.InitOptions = all.InitOptions
	o.CalculationParams = all.CalculationParams
	o.SharedSpecParams = all.SharedSpecParams

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
