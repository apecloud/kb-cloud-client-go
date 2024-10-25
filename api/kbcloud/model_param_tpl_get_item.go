// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParamTplGetItem paramTplGetItem is a list of get parameter template and parameterSpec

type ParamTplGetItem struct {
	// The name of the configuration spec
	SpecName string `json:"specName"`
	// Cluster parameters configuration, include the file name and content of the parameters
	Config ConfigurationWithRegex `json:"config"`
	// With the list of parameterSpecs and the configuration file name
	ParameterSpec ParameterSpecListItem `json:"parameterSpec"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParamTplGetItem instantiates a new ParamTplGetItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParamTplGetItem(specName string, config ConfigurationWithRegex, parameterSpec ParameterSpecListItem) *ParamTplGetItem {
	this := ParamTplGetItem{}
	this.SpecName = specName
	this.Config = config
	this.ParameterSpec = parameterSpec
	return &this
}

// NewParamTplGetItemWithDefaults instantiates a new ParamTplGetItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParamTplGetItemWithDefaults() *ParamTplGetItem {
	this := ParamTplGetItem{}
	return &this
}

// GetSpecName returns the SpecName field value.
func (o *ParamTplGetItem) GetSpecName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpecName
}

// GetSpecNameOk returns a tuple with the SpecName field value
// and a boolean to check if the value has been set.
func (o *ParamTplGetItem) GetSpecNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecName, true
}

// SetSpecName sets field value.
func (o *ParamTplGetItem) SetSpecName(v string) {
	o.SpecName = v
}

// GetConfig returns the Config field value.
func (o *ParamTplGetItem) GetConfig() ConfigurationWithRegex {
	if o == nil {
		var ret ConfigurationWithRegex
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *ParamTplGetItem) GetConfigOk() (*ConfigurationWithRegex, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value.
func (o *ParamTplGetItem) SetConfig(v ConfigurationWithRegex) {
	o.Config = v
}

// GetParameterSpec returns the ParameterSpec field value.
func (o *ParamTplGetItem) GetParameterSpec() ParameterSpecListItem {
	if o == nil {
		var ret ParameterSpecListItem
		return ret
	}
	return o.ParameterSpec
}

// GetParameterSpecOk returns a tuple with the ParameterSpec field value
// and a boolean to check if the value has been set.
func (o *ParamTplGetItem) GetParameterSpecOk() (*ParameterSpecListItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParameterSpec, true
}

// SetParameterSpec sets field value.
func (o *ParamTplGetItem) SetParameterSpec(v ParameterSpecListItem) {
	o.ParameterSpec = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParamTplGetItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["specName"] = o.SpecName
	toSerialize["config"] = o.Config
	toSerialize["parameterSpec"] = o.ParameterSpec

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParamTplGetItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SpecName      *string                 `json:"specName"`
		Config        *ConfigurationWithRegex `json:"config"`
		ParameterSpec *ParameterSpecListItem  `json:"parameterSpec"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SpecName == nil {
		return fmt.Errorf("required field specName missing")
	}
	if all.Config == nil {
		return fmt.Errorf("required field config missing")
	}
	if all.ParameterSpec == nil {
		return fmt.Errorf("required field parameterSpec missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"specName", "config", "parameterSpec"})
	} else {
		return err
	}

	hasInvalidField := false
	o.SpecName = *all.SpecName
	if all.Config.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Config = *all.Config
	if all.ParameterSpec.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ParameterSpec = *all.ParameterSpec

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
