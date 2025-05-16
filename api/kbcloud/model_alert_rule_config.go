// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	_io "io"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AlertRuleConfig struct {
	// YAML file content containing the new alert rule configuration
	Content _io.Reader `json:"content"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertRuleConfig instantiates a new AlertRuleConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertRuleConfig(content _io.Reader) *AlertRuleConfig {
	this := AlertRuleConfig{}
	this.Content = content
	return &this
}

// NewAlertRuleConfigWithDefaults instantiates a new AlertRuleConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertRuleConfigWithDefaults() *AlertRuleConfig {
	this := AlertRuleConfig{}
	return &this
}

// GetContent returns the Content field value.
func (o *AlertRuleConfig) GetContent() _io.Reader {
	if o == nil {
		var ret _io.Reader
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AlertRuleConfig) GetContentOk() (*_io.Reader, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *AlertRuleConfig) SetContent(v _io.Reader) {
	o.Content = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertRuleConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["content"] = o.Content

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertRuleConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content *_io.Reader `json:"content"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"content"})
	} else {
		return err
	}
	o.Content = *all.Content

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
