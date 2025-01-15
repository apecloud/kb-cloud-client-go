// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type Tenant struct {
	Id               *string   `json:"id,omitempty"`
	Name             *string   `json:"name,omitempty"`
	Mode             *string   `json:"mode,omitempty"`
	CreateAt         *string   `json:"create_at,omitempty"`
	Role             *string   `json:"role,omitempty"`
	PrimaryZone      *string   `json:"primary_zone,omitempty"`
	Status           *string   `json:"status,omitempty"`
	PrimaryZoneProxy *string   `json:"primary_zone_proxy,omitempty"`
	ChartSet         *string   `json:"chart_set,omitempty"`
	Resource         *string   `json:"resource,omitempty"`
	Cpu              *CPU      `json:"cpu,omitempty"`
	Memory           *Memory   `json:"memory,omitempty"`
	LogDisk          *LogDisk  `json:"log_disk,omitempty"`
	DataDisk         *DataDisk `json:"data_disk,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTenant instantiates a new Tenant object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTenant() *Tenant {
	this := Tenant{}
	return &this
}

// NewTenantWithDefaults instantiates a new Tenant object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTenantWithDefaults() *Tenant {
	this := Tenant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Tenant) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Tenant) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Tenant) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Tenant) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Tenant) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Tenant) SetName(v string) {
	o.Name = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *Tenant) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *Tenant) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *Tenant) SetMode(v string) {
	o.Mode = &v
}

// GetCreateAt returns the CreateAt field value if set, zero value otherwise.
func (o *Tenant) GetCreateAt() string {
	if o == nil || o.CreateAt == nil {
		var ret string
		return ret
	}
	return *o.CreateAt
}

// GetCreateAtOk returns a tuple with the CreateAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCreateAtOk() (*string, bool) {
	if o == nil || o.CreateAt == nil {
		return nil, false
	}
	return o.CreateAt, true
}

// HasCreateAt returns a boolean if a field has been set.
func (o *Tenant) HasCreateAt() bool {
	return o != nil && o.CreateAt != nil
}

// SetCreateAt gets a reference to the given string and assigns it to the CreateAt field.
func (o *Tenant) SetCreateAt(v string) {
	o.CreateAt = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Tenant) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *Tenant) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *Tenant) SetRole(v string) {
	o.Role = &v
}

// GetPrimaryZone returns the PrimaryZone field value if set, zero value otherwise.
func (o *Tenant) GetPrimaryZone() string {
	if o == nil || o.PrimaryZone == nil {
		var ret string
		return ret
	}
	return *o.PrimaryZone
}

// GetPrimaryZoneOk returns a tuple with the PrimaryZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetPrimaryZoneOk() (*string, bool) {
	if o == nil || o.PrimaryZone == nil {
		return nil, false
	}
	return o.PrimaryZone, true
}

// HasPrimaryZone returns a boolean if a field has been set.
func (o *Tenant) HasPrimaryZone() bool {
	return o != nil && o.PrimaryZone != nil
}

// SetPrimaryZone gets a reference to the given string and assigns it to the PrimaryZone field.
func (o *Tenant) SetPrimaryZone(v string) {
	o.PrimaryZone = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Tenant) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Tenant) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Tenant) SetStatus(v string) {
	o.Status = &v
}

// GetPrimaryZoneProxy returns the PrimaryZoneProxy field value if set, zero value otherwise.
func (o *Tenant) GetPrimaryZoneProxy() string {
	if o == nil || o.PrimaryZoneProxy == nil {
		var ret string
		return ret
	}
	return *o.PrimaryZoneProxy
}

// GetPrimaryZoneProxyOk returns a tuple with the PrimaryZoneProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetPrimaryZoneProxyOk() (*string, bool) {
	if o == nil || o.PrimaryZoneProxy == nil {
		return nil, false
	}
	return o.PrimaryZoneProxy, true
}

// HasPrimaryZoneProxy returns a boolean if a field has been set.
func (o *Tenant) HasPrimaryZoneProxy() bool {
	return o != nil && o.PrimaryZoneProxy != nil
}

// SetPrimaryZoneProxy gets a reference to the given string and assigns it to the PrimaryZoneProxy field.
func (o *Tenant) SetPrimaryZoneProxy(v string) {
	o.PrimaryZoneProxy = &v
}

// GetChartSet returns the ChartSet field value if set, zero value otherwise.
func (o *Tenant) GetChartSet() string {
	if o == nil || o.ChartSet == nil {
		var ret string
		return ret
	}
	return *o.ChartSet
}

// GetChartSetOk returns a tuple with the ChartSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetChartSetOk() (*string, bool) {
	if o == nil || o.ChartSet == nil {
		return nil, false
	}
	return o.ChartSet, true
}

// HasChartSet returns a boolean if a field has been set.
func (o *Tenant) HasChartSet() bool {
	return o != nil && o.ChartSet != nil
}

// SetChartSet gets a reference to the given string and assigns it to the ChartSet field.
func (o *Tenant) SetChartSet(v string) {
	o.ChartSet = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *Tenant) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *Tenant) HasResource() bool {
	return o != nil && o.Resource != nil
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *Tenant) SetResource(v string) {
	o.Resource = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Tenant) GetCpu() CPU {
	if o == nil || o.Cpu == nil {
		var ret CPU
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCpuOk() (*CPU, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Tenant) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given CPU and assigns it to the Cpu field.
func (o *Tenant) SetCpu(v CPU) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Tenant) GetMemory() Memory {
	if o == nil || o.Memory == nil {
		var ret Memory
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetMemoryOk() (*Memory, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Tenant) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given Memory and assigns it to the Memory field.
func (o *Tenant) SetMemory(v Memory) {
	o.Memory = &v
}

// GetLogDisk returns the LogDisk field value if set, zero value otherwise.
func (o *Tenant) GetLogDisk() LogDisk {
	if o == nil || o.LogDisk == nil {
		var ret LogDisk
		return ret
	}
	return *o.LogDisk
}

// GetLogDiskOk returns a tuple with the LogDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetLogDiskOk() (*LogDisk, bool) {
	if o == nil || o.LogDisk == nil {
		return nil, false
	}
	return o.LogDisk, true
}

// HasLogDisk returns a boolean if a field has been set.
func (o *Tenant) HasLogDisk() bool {
	return o != nil && o.LogDisk != nil
}

// SetLogDisk gets a reference to the given LogDisk and assigns it to the LogDisk field.
func (o *Tenant) SetLogDisk(v LogDisk) {
	o.LogDisk = &v
}

// GetDataDisk returns the DataDisk field value if set, zero value otherwise.
func (o *Tenant) GetDataDisk() DataDisk {
	if o == nil || o.DataDisk == nil {
		var ret DataDisk
		return ret
	}
	return *o.DataDisk
}

// GetDataDiskOk returns a tuple with the DataDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetDataDiskOk() (*DataDisk, bool) {
	if o == nil || o.DataDisk == nil {
		return nil, false
	}
	return o.DataDisk, true
}

// HasDataDisk returns a boolean if a field has been set.
func (o *Tenant) HasDataDisk() bool {
	return o != nil && o.DataDisk != nil
}

// SetDataDisk gets a reference to the given DataDisk and assigns it to the DataDisk field.
func (o *Tenant) SetDataDisk(v DataDisk) {
	o.DataDisk = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Tenant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.CreateAt != nil {
		toSerialize["create_at"] = o.CreateAt
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.PrimaryZone != nil {
		toSerialize["primary_zone"] = o.PrimaryZone
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.PrimaryZoneProxy != nil {
		toSerialize["primary_zone_proxy"] = o.PrimaryZoneProxy
	}
	if o.ChartSet != nil {
		toSerialize["chart_set"] = o.ChartSet
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.LogDisk != nil {
		toSerialize["log_disk"] = o.LogDisk
	}
	if o.DataDisk != nil {
		toSerialize["data_disk"] = o.DataDisk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Tenant) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id               *string   `json:"id,omitempty"`
		Name             *string   `json:"name,omitempty"`
		Mode             *string   `json:"mode,omitempty"`
		CreateAt         *string   `json:"create_at,omitempty"`
		Role             *string   `json:"role,omitempty"`
		PrimaryZone      *string   `json:"primary_zone,omitempty"`
		Status           *string   `json:"status,omitempty"`
		PrimaryZoneProxy *string   `json:"primary_zone_proxy,omitempty"`
		ChartSet         *string   `json:"chart_set,omitempty"`
		Resource         *string   `json:"resource,omitempty"`
		Cpu              *CPU      `json:"cpu,omitempty"`
		Memory           *Memory   `json:"memory,omitempty"`
		LogDisk          *LogDisk  `json:"log_disk,omitempty"`
		DataDisk         *DataDisk `json:"data_disk,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "mode", "create_at", "role", "primary_zone", "status", "primary_zone_proxy", "chart_set", "resource", "cpu", "memory", "log_disk", "data_disk"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = all.Name
	o.Mode = all.Mode
	o.CreateAt = all.CreateAt
	o.Role = all.Role
	o.PrimaryZone = all.PrimaryZone
	o.Status = all.Status
	o.PrimaryZoneProxy = all.PrimaryZoneProxy
	o.ChartSet = all.ChartSet
	o.Resource = all.Resource
	if all.Cpu != nil && all.Cpu.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cpu = all.Cpu
	if all.Memory != nil && all.Memory.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Memory = all.Memory
	if all.LogDisk != nil && all.LogDisk.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogDisk = all.LogDisk
	if all.DataDisk != nil && all.DataDisk.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DataDisk = all.DataDisk

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
