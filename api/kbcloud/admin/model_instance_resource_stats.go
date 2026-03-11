// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceResourceStats InstanceResourceStats holds the requests, limits, and available stats for an instance.
type InstanceResourceStats struct {
	// ResourceStats holds the requests, limits, and available stats for a resource.
	CpuStats ResourceStats `json:"cpuStats"`
	// ResourceStats holds the requests, limits, and available stats for a resource.
	MemoryStats ResourceStats `json:"memoryStats"`
	// ResourceStats holds the requests, limits, and available stats for a resource.
	EphemeralStorageStats *ResourceStats `json:"ephemeralStorageStats,omitempty"`
	// Name of the instance.
	Name string `json:"name"`
	// Type of the instance, choose cluster or system
	Type *InstanceResourceStatsType `json:"type,omitempty"`
	// Engine type of the instance (e.g., postgresql, mysql).
	Engine *string `json:"engine,omitempty"`
	// Engine version of the instance (e.g., postgresql-18-1.0.2).
	EngineVersion *string `json:"engineVersion,omitempty"`
	// Role of the instance (e.g., primary, secondary).
	Role *string `json:"role,omitempty"`
	// Name of the cluster this instance belongs to.
	ClusterName *string `json:"clusterName,omitempty"`
	// Database status (e.g., running, stopped).
	DbStatus *string `json:"dbStatus,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceResourceStats instantiates a new InstanceResourceStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceResourceStats(cpuStats ResourceStats, memoryStats ResourceStats, name string) *InstanceResourceStats {
	this := InstanceResourceStats{}
	this.CpuStats = cpuStats
	this.MemoryStats = memoryStats
	this.Name = name
	return &this
}

// NewInstanceResourceStatsWithDefaults instantiates a new InstanceResourceStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceResourceStatsWithDefaults() *InstanceResourceStats {
	this := InstanceResourceStats{}
	return &this
}

// GetCpuStats returns the CpuStats field value.
func (o *InstanceResourceStats) GetCpuStats() ResourceStats {
	if o == nil {
		var ret ResourceStats
		return ret
	}
	return o.CpuStats
}

// GetCpuStatsOk returns a tuple with the CpuStats field value
// and a boolean to check if the value has been set.
func (o *InstanceResourceStats) GetCpuStatsOk() (*ResourceStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuStats, true
}

// SetCpuStats sets field value.
func (o *InstanceResourceStats) SetCpuStats(v ResourceStats) {
	o.CpuStats = v
}

// GetMemoryStats returns the MemoryStats field value.
func (o *InstanceResourceStats) GetMemoryStats() ResourceStats {
	if o == nil {
		var ret ResourceStats
		return ret
	}
	return o.MemoryStats
}

// GetMemoryStatsOk returns a tuple with the MemoryStats field value
// and a boolean to check if the value has been set.
func (o *InstanceResourceStats) GetMemoryStatsOk() (*ResourceStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryStats, true
}

// SetMemoryStats sets field value.
func (o *InstanceResourceStats) SetMemoryStats(v ResourceStats) {
	o.MemoryStats = v
}

// GetEphemeralStorageStats returns the EphemeralStorageStats field value if set, zero value otherwise.
func (o *InstanceResourceStats) GetEphemeralStorageStats() ResourceStats {
	if o == nil || o.EphemeralStorageStats == nil {
		var ret ResourceStats
		return ret
	}
	return *o.EphemeralStorageStats
}

// GetEphemeralStorageStatsOk returns a tuple with the EphemeralStorageStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResourceStats) GetEphemeralStorageStatsOk() (*ResourceStats, bool) {
	if o == nil || o.EphemeralStorageStats == nil {
		return nil, false
	}
	return o.EphemeralStorageStats, true
}

// HasEphemeralStorageStats returns a boolean if a field has been set.
func (o *InstanceResourceStats) HasEphemeralStorageStats() bool {
	return o != nil && o.EphemeralStorageStats != nil
}

// SetEphemeralStorageStats gets a reference to the given ResourceStats and assigns it to the EphemeralStorageStats field.
func (o *InstanceResourceStats) SetEphemeralStorageStats(v ResourceStats) {
	o.EphemeralStorageStats = &v
}

// GetName returns the Name field value.
func (o *InstanceResourceStats) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceResourceStats) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InstanceResourceStats) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InstanceResourceStats) GetType() InstanceResourceStatsType {
	if o == nil || o.Type == nil {
		var ret InstanceResourceStatsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResourceStats) GetTypeOk() (*InstanceResourceStatsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InstanceResourceStats) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given InstanceResourceStatsType and assigns it to the Type field.
func (o *InstanceResourceStats) SetType(v InstanceResourceStatsType) {
	o.Type = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *InstanceResourceStats) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResourceStats) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *InstanceResourceStats) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *InstanceResourceStats) SetEngine(v string) {
	o.Engine = &v
}

// GetEngineVersion returns the EngineVersion field value if set, zero value otherwise.
func (o *InstanceResourceStats) GetEngineVersion() string {
	if o == nil || o.EngineVersion == nil {
		var ret string
		return ret
	}
	return *o.EngineVersion
}

// GetEngineVersionOk returns a tuple with the EngineVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResourceStats) GetEngineVersionOk() (*string, bool) {
	if o == nil || o.EngineVersion == nil {
		return nil, false
	}
	return o.EngineVersion, true
}

// HasEngineVersion returns a boolean if a field has been set.
func (o *InstanceResourceStats) HasEngineVersion() bool {
	return o != nil && o.EngineVersion != nil
}

// SetEngineVersion gets a reference to the given string and assigns it to the EngineVersion field.
func (o *InstanceResourceStats) SetEngineVersion(v string) {
	o.EngineVersion = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *InstanceResourceStats) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResourceStats) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *InstanceResourceStats) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *InstanceResourceStats) SetRole(v string) {
	o.Role = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *InstanceResourceStats) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResourceStats) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *InstanceResourceStats) HasClusterName() bool {
	return o != nil && o.ClusterName != nil
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *InstanceResourceStats) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetDbStatus returns the DbStatus field value if set, zero value otherwise.
func (o *InstanceResourceStats) GetDbStatus() string {
	if o == nil || o.DbStatus == nil {
		var ret string
		return ret
	}
	return *o.DbStatus
}

// GetDbStatusOk returns a tuple with the DbStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceResourceStats) GetDbStatusOk() (*string, bool) {
	if o == nil || o.DbStatus == nil {
		return nil, false
	}
	return o.DbStatus, true
}

// HasDbStatus returns a boolean if a field has been set.
func (o *InstanceResourceStats) HasDbStatus() bool {
	return o != nil && o.DbStatus != nil
}

// SetDbStatus gets a reference to the given string and assigns it to the DbStatus field.
func (o *InstanceResourceStats) SetDbStatus(v string) {
	o.DbStatus = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceResourceStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cpuStats"] = o.CpuStats
	toSerialize["memoryStats"] = o.MemoryStats
	if o.EphemeralStorageStats != nil {
		toSerialize["ephemeralStorageStats"] = o.EphemeralStorageStats
	}
	toSerialize["name"] = o.Name
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.EngineVersion != nil {
		toSerialize["engineVersion"] = o.EngineVersion
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.DbStatus != nil {
		toSerialize["dbStatus"] = o.DbStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceResourceStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CpuStats              *ResourceStats             `json:"cpuStats"`
		MemoryStats           *ResourceStats             `json:"memoryStats"`
		EphemeralStorageStats *ResourceStats             `json:"ephemeralStorageStats,omitempty"`
		Name                  *string                    `json:"name"`
		Type                  *InstanceResourceStatsType `json:"type,omitempty"`
		Engine                *string                    `json:"engine,omitempty"`
		EngineVersion         *string                    `json:"engineVersion,omitempty"`
		Role                  *string                    `json:"role,omitempty"`
		ClusterName           *string                    `json:"clusterName,omitempty"`
		DbStatus              *string                    `json:"dbStatus,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CpuStats == nil {
		return fmt.Errorf("required field cpuStats missing")
	}
	if all.MemoryStats == nil {
		return fmt.Errorf("required field memoryStats missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cpuStats", "memoryStats", "ephemeralStorageStats", "name", "type", "engine", "engineVersion", "role", "clusterName", "dbStatus"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CpuStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CpuStats = *all.CpuStats
	if all.MemoryStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MemoryStats = *all.MemoryStats
	if all.EphemeralStorageStats != nil && all.EphemeralStorageStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.EphemeralStorageStats = all.EphemeralStorageStats
	o.Name = *all.Name
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Engine = all.Engine
	o.EngineVersion = all.EngineVersion
	o.Role = all.Role
	o.ClusterName = all.ClusterName
	o.DbStatus = all.DbStatus

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
