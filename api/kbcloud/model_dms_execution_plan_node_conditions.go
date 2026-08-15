// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanNodeConditions struct {
	Filter    common.NullableString `json:"filter,omitempty"`
	IndexCond common.NullableString `json:"indexCond,omitempty"`
	JoinCond  common.NullableString `json:"joinCond,omitempty"`
	HashCond  common.NullableString `json:"hashCond,omitempty"`
	SortKey   []string              `json:"sortKey,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExecutionPlanNodeConditions instantiates a new DmsExecutionPlanNodeConditions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExecutionPlanNodeConditions() *DmsExecutionPlanNodeConditions {
	this := DmsExecutionPlanNodeConditions{}
	return &this
}

// NewDmsExecutionPlanNodeConditionsWithDefaults instantiates a new DmsExecutionPlanNodeConditions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExecutionPlanNodeConditionsWithDefaults() *DmsExecutionPlanNodeConditions {
	this := DmsExecutionPlanNodeConditions{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeConditions) GetFilter() string {
	if o == nil || o.Filter.Get() == nil {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeConditions) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeConditions) HasFilter() bool {
	return o != nil && o.Filter.IsSet()
}

// SetFilter gets a reference to the given common.NullableString and assigns it to the Filter field.
func (o *DmsExecutionPlanNodeConditions) SetFilter(v string) {
	o.Filter.Set(&v)
}

// SetFilterNil sets the value for Filter to be an explicit nil.
func (o *DmsExecutionPlanNodeConditions) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil.
func (o *DmsExecutionPlanNodeConditions) UnsetFilter() {
	o.Filter.Unset()
}

// GetIndexCond returns the IndexCond field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeConditions) GetIndexCond() string {
	if o == nil || o.IndexCond.Get() == nil {
		var ret string
		return ret
	}
	return *o.IndexCond.Get()
}

// GetIndexCondOk returns a tuple with the IndexCond field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeConditions) GetIndexCondOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexCond.Get(), o.IndexCond.IsSet()
}

// HasIndexCond returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeConditions) HasIndexCond() bool {
	return o != nil && o.IndexCond.IsSet()
}

// SetIndexCond gets a reference to the given common.NullableString and assigns it to the IndexCond field.
func (o *DmsExecutionPlanNodeConditions) SetIndexCond(v string) {
	o.IndexCond.Set(&v)
}

// SetIndexCondNil sets the value for IndexCond to be an explicit nil.
func (o *DmsExecutionPlanNodeConditions) SetIndexCondNil() {
	o.IndexCond.Set(nil)
}

// UnsetIndexCond ensures that no value is present for IndexCond, not even an explicit nil.
func (o *DmsExecutionPlanNodeConditions) UnsetIndexCond() {
	o.IndexCond.Unset()
}

// GetJoinCond returns the JoinCond field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeConditions) GetJoinCond() string {
	if o == nil || o.JoinCond.Get() == nil {
		var ret string
		return ret
	}
	return *o.JoinCond.Get()
}

// GetJoinCondOk returns a tuple with the JoinCond field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeConditions) GetJoinCondOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JoinCond.Get(), o.JoinCond.IsSet()
}

// HasJoinCond returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeConditions) HasJoinCond() bool {
	return o != nil && o.JoinCond.IsSet()
}

// SetJoinCond gets a reference to the given common.NullableString and assigns it to the JoinCond field.
func (o *DmsExecutionPlanNodeConditions) SetJoinCond(v string) {
	o.JoinCond.Set(&v)
}

// SetJoinCondNil sets the value for JoinCond to be an explicit nil.
func (o *DmsExecutionPlanNodeConditions) SetJoinCondNil() {
	o.JoinCond.Set(nil)
}

// UnsetJoinCond ensures that no value is present for JoinCond, not even an explicit nil.
func (o *DmsExecutionPlanNodeConditions) UnsetJoinCond() {
	o.JoinCond.Unset()
}

// GetHashCond returns the HashCond field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeConditions) GetHashCond() string {
	if o == nil || o.HashCond.Get() == nil {
		var ret string
		return ret
	}
	return *o.HashCond.Get()
}

// GetHashCondOk returns a tuple with the HashCond field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeConditions) GetHashCondOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HashCond.Get(), o.HashCond.IsSet()
}

// HasHashCond returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeConditions) HasHashCond() bool {
	return o != nil && o.HashCond.IsSet()
}

// SetHashCond gets a reference to the given common.NullableString and assigns it to the HashCond field.
func (o *DmsExecutionPlanNodeConditions) SetHashCond(v string) {
	o.HashCond.Set(&v)
}

// SetHashCondNil sets the value for HashCond to be an explicit nil.
func (o *DmsExecutionPlanNodeConditions) SetHashCondNil() {
	o.HashCond.Set(nil)
}

// UnsetHashCond ensures that no value is present for HashCond, not even an explicit nil.
func (o *DmsExecutionPlanNodeConditions) UnsetHashCond() {
	o.HashCond.Unset()
}

// GetSortKey returns the SortKey field value if set, zero value otherwise.
func (o *DmsExecutionPlanNodeConditions) GetSortKey() []string {
	if o == nil || o.SortKey == nil {
		var ret []string
		return ret
	}
	return o.SortKey
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNodeConditions) GetSortKeyOk() (*[]string, bool) {
	if o == nil || o.SortKey == nil {
		return nil, false
	}
	return &o.SortKey, true
}

// HasSortKey returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeConditions) HasSortKey() bool {
	return o != nil && o.SortKey != nil
}

// SetSortKey gets a reference to the given []string and assigns it to the SortKey field.
func (o *DmsExecutionPlanNodeConditions) SetSortKey(v []string) {
	o.SortKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExecutionPlanNodeConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if o.IndexCond.IsSet() {
		toSerialize["indexCond"] = o.IndexCond.Get()
	}
	if o.JoinCond.IsSet() {
		toSerialize["joinCond"] = o.JoinCond.Get()
	}
	if o.HashCond.IsSet() {
		toSerialize["hashCond"] = o.HashCond.Get()
	}
	if o.SortKey != nil {
		toSerialize["sortKey"] = o.SortKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExecutionPlanNodeConditions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter    common.NullableString `json:"filter,omitempty"`
		IndexCond common.NullableString `json:"indexCond,omitempty"`
		JoinCond  common.NullableString `json:"joinCond,omitempty"`
		HashCond  common.NullableString `json:"hashCond,omitempty"`
		SortKey   []string              `json:"sortKey,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"filter", "indexCond", "joinCond", "hashCond", "sortKey"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.IndexCond = all.IndexCond
	o.JoinCond = all.JoinCond
	o.HashCond = all.HashCond
	o.SortKey = all.SortKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
