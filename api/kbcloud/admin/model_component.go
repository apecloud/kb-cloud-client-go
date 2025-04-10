// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Component Component info in environment
type Component struct {
	// Kubernetes version of the environment
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`
	// KubeBlocks version of the environment
	KbVersion *string `json:"kbVersion,omitempty"`
	// Gemini version of the environment
	GeminiVersion *string `json:"geminiVersion,omitempty"`
	// Oteld version of the environment
	OteldVersion *string `json:"oteldVersion,omitempty"`
	// Image registry used by the environment
	ImageRegistry *string `json:"imageRegistry,omitempty"`
	// Default storage class that used by KubeBlocks to create cluster
	DefaultStorageClass string `json:"defaultStorageClass"`
	// CPU overcommit ratio of this environment
	CpuOvercommitRatio *float64 `json:"cpuOvercommitRatio,omitempty"`
	// Memory overcommit ratio of this environment
	MemoryOvercommitRatio *float64 `json:"memoryOvercommitRatio,omitempty"`
	// the replicas of core componment, such as kubeblocks and gemini
	Replicas *int32 `json:"replicas,omitempty"`
	// Namespace info for environment
	Namespaces []string `json:"namespaces,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponent instantiates a new Component object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponent(defaultStorageClass string) *Component {
	this := Component{}
	this.DefaultStorageClass = defaultStorageClass
	var replicas int32 = 1
	this.Replicas = &replicas
	return &this
}

// NewComponentWithDefaults instantiates a new Component object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentWithDefaults() *Component {
	this := Component{}
	var replicas int32 = 1
	this.Replicas = &replicas
	return &this
}

// GetKubernetesVersion returns the KubernetesVersion field value if set, zero value otherwise.
func (o *Component) GetKubernetesVersion() string {
	if o == nil || o.KubernetesVersion == nil {
		var ret string
		return ret
	}
	return *o.KubernetesVersion
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetKubernetesVersionOk() (*string, bool) {
	if o == nil || o.KubernetesVersion == nil {
		return nil, false
	}
	return o.KubernetesVersion, true
}

// HasKubernetesVersion returns a boolean if a field has been set.
func (o *Component) HasKubernetesVersion() bool {
	return o != nil && o.KubernetesVersion != nil
}

// SetKubernetesVersion gets a reference to the given string and assigns it to the KubernetesVersion field.
func (o *Component) SetKubernetesVersion(v string) {
	o.KubernetesVersion = &v
}

// GetKbVersion returns the KbVersion field value if set, zero value otherwise.
func (o *Component) GetKbVersion() string {
	if o == nil || o.KbVersion == nil {
		var ret string
		return ret
	}
	return *o.KbVersion
}

// GetKbVersionOk returns a tuple with the KbVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetKbVersionOk() (*string, bool) {
	if o == nil || o.KbVersion == nil {
		return nil, false
	}
	return o.KbVersion, true
}

// HasKbVersion returns a boolean if a field has been set.
func (o *Component) HasKbVersion() bool {
	return o != nil && o.KbVersion != nil
}

// SetKbVersion gets a reference to the given string and assigns it to the KbVersion field.
func (o *Component) SetKbVersion(v string) {
	o.KbVersion = &v
}

// GetGeminiVersion returns the GeminiVersion field value if set, zero value otherwise.
func (o *Component) GetGeminiVersion() string {
	if o == nil || o.GeminiVersion == nil {
		var ret string
		return ret
	}
	return *o.GeminiVersion
}

// GetGeminiVersionOk returns a tuple with the GeminiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetGeminiVersionOk() (*string, bool) {
	if o == nil || o.GeminiVersion == nil {
		return nil, false
	}
	return o.GeminiVersion, true
}

// HasGeminiVersion returns a boolean if a field has been set.
func (o *Component) HasGeminiVersion() bool {
	return o != nil && o.GeminiVersion != nil
}

// SetGeminiVersion gets a reference to the given string and assigns it to the GeminiVersion field.
func (o *Component) SetGeminiVersion(v string) {
	o.GeminiVersion = &v
}

// GetOteldVersion returns the OteldVersion field value if set, zero value otherwise.
func (o *Component) GetOteldVersion() string {
	if o == nil || o.OteldVersion == nil {
		var ret string
		return ret
	}
	return *o.OteldVersion
}

// GetOteldVersionOk returns a tuple with the OteldVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetOteldVersionOk() (*string, bool) {
	if o == nil || o.OteldVersion == nil {
		return nil, false
	}
	return o.OteldVersion, true
}

// HasOteldVersion returns a boolean if a field has been set.
func (o *Component) HasOteldVersion() bool {
	return o != nil && o.OteldVersion != nil
}

// SetOteldVersion gets a reference to the given string and assigns it to the OteldVersion field.
func (o *Component) SetOteldVersion(v string) {
	o.OteldVersion = &v
}

// GetImageRegistry returns the ImageRegistry field value if set, zero value otherwise.
func (o *Component) GetImageRegistry() string {
	if o == nil || o.ImageRegistry == nil {
		var ret string
		return ret
	}
	return *o.ImageRegistry
}

// GetImageRegistryOk returns a tuple with the ImageRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetImageRegistryOk() (*string, bool) {
	if o == nil || o.ImageRegistry == nil {
		return nil, false
	}
	return o.ImageRegistry, true
}

// HasImageRegistry returns a boolean if a field has been set.
func (o *Component) HasImageRegistry() bool {
	return o != nil && o.ImageRegistry != nil
}

