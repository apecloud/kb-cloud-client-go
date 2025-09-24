// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PageResult PageResult info
type PageResult struct {
	// a link to the first page of results. This link is optional for collections that cannot be indexed directly to a given page
	First *string `json:"first,omitempty"`
	// a link to the last page of results. This link is optional for collections that cannot be indexed directly to a given page
	Last *string `json:"last,omitempty"`
	// a link to the next page of results. A response that does not contain a next link does not have further data to fetch
	Next *string `json:"next,omitempty"`
	// a link to the previous page of results. A response that does not contain a prev link has no previous data. This link is optional for collections that cannot be traversed backward
	Prev *string `json:"prev,omitempty"`
	// the total count of items in the list irrespective of pagination
	TotalSize *int64 `json:"totalSize,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPageResult instantiates a new PageResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPageResult() *PageResult {
	this := PageResult{}
	return &this
}

// NewPageResultWithDefaults instantiates a new PageResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPageResultWithDefaults() *PageResult {
	this := PageResult{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise.
func (o *PageResult) GetFirst() string {
	if o == nil || o.First == nil {
		var ret string
		return ret
	}
	return *o.First
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageResult) GetFirstOk() (*string, bool) {
	if o == nil || o.First == nil {
		return nil, false
	}
	return o.First, true
}

// HasFirst returns a boolean if a field has been set.
func (o *PageResult) HasFirst() bool {
	return o != nil && o.First != nil
}

// SetFirst gets a reference to the given string and assigns it to the First field.
func (o *PageResult) SetFirst(v string) {
	o.First = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *PageResult) GetLast() string {
	if o == nil || o.Last == nil {
		var ret string
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageResult) GetLastOk() (*string, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *PageResult) HasLast() bool {
	return o != nil && o.Last != nil
}

// SetLast gets a reference to the given string and assigns it to the Last field.
func (o *PageResult) SetLast(v string) {
	o.Last = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PageResult) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageResult) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PageResult) HasNext() bool {
	return o != nil && o.Next != nil
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PageResult) SetNext(v string) {
	o.Next = &v
}

// GetPrev returns the Prev field value if set, zero value otherwise.
func (o *PageResult) GetPrev() string {
	if o == nil || o.Prev == nil {
		var ret string
		return ret
	}
	return *o.Prev
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageResult) GetPrevOk() (*string, bool) {
	if o == nil || o.Prev == nil {
		return nil, false
	}
	return o.Prev, true
}

// HasPrev returns a boolean if a field has been set.
func (o *PageResult) HasPrev() bool {
	return o != nil && o.Prev != nil
}

// SetPrev gets a reference to the given string and assigns it to the Prev field.
func (o *PageResult) SetPrev(v string) {
	o.Prev = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *PageResult) GetTotalSize() int64 {
	if o == nil || o.TotalSize == nil {
		var ret int64
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageResult) GetTotalSizeOk() (*int64, bool) {
	if o == nil || o.TotalSize == nil {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *PageResult) HasTotalSize() bool {
	return o != nil && o.TotalSize != nil
}

// SetTotalSize gets a reference to the given int64 and assigns it to the TotalSize field.
func (o *PageResult) SetTotalSize(v int64) {
	o.TotalSize = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PageResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.First != nil {
		toSerialize["first"] = o.First
	}
	if o.Last != nil {
		toSerialize["last"] = o.Last
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Prev != nil {
		toSerialize["prev"] = o.Prev
	}
	if o.TotalSize != nil {
		toSerialize["totalSize"] = o.TotalSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PageResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		First     *string `json:"first,omitempty"`
		Last      *string `json:"last,omitempty"`
		Next      *string `json:"next,omitempty"`
		Prev      *string `json:"prev,omitempty"`
		TotalSize *int64  `json:"totalSize,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"first", "last", "next", "prev", "totalSize"})
	} else {
		return err
	}
	o.First = all.First
	o.Last = all.Last
	o.Next = all.Next
	o.Prev = all.Prev
	o.TotalSize = all.TotalSize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
