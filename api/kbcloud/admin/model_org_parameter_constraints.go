// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

// OrgParameterConstraints org parameter constraints including min, max, enum, default value
type OrgParameterConstraints struct {
	// org parameter min value
	Min *string `json:"min,omitempty"`
	// org parameter max value
	Max *string `json:"max,omitempty"`
	// org parameter enum constraints
	Enum []string `json:"enum,omitempty"`
	// org parameter default value
	Default *string `json:"default,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgParameterConstraints instantiates a new OrgParameterConstraints object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgParameterConstraints() *OrgParameterConstraints {
	this := OrgParameterConstraints{}
	return &this
}

// NewOrgParameterConstraintsWithDefaults instantiates a new OrgParameterConstraints object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgParameterConstraintsWithDefaults() *OrgParameterConstraints {
	this := OrgParameterConstraints{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *OrgParameterConstraints) GetMin() string {
	if o == nil || o.Min == nil {
		var ret string
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgParameterConstraints) GetMinOk() (*string, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *OrgParameterConstraints) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given string and assigns it to the Min field.
func (o *OrgParameterConstraints) SetMin(v string) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *OrgParameterConstraints) GetMax() string {
	if o == nil || o.Max == nil {
		var ret string
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgParameterConstraints) GetMaxOk() (*string, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *OrgParameterConstraints) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given string and assigns it to the Max field.
func (o *OrgParameterConstraints) SetMax(v string) {
	o.Max = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *OrgParameterConstraints) GetEnum() []string {
	if o == nil || o.Enum == nil {
		var ret []string
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgParameterConstraints) GetEnumOk() (*[]string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return &o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *OrgParameterConstraints) HasEnum() bool {
	return o != nil && o.Enum != nil
}

// SetEnum gets a reference to the given []string and assigns it to the Enum field.
func (o *OrgParameterConstraints) SetEnum(v []string) {
	o.Enum = v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *OrgParameterConstraints) GetDefault() string {
	if o == nil || o.Default == nil {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgParameterConstraints) GetDefaultOk() (*string, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *OrgParameterConstraints) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *OrgParameterConstraints) SetDefault(v string) {
	o.Default = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgParameterConstraints) MarshalJSON() ([]byte, error) {
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
func (o *OrgParameterConstraints) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Min     *string  `json:"min,omitempty"`
		Max     *string  `json:"max,omitempty"`
		Enum    []string `json:"enum,omitempty"`
		Default *string  `json:"default,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
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
