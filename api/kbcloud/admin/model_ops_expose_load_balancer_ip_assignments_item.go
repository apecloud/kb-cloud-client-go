// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OpsExposeLoadBalancerIPAssignmentsItem struct {
	// The IP family of the LoadBalancer IP. Valid values are IPv4 and IPv6. If omitted, it will be inferred from loadBalancerIP.
	IpFamily             common.NullableString `json:"ipFamily,omitempty"`
	LoadBalancerIp       string                `json:"loadBalancerIP"`
	LoadBalancerIpPoolId string                `json:"loadBalancerIPPoolID"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsExposeLoadBalancerIPAssignmentsItem instantiates a new OpsExposeLoadBalancerIPAssignmentsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsExposeLoadBalancerIPAssignmentsItem(loadBalancerIp string, loadBalancerIpPoolId string) *OpsExposeLoadBalancerIPAssignmentsItem {
	this := OpsExposeLoadBalancerIPAssignmentsItem{}
	this.LoadBalancerIp = loadBalancerIp
	this.LoadBalancerIpPoolId = loadBalancerIpPoolId
	return &this
}

// NewOpsExposeLoadBalancerIPAssignmentsItemWithDefaults instantiates a new OpsExposeLoadBalancerIPAssignmentsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsExposeLoadBalancerIPAssignmentsItemWithDefaults() *OpsExposeLoadBalancerIPAssignmentsItem {
	this := OpsExposeLoadBalancerIPAssignmentsItem{}
	return &this
}

// GetIpFamily returns the IpFamily field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpsExposeLoadBalancerIPAssignmentsItem) GetIpFamily() string {
	if o == nil || o.IpFamily.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpFamily.Get()
}

// GetIpFamilyOk returns a tuple with the IpFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) GetIpFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpFamily.Get(), o.IpFamily.IsSet()
}

// HasIpFamily returns a boolean if a field has been set.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) HasIpFamily() bool {
	return o != nil && o.IpFamily.IsSet()
}

// SetIpFamily gets a reference to the given common.NullableString and assigns it to the IpFamily field.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) SetIpFamily(v string) {
	o.IpFamily.Set(&v)
}

// SetIpFamilyNil sets the value for IpFamily to be an explicit nil.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) SetIpFamilyNil() {
	o.IpFamily.Set(nil)
}

// UnsetIpFamily ensures that no value is present for IpFamily, not even an explicit nil.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) UnsetIpFamily() {
	o.IpFamily.Unset()
}

// GetLoadBalancerIp returns the LoadBalancerIp field value.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) GetLoadBalancerIp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LoadBalancerIp
}

// GetLoadBalancerIpOk returns a tuple with the LoadBalancerIp field value
// and a boolean to check if the value has been set.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) GetLoadBalancerIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerIp, true
}

// SetLoadBalancerIp sets field value.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) SetLoadBalancerIp(v string) {
	o.LoadBalancerIp = v
}

// GetLoadBalancerIpPoolId returns the LoadBalancerIpPoolId field value.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) GetLoadBalancerIpPoolId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LoadBalancerIpPoolId
}

// GetLoadBalancerIpPoolIdOk returns a tuple with the LoadBalancerIpPoolId field value
// and a boolean to check if the value has been set.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) GetLoadBalancerIpPoolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerIpPoolId, true
}

// SetLoadBalancerIpPoolId sets field value.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) SetLoadBalancerIpPoolId(v string) {
	o.LoadBalancerIpPoolId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsExposeLoadBalancerIPAssignmentsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.IpFamily.IsSet() {
		toSerialize["ipFamily"] = o.IpFamily.Get()
	}
	toSerialize["loadBalancerIP"] = o.LoadBalancerIp
	toSerialize["loadBalancerIPPoolID"] = o.LoadBalancerIpPoolId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsExposeLoadBalancerIPAssignmentsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IpFamily             common.NullableString `json:"ipFamily,omitempty"`
		LoadBalancerIp       *string               `json:"loadBalancerIP"`
		LoadBalancerIpPoolId *string               `json:"loadBalancerIPPoolID"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.LoadBalancerIp == nil {
		return fmt.Errorf("required field loadBalancerIP missing")
	}
	if all.LoadBalancerIpPoolId == nil {
		return fmt.Errorf("required field loadBalancerIPPoolID missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ipFamily", "loadBalancerIP", "loadBalancerIPPoolID"})
	} else {
		return err
	}
	o.IpFamily = all.IpFamily
	o.LoadBalancerIp = *all.LoadBalancerIp
	o.LoadBalancerIpPoolId = *all.LoadBalancerIpPoolId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
