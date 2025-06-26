// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterCreate KubeBlocks cluster information
type ClusterCreate struct {
	// When two clusters have a relationship, parentId records the parent cluster id.Can be empty when there is no relationship
	ParentId common.NullableString `json:"parentId,omitempty"`
	// Describes the type of cluster, [Normal] normal cluster; [DisasterRecovery] disaster recovery cluster
	ClusterType NullableClusterType `json:"clusterType,omitempty"`
	// Org Name
	OrgName *string `json:"orgName,omitempty"`
	// Environment Name
	EnvironmentName string `json:"environmentName"`
	// Name of project, it is the alias of environment namespace
	Project *string `json:"project,omitempty"`
	// Name of cluster. Name must be unique within an Org
	Name string `json:"name"`
	// Cluster Application Engine
	Engine  string          `json:"engine"`
	License *ClusterLicense `json:"license,omitempty"`
	// Items is the list of parameter template in the list
	ParamTpls []ParamTplsItem `json:"paramTpls,omitempty"`
	// Cluster Application Version
	Version *string `json:"version,omitempty"`
	// The termination policy of cluster.
	TerminationPolicy *ClusterTerminationPolicy `json:"terminationPolicy,omitempty"`
	// Cluster topology mode
	Mode *string `json:"mode,omitempty"`
	// Components is the list of components
	Components []ComponentItemCreate `json:"components,omitempty"`
	// Extra configuration for cluster
	Extra map[string]interface{} `json:"extra,omitempty"`
	// InitOptions is the list of init option
	InitOptions []InitOptionItem `json:"initOptions,omitempty"`
	// Use single availability zones
	SingleZone *bool `json:"singleZone,omitempty"`
	// Availability Zones
	AvailabilityZones []string `json:"availabilityZones,omitempty"`
	// clusterBackup is the payload for cluster backup
	Backup *ClusterBackup `json:"backup,omitempty"`
	// Specify a NodeGroup for the cluster
	NodeGroup common.NullableString `json:"nodeGroup,omitempty"`
	// Display name of cluster.
	DisplayName *string `json:"displayName,omitempty"`
	// if cluster is static cluster
	Static      *bool        `json:"static,omitempty"`
	NetworkMode *NetworkMode `json:"networkMode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterCreate instantiates a new ClusterCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterCreate(environmentName string, name string, engine string) *ClusterCreate {
	this := ClusterCreate{}
	var clusterType ClusterType = ClusterTypeNormal
	this.ClusterType = *NewNullableClusterType(&clusterType)
	this.EnvironmentName = environmentName
	var project string = "kubeblocks-cloud-ns"
	this.Project = &project
	this.Name = name
	this.Engine = engine
	var terminationPolicy ClusterTerminationPolicy = ClusterTerminationPolicyDelete
	this.TerminationPolicy = &terminationPolicy
	var singleZone bool = false
	this.SingleZone = &singleZone
	return &this
}

// NewClusterCreateWithDefaults instantiates a new ClusterCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterCreateWithDefaults() *ClusterCreate {
	this := ClusterCreate{}
	var clusterType ClusterType = ClusterTypeNormal
	this.ClusterType = *NewNullableClusterType(&clusterType)
	var project string = "kubeblocks-cloud-ns"
	this.Project = &project
	var terminationPolicy ClusterTerminationPolicy = ClusterTerminationPolicyDelete
	this.TerminationPolicy = &terminationPolicy
	var singleZone bool = false
	this.SingleZone = &singleZone
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterCreate) GetParentId() string {
	if o == nil || o.ParentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterCreate) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *ClusterCreate) HasParentId() bool {
	return o != nil && o.ParentId.IsSet()
}

// SetParentId gets a reference to the given common.NullableString and assigns it to the ParentId field.
func (o *ClusterCreate) SetParentId(v string) {
	o.ParentId.Set(&v)
}

// SetParentIdNil sets the value for ParentId to be an explicit nil.
func (o *ClusterCreate) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil.
func (o *ClusterCreate) UnsetParentId() {
	o.ParentId.Unset()
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterCreate) GetClusterType() ClusterType {
	if o == nil || o.ClusterType.Get() == nil {
		var ret ClusterType
		return ret
	}
	return *o.ClusterType.Get()
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterCreate) GetClusterTypeOk() (*ClusterType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterType.Get(), o.ClusterType.IsSet()
}

// HasClusterType returns a boolean if a field has been set.
func (o *ClusterCreate) HasClusterType() bool {
	return o != nil && o.ClusterType.IsSet()
}

// SetClusterType gets a reference to the given NullableClusterType and assigns it to the ClusterType field.
func (o *ClusterCreate) SetClusterType(v ClusterType) {
	o.ClusterType.Set(&v)
}

// SetClusterTypeNil sets the value for ClusterType to be an explicit nil.
func (o *ClusterCreate) SetClusterTypeNil() {
	o.ClusterType.Set(nil)
}

// UnsetClusterType ensures that no value is present for ClusterType, not even an explicit nil.
func (o *ClusterCreate) UnsetClusterType() {
	o.ClusterType.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *ClusterCreate) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *ClusterCreate) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *ClusterCreate) SetOrgName(v string) {
	o.OrgName = &v
}

// GetEnvironmentName returns the EnvironmentName field value.
func (o *ClusterCreate) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *ClusterCreate) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ClusterCreate) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ClusterCreate) HasProject() bool {
	return o != nil && o.Project != nil
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *ClusterCreate) SetProject(v string) {
	o.Project = &v
}

// GetName returns the Name field value.
func (o *ClusterCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ClusterCreate) SetName(v string) {
	o.Name = v
}

// GetEngine returns the Engine field value.
func (o *ClusterCreate) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *ClusterCreate) SetEngine(v string) {
	o.Engine = v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *ClusterCreate) GetLicense() ClusterLicense {
	if o == nil || o.License == nil {
		var ret ClusterLicense
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetLicenseOk() (*ClusterLicense, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *ClusterCreate) HasLicense() bool {
	return o != nil && o.License != nil
}

// SetLicense gets a reference to the given ClusterLicense and assigns it to the License field.
func (o *ClusterCreate) SetLicense(v ClusterLicense) {
	o.License = &v
}

// GetParamTpls returns the ParamTpls field value if set, zero value otherwise.
func (o *ClusterCreate) GetParamTpls() []ParamTplsItem {
	if o == nil || o.ParamTpls == nil {
		var ret []ParamTplsItem
		return ret
	}
	return o.ParamTpls
}

// GetParamTplsOk returns a tuple with the ParamTpls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetParamTplsOk() (*[]ParamTplsItem, bool) {
	if o == nil || o.ParamTpls == nil {
		return nil, false
	}
	return &o.ParamTpls, true
}

// HasParamTpls returns a boolean if a field has been set.
func (o *ClusterCreate) HasParamTpls() bool {
	return o != nil && o.ParamTpls != nil
}

// SetParamTpls gets a reference to the given []ParamTplsItem and assigns it to the ParamTpls field.
func (o *ClusterCreate) SetParamTpls(v []ParamTplsItem) {
	o.ParamTpls = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ClusterCreate) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ClusterCreate) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ClusterCreate) SetVersion(v string) {
	o.Version = &v
}

// GetTerminationPolicy returns the TerminationPolicy field value if set, zero value otherwise.
func (o *ClusterCreate) GetTerminationPolicy() ClusterTerminationPolicy {
	if o == nil || o.TerminationPolicy == nil {
		var ret ClusterTerminationPolicy
		return ret
	}
	return *o.TerminationPolicy
}

// GetTerminationPolicyOk returns a tuple with the TerminationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetTerminationPolicyOk() (*ClusterTerminationPolicy, bool) {
	if o == nil || o.TerminationPolicy == nil {
		return nil, false
	}
	return o.TerminationPolicy, true
}

// HasTerminationPolicy returns a boolean if a field has been set.
func (o *ClusterCreate) HasTerminationPolicy() bool {
	return o != nil && o.TerminationPolicy != nil
}

// SetTerminationPolicy gets a reference to the given ClusterTerminationPolicy and assigns it to the TerminationPolicy field.
func (o *ClusterCreate) SetTerminationPolicy(v ClusterTerminationPolicy) {
	o.TerminationPolicy = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ClusterCreate) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ClusterCreate) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ClusterCreate) SetMode(v string) {
	o.Mode = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ClusterCreate) GetComponents() []ComponentItemCreate {
	if o == nil || o.Components == nil {
		var ret []ComponentItemCreate
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetComponentsOk() (*[]ComponentItemCreate, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ClusterCreate) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []ComponentItemCreate and assigns it to the Components field.
func (o *ClusterCreate) SetComponents(v []ComponentItemCreate) {
	o.Components = v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *ClusterCreate) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return &o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *ClusterCreate) HasExtra() bool {
	return o != nil && o.Extra != nil
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *ClusterCreate) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// GetInitOptions returns the InitOptions field value if set, zero value otherwise.
func (o *ClusterCreate) GetInitOptions() []InitOptionItem {
	if o == nil || o.InitOptions == nil {
		var ret []InitOptionItem
		return ret
	}
	return o.InitOptions
}

// GetInitOptionsOk returns a tuple with the InitOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetInitOptionsOk() (*[]InitOptionItem, bool) {
	if o == nil || o.InitOptions == nil {
		return nil, false
	}
	return &o.InitOptions, true
}

// HasInitOptions returns a boolean if a field has been set.
func (o *ClusterCreate) HasInitOptions() bool {
	return o != nil && o.InitOptions != nil
}

// SetInitOptions gets a reference to the given []InitOptionItem and assigns it to the InitOptions field.
func (o *ClusterCreate) SetInitOptions(v []InitOptionItem) {
	o.InitOptions = v
}

// GetSingleZone returns the SingleZone field value if set, zero value otherwise.
func (o *ClusterCreate) GetSingleZone() bool {
	if o == nil || o.SingleZone == nil {
		var ret bool
		return ret
	}
	return *o.SingleZone
}

// GetSingleZoneOk returns a tuple with the SingleZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetSingleZoneOk() (*bool, bool) {
	if o == nil || o.SingleZone == nil {
		return nil, false
	}
	return o.SingleZone, true
}

// HasSingleZone returns a boolean if a field has been set.
func (o *ClusterCreate) HasSingleZone() bool {
	return o != nil && o.SingleZone != nil
}

// SetSingleZone gets a reference to the given bool and assigns it to the SingleZone field.
func (o *ClusterCreate) SetSingleZone(v bool) {
	o.SingleZone = &v
}

// GetAvailabilityZones returns the AvailabilityZones field value if set, zero value otherwise.
func (o *ClusterCreate) GetAvailabilityZones() []string {
	if o == nil || o.AvailabilityZones == nil {
		var ret []string
		return ret
	}
	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetAvailabilityZonesOk() (*[]string, bool) {
	if o == nil || o.AvailabilityZones == nil {
		return nil, false
	}
	return &o.AvailabilityZones, true
}

// HasAvailabilityZones returns a boolean if a field has been set.
func (o *ClusterCreate) HasAvailabilityZones() bool {
	return o != nil && o.AvailabilityZones != nil
}

// SetAvailabilityZones gets a reference to the given []string and assigns it to the AvailabilityZones field.
func (o *ClusterCreate) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = v
}

// GetBackup returns the Backup field value if set, zero value otherwise.
func (o *ClusterCreate) GetBackup() ClusterBackup {
	if o == nil || o.Backup == nil {
		var ret ClusterBackup
		return ret
	}
	return *o.Backup
}

// GetBackupOk returns a tuple with the Backup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetBackupOk() (*ClusterBackup, bool) {
	if o == nil || o.Backup == nil {
		return nil, false
	}
	return o.Backup, true
}

// HasBackup returns a boolean if a field has been set.
func (o *ClusterCreate) HasBackup() bool {
	return o != nil && o.Backup != nil
}

// SetBackup gets a reference to the given ClusterBackup and assigns it to the Backup field.
func (o *ClusterCreate) SetBackup(v ClusterBackup) {
	o.Backup = &v
}

// GetNodeGroup returns the NodeGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterCreate) GetNodeGroup() string {
	if o == nil || o.NodeGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeGroup.Get()
}

// GetNodeGroupOk returns a tuple with the NodeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterCreate) GetNodeGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeGroup.Get(), o.NodeGroup.IsSet()
}

// HasNodeGroup returns a boolean if a field has been set.
func (o *ClusterCreate) HasNodeGroup() bool {
	return o != nil && o.NodeGroup.IsSet()
}

// SetNodeGroup gets a reference to the given common.NullableString and assigns it to the NodeGroup field.
func (o *ClusterCreate) SetNodeGroup(v string) {
	o.NodeGroup.Set(&v)
}

// SetNodeGroupNil sets the value for NodeGroup to be an explicit nil.
func (o *ClusterCreate) SetNodeGroupNil() {
	o.NodeGroup.Set(nil)
}

// UnsetNodeGroup ensures that no value is present for NodeGroup, not even an explicit nil.
func (o *ClusterCreate) UnsetNodeGroup() {
	o.NodeGroup.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ClusterCreate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterCreate) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ClusterCreate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetStatic returns the Static field value if set, zero value otherwise.
func (o *ClusterCreate) GetStatic() bool {
	if o == nil || o.Static == nil {
		var ret bool
		return ret
	}
	return *o.Static
}

// GetStaticOk returns a tuple with the Static field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetStaticOk() (*bool, bool) {
	if o == nil || o.Static == nil {
		return nil, false
	}
	return o.Static, true
}

// HasStatic returns a boolean if a field has been set.
func (o *ClusterCreate) HasStatic() bool {
	return o != nil && o.Static != nil
}

// SetStatic gets a reference to the given bool and assigns it to the Static field.
func (o *ClusterCreate) SetStatic(v bool) {
	o.Static = &v
}

// GetNetworkMode returns the NetworkMode field value if set, zero value otherwise.
func (o *ClusterCreate) GetNetworkMode() NetworkMode {
	if o == nil || o.NetworkMode == nil {
		var ret NetworkMode
		return ret
	}
	return *o.NetworkMode
}

// GetNetworkModeOk returns a tuple with the NetworkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCreate) GetNetworkModeOk() (*NetworkMode, bool) {
	if o == nil || o.NetworkMode == nil {
		return nil, false
	}
	return o.NetworkMode, true
}

// HasNetworkMode returns a boolean if a field has been set.
func (o *ClusterCreate) HasNetworkMode() bool {
	return o != nil && o.NetworkMode != nil
}

// SetNetworkMode gets a reference to the given NetworkMode and assigns it to the NetworkMode field.
func (o *ClusterCreate) SetNetworkMode(v NetworkMode) {
	o.NetworkMode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ParentId.IsSet() {
		toSerialize["parentId"] = o.ParentId.Get()
	}
	if o.ClusterType.IsSet() {
		toSerialize["clusterType"] = o.ClusterType.Get()
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	toSerialize["environmentName"] = o.EnvironmentName
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	toSerialize["name"] = o.Name
	toSerialize["engine"] = o.Engine
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.ParamTpls != nil {
		toSerialize["paramTpls"] = o.ParamTpls
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.TerminationPolicy != nil {
		toSerialize["terminationPolicy"] = o.TerminationPolicy
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
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
	if o.SingleZone != nil {
		toSerialize["singleZone"] = o.SingleZone
	}
	if o.AvailabilityZones != nil {
		toSerialize["availabilityZones"] = o.AvailabilityZones
	}
	if o.Backup != nil {
		toSerialize["backup"] = o.Backup
	}
	if o.NodeGroup.IsSet() {
		toSerialize["nodeGroup"] = o.NodeGroup.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Static != nil {
		toSerialize["static"] = o.Static
	}
	if o.NetworkMode != nil {
		toSerialize["networkMode"] = o.NetworkMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ParentId          common.NullableString     `json:"parentId,omitempty"`
		ClusterType       NullableClusterType       `json:"clusterType,omitempty"`
		OrgName           *string                   `json:"orgName,omitempty"`
		EnvironmentName   *string                   `json:"environmentName"`
		Project           *string                   `json:"project,omitempty"`
		Name              *string                   `json:"name"`
		Engine            *string                   `json:"engine"`
		License           *ClusterLicense           `json:"license,omitempty"`
		ParamTpls         []ParamTplsItem           `json:"paramTpls,omitempty"`
		Version           *string                   `json:"version,omitempty"`
		TerminationPolicy *ClusterTerminationPolicy `json:"terminationPolicy,omitempty"`
		Mode              *string                   `json:"mode,omitempty"`
		Components        []ComponentItemCreate     `json:"components,omitempty"`
		Extra             map[string]interface{}    `json:"extra,omitempty"`
		InitOptions       []InitOptionItem          `json:"initOptions,omitempty"`
		SingleZone        *bool                     `json:"singleZone,omitempty"`
		AvailabilityZones []string                  `json:"availabilityZones,omitempty"`
		Backup            *ClusterBackup            `json:"backup,omitempty"`
		NodeGroup         common.NullableString     `json:"nodeGroup,omitempty"`
		DisplayName       *string                   `json:"displayName,omitempty"`
		Static            *bool                     `json:"static,omitempty"`
		NetworkMode       *NetworkMode              `json:"networkMode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
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
		common.DeleteKeys(additionalProperties, &[]string{"parentId", "clusterType", "orgName", "environmentName", "project", "name", "engine", "license", "paramTpls", "version", "terminationPolicy", "mode", "components", "extra", "initOptions", "singleZone", "availabilityZones", "backup", "nodeGroup", "displayName", "static", "networkMode"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ParentId = all.ParentId
	if all.ClusterType.Get() != nil && !all.ClusterType.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.ClusterType = all.ClusterType
	}
	o.OrgName = all.OrgName
	o.EnvironmentName = *all.EnvironmentName
	o.Project = all.Project
	o.Name = *all.Name
	o.Engine = *all.Engine
	if all.License != nil && all.License.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.License = all.License
	o.ParamTpls = all.ParamTpls
	o.Version = all.Version
	if all.TerminationPolicy != nil && !all.TerminationPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.TerminationPolicy = all.TerminationPolicy
	}
	o.Mode = all.Mode
	o.Components = all.Components
	o.Extra = all.Extra
	o.InitOptions = all.InitOptions
	o.SingleZone = all.SingleZone
	o.AvailabilityZones = all.AvailabilityZones
	if all.Backup != nil && all.Backup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Backup = all.Backup
	o.NodeGroup = all.NodeGroup
	o.DisplayName = all.DisplayName
	o.Static = all.Static
	if all.NetworkMode != nil && !all.NetworkMode.IsValid() {
		hasInvalidField = true
	} else {
		o.NetworkMode = all.NetworkMode
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
