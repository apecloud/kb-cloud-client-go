// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// IpPoolSelection Provider-qualified Pod IP pool selection for component pod IP allocation.
type IpPoolSelection struct {
	// KBE provider used for Pod IP pool discovery and explicit selection.
	Provider IpPoolProvider `json:"provider"`
	// IPv4 pool names resolved within the selected provider.
	Ipv4Pools []string `json:"ipv4Pools,omitempty"`
	// IPv6 pool names resolved within the selected provider.
	Ipv6Pools []string `json:"ipv6Pools,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIpPoolSelection instantiates a new IpPoolSelection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIpPoolSelection(provider IpPoolProvider) *IpPoolSelection {
	this := IpPoolSelection{}
	this.Provider = provider
	return &this
}

// NewIpPoolSelectionWithDefaults instantiates a new IpPoolSelection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIpPoolSelectionWithDefaults() *IpPoolSelection {
	this := IpPoolSelection{}
	return &this
}

// GetProvider returns the Provider field value.
func (o *IpPoolSelection) GetProvider() IpPoolProvider {
	if o == nil {
		var ret IpPoolProvider
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *IpPoolSelection) GetProviderOk() (*IpPoolProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *IpPoolSelection) SetProvider(v IpPoolProvider) {
	o.Provider = v
}

// GetIpv4Pools returns the Ipv4Pools field value if set, zero value otherwise.
func (o *IpPoolSelection) GetIpv4Pools() []string {
	if o == nil || o.Ipv4Pools == nil {
		var ret []string
		return ret
	}
	return o.Ipv4Pools
}

// GetIpv4PoolsOk returns a tuple with the Ipv4Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolSelection) GetIpv4PoolsOk() (*[]string, bool) {
	if o == nil || o.Ipv4Pools == nil {
		return nil, false
	}
	return &o.Ipv4Pools, true
}

// HasIpv4Pools returns a boolean if a field has been set.
func (o *IpPoolSelection) HasIpv4Pools() bool {
	return o != nil && o.Ipv4Pools != nil
}

// SetIpv4Pools gets a reference to the given []string and assigns it to the Ipv4Pools field.
func (o *IpPoolSelection) SetIpv4Pools(v []string) {
	o.Ipv4Pools = v
}

// GetIpv6Pools returns the Ipv6Pools field value if set, zero value otherwise.
func (o *IpPoolSelection) GetIpv6Pools() []string {
	if o == nil || o.Ipv6Pools == nil {
		var ret []string
		return ret
	}
	return o.Ipv6Pools
}

// GetIpv6PoolsOk returns a tuple with the Ipv6Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPoolSelection) GetIpv6PoolsOk() (*[]string, bool) {
	if o == nil || o.Ipv6Pools == nil {
		return nil, false
	}
	return &o.Ipv6Pools, true
}

// HasIpv6Pools returns a boolean if a field has been set.
func (o *IpPoolSelection) HasIpv6Pools() bool {
	return o != nil && o.Ipv6Pools != nil
}

// SetIpv6Pools gets a reference to the given []string and assigns it to the Ipv6Pools field.
func (o *IpPoolSelection) SetIpv6Pools(v []string) {
	o.Ipv6Pools = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IpPoolSelection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["provider"] = o.Provider
	if o.Ipv4Pools != nil {
		toSerialize["ipv4Pools"] = o.Ipv4Pools
	}
	if o.Ipv6Pools != nil {
		toSerialize["ipv6Pools"] = o.Ipv6Pools
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IpPoolSelection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Provider  *IpPoolProvider `json:"provider"`
		Ipv4Pools []string        `json:"ipv4Pools,omitempty"`
		Ipv6Pools []string        `json:"ipv6Pools,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"provider", "ipv4Pools", "ipv6Pools"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Provider.IsValid() {
		hasInvalidField = true
	} else {
		o.Provider = *all.Provider
	}
	o.Ipv4Pools = all.Ipv4Pools
	o.Ipv6Pools = all.Ipv6Pools

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
