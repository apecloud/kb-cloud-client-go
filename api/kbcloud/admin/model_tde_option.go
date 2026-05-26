// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TdeOption Transparent data encryption configuration.
type TdeOption struct {
	// Component type
	Component string `json:"component"`
	// Parameter configuration for querying and controlling TDE.
	Parameter TdeParameterOption `json:"parameter"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTdeOption instantiates a new TdeOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTdeOption(component string, parameter TdeParameterOption) *TdeOption {
	this := TdeOption{}
	this.Component = component
	this.Parameter = parameter
	return &this
}

// NewTdeOptionWithDefaults instantiates a new TdeOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTdeOptionWithDefaults() *TdeOption {
	this := TdeOption{}
	return &this
}

// GetComponent returns the Component field value.
func (o *TdeOption) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *TdeOption) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *TdeOption) SetComponent(v string) {
	o.Component = v
}

// GetParameter returns the Parameter field value.
func (o *TdeOption) GetParameter() TdeParameterOption {
	if o == nil {
		var ret TdeParameterOption
		return ret
	}
	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *TdeOption) GetParameterOk() (*TdeParameterOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value.
func (o *TdeOption) SetParameter(v TdeParameterOption) {
	o.Parameter = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TdeOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["parameter"] = o.Parameter

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TdeOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string             `json:"component"`
		Parameter *TdeParameterOption `json:"parameter"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Parameter == nil {
		return fmt.Errorf("required field parameter missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "parameter"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Component = *all.Component
	if all.Parameter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Parameter = *all.Parameter

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
