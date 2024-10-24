// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RecycleBinClusterListItem information of kubeblocks cluster in recycle bin
// NODESCRIPTION RecycleBinClusterListItem
//
// Deprecated: This model is deprecated.
type RecycleBinClusterListItem struct {
	// Cloud Provider
	CloudProvider string `json:"cloudProvider"`
	// Cloud Provider
	CloudRegion *string `json:"cloudRegion,omitempty"`
	// Availability Zones
	AvailabilityZones []string `json:"availabilityZones,omitempty"`
	// Display name of cluster.
	DisplayName *string `json:"displayName,omitempty"`
	// Cluster Application Engine
	Engine string `json:"engine"`
	// Cluster topology mode
	Mode *string `json:"mode,omitempty"`
	// Environment Name
	EnvironmentName string `json:"environmentName"`
	// Cluster Recycle Bin ID
	Id string `json:"id"`
	// Cluster ID
	ClusterId *string `json:"clusterId,omitempty"`
	// Name of cluster. Name must be unique within an Org
	Name string `json:"name"`
	// Cluster Application Version
	Version string `json:"version"`
	// Cluster main component classCode
	ClassCode *string `json:"classCode,omitempty"`
	// Cluster main component storage
	Storage *string `json:"storage,omitempty"`
	// Name of the Org
	OrgName *string `json:"orgName,omitempty"`
	// state represents whether the cluster has been deleted by a stop opsRequest,  and therefore, whether it is in the recycle bin. It is differnt from the Status.Phase of the cluster in k8s.
	State string `json:"state"`
	// StoppedAt is a timestamp representing the server time when this object was stopped and moved to the recycle bin.
	StoppedAt time.Time `json:"stoppedAt"`
	// CreatedAt is a timestamp representing the server time when this object was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// ExpiredAt is a timestamp representing the server time when this object will be expired, and deleted automatically.
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRecycleBinClusterListItem instantiates a new RecycleBinClusterListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecycleBinClusterListItem(cloudProvider string, engine string, environmentName string, id string, name string, version string, state string, stoppedAt time.Time) *RecycleBinClusterListItem {
	this := RecycleBinClusterListItem{}
	this.CloudProvider = cloudProvider
	this.Engine = engine
	this.EnvironmentName = environmentName
	this.Id = id
	this.Name = name
	this.Version = version
	this.State = state
	this.StoppedAt = stoppedAt
	return &this
}

// NewRecycleBinClusterListItemWithDefaults instantiates a new RecycleBinClusterListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecycleBinClusterListItemWithDefaults() *RecycleBinClusterListItem {
	this := RecycleBinClusterListItem{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value.
func (o *RecycleBinClusterListItem) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value.
func (o *RecycleBinClusterListItem) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetCloudRegion returns the CloudRegion field value if set, zero value otherwise.
func (o *RecycleBinClusterListItem) GetCloudRegion() string {
	if o == nil || o.CloudRegion == nil {
		var ret string
		return ret
	}
	return *o.CloudRegion
}

// GetCloudRegionOk returns a tuple with the CloudRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetCloudRegionOk() (*string, bool) {
	if o == nil || o.CloudRegion == nil {
		return nil, false
	}
	return o.CloudRegion, true
}

// HasCloudRegion returns a boolean if a field has been set.
func (o *RecycleBinClusterListItem) HasCloudRegion() bool {
	return o != nil && o.CloudRegion != nil
}

// SetCloudRegion gets a reference to the given string and assigns it to the CloudRegion field.
func (o *RecycleBinClusterListItem) SetCloudRegion(v string) {
	o.CloudRegion = &v
}

// GetAvailabilityZones returns the AvailabilityZones field value if set, zero value otherwise.
func (o *RecycleBinClusterListItem) GetAvailabilityZones() []string {
	if o == nil || o.AvailabilityZones == nil {
		var ret []string
		return ret
	}
	return o.AvailabilityZones
}

// GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetAvailabilityZonesOk() (*[]string, bool) {
	if o == nil || o.AvailabilityZones == nil {
		return nil, false
	}
	return &o.AvailabilityZones, true
}

// HasAvailabilityZones returns a boolean if a field has been set.
func (o *RecycleBinClusterListItem) HasAvailabilityZones() bool {
	return o != nil && o.AvailabilityZones != nil
}

// SetAvailabilityZones gets a reference to the given []string and assigns it to the AvailabilityZones field.
func (o *RecycleBinClusterListItem) SetAvailabilityZones(v []string) {
	o.AvailabilityZones = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *RecycleBinClusterListItem) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RecycleBinClusterListItem) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *RecycleBinClusterListItem) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEngine returns the Engine field value.
func (o *RecycleBinClusterListItem) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *RecycleBinClusterListItem) SetEngine(v string) {
	o.Engine = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *RecycleBinClusterListItem) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *RecycleBinClusterListItem) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *RecycleBinClusterListItem) SetMode(v string) {
	o.Mode = &v
}

// GetEnvironmentName returns the EnvironmentName field value.
func (o *RecycleBinClusterListItem) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *RecycleBinClusterListItem) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetId returns the Id field value.
func (o *RecycleBinClusterListItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RecycleBinClusterListItem) SetId(v string) {
	o.Id = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *RecycleBinClusterListItem) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *RecycleBinClusterListItem) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *RecycleBinClusterListItem) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetName returns the Name field value.
