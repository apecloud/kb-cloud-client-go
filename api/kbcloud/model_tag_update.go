// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TagUpdate Tag update
type TagUpdate struct {
	// The key of the tag
	Key string `json:"key"`
	// the value corresponding th the key
	Value string `json:"value"`
	// the id of the tag
	Id int64 `json:"id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagUpdate instantiates a new TagUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagUpdate(key string, value string, id int64) *TagUpdate {
	this := TagUpdate{}
	this.Key = key
	this.Value = value
	this.Id = id
	return &this
}

// NewTagUpdateWithDefaults instantiates a new TagUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagUpdateWithDefaults() *TagUpdate {
	this := TagUpdate{}
	return &this
}

// GetKey returns the Key field value.
func (o *TagUpdate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TagUpdate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *TagUpdate) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value.
func (o *TagUpdate) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TagUpdate) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *TagUpdate) SetValue(v string) {
	o.Value = v
}

// GetId returns the Id field value.
func (o *TagUpdate) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagUpdate) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TagUpdate) SetId(v int64) {
	o.Id = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key   *string `json:"key"`
		Value *string `json:"value"`
		Id    *int64  `json:"id"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key", "value", "id"})
	} else {
		return err
	}
	o.Key = *all.Key
	o.Value = *all.Value
	o.Id = *all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
