// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsPagination struct {
	RowsCount  *int32 `json:"rows_count,omitempty"`
	Page       *int32 `json:"page,omitempty"`
	PagesCount *int32 `json:"pages_count,omitempty"`
	PerPage    *int32 `json:"per_page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsPagination instantiates a new DmsPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsPagination() *DmsPagination {
	this := DmsPagination{}
	return &this
}

// NewDmsPaginationWithDefaults instantiates a new DmsPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsPaginationWithDefaults() *DmsPagination {
	this := DmsPagination{}
	return &this
}

// GetRowsCount returns the RowsCount field value if set, zero value otherwise.
func (o *DmsPagination) GetRowsCount() int32 {
	if o == nil || o.RowsCount == nil {
		var ret int32
		return ret
	}
	return *o.RowsCount
}

// GetRowsCountOk returns a tuple with the RowsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsPagination) GetRowsCountOk() (*int32, bool) {
	if o == nil || o.RowsCount == nil {
		return nil, false
	}
	return o.RowsCount, true
}

// HasRowsCount returns a boolean if a field has been set.
func (o *DmsPagination) HasRowsCount() bool {
	return o != nil && o.RowsCount != nil
}

// SetRowsCount gets a reference to the given int32 and assigns it to the RowsCount field.
func (o *DmsPagination) SetRowsCount(v int32) {
	o.RowsCount = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DmsPagination) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsPagination) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DmsPagination) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *DmsPagination) SetPage(v int32) {
	o.Page = &v
}

// GetPagesCount returns the PagesCount field value if set, zero value otherwise.
func (o *DmsPagination) GetPagesCount() int32 {
	if o == nil || o.PagesCount == nil {
		var ret int32
		return ret
	}
	return *o.PagesCount
}

// GetPagesCountOk returns a tuple with the PagesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsPagination) GetPagesCountOk() (*int32, bool) {
	if o == nil || o.PagesCount == nil {
		return nil, false
	}
	return o.PagesCount, true
}

// HasPagesCount returns a boolean if a field has been set.
func (o *DmsPagination) HasPagesCount() bool {
	return o != nil && o.PagesCount != nil
}

// SetPagesCount gets a reference to the given int32 and assigns it to the PagesCount field.
func (o *DmsPagination) SetPagesCount(v int32) {
	o.PagesCount = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *DmsPagination) GetPerPage() int32 {
	if o == nil || o.PerPage == nil {
		var ret int32
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsPagination) GetPerPageOk() (*int32, bool) {
	if o == nil || o.PerPage == nil {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *DmsPagination) HasPerPage() bool {
	return o != nil && o.PerPage != nil
}

// SetPerPage gets a reference to the given int32 and assigns it to the PerPage field.
func (o *DmsPagination) SetPerPage(v int32) {
	o.PerPage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.RowsCount != nil {
		toSerialize["rows_count"] = o.RowsCount
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PagesCount != nil {
		toSerialize["pages_count"] = o.PagesCount
	}
	if o.PerPage != nil {
		toSerialize["per_page"] = o.PerPage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RowsCount  *int32 `json:"rows_count,omitempty"`
		Page       *int32 `json:"page,omitempty"`
		PagesCount *int32 `json:"pages_count,omitempty"`
		PerPage    *int32 `json:"per_page,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"rows_count", "page", "pages_count", "per_page"})
	} else {
		return err
	}
	o.RowsCount = all.RowsCount
	o.Page = all.Page
	o.PagesCount = all.PagesCount
	o.PerPage = all.PerPage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
