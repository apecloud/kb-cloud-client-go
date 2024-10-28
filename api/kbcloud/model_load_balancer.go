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
	Available LoadBalancerStatus `json:"available"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLoadBalancer instantiates a new LoadBalancer object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLoadBalancer(available LoadBalancerStatus) *LoadBalancer {
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
func (o *LoadBalancer) GetAvailable() LoadBalancerStatus {
	if o == nil {
		var ret LoadBalancerStatus
		return ret
	}
	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *LoadBalancer) GetAvailableOk() (*LoadBalancerStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value.
func (o *LoadBalancer) SetAvailable(v LoadBalancerStatus) {
	o.Available = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LoadBalancer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["available"] = o.Available

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LoadBalancer) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Available *LoadBalancerStatus `json:"available"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Available == nil {
		return fmt.Errorf("required field available missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"available"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Available.IsValid() {
		hasInvalidField = true
	} else {
		o.Available = *all.Available
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
