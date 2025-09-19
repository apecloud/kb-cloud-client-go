// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ParameterConfig struct {
	// name of config spec, included in configSpecs
	ConfigSpecName string `json:"configSpecName"`
	// name of the referenced configuration template file name, all files in parameter template fold under stores
	TemplateRef string `json:"templateRef"`
	// name of the configuration file, such as my.cnf, postgresql.conf
	ConfigFileName string `json:"configFileName"`
	// format of the configuration file, such as ini, yaml, conf
	Format string `json:"format"`
	// section of the configuration file, such as mysqld
	Section *string `json:"section,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterConfig instantiates a new ParameterConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterConfig(configSpecName string, templateRef string, configFileName string, format string) *ParameterConfig {
	this := ParameterConfig{}
	this.ConfigSpecName = configSpecName
	this.TemplateRef = templateRef
	this.ConfigFileName = configFileName
	this.Format = format
	return &this
}

// NewParameterConfigWithDefaults instantiates a new ParameterConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterConfigWithDefaults() *ParameterConfig {
	this := ParameterConfig{}
	return &this
}

// GetConfigSpecName returns the ConfigSpecName field value.
func (o *ParameterConfig) GetConfigSpecName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigSpecName
}

// GetConfigSpecNameOk returns a tuple with the ConfigSpecName field value
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetConfigSpecNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpecName, true
}

// SetConfigSpecName sets field value.
func (o *ParameterConfig) SetConfigSpecName(v string) {
	o.ConfigSpecName = v
}

// GetTemplateRef returns the TemplateRef field value.
func (o *ParameterConfig) GetTemplateRef() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TemplateRef
}

// GetTemplateRefOk returns a tuple with the TemplateRef field value
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetTemplateRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateRef, true
}

// SetTemplateRef sets field value.
func (o *ParameterConfig) SetTemplateRef(v string) {
	o.TemplateRef = v
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

// GetFormat returns the Format field value.
func (o *ParameterConfig) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *ParameterConfig) SetFormat(v string) {
	o.Format = v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *ParameterConfig) GetSection() string {
	if o == nil || o.Section == nil {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterConfig) GetSectionOk() (*string, bool) {
	if o == nil || o.Section == nil {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *ParameterConfig) HasSection() bool {
	return o != nil && o.Section != nil
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *ParameterConfig) SetSection(v string) {
	o.Section = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["configSpecName"] = o.ConfigSpecName
	toSerialize["templateRef"] = o.TemplateRef
	toSerialize["configFileName"] = o.ConfigFileName
	toSerialize["format"] = o.Format
	if o.Section != nil {
		toSerialize["section"] = o.Section
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigSpecName *string `json:"configSpecName"`
		TemplateRef    *string `json:"templateRef"`
		ConfigFileName *string `json:"configFileName"`
		Format         *string `json:"format"`
		Section        *string `json:"section,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ConfigSpecName == nil {
		return fmt.Errorf("required field configSpecName missing")
	}
	if all.TemplateRef == nil {
		return fmt.Errorf("required field templateRef missing")
	}
	if all.ConfigFileName == nil {
		return fmt.Errorf("required field configFileName missing")
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"configSpecName", "templateRef", "configFileName", "format", "section"})
	} else {
		return err
	}
	o.ConfigSpecName = *all.ConfigSpecName
	o.TemplateRef = *all.TemplateRef
	o.ConfigFileName = *all.ConfigFileName
	o.Format = *all.Format
	o.Section = all.Section

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
