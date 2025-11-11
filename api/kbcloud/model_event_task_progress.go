// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventTaskProgress eventTaskProgress is the information of the event task progress
type EventTaskProgress struct {
	// conponent name
	Name *string `json:"name,omitempty"`
	// group name
	Group *string `json:"group,omitempty"`
	// the resource name
	ObjectKey common.NullableString `json:"objectKey,omitempty"`
	// message of the task progress
	Message *string `json:"message,omitempty"`
	// status of the task progress
	Status *string `json:"status,omitempty"`
	// start time of the event task progress
	StartTime common.NullableTime `json:"startTime,omitempty"`
	// end time of the event task progress
	EndTime common.NullableTime `json:"endTime,omitempty"`
	// Indicates the name of an OpsAction, Either `objectKey` or `customOpsName` must be provided. This field is provided when ops is `custom`.
	CustomOpsName common.NullableString `json:"customOpsName,omitempty"`
	// customOpsTasks is a list of custom ops task. This field is provided when ops is `custom`.
	CustomOpsTasks *CustomOpsTasks `json:"customOpsTasks,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventTaskProgress instantiates a new EventTaskProgress object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventTaskProgress() *EventTaskProgress {
	this := EventTaskProgress{}
	return &this
}

// NewEventTaskProgressWithDefaults instantiates a new EventTaskProgress object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventTaskProgressWithDefaults() *EventTaskProgress {
	this := EventTaskProgress{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventTaskProgress) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskProgress) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventTaskProgress) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventTaskProgress) SetName(v string) {
	o.Name = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *EventTaskProgress) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskProgress) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *EventTaskProgress) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *EventTaskProgress) SetGroup(v string) {
	o.Group = &v
}

// GetObjectKey returns the ObjectKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTaskProgress) GetObjectKey() string {
	if o == nil || o.ObjectKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectKey.Get()
}

// GetObjectKeyOk returns a tuple with the ObjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventTaskProgress) GetObjectKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectKey.Get(), o.ObjectKey.IsSet()
}

// HasObjectKey returns a boolean if a field has been set.
func (o *EventTaskProgress) HasObjectKey() bool {
	return o != nil && o.ObjectKey.IsSet()
}

// SetObjectKey gets a reference to the given common.NullableString and assigns it to the ObjectKey field.
func (o *EventTaskProgress) SetObjectKey(v string) {
	o.ObjectKey.Set(&v)
}

// SetObjectKeyNil sets the value for ObjectKey to be an explicit nil.
func (o *EventTaskProgress) SetObjectKeyNil() {
	o.ObjectKey.Set(nil)
}

// UnsetObjectKey ensures that no value is present for ObjectKey, not even an explicit nil.
func (o *EventTaskProgress) UnsetObjectKey() {
	o.ObjectKey.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EventTaskProgress) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskProgress) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EventTaskProgress) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EventTaskProgress) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventTaskProgress) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskProgress) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventTaskProgress) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EventTaskProgress) SetStatus(v string) {
	o.Status = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTaskProgress) GetStartTime() time.Time {
	if o == nil || o.StartTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime.Get()
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventTaskProgress) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTime.Get(), o.StartTime.IsSet()
}

// HasStartTime returns a boolean if a field has been set.
func (o *EventTaskProgress) HasStartTime() bool {
	return o != nil && o.StartTime.IsSet()
}

// SetStartTime gets a reference to the given common.NullableTime and assigns it to the StartTime field.
func (o *EventTaskProgress) SetStartTime(v time.Time) {
	o.StartTime.Set(&v)
}

// SetStartTimeNil sets the value for StartTime to be an explicit nil.
func (o *EventTaskProgress) SetStartTimeNil() {
	o.StartTime.Set(nil)
}

// UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil.
func (o *EventTaskProgress) UnsetStartTime() {
	o.StartTime.Unset()
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTaskProgress) GetEndTime() time.Time {
	if o == nil || o.EndTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventTaskProgress) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// HasEndTime returns a boolean if a field has been set.
func (o *EventTaskProgress) HasEndTime() bool {
	return o != nil && o.EndTime.IsSet()
}

// SetEndTime gets a reference to the given common.NullableTime and assigns it to the EndTime field.
func (o *EventTaskProgress) SetEndTime(v time.Time) {
	o.EndTime.Set(&v)
}

// SetEndTimeNil sets the value for EndTime to be an explicit nil.
func (o *EventTaskProgress) SetEndTimeNil() {
	o.EndTime.Set(nil)
}

// UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil.
func (o *EventTaskProgress) UnsetEndTime() {
	o.EndTime.Unset()
}

// GetCustomOpsName returns the CustomOpsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventTaskProgress) GetCustomOpsName() string {
	if o == nil || o.CustomOpsName.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomOpsName.Get()
}

// GetCustomOpsNameOk returns a tuple with the CustomOpsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventTaskProgress) GetCustomOpsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomOpsName.Get(), o.CustomOpsName.IsSet()
}

// HasCustomOpsName returns a boolean if a field has been set.
func (o *EventTaskProgress) HasCustomOpsName() bool {
	return o != nil && o.CustomOpsName.IsSet()
}

// SetCustomOpsName gets a reference to the given common.NullableString and assigns it to the CustomOpsName field.
func (o *EventTaskProgress) SetCustomOpsName(v string) {
	o.CustomOpsName.Set(&v)
}

// SetCustomOpsNameNil sets the value for CustomOpsName to be an explicit nil.
func (o *EventTaskProgress) SetCustomOpsNameNil() {
	o.CustomOpsName.Set(nil)
}

// UnsetCustomOpsName ensures that no value is present for CustomOpsName, not even an explicit nil.
func (o *EventTaskProgress) UnsetCustomOpsName() {
	o.CustomOpsName.Unset()
}

// GetCustomOpsTasks returns the CustomOpsTasks field value if set, zero value otherwise.
func (o *EventTaskProgress) GetCustomOpsTasks() CustomOpsTasks {
	if o == nil || o.CustomOpsTasks == nil {
		var ret CustomOpsTasks
		return ret
	}
	return *o.CustomOpsTasks
}

// GetCustomOpsTasksOk returns a tuple with the CustomOpsTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskProgress) GetCustomOpsTasksOk() (*CustomOpsTasks, bool) {
	if o == nil || o.CustomOpsTasks == nil {
		return nil, false
	}
	return o.CustomOpsTasks, true
}

// HasCustomOpsTasks returns a boolean if a field has been set.
func (o *EventTaskProgress) HasCustomOpsTasks() bool {
	return o != nil && o.CustomOpsTasks != nil
}

// SetCustomOpsTasks gets a reference to the given CustomOpsTasks and assigns it to the CustomOpsTasks field.
func (o *EventTaskProgress) SetCustomOpsTasks(v CustomOpsTasks) {
	o.CustomOpsTasks = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventTaskProgress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.ObjectKey.IsSet() {
		toSerialize["objectKey"] = o.ObjectKey.Get()
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StartTime.IsSet() {
		toSerialize["startTime"] = o.StartTime.Get()
	}
	if o.EndTime.IsSet() {
		toSerialize["endTime"] = o.EndTime.Get()
	}
	if o.CustomOpsName.IsSet() {
		toSerialize["customOpsName"] = o.CustomOpsName.Get()
	}
	if o.CustomOpsTasks != nil {
		toSerialize["customOpsTasks"] = o.CustomOpsTasks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventTaskProgress) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string               `json:"name,omitempty"`
		Group          *string               `json:"group,omitempty"`
		ObjectKey      common.NullableString `json:"objectKey,omitempty"`
		Message        *string               `json:"message,omitempty"`
		Status         *string               `json:"status,omitempty"`
		StartTime      common.NullableTime   `json:"startTime,omitempty"`
		EndTime        common.NullableTime   `json:"endTime,omitempty"`
		CustomOpsName  common.NullableString `json:"customOpsName,omitempty"`
		CustomOpsTasks *CustomOpsTasks       `json:"customOpsTasks,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "group", "objectKey", "message", "status", "startTime", "endTime", "customOpsName", "customOpsTasks"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Group = all.Group
	o.ObjectKey = all.ObjectKey
	o.Message = all.Message
	o.Status = all.Status
	o.StartTime = all.StartTime
	o.EndTime = all.EndTime
	o.CustomOpsName = all.CustomOpsName
	if all.CustomOpsTasks != nil && all.CustomOpsTasks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CustomOpsTasks = all.CustomOpsTasks

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
