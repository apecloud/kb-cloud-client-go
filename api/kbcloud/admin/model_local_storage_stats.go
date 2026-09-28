// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// LocalStorageStats Deduplicated apelocal statistics. Does not include other storage drivers. Missing measurements are omitted, never treated as zero.
type LocalStorageStats struct {
	// Conflicting shared quotas cannot form a single logical total.
	CapacityMode *StorageCapacitySummaryMode `json:"capacityMode,omitempty"`
	// Whether all required measurements for the selected capacity mode are present.
	Complete             *bool  `json:"complete,omitempty"`
	NodeCount            *int32 `json:"nodeCount,omitempty"`
	UnavailableNodeCount *int32 `json:"unavailableNodeCount,omitempty"`
	// Physical capacity in GiB, deduplicated by backing filesystem.
	PhysicalCapacity *float64 `json:"physicalCapacity,omitempty"`
	// Physical filesystem usage in GiB, including other files.
	PhysicalUsage *float64 `json:"physicalUsage,omitempty"`
	// Logical total in GiB; omitted for unlimited or conflicting quotas.
	AllocatableCapacity *float64 `json:"allocatableCapacity,omitempty"`
	// CSI allocated capacity in GiB, deduplicated by backing allocation resource.
	AllocatedCapacity *float64 `json:"allocatedCapacity,omitempty"`
	// Remaining logical capacity in GiB on ready schedulable nodes only.
	AvailableCapacity *float64 `json:"availableCapacity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLocalStorageStats instantiates a new LocalStorageStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLocalStorageStats() *LocalStorageStats {
	this := LocalStorageStats{}
	return &this
}

// NewLocalStorageStatsWithDefaults instantiates a new LocalStorageStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLocalStorageStatsWithDefaults() *LocalStorageStats {
	this := LocalStorageStats{}
	return &this
}

// GetCapacityMode returns the CapacityMode field value if set, zero value otherwise.
func (o *LocalStorageStats) GetCapacityMode() StorageCapacitySummaryMode {
	if o == nil || o.CapacityMode == nil {
		var ret StorageCapacitySummaryMode
		return ret
	}
	return *o.CapacityMode
}

// GetCapacityModeOk returns a tuple with the CapacityMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalStorageStats) GetCapacityModeOk() (*StorageCapacitySummaryMode, bool) {
	if o == nil || o.CapacityMode == nil {
		return nil, false
	}
	return o.CapacityMode, true
}

// HasCapacityMode returns a boolean if a field has been set.
func (o *LocalStorageStats) HasCapacityMode() bool {
	return o != nil && o.CapacityMode != nil
}

// SetCapacityMode gets a reference to the given StorageCapacitySummaryMode and assigns it to the CapacityMode field.
func (o *LocalStorageStats) SetCapacityMode(v StorageCapacitySummaryMode) {
	o.CapacityMode = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *LocalStorageStats) GetComplete() bool {
	if o == nil || o.Complete == nil {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalStorageStats) GetCompleteOk() (*bool, bool) {
	if o == nil || o.Complete == nil {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *LocalStorageStats) HasComplete() bool {
	return o != nil && o.Complete != nil
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *LocalStorageStats) SetComplete(v bool) {
	o.Complete = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *LocalStorageStats) GetNodeCount() int32 {
	if o == nil || o.NodeCount == nil {
		var ret int32
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalStorageStats) GetNodeCountOk() (*int32, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *LocalStorageStats) HasNodeCount() bool {
	return o != nil && o.NodeCount != nil
}

// SetNodeCount gets a reference to the given int32 and assigns it to the NodeCount field.
func (o *LocalStorageStats) SetNodeCount(v int32) {
	o.NodeCount = &v
}

// GetUnavailableNodeCount returns the UnavailableNodeCount field value if set, zero value otherwise.
func (o *LocalStorageStats) GetUnavailableNodeCount() int32 {
	if o == nil || o.UnavailableNodeCount == nil {
		var ret int32
		return ret
	}
	return *o.UnavailableNodeCount
}

// GetUnavailableNodeCountOk returns a tuple with the UnavailableNodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalStorageStats) GetUnavailableNodeCountOk() (*int32, bool) {
	if o == nil || o.UnavailableNodeCount == nil {
		return nil, false
	}
	return o.UnavailableNodeCount, true
}

// HasUnavailableNodeCount returns a boolean if a field has been set.
func (o *LocalStorageStats) HasUnavailableNodeCount() bool {
	return o != nil && o.UnavailableNodeCount != nil
}

// SetUnavailableNodeCount gets a reference to the given int32 and assigns it to the UnavailableNodeCount field.
func (o *LocalStorageStats) SetUnavailableNodeCount(v int32) {
	o.UnavailableNodeCount = &v
}

// GetPhysicalCapacity returns the PhysicalCapacity field value if set, zero value otherwise.
func (o *LocalStorageStats) GetPhysicalCapacity() float64 {
	if o == nil || o.PhysicalCapacity == nil {
		var ret float64
		return ret
	}
	return *o.PhysicalCapacity
}

// GetPhysicalCapacityOk returns a tuple with the PhysicalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalStorageStats) GetPhysicalCapacityOk() (*float64, bool) {
	if o == nil || o.PhysicalCapacity == nil {
		return nil, false
	}
	return o.PhysicalCapacity, true
}

// HasPhysicalCapacity returns a boolean if a field has been set.
func (o *LocalStorageStats) HasPhysicalCapacity() bool {
	return o != nil && o.PhysicalCapacity != nil
}

// SetPhysicalCapacity gets a reference to the given float64 and assigns it to the PhysicalCapacity field.
func (o *LocalStorageStats) SetPhysicalCapacity(v float64) {
	o.PhysicalCapacity = &v
}

// GetPhysicalUsage returns the PhysicalUsage field value if set, zero value otherwise.
func (o *LocalStorageStats) GetPhysicalUsage() float64 {
	if o == nil || o.PhysicalUsage == nil {
		var ret float64
		return ret
	}
	return *o.PhysicalUsage
}

// GetPhysicalUsageOk returns a tuple with the PhysicalUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalStorageStats) GetPhysicalUsageOk() (*float64, bool) {
	if o == nil || o.PhysicalUsage == nil {
		return nil, false
	}
	return o.PhysicalUsage, true
}

// HasPhysicalUsage returns a boolean if a field has been set.
func (o *LocalStorageStats) HasPhysicalUsage() bool {
	return o != nil && o.PhysicalUsage != nil
}

// SetPhysicalUsage gets a reference to the given float64 and assigns it to the PhysicalUsage field.
func (o *LocalStorageStats) SetPhysicalUsage(v float64) {
	o.PhysicalUsage = &v
}

// GetAllocatableCapacity returns the AllocatableCapacity field value if set, zero value otherwise.
func (o *LocalStorageStats) GetAllocatableCapacity() float64 {
	if o == nil || o.AllocatableCapacity == nil {
		var ret float64
		return ret
	}
	return *o.AllocatableCapacity
}

// GetAllocatableCapacityOk returns a tuple with the AllocatableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalStorageStats) GetAllocatableCapacityOk() (*float64, bool) {
	if o == nil || o.AllocatableCapacity == nil {
		return nil, false
	}
	return o.AllocatableCapacity, true
}

// HasAllocatableCapacity returns a boolean if a field has been set.
func (o *LocalStorageStats) HasAllocatableCapacity() bool {
	return o != nil && o.AllocatableCapacity != nil
}

// SetAllocatableCapacity gets a reference to the given float64 and assigns it to the AllocatableCapacity field.
func (o *LocalStorageStats) SetAllocatableCapacity(v float64) {
	o.AllocatableCapacity = &v
}

// GetAllocatedCapacity returns the AllocatedCapacity field value if set, zero value otherwise.
func (o *LocalStorageStats) GetAllocatedCapacity() float64 {
	if o == nil || o.AllocatedCapacity == nil {
		var ret float64
		return ret
	}
	return *o.AllocatedCapacity
}

// GetAllocatedCapacityOk returns a tuple with the AllocatedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalStorageStats) GetAllocatedCapacityOk() (*float64, bool) {
	if o == nil || o.AllocatedCapacity == nil {
		return nil, false
	}
	return o.AllocatedCapacity, true
}

// HasAllocatedCapacity returns a boolean if a field has been set.
func (o *LocalStorageStats) HasAllocatedCapacity() bool {
	return o != nil && o.AllocatedCapacity != nil
}

// SetAllocatedCapacity gets a reference to the given float64 and assigns it to the AllocatedCapacity field.
func (o *LocalStorageStats) SetAllocatedCapacity(v float64) {
	o.AllocatedCapacity = &v
}

// GetAvailableCapacity returns the AvailableCapacity field value if set, zero value otherwise.
func (o *LocalStorageStats) GetAvailableCapacity() float64 {
	if o == nil || o.AvailableCapacity == nil {
		var ret float64
		return ret
	}
	return *o.AvailableCapacity
}

// GetAvailableCapacityOk returns a tuple with the AvailableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalStorageStats) GetAvailableCapacityOk() (*float64, bool) {
	if o == nil || o.AvailableCapacity == nil {
		return nil, false
	}
	return o.AvailableCapacity, true
}

// HasAvailableCapacity returns a boolean if a field has been set.
func (o *LocalStorageStats) HasAvailableCapacity() bool {
	return o != nil && o.AvailableCapacity != nil
}

// SetAvailableCapacity gets a reference to the given float64 and assigns it to the AvailableCapacity field.
func (o *LocalStorageStats) SetAvailableCapacity(v float64) {
	o.AvailableCapacity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LocalStorageStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CapacityMode != nil {
		toSerialize["capacityMode"] = o.CapacityMode
	}
	if o.Complete != nil {
		toSerialize["complete"] = o.Complete
	}
	if o.NodeCount != nil {
		toSerialize["nodeCount"] = o.NodeCount
	}
	if o.UnavailableNodeCount != nil {
		toSerialize["unavailableNodeCount"] = o.UnavailableNodeCount
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
	if o.AllocatedCapacity != nil {
		toSerialize["allocatedCapacity"] = o.AllocatedCapacity
	}
	if o.AvailableCapacity != nil {
		toSerialize["availableCapacity"] = o.AvailableCapacity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LocalStorageStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CapacityMode         *StorageCapacitySummaryMode `json:"capacityMode,omitempty"`
		Complete             *bool                       `json:"complete,omitempty"`
		NodeCount            *int32                      `json:"nodeCount,omitempty"`
		UnavailableNodeCount *int32                      `json:"unavailableNodeCount,omitempty"`
		PhysicalCapacity     *float64                    `json:"physicalCapacity,omitempty"`
		PhysicalUsage        *float64                    `json:"physicalUsage,omitempty"`
		AllocatableCapacity  *float64                    `json:"allocatableCapacity,omitempty"`
		AllocatedCapacity    *float64                    `json:"allocatedCapacity,omitempty"`
		AvailableCapacity    *float64                    `json:"availableCapacity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"capacityMode", "complete", "nodeCount", "unavailableNodeCount", "physicalCapacity", "physicalUsage", "allocatableCapacity", "allocatedCapacity", "availableCapacity"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CapacityMode != nil && !all.CapacityMode.IsValid() {
		hasInvalidField = true
	} else {
		o.CapacityMode = all.CapacityMode
	}
	o.Complete = all.Complete
	o.NodeCount = all.NodeCount
	o.UnavailableNodeCount = all.UnavailableNodeCount
	o.PhysicalCapacity = all.PhysicalCapacity
	o.PhysicalUsage = all.PhysicalUsage
	o.AllocatableCapacity = all.AllocatableCapacity
	o.AllocatedCapacity = all.AllocatedCapacity
	o.AvailableCapacity = all.AvailableCapacity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
