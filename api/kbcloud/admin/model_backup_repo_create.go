// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// BackupRepoCreate BackupRepoCreate is the payload to create a KubeBlocks cluster backup repo
type BackupRepoCreate struct {
	// the id of storage that backup repo used
	StorageId string `json:"storageID"`
	// the access method for backup repo
	AccessMethod *BackupRepoAccessMethod `json:"accessMethod,omitempty"`
	// default specifies whether the backupRepo is the default backupRepo
	Default *bool `json:"default,omitempty"`
	// name of the backupRepo
	Name string `json:"name"`
	// the parameters to create the storage
	Params map[string]string `json:"params,omitempty"`
	// Specify the reclaim policy for PVs created by this backup repo
	PvReclaimPolicy *BackupRepoPVReclaimPolicy `json:"pvReclaimPolicy,omitempty"`
	// Specify the capacity of the new created PVC
	VolumeCapacity *string `json:"volumeCapacity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupRepoCreate instantiates a new BackupRepoCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupRepoCreate(storageId string, name string) *BackupRepoCreate {
	this := BackupRepoCreate{}
	this.StorageId = storageId
	var accessMethod BackupRepoAccessMethod = BackupRepoAccessMethodTool
	this.AccessMethod = &accessMethod
	this.Name = name
	return &this
}

// NewBackupRepoCreateWithDefaults instantiates a new BackupRepoCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupRepoCreateWithDefaults() *BackupRepoCreate {
	this := BackupRepoCreate{}
	var accessMethod BackupRepoAccessMethod = BackupRepoAccessMethodTool
	this.AccessMethod = &accessMethod
	return &this
}

// GetStorageId returns the StorageId field value.
func (o *BackupRepoCreate) GetStorageId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value
// and a boolean to check if the value has been set.
func (o *BackupRepoCreate) GetStorageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageId, true
}

// SetStorageId sets field value.
func (o *BackupRepoCreate) SetStorageId(v string) {
	o.StorageId = v
}

// GetAccessMethod returns the AccessMethod field value if set, zero value otherwise.
func (o *BackupRepoCreate) GetAccessMethod() BackupRepoAccessMethod {
	if o == nil || o.AccessMethod == nil {
		var ret BackupRepoAccessMethod
		return ret
	}
	return *o.AccessMethod
}

// GetAccessMethodOk returns a tuple with the AccessMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoCreate) GetAccessMethodOk() (*BackupRepoAccessMethod, bool) {
	if o == nil || o.AccessMethod == nil {
		return nil, false
	}
	return o.AccessMethod, true
}

// HasAccessMethod returns a boolean if a field has been set.
func (o *BackupRepoCreate) HasAccessMethod() bool {
	return o != nil && o.AccessMethod != nil
}

// SetAccessMethod gets a reference to the given BackupRepoAccessMethod and assigns it to the AccessMethod field.
func (o *BackupRepoCreate) SetAccessMethod(v BackupRepoAccessMethod) {
	o.AccessMethod = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *BackupRepoCreate) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoCreate) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *BackupRepoCreate) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *BackupRepoCreate) SetDefault(v bool) {
	o.Default = &v
}

// GetName returns the Name field value.
func (o *BackupRepoCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BackupRepoCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *BackupRepoCreate) SetName(v string) {
	o.Name = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *BackupRepoCreate) GetParams() map[string]string {
	if o == nil || o.Params == nil {
		var ret map[string]string
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoCreate) GetParamsOk() (*map[string]string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return &o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *BackupRepoCreate) HasParams() bool {
	return o != nil && o.Params != nil
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *BackupRepoCreate) SetParams(v map[string]string) {
	o.Params = v
}

// GetPvReclaimPolicy returns the PvReclaimPolicy field value if set, zero value otherwise.
func (o *BackupRepoCreate) GetPvReclaimPolicy() BackupRepoPVReclaimPolicy {
	if o == nil || o.PvReclaimPolicy == nil {
		var ret BackupRepoPVReclaimPolicy
		return ret
	}
	return *o.PvReclaimPolicy
}

// GetPvReclaimPolicyOk returns a tuple with the PvReclaimPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoCreate) GetPvReclaimPolicyOk() (*BackupRepoPVReclaimPolicy, bool) {
	if o == nil || o.PvReclaimPolicy == nil {
		return nil, false
	}
	return o.PvReclaimPolicy, true
}

// HasPvReclaimPolicy returns a boolean if a field has been set.
func (o *BackupRepoCreate) HasPvReclaimPolicy() bool {
	return o != nil && o.PvReclaimPolicy != nil
}

// SetPvReclaimPolicy gets a reference to the given BackupRepoPVReclaimPolicy and assigns it to the PvReclaimPolicy field.
func (o *BackupRepoCreate) SetPvReclaimPolicy(v BackupRepoPVReclaimPolicy) {
	o.PvReclaimPolicy = &v
}

// GetVolumeCapacity returns the VolumeCapacity field value if set, zero value otherwise.
func (o *BackupRepoCreate) GetVolumeCapacity() string {
	if o == nil || o.VolumeCapacity == nil {
		var ret string
		return ret
	}
	return *o.VolumeCapacity
}

// GetVolumeCapacityOk returns a tuple with the VolumeCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepoCreate) GetVolumeCapacityOk() (*string, bool) {
	if o == nil || o.VolumeCapacity == nil {
		return nil, false
	}
	return o.VolumeCapacity, true
}

// HasVolumeCapacity returns a boolean if a field has been set.
func (o *BackupRepoCreate) HasVolumeCapacity() bool {
	return o != nil && o.VolumeCapacity != nil
}

// SetVolumeCapacity gets a reference to the given string and assigns it to the VolumeCapacity field.
func (o *BackupRepoCreate) SetVolumeCapacity(v string) {
	o.VolumeCapacity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupRepoCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["storageID"] = o.StorageId
	if o.AccessMethod != nil {
		toSerialize["accessMethod"] = o.AccessMethod
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	toSerialize["name"] = o.Name
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.PvReclaimPolicy != nil {
		toSerialize["pvReclaimPolicy"] = o.PvReclaimPolicy
	}
	if o.VolumeCapacity != nil {
		toSerialize["volumeCapacity"] = o.VolumeCapacity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupRepoCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StorageId       *string                    `json:"storageID"`
		AccessMethod    *BackupRepoAccessMethod    `json:"accessMethod,omitempty"`
		Default         *bool                      `json:"default,omitempty"`
		Name            *string                    `json:"name"`
		Params          map[string]string          `json:"params,omitempty"`
		PvReclaimPolicy *BackupRepoPVReclaimPolicy `json:"pvReclaimPolicy,omitempty"`
		VolumeCapacity  *string                    `json:"volumeCapacity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.StorageId == nil {
		return fmt.Errorf("required field storageID missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"storageID", "accessMethod", "default", "name", "params", "pvReclaimPolicy", "volumeCapacity"})
	} else {
		return err
	}

	hasInvalidField := false
	o.StorageId = *all.StorageId
	if all.AccessMethod != nil && !all.AccessMethod.IsValid() {
		hasInvalidField = true
	} else {
		o.AccessMethod = all.AccessMethod
	}
	o.Default = all.Default
	o.Name = *all.Name
	o.Params = all.Params
	if all.PvReclaimPolicy != nil && !all.PvReclaimPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.PvReclaimPolicy = all.PvReclaimPolicy
	}
	o.VolumeCapacity = all.VolumeCapacity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
