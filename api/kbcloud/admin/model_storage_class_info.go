// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"
)

// StorageClassInfo StorageClassInfo provides detailed information about a specific storage class.
type StorageClassInfo struct {
	// the name of the storage class
	Name string `json:"name"`
	// the creation time of the storage class
	CreationTimestamp string `json:"creationTimestamp"`
	// the provisioner of the storage class
	Provisioner string `json:"provisioner"`
	// the parameters of the storage class
	Parameters map[string]string `json:"parameters,omitempty"`
	// the labels of the storage class
	Labels map[string]string `json:"labels,omitempty"`
	// the annotations of the storage class
	Annotations map[string]string `json:"annotations,omitempty"`
	// the reclaim policy of the storage class
	ReclaimPolicy string `json:"reclaimPolicy"`
	// whether allow volume expansion
	AllowVolumeExpansion bool `json:"allowVolumeExpansion"`
	// the volume binding mode of the storage class
	VolumeBindingMode string `json:"volumeBindingMode"`
	// the number of PVCs
	PvcCount string `json:"pvcCount"`
	// whether allow clone
	AllowClone bool `json:"allowClone"`
	// whether allow snapshot
	AllowSnapshot bool `json:"allowSnapshot"`
	// whether is default class
	IsDefaultClass bool `json:"isDefaultClass"`
	// the type of the storage class
	Type string `json:"type"`
	// the host path when using local storage provisioner
	HostPath *string `json:"hostPath,omitempty"`
	// the mount options of the storage class
	MountOptions []string `json:"mountOptions,omitempty"`
	// the creation time of the storage class
	CreatedAt common.NullableTime `json:"createdAt,omitempty"`
	// the description of the storage class
	Description string `json:"description"`
	// the display name of the storage class
	DisplayName string `json:"displayName"`
	// whether the storage class is enabled
	Enabled bool `json:"enabled"`
	// the id of the storage class
	Id string `json:"id"`
	// the update time of the storage class
	UpdatedAt common.NullableTime `json:"updatedAt,omitempty"`
	// the List stands for stats for the storage volumes of nodes.
	StatsByNodeList *StorageClassInfoStatsByNodeList `json:"statsByNodeList,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageClassInfo instantiates a new StorageClassInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageClassInfo(name string, creationTimestamp string, provisioner string, reclaimPolicy string, allowVolumeExpansion bool, volumeBindingMode string, pvcCount string, allowClone bool, allowSnapshot bool, isDefaultClass bool, typeVar string, description string, displayName string, enabled bool, id string) *StorageClassInfo {
	this := StorageClassInfo{}
	this.Name = name
	this.CreationTimestamp = creationTimestamp
	this.Provisioner = provisioner
	this.ReclaimPolicy = reclaimPolicy
	this.AllowVolumeExpansion = allowVolumeExpansion
	this.VolumeBindingMode = volumeBindingMode
	this.PvcCount = pvcCount
	this.AllowClone = allowClone
	this.AllowSnapshot = allowSnapshot
	this.IsDefaultClass = isDefaultClass
	this.Type = typeVar
	this.Description = description
	this.DisplayName = displayName
	this.Enabled = enabled
	this.Id = id
	return &this
}

// NewStorageClassInfoWithDefaults instantiates a new StorageClassInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageClassInfoWithDefaults() *StorageClassInfo {
	this := StorageClassInfo{}
	return &this
}

// GetName returns the Name field value.
func (o *StorageClassInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *StorageClassInfo) SetName(v string) {
	o.Name = v
}

// GetCreationTimestamp returns the CreationTimestamp field value.
func (o *StorageClassInfo) GetCreationTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetCreationTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationTimestamp, true
}

// SetCreationTimestamp sets field value.
func (o *StorageClassInfo) SetCreationTimestamp(v string) {
	o.CreationTimestamp = v
}

// GetProvisioner returns the Provisioner field value.
func (o *StorageClassInfo) GetProvisioner() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provisioner
}

// GetProvisionerOk returns a tuple with the Provisioner field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetProvisionerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provisioner, true
}

// SetProvisioner sets field value.
func (o *StorageClassInfo) SetProvisioner(v string) {
	o.Provisioner = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *StorageClassInfo) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *StorageClassInfo) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *StorageClassInfo) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *StorageClassInfo) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *StorageClassInfo) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *StorageClassInfo) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *StorageClassInfo) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *StorageClassInfo) HasAnnotations() bool {
	return o != nil && o.Annotations != nil
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *StorageClassInfo) SetAnnotations(v map[string]string) {
	o.Annotations = v
}

// GetReclaimPolicy returns the ReclaimPolicy field value.
func (o *StorageClassInfo) GetReclaimPolicy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReclaimPolicy
}

// GetReclaimPolicyOk returns a tuple with the ReclaimPolicy field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetReclaimPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReclaimPolicy, true
}

// SetReclaimPolicy sets field value.
func (o *StorageClassInfo) SetReclaimPolicy(v string) {
	o.ReclaimPolicy = v
}

// GetAllowVolumeExpansion returns the AllowVolumeExpansion field value.
func (o *StorageClassInfo) GetAllowVolumeExpansion() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AllowVolumeExpansion
}

// GetAllowVolumeExpansionOk returns a tuple with the AllowVolumeExpansion field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetAllowVolumeExpansionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowVolumeExpansion, true
}

// SetAllowVolumeExpansion sets field value.
func (o *StorageClassInfo) SetAllowVolumeExpansion(v bool) {
	o.AllowVolumeExpansion = v
}

// GetVolumeBindingMode returns the VolumeBindingMode field value.
func (o *StorageClassInfo) GetVolumeBindingMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.VolumeBindingMode
}

// GetVolumeBindingModeOk returns a tuple with the VolumeBindingMode field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetVolumeBindingModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeBindingMode, true
}

// SetVolumeBindingMode sets field value.
func (o *StorageClassInfo) SetVolumeBindingMode(v string) {
	o.VolumeBindingMode = v
}

// GetPvcCount returns the PvcCount field value.
func (o *StorageClassInfo) GetPvcCount() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PvcCount
}

// GetPvcCountOk returns a tuple with the PvcCount field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetPvcCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PvcCount, true
}

// SetPvcCount sets field value.
func (o *StorageClassInfo) SetPvcCount(v string) {
	o.PvcCount = v
}

// GetAllowClone returns the AllowClone field value.
func (o *StorageClassInfo) GetAllowClone() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AllowClone
}

// GetAllowCloneOk returns a tuple with the AllowClone field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetAllowCloneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowClone, true
}

// SetAllowClone sets field value.
func (o *StorageClassInfo) SetAllowClone(v bool) {
	o.AllowClone = v
}

// GetAllowSnapshot returns the AllowSnapshot field value.
func (o *StorageClassInfo) GetAllowSnapshot() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AllowSnapshot
}

// GetAllowSnapshotOk returns a tuple with the AllowSnapshot field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetAllowSnapshotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowSnapshot, true
}

// SetAllowSnapshot sets field value.
func (o *StorageClassInfo) SetAllowSnapshot(v bool) {
	o.AllowSnapshot = v
}

// GetIsDefaultClass returns the IsDefaultClass field value.
func (o *StorageClassInfo) GetIsDefaultClass() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDefaultClass
}

// GetIsDefaultClassOk returns a tuple with the IsDefaultClass field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetIsDefaultClassOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefaultClass, true
}

// SetIsDefaultClass sets field value.
func (o *StorageClassInfo) SetIsDefaultClass(v bool) {
	o.IsDefaultClass = v
}

// GetType returns the Type field value.
func (o *StorageClassInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *StorageClassInfo) SetType(v string) {
	o.Type = v
}

// GetHostPath returns the HostPath field value if set, zero value otherwise.
func (o *StorageClassInfo) GetHostPath() string {
	if o == nil || o.HostPath == nil {
		var ret string
		return ret
	}
	return *o.HostPath
}

// GetHostPathOk returns a tuple with the HostPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetHostPathOk() (*string, bool) {
	if o == nil || o.HostPath == nil {
		return nil, false
	}
	return o.HostPath, true
}

// HasHostPath returns a boolean if a field has been set.
func (o *StorageClassInfo) HasHostPath() bool {
	return o != nil && o.HostPath != nil
}

// SetHostPath gets a reference to the given string and assigns it to the HostPath field.
func (o *StorageClassInfo) SetHostPath(v string) {
	o.HostPath = &v
}

// GetMountOptions returns the MountOptions field value if set, zero value otherwise.
func (o *StorageClassInfo) GetMountOptions() []string {
	if o == nil || o.MountOptions == nil {
		var ret []string
		return ret
	}
	return o.MountOptions
}

// GetMountOptionsOk returns a tuple with the MountOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetMountOptionsOk() (*[]string, bool) {
	if o == nil || o.MountOptions == nil {
		return nil, false
	}
	return &o.MountOptions, true
}

// HasMountOptions returns a boolean if a field has been set.
func (o *StorageClassInfo) HasMountOptions() bool {
	return o != nil && o.MountOptions != nil
}

// SetMountOptions gets a reference to the given []string and assigns it to the MountOptions field.
func (o *StorageClassInfo) SetMountOptions(v []string) {
	o.MountOptions = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageClassInfo) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StorageClassInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StorageClassInfo) HasCreatedAt() bool {
	return o != nil && o.CreatedAt.IsSet()
}

// SetCreatedAt gets a reference to the given common.NullableTime and assigns it to the CreatedAt field.
func (o *StorageClassInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil.
func (o *StorageClassInfo) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil.
func (o *StorageClassInfo) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDescription returns the Description field value.
func (o *StorageClassInfo) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *StorageClassInfo) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value.
func (o *StorageClassInfo) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *StorageClassInfo) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEnabled returns the Enabled field value.
func (o *StorageClassInfo) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *StorageClassInfo) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *StorageClassInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *StorageClassInfo) SetId(v string) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageClassInfo) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StorageClassInfo) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *StorageClassInfo) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt.IsSet()
}

// SetUpdatedAt gets a reference to the given common.NullableTime and assigns it to the UpdatedAt field.
func (o *StorageClassInfo) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil.
func (o *StorageClassInfo) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil.
func (o *StorageClassInfo) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetStatsByNodeList returns the StatsByNodeList field value if set, zero value otherwise.
func (o *StorageClassInfo) GetStatsByNodeList() StorageClassInfoStatsByNodeList {
	if o == nil || o.StatsByNodeList == nil {
		var ret StorageClassInfoStatsByNodeList
		return ret
	}
	return *o.StatsByNodeList
}

// GetStatsByNodeListOk returns a tuple with the StatsByNodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassInfo) GetStatsByNodeListOk() (*StorageClassInfoStatsByNodeList, bool) {
	if o == nil || o.StatsByNodeList == nil {
		return nil, false
	}
	return o.StatsByNodeList, true
}

// HasStatsByNodeList returns a boolean if a field has been set.
func (o *StorageClassInfo) HasStatsByNodeList() bool {
	return o != nil && o.StatsByNodeList != nil
}

// SetStatsByNodeList gets a reference to the given StorageClassInfoStatsByNodeList and assigns it to the StatsByNodeList field.
func (o *StorageClassInfo) SetStatsByNodeList(v StorageClassInfoStatsByNodeList) {
	o.StatsByNodeList = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageClassInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["creationTimestamp"] = o.CreationTimestamp
	toSerialize["provisioner"] = o.Provisioner
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	toSerialize["reclaimPolicy"] = o.ReclaimPolicy
	toSerialize["allowVolumeExpansion"] = o.AllowVolumeExpansion
	toSerialize["volumeBindingMode"] = o.VolumeBindingMode
	toSerialize["pvcCount"] = o.PvcCount
	toSerialize["allowClone"] = o.AllowClone
	toSerialize["allowSnapshot"] = o.AllowSnapshot
	toSerialize["isDefaultClass"] = o.IsDefaultClass
	toSerialize["type"] = o.Type
	if o.HostPath != nil {
		toSerialize["hostPath"] = o.HostPath
	}
	if o.MountOptions != nil {
		toSerialize["mountOptions"] = o.MountOptions
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	toSerialize["description"] = o.Description
	toSerialize["displayName"] = o.DisplayName
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	if o.UpdatedAt.IsSet() {
		toSerialize["updatedAt"] = o.UpdatedAt.Get()
	}
	if o.StatsByNodeList != nil {
		toSerialize["statsByNodeList"] = o.StatsByNodeList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageClassInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                 *string                          `json:"name"`
		CreationTimestamp    *string                          `json:"creationTimestamp"`
		Provisioner          *string                          `json:"provisioner"`
		Parameters           map[string]string                `json:"parameters,omitempty"`
		Labels               map[string]string                `json:"labels,omitempty"`
		Annotations          map[string]string                `json:"annotations,omitempty"`
		ReclaimPolicy        *string                          `json:"reclaimPolicy"`
		AllowVolumeExpansion *bool                            `json:"allowVolumeExpansion"`
		VolumeBindingMode    *string                          `json:"volumeBindingMode"`
		PvcCount             *string                          `json:"pvcCount"`
		AllowClone           *bool                            `json:"allowClone"`
		AllowSnapshot        *bool                            `json:"allowSnapshot"`
		IsDefaultClass       *bool                            `json:"isDefaultClass"`
		Type                 *string                          `json:"type"`
		HostPath             *string                          `json:"hostPath,omitempty"`
		MountOptions         []string                         `json:"mountOptions,omitempty"`
		CreatedAt            common.NullableTime              `json:"createdAt,omitempty"`
		Description          *string                          `json:"description"`
		DisplayName          *string                          `json:"displayName"`
		Enabled              *bool                            `json:"enabled"`
		Id                   *string                          `json:"id"`
		UpdatedAt            common.NullableTime              `json:"updatedAt,omitempty"`
		StatsByNodeList      *StorageClassInfoStatsByNodeList `json:"statsByNodeList,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.CreationTimestamp == nil {
		return fmt.Errorf("required field creationTimestamp missing")
	}
	if all.Provisioner == nil {
		return fmt.Errorf("required field provisioner missing")
	}
	if all.ReclaimPolicy == nil {
		return fmt.Errorf("required field reclaimPolicy missing")
	}
	if all.AllowVolumeExpansion == nil {
		return fmt.Errorf("required field allowVolumeExpansion missing")
	}
	if all.VolumeBindingMode == nil {
		return fmt.Errorf("required field volumeBindingMode missing")
	}
	if all.PvcCount == nil {
		return fmt.Errorf("required field pvcCount missing")
	}
	if all.AllowClone == nil {
		return fmt.Errorf("required field allowClone missing")
	}
	if all.AllowSnapshot == nil {
		return fmt.Errorf("required field allowSnapshot missing")
	}
	if all.IsDefaultClass == nil {
		return fmt.Errorf("required field isDefaultClass missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field displayName missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "creationTimestamp", "provisioner", "parameters", "labels", "annotations", "reclaimPolicy", "allowVolumeExpansion", "volumeBindingMode", "pvcCount", "allowClone", "allowSnapshot", "isDefaultClass", "type", "hostPath", "mountOptions", "createdAt", "description", "displayName", "enabled", "id", "updatedAt", "statsByNodeList"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.CreationTimestamp = *all.CreationTimestamp
	o.Provisioner = *all.Provisioner
	o.Parameters = all.Parameters
	o.Labels = all.Labels
	o.Annotations = all.Annotations
	o.ReclaimPolicy = *all.ReclaimPolicy
	o.AllowVolumeExpansion = *all.AllowVolumeExpansion
	o.VolumeBindingMode = *all.VolumeBindingMode
	o.PvcCount = *all.PvcCount
	o.AllowClone = *all.AllowClone
	o.AllowSnapshot = *all.AllowSnapshot
	o.IsDefaultClass = *all.IsDefaultClass
	o.Type = *all.Type
	o.HostPath = all.HostPath
	o.MountOptions = all.MountOptions
	o.CreatedAt = all.CreatedAt
	o.Description = *all.Description
	o.DisplayName = *all.DisplayName
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.UpdatedAt = all.UpdatedAt
	if all.StatsByNodeList != nil && all.StatsByNodeList.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatsByNodeList = all.StatsByNodeList

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
