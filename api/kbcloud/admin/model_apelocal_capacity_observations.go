// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ApelocalCapacityObservations Known-node aggregates after resource deduplication. Capacity values are in GiB and ratios are percentages using only nodes with both inputs. Independent measurements can cover different nodes. Quota, remaining quota and allocation percentage objects are omitted for unlimited or conflicting policies. Remaining quota covers only ready schedulable nodes. Physical remaining is capacity minus usage from nodes with both measurements and usage not exceeding capacity; it is not a guaranteed writable or schedulable quota. Existing summary totals remain omitted when incomplete.
type ApelocalCapacityObservations struct {
	// A measurement over known nodes. Missing nodes are not treated as zero. Coverage counts nodes whose complete per-node inputs are known, not individual storage resources.
	PhysicalCapacity *ApelocalCapacityObservation `json:"physicalCapacity,omitempty"`
	// A measurement over known nodes. Missing nodes are not treated as zero. Coverage counts nodes whose complete per-node inputs are known, not individual storage resources.
	PhysicalUsage *ApelocalCapacityObservation `json:"physicalUsage,omitempty"`
	// A measurement over known nodes. Missing nodes are not treated as zero. Coverage counts nodes whose complete per-node inputs are known, not individual storage resources.
	PhysicalRemaining *ApelocalCapacityObservation `json:"physicalRemaining,omitempty"`
	// A measurement over known nodes. Missing nodes are not treated as zero. Coverage counts nodes whose complete per-node inputs are known, not individual storage resources.
	AllocatedCapacity *ApelocalCapacityObservation `json:"allocatedCapacity,omitempty"`
	// A measurement over known nodes. Missing nodes are not treated as zero. Coverage counts nodes whose complete per-node inputs are known, not individual storage resources.
	AllocatableCapacity *ApelocalCapacityObservation `json:"allocatableCapacity,omitempty"`
	// A measurement over known nodes. Missing nodes are not treated as zero. Coverage counts nodes whose complete per-node inputs are known, not individual storage resources.
	AvailableCapacity *ApelocalCapacityObservation `json:"availableCapacity,omitempty"`
	// A measurement over known nodes. Missing nodes are not treated as zero. Coverage counts nodes whose complete per-node inputs are known, not individual storage resources.
	PhysicalUsagePercent *ApelocalCapacityObservation `json:"physicalUsagePercent,omitempty"`
	// A measurement over known nodes. Missing nodes are not treated as zero. Coverage counts nodes whose complete per-node inputs are known, not individual storage resources.
	AllocationPercent *ApelocalCapacityObservation `json:"allocationPercent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApelocalCapacityObservations instantiates a new ApelocalCapacityObservations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApelocalCapacityObservations() *ApelocalCapacityObservations {
	this := ApelocalCapacityObservations{}
	return &this
}

// NewApelocalCapacityObservationsWithDefaults instantiates a new ApelocalCapacityObservations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApelocalCapacityObservationsWithDefaults() *ApelocalCapacityObservations {
	this := ApelocalCapacityObservations{}
	return &this
}

// GetPhysicalCapacity returns the PhysicalCapacity field value if set, zero value otherwise.
func (o *ApelocalCapacityObservations) GetPhysicalCapacity() ApelocalCapacityObservation {
	if o == nil || o.PhysicalCapacity == nil {
		var ret ApelocalCapacityObservation
		return ret
	}
	return *o.PhysicalCapacity
}

// GetPhysicalCapacityOk returns a tuple with the PhysicalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservations) GetPhysicalCapacityOk() (*ApelocalCapacityObservation, bool) {
	if o == nil || o.PhysicalCapacity == nil {
		return nil, false
	}
	return o.PhysicalCapacity, true
}

// HasPhysicalCapacity returns a boolean if a field has been set.
func (o *ApelocalCapacityObservations) HasPhysicalCapacity() bool {
	return o != nil && o.PhysicalCapacity != nil
}

// SetPhysicalCapacity gets a reference to the given ApelocalCapacityObservation and assigns it to the PhysicalCapacity field.
func (o *ApelocalCapacityObservations) SetPhysicalCapacity(v ApelocalCapacityObservation) {
	o.PhysicalCapacity = &v
}

// GetPhysicalUsage returns the PhysicalUsage field value if set, zero value otherwise.
func (o *ApelocalCapacityObservations) GetPhysicalUsage() ApelocalCapacityObservation {
	if o == nil || o.PhysicalUsage == nil {
		var ret ApelocalCapacityObservation
		return ret
	}
	return *o.PhysicalUsage
}

// GetPhysicalUsageOk returns a tuple with the PhysicalUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservations) GetPhysicalUsageOk() (*ApelocalCapacityObservation, bool) {
	if o == nil || o.PhysicalUsage == nil {
		return nil, false
	}
	return o.PhysicalUsage, true
}

// HasPhysicalUsage returns a boolean if a field has been set.
func (o *ApelocalCapacityObservations) HasPhysicalUsage() bool {
	return o != nil && o.PhysicalUsage != nil
}

// SetPhysicalUsage gets a reference to the given ApelocalCapacityObservation and assigns it to the PhysicalUsage field.
func (o *ApelocalCapacityObservations) SetPhysicalUsage(v ApelocalCapacityObservation) {
	o.PhysicalUsage = &v
}

// GetPhysicalRemaining returns the PhysicalRemaining field value if set, zero value otherwise.
func (o *ApelocalCapacityObservations) GetPhysicalRemaining() ApelocalCapacityObservation {
	if o == nil || o.PhysicalRemaining == nil {
		var ret ApelocalCapacityObservation
		return ret
	}
	return *o.PhysicalRemaining
}

// GetPhysicalRemainingOk returns a tuple with the PhysicalRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservations) GetPhysicalRemainingOk() (*ApelocalCapacityObservation, bool) {
	if o == nil || o.PhysicalRemaining == nil {
		return nil, false
	}
	return o.PhysicalRemaining, true
}

// HasPhysicalRemaining returns a boolean if a field has been set.
func (o *ApelocalCapacityObservations) HasPhysicalRemaining() bool {
	return o != nil && o.PhysicalRemaining != nil
}

// SetPhysicalRemaining gets a reference to the given ApelocalCapacityObservation and assigns it to the PhysicalRemaining field.
func (o *ApelocalCapacityObservations) SetPhysicalRemaining(v ApelocalCapacityObservation) {
	o.PhysicalRemaining = &v
}

// GetAllocatedCapacity returns the AllocatedCapacity field value if set, zero value otherwise.
func (o *ApelocalCapacityObservations) GetAllocatedCapacity() ApelocalCapacityObservation {
	if o == nil || o.AllocatedCapacity == nil {
		var ret ApelocalCapacityObservation
		return ret
	}
	return *o.AllocatedCapacity
}

// GetAllocatedCapacityOk returns a tuple with the AllocatedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservations) GetAllocatedCapacityOk() (*ApelocalCapacityObservation, bool) {
	if o == nil || o.AllocatedCapacity == nil {
		return nil, false
	}
	return o.AllocatedCapacity, true
}

// HasAllocatedCapacity returns a boolean if a field has been set.
func (o *ApelocalCapacityObservations) HasAllocatedCapacity() bool {
	return o != nil && o.AllocatedCapacity != nil
}

// SetAllocatedCapacity gets a reference to the given ApelocalCapacityObservation and assigns it to the AllocatedCapacity field.
func (o *ApelocalCapacityObservations) SetAllocatedCapacity(v ApelocalCapacityObservation) {
	o.AllocatedCapacity = &v
}

// GetAllocatableCapacity returns the AllocatableCapacity field value if set, zero value otherwise.
func (o *ApelocalCapacityObservations) GetAllocatableCapacity() ApelocalCapacityObservation {
	if o == nil || o.AllocatableCapacity == nil {
		var ret ApelocalCapacityObservation
		return ret
	}
	return *o.AllocatableCapacity
}

// GetAllocatableCapacityOk returns a tuple with the AllocatableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservations) GetAllocatableCapacityOk() (*ApelocalCapacityObservation, bool) {
	if o == nil || o.AllocatableCapacity == nil {
		return nil, false
	}
	return o.AllocatableCapacity, true
}

// HasAllocatableCapacity returns a boolean if a field has been set.
func (o *ApelocalCapacityObservations) HasAllocatableCapacity() bool {
	return o != nil && o.AllocatableCapacity != nil
}

// SetAllocatableCapacity gets a reference to the given ApelocalCapacityObservation and assigns it to the AllocatableCapacity field.
func (o *ApelocalCapacityObservations) SetAllocatableCapacity(v ApelocalCapacityObservation) {
	o.AllocatableCapacity = &v
}

// GetAvailableCapacity returns the AvailableCapacity field value if set, zero value otherwise.
func (o *ApelocalCapacityObservations) GetAvailableCapacity() ApelocalCapacityObservation {
	if o == nil || o.AvailableCapacity == nil {
		var ret ApelocalCapacityObservation
		return ret
	}
	return *o.AvailableCapacity
}

// GetAvailableCapacityOk returns a tuple with the AvailableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservations) GetAvailableCapacityOk() (*ApelocalCapacityObservation, bool) {
	if o == nil || o.AvailableCapacity == nil {
		return nil, false
	}
	return o.AvailableCapacity, true
}

// HasAvailableCapacity returns a boolean if a field has been set.
func (o *ApelocalCapacityObservations) HasAvailableCapacity() bool {
	return o != nil && o.AvailableCapacity != nil
}

// SetAvailableCapacity gets a reference to the given ApelocalCapacityObservation and assigns it to the AvailableCapacity field.
func (o *ApelocalCapacityObservations) SetAvailableCapacity(v ApelocalCapacityObservation) {
	o.AvailableCapacity = &v
}

// GetPhysicalUsagePercent returns the PhysicalUsagePercent field value if set, zero value otherwise.
func (o *ApelocalCapacityObservations) GetPhysicalUsagePercent() ApelocalCapacityObservation {
	if o == nil || o.PhysicalUsagePercent == nil {
		var ret ApelocalCapacityObservation
		return ret
	}
	return *o.PhysicalUsagePercent
}

// GetPhysicalUsagePercentOk returns a tuple with the PhysicalUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservations) GetPhysicalUsagePercentOk() (*ApelocalCapacityObservation, bool) {
	if o == nil || o.PhysicalUsagePercent == nil {
		return nil, false
	}
	return o.PhysicalUsagePercent, true
}

// HasPhysicalUsagePercent returns a boolean if a field has been set.
func (o *ApelocalCapacityObservations) HasPhysicalUsagePercent() bool {
	return o != nil && o.PhysicalUsagePercent != nil
}

// SetPhysicalUsagePercent gets a reference to the given ApelocalCapacityObservation and assigns it to the PhysicalUsagePercent field.
func (o *ApelocalCapacityObservations) SetPhysicalUsagePercent(v ApelocalCapacityObservation) {
	o.PhysicalUsagePercent = &v
}

// GetAllocationPercent returns the AllocationPercent field value if set, zero value otherwise.
func (o *ApelocalCapacityObservations) GetAllocationPercent() ApelocalCapacityObservation {
	if o == nil || o.AllocationPercent == nil {
		var ret ApelocalCapacityObservation
		return ret
	}
	return *o.AllocationPercent
}

// GetAllocationPercentOk returns a tuple with the AllocationPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApelocalCapacityObservations) GetAllocationPercentOk() (*ApelocalCapacityObservation, bool) {
	if o == nil || o.AllocationPercent == nil {
		return nil, false
	}
	return o.AllocationPercent, true
}

// HasAllocationPercent returns a boolean if a field has been set.
func (o *ApelocalCapacityObservations) HasAllocationPercent() bool {
	return o != nil && o.AllocationPercent != nil
}

// SetAllocationPercent gets a reference to the given ApelocalCapacityObservation and assigns it to the AllocationPercent field.
func (o *ApelocalCapacityObservations) SetAllocationPercent(v ApelocalCapacityObservation) {
	o.AllocationPercent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApelocalCapacityObservations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.PhysicalCapacity != nil {
		toSerialize["physicalCapacity"] = o.PhysicalCapacity
	}
	if o.PhysicalUsage != nil {
		toSerialize["physicalUsage"] = o.PhysicalUsage
	}
	if o.PhysicalRemaining != nil {
		toSerialize["physicalRemaining"] = o.PhysicalRemaining
	}
	if o.AllocatedCapacity != nil {
		toSerialize["allocatedCapacity"] = o.AllocatedCapacity
	}
	if o.AllocatableCapacity != nil {
		toSerialize["allocatableCapacity"] = o.AllocatableCapacity
	}
	if o.AvailableCapacity != nil {
		toSerialize["availableCapacity"] = o.AvailableCapacity
	}
	if o.PhysicalUsagePercent != nil {
		toSerialize["physicalUsagePercent"] = o.PhysicalUsagePercent
	}
	if o.AllocationPercent != nil {
		toSerialize["allocationPercent"] = o.AllocationPercent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApelocalCapacityObservations) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PhysicalCapacity     *ApelocalCapacityObservation `json:"physicalCapacity,omitempty"`
		PhysicalUsage        *ApelocalCapacityObservation `json:"physicalUsage,omitempty"`
		PhysicalRemaining    *ApelocalCapacityObservation `json:"physicalRemaining,omitempty"`
		AllocatedCapacity    *ApelocalCapacityObservation `json:"allocatedCapacity,omitempty"`
		AllocatableCapacity  *ApelocalCapacityObservation `json:"allocatableCapacity,omitempty"`
		AvailableCapacity    *ApelocalCapacityObservation `json:"availableCapacity,omitempty"`
		PhysicalUsagePercent *ApelocalCapacityObservation `json:"physicalUsagePercent,omitempty"`
		AllocationPercent    *ApelocalCapacityObservation `json:"allocationPercent,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"physicalCapacity", "physicalUsage", "physicalRemaining", "allocatedCapacity", "allocatableCapacity", "availableCapacity", "physicalUsagePercent", "allocationPercent"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.PhysicalCapacity != nil && all.PhysicalCapacity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PhysicalCapacity = all.PhysicalCapacity
	if all.PhysicalUsage != nil && all.PhysicalUsage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PhysicalUsage = all.PhysicalUsage
	if all.PhysicalRemaining != nil && all.PhysicalRemaining.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PhysicalRemaining = all.PhysicalRemaining
	if all.AllocatedCapacity != nil && all.AllocatedCapacity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AllocatedCapacity = all.AllocatedCapacity
	if all.AllocatableCapacity != nil && all.AllocatableCapacity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AllocatableCapacity = all.AllocatableCapacity
	if all.AvailableCapacity != nil && all.AvailableCapacity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AvailableCapacity = all.AvailableCapacity
	if all.PhysicalUsagePercent != nil && all.PhysicalUsagePercent.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PhysicalUsagePercent = all.PhysicalUsagePercent
	if all.AllocationPercent != nil && all.AllocationPercent.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AllocationPercent = all.AllocationPercent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
