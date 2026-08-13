// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanEdgeRows struct {
	Estimated common.NullableFloat64 `json:"estimated,omitempty"`
	Actual    common.NullableFloat64 `json:"actual,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExecutionPlanEdgeRows instantiates a new DmsExecutionPlanEdgeRows object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExecutionPlanEdgeRows() *DmsExecutionPlanEdgeRows {
	this := DmsExecutionPlanEdgeRows{}
	return &this
}

// NewDmsExecutionPlanEdgeRowsWithDefaults instantiates a new DmsExecutionPlanEdgeRows object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExecutionPlanEdgeRowsWithDefaults() *DmsExecutionPlanEdgeRows {
	this := DmsExecutionPlanEdgeRows{}
	return &this
}

// GetEstimated returns the Estimated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanEdgeRows) GetEstimated() float64 {
	if o == nil || o.Estimated.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Estimated.Get()
}

// GetEstimatedOk returns a tuple with the Estimated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanEdgeRows) GetEstimatedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Estimated.Get(), o.Estimated.IsSet()
}

// HasEstimated returns a boolean if a field has been set.
func (o *DmsExecutionPlanEdgeRows) HasEstimated() bool {
	return o != nil && o.Estimated.IsSet()
}

// SetEstimated gets a reference to the given common.NullableFloat64 and assigns it to the Estimated field.
func (o *DmsExecutionPlanEdgeRows) SetEstimated(v float64) {
	o.Estimated.Set(&v)
}

// SetEstimatedNil sets the value for Estimated to be an explicit nil.
func (o *DmsExecutionPlanEdgeRows) SetEstimatedNil() {
	o.Estimated.Set(nil)
}

// UnsetEstimated ensures that no value is present for Estimated, not even an explicit nil.
func (o *DmsExecutionPlanEdgeRows) UnsetEstimated() {
	o.Estimated.Unset()
}

// GetActual returns the Actual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanEdgeRows) GetActual() float64 {
	if o == nil || o.Actual.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Actual.Get()
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanEdgeRows) GetActualOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actual.Get(), o.Actual.IsSet()
}

// HasActual returns a boolean if a field has been set.
func (o *DmsExecutionPlanEdgeRows) HasActual() bool {
	return o != nil && o.Actual.IsSet()
}

// SetActual gets a reference to the given common.NullableFloat64 and assigns it to the Actual field.
func (o *DmsExecutionPlanEdgeRows) SetActual(v float64) {
	o.Actual.Set(&v)
}

// SetActualNil sets the value for Actual to be an explicit nil.
func (o *DmsExecutionPlanEdgeRows) SetActualNil() {
	o.Actual.Set(nil)
}

// UnsetActual ensures that no value is present for Actual, not even an explicit nil.
func (o *DmsExecutionPlanEdgeRows) UnsetActual() {
	o.Actual.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExecutionPlanEdgeRows) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExecutionPlanEdgeRows) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Estimated common.NullableFloat64 `json:"estimated,omitempty"`
		Actual    common.NullableFloat64 `json:"actual,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"estimated", "actual"})
	} else {
		return err
	}
	o.Estimated = all.Estimated
	o.Actual = all.Actual

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
