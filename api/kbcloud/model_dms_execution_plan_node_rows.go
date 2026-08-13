// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanNodeRows struct {
	Estimated       common.NullableFloat64 `json:"estimated,omitempty"`
	Actual          common.NullableFloat64 `json:"actual,omitempty"`
	ActualPerLoop   common.NullableFloat64 `json:"actualPerLoop,omitempty"`
	RemovedByFilter common.NullableFloat64 `json:"removedByFilter,omitempty"`
	MismatchRatio   common.NullableFloat64 `json:"mismatchRatio,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExecutionPlanNodeRows instantiates a new DmsExecutionPlanNodeRows object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExecutionPlanNodeRows() *DmsExecutionPlanNodeRows {
	this := DmsExecutionPlanNodeRows{}
	return &this
}

// NewDmsExecutionPlanNodeRowsWithDefaults instantiates a new DmsExecutionPlanNodeRows object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExecutionPlanNodeRowsWithDefaults() *DmsExecutionPlanNodeRows {
	this := DmsExecutionPlanNodeRows{}
	return &this
}

// GetEstimated returns the Estimated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeRows) GetEstimated() float64 {
	if o == nil || o.Estimated.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Estimated.Get()
}

// GetEstimatedOk returns a tuple with the Estimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeRows) GetEstimatedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Estimated.Get(), o.Estimated.IsSet()
}

// HasEstimated returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeRows) HasEstimated() bool {
	return o != nil && o.Estimated.IsSet()
}

// SetEstimated gets a reference to the given common.NullableFloat64 and assigns it to the Estimated field.
func (o *DmsExecutionPlanNodeRows) SetEstimated(v float64) {
	o.Estimated.Set(&v)
}

// SetEstimatedNil sets the value for Estimated to be an explicit nil.
func (o *DmsExecutionPlanNodeRows) SetEstimatedNil() {
	o.Estimated.Set(nil)
}

// UnsetEstimated ensures that no value is present for Estimated, not even an explicit nil.
func (o *DmsExecutionPlanNodeRows) UnsetEstimated() {
	o.Estimated.Unset()
}

// GetActual returns the Actual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeRows) GetActual() float64 {
	if o == nil || o.Actual.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Actual.Get()
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeRows) GetActualOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actual.Get(), o.Actual.IsSet()
}

// HasActual returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeRows) HasActual() bool {
	return o != nil && o.Actual.IsSet()
}

// SetActual gets a reference to the given common.NullableFloat64 and assigns it to the Actual field.
func (o *DmsExecutionPlanNodeRows) SetActual(v float64) {
	o.Actual.Set(&v)
}

// SetActualNil sets the value for Actual to be an explicit nil.
func (o *DmsExecutionPlanNodeRows) SetActualNil() {
	o.Actual.Set(nil)
}

// UnsetActual ensures that no value is present for Actual, not even an explicit nil.
func (o *DmsExecutionPlanNodeRows) UnsetActual() {
	o.Actual.Unset()
}

// GetActualPerLoop returns the ActualPerLoop field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeRows) GetActualPerLoop() float64 {
	if o == nil || o.ActualPerLoop.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ActualPerLoop.Get()
}

// GetActualPerLoopOk returns a tuple with the ActualPerLoop field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeRows) GetActualPerLoopOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActualPerLoop.Get(), o.ActualPerLoop.IsSet()
}

// HasActualPerLoop returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeRows) HasActualPerLoop() bool {
	return o != nil && o.ActualPerLoop.IsSet()
}

// SetActualPerLoop gets a reference to the given common.NullableFloat64 and assigns it to the ActualPerLoop field.
func (o *DmsExecutionPlanNodeRows) SetActualPerLoop(v float64) {
	o.ActualPerLoop.Set(&v)
}

// SetActualPerLoopNil sets the value for ActualPerLoop to be an explicit nil.
func (o *DmsExecutionPlanNodeRows) SetActualPerLoopNil() {
	o.ActualPerLoop.Set(nil)
}

// UnsetActualPerLoop ensures that no value is present for ActualPerLoop, not even an explicit nil.
func (o *DmsExecutionPlanNodeRows) UnsetActualPerLoop() {
	o.ActualPerLoop.Unset()
}

// GetRemovedByFilter returns the RemovedByFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeRows) GetRemovedByFilter() float64 {
	if o == nil || o.RemovedByFilter.Get() == nil {
		var ret float64
		return ret
	}
	return *o.RemovedByFilter.Get()
}

// GetRemovedByFilterOk returns a tuple with the RemovedByFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeRows) GetRemovedByFilterOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovedByFilter.Get(), o.RemovedByFilter.IsSet()
}

// HasRemovedByFilter returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeRows) HasRemovedByFilter() bool {
	return o != nil && o.RemovedByFilter.IsSet()
}

// SetRemovedByFilter gets a reference to the given common.NullableFloat64 and assigns it to the RemovedByFilter field.
func (o *DmsExecutionPlanNodeRows) SetRemovedByFilter(v float64) {
	o.RemovedByFilter.Set(&v)
}

// SetRemovedByFilterNil sets the value for RemovedByFilter to be an explicit nil.
func (o *DmsExecutionPlanNodeRows) SetRemovedByFilterNil() {
	o.RemovedByFilter.Set(nil)
}

// UnsetRemovedByFilter ensures that no value is present for RemovedByFilter, not even an explicit nil.
func (o *DmsExecutionPlanNodeRows) UnsetRemovedByFilter() {
	o.RemovedByFilter.Unset()
}

// GetMismatchRatio returns the MismatchRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeRows) GetMismatchRatio() float64 {
	if o == nil || o.MismatchRatio.Get() == nil {
		var ret float64
		return ret
	}
	return *o.MismatchRatio.Get()
}

// GetMismatchRatioOk returns a tuple with the MismatchRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeRows) GetMismatchRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MismatchRatio.Get(), o.MismatchRatio.IsSet()
}

// HasMismatchRatio returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeRows) HasMismatchRatio() bool {
	return o != nil && o.MismatchRatio.IsSet()
}

// SetMismatchRatio gets a reference to the given common.NullableFloat64 and assigns it to the MismatchRatio field.
func (o *DmsExecutionPlanNodeRows) SetMismatchRatio(v float64) {
	o.MismatchRatio.Set(&v)
}

// SetMismatchRatioNil sets the value for MismatchRatio to be an explicit nil.
func (o *DmsExecutionPlanNodeRows) SetMismatchRatioNil() {
	o.MismatchRatio.Set(nil)
}

// UnsetMismatchRatio ensures that no value is present for MismatchRatio, not even an explicit nil.
func (o *DmsExecutionPlanNodeRows) UnsetMismatchRatio() {
	o.MismatchRatio.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExecutionPlanNodeRows) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Estimated.IsSet() {
		toSerialize["estimated"] = o.Estimated.Get()
	}
	if o.Actual.IsSet() {
		toSerialize["actual"] = o.Actual.Get()
	}
	if o.ActualPerLoop.IsSet() {
		toSerialize["actualPerLoop"] = o.ActualPerLoop.Get()
	}
	if o.RemovedByFilter.IsSet() {
		toSerialize["removedByFilter"] = o.RemovedByFilter.Get()
	}
	if o.MismatchRatio.IsSet() {
		toSerialize["mismatchRatio"] = o.MismatchRatio.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExecutionPlanNodeRows) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Estimated       common.NullableFloat64 `json:"estimated,omitempty"`
		Actual          common.NullableFloat64 `json:"actual,omitempty"`
		ActualPerLoop   common.NullableFloat64 `json:"actualPerLoop,omitempty"`
		RemovedByFilter common.NullableFloat64 `json:"removedByFilter,omitempty"`
		MismatchRatio   common.NullableFloat64 `json:"mismatchRatio,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"estimated", "actual", "actualPerLoop", "removedByFilter", "mismatchRatio"})
	} else {
		return err
	}
	o.Estimated = all.Estimated
	o.Actual = all.Actual
	o.ActualPerLoop = all.ActualPerLoop
	o.RemovedByFilter = all.RemovedByFilter
	o.MismatchRatio = all.MismatchRatio

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
