// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"
)

// Cluster KubeBlocks cluster information
type Cluster struct {
	// Cluster ID
	Id interface{} `json:"id,omitempty"`
	// When two clusters have a relationship, parentId records the parent cluster id.Can be empty when there is no relationship
	ParentId common.NullableInt64 `json:"parentId,omitempty"`
	// the name of parent cluster
	ParentName common.NullableString `json:"parentName,omitempty"`
	// the display name of parent cluster
	ParentDisplayName common.NullableString `json:"parentDisplayName,omitempty"`
	// Describes the type of cluster, [Normal] normal cluster; [DisasterRecovery] disaster recovery cluster
	ClusterType NullableClusterType    `json:"clusterType,omitempty"`
	Delay       common.NullableFloat64 `json:"delay,omitempty"`
	// Org Name
	OrgName *string `json:"orgName,omitempty"`
	// Cloud Provider
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// ID of the environment
	EnvironmentId *string `json:"environmentId,omitempty"`
	// Environment Name
	EnvironmentName string `json:"environmentName"`
	// Environment Types
	EnvironmentType *string `json:"environmentType,omitempty"`
	// Cloud Region
	CloudRegion *string `json:"cloudRegion,omitempty"`
	// Environment Namespace
	// Deprecated
	Namespace *string `json:"namespace,omitempty"`
	// Name of project
	Project *string `json:"project,omitempty"`
	// Name of cluster. Name must be unique within an Org
	Name string `json:"name"`
	// Hash of cluster. Name must be unique within an Org
	Hash *string `json:"hash,omitempty"`
	// Cluster Application Engine
	Engine  string          `json:"engine"`
	License *ClusterLicense `json:"license,omitempty"`
	// Items is the list of parameter template in the list
	ParamTpls []ParamTplsItem `json:"paramTpls,omitempty"`
	// Values
	Values map[string]interface{} `json:"Values,omitempty"`
	// Cluster Application Version
	Version *string `json:"version,omitempty"`
	// The termination policy of cluster.
	TerminationPolicy *ClusterTerminationPolicy `json:"terminationPolicy,omitempty"`
	// Specify whether the cluster enable monitoring.
	MonitorEnabled *bool `json:"monitorEnabled,omitempty"`
	// Specify whether the cluster can be accessed from within the VPC.
	VpcEndpointEnabled *bool `json:"vpcEndpointEnabled,omitempty"`
	// Specify whether the cluster can be accessed from the public internet.
	InternetEndpointEnabled *bool `json:"internetEndpointEnabled,omitempty"`
	// Enable the cluster to provide TLS connection or not.
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`
	// Enable the cluster to provide NodePort service or not.
	NodePortEnabled *bool `json:"nodePortEnabled,omitempty"`
	// Cluster Status
	Status *string `json:"status,omitempty"`
	// CreatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// UpdatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// CreateOnlySet
	CreateOnlySet *bool `json:"createOnlySet,omitempty"`
	// Cluster topology mode
	Mode *string `json:"mode,omitempty"`
	// Proxy Enabled
	ProxyEnabled *bool `json:"proxyEnabled,omitempty"`
	// Items is the list of ComponentSpec in the list
	Components []ComponentsItem `json:"components,omitempty"`
	// Values
	Extra map[string]interface{} `json:"extra,omitempty"`
	// Items is the list of init Options in the list
	InitOptions []InitOptionsItem `json:"initOptions,omitempty"`
	// ExtraInfo
	ExtraInfo map[string]interface{} `json:"extraInfo,omitempty"`
	// Tolerations of cluster
	Tolerations *string `json:"tolerations,omitempty"`
	// Use single availability zones
	SingleZone *bool `json:"singleZone,omitempty"`
	// Availability Zones
	AvailabilityZones []string `json:"availabilityZones,omitempty"`
	// Enable pod antiaffinity for cluster
	PodAntiAffinityEnabled *bool `json:"podAntiAffinityEnabled,omitempty"`
	// clusterBackup is the payload for cluster backup
	Backup *ClusterBackup `json:"backup,omitempty"`
	// Specify a NodeGroup for the cluster
	NodeGroup common.NullableString `json:"nodeGroup,omitempty"`
	// Cluster main component codeShort
	CodeShort *string `json:"codeShort,omitempty"`
	// Display name of cluster.
	DisplayName *string `json:"displayName,omitempty"`
	// if cluster is static cluster
	Static *bool `json:"static,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCluster instantiates a new Cluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCluster(environmentName string, name string, engine string) *Cluster {
	this := Cluster{}
	var clusterType ClusterType = ClusterTypeNormal
	this.ClusterType = *NewNullableClusterType(&clusterType)
	this.EnvironmentName = environmentName
	var namespace string = "kubeblocks-cloud-ns"
	this.Namespace = &namespace
	var project string = "kubeblocks-cloud-ns"
	this.Project = &project
	this.Name = name
	this.Engine = engine
	var terminationPolicy ClusterTerminationPolicy = ClusterTerminationPolicyDelete
	this.TerminationPolicy = &terminationPolicy
	var monitorEnabled bool = false
	this.MonitorEnabled = &monitorEnabled
	var vpcEndpointEnabled bool = false
	this.VpcEndpointEnabled = &vpcEndpointEnabled
	var internetEndpointEnabled bool = false
	this.InternetEndpointEnabled = &internetEndpointEnabled
	var tlsEnabled bool = false
	this.TlsEnabled = &tlsEnabled
	var nodePortEnabled bool = false
	this.NodePortEnabled = &nodePortEnabled
	var proxyEnabled bool = false
	this.ProxyEnabled = &proxyEnabled
	var singleZone bool = false
	this.SingleZone = &singleZone
	var podAntiAffinityEnabled bool = true
	this.PodAntiAffinityEnabled = &podAntiAffinityEnabled
	return &this
}

