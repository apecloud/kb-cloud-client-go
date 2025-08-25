// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LoadBalancer The load balancer info
type LoadBalancer struct {
	// Whether the loadbalancer is available in the environment.
	Available LoadBalancerAvailableType `json:"available"`
	VipPools  []VipPool                 `json:"vipPools,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLoadBalancer instantiates a new LoadBalancer object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLoadBalancer(available LoadBalancerAvailableType) *LoadBalancer {
	this := LoadBalancer{}
	this.Available = available
	return &this
}

// NewLoadBalancerWithDefaults instantiates a new LoadBalancer object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLoadBalancerWithDefaults() *LoadBalancer {
	this := LoadBalancer{}
	return &this
}

// GetAvailable returns the Available field value.
func (o *LoadBalancer) GetAvailable() LoadBalancerAvailableType {
	if o == nil {
		var ret LoadBalancerAvailableType
		return ret
	}
	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetAvailableOk() (*LoadBalancerAvailableType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value.
func (o *LoadBalancer) SetAvailable(v LoadBalancerAvailableType) {
	o.Available = v
}

// GetVipPools returns the VipPools field value if set, zero value otherwise.
func (o *LoadBalancer) GetVipPools() []VipPool {
	if o == nil || o.VipPools == nil {
		var ret []VipPool
		return ret
	}
	return o.VipPools
}

// GetVipPoolsOk returns a tuple with the VipPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetVipPoolsOk() (*[]VipPool, bool) {
	if o == nil || o.VipPools == nil {
		return nil, false
	}
	return &o.VipPools, true
}

// HasVipPools returns a boolean if a field has been set.
func (o *LoadBalancer) HasVipPools() bool {
	return o != nil && o.VipPools != nil
}

// SetVipPools gets a reference to the given []VipPool and assigns it to the VipPools field.
func (o *LoadBalancer) SetVipPools(v []VipPool) {
	o.VipPools = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["available"] = o.Available
	if o.VipPools != nil {
		toSerialize["vipPools"] = o.VipPools
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LoadBalancer) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Available *LoadBalancerAvailableType `json:"available"`
		VipPools  []VipPool                  `json:"vipPools,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Available == nil {
		return fmt.Errorf("required field available missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"available", "vipPools"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Available.IsValid() {
		hasInvalidField = true
	} else {
		o.Available = *all.Available
	}
	o.VipPools = all.VipPools

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