// SetImageRegistry gets a reference to the given string and assigns it to the ImageRegistry field.
func (o *Component) SetImageRegistry(v string) {
	o.ImageRegistry = &v
}

// GetDefaultStorageClass returns the DefaultStorageClass field value.
func (o *Component) GetDefaultStorageClass() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefaultStorageClass
}

// GetDefaultStorageClassOk returns a tuple with the DefaultStorageClass field value
// and a boolean to check if the value has been set.
func (o *Component) GetDefaultStorageClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultStorageClass, true
}

// SetDefaultStorageClass sets field value.
func (o *Component) SetDefaultStorageClass(v string) {
	o.DefaultStorageClass = v
}

// GetCpuOvercommitRatio returns the CpuOvercommitRatio field value if set, zero value otherwise.
func (o *Component) GetCpuOvercommitRatio() float64 {
	if o == nil || o.CpuOvercommitRatio == nil {
		var ret float64
		return ret
	}
	return *o.CpuOvercommitRatio
}

// GetCpuOvercommitRatioOk returns a tuple with the CpuOvercommitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetCpuOvercommitRatioOk() (*float64, bool) {
	if o == nil || o.CpuOvercommitRatio == nil {
		return nil, false
	}
	return o.CpuOvercommitRatio, true
}

// HasCpuOvercommitRatio returns a boolean if a field has been set.
func (o *Component) HasCpuOvercommitRatio() bool {
	return o != nil && o.CpuOvercommitRatio != nil
}

// SetCpuOvercommitRatio gets a reference to the given float64 and assigns it to the CpuOvercommitRatio field.
func (o *Component) SetCpuOvercommitRatio(v float64) {
	o.CpuOvercommitRatio = &v
}

// GetMemoryOvercommitRatio returns the MemoryOvercommitRatio field value if set, zero value otherwise.
func (o *Component) GetMemoryOvercommitRatio() float64 {
	if o == nil || o.MemoryOvercommitRatio == nil {
		var ret float64
		return ret
	}
	return *o.MemoryOvercommitRatio
}

// GetMemoryOvercommitRatioOk returns a tuple with the MemoryOvercommitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetMemoryOvercommitRatioOk() (*float64, bool) {
	if o == nil || o.MemoryOvercommitRatio == nil {
		return nil, false
	}
	return o.MemoryOvercommitRatio, true
}

// HasMemoryOvercommitRatio returns a boolean if a field has been set.
func (o *Component) HasMemoryOvercommitRatio() bool {
	return o != nil && o.MemoryOvercommitRatio != nil
}

// SetMemoryOvercommitRatio gets a reference to the given float64 and assigns it to the MemoryOvercommitRatio field.
func (o *Component) SetMemoryOvercommitRatio(v float64) {
	o.MemoryOvercommitRatio = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *Component) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *Component) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *Component) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *Component) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetNamespacesOk() (*[]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return &o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *Component) HasNamespaces() bool {
	return o != nil && o.Namespaces != nil
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *Component) SetNamespaces(v []string) {
	o.Namespaces = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Component) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.KubernetesVersion != nil {
		toSerialize["kubernetesVersion"] = o.KubernetesVersion
	}
	if o.KbVersion != nil {
		toSerialize["kbVersion"] = o.KbVersion
	}
	if o.GeminiVersion != nil {
		toSerialize["geminiVersion"] = o.GeminiVersion
	}
	if o.OteldVersion != nil {
		toSerialize["oteldVersion"] = o.OteldVersion
	}
	if o.ImageRegistry != nil {
		toSerialize["imageRegistry"] = o.ImageRegistry
	}
	toSerialize["defaultStorageClass"] = o.DefaultStorageClass
	if o.CpuOvercommitRatio != nil {
		toSerialize["cpuOvercommitRatio"] = o.CpuOvercommitRatio
	}
	if o.MemoryOvercommitRatio != nil {
		toSerialize["memoryOvercommitRatio"] = o.MemoryOvercommitRatio
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Component) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		KubernetesVersion     *string  `json:"kubernetesVersion,omitempty"`
		KbVersion             *string  `json:"kbVersion,omitempty"`
		GeminiVersion         *string  `json:"geminiVersion,omitempty"`
		OteldVersion          *string  `json:"oteldVersion,omitempty"`
		ImageRegistry         *string  `json:"imageRegistry,omitempty"`
		DefaultStorageClass   *string  `json:"defaultStorageClass"`
		CpuOvercommitRatio    *float64 `json:"cpuOvercommitRatio,omitempty"`
		MemoryOvercommitRatio *float64 `json:"memoryOvercommitRatio,omitempty"`
		Replicas              *int32   `json:"replicas,omitempty"`
		Namespaces            []string `json:"namespaces,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DefaultStorageClass == nil {
		return fmt.Errorf("required field defaultStorageClass missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kubernetesVersion", "kbVersion", "geminiVersion", "oteldVersion", "imageRegistry", "defaultStorageClass", "cpuOvercommitRatio", "memoryOvercommitRatio", "replicas", "namespaces"})
	} else {
		return err
	}
	o.KubernetesVersion = all.KubernetesVersion
	o.KbVersion = all.KbVersion
	o.GeminiVersion = all.GeminiVersion
	o.OteldVersion = all.OteldVersion
	o.ImageRegistry = all.ImageRegistry
	o.DefaultStorageClass = *all.DefaultStorageClass
	o.CpuOvercommitRatio = all.CpuOvercommitRatio
	o.MemoryOvercommitRatio = all.MemoryOvercommitRatio
	o.Replicas = all.Replicas
	o.Namespaces = all.Namespaces

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
