// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ParameterConfig struct {
	// NODESCRIPTION ConfigName
	ConfigName string `json:"configName"`
	// NODESCRIPTION SpecName
	SpecName string `json:"specName"`
	// NODESCRIPTION Constraint
	Constraint string `json:"constraint"`
	// NODESCRIPTION Regex
	Regex string `json:"regex"`
	// NODESCRIPTION ConfigTplName
	ConfigTplName string `json:"configTplName"`
	// NODESCRIPTION InitOptions
	InitOptions map[string]interface{} `json:"initOptions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterConfig instantiates a new ParameterConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterConfig(configName string, specName string, constraint string, regex string, configTplName string) *ParameterConfig {
	this := ParameterConfig{}
	this.ConfigName = configName
	this.SpecName = specName
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

// GetConfigName returns the ConfigName field value.
func (o *ParameterConfig) GetConfigName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigName
}

// GetConfigNameOk returns a tuple with the ConfigName field value
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetConfigNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigName, true
}

// SetConfigName sets field value.
func (o *ParameterConfig) SetConfigName(v string) {
	o.ConfigName = v
}

// GetSpecName returns the SpecName field value.
func (o *ParameterConfig) GetSpecName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpecName
}

// GetSpecNameOk returns a tuple with the SpecName field value
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetSpecNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecName, true
}

// SetSpecName sets field value.
func (o *ParameterConfig) SetSpecName(v string) {
	o.SpecName = v
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

// MarshalJSON serializes the struct using spec logic.
func (o ParameterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["configName"] = o.ConfigName
	toSerialize["specName"] = o.SpecName
	toSerialize["constraint"] = o.Constraint
	toSerialize["regex"] = o.Regex
	toSerialize["configTplName"] = o.ConfigTplName
	if o.InitOptions != nil {
		toSerialize["initOptions"] = o.InitOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigName    *string                `json:"configName"`
		SpecName      *string                `json:"specName"`
		Constraint    *string                `json:"constraint"`
		Regex         *string                `json:"regex"`
		ConfigTplName *string                `json:"configTplName"`
		InitOptions   map[string]interface{} `json:"initOptions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigName == nil {
		return fmt.Errorf("required field configName missing")
	}
	if all.SpecName == nil {
		return fmt.Errorf("required field specName missing")
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
		common.DeleteKeys(additionalProperties, &[]string{"configName", "specName", "constraint", "regex", "configTplName", "initOptions"})
	} else {
		return err
	}
	o.ConfigName = *all.ConfigName
	o.SpecName = *all.SpecName
	o.Constraint = *all.Constraint
	o.Regex = *all.Regex
	o.ConfigTplName = *all.ConfigTplName
	o.InitOptions = all.InitOptions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
