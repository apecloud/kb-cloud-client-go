// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VipPortCheckResultOccupiedPortsItem A port already used on the VIP
type VipPortCheckResultOccupiedPortsItem struct {
	// Port number
	Port int32 `json:"port"`
	// Port protocol
	Protocol string `json:"protocol"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVipPortCheckResultOccupiedPortsItem instantiates a new VipPortCheckResultOccupiedPortsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVipPortCheckResultOccupiedPortsItem(port int32, protocol string) *VipPortCheckResultOccupiedPortsItem {
	this := VipPortCheckResultOccupiedPortsItem{}
	this.Port = port
	this.Protocol = protocol
	return &this
}

// NewVipPortCheckResultOccupiedPortsItemWithDefaults instantiates a new VipPortCheckResultOccupiedPortsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVipPortCheckResultOccupiedPortsItemWithDefaults() *VipPortCheckResultOccupiedPortsItem {
	this := VipPortCheckResultOccupiedPortsItem{}
	return &this
}

// GetPort returns the Port field value.
func (o *VipPortCheckResultOccupiedPortsItem) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *VipPortCheckResultOccupiedPortsItem) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value.
func (o *VipPortCheckResultOccupiedPortsItem) SetPort(v int32) {
	o.Port = v
}

// GetProtocol returns the Protocol field value.
func (o *VipPortCheckResultOccupiedPortsItem) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *VipPortCheckResultOccupiedPortsItem) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value.
func (o *VipPortCheckResultOccupiedPortsItem) SetProtocol(v string) {
	o.Protocol = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VipPortCheckResultOccupiedPortsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["port"] = o.Port
	toSerialize["protocol"] = o.Protocol

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VipPortCheckResultOccupiedPortsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Port     *int32  `json:"port"`
		Protocol *string `json:"protocol"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Port == nil {
		return fmt.Errorf("required field port missing")
	}
	if all.Protocol == nil {
		return fmt.Errorf("required field protocol missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"port", "protocol"})
	} else {
		return err
	}
	o.Port = *all.Port
	o.Protocol = *all.Protocol

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
