// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DisasterRecoveryHistoryItem DisasterRecovery history detail for Cluster
type DisasterRecoveryHistoryItem struct {
	// the ID of disaster recovery task in TaskFlow
	TaskId common.NullableInt32 `json:"taskID,omitempty"`
	// the ID of parent cluster
	ParentClusterId common.NullableInt32 `json:"parentClusterID,omitempty"`
	// the Name of parent cluster
	ParentClusterName common.NullableString `json:"parentClusterName,omitempty"`
	// parent env name
	ParentEnvName common.NullableString `json:"parentEnvName,omitempty"`
	// the ID of promote cluster
	ClusterId common.NullableInt32 `json:"clusterID,omitempty"`
	// the Name of promote cluster
	ClusterName common.NullableString `json:"clusterName,omitempty"`
	// env name
	EnvName common.NullableString `json:"envName,omitempty"`
	// the event type of disasterRecovery history, support values: [CreateInstance, DeleteInstance, Promote]
	EventType NullableDisasterRecoveryEventType `json:"eventType,omitempty"`
	// the reason of promote
	Reason common.NullableString `json:"reason,omitempty"`
	// the operator name
	Operator *string `json:"operator,omitempty"`
	// the user ID of the operator
	OperatorId *string `json:"operatorId,omitempty"`
	// the create time of promote event
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// the update time of promote event
	UpdateAt *time.Time `json:"updateAt,omitempty"`
	// the status of promote event, support values: [Succeed, Failed, Running, Unknown]
	Status NullableDisasterRecoveryStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDisasterRecoveryHistoryItem instantiates a new DisasterRecoveryHistoryItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDisasterRecoveryHistoryItem() *DisasterRecoveryHistoryItem {
	this := DisasterRecoveryHistoryItem{}
	var status DisasterRecoveryStatus = DisasterRecoveryStatusUnknown
	this.Status = *NewNullableDisasterRecoveryStatus(&status)
	return &this
}

// NewDisasterRecoveryHistoryItemWithDefaults instantiates a new DisasterRecoveryHistoryItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDisasterRecoveryHistoryItemWithDefaults() *DisasterRecoveryHistoryItem {
	this := DisasterRecoveryHistoryItem{}
	var status DisasterRecoveryStatus = DisasterRecoveryStatusUnknown
	this.Status = *NewNullableDisasterRecoveryStatus(&status)
	return &this
}

// GetTaskId returns the TaskId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryHistoryItem) GetTaskId() int32 {
	if o == nil || o.TaskId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TaskId.Get()
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryHistoryItem) GetTaskIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskId.Get(), o.TaskId.IsSet()
}

// HasTaskId returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasTaskId() bool {
	return o != nil && o.TaskId.IsSet()
}

// SetTaskId gets a reference to the given common.NullableInt32 and assigns it to the TaskId field.
func (o *DisasterRecoveryHistoryItem) SetTaskId(v int32) {
	o.TaskId.Set(&v)
}

// SetTaskIdNil sets the value for TaskId to be an explicit nil.
func (o *DisasterRecoveryHistoryItem) SetTaskIdNil() {
	o.TaskId.Set(nil)
}

// UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil.
func (o *DisasterRecoveryHistoryItem) UnsetTaskId() {
	o.TaskId.Unset()
}

// GetParentClusterId returns the ParentClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryHistoryItem) GetParentClusterId() int32 {
	if o == nil || o.ParentClusterId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ParentClusterId.Get()
}

// GetParentClusterIdOk returns a tuple with the ParentClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryHistoryItem) GetParentClusterIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentClusterId.Get(), o.ParentClusterId.IsSet()
}

// HasParentClusterId returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasParentClusterId() bool {
	return o != nil && o.ParentClusterId.IsSet()
}

// SetParentClusterId gets a reference to the given common.NullableInt32 and assigns it to the ParentClusterId field.
func (o *DisasterRecoveryHistoryItem) SetParentClusterId(v int32) {
	o.ParentClusterId.Set(&v)
}

// SetParentClusterIdNil sets the value for ParentClusterId to be an explicit nil.
func (o *DisasterRecoveryHistoryItem) SetParentClusterIdNil() {
	o.ParentClusterId.Set(nil)
}

