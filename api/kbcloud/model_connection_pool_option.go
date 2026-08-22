// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ConnectionPoolOption struct {
	// Whether this engine supports an independent connection pool component.
	Supported bool `json:"supported"`
	// Connection pool implementation managed by KBE.
	Provider ClusterConnectionPoolProvider `json:"provider"`
	// EngineOption component name used by the connection pool.
	Component string `json:"component"`
	// Connection reuse mode. The MVP supports session pooling only.
	DefaultPoolMode ClusterConnectionPoolMode   `json:"defaultPoolMode"`
	PoolModes       []ClusterConnectionPoolMode `json:"poolModes"`
	// Compatible KubeBlocks major versions. Empty means all versions.
	CompatibleKubeblocksVersions []string `json:"compatibleKubeblocksVersions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConnectionPoolOption instantiates a new ConnectionPoolOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConnectionPoolOption(supported bool, provider ClusterConnectionPoolProvider, component string, defaultPoolMode ClusterConnectionPoolMode, poolModes []ClusterConnectionPoolMode) *ConnectionPoolOption {
	this := ConnectionPoolOption{}
	this.Supported = supported
	this.Provider = provider
	this.Component = component
	this.DefaultPoolMode = defaultPoolMode
	this.PoolModes = poolModes
	return &this
}

// NewConnectionPoolOptionWithDefaults instantiates a new ConnectionPoolOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConnectionPoolOptionWithDefaults() *ConnectionPoolOption {
	this := ConnectionPoolOption{}
	return &this
}

// GetSupported returns the Supported field value.
func (o *ConnectionPoolOption) GetSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *ConnectionPoolOption) GetSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supported, true
}

// SetSupported sets field value.
func (o *ConnectionPoolOption) SetSupported(v bool) {
	o.Supported = v
}

// GetProvider returns the Provider field value.
func (o *ConnectionPoolOption) GetProvider() ClusterConnectionPoolProvider {
	if o == nil {
		var ret ClusterConnectionPoolProvider
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ConnectionPoolOption) GetProviderOk() (*ClusterConnectionPoolProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *ConnectionPoolOption) SetProvider(v ClusterConnectionPoolProvider) {
	o.Provider = v
}

// GetComponent returns the Component field value.
func (o *ConnectionPoolOption) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ConnectionPoolOption) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ConnectionPoolOption) SetComponent(v string) {
	o.Component = v
}

// GetDefaultPoolMode returns the DefaultPoolMode field value.
func (o *ConnectionPoolOption) GetDefaultPoolMode() ClusterConnectionPoolMode {
	if o == nil {
		var ret ClusterConnectionPoolMode
		return ret
	}
	return o.DefaultPoolMode
}

// GetDefaultPoolModeOk returns a tuple with the DefaultPoolMode field value
// and a boolean to check if the value has been set.
func (o *ConnectionPoolOption) GetDefaultPoolModeOk() (*ClusterConnectionPoolMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultPoolMode, true
}

// SetDefaultPoolMode sets field value.
func (o *ConnectionPoolOption) SetDefaultPoolMode(v ClusterConnectionPoolMode) {
	o.DefaultPoolMode = v
}

// GetPoolModes returns the PoolModes field value.
func (o *ConnectionPoolOption) GetPoolModes() []ClusterConnectionPoolMode {
	if o == nil {
		var ret []ClusterConnectionPoolMode
		return ret
	}
	return o.PoolModes
}

// GetPoolModesOk returns a tuple with the PoolModes field value
// and a boolean to check if the value has been set.
func (o *ConnectionPoolOption) GetPoolModesOk() (*[]ClusterConnectionPoolMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolModes, true
}

// SetPoolModes sets field value.
func (o *ConnectionPoolOption) SetPoolModes(v []ClusterConnectionPoolMode) {
	o.PoolModes = v
}

// GetCompatibleKubeblocksVersions returns the CompatibleKubeblocksVersions field value if set, zero value otherwise.
func (o *ConnectionPoolOption) GetCompatibleKubeblocksVersions() []string {
	if o == nil || o.CompatibleKubeblocksVersions == nil {
		var ret []string
		return ret
	}
	return o.CompatibleKubeblocksVersions
}

// GetCompatibleKubeblocksVersionsOk returns a tuple with the CompatibleKubeblocksVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionPoolOption) GetCompatibleKubeblocksVersionsOk() (*[]string, bool) {
	if o == nil || o.CompatibleKubeblocksVersions == nil {
		return nil, false
	}
	return &o.CompatibleKubeblocksVersions, true
}

// HasCompatibleKubeblocksVersions returns a boolean if a field has been set.
func (o *ConnectionPoolOption) HasCompatibleKubeblocksVersions() bool {
	return o != nil && o.CompatibleKubeblocksVersions != nil
}

// SetCompatibleKubeblocksVersions gets a reference to the given []string and assigns it to the CompatibleKubeblocksVersions field.
func (o *ConnectionPoolOption) SetCompatibleKubeblocksVersions(v []string) {
	o.CompatibleKubeblocksVersions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConnectionPoolOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["supported"] = o.Supported
	toSerialize["provider"] = o.Provider
	toSerialize["component"] = o.Component
	toSerialize["defaultPoolMode"] = o.DefaultPoolMode
	toSerialize["poolModes"] = o.PoolModes
	if o.CompatibleKubeblocksVersions != nil {
		toSerialize["compatibleKubeblocksVersions"] = o.CompatibleKubeblocksVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConnectionPoolOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Supported                    *bool                          `json:"supported"`
		Provider                     *ClusterConnectionPoolProvider `json:"provider"`
		Component                    *string                        `json:"component"`
		DefaultPoolMode              *ClusterConnectionPoolMode     `json:"defaultPoolMode"`
		PoolModes                    *[]ClusterConnectionPoolMode   `json:"poolModes"`
		CompatibleKubeblocksVersions []string                       `json:"compatibleKubeblocksVersions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Supported == nil {
		return fmt.Errorf("required field supported missing")
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.DefaultPoolMode == nil {
		return fmt.Errorf("required field defaultPoolMode missing")
	}
	if all.PoolModes == nil {
		return fmt.Errorf("required field poolModes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"supported", "provider", "component", "defaultPoolMode", "poolModes", "compatibleKubeblocksVersions"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Supported = *all.Supported
	if !all.Provider.IsValid() {
		hasInvalidField = true
	} else {
		o.Provider = *all.Provider
	}
	o.Component = *all.Component
	if !all.DefaultPoolMode.IsValid() {
		hasInvalidField = true
	} else {
		o.DefaultPoolMode = *all.DefaultPoolMode
	}
	o.PoolModes = *all.PoolModes
	o.CompatibleKubeblocksVersions = all.CompatibleKubeblocksVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
