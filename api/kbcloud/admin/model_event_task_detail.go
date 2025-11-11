// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventTaskDetail eventTaskDetail is the information of the event task detail
type EventTaskDetail struct {
	// reason of the event task detail
	Reason *string `json:"reason,omitempty"`
	// type of the event task detail
	Type *string `json:"type,omitempty"`
	// status of the event task detail
	Status *string `json:"status,omitempty"`
	// message of the event task detail
	Message *string `json:"message,omitempty"`
	// last transition time of the event task detail
	LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventTaskDetail instantiates a new EventTaskDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventTaskDetail() *EventTaskDetail {
	this := EventTaskDetail{}
	return &this
}

// NewEventTaskDetailWithDefaults instantiates a new EventTaskDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventTaskDetailWithDefaults() *EventTaskDetail {
	this := EventTaskDetail{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *EventTaskDetail) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskDetail) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *EventTaskDetail) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *EventTaskDetail) SetReason(v string) {
	o.Reason = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventTaskDetail) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskDetail) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventTaskDetail) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventTaskDetail) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventTaskDetail) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskDetail) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventTaskDetail) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EventTaskDetail) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EventTaskDetail) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskDetail) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EventTaskDetail) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EventTaskDetail) SetMessage(v string) {
	o.Message = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *EventTaskDetail) GetLastTransitionTime() time.Time {
	if o == nil || o.LastTransitionTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskDetail) GetLastTransitionTimeOk() (*time.Time, bool) {
	if o == nil || o.LastTransitionTime == nil {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *EventTaskDetail) HasLastTransitionTime() bool {
	return o != nil && o.LastTransitionTime != nil
}

// SetLastTransitionTime gets a reference to the given time.Time and assigns it to the LastTransitionTime field.
func (o *EventTaskDetail) SetLastTransitionTime(v time.Time) {
	o.LastTransitionTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventTaskDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.LastTransitionTime != nil {
		if o.LastTransitionTime.Nanosecond() == 0 {
			toSerialize["lastTransitionTime"] = o.LastTransitionTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["lastTransitionTime"] = o.LastTransitionTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventTaskDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Reason             *string    `json:"reason,omitempty"`
		Type               *string    `json:"type,omitempty"`
		Status             *string    `json:"status,omitempty"`
		Message            *string    `json:"message,omitempty"`
		LastTransitionTime *time.Time `json:"lastTransitionTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"reason", "type", "status", "message", "lastTransitionTime"})
	} else {
		return err
	}
	o.Reason = all.Reason
	o.Type = all.Type
	o.Status = all.Status
	o.Message = all.Message
	o.LastTransitionTime = all.LastTransitionTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
