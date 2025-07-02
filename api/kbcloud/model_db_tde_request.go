// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DbTDERequest struct {
	Enable  []string `json:"enable"`
	Disable []string `json:"disable"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDbTDERequest instantiates a new DbTDERequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDbTDERequest(enable []string, disable []string) *DbTDERequest {
	this := DbTDERequest{}
	this.Enable = enable
	this.Disable = disable
	return &this
}

// NewDbTDERequestWithDefaults instantiates a new DbTDERequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDbTDERequestWithDefaults() *DbTDERequest {
	this := DbTDERequest{}
	return &this
}

// GetEnable returns the Enable field value.
func (o *DbTDERequest) GetEnable() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *DbTDERequest) GetEnableOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value.
func (o *DbTDERequest) SetEnable(v []string) {
	o.Enable = v
}

// GetDisable returns the Disable field value.
func (o *DbTDERequest) GetDisable() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Disable
}

// GetDisableOk returns a tuple with the Disable field value
// and a boolean to check if the value has been set.
func (o *DbTDERequest) GetDisableOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disable, true
}

// SetDisable sets field value.
func (o *DbTDERequest) SetDisable(v []string) {
	o.Disable = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DbTDERequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enable"] = o.Enable
	toSerialize["disable"] = o.Disable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DbTDERequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enable  *[]string `json:"enable"`
		Disable *[]string `json:"disable"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enable == nil {
		return fmt.Errorf("required field enable missing")
	}
	if all.Disable == nil {
		return fmt.Errorf("required field disable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enable", "disable"})
	} else {
		return err
	}
	o.Enable = *all.Enable
	o.Disable = *all.Disable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
