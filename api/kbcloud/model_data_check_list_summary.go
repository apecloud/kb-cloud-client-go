// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataCheckListSummary struct {
	IsConsistent common.NullableBool `json:"is_consistent,omitempty"`
	MissCount    *int64              `json:"miss_count,omitempty"`
	DiffCount    *int64              `json:"diff_count,omitempty"`
	SkipCount    *int64              `json:"skip_count,omitempty"`
	SqlCount     *int64              `json:"sql_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataCheckListSummary instantiates a new DataCheckListSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataCheckListSummary() *DataCheckListSummary {
	this := DataCheckListSummary{}
	return &this
}

// NewDataCheckListSummaryWithDefaults instantiates a new DataCheckListSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataCheckListSummaryWithDefaults() *DataCheckListSummary {
	this := DataCheckListSummary{}
	return &this
}

// GetIsConsistent returns the IsConsistent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataCheckListSummary) GetIsConsistent() bool {
	if o == nil || o.IsConsistent.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsConsistent.Get()
}

// GetIsConsistentOk returns a tuple with the IsConsistent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DataCheckListSummary) GetIsConsistentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsConsistent.Get(), o.IsConsistent.IsSet()
}

// HasIsConsistent returns a boolean if a field has been set.
func (o *DataCheckListSummary) HasIsConsistent() bool {
	return o != nil && o.IsConsistent.IsSet()
}

// SetIsConsistent gets a reference to the given common.NullableBool and assigns it to the IsConsistent field.
func (o *DataCheckListSummary) SetIsConsistent(v bool) {
	o.IsConsistent.Set(&v)
}

// SetIsConsistentNil sets the value for IsConsistent to be an explicit nil.
func (o *DataCheckListSummary) SetIsConsistentNil() {
	o.IsConsistent.Set(nil)
}

// UnsetIsConsistent ensures that no value is present for IsConsistent, not even an explicit nil.
func (o *DataCheckListSummary) UnsetIsConsistent() {
	o.IsConsistent.Unset()
}

// GetMissCount returns the MissCount field value if set, zero value otherwise.
func (o *DataCheckListSummary) GetMissCount() int64 {
	if o == nil || o.MissCount == nil {
		var ret int64
		return ret
	}
	return *o.MissCount
}

// GetMissCountOk returns a tuple with the MissCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListSummary) GetMissCountOk() (*int64, bool) {
	if o == nil || o.MissCount == nil {
		return nil, false
	}
	return o.MissCount, true
}

// HasMissCount returns a boolean if a field has been set.
func (o *DataCheckListSummary) HasMissCount() bool {
	return o != nil && o.MissCount != nil
}

// SetMissCount gets a reference to the given int64 and assigns it to the MissCount field.
func (o *DataCheckListSummary) SetMissCount(v int64) {
	o.MissCount = &v
}

// GetDiffCount returns the DiffCount field value if set, zero value otherwise.
func (o *DataCheckListSummary) GetDiffCount() int64 {
	if o == nil || o.DiffCount == nil {
		var ret int64
		return ret
	}
	return *o.DiffCount
}

// GetDiffCountOk returns a tuple with the DiffCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListSummary) GetDiffCountOk() (*int64, bool) {
	if o == nil || o.DiffCount == nil {
		return nil, false
	}
	return o.DiffCount, true
}

// HasDiffCount returns a boolean if a field has been set.
func (o *DataCheckListSummary) HasDiffCount() bool {
	return o != nil && o.DiffCount != nil
}

// SetDiffCount gets a reference to the given int64 and assigns it to the DiffCount field.
func (o *DataCheckListSummary) SetDiffCount(v int64) {
	o.DiffCount = &v
}

// GetSkipCount returns the SkipCount field value if set, zero value otherwise.
func (o *DataCheckListSummary) GetSkipCount() int64 {
	if o == nil || o.SkipCount == nil {
		var ret int64
		return ret
	}
	return *o.SkipCount
}

// GetSkipCountOk returns a tuple with the SkipCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListSummary) GetSkipCountOk() (*int64, bool) {
	if o == nil || o.SkipCount == nil {
		return nil, false
	}
	return o.SkipCount, true
}

// HasSkipCount returns a boolean if a field has been set.
func (o *DataCheckListSummary) HasSkipCount() bool {
	return o != nil && o.SkipCount != nil
}

// SetSkipCount gets a reference to the given int64 and assigns it to the SkipCount field.
func (o *DataCheckListSummary) SetSkipCount(v int64) {
	o.SkipCount = &v
}

// GetSqlCount returns the SqlCount field value if set, zero value otherwise.
func (o *DataCheckListSummary) GetSqlCount() int64 {
	if o == nil || o.SqlCount == nil {
		var ret int64
		return ret
	}
	return *o.SqlCount
}

// GetSqlCountOk returns a tuple with the SqlCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckListSummary) GetSqlCountOk() (*int64, bool) {
	if o == nil || o.SqlCount == nil {
		return nil, false
	}
	return o.SqlCount, true
}

// HasSqlCount returns a boolean if a field has been set.
func (o *DataCheckListSummary) HasSqlCount() bool {
	return o != nil && o.SqlCount != nil
}

// SetSqlCount gets a reference to the given int64 and assigns it to the SqlCount field.
func (o *DataCheckListSummary) SetSqlCount(v int64) {
	o.SqlCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataCheckListSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.IsConsistent.IsSet() {
		toSerialize["is_consistent"] = o.IsConsistent.Get()
	}
	if o.MissCount != nil {
		toSerialize["miss_count"] = o.MissCount
	}
	if o.DiffCount != nil {
		toSerialize["diff_count"] = o.DiffCount
	}
	if o.SkipCount != nil {
		toSerialize["skip_count"] = o.SkipCount
	}
	if o.SqlCount != nil {
		toSerialize["sql_count"] = o.SqlCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataCheckListSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsConsistent common.NullableBool `json:"is_consistent,omitempty"`
		MissCount    *int64              `json:"miss_count,omitempty"`
		DiffCount    *int64              `json:"diff_count,omitempty"`
		SkipCount    *int64              `json:"skip_count,omitempty"`
		SqlCount     *int64              `json:"sql_count,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"is_consistent", "miss_count", "diff_count", "skip_count", "sql_count"})
	} else {
		return err
	}
	o.IsConsistent = all.IsConsistent
	o.MissCount = all.MissCount
	o.DiffCount = all.DiffCount
	o.SkipCount = all.SkipCount
	o.SqlCount = all.SqlCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
