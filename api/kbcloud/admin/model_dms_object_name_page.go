// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsObjectNamePage struct {
	Items []string `json:"items,omitempty"`
	// Total number of object names after applying search filters.
	Total *int64 `json:"total,omitempty"`
	// Number of object names returned in this page.
	Count *int64 `json:"count,omitempty"`
	// Requested page size.
	Limit *int64 `json:"limit,omitempty"`
	// Requested offset.
	Offset *int64 `json:"offset,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsObjectNamePage instantiates a new DmsObjectNamePage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsObjectNamePage() *DmsObjectNamePage {
	this := DmsObjectNamePage{}
	return &this
}

// NewDmsObjectNamePageWithDefaults instantiates a new DmsObjectNamePage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsObjectNamePageWithDefaults() *DmsObjectNamePage {
	this := DmsObjectNamePage{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DmsObjectNamePage) GetItems() []string {
	if o == nil || o.Items == nil {
		var ret []string
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObjectNamePage) GetItemsOk() (*[]string, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DmsObjectNamePage) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []string and assigns it to the Items field.
func (o *DmsObjectNamePage) SetItems(v []string) {
	o.Items = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DmsObjectNamePage) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObjectNamePage) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DmsObjectNamePage) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *DmsObjectNamePage) SetTotal(v int64) {
	o.Total = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *DmsObjectNamePage) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObjectNamePage) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *DmsObjectNamePage) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *DmsObjectNamePage) SetCount(v int64) {
	o.Count = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *DmsObjectNamePage) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObjectNamePage) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *DmsObjectNamePage) HasLimit() bool {
	return o != nil && o.Limit != nil
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *DmsObjectNamePage) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *DmsObjectNamePage) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObjectNamePage) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *DmsObjectNamePage) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *DmsObjectNamePage) SetOffset(v int64) {
	o.Offset = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsObjectNamePage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsObjectNamePage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items  []string `json:"items,omitempty"`
		Total  *int64   `json:"total,omitempty"`
		Count  *int64   `json:"count,omitempty"`
		Limit  *int64   `json:"limit,omitempty"`
		Offset *int64   `json:"offset,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "total", "count", "limit", "offset"})
	} else {
		return err
	}
	o.Items = all.Items
	o.Total = all.Total
	o.Count = all.Count
	o.Limit = all.Limit
	o.Offset = all.Offset

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
