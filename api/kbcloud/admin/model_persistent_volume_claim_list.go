// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// PersistentVolumeClaimList the List stands for stats for persistentvolumeclaims.
type PersistentVolumeClaimList struct {
	Items []PersistentVolumeClaim `json:"items,omitempty"`
	// the current page number
	CurrentPage common.NullableInt64 `json:"currentPage,omitempty"`
	// the total page number
	TotalPage common.NullableInt64 `json:"totalPage,omitempty"`
	// the page size
	PageSize common.NullableInt64 `json:"pageSize,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPersistentVolumeClaimList instantiates a new PersistentVolumeClaimList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPersistentVolumeClaimList() *PersistentVolumeClaimList {
	this := PersistentVolumeClaimList{}
	return &this
}

// NewPersistentVolumeClaimListWithDefaults instantiates a new PersistentVolumeClaimList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPersistentVolumeClaimListWithDefaults() *PersistentVolumeClaimList {
	this := PersistentVolumeClaimList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PersistentVolumeClaimList) GetItems() []PersistentVolumeClaim {
	if o == nil || o.Items == nil {
		var ret []PersistentVolumeClaim
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersistentVolumeClaimList) GetItemsOk() (*[]PersistentVolumeClaim, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PersistentVolumeClaimList) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []PersistentVolumeClaim and assigns it to the Items field.
func (o *PersistentVolumeClaimList) SetItems(v []PersistentVolumeClaim) {
	o.Items = v
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistentVolumeClaimList) GetCurrentPage() int64 {
	if o == nil || o.CurrentPage.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CurrentPage.Get()
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaimList) GetCurrentPageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPage.Get(), o.CurrentPage.IsSet()
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *PersistentVolumeClaimList) HasCurrentPage() bool {
	return o != nil && o.CurrentPage.IsSet()
}

// SetCurrentPage gets a reference to the given common.NullableInt64 and assigns it to the CurrentPage field.
func (o *PersistentVolumeClaimList) SetCurrentPage(v int64) {
	o.CurrentPage.Set(&v)
}

// SetCurrentPageNil sets the value for CurrentPage to be an explicit nil.
func (o *PersistentVolumeClaimList) SetCurrentPageNil() {
	o.CurrentPage.Set(nil)
}

// UnsetCurrentPage ensures that no value is present for CurrentPage, not even an explicit nil.
func (o *PersistentVolumeClaimList) UnsetCurrentPage() {
	o.CurrentPage.Unset()
}

// GetTotalPage returns the TotalPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistentVolumeClaimList) GetTotalPage() int64 {
	if o == nil || o.TotalPage.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalPage.Get()
}

// GetTotalPageOk returns a tuple with the TotalPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaimList) GetTotalPageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalPage.Get(), o.TotalPage.IsSet()
}

// HasTotalPage returns a boolean if a field has been set.
func (o *PersistentVolumeClaimList) HasTotalPage() bool {
	return o != nil && o.TotalPage.IsSet()
}

// SetTotalPage gets a reference to the given common.NullableInt64 and assigns it to the TotalPage field.
func (o *PersistentVolumeClaimList) SetTotalPage(v int64) {
	o.TotalPage.Set(&v)
}

// SetTotalPageNil sets the value for TotalPage to be an explicit nil.
func (o *PersistentVolumeClaimList) SetTotalPageNil() {
	o.TotalPage.Set(nil)
}

// UnsetTotalPage ensures that no value is present for TotalPage, not even an explicit nil.
func (o *PersistentVolumeClaimList) UnsetTotalPage() {
	o.TotalPage.Unset()
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistentVolumeClaimList) GetPageSize() int64 {
	if o == nil || o.PageSize.Get() == nil {
		var ret int64
		return ret
	}
	return *o.PageSize.Get()
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PersistentVolumeClaimList) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSize.Get(), o.PageSize.IsSet()
}

// HasPageSize returns a boolean if a field has been set.
func (o *PersistentVolumeClaimList) HasPageSize() bool {
	return o != nil && o.PageSize.IsSet()
}

// SetPageSize gets a reference to the given common.NullableInt64 and assigns it to the PageSize field.
func (o *PersistentVolumeClaimList) SetPageSize(v int64) {
	o.PageSize.Set(&v)
}

// SetPageSizeNil sets the value for PageSize to be an explicit nil.
func (o *PersistentVolumeClaimList) SetPageSizeNil() {
	o.PageSize.Set(nil)
}

// UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil.
func (o *PersistentVolumeClaimList) UnsetPageSize() {
	o.PageSize.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o PersistentVolumeClaimList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.CurrentPage.IsSet() {
		toSerialize["currentPage"] = o.CurrentPage.Get()
	}
	if o.TotalPage.IsSet() {
		toSerialize["totalPage"] = o.TotalPage.Get()
	}
	if o.PageSize.IsSet() {
		toSerialize["pageSize"] = o.PageSize.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PersistentVolumeClaimList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items       []PersistentVolumeClaim `json:"items,omitempty"`
		CurrentPage common.NullableInt64    `json:"currentPage,omitempty"`
		TotalPage   common.NullableInt64    `json:"totalPage,omitempty"`
		PageSize    common.NullableInt64    `json:"pageSize,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "currentPage", "totalPage", "pageSize"})
	} else {
		return err
	}
	o.Items = all.Items
	o.CurrentPage = all.CurrentPage
	o.TotalPage = all.TotalPage
	o.PageSize = all.PageSize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
