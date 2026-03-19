// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/google/uuid"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MeteringTrack RTS metering track information
type MeteringTrack struct {
	// The ID of cluster
	ClusterId common.NullableString `json:"clusterID,omitempty"`
	// The Name of organization
	OrgName common.NullableString `json:"orgName,omitempty"`
	// The Name of project
	ProjectName common.NullableString `json:"projectName,omitempty"`
	// The ID of environment
	EnvironmentId common.NullableString `json:"environmentID,omitempty"`
	// The clipped start time of the metering track in Unix timestamp (seconds since epoch)
	StartTime *int64 `json:"startTime,omitempty"`
	// The clipped end time of the metering track in Unix timestamp (seconds since epoch)
	EndTime *int64 `json:"endTime,omitempty"`
	// The CPU usage tracked for the interval
	Cpu *string `json:"cpu,omitempty"`
	// The memory usage tracked for the interval
	Memory *string `json:"memory,omitempty"`
	// The storage usage tracked for the interval
	Storage *string `json:"storage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMeteringTrack instantiates a new MeteringTrack object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMeteringTrack() *MeteringTrack {
	this := MeteringTrack{}
	return &this
}

// NewMeteringTrackWithDefaults instantiates a new MeteringTrack object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMeteringTrackWithDefaults() *MeteringTrack {
	this := MeteringTrack{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MeteringTrack) GetClusterId() string {
	if o == nil || o.ClusterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterId.Get()
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MeteringTrack) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterId.Get(), o.ClusterId.IsSet()
}

// HasClusterId returns a boolean if a field has been set.
func (o *MeteringTrack) HasClusterId() bool {
	return o != nil && o.ClusterId.IsSet()
}

// SetClusterId gets a reference to the given common.NullableString and assigns it to the ClusterId field.
func (o *MeteringTrack) SetClusterId(v string) {
	o.ClusterId.Set(&v)
}

// SetClusterIdNil sets the value for ClusterId to be an explicit nil.
func (o *MeteringTrack) SetClusterIdNil() {
	o.ClusterId.Set(nil)
}

// UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil.
func (o *MeteringTrack) UnsetClusterId() {
	o.ClusterId.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MeteringTrack) GetOrgName() string {
	if o == nil || o.OrgName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrgName.Get()
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MeteringTrack) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrgName.Get(), o.OrgName.IsSet()
}

// HasOrgName returns a boolean if a field has been set.
func (o *MeteringTrack) HasOrgName() bool {
	return o != nil && o.OrgName.IsSet()
}

// SetOrgName gets a reference to the given common.NullableString and assigns it to the OrgName field.
func (o *MeteringTrack) SetOrgName(v string) {
	o.OrgName.Set(&v)
}

// SetOrgNameNil sets the value for OrgName to be an explicit nil.
func (o *MeteringTrack) SetOrgNameNil() {
	o.OrgName.Set(nil)
}

// UnsetOrgName ensures that no value is present for OrgName, not even an explicit nil.
func (o *MeteringTrack) UnsetOrgName() {
	o.OrgName.Unset()
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MeteringTrack) GetProjectName() string {
	if o == nil || o.ProjectName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MeteringTrack) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *MeteringTrack) HasProjectName() bool {
	return o != nil && o.ProjectName.IsSet()
}

// SetProjectName gets a reference to the given common.NullableString and assigns it to the ProjectName field.
func (o *MeteringTrack) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}

// SetProjectNameNil sets the value for ProjectName to be an explicit nil.
func (o *MeteringTrack) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil.
func (o *MeteringTrack) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MeteringTrack) GetEnvironmentId() uuid.UUID {
	if o == nil || o.EnvironmentId.Get() == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MeteringTrack) GetEnvironmentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *MeteringTrack) HasEnvironmentId() bool {
	return o != nil && o.EnvironmentId.IsSet()
}

// SetEnvironmentId gets a reference to the given common.NullableString and assigns it to the EnvironmentId field.
func (o *MeteringTrack) SetEnvironmentId(v uuid.UUID) {
	o.EnvironmentId.Set(&v)
}

// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil.
func (o *MeteringTrack) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil.
func (o *MeteringTrack) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *MeteringTrack) GetStartTime() int64 {
	if o == nil || o.StartTime == nil {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringTrack) GetStartTimeOk() (*int64, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *MeteringTrack) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *MeteringTrack) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *MeteringTrack) GetEndTime() int64 {
	if o == nil || o.EndTime == nil {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringTrack) GetEndTimeOk() (*int64, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *MeteringTrack) HasEndTime() bool {
	return o != nil && o.EndTime != nil
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *MeteringTrack) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *MeteringTrack) GetCpu() string {
	if o == nil || o.Cpu == nil {
		var ret string
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringTrack) GetCpuOk() (*string, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *MeteringTrack) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given string and assigns it to the Cpu field.
func (o *MeteringTrack) SetCpu(v string) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *MeteringTrack) GetMemory() string {
	if o == nil || o.Memory == nil {
		var ret string
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringTrack) GetMemoryOk() (*string, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *MeteringTrack) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given string and assigns it to the Memory field.
func (o *MeteringTrack) SetMemory(v string) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *MeteringTrack) GetStorage() string {
	if o == nil || o.Storage == nil {
		var ret string
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringTrack) GetStorageOk() (*string, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *MeteringTrack) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given string and assigns it to the Storage field.
func (o *MeteringTrack) SetStorage(v string) {
	o.Storage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MeteringTrack) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ClusterId.IsSet() {
		toSerialize["clusterID"] = o.ClusterId.Get()
	}
	if o.OrgName.IsSet() {
		toSerialize["orgName"] = o.OrgName.Get()
	}
	if o.ProjectName.IsSet() {
		toSerialize["projectName"] = o.ProjectName.Get()
	}
	if o.EnvironmentId.IsSet() {
		toSerialize["environmentID"] = o.EnvironmentId.Get()
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MeteringTrack) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterId     common.NullableString `json:"clusterID,omitempty"`
		OrgName       common.NullableString `json:"orgName,omitempty"`
		ProjectName   common.NullableString `json:"projectName,omitempty"`
		EnvironmentId common.NullableString `json:"environmentID,omitempty"`
		StartTime     *int64                `json:"startTime,omitempty"`
		EndTime       *int64                `json:"endTime,omitempty"`
		Cpu           *string               `json:"cpu,omitempty"`
		Memory        *string               `json:"memory,omitempty"`
		Storage       *string               `json:"storage,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterID", "orgName", "projectName", "environmentID", "startTime", "endTime", "cpu", "memory", "storage"})
	} else {
		return err
	}
	o.ClusterId = all.ClusterId
	o.OrgName = all.OrgName
	o.ProjectName = all.ProjectName
	o.EnvironmentId = all.EnvironmentId
	o.StartTime = all.StartTime
	o.EndTime = all.EndTime
	o.Cpu = all.Cpu
	o.Memory = all.Memory
	o.Storage = all.Storage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
