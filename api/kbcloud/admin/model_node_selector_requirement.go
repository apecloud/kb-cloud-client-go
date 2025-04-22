// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NodeSelectorRequirement Defines a requirement for a node selector.
type NodeSelectorRequirement struct {
	Key      string               `json:"key"`
	Operator NodeSelectorOperator `json:"operator"`
	Values   []string             `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNodeSelectorRequirement instantiates a new NodeSelectorRequirement object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNodeSelectorRequirement(key string, operator NodeSelectorOperator) *NodeSelectorRequirement {
	this := NodeSelectorRequirement{}
	this.Key = key
	this.Operator = operator
	return &this
}

// NewNodeSelectorRequirementWithDefaults instantiates a new NodeSelectorRequirement object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodeSelectorRequirementWithDefaults() *NodeSelectorRequirement {
	this := NodeSelectorRequirement{}
	return &this
}

// GetKey returns the Key field value.
func (o *NodeSelectorRequirement) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *NodeSelectorRequirement) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *NodeSelectorRequirement) SetKey(v string) {
	o.Key = v
}

// GetOperator returns the Operator field value.
func (o *NodeSelectorRequirement) GetOperator() NodeSelectorOperator {
	if o == nil {
		var ret NodeSelectorOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *NodeSelectorRequirement) GetOperatorOk() (*NodeSelectorOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *NodeSelectorRequirement) SetOperator(v NodeSelectorOperator) {
	o.Operator = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *NodeSelectorRequirement) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeSelectorRequirement) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *NodeSelectorRequirement) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *NodeSelectorRequirement) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o NodeSelectorRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	toSerialize["operator"] = o.Operator
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NodeSelectorRequirement) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key      *string               `json:"key"`
		Operator *NodeSelectorOperator `json:"operator"`
		Values   []string              `json:"values,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key", "operator", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Key = *all.Key
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
