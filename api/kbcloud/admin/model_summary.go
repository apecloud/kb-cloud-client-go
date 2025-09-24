// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Summary struct {
	// The namespace of cluster
	Namespace *string `json:"namespace,omitempty"`
	// The name of cluster
	Name *string `json:"name,omitempty"`
	// The CPU requests in configuration
	CpuRequests *string `json:"cpuRequests,omitempty"`
	// The CPU limits in configuration
	CpuLimits *string `json:"cpuLimits,omitempty"`
	// The memory requests in configuration
	MemoryRequests *string `json:"memoryRequests,omitempty"`
	// The memory limits in configuration
	MemoryLimits *string `json:"memoryLimits,omitempty"`
	// The storage size of volume
	StorageSize *string `json:"storageSize,omitempty"`
	// The number of postgresql pods
	Replicas *int32 `json:"replicas,omitempty"`
	// The endpoint of backup
	BackupEndpoint *string `json:"backupEndpoint,omitempty"`
	// The path where to store backup
	BackupPath *string `json:"backupPath,omitempty"`
	// The backup schedule
	BackupSchedule *string `json:"backupSchedule,omitempty"`
	// The status of cluster
	Status *string `json:"status,omitempty"`
	// The creation time of cluster
	CreationTime *string `json:"creationTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSummary instantiates a new Summary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSummary() *Summary {
	this := Summary{}
	return &this
}

// NewSummaryWithDefaults instantiates a new Summary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSummaryWithDefaults() *Summary {
	this := Summary{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Summary) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Summary) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Summary) SetNamespace(v string) {
	o.Namespace = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Summary) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Summary) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Summary) SetName(v string) {
	o.Name = &v
}

// GetCpuRequests returns the CpuRequests field value if set, zero value otherwise.
func (o *Summary) GetCpuRequests() string {
	if o == nil || o.CpuRequests == nil {
		var ret string
		return ret
	}
	return *o.CpuRequests
}

// GetCpuRequestsOk returns a tuple with the CpuRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetCpuRequestsOk() (*string, bool) {
	if o == nil || o.CpuRequests == nil {
		return nil, false
	}
	return o.CpuRequests, true
}

// HasCpuRequests returns a boolean if a field has been set.
func (o *Summary) HasCpuRequests() bool {
	return o != nil && o.CpuRequests != nil
}

// SetCpuRequests gets a reference to the given string and assigns it to the CpuRequests field.
func (o *Summary) SetCpuRequests(v string) {
	o.CpuRequests = &v
}

// GetCpuLimits returns the CpuLimits field value if set, zero value otherwise.
func (o *Summary) GetCpuLimits() string {
	if o == nil || o.CpuLimits == nil {
		var ret string
		return ret
	}
	return *o.CpuLimits
}

// GetCpuLimitsOk returns a tuple with the CpuLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetCpuLimitsOk() (*string, bool) {
	if o == nil || o.CpuLimits == nil {
		return nil, false
	}
	return o.CpuLimits, true
}

// HasCpuLimits returns a boolean if a field has been set.
func (o *Summary) HasCpuLimits() bool {
	return o != nil && o.CpuLimits != nil
}

// SetCpuLimits gets a reference to the given string and assigns it to the CpuLimits field.
func (o *Summary) SetCpuLimits(v string) {
	o.CpuLimits = &v
}

// GetMemoryRequests returns the MemoryRequests field value if set, zero value otherwise.
func (o *Summary) GetMemoryRequests() string {
	if o == nil || o.MemoryRequests == nil {
		var ret string
		return ret
	}
	return *o.MemoryRequests
}

// GetMemoryRequestsOk returns a tuple with the MemoryRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetMemoryRequestsOk() (*string, bool) {
	if o == nil || o.MemoryRequests == nil {
		return nil, false
	}
	return o.MemoryRequests, true
}

// HasMemoryRequests returns a boolean if a field has been set.
func (o *Summary) HasMemoryRequests() bool {
	return o != nil && o.MemoryRequests != nil
}

// SetMemoryRequests gets a reference to the given string and assigns it to the MemoryRequests field.
func (o *Summary) SetMemoryRequests(v string) {
	o.MemoryRequests = &v
}

// GetMemoryLimits returns the MemoryLimits field value if set, zero value otherwise.
func (o *Summary) GetMemoryLimits() string {
	if o == nil || o.MemoryLimits == nil {
		var ret string
		return ret
	}
	return *o.MemoryLimits
}

// GetMemoryLimitsOk returns a tuple with the MemoryLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetMemoryLimitsOk() (*string, bool) {
	if o == nil || o.MemoryLimits == nil {
		return nil, false
	}
	return o.MemoryLimits, true
}

// HasMemoryLimits returns a boolean if a field has been set.
func (o *Summary) HasMemoryLimits() bool {
	return o != nil && o.MemoryLimits != nil
}

// SetMemoryLimits gets a reference to the given string and assigns it to the MemoryLimits field.
func (o *Summary) SetMemoryLimits(v string) {
	o.MemoryLimits = &v
}

// GetStorageSize returns the StorageSize field value if set, zero value otherwise.
func (o *Summary) GetStorageSize() string {
	if o == nil || o.StorageSize == nil {
		var ret string
		return ret
	}
	return *o.StorageSize
}

// GetStorageSizeOk returns a tuple with the StorageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetStorageSizeOk() (*string, bool) {
	if o == nil || o.StorageSize == nil {
		return nil, false
	}
	return o.StorageSize, true
}

// HasStorageSize returns a boolean if a field has been set.
func (o *Summary) HasStorageSize() bool {
	return o != nil && o.StorageSize != nil
}

// SetStorageSize gets a reference to the given string and assigns it to the StorageSize field.
func (o *Summary) SetStorageSize(v string) {
	o.StorageSize = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *Summary) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *Summary) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *Summary) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetBackupEndpoint returns the BackupEndpoint field value if set, zero value otherwise.
func (o *Summary) GetBackupEndpoint() string {
	if o == nil || o.BackupEndpoint == nil {
		var ret string
		return ret
	}
	return *o.BackupEndpoint
}

// GetBackupEndpointOk returns a tuple with the BackupEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetBackupEndpointOk() (*string, bool) {
	if o == nil || o.BackupEndpoint == nil {
		return nil, false
	}
	return o.BackupEndpoint, true
}

// HasBackupEndpoint returns a boolean if a field has been set.
func (o *Summary) HasBackupEndpoint() bool {
	return o != nil && o.BackupEndpoint != nil
}

// SetBackupEndpoint gets a reference to the given string and assigns it to the BackupEndpoint field.
func (o *Summary) SetBackupEndpoint(v string) {
	o.BackupEndpoint = &v
}

// GetBackupPath returns the BackupPath field value if set, zero value otherwise.
func (o *Summary) GetBackupPath() string {
	if o == nil || o.BackupPath == nil {
		var ret string
		return ret
	}
	return *o.BackupPath
}

// GetBackupPathOk returns a tuple with the BackupPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetBackupPathOk() (*string, bool) {
	if o == nil || o.BackupPath == nil {
		return nil, false
	}
	return o.BackupPath, true
}

// HasBackupPath returns a boolean if a field has been set.
func (o *Summary) HasBackupPath() bool {
	return o != nil && o.BackupPath != nil
}

// SetBackupPath gets a reference to the given string and assigns it to the BackupPath field.
func (o *Summary) SetBackupPath(v string) {
	o.BackupPath = &v
}

// GetBackupSchedule returns the BackupSchedule field value if set, zero value otherwise.
func (o *Summary) GetBackupSchedule() string {
	if o == nil || o.BackupSchedule == nil {
		var ret string
		return ret
	}
	return *o.BackupSchedule
}

// GetBackupScheduleOk returns a tuple with the BackupSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetBackupScheduleOk() (*string, bool) {
	if o == nil || o.BackupSchedule == nil {
		return nil, false
	}
	return o.BackupSchedule, true
}

// HasBackupSchedule returns a boolean if a field has been set.
func (o *Summary) HasBackupSchedule() bool {
	return o != nil && o.BackupSchedule != nil
}

// SetBackupSchedule gets a reference to the given string and assigns it to the BackupSchedule field.
func (o *Summary) SetBackupSchedule(v string) {
	o.BackupSchedule = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Summary) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Summary) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Summary) SetStatus(v string) {
	o.Status = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Summary) GetCreationTime() string {
	if o == nil || o.CreationTime == nil {
		var ret string
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetCreationTimeOk() (*string, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Summary) HasCreationTime() bool {
	return o != nil && o.CreationTime != nil
}

// SetCreationTime gets a reference to the given string and assigns it to the CreationTime field.
func (o *Summary) SetCreationTime(v string) {
	o.CreationTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Summary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CpuRequests != nil {
		toSerialize["cpuRequests"] = o.CpuRequests
	}
	if o.CpuLimits != nil {
		toSerialize["cpuLimits"] = o.CpuLimits
	}
	if o.MemoryRequests != nil {
		toSerialize["memoryRequests"] = o.MemoryRequests
	}
	if o.MemoryLimits != nil {
		toSerialize["memoryLimits"] = o.MemoryLimits
	}
	if o.StorageSize != nil {
		toSerialize["storageSize"] = o.StorageSize
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.BackupEndpoint != nil {
		toSerialize["backupEndpoint"] = o.BackupEndpoint
	}
	if o.BackupPath != nil {
		toSerialize["backupPath"] = o.BackupPath
	}
	if o.BackupSchedule != nil {
		toSerialize["backupSchedule"] = o.BackupSchedule
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.CreationTime != nil {
		toSerialize["creationTime"] = o.CreationTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Summary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Namespace      *string `json:"namespace,omitempty"`
		Name           *string `json:"name,omitempty"`
		CpuRequests    *string `json:"cpuRequests,omitempty"`
		CpuLimits      *string `json:"cpuLimits,omitempty"`
		MemoryRequests *string `json:"memoryRequests,omitempty"`
		MemoryLimits   *string `json:"memoryLimits,omitempty"`
		StorageSize    *string `json:"storageSize,omitempty"`
		Replicas       *int32  `json:"replicas,omitempty"`
		BackupEndpoint *string `json:"backupEndpoint,omitempty"`
		BackupPath     *string `json:"backupPath,omitempty"`
		BackupSchedule *string `json:"backupSchedule,omitempty"`
		Status         *string `json:"status,omitempty"`
		CreationTime   *string `json:"creationTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"namespace", "name", "cpuRequests", "cpuLimits", "memoryRequests", "memoryLimits", "storageSize", "replicas", "backupEndpoint", "backupPath", "backupSchedule", "status", "creationTime"})
	} else {
		return err
	}
	o.Namespace = all.Namespace
	o.Name = all.Name
	o.CpuRequests = all.CpuRequests
	o.CpuLimits = all.CpuLimits
	o.MemoryRequests = all.MemoryRequests
	o.MemoryLimits = all.MemoryLimits
	o.StorageSize = all.StorageSize
	o.Replicas = all.Replicas
	o.BackupEndpoint = all.BackupEndpoint
	o.BackupPath = all.BackupPath
	o.BackupSchedule = all.BackupSchedule
	o.Status = all.Status
	o.CreationTime = all.CreationTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
