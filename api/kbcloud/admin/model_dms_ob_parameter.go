// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsObParameter struct {
	// The name of the parameter
	Name string `json:"name"`
	// The value of the parameter
	Value string `json:"value"`
	// The type of the parameter value
	DataType *string `json:"dataType,omitempty"`
	// The description of the parameter
	Description string `json:"description"`
	// The value options of the parameter
	Enum []interface{} `json:"enum,omitempty"`
	// The maximum value of the parameter
	Maximum *string `json:"maximum,omitempty"`
	// The minimum value of the parameter
	Minimum *string `json:"minimum,omitempty"`
	// Whether the parameter is an immutable parameter, immutable parameters cannot be modified
	Immutable *bool `json:"immutable,omitempty"`
	// Whether the parameter is variable or a parameter
	IsVariable *bool `json:"isVariable,omitempty"`
	// EditLevel represents the way the configuration item is modified.
	EditLevel *string `json:"editLevel,omitempty"`
	// Whether the parameter is read-only
	ReadOnly *bool `json:"readOnly,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsObParameter instantiates a new DmsObParameter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsObParameter(name string, value string, description string) *DmsObParameter {
	this := DmsObParameter{}
	this.Name = name
	this.Value = value
	this.Description = description
	return &this
}

// NewDmsObParameterWithDefaults instantiates a new DmsObParameter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsObParameterWithDefaults() *DmsObParameter {
	this := DmsObParameter{}
	return &this
}

// GetName returns the Name field value.
func (o *DmsObParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DmsObParameter) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value.
func (o *DmsObParameter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *DmsObParameter) SetValue(v string) {
	o.Value = v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *DmsObParameter) GetDataType() string {
	if o == nil || o.DataType == nil {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetDataTypeOk() (*string, bool) {
	if o == nil || o.DataType == nil {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *DmsObParameter) HasDataType() bool {
	return o != nil && o.DataType != nil
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *DmsObParameter) SetDataType(v string) {
	o.DataType = &v
}

// GetDescription returns the Description field value.
func (o *DmsObParameter) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *DmsObParameter) SetDescription(v string) {
	o.Description = v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *DmsObParameter) GetEnum() []interface{} {
	if o == nil || o.Enum == nil {
		var ret []interface{}
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetEnumOk() (*[]interface{}, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return &o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *DmsObParameter) HasEnum() bool {
	return o != nil && o.Enum != nil
}

// SetEnum gets a reference to the given []interface{} and assigns it to the Enum field.
func (o *DmsObParameter) SetEnum(v []interface{}) {
	o.Enum = v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *DmsObParameter) GetMaximum() string {
	if o == nil || o.Maximum == nil {
		var ret string
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetMaximumOk() (*string, bool) {
	if o == nil || o.Maximum == nil {
		return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *DmsObParameter) HasMaximum() bool {
	return o != nil && o.Maximum != nil
}

// SetMaximum gets a reference to the given string and assigns it to the Maximum field.
func (o *DmsObParameter) SetMaximum(v string) {
	o.Maximum = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *DmsObParameter) GetMinimum() string {
	if o == nil || o.Minimum == nil {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetMinimumOk() (*string, bool) {
	if o == nil || o.Minimum == nil {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *DmsObParameter) HasMinimum() bool {
	return o != nil && o.Minimum != nil
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *DmsObParameter) SetMinimum(v string) {
	o.Minimum = &v
}

// GetImmutable returns the Immutable field value if set, zero value otherwise.
func (o *DmsObParameter) GetImmutable() bool {
	if o == nil || o.Immutable == nil {
		var ret bool
		return ret
	}
	return *o.Immutable
}

// GetImmutableOk returns a tuple with the Immutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetImmutableOk() (*bool, bool) {
	if o == nil || o.Immutable == nil {
		return nil, false
	}
	return o.Immutable, true
}

// HasImmutable returns a boolean if a field has been set.
func (o *DmsObParameter) HasImmutable() bool {
	return o != nil && o.Immutable != nil
}

// SetImmutable gets a reference to the given bool and assigns it to the Immutable field.
func (o *DmsObParameter) SetImmutable(v bool) {
	o.Immutable = &v
}

// GetIsVariable returns the IsVariable field value if set, zero value otherwise.
func (o *DmsObParameter) GetIsVariable() bool {
	if o == nil || o.IsVariable == nil {
		var ret bool
		return ret
	}
	return *o.IsVariable
}

// GetIsVariableOk returns a tuple with the IsVariable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetIsVariableOk() (*bool, bool) {
	if o == nil || o.IsVariable == nil {
		return nil, false
	}
	return o.IsVariable, true
}

// HasIsVariable returns a boolean if a field has been set.
func (o *DmsObParameter) HasIsVariable() bool {
	return o != nil && o.IsVariable != nil
}

// SetIsVariable gets a reference to the given bool and assigns it to the IsVariable field.
func (o *DmsObParameter) SetIsVariable(v bool) {
	o.IsVariable = &v
}

// GetEditLevel returns the EditLevel field value if set, zero value otherwise.
func (o *DmsObParameter) GetEditLevel() string {
	if o == nil || o.EditLevel == nil {
		var ret string
		return ret
	}
	return *o.EditLevel
}

// GetEditLevelOk returns a tuple with the EditLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetEditLevelOk() (*string, bool) {
	if o == nil || o.EditLevel == nil {
		return nil, false
	}
	return o.EditLevel, true
}

// HasEditLevel returns a boolean if a field has been set.
func (o *DmsObParameter) HasEditLevel() bool {
	return o != nil && o.EditLevel != nil
}

// SetEditLevel gets a reference to the given string and assigns it to the EditLevel field.
func (o *DmsObParameter) SetEditLevel(v string) {
	o.EditLevel = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *DmsObParameter) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsObParameter) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *DmsObParameter) HasReadOnly() bool {
	return o != nil && o.ReadOnly != nil
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *DmsObParameter) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsObParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	if o.DataType != nil {
		toSerialize["dataType"] = o.DataType
	}
	toSerialize["description"] = o.Description
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}
	if o.Maximum != nil {
		toSerialize["maximum"] = o.Maximum
	}
	if o.Minimum != nil {
		toSerialize["minimum"] = o.Minimum
	}
	if o.Immutable != nil {
		toSerialize["immutable"] = o.Immutable
	}
	if o.IsVariable != nil {
		toSerialize["isVariable"] = o.IsVariable
	}
	if o.EditLevel != nil {
		toSerialize["editLevel"] = o.EditLevel
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsObParameter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string       `json:"name"`
		Value       *string       `json:"value"`
		DataType    *string       `json:"dataType,omitempty"`
		Description *string       `json:"description"`
		Enum        []interface{} `json:"enum,omitempty"`
		Maximum     *string       `json:"maximum,omitempty"`
		Minimum     *string       `json:"minimum,omitempty"`
		Immutable   *bool         `json:"immutable,omitempty"`
		IsVariable  *bool         `json:"isVariable,omitempty"`
		EditLevel   *string       `json:"editLevel,omitempty"`
		ReadOnly    *bool         `json:"readOnly,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "value", "dataType", "description", "enum", "maximum", "minimum", "immutable", "isVariable", "editLevel", "readOnly"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Value = *all.Value
	o.DataType = all.DataType
	o.Description = *all.Description
	o.Enum = all.Enum
	o.Maximum = all.Maximum
	o.Minimum = all.Minimum
	o.Immutable = all.Immutable
	o.IsVariable = all.IsVariable
	o.EditLevel = all.EditLevel
	o.ReadOnly = all.ReadOnly

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
