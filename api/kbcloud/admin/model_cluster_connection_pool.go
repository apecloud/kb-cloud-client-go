// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterConnectionPool PgBouncer connection pool state. PostgreSQL keeps its existing topology and
// always models PgBouncer as an independent component when the installed
// Addon supports it. A desired replica count greater than zero means enabled;
// available is reported separately from the reconciled workload and endpoint.
type ClusterConnectionPool struct {
	// Whether the desired PgBouncer replica count is greater than zero.
	Enabled bool `json:"enabled"`
	// Whether PgBouncer is enabled and its Cluster and InstanceSet generations,
	// ready/available/updated replicas, and ready endpoint addresses have all
	// converged to the desired replica count.
	//
	Available common.NullableBool `json:"available,omitempty"`
	// Connection pool implementation managed by KBE.
	Provider *ClusterConnectionPoolProvider `json:"provider,omitempty"`
	// Connection reuse mode. The MVP supports session pooling only.
	PoolMode *ClusterConnectionPoolMode `json:"poolMode,omitempty"`
	// Desired PgBouncer replicas. Zero disables the connection pool.
	Replicas *int32 `json:"replicas,omitempty"`
	// Resource class for the connection pool component.
	ClassCode *string `json:"classCode,omitempty"`
	// Connection pool CPU cores when classCode is not used.
	Cpu *float64 `json:"cpu,omitempty"`
	// Connection pool memory in GiB when classCode is not used.
	Memory *float64 `json:"memory,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterConnectionPool instantiates a new ClusterConnectionPool object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterConnectionPool(enabled bool) *ClusterConnectionPool {
	this := ClusterConnectionPool{}
	this.Enabled = enabled
	return &this
}

// NewClusterConnectionPoolWithDefaults instantiates a new ClusterConnectionPool object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterConnectionPoolWithDefaults() *ClusterConnectionPool {
	this := ClusterConnectionPool{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *ClusterConnectionPool) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ClusterConnectionPool) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ClusterConnectionPool) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAvailable returns the Available field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterConnectionPool) GetAvailable() bool {
	if o == nil || o.Available.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Available.Get()
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterConnectionPool) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Available.Get(), o.Available.IsSet()
}

// HasAvailable returns a boolean if a field has been set.
func (o *ClusterConnectionPool) HasAvailable() bool {
	return o != nil && o.Available.IsSet()
}

// SetAvailable gets a reference to the given common.NullableBool and assigns it to the Available field.
func (o *ClusterConnectionPool) SetAvailable(v bool) {
	o.Available.Set(&v)
}

// SetAvailableNil sets the value for Available to be an explicit nil.
func (o *ClusterConnectionPool) SetAvailableNil() {
	o.Available.Set(nil)
}

// UnsetAvailable ensures that no value is present for Available, not even an explicit nil.
func (o *ClusterConnectionPool) UnsetAvailable() {
	o.Available.Unset()
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ClusterConnectionPool) GetProvider() ClusterConnectionPoolProvider {
	if o == nil || o.Provider == nil {
		var ret ClusterConnectionPoolProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionPool) GetProviderOk() (*ClusterConnectionPoolProvider, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ClusterConnectionPool) HasProvider() bool {
	return o != nil && o.Provider != nil
}

// SetProvider gets a reference to the given ClusterConnectionPoolProvider and assigns it to the Provider field.
func (o *ClusterConnectionPool) SetProvider(v ClusterConnectionPoolProvider) {
	o.Provider = &v
}

// GetPoolMode returns the PoolMode field value if set, zero value otherwise.
func (o *ClusterConnectionPool) GetPoolMode() ClusterConnectionPoolMode {
	if o == nil || o.PoolMode == nil {
		var ret ClusterConnectionPoolMode
		return ret
	}
	return *o.PoolMode
}

// GetPoolModeOk returns a tuple with the PoolMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionPool) GetPoolModeOk() (*ClusterConnectionPoolMode, bool) {
	if o == nil || o.PoolMode == nil {
		return nil, false
	}
	return o.PoolMode, true
}

// HasPoolMode returns a boolean if a field has been set.
func (o *ClusterConnectionPool) HasPoolMode() bool {
	return o != nil && o.PoolMode != nil
}

// SetPoolMode gets a reference to the given ClusterConnectionPoolMode and assigns it to the PoolMode field.
func (o *ClusterConnectionPool) SetPoolMode(v ClusterConnectionPoolMode) {
	o.PoolMode = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ClusterConnectionPool) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionPool) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ClusterConnectionPool) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *ClusterConnectionPool) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *ClusterConnectionPool) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionPool) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *ClusterConnectionPool) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *ClusterConnectionPool) SetClassCode(v string) {
	o.ClassCode = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ClusterConnectionPool) GetCpu() float64 {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionPool) GetCpuOk() (*float64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ClusterConnectionPool) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given float64 and assigns it to the Cpu field.
func (o *ClusterConnectionPool) SetCpu(v float64) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ClusterConnectionPool) GetMemory() float64 {
	if o == nil || o.Memory == nil {
		var ret float64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionPool) GetMemoryOk() (*float64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ClusterConnectionPool) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given float64 and assigns it to the Memory field.
func (o *ClusterConnectionPool) SetMemory(v float64) {
	o.Memory = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterConnectionPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.Available.IsSet() {
		toSerialize["available"] = o.Available.Get()
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.PoolMode != nil {
		toSerialize["poolMode"] = o.PoolMode
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterConnectionPool) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled   *bool                          `json:"enabled"`
		Available common.NullableBool            `json:"available,omitempty"`
		Provider  *ClusterConnectionPoolProvider `json:"provider,omitempty"`
		PoolMode  *ClusterConnectionPoolMode     `json:"poolMode,omitempty"`
		Replicas  *int32                         `json:"replicas,omitempty"`
		ClassCode *string                        `json:"classCode,omitempty"`
		Cpu       *float64                       `json:"cpu,omitempty"`
		Memory    *float64                       `json:"memory,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "available", "provider", "poolMode", "replicas", "classCode", "cpu", "memory"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = *all.Enabled
	o.Available = all.Available
	if all.Provider != nil && !all.Provider.IsValid() {
		hasInvalidField = true
	} else {
		o.Provider = all.Provider
	}
	if all.PoolMode != nil && !all.PoolMode.IsValid() {
		hasInvalidField = true
	} else {
		o.PoolMode = all.PoolMode
	}
	o.Replicas = all.Replicas
	o.ClassCode = all.ClassCode
	o.Cpu = all.Cpu
	o.Memory = all.Memory

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
