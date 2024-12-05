// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterListItem KubeBlocks cluster information
type ClusterListItem struct {
	// Cloud Provider
	CloudProvider string `json:"cloudProvider"`
	// Cloud Provider
	CloudRegion *string `json:"cloudRegion,omitempty"`
	// Availability Zones
	AvailabilityZones []string `json:"availabilityZones,omitempty"`
	// CreatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists
	CreatedAt time.Time `json:"createdAt"`
	// Display name of cluster.
	DisplayName *string `json:"displayName,omitempty"`
	// Cluster Application Engine
	Engine string `json:"engine"`
	// Cluster topology mode
	Mode *string `json:"mode,omitempty"`
	// Environment Name
	EnvironmentName string `json:"environmentName"`
	// ID of cluster
	Id string `json:"id"`
	// Name of cluster. Name must be unique within an Org
	Name string `json:"name"`
	// When two clusters have a relationship, parentId records the parent cluster id.Can be empty when there is no relationship
	ParentId common.NullableInt64 `json:"parentId,omitempty"`
	// the name of parent cluster
	ParentName common.NullableString `json:"parentName,omitempty"`
	// the display name of parent cluster
	ParentDisplayName common.NullableString `json:"parentDisplayName,omitempty"`
	// Describes the type of cluster, [Normal] normal cluster; [DisasterRecovery] disaster recovery cluster
	ClusterType NullableClusterType    `json:"clusterType,omitempty"`
	Delay       common.NullableFloat64 `json:"delay,omitempty"`
	// Cluster Status
	Status string `json:"status"`
	// Cluster termination policy
	TerminationPolicy string `json:"terminationPolicy"`
	// UpdatedAt is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC. Populated by the system. Read-only. Null for lists
	UpdatedAt time.Time `json:"updatedAt"`
	// Cluster Application Version
	Version string `json:"version"`
	// Cluster main component classCode
	ClassCode *string `json:"classCode,omitempty"`
	// Cluster main component storage
	Storage *string `json:"storage,omitempty"`
	// Cluster main component codeShort
	CodeShort *string `json:"codeShort,omitempty"`
	// Org Name
	OrgName *string `json:"orgName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterListItem instantiates a new ClusterListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterListItem(cloudProvider string, createdAt time.Time, engine string, environmentName string, id string, name string, status string, terminationPolicy string, updatedAt time.Time, version string) *ClusterListItem {
	this := ClusterListItem{}
	this.CloudProvider = cloudProvider
	this.CreatedAt = createdAt
	this.Engine = engine
	this.EnvironmentName = environmentName
	this.Id = id
	this.Name = name
	var clusterType ClusterType = ClusterTypeNormal
	this.ClusterType = *NewNullableClusterType(&clusterType)
	this.Status = status
	this.TerminationPolicy = terminationPolicy
	this.UpdatedAt = updatedAt
	this.Version = version
	return &this
}

// NewClusterListItemWithDefaults instantiates a new ClusterListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterListItemWithDefaults() *ClusterListItem {
	this := ClusterListItem{}
	var clusterType ClusterType = ClusterTypeNormal
	this.ClusterType = *NewNullableClusterType(&clusterType)
	return &this
}

// GetCloudProvider returns the CloudProvider field value.
func (o *ClusterListItem) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value.
func (o *ClusterListItem) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetCloudRegion returns the CloudRegion field value if set, zero value otherwise.
func (o *ClusterListItem) GetCloudRegion() string {
	if o == nil || o.CloudRegion == nil {
		var ret string
		return ret
	}
	return *o.CloudRegion
}

// GetCloudRegionOk returns a tuple with the CloudRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetCloudRegionOk() (*string, bool) {
	if o == nil || o.CloudRegion == nil {
		return nil, false
	}
	return o.CloudRegion, true
}

// HasCloudRegion returns a boolean if a field has been set.
func (o *ClusterListItem) HasCloudRegion() bool {
	return o != nil && o.CloudRegion != nil
}

// SetCloudRegion gets a reference to the given string and assigns it to the CloudRegion field.
func (o *ClusterListItem) SetCloudRegion(v string) {
	o.CloudRegion = &v
}

// GetAvailabilityZones returns the AvailabilityZones field value if set, zero value otherwise.
func (o *ClusterListItem) GetAvailabilityZones() []string {
	if o == nil || o.AvailabilityZones == nil {
		var ret []string
		return ret
	}
	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetAvailabilityZonesOk() (*[]string, bool) {
	if o == nil || o.AvailabilityZones == nil {
		return nil, false
	}
	return &o.AvailabilityZones, true
}

// HasAvailabilityZones returns a boolean if a field has been set.
func (o *ClusterListItem) HasAvailabilityZones() bool {
	return o != nil && o.AvailabilityZones != nil
}

// SetAvailabilityZones gets a reference to the given []string and assigns it to the AvailabilityZones field.
func (o *ClusterListItem) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ClusterListItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ClusterListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ClusterListItem) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterListItem) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ClusterListItem) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEngine returns the Engine field value.
func (o *ClusterListItem) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *ClusterListItem) SetEngine(v string) {
	o.Engine = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ClusterListItem) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ClusterListItem) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ClusterListItem) SetMode(v string) {
	o.Mode = &v
}

// GetEnvironmentName returns the EnvironmentName field value.
func (o *ClusterListItem) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *ClusterListItem) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetId returns the Id field value.
func (o *ClusterListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ClusterListItem) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *ClusterListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ClusterListItem) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterListItem) GetParentId() int64 {
	if o == nil || o.ParentId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterListItem) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *ClusterListItem) HasParentId() bool {
	return o != nil && o.ParentId.IsSet()
}

// SetParentId gets a reference to the given common.NullableInt64 and assigns it to the ParentId field.
func (o *ClusterListItem) SetParentId(v int64) {
	o.ParentId.Set(&v)
}

// SetParentIdNil sets the value for ParentId to be an explicit nil.
func (o *ClusterListItem) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil.
func (o *ClusterListItem) UnsetParentId() {
	o.ParentId.Unset()
}

// GetParentName returns the ParentName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterListItem) GetParentName() string {
	if o == nil || o.ParentName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentName.Get()
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterListItem) GetParentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentName.Get(), o.ParentName.IsSet()
}

// HasParentName returns a boolean if a field has been set.
func (o *ClusterListItem) HasParentName() bool {
	return o != nil && o.ParentName.IsSet()
}

// SetParentName gets a reference to the given common.NullableString and assigns it to the ParentName field.
func (o *ClusterListItem) SetParentName(v string) {
	o.ParentName.Set(&v)
}

// SetParentNameNil sets the value for ParentName to be an explicit nil.
func (o *ClusterListItem) SetParentNameNil() {
	o.ParentName.Set(nil)
}

// UnsetParentName ensures that no value is present for ParentName, not even an explicit nil.
func (o *ClusterListItem) UnsetParentName() {
	o.ParentName.Unset()
}

// GetParentDisplayName returns the ParentDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterListItem) GetParentDisplayName() string {
	if o == nil || o.ParentDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentDisplayName.Get()
}

// GetParentDisplayNameOk returns a tuple with the ParentDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterListItem) GetParentDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentDisplayName.Get(), o.ParentDisplayName.IsSet()
}

// HasParentDisplayName returns a boolean if a field has been set.
func (o *ClusterListItem) HasParentDisplayName() bool {
	return o != nil && o.ParentDisplayName.IsSet()
}

// SetParentDisplayName gets a reference to the given common.NullableString and assigns it to the ParentDisplayName field.
func (o *ClusterListItem) SetParentDisplayName(v string) {
	o.ParentDisplayName.Set(&v)
}

// SetParentDisplayNameNil sets the value for ParentDisplayName to be an explicit nil.
func (o *ClusterListItem) SetParentDisplayNameNil() {
	o.ParentDisplayName.Set(nil)
}

// UnsetParentDisplayName ensures that no value is present for ParentDisplayName, not even an explicit nil.
func (o *ClusterListItem) UnsetParentDisplayName() {
	o.ParentDisplayName.Unset()
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterListItem) GetClusterType() ClusterType {
	if o == nil || o.ClusterType.Get() == nil {
		var ret ClusterType
		return ret
	}
	return *o.ClusterType.Get()
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterListItem) GetClusterTypeOk() (*ClusterType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterType.Get(), o.ClusterType.IsSet()
}

// HasClusterType returns a boolean if a field has been set.
func (o *ClusterListItem) HasClusterType() bool {
	return o != nil && o.ClusterType.IsSet()
}

// SetClusterType gets a reference to the given NullableClusterType and assigns it to the ClusterType field.
func (o *ClusterListItem) SetClusterType(v ClusterType) {
	o.ClusterType.Set(&v)
}

// SetClusterTypeNil sets the value for ClusterType to be an explicit nil.
func (o *ClusterListItem) SetClusterTypeNil() {
	o.ClusterType.Set(nil)
}

// UnsetClusterType ensures that no value is present for ClusterType, not even an explicit nil.
func (o *ClusterListItem) UnsetClusterType() {
	o.ClusterType.Unset()
}

// GetDelay returns the Delay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterListItem) GetDelay() float64 {
	if o == nil || o.Delay.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Delay.Get()
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterListItem) GetDelayOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Delay.Get(), o.Delay.IsSet()
}

// HasDelay returns a boolean if a field has been set.
func (o *ClusterListItem) HasDelay() bool {
	return o != nil && o.Delay.IsSet()
}

// SetDelay gets a reference to the given common.NullableFloat64 and assigns it to the Delay field.
func (o *ClusterListItem) SetDelay(v float64) {
	o.Delay.Set(&v)
}

// SetDelayNil sets the value for Delay to be an explicit nil.
func (o *ClusterListItem) SetDelayNil() {
	o.Delay.Set(nil)
}

// UnsetDelay ensures that no value is present for Delay, not even an explicit nil.
func (o *ClusterListItem) UnsetDelay() {
	o.Delay.Unset()
}

// GetStatus returns the Status field value.
func (o *ClusterListItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ClusterListItem) SetStatus(v string) {
	o.Status = v
}

// GetTerminationPolicy returns the TerminationPolicy field value.
func (o *ClusterListItem) GetTerminationPolicy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TerminationPolicy
}

// GetTerminationPolicyOk returns a tuple with the TerminationPolicy field value
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetTerminationPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationPolicy, true
}

// SetTerminationPolicy sets field value.
func (o *ClusterListItem) SetTerminationPolicy(v string) {
	o.TerminationPolicy = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ClusterListItem) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ClusterListItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVersion returns the Version field value.
func (o *ClusterListItem) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *ClusterListItem) SetVersion(v string) {
	o.Version = v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *ClusterListItem) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *ClusterListItem) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *ClusterListItem) SetClassCode(v string) {
	o.ClassCode = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ClusterListItem) GetStorage() string {
	if o == nil || o.Storage == nil {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetStorageOk() (*string, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ClusterListItem) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *ClusterListItem) SetStorage(v string) {
	o.Storage = &v
}

// GetCodeShort returns the CodeShort field value if set, zero value otherwise.
func (o *ClusterListItem) GetCodeShort() string {
	if o == nil || o.CodeShort == nil {
		var ret string
		return ret
	}
	return *o.CodeShort
}

// GetCodeShortOk returns a tuple with the CodeShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetCodeShortOk() (*string, bool) {
	if o == nil || o.CodeShort == nil {
		return nil, false
	}
	return o.CodeShort, true
}

// HasCodeShort returns a boolean if a field has been set.
func (o *ClusterListItem) HasCodeShort() bool {
	return o != nil && o.CodeShort != nil
}

// SetCodeShort gets a reference to the given string and assigns it to the CodeShort field.
func (o *ClusterListItem) SetCodeShort(v string) {
	o.CodeShort = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *ClusterListItem) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterListItem) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *ClusterListItem) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *ClusterListItem) SetOrgName(v string) {
	o.OrgName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cloudProvider"] = o.CloudProvider
	if o.CloudRegion != nil {
		toSerialize["cloudRegion"] = o.CloudRegion
	}
	if o.AvailabilityZones != nil {
		toSerialize["availabilityZones"] = o.AvailabilityZones
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["engine"] = o.Engine
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["environmentName"] = o.EnvironmentName
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
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
	toSerialize["status"] = o.Status
	toSerialize["terminationPolicy"] = o.TerminationPolicy
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["version"] = o.Version
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.CodeShort != nil {
		toSerialize["codeShort"] = o.CodeShort
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CloudProvider     *string                `json:"cloudProvider"`
		CloudRegion       *string                `json:"cloudRegion,omitempty"`
		AvailabilityZones []string               `json:"availabilityZones,omitempty"`
		CreatedAt         *time.Time             `json:"createdAt"`
		DisplayName       *string                `json:"displayName,omitempty"`
		Engine            *string                `json:"engine"`
		Mode              *string                `json:"mode,omitempty"`
		EnvironmentName   *string                `json:"environmentName"`
		Id                *string                `json:"id"`
		Name              *string                `json:"name"`
		ParentId          common.NullableInt64   `json:"parentId,omitempty"`
		ParentName        common.NullableString  `json:"parentName,omitempty"`
		ParentDisplayName common.NullableString  `json:"parentDisplayName,omitempty"`
		ClusterType       NullableClusterType    `json:"clusterType,omitempty"`
		Delay             common.NullableFloat64 `json:"delay,omitempty"`
		Status            *string                `json:"status"`
		TerminationPolicy *string                `json:"terminationPolicy"`
		UpdatedAt         *time.Time             `json:"updatedAt"`
		Version           *string                `json:"version"`
		ClassCode         *string                `json:"classCode,omitempty"`
		Storage           *string                `json:"storage,omitempty"`
		CodeShort         *string                `json:"codeShort,omitempty"`
		OrgName           *string                `json:"orgName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CloudProvider == nil {
		return fmt.Errorf("required field cloudProvider missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.EnvironmentName == nil {
		return fmt.Errorf("required field environmentName missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.TerminationPolicy == nil {
		return fmt.Errorf("required field terminationPolicy missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cloudProvider", "cloudRegion", "availabilityZones", "createdAt", "displayName", "engine", "mode", "environmentName", "id", "name", "parentId", "parentName", "parentDisplayName", "clusterType", "delay", "status", "terminationPolicy", "updatedAt", "version", "classCode", "storage", "codeShort", "orgName"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CloudProvider = *all.CloudProvider
	o.CloudRegion = all.CloudRegion
	o.AvailabilityZones = all.AvailabilityZones
	o.CreatedAt = *all.CreatedAt
	o.DisplayName = all.DisplayName
	o.Engine = *all.Engine
	o.Mode = all.Mode
	o.EnvironmentName = *all.EnvironmentName
	o.Id = *all.Id
	o.Name = *all.Name
	o.ParentId = all.ParentId
	o.ParentName = all.ParentName
	o.ParentDisplayName = all.ParentDisplayName
	if all.ClusterType.Get() != nil && !all.ClusterType.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.ClusterType = all.ClusterType
	}
	o.Delay = all.Delay
	o.Status = *all.Status
	o.TerminationPolicy = *all.TerminationPolicy
	o.UpdatedAt = *all.UpdatedAt
	o.Version = *all.Version
	o.ClassCode = all.ClassCode
	o.Storage = all.Storage
	o.CodeShort = all.CodeShort
	o.OrgName = all.OrgName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
