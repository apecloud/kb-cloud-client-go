// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ImportNumericValidation Validation rules for numeric field type
type ImportNumericValidation struct {
	// Minimum value
	Min *float64 `json:"min,omitempty"`
	// Maximum value
	Max *float64 `json:"max,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportNumericValidation instantiates a new ImportNumericValidation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportNumericValidation() *ImportNumericValidation {
	this := ImportNumericValidation{}
	return &this
}

// NewImportNumericValidationWithDefaults instantiates a new ImportNumericValidation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportNumericValidationWithDefaults() *ImportNumericValidation {
	this := ImportNumericValidation{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ImportNumericValidation) GetMin() float64 {
	if o == nil || o.Min == nil {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumericValidation) GetMinOk() (*float64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ImportNumericValidation) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *ImportNumericValidation) SetMin(v float64) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ImportNumericValidation) GetMax() float64 {
	if o == nil || o.Max == nil {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportNumericValidation) GetMaxOk() (*float64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ImportNumericValidation) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *ImportNumericValidation) SetMax(v float64) {
	o.Max = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportNumericValidation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportNumericValidation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Min *float64 `json:"min,omitempty"`
		Max *float64 `json:"max,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"min", "max"})
	} else {
		return err
	}
	o.Min = all.Min
	o.Max = all.Max

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
