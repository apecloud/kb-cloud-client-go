// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventFilterOptionList A collection of filter results.
type EventFilterOptionList struct {
	// An array containing individual filter result items.
	Items []EventFilterOption `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventFilterOptionList instantiates a new EventFilterOptionList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventFilterOptionList(items []EventFilterOption) *EventFilterOptionList {
	this := EventFilterOptionList{}
	this.Items = items
	return &this
}

// NewEventFilterOptionListWithDefaults instantiates a new EventFilterOptionList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventFilterOptionListWithDefaults() *EventFilterOptionList {
	this := EventFilterOptionList{}
	return &this
}

// GetItems returns the Items field value.
func (o *EventFilterOptionList) GetItems() []EventFilterOption {
	if o == nil {
		var ret []EventFilterOption
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EventFilterOptionList) GetItemsOk() (*[]EventFilterOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *EventFilterOptionList) SetItems(v []EventFilterOption) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventFilterOptionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventFilterOptionList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items *[]EventFilterOption `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items"})
	} else {
		return err
	}
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
