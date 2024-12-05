// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RestoreStatusActionsItem struct {
	Message *string `json:"message,omitempty"`
	// action name
	Name *string `json:"name,omitempty"`
	// action status, enum values: [Processing, Completed, Failed]
	Status     *string    `json:"status,omitempty"`
	BackupName *string    `json:"backupName,omitempty"`
	ObjectKey  *string    `json:"objectKey,omitempty"`
	StartTime  *time.Time `json:"startTime,omitempty"`
	EndTime    *time.Time `json:"endTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestoreStatusActionsItem instantiates a new RestoreStatusActionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestoreStatusActionsItem() *RestoreStatusActionsItem {
	this := RestoreStatusActionsItem{}
	return &this
}

// NewRestoreStatusActionsItemWithDefaults instantiates a new RestoreStatusActionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestoreStatusActionsItemWithDefaults() *RestoreStatusActionsItem {
	this := RestoreStatusActionsItem{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RestoreStatusActionsItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusActionsItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RestoreStatusActionsItem) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RestoreStatusActionsItem) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RestoreStatusActionsItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusActionsItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RestoreStatusActionsItem) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RestoreStatusActionsItem) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RestoreStatusActionsItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusActionsItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RestoreStatusActionsItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RestoreStatusActionsItem) SetStatus(v string) {
	o.Status = &v
}

// GetBackupName returns the BackupName field value if set, zero value otherwise.
func (o *RestoreStatusActionsItem) GetBackupName() string {
	if o == nil || o.BackupName == nil {
		var ret string
		return ret
	}
	return *o.BackupName
}

// GetBackupNameOk returns a tuple with the BackupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusActionsItem) GetBackupNameOk() (*string, bool) {
	if o == nil || o.BackupName == nil {
		return nil, false
	}
	return o.BackupName, true
}

// HasBackupName returns a boolean if a field has been set.
func (o *RestoreStatusActionsItem) HasBackupName() bool {
	return o != nil && o.BackupName != nil
}

// SetBackupName gets a reference to the given string and assigns it to the BackupName field.
func (o *RestoreStatusActionsItem) SetBackupName(v string) {
	o.BackupName = &v
}

// GetObjectKey returns the ObjectKey field value if set, zero value otherwise.
func (o *RestoreStatusActionsItem) GetObjectKey() string {
	if o == nil || o.ObjectKey == nil {
		var ret string
		return ret
	}
	return *o.ObjectKey
}

// GetObjectKeyOk returns a tuple with the ObjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusActionsItem) GetObjectKeyOk() (*string, bool) {
	if o == nil || o.ObjectKey == nil {
		return nil, false
	}
	return o.ObjectKey, true
}

// HasObjectKey returns a boolean if a field has been set.
func (o *RestoreStatusActionsItem) HasObjectKey() bool {
	return o != nil && o.ObjectKey != nil
}

// SetObjectKey gets a reference to the given string and assigns it to the ObjectKey field.
func (o *RestoreStatusActionsItem) SetObjectKey(v string) {
	o.ObjectKey = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RestoreStatusActionsItem) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusActionsItem) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RestoreStatusActionsItem) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *RestoreStatusActionsItem) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *RestoreStatusActionsItem) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestoreStatusActionsItem) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *RestoreStatusActionsItem) HasEndTime() bool {
	return o != nil && o.EndTime != nil
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *RestoreStatusActionsItem) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestoreStatusActionsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.BackupName != nil {
		toSerialize["backupName"] = o.BackupName
	}
	if o.ObjectKey != nil {
		toSerialize["objectKey"] = o.ObjectKey
	}
	if o.StartTime != nil {
		if o.StartTime.Nanosecond() == 0 {
			toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startTime"] = o.StartTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.EndTime != nil {
		if o.EndTime.Nanosecond() == 0 {
			toSerialize["endTime"] = o.EndTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["endTime"] = o.EndTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestoreStatusActionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message    *string    `json:"message,omitempty"`
		Name       *string    `json:"name,omitempty"`
		Status     *string    `json:"status,omitempty"`
		BackupName *string    `json:"backupName,omitempty"`
		ObjectKey  *string    `json:"objectKey,omitempty"`
		StartTime  *time.Time `json:"startTime,omitempty"`
		EndTime    *time.Time `json:"endTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"message", "name", "status", "backupName", "objectKey", "startTime", "endTime"})
	} else {
		return err
	}
	o.Message = all.Message
	o.Name = all.Name
	o.Status = all.Status
	o.BackupName = all.BackupName
	o.ObjectKey = all.ObjectKey
	o.StartTime = all.StartTime
	o.EndTime = all.EndTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
