// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DatabaseList DatabaesList is a list of Databaes

type DatabaseList struct {
	// Items is the list of Databases objects in the list
	Items []DatabaseItem `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatabaseList instantiates a new DatabaseList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabaseList(items []DatabaseItem) *DatabaseList {
	this := DatabaseList{}
	this.Items = items
	return &this
}

// NewDatabaseListWithDefaults instantiates a new DatabaseList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabaseListWithDefaults() *DatabaseList {
	this := DatabaseList{}
	return &this
}

// GetItems returns the Items field value.
func (o *DatabaseList) GetItems() []DatabaseItem {
	if o == nil {
		var ret []DatabaseItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *DatabaseList) GetItemsOk() (*[]DatabaseItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *DatabaseList) SetItems(v []DatabaseItem) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabaseList) MarshalJSON() ([]byte, error) {
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
func (o *DatabaseList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items *[]DatabaseItem `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
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
