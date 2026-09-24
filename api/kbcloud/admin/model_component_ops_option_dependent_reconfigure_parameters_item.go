// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ComponentOpsOptionDependentReconfigureParametersItem struct {
	// parameter name.
	Key string `json:"key"`
	// Go template expression. Same built-in objects as dependentCustomOps params.
	// Arithmetic uses sprig. quantity parses a Kubernetes quantity string from the spec expression into bytes.
	//
	Value string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOpsOptionDependentReconfigureParametersItem instantiates a new ComponentOpsOptionDependentReconfigureParametersItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOpsOptionDependentReconfigureParametersItem(key string, value string) *ComponentOpsOptionDependentReconfigureParametersItem {
	this := ComponentOpsOptionDependentReconfigureParametersItem{}
	this.Key = key
	this.Value = value
	return &this
}

// NewComponentOpsOptionDependentReconfigureParametersItemWithDefaults instantiates a new ComponentOpsOptionDependentReconfigureParametersItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOpsOptionDependentReconfigureParametersItemWithDefaults() *ComponentOpsOptionDependentReconfigureParametersItem {
	this := ComponentOpsOptionDependentReconfigureParametersItem{}
	return &this
}

// GetKey returns the Key field value.
func (o *ComponentOpsOptionDependentReconfigureParametersItem) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentReconfigureParametersItem) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *ComponentOpsOptionDependentReconfigureParametersItem) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value.
func (o *ComponentOpsOptionDependentReconfigureParametersItem) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ComponentOpsOptionDependentReconfigureParametersItem) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ComponentOpsOptionDependentReconfigureParametersItem) SetValue(v string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOpsOptionDependentReconfigureParametersItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOpsOptionDependentReconfigureParametersItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key   *string `json:"key"`
		Value *string `json:"value"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key", "value"})
	} else {
		return err
	}
	o.Key = *all.Key
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
