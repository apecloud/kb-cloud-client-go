// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventList EventList is a list of operation event objects
type EventList struct {
	// Items is the list of operation event objects in the list
	Items []Event `json:"items"`
	// Pagination information
	Pagination *PaginationResult `json:"pagination,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventList instantiates a new EventList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventList(items []Event) *EventList {
	this := EventList{}
	this.Items = items
	return &this
}

// NewEventListWithDefaults instantiates a new EventList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventListWithDefaults() *EventList {
	this := EventList{}
	return &this
}

// GetItems returns the Items field value.
func (o *EventList) GetItems() []Event {
	if o == nil {
		var ret []Event
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EventList) GetItemsOk() (*[]Event, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *EventList) SetItems(v []Event) {
	o.Items = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *EventList) GetPagination() PaginationResult {
	if o == nil || o.Pagination == nil {
		var ret PaginationResult
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventList) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *EventList) HasPagination() bool {
	return o != nil && o.Pagination != nil
}

// SetPagination gets a reference to the given PaginationResult and assigns it to the Pagination field.
func (o *EventList) SetPagination(v PaginationResult) {
	o.Pagination = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      *[]Event          `json:"items"`
		Pagination *PaginationResult `json:"pagination,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "pagination"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Items = *all.Items
	if all.Pagination != nil && all.Pagination.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagination = all.Pagination

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
