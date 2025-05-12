// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type IntegerOption struct {
	Min      int32   `json:"min"`
	Max      int32   `json:"max"`
	Default  int32   `json:"default"`
	Step     int32   `json:"step"`
	Excludes []int32 `json:"excludes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegerOption instantiates a new IntegerOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegerOption(min int32, max int32, defaultVar int32, step int32) *IntegerOption {
	this := IntegerOption{}
	this.Min = min
	this.Max = max
	this.Default = defaultVar
	this.Step = step
	return &this
}

// NewIntegerOptionWithDefaults instantiates a new IntegerOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegerOptionWithDefaults() *IntegerOption {
	this := IntegerOption{}
	return &this
}

// GetMin returns the Min field value.
func (o *IntegerOption) GetMin() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *IntegerOption) GetMinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Min, true
}

// SetMin sets field value.
func (o *IntegerOption) SetMin(v int32) {
	o.Min = v
}

// GetMax returns the Max field value.
func (o *IntegerOption) GetMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *IntegerOption) GetMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value.
func (o *IntegerOption) SetMax(v int32) {
	o.Max = v
}

// GetDefault returns the Default field value.
func (o *IntegerOption) GetDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *IntegerOption) GetDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value.
func (o *IntegerOption) SetDefault(v int32) {
	o.Default = v
}

// GetStep returns the Step field value.
func (o *IntegerOption) GetStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Step
}

// GetStepOk returns a tuple with the Step field value
// and a boolean to check if the value has been set.
func (o *IntegerOption) GetStepOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Step, true
}

// SetStep sets field value.
func (o *IntegerOption) SetStep(v int32) {
	o.Step = v
}

// GetExcludes returns the Excludes field value if set, zero value otherwise.
func (o *IntegerOption) GetExcludes() []int32 {
	if o == nil || o.Excludes == nil {
		var ret []int32
		return ret
	}
	return o.Excludes
}

// GetExcludesOk returns a tuple with the Excludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegerOption) GetExcludesOk() (*[]int32, bool) {
	if o == nil || o.Excludes == nil {
		return nil, false
	}
	return &o.Excludes, true
}

// HasExcludes returns a boolean if a field has been set.
func (o *IntegerOption) HasExcludes() bool {
	return o != nil && o.Excludes != nil
}

// SetExcludes gets a reference to the given []int32 and assigns it to the Excludes field.
func (o *IntegerOption) SetExcludes(v []int32) {
	o.Excludes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegerOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["min"] = o.Min
	toSerialize["max"] = o.Max
	toSerialize["default"] = o.Default
	toSerialize["step"] = o.Step
	if o.Excludes != nil {
		toSerialize["excludes"] = o.Excludes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegerOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Min      *int32  `json:"min"`
		Max      *int32  `json:"max"`
		Default  *int32  `json:"default"`
		Step     *int32  `json:"step"`
		Excludes []int32 `json:"excludes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
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
		common.DeleteKeys(additionalProperties, &[]string{"min", "max", "default", "step", "excludes"})
	} else {
		return err
	}
	o.Min = *all.Min
	o.Max = *all.Max
	o.Default = *all.Default
	o.Step = *all.Step
	o.Excludes = all.Excludes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
