// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TaskSchedule struct {
	// Schedule type for task execution
	Type TaskScheduleType `json:"type"`
	// Start time of the execution window (required for scheduled type)
	StartTime *time.Time `json:"startTime,omitempty"`
	// End time of the execution window (required for scheduled type)
	EndTime *time.Time `json:"endTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTaskSchedule instantiates a new TaskSchedule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTaskSchedule(typeVar TaskScheduleType) *TaskSchedule {
	this := TaskSchedule{}
	this.Type = typeVar
	return &this
}

// NewTaskScheduleWithDefaults instantiates a new TaskSchedule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTaskScheduleWithDefaults() *TaskSchedule {
	this := TaskSchedule{}
	return &this
}

// GetType returns the Type field value.
func (o *TaskSchedule) GetType() TaskScheduleType {
	if o == nil {
		var ret TaskScheduleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaskSchedule) GetTypeOk() (*TaskScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TaskSchedule) SetType(v TaskScheduleType) {
	o.Type = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TaskSchedule) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedule) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TaskSchedule) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *TaskSchedule) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TaskSchedule) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskSchedule) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TaskSchedule) HasEndTime() bool {
	return o != nil && o.EndTime != nil
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *TaskSchedule) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TaskSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
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
func (o *TaskSchedule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type      *TaskScheduleType `json:"type"`
		StartTime *time.Time        `json:"startTime,omitempty"`
		EndTime   *time.Time        `json:"endTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "startTime", "endTime"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.StartTime = all.StartTime
	o.EndTime = all.EndTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
