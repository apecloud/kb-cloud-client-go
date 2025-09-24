// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PlatformParameterConstraints platformParameter constraints including min, max, enum, default value
type PlatformParameterConstraints struct {
	// platformParameter min value
	Min *string `json:"min,omitempty"`
	// platformParameter max value
	Max *string `json:"max,omitempty"`
	// platformParameter enum constraints
	Enum []string `json:"enum,omitempty"`
	// platformParameter default value
	Default *string `json:"default,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPlatformParameterConstraints instantiates a new PlatformParameterConstraints object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPlatformParameterConstraints() *PlatformParameterConstraints {
	this := PlatformParameterConstraints{}
	return &this
}

// NewPlatformParameterConstraintsWithDefaults instantiates a new PlatformParameterConstraints object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPlatformParameterConstraintsWithDefaults() *PlatformParameterConstraints {
	this := PlatformParameterConstraints{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *PlatformParameterConstraints) GetMin() string {
	if o == nil || o.Min == nil {
		var ret string
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformParameterConstraints) GetMinOk() (*string, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *PlatformParameterConstraints) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given string and assigns it to the Min field.
func (o *PlatformParameterConstraints) SetMin(v string) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *PlatformParameterConstraints) GetMax() string {
	if o == nil || o.Max == nil {
		var ret string
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformParameterConstraints) GetMaxOk() (*string, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *PlatformParameterConstraints) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given string and assigns it to the Max field.
func (o *PlatformParameterConstraints) SetMax(v string) {
	o.Max = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *PlatformParameterConstraints) GetEnum() []string {
	if o == nil || o.Enum == nil {
		var ret []string
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformParameterConstraints) GetEnumOk() (*[]string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return &o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *PlatformParameterConstraints) HasEnum() bool {
	return o != nil && o.Enum != nil
}

// SetEnum gets a reference to the given []string and assigns it to the Enum field.
func (o *PlatformParameterConstraints) SetEnum(v []string) {
	o.Enum = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PlatformParameterConstraints) GetDefault() string {
	if o == nil || o.Default == nil {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformParameterConstraints) GetDefaultOk() (*string, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PlatformParameterConstraints) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *PlatformParameterConstraints) SetDefault(v string) {
	o.Default = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PlatformParameterConstraints) MarshalJSON() ([]byte, error) {
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
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PlatformParameterConstraints) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Min     *string  `json:"min,omitempty"`
		Max     *string  `json:"max,omitempty"`
		Enum    []string `json:"enum,omitempty"`
		Default *string  `json:"default,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"min", "max", "enum", "default"})
	} else {
		return err
	}
	o.Min = all.Min
	o.Max = all.Max
	o.Enum = all.Enum
	o.Default = all.Default

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
