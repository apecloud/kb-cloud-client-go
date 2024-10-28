// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
	"github.com/google/uuid"
)

// BackupRepo backupRepo is the payload for KubeBlocks cluster backup repo
type BackupRepo struct {
	// the access method for backup repo
	AccessMethod BackupRepoAccessMethod `json:"accessMethod"`
	// the id of backup repo
	Id *string `json:"id,omitempty"`
	// backupNums specifies the number of backups in the backupRepo
	BackupNums *int32 `json:"backupNums,omitempty"`
	// config specifies the configuration of the backupRepo
	Config map[string]string `json:"config"`
	// createdAt specifies the creation time of the backupRepo
	CreatedAt time.Time `json:"createdAt"`
	// default specifies whether the backupRepo is the default backupRepo
	Default bool `json:"default"`
	// environmentId of the backupRepo
	EnvironmentId uuid.UUID `json:"environmentId"`
	// environmentName of the backupRepo
	EnvironmentName string `json:"environmentName"`
	// name of the backupRepo
	Name string `json:"name"`
	// status specifies the status of the backupRepo
	Status string `json:"status"`
	// the id of storage used by backup repo
	StorageId *string `json:"storageID,omitempty"`
	// storageProvider specifies the storage provider of the backupRepo
	StorageProvider string `json:"storageProvider"`
	// totalSize specifies the total size of backups in the backupRepo
	TotalSize *string `json:"totalSize,omitempty"`
	// failedReason specifies the reason of the backupRepo failure
	FailedReason *string `json:"failedReason,omitempty"`
	// failedMessage specifies the message of the backupRepo failure
	FailedMessage *string `json:"failedMessage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBackupRepo instantiates a new BackupRepo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBackupRepo(accessMethod BackupRepoAccessMethod, config map[string]string, createdAt time.Time, defaultVar bool, environmentId uuid.UUID, environmentName string, name string, status string, storageProvider string) *BackupRepo {
	this := BackupRepo{}
	this.AccessMethod = accessMethod
	this.Config = config
	this.CreatedAt = createdAt
	this.Default = defaultVar
	this.EnvironmentId = environmentId
	this.EnvironmentName = environmentName
	this.Name = name
	this.Status = status
	this.StorageProvider = storageProvider
	return &this
}

// NewBackupRepoWithDefaults instantiates a new BackupRepo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBackupRepoWithDefaults() *BackupRepo {
	this := BackupRepo{}
	var accessMethod BackupRepoAccessMethod = BackupRepoAccessMethodTool
	this.AccessMethod = accessMethod
	return &this
}

// GetAccessMethod returns the AccessMethod field value.
func (o *BackupRepo) GetAccessMethod() BackupRepoAccessMethod {
	if o == nil {
		var ret BackupRepoAccessMethod
		return ret
	}
	return o.AccessMethod
}

// GetAccessMethodOk returns a tuple with the AccessMethod field value
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetAccessMethodOk() (*BackupRepoAccessMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessMethod, true
}

// SetAccessMethod sets field value.
func (o *BackupRepo) SetAccessMethod(v BackupRepoAccessMethod) {
	o.AccessMethod = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BackupRepo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BackupRepo) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BackupRepo) SetId(v string) {
	o.Id = &v
}

// GetBackupNums returns the BackupNums field value if set, zero value otherwise.
func (o *BackupRepo) GetBackupNums() int32 {
	if o == nil || o.BackupNums == nil {
		var ret int32
		return ret
	}
	return *o.BackupNums
}

// GetBackupNumsOk returns a tuple with the BackupNums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetBackupNumsOk() (*int32, bool) {
	if o == nil || o.BackupNums == nil {
		return nil, false
	}
	return o.BackupNums, true
}

// HasBackupNums returns a boolean if a field has been set.
func (o *BackupRepo) HasBackupNums() bool {
	return o != nil && o.BackupNums != nil
}

// SetBackupNums gets a reference to the given int32 and assigns it to the BackupNums field.
func (o *BackupRepo) SetBackupNums(v int32) {
	o.BackupNums = &v
}

// GetConfig returns the Config field value.
func (o *BackupRepo) GetConfig() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetConfigOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value.
func (o *BackupRepo) SetConfig(v map[string]string) {
	o.Config = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *BackupRepo) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *BackupRepo) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDefault returns the Default field value.
func (o *BackupRepo) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value.
func (o *BackupRepo) SetDefault(v bool) {
	o.Default = v
}

// GetEnvironmentId returns the EnvironmentId field value.
func (o *BackupRepo) GetEnvironmentId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value.
func (o *BackupRepo) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId = v
}

// GetEnvironmentName returns the EnvironmentName field value.
func (o *BackupRepo) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value.
func (o *BackupRepo) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetName returns the Name field value.
func (o *BackupRepo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *BackupRepo) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value.
func (o *BackupRepo) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *BackupRepo) SetStatus(v string) {
	o.Status = v
}

// GetStorageId returns the StorageId field value if set, zero value otherwise.
func (o *BackupRepo) GetStorageId() string {
	if o == nil || o.StorageId == nil {
		var ret string
		return ret
	}
	return *o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetStorageIdOk() (*string, bool) {
	if o == nil || o.StorageId == nil {
		return nil, false
	}
	return o.StorageId, true
}

// HasStorageId returns a boolean if a field has been set.
func (o *BackupRepo) HasStorageId() bool {
	return o != nil && o.StorageId != nil
}

// SetStorageId gets a reference to the given string and assigns it to the StorageId field.
func (o *BackupRepo) SetStorageId(v string) {
	o.StorageId = &v
}

// GetStorageProvider returns the StorageProvider field value.
func (o *BackupRepo) GetStorageProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageProvider
}

// GetStorageProviderOk returns a tuple with the StorageProvider field value
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetStorageProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageProvider, true
}

// SetStorageProvider sets field value.
func (o *BackupRepo) SetStorageProvider(v string) {
	o.StorageProvider = v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *BackupRepo) GetTotalSize() string {
	if o == nil || o.TotalSize == nil {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetTotalSizeOk() (*string, bool) {
	if o == nil || o.TotalSize == nil {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *BackupRepo) HasTotalSize() bool {
	return o != nil && o.TotalSize != nil
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *BackupRepo) SetTotalSize(v string) {
	o.TotalSize = &v
}

// GetFailedReason returns the FailedReason field value if set, zero value otherwise.
func (o *BackupRepo) GetFailedReason() string {
	if o == nil || o.FailedReason == nil {
		var ret string
		return ret
	}
	return *o.FailedReason
}

// GetFailedReasonOk returns a tuple with the FailedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetFailedReasonOk() (*string, bool) {
	if o == nil || o.FailedReason == nil {
		return nil, false
	}
	return o.FailedReason, true
}

// HasFailedReason returns a boolean if a field has been set.
func (o *BackupRepo) HasFailedReason() bool {
	return o != nil && o.FailedReason != nil
}

// SetFailedReason gets a reference to the given string and assigns it to the FailedReason field.
func (o *BackupRepo) SetFailedReason(v string) {
	o.FailedReason = &v
}

// GetFailedMessage returns the FailedMessage field value if set, zero value otherwise.
func (o *BackupRepo) GetFailedMessage() string {
	if o == nil || o.FailedMessage == nil {
		var ret string
		return ret
	}
	return *o.FailedMessage
}

// GetFailedMessageOk returns a tuple with the FailedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRepo) GetFailedMessageOk() (*string, bool) {
	if o == nil || o.FailedMessage == nil {
		return nil, false
	}
	return o.FailedMessage, true
}

// HasFailedMessage returns a boolean if a field has been set.
func (o *BackupRepo) HasFailedMessage() bool {
	return o != nil && o.FailedMessage != nil
}

// SetFailedMessage gets a reference to the given string and assigns it to the FailedMessage field.
func (o *BackupRepo) SetFailedMessage(v string) {
	o.FailedMessage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BackupRepo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["accessMethod"] = o.AccessMethod
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.BackupNums != nil {
		toSerialize["backupNums"] = o.BackupNums
	}
	toSerialize["config"] = o.Config
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["default"] = o.Default
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["environmentName"] = o.EnvironmentName
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	if o.StorageId != nil {
		toSerialize["storageID"] = o.StorageId
	}
	toSerialize["storageProvider"] = o.StorageProvider
	if o.TotalSize != nil {
		toSerialize["totalSize"] = o.TotalSize
	}
	if o.FailedReason != nil {
		toSerialize["failedReason"] = o.FailedReason
	}
	if o.FailedMessage != nil {
		toSerialize["failedMessage"] = o.FailedMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BackupRepo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessMethod    *BackupRepoAccessMethod `json:"accessMethod"`
		Id              *string                 `json:"id,omitempty"`
		BackupNums      *int32                  `json:"backupNums,omitempty"`
		Config          *map[string]string      `json:"config"`
		CreatedAt       *time.Time              `json:"createdAt"`
		Default         *bool                   `json:"default"`
		EnvironmentId   *uuid.UUID              `json:"environmentId"`
		EnvironmentName *string                 `json:"environmentName"`
		Name            *string                 `json:"name"`
		Status          *string                 `json:"status"`
		StorageId       *string                 `json:"storageID,omitempty"`
		StorageProvider *string                 `json:"storageProvider"`
		TotalSize       *string                 `json:"totalSize,omitempty"`
		FailedReason    *string                 `json:"failedReason,omitempty"`
		FailedMessage   *string                 `json:"failedMessage,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessMethod == nil {
		return fmt.Errorf("required field accessMethod missing")
	}
	if all.Config == nil {
		return fmt.Errorf("required field config missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.Default == nil {
		return fmt.Errorf("required field default missing")
	}
	if all.EnvironmentId == nil {
		return fmt.Errorf("required field environmentId missing")
	}
	if all.EnvironmentName == nil {
		return fmt.Errorf("required field environmentName missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.StorageProvider == nil {
		return fmt.Errorf("required field storageProvider missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"accessMethod", "id", "backupNums", "config", "createdAt", "default", "environmentId", "environmentName", "name", "status", "storageID", "storageProvider", "totalSize", "failedReason", "failedMessage"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AccessMethod.IsValid() {
		hasInvalidField = true
	} else {
		o.AccessMethod = *all.AccessMethod
	}
	o.Id = all.Id
	o.BackupNums = all.BackupNums
	o.Config = *all.Config
	o.CreatedAt = *all.CreatedAt
	o.Default = *all.Default
	o.EnvironmentId = *all.EnvironmentId
	o.EnvironmentName = *all.EnvironmentName
	o.Name = *all.Name
	o.Status = *all.Status
	o.StorageId = all.StorageId
	o.StorageProvider = *all.StorageProvider
	o.TotalSize = all.TotalSize
	o.FailedReason = all.FailedReason
	o.FailedMessage = all.FailedMessage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
