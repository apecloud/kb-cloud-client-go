// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RecycleBinCluster KubeBlocks cluster(in recycle bin) information
// NODESCRIPTION RecycleBinCluster
//
// Deprecated: This model is deprecated.
type RecycleBinCluster struct {
	// Cluster Recycle Bin ID
	Id *string `json:"id,omitempty"`
	// Cluster ID
	ClusterId *string `json:"clusterId,omitempty"`
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
	Namespace *string `json:"namespace,omitempty"`
	// Name of cluster. Name must be unique within an Org
	Name string `json:"name"`
	// Hash of cluster. Name must be unique within an Org
	Hash *string `json:"hash,omitempty"`
	// Cluster Application Engine
	Engine string `json:"engine"`
	// Values
	Values map[string]interface{} `json:"values,omitempty"`
	// Cluster Application Version
	Version *string `json:"version,omitempty"`
	// The number of replicas, for standalone mode, the replicas is 1, for raftGroup mode, the default replicas is 3.
	Replicas *int32 `json:"replicas,omitempty"`
	// CPU cores.
	Cpu *float `json:"cpu,omitempty"`
	// Memory, the unit is Gi.
	Memory *float `json:"memory,omitempty"`
	// Storage size, the unit is Gi.
	Storage *float `json:"storage,omitempty"`
	// status represents the actual status of the cluster in k8s,  it is different from the state of the cluster in recycle bin,  which means whether the cluster has been deleted by a stop opsRequest.
	Status *string `json:"status,omitempty"`
	// state represents whether the cluster has been deleted by a stop opsRequest,  and therefore, whether it is in the recycle bin. It is differnt from the Status of the cluster undeleted.
	State *string `json:"state,omitempty"`
	// Cluster topology mode
	Mode *string `json:"mode,omitempty"`
	// Items is the list of ComponentSpec in the list
	Components []ComponentsItem `json:"components,omitempty"`
	// Use single availability zones
	SingleZone *bool `json:"singleZone,omitempty"`
	// Availability Zones
	AvailabilityZones []string `json:"availabilityZones,omitempty"`
	// clusterBackup is the payload for cluster backup
	Backup *ClusterBackup `json:"backup,omitempty"`
	// StoppedAt is a timestamp representing the server time when this object was stopped and moved to the recycle bin.
	StoppedAt *time.Time `json:"stoppedAt,omitempty"`
	// CreatedAt is a timestamp representing the server time when this object was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// ExpiredAt is a timestamp representing the server time when this object will be expired, and deleted automatically.
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRecycleBinCluster instantiates a new RecycleBinCluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecycleBinCluster(environmentName string, name string, engine string) *RecycleBinCluster {
	this := RecycleBinCluster{}
	this.EnvironmentName = environmentName
	var namespace string = "kubeblocks-cloud-ns"
	this.Namespace = &namespace
	this.Name = name
	this.Engine = engine
	var singleZone bool = false
	this.SingleZone = &singleZone
	return &this
}

// NewRecycleBinClusterWithDefaults instantiates a new RecycleBinCluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecycleBinClusterWithDefaults() *RecycleBinCluster {
	this := RecycleBinCluster{}
	var namespace string = "kubeblocks-cloud-ns"
	this.Namespace = &namespace
	var singleZone bool = false
	this.SingleZone = &singleZone
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RecycleBinCluster) SetId(v string) {
	o.Id = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *RecycleBinCluster) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *RecycleBinCluster) SetOrgName(v string) {
	o.OrgName = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasCloudProvider() bool {
	return o != nil && o.CloudProvider != nil
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *RecycleBinCluster) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId != nil
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *RecycleBinCluster) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetEnvironmentName returns the EnvironmentName field value.
func (o *RecycleBinCluster) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *RecycleBinCluster) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetEnvironmentType() string {
	if o == nil || o.EnvironmentType == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || o.EnvironmentType == nil {
		return nil, false
	}
	return o.EnvironmentType, true
}

// HasEnvironmentType returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasEnvironmentType() bool {
	return o != nil && o.EnvironmentType != nil
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *RecycleBinCluster) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetCloudRegion returns the CloudRegion field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetCloudRegion() string {
	if o == nil || o.CloudRegion == nil {
		var ret string
		return ret
	}
	return *o.CloudRegion
}

// GetCloudRegionOk returns a tuple with the CloudRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetCloudRegionOk() (*string, bool) {
	if o == nil || o.CloudRegion == nil {
		return nil, false
	}
	return o.CloudRegion, true
}

