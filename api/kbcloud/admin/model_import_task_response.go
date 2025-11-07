// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportTaskResponse Import task response
type ImportTaskResponse struct {
	// Import task ID
	Id string `json:"id"`
	// Import name
	Name string `json:"name"`
	// Import task status
	Status ImportTaskStatus `json:"status"`
	// Source endpoint address
	SourceEndpoint string `json:"sourceEndpoint"`
	// Target cluster name
	TargetClusterName string `json:"targetClusterName"`
	// Overall progress
	Progress int32 `json:"progress"`
	// Data source type
	SourceType string `json:"sourceType"`
	// Target engine type
	TargetEngine string `json:"targetEngine"`
	// Sync timestamp
	SyncTimestamp common.NullableTime `json:"syncTimestamp,omitempty"`
	// Delay seconds
	DelaySecond common.NullableInt32 `json:"delaySecond,omitempty"`
	// Creation time
	CreatedAt time.Time `json:"createdAt"`
	// Update time
	UpdatedAt time.Time `json:"updatedAt"`
	// Finish time
	FinishedAt common.NullableTime `json:"finishedAt,omitempty"`
	// Operation result message
	Message common.NullableString `json:"message"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportTaskResponse instantiates a new ImportTaskResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportTaskResponse(id string, name string, status ImportTaskStatus, sourceEndpoint string, targetClusterName string, progress int32, sourceType string, targetEngine string, createdAt time.Time, updatedAt time.Time, message common.NullableString) *ImportTaskResponse {
	this := ImportTaskResponse{}
	this.Id = id
	this.Name = name
	this.Status = status
	this.SourceEndpoint = sourceEndpoint
	this.TargetClusterName = targetClusterName
	this.Progress = progress
	this.SourceType = sourceType
	this.TargetEngine = targetEngine
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Message = message
	return &this
}

// NewImportTaskResponseWithDefaults instantiates a new ImportTaskResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportTaskResponseWithDefaults() *ImportTaskResponse {
	this := ImportTaskResponse{}
	var progress int32 = 0
	this.Progress = progress
	return &this
}

// GetId returns the Id field value.
func (o *ImportTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImportTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ImportTaskResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *ImportTaskResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImportTaskResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ImportTaskResponse) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value.
func (o *ImportTaskResponse) GetStatus() ImportTaskStatus {
	if o == nil {
		var ret ImportTaskStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ImportTaskResponse) GetStatusOk() (*ImportTaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ImportTaskResponse) SetStatus(v ImportTaskStatus) {
	o.Status = v
}

// GetSourceEndpoint returns the SourceEndpoint field value.
func (o *ImportTaskResponse) GetSourceEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceEndpoint
}

// GetSourceEndpointOk returns a tuple with the SourceEndpoint field value
// and a boolean to check if the value has been set.
func (o *ImportTaskResponse) GetSourceEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceEndpoint, true
}

// SetSourceEndpoint sets field value.
func (o *ImportTaskResponse) SetSourceEndpoint(v string) {
	o.SourceEndpoint = v
}

// GetTargetClusterName returns the TargetClusterName field value.
func (o *ImportTaskResponse) GetTargetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetClusterName
}

// GetTargetClusterNameOk returns a tuple with the TargetClusterName field value
// and a boolean to check if the value has been set.
func (o *ImportTaskResponse) GetTargetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetClusterName, true
}

// SetTargetClusterName sets field value.
func (o *ImportTaskResponse) SetTargetClusterName(v string) {
	o.TargetClusterName = v
}

// GetProgress returns the Progress field value.
func (o *ImportTaskResponse) GetProgress() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *ImportTaskResponse) GetProgressOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value.
func (o *ImportTaskResponse) SetProgress(v int32) {
	o.Progress = v
}

// GetSourceType returns the SourceType field value.
func (o *ImportTaskResponse) GetSourceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *ImportTaskResponse) GetSourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value.
func (o *ImportTaskResponse) SetSourceType(v string) {
	o.SourceType = v
}

// GetTargetEngine returns the TargetEngine field value.
func (o *ImportTaskResponse) GetTargetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetEngine
}

// GetTargetEngineOk returns a tuple with the TargetEngine field value
// and a boolean to check if the value has been set.
func (o *ImportTaskResponse) GetTargetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetEngine, true
}

// SetTargetEngine sets field value.
func (o *ImportTaskResponse) SetTargetEngine(v string) {
	o.TargetEngine = v
}

// GetSyncTimestamp returns the SyncTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportTaskResponse) GetSyncTimestamp() time.Time {
	if o == nil || o.SyncTimestamp.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SyncTimestamp.Get()
}

// GetSyncTimestampOk returns a tuple with the SyncTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ImportTaskResponse) GetSyncTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyncTimestamp.Get(), o.SyncTimestamp.IsSet()
}

// HasSyncTimestamp returns a boolean if a field has been set.
func (o *ImportTaskResponse) HasSyncTimestamp() bool {
	return o != nil && o.SyncTimestamp.IsSet()
}

// SetSyncTimestamp gets a reference to the given common.NullableTime and assigns it to the SyncTimestamp field.
func (o *ImportTaskResponse) SetSyncTimestamp(v time.Time) {
	o.SyncTimestamp.Set(&v)
}

// SetSyncTimestampNil sets the value for SyncTimestamp to be an explicit nil.
func (o *ImportTaskResponse) SetSyncTimestampNil() {
	o.SyncTimestamp.Set(nil)
}

// UnsetSyncTimestamp ensures that no value is present for SyncTimestamp, not even an explicit nil.
func (o *ImportTaskResponse) UnsetSyncTimestamp() {
	o.SyncTimestamp.Unset()
}

// GetDelaySecond returns the DelaySecond field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportTaskResponse) GetDelaySecond() int32 {
	if o == nil || o.DelaySecond.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DelaySecond.Get()
}

// GetDelaySecondOk returns a tuple with the DelaySecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ImportTaskResponse) GetDelaySecondOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DelaySecond.Get(), o.DelaySecond.IsSet()
}

// HasDelaySecond returns a boolean if a field has been set.
func (o *ImportTaskResponse) HasDelaySecond() bool {
	return o != nil && o.DelaySecond.IsSet()
}

// SetDelaySecond gets a reference to the given common.NullableInt32 and assigns it to the DelaySecond field.
func (o *ImportTaskResponse) SetDelaySecond(v int32) {
	o.DelaySecond.Set(&v)
}

// SetDelaySecondNil sets the value for DelaySecond to be an explicit nil.
func (o *ImportTaskResponse) SetDelaySecondNil() {
	o.DelaySecond.Set(nil)
}

// UnsetDelaySecond ensures that no value is present for DelaySecond, not even an explicit nil.
func (o *ImportTaskResponse) UnsetDelaySecond() {
	o.DelaySecond.Unset()
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ImportTaskResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ImportTaskResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ImportTaskResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *ImportTaskResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ImportTaskResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *ImportTaskResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportTaskResponse) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt.Get()
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ImportTaskResponse) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FinishedAt.Get(), o.FinishedAt.IsSet()
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *ImportTaskResponse) HasFinishedAt() bool {
	return o != nil && o.FinishedAt.IsSet()
}

// SetFinishedAt gets a reference to the given common.NullableTime and assigns it to the FinishedAt field.
func (o *ImportTaskResponse) SetFinishedAt(v time.Time) {
	o.FinishedAt.Set(&v)
}

// SetFinishedAtNil sets the value for FinishedAt to be an explicit nil.
func (o *ImportTaskResponse) SetFinishedAtNil() {
	o.FinishedAt.Set(nil)
}

// UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil.
func (o *ImportTaskResponse) UnsetFinishedAt() {
	o.FinishedAt.Unset()
}

// GetMessage returns the Message field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ImportTaskResponse) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ImportTaskResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value.
func (o *ImportTaskResponse) SetMessage(v string) {
	o.Message.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["sourceEndpoint"] = o.SourceEndpoint
	toSerialize["targetClusterName"] = o.TargetClusterName
	toSerialize["progress"] = o.Progress
	toSerialize["sourceType"] = o.SourceType
	toSerialize["targetEngine"] = o.TargetEngine
	if o.SyncTimestamp.IsSet() {
		toSerialize["syncTimestamp"] = o.SyncTimestamp.Get()
	}
	if o.DelaySecond.IsSet() {
		toSerialize["delaySecond"] = o.DelaySecond.Get()
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.FinishedAt.IsSet() {
		toSerialize["finishedAt"] = o.FinishedAt.Get()
	}
	toSerialize["message"] = o.Message.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportTaskResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                *string               `json:"id"`
		Name              *string               `json:"name"`
		Status            *ImportTaskStatus     `json:"status"`
		SourceEndpoint    *string               `json:"sourceEndpoint"`
		TargetClusterName *string               `json:"targetClusterName"`
		Progress          *int32                `json:"progress"`
		SourceType        *string               `json:"sourceType"`
		TargetEngine      *string               `json:"targetEngine"`
		SyncTimestamp     common.NullableTime   `json:"syncTimestamp,omitempty"`
		DelaySecond       common.NullableInt32  `json:"delaySecond,omitempty"`
		CreatedAt         *time.Time            `json:"createdAt"`
		UpdatedAt         *time.Time            `json:"updatedAt"`
		FinishedAt        common.NullableTime   `json:"finishedAt,omitempty"`
		Message           common.NullableString `json:"message"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.SourceEndpoint == nil {
		return fmt.Errorf("required field sourceEndpoint missing")
	}
	if all.TargetClusterName == nil {
		return fmt.Errorf("required field targetClusterName missing")
	}
	if all.Progress == nil {
		return fmt.Errorf("required field progress missing")
	}
	if all.SourceType == nil {
		return fmt.Errorf("required field sourceType missing")
	}
	if all.TargetEngine == nil {
		return fmt.Errorf("required field targetEngine missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	if !all.Message.IsSet() {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "status", "sourceEndpoint", "targetClusterName", "progress", "sourceType", "targetEngine", "syncTimestamp", "delaySecond", "createdAt", "updatedAt", "finishedAt", "message"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Name = *all.Name
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.SourceEndpoint = *all.SourceEndpoint
	o.TargetClusterName = *all.TargetClusterName
	o.Progress = *all.Progress
	o.SourceType = *all.SourceType
	o.TargetEngine = *all.TargetEngine
	o.SyncTimestamp = all.SyncTimestamp
	o.DelaySecond = all.DelaySecond
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt
	o.FinishedAt = all.FinishedAt
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
