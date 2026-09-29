// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsPagination struct {
	// Zero-based offset of this showData page.
	Offset common.NullableInt64 `json:"offset,omitempty"`
	Limit  common.NullableInt64 `json:"limit,omitempty"`
	// Whether another showData page exists; determined by reading one extra row.
	HasMore common.NullableBool `json:"hasMore,omitempty"`
	// Present only when hasMore is true.
	NextOffset common.NullableInt64 `json:"nextOffset,omitempty"`
	// Exact count returned only by countOnly requests (including zero).
	TotalRows common.NullableInt64 `json:"totalRows,omitempty"`
	// Primary-key columns used for ordering; empty means ordering is not guaranteed. Pages are not a fixed snapshot.
	OrderColumns []string `json:"orderColumns,omitempty"`
	RowsCount    *int32   `json:"rows_count,omitempty"`
	Page         *int32   `json:"page,omitempty"`
	PagesCount   *int32   `json:"pages_count,omitempty"`
	PerPage      *int32   `json:"per_page,omitempty"`
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

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsPagination) GetOffset() int64 {
	if o == nil || o.Offset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsPagination) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *DmsPagination) HasOffset() bool {
	return o != nil && o.Offset.IsSet()
}

// SetOffset gets a reference to the given common.NullableInt64 and assigns it to the Offset field.
func (o *DmsPagination) SetOffset(v int64) {
	o.Offset.Set(&v)
}

// SetOffsetNil sets the value for Offset to be an explicit nil.
func (o *DmsPagination) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil.
func (o *DmsPagination) UnsetOffset() {
	o.Offset.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsPagination) GetLimit() int64 {
	if o == nil || o.Limit.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsPagination) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *DmsPagination) HasLimit() bool {
	return o != nil && o.Limit.IsSet()
}

// SetLimit gets a reference to the given common.NullableInt64 and assigns it to the Limit field.
func (o *DmsPagination) SetLimit(v int64) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil.
func (o *DmsPagination) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil.
func (o *DmsPagination) UnsetLimit() {
	o.Limit.Unset()
}

// GetHasMore returns the HasMore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsPagination) GetHasMore() bool {
	if o == nil || o.HasMore.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasMore.Get()
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsPagination) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasMore.Get(), o.HasMore.IsSet()
}

// HasHasMore returns a boolean if a field has been set.
func (o *DmsPagination) HasHasMore() bool {
	return o != nil && o.HasMore.IsSet()
}

// SetHasMore gets a reference to the given common.NullableBool and assigns it to the HasMore field.
func (o *DmsPagination) SetHasMore(v bool) {
	o.HasMore.Set(&v)
}

// SetHasMoreNil sets the value for HasMore to be an explicit nil.
func (o *DmsPagination) SetHasMoreNil() {
	o.HasMore.Set(nil)
}

// UnsetHasMore ensures that no value is present for HasMore, not even an explicit nil.
func (o *DmsPagination) UnsetHasMore() {
	o.HasMore.Unset()
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsPagination) GetNextOffset() int64 {
	if o == nil || o.NextOffset.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextOffset.Get()
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsPagination) GetNextOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextOffset.Get(), o.NextOffset.IsSet()
}

// HasNextOffset returns a boolean if a field has been set.
func (o *DmsPagination) HasNextOffset() bool {
	return o != nil && o.NextOffset.IsSet()
}

// SetNextOffset gets a reference to the given common.NullableInt64 and assigns it to the NextOffset field.
func (o *DmsPagination) SetNextOffset(v int64) {
	o.NextOffset.Set(&v)
}

// SetNextOffsetNil sets the value for NextOffset to be an explicit nil.
func (o *DmsPagination) SetNextOffsetNil() {
	o.NextOffset.Set(nil)
}

// UnsetNextOffset ensures that no value is present for NextOffset, not even an explicit nil.
func (o *DmsPagination) UnsetNextOffset() {
	o.NextOffset.Unset()
}

// GetTotalRows returns the TotalRows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsPagination) GetTotalRows() int64 {
	if o == nil || o.TotalRows.Get() == nil {
		var ret int64
		return ret
	}
	return *o.TotalRows.Get()
}

// GetTotalRowsOk returns a tuple with the TotalRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsPagination) GetTotalRowsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalRows.Get(), o.TotalRows.IsSet()
}

// HasTotalRows returns a boolean if a field has been set.
func (o *DmsPagination) HasTotalRows() bool {
	return o != nil && o.TotalRows.IsSet()
}

// SetTotalRows gets a reference to the given common.NullableInt64 and assigns it to the TotalRows field.
func (o *DmsPagination) SetTotalRows(v int64) {
	o.TotalRows.Set(&v)
}

// SetTotalRowsNil sets the value for TotalRows to be an explicit nil.
func (o *DmsPagination) SetTotalRowsNil() {
	o.TotalRows.Set(nil)
}

// UnsetTotalRows ensures that no value is present for TotalRows, not even an explicit nil.
func (o *DmsPagination) UnsetTotalRows() {
	o.TotalRows.Unset()
}

// GetOrderColumns returns the OrderColumns field value if set, zero value otherwise.
func (o *DmsPagination) GetOrderColumns() []string {
	if o == nil || o.OrderColumns == nil {
		var ret []string
		return ret
	}
	return o.OrderColumns
}

// GetOrderColumnsOk returns a tuple with the OrderColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsPagination) GetOrderColumnsOk() (*[]string, bool) {
	if o == nil || o.OrderColumns == nil {
		return nil, false
	}
	return &o.OrderColumns, true
}

// HasOrderColumns returns a boolean if a field has been set.
func (o *DmsPagination) HasOrderColumns() bool {
	return o != nil && o.OrderColumns != nil
}

// SetOrderColumns gets a reference to the given []string and assigns it to the OrderColumns field.
func (o *DmsPagination) SetOrderColumns(v []string) {
	o.OrderColumns = v
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
	if o.Offset.IsSet() {
		toSerialize["offset"] = o.Offset.Get()
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.HasMore.IsSet() {
		toSerialize["hasMore"] = o.HasMore.Get()
	}
	if o.NextOffset.IsSet() {
		toSerialize["nextOffset"] = o.NextOffset.Get()
	}
	if o.TotalRows.IsSet() {
		toSerialize["totalRows"] = o.TotalRows.Get()
	}
	if o.OrderColumns != nil {
		toSerialize["orderColumns"] = o.OrderColumns
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
		Offset       common.NullableInt64 `json:"offset,omitempty"`
		Limit        common.NullableInt64 `json:"limit,omitempty"`
		HasMore      common.NullableBool  `json:"hasMore,omitempty"`
		NextOffset   common.NullableInt64 `json:"nextOffset,omitempty"`
		TotalRows    common.NullableInt64 `json:"totalRows,omitempty"`
		OrderColumns []string             `json:"orderColumns,omitempty"`
		RowsCount    *int32               `json:"rows_count,omitempty"`
		Page         *int32               `json:"page,omitempty"`
		PagesCount   *int32               `json:"pages_count,omitempty"`
		PerPage      *int32               `json:"per_page,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"offset", "limit", "hasMore", "nextOffset", "totalRows", "orderColumns", "rows_count", "page", "pages_count", "per_page"})
	} else {
		return err
	}
	o.Offset = all.Offset
	o.Limit = all.Limit
	o.HasMore = all.HasMore
	o.NextOffset = all.NextOffset
	o.TotalRows = all.TotalRows
	o.OrderColumns = all.OrderColumns
	o.RowsCount = all.RowsCount
	o.Page = all.Page
	o.PagesCount = all.PagesCount
	o.PerPage = all.PerPage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
