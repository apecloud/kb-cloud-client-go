// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// IpPool Provider-neutral Pod IP pool summary.
type IpPool struct {
	Name             *string         `json:"name,omitempty"`
	IpFamily         *IpPoolIPFamily `json:"ipFamily,omitempty"`
	Subnet           *string         `json:"subnet,omitempty"`
	Disabled         *bool           `json:"disabled,omitempty"`
	Default          *bool           `json:"default,omitempty"`
	AllocatedIpCount *int64          `json:"allocatedIPCount,omitempty"`
	TotalIpCount     *int64          `json:"totalIPCount,omitempty"`
	AvailableIpCount *int64          `json:"availableIPCount,omitempty"`
	CapacityReliable *bool           `json:"capacityReliable,omitempty"`
	AffinitySummary  []string        `json:"affinitySummary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIpPool instantiates a new IpPool object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIpPool() *IpPool {
	this := IpPool{}
	return &this
}

// NewIpPoolWithDefaults instantiates a new IpPool object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIpPoolWithDefaults() *IpPool {
	this := IpPool{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IpPool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IpPool) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IpPool) SetName(v string) {
	o.Name = &v
}

// GetIpFamily returns the IpFamily field value if set, zero value otherwise.
func (o *IpPool) GetIpFamily() IpPoolIPFamily {
	if o == nil || o.IpFamily == nil {
		var ret IpPoolIPFamily
		return ret
	}
	return *o.IpFamily
}

// GetIpFamilyOk returns a tuple with the IpFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetIpFamilyOk() (*IpPoolIPFamily, bool) {
	if o == nil || o.IpFamily == nil {
		return nil, false
	}
	return o.IpFamily, true
}

// HasIpFamily returns a boolean if a field has been set.
func (o *IpPool) HasIpFamily() bool {
	return o != nil && o.IpFamily != nil
}

// SetIpFamily gets a reference to the given IpPoolIPFamily and assigns it to the IpFamily field.
func (o *IpPool) SetIpFamily(v IpPoolIPFamily) {
	o.IpFamily = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *IpPool) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *IpPool) HasSubnet() bool {
	return o != nil && o.Subnet != nil
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *IpPool) SetSubnet(v string) {
	o.Subnet = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *IpPool) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *IpPool) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *IpPool) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *IpPool) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *IpPool) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *IpPool) SetDefault(v bool) {
	o.Default = &v
}

// GetAllocatedIpCount returns the AllocatedIpCount field value if set, zero value otherwise.
func (o *IpPool) GetAllocatedIpCount() int64 {
	if o == nil || o.AllocatedIpCount == nil {
		var ret int64
		return ret
	}
	return *o.AllocatedIpCount
}

// GetAllocatedIpCountOk returns a tuple with the AllocatedIpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetAllocatedIpCountOk() (*int64, bool) {
	if o == nil || o.AllocatedIpCount == nil {
		return nil, false
	}
	return o.AllocatedIpCount, true
}

// HasAllocatedIpCount returns a boolean if a field has been set.
func (o *IpPool) HasAllocatedIpCount() bool {
	return o != nil && o.AllocatedIpCount != nil
}

// SetAllocatedIpCount gets a reference to the given int64 and assigns it to the AllocatedIpCount field.
func (o *IpPool) SetAllocatedIpCount(v int64) {
	o.AllocatedIpCount = &v
}

// GetTotalIpCount returns the TotalIpCount field value if set, zero value otherwise.
func (o *IpPool) GetTotalIpCount() int64 {
	if o == nil || o.TotalIpCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalIpCount
}

// GetTotalIpCountOk returns a tuple with the TotalIpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetTotalIpCountOk() (*int64, bool) {
	if o == nil || o.TotalIpCount == nil {
		return nil, false
	}
	return o.TotalIpCount, true
}

// HasTotalIpCount returns a boolean if a field has been set.
func (o *IpPool) HasTotalIpCount() bool {
	return o != nil && o.TotalIpCount != nil
}

// SetTotalIpCount gets a reference to the given int64 and assigns it to the TotalIpCount field.
func (o *IpPool) SetTotalIpCount(v int64) {
	o.TotalIpCount = &v
}

// GetAvailableIpCount returns the AvailableIpCount field value if set, zero value otherwise.
func (o *IpPool) GetAvailableIpCount() int64 {
	if o == nil || o.AvailableIpCount == nil {
		var ret int64
		return ret
	}
	return *o.AvailableIpCount
}

// GetAvailableIpCountOk returns a tuple with the AvailableIpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetAvailableIpCountOk() (*int64, bool) {
	if o == nil || o.AvailableIpCount == nil {
		return nil, false
	}
	return o.AvailableIpCount, true
}

// HasAvailableIpCount returns a boolean if a field has been set.
func (o *IpPool) HasAvailableIpCount() bool {
	return o != nil && o.AvailableIpCount != nil
}

// SetAvailableIpCount gets a reference to the given int64 and assigns it to the AvailableIpCount field.
func (o *IpPool) SetAvailableIpCount(v int64) {
	o.AvailableIpCount = &v
}

// GetCapacityReliable returns the CapacityReliable field value if set, zero value otherwise.
func (o *IpPool) GetCapacityReliable() bool {
	if o == nil || o.CapacityReliable == nil {
		var ret bool
		return ret
	}
	return *o.CapacityReliable
}

// GetCapacityReliableOk returns a tuple with the CapacityReliable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetCapacityReliableOk() (*bool, bool) {
	if o == nil || o.CapacityReliable == nil {
		return nil, false
	}
	return o.CapacityReliable, true
}

// HasCapacityReliable returns a boolean if a field has been set.
func (o *IpPool) HasCapacityReliable() bool {
	return o != nil && o.CapacityReliable != nil
}

// SetCapacityReliable gets a reference to the given bool and assigns it to the CapacityReliable field.
func (o *IpPool) SetCapacityReliable(v bool) {
	o.CapacityReliable = &v
}

// GetAffinitySummary returns the AffinitySummary field value if set, zero value otherwise.
func (o *IpPool) GetAffinitySummary() []string {
	if o == nil || o.AffinitySummary == nil {
		var ret []string
		return ret
	}
	return o.AffinitySummary
}

// GetAffinitySummaryOk returns a tuple with the AffinitySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetAffinitySummaryOk() (*[]string, bool) {
	if o == nil || o.AffinitySummary == nil {
		return nil, false
	}
	return &o.AffinitySummary, true
}

// HasAffinitySummary returns a boolean if a field has been set.
func (o *IpPool) HasAffinitySummary() bool {
	return o != nil && o.AffinitySummary != nil
}

// SetAffinitySummary gets a reference to the given []string and assigns it to the AffinitySummary field.
func (o *IpPool) SetAffinitySummary(v []string) {
	o.AffinitySummary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IpPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.IpFamily != nil {
		toSerialize["ipFamily"] = o.IpFamily
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.AllocatedIpCount != nil {
		toSerialize["allocatedIPCount"] = o.AllocatedIpCount
	}
	if o.TotalIpCount != nil {
		toSerialize["totalIPCount"] = o.TotalIpCount
	}
	if o.AvailableIpCount != nil {
		toSerialize["availableIPCount"] = o.AvailableIpCount
	}
	if o.CapacityReliable != nil {
		toSerialize["capacityReliable"] = o.CapacityReliable
	}
	if o.AffinitySummary != nil {
		toSerialize["affinitySummary"] = o.AffinitySummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IpPool) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string         `json:"name,omitempty"`
		IpFamily         *IpPoolIPFamily `json:"ipFamily,omitempty"`
		Subnet           *string         `json:"subnet,omitempty"`
		Disabled         *bool           `json:"disabled,omitempty"`
		Default          *bool           `json:"default,omitempty"`
		AllocatedIpCount *int64          `json:"allocatedIPCount,omitempty"`
		TotalIpCount     *int64          `json:"totalIPCount,omitempty"`
		AvailableIpCount *int64          `json:"availableIPCount,omitempty"`
		CapacityReliable *bool           `json:"capacityReliable,omitempty"`
		AffinitySummary  []string        `json:"affinitySummary,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "ipFamily", "subnet", "disabled", "default", "allocatedIPCount", "totalIPCount", "availableIPCount", "capacityReliable", "affinitySummary"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.IpFamily != nil && !all.IpFamily.IsValid() {
		hasInvalidField = true
	} else {
		o.IpFamily = all.IpFamily
	}
	o.Subnet = all.Subnet
	o.Disabled = all.Disabled
	o.Default = all.Default
	o.AllocatedIpCount = all.AllocatedIpCount
	o.TotalIpCount = all.TotalIpCount
	o.AvailableIpCount = all.AvailableIpCount
	o.CapacityReliable = all.CapacityReliable
	o.AffinitySummary = all.AffinitySummary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
