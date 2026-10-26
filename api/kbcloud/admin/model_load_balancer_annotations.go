// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// LoadBalancerAnnotations Provider-neutral annotations applied to load balancer Services by exposure type.
type LoadBalancerAnnotations struct {
	// Annotations used for load balancers exposed inside the VPC.
	Vpc map[string]string `json:"vpc"`
	// Annotations used for load balancers exposed to the Internet.
	Internet map[string]string `json:"internet"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLoadBalancerAnnotations instantiates a new LoadBalancerAnnotations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLoadBalancerAnnotations(vpc map[string]string, internet map[string]string) *LoadBalancerAnnotations {
	this := LoadBalancerAnnotations{}
	this.Vpc = vpc
	this.Internet = internet
	return &this
}

// NewLoadBalancerAnnotationsWithDefaults instantiates a new LoadBalancerAnnotations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLoadBalancerAnnotationsWithDefaults() *LoadBalancerAnnotations {
	this := LoadBalancerAnnotations{}
	return &this
}

// GetVpc returns the Vpc field value.
func (o *LoadBalancerAnnotations) GetVpc() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Vpc
}

// GetVpcOk returns a tuple with the Vpc field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerAnnotations) GetVpcOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vpc, true
}

// SetVpc sets field value.
func (o *LoadBalancerAnnotations) SetVpc(v map[string]string) {
	o.Vpc = v
}

// GetInternet returns the Internet field value.
func (o *LoadBalancerAnnotations) GetInternet() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Internet
}

// GetInternetOk returns a tuple with the Internet field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerAnnotations) GetInternetOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internet, true
}

// SetInternet sets field value.
func (o *LoadBalancerAnnotations) SetInternet(v map[string]string) {
	o.Internet = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LoadBalancerAnnotations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["vpc"] = o.Vpc
	toSerialize["internet"] = o.Internet

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LoadBalancerAnnotations) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Vpc      *map[string]string `json:"vpc"`
		Internet *map[string]string `json:"internet"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Vpc == nil {
		return fmt.Errorf("required field vpc missing")
	}
	if all.Internet == nil {
		return fmt.Errorf("required field internet missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"vpc", "internet"})
	} else {
		return err
	}
	o.Vpc = *all.Vpc
	o.Internet = *all.Internet

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
