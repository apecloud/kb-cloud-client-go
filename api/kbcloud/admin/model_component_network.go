// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ComponentNetwork Per-component network configuration.
type ComponentNetwork struct {
	// Spiderpool IPv4 IPPool selection for component pod IP allocation.
	SpiderPool *SpiderPoolConfig `json:"spiderPool,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentNetwork instantiates a new ComponentNetwork object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentNetwork() *ComponentNetwork {
	this := ComponentNetwork{}
	return &this
}

// NewComponentNetworkWithDefaults instantiates a new ComponentNetwork object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentNetworkWithDefaults() *ComponentNetwork {
	this := ComponentNetwork{}
	return &this
}

// GetSpiderPool returns the SpiderPool field value if set, zero value otherwise.
func (o *ComponentNetwork) GetSpiderPool() SpiderPoolConfig {
	if o == nil || o.SpiderPool == nil {
		var ret SpiderPoolConfig
		return ret
	}
	return *o.SpiderPool
}

// GetSpiderPoolOk returns a tuple with the SpiderPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentNetwork) GetSpiderPoolOk() (*SpiderPoolConfig, bool) {
	if o == nil || o.SpiderPool == nil {
		return nil, false
	}
	return o.SpiderPool, true
}

// HasSpiderPool returns a boolean if a field has been set.
func (o *ComponentNetwork) HasSpiderPool() bool {
	return o != nil && o.SpiderPool != nil
}

// SetSpiderPool gets a reference to the given SpiderPoolConfig and assigns it to the SpiderPool field.
func (o *ComponentNetwork) SetSpiderPool(v SpiderPoolConfig) {
	o.SpiderPool = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SpiderPool != nil {
		toSerialize["spiderPool"] = o.SpiderPool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentNetwork) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SpiderPool *SpiderPoolConfig `json:"spiderPool,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"spiderPool"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.SpiderPool != nil && all.SpiderPool.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SpiderPool = all.SpiderPool

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
