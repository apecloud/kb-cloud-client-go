// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type FloatOption struct {
	// NODESCRIPTION Min
	Min float64 `json:"min"`
	// NODESCRIPTION Max
	Max float64 `json:"max"`
	// NODESCRIPTION Default
	Default float64 `json:"default"`
	// NODESCRIPTION Step
	Step float64 `json:"step"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFloatOption instantiates a new FloatOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFloatOption(min float64, max float64, defaultVar float64, step float64) *FloatOption {
	this := FloatOption{}
	this.Min = min
	this.Max = max
	this.Default = defaultVar
	this.Step = step
	return &this
}

// NewFloatOptionWithDefaults instantiates a new FloatOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFloatOptionWithDefaults() *FloatOption {
	this := FloatOption{}
	return &this
}

// GetMin returns the Min field value.
func (o *FloatOption) GetMin() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *FloatOption) GetMinOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value.
func (o *FloatOption) SetMin(v float64) {
	o.Min = v
}

// GetMax returns the Max field value.
func (o *FloatOption) GetMax() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *FloatOption) GetMaxOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value.
func (o *FloatOption) SetMax(v float64) {
	o.Max = v
}

// GetDefault returns the Default field value.
func (o *FloatOption) GetDefault() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *FloatOption) GetDefaultOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value.
func (o *FloatOption) SetDefault(v float64) {
	o.Default = v
}

// GetStep returns the Step field value.
func (o *FloatOption) GetStep() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Step
}

// GetStepOk returns a tuple with the Step field value
// and a boolean to check if the value has been set.
func (o *FloatOption) GetStepOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Step, true
}

// SetStep sets field value.
func (o *FloatOption) SetStep(v float64) {
	o.Step = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FloatOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["min"] = o.Min
	toSerialize["max"] = o.Max
	toSerialize["default"] = o.Default
	toSerialize["step"] = o.Step

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FloatOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Min     *float64 `json:"min"`
		Max     *float64 `json:"max"`
		Default *float64 `json:"default"`
		Step    *float64 `json:"step"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Min == nil {
		return fmt.Errorf("required field min missing")
	}
	if all.Max == nil {
		return fmt.Errorf("required field max missing")
	}
	if all.Default == nil {
		return fmt.Errorf("required field default missing")
	}
	if all.Step == nil {
		return fmt.Errorf("required field step missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"min", "max", "default", "step"})
	} else {
		return err
	}
	o.Min = *all.Min
	o.Max = *all.Max
	o.Default = *all.Default
	o.Step = *all.Step

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
