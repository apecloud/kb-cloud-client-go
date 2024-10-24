// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RestoreCreate RestoreCreate is the payload to restore a KubeBlocks cluster
// NODESCRIPTION RestoreCreate
//
// Deprecated: This model is deprecated.
type RestoreCreate struct {
	// the env name of the target environment to restore
	EnvironmentName string `json:"environmentName"`
	// the id of backup to restore
	BackupId string `json:"backupId"`
	// KubeBlocks cluster information
	Cluster Cluster `json:"cluster"`
	// restoreTime point to restore
	RestoreTimeStr *string `json:"restoreTimeStr,omitempty"`
	// the recovery process in the PostReady phase will be performed after the cluster is running successfully.
	DoReadyRestoreAfterClusterRunning *bool `json:"doReadyRestoreAfterClusterRunning,omitempty"`
	// the volume claim restore policy, support values: [Serial, Parallel]
	VolumeRestorePolicy *VolumeRestorePolicy `json:"volumeRestorePolicy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestoreCreate instantiates a new RestoreCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestoreCreate(environmentName string, backupId string, cluster Cluster) *RestoreCreate {
	this := RestoreCreate{}
	this.EnvironmentName = environmentName
	this.BackupId = backupId
	this.Cluster = cluster
	var volumeRestorePolicy VolumeRestorePolicy = VOLUMERESTOREPOLICY_SERIAL
	this.VolumeRestorePolicy = &volumeRestorePolicy
	return &this
}

// NewRestoreCreateWithDefaults instantiates a new RestoreCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestoreCreateWithDefaults() *RestoreCreate {
	this := RestoreCreate{}
	var volumeRestorePolicy VolumeRestorePolicy = VOLUMERESTOREPOLICY_SERIAL
	this.VolumeRestorePolicy = &volumeRestorePolicy
	return &this
}

// GetEnvironmentName returns the EnvironmentName field value.
func (o *RestoreCreate) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *RestoreCreate) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *RestoreCreate) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetBackupId returns the BackupId field value.
func (o *RestoreCreate) GetBackupId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value
// and a boolean to check if the value has been set.
func (o *RestoreCreate) GetBackupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupId, true
}

// SetBackupId sets field value.
func (o *RestoreCreate) SetBackupId(v string) {
	o.BackupId = v
}

// GetCluster returns the Cluster field value.
func (o *RestoreCreate) GetCluster() Cluster {
	if o == nil {
		var ret Cluster
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *RestoreCreate) GetClusterOk() (*Cluster, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *RestoreCreate) SetCluster(v Cluster) {
	o.Cluster = v
}

// GetRestoreTimeStr returns the RestoreTimeStr field value if set, zero value otherwise.
func (o *RestoreCreate) GetRestoreTimeStr() string {
	if o == nil || o.RestoreTimeStr == nil {
		var ret string
		return ret
	}
	return *o.RestoreTimeStr
}

// GetRestoreTimeStrOk returns a tuple with the RestoreTimeStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreCreate) GetRestoreTimeStrOk() (*string, bool) {
	if o == nil || o.RestoreTimeStr == nil {
		return nil, false
	}
	return o.RestoreTimeStr, true
}

// HasRestoreTimeStr returns a boolean if a field has been set.
func (o *RestoreCreate) HasRestoreTimeStr() bool {
	return o != nil && o.RestoreTimeStr != nil
}

// SetRestoreTimeStr gets a reference to the given string and assigns it to the RestoreTimeStr field.
func (o *RestoreCreate) SetRestoreTimeStr(v string) {
	o.RestoreTimeStr = &v
}

// GetDoReadyRestoreAfterClusterRunning returns the DoReadyRestoreAfterClusterRunning field value if set, zero value otherwise.
func (o *RestoreCreate) GetDoReadyRestoreAfterClusterRunning() bool {
	if o == nil || o.DoReadyRestoreAfterClusterRunning == nil {
		var ret bool
		return ret
	}
	return *o.DoReadyRestoreAfterClusterRunning
}

// GetDoReadyRestoreAfterClusterRunningOk returns a tuple with the DoReadyRestoreAfterClusterRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreCreate) GetDoReadyRestoreAfterClusterRunningOk() (*bool, bool) {
	if o == nil || o.DoReadyRestoreAfterClusterRunning == nil {
		return nil, false
	}
	return o.DoReadyRestoreAfterClusterRunning, true
}

// HasDoReadyRestoreAfterClusterRunning returns a boolean if a field has been set.
func (o *RestoreCreate) HasDoReadyRestoreAfterClusterRunning() bool {
	return o != nil && o.DoReadyRestoreAfterClusterRunning != nil
}

// SetDoReadyRestoreAfterClusterRunning gets a reference to the given bool and assigns it to the DoReadyRestoreAfterClusterRunning field.
func (o *RestoreCreate) SetDoReadyRestoreAfterClusterRunning(v bool) {
	o.DoReadyRestoreAfterClusterRunning = &v
}

// GetVolumeRestorePolicy returns the VolumeRestorePolicy field value if set, zero value otherwise.
func (o *RestoreCreate) GetVolumeRestorePolicy() VolumeRestorePolicy {
	if o == nil || o.VolumeRestorePolicy == nil {
		var ret VolumeRestorePolicy
		return ret
	}
	return *o.VolumeRestorePolicy
}

// GetVolumeRestorePolicyOk returns a tuple with the VolumeRestorePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreCreate) GetVolumeRestorePolicyOk() (*VolumeRestorePolicy, bool) {
	if o == nil || o.VolumeRestorePolicy == nil {
		return nil, false
	}
	return o.VolumeRestorePolicy, true
}

// HasVolumeRestorePolicy returns a boolean if a field has been set.
func (o *RestoreCreate) HasVolumeRestorePolicy() bool {
	return o != nil && o.VolumeRestorePolicy != nil
}

// SetVolumeRestorePolicy gets a reference to the given VolumeRestorePolicy and assigns it to the VolumeRestorePolicy field.
func (o *RestoreCreate) SetVolumeRestorePolicy(v VolumeRestorePolicy) {
	o.VolumeRestorePolicy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestoreCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["environmentName"] = o.EnvironmentName
	toSerialize["backupId"] = o.BackupId
	toSerialize["cluster"] = o.Cluster
	if o.RestoreTimeStr != nil {
		toSerialize["restoreTimeStr"] = o.RestoreTimeStr
	}
	if o.DoReadyRestoreAfterClusterRunning != nil {
		toSerialize["doReadyRestoreAfterClusterRunning"] = o.DoReadyRestoreAfterClusterRunning
	}
	if o.VolumeRestorePolicy != nil {
		toSerialize["volumeRestorePolicy"] = o.VolumeRestorePolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestoreCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EnvironmentName                   *string              `json:"environmentName"`
		BackupId                          *string              `json:"backupId"`
		Cluster                           *Cluster             `json:"cluster"`
		RestoreTimeStr                    *string              `json:"restoreTimeStr,omitempty"`
		DoReadyRestoreAfterClusterRunning *bool                `json:"doReadyRestoreAfterClusterRunning,omitempty"`
		VolumeRestorePolicy               *VolumeRestorePolicy `json:"volumeRestorePolicy,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EnvironmentName == nil {
		return fmt.Errorf("required field environmentName missing")
	}
	if all.BackupId == nil {
		return fmt.Errorf("required field backupId missing")
	}
	if all.Cluster == nil {
		return fmt.Errorf("required field cluster missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"environmentName", "backupId", "cluster", "restoreTimeStr", "doReadyRestoreAfterClusterRunning", "volumeRestorePolicy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EnvironmentName = *all.EnvironmentName
	o.BackupId = *all.BackupId
	if all.Cluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cluster = *all.Cluster
	o.RestoreTimeStr = all.RestoreTimeStr
	o.DoReadyRestoreAfterClusterRunning = all.DoReadyRestoreAfterClusterRunning
	if all.VolumeRestorePolicy != nil && !all.VolumeRestorePolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.VolumeRestorePolicy = all.VolumeRestorePolicy
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
