// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LoadBalancerInstall Install the load balancer in the environment

type LoadBalancerInstall struct {
	// Type of the load balancer
	Type LoadBalancerType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLoadBalancerInstall instantiates a new LoadBalancerInstall object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLoadBalancerInstall(typeVar LoadBalancerType) *LoadBalancerInstall {
	this := LoadBalancerInstall{}
	this.Type = typeVar
	return &this
}

// NewLoadBalancerInstallWithDefaults instantiates a new LoadBalancerInstall object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLoadBalancerInstallWithDefaults() *LoadBalancerInstall {
	this := LoadBalancerInstall{}
	return &this
}

// GetType returns the Type field value.
func (o *LoadBalancerInstall) GetType() LoadBalancerType {
	if o == nil {
		var ret LoadBalancerType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerInstall) GetTypeOk() (*LoadBalancerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LoadBalancerInstall) SetType(v LoadBalancerType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LoadBalancerInstall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LoadBalancerInstall) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *LoadBalancerType `json:"type"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