// NewClusterWithDefaults instantiates a new Cluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterWithDefaults() *Cluster {
	this := Cluster{}
	var clusterType ClusterType = ClusterTypeNormal
	this.ClusterType = *NewNullableClusterType(&clusterType)
	var namespace string = "kubeblocks-cloud-ns"
	this.Namespace = &namespace
	var project string = "kubeblocks-cloud-ns"
	this.Project = &project
	var terminationPolicy ClusterTerminationPolicy = ClusterTerminationPolicyDelete
	this.TerminationPolicy = &terminationPolicy
	var monitorEnabled bool = false
	this.MonitorEnabled = &monitorEnabled
	var vpcEndpointEnabled bool = false
	this.VpcEndpointEnabled = &vpcEndpointEnabled
	var internetEndpointEnabled bool = false
	this.InternetEndpointEnabled = &internetEndpointEnabled
	var tlsEnabled bool = false
	this.TlsEnabled = &tlsEnabled
	var nodePortEnabled bool = false
	this.NodePortEnabled = &nodePortEnabled
	var proxyEnabled bool = false
	this.ProxyEnabled = &proxyEnabled
	var singleZone bool = false
	this.SingleZone = &singleZone
	var podAntiAffinityEnabled bool = true
	this.PodAntiAffinityEnabled = &podAntiAffinityEnabled
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Cluster) GetId() interface{} {
	if o == nil || o.Id == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Cluster) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *Cluster) SetId(v interface{}) {
	o.Id = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetParentId() int64 {
	if o == nil || o.ParentId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Cluster) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *Cluster) HasParentId() bool {
	return o != nil && o.ParentId.IsSet()
}

// SetParentId gets a reference to the given common.NullableInt64 and assigns it to the ParentId field.
func (o *Cluster) SetParentId(v int64) {
	o.ParentId.Set(&v)
}

// SetParentIdNil sets the value for ParentId to be an explicit nil.
func (o *Cluster) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil.
func (o *Cluster) UnsetParentId() {
	o.ParentId.Unset()
}

// GetParentName returns the ParentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetParentName() string {
	if o == nil || o.ParentName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentName.Get()
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Cluster) GetParentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentName.Get(), o.ParentName.IsSet()
}

// HasParentName returns a boolean if a field has been set.
func (o *Cluster) HasParentName() bool {
	return o != nil && o.ParentName.IsSet()
}

// SetParentName gets a reference to the given common.NullableString and assigns it to the ParentName field.
func (o *Cluster) SetParentName(v string) {
	o.ParentName.Set(&v)
}

// SetParentNameNil sets the value for ParentName to be an explicit nil.
func (o *Cluster) SetParentNameNil() {
	o.ParentName.Set(nil)
}

// UnsetParentName ensures that no value is present for ParentName, not even an explicit nil.
func (o *Cluster) UnsetParentName() {
	o.ParentName.Unset()
}

// GetParentDisplayName returns the ParentDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetParentDisplayName() string {
	if o == nil || o.ParentDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentDisplayName.Get()
}

// GetParentDisplayNameOk returns a tuple with the ParentDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Cluster) GetParentDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentDisplayName.Get(), o.ParentDisplayName.IsSet()
}

// HasParentDisplayName returns a boolean if a field has been set.
func (o *Cluster) HasParentDisplayName() bool {
	return o != nil && o.ParentDisplayName.IsSet()
}

// SetParentDisplayName gets a reference to the given common.NullableString and assigns it to the ParentDisplayName field.
func (o *Cluster) SetParentDisplayName(v string) {
	o.ParentDisplayName.Set(&v)
}

// SetParentDisplayNameNil sets the value for ParentDisplayName to be an explicit nil.
func (o *Cluster) SetParentDisplayNameNil() {
	o.ParentDisplayName.Set(nil)
}

// UnsetParentDisplayName ensures that no value is present for ParentDisplayName, not even an explicit nil.
func (o *Cluster) UnsetParentDisplayName() {
	o.ParentDisplayName.Unset()
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetClusterType() ClusterType {
	if o == nil || o.ClusterType.Get() == nil {
		var ret ClusterType
		return ret
	}
	return *o.ClusterType.Get()
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Cluster) GetClusterTypeOk() (*ClusterType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterType.Get(), o.ClusterType.IsSet()
}

// HasClusterType returns a boolean if a field has been set.
func (o *Cluster) HasClusterType() bool {
	return o != nil && o.ClusterType.IsSet()
}

// SetClusterType gets a reference to the given NullableClusterType and assigns it to the ClusterType field.
func (o *Cluster) SetClusterType(v ClusterType) {
	o.ClusterType.Set(&v)
}

// SetClusterTypeNil sets the value for ClusterType to be an explicit nil.
func (o *Cluster) SetClusterTypeNil() {
	o.ClusterType.Set(nil)
}

// UnsetClusterType ensures that no value is present for ClusterType, not even an explicit nil.
func (o *Cluster) UnsetClusterType() {
	o.ClusterType.Unset()
}

// GetDelay returns the Delay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetDelay() float64 {
	if o == nil || o.Delay.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Delay.Get()
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Cluster) GetDelayOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Delay.Get(), o.Delay.IsSet()
}

