// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type EventObject struct {
	EventType  *EventType `json:"eventType,omitempty"`
	EventValue *string    `json:"eventValue,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventObject instantiates a new EventObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventObject() *EventObject {
	this := EventObject{}
	return &this
}

// NewEventObjectWithDefaults instantiates a new EventObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventObjectWithDefaults() *EventObject {
	this := EventObject{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EventObject) GetEventType() EventType {
	if o == nil || o.EventType == nil {
		var ret EventType
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventObject) GetEventTypeOk() (*EventType, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EventObject) HasEventType() bool {
	return o != nil && o.EventType != nil
}

// SetEventType gets a reference to the given EventType and assigns it to the EventType field.
func (o *EventObject) SetEventType(v EventType) {
	o.EventType = &v
}

// GetEventValue returns the EventValue field value if set, zero value otherwise.
func (o *EventObject) GetEventValue() string {
	if o == nil || o.EventValue == nil {
		var ret string
		return ret
	}
	return *o.EventValue
}

// GetEventValueOk returns a tuple with the EventValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventObject) GetEventValueOk() (*string, bool) {
	if o == nil || o.EventValue == nil {
		return nil, false
	}
	return o.EventValue, true
}

// HasEventValue returns a boolean if a field has been set.
func (o *EventObject) HasEventValue() bool {
	return o != nil && o.EventValue != nil
}

// SetEventValue gets a reference to the given string and assigns it to the EventValue field.
func (o *EventObject) SetEventValue(v string) {
	o.EventValue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.EventValue != nil {
		toSerialize["eventValue"] = o.EventValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventType  *EventType `json:"eventType,omitempty"`
		EventValue *string    `json:"eventValue,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"eventType", "eventValue"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.EventType != nil && !all.EventType.IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = all.EventType
	}
	o.EventValue = all.EventValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
