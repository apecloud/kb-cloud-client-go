// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageClassNodeStats storageClassNodeStats is a storage class node stats.
type StorageClassNodeStats struct {
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
	// the name of the node
	NodeName string `json:"nodeName"`
	// the status of the node
	NodeStatus string `json:"nodeStatus"`
	// the total requests size of all PVCs on the node
	Requests float64 `json:"requests"`
	// the actual disk usage of hostpath on the node
	Usage float64 `json:"usage"`
	// the capacity of hostpath on the node
	Capacity float64 `json:"capacity"`
	// the number of PVCs on the node
	PvcCount int32 `json:"pvcCount"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageClassNodeStats instantiates a new StorageClassNodeStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageClassNodeStats(nodeName string, nodeStatus string, requests float64, usage float64, capacity float64, pvcCount int32) *StorageClassNodeStats {
	this := StorageClassNodeStats{}
	this.NodeName = nodeName
	this.NodeStatus = nodeStatus
	this.Requests = requests
	this.Usage = usage
	this.Capacity = capacity
	this.PvcCount = pvcCount
	return &this
}

// NewStorageClassNodeStatsWithDefaults instantiates a new StorageClassNodeStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageClassNodeStatsWithDefaults() *StorageClassNodeStats {
	this := StorageClassNodeStats{}
	return &this
}

// GetAllocationResourceId returns the AllocationResourceId field value if set, zero value otherwise.
func (o *StorageClassNodeStats) GetAllocationResourceId() string {
	if o == nil || o.AllocationResourceId == nil {
		var ret string
		return ret
	}
	return *o.AllocationResourceId
}

// GetAllocationResourceIdOk returns a tuple with the AllocationResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetAllocationResourceIdOk() (*string, bool) {
	if o == nil || o.AllocationResourceId == nil {
		return nil, false
	}
	return o.AllocationResourceId, true
}

// HasAllocationResourceId returns a boolean if a field has been set.
func (o *StorageClassNodeStats) HasAllocationResourceId() bool {
	return o != nil && o.AllocationResourceId != nil
}

// SetAllocationResourceId gets a reference to the given string and assigns it to the AllocationResourceId field.
func (o *StorageClassNodeStats) SetAllocationResourceId(v string) {
	o.AllocationResourceId = &v
}

// GetPhysicalResourceId returns the PhysicalResourceId field value if set, zero value otherwise.
func (o *StorageClassNodeStats) GetPhysicalResourceId() string {
	if o == nil || o.PhysicalResourceId == nil {
		var ret string
		return ret
	}
	return *o.PhysicalResourceId
}

// GetPhysicalResourceIdOk returns a tuple with the PhysicalResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetPhysicalResourceIdOk() (*string, bool) {
	if o == nil || o.PhysicalResourceId == nil {
		return nil, false
	}
	return o.PhysicalResourceId, true
}

// HasPhysicalResourceId returns a boolean if a field has been set.
func (o *StorageClassNodeStats) HasPhysicalResourceId() bool {
	return o != nil && o.PhysicalResourceId != nil
}

// SetPhysicalResourceId gets a reference to the given string and assigns it to the PhysicalResourceId field.
func (o *StorageClassNodeStats) SetPhysicalResourceId(v string) {
	o.PhysicalResourceId = &v
}

// GetPhysicalCapacity returns the PhysicalCapacity field value if set, zero value otherwise.
func (o *StorageClassNodeStats) GetPhysicalCapacity() float64 {
	if o == nil || o.PhysicalCapacity == nil {
		var ret float64
		return ret
	}
	return *o.PhysicalCapacity
}

// GetPhysicalCapacityOk returns a tuple with the PhysicalCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetPhysicalCapacityOk() (*float64, bool) {
	if o == nil || o.PhysicalCapacity == nil {
		return nil, false
	}
	return o.PhysicalCapacity, true
}

// HasPhysicalCapacity returns a boolean if a field has been set.
func (o *StorageClassNodeStats) HasPhysicalCapacity() bool {
	return o != nil && o.PhysicalCapacity != nil
}

// SetPhysicalCapacity gets a reference to the given float64 and assigns it to the PhysicalCapacity field.
func (o *StorageClassNodeStats) SetPhysicalCapacity(v float64) {
	o.PhysicalCapacity = &v
}

// GetPhysicalUsage returns the PhysicalUsage field value if set, zero value otherwise.
func (o *StorageClassNodeStats) GetPhysicalUsage() float64 {
	if o == nil || o.PhysicalUsage == nil {
		var ret float64
		return ret
	}
	return *o.PhysicalUsage
}

// GetPhysicalUsageOk returns a tuple with the PhysicalUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetPhysicalUsageOk() (*float64, bool) {
	if o == nil || o.PhysicalUsage == nil {
		return nil, false
	}
	return o.PhysicalUsage, true
}

// HasPhysicalUsage returns a boolean if a field has been set.
func (o *StorageClassNodeStats) HasPhysicalUsage() bool {
	return o != nil && o.PhysicalUsage != nil
}

// SetPhysicalUsage gets a reference to the given float64 and assigns it to the PhysicalUsage field.
func (o *StorageClassNodeStats) SetPhysicalUsage(v float64) {
	o.PhysicalUsage = &v
}

// GetAllocatableCapacity returns the AllocatableCapacity field value if set, zero value otherwise.
func (o *StorageClassNodeStats) GetAllocatableCapacity() float64 {
	if o == nil || o.AllocatableCapacity == nil {
		var ret float64
		return ret
	}
	return *o.AllocatableCapacity
}

// GetAllocatableCapacityOk returns a tuple with the AllocatableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetAllocatableCapacityOk() (*float64, bool) {
	if o == nil || o.AllocatableCapacity == nil {
		return nil, false
	}
	return o.AllocatableCapacity, true
}

// HasAllocatableCapacity returns a boolean if a field has been set.
func (o *StorageClassNodeStats) HasAllocatableCapacity() bool {
	return o != nil && o.AllocatableCapacity != nil
}

// SetAllocatableCapacity gets a reference to the given float64 and assigns it to the AllocatableCapacity field.
func (o *StorageClassNodeStats) SetAllocatableCapacity(v float64) {
	o.AllocatableCapacity = &v
}

// GetAvailableCapacity returns the AvailableCapacity field value if set, zero value otherwise.
func (o *StorageClassNodeStats) GetAvailableCapacity() float64 {
	if o == nil || o.AvailableCapacity == nil {
		var ret float64
		return ret
	}
	return *o.AvailableCapacity
}

// GetAvailableCapacityOk returns a tuple with the AvailableCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetAvailableCapacityOk() (*float64, bool) {
	if o == nil || o.AvailableCapacity == nil {
		return nil, false
	}
	return o.AvailableCapacity, true
}

// HasAvailableCapacity returns a boolean if a field has been set.
func (o *StorageClassNodeStats) HasAvailableCapacity() bool {
	return o != nil && o.AvailableCapacity != nil
}

// SetAvailableCapacity gets a reference to the given float64 and assigns it to the AvailableCapacity field.
func (o *StorageClassNodeStats) SetAvailableCapacity(v float64) {
	o.AvailableCapacity = &v
}

// GetAllocatedCapacity returns the AllocatedCapacity field value if set, zero value otherwise.
func (o *StorageClassNodeStats) GetAllocatedCapacity() float64 {
	if o == nil || o.AllocatedCapacity == nil {
		var ret float64
		return ret
	}
	return *o.AllocatedCapacity
}

// GetAllocatedCapacityOk returns a tuple with the AllocatedCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetAllocatedCapacityOk() (*float64, bool) {
	if o == nil || o.AllocatedCapacity == nil {
		return nil, false
	}
	return o.AllocatedCapacity, true
}

// HasAllocatedCapacity returns a boolean if a field has been set.
func (o *StorageClassNodeStats) HasAllocatedCapacity() bool {
	return o != nil && o.AllocatedCapacity != nil
}

// SetAllocatedCapacity gets a reference to the given float64 and assigns it to the AllocatedCapacity field.
func (o *StorageClassNodeStats) SetAllocatedCapacity(v float64) {
	o.AllocatedCapacity = &v
}

// GetMaximumVolumeSize returns the MaximumVolumeSize field value if set, zero value otherwise.
func (o *StorageClassNodeStats) GetMaximumVolumeSize() float64 {
	if o == nil || o.MaximumVolumeSize == nil {
		var ret float64
		return ret
	}
	return *o.MaximumVolumeSize
}

// GetMaximumVolumeSizeOk returns a tuple with the MaximumVolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetMaximumVolumeSizeOk() (*float64, bool) {
	if o == nil || o.MaximumVolumeSize == nil {
		return nil, false
	}
	return o.MaximumVolumeSize, true
}

// HasMaximumVolumeSize returns a boolean if a field has been set.
func (o *StorageClassNodeStats) HasMaximumVolumeSize() bool {
	return o != nil && o.MaximumVolumeSize != nil
}

// SetMaximumVolumeSize gets a reference to the given float64 and assigns it to the MaximumVolumeSize field.
func (o *StorageClassNodeStats) SetMaximumVolumeSize(v float64) {
	o.MaximumVolumeSize = &v
}

// GetNodeName returns the NodeName field value.
func (o *StorageClassNodeStats) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value.
func (o *StorageClassNodeStats) SetNodeName(v string) {
	o.NodeName = v
}

// GetNodeStatus returns the NodeStatus field value.
func (o *StorageClassNodeStats) GetNodeStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NodeStatus
}

// GetNodeStatusOk returns a tuple with the NodeStatus field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetNodeStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeStatus, true
}

// SetNodeStatus sets field value.
func (o *StorageClassNodeStats) SetNodeStatus(v string) {
	o.NodeStatus = v
}

// GetRequests returns the Requests field value.
func (o *StorageClassNodeStats) GetRequests() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetRequestsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *StorageClassNodeStats) SetRequests(v float64) {
	o.Requests = v
}

// GetUsage returns the Usage field value.
func (o *StorageClassNodeStats) GetUsage() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetUsageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value.
func (o *StorageClassNodeStats) SetUsage(v float64) {
	o.Usage = v
}

// GetCapacity returns the Capacity field value.
func (o *StorageClassNodeStats) GetCapacity() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetCapacityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value.
func (o *StorageClassNodeStats) SetCapacity(v float64) {
	o.Capacity = v
}

// GetPvcCount returns the PvcCount field value.
func (o *StorageClassNodeStats) GetPvcCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.PvcCount
}

// GetPvcCountOk returns a tuple with the PvcCount field value
// and a boolean to check if the value has been set.
func (o *StorageClassNodeStats) GetPvcCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PvcCount, true
}

// SetPvcCount sets field value.
func (o *StorageClassNodeStats) SetPvcCount(v int32) {
	o.PvcCount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageClassNodeStats) MarshalJSON() ([]byte, error) {
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
	toSerialize["nodeName"] = o.NodeName
	toSerialize["nodeStatus"] = o.NodeStatus
	toSerialize["requests"] = o.Requests
	toSerialize["usage"] = o.Usage
	toSerialize["capacity"] = o.Capacity
	toSerialize["pvcCount"] = o.PvcCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageClassNodeStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllocationResourceId *string  `json:"allocationResourceID,omitempty"`
		PhysicalResourceId   *string  `json:"physicalResourceID,omitempty"`
		PhysicalCapacity     *float64 `json:"physicalCapacity,omitempty"`
		PhysicalUsage        *float64 `json:"physicalUsage,omitempty"`
		AllocatableCapacity  *float64 `json:"allocatableCapacity,omitempty"`
		AvailableCapacity    *float64 `json:"availableCapacity,omitempty"`
		AllocatedCapacity    *float64 `json:"allocatedCapacity,omitempty"`
		MaximumVolumeSize    *float64 `json:"maximumVolumeSize,omitempty"`
		NodeName             *string  `json:"nodeName"`
		NodeStatus           *string  `json:"nodeStatus"`
		Requests             *float64 `json:"requests"`
		Usage                *float64 `json:"usage"`
		Capacity             *float64 `json:"capacity"`
		PvcCount             *int32   `json:"pvcCount"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.NodeName == nil {
		return fmt.Errorf("required field nodeName missing")
	}
	if all.NodeStatus == nil {
		return fmt.Errorf("required field nodeStatus missing")
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.Usage == nil {
		return fmt.Errorf("required field usage missing")
	}
	if all.Capacity == nil {
		return fmt.Errorf("required field capacity missing")
	}
	if all.PvcCount == nil {
		return fmt.Errorf("required field pvcCount missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"allocationResourceID", "physicalResourceID", "physicalCapacity", "physicalUsage", "allocatableCapacity", "availableCapacity", "allocatedCapacity", "maximumVolumeSize", "nodeName", "nodeStatus", "requests", "usage", "capacity", "pvcCount"})
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
	o.NodeName = *all.NodeName
	o.NodeStatus = *all.NodeStatus
	o.Requests = *all.Requests
	o.Usage = *all.Usage
	o.Capacity = *all.Capacity
	o.PvcCount = *all.PvcCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
