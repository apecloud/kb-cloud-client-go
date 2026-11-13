// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VipPortCheckRequest VIP port check request
type VipPortCheckRequest struct {
	// VIP address to check
	Ip string `json:"ip"`
	// TCP ports to check. When empty, only the occupied ports are returned.
	Ports []int32 `json:"ports,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVipPortCheckRequest instantiates a new VipPortCheckRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVipPortCheckRequest(ip string) *VipPortCheckRequest {
	this := VipPortCheckRequest{}
	this.Ip = ip
	return &this
}

// NewVipPortCheckRequestWithDefaults instantiates a new VipPortCheckRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVipPortCheckRequestWithDefaults() *VipPortCheckRequest {
	this := VipPortCheckRequest{}
	return &this
}

// GetIp returns the Ip field value.
func (o *VipPortCheckRequest) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *VipPortCheckRequest) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value.
func (o *VipPortCheckRequest) SetIp(v string) {
	o.Ip = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *VipPortCheckRequest) GetPorts() []int32 {
	if o == nil || o.Ports == nil {
		var ret []int32
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipPortCheckRequest) GetPortsOk() (*[]int32, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return &o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *VipPortCheckRequest) HasPorts() bool {
	return o != nil && o.Ports != nil
}

// SetPorts gets a reference to the given []int32 and assigns it to the Ports field.
func (o *VipPortCheckRequest) SetPorts(v []int32) {
	o.Ports = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VipPortCheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["ip"] = o.Ip
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VipPortCheckRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ip    *string `json:"ip"`
		Ports []int32 `json:"ports,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Ip == nil {
		return fmt.Errorf("required field ip missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ip", "ports"})
	} else {
		return err
	}
	o.Ip = *all.Ip
	o.Ports = all.Ports

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