// HasDelay returns a boolean if a field has been set.
func (o *Cluster) HasDelay() bool {
	return o != nil && o.Delay.IsSet()
}

// SetDelay gets a reference to the given common.NullableFloat64 and assigns it to the Delay field.
func (o *Cluster) SetDelay(v float64) {
	o.Delay.Set(&v)
}

// SetDelayNil sets the value for Delay to be an explicit nil.
func (o *Cluster) SetDelayNil() {
	o.Delay.Set(nil)
}

// UnsetDelay ensures that no value is present for Delay, not even an explicit nil.
func (o *Cluster) UnsetDelay() {
	o.Delay.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Cluster) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Cluster) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Cluster) SetOrgName(v string) {
	o.OrgName = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *Cluster) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *Cluster) HasCloudProvider() bool {
	return o != nil && o.CloudProvider != nil
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *Cluster) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *Cluster) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *Cluster) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *Cluster) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetEnvironmentName returns the EnvironmentName field value.
func (o *Cluster) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *Cluster) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *Cluster) GetEnvironmentType() string {
	if o == nil || o.EnvironmentType == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || o.EnvironmentType == nil {
		return nil, false
	}
	return o.EnvironmentType, true
}

// HasEnvironmentType returns a boolean if a field has been set.
func (o *Cluster) HasEnvironmentType() bool {
	return o != nil && o.EnvironmentType != nil
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *Cluster) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetCloudRegion returns the CloudRegion field value if set, zero value otherwise.
func (o *Cluster) GetCloudRegion() string {
	if o == nil || o.CloudRegion == nil {
		var ret string
		return ret
	}
	return *o.CloudRegion
}

// GetCloudRegionOk returns a tuple with the CloudRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCloudRegionOk() (*string, bool) {
	if o == nil || o.CloudRegion == nil {
		return nil, false
	}
	return o.CloudRegion, true
}

// HasCloudRegion returns a boolean if a field has been set.
func (o *Cluster) HasCloudRegion() bool {
	return o != nil && o.CloudRegion != nil
}

// SetCloudRegion gets a reference to the given string and assigns it to the CloudRegion field.
func (o *Cluster) SetCloudRegion(v string) {
	o.CloudRegion = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
// Deprecated
func (o *Cluster) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Cluster) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Cluster) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
// Deprecated
func (o *Cluster) SetNamespace(v string) {
	o.Namespace = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Cluster) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Cluster) HasProject() bool {
	return o != nil && o.Project != nil
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *Cluster) SetProject(v string) {
	o.Project = &v
}

// GetName returns the Name field value.
func (o *Cluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Cluster) SetName(v string) {
	o.Name = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *Cluster) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *Cluster) HasHash() bool {
	return o != nil && o.Hash != nil
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *Cluster) SetHash(v string) {
	o.Hash = &v
}

// GetEngine returns the Engine field value.
func (o *Cluster) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *Cluster) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *Cluster) SetEngine(v string) {
	o.Engine = v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *Cluster) GetLicense() ClusterLicense {
	if o == nil || o.License == nil {
		var ret ClusterLicense
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetLicenseOk() (*ClusterLicense, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *Cluster) HasLicense() bool {
	return o != nil && o.License != nil
}

// SetLicense gets a reference to the given ClusterLicense and assigns it to the License field.
func (o *Cluster) SetLicense(v ClusterLicense) {
	o.License = &v
}

// GetParamTpls returns the ParamTpls field value if set, zero value otherwise.
func (o *Cluster) GetParamTpls() []ParamTplsItem {
	if o == nil || o.ParamTpls == nil {
		var ret []ParamTplsItem
		return ret
	}
	return o.ParamTpls
}

// GetParamTplsOk returns a tuple with the ParamTpls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetParamTplsOk() (*[]ParamTplsItem, bool) {
	if o == nil || o.ParamTpls == nil {
		return nil, false
	}
	return &o.ParamTpls, true
}

// HasParamTpls returns a boolean if a field has been set.
func (o *Cluster) HasParamTpls() bool {
	return o != nil && o.ParamTpls != nil
}

// SetParamTpls gets a reference to the given []ParamTplsItem and assigns it to the ParamTpls field.
func (o *Cluster) SetParamTpls(v []ParamTplsItem) {
	o.ParamTpls = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Cluster) GetValues() map[string]interface{} {
	if o == nil || o.Values == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetValuesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Cluster) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the Values field.
func (o *Cluster) SetValues(v map[string]interface{}) {
	o.Values = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Cluster) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Cluster) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Cluster) SetVersion(v string) {
	o.Version = &v
}

