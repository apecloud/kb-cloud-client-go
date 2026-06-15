// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SLACluster The cluster summary used by SLA calculation results. SLA responses can be built from historical outage data, so general cluster list fields may be unavailable.
type SLACluster struct {
	// ID of cluster.
	Id *string `json:"id,omitempty"`
	// Name of cluster.
	Name *string `json:"name,omitempty"`
	// Display name of cluster.
	DisplayName *string `json:"displayName,omitempty"`
	// Org Name.
	OrgName *string `json:"orgName,omitempty"`
	// Environment Name.
	EnvironmentName *string `json:"environmentName,omitempty"`
	// Cluster Application Engine.
	Engine *string `json:"engine,omitempty"`
	// Describes the type of cluster
	ClusterType NullableClusterType `json:"clusterType,omitempty"`
	// The timestamp when the cluster was created, when available.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The SLA calculation status for the cluster.
	Status *string `json:"status,omitempty"`
	// Cluster Application Version.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSLACluster instantiates a new SLACluster object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLACluster() *SLACluster {
	this := SLACluster{}
	var clusterType ClusterType = ClusterTypeNormal
	this.ClusterType = *NewNullableClusterType(&clusterType)
	return &this
}

// NewSLAClusterWithDefaults instantiates a new SLACluster object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLAClusterWithDefaults() *SLACluster {
	this := SLACluster{}
	var clusterType ClusterType = ClusterTypeNormal
	this.ClusterType = *NewNullableClusterType(&clusterType)
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SLACluster) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLACluster) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SLACluster) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SLACluster) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SLACluster) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLACluster) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SLACluster) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SLACluster) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SLACluster) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLACluster) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SLACluster) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SLACluster) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *SLACluster) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLACluster) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *SLACluster) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *SLACluster) SetOrgName(v string) {
	o.OrgName = &v
}

// GetEnvironmentName returns the EnvironmentName field value if set, zero value otherwise.
func (o *SLACluster) GetEnvironmentName() string {
	if o == nil || o.EnvironmentName == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLACluster) GetEnvironmentNameOk() (*string, bool) {
	if o == nil || o.EnvironmentName == nil {
		return nil, false
	}
	return o.EnvironmentName, true
}

// HasEnvironmentName returns a boolean if a field has been set.
func (o *SLACluster) HasEnvironmentName() bool {
	return o != nil && o.EnvironmentName != nil
}

// SetEnvironmentName gets a reference to the given string and assigns it to the EnvironmentName field.
func (o *SLACluster) SetEnvironmentName(v string) {
	o.EnvironmentName = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *SLACluster) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLACluster) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *SLACluster) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *SLACluster) SetEngine(v string) {
	o.Engine = &v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SLACluster) GetClusterType() ClusterType {
	if o == nil || o.ClusterType.Get() == nil {
		var ret ClusterType
		return ret
	}
	return *o.ClusterType.Get()
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SLACluster) GetClusterTypeOk() (*ClusterType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterType.Get(), o.ClusterType.IsSet()
}

// HasClusterType returns a boolean if a field has been set.
func (o *SLACluster) HasClusterType() bool {
	return o != nil && o.ClusterType.IsSet()
}

// SetClusterType gets a reference to the given NullableClusterType and assigns it to the ClusterType field.
func (o *SLACluster) SetClusterType(v ClusterType) {
	o.ClusterType.Set(&v)
}

// SetClusterTypeNil sets the value for ClusterType to be an explicit nil.
func (o *SLACluster) SetClusterTypeNil() {
	o.ClusterType.Set(nil)
}

// UnsetClusterType ensures that no value is present for ClusterType, not even an explicit nil.
func (o *SLACluster) UnsetClusterType() {
	o.ClusterType.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SLACluster) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLACluster) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SLACluster) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SLACluster) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SLACluster) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLACluster) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SLACluster) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SLACluster) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SLACluster) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLACluster) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SLACluster) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SLACluster) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLACluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.EnvironmentName != nil {
		toSerialize["environmentName"] = o.EnvironmentName
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.ClusterType.IsSet() {
		toSerialize["clusterType"] = o.ClusterType.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLACluster) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id              *string             `json:"id,omitempty"`
		Name            *string             `json:"name,omitempty"`
		DisplayName     *string             `json:"displayName,omitempty"`
		OrgName         *string             `json:"orgName,omitempty"`
		EnvironmentName *string             `json:"environmentName,omitempty"`
		Engine          *string             `json:"engine,omitempty"`
		ClusterType     NullableClusterType `json:"clusterType,omitempty"`
		CreatedAt       *time.Time          `json:"createdAt,omitempty"`
		Status          *string             `json:"status,omitempty"`
		Version         *string             `json:"version,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "displayName", "orgName", "environmentName", "engine", "clusterType", "createdAt", "status", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = all.Name
	o.DisplayName = all.DisplayName
	o.OrgName = all.OrgName
	o.EnvironmentName = all.EnvironmentName
	o.Engine = all.Engine
	if all.ClusterType.Get() != nil && !all.ClusterType.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.ClusterType = all.ClusterType
	}
	o.CreatedAt = all.CreatedAt
	o.Status = all.Status
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
