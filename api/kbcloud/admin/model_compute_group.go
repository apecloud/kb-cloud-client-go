// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ComputeGroup struct {
	Name          *string `json:"name,omitempty"`
	ComponentName *string `json:"componentName,omitempty"`
	Managed       *bool   `json:"managed,omitempty"`
	IsDefault     *bool   `json:"isDefault,omitempty"`
	// Kubernetes component phase, independent of engine membership.
	Phase             *string               `json:"phase,omitempty"`
	Stopped           *bool                 `json:"stopped,omitempty"`
	Replicas          *int64                `json:"replicas,omitempty"`
	ClassCode         *string               `json:"classCode,omitempty"`
	Cpu               *float64              `json:"cpu,omitempty"`
	Memory            *float64              `json:"memory,omitempty"`
	BackendCount      *int64                `json:"backendCount,omitempty"`
	AliveBackendCount *int64                `json:"aliveBackendCount,omitempty"`
	DefaultUsers      []string              `json:"defaultUsers,omitempty"`
	Backends          []ComputeGroupBackend `json:"backends,omitempty"`
	AllowedActions    []string              `json:"allowedActions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComputeGroup instantiates a new ComputeGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComputeGroup() *ComputeGroup {
	this := ComputeGroup{}
	return &this
}

// NewComputeGroupWithDefaults instantiates a new ComputeGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComputeGroupWithDefaults() *ComputeGroup {
	this := ComputeGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeGroup) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeGroup) SetName(v string) {
	o.Name = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *ComputeGroup) GetComponentName() string {
	if o == nil || o.ComponentName == nil {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetComponentNameOk() (*string, bool) {
	if o == nil || o.ComponentName == nil {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *ComputeGroup) HasComponentName() bool {
	return o != nil && o.ComponentName != nil
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *ComputeGroup) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *ComputeGroup) GetManaged() bool {
	if o == nil || o.Managed == nil {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetManagedOk() (*bool, bool) {
	if o == nil || o.Managed == nil {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *ComputeGroup) HasManaged() bool {
	return o != nil && o.Managed != nil
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *ComputeGroup) SetManaged(v bool) {
	o.Managed = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ComputeGroup) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ComputeGroup) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ComputeGroup) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *ComputeGroup) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *ComputeGroup) HasPhase() bool {
	return o != nil && o.Phase != nil
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *ComputeGroup) SetPhase(v string) {
	o.Phase = &v
}

// GetStopped returns the Stopped field value if set, zero value otherwise.
func (o *ComputeGroup) GetStopped() bool {
	if o == nil || o.Stopped == nil {
		var ret bool
		return ret
	}
	return *o.Stopped
}

// GetStoppedOk returns a tuple with the Stopped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetStoppedOk() (*bool, bool) {
	if o == nil || o.Stopped == nil {
		return nil, false
	}
	return o.Stopped, true
}

// HasStopped returns a boolean if a field has been set.
func (o *ComputeGroup) HasStopped() bool {
	return o != nil && o.Stopped != nil
}

// SetStopped gets a reference to the given bool and assigns it to the Stopped field.
func (o *ComputeGroup) SetStopped(v bool) {
	o.Stopped = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ComputeGroup) GetReplicas() int64 {
	if o == nil || o.Replicas == nil {
		var ret int64
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetReplicasOk() (*int64, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ComputeGroup) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int64 and assigns it to the Replicas field.
func (o *ComputeGroup) SetReplicas(v int64) {
	o.Replicas = &v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *ComputeGroup) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *ComputeGroup) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *ComputeGroup) SetClassCode(v string) {
	o.ClassCode = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ComputeGroup) GetCpu() float64 {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetCpuOk() (*float64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ComputeGroup) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given float64 and assigns it to the Cpu field.
func (o *ComputeGroup) SetCpu(v float64) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ComputeGroup) GetMemory() float64 {
	if o == nil || o.Memory == nil {
		var ret float64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetMemoryOk() (*float64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ComputeGroup) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given float64 and assigns it to the Memory field.
func (o *ComputeGroup) SetMemory(v float64) {
	o.Memory = &v
}

// GetBackendCount returns the BackendCount field value if set, zero value otherwise.
func (o *ComputeGroup) GetBackendCount() int64 {
	if o == nil || o.BackendCount == nil {
		var ret int64
		return ret
	}
	return *o.BackendCount
}

// GetBackendCountOk returns a tuple with the BackendCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetBackendCountOk() (*int64, bool) {
	if o == nil || o.BackendCount == nil {
		return nil, false
	}
	return o.BackendCount, true
}

// HasBackendCount returns a boolean if a field has been set.
func (o *ComputeGroup) HasBackendCount() bool {
	return o != nil && o.BackendCount != nil
}

// SetBackendCount gets a reference to the given int64 and assigns it to the BackendCount field.
func (o *ComputeGroup) SetBackendCount(v int64) {
	o.BackendCount = &v
}

// GetAliveBackendCount returns the AliveBackendCount field value if set, zero value otherwise.
func (o *ComputeGroup) GetAliveBackendCount() int64 {
	if o == nil || o.AliveBackendCount == nil {
		var ret int64
		return ret
	}
	return *o.AliveBackendCount
}

// GetAliveBackendCountOk returns a tuple with the AliveBackendCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetAliveBackendCountOk() (*int64, bool) {
	if o == nil || o.AliveBackendCount == nil {
		return nil, false
	}
	return o.AliveBackendCount, true
}

// HasAliveBackendCount returns a boolean if a field has been set.
func (o *ComputeGroup) HasAliveBackendCount() bool {
	return o != nil && o.AliveBackendCount != nil
}

// SetAliveBackendCount gets a reference to the given int64 and assigns it to the AliveBackendCount field.
func (o *ComputeGroup) SetAliveBackendCount(v int64) {
	o.AliveBackendCount = &v
}

// GetDefaultUsers returns the DefaultUsers field value if set, zero value otherwise.
func (o *ComputeGroup) GetDefaultUsers() []string {
	if o == nil || o.DefaultUsers == nil {
		var ret []string
		return ret
	}
	return o.DefaultUsers
}

// GetDefaultUsersOk returns a tuple with the DefaultUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetDefaultUsersOk() (*[]string, bool) {
	if o == nil || o.DefaultUsers == nil {
		return nil, false
	}
	return &o.DefaultUsers, true
}

// HasDefaultUsers returns a boolean if a field has been set.
func (o *ComputeGroup) HasDefaultUsers() bool {
	return o != nil && o.DefaultUsers != nil
}

// SetDefaultUsers gets a reference to the given []string and assigns it to the DefaultUsers field.
func (o *ComputeGroup) SetDefaultUsers(v []string) {
	o.DefaultUsers = v
}

// GetBackends returns the Backends field value if set, zero value otherwise.
func (o *ComputeGroup) GetBackends() []ComputeGroupBackend {
	if o == nil || o.Backends == nil {
		var ret []ComputeGroupBackend
		return ret
	}
	return o.Backends
}

// GetBackendsOk returns a tuple with the Backends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetBackendsOk() (*[]ComputeGroupBackend, bool) {
	if o == nil || o.Backends == nil {
		return nil, false
	}
	return &o.Backends, true
}

// HasBackends returns a boolean if a field has been set.
func (o *ComputeGroup) HasBackends() bool {
	return o != nil && o.Backends != nil
}

// SetBackends gets a reference to the given []ComputeGroupBackend and assigns it to the Backends field.
func (o *ComputeGroup) SetBackends(v []ComputeGroupBackend) {
	o.Backends = v
}

// GetAllowedActions returns the AllowedActions field value if set, zero value otherwise.
func (o *ComputeGroup) GetAllowedActions() []string {
	if o == nil || o.AllowedActions == nil {
		var ret []string
		return ret
	}
	return o.AllowedActions
}

// GetAllowedActionsOk returns a tuple with the AllowedActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroup) GetAllowedActionsOk() (*[]string, bool) {
	if o == nil || o.AllowedActions == nil {
		return nil, false
	}
	return &o.AllowedActions, true
}

// HasAllowedActions returns a boolean if a field has been set.
func (o *ComputeGroup) HasAllowedActions() bool {
	return o != nil && o.AllowedActions != nil
}

// SetAllowedActions gets a reference to the given []string and assigns it to the AllowedActions field.
func (o *ComputeGroup) SetAllowedActions(v []string) {
	o.AllowedActions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComputeGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ComponentName != nil {
		toSerialize["componentName"] = o.ComponentName
	}
	if o.Managed != nil {
		toSerialize["managed"] = o.Managed
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.Phase != nil {
		toSerialize["phase"] = o.Phase
	}
	if o.Stopped != nil {
		toSerialize["stopped"] = o.Stopped
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
	if o.BackendCount != nil {
		toSerialize["backendCount"] = o.BackendCount
	}
	if o.AliveBackendCount != nil {
		toSerialize["aliveBackendCount"] = o.AliveBackendCount
	}
	if o.DefaultUsers != nil {
		toSerialize["defaultUsers"] = o.DefaultUsers
	}
	if o.Backends != nil {
		toSerialize["backends"] = o.Backends
	}
	if o.AllowedActions != nil {
		toSerialize["allowedActions"] = o.AllowedActions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComputeGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string               `json:"name,omitempty"`
		ComponentName     *string               `json:"componentName,omitempty"`
		Managed           *bool                 `json:"managed,omitempty"`
		IsDefault         *bool                 `json:"isDefault,omitempty"`
		Phase             *string               `json:"phase,omitempty"`
		Stopped           *bool                 `json:"stopped,omitempty"`
		Replicas          *int64                `json:"replicas,omitempty"`
		ClassCode         *string               `json:"classCode,omitempty"`
		Cpu               *float64              `json:"cpu,omitempty"`
		Memory            *float64              `json:"memory,omitempty"`
		BackendCount      *int64                `json:"backendCount,omitempty"`
		AliveBackendCount *int64                `json:"aliveBackendCount,omitempty"`
		DefaultUsers      []string              `json:"defaultUsers,omitempty"`
		Backends          []ComputeGroupBackend `json:"backends,omitempty"`
		AllowedActions    []string              `json:"allowedActions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "componentName", "managed", "isDefault", "phase", "stopped", "replicas", "classCode", "cpu", "memory", "backendCount", "aliveBackendCount", "defaultUsers", "backends", "allowedActions"})
	} else {
		return err
	}
	o.Name = all.Name
	o.ComponentName = all.ComponentName
	o.Managed = all.Managed
	o.IsDefault = all.IsDefault
	o.Phase = all.Phase
	o.Stopped = all.Stopped
	o.Replicas = all.Replicas
	o.ClassCode = all.ClassCode
	o.Cpu = all.Cpu
	o.Memory = all.Memory
	o.BackendCount = all.BackendCount
	o.AliveBackendCount = all.AliveBackendCount
	o.DefaultUsers = all.DefaultUsers
	o.Backends = all.Backends
	o.AllowedActions = all.AllowedActions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
