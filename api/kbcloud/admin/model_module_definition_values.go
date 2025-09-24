// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ModuleDefinitionValues struct {
	ModuleValue *string                `json:"moduleValue,omitempty"`
	Order       *int32                 `json:"order,omitempty"`
	Weight      common.NullableFloat64 `json:"weight,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModuleDefinitionValues instantiates a new ModuleDefinitionValues object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModuleDefinitionValues() *ModuleDefinitionValues {
	this := ModuleDefinitionValues{}
	return &this
}

// NewModuleDefinitionValuesWithDefaults instantiates a new ModuleDefinitionValues object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModuleDefinitionValuesWithDefaults() *ModuleDefinitionValues {
	this := ModuleDefinitionValues{}
	return &this
}

// GetModuleValue returns the ModuleValue field value if set, zero value otherwise.
func (o *ModuleDefinitionValues) GetModuleValue() string {
	if o == nil || o.ModuleValue == nil {
		var ret string
		return ret
	}
	return *o.ModuleValue
}

// GetModuleValueOk returns a tuple with the ModuleValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleDefinitionValues) GetModuleValueOk() (*string, bool) {
	if o == nil || o.ModuleValue == nil {
		return nil, false
	}
	return o.ModuleValue, true
}

// HasModuleValue returns a boolean if a field has been set.
func (o *ModuleDefinitionValues) HasModuleValue() bool {
	return o != nil && o.ModuleValue != nil
}

// SetModuleValue gets a reference to the given string and assigns it to the ModuleValue field.
func (o *ModuleDefinitionValues) SetModuleValue(v string) {
	o.ModuleValue = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ModuleDefinitionValues) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleDefinitionValues) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ModuleDefinitionValues) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *ModuleDefinitionValues) SetOrder(v int32) {
	o.Order = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModuleDefinitionValues) GetWeight() float64 {
	if o == nil || o.Weight.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Weight.Get()
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ModuleDefinitionValues) GetWeightOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Weight.Get(), o.Weight.IsSet()
}

// HasWeight returns a boolean if a field has been set.
func (o *ModuleDefinitionValues) HasWeight() bool {
	return o != nil && o.Weight.IsSet()
}

// SetWeight gets a reference to the given common.NullableFloat64 and assigns it to the Weight field.
func (o *ModuleDefinitionValues) SetWeight(v float64) {
	o.Weight.Set(&v)
}

// SetWeightNil sets the value for Weight to be an explicit nil.
func (o *ModuleDefinitionValues) SetWeightNil() {
	o.Weight.Set(nil)
}

// UnsetWeight ensures that no value is present for Weight, not even an explicit nil.
func (o *ModuleDefinitionValues) UnsetWeight() {
	o.Weight.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ModuleDefinitionValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ModuleValue != nil {
		toSerialize["moduleValue"] = o.ModuleValue
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Weight.IsSet() {
		toSerialize["weight"] = o.Weight.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModuleDefinitionValues) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModuleValue *string                `json:"moduleValue,omitempty"`
		Order       *int32                 `json:"order,omitempty"`
		Weight      common.NullableFloat64 `json:"weight,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"moduleValue", "order", "weight"})
	} else {
		return err
	}
	o.ModuleValue = all.ModuleValue
	o.Order = all.Order
	o.Weight = all.Weight

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
