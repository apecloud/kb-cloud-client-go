// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type UpdateTopicConfigRequest struct {
	Configs map[string]string `json:"configs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateTopicConfigRequest instantiates a new UpdateTopicConfigRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateTopicConfigRequest() *UpdateTopicConfigRequest {
	this := UpdateTopicConfigRequest{}
	return &this
}

// NewUpdateTopicConfigRequestWithDefaults instantiates a new UpdateTopicConfigRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateTopicConfigRequestWithDefaults() *UpdateTopicConfigRequest {
	this := UpdateTopicConfigRequest{}
	return &this
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *UpdateTopicConfigRequest) GetConfigs() map[string]string {
	if o == nil || o.Configs == nil {
		var ret map[string]string
		return ret
	}
	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTopicConfigRequest) GetConfigsOk() (*map[string]string, bool) {
	if o == nil || o.Configs == nil {
		return nil, false
	}
	return &o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *UpdateTopicConfigRequest) HasConfigs() bool {
	return o != nil && o.Configs != nil
}

// SetConfigs gets a reference to the given map[string]string and assigns it to the Configs field.
func (o *UpdateTopicConfigRequest) SetConfigs(v map[string]string) {
	o.Configs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateTopicConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Configs != nil {
		toSerialize["configs"] = o.Configs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateTopicConfigRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Configs map[string]string `json:"configs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"configs"})
	} else {
		return err
	}
	o.Configs = all.Configs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
