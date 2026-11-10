// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// VipPortCheckResult VIP port check result
type VipPortCheckResult struct {
	// Requested ports that are already in use
	Conflicts []int32 `json:"conflicts,omitempty"`
	// Checked VIP address
	Ip string `json:"ip"`
	// Ports already used on the VIP
	OccupiedPorts []VipPortOccupancy `json:"occupiedPorts"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVipPortCheckResult instantiates a new VipPortCheckResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVipPortCheckResult(ip string, occupiedPorts []VipPortOccupancy) *VipPortCheckResult {
	this := VipPortCheckResult{}
	this.Ip = ip
	this.OccupiedPorts = occupiedPorts
	return &this
}

// NewVipPortCheckResultWithDefaults instantiates a new VipPortCheckResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVipPortCheckResultWithDefaults() *VipPortCheckResult {
	this := VipPortCheckResult{}
	return &this
}

// GetConflicts returns the Conflicts field value if set, zero value otherwise.
func (o *VipPortCheckResult) GetConflicts() []int32 {
	if o == nil || o.Conflicts == nil {
		var ret []int32
		return ret
	}
	return o.Conflicts
}

// GetConflictsOk returns a tuple with the Conflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VipPortCheckResult) GetConflictsOk() (*[]int32, bool) {
	if o == nil || o.Conflicts == nil {
		return nil, false
	}
	return &o.Conflicts, true
}

// HasConflicts returns a boolean if a field has been set.
func (o *VipPortCheckResult) HasConflicts() bool {
	return o != nil && o.Conflicts != nil
}

// SetConflicts gets a reference to the given []int32 and assigns it to the Conflicts field.
func (o *VipPortCheckResult) SetConflicts(v []int32) {
	o.Conflicts = v
}

// GetIp returns the Ip field value.
func (o *VipPortCheckResult) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *VipPortCheckResult) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value.
func (o *VipPortCheckResult) SetIp(v string) {
	o.Ip = v
}

// GetOccupiedPorts returns the OccupiedPorts field value.
func (o *VipPortCheckResult) GetOccupiedPorts() []VipPortOccupancy {
	if o == nil {
		var ret []VipPortOccupancy
		return ret
	}
	return o.OccupiedPorts
}

// GetOccupiedPortsOk returns a tuple with the OccupiedPorts field value
// and a boolean to check if the value has been set.
func (o *VipPortCheckResult) GetOccupiedPortsOk() (*[]VipPortOccupancy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OccupiedPorts, true
}

// SetOccupiedPorts sets field value.
func (o *VipPortCheckResult) SetOccupiedPorts(v []VipPortOccupancy) {
	o.OccupiedPorts = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VipPortCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Conflicts != nil {
		toSerialize["conflicts"] = o.Conflicts
	}
	toSerialize["ip"] = o.Ip
	toSerialize["occupiedPorts"] = o.OccupiedPorts

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VipPortCheckResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conflicts     []int32             `json:"conflicts,omitempty"`
		Ip            *string             `json:"ip"`
		OccupiedPorts *[]VipPortOccupancy `json:"occupiedPorts"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Ip == nil {
		return fmt.Errorf("required field ip missing")
	}
	if all.OccupiedPorts == nil {
		return fmt.Errorf("required field occupiedPorts missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"conflicts", "ip", "occupiedPorts"})
	} else {
		return err
	}
	o.Conflicts = all.Conflicts
	o.Ip = *all.Ip
	o.OccupiedPorts = *all.OccupiedPorts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
