// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageClassCreate StorageClassCreate provides detailed creation information about a storage class.
// NODESCRIPTION StorageClassCreate
//
// Deprecated: This model is deprecated.
type StorageClassCreate struct {
	// the name of the storage class
	Name string `json:"name"`
	// the provisioner of the storage class
	Provisioner *string `json:"provisioner,omitempty"`
	// the labels of the storage class
	Labels map[string]string `json:"labels,omitempty"`
	// the annotations of the storage class
	Annotations map[string]string `json:"annotations,omitempty"`
	// the parameters of the storage class
	Parameters map[string]string `json:"parameters,omitempty"`
	// reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
	ReclaimPolicy StorageClassReclaimPolicy `json:"reclaimPolicy"`
	// whether allow volume expansion
	AllowVolumeExpansion *bool `json:"allowVolumeExpansion,omitempty"`
	// volumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound. Defaults to Immediate.
	VolumeBindingMode StorageClassVolumeBindingMode `json:"volumeBindingMode"`
	// whether allow clone
	AllowClone *bool `json:"allowClone,omitempty"`
	// whether allow snapshot
	AllowSnapshot *bool `json:"allowSnapshot,omitempty"`
	// whether is default class
	IsDefaultClass *bool `json:"isDefaultClass,omitempty"`
	// the type of the storage class
	Type *string `json:"type,omitempty"`
	// the host path when using local storage provisioner
	HostPath *string `json:"hostPath,omitempty"`
	// the mount options of the storage class
	MountOptions []string `json:"mountOptions,omitempty"`
	// the description of the storage class
	Description string `json:"description"`
	// the display name of the storage class
	DisplayName string `json:"displayName"`
	// whether the storage class is enabled
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageClassCreate instantiates a new StorageClassCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageClassCreate(name string, reclaimPolicy StorageClassReclaimPolicy, volumeBindingMode StorageClassVolumeBindingMode, description string, displayName string, enabled bool) *StorageClassCreate {
	this := StorageClassCreate{}
	this.Name = name
	this.ReclaimPolicy = reclaimPolicy
	this.VolumeBindingMode = volumeBindingMode
	this.Description = description
	this.DisplayName = displayName
	this.Enabled = enabled
	return &this
}

// NewStorageClassCreateWithDefaults instantiates a new StorageClassCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageClassCreateWithDefaults() *StorageClassCreate {
	this := StorageClassCreate{}
	return &this
}

// GetName returns the Name field value.
func (o *StorageClassCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *StorageClassCreate) SetName(v string) {
	o.Name = v
}

// GetProvisioner returns the Provisioner field value if set, zero value otherwise.
func (o *StorageClassCreate) GetProvisioner() string {
	if o == nil || o.Provisioner == nil {
		var ret string
		return ret
	}
	return *o.Provisioner
}

// GetProvisionerOk returns a tuple with the Provisioner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetProvisionerOk() (*string, bool) {
	if o == nil || o.Provisioner == nil {
		return nil, false
	}
	return o.Provisioner, true
}

// HasProvisioner returns a boolean if a field has been set.
func (o *StorageClassCreate) HasProvisioner() bool {
	return o != nil && o.Provisioner != nil
}

// SetProvisioner gets a reference to the given string and assigns it to the Provisioner field.
func (o *StorageClassCreate) SetProvisioner(v string) {
	o.Provisioner = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *StorageClassCreate) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return &o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *StorageClassCreate) HasLabels() bool {
	return o != nil && o.Labels != nil
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *StorageClassCreate) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *StorageClassCreate) GetAnnotations() map[string]string {
	if o == nil || o.Annotations == nil {
		var ret map[string]string
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *StorageClassCreate) HasAnnotations() bool {
	return o != nil && o.Annotations != nil
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *StorageClassCreate) SetAnnotations(v map[string]string) {
	o.Annotations = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *StorageClassCreate) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *StorageClassCreate) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *StorageClassCreate) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetReclaimPolicy returns the ReclaimPolicy field value.
func (o *StorageClassCreate) GetReclaimPolicy() StorageClassReclaimPolicy {
	if o == nil {
		var ret StorageClassReclaimPolicy
		return ret
	}
	return o.ReclaimPolicy
}

// GetReclaimPolicyOk returns a tuple with the ReclaimPolicy field value
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetReclaimPolicyOk() (*StorageClassReclaimPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReclaimPolicy, true
}

// SetReclaimPolicy sets field value.
func (o *StorageClassCreate) SetReclaimPolicy(v StorageClassReclaimPolicy) {
	o.ReclaimPolicy = v
}

// GetAllowVolumeExpansion returns the AllowVolumeExpansion field value if set, zero value otherwise.
func (o *StorageClassCreate) GetAllowVolumeExpansion() bool {
	if o == nil || o.AllowVolumeExpansion == nil {
		var ret bool
		return ret
	}
	return *o.AllowVolumeExpansion
}