// GetTerminationPolicy returns the TerminationPolicy field value if set, zero value otherwise.
func (o *Cluster) GetTerminationPolicy() ClusterTerminationPolicy {
	if o == nil || o.TerminationPolicy == nil {
		var ret ClusterTerminationPolicy
		return ret
	}
	return *o.TerminationPolicy
}

// GetTerminationPolicyOk returns a tuple with the TerminationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetTerminationPolicyOk() (*ClusterTerminationPolicy, bool) {
	if o == nil || o.TerminationPolicy == nil {
		return nil, false
	}
	return o.TerminationPolicy, true
}

// HasTerminationPolicy returns a boolean if a field has been set.
func (o *Cluster) HasTerminationPolicy() bool {
	return o != nil && o.TerminationPolicy != nil
}

// SetTerminationPolicy gets a reference to the given ClusterTerminationPolicy and assigns it to the TerminationPolicy field.
func (o *Cluster) SetTerminationPolicy(v ClusterTerminationPolicy) {
	o.TerminationPolicy = &v
}

// GetMonitorEnabled returns the MonitorEnabled field value if set, zero value otherwise.
func (o *Cluster) GetMonitorEnabled() bool {
	if o == nil || o.MonitorEnabled == nil {
		var ret bool
		return ret
	}
	return *o.MonitorEnabled
}

// GetMonitorEnabledOk returns a tuple with the MonitorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetMonitorEnabledOk() (*bool, bool) {
	if o == nil || o.MonitorEnabled == nil {
		return nil, false
	}
	return o.MonitorEnabled, true
}

// HasMonitorEnabled returns a boolean if a field has been set.
func (o *Cluster) HasMonitorEnabled() bool {
	return o != nil && o.MonitorEnabled != nil
}

// SetMonitorEnabled gets a reference to the given bool and assigns it to the MonitorEnabled field.
func (o *Cluster) SetMonitorEnabled(v bool) {
	o.MonitorEnabled = &v
}

// GetVpcEndpointEnabled returns the VpcEndpointEnabled field value if set, zero value otherwise.
func (o *Cluster) GetVpcEndpointEnabled() bool {
	if o == nil || o.VpcEndpointEnabled == nil {
		var ret bool
		return ret
	}
	return *o.VpcEndpointEnabled
}

// GetVpcEndpointEnabledOk returns a tuple with the VpcEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetVpcEndpointEnabledOk() (*bool, bool) {
	if o == nil || o.VpcEndpointEnabled == nil {
		return nil, false
	}
	return o.VpcEndpointEnabled, true
}

// HasVpcEndpointEnabled returns a boolean if a field has been set.
func (o *Cluster) HasVpcEndpointEnabled() bool {
	return o != nil && o.VpcEndpointEnabled != nil
}

// SetVpcEndpointEnabled gets a reference to the given bool and assigns it to the VpcEndpointEnabled field.
func (o *Cluster) SetVpcEndpointEnabled(v bool) {
	o.VpcEndpointEnabled = &v
}

// GetInternetEndpointEnabled returns the InternetEndpointEnabled field value if set, zero value otherwise.
func (o *Cluster) GetInternetEndpointEnabled() bool {
	if o == nil || o.InternetEndpointEnabled == nil {
		var ret bool
		return ret
	}
	return *o.InternetEndpointEnabled
}

// GetInternetEndpointEnabledOk returns a tuple with the InternetEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetInternetEndpointEnabledOk() (*bool, bool) {
	if o == nil || o.InternetEndpointEnabled == nil {
		return nil, false
	}
	return o.InternetEndpointEnabled, true
}

// HasInternetEndpointEnabled returns a boolean if a field has been set.
func (o *Cluster) HasInternetEndpointEnabled() bool {
	return o != nil && o.InternetEndpointEnabled != nil
}

// SetInternetEndpointEnabled gets a reference to the given bool and assigns it to the InternetEndpointEnabled field.
func (o *Cluster) SetInternetEndpointEnabled(v bool) {
	o.InternetEndpointEnabled = &v
}

// GetTlsEnabled returns the TlsEnabled field value if set, zero value otherwise.
func (o *Cluster) GetTlsEnabled() bool {
	if o == nil || o.TlsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TlsEnabled
}

// GetTlsEnabledOk returns a tuple with the TlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetTlsEnabledOk() (*bool, bool) {
	if o == nil || o.TlsEnabled == nil {
		return nil, false
	}
	return o.TlsEnabled, true
}

// HasTlsEnabled returns a boolean if a field has been set.
func (o *Cluster) HasTlsEnabled() bool {
	return o != nil && o.TlsEnabled != nil
}

// SetTlsEnabled gets a reference to the given bool and assigns it to the TlsEnabled field.
func (o *Cluster) SetTlsEnabled(v bool) {
	o.TlsEnabled = &v
}

// GetNodePortEnabled returns the NodePortEnabled field value if set, zero value otherwise.
func (o *Cluster) GetNodePortEnabled() bool {
	if o == nil || o.NodePortEnabled == nil {
		var ret bool
		return ret
	}
	return *o.NodePortEnabled
}

