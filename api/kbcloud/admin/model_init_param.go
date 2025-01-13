// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InitParam struct {
	// name of the init parameter
	Name string `json:"name"`
	// name of config spec, which the parameter belongs to, equal to componentDefinition.configs[x].name
	ConfigSpecName string `json:"configSpecName"`
	// whether the parameter is required
	Required bool `json:"required"`
	// label of the parameter
	Label string `json:"label"`
	// type of the parameter
	Type string `json:"type"`
	// default value of the parameter
	DefaultValue *string `json:"defaultValue,omitempty"`
	// maximum value of the parameter
	Maximum *string `json:"maximum,omitempty"`
	// minimum value of the parameter
	Minimum *string `json:"minimum,omitempty"`
	// enum value of the parameter
	Enum []string `json:"enum,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInitParam instantiates a new InitParam object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInitParam(name string, configSpecName string, required bool, label string, typeVar string) *InitParam {
	this := InitParam{}
	this.Name = name
	this.ConfigSpecName = configSpecName
	this.Required = required
	this.Label = label
	this.Type = typeVar
	return &this
}

// NewInitParamWithDefaults instantiates a new InitParam object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInitParamWithDefaults() *InitParam {
	this := InitParam{}
	return &this
}

// GetName returns the Name field value.
func (o *InitParam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InitParam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InitParam) SetName(v string) {
	o.Name = v
}

// GetConfigSpecName returns the ConfigSpecName field value.
func (o *InitParam) GetConfigSpecName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigSpecName
}

// GetConfigSpecNameOk returns a tuple with the ConfigSpecName field value
// and a boolean to check if the value has been set.
func (o *InitParam) GetConfigSpecNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpecName, true
}

// SetConfigSpecName sets field value.
func (o *InitParam) SetConfigSpecName(v string) {
	o.ConfigSpecName = v
}

// GetRequired returns the Required field value.
func (o *InitParam) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *InitParam) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value.
func (o *InitParam) SetRequired(v bool) {
	o.Required = v
}

// GetLabel returns the Label field value.
func (o *InitParam) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *InitParam) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *InitParam) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value.
func (o *InitParam) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InitParam) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *InitParam) SetType(v string) {
	o.Type = v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *InitParam) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitParam) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *InitParam) HasDefaultValue() bool {
	return o != nil && o.DefaultValue != nil
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *InitParam) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *InitParam) GetMaximum() string {
	if o == nil || o.Maximum == nil {
		var ret string
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitParam) GetMaximumOk() (*string, bool) {
	if o == nil || o.Maximum == nil {
		return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *InitParam) HasMaximum() bool {
	return o != nil && o.Maximum != nil
}

// SetMaximum gets a reference to the given string and assigns it to the Maximum field.
func (o *InitParam) SetMaximum(v string) {
	o.Maximum = &v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *InitParam) GetMinimum() string {
	if o == nil || o.Minimum == nil {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitParam) GetMinimumOk() (*string, bool) {
	if o == nil || o.Minimum == nil {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *InitParam) HasMinimum() bool {
	return o != nil && o.Minimum != nil
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *InitParam) SetMinimum(v string) {
	o.Minimum = &v
}

// GetEnum returns the Enum field value if set, zero value otherwise.
func (o *InitParam) GetEnum() []string {
	if o == nil || o.Enum == nil {
		var ret []string
		return ret
	}
	return o.Enum
}

// GetEnumOk returns a tuple with the Enum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitParam) GetEnumOk() (*[]string, bool) {
	if o == nil || o.Enum == nil {
		return nil, false
	}
	return &o.Enum, true
}

// HasEnum returns a boolean if a field has been set.
func (o *InitParam) HasEnum() bool {
	return o != nil && o.Enum != nil
}

// SetEnum gets a reference to the given []string and assigns it to the Enum field.
func (o *InitParam) SetEnum(v []string) {
	o.Enum = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InitParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["configSpecName"] = o.ConfigSpecName
	toSerialize["required"] = o.Required
	toSerialize["label"] = o.Label
	toSerialize["type"] = o.Type
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.Maximum != nil {
		toSerialize["maximum"] = o.Maximum
	}
	if o.Minimum != nil {
		toSerialize["minimum"] = o.Minimum
	}
	if o.Enum != nil {
		toSerialize["enum"] = o.Enum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InitParam) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string  `json:"name"`
		ConfigSpecName *string  `json:"configSpecName"`
		Required       *bool    `json:"required"`
		Label          *string  `json:"label"`
		Type           *string  `json:"type"`
		DefaultValue   *string  `json:"defaultValue,omitempty"`
		Maximum        *string  `json:"maximum,omitempty"`
		Minimum        *string  `json:"minimum,omitempty"`
		Enum           []string `json:"enum,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ConfigSpecName == nil {
		return fmt.Errorf("required field configSpecName missing")
	}
	if all.Required == nil {
		return fmt.Errorf("required field required missing")
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "configSpecName", "required", "label", "type", "defaultValue", "maximum", "minimum", "enum"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.ConfigSpecName = *all.ConfigSpecName
	o.Required = *all.Required
	o.Label = *all.Label
	o.Type = *all.Type
	o.DefaultValue = all.DefaultValue
	o.Maximum = all.Maximum
	o.Minimum = all.Minimum
	o.Enum = all.Enum

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
