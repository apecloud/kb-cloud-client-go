// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PaginationResult Pagination information
type PaginationResult struct {
	// Current page number (1-based)
	Page int32 `json:"page"`
	// Number of records per page
	PageSize int32 `json:"pageSize"`
	// Total number of records
	Total int32 `json:"total"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPaginationResult instantiates a new PaginationResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPaginationResult(page int32, pageSize int32, total int32) *PaginationResult {
	this := PaginationResult{}
	this.Page = page
	this.PageSize = pageSize
	this.Total = total
	return &this
}

// NewPaginationResultWithDefaults instantiates a new PaginationResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPaginationResultWithDefaults() *PaginationResult {
	this := PaginationResult{}
	return &this
}

// GetPage returns the Page field value.
func (o *PaginationResult) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PaginationResult) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value.
func (o *PaginationResult) SetPage(v int32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value.
func (o *PaginationResult) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *PaginationResult) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value.
func (o *PaginationResult) SetPageSize(v int32) {
	o.PageSize = v
}

// GetTotal returns the Total field value.
func (o *PaginationResult) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PaginationResult) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *PaginationResult) SetTotal(v int32) {
	o.Total = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PaginationResult) MarshalJSON() ([]byte, error) {
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
func (o *PaginationResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page     *int32 `json:"page"`
		PageSize *int32 `json:"pageSize"`
		Total    *int32 `json:"total"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
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
