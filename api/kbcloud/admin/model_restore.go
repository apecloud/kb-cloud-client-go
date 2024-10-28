// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Restore create a KubeBlocks restore API

type Restore struct {
	// organization name
	OrgName *string `json:"orgName,omitempty"`
	// backup name
	BackupName string `json:"backupName"`
	// kubeBlocks cluster name
	ClusterName string `json:"clusterName"`
	// component name of cluster
	ComponentName string `json:"componentName"`
	// target pod name
	TargetPodName *string `json:"targetPodName,omitempty"`
	// NODESCRIPTION CreatedAt
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// kubeBlocks restore name
	Name *string `json:"name,omitempty"`
	// restore parameters to inject env of the restore CR.
	Parameters *RestoreParameters `json:"parameters,omitempty"`
	// restore time
	RestoreTime *string `json:"restoreTime,omitempty"`
	// restore status
	Status *RestoreStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestore instantiates a new Restore object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestore(backupName string, clusterName string, componentName string) *Restore {
	this := Restore{}
	this.BackupName = backupName
	this.ClusterName = clusterName
	this.ComponentName = componentName
	return &this
}

// NewRestoreWithDefaults instantiates a new Restore object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestoreWithDefaults() *Restore {
	this := Restore{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Restore) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restore) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Restore) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Restore) SetOrgName(v string) {
	o.OrgName = &v
}

// GetBackupName returns the BackupName field value.
func (o *Restore) GetBackupName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BackupName
}

// GetBackupNameOk returns a tuple with the BackupName field value
// and a boolean to check if the value has been set.
func (o *Restore) GetBackupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackupName, true
}

// SetBackupName sets field value.
func (o *Restore) SetBackupName(v string) {
	o.BackupName = v
}

// GetClusterName returns the ClusterName field value.
func (o *Restore) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *Restore) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value.
func (o *Restore) SetClusterName(v string) {
	o.ClusterName = v
}

// GetComponentName returns the ComponentName field value.
func (o *Restore) GetComponentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value
// and a boolean to check if the value has been set.
func (o *Restore) GetComponentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentName, true
}

// SetComponentName sets field value.
func (o *Restore) SetComponentName(v string) {
	o.ComponentName = v
}

// GetTargetPodName returns the TargetPodName field value if set, zero value otherwise.
func (o *Restore) GetTargetPodName() string {
	if o == nil || o.TargetPodName == nil {
		var ret string
		return ret
	}
	return *o.TargetPodName
}

// GetTargetPodNameOk returns a tuple with the TargetPodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restore) GetTargetPodNameOk() (*string, bool) {
	if o == nil || o.TargetPodName == nil {
		return nil, false
	}
	return o.TargetPodName, true
}

// HasTargetPodName returns a boolean if a field has been set.
func (o *Restore) HasTargetPodName() bool {
	return o != nil && o.TargetPodName != nil
}

// SetTargetPodName gets a reference to the given string and assigns it to the TargetPodName field.
func (o *Restore) SetTargetPodName(v string) {
	o.TargetPodName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Restore) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restore) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Restore) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Restore) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Restore) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restore) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Restore) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Restore) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Restore) GetParameters() RestoreParameters {
	if o == nil || o.Parameters == nil {
		var ret RestoreParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restore) GetParametersOk() (*RestoreParameters, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Restore) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given RestoreParameters and assigns it to the Parameters field.
func (o *Restore) SetParameters(v RestoreParameters) {
	o.Parameters = &v
}

// GetRestoreTime returns the RestoreTime field value if set, zero value otherwise.
func (o *Restore) GetRestoreTime() string {
	if o == nil || o.RestoreTime == nil {
		var ret string
		return ret
	}
	return *o.RestoreTime
}

// GetRestoreTimeOk returns a tuple with the RestoreTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restore) GetRestoreTimeOk() (*string, bool) {
	if o == nil || o.RestoreTime == nil {
		return nil, false
	}
	return o.RestoreTime, true
}

// HasRestoreTime returns a boolean if a field has been set.
func (o *Restore) HasRestoreTime() bool {
	return o != nil && o.RestoreTime != nil
}

// SetRestoreTime gets a reference to the given string and assigns it to the RestoreTime field.
func (o *Restore) SetRestoreTime(v string) {
	o.RestoreTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Restore) GetStatus() RestoreStatus {
	if o == nil || o.Status == nil {
		var ret RestoreStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restore) GetStatusOk() (*RestoreStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Restore) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given RestoreStatus and assigns it to the Status field.
func (o *Restore) SetStatus(v RestoreStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Restore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	toSerialize["backupName"] = o.BackupName
	toSerialize["clusterName"] = o.ClusterName
	toSerialize["componentName"] = o.ComponentName
	if o.TargetPodName != nil {
		toSerialize["targetPodName"] = o.TargetPodName
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.RestoreTime != nil {
		toSerialize["restoreTime"] = o.RestoreTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Restore) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName       *string            `json:"orgName,omitempty"`
		BackupName    *string            `json:"backupName"`
		ClusterName   *string            `json:"clusterName"`
		ComponentName *string            `json:"componentName"`
		TargetPodName *string            `json:"targetPodName,omitempty"`
		CreatedAt     *time.Time         `json:"createdAt,omitempty"`
		Name          *string            `json:"name,omitempty"`
		Parameters    *RestoreParameters `json:"parameters,omitempty"`
		RestoreTime   *string            `json:"restoreTime,omitempty"`
		Status        *RestoreStatus     `json:"status,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BackupName == nil {
		return fmt.Errorf("required field backupName missing")
	}
	if all.ClusterName == nil {
		return fmt.Errorf("required field clusterName missing")
	}
	if all.ComponentName == nil {
		return fmt.Errorf("required field componentName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgName", "backupName", "clusterName", "componentName", "targetPodName", "createdAt", "name", "parameters", "restoreTime", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.OrgName = all.OrgName
	o.BackupName = *all.BackupName
	o.ClusterName = *all.ClusterName
	o.ComponentName = *all.ComponentName
	o.TargetPodName = all.TargetPodName
	o.CreatedAt = all.CreatedAt
	o.Name = all.Name
	if all.Parameters != nil && all.Parameters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Parameters = all.Parameters
	o.RestoreTime = all.RestoreTime
	if all.Status != nil && all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
