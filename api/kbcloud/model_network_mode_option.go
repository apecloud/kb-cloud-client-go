// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type NetworkModeOption struct {
	Supported []NetworkMode `json:"supported"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNetworkModeOption instantiates a new NetworkModeOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNetworkModeOption(supported []NetworkMode) *NetworkModeOption {
	this := NetworkModeOption{}
	this.Supported = supported
	return &this
}

// NewNetworkModeOptionWithDefaults instantiates a new NetworkModeOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNetworkModeOptionWithDefaults() *NetworkModeOption {
	this := NetworkModeOption{}
	return &this
}

// GetSupported returns the Supported field value.
func (o *NetworkModeOption) GetSupported() []NetworkMode {
	if o == nil {
		var ret []NetworkMode
		return ret
	}
	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *NetworkModeOption) GetSupportedOk() (*[]NetworkMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supported, true
}

// SetSupported sets field value.
func (o *NetworkModeOption) SetSupported(v []NetworkMode) {
	o.Supported = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NetworkModeOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["supported"] = o.Supported

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NetworkModeOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Supported *[]NetworkMode `json:"supported"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Supported == nil {
		return fmt.Errorf("required field supported missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"supported"})
	} else {
		return err
	}
	o.Supported = *all.Supported

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