// UnsetParentClusterId ensures that no value is present for ParentClusterId, not even an explicit nil.
func (o *DisasterRecoveryHistoryItem) UnsetParentClusterId() {
	o.ParentClusterId.Unset()
}

// GetParentClusterName returns the ParentClusterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryHistoryItem) GetParentClusterName() string {
	if o == nil || o.ParentClusterName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentClusterName.Get()
}

// GetParentClusterNameOk returns a tuple with the ParentClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryHistoryItem) GetParentClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentClusterName.Get(), o.ParentClusterName.IsSet()
}

// HasParentClusterName returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasParentClusterName() bool {
	return o != nil && o.ParentClusterName.IsSet()
}

// SetParentClusterName gets a reference to the given common.NullableString and assigns it to the ParentClusterName field.
func (o *DisasterRecoveryHistoryItem) SetParentClusterName(v string) {
	o.ParentClusterName.Set(&v)
}

// SetParentClusterNameNil sets the value for ParentClusterName to be an explicit nil.
func (o *DisasterRecoveryHistoryItem) SetParentClusterNameNil() {
	o.ParentClusterName.Set(nil)
}

// UnsetParentClusterName ensures that no value is present for ParentClusterName, not even an explicit nil.
func (o *DisasterRecoveryHistoryItem) UnsetParentClusterName() {
	o.ParentClusterName.Unset()
}

// GetParentEnvName returns the ParentEnvName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryHistoryItem) GetParentEnvName() string {
	if o == nil || o.ParentEnvName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentEnvName.Get()
}

// GetParentEnvNameOk returns a tuple with the ParentEnvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryHistoryItem) GetParentEnvNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentEnvName.Get(), o.ParentEnvName.IsSet()
}

// HasParentEnvName returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasParentEnvName() bool {
	return o != nil && o.ParentEnvName.IsSet()
}

// SetParentEnvName gets a reference to the given common.NullableString and assigns it to the ParentEnvName field.
func (o *DisasterRecoveryHistoryItem) SetParentEnvName(v string) {
	o.ParentEnvName.Set(&v)
}

// SetParentEnvNameNil sets the value for ParentEnvName to be an explicit nil.
func (o *DisasterRecoveryHistoryItem) SetParentEnvNameNil() {
	o.ParentEnvName.Set(nil)
}

// UnsetParentEnvName ensures that no value is present for ParentEnvName, not even an explicit nil.
func (o *DisasterRecoveryHistoryItem) UnsetParentEnvName() {
	o.ParentEnvName.Unset()
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryHistoryItem) GetClusterId() int32 {
	if o == nil || o.ClusterId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ClusterId.Get()
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryHistoryItem) GetClusterIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterId.Get(), o.ClusterId.IsSet()
}

// HasClusterId returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasClusterId() bool {
	return o != nil && o.ClusterId.IsSet()
}

// SetClusterId gets a reference to the given common.NullableInt32 and assigns it to the ClusterId field.
func (o *DisasterRecoveryHistoryItem) SetClusterId(v int32) {
	o.ClusterId.Set(&v)
}

// SetClusterIdNil sets the value for ClusterId to be an explicit nil.
func (o *DisasterRecoveryHistoryItem) SetClusterIdNil() {
	o.ClusterId.Set(nil)
}

// UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil.
func (o *DisasterRecoveryHistoryItem) UnsetClusterId() {
	o.ClusterId.Unset()
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryHistoryItem) GetClusterName() string {
	if o == nil || o.ClusterName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClusterName.Get()
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryHistoryItem) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClusterName.Get(), o.ClusterName.IsSet()
}

// HasClusterName returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasClusterName() bool {
	return o != nil && o.ClusterName.IsSet()
}

// SetClusterName gets a reference to the given common.NullableString and assigns it to the ClusterName field.
func (o *DisasterRecoveryHistoryItem) SetClusterName(v string) {
	o.ClusterName.Set(&v)
}

// SetClusterNameNil sets the value for ClusterName to be an explicit nil.
func (o *DisasterRecoveryHistoryItem) SetClusterNameNil() {
	o.ClusterName.Set(nil)
}

// UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil.
func (o *DisasterRecoveryHistoryItem) UnsetClusterName() {
	o.ClusterName.Unset()
}

// GetEnvName returns the EnvName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryHistoryItem) GetEnvName() string {
	if o == nil || o.EnvName.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnvName.Get()
}

// GetEnvNameOk returns a tuple with the EnvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryHistoryItem) GetEnvNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvName.Get(), o.EnvName.IsSet()
}

// HasEnvName returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasEnvName() bool {
	return o != nil && o.EnvName.IsSet()
}

// SetEnvName gets a reference to the given common.NullableString and assigns it to the EnvName field.
func (o *DisasterRecoveryHistoryItem) SetEnvName(v string) {
	o.EnvName.Set(&v)
}

// SetEnvNameNil sets the value for EnvName to be an explicit nil.
func (o *DisasterRecoveryHistoryItem) SetEnvNameNil() {
	o.EnvName.Set(nil)
}

// UnsetEnvName ensures that no value is present for EnvName, not even an explicit nil.
func (o *DisasterRecoveryHistoryItem) UnsetEnvName() {
	o.EnvName.Unset()
}

// GetEventType returns the EventType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryHistoryItem) GetEventType() DisasterRecoveryEventType {
	if o == nil || o.EventType.Get() == nil {
		var ret DisasterRecoveryEventType
		return ret
	}
	return *o.EventType.Get()
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryHistoryItem) GetEventTypeOk() (*DisasterRecoveryEventType, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventType.Get(), o.EventType.IsSet()
}

// HasEventType returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasEventType() bool {
	return o != nil && o.EventType.IsSet()
}

// SetEventType gets a reference to the given NullableDisasterRecoveryEventType and assigns it to the EventType field.
func (o *DisasterRecoveryHistoryItem) SetEventType(v DisasterRecoveryEventType) {
	o.EventType.Set(&v)
}

// SetEventTypeNil sets the value for EventType to be an explicit nil.
func (o *DisasterRecoveryHistoryItem) SetEventTypeNil() {
	o.EventType.Set(nil)
}

// UnsetEventType ensures that no value is present for EventType, not even an explicit nil.
func (o *DisasterRecoveryHistoryItem) UnsetEventType() {
	o.EventType.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryHistoryItem) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryHistoryItem) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasReason() bool {
	return o != nil && o.Reason.IsSet()
}

// SetReason gets a reference to the given common.NullableString and assigns it to the Reason field.
func (o *DisasterRecoveryHistoryItem) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil.
func (o *DisasterRecoveryHistoryItem) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil.
func (o *DisasterRecoveryHistoryItem) UnsetReason() {
	o.Reason.Unset()
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *DisasterRecoveryHistoryItem) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryHistoryItem) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *DisasterRecoveryHistoryItem) SetOperator(v string) {
	o.Operator = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise.
func (o *DisasterRecoveryHistoryItem) GetOperatorId() string {
	if o == nil || o.OperatorId == nil {
		var ret string
		return ret
	}
	return *o.OperatorId
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryHistoryItem) GetOperatorIdOk() (*string, bool) {
	if o == nil || o.OperatorId == nil {
		return nil, false
	}
	return o.OperatorId, true
}

// HasOperatorId returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasOperatorId() bool {
	return o != nil && o.OperatorId != nil
}

// SetOperatorId gets a reference to the given string and assigns it to the OperatorId field.
func (o *DisasterRecoveryHistoryItem) SetOperatorId(v string) {
	o.OperatorId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DisasterRecoveryHistoryItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryHistoryItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DisasterRecoveryHistoryItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdateAt returns the UpdateAt field value if set, zero value otherwise.
func (o *DisasterRecoveryHistoryItem) GetUpdateAt() time.Time {
	if o == nil || o.UpdateAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateAt
}

// GetUpdateAtOk returns a tuple with the UpdateAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryHistoryItem) GetUpdateAtOk() (*time.Time, bool) {
	if o == nil || o.UpdateAt == nil {
		return nil, false
	}
	return o.UpdateAt, true
}