// GetNodePortEnabledOk returns a tuple with the NodePortEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetNodePortEnabledOk() (*bool, bool) {
	if o == nil || o.NodePortEnabled == nil {
		return nil, false
	}
	return o.NodePortEnabled, true
}

// HasNodePortEnabled returns a boolean if a field has been set.
func (o *Cluster) HasNodePortEnabled() bool {
	return o != nil && o.NodePortEnabled != nil
}

// SetNodePortEnabled gets a reference to the given bool and assigns it to the NodePortEnabled field.
func (o *Cluster) SetNodePortEnabled(v bool) {
	o.NodePortEnabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Cluster) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Cluster) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Cluster) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Cluster) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Cluster) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Cluster) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Cluster) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Cluster) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Cluster) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreateOnlySet returns the CreateOnlySet field value if set, zero value otherwise.
func (o *Cluster) GetCreateOnlySet() bool {
	if o == nil || o.CreateOnlySet == nil {
		var ret bool
		return ret
	}
	return *o.CreateOnlySet
}

// GetCreateOnlySetOk returns a tuple with the CreateOnlySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCreateOnlySetOk() (*bool, bool) {
	if o == nil || o.CreateOnlySet == nil {
		return nil, false
	}
	return o.CreateOnlySet, true
}

// HasCreateOnlySet returns a boolean if a field has been set.
func (o *Cluster) HasCreateOnlySet() bool {
	return o != nil && o.CreateOnlySet != nil
}

// SetCreateOnlySet gets a reference to the given bool and assigns it to the CreateOnlySet field.
func (o *Cluster) SetCreateOnlySet(v bool) {
	o.CreateOnlySet = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *Cluster) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *Cluster) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *Cluster) SetMode(v string) {
	o.Mode = &v
}

// GetProxyEnabled returns the ProxyEnabled field value if set, zero value otherwise.
func (o *Cluster) GetProxyEnabled() bool {
	if o == nil || o.ProxyEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ProxyEnabled
}

// GetProxyEnabledOk returns a tuple with the ProxyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetProxyEnabledOk() (*bool, bool) {
	if o == nil || o.ProxyEnabled == nil {
		return nil, false
	}
	return o.ProxyEnabled, true
}

// HasProxyEnabled returns a boolean if a field has been set.
func (o *Cluster) HasProxyEnabled() bool {
	return o != nil && o.ProxyEnabled != nil
}

// SetProxyEnabled gets a reference to the given bool and assigns it to the ProxyEnabled field.
func (o *Cluster) SetProxyEnabled(v bool) {
	o.ProxyEnabled = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *Cluster) GetComponents() []ComponentsItem {
	if o == nil || o.Components == nil {
		var ret []ComponentsItem
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetComponentsOk() (*[]ComponentsItem, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *Cluster) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []ComponentsItem and assigns it to the Components field.
func (o *Cluster) SetComponents(v []ComponentsItem) {
	o.Components = v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Cluster) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return &o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Cluster) HasExtra() bool {
	return o != nil && o.Extra != nil
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *Cluster) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// GetInitOptions returns the InitOptions field value if set, zero value otherwise.
func (o *Cluster) GetInitOptions() []InitOptionsItem {
	if o == nil || o.InitOptions == nil {
		var ret []InitOptionsItem
		return ret
	}
	return o.InitOptions
}

// GetInitOptionsOk returns a tuple with the InitOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetInitOptionsOk() (*[]InitOptionsItem, bool) {
	if o == nil || o.InitOptions == nil {
		return nil, false
	}
	return &o.InitOptions, true
}

// HasInitOptions returns a boolean if a field has been set.
func (o *Cluster) HasInitOptions() bool {
	return o != nil && o.InitOptions != nil
}

// SetInitOptions gets a reference to the given []InitOptionsItem and assigns it to the InitOptions field.
func (o *Cluster) SetInitOptions(v []InitOptionsItem) {
	o.InitOptions = v
}

// GetExtraInfo returns the ExtraInfo field value if set, zero value otherwise.
func (o *Cluster) GetExtraInfo() map[string]interface{} {
	if o == nil || o.ExtraInfo == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraInfo
}

// GetExtraInfoOk returns a tuple with the ExtraInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetExtraInfoOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExtraInfo == nil {
		return nil, false
	}
	return &o.ExtraInfo, true
}

// HasExtraInfo returns a boolean if a field has been set.
func (o *Cluster) HasExtraInfo() bool {
	return o != nil && o.ExtraInfo != nil
}

// SetExtraInfo gets a reference to the given map[string]interface{} and assigns it to the ExtraInfo field.
func (o *Cluster) SetExtraInfo(v map[string]interface{}) {
	o.ExtraInfo = v
}

// GetTolerations returns the Tolerations field value if set, zero value otherwise.
func (o *Cluster) GetTolerations() string {
	if o == nil || o.Tolerations == nil {
		var ret string
		return ret
	}
	return *o.Tolerations
}

// GetTolerationsOk returns a tuple with the Tolerations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetTolerationsOk() (*string, bool) {
	if o == nil || o.Tolerations == nil {
		return nil, false
	}
	return o.Tolerations, true
}

// HasTolerations returns a boolean if a field has been set.
func (o *Cluster) HasTolerations() bool {
	return o != nil && o.Tolerations != nil
}

