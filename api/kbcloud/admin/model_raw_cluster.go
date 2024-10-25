// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RawCluster cluster info

type RawCluster struct {
	// the cluster name in k8s
	KubernetesName *string `json:"kubernetesName,omitempty"`
	// the namespace of cluster
	Namespace *string `json:"namespace,omitempty"`
	// the replicas of cluster
	Replicas *int32 `json:"replicas,omitempty"`
	// Storage size, the unit is Gi
	Storage *float64 `json:"storage,omitempty"`
	// CPU cores.
	Cpu *float64 `json:"cpu,omitempty"`
	// Memory, the unit is Gi.
	Memory *float64 `json:"memory,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRawCluster instantiates a new RawCluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRawCluster() *RawCluster {
	this := RawCluster{}
	return &this
}

// NewRawClusterWithDefaults instantiates a new RawCluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRawClusterWithDefaults() *RawCluster {
	this := RawCluster{}
	return &this
}

// GetKubernetesName returns the KubernetesName field value if set, zero value otherwise.
func (o *RawCluster) GetKubernetesName() string {
	if o == nil || o.KubernetesName == nil {
		var ret string
		return ret
	}
	return *o.KubernetesName
}

// GetKubernetesNameOk returns a tuple with the KubernetesName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawCluster) GetKubernetesNameOk() (*string, bool) {
	if o == nil || o.KubernetesName == nil {
		return nil, false
	}
	return o.KubernetesName, true
}

// HasKubernetesName returns a boolean if a field has been set.
func (o *RawCluster) HasKubernetesName() bool {
	return o != nil && o.KubernetesName != nil
}

// SetKubernetesName gets a reference to the given string and assigns it to the KubernetesName field.
func (o *RawCluster) SetKubernetesName(v string) {
	o.KubernetesName = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *RawCluster) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawCluster) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *RawCluster) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *RawCluster) SetNamespace(v string) {
	o.Namespace = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *RawCluster) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawCluster) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *RawCluster) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *RawCluster) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *RawCluster) GetStorage() float64 {
	if o == nil || o.Storage == nil {
		var ret float64
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawCluster) GetStorageOk() (*float64, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *RawCluster) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given float64 and assigns it to the Storage field.
func (o *RawCluster) SetStorage(v float64) {
	o.Storage = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *RawCluster) GetCpu() float64 {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawCluster) GetCpuOk() (*float64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *RawCluster) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given float64 and assigns it to the Cpu field.
func (o *RawCluster) SetCpu(v float64) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *RawCluster) GetMemory() float64 {
	if o == nil || o.Memory == nil {
		var ret float64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawCluster) GetMemoryOk() (*float64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *RawCluster) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given float64 and assigns it to the Memory field.
func (o *RawCluster) SetMemory(v float64) {
	o.Memory = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RawCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.KubernetesName != nil {
		toSerialize["kubernetesName"] = o.KubernetesName
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
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
func (o *RawCluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		KubernetesName *string  `json:"kubernetesName,omitempty"`
		Namespace      *string  `json:"namespace,omitempty"`
		Replicas       *int32   `json:"replicas,omitempty"`
		Storage        *float64 `json:"storage,omitempty"`
		Cpu            *float64 `json:"cpu,omitempty"`
		Memory         *float64 `json:"memory,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kubernetesName", "namespace", "replicas", "storage", "cpu", "memory"})
	} else {
		return err
	}
	o.KubernetesName = all.KubernetesName
	o.Namespace = all.Namespace
	o.Replicas = all.Replicas
	o.Storage = all.Storage
	o.Cpu = all.Cpu
	o.Memory = all.Memory

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
