// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OccupationParam struct {
	// name of the parameter
	Name string `json:"name"`
	// expression of the parameter value
	Expression string `json:"expression"`
	// name of config spec, included in configSpecs, equal to componentDefinition.configs[x].name
	ConfigSpecName string `json:"configSpecName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOccupationParam instantiates a new OccupationParam object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOccupationParam(name string, expression string, configSpecName string) *OccupationParam {
	this := OccupationParam{}
	this.Name = name
	this.Expression = expression
	this.ConfigSpecName = configSpecName
	return &this
}

// NewOccupationParamWithDefaults instantiates a new OccupationParam object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOccupationParamWithDefaults() *OccupationParam {
	this := OccupationParam{}
	return &this
}

// GetName returns the Name field value.
func (o *OccupationParam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OccupationParam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OccupationParam) SetName(v string) {
	o.Name = v
}

// GetExpression returns the Expression field value.
func (o *OccupationParam) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *OccupationParam) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *OccupationParam) SetExpression(v string) {
	o.Expression = v
}

// GetConfigSpecName returns the ConfigSpecName field value.
func (o *OccupationParam) GetConfigSpecName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigSpecName
}

// GetConfigSpecNameOk returns a tuple with the ConfigSpecName field value
// and a boolean to check if the value has been set.
func (o *OccupationParam) GetConfigSpecNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpecName, true
}

// SetConfigSpecName sets field value.
func (o *OccupationParam) SetConfigSpecName(v string) {
	o.ConfigSpecName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OccupationParam) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["expression"] = o.Expression
	toSerialize["configSpecName"] = o.ConfigSpecName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OccupationParam) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string `json:"name"`
		Expression     *string `json:"expression"`
		ConfigSpecName *string `json:"configSpecName"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	if all.ConfigSpecName == nil {
		return fmt.Errorf("required field configSpecName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "expression", "configSpecName"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Expression = *all.Expression
	o.ConfigSpecName = *all.ConfigSpecName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