// GetAllowVolumeExpansionOk returns a tuple with the AllowVolumeExpansion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetAllowVolumeExpansionOk() (*bool, bool) {
	if o == nil || o.AllowVolumeExpansion == nil {
		return nil, false
	}
	return o.AllowVolumeExpansion, true
}

// HasAllowVolumeExpansion returns a boolean if a field has been set.
func (o *StorageClassCreate) HasAllowVolumeExpansion() bool {
	return o != nil && o.AllowVolumeExpansion != nil
}

// SetAllowVolumeExpansion gets a reference to the given bool and assigns it to the AllowVolumeExpansion field.
func (o *StorageClassCreate) SetAllowVolumeExpansion(v bool) {
	o.AllowVolumeExpansion = &v
}

// GetVolumeBindingMode returns the VolumeBindingMode field value.
func (o *StorageClassCreate) GetVolumeBindingMode() StorageClassVolumeBindingMode {
	if o == nil {
		var ret StorageClassVolumeBindingMode
		return ret
	}
	return o.VolumeBindingMode
}

// GetVolumeBindingModeOk returns a tuple with the VolumeBindingMode field value
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetVolumeBindingModeOk() (*StorageClassVolumeBindingMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeBindingMode, true
}

// SetVolumeBindingMode sets field value.
func (o *StorageClassCreate) SetVolumeBindingMode(v StorageClassVolumeBindingMode) {
	o.VolumeBindingMode = v
}

// GetAllowClone returns the AllowClone field value if set, zero value otherwise.
func (o *StorageClassCreate) GetAllowClone() bool {
	if o == nil || o.AllowClone == nil {
		var ret bool
		return ret
	}
	return *o.AllowClone
}

// GetAllowCloneOk returns a tuple with the AllowClone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetAllowCloneOk() (*bool, bool) {
	if o == nil || o.AllowClone == nil {
		return nil, false
	}
	return o.AllowClone, true
}

// HasAllowClone returns a boolean if a field has been set.
func (o *StorageClassCreate) HasAllowClone() bool {
	return o != nil && o.AllowClone != nil
}

// SetAllowClone gets a reference to the given bool and assigns it to the AllowClone field.
func (o *StorageClassCreate) SetAllowClone(v bool) {
	o.AllowClone = &v
}

// GetAllowSnapshot returns the AllowSnapshot field value if set, zero value otherwise.
func (o *StorageClassCreate) GetAllowSnapshot() bool {
	if o == nil || o.AllowSnapshot == nil {
		var ret bool
		return ret
	}
	return *o.AllowSnapshot
}

// GetAllowSnapshotOk returns a tuple with the AllowSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetAllowSnapshotOk() (*bool, bool) {
	if o == nil || o.AllowSnapshot == nil {
		return nil, false
	}
	return o.AllowSnapshot, true
}

// HasAllowSnapshot returns a boolean if a field has been set.
func (o *StorageClassCreate) HasAllowSnapshot() bool {
	return o != nil && o.AllowSnapshot != nil
}

// SetAllowSnapshot gets a reference to the given bool and assigns it to the AllowSnapshot field.
func (o *StorageClassCreate) SetAllowSnapshot(v bool) {
	o.AllowSnapshot = &v
}

// GetIsDefaultClass returns the IsDefaultClass field value if set, zero value otherwise.
func (o *StorageClassCreate) GetIsDefaultClass() bool {
	if o == nil || o.IsDefaultClass == nil {
		var ret bool
		return ret
	}
	return *o.IsDefaultClass
}

// GetIsDefaultClassOk returns a tuple with the IsDefaultClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetIsDefaultClassOk() (*bool, bool) {
	if o == nil || o.IsDefaultClass == nil {
		return nil, false
	}
	return o.IsDefaultClass, true
}

// HasIsDefaultClass returns a boolean if a field has been set.
func (o *StorageClassCreate) HasIsDefaultClass() bool {
	return o != nil && o.IsDefaultClass != nil
}

// SetIsDefaultClass gets a reference to the given bool and assigns it to the IsDefaultClass field.
func (o *StorageClassCreate) SetIsDefaultClass(v bool) {
	o.IsDefaultClass = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StorageClassCreate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StorageClassCreate) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StorageClassCreate) SetType(v string) {
	o.Type = &v
}

// GetHostPath returns the HostPath field value if set, zero value otherwise.
func (o *StorageClassCreate) GetHostPath() string {
	if o == nil || o.HostPath == nil {
		var ret string
		return ret
	}
	return *o.HostPath
}

// GetHostPathOk returns a tuple with the HostPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetHostPathOk() (*string, bool) {
	if o == nil || o.HostPath == nil {
		return nil, false
	}
	return o.HostPath, true
}

// HasHostPath returns a boolean if a field has been set.
func (o *StorageClassCreate) HasHostPath() bool {
	return o != nil && o.HostPath != nil
}

