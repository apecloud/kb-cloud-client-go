// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ConfigSpecOption struct {
	// name of config spec
	ConfigSpecName string `json:"configSpecName"`
	// format of the configuration file, such as ini, yaml, conf
	Format string `json:"format"`
	// name of the configuration file, such as my.cnf, postgresql.conf
	ConfigFileName string `json:"configFileName"`
	// section of the configuration file, such as mysqld
	Section *string `json:"section,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConfigSpecOption instantiates a new ConfigSpecOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfigSpecOption(configSpecName string, format string, configFileName string) *ConfigSpecOption {
	this := ConfigSpecOption{}
	this.ConfigSpecName = configSpecName
	this.Format = format
	this.ConfigFileName = configFileName
	return &this
}

// NewConfigSpecOptionWithDefaults instantiates a new ConfigSpecOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfigSpecOptionWithDefaults() *ConfigSpecOption {
	this := ConfigSpecOption{}
	return &this
}

// GetConfigSpecName returns the ConfigSpecName field value.
func (o *ConfigSpecOption) GetConfigSpecName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigSpecName
}

// GetConfigSpecNameOk returns a tuple with the ConfigSpecName field value
// and a boolean to check if the value has been set.
func (o *ConfigSpecOption) GetConfigSpecNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpecName, true
}

// SetConfigSpecName sets field value.
func (o *ConfigSpecOption) SetConfigSpecName(v string) {
	o.ConfigSpecName = v
}

// GetFormat returns the Format field value.
func (o *ConfigSpecOption) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *ConfigSpecOption) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *ConfigSpecOption) SetFormat(v string) {
	o.Format = v
}

// GetConfigFileName returns the ConfigFileName field value.
func (o *ConfigSpecOption) GetConfigFileName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigFileName
}

// GetConfigFileNameOk returns a tuple with the ConfigFileName field value
// and a boolean to check if the value has been set.
func (o *ConfigSpecOption) GetConfigFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigFileName, true
}

// SetConfigFileName sets field value.
func (o *ConfigSpecOption) SetConfigFileName(v string) {
	o.ConfigFileName = v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *ConfigSpecOption) GetSection() string {
	if o == nil || o.Section == nil {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigSpecOption) GetSectionOk() (*string, bool) {
	if o == nil || o.Section == nil {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *ConfigSpecOption) HasSection() bool {
	return o != nil && o.Section != nil
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *ConfigSpecOption) SetSection(v string) {
	o.Section = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConfigSpecOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["configSpecName"] = o.ConfigSpecName
	toSerialize["format"] = o.Format
	toSerialize["configFileName"] = o.ConfigFileName
	if o.Section != nil {
		toSerialize["section"] = o.Section
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfigSpecOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigSpecName *string `json:"configSpecName"`
		Format         *string `json:"format"`
		ConfigFileName *string `json:"configFileName"`
		Section        *string `json:"section,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ConfigSpecName == nil {
		return fmt.Errorf("required field configSpecName missing")
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	if all.ConfigFileName == nil {
		return fmt.Errorf("required field configFileName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"configSpecName", "format", "configFileName", "section"})
	} else {
		return err
	}
	o.ConfigSpecName = *all.ConfigSpecName
	o.Format = *all.Format
	o.ConfigFileName = *all.ConfigFileName
	o.Section = all.Section

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
