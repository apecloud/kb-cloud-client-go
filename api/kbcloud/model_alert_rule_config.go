// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	_io "io"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AlertRuleConfig struct {
	// YAML file content containing the new alert rule configuration
	Content *_io.Reader `json:"content,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertRuleConfig instantiates a new AlertRuleConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertRuleConfig() *AlertRuleConfig {
	this := AlertRuleConfig{}
	return &this
}

// NewAlertRuleConfigWithDefaults instantiates a new AlertRuleConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertRuleConfigWithDefaults() *AlertRuleConfig {
	this := AlertRuleConfig{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AlertRuleConfig) GetContent() _io.Reader {
	if o == nil || o.Content == nil {
		var ret _io.Reader
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRuleConfig) GetContentOk() (*_io.Reader, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AlertRuleConfig) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given _io.Reader and assigns it to the Content field.
func (o *AlertRuleConfig) SetContent(v _io.Reader) {
	o.Content = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertRuleConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertRuleConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content *_io.Reader `json:"content,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"content"})
	} else {
		return err
	}
	o.Content = all.Content

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
