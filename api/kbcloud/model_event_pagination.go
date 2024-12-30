// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EventPagination event pagination
type EventPagination struct {
	// current page
	Page int32 `json:"page"`
	// page size
	PageSize int32 `json:"pageSize"`
	// total records
	Total int32 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventPagination instantiates a new EventPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventPagination(page int32, pageSize int32, total int32) *EventPagination {
	this := EventPagination{}
	this.Page = page
	this.PageSize = pageSize
	this.Total = total
	return &this
}

// NewEventPaginationWithDefaults instantiates a new EventPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventPaginationWithDefaults() *EventPagination {
	this := EventPagination{}
	return &this
}

// GetPage returns the Page field value.
func (o *EventPagination) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *EventPagination) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value.
func (o *EventPagination) SetPage(v int32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value.
func (o *EventPagination) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *EventPagination) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value.
func (o *EventPagination) SetPageSize(v int32) {
	o.PageSize = v
}

// GetTotal returns the Total field value.
func (o *EventPagination) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *EventPagination) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *EventPagination) SetTotal(v int32) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["page"] = o.Page
	toSerialize["pageSize"] = o.PageSize
	toSerialize["total"] = o.Total

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page     *int32 `json:"page"`
		PageSize *int32 `json:"pageSize"`
		Total    *int32 `json:"total"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Page == nil {
		return fmt.Errorf("required field page missing")
	}
	if all.PageSize == nil {
		return fmt.Errorf("required field pageSize missing")
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"page", "pageSize", "total"})
	} else {
		return err
	}
	o.Page = *all.Page
	o.PageSize = *all.PageSize
	o.Total = *all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
