// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanEdge struct {
	Id          *string                   `json:"id,omitempty"`
	Source      *string                   `json:"source,omitempty"`
	Target      *string                   `json:"target,omitempty"`
	Rows        *DmsExecutionPlanEdgeRows `json:"rows,omitempty"`
	WidthWeight common.NullableFloat64    `json:"widthWeight,omitempty"`
	Label       common.NullableString     `json:"label,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExecutionPlanEdge instantiates a new DmsExecutionPlanEdge object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExecutionPlanEdge() *DmsExecutionPlanEdge {
	this := DmsExecutionPlanEdge{}
	return &this
}

// NewDmsExecutionPlanEdgeWithDefaults instantiates a new DmsExecutionPlanEdge object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExecutionPlanEdgeWithDefaults() *DmsExecutionPlanEdge {
	this := DmsExecutionPlanEdge{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DmsExecutionPlanEdge) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanEdge) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DmsExecutionPlanEdge) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DmsExecutionPlanEdge) SetId(v string) {
	o.Id = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DmsExecutionPlanEdge) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanEdge) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DmsExecutionPlanEdge) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *DmsExecutionPlanEdge) SetSource(v string) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *DmsExecutionPlanEdge) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanEdge) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *DmsExecutionPlanEdge) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *DmsExecutionPlanEdge) SetTarget(v string) {
	o.Target = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *DmsExecutionPlanEdge) GetRows() DmsExecutionPlanEdgeRows {
	if o == nil || o.Rows == nil {
		var ret DmsExecutionPlanEdgeRows
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanEdge) GetRowsOk() (*DmsExecutionPlanEdgeRows, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *DmsExecutionPlanEdge) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given DmsExecutionPlanEdgeRows and assigns it to the Rows field.
func (o *DmsExecutionPlanEdge) SetRows(v DmsExecutionPlanEdgeRows) {
	o.Rows = &v
}

// GetWidthWeight returns the WidthWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanEdge) GetWidthWeight() float64 {
	if o == nil || o.WidthWeight.Get() == nil {
		var ret float64
		return ret
	}
	return *o.WidthWeight.Get()
}

// GetWidthWeightOk returns a tuple with the WidthWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanEdge) GetWidthWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WidthWeight.Get(), o.WidthWeight.IsSet()
}

// HasWidthWeight returns a boolean if a field has been set.
func (o *DmsExecutionPlanEdge) HasWidthWeight() bool {
	return o != nil && o.WidthWeight.IsSet()
}

// SetWidthWeight gets a reference to the given common.NullableFloat64 and assigns it to the WidthWeight field.
func (o *DmsExecutionPlanEdge) SetWidthWeight(v float64) {
	o.WidthWeight.Set(&v)
}

// SetWidthWeightNil sets the value for WidthWeight to be an explicit nil.
func (o *DmsExecutionPlanEdge) SetWidthWeightNil() {
	o.WidthWeight.Set(nil)
}

// UnsetWidthWeight ensures that no value is present for WidthWeight, not even an explicit nil.
func (o *DmsExecutionPlanEdge) UnsetWidthWeight() {
	o.WidthWeight.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanEdge) GetLabel() string {
	if o == nil || o.Label.Get() == nil {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanEdge) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *DmsExecutionPlanEdge) HasLabel() bool {
	return o != nil && o.Label.IsSet()
}

// SetLabel gets a reference to the given common.NullableString and assigns it to the Label field.
func (o *DmsExecutionPlanEdge) SetLabel(v string) {
	o.Label.Set(&v)
}

// SetLabelNil sets the value for Label to be an explicit nil.
func (o *DmsExecutionPlanEdge) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil.
func (o *DmsExecutionPlanEdge) UnsetLabel() {
	o.Label.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExecutionPlanEdge) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.WidthWeight.IsSet() {
		toSerialize["widthWeight"] = o.WidthWeight.Get()
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExecutionPlanEdge) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *string                   `json:"id,omitempty"`
		Source      *string                   `json:"source,omitempty"`
		Target      *string                   `json:"target,omitempty"`
		Rows        *DmsExecutionPlanEdgeRows `json:"rows,omitempty"`
		WidthWeight common.NullableFloat64    `json:"widthWeight,omitempty"`
		Label       common.NullableString     `json:"label,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "source", "target", "rows", "widthWeight", "label"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Source = all.Source
	o.Target = all.Target
	if all.Rows != nil && all.Rows.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rows = all.Rows
	o.WidthWeight = all.WidthWeight
	o.Label = all.Label

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