// HasUpdateAt returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasUpdateAt() bool {
	return o != nil && o.UpdateAt != nil
}

// SetUpdateAt gets a reference to the given time.Time and assigns it to the UpdateAt field.
func (o *DisasterRecoveryHistoryItem) SetUpdateAt(v time.Time) {
	o.UpdateAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisasterRecoveryHistoryItem) GetStatus() DisasterRecoveryStatus {
	if o == nil || o.Status.Get() == nil {
		var ret DisasterRecoveryStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DisasterRecoveryHistoryItem) GetStatusOk() (*DisasterRecoveryStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *DisasterRecoveryHistoryItem) HasStatus() bool {
	return o != nil && o.Status.IsSet()
}

// SetStatus gets a reference to the given NullableDisasterRecoveryStatus and assigns it to the Status field.
func (o *DisasterRecoveryHistoryItem) SetStatus(v DisasterRecoveryStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil.
func (o *DisasterRecoveryHistoryItem) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil.
func (o *DisasterRecoveryHistoryItem) UnsetStatus() {
	o.Status.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DisasterRecoveryHistoryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TaskId.IsSet() {
		toSerialize["taskID"] = o.TaskId.Get()
	}
	if o.ParentClusterId.IsSet() {
		toSerialize["parentClusterID"] = o.ParentClusterId.Get()
	}
	if o.ParentClusterName.IsSet() {
		toSerialize["parentClusterName"] = o.ParentClusterName.Get()
	}
	if o.ParentEnvName.IsSet() {
		toSerialize["parentEnvName"] = o.ParentEnvName.Get()
	}
	if o.ClusterId.IsSet() {
		toSerialize["clusterID"] = o.ClusterId.Get()
	}
	if o.ClusterName.IsSet() {
		toSerialize["clusterName"] = o.ClusterName.Get()
	}
	if o.EnvName.IsSet() {
		toSerialize["envName"] = o.EnvName.Get()
	}
	if o.EventType.IsSet() {
		toSerialize["eventType"] = o.EventType.Get()
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.OperatorId != nil {
		toSerialize["operatorId"] = o.OperatorId
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdateAt != nil {
		if o.UpdateAt.Nanosecond() == 0 {
			toSerialize["updateAt"] = o.UpdateAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updateAt"] = o.UpdateAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DisasterRecoveryHistoryItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId            common.NullableInt32              `json:"taskID,omitempty"`
		ParentClusterId   common.NullableInt32              `json:"parentClusterID,omitempty"`
		ParentClusterName common.NullableString             `json:"parentClusterName,omitempty"`
		ParentEnvName     common.NullableString             `json:"parentEnvName,omitempty"`
		ClusterId         common.NullableInt32              `json:"clusterID,omitempty"`
		ClusterName       common.NullableString             `json:"clusterName,omitempty"`
		EnvName           common.NullableString             `json:"envName,omitempty"`
		EventType         NullableDisasterRecoveryEventType `json:"eventType,omitempty"`
		Reason            common.NullableString             `json:"reason,omitempty"`
		Operator          *string                           `json:"operator,omitempty"`
		OperatorId        *string                           `json:"operatorId,omitempty"`
		CreatedAt         *time.Time                        `json:"createdAt,omitempty"`
		UpdateAt          *time.Time                        `json:"updateAt,omitempty"`
		Status            NullableDisasterRecoveryStatus    `json:"status,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskID", "parentClusterID", "parentClusterName", "parentEnvName", "clusterID", "clusterName", "envName", "eventType", "reason", "operator", "operatorId", "createdAt", "updateAt", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TaskId = all.TaskId
	o.ParentClusterId = all.ParentClusterId
	o.ParentClusterName = all.ParentClusterName
	o.ParentEnvName = all.ParentEnvName
	o.ClusterId = all.ClusterId
	o.ClusterName = all.ClusterName
	o.EnvName = all.EnvName
	if all.EventType.Get() != nil && !all.EventType.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = all.EventType
	}
	o.Reason = all.Reason
	o.Operator = all.Operator
	o.OperatorId = all.OperatorId
	o.CreatedAt = all.CreatedAt
	o.UpdateAt = all.UpdateAt
	if all.Status.Get() != nil && !all.Status.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
