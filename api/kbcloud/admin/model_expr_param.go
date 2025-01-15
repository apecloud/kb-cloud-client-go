// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ExprParam struct {
	// name of the parameter
	Name string `json:"name"`
	// description of each variables and expression
	Description string `json:"description"`
	// name of config spec, included in configSpecs, equal to componentDefinition.configs[x].name
	ConfigSpecName string `json:"configSpecName"`
	// valid variables in the expression
	ValidVariables []string `json:"validVariables"`
	// default expression of the parameter
	DefaultExpression string `json:"defaultExpression"`
	// unit of parameter
	Unit *string `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExprParam instantiates a new ExprParam object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExprParam(name string, description string, configSpecName string, validVariables []string, defaultExpression string) *ExprParam {
	this := ExprParam{}
	this.Name = name
	this.Description = description
	this.ConfigSpecName = configSpecName
	this.ValidVariables = validVariables
	this.DefaultExpression = defaultExpression
	return &this
}

// NewExprParamWithDefaults instantiates a new ExprParam object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExprParamWithDefaults() *ExprParam {
	this := ExprParam{}
	return &this
}

// GetName returns the Name field value.
func (o *ExprParam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExprParam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ExprParam) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value.
func (o *ExprParam) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ExprParam) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ExprParam) SetDescription(v string) {
	o.Description = v
}

// GetConfigSpecName returns the ConfigSpecName field value.
func (o *ExprParam) GetConfigSpecName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigSpecName
}

// GetConfigSpecNameOk returns a tuple with the ConfigSpecName field value
// and a boolean to check if the value has been set.
func (o *ExprParam) GetConfigSpecNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpecName, true
}

// SetConfigSpecName sets field value.
func (o *ExprParam) SetConfigSpecName(v string) {
	o.ConfigSpecName = v
}

// GetValidVariables returns the ValidVariables field value.
func (o *ExprParam) GetValidVariables() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ValidVariables
}

// GetValidVariablesOk returns a tuple with the ValidVariables field value
// and a boolean to check if the value has been set.
func (o *ExprParam) GetValidVariablesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidVariables, true
}

// SetValidVariables sets field value.
func (o *ExprParam) SetValidVariables(v []string) {
	o.ValidVariables = v
}

// GetDefaultExpression returns the DefaultExpression field value.
func (o *ExprParam) GetDefaultExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefaultExpression
}

// GetDefaultExpressionOk returns a tuple with the DefaultExpression field value
// and a boolean to check if the value has been set.
func (o *ExprParam) GetDefaultExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultExpression, true
}

// SetDefaultExpression sets field value.
func (o *ExprParam) SetDefaultExpression(v string) {
	o.DefaultExpression = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ExprParam) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExprParam) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ExprParam) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *ExprParam) SetUnit(v string) {
	o.Unit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExprParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["configSpecName"] = o.ConfigSpecName
	toSerialize["validVariables"] = o.ValidVariables
	toSerialize["defaultExpression"] = o.DefaultExpression
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExprParam) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string   `json:"name"`
		Description       *string   `json:"description"`
		ConfigSpecName    *string   `json:"configSpecName"`
		ValidVariables    *[]string `json:"validVariables"`
		DefaultExpression *string   `json:"defaultExpression"`
		Unit              *string   `json:"unit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.ConfigSpecName == nil {
		return fmt.Errorf("required field configSpecName missing")
	}
	if all.ValidVariables == nil {
		return fmt.Errorf("required field validVariables missing")
	}
	if all.DefaultExpression == nil {
		return fmt.Errorf("required field defaultExpression missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "description", "configSpecName", "validVariables", "defaultExpression", "unit"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Description = *all.Description
	o.ConfigSpecName = *all.ConfigSpecName
	o.ValidVariables = *all.ValidVariables
	o.DefaultExpression = *all.DefaultExpression
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
