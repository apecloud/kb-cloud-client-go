// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TaskList struct {
	Items []TaskListItem `json:"items"`
	// Pagination information
	Pagination *PaginationResult `json:"pagination,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTaskList instantiates a new TaskList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTaskList(items []TaskListItem) *TaskList {
	this := TaskList{}
	this.Items = items
	return &this
}

// NewTaskListWithDefaults instantiates a new TaskList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTaskListWithDefaults() *TaskList {
	this := TaskList{}
	return &this
}

// GetItems returns the Items field value.
func (o *TaskList) GetItems() []TaskListItem {
	if o == nil {
		var ret []TaskListItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TaskList) GetItemsOk() (*[]TaskListItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *TaskList) SetItems(v []TaskListItem) {
	o.Items = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *TaskList) GetPagination() PaginationResult {
	if o == nil || o.Pagination == nil {
		var ret PaginationResult
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskList) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *TaskList) HasPagination() bool {
	return o != nil && o.Pagination != nil
}

// SetPagination gets a reference to the given PaginationResult and assigns it to the Pagination field.
func (o *TaskList) SetPagination(v PaginationResult) {
	o.Pagination = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TaskList) MarshalJSON() ([]byte, error) {
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
func (o *TaskList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      *[]TaskListItem   `json:"items"`
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
