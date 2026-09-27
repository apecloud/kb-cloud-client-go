// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VipPoolUpdate VIP Pool update
type VipPoolUpdate struct {
	// IP Addresses
	Addresses string `json:"addresses"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVipPoolUpdate instantiates a new VipPoolUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVipPoolUpdate(addresses string) *VipPoolUpdate {
	this := VipPoolUpdate{}
	this.Addresses = addresses
	return &this
}

// NewVipPoolUpdateWithDefaults instantiates a new VipPoolUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVipPoolUpdateWithDefaults() *VipPoolUpdate {
	this := VipPoolUpdate{}
	return &this
}

// GetAddresses returns the Addresses field value.
func (o *VipPoolUpdate) GetAddresses() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *VipPoolUpdate) GetAddressesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value.
func (o *VipPoolUpdate) SetAddresses(v string) {
	o.Addresses = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VipPoolUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["addresses"] = o.Addresses

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VipPoolUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Addresses *string `json:"addresses"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Addresses == nil {
		return fmt.Errorf("required field addresses missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"addresses"})
	} else {
		return err
	}
	o.Addresses = *all.Addresses

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
