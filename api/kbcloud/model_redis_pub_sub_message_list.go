// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisPubSubMessageList struct {
	Items      []RedisPubSubMessage `json:"items,omitempty"`
	NextCursor *string              `json:"nextCursor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisPubSubMessageList instantiates a new RedisPubSubMessageList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisPubSubMessageList() *RedisPubSubMessageList {
	this := RedisPubSubMessageList{}
	return &this
}

// NewRedisPubSubMessageListWithDefaults instantiates a new RedisPubSubMessageList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisPubSubMessageListWithDefaults() *RedisPubSubMessageList {
	this := RedisPubSubMessageList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *RedisPubSubMessageList) GetItems() []RedisPubSubMessage {
	if o == nil || o.Items == nil {
		var ret []RedisPubSubMessage
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubMessageList) GetItemsOk() (*[]RedisPubSubMessage, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *RedisPubSubMessageList) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []RedisPubSubMessage and assigns it to the Items field.
func (o *RedisPubSubMessageList) SetItems(v []RedisPubSubMessage) {
	o.Items = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *RedisPubSubMessageList) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisPubSubMessageList) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *RedisPubSubMessageList) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *RedisPubSubMessageList) SetNextCursor(v string) {
	o.NextCursor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisPubSubMessageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextCursor != nil {
		toSerialize["nextCursor"] = o.NextCursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisPubSubMessageList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      []RedisPubSubMessage `json:"items,omitempty"`
		NextCursor *string              `json:"nextCursor,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "nextCursor"})
	} else {
		return err
	}
	o.Items = all.Items
	o.NextCursor = all.NextCursor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
