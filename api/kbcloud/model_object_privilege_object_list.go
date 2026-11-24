// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeObjectList struct {
	Items   []ObjectPrivilegeObject `json:"items"`
	HasMore bool                    `json:"hasMore"`
	// Offset for the next page; omitted when exhausted.
	NextOffset *int64 `json:"nextOffset,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeObjectList instantiates a new ObjectPrivilegeObjectList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeObjectList(items []ObjectPrivilegeObject, hasMore bool) *ObjectPrivilegeObjectList {
	this := ObjectPrivilegeObjectList{}
	this.Items = items
	this.HasMore = hasMore
	return &this
}

// NewObjectPrivilegeObjectListWithDefaults instantiates a new ObjectPrivilegeObjectList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeObjectListWithDefaults() *ObjectPrivilegeObjectList {
	this := ObjectPrivilegeObjectList{}
	return &this
}

// GetItems returns the Items field value.
func (o *ObjectPrivilegeObjectList) GetItems() []ObjectPrivilegeObject {
	if o == nil {
		var ret []ObjectPrivilegeObject
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeObjectList) GetItemsOk() (*[]ObjectPrivilegeObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *ObjectPrivilegeObjectList) SetItems(v []ObjectPrivilegeObject) {
	o.Items = v
}

// GetHasMore returns the HasMore field value.
func (o *ObjectPrivilegeObjectList) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeObjectList) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value.
func (o *ObjectPrivilegeObjectList) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise.
func (o *ObjectPrivilegeObjectList) GetNextOffset() int64 {
	if o == nil || o.NextOffset == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeObjectList) GetNextOffsetOk() (*int64, bool) {
	if o == nil || o.NextOffset == nil {
		return nil, false
	}
	return o.NextOffset, true
}

// HasNextOffset returns a boolean if a field has been set.
func (o *ObjectPrivilegeObjectList) HasNextOffset() bool {
	return o != nil && o.NextOffset != nil
}

// SetNextOffset gets a reference to the given int64 and assigns it to the NextOffset field.
func (o *ObjectPrivilegeObjectList) SetNextOffset(v int64) {
	o.NextOffset = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeObjectList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	toSerialize["hasMore"] = o.HasMore
	if o.NextOffset != nil {
		toSerialize["nextOffset"] = o.NextOffset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeObjectList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      *[]ObjectPrivilegeObject `json:"items"`
		HasMore    *bool                    `json:"hasMore"`
		NextOffset *int64                   `json:"nextOffset,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.HasMore == nil {
		return fmt.Errorf("required field hasMore missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "hasMore", "nextOffset"})
	} else {
		return err
	}
	o.Items = *all.Items
	o.HasMore = *all.HasMore
	o.NextOffset = all.NextOffset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
