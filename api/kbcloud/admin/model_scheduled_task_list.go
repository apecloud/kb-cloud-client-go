// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ScheduledTaskList struct {
	Items []ScheduledTaskSummary `json:"items"`
	// Pagination information
	Pagination PaginationResult `json:"pagination"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduledTaskList instantiates a new ScheduledTaskList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduledTaskList(items []ScheduledTaskSummary, pagination PaginationResult) *ScheduledTaskList {
	this := ScheduledTaskList{}
	this.Items = items
	this.Pagination = pagination
	return &this
}

// NewScheduledTaskListWithDefaults instantiates a new ScheduledTaskList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduledTaskListWithDefaults() *ScheduledTaskList {
	this := ScheduledTaskList{}
	return &this
}

// GetItems returns the Items field value.
func (o *ScheduledTaskList) GetItems() []ScheduledTaskSummary {
	if o == nil {
		var ret []ScheduledTaskSummary
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ScheduledTaskList) GetItemsOk() (*[]ScheduledTaskSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *ScheduledTaskList) SetItems(v []ScheduledTaskSummary) {
	o.Items = v
}

// GetPagination returns the Pagination field value.
func (o *ScheduledTaskList) GetPagination() PaginationResult {
	if o == nil {
		var ret PaginationResult
		return ret
	}
	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ScheduledTaskList) GetPaginationOk() (*PaginationResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value.
func (o *ScheduledTaskList) SetPagination(v PaginationResult) {
	o.Pagination = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduledTaskList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	toSerialize["pagination"] = o.Pagination

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduledTaskList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      *[]ScheduledTaskSummary `json:"items"`
		Pagination *PaginationResult       `json:"pagination"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.Pagination == nil {
		return fmt.Errorf("required field pagination missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "pagination"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Items = *all.Items
	if all.Pagination.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagination = *all.Pagination

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
