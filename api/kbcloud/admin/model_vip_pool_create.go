// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VipPoolCreate VIP Pool create
type VipPoolCreate struct {
	// IP Addresses
	Addresses string `json:"addresses"`
	// Name of the VIP Pool. It is what a cluster selects the pool by, so it is not restricted by Kubernetes resource naming rules. The MetalLB IPAddressPool resource name stays the generated kb-<id>.
	PoolName string `json:"poolName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVipPoolCreate instantiates a new VipPoolCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVipPoolCreate(addresses string, poolName string) *VipPoolCreate {
	this := VipPoolCreate{}
	this.Addresses = addresses
	this.PoolName = poolName
	return &this
}

// NewVipPoolCreateWithDefaults instantiates a new VipPoolCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVipPoolCreateWithDefaults() *VipPoolCreate {
	this := VipPoolCreate{}
	return &this
}

// GetAddresses returns the Addresses field value.
func (o *VipPoolCreate) GetAddresses() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *VipPoolCreate) GetAddressesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value.
func (o *VipPoolCreate) SetAddresses(v string) {
	o.Addresses = v
}

// GetPoolName returns the PoolName field value.
func (o *VipPoolCreate) GetPoolName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value
// and a boolean to check if the value has been set.
func (o *VipPoolCreate) GetPoolNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolName, true
}

// SetPoolName sets field value.
func (o *VipPoolCreate) SetPoolName(v string) {
	o.PoolName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VipPoolCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["addresses"] = o.Addresses
	toSerialize["poolName"] = o.PoolName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VipPoolCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Addresses *string `json:"addresses"`
		PoolName  *string `json:"poolName"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Addresses == nil {
		return fmt.Errorf("required field addresses missing")
	}
	if all.PoolName == nil {
		return fmt.Errorf("required field poolName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"addresses", "poolName"})
	} else {
		return err
	}
	o.Addresses = *all.Addresses
	o.PoolName = *all.PoolName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