func (o *RecycleBinClusterListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RecycleBinClusterListItem) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value.
func (o *RecycleBinClusterListItem) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *RecycleBinClusterListItem) SetVersion(v string) {
	o.Version = v
}

// GetClassCode returns the ClassCode field value if set, zero value otherwise.
func (o *RecycleBinClusterListItem) GetClassCode() string {
	if o == nil || o.ClassCode == nil {
		var ret string
		return ret
	}
	return *o.ClassCode
}

// GetClassCodeOk returns a tuple with the ClassCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetClassCodeOk() (*string, bool) {
	if o == nil || o.ClassCode == nil {
		return nil, false
	}
	return o.ClassCode, true
}

// HasClassCode returns a boolean if a field has been set.
func (o *RecycleBinClusterListItem) HasClassCode() bool {
	return o != nil && o.ClassCode != nil
}

// SetClassCode gets a reference to the given string and assigns it to the ClassCode field.
func (o *RecycleBinClusterListItem) SetClassCode(v string) {
	o.ClassCode = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *RecycleBinClusterListItem) GetStorage() string {
	if o == nil || o.Storage == nil {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetStorageOk() (*string, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *RecycleBinClusterListItem) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *RecycleBinClusterListItem) SetStorage(v string) {
	o.Storage = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *RecycleBinClusterListItem) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *RecycleBinClusterListItem) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *RecycleBinClusterListItem) SetOrgName(v string) {
	o.OrgName = &v
}

// GetState returns the State field value.
func (o *RecycleBinClusterListItem) GetState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *RecycleBinClusterListItem) SetState(v string) {
	o.State = v
}

// GetStoppedAt returns the StoppedAt field value.
func (o *RecycleBinClusterListItem) GetStoppedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StoppedAt
}

// GetStoppedAtOk returns a tuple with the StoppedAt field value
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetStoppedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoppedAt, true
}

// SetStoppedAt sets field value.
func (o *RecycleBinClusterListItem) SetStoppedAt(v time.Time) {
	o.StoppedAt = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RecycleBinClusterListItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RecycleBinClusterListItem) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RecycleBinClusterListItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *RecycleBinClusterListItem) GetExpiredAt() time.Time {
	if o == nil || o.ExpiredAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecycleBinClusterListItem) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiredAt == nil {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *RecycleBinClusterListItem) HasExpiredAt() bool {
	return o != nil && o.ExpiredAt != nil
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *RecycleBinClusterListItem) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecycleBinClusterListItem) MarshalJSON() ([]byte, error) {
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
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["engine"] = o.Engine
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["environmentName"] = o.EnvironmentName
	toSerialize["id"] = o.Id
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version
	if o.ClassCode != nil {
		toSerialize["classCode"] = o.ClassCode
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	toSerialize["state"] = o.State
	if o.StoppedAt.Nanosecond() == 0 {
		toSerialize["stoppedAt"] = o.StoppedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["stoppedAt"] = o.StoppedAt.Format("2006-01-02T15:04:05.000Z07:00")
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
func (o *RecycleBinClusterListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CloudProvider     *string    `json:"cloudProvider"`
		CloudRegion       *string    `json:"cloudRegion,omitempty"`
		AvailabilityZones []string   `json:"availabilityZones,omitempty"`
		DisplayName       *string    `json:"displayName,omitempty"`
		Engine            *string    `json:"engine"`
		Mode              *string    `json:"mode,omitempty"`
		EnvironmentName   *string    `json:"environmentName"`
		Id                *string    `json:"id"`
		ClusterId         *string    `json:"clusterId,omitempty"`
		Name              *string    `json:"name"`
		Version           *string    `json:"version"`
		ClassCode         *string    `json:"classCode,omitempty"`
		Storage           *string    `json:"storage,omitempty"`
		OrgName           *string    `json:"orgName,omitempty"`
		State             *string    `json:"state"`
		StoppedAt         *time.Time `json:"stoppedAt"`
		CreatedAt         *time.Time `json:"createdAt,omitempty"`
		ExpiredAt         *time.Time `json:"expiredAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CloudProvider == nil {
		return fmt.Errorf("required field cloudProvider missing")
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
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.StoppedAt == nil {
		return fmt.Errorf("required field stoppedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cloudProvider", "cloudRegion", "availabilityZones", "displayName", "engine", "mode", "environmentName", "id", "clusterId", "name", "version", "classCode", "storage", "orgName", "state", "stoppedAt", "createdAt", "expiredAt"})
	} else {
		return err
	}
	o.CloudProvider = *all.CloudProvider
	o.CloudRegion = all.CloudRegion
	o.AvailabilityZones = all.AvailabilityZones
	o.DisplayName = all.DisplayName
	o.Engine = *all.Engine
	o.Mode = all.Mode
	o.EnvironmentName = *all.EnvironmentName
	o.Id = *all.Id
	o.ClusterId = all.ClusterId
	o.Name = *all.Name
	o.Version = *all.Version
	o.ClassCode = all.ClassCode
	o.Storage = all.Storage
	o.OrgName = all.OrgName
	o.State = *all.State
	o.StoppedAt = *all.StoppedAt
	o.CreatedAt = all.CreatedAt
	o.ExpiredAt = all.ExpiredAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
