// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanSummary struct {
	TotalCost         common.NullableFloat64 `json:"totalCost,omitempty"`
	TotalActualTimeMs common.NullableFloat64 `json:"totalActualTimeMs,omitempty"`
	PlanRows          common.NullableFloat64 `json:"planRows,omitempty"`
	ActualRows        common.NullableFloat64 `json:"actualRows,omitempty"`
	NodeCount         *int64                 `json:"nodeCount,omitempty"`
	MaxCostNodeId     common.NullableString  `json:"maxCostNodeId,omitempty"`
	MaxRowsNodeId     common.NullableString  `json:"maxRowsNodeId,omitempty"`
	MaxTimeNodeId     common.NullableString  `json:"maxTimeNodeId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExecutionPlanSummary instantiates a new DmsExecutionPlanSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExecutionPlanSummary() *DmsExecutionPlanSummary {
	this := DmsExecutionPlanSummary{}
	return &this
}

// NewDmsExecutionPlanSummaryWithDefaults instantiates a new DmsExecutionPlanSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExecutionPlanSummaryWithDefaults() *DmsExecutionPlanSummary {
	this := DmsExecutionPlanSummary{}
	return &this
}

// GetTotalCost returns the TotalCost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanSummary) GetTotalCost() float64 {
	if o == nil || o.TotalCost.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TotalCost.Get()
}

// GetTotalCostOk returns a tuple with the TotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanSummary) GetTotalCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCost.Get(), o.TotalCost.IsSet()
}

// HasTotalCost returns a boolean if a field has been set.
func (o *DmsExecutionPlanSummary) HasTotalCost() bool {
	return o != nil && o.TotalCost.IsSet()
}

// SetTotalCost gets a reference to the given common.NullableFloat64 and assigns it to the TotalCost field.
func (o *DmsExecutionPlanSummary) SetTotalCost(v float64) {
	o.TotalCost.Set(&v)
}

// SetTotalCostNil sets the value for TotalCost to be an explicit nil.
func (o *DmsExecutionPlanSummary) SetTotalCostNil() {
	o.TotalCost.Set(nil)
}

// UnsetTotalCost ensures that no value is present for TotalCost, not even an explicit nil.
func (o *DmsExecutionPlanSummary) UnsetTotalCost() {
	o.TotalCost.Unset()
}

// GetTotalActualTimeMs returns the TotalActualTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanSummary) GetTotalActualTimeMs() float64 {
	if o == nil || o.TotalActualTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TotalActualTimeMs.Get()
}

// GetTotalActualTimeMsOk returns a tuple with the TotalActualTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanSummary) GetTotalActualTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalActualTimeMs.Get(), o.TotalActualTimeMs.IsSet()
}

// HasTotalActualTimeMs returns a boolean if a field has been set.
func (o *DmsExecutionPlanSummary) HasTotalActualTimeMs() bool {
	return o != nil && o.TotalActualTimeMs.IsSet()
}

// SetTotalActualTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the TotalActualTimeMs field.
func (o *DmsExecutionPlanSummary) SetTotalActualTimeMs(v float64) {
	o.TotalActualTimeMs.Set(&v)
}

// SetTotalActualTimeMsNil sets the value for TotalActualTimeMs to be an explicit nil.
func (o *DmsExecutionPlanSummary) SetTotalActualTimeMsNil() {
	o.TotalActualTimeMs.Set(nil)
}

// UnsetTotalActualTimeMs ensures that no value is present for TotalActualTimeMs, not even an explicit nil.
func (o *DmsExecutionPlanSummary) UnsetTotalActualTimeMs() {
	o.TotalActualTimeMs.Unset()
}

// GetPlanRows returns the PlanRows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanSummary) GetPlanRows() float64 {
	if o == nil || o.PlanRows.Get() == nil {
		var ret float64
		return ret
	}
	return *o.PlanRows.Get()
}

// GetPlanRowsOk returns a tuple with the PlanRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanSummary) GetPlanRowsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanRows.Get(), o.PlanRows.IsSet()
}

// HasPlanRows returns a boolean if a field has been set.
func (o *DmsExecutionPlanSummary) HasPlanRows() bool {
	return o != nil && o.PlanRows.IsSet()
}

// SetPlanRows gets a reference to the given common.NullableFloat64 and assigns it to the PlanRows field.
func (o *DmsExecutionPlanSummary) SetPlanRows(v float64) {
	o.PlanRows.Set(&v)
}

// SetPlanRowsNil sets the value for PlanRows to be an explicit nil.
func (o *DmsExecutionPlanSummary) SetPlanRowsNil() {
	o.PlanRows.Set(nil)
}

// UnsetPlanRows ensures that no value is present for PlanRows, not even an explicit nil.
func (o *DmsExecutionPlanSummary) UnsetPlanRows() {
	o.PlanRows.Unset()
}

// GetActualRows returns the ActualRows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanSummary) GetActualRows() float64 {
	if o == nil || o.ActualRows.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ActualRows.Get()
}

// GetActualRowsOk returns a tuple with the ActualRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanSummary) GetActualRowsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActualRows.Get(), o.ActualRows.IsSet()
}

// HasActualRows returns a boolean if a field has been set.
func (o *DmsExecutionPlanSummary) HasActualRows() bool {
	return o != nil && o.ActualRows.IsSet()
}

// SetActualRows gets a reference to the given common.NullableFloat64 and assigns it to the ActualRows field.
func (o *DmsExecutionPlanSummary) SetActualRows(v float64) {
	o.ActualRows.Set(&v)
}

// SetActualRowsNil sets the value for ActualRows to be an explicit nil.
func (o *DmsExecutionPlanSummary) SetActualRowsNil() {
	o.ActualRows.Set(nil)
}

// UnsetActualRows ensures that no value is present for ActualRows, not even an explicit nil.
func (o *DmsExecutionPlanSummary) UnsetActualRows() {
	o.ActualRows.Unset()
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *DmsExecutionPlanSummary) GetNodeCount() int64 {
	if o == nil || o.NodeCount == nil {
		var ret int64
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanSummary) GetNodeCountOk() (*int64, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *DmsExecutionPlanSummary) HasNodeCount() bool {
	return o != nil && o.NodeCount != nil
}

// SetNodeCount gets a reference to the given int64 and assigns it to the NodeCount field.
func (o *DmsExecutionPlanSummary) SetNodeCount(v int64) {
	o.NodeCount = &v
}

// GetMaxCostNodeId returns the MaxCostNodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanSummary) GetMaxCostNodeId() string {
	if o == nil || o.MaxCostNodeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaxCostNodeId.Get()
}

// GetMaxCostNodeIdOk returns a tuple with the MaxCostNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanSummary) GetMaxCostNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCostNodeId.Get(), o.MaxCostNodeId.IsSet()
}

// HasMaxCostNodeId returns a boolean if a field has been set.
func (o *DmsExecutionPlanSummary) HasMaxCostNodeId() bool {
	return o != nil && o.MaxCostNodeId.IsSet()
}

// SetMaxCostNodeId gets a reference to the given common.NullableString and assigns it to the MaxCostNodeId field.
func (o *DmsExecutionPlanSummary) SetMaxCostNodeId(v string) {
	o.MaxCostNodeId.Set(&v)
}

// SetMaxCostNodeIdNil sets the value for MaxCostNodeId to be an explicit nil.
func (o *DmsExecutionPlanSummary) SetMaxCostNodeIdNil() {
	o.MaxCostNodeId.Set(nil)
}

// UnsetMaxCostNodeId ensures that no value is present for MaxCostNodeId, not even an explicit nil.
func (o *DmsExecutionPlanSummary) UnsetMaxCostNodeId() {
	o.MaxCostNodeId.Unset()
}

// GetMaxRowsNodeId returns the MaxRowsNodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanSummary) GetMaxRowsNodeId() string {
	if o == nil || o.MaxRowsNodeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaxRowsNodeId.Get()
}

// GetMaxRowsNodeIdOk returns a tuple with the MaxRowsNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanSummary) GetMaxRowsNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRowsNodeId.Get(), o.MaxRowsNodeId.IsSet()
}

// HasMaxRowsNodeId returns a boolean if a field has been set.
func (o *DmsExecutionPlanSummary) HasMaxRowsNodeId() bool {
	return o != nil && o.MaxRowsNodeId.IsSet()
}

// SetMaxRowsNodeId gets a reference to the given common.NullableString and assigns it to the MaxRowsNodeId field.
func (o *DmsExecutionPlanSummary) SetMaxRowsNodeId(v string) {
	o.MaxRowsNodeId.Set(&v)
}

// SetMaxRowsNodeIdNil sets the value for MaxRowsNodeId to be an explicit nil.
func (o *DmsExecutionPlanSummary) SetMaxRowsNodeIdNil() {
	o.MaxRowsNodeId.Set(nil)
}

// UnsetMaxRowsNodeId ensures that no value is present for MaxRowsNodeId, not even an explicit nil.
func (o *DmsExecutionPlanSummary) UnsetMaxRowsNodeId() {
	o.MaxRowsNodeId.Unset()
}

// GetMaxTimeNodeId returns the MaxTimeNodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanSummary) GetMaxTimeNodeId() string {
	if o == nil || o.MaxTimeNodeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MaxTimeNodeId.Get()
}

// GetMaxTimeNodeIdOk returns a tuple with the MaxTimeNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanSummary) GetMaxTimeNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTimeNodeId.Get(), o.MaxTimeNodeId.IsSet()
}

// HasMaxTimeNodeId returns a boolean if a field has been set.
func (o *DmsExecutionPlanSummary) HasMaxTimeNodeId() bool {
	return o != nil && o.MaxTimeNodeId.IsSet()
}

// SetMaxTimeNodeId gets a reference to the given common.NullableString and assigns it to the MaxTimeNodeId field.
func (o *DmsExecutionPlanSummary) SetMaxTimeNodeId(v string) {
	o.MaxTimeNodeId.Set(&v)
}

// SetMaxTimeNodeIdNil sets the value for MaxTimeNodeId to be an explicit nil.
func (o *DmsExecutionPlanSummary) SetMaxTimeNodeIdNil() {
	o.MaxTimeNodeId.Set(nil)
}

// UnsetMaxTimeNodeId ensures that no value is present for MaxTimeNodeId, not even an explicit nil.
func (o *DmsExecutionPlanSummary) UnsetMaxTimeNodeId() {
	o.MaxTimeNodeId.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExecutionPlanSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TotalCost.IsSet() {
		toSerialize["totalCost"] = o.TotalCost.Get()
	}
	if o.TotalActualTimeMs.IsSet() {
		toSerialize["totalActualTimeMs"] = o.TotalActualTimeMs.Get()
	}
	if o.PlanRows.IsSet() {
		toSerialize["planRows"] = o.PlanRows.Get()
	}
	if o.ActualRows.IsSet() {
		toSerialize["actualRows"] = o.ActualRows.Get()
	}
	if o.NodeCount != nil {
		toSerialize["nodeCount"] = o.NodeCount
	}
	if o.MaxCostNodeId.IsSet() {
		toSerialize["maxCostNodeId"] = o.MaxCostNodeId.Get()
	}
	if o.MaxRowsNodeId.IsSet() {
		toSerialize["maxRowsNodeId"] = o.MaxRowsNodeId.Get()
	}
	if o.MaxTimeNodeId.IsSet() {
		toSerialize["maxTimeNodeId"] = o.MaxTimeNodeId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExecutionPlanSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalCost         common.NullableFloat64 `json:"totalCost,omitempty"`
		TotalActualTimeMs common.NullableFloat64 `json:"totalActualTimeMs,omitempty"`
		PlanRows          common.NullableFloat64 `json:"planRows,omitempty"`
		ActualRows        common.NullableFloat64 `json:"actualRows,omitempty"`
		NodeCount         *int64                 `json:"nodeCount,omitempty"`
		MaxCostNodeId     common.NullableString  `json:"maxCostNodeId,omitempty"`
		MaxRowsNodeId     common.NullableString  `json:"maxRowsNodeId,omitempty"`
		MaxTimeNodeId     common.NullableString  `json:"maxTimeNodeId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"totalCost", "totalActualTimeMs", "planRows", "actualRows", "nodeCount", "maxCostNodeId", "maxRowsNodeId", "maxTimeNodeId"})
	} else {
		return err
	}
	o.TotalCost = all.TotalCost
	o.TotalActualTimeMs = all.TotalActualTimeMs
	o.PlanRows = all.PlanRows
	o.ActualRows = all.ActualRows
	o.NodeCount = all.NodeCount
	o.MaxCostNodeId = all.MaxCostNodeId
	o.MaxRowsNodeId = all.MaxRowsNodeId
	o.MaxTimeNodeId = all.MaxTimeNodeId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
