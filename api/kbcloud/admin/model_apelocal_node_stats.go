// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ApelocalNodeStats apelocal-specific capacity measurements for one node. Missing measurements are omitted.
type ApelocalNodeStats struct {
	// Node-local backing allocation resource identity, used to deduplicate SC aliases.
	AllocationResourceId *string `json:"allocationResourceID,omitempty"`
	// Node-local observed filesystem identity, used to deduplicate shared physical capacity.
	PhysicalResourceId *string `json:"physicalResourceID,omitempty"`
	// Observed filesystem capacity in GiB; omitted when the backing filesystem cannot be identified completely.
	PhysicalCapacity *float64 `json:"physicalCapacity,omitempty"`
	// Observed filesystem used bytes in GiB, including other data on the same filesystem.
	PhysicalUsage *float64 `json:"physicalUsage,omitempty"`
	// CSI total allocatable capacity in GiB; omitted for unlimited capacity.
	AllocatableCapacity *float64 `json:"allocatableCapacity,omitempty"`
	// CSI remaining allocatable capacity in GiB; omitted for unlimited capacity.
	AvailableCapacity *float64 `json:"availableCapacity,omitempty"`
	// Capacity accounted for by CSI in GiB, including shared definitions; omitted when it cannot be determined exactly.
	AllocatedCapacity *float64 `json:"allocatedCapacity,omitempty"`
	// Largest volume CSI can allocate on this node in GiB.
	MaximumVolumeSize *float64 `json:"maximumVolumeSize,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApelocalNodeStats instantiates a new ApelocalNodeStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApelocalNodeStats() *ApelocalNodeStats {
	this := ApelocalNodeStats{}
	return &this
}

// NewApelocalNodeStatsWithDefaults instantiates a new ApelocalNodeStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApelocalNodeStatsWithDefaults() *ApelocalNodeStats {
	this := ApelocalNodeStats{}
	return &this
}

// GetAllocationResourceId returns the AllocationResourceId field value if set, zero value otherwise.
func (o *ApelocalNodeStats) GetAllocationResourceId() string {
	if o == nil || o.AllocationResourceId == nil {
		var ret string
		return ret
	}
	return *o.AllocationResourceId
}

// GetAllocationResourceIdOk returns a tuple with the AllocationResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalNodeStats) GetAllocationResourceIdOk() (*string, bool) {
	if o == nil || o.AllocationResourceId == nil {
		return nil, false
	}
	return o.AllocationResourceId, true
}

// HasAllocationResourceId returns a boolean if a field has been set.
func (o *ApelocalNodeStats) HasAllocationResourceId() bool {
	return o != nil && o.AllocationResourceId != nil
}

// SetAllocationResourceId gets a reference to the given string and assigns it to the AllocationResourceId field.
func (o *ApelocalNodeStats) SetAllocationResourceId(v string) {
	o.AllocationResourceId = &v
}

// GetPhysicalResourceId returns the PhysicalResourceId field value if set, zero value otherwise.
func (o *ApelocalNodeStats) GetPhysicalResourceId() string {
	if o == nil || o.PhysicalResourceId == nil {
		var ret string
		return ret
	}
	return *o.PhysicalResourceId
}

// GetPhysicalResourceIdOk returns a tuple with the PhysicalResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalNodeStats) GetPhysicalResourceIdOk() (*string, bool) {
	if o == nil || o.PhysicalResourceId == nil {
		return nil, false
	}
	return o.PhysicalResourceId, true
}

// HasPhysicalResourceId returns a boolean if a field has been set.
func (o *ApelocalNodeStats) HasPhysicalResourceId() bool {
	return o != nil && o.PhysicalResourceId != nil
}

// SetPhysicalResourceId gets a reference to the given string and assigns it to the PhysicalResourceId field.
func (o *ApelocalNodeStats) SetPhysicalResourceId(v string) {
	o.PhysicalResourceId = &v
}

// GetPhysicalCapacity returns the PhysicalCapacity field value if set, zero value otherwise.
func (o *ApelocalNodeStats) GetPhysicalCapacity() float64 {
	if o == nil || o.PhysicalCapacity == nil {
		var ret float64
		return ret
	}
	return *o.PhysicalCapacity
}

// GetPhysicalCapacityOk returns a tuple with the PhysicalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalNodeStats) GetPhysicalCapacityOk() (*float64, bool) {
	if o == nil || o.PhysicalCapacity == nil {
		return nil, false
	}
	return o.PhysicalCapacity, true
}

// HasPhysicalCapacity returns a boolean if a field has been set.
func (o *ApelocalNodeStats) HasPhysicalCapacity() bool {
	return o != nil && o.PhysicalCapacity != nil
}

// SetPhysicalCapacity gets a reference to the given float64 and assigns it to the PhysicalCapacity field.
func (o *ApelocalNodeStats) SetPhysicalCapacity(v float64) {
	o.PhysicalCapacity = &v
}

// GetPhysicalUsage returns the PhysicalUsage field value if set, zero value otherwise.
func (o *ApelocalNodeStats) GetPhysicalUsage() float64 {
	if o == nil || o.PhysicalUsage == nil {
		var ret float64
		return ret
	}
	return *o.PhysicalUsage
}

// GetPhysicalUsageOk returns a tuple with the PhysicalUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalNodeStats) GetPhysicalUsageOk() (*float64, bool) {
	if o == nil || o.PhysicalUsage == nil {
		return nil, false
	}
	return o.PhysicalUsage, true
}

// HasPhysicalUsage returns a boolean if a field has been set.
func (o *ApelocalNodeStats) HasPhysicalUsage() bool {
	return o != nil && o.PhysicalUsage != nil
}

// SetPhysicalUsage gets a reference to the given float64 and assigns it to the PhysicalUsage field.
func (o *ApelocalNodeStats) SetPhysicalUsage(v float64) {
	o.PhysicalUsage = &v
}

// GetAllocatableCapacity returns the AllocatableCapacity field value if set, zero value otherwise.
func (o *ApelocalNodeStats) GetAllocatableCapacity() float64 {
	if o == nil || o.AllocatableCapacity == nil {
		var ret float64
		return ret
	}
	return *o.AllocatableCapacity
}

// GetAllocatableCapacityOk returns a tuple with the AllocatableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalNodeStats) GetAllocatableCapacityOk() (*float64, bool) {
	if o == nil || o.AllocatableCapacity == nil {
		return nil, false
	}
	return o.AllocatableCapacity, true
}

// HasAllocatableCapacity returns a boolean if a field has been set.
func (o *ApelocalNodeStats) HasAllocatableCapacity() bool {
	return o != nil && o.AllocatableCapacity != nil
}

// SetAllocatableCapacity gets a reference to the given float64 and assigns it to the AllocatableCapacity field.
func (o *ApelocalNodeStats) SetAllocatableCapacity(v float64) {
	o.AllocatableCapacity = &v
}

// GetAvailableCapacity returns the AvailableCapacity field value if set, zero value otherwise.
func (o *ApelocalNodeStats) GetAvailableCapacity() float64 {
	if o == nil || o.AvailableCapacity == nil {
		var ret float64
		return ret
	}
	return *o.AvailableCapacity
}

// GetAvailableCapacityOk returns a tuple with the AvailableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalNodeStats) GetAvailableCapacityOk() (*float64, bool) {
	if o == nil || o.AvailableCapacity == nil {
		return nil, false
	}
	return o.AvailableCapacity, true
}

// HasAvailableCapacity returns a boolean if a field has been set.
func (o *ApelocalNodeStats) HasAvailableCapacity() bool {
	return o != nil && o.AvailableCapacity != nil
}

// SetAvailableCapacity gets a reference to the given float64 and assigns it to the AvailableCapacity field.
func (o *ApelocalNodeStats) SetAvailableCapacity(v float64) {
	o.AvailableCapacity = &v
}

// GetAllocatedCapacity returns the AllocatedCapacity field value if set, zero value otherwise.
func (o *ApelocalNodeStats) GetAllocatedCapacity() float64 {
	if o == nil || o.AllocatedCapacity == nil {
		var ret float64
		return ret
	}
	return *o.AllocatedCapacity
}

// GetAllocatedCapacityOk returns a tuple with the AllocatedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalNodeStats) GetAllocatedCapacityOk() (*float64, bool) {
	if o == nil || o.AllocatedCapacity == nil {
		return nil, false
	}
	return o.AllocatedCapacity, true
}

// HasAllocatedCapacity returns a boolean if a field has been set.
func (o *ApelocalNodeStats) HasAllocatedCapacity() bool {
	return o != nil && o.AllocatedCapacity != nil
}

// SetAllocatedCapacity gets a reference to the given float64 and assigns it to the AllocatedCapacity field.
func (o *ApelocalNodeStats) SetAllocatedCapacity(v float64) {
	o.AllocatedCapacity = &v
}

// GetMaximumVolumeSize returns the MaximumVolumeSize field value if set, zero value otherwise.
func (o *ApelocalNodeStats) GetMaximumVolumeSize() float64 {
	if o == nil || o.MaximumVolumeSize == nil {
		var ret float64
		return ret
	}
	return *o.MaximumVolumeSize
}

// GetMaximumVolumeSizeOk returns a tuple with the MaximumVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalNodeStats) GetMaximumVolumeSizeOk() (*float64, bool) {
	if o == nil || o.MaximumVolumeSize == nil {
		return nil, false
	}
	return o.MaximumVolumeSize, true
}

// HasMaximumVolumeSize returns a boolean if a field has been set.
func (o *ApelocalNodeStats) HasMaximumVolumeSize() bool {
	return o != nil && o.MaximumVolumeSize != nil
}

// SetMaximumVolumeSize gets a reference to the given float64 and assigns it to the MaximumVolumeSize field.
func (o *ApelocalNodeStats) SetMaximumVolumeSize(v float64) {
	o.MaximumVolumeSize = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApelocalNodeStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.AllocationResourceId != nil {
		toSerialize["allocationResourceID"] = o.AllocationResourceId
	}
	if o.PhysicalResourceId != nil {
		toSerialize["physicalResourceID"] = o.PhysicalResourceId
	}
	if o.PhysicalCapacity != nil {
		toSerialize["physicalCapacity"] = o.PhysicalCapacity
	}
	if o.PhysicalUsage != nil {
		toSerialize["physicalUsage"] = o.PhysicalUsage
	}
	if o.AllocatableCapacity != nil {
		toSerialize["allocatableCapacity"] = o.AllocatableCapacity
	}
	if o.AvailableCapacity != nil {
		toSerialize["availableCapacity"] = o.AvailableCapacity
	}
	if o.AllocatedCapacity != nil {
		toSerialize["allocatedCapacity"] = o.AllocatedCapacity
	}
	if o.MaximumVolumeSize != nil {
		toSerialize["maximumVolumeSize"] = o.MaximumVolumeSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApelocalNodeStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllocationResourceId *string  `json:"allocationResourceID,omitempty"`
		PhysicalResourceId   *string  `json:"physicalResourceID,omitempty"`
		PhysicalCapacity     *float64 `json:"physicalCapacity,omitempty"`
		PhysicalUsage        *float64 `json:"physicalUsage,omitempty"`
		AllocatableCapacity  *float64 `json:"allocatableCapacity,omitempty"`
		AvailableCapacity    *float64 `json:"availableCapacity,omitempty"`
		AllocatedCapacity    *float64 `json:"allocatedCapacity,omitempty"`
		MaximumVolumeSize    *float64 `json:"maximumVolumeSize,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"allocationResourceID", "physicalResourceID", "physicalCapacity", "physicalUsage", "allocatableCapacity", "availableCapacity", "allocatedCapacity", "maximumVolumeSize"})
	} else {
		return err
	}
	o.AllocationResourceId = all.AllocationResourceId
	o.PhysicalResourceId = all.PhysicalResourceId
	o.PhysicalCapacity = all.PhysicalCapacity
	o.PhysicalUsage = all.PhysicalUsage
	o.AllocatableCapacity = all.AllocatableCapacity
	o.AvailableCapacity = all.AvailableCapacity
	o.AllocatedCapacity = all.AllocatedCapacity
	o.MaximumVolumeSize = all.MaximumVolumeSize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
