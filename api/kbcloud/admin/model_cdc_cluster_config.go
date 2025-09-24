// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type CdcClusterConfig struct {
	Component  *string           `json:"component,omitempty"`
	ConfigName *string           `json:"configName,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCdcClusterConfig instantiates a new CdcClusterConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCdcClusterConfig() *CdcClusterConfig {
	this := CdcClusterConfig{}
	return &this
}

// NewCdcClusterConfigWithDefaults instantiates a new CdcClusterConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCdcClusterConfigWithDefaults() *CdcClusterConfig {
	this := CdcClusterConfig{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CdcClusterConfig) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcClusterConfig) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CdcClusterConfig) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *CdcClusterConfig) SetComponent(v string) {
	o.Component = &v
}

// GetConfigName returns the ConfigName field value if set, zero value otherwise.
func (o *CdcClusterConfig) GetConfigName() string {
	if o == nil || o.ConfigName == nil {
		var ret string
		return ret
	}
	return *o.ConfigName
}

// GetConfigNameOk returns a tuple with the ConfigName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcClusterConfig) GetConfigNameOk() (*string, bool) {
	if o == nil || o.ConfigName == nil {
		return nil, false
	}
	return o.ConfigName, true
}

// HasConfigName returns a boolean if a field has been set.
func (o *CdcClusterConfig) HasConfigName() bool {
	return o != nil && o.ConfigName != nil
}

// SetConfigName gets a reference to the given string and assigns it to the ConfigName field.
func (o *CdcClusterConfig) SetConfigName(v string) {
	o.ConfigName = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *CdcClusterConfig) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdcClusterConfig) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CdcClusterConfig) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *CdcClusterConfig) SetParameters(v map[string]string) {
	o.Parameters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CdcClusterConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.ConfigName != nil {
		toSerialize["configName"] = o.ConfigName
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CdcClusterConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component  *string           `json:"component,omitempty"`
		ConfigName *string           `json:"configName,omitempty"`
		Parameters map[string]string `json:"parameters,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "configName", "parameters"})
	} else {
		return err
	}
	o.Component = all.Component
	o.ConfigName = all.ConfigName
	o.Parameters = all.Parameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
