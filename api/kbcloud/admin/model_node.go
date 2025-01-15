// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Node node info
type Node struct {
	// CPU cores of the node
	Cpu *int64 `json:"cpu,omitempty"`
	// ResourceStats holds the requests, limits, and available stats for a resource.
	CpuStats *ResourceStats `json:"cpuStats,omitempty"`
	// CreatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system. Read-only. Null for lists
	CreatedAt time.Time `json:"createdAt"`
	// Host name of the node
	HostName string `json:"hostName"`
	// Instance type of the node
	InstanceType *string `json:"instanceType,omitempty"`
	// IP address of the node
	Ip string `json:"ip"`
	// Memory of the node, unit is GiB
	Memory *int64 `json:"memory,omitempty"`
	// ResourceStats holds the requests, limits, and available stats for a resource.
	MemoryStats *ResourceStats `json:"memoryStats,omitempty"`
	// Zone of the node
	Zone *string `json:"zone,omitempty"`
	// Region of the node
	Region *string `json:"region,omitempty"`
	// whether node is unschedulable
	Unschedulable *bool `json:"unschedulable,omitempty"`
	// Node Group of the node
	NodeGroup *string `json:"nodeGroup,omitempty"`
	// node is in control plane
	ControlPlane *bool `json:"controlPlane,omitempty"`
	// node is in data plane
	DataPlane *bool `json:"dataPlane,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNode instantiates a new Node object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNode(createdAt time.Time, hostName string, ip string) *Node {
	this := Node{}
	this.CreatedAt = createdAt
	this.HostName = hostName
	this.Ip = ip
	return &this
}

// NewNodeWithDefaults instantiates a new Node object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodeWithDefaults() *Node {
	this := Node{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Node) GetCpu() int64 {
	if o == nil || o.Cpu == nil {
		var ret int64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetCpuOk() (*int64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Node) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *Node) SetCpu(v int64) {
	o.Cpu = &v
}

// GetCpuStats returns the CpuStats field value if set, zero value otherwise.
func (o *Node) GetCpuStats() ResourceStats {
	if o == nil || o.CpuStats == nil {
		var ret ResourceStats
		return ret
	}
	return *o.CpuStats
}

// GetCpuStatsOk returns a tuple with the CpuStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetCpuStatsOk() (*ResourceStats, bool) {
	if o == nil || o.CpuStats == nil {
		return nil, false
	}
	return o.CpuStats, true
}

// HasCpuStats returns a boolean if a field has been set.
func (o *Node) HasCpuStats() bool {
	return o != nil && o.CpuStats != nil
}

// SetCpuStats gets a reference to the given ResourceStats and assigns it to the CpuStats field.
func (o *Node) SetCpuStats(v ResourceStats) {
	o.CpuStats = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Node) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Node) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Node) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetHostName returns the HostName field value.
func (o *Node) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *Node) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value.
func (o *Node) SetHostName(v string) {
	o.HostName = v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *Node) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *Node) HasInstanceType() bool {
	return o != nil && o.InstanceType != nil
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *Node) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetIp returns the Ip field value.
func (o *Node) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *Node) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value.
func (o *Node) SetIp(v string) {
	o.Ip = v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Node) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Node) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *Node) SetMemory(v int64) {
	o.Memory = &v
}

// GetMemoryStats returns the MemoryStats field value if set, zero value otherwise.
func (o *Node) GetMemoryStats() ResourceStats {
	if o == nil || o.MemoryStats == nil {
		var ret ResourceStats
		return ret
	}
	return *o.MemoryStats
}

// GetMemoryStatsOk returns a tuple with the MemoryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetMemoryStatsOk() (*ResourceStats, bool) {
	if o == nil || o.MemoryStats == nil {
		return nil, false
	}
	return o.MemoryStats, true
}

// HasMemoryStats returns a boolean if a field has been set.
func (o *Node) HasMemoryStats() bool {
	return o != nil && o.MemoryStats != nil
}

// SetMemoryStats gets a reference to the given ResourceStats and assigns it to the MemoryStats field.
func (o *Node) SetMemoryStats(v ResourceStats) {
	o.MemoryStats = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *Node) GetZone() string {
	if o == nil || o.Zone == nil {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetZoneOk() (*string, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *Node) HasZone() bool {
	return o != nil && o.Zone != nil
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *Node) SetZone(v string) {
	o.Zone = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Node) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Node) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Node) SetRegion(v string) {
	o.Region = &v
}

// GetUnschedulable returns the Unschedulable field value if set, zero value otherwise.
func (o *Node) GetUnschedulable() bool {
	if o == nil || o.Unschedulable == nil {
		var ret bool
		return ret
	}
	return *o.Unschedulable
}

// GetUnschedulableOk returns a tuple with the Unschedulable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetUnschedulableOk() (*bool, bool) {
	if o == nil || o.Unschedulable == nil {
		return nil, false
	}
	return o.Unschedulable, true
}

// HasUnschedulable returns a boolean if a field has been set.
func (o *Node) HasUnschedulable() bool {
	return o != nil && o.Unschedulable != nil
}

// SetUnschedulable gets a reference to the given bool and assigns it to the Unschedulable field.
func (o *Node) SetUnschedulable(v bool) {
	o.Unschedulable = &v
}

// GetNodeGroup returns the NodeGroup field value if set, zero value otherwise.
func (o *Node) GetNodeGroup() string {
	if o == nil || o.NodeGroup == nil {
		var ret string
		return ret
	}
	return *o.NodeGroup
}

// GetNodeGroupOk returns a tuple with the NodeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetNodeGroupOk() (*string, bool) {
	if o == nil || o.NodeGroup == nil {
		return nil, false
	}
	return o.NodeGroup, true
}

// HasNodeGroup returns a boolean if a field has been set.
func (o *Node) HasNodeGroup() bool {
	return o != nil && o.NodeGroup != nil
}

// SetNodeGroup gets a reference to the given string and assigns it to the NodeGroup field.
func (o *Node) SetNodeGroup(v string) {
	o.NodeGroup = &v
}

// GetControlPlane returns the ControlPlane field value if set, zero value otherwise.
func (o *Node) GetControlPlane() bool {
	if o == nil || o.ControlPlane == nil {
		var ret bool
		return ret
	}
	return *o.ControlPlane
}

// GetControlPlaneOk returns a tuple with the ControlPlane field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetControlPlaneOk() (*bool, bool) {
	if o == nil || o.ControlPlane == nil {
		return nil, false
	}
	return o.ControlPlane, true
}

// HasControlPlane returns a boolean if a field has been set.
func (o *Node) HasControlPlane() bool {
	return o != nil && o.ControlPlane != nil
}

// SetControlPlane gets a reference to the given bool and assigns it to the ControlPlane field.
func (o *Node) SetControlPlane(v bool) {
	o.ControlPlane = &v
}

// GetDataPlane returns the DataPlane field value if set, zero value otherwise.
func (o *Node) GetDataPlane() bool {
	if o == nil || o.DataPlane == nil {
		var ret bool
		return ret
	}
	return *o.DataPlane
}

// GetDataPlaneOk returns a tuple with the DataPlane field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Node) GetDataPlaneOk() (*bool, bool) {
	if o == nil || o.DataPlane == nil {
		return nil, false
	}
	return o.DataPlane, true
}

// HasDataPlane returns a boolean if a field has been set.
func (o *Node) HasDataPlane() bool {
	return o != nil && o.DataPlane != nil
}

// SetDataPlane gets a reference to the given bool and assigns it to the DataPlane field.
func (o *Node) SetDataPlane(v bool) {
	o.DataPlane = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Node) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.CpuStats != nil {
		toSerialize["cpuStats"] = o.CpuStats
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["hostName"] = o.HostName
	if o.InstanceType != nil {
		toSerialize["instanceType"] = o.InstanceType
	}
	toSerialize["ip"] = o.Ip
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MemoryStats != nil {
		toSerialize["memoryStats"] = o.MemoryStats
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Unschedulable != nil {
		toSerialize["unschedulable"] = o.Unschedulable
	}
	if o.NodeGroup != nil {
		toSerialize["nodeGroup"] = o.NodeGroup
	}
	if o.ControlPlane != nil {
		toSerialize["controlPlane"] = o.ControlPlane
	}
	if o.DataPlane != nil {
		toSerialize["dataPlane"] = o.DataPlane
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Node) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cpu           *int64         `json:"cpu,omitempty"`
		CpuStats      *ResourceStats `json:"cpuStats,omitempty"`
		CreatedAt     *time.Time     `json:"createdAt"`
		HostName      *string        `json:"hostName"`
		InstanceType  *string        `json:"instanceType,omitempty"`
		Ip            *string        `json:"ip"`
		Memory        *int64         `json:"memory,omitempty"`
		MemoryStats   *ResourceStats `json:"memoryStats,omitempty"`
		Zone          *string        `json:"zone,omitempty"`
		Region        *string        `json:"region,omitempty"`
		Unschedulable *bool          `json:"unschedulable,omitempty"`
		NodeGroup     *string        `json:"nodeGroup,omitempty"`
		ControlPlane  *bool          `json:"controlPlane,omitempty"`
		DataPlane     *bool          `json:"dataPlane,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.HostName == nil {
		return fmt.Errorf("required field hostName missing")
	}
	if all.Ip == nil {
		return fmt.Errorf("required field ip missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cpu", "cpuStats", "createdAt", "hostName", "instanceType", "ip", "memory", "memoryStats", "zone", "region", "unschedulable", "nodeGroup", "controlPlane", "dataPlane"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Cpu = all.Cpu
	if all.CpuStats != nil && all.CpuStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CpuStats = all.CpuStats
	o.CreatedAt = *all.CreatedAt
	o.HostName = *all.HostName
	o.InstanceType = all.InstanceType
	o.Ip = *all.Ip
	o.Memory = all.Memory
	if all.MemoryStats != nil && all.MemoryStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MemoryStats = all.MemoryStats
	o.Zone = all.Zone
	o.Region = all.Region
	o.Unschedulable = all.Unschedulable
	o.NodeGroup = all.NodeGroup
	o.ControlPlane = all.ControlPlane
	o.DataPlane = all.DataPlane

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
