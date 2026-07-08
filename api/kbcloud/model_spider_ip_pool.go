// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// SpiderIPPool Spiderpool IPv4 IPPool summary.
type SpiderIPPool struct {
	Name             *string  `json:"name,omitempty"`
	IpVersion        *int32   `json:"ipVersion,omitempty"`
	Subnet           *string  `json:"subnet,omitempty"`
	Disabled         *bool    `json:"disabled,omitempty"`
	Default          *bool    `json:"default,omitempty"`
	AllocatedIpCount *int64   `json:"allocatedIPCount,omitempty"`
	TotalIpCount     *int64   `json:"totalIPCount,omitempty"`
	AvailableIpCount *int64   `json:"availableIPCount,omitempty"`
	CapacityReliable *bool    `json:"capacityReliable,omitempty"`
	AffinitySummary  []string `json:"affinitySummary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSpiderIPPool instantiates a new SpiderIPPool object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSpiderIPPool() *SpiderIPPool {
	this := SpiderIPPool{}
	return &this
}

// NewSpiderIPPoolWithDefaults instantiates a new SpiderIPPool object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSpiderIPPoolWithDefaults() *SpiderIPPool {
	this := SpiderIPPool{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SpiderIPPool) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderIPPool) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SpiderIPPool) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SpiderIPPool) SetName(v string) {
	o.Name = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *SpiderIPPool) GetIpVersion() int32 {
	if o == nil || o.IpVersion == nil {
		var ret int32
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderIPPool) GetIpVersionOk() (*int32, bool) {
	if o == nil || o.IpVersion == nil {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *SpiderIPPool) HasIpVersion() bool {
	return o != nil && o.IpVersion != nil
}

// SetIpVersion gets a reference to the given int32 and assigns it to the IpVersion field.
func (o *SpiderIPPool) SetIpVersion(v int32) {
	o.IpVersion = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *SpiderIPPool) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderIPPool) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *SpiderIPPool) HasSubnet() bool {
	return o != nil && o.Subnet != nil
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *SpiderIPPool) SetSubnet(v string) {
	o.Subnet = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *SpiderIPPool) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderIPPool) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *SpiderIPPool) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *SpiderIPPool) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *SpiderIPPool) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderIPPool) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *SpiderIPPool) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *SpiderIPPool) SetDefault(v bool) {
	o.Default = &v
}

// GetAllocatedIpCount returns the AllocatedIpCount field value if set, zero value otherwise.
func (o *SpiderIPPool) GetAllocatedIpCount() int64 {
	if o == nil || o.AllocatedIpCount == nil {
		var ret int64
		return ret
	}
	return *o.AllocatedIpCount
}

// GetAllocatedIpCountOk returns a tuple with the AllocatedIpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderIPPool) GetAllocatedIpCountOk() (*int64, bool) {
	if o == nil || o.AllocatedIpCount == nil {
		return nil, false
	}
	return o.AllocatedIpCount, true
}

// HasAllocatedIpCount returns a boolean if a field has been set.
func (o *SpiderIPPool) HasAllocatedIpCount() bool {
	return o != nil && o.AllocatedIpCount != nil
}

// SetAllocatedIpCount gets a reference to the given int64 and assigns it to the AllocatedIpCount field.
func (o *SpiderIPPool) SetAllocatedIpCount(v int64) {
	o.AllocatedIpCount = &v
}

// GetTotalIpCount returns the TotalIpCount field value if set, zero value otherwise.
func (o *SpiderIPPool) GetTotalIpCount() int64 {
	if o == nil || o.TotalIpCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalIpCount
}

// GetTotalIpCountOk returns a tuple with the TotalIpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderIPPool) GetTotalIpCountOk() (*int64, bool) {
	if o == nil || o.TotalIpCount == nil {
		return nil, false
	}
	return o.TotalIpCount, true
}

// HasTotalIpCount returns a boolean if a field has been set.
func (o *SpiderIPPool) HasTotalIpCount() bool {
	return o != nil && o.TotalIpCount != nil
}

// SetTotalIpCount gets a reference to the given int64 and assigns it to the TotalIpCount field.
func (o *SpiderIPPool) SetTotalIpCount(v int64) {
	o.TotalIpCount = &v
}

// GetAvailableIpCount returns the AvailableIpCount field value if set, zero value otherwise.
func (o *SpiderIPPool) GetAvailableIpCount() int64 {
	if o == nil || o.AvailableIpCount == nil {
		var ret int64
		return ret
	}
	return *o.AvailableIpCount
}

// GetAvailableIpCountOk returns a tuple with the AvailableIpCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderIPPool) GetAvailableIpCountOk() (*int64, bool) {
	if o == nil || o.AvailableIpCount == nil {
		return nil, false
	}
	return o.AvailableIpCount, true
}

// HasAvailableIpCount returns a boolean if a field has been set.
func (o *SpiderIPPool) HasAvailableIpCount() bool {
	return o != nil && o.AvailableIpCount != nil
}

// SetAvailableIpCount gets a reference to the given int64 and assigns it to the AvailableIpCount field.
func (o *SpiderIPPool) SetAvailableIpCount(v int64) {
	o.AvailableIpCount = &v
}

// GetCapacityReliable returns the CapacityReliable field value if set, zero value otherwise.
func (o *SpiderIPPool) GetCapacityReliable() bool {
	if o == nil || o.CapacityReliable == nil {
		var ret bool
		return ret
	}
	return *o.CapacityReliable
}

// GetCapacityReliableOk returns a tuple with the CapacityReliable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderIPPool) GetCapacityReliableOk() (*bool, bool) {
	if o == nil || o.CapacityReliable == nil {
		return nil, false
	}
	return o.CapacityReliable, true
}

// HasCapacityReliable returns a boolean if a field has been set.
func (o *SpiderIPPool) HasCapacityReliable() bool {
	return o != nil && o.CapacityReliable != nil
}

// SetCapacityReliable gets a reference to the given bool and assigns it to the CapacityReliable field.
func (o *SpiderIPPool) SetCapacityReliable(v bool) {
	o.CapacityReliable = &v
}

// GetAffinitySummary returns the AffinitySummary field value if set, zero value otherwise.
func (o *SpiderIPPool) GetAffinitySummary() []string {
	if o == nil || o.AffinitySummary == nil {
		var ret []string
		return ret
	}
	return o.AffinitySummary
}

// GetAffinitySummaryOk returns a tuple with the AffinitySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpiderIPPool) GetAffinitySummaryOk() (*[]string, bool) {
	if o == nil || o.AffinitySummary == nil {
		return nil, false
	}
	return &o.AffinitySummary, true
}

// HasAffinitySummary returns a boolean if a field has been set.
func (o *SpiderIPPool) HasAffinitySummary() bool {
	return o != nil && o.AffinitySummary != nil
}

// SetAffinitySummary gets a reference to the given []string and assigns it to the AffinitySummary field.
func (o *SpiderIPPool) SetAffinitySummary(v []string) {
	o.AffinitySummary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SpiderIPPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.IpVersion != nil {
		toSerialize["ipVersion"] = o.IpVersion
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
func (o *SpiderIPPool) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string  `json:"name,omitempty"`
		IpVersion        *int32   `json:"ipVersion,omitempty"`
		Subnet           *string  `json:"subnet,omitempty"`
		Disabled         *bool    `json:"disabled,omitempty"`
		Default          *bool    `json:"default,omitempty"`
		AllocatedIpCount *int64   `json:"allocatedIPCount,omitempty"`
		TotalIpCount     *int64   `json:"totalIPCount,omitempty"`
		AvailableIpCount *int64   `json:"availableIPCount,omitempty"`
		CapacityReliable *bool    `json:"capacityReliable,omitempty"`
		AffinitySummary  []string `json:"affinitySummary,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "ipVersion", "subnet", "disabled", "default", "allocatedIPCount", "totalIPCount", "availableIPCount", "capacityReliable", "affinitySummary"})
	} else {
		return err
	}
	o.Name = all.Name
	o.IpVersion = all.IpVersion
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

	return nil
}
