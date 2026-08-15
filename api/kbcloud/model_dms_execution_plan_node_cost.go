// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanNodeCost struct {
	Startup common.NullableFloat64 `json:"startup,omitempty"`
	Total   common.NullableFloat64 `json:"total,omitempty"`
	Self    common.NullableFloat64 `json:"self,omitempty"`
	Percent common.NullableFloat64 `json:"percent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExecutionPlanNodeCost instantiates a new DmsExecutionPlanNodeCost object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExecutionPlanNodeCost() *DmsExecutionPlanNodeCost {
	this := DmsExecutionPlanNodeCost{}
	return &this
}

// NewDmsExecutionPlanNodeCostWithDefaults instantiates a new DmsExecutionPlanNodeCost object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExecutionPlanNodeCostWithDefaults() *DmsExecutionPlanNodeCost {
	this := DmsExecutionPlanNodeCost{}
	return &this
}

// GetStartup returns the Startup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeCost) GetStartup() float64 {
	if o == nil || o.Startup.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Startup.Get()
}

// GetStartupOk returns a tuple with the Startup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeCost) GetStartupOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Startup.Get(), o.Startup.IsSet()
}

// HasStartup returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeCost) HasStartup() bool {
	return o != nil && o.Startup.IsSet()
}

// SetStartup gets a reference to the given common.NullableFloat64 and assigns it to the Startup field.
func (o *DmsExecutionPlanNodeCost) SetStartup(v float64) {
	o.Startup.Set(&v)
}

// SetStartupNil sets the value for Startup to be an explicit nil.
func (o *DmsExecutionPlanNodeCost) SetStartupNil() {
	o.Startup.Set(nil)
}

// UnsetStartup ensures that no value is present for Startup, not even an explicit nil.
func (o *DmsExecutionPlanNodeCost) UnsetStartup() {
	o.Startup.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeCost) GetTotal() float64 {
	if o == nil || o.Total.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeCost) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeCost) HasTotal() bool {
	return o != nil && o.Total.IsSet()
}

// SetTotal gets a reference to the given common.NullableFloat64 and assigns it to the Total field.
func (o *DmsExecutionPlanNodeCost) SetTotal(v float64) {
	o.Total.Set(&v)
}

// SetTotalNil sets the value for Total to be an explicit nil.
func (o *DmsExecutionPlanNodeCost) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil.
func (o *DmsExecutionPlanNodeCost) UnsetTotal() {
	o.Total.Unset()
}

// GetSelf returns the Self field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeCost) GetSelf() float64 {
	if o == nil || o.Self.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Self.Get()
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeCost) GetSelfOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Self.Get(), o.Self.IsSet()
}

// HasSelf returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeCost) HasSelf() bool {
	return o != nil && o.Self.IsSet()
}

// SetSelf gets a reference to the given common.NullableFloat64 and assigns it to the Self field.
func (o *DmsExecutionPlanNodeCost) SetSelf(v float64) {
	o.Self.Set(&v)
}

// SetSelfNil sets the value for Self to be an explicit nil.
func (o *DmsExecutionPlanNodeCost) SetSelfNil() {
	o.Self.Set(nil)
}

// UnsetSelf ensures that no value is present for Self, not even an explicit nil.
func (o *DmsExecutionPlanNodeCost) UnsetSelf() {
	o.Self.Unset()
}

// GetPercent returns the Percent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeCost) GetPercent() float64 {
	if o == nil || o.Percent.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Percent.Get()
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeCost) GetPercentOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Percent.Get(), o.Percent.IsSet()
}

// HasPercent returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeCost) HasPercent() bool {
	return o != nil && o.Percent.IsSet()
}

// SetPercent gets a reference to the given common.NullableFloat64 and assigns it to the Percent field.
func (o *DmsExecutionPlanNodeCost) SetPercent(v float64) {
	o.Percent.Set(&v)
}

// SetPercentNil sets the value for Percent to be an explicit nil.
func (o *DmsExecutionPlanNodeCost) SetPercentNil() {
	o.Percent.Set(nil)
}

// UnsetPercent ensures that no value is present for Percent, not even an explicit nil.
func (o *DmsExecutionPlanNodeCost) UnsetPercent() {
	o.Percent.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExecutionPlanNodeCost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Startup.IsSet() {
		toSerialize["startup"] = o.Startup.Get()
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}
	if o.Self.IsSet() {
		toSerialize["self"] = o.Self.Get()
	}
	if o.Percent.IsSet() {
		toSerialize["percent"] = o.Percent.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExecutionPlanNodeCost) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Startup common.NullableFloat64 `json:"startup,omitempty"`
		Total   common.NullableFloat64 `json:"total,omitempty"`
		Self    common.NullableFloat64 `json:"self,omitempty"`
		Percent common.NullableFloat64 `json:"percent,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"startup", "total", "self", "percent"})
	} else {
		return err
	}
	o.Startup = all.Startup
	o.Total = all.Total
	o.Self = all.Self
	o.Percent = all.Percent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
