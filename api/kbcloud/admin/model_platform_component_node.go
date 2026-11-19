// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// PlatformComponentNode A node hosting platform components. CPU and memory report capacity, allocatable capacity and actual usage in cores and GiB respectively. Filesystem statistics report root filesystem capacity and usage in GiB. Unavailable samples are -1.
type PlatformComponentNode struct {
	Name *string `json:"name,omitempty"`
	// Internal node IP, falling back to external IP
	Ip *string `json:"ip,omitempty"`
	// Kubernetes node role, control-plane or worker
	Role *string `json:"role,omitempty"`
	// Ready when the NodeReady condition is True, otherwise NotReady
	Status *string `json:"status,omitempty"`
	// Names of platform components with Pods assigned to this node
	Components []string `json:"components,omitempty"`
	// ResourceStats holds the requests, limits, and available stats for a resource.
	CpuStats *ResourceStats `json:"cpuStats,omitempty"`
	// ResourceStats holds the requests, limits, and available stats for a resource.
	MemoryStats *ResourceStats `json:"memoryStats,omitempty"`
	// ResourceStats holds the requests, limits, and available stats for a resource.
	FilesystemStats *ResourceStats `json:"filesystemStats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlatformComponentNode instantiates a new PlatformComponentNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlatformComponentNode() *PlatformComponentNode {
	this := PlatformComponentNode{}
	return &this
}

// NewPlatformComponentNodeWithDefaults instantiates a new PlatformComponentNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlatformComponentNodeWithDefaults() *PlatformComponentNode {
	this := PlatformComponentNode{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlatformComponentNode) SetName(v string) {
	o.Name = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasIp() bool {
	return o != nil && o.Ip != nil
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *PlatformComponentNode) SetIp(v string) {
	o.Ip = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *PlatformComponentNode) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PlatformComponentNode) SetStatus(v string) {
	o.Status = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetComponents() []string {
	if o == nil || o.Components == nil {
		var ret []string
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetComponentsOk() (*[]string, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []string and assigns it to the Components field.
func (o *PlatformComponentNode) SetComponents(v []string) {
	o.Components = v
}

// GetCpuStats returns the CpuStats field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetCpuStats() ResourceStats {
	if o == nil || o.CpuStats == nil {
		var ret ResourceStats
		return ret
	}
	return *o.CpuStats
}

// GetCpuStatsOk returns a tuple with the CpuStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetCpuStatsOk() (*ResourceStats, bool) {
	if o == nil || o.CpuStats == nil {
		return nil, false
	}
	return o.CpuStats, true
}

// HasCpuStats returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasCpuStats() bool {
	return o != nil && o.CpuStats != nil
}

// SetCpuStats gets a reference to the given ResourceStats and assigns it to the CpuStats field.
func (o *PlatformComponentNode) SetCpuStats(v ResourceStats) {
	o.CpuStats = &v
}

// GetMemoryStats returns the MemoryStats field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetMemoryStats() ResourceStats {
	if o == nil || o.MemoryStats == nil {
		var ret ResourceStats
		return ret
	}
	return *o.MemoryStats
}

// GetMemoryStatsOk returns a tuple with the MemoryStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetMemoryStatsOk() (*ResourceStats, bool) {
	if o == nil || o.MemoryStats == nil {
		return nil, false
	}
	return o.MemoryStats, true
}

// HasMemoryStats returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasMemoryStats() bool {
	return o != nil && o.MemoryStats != nil
}

// SetMemoryStats gets a reference to the given ResourceStats and assigns it to the MemoryStats field.
func (o *PlatformComponentNode) SetMemoryStats(v ResourceStats) {
	o.MemoryStats = &v
}

// GetFilesystemStats returns the FilesystemStats field value if set, zero value otherwise.
func (o *PlatformComponentNode) GetFilesystemStats() ResourceStats {
	if o == nil || o.FilesystemStats == nil {
		var ret ResourceStats
		return ret
	}
	return *o.FilesystemStats
}

// GetFilesystemStatsOk returns a tuple with the FilesystemStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformComponentNode) GetFilesystemStatsOk() (*ResourceStats, bool) {
	if o == nil || o.FilesystemStats == nil {
		return nil, false
	}
	return o.FilesystemStats, true
}

// HasFilesystemStats returns a boolean if a field has been set.
func (o *PlatformComponentNode) HasFilesystemStats() bool {
	return o != nil && o.FilesystemStats != nil
}

// SetFilesystemStats gets a reference to the given ResourceStats and assigns it to the FilesystemStats field.
func (o *PlatformComponentNode) SetFilesystemStats(v ResourceStats) {
	o.FilesystemStats = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlatformComponentNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.CpuStats != nil {
		toSerialize["cpuStats"] = o.CpuStats
	}
	if o.MemoryStats != nil {
		toSerialize["memoryStats"] = o.MemoryStats
	}
	if o.FilesystemStats != nil {
		toSerialize["filesystemStats"] = o.FilesystemStats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlatformComponentNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name            *string        `json:"name,omitempty"`
		Ip              *string        `json:"ip,omitempty"`
		Role            *string        `json:"role,omitempty"`
		Status          *string        `json:"status,omitempty"`
		Components      []string       `json:"components,omitempty"`
		CpuStats        *ResourceStats `json:"cpuStats,omitempty"`
		MemoryStats     *ResourceStats `json:"memoryStats,omitempty"`
		FilesystemStats *ResourceStats `json:"filesystemStats,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "ip", "role", "status", "components", "cpuStats", "memoryStats", "filesystemStats"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Ip = all.Ip
	o.Role = all.Role
	o.Status = all.Status
	o.Components = all.Components
	if all.CpuStats != nil && all.CpuStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CpuStats = all.CpuStats
	if all.MemoryStats != nil && all.MemoryStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MemoryStats = all.MemoryStats
	if all.FilesystemStats != nil && all.FilesystemStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FilesystemStats = all.FilesystemStats

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
