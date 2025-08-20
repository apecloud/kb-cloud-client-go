// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VipPool VIP Pool
type VipPool struct {
	// IP Addresses
	Addresses string `json:"addresses"`
	// ID of VIP Pool
	Id interface{} `json:"id,omitempty"`
	// Total number of IP addresses
	Total int64 `json:"total"`
	// Used number of IP addresses
	Used int64 `json:"used"`
	// Used IP addresses
	UsedIPs []string `json:"usedIPs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVipPool instantiates a new VipPool object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVipPool(addresses string, total int64, used int64) *VipPool {
	this := VipPool{}
	this.Addresses = addresses
	this.Total = total
	this.Used = used
	return &this
}

// NewVipPoolWithDefaults instantiates a new VipPool object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVipPoolWithDefaults() *VipPool {
	this := VipPool{}
	return &this
}

// GetAddresses returns the Addresses field value.
func (o *VipPool) GetAddresses() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *VipPool) GetAddressesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addresses, true
}

// SetAddresses sets field value.
func (o *VipPool) SetAddresses(v string) {
	o.Addresses = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VipPool) GetId() interface{} {
	if o == nil || o.Id == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipPool) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VipPool) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *VipPool) SetId(v interface{}) {
	o.Id = v
}

// GetTotal returns the Total field value.
func (o *VipPool) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *VipPool) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *VipPool) SetTotal(v int64) {
	o.Total = v
}

// GetUsed returns the Used field value.
func (o *VipPool) GetUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *VipPool) GetUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value.
func (o *VipPool) SetUsed(v int64) {
	o.Used = v
}

// GetUsedIPs returns the UsedIPs field value if set, zero value otherwise.
func (o *VipPool) GetUsedIPs() []string {
	if o == nil || o.UsedIPs == nil {
		var ret []string
		return ret
	}
	return o.UsedIPs
}

// GetUsedIPsOk returns a tuple with the UsedIPs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipPool) GetUsedIPsOk() (*[]string, bool) {
	if o == nil || o.UsedIPs == nil {
		return nil, false
	}
	return &o.UsedIPs, true
}

// HasUsedIPs returns a boolean if a field has been set.
func (o *VipPool) HasUsedIPs() bool {
	return o != nil && o.UsedIPs != nil
}

// SetUsedIPs gets a reference to the given []string and assigns it to the UsedIPs field.
func (o *VipPool) SetUsedIPs(v []string) {
	o.UsedIPs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VipPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["addresses"] = o.Addresses
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["total"] = o.Total
	toSerialize["used"] = o.Used
	if o.UsedIPs != nil {
		toSerialize["usedIPs"] = o.UsedIPs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VipPool) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Addresses *string     `json:"addresses"`
		Id        interface{} `json:"id,omitempty"`
		Total     *int64      `json:"total"`
		Used      *int64      `json:"used"`
		UsedIPs   []string    `json:"usedIPs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Addresses == nil {
		return fmt.Errorf("required field addresses missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	if all.Used == nil {
		return fmt.Errorf("required field used missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"addresses", "id", "total", "used", "usedIPs"})
	} else {
		return err
	}
	o.Addresses = *all.Addresses
	o.Id = all.Id
	o.Total = *all.Total
	o.Used = *all.Used
	o.UsedIPs = all.UsedIPs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
