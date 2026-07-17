// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ComponentNetwork Per-component network configuration.
type ComponentNetwork struct {
	// Provider-qualified Pod IP pool selection for component pod IP allocation.
	IpPool *IpPoolSelection `json:"ipPool,omitempty"`
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

// GetIpPool returns the IpPool field value if set, zero value otherwise.
func (o *ComponentNetwork) GetIpPool() IpPoolSelection {
	if o == nil || o.IpPool == nil {
		var ret IpPoolSelection
		return ret
	}
	return *o.IpPool
}

// GetIpPoolOk returns a tuple with the IpPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentNetwork) GetIpPoolOk() (*IpPoolSelection, bool) {
	if o == nil || o.IpPool == nil {
		return nil, false
	}
	return o.IpPool, true
}

// HasIpPool returns a boolean if a field has been set.
func (o *ComponentNetwork) HasIpPool() bool {
	return o != nil && o.IpPool != nil
}

// SetIpPool gets a reference to the given IpPoolSelection and assigns it to the IpPool field.
func (o *ComponentNetwork) SetIpPool(v IpPoolSelection) {
	o.IpPool = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.IpPool != nil {
		toSerialize["ipPool"] = o.IpPool
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentNetwork) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IpPool *IpPoolSelection `json:"ipPool,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ipPool"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.IpPool != nil && all.IpPool.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IpPool = all.IpPool

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
