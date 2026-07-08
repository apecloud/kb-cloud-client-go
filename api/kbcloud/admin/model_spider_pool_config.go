// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// SpiderPoolConfig Spiderpool IPv4 IPPool selection for component pod IP allocation.
type SpiderPoolConfig struct {
	// IPv4 SpiderIPPool names used to generate ipam.spidernet.io/ippool.
	Ipv4Pools []string `json:"ipv4Pools,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpiderPoolConfig instantiates a new SpiderPoolConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpiderPoolConfig() *SpiderPoolConfig {
	this := SpiderPoolConfig{}
	return &this
}

// NewSpiderPoolConfigWithDefaults instantiates a new SpiderPoolConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpiderPoolConfigWithDefaults() *SpiderPoolConfig {
	this := SpiderPoolConfig{}
	return &this
}

// GetIpv4Pools returns the Ipv4Pools field value if set, zero value otherwise.
func (o *SpiderPoolConfig) GetIpv4Pools() []string {
	if o == nil || o.Ipv4Pools == nil {
		var ret []string
		return ret
	}
	return o.Ipv4Pools
}

// GetIpv4PoolsOk returns a tuple with the Ipv4Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderPoolConfig) GetIpv4PoolsOk() (*[]string, bool) {
	if o == nil || o.Ipv4Pools == nil {
		return nil, false
	}
	return &o.Ipv4Pools, true
}

// HasIpv4Pools returns a boolean if a field has been set.
func (o *SpiderPoolConfig) HasIpv4Pools() bool {
	return o != nil && o.Ipv4Pools != nil
}

// SetIpv4Pools gets a reference to the given []string and assigns it to the Ipv4Pools field.
func (o *SpiderPoolConfig) SetIpv4Pools(v []string) {
	o.Ipv4Pools = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpiderPoolConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Ipv4Pools != nil {
		toSerialize["ipv4Pools"] = o.Ipv4Pools
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SpiderPoolConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ipv4Pools []string `json:"ipv4Pools,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ipv4Pools"})
	} else {
		return err
	}
	o.Ipv4Pools = all.Ipv4Pools

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
