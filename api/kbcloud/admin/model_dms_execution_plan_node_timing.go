// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanNodeTiming struct {
	StartupMs common.NullableFloat64 `json:"startupMs,omitempty"`
	TotalMs   common.NullableFloat64 `json:"totalMs,omitempty"`
	SelfMs    common.NullableFloat64 `json:"selfMs,omitempty"`
	Loops     common.NullableFloat64 `json:"loops,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExecutionPlanNodeTiming instantiates a new DmsExecutionPlanNodeTiming object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExecutionPlanNodeTiming() *DmsExecutionPlanNodeTiming {
	this := DmsExecutionPlanNodeTiming{}
	return &this
}

// NewDmsExecutionPlanNodeTimingWithDefaults instantiates a new DmsExecutionPlanNodeTiming object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExecutionPlanNodeTimingWithDefaults() *DmsExecutionPlanNodeTiming {
	this := DmsExecutionPlanNodeTiming{}
	return &this
}

// GetStartupMs returns the StartupMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeTiming) GetStartupMs() float64 {
	if o == nil || o.StartupMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.StartupMs.Get()
}

// GetStartupMsOk returns a tuple with the StartupMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeTiming) GetStartupMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartupMs.Get(), o.StartupMs.IsSet()
}

// HasStartupMs returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeTiming) HasStartupMs() bool {
	return o != nil && o.StartupMs.IsSet()
}

// SetStartupMs gets a reference to the given common.NullableFloat64 and assigns it to the StartupMs field.
func (o *DmsExecutionPlanNodeTiming) SetStartupMs(v float64) {
	o.StartupMs.Set(&v)
}

// SetStartupMsNil sets the value for StartupMs to be an explicit nil.
func (o *DmsExecutionPlanNodeTiming) SetStartupMsNil() {
	o.StartupMs.Set(nil)
}

// UnsetStartupMs ensures that no value is present for StartupMs, not even an explicit nil.
func (o *DmsExecutionPlanNodeTiming) UnsetStartupMs() {
	o.StartupMs.Unset()
}

// GetTotalMs returns the TotalMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeTiming) GetTotalMs() float64 {
	if o == nil || o.TotalMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TotalMs.Get()
}

// GetTotalMsOk returns a tuple with the TotalMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeTiming) GetTotalMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalMs.Get(), o.TotalMs.IsSet()
}

// HasTotalMs returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeTiming) HasTotalMs() bool {
	return o != nil && o.TotalMs.IsSet()
}

// SetTotalMs gets a reference to the given common.NullableFloat64 and assigns it to the TotalMs field.
func (o *DmsExecutionPlanNodeTiming) SetTotalMs(v float64) {
	o.TotalMs.Set(&v)
}

// SetTotalMsNil sets the value for TotalMs to be an explicit nil.
func (o *DmsExecutionPlanNodeTiming) SetTotalMsNil() {
	o.TotalMs.Set(nil)
}

// UnsetTotalMs ensures that no value is present for TotalMs, not even an explicit nil.
func (o *DmsExecutionPlanNodeTiming) UnsetTotalMs() {
	o.TotalMs.Unset()
}

// GetSelfMs returns the SelfMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeTiming) GetSelfMs() float64 {
	if o == nil || o.SelfMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SelfMs.Get()
}

// GetSelfMsOk returns a tuple with the SelfMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeTiming) GetSelfMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelfMs.Get(), o.SelfMs.IsSet()
}

// HasSelfMs returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeTiming) HasSelfMs() bool {
	return o != nil && o.SelfMs.IsSet()
}

// SetSelfMs gets a reference to the given common.NullableFloat64 and assigns it to the SelfMs field.
func (o *DmsExecutionPlanNodeTiming) SetSelfMs(v float64) {
	o.SelfMs.Set(&v)
}

// SetSelfMsNil sets the value for SelfMs to be an explicit nil.
func (o *DmsExecutionPlanNodeTiming) SetSelfMsNil() {
	o.SelfMs.Set(nil)
}

// UnsetSelfMs ensures that no value is present for SelfMs, not even an explicit nil.
func (o *DmsExecutionPlanNodeTiming) UnsetSelfMs() {
	o.SelfMs.Unset()
}

// GetLoops returns the Loops field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNodeTiming) GetLoops() float64 {
	if o == nil || o.Loops.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Loops.Get()
}

// GetLoopsOk returns a tuple with the Loops field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNodeTiming) GetLoopsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Loops.Get(), o.Loops.IsSet()
}

// HasLoops returns a boolean if a field has been set.
func (o *DmsExecutionPlanNodeTiming) HasLoops() bool {
	return o != nil && o.Loops.IsSet()
}

// SetLoops gets a reference to the given common.NullableFloat64 and assigns it to the Loops field.
func (o *DmsExecutionPlanNodeTiming) SetLoops(v float64) {
	o.Loops.Set(&v)
}

// SetLoopsNil sets the value for Loops to be an explicit nil.
func (o *DmsExecutionPlanNodeTiming) SetLoopsNil() {
	o.Loops.Set(nil)
}

// UnsetLoops ensures that no value is present for Loops, not even an explicit nil.
func (o *DmsExecutionPlanNodeTiming) UnsetLoops() {
	o.Loops.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExecutionPlanNodeTiming) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.StartupMs.IsSet() {
		toSerialize["startupMs"] = o.StartupMs.Get()
	}
	if o.TotalMs.IsSet() {
		toSerialize["totalMs"] = o.TotalMs.Get()
	}
	if o.SelfMs.IsSet() {
		toSerialize["selfMs"] = o.SelfMs.Get()
	}
	if o.Loops.IsSet() {
		toSerialize["loops"] = o.Loops.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExecutionPlanNodeTiming) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StartupMs common.NullableFloat64 `json:"startupMs,omitempty"`
		TotalMs   common.NullableFloat64 `json:"totalMs,omitempty"`
		SelfMs    common.NullableFloat64 `json:"selfMs,omitempty"`
		Loops     common.NullableFloat64 `json:"loops,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"startupMs", "totalMs", "selfMs", "loops"})
	} else {
		return err
	}
	o.StartupMs = all.StartupMs
	o.TotalMs = all.TotalMs
	o.SelfMs = all.SelfMs
	o.Loops = all.Loops

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
