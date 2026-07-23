// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// IpPoolList Pod IP pool discovery result for an environment.
type IpPoolList struct {
	Providers []IpPoolProviderList `json:"providers"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIpPoolList instantiates a new IpPoolList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIpPoolList(providers []IpPoolProviderList) *IpPoolList {
	this := IpPoolList{}
	this.Providers = providers
	return &this
}

// NewIpPoolListWithDefaults instantiates a new IpPoolList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIpPoolListWithDefaults() *IpPoolList {
	this := IpPoolList{}
	return &this
}

// GetProviders returns the Providers field value.
func (o *IpPoolList) GetProviders() []IpPoolProviderList {
	if o == nil {
		var ret []IpPoolProviderList
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *IpPoolList) GetProvidersOk() (*[]IpPoolProviderList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Providers, true
}

// SetProviders sets field value.
func (o *IpPoolList) SetProviders(v []IpPoolProviderList) {
	o.Providers = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IpPoolList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["providers"] = o.Providers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IpPoolList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Providers *[]IpPoolProviderList `json:"providers"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Providers == nil {
		return fmt.Errorf("required field providers missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"providers"})
	} else {
		return err
	}
	o.Providers = *all.Providers

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
