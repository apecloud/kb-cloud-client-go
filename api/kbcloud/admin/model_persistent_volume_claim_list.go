// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PersistentVolumeClaimList the List stands for stats for persistentvolumeclaims.
// NODESCRIPTION PersistentVolumeClaimList
//
// Deprecated: This model is deprecated.
type PersistentVolumeClaimList struct {
	// NODESCRIPTION Items
	Items []PersistentVolumeClaim `json:"items"`
	// the current page number
	CurrentPage int64 `json:"currentPage"`
	// the total page number
	TotalPage int64 `json:"totalPage"`
	// the page size
	PageSize int64 `json:"pageSize"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPersistentVolumeClaimList instantiates a new PersistentVolumeClaimList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPersistentVolumeClaimList(items []PersistentVolumeClaim, currentPage int64, totalPage int64, pageSize int64) *PersistentVolumeClaimList {
	this := PersistentVolumeClaimList{}
	this.Items = items
	this.CurrentPage = currentPage
	this.TotalPage = totalPage
	this.PageSize = pageSize
	return &this
}

// NewPersistentVolumeClaimListWithDefaults instantiates a new PersistentVolumeClaimList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPersistentVolumeClaimListWithDefaults() *PersistentVolumeClaimList {
	this := PersistentVolumeClaimList{}
	return &this
}

// GetItems returns the Items field value.
func (o *PersistentVolumeClaimList) GetItems() []PersistentVolumeClaim {
	if o == nil {
		var ret []PersistentVolumeClaim
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaimList) GetItemsOk() (*[]PersistentVolumeClaim, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *PersistentVolumeClaimList) SetItems(v []PersistentVolumeClaim) {
	o.Items = v
}

// GetCurrentPage returns the CurrentPage field value.
func (o *PersistentVolumeClaimList) GetCurrentPage() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaimList) GetCurrentPageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPage, true
}

// SetCurrentPage sets field value.
func (o *PersistentVolumeClaimList) SetCurrentPage(v int64) {
	o.CurrentPage = v
}

// GetTotalPage returns the TotalPage field value.
func (o *PersistentVolumeClaimList) GetTotalPage() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalPage
}

// GetTotalPageOk returns a tuple with the TotalPage field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaimList) GetTotalPageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPage, true
}

// SetTotalPage sets field value.
func (o *PersistentVolumeClaimList) SetTotalPage(v int64) {
	o.TotalPage = v
}

// GetPageSize returns the PageSize field value.
func (o *PersistentVolumeClaimList) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaimList) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value.
func (o *PersistentVolumeClaimList) SetPageSize(v int64) {
	o.PageSize = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PersistentVolumeClaimList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	toSerialize["currentPage"] = o.CurrentPage
	toSerialize["totalPage"] = o.TotalPage
	toSerialize["pageSize"] = o.PageSize

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PersistentVolumeClaimList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items       *[]PersistentVolumeClaim `json:"items"`
		CurrentPage *int64                   `json:"currentPage"`
		TotalPage   *int64                   `json:"totalPage"`
		PageSize    *int64                   `json:"pageSize"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.CurrentPage == nil {
		return fmt.Errorf("required field currentPage missing")
	}
	if all.TotalPage == nil {
		return fmt.Errorf("required field totalPage missing")
	}
	if all.PageSize == nil {
		return fmt.Errorf("required field pageSize missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "currentPage", "totalPage", "pageSize"})
	} else {
		return err
	}
	o.Items = *all.Items
	o.CurrentPage = *all.CurrentPage
	o.TotalPage = *all.TotalPage
	o.PageSize = *all.PageSize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
