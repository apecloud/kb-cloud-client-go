// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// ConfigurationWithRegex Cluster parameters configuration, include the file name and content of the parameters 
type ConfigurationWithRegex struct {
	// The name of the configuration file
	FileName string `json:"fileName"`
	// The content of the configuration file
	Content string `json:"content"`
	// the template regex
	Regex string `json:"regex"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewConfigurationWithRegex instantiates a new ConfigurationWithRegex object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfigurationWithRegex(fileName string, content string, regex string) *ConfigurationWithRegex {
	this := ConfigurationWithRegex{}
	this.FileName = fileName
	this.Content = content
	this.Regex = regex
	return &this
}

// NewConfigurationWithRegexWithDefaults instantiates a new ConfigurationWithRegex object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfigurationWithRegexWithDefaults() *ConfigurationWithRegex {
	this := ConfigurationWithRegex{}
	return &this
}
// GetFileName returns the FileName field value.
func (o *ConfigurationWithRegex) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *ConfigurationWithRegex) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value.
func (o *ConfigurationWithRegex) SetFileName(v string) {
	o.FileName = v
}


// GetContent returns the Content field value.
func (o *ConfigurationWithRegex) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ConfigurationWithRegex) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *ConfigurationWithRegex) SetContent(v string) {
	o.Content = v
}


// GetRegex returns the Regex field value.
func (o *ConfigurationWithRegex) GetRegex() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Regex
}

// GetRegexOk returns a tuple with the Regex field value
// and a boolean to check if the value has been set.
func (o *ConfigurationWithRegex) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Regex, true
}

// SetRegex sets field value.
func (o *ConfigurationWithRegex) SetRegex(v string) {
	o.Regex = v
}



// MarshalJSON serializes the struct using spec logic.
func (o ConfigurationWithRegex) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["fileName"] = o.FileName
	toSerialize["content"] = o.Content
	toSerialize["regex"] = o.Regex

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConfigurationWithRegex) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FileName *string `json:"fileName"`
		Content *string `json:"content"`
		Regex *string `json:"regex"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FileName == nil {
		return fmt.Errorf("required field fileName missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.Regex == nil {
		return fmt.Errorf("required field regex missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "fileName", "content", "regex",  })
	} else {
		return err
	}
	o.FileName = *all.FileName
	o.Content = *all.Content
	o.Regex = *all.Regex

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
