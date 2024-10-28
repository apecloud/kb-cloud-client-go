// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObParameter struct {
	// The name of the parameter
	Name string `json:"name"`
	// The value of the parameter
	Value string `json:"value"`
	// The type of the parameter value
	DataType string `json:"dataType"`
	// The description of the parameter
	Description string `json:"description"`
	// The value options of the parameter
	Enum []map[string]interface{} `json:"enum"`
	// The maximum value of the parameter
	Maximum float64 `json:"maximum"`
	// The minimum value of the parameter
	Minimum float64 `json:"minimum"`
	// Whether the parameter is an immutable parameter, immutable parameters cannot be modified
	Immutable *bool `json:"immutable,omitempty"`
	// Whether the parameter is variable or a parameter
	IsVariable *bool `json:"isVariable,omitempty"`
	// EditLevel represents the way the configuration item is modified.
	EditLevel *string `json:"editLevel,omitempty"`
	// Whether the parameter is read-only
	ReadOnly bool `json:"readOnly"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObParameter instantiates a new ObParameter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObParameter(name string, value string, dataType string, description string, enum []map[string]interface{}, maximum float64, minimum float64, readOnly bool) *ObParameter {
	this := ObParameter{}
	this.Name = name
	this.Value = value
	this.DataType = dataType
	this.Description = description
	this.Enum = enum
	this.Maximum = maximum
	this.Minimum = minimum
	this.ReadOnly = readOnly
	return &this
}

// NewObParameterWithDefaults instantiates a new ObParameter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObParameterWithDefaults() *ObParameter {
	this := ObParameter{}
	return &this
}

// GetName returns the Name field value.
func (o *ObParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObParameter) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value.
func (o *ObParameter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ObParameter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ObParameter) SetValue(v string) {
	o.Value = v
}

// GetDataType returns the DataType field value.
func (o *ObParameter) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *ObParameter) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value.
func (o *ObParameter) SetDataType(v string) {
	o.DataType = v
}

// GetDescription returns the Description field value.
func (o *ObParameter) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ObParameter) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ObParameter) SetDescription(v string) {
	o.Description = v
}

// GetEnum returns the Enum field value.
func (o *ObParameter) GetEnum() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value
// and a boolean to check if the value has been set.
func (o *ObParameter) GetEnumOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enum, true
}

// SetEnum sets field value.
func (o *ObParameter) SetEnum(v []map[string]interface{}) {
	o.Enum = v
}

// GetMaximum returns the Maximum field value.
func (o *ObParameter) GetMaximum() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value
// and a boolean to check if the value has been set.
func (o *ObParameter) GetMaximumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maximum, true
}

// SetMaximum sets field value.
func (o *ObParameter) SetMaximum(v float64) {
	o.Maximum = v
}

// GetMinimum returns the Minimum field value.
func (o *ObParameter) GetMinimum() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value
// and a boolean to check if the value has been set.
func (o *ObParameter) GetMinimumOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minimum, true
}

// SetMinimum sets field value.
func (o *ObParameter) SetMinimum(v float64) {
	o.Minimum = v
}

// GetImmutable returns the Immutable field value if set, zero value otherwise.
func (o *ObParameter) GetImmutable() bool {
	if o == nil || o.Immutable == nil {
		var ret bool
		return ret
	}
	return *o.Immutable
}

// GetImmutableOk returns a tuple with the Immutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObParameter) GetImmutableOk() (*bool, bool) {
	if o == nil || o.Immutable == nil {
		return nil, false
	}
	return o.Immutable, true
}

// HasImmutable returns a boolean if a field has been set.
func (o *ObParameter) HasImmutable() bool {
	return o != nil && o.Immutable != nil
}

// SetImmutable gets a reference to the given bool and assigns it to the Immutable field.
func (o *ObParameter) SetImmutable(v bool) {
	o.Immutable = &v
}

// GetIsVariable returns the IsVariable field value if set, zero value otherwise.
func (o *ObParameter) GetIsVariable() bool {
	if o == nil || o.IsVariable == nil {
		var ret bool
		return ret
	}
	return *o.IsVariable
}

// GetIsVariableOk returns a tuple with the IsVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObParameter) GetIsVariableOk() (*bool, bool) {
	if o == nil || o.IsVariable == nil {
		return nil, false
	}
	return o.IsVariable, true
}

// HasIsVariable returns a boolean if a field has been set.
func (o *ObParameter) HasIsVariable() bool {
	return o != nil && o.IsVariable != nil
}

// SetIsVariable gets a reference to the given bool and assigns it to the IsVariable field.
func (o *ObParameter) SetIsVariable(v bool) {
	o.IsVariable = &v
}

// GetEditLevel returns the EditLevel field value if set, zero value otherwise.
func (o *ObParameter) GetEditLevel() string {
	if o == nil || o.EditLevel == nil {
		var ret string
		return ret
	}
	return *o.EditLevel
}

// GetEditLevelOk returns a tuple with the EditLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObParameter) GetEditLevelOk() (*string, bool) {
	if o == nil || o.EditLevel == nil {
		return nil, false
	}
	return o.EditLevel, true
}

// HasEditLevel returns a boolean if a field has been set.
func (o *ObParameter) HasEditLevel() bool {
	return o != nil && o.EditLevel != nil
}

// SetEditLevel gets a reference to the given string and assigns it to the EditLevel field.
func (o *ObParameter) SetEditLevel(v string) {
	o.EditLevel = &v
}

// GetReadOnly returns the ReadOnly field value.
func (o *ObParameter) GetReadOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value
// and a boolean to check if the value has been set.
func (o *ObParameter) GetReadOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadOnly, true
}

// SetReadOnly sets field value.
func (o *ObParameter) SetReadOnly(v bool) {
	o.ReadOnly = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	toSerialize["dataType"] = o.DataType
	toSerialize["description"] = o.Description
	toSerialize["enum"] = o.Enum
	toSerialize["maximum"] = o.Maximum
	toSerialize["minimum"] = o.Minimum
	if o.Immutable != nil {
		toSerialize["immutable"] = o.Immutable
	}
	if o.IsVariable != nil {
		toSerialize["isVariable"] = o.IsVariable
	}
	if o.EditLevel != nil {
		toSerialize["editLevel"] = o.EditLevel
	}
	toSerialize["readOnly"] = o.ReadOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObParameter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string                   `json:"name"`
		Value       *string                   `json:"value"`
		DataType    *string                   `json:"dataType"`
		Description *string                   `json:"description"`
		Enum        *[]map[string]interface{} `json:"enum"`
		Maximum     *float64                  `json:"maximum"`
		Minimum     *float64                  `json:"minimum"`
		Immutable   *bool                     `json:"immutable,omitempty"`
		IsVariable  *bool                     `json:"isVariable,omitempty"`
		EditLevel   *string                   `json:"editLevel,omitempty"`
		ReadOnly    *bool                     `json:"readOnly"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	if all.DataType == nil {
		return fmt.Errorf("required field dataType missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Enum == nil {
		return fmt.Errorf("required field enum missing")
	}
	if all.Maximum == nil {
		return fmt.Errorf("required field maximum missing")
	}
	if all.Minimum == nil {
		return fmt.Errorf("required field minimum missing")
	}
	if all.ReadOnly == nil {
		return fmt.Errorf("required field readOnly missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "value", "dataType", "description", "enum", "maximum", "minimum", "immutable", "isVariable", "editLevel", "readOnly"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Value = *all.Value
	o.DataType = *all.DataType
	o.Description = *all.Description
	o.Enum = *all.Enum
	o.Maximum = *all.Maximum
	o.Minimum = *all.Minimum
	o.Immutable = all.Immutable
	o.IsVariable = all.IsVariable
	o.EditLevel = all.EditLevel
	o.ReadOnly = *all.ReadOnly

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
