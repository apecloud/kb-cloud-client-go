// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package data

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// ClusterTaskProgress clusterTaskProgress is the information of the task progress 
type ClusterTaskProgress struct {
	// conponent name
	Name *string `json:"name,omitempty"`
	// group name
	Group *string `json:"group,omitempty"`
	// the instance name
	ObjectKey common.NullableString `json:"objectKey,omitempty"`
	// message of the task progress
	Message *string `json:"message,omitempty"`
	// status of the task progress
	Status *string `json:"status,omitempty"`
	// start time of the task progress
	StartTime *time.Time `json:"startTime,omitempty"`
	// end time of the task progress
	EndTime common.NullableTime `json:"endTime,omitempty"`
	// Indicates the name of an OpsAction, Either `objectKey` or `customOpsName` must be provided. This field is provided when ops is `custom`.
	CustomOpsName common.NullableString `json:"customOpsName,omitempty"`
	// customOpsTasks is a list of custom ops task. This field is provided when ops is `custom`.
	CustomOpsTasks *CustomOpsTasks `json:"customOpsTasks,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewClusterTaskProgress instantiates a new ClusterTaskProgress object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterTaskProgress() *ClusterTaskProgress {
	this := ClusterTaskProgress{}
	return &this
}

// NewClusterTaskProgressWithDefaults instantiates a new ClusterTaskProgress object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterTaskProgressWithDefaults() *ClusterTaskProgress {
	this := ClusterTaskProgress{}
	return &this
}
// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterTaskProgress) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskProgress) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterTaskProgress) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterTaskProgress) SetName(v string) {
	o.Name = &v
}


// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ClusterTaskProgress) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskProgress) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ClusterTaskProgress) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ClusterTaskProgress) SetGroup(v string) {
	o.Group = &v
}


// GetObjectKey returns the ObjectKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterTaskProgress) GetObjectKey() string {
	if o == nil || o.ObjectKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectKey.Get()
}

// GetObjectKeyOk returns a tuple with the ObjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterTaskProgress) GetObjectKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectKey.Get(), o.ObjectKey.IsSet()
}

// HasObjectKey returns a boolean if a field has been set.
func (o *ClusterTaskProgress) HasObjectKey() bool {
	return o != nil && o.ObjectKey.IsSet()
}

// SetObjectKey gets a reference to the given common.NullableString and assigns it to the ObjectKey field.
func (o *ClusterTaskProgress) SetObjectKey(v string) {
	o.ObjectKey.Set(&v)
}
// SetObjectKeyNil sets the value for ObjectKey to be an explicit nil.
func (o *ClusterTaskProgress) SetObjectKeyNil() {
	o.ObjectKey.Set(nil)
}

// UnsetObjectKey ensures that no value is present for ObjectKey, not even an explicit nil.
func (o *ClusterTaskProgress) UnsetObjectKey() {
	o.ObjectKey.Unset()
}


// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClusterTaskProgress) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskProgress) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClusterTaskProgress) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ClusterTaskProgress) SetMessage(v string) {
	o.Message = &v
}


// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterTaskProgress) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskProgress) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterTaskProgress) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ClusterTaskProgress) SetStatus(v string) {
	o.Status = &v
}


// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ClusterTaskProgress) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskProgress) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ClusterTaskProgress) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ClusterTaskProgress) SetStartTime(v time.Time) {
	o.StartTime = &v
}


// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterTaskProgress) GetEndTime() time.Time {
	if o == nil || o.EndTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterTaskProgress) GetEndTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// HasEndTime returns a boolean if a field has been set.
func (o *ClusterTaskProgress) HasEndTime() bool {
	return o != nil && o.EndTime.IsSet()
}

// SetEndTime gets a reference to the given common.NullableTime and assigns it to the EndTime field.
func (o *ClusterTaskProgress) SetEndTime(v time.Time) {
	o.EndTime.Set(&v)
}
// SetEndTimeNil sets the value for EndTime to be an explicit nil.
func (o *ClusterTaskProgress) SetEndTimeNil() {
	o.EndTime.Set(nil)
}

// UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil.
func (o *ClusterTaskProgress) UnsetEndTime() {
	o.EndTime.Unset()
}


// GetCustomOpsName returns the CustomOpsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterTaskProgress) GetCustomOpsName() string {
	if o == nil || o.CustomOpsName.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomOpsName.Get()
}

// GetCustomOpsNameOk returns a tuple with the CustomOpsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ClusterTaskProgress) GetCustomOpsNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomOpsName.Get(), o.CustomOpsName.IsSet()
}

// HasCustomOpsName returns a boolean if a field has been set.
func (o *ClusterTaskProgress) HasCustomOpsName() bool {
	return o != nil && o.CustomOpsName.IsSet()
}

// SetCustomOpsName gets a reference to the given common.NullableString and assigns it to the CustomOpsName field.
func (o *ClusterTaskProgress) SetCustomOpsName(v string) {
	o.CustomOpsName.Set(&v)
}
// SetCustomOpsNameNil sets the value for CustomOpsName to be an explicit nil.
func (o *ClusterTaskProgress) SetCustomOpsNameNil() {
	o.CustomOpsName.Set(nil)
}

// UnsetCustomOpsName ensures that no value is present for CustomOpsName, not even an explicit nil.
func (o *ClusterTaskProgress) UnsetCustomOpsName() {
	o.CustomOpsName.Unset()
}


// GetCustomOpsTasks returns the CustomOpsTasks field value if set, zero value otherwise.
func (o *ClusterTaskProgress) GetCustomOpsTasks() CustomOpsTasks {
	if o == nil || o.CustomOpsTasks == nil {
		var ret CustomOpsTasks
		return ret
	}
	return *o.CustomOpsTasks
}

// GetCustomOpsTasksOk returns a tuple with the CustomOpsTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTaskProgress) GetCustomOpsTasksOk() (*CustomOpsTasks, bool) {
	if o == nil || o.CustomOpsTasks == nil {
		return nil, false
	}
	return o.CustomOpsTasks, true
}

// HasCustomOpsTasks returns a boolean if a field has been set.
func (o *ClusterTaskProgress) HasCustomOpsTasks() bool {
	return o != nil && o.CustomOpsTasks != nil
}

// SetCustomOpsTasks gets a reference to the given CustomOpsTasks and assigns it to the CustomOpsTasks field.
func (o *ClusterTaskProgress) SetCustomOpsTasks(v CustomOpsTasks) {
	o.CustomOpsTasks = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o ClusterTaskProgress) MarshalJSON() ([]byte, error) {
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
	if o.StartTime != nil {
		if o.StartTime.Nanosecond() == 0 {
			toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *ClusterTaskProgress) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name,omitempty"`
		Group *string `json:"group,omitempty"`
		ObjectKey common.NullableString `json:"objectKey,omitempty"`
		Message *string `json:"message,omitempty"`
		Status *string `json:"status,omitempty"`
		StartTime *time.Time `json:"startTime,omitempty"`
		EndTime common.NullableTime `json:"endTime,omitempty"`
		CustomOpsName common.NullableString `json:"customOpsName,omitempty"`
		CustomOpsTasks *CustomOpsTasks `json:"customOpsTasks,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "name", "group", "objectKey", "message", "status", "startTime", "endTime", "customOpsName", "customOpsTasks",  })
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
	if  all.CustomOpsTasks != nil && all.CustomOpsTasks.UnparsedObject != nil && o.UnparsedObject == nil {
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
