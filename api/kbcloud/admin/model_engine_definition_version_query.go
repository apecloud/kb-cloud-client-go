// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

type EngineDefinitionVersionQuery struct {
	// the query to get the version, if not provided, will use inherited query from engine definition
	Sql    common.NullableString `json:"sql,omitempty"`
	Column common.NullableString `json:"column,omitempty"`
	Regex  common.NullableString `json:"regex,omitempty"`
	Min    common.NullableFloat  `json:"min,omitempty"`
	Max    common.NullableFloat  `json:"max,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineDefinitionVersionQuery instantiates a new EngineDefinitionVersionQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineDefinitionVersionQuery() *EngineDefinitionVersionQuery {
	this := EngineDefinitionVersionQuery{}
	return &this
}

// NewEngineDefinitionVersionQueryWithDefaults instantiates a new EngineDefinitionVersionQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineDefinitionVersionQueryWithDefaults() *EngineDefinitionVersionQuery {
	this := EngineDefinitionVersionQuery{}
	return &this
}

// GetSql returns the Sql field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineDefinitionVersionQuery) GetSql() string {
	if o == nil || o.Sql.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sql.Get()
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineDefinitionVersionQuery) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sql.Get(), o.Sql.IsSet()
}

// HasSql returns a boolean if a field has been set.
func (o *EngineDefinitionVersionQuery) HasSql() bool {
	return o != nil && o.Sql.IsSet()
}

// SetSql gets a reference to the given common.NullableString and assigns it to the Sql field.
func (o *EngineDefinitionVersionQuery) SetSql(v string) {
	o.Sql.Set(&v)
}

// SetSqlNil sets the value for Sql to be an explicit nil.
func (o *EngineDefinitionVersionQuery) SetSqlNil() {
	o.Sql.Set(nil)
}

// UnsetSql ensures that no value is present for Sql, not even an explicit nil.
func (o *EngineDefinitionVersionQuery) UnsetSql() {
	o.Sql.Unset()
}

// GetColumn returns the Column field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineDefinitionVersionQuery) GetColumn() string {
	if o == nil || o.Column.Get() == nil {
		var ret string
		return ret
	}
	return *o.Column.Get()
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineDefinitionVersionQuery) GetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Column.Get(), o.Column.IsSet()
}

// HasColumn returns a boolean if a field has been set.
func (o *EngineDefinitionVersionQuery) HasColumn() bool {
	return o != nil && o.Column.IsSet()
}

// SetColumn gets a reference to the given common.NullableString and assigns it to the Column field.
func (o *EngineDefinitionVersionQuery) SetColumn(v string) {
	o.Column.Set(&v)
}

// SetColumnNil sets the value for Column to be an explicit nil.
func (o *EngineDefinitionVersionQuery) SetColumnNil() {
	o.Column.Set(nil)
}

// UnsetColumn ensures that no value is present for Column, not even an explicit nil.
func (o *EngineDefinitionVersionQuery) UnsetColumn() {
	o.Column.Unset()
}

// GetRegex returns the Regex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineDefinitionVersionQuery) GetRegex() string {
	if o == nil || o.Regex.Get() == nil {
		var ret string
		return ret
	}
	return *o.Regex.Get()
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineDefinitionVersionQuery) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regex.Get(), o.Regex.IsSet()
}

// HasRegex returns a boolean if a field has been set.
func (o *EngineDefinitionVersionQuery) HasRegex() bool {
	return o != nil && o.Regex.IsSet()
}

// SetRegex gets a reference to the given common.NullableString and assigns it to the Regex field.
func (o *EngineDefinitionVersionQuery) SetRegex(v string) {
	o.Regex.Set(&v)
}

// SetRegexNil sets the value for Regex to be an explicit nil.
func (o *EngineDefinitionVersionQuery) SetRegexNil() {
	o.Regex.Set(nil)
}

// UnsetRegex ensures that no value is present for Regex, not even an explicit nil.
func (o *EngineDefinitionVersionQuery) UnsetRegex() {
	o.Regex.Unset()
}

// GetMin returns the Min field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineDefinitionVersionQuery) GetMin() float {
	if o == nil || o.Min.Get() == nil {
		var ret float
		return ret
	}
	return *o.Min.Get()
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineDefinitionVersionQuery) GetMinOk() (*float, bool) {
	if o == nil {
		return nil, false
	}
	return o.Min.Get(), o.Min.IsSet()
}

// HasMin returns a boolean if a field has been set.
func (o *EngineDefinitionVersionQuery) HasMin() bool {
	return o != nil && o.Min.IsSet()
}

// SetMin gets a reference to the given common.NullableFloat and assigns it to the Min field.
func (o *EngineDefinitionVersionQuery) SetMin(v float) {
	o.Min.Set(&v)
}

// SetMinNil sets the value for Min to be an explicit nil.
func (o *EngineDefinitionVersionQuery) SetMinNil() {
	o.Min.Set(nil)
}

// UnsetMin ensures that no value is present for Min, not even an explicit nil.
func (o *EngineDefinitionVersionQuery) UnsetMin() {
	o.Min.Unset()
}

// GetMax returns the Max field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineDefinitionVersionQuery) GetMax() float {
	if o == nil || o.Max.Get() == nil {
		var ret float
		return ret
	}
	return *o.Max.Get()
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineDefinitionVersionQuery) GetMaxOk() (*float, bool) {
	if o == nil {
		return nil, false
	}
	return o.Max.Get(), o.Max.IsSet()
}

// HasMax returns a boolean if a field has been set.
func (o *EngineDefinitionVersionQuery) HasMax() bool {
	return o != nil && o.Max.IsSet()
}

// SetMax gets a reference to the given common.NullableFloat and assigns it to the Max field.
func (o *EngineDefinitionVersionQuery) SetMax(v float) {
	o.Max.Set(&v)
}

// SetMaxNil sets the value for Max to be an explicit nil.
func (o *EngineDefinitionVersionQuery) SetMaxNil() {
	o.Max.Set(nil)
}

// UnsetMax ensures that no value is present for Max, not even an explicit nil.
func (o *EngineDefinitionVersionQuery) UnsetMax() {
	o.Max.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineDefinitionVersionQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Sql.IsSet() {
		toSerialize["sql"] = o.Sql.Get()
	}
	if o.Column.IsSet() {
		toSerialize["column"] = o.Column.Get()
	}
	if o.Regex.IsSet() {
		toSerialize["regex"] = o.Regex.Get()
	}
	if o.Min.IsSet() {
		toSerialize["min"] = o.Min.Get()
	}
	if o.Max.IsSet() {
		toSerialize["max"] = o.Max.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineDefinitionVersionQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Sql    common.NullableString `json:"sql,omitempty"`
		Column common.NullableString `json:"column,omitempty"`
		Regex  common.NullableString `json:"regex,omitempty"`
		Min    common.NullableFloat  `json:"min,omitempty"`
		Max    common.NullableFloat  `json:"max,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sql", "column", "regex", "min", "max"})
	} else {
		return err
	}
	o.Sql = all.Sql
	o.Column = all.Column
	o.Regex = all.Regex
	o.Min = all.Min
	o.Max = all.Max

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
