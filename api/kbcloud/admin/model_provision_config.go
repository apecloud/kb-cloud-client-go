// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ProvisionConfig Configuration to provision infrastructure for this environment
type ProvisionConfig struct {
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
	// Namespace info for environment
	Namespaces []string `json:"namespaces,omitempty"`
	// Register manifest for environment
	Register   Register    `json:"register"`
	Components []Component `json:"components"`
	// Create your node plan, and the selected nodes will be used for pod scheduling
	NodePool []NodePoolNode `json:"nodePool,omitempty"`
	// Storage config for environment
	Storage StorageConfig `json:"storage"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProvisionConfig instantiates a new ProvisionConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProvisionConfig(defaultStorageClass string, register Register, components []Component, storage StorageConfig) *ProvisionConfig {
	this := ProvisionConfig{}
	this.DefaultStorageClass = defaultStorageClass
	this.Register = register
	this.Components = components
	this.Storage = storage
	return &this
}

// NewProvisionConfigWithDefaults instantiates a new ProvisionConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProvisionConfigWithDefaults() *ProvisionConfig {
	this := ProvisionConfig{}
	return &this
}

// GetKubernetesVersion returns the KubernetesVersion field value if set, zero value otherwise.
func (o *ProvisionConfig) GetKubernetesVersion() string {
	if o == nil || o.KubernetesVersion == nil {
		var ret string
		return ret
	}
	return *o.KubernetesVersion
}

// GetKubernetesVersionOk returns a tuple with the KubernetesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetKubernetesVersionOk() (*string, bool) {
	if o == nil || o.KubernetesVersion == nil {
		return nil, false
	}
	return o.KubernetesVersion, true
}

// HasKubernetesVersion returns a boolean if a field has been set.
func (o *ProvisionConfig) HasKubernetesVersion() bool {
	return o != nil && o.KubernetesVersion != nil
}

// SetKubernetesVersion gets a reference to the given string and assigns it to the KubernetesVersion field.
func (o *ProvisionConfig) SetKubernetesVersion(v string) {
	o.KubernetesVersion = &v
}

// GetKbVersion returns the KbVersion field value if set, zero value otherwise.
func (o *ProvisionConfig) GetKbVersion() string {
	if o == nil || o.KbVersion == nil {
		var ret string
		return ret
	}
	return *o.KbVersion
}

// GetKbVersionOk returns a tuple with the KbVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetKbVersionOk() (*string, bool) {
	if o == nil || o.KbVersion == nil {
		return nil, false
	}
	return o.KbVersion, true
}

// HasKbVersion returns a boolean if a field has been set.
func (o *ProvisionConfig) HasKbVersion() bool {
	return o != nil && o.KbVersion != nil
}

// SetKbVersion gets a reference to the given string and assigns it to the KbVersion field.
func (o *ProvisionConfig) SetKbVersion(v string) {
	o.KbVersion = &v
}

// GetGeminiVersion returns the GeminiVersion field value if set, zero value otherwise.
func (o *ProvisionConfig) GetGeminiVersion() string {
	if o == nil || o.GeminiVersion == nil {
		var ret string
		return ret
	}
	return *o.GeminiVersion
}

// GetGeminiVersionOk returns a tuple with the GeminiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetGeminiVersionOk() (*string, bool) {
	if o == nil || o.GeminiVersion == nil {
		return nil, false
	}
	return o.GeminiVersion, true
}

// HasGeminiVersion returns a boolean if a field has been set.
func (o *ProvisionConfig) HasGeminiVersion() bool {
	return o != nil && o.GeminiVersion != nil
}

// SetGeminiVersion gets a reference to the given string and assigns it to the GeminiVersion field.
func (o *ProvisionConfig) SetGeminiVersion(v string) {
	o.GeminiVersion = &v
}

// GetOteldVersion returns the OteldVersion field value if set, zero value otherwise.
func (o *ProvisionConfig) GetOteldVersion() string {
	if o == nil || o.OteldVersion == nil {
		var ret string
		return ret
	}
	return *o.OteldVersion
}

// GetOteldVersionOk returns a tuple with the OteldVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetOteldVersionOk() (*string, bool) {
	if o == nil || o.OteldVersion == nil {
		return nil, false
	}
	return o.OteldVersion, true
}

// HasOteldVersion returns a boolean if a field has been set.
func (o *ProvisionConfig) HasOteldVersion() bool {
	return o != nil && o.OteldVersion != nil
}

// SetOteldVersion gets a reference to the given string and assigns it to the OteldVersion field.
func (o *ProvisionConfig) SetOteldVersion(v string) {
	o.OteldVersion = &v
}

// GetImageRegistry returns the ImageRegistry field value if set, zero value otherwise.
func (o *ProvisionConfig) GetImageRegistry() string {
	if o == nil || o.ImageRegistry == nil {
		var ret string
		return ret
	}
	return *o.ImageRegistry
}

// GetImageRegistryOk returns a tuple with the ImageRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetImageRegistryOk() (*string, bool) {
	if o == nil || o.ImageRegistry == nil {
		return nil, false
	}
	return o.ImageRegistry, true
}

// HasImageRegistry returns a boolean if a field has been set.
func (o *ProvisionConfig) HasImageRegistry() bool {
	return o != nil && o.ImageRegistry != nil
}

// SetImageRegistry gets a reference to the given string and assigns it to the ImageRegistry field.
func (o *ProvisionConfig) SetImageRegistry(v string) {
	o.ImageRegistry = &v
}

// GetDefaultStorageClass returns the DefaultStorageClass field value.
func (o *ProvisionConfig) GetDefaultStorageClass() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefaultStorageClass
}

// GetDefaultStorageClassOk returns a tuple with the DefaultStorageClass field value
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetDefaultStorageClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultStorageClass, true
}

// SetDefaultStorageClass sets field value.
func (o *ProvisionConfig) SetDefaultStorageClass(v string) {
	o.DefaultStorageClass = v
}

// GetCpuOvercommitRatio returns the CpuOvercommitRatio field value if set, zero value otherwise.
func (o *ProvisionConfig) GetCpuOvercommitRatio() float64 {
	if o == nil || o.CpuOvercommitRatio == nil {
		var ret float64
		return ret
	}
	return *o.CpuOvercommitRatio
}

// GetCpuOvercommitRatioOk returns a tuple with the CpuOvercommitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetCpuOvercommitRatioOk() (*float64, bool) {
	if o == nil || o.CpuOvercommitRatio == nil {
		return nil, false
	}
	return o.CpuOvercommitRatio, true
}

// HasCpuOvercommitRatio returns a boolean if a field has been set.
func (o *ProvisionConfig) HasCpuOvercommitRatio() bool {
	return o != nil && o.CpuOvercommitRatio != nil
}

// SetCpuOvercommitRatio gets a reference to the given float64 and assigns it to the CpuOvercommitRatio field.
func (o *ProvisionConfig) SetCpuOvercommitRatio(v float64) {
	o.CpuOvercommitRatio = &v
}

// GetMemoryOvercommitRatio returns the MemoryOvercommitRatio field value if set, zero value otherwise.
func (o *ProvisionConfig) GetMemoryOvercommitRatio() float64 {
	if o == nil || o.MemoryOvercommitRatio == nil {
		var ret float64
		return ret
	}
	return *o.MemoryOvercommitRatio
}

// GetMemoryOvercommitRatioOk returns a tuple with the MemoryOvercommitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetMemoryOvercommitRatioOk() (*float64, bool) {
	if o == nil || o.MemoryOvercommitRatio == nil {
		return nil, false
	}
	return o.MemoryOvercommitRatio, true
}

// HasMemoryOvercommitRatio returns a boolean if a field has been set.
func (o *ProvisionConfig) HasMemoryOvercommitRatio() bool {
	return o != nil && o.MemoryOvercommitRatio != nil
}

// SetMemoryOvercommitRatio gets a reference to the given float64 and assigns it to the MemoryOvercommitRatio field.
func (o *ProvisionConfig) SetMemoryOvercommitRatio(v float64) {
	o.MemoryOvercommitRatio = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *ProvisionConfig) GetNamespaces() []string {
	if o == nil || o.Namespaces == nil {
		var ret []string
		return ret
	}
	return o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetNamespacesOk() (*[]string, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return &o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *ProvisionConfig) HasNamespaces() bool {
	return o != nil && o.Namespaces != nil
}

// SetNamespaces gets a reference to the given []string and assigns it to the Namespaces field.
func (o *ProvisionConfig) SetNamespaces(v []string) {
	o.Namespaces = v
}

// GetRegister returns the Register field value.
func (o *ProvisionConfig) GetRegister() Register {
	if o == nil {
		var ret Register
		return ret
	}
	return o.Register
}

// GetRegisterOk returns a tuple with the Register field value
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetRegisterOk() (*Register, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Register, true
}

// SetRegister sets field value.
func (o *ProvisionConfig) SetRegister(v Register) {
	o.Register = v
}

// GetComponents returns the Components field value.
func (o *ProvisionConfig) GetComponents() []Component {
	if o == nil {
		var ret []Component
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetComponentsOk() (*[]Component, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value.
func (o *ProvisionConfig) SetComponents(v []Component) {
	o.Components = v
}

// GetNodePool returns the NodePool field value if set, zero value otherwise.
func (o *ProvisionConfig) GetNodePool() []NodePoolNode {
	if o == nil || o.NodePool == nil {
		var ret []NodePoolNode
		return ret
	}
	return o.NodePool
}

// GetNodePoolOk returns a tuple with the NodePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetNodePoolOk() (*[]NodePoolNode, bool) {
	if o == nil || o.NodePool == nil {
		return nil, false
	}
	return &o.NodePool, true
}

// HasNodePool returns a boolean if a field has been set.
func (o *ProvisionConfig) HasNodePool() bool {
	return o != nil && o.NodePool != nil
}

// SetNodePool gets a reference to the given []NodePoolNode and assigns it to the NodePool field.
func (o *ProvisionConfig) SetNodePool(v []NodePoolNode) {
	o.NodePool = v
}

// GetStorage returns the Storage field value.
func (o *ProvisionConfig) GetStorage() StorageConfig {
	if o == nil {
		var ret StorageConfig
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetStorageOk() (*StorageConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value.
func (o *ProvisionConfig) SetStorage(v StorageConfig) {
	o.Storage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProvisionConfig) MarshalJSON() ([]byte, error) {
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
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	toSerialize["register"] = o.Register
	toSerialize["components"] = o.Components
	if o.NodePool != nil {
		toSerialize["nodePool"] = o.NodePool
	}
	toSerialize["storage"] = o.Storage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProvisionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		KubernetesVersion     *string        `json:"kubernetesVersion,omitempty"`
		KbVersion             *string        `json:"kbVersion,omitempty"`
		GeminiVersion         *string        `json:"geminiVersion,omitempty"`
		OteldVersion          *string        `json:"oteldVersion,omitempty"`
		ImageRegistry         *string        `json:"imageRegistry,omitempty"`
		DefaultStorageClass   *string        `json:"defaultStorageClass"`
		CpuOvercommitRatio    *float64       `json:"cpuOvercommitRatio,omitempty"`
		MemoryOvercommitRatio *float64       `json:"memoryOvercommitRatio,omitempty"`
		Namespaces            []string       `json:"namespaces,omitempty"`
		Register              *Register      `json:"register"`
		Components            *[]Component   `json:"components"`
		NodePool              []NodePoolNode `json:"nodePool,omitempty"`
		Storage               *StorageConfig `json:"storage"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DefaultStorageClass == nil {
		return fmt.Errorf("required field defaultStorageClass missing")
	}
	if all.Register == nil {
		return fmt.Errorf("required field register missing")
	}
	if all.Components == nil {
		return fmt.Errorf("required field components missing")
	}
	if all.Storage == nil {
		return fmt.Errorf("required field storage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kubernetesVersion", "kbVersion", "geminiVersion", "oteldVersion", "imageRegistry", "defaultStorageClass", "cpuOvercommitRatio", "memoryOvercommitRatio", "namespaces", "register", "components", "nodePool", "storage"})
	} else {
		return err
	}

	hasInvalidField := false
	o.KubernetesVersion = all.KubernetesVersion
	o.KbVersion = all.KbVersion
	o.GeminiVersion = all.GeminiVersion
	o.OteldVersion = all.OteldVersion
	o.ImageRegistry = all.ImageRegistry
	o.DefaultStorageClass = *all.DefaultStorageClass
	o.CpuOvercommitRatio = all.CpuOvercommitRatio
	o.MemoryOvercommitRatio = all.MemoryOvercommitRatio
	o.Namespaces = all.Namespaces
	if all.Register.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Register = *all.Register
	o.Components = *all.Components
	o.NodePool = all.NodePool
	if all.Storage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Storage = *all.Storage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