// SetTolerations gets a reference to the given string and assigns it to the Tolerations field.
func (o *Cluster) SetTolerations(v string) {
	o.Tolerations = &v
}

// GetSingleZone returns the SingleZone field value if set, zero value otherwise.
func (o *Cluster) GetSingleZone() bool {
	if o == nil || o.SingleZone == nil {
		var ret bool
		return ret
	}
	return *o.SingleZone
}

// GetSingleZoneOk returns a tuple with the SingleZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetSingleZoneOk() (*bool, bool) {
	if o == nil || o.SingleZone == nil {
		return nil, false
	}
	return o.SingleZone, true
}

// HasSingleZone returns a boolean if a field has been set.
func (o *Cluster) HasSingleZone() bool {
	return o != nil && o.SingleZone != nil
}

// SetSingleZone gets a reference to the given bool and assigns it to the SingleZone field.
func (o *Cluster) SetSingleZone(v bool) {
	o.SingleZone = &v
}

// GetAvailabilityZones returns the AvailabilityZones field value if set, zero value otherwise.
func (o *Cluster) GetAvailabilityZones() []string {
	if o == nil || o.AvailabilityZones == nil {
		var ret []string
		return ret
	}
	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetAvailabilityZonesOk() (*[]string, bool) {
	if o == nil || o.AvailabilityZones == nil {
		return nil, false
	}
	return &o.AvailabilityZones, true
}

// HasAvailabilityZones returns a boolean if a field has been set.
func (o *Cluster) HasAvailabilityZones() bool {
	return o != nil && o.AvailabilityZones != nil
}

// SetAvailabilityZones gets a reference to the given []string and assigns it to the AvailabilityZones field.
func (o *Cluster) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = v
}

// GetPodAntiAffinityEnabled returns the PodAntiAffinityEnabled field value if set, zero value otherwise.
func (o *Cluster) GetPodAntiAffinityEnabled() bool {
	if o == nil || o.PodAntiAffinityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PodAntiAffinityEnabled
}

// GetPodAntiAffinityEnabledOk returns a tuple with the PodAntiAffinityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetPodAntiAffinityEnabledOk() (*bool, bool) {
	if o == nil || o.PodAntiAffinityEnabled == nil {
		return nil, false
	}
	return o.PodAntiAffinityEnabled, true
}

// HasPodAntiAffinityEnabled returns a boolean if a field has been set.
func (o *Cluster) HasPodAntiAffinityEnabled() bool {
	return o != nil && o.PodAntiAffinityEnabled != nil
}

// SetPodAntiAffinityEnabled gets a reference to the given bool and assigns it to the PodAntiAffinityEnabled field.
func (o *Cluster) SetPodAntiAffinityEnabled(v bool) {
	o.PodAntiAffinityEnabled = &v
}

// GetBackup returns the Backup field value if set, zero value otherwise.
func (o *Cluster) GetBackup() ClusterBackup {
	if o == nil || o.Backup == nil {
		var ret ClusterBackup
		return ret
	}
	return *o.Backup
}

// GetBackupOk returns a tuple with the Backup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetBackupOk() (*ClusterBackup, bool) {
	if o == nil || o.Backup == nil {
		return nil, false
	}
	return o.Backup, true
}

// HasBackup returns a boolean if a field has been set.
func (o *Cluster) HasBackup() bool {
	return o != nil && o.Backup != nil
}

// SetBackup gets a reference to the given ClusterBackup and assigns it to the Backup field.
func (o *Cluster) SetBackup(v ClusterBackup) {
	o.Backup = &v
}

// GetNodeGroup returns the NodeGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cluster) GetNodeGroup() string {
	if o == nil || o.NodeGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeGroup.Get()
}

// GetNodeGroupOk returns a tuple with the NodeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *Cluster) GetNodeGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeGroup.Get(), o.NodeGroup.IsSet()
}

// HasNodeGroup returns a boolean if a field has been set.
func (o *Cluster) HasNodeGroup() bool {
	return o != nil && o.NodeGroup.IsSet()
}

// SetNodeGroup gets a reference to the given common.NullableString and assigns it to the NodeGroup field.
func (o *Cluster) SetNodeGroup(v string) {
	o.NodeGroup.Set(&v)
}

// SetNodeGroupNil sets the value for NodeGroup to be an explicit nil.
func (o *Cluster) SetNodeGroupNil() {
	o.NodeGroup.Set(nil)
}

// UnsetNodeGroup ensures that no value is present for NodeGroup, not even an explicit nil.
func (o *Cluster) UnsetNodeGroup() {
	o.NodeGroup.Unset()
}

// GetCodeShort returns the CodeShort field value if set, zero value otherwise.
func (o *Cluster) GetCodeShort() string {
	if o == nil || o.CodeShort == nil {
		var ret string
		return ret
	}
	return *o.CodeShort
}

// GetCodeShortOk returns a tuple with the CodeShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetCodeShortOk() (*string, bool) {
	if o == nil || o.CodeShort == nil {
		return nil, false
	}
	return o.CodeShort, true
}

// HasCodeShort returns a boolean if a field has been set.
func (o *Cluster) HasCodeShort() bool {
	return o != nil && o.CodeShort != nil
}

