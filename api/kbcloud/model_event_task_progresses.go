// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventTaskProgresses eventTaskProgresses is a list of event task progress detail
type EventTaskProgresses struct {
	// Items is the list of event task progress detail in the list
	Items []EventTaskProgress `json:"items,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventTaskProgresses instantiates a new EventTaskProgresses object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventTaskProgresses() *EventTaskProgresses {
	this := EventTaskProgresses{}
	return &this
}

// NewEventTaskProgressesWithDefaults instantiates a new EventTaskProgresses object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventTaskProgressesWithDefaults() *EventTaskProgresses {
	this := EventTaskProgresses{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *EventTaskProgresses) GetItems() []EventTaskProgress {
	if o == nil || o.Items == nil {
		var ret []EventTaskProgress
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTaskProgresses) GetItemsOk() (*[]EventTaskProgress, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *EventTaskProgresses) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []EventTaskProgress and assigns it to the Items field.
func (o *EventTaskProgresses) SetItems(v []EventTaskProgress) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventTaskProgresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventTaskProgresses) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items []EventTaskProgress `json:"items,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items"})
	} else {
		return err
	}
	o.Items = all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
