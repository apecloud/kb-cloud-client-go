// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentUpdate Environment info
type EnvironmentUpdate struct {
	// The description of the organization
	Description common.NullableString `json:"description,omitempty"`
	// The display name of the context
	DisplayName common.NullableString `json:"displayName,omitempty"`
	// Type of this environment
	Type *EnvironmentType `json:"type,omitempty"`
	// Organizations that have access for this environment
	Organizations common.NullableList[string] `json:"organizations,omitempty"`
	// Organizations that have access for this environment
	Namespaces common.NullableList[string] `json:"namespaces,omitempty"`
	// the over commit ratio of cpu
	CpuOverCommitRatio common.NullableFloat64 `json:"cpuOverCommitRatio,omitempty"`
	// the over commit ratio of memory
	MemoryOverCommitRatio common.NullableFloat64 `json:"memoryOverCommitRatio,omitempty"`
	// cluster instance autohealing process config
	AutohealingConfig *AutohealingConfig `json:"autohealingConfig,omitempty"`
	// the default storage class of this environment
	DefaultStorageClass common.NullableString `json:"defaultStorageClass,omitempty"`
	// Enable pod antiaffinity for cluster
	PodAntiAffinityEnabled *bool `json:"podAntiAffinityEnabled,omitempty"`
	// the image registry used to pull image
	ImageRegistry common.NullableString `json:"imageRegistry,omitempty"`
	// Enable node port service for this environment
	NodePortEnabled *bool `json:"nodePortEnabled,omitempty"`
	// Enable load balancer service for this environment
	LbEnabled *bool `json:"lbEnabled,omitempty"`
	// Enable the Internet load balancer service for this environment
	InternetLbEnabled *bool `json:"internetLBEnabled,omitempty"`
	// Network modes of the environment
	NetworkModes []NetworkMode `json:"networkModes,omitempty"`
	// Environment delete policy to protect environment from false delete
	DeletePolicy *EnvironmentDeletePolicy `json:"deletePolicy,omitempty"`
	// Cluster operation validation policy, such as create, hscale, vscale, etc.
	ClusterValidationPolicy *ClusterValidationPolicy `json:"clusterValidationPolicy,omitempty"`
	// Cloud Provider
	Provider common.NullableString `json:"provider,omitempty"`
	// whether to enable calculate the cluster SLA for the environment
	SlaEnabled *bool `json:"slaEnabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentUpdate instantiates a new EnvironmentUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentUpdate() *EnvironmentUpdate {
	this := EnvironmentUpdate{}
	var podAntiAffinityEnabled bool = true
	this.PodAntiAffinityEnabled = &podAntiAffinityEnabled
	var nodePortEnabled bool = true
	this.NodePortEnabled = &nodePortEnabled
	var lbEnabled bool = true
	this.LbEnabled = &lbEnabled
	var internetLbEnabled bool = true
	this.InternetLbEnabled = &internetLbEnabled
	var deletePolicy EnvironmentDeletePolicy = EnvironmentDeletePolicyDoNotDelete
	this.DeletePolicy = &deletePolicy
	var clusterValidationPolicy ClusterValidationPolicy = ClusterValidationPolicyValidateOnly
	this.ClusterValidationPolicy = &clusterValidationPolicy
	var slaEnabled bool = false
	this.SlaEnabled = &slaEnabled
	return &this
}

// NewEnvironmentUpdateWithDefaults instantiates a new EnvironmentUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentUpdateWithDefaults() *EnvironmentUpdate {
	this := EnvironmentUpdate{}
	var podAntiAffinityEnabled bool = true
	this.PodAntiAffinityEnabled = &podAntiAffinityEnabled
	var nodePortEnabled bool = true
	this.NodePortEnabled = &nodePortEnabled
	var lbEnabled bool = true
	this.LbEnabled = &lbEnabled
	var internetLbEnabled bool = true
	this.InternetLbEnabled = &internetLbEnabled
	var deletePolicy EnvironmentDeletePolicy = EnvironmentDeletePolicyDoNotDelete
	this.DeletePolicy = &deletePolicy
	var clusterValidationPolicy ClusterValidationPolicy = ClusterValidationPolicyValidateOnly
	this.ClusterValidationPolicy = &clusterValidationPolicy
	var slaEnabled bool = false
	this.SlaEnabled = &slaEnabled
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentUpdate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given common.NullableString and assigns it to the Description field.
func (o *EnvironmentUpdate) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *EnvironmentUpdate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *EnvironmentUpdate) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasDisplayName() bool {
	return o != nil && o.DisplayName.IsSet()
}

// SetDisplayName gets a reference to the given common.NullableString and assigns it to the DisplayName field.
func (o *EnvironmentUpdate) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil.
func (o *EnvironmentUpdate) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil.
func (o *EnvironmentUpdate) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetType() EnvironmentType {
	if o == nil || o.Type == nil {
		var ret EnvironmentType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetTypeOk() (*EnvironmentType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given EnvironmentType and assigns it to the Type field.
func (o *EnvironmentUpdate) SetType(v EnvironmentType) {
	o.Type = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentUpdate) GetOrganizations() []string {
	if o == nil || o.Organizations.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Organizations.Get()
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentUpdate) GetOrganizationsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organizations.Get(), o.Organizations.IsSet()
}

// HasOrganizations returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasOrganizations() bool {
	return o != nil && o.Organizations.IsSet()
}

// SetOrganizations gets a reference to the given common.NullableList[string] and assigns it to the Organizations field.
func (o *EnvironmentUpdate) SetOrganizations(v []string) {
	o.Organizations.Set(&v)
}

// SetOrganizationsNil sets the value for Organizations to be an explicit nil.
func (o *EnvironmentUpdate) SetOrganizationsNil() {
	o.Organizations.Set(nil)
}

// UnsetOrganizations ensures that no value is present for Organizations, not even an explicit nil.
func (o *EnvironmentUpdate) UnsetOrganizations() {
	o.Organizations.Unset()
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentUpdate) GetNamespaces() []string {
	if o == nil || o.Namespaces.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Namespaces.Get()
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentUpdate) GetNamespacesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespaces.Get(), o.Namespaces.IsSet()
}

// HasNamespaces returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasNamespaces() bool {
	return o != nil && o.Namespaces.IsSet()
}

// SetNamespaces gets a reference to the given common.NullableList[string] and assigns it to the Namespaces field.
func (o *EnvironmentUpdate) SetNamespaces(v []string) {
	o.Namespaces.Set(&v)
}

// SetNamespacesNil sets the value for Namespaces to be an explicit nil.
func (o *EnvironmentUpdate) SetNamespacesNil() {
	o.Namespaces.Set(nil)
}

// UnsetNamespaces ensures that no value is present for Namespaces, not even an explicit nil.
func (o *EnvironmentUpdate) UnsetNamespaces() {
	o.Namespaces.Unset()
}

// GetCpuOverCommitRatio returns the CpuOverCommitRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentUpdate) GetCpuOverCommitRatio() float64 {
	if o == nil || o.CpuOverCommitRatio.Get() == nil {
		var ret float64
		return ret
	}
	return *o.CpuOverCommitRatio.Get()
}

// GetCpuOverCommitRatioOk returns a tuple with the CpuOverCommitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentUpdate) GetCpuOverCommitRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuOverCommitRatio.Get(), o.CpuOverCommitRatio.IsSet()
}

// HasCpuOverCommitRatio returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasCpuOverCommitRatio() bool {
	return o != nil && o.CpuOverCommitRatio.IsSet()
}

// SetCpuOverCommitRatio gets a reference to the given common.NullableFloat64 and assigns it to the CpuOverCommitRatio field.
func (o *EnvironmentUpdate) SetCpuOverCommitRatio(v float64) {
	o.CpuOverCommitRatio.Set(&v)
}

// SetCpuOverCommitRatioNil sets the value for CpuOverCommitRatio to be an explicit nil.
func (o *EnvironmentUpdate) SetCpuOverCommitRatioNil() {
	o.CpuOverCommitRatio.Set(nil)
}

// UnsetCpuOverCommitRatio ensures that no value is present for CpuOverCommitRatio, not even an explicit nil.
func (o *EnvironmentUpdate) UnsetCpuOverCommitRatio() {
	o.CpuOverCommitRatio.Unset()
}

// GetMemoryOverCommitRatio returns the MemoryOverCommitRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentUpdate) GetMemoryOverCommitRatio() float64 {
	if o == nil || o.MemoryOverCommitRatio.Get() == nil {
		var ret float64
		return ret
	}
	return *o.MemoryOverCommitRatio.Get()
}

// GetMemoryOverCommitRatioOk returns a tuple with the MemoryOverCommitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentUpdate) GetMemoryOverCommitRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemoryOverCommitRatio.Get(), o.MemoryOverCommitRatio.IsSet()
}

// HasMemoryOverCommitRatio returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasMemoryOverCommitRatio() bool {
	return o != nil && o.MemoryOverCommitRatio.IsSet()
}

// SetMemoryOverCommitRatio gets a reference to the given common.NullableFloat64 and assigns it to the MemoryOverCommitRatio field.
func (o *EnvironmentUpdate) SetMemoryOverCommitRatio(v float64) {
	o.MemoryOverCommitRatio.Set(&v)
}

// SetMemoryOverCommitRatioNil sets the value for MemoryOverCommitRatio to be an explicit nil.
func (o *EnvironmentUpdate) SetMemoryOverCommitRatioNil() {
	o.MemoryOverCommitRatio.Set(nil)
}

// UnsetMemoryOverCommitRatio ensures that no value is present for MemoryOverCommitRatio, not even an explicit nil.
func (o *EnvironmentUpdate) UnsetMemoryOverCommitRatio() {
	o.MemoryOverCommitRatio.Unset()
}

// GetAutohealingConfig returns the AutohealingConfig field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetAutohealingConfig() AutohealingConfig {
	if o == nil || o.AutohealingConfig == nil {
		var ret AutohealingConfig
		return ret
	}
	return *o.AutohealingConfig
}

// GetAutohealingConfigOk returns a tuple with the AutohealingConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetAutohealingConfigOk() (*AutohealingConfig, bool) {
	if o == nil || o.AutohealingConfig == nil {
		return nil, false
	}
	return o.AutohealingConfig, true
}

// HasAutohealingConfig returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasAutohealingConfig() bool {
	return o != nil && o.AutohealingConfig != nil
}

// SetAutohealingConfig gets a reference to the given AutohealingConfig and assigns it to the AutohealingConfig field.
func (o *EnvironmentUpdate) SetAutohealingConfig(v AutohealingConfig) {
	o.AutohealingConfig = &v
}

// GetDefaultStorageClass returns the DefaultStorageClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentUpdate) GetDefaultStorageClass() string {
	if o == nil || o.DefaultStorageClass.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultStorageClass.Get()
}

// GetDefaultStorageClassOk returns a tuple with the DefaultStorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentUpdate) GetDefaultStorageClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultStorageClass.Get(), o.DefaultStorageClass.IsSet()
}

// HasDefaultStorageClass returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasDefaultStorageClass() bool {
	return o != nil && o.DefaultStorageClass.IsSet()
}

// SetDefaultStorageClass gets a reference to the given common.NullableString and assigns it to the DefaultStorageClass field.
func (o *EnvironmentUpdate) SetDefaultStorageClass(v string) {
	o.DefaultStorageClass.Set(&v)
}

// SetDefaultStorageClassNil sets the value for DefaultStorageClass to be an explicit nil.
func (o *EnvironmentUpdate) SetDefaultStorageClassNil() {
	o.DefaultStorageClass.Set(nil)
}

// UnsetDefaultStorageClass ensures that no value is present for DefaultStorageClass, not even an explicit nil.
func (o *EnvironmentUpdate) UnsetDefaultStorageClass() {
	o.DefaultStorageClass.Unset()
}

// GetPodAntiAffinityEnabled returns the PodAntiAffinityEnabled field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetPodAntiAffinityEnabled() bool {
	if o == nil || o.PodAntiAffinityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PodAntiAffinityEnabled
}

// GetPodAntiAffinityEnabledOk returns a tuple with the PodAntiAffinityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetPodAntiAffinityEnabledOk() (*bool, bool) {
	if o == nil || o.PodAntiAffinityEnabled == nil {
		return nil, false
	}
	return o.PodAntiAffinityEnabled, true
}

// HasPodAntiAffinityEnabled returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasPodAntiAffinityEnabled() bool {
	return o != nil && o.PodAntiAffinityEnabled != nil
}

// SetPodAntiAffinityEnabled gets a reference to the given bool and assigns it to the PodAntiAffinityEnabled field.
func (o *EnvironmentUpdate) SetPodAntiAffinityEnabled(v bool) {
	o.PodAntiAffinityEnabled = &v
}

// GetImageRegistry returns the ImageRegistry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentUpdate) GetImageRegistry() string {
	if o == nil || o.ImageRegistry.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImageRegistry.Get()
}

// GetImageRegistryOk returns a tuple with the ImageRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentUpdate) GetImageRegistryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageRegistry.Get(), o.ImageRegistry.IsSet()
}

// HasImageRegistry returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasImageRegistry() bool {
	return o != nil && o.ImageRegistry.IsSet()
}

// SetImageRegistry gets a reference to the given common.NullableString and assigns it to the ImageRegistry field.
func (o *EnvironmentUpdate) SetImageRegistry(v string) {
	o.ImageRegistry.Set(&v)
}

// SetImageRegistryNil sets the value for ImageRegistry to be an explicit nil.
func (o *EnvironmentUpdate) SetImageRegistryNil() {
	o.ImageRegistry.Set(nil)
}

// UnsetImageRegistry ensures that no value is present for ImageRegistry, not even an explicit nil.
func (o *EnvironmentUpdate) UnsetImageRegistry() {
	o.ImageRegistry.Unset()
}

// GetNodePortEnabled returns the NodePortEnabled field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetNodePortEnabled() bool {
	if o == nil || o.NodePortEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NodePortEnabled
}

// GetNodePortEnabledOk returns a tuple with the NodePortEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetNodePortEnabledOk() (*bool, bool) {
	if o == nil || o.NodePortEnabled == nil {
		return nil, false
	}
	return o.NodePortEnabled, true
}

// HasNodePortEnabled returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasNodePortEnabled() bool {
	return o != nil && o.NodePortEnabled != nil
}

// SetNodePortEnabled gets a reference to the given bool and assigns it to the NodePortEnabled field.
func (o *EnvironmentUpdate) SetNodePortEnabled(v bool) {
	o.NodePortEnabled = &v
}

// GetLbEnabled returns the LbEnabled field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetLbEnabled() bool {
	if o == nil || o.LbEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LbEnabled
}

// GetLbEnabledOk returns a tuple with the LbEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetLbEnabledOk() (*bool, bool) {
	if o == nil || o.LbEnabled == nil {
		return nil, false
	}
	return o.LbEnabled, true
}

// HasLbEnabled returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasLbEnabled() bool {
	return o != nil && o.LbEnabled != nil
}

// SetLbEnabled gets a reference to the given bool and assigns it to the LbEnabled field.
func (o *EnvironmentUpdate) SetLbEnabled(v bool) {
	o.LbEnabled = &v
}

// GetInternetLbEnabled returns the InternetLbEnabled field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetInternetLbEnabled() bool {
	if o == nil || o.InternetLbEnabled == nil {
		var ret bool
		return ret
	}
	return *o.InternetLbEnabled
}

// GetInternetLbEnabledOk returns a tuple with the InternetLbEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetInternetLbEnabledOk() (*bool, bool) {
	if o == nil || o.InternetLbEnabled == nil {
		return nil, false
	}
	return o.InternetLbEnabled, true
}

// HasInternetLbEnabled returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasInternetLbEnabled() bool {
	return o != nil && o.InternetLbEnabled != nil
}

// SetInternetLbEnabled gets a reference to the given bool and assigns it to the InternetLbEnabled field.
func (o *EnvironmentUpdate) SetInternetLbEnabled(v bool) {
	o.InternetLbEnabled = &v
}

// GetNetworkModes returns the NetworkModes field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetNetworkModes() []NetworkMode {
	if o == nil || o.NetworkModes == nil {
		var ret []NetworkMode
		return ret
	}
	return o.NetworkModes
}

// GetNetworkModesOk returns a tuple with the NetworkModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetNetworkModesOk() (*[]NetworkMode, bool) {
	if o == nil || o.NetworkModes == nil {
		return nil, false
	}
	return &o.NetworkModes, true
}

// HasNetworkModes returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasNetworkModes() bool {
	return o != nil && o.NetworkModes != nil
}

// SetNetworkModes gets a reference to the given []NetworkMode and assigns it to the NetworkModes field.
func (o *EnvironmentUpdate) SetNetworkModes(v []NetworkMode) {
	o.NetworkModes = v
}

// GetDeletePolicy returns the DeletePolicy field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetDeletePolicy() EnvironmentDeletePolicy {
	if o == nil || o.DeletePolicy == nil {
		var ret EnvironmentDeletePolicy
		return ret
	}
	return *o.DeletePolicy
}

// GetDeletePolicyOk returns a tuple with the DeletePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetDeletePolicyOk() (*EnvironmentDeletePolicy, bool) {
	if o == nil || o.DeletePolicy == nil {
		return nil, false
	}
	return o.DeletePolicy, true
}

// HasDeletePolicy returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasDeletePolicy() bool {
	return o != nil && o.DeletePolicy != nil
}

// SetDeletePolicy gets a reference to the given EnvironmentDeletePolicy and assigns it to the DeletePolicy field.
func (o *EnvironmentUpdate) SetDeletePolicy(v EnvironmentDeletePolicy) {
	o.DeletePolicy = &v
}

// GetClusterValidationPolicy returns the ClusterValidationPolicy field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetClusterValidationPolicy() ClusterValidationPolicy {
	if o == nil || o.ClusterValidationPolicy == nil {
		var ret ClusterValidationPolicy
		return ret
	}
	return *o.ClusterValidationPolicy
}

// GetClusterValidationPolicyOk returns a tuple with the ClusterValidationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetClusterValidationPolicyOk() (*ClusterValidationPolicy, bool) {
	if o == nil || o.ClusterValidationPolicy == nil {
		return nil, false
	}
	return o.ClusterValidationPolicy, true
}

// HasClusterValidationPolicy returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasClusterValidationPolicy() bool {
	return o != nil && o.ClusterValidationPolicy != nil
}

// SetClusterValidationPolicy gets a reference to the given ClusterValidationPolicy and assigns it to the ClusterValidationPolicy field.
func (o *EnvironmentUpdate) SetClusterValidationPolicy(v ClusterValidationPolicy) {
	o.ClusterValidationPolicy = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentUpdate) GetProvider() string {
	if o == nil || o.Provider.Get() == nil {
		var ret string
		return ret
	}
	return *o.Provider.Get()
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentUpdate) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Provider.Get(), o.Provider.IsSet()
}

// HasProvider returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasProvider() bool {
	return o != nil && o.Provider.IsSet()
}

// SetProvider gets a reference to the given common.NullableString and assigns it to the Provider field.
func (o *EnvironmentUpdate) SetProvider(v string) {
	o.Provider.Set(&v)
}

// SetProviderNil sets the value for Provider to be an explicit nil.
func (o *EnvironmentUpdate) SetProviderNil() {
	o.Provider.Set(nil)
}

// UnsetProvider ensures that no value is present for Provider, not even an explicit nil.
func (o *EnvironmentUpdate) UnsetProvider() {
	o.Provider.Unset()
}

// GetSlaEnabled returns the SlaEnabled field value if set, zero value otherwise.
func (o *EnvironmentUpdate) GetSlaEnabled() bool {
	if o == nil || o.SlaEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SlaEnabled
}

// GetSlaEnabledOk returns a tuple with the SlaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentUpdate) GetSlaEnabledOk() (*bool, bool) {
	if o == nil || o.SlaEnabled == nil {
		return nil, false
	}
	return o.SlaEnabled, true
}

// HasSlaEnabled returns a boolean if a field has been set.
func (o *EnvironmentUpdate) HasSlaEnabled() bool {
	return o != nil && o.SlaEnabled != nil
}

// SetSlaEnabled gets a reference to the given bool and assigns it to the SlaEnabled field.
func (o *EnvironmentUpdate) SetSlaEnabled(v bool) {
	o.SlaEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Organizations.IsSet() {
		toSerialize["organizations"] = o.Organizations.Get()
	}
	if o.Namespaces.IsSet() {
		toSerialize["namespaces"] = o.Namespaces.Get()
	}
	if o.CpuOverCommitRatio.IsSet() {
		toSerialize["cpuOverCommitRatio"] = o.CpuOverCommitRatio.Get()
	}
	if o.MemoryOverCommitRatio.IsSet() {
		toSerialize["memoryOverCommitRatio"] = o.MemoryOverCommitRatio.Get()
	}
	if o.AutohealingConfig != nil {
		toSerialize["autohealingConfig"] = o.AutohealingConfig
	}
	if o.DefaultStorageClass.IsSet() {
		toSerialize["defaultStorageClass"] = o.DefaultStorageClass.Get()
	}
	if o.PodAntiAffinityEnabled != nil {
		toSerialize["podAntiAffinityEnabled"] = o.PodAntiAffinityEnabled
	}
	if o.ImageRegistry.IsSet() {
		toSerialize["imageRegistry"] = o.ImageRegistry.Get()
	}
	if o.NodePortEnabled != nil {
		toSerialize["nodePortEnabled"] = o.NodePortEnabled
	}
	if o.LbEnabled != nil {
		toSerialize["lbEnabled"] = o.LbEnabled
	}
	if o.InternetLbEnabled != nil {
		toSerialize["internetLBEnabled"] = o.InternetLbEnabled
	}
	if o.NetworkModes != nil {
		toSerialize["networkModes"] = o.NetworkModes
	}
	if o.DeletePolicy != nil {
		toSerialize["deletePolicy"] = o.DeletePolicy
	}
	if o.ClusterValidationPolicy != nil {
		toSerialize["clusterValidationPolicy"] = o.ClusterValidationPolicy
	}
	if o.Provider.IsSet() {
		toSerialize["provider"] = o.Provider.Get()
	}
	if o.SlaEnabled != nil {
		toSerialize["slaEnabled"] = o.SlaEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description             common.NullableString       `json:"description,omitempty"`
		DisplayName             common.NullableString       `json:"displayName,omitempty"`
		Type                    *EnvironmentType            `json:"type,omitempty"`
		Organizations           common.NullableList[string] `json:"organizations,omitempty"`
		Namespaces              common.NullableList[string] `json:"namespaces,omitempty"`
		CpuOverCommitRatio      common.NullableFloat64      `json:"cpuOverCommitRatio,omitempty"`
		MemoryOverCommitRatio   common.NullableFloat64      `json:"memoryOverCommitRatio,omitempty"`
		AutohealingConfig       *AutohealingConfig          `json:"autohealingConfig,omitempty"`
		DefaultStorageClass     common.NullableString       `json:"defaultStorageClass,omitempty"`
		PodAntiAffinityEnabled  *bool                       `json:"podAntiAffinityEnabled,omitempty"`
		ImageRegistry           common.NullableString       `json:"imageRegistry,omitempty"`
		NodePortEnabled         *bool                       `json:"nodePortEnabled,omitempty"`
		LbEnabled               *bool                       `json:"lbEnabled,omitempty"`
		InternetLbEnabled       *bool                       `json:"internetLBEnabled,omitempty"`
		NetworkModes            []NetworkMode               `json:"networkModes,omitempty"`
		DeletePolicy            *EnvironmentDeletePolicy    `json:"deletePolicy,omitempty"`
		ClusterValidationPolicy *ClusterValidationPolicy    `json:"clusterValidationPolicy,omitempty"`
		Provider                common.NullableString       `json:"provider,omitempty"`
		SlaEnabled              *bool                       `json:"slaEnabled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "displayName", "type", "organizations", "namespaces", "cpuOverCommitRatio", "memoryOverCommitRatio", "autohealingConfig", "defaultStorageClass", "podAntiAffinityEnabled", "imageRegistry", "nodePortEnabled", "lbEnabled", "internetLBEnabled", "networkModes", "deletePolicy", "clusterValidationPolicy", "provider", "slaEnabled"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.DisplayName = all.DisplayName
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Organizations = all.Organizations
	o.Namespaces = all.Namespaces
	o.CpuOverCommitRatio = all.CpuOverCommitRatio
	o.MemoryOverCommitRatio = all.MemoryOverCommitRatio
	if all.AutohealingConfig != nil && all.AutohealingConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AutohealingConfig = all.AutohealingConfig
	o.DefaultStorageClass = all.DefaultStorageClass
	o.PodAntiAffinityEnabled = all.PodAntiAffinityEnabled
	o.ImageRegistry = all.ImageRegistry
	o.NodePortEnabled = all.NodePortEnabled
	o.LbEnabled = all.LbEnabled
	o.InternetLbEnabled = all.InternetLbEnabled
	o.NetworkModes = all.NetworkModes
	if all.DeletePolicy != nil && !all.DeletePolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.DeletePolicy = all.DeletePolicy
	}
	if all.ClusterValidationPolicy != nil && !all.ClusterValidationPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.ClusterValidationPolicy = all.ClusterValidationPolicy
	}
	o.Provider = all.Provider
	o.SlaEnabled = all.SlaEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