// HasCloudRegion returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasCloudRegion() bool {
	return o != nil && o.CloudRegion != nil
}

// SetCloudRegion gets a reference to the given string and assigns it to the CloudRegion field.
func (o *RecycleBinCluster) SetCloudRegion(v string) {
	o.CloudRegion = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *RecycleBinCluster) SetNamespace(v string) {
	o.Namespace = &v
}

// GetName returns the Name field value.
func (o *RecycleBinCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RecycleBinCluster) SetName(v string) {
	o.Name = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasHash() bool {
	return o != nil && o.Hash != nil
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *RecycleBinCluster) SetHash(v string) {
	o.Hash = &v
}

// GetEngine returns the Engine field value.
func (o *RecycleBinCluster) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *RecycleBinCluster) SetEngine(v string) {
	o.Engine = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetValues() map[string]interface{} {
	if o == nil || o.Values == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetValuesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given map[string]interface{} and assigns it to the Values field.
func (o *RecycleBinCluster) SetValues(v map[string]interface{}) {
	o.Values = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RecycleBinCluster) SetVersion(v string) {
	o.Version = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *RecycleBinCluster) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetCpu() float {
	if o == nil || o.Cpu == nil {
		var ret float
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetCpuOk() (*float, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given float and assigns it to the Cpu field.
func (o *RecycleBinCluster) SetCpu(v float) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetMemory() float {
	if o == nil || o.Memory == nil {
		var ret float
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetMemoryOk() (*float, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given float and assigns it to the Memory field.
func (o *RecycleBinCluster) SetMemory(v float) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetStorage() float {
	if o == nil || o.Storage == nil {
		var ret float
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetStorageOk() (*float, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given float and assigns it to the Storage field.
func (o *RecycleBinCluster) SetStorage(v float) {
	o.Storage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RecycleBinCluster) SetStatus(v string) {
	o.Status = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RecycleBinCluster) SetState(v string) {
	o.State = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *RecycleBinCluster) SetMode(v string) {
	o.Mode = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetComponents() []ComponentsItem {
	if o == nil || o.Components == nil {
		var ret []ComponentsItem
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetComponentsOk() (*[]ComponentsItem, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []ComponentsItem and assigns it to the Components field.
func (o *RecycleBinCluster) SetComponents(v []ComponentsItem) {
	o.Components = v
}

// GetSingleZone returns the SingleZone field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetSingleZone() bool {
	if o == nil || o.SingleZone == nil {
		var ret bool
		return ret
	}
	return *o.SingleZone
}

// GetSingleZoneOk returns a tuple with the SingleZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetSingleZoneOk() (*bool, bool) {
	if o == nil || o.SingleZone == nil {
		return nil, false
	}
	return o.SingleZone, true
}

// HasSingleZone returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasSingleZone() bool {
	return o != nil && o.SingleZone != nil
}

// SetSingleZone gets a reference to the given bool and assigns it to the SingleZone field.
func (o *RecycleBinCluster) SetSingleZone(v bool) {
	o.SingleZone = &v
}

// GetAvailabilityZones returns the AvailabilityZones field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetAvailabilityZones() []string {
	if o == nil || o.AvailabilityZones == nil {
		var ret []string
		return ret
	}
	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetAvailabilityZonesOk() (*[]string, bool) {
	if o == nil || o.AvailabilityZones == nil {
		return nil, false
	}
	return &o.AvailabilityZones, true
}

// HasAvailabilityZones returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasAvailabilityZones() bool {
	return o != nil && o.AvailabilityZones != nil
}

// SetAvailabilityZones gets a reference to the given []string and assigns it to the AvailabilityZones field.
func (o *RecycleBinCluster) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = v
}

// GetBackup returns the Backup field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetBackup() ClusterBackup {
	if o == nil || o.Backup == nil {
		var ret ClusterBackup
		return ret
	}
	return *o.Backup
}

// GetBackupOk returns a tuple with the Backup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetBackupOk() (*ClusterBackup, bool) {
	if o == nil || o.Backup == nil {
		return nil, false
	}
	return o.Backup, true
}

// HasBackup returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasBackup() bool {
	return o != nil && o.Backup != nil
}

// SetBackup gets a reference to the given ClusterBackup and assigns it to the Backup field.
func (o *RecycleBinCluster) SetBackup(v ClusterBackup) {
	o.Backup = &v
}

// GetStoppedAt returns the StoppedAt field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetStoppedAt() time.Time {
	if o == nil || o.StoppedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StoppedAt
}

// GetStoppedAtOk returns a tuple with the StoppedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetStoppedAtOk() (*time.Time, bool) {
	if o == nil || o.StoppedAt == nil {
		return nil, false
	}
	return o.StoppedAt, true
}

// HasStoppedAt returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasStoppedAt() bool {
	return o != nil && o.StoppedAt != nil
}

// SetStoppedAt gets a reference to the given time.Time and assigns it to the StoppedAt field.
func (o *RecycleBinCluster) SetStoppedAt(v time.Time) {
	o.StoppedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RecycleBinCluster) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *RecycleBinCluster) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinCluster) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiredAt == nil {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *RecycleBinCluster) HasExpiredAt() bool {
	return o != nil && o.ExpiredAt != nil
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *RecycleBinCluster) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecycleBinCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
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
	toSerialize["name"] = o.Name
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	toSerialize["engine"] = o.Engine
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
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
	if o.StoppedAt != nil {
		if o.StoppedAt.Nanosecond() == 0 {
			toSerialize["stoppedAt"] = o.StoppedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["stoppedAt"] = o.StoppedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ExpiredAt != nil {
		if o.ExpiredAt.Nanosecond() == 0 {
			toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expiredAt"] = o.ExpiredAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RecycleBinCluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                *string                `json:"id,omitempty"`
		ClusterId         *string                `json:"clusterId,omitempty"`
		OrgName           *string                `json:"orgName,omitempty"`
		CloudProvider     *string                `json:"cloudProvider,omitempty"`
		EnvironmentId     *string                `json:"environmentId,omitempty"`
		EnvironmentName   *string                `json:"environmentName"`
		EnvironmentType   *string                `json:"environmentType,omitempty"`
		CloudRegion       *string                `json:"cloudRegion,omitempty"`
		Namespace         *string                `json:"namespace,omitempty"`
		Name              *string                `json:"name"`
		Hash              *string                `json:"hash,omitempty"`
		Engine            *string                `json:"engine"`
		Values            map[string]interface{} `json:"values,omitempty"`
		Version           *string                `json:"version,omitempty"`
		Replicas          *int32                 `json:"replicas,omitempty"`
		Cpu               *float                 `json:"cpu,omitempty"`
		Memory            *float                 `json:"memory,omitempty"`
		Storage           *float                 `json:"storage,omitempty"`
		Status            *string                `json:"status,omitempty"`
		State             *string                `json:"state,omitempty"`
		Mode              *string                `json:"mode,omitempty"`
		Components        []ComponentsItem       `json:"components,omitempty"`
		SingleZone        *bool                  `json:"singleZone,omitempty"`
		AvailabilityZones []string               `json:"availabilityZones,omitempty"`
		Backup            *ClusterBackup         `json:"backup,omitempty"`
		StoppedAt         *time.Time             `json:"stoppedAt,omitempty"`
		CreatedAt         *time.Time             `json:"createdAt,omitempty"`
		ExpiredAt         *time.Time             `json:"expiredAt,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"id", "clusterId", "orgName", "cloudProvider", "environmentId", "environmentName", "environmentType", "cloudRegion", "namespace", "name", "hash", "engine", "values", "version", "replicas", "cpu", "memory", "storage", "status", "state", "mode", "components", "singleZone", "availabilityZones", "backup", "stoppedAt", "createdAt", "expiredAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.ClusterId = all.ClusterId
	o.OrgName = all.OrgName
	o.CloudProvider = all.CloudProvider
	o.EnvironmentId = all.EnvironmentId
	o.EnvironmentName = *all.EnvironmentName
	o.EnvironmentType = all.EnvironmentType
	o.CloudRegion = all.CloudRegion
	o.Namespace = all.Namespace
	o.Name = *all.Name
	o.Hash = all.Hash
	o.Engine = *all.Engine
	o.Values = all.Values
	o.Version = all.Version
	o.Replicas = all.Replicas
	o.Cpu = all.Cpu
	o.Memory = all.Memory
	o.Storage = all.Storage
	o.Status = all.Status
	o.State = all.State
	o.Mode = all.Mode
	o.Components = all.Components
	o.SingleZone = all.SingleZone
	o.AvailabilityZones = all.AvailabilityZones
	if all.Backup != nil && all.Backup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Backup = all.Backup
	o.StoppedAt = all.StoppedAt
	o.CreatedAt = all.CreatedAt
	o.ExpiredAt = all.ExpiredAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
