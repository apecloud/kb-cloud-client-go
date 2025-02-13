// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventQueryFilterOption Represents a single filter result item.
type EventQueryFilterOption struct {
	// ID of the parameter, some parameters may not have an ID.
	Id *string `json:"id,omitempty"`
	// Text representation of the parameter.
	Text string `json:"text"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventQueryFilterOption instantiates a new EventQueryFilterOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventQueryFilterOption(text string) *EventQueryFilterOption {
	this := EventQueryFilterOption{}
	this.Text = text
	return &this
}

// NewEventQueryFilterOptionWithDefaults instantiates a new EventQueryFilterOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventQueryFilterOptionWithDefaults() *EventQueryFilterOption {
	this := EventQueryFilterOption{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventQueryFilterOption) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventQueryFilterOption) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventQueryFilterOption) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventQueryFilterOption) SetId(v string) {
	o.Id = &v
}

// GetText returns the Text field value.
func (o *EventQueryFilterOption) GetText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *EventQueryFilterOption) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value.
func (o *EventQueryFilterOption) SetText(v string) {
	o.Text = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventQueryFilterOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["text"] = o.Text

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventQueryFilterOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string `json:"id,omitempty"`
		Text *string `json:"text"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Text == nil {
		return fmt.Errorf("required field text missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "text"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Text = *all.Text

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