// SetCodeShort gets a reference to the given string and assigns it to the CodeShort field.
func (o *Cluster) SetCodeShort(v string) {
	o.CodeShort = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Cluster) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Cluster) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Cluster) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetStatic returns the Static field value if set, zero value otherwise.
func (o *Cluster) GetStatic() bool {
	if o == nil || o.Static == nil {
		var ret bool
		return ret
	}
	return *o.Static
}

// GetStaticOk returns a tuple with the Static field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cluster) GetStaticOk() (*bool, bool) {
	if o == nil || o.Static == nil {
		return nil, false
	}
	return o.Static, true
}

// HasStatic returns a boolean if a field has been set.
func (o *Cluster) HasStatic() bool {
	return o != nil && o.Static != nil
}

// SetStatic gets a reference to the given bool and assigns it to the Static field.
func (o *Cluster) SetStatic(v bool) {
	o.Static = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Cluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if o.ParentName.IsSet() {
		toSerialize["parentName"] = o.ParentName.Get()
	}
	if o.ParentDisplayName.IsSet() {
		toSerialize["parentDisplayName"] = o.ParentDisplayName.Get()
	}
	if o.ClusterType.IsSet() {
		toSerialize["clusterType"] = o.ClusterType.Get()
	}
	if o.Delay.IsSet() {
		toSerialize["delay"] = o.Delay.Get()
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.CloudProvider != nil {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if o.EnvironmentId != nil {
		toSerialize["environmentId"] = o.EnvironmentId
	}
	toSerialize["environmentName"] = o.EnvironmentName
	if o.EnvironmentType != nil {
		toSerialize["environmentType"] = o.EnvironmentType
	}
	if o.CloudRegion != nil {
		toSerialize["cloudRegion"] = o.CloudRegion
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	toSerialize["name"] = o.Name
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	toSerialize["engine"] = o.Engine
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.ParamTpls != nil {
		toSerialize["paramTpls"] = o.ParamTpls
	}
	if o.Values != nil {
		toSerialize["Values"] = o.Values
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.TerminationPolicy != nil {
		toSerialize["terminationPolicy"] = o.TerminationPolicy
	}
	if o.MonitorEnabled != nil {
		toSerialize["monitorEnabled"] = o.MonitorEnabled
	}
	if o.VpcEndpointEnabled != nil {
		toSerialize["vpcEndpointEnabled"] = o.VpcEndpointEnabled
	}
	if o.InternetEndpointEnabled != nil {
		toSerialize["internetEndpointEnabled"] = o.InternetEndpointEnabled
	}
	if o.TlsEnabled != nil {
		toSerialize["tlsEnabled"] = o.TlsEnabled
	}
	if o.NodePortEnabled != nil {
		toSerialize["nodePortEnabled"] = o.NodePortEnabled
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreateOnlySet != nil {
		toSerialize["createOnlySet"] = o.CreateOnlySet
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.ProxyEnabled != nil {
		toSerialize["proxyEnabled"] = o.ProxyEnabled
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.InitOptions != nil {
		toSerialize["initOptions"] = o.InitOptions
	}
	if o.ExtraInfo != nil {
		toSerialize["extraInfo"] = o.ExtraInfo
	}
	if o.Tolerations != nil {
		toSerialize["tolerations"] = o.Tolerations
	}
	if o.SingleZone != nil {
		toSerialize["singleZone"] = o.SingleZone
	}
	if o.AvailabilityZones != nil {
		toSerialize["availabilityZones"] = o.AvailabilityZones
	}
	if o.PodAntiAffinityEnabled != nil {
		toSerialize["podAntiAffinityEnabled"] = o.PodAntiAffinityEnabled
	}
	if o.Backup != nil {
		toSerialize["backup"] = o.Backup
	}
	if o.NodeGroup.IsSet() {
		toSerialize["nodeGroup"] = o.NodeGroup.Get()
	}
	if o.CodeShort != nil {
		toSerialize["codeShort"] = o.CodeShort
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Static != nil {
		toSerialize["static"] = o.Static
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Cluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                      interface{}               `json:"id,omitempty"`
		ParentId                common.NullableInt64      `json:"parentId,omitempty"`
		ParentName              common.NullableString     `json:"parentName,omitempty"`
		ParentDisplayName       common.NullableString     `json:"parentDisplayName,omitempty"`
		ClusterType             NullableClusterType       `json:"clusterType,omitempty"`
		Delay                   common.NullableFloat64    `json:"delay,omitempty"`
		OrgName                 *string                   `json:"orgName,omitempty"`
		CloudProvider           *string                   `json:"cloudProvider,omitempty"`
		EnvironmentId           *string                   `json:"environmentId,omitempty"`
		EnvironmentName         *string                   `json:"environmentName"`
		EnvironmentType         *string                   `json:"environmentType,omitempty"`
		CloudRegion             *string                   `json:"cloudRegion,omitempty"`
		Namespace               *string                   `json:"namespace,omitempty"`
		Project                 *string                   `json:"project,omitempty"`
		Name                    *string                   `json:"name"`
		Hash                    *string                   `json:"hash,omitempty"`
		Engine                  *string                   `json:"engine"`
		License                 *ClusterLicense           `json:"license,omitempty"`
		ParamTpls               []ParamTplsItem           `json:"paramTpls,omitempty"`
		Values                  map[string]interface{}    `json:"Values,omitempty"`
		Version                 *string                   `json:"version,omitempty"`
		TerminationPolicy       *ClusterTerminationPolicy `json:"terminationPolicy,omitempty"`
		MonitorEnabled          *bool                     `json:"monitorEnabled,omitempty"`
		VpcEndpointEnabled      *bool                     `json:"vpcEndpointEnabled,omitempty"`
		InternetEndpointEnabled *bool                     `json:"internetEndpointEnabled,omitempty"`
		TlsEnabled              *bool                     `json:"tlsEnabled,omitempty"`
		NodePortEnabled         *bool                     `json:"nodePortEnabled,omitempty"`
		Status                  *string                   `json:"status,omitempty"`
		CreatedAt               *time.Time                `json:"createdAt,omitempty"`
		UpdatedAt               *time.Time                `json:"updatedAt,omitempty"`
		CreateOnlySet           *bool                     `json:"createOnlySet,omitempty"`
		Mode                    *string                   `json:"mode,omitempty"`
		ProxyEnabled            *bool                     `json:"proxyEnabled,omitempty"`
		Components              []ComponentsItem          `json:"components,omitempty"`
		Extra                   map[string]interface{}    `json:"extra,omitempty"`
		InitOptions             []InitOptionsItem         `json:"initOptions,omitempty"`
		ExtraInfo               map[string]interface{}    `json:"extraInfo,omitempty"`
		Tolerations             *string                   `json:"tolerations,omitempty"`
		SingleZone              *bool                     `json:"singleZone,omitempty"`
		AvailabilityZones       []string                  `json:"availabilityZones,omitempty"`
		PodAntiAffinityEnabled  *bool                     `json:"podAntiAffinityEnabled,omitempty"`
		Backup                  *ClusterBackup            `json:"backup,omitempty"`
		NodeGroup               common.NullableString     `json:"nodeGroup,omitempty"`
		CodeShort               *string                   `json:"codeShort,omitempty"`
		DisplayName             *string                   `json:"displayName,omitempty"`
		Static                  *bool                     `json:"static,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EnvironmentName == nil {
		return fmt.Errorf("required field environmentName missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "parentId", "parentName", "parentDisplayName", "clusterType", "delay", "orgName", "cloudProvider", "environmentId", "environmentName", "environmentType", "cloudRegion", "namespace", "project", "name", "hash", "engine", "license", "paramTpls", "Values", "version", "terminationPolicy", "monitorEnabled", "vpcEndpointEnabled", "internetEndpointEnabled", "tlsEnabled", "nodePortEnabled", "status", "createdAt", "updatedAt", "createOnlySet", "mode", "proxyEnabled", "components", "extra", "initOptions", "extraInfo", "tolerations", "singleZone", "availabilityZones", "podAntiAffinityEnabled", "backup", "nodeGroup", "codeShort", "displayName", "static"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.ParentId = all.ParentId
	o.ParentName = all.ParentName
	o.ParentDisplayName = all.ParentDisplayName
	if all.ClusterType.Get() != nil && !all.ClusterType.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.ClusterType = all.ClusterType
	}
	o.Delay = all.Delay
	o.OrgName = all.OrgName
	o.CloudProvider = all.CloudProvider
	o.EnvironmentId = all.EnvironmentId
	o.EnvironmentName = *all.EnvironmentName
	o.EnvironmentType = all.EnvironmentType
	o.CloudRegion = all.CloudRegion
	o.Namespace = all.Namespace
	o.Project = all.Project
	o.Name = *all.Name
	o.Hash = all.Hash
	o.Engine = *all.Engine
	if all.License != nil && all.License.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.License = all.License
	o.ParamTpls = all.ParamTpls
	o.Values = all.Values
	o.Version = all.Version
	if all.TerminationPolicy != nil && !all.TerminationPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.TerminationPolicy = all.TerminationPolicy
	}
	o.MonitorEnabled = all.MonitorEnabled
	o.VpcEndpointEnabled = all.VpcEndpointEnabled
	o.InternetEndpointEnabled = all.InternetEndpointEnabled
	o.TlsEnabled = all.TlsEnabled
	o.NodePortEnabled = all.NodePortEnabled
	o.Status = all.Status
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	o.CreateOnlySet = all.CreateOnlySet
	o.Mode = all.Mode
	o.ProxyEnabled = all.ProxyEnabled
	o.Components = all.Components
	o.Extra = all.Extra
	o.InitOptions = all.InitOptions
	o.ExtraInfo = all.ExtraInfo
	o.Tolerations = all.Tolerations
	o.SingleZone = all.SingleZone
	o.AvailabilityZones = all.AvailabilityZones
	o.PodAntiAffinityEnabled = all.PodAntiAffinityEnabled
	if all.Backup != nil && all.Backup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Backup = all.Backup
	o.NodeGroup = all.NodeGroup
	o.CodeShort = all.CodeShort
	o.DisplayName = all.DisplayName
	o.Static = all.Static

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
