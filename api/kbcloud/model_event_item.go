// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EventItem struct {
	EventId      *string               `json:"eventID,omitempty"`
	EventName    *string               `json:"eventName,omitempty"`
	EventTime    *time.Time            `json:"eventTime,omitempty"`
	EventMessage *string               `json:"eventMessage,omitempty"`
	Operator     *string               `json:"operator,omitempty"`
	OperatorId   common.NullableString `json:"operatorID,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventItem instantiates a new EventItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventItem() *EventItem {
	this := EventItem{}
	return &this
}

// NewEventItemWithDefaults instantiates a new EventItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventItemWithDefaults() *EventItem {
	this := EventItem{}
	return &this
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *EventItem) GetEventId() string {
	if o == nil || o.EventId == nil {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventItem) GetEventIdOk() (*string, bool) {
	if o == nil || o.EventId == nil {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *EventItem) HasEventId() bool {
	return o != nil && o.EventId != nil
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *EventItem) SetEventId(v string) {
	o.EventId = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *EventItem) GetEventName() string {
	if o == nil || o.EventName == nil {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventItem) GetEventNameOk() (*string, bool) {
	if o == nil || o.EventName == nil {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *EventItem) HasEventName() bool {
	return o != nil && o.EventName != nil
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *EventItem) SetEventName(v string) {
	o.EventName = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *EventItem) GetEventTime() time.Time {
	if o == nil || o.EventTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventItem) GetEventTimeOk() (*time.Time, bool) {
	if o == nil || o.EventTime == nil {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *EventItem) HasEventTime() bool {
	return o != nil && o.EventTime != nil
}

// SetEventTime gets a reference to the given time.Time and assigns it to the EventTime field.
func (o *EventItem) SetEventTime(v time.Time) {
	o.EventTime = &v
}

// GetEventMessage returns the EventMessage field value if set, zero value otherwise.
func (o *EventItem) GetEventMessage() string {
	if o == nil || o.EventMessage == nil {
		var ret string
		return ret
	}
	return *o.EventMessage
}

// GetEventMessageOk returns a tuple with the EventMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventItem) GetEventMessageOk() (*string, bool) {
	if o == nil || o.EventMessage == nil {
		return nil, false
	}
	return o.EventMessage, true
}

// HasEventMessage returns a boolean if a field has been set.
func (o *EventItem) HasEventMessage() bool {
	return o != nil && o.EventMessage != nil
}

// SetEventMessage gets a reference to the given string and assigns it to the EventMessage field.
func (o *EventItem) SetEventMessage(v string) {
	o.EventMessage = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *EventItem) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventItem) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *EventItem) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *EventItem) SetOperator(v string) {
	o.Operator = &v
}

// GetOperatorId returns the OperatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventItem) GetOperatorId() string {
	if o == nil || o.OperatorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.OperatorId.Get()
}

// GetOperatorIdOk returns a tuple with the OperatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EventItem) GetOperatorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperatorId.Get(), o.OperatorId.IsSet()
}

// HasOperatorId returns a boolean if a field has been set.
func (o *EventItem) HasOperatorId() bool {
	return o != nil && o.OperatorId.IsSet()
}

// SetOperatorId gets a reference to the given common.NullableString and assigns it to the OperatorId field.
func (o *EventItem) SetOperatorId(v string) {
	o.OperatorId.Set(&v)
}

// SetOperatorIdNil sets the value for OperatorId to be an explicit nil.
func (o *EventItem) SetOperatorIdNil() {
	o.OperatorId.Set(nil)
}

// UnsetOperatorId ensures that no value is present for OperatorId, not even an explicit nil.
func (o *EventItem) UnsetOperatorId() {
	o.OperatorId.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EventItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EventId != nil {
		toSerialize["eventID"] = o.EventId
	}
	if o.EventName != nil {
		toSerialize["eventName"] = o.EventName
	}
	if o.EventTime != nil {
		if o.EventTime.Nanosecond() == 0 {
			toSerialize["eventTime"] = o.EventTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["eventTime"] = o.EventTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.EventMessage != nil {
		toSerialize["eventMessage"] = o.EventMessage
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.OperatorId.IsSet() {
		toSerialize["operatorID"] = o.OperatorId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventId      *string               `json:"eventID,omitempty"`
		EventName    *string               `json:"eventName,omitempty"`
		EventTime    *time.Time            `json:"eventTime,omitempty"`
		EventMessage *string               `json:"eventMessage,omitempty"`
		Operator     *string               `json:"operator,omitempty"`
		OperatorId   common.NullableString `json:"operatorID,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"eventID", "eventName", "eventTime", "eventMessage", "operator", "operatorID"})
	} else {
		return err
	}
	o.EventId = all.EventId
	o.EventName = all.EventName
	o.EventTime = all.EventTime
	o.EventMessage = all.EventMessage
	o.Operator = all.Operator
	o.OperatorId = all.OperatorId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
