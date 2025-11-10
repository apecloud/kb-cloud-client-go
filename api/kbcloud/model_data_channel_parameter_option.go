// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// DataChannelParameterOption module parameter template, the key is the module name, the value is the parameter template
type DataChannelParameterOption struct {
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{}            `json:"-"`
	AdditionalProperties map[string][]DataChannelParameter `json:"-"`
}

// NewDataChannelParameterOption instantiates a new DataChannelParameterOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataChannelParameterOption() *DataChannelParameterOption {
	this := DataChannelParameterOption{}
	return &this
}

// NewDataChannelParameterOptionWithDefaults instantiates a new DataChannelParameterOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataChannelParameterOptionWithDefaults() *DataChannelParameterOption {
	this := DataChannelParameterOption{}
	return &this
}

// MarshalJSON serializes the struct using spec logic.
func (o DataChannelParameterOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataChannelParameterOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string][]DataChannelParameter)
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{})
	} else {
		return err
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
