// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// IpPoolSelectionPolicy Per-address-family requirements for explicit IP pool selection.
type IpPoolSelectionPolicy struct {
	Ipv4 IpPoolSelectionRequirement `json:"ipv4"`
	Ipv6 IpPoolSelectionRequirement `json:"ipv6"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIpPoolSelectionPolicy instantiates a new IpPoolSelectionPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIpPoolSelectionPolicy(ipv4 IpPoolSelectionRequirement, ipv6 IpPoolSelectionRequirement) *IpPoolSelectionPolicy {
	this := IpPoolSelectionPolicy{}
	this.Ipv4 = ipv4
	this.Ipv6 = ipv6
	return &this
}

// NewIpPoolSelectionPolicyWithDefaults instantiates a new IpPoolSelectionPolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIpPoolSelectionPolicyWithDefaults() *IpPoolSelectionPolicy {
	this := IpPoolSelectionPolicy{}
	return &this
}

// GetIpv4 returns the Ipv4 field value.
func (o *IpPoolSelectionPolicy) GetIpv4() IpPoolSelectionRequirement {
	if o == nil {
		var ret IpPoolSelectionRequirement
		return ret
	}
	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value
// and a boolean to check if the value has been set.
func (o *IpPoolSelectionPolicy) GetIpv4Ok() (*IpPoolSelectionRequirement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ipv4, true
}

// SetIpv4 sets field value.
func (o *IpPoolSelectionPolicy) SetIpv4(v IpPoolSelectionRequirement) {
	o.Ipv4 = v
}

// GetIpv6 returns the Ipv6 field value.
func (o *IpPoolSelectionPolicy) GetIpv6() IpPoolSelectionRequirement {
	if o == nil {
		var ret IpPoolSelectionRequirement
		return ret
	}
	return o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value
// and a boolean to check if the value has been set.
func (o *IpPoolSelectionPolicy) GetIpv6Ok() (*IpPoolSelectionRequirement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ipv6, true
}

// SetIpv6 sets field value.
func (o *IpPoolSelectionPolicy) SetIpv6(v IpPoolSelectionRequirement) {
	o.Ipv6 = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IpPoolSelectionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["ipv4"] = o.Ipv4
	toSerialize["ipv6"] = o.Ipv6

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IpPoolSelectionPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ipv4 *IpPoolSelectionRequirement `json:"ipv4"`
		Ipv6 *IpPoolSelectionRequirement `json:"ipv6"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Ipv4 == nil {
		return fmt.Errorf("required field ipv4 missing")
	}
	if all.Ipv6 == nil {
		return fmt.Errorf("required field ipv6 missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ipv4", "ipv6"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Ipv4.IsValid() {
		hasInvalidField = true
	} else {
		o.Ipv4 = *all.Ipv4
	}
	if !all.Ipv6.IsValid() {
		hasInvalidField = true
	} else {
		o.Ipv6 = *all.Ipv6
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
