// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DataCheckTableSummary struct {
	Schema       *string `json:"schema,omitempty"`
	Tb           *string `json:"tb,omitempty"`
	TargetSchema *string `json:"target_schema,omitempty"`
	TargetTb     *string `json:"target_tb,omitempty"`
	CheckedCount *int64  `json:"checked_count,omitempty"`
	MissCount    *int64  `json:"miss_count,omitempty"`
	DiffCount    *int64  `json:"diff_count,omitempty"`
	SkipCount    *int64  `json:"skip_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataCheckTableSummary instantiates a new DataCheckTableSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataCheckTableSummary() *DataCheckTableSummary {
	this := DataCheckTableSummary{}
	return &this
}

// NewDataCheckTableSummaryWithDefaults instantiates a new DataCheckTableSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataCheckTableSummaryWithDefaults() *DataCheckTableSummary {
	this := DataCheckTableSummary{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *DataCheckTableSummary) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckTableSummary) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *DataCheckTableSummary) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *DataCheckTableSummary) SetSchema(v string) {
	o.Schema = &v
}

// GetTb returns the Tb field value if set, zero value otherwise.
func (o *DataCheckTableSummary) GetTb() string {
	if o == nil || o.Tb == nil {
		var ret string
		return ret
	}
	return *o.Tb
}

// GetTbOk returns a tuple with the Tb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckTableSummary) GetTbOk() (*string, bool) {
	if o == nil || o.Tb == nil {
		return nil, false
	}
	return o.Tb, true
}

// HasTb returns a boolean if a field has been set.
func (o *DataCheckTableSummary) HasTb() bool {
	return o != nil && o.Tb != nil
}

// SetTb gets a reference to the given string and assigns it to the Tb field.
func (o *DataCheckTableSummary) SetTb(v string) {
	o.Tb = &v
}

// GetTargetSchema returns the TargetSchema field value if set, zero value otherwise.
func (o *DataCheckTableSummary) GetTargetSchema() string {
	if o == nil || o.TargetSchema == nil {
		var ret string
		return ret
	}
	return *o.TargetSchema
}

// GetTargetSchemaOk returns a tuple with the TargetSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckTableSummary) GetTargetSchemaOk() (*string, bool) {
	if o == nil || o.TargetSchema == nil {
		return nil, false
	}
	return o.TargetSchema, true
}

// HasTargetSchema returns a boolean if a field has been set.
func (o *DataCheckTableSummary) HasTargetSchema() bool {
	return o != nil && o.TargetSchema != nil
}

// SetTargetSchema gets a reference to the given string and assigns it to the TargetSchema field.
func (o *DataCheckTableSummary) SetTargetSchema(v string) {
	o.TargetSchema = &v
}

// GetTargetTb returns the TargetTb field value if set, zero value otherwise.
func (o *DataCheckTableSummary) GetTargetTb() string {
	if o == nil || o.TargetTb == nil {
		var ret string
		return ret
	}
	return *o.TargetTb
}

// GetTargetTbOk returns a tuple with the TargetTb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckTableSummary) GetTargetTbOk() (*string, bool) {
	if o == nil || o.TargetTb == nil {
		return nil, false
	}
	return o.TargetTb, true
}

// HasTargetTb returns a boolean if a field has been set.
func (o *DataCheckTableSummary) HasTargetTb() bool {
	return o != nil && o.TargetTb != nil
}

// SetTargetTb gets a reference to the given string and assigns it to the TargetTb field.
func (o *DataCheckTableSummary) SetTargetTb(v string) {
	o.TargetTb = &v
}

// GetCheckedCount returns the CheckedCount field value if set, zero value otherwise.
func (o *DataCheckTableSummary) GetCheckedCount() int64 {
	if o == nil || o.CheckedCount == nil {
		var ret int64
		return ret
	}
	return *o.CheckedCount
}

// GetCheckedCountOk returns a tuple with the CheckedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckTableSummary) GetCheckedCountOk() (*int64, bool) {
	if o == nil || o.CheckedCount == nil {
		return nil, false
	}
	return o.CheckedCount, true
}

// HasCheckedCount returns a boolean if a field has been set.
func (o *DataCheckTableSummary) HasCheckedCount() bool {
	return o != nil && o.CheckedCount != nil
}

// SetCheckedCount gets a reference to the given int64 and assigns it to the CheckedCount field.
func (o *DataCheckTableSummary) SetCheckedCount(v int64) {
	o.CheckedCount = &v
}

// GetMissCount returns the MissCount field value if set, zero value otherwise.
func (o *DataCheckTableSummary) GetMissCount() int64 {
	if o == nil || o.MissCount == nil {
		var ret int64
		return ret
	}
	return *o.MissCount
}

// GetMissCountOk returns a tuple with the MissCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckTableSummary) GetMissCountOk() (*int64, bool) {
	if o == nil || o.MissCount == nil {
		return nil, false
	}
	return o.MissCount, true
}

// HasMissCount returns a boolean if a field has been set.
func (o *DataCheckTableSummary) HasMissCount() bool {
	return o != nil && o.MissCount != nil
}

// SetMissCount gets a reference to the given int64 and assigns it to the MissCount field.
func (o *DataCheckTableSummary) SetMissCount(v int64) {
	o.MissCount = &v
}

// GetDiffCount returns the DiffCount field value if set, zero value otherwise.
func (o *DataCheckTableSummary) GetDiffCount() int64 {
	if o == nil || o.DiffCount == nil {
		var ret int64
		return ret
	}
	return *o.DiffCount
}

// GetDiffCountOk returns a tuple with the DiffCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckTableSummary) GetDiffCountOk() (*int64, bool) {
	if o == nil || o.DiffCount == nil {
		return nil, false
	}
	return o.DiffCount, true
}

// HasDiffCount returns a boolean if a field has been set.
func (o *DataCheckTableSummary) HasDiffCount() bool {
	return o != nil && o.DiffCount != nil
}

// SetDiffCount gets a reference to the given int64 and assigns it to the DiffCount field.
func (o *DataCheckTableSummary) SetDiffCount(v int64) {
	o.DiffCount = &v
}

// GetSkipCount returns the SkipCount field value if set, zero value otherwise.
func (o *DataCheckTableSummary) GetSkipCount() int64 {
	if o == nil || o.SkipCount == nil {
		var ret int64
		return ret
	}
	return *o.SkipCount
}

// GetSkipCountOk returns a tuple with the SkipCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCheckTableSummary) GetSkipCountOk() (*int64, bool) {
	if o == nil || o.SkipCount == nil {
		return nil, false
	}
	return o.SkipCount, true
}

// HasSkipCount returns a boolean if a field has been set.
func (o *DataCheckTableSummary) HasSkipCount() bool {
	return o != nil && o.SkipCount != nil
}

// SetSkipCount gets a reference to the given int64 and assigns it to the SkipCount field.
func (o *DataCheckTableSummary) SetSkipCount(v int64) {
	o.SkipCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataCheckTableSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Tb != nil {
		toSerialize["tb"] = o.Tb
	}
	if o.TargetSchema != nil {
		toSerialize["target_schema"] = o.TargetSchema
	}
	if o.TargetTb != nil {
		toSerialize["target_tb"] = o.TargetTb
	}
	if o.CheckedCount != nil {
		toSerialize["checked_count"] = o.CheckedCount
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataCheckTableSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Schema       *string `json:"schema,omitempty"`
		Tb           *string `json:"tb,omitempty"`
		TargetSchema *string `json:"target_schema,omitempty"`
		TargetTb     *string `json:"target_tb,omitempty"`
		CheckedCount *int64  `json:"checked_count,omitempty"`
		MissCount    *int64  `json:"miss_count,omitempty"`
		DiffCount    *int64  `json:"diff_count,omitempty"`
		SkipCount    *int64  `json:"skip_count,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"schema", "tb", "target_schema", "target_tb", "checked_count", "miss_count", "diff_count", "skip_count"})
	} else {
		return err
	}
	o.Schema = all.Schema
	o.Tb = all.Tb
	o.TargetSchema = all.TargetSchema
	o.TargetTb = all.TargetTb
	o.CheckedCount = all.CheckedCount
	o.MissCount = all.MissCount
	o.DiffCount = all.DiffCount
	o.SkipCount = all.SkipCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
