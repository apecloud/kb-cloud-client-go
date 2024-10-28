// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION LocalizedDescription
type LocalizedDescription struct {
	// NODESCRIPTION ZhCn
	ZhCn string `json:"zh-CN"`
	// NODESCRIPTION EnUs
	EnUs string `json:"en-US"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLocalizedDescription instantiates a new LocalizedDescription object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLocalizedDescription(zhCn string, enUs string) *LocalizedDescription {
	this := LocalizedDescription{}
	this.ZhCn = zhCn
	this.EnUs = enUs
	return &this
}

// NewLocalizedDescriptionWithDefaults instantiates a new LocalizedDescription object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLocalizedDescriptionWithDefaults() *LocalizedDescription {
	this := LocalizedDescription{}
	return &this
}

// GetZhCn returns the ZhCn field value.
func (o *LocalizedDescription) GetZhCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ZhCn
}

// GetZhCnOk returns a tuple with the ZhCn field value
// and a boolean to check if the value has been set.
func (o *LocalizedDescription) GetZhCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZhCn, true
}

// SetZhCn sets field value.
func (o *LocalizedDescription) SetZhCn(v string) {
	o.ZhCn = v
}

// GetEnUs returns the EnUs field value.
func (o *LocalizedDescription) GetEnUs() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnUs
}

// GetEnUsOk returns a tuple with the EnUs field value
// and a boolean to check if the value has been set.
func (o *LocalizedDescription) GetEnUsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnUs, true
}

// SetEnUs sets field value.
func (o *LocalizedDescription) SetEnUs(v string) {
	o.EnUs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LocalizedDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["zh-CN"] = o.ZhCn
	toSerialize["en-US"] = o.EnUs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LocalizedDescription) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ZhCn *string `json:"zh-CN"`
		EnUs *string `json:"en-US"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ZhCn == nil {
		return fmt.Errorf("required field zh-CN missing")
	}
	if all.EnUs == nil {
		return fmt.Errorf("required field en-US missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"zh-CN", "en-US"})
	} else {
		return err
	}
	o.ZhCn = *all.ZhCn
	o.EnUs = *all.EnUs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
