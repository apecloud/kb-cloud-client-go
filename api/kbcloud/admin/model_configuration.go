// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Configuration Cluster parameters configuration, include the file name and content of the parameters
type Configuration struct {
	// The name of the configuration file
	FileName string `json:"fileName"`
	// The content of the configuration file
	Content string `json:"content"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConfiguration instantiates a new Configuration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConfiguration(fileName string, content string) *Configuration {
	this := Configuration{}
	this.FileName = fileName
	this.Content = content
	return &this
}

// NewConfigurationWithDefaults instantiates a new Configuration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConfigurationWithDefaults() *Configuration {
	this := Configuration{}
	return &this
}

// GetFileName returns the FileName field value.
func (o *Configuration) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *Configuration) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value.
func (o *Configuration) SetFileName(v string) {
	o.FileName = v
}

// GetContent returns the Content field value.
func (o *Configuration) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *Configuration) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *Configuration) SetContent(v string) {
	o.Content = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Configuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["fileName"] = o.FileName
	toSerialize["content"] = o.Content

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Configuration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FileName *string `json:"fileName"`
		Content  *string `json:"content"`
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
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"fileName", "content"})
	} else {
		return err
	}
	o.FileName = *all.FileName
	o.Content = *all.Content

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
