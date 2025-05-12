// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	_io "io"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AlertRuleYamlContent The content of the alert rule configuration file in YAML format, including organization names.
type AlertRuleYamlContent struct {
	// Organization name for the alert rule configuration
	OrgName *string `json:"orgName,omitempty"`
	// YAML file content containing the new alert rule configuration
	Content *_io.Reader `json:"content,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertRuleYamlContent instantiates a new AlertRuleYamlContent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertRuleYamlContent() *AlertRuleYamlContent {
	this := AlertRuleYamlContent{}
	return &this
}

// NewAlertRuleYamlContentWithDefaults instantiates a new AlertRuleYamlContent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertRuleYamlContentWithDefaults() *AlertRuleYamlContent {
	this := AlertRuleYamlContent{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AlertRuleYamlContent) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRuleYamlContent) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AlertRuleYamlContent) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AlertRuleYamlContent) SetOrgName(v string) {
	o.OrgName = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AlertRuleYamlContent) GetContent() _io.Reader {
	if o == nil || o.Content == nil {
		var ret _io.Reader
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertRuleYamlContent) GetContentOk() (*_io.Reader, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AlertRuleYamlContent) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given _io.Reader and assigns it to the Content field.
func (o *AlertRuleYamlContent) SetContent(v _io.Reader) {
	o.Content = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertRuleYamlContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
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
func (o *AlertRuleYamlContent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName *string     `json:"orgName,omitempty"`
		Content *_io.Reader `json:"content,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgName", "content"})
	} else {
		return err
	}
	o.OrgName = all.OrgName
	o.Content = all.Content

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
