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
	// Register manifest for environment
	Register Register `json:"register"`
	// Component info in environment
	Component Component `json:"component"`
	// Create your node plan, and the selected nodes will be used for pod scheduling
	NodePool []NodePoolNode `json:"nodePool,omitempty"`
	// Storage config for environment
	Storage *StorageConfig `json:"storage,omitempty"`
	// option modules of environment
	Modules []EnvironmentModule `json:"modules,omitempty"`
	// Configuration for log-related components (e.g., Loki)
	LogsConfig *LogsConfig `json:"logsConfig,omitempty"`
	// Configuration for metrics-related components (e.g., Victoriametrics)
	MetricsConfig *MetricsConfig `json:"metricsConfig,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProvisionConfig instantiates a new ProvisionConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProvisionConfig(register Register, component Component) *ProvisionConfig {
	this := ProvisionConfig{}
	this.Register = register
	this.Component = component
	return &this
}

// NewProvisionConfigWithDefaults instantiates a new ProvisionConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProvisionConfigWithDefaults() *ProvisionConfig {
	this := ProvisionConfig{}
	return &this
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

// GetComponent returns the Component field value.
func (o *ProvisionConfig) GetComponent() Component {
	if o == nil {
		var ret Component
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetComponentOk() (*Component, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ProvisionConfig) SetComponent(v Component) {
	o.Component = v
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

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ProvisionConfig) GetStorage() StorageConfig {
	if o == nil || o.Storage == nil {
		var ret StorageConfig
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetStorageOk() (*StorageConfig, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ProvisionConfig) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given StorageConfig and assigns it to the Storage field.
func (o *ProvisionConfig) SetStorage(v StorageConfig) {
	o.Storage = &v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *ProvisionConfig) GetModules() []EnvironmentModule {
	if o == nil || o.Modules == nil {
		var ret []EnvironmentModule
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetModulesOk() (*[]EnvironmentModule, bool) {
	if o == nil || o.Modules == nil {
		return nil, false
	}
	return &o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *ProvisionConfig) HasModules() bool {
	return o != nil && o.Modules != nil
}

// SetModules gets a reference to the given []EnvironmentModule and assigns it to the Modules field.
func (o *ProvisionConfig) SetModules(v []EnvironmentModule) {
	o.Modules = v
}

// GetLogsConfig returns the LogsConfig field value if set, zero value otherwise.
func (o *ProvisionConfig) GetLogsConfig() LogsConfig {
	if o == nil || o.LogsConfig == nil {
		var ret LogsConfig
		return ret
	}
	return *o.LogsConfig
}

// GetLogsConfigOk returns a tuple with the LogsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetLogsConfigOk() (*LogsConfig, bool) {
	if o == nil || o.LogsConfig == nil {
		return nil, false
	}
	return o.LogsConfig, true
}

// HasLogsConfig returns a boolean if a field has been set.
func (o *ProvisionConfig) HasLogsConfig() bool {
	return o != nil && o.LogsConfig != nil
}

// SetLogsConfig gets a reference to the given LogsConfig and assigns it to the LogsConfig field.
func (o *ProvisionConfig) SetLogsConfig(v LogsConfig) {
	o.LogsConfig = &v
}

// GetMetricsConfig returns the MetricsConfig field value if set, zero value otherwise.
func (o *ProvisionConfig) GetMetricsConfig() MetricsConfig {
	if o == nil || o.MetricsConfig == nil {
		var ret MetricsConfig
		return ret
	}
	return *o.MetricsConfig
}

// GetMetricsConfigOk returns a tuple with the MetricsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionConfig) GetMetricsConfigOk() (*MetricsConfig, bool) {
	if o == nil || o.MetricsConfig == nil {
		return nil, false
	}
	return o.MetricsConfig, true
}

// HasMetricsConfig returns a boolean if a field has been set.
func (o *ProvisionConfig) HasMetricsConfig() bool {
	return o != nil && o.MetricsConfig != nil
}

// SetMetricsConfig gets a reference to the given MetricsConfig and assigns it to the MetricsConfig field.
func (o *ProvisionConfig) SetMetricsConfig(v MetricsConfig) {
	o.MetricsConfig = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProvisionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["register"] = o.Register
	toSerialize["component"] = o.Component
	if o.NodePool != nil {
		toSerialize["nodePool"] = o.NodePool
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Modules != nil {
		toSerialize["modules"] = o.Modules
	}
	if o.LogsConfig != nil {
		toSerialize["logsConfig"] = o.LogsConfig
	}
	if o.MetricsConfig != nil {
		toSerialize["metricsConfig"] = o.MetricsConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProvisionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Register      *Register           `json:"register"`
		Component     *Component          `json:"component"`
		NodePool      []NodePoolNode      `json:"nodePool,omitempty"`
		Storage       *StorageConfig      `json:"storage,omitempty"`
		Modules       []EnvironmentModule `json:"modules,omitempty"`
		LogsConfig    *LogsConfig         `json:"logsConfig,omitempty"`
		MetricsConfig *MetricsConfig      `json:"metricsConfig,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Register == nil {
		return fmt.Errorf("required field register missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"register", "component", "nodePool", "storage", "modules", "logsConfig", "metricsConfig"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Register.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Register = *all.Register
	if all.Component.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Component = *all.Component
	o.NodePool = all.NodePool
	if all.Storage != nil && all.Storage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Storage = all.Storage
	o.Modules = all.Modules
	if all.LogsConfig != nil && all.LogsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogsConfig = all.LogsConfig
	if all.MetricsConfig != nil && all.MetricsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricsConfig = all.MetricsConfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