// SetHostPath gets a reference to the given string and assigns it to the HostPath field.
func (o *StorageClassCreate) SetHostPath(v string) {
	o.HostPath = &v
}

// GetMountOptions returns the MountOptions field value if set, zero value otherwise.
func (o *StorageClassCreate) GetMountOptions() []string {
	if o == nil || o.MountOptions == nil {
		var ret []string
		return ret
	}
	return o.MountOptions
}

// GetMountOptionsOk returns a tuple with the MountOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetMountOptionsOk() (*[]string, bool) {
	if o == nil || o.MountOptions == nil {
		return nil, false
	}
	return &o.MountOptions, true
}

// HasMountOptions returns a boolean if a field has been set.
func (o *StorageClassCreate) HasMountOptions() bool {
	return o != nil && o.MountOptions != nil
}

// SetMountOptions gets a reference to the given []string and assigns it to the MountOptions field.
func (o *StorageClassCreate) SetMountOptions(v []string) {
	o.MountOptions = v
}

// GetDescription returns the Description field value.
func (o *StorageClassCreate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *StorageClassCreate) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value.
func (o *StorageClassCreate) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *StorageClassCreate) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEnabled returns the Enabled field value.
func (o *StorageClassCreate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *StorageClassCreate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *StorageClassCreate) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageClassCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Provisioner != nil {
		toSerialize["provisioner"] = o.Provisioner
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Annotations != nil {
		toSerialize["annotations"] = o.Annotations
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	toSerialize["reclaimPolicy"] = o.ReclaimPolicy
	if o.AllowVolumeExpansion != nil {
		toSerialize["allowVolumeExpansion"] = o.AllowVolumeExpansion
	}
	toSerialize["volumeBindingMode"] = o.VolumeBindingMode
	if o.AllowClone != nil {
		toSerialize["allowClone"] = o.AllowClone
	}
	if o.AllowSnapshot != nil {
		toSerialize["allowSnapshot"] = o.AllowSnapshot
	}
	if o.IsDefaultClass != nil {
		toSerialize["isDefaultClass"] = o.IsDefaultClass
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.HostPath != nil {
		toSerialize["hostPath"] = o.HostPath
	}
	if o.MountOptions != nil {
		toSerialize["mountOptions"] = o.MountOptions
	}
	toSerialize["description"] = o.Description
	toSerialize["displayName"] = o.DisplayName
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageClassCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                 *string                        `json:"name"`
		Provisioner          *string                        `json:"provisioner,omitempty"`
		Labels               map[string]string              `json:"labels,omitempty"`
		Annotations          map[string]string              `json:"annotations,omitempty"`
		Parameters           map[string]string              `json:"parameters,omitempty"`
		ReclaimPolicy        *StorageClassReclaimPolicy     `json:"reclaimPolicy"`
		AllowVolumeExpansion *bool                          `json:"allowVolumeExpansion,omitempty"`
		VolumeBindingMode    *StorageClassVolumeBindingMode `json:"volumeBindingMode"`
		AllowClone           *bool                          `json:"allowClone,omitempty"`
		AllowSnapshot        *bool                          `json:"allowSnapshot,omitempty"`
		IsDefaultClass       *bool                          `json:"isDefaultClass,omitempty"`
		Type                 *string                        `json:"type,omitempty"`
		HostPath             *string                        `json:"hostPath,omitempty"`
		MountOptions         []string                       `json:"mountOptions,omitempty"`
		Description          *string                        `json:"description"`
		DisplayName          *string                        `json:"displayName"`
		Enabled              *bool                          `json:"enabled"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ReclaimPolicy == nil {
		return fmt.Errorf("required field reclaimPolicy missing")
	}
	if all.VolumeBindingMode == nil {
		return fmt.Errorf("required field volumeBindingMode missing")
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
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "provisioner", "labels", "annotations", "parameters", "reclaimPolicy", "allowVolumeExpansion", "volumeBindingMode", "allowClone", "allowSnapshot", "isDefaultClass", "type", "hostPath", "mountOptions", "description", "displayName", "enabled"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	o.Provisioner = all.Provisioner
	o.Labels = all.Labels
	o.Annotations = all.Annotations
	o.Parameters = all.Parameters
	if !all.ReclaimPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.ReclaimPolicy = *all.ReclaimPolicy
	}
	o.AllowVolumeExpansion = all.AllowVolumeExpansion
	if !all.VolumeBindingMode.IsValid() {
		hasInvalidField = true
	} else {
		o.VolumeBindingMode = *all.VolumeBindingMode
	}
	o.AllowClone = all.AllowClone
	o.AllowSnapshot = all.AllowSnapshot
	o.IsDefaultClass = all.IsDefaultClass
	o.Type = all.Type
	o.HostPath = all.HostPath
	o.MountOptions = all.MountOptions
	o.Description = *all.Description
	o.DisplayName = *all.DisplayName
	o.Enabled = *all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
