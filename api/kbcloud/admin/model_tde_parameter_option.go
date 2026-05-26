// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TdeParameterOption Parameter configuration for querying and controlling TDE.
type TdeParameterOption struct {
	// Configuration spec name
	ConfigSpec string `json:"configSpec"`
	// Configuration file name
	ConfigFileName string `json:"configFileName"`
	// Primary parameter key used to query TDE status
	Key string `json:"key"`
	// List of parameter values that indicate TDE is enabled
	EnabledValues []string `json:"enabledValues,omitempty"`
	// List of parameter values that indicate TDE is disabled
	DisabledValues []string `json:"disabledValues,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTdeParameterOption instantiates a new TdeParameterOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTdeParameterOption(configSpec string, configFileName string, key string) *TdeParameterOption {
	this := TdeParameterOption{}
	this.ConfigSpec = configSpec
	this.ConfigFileName = configFileName
	this.Key = key
	return &this
}

// NewTdeParameterOptionWithDefaults instantiates a new TdeParameterOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTdeParameterOptionWithDefaults() *TdeParameterOption {
	this := TdeParameterOption{}
	return &this
}

// GetConfigSpec returns the ConfigSpec field value.
func (o *TdeParameterOption) GetConfigSpec() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigSpec
}

// GetConfigSpecOk returns a tuple with the ConfigSpec field value
// and a boolean to check if the value has been set.
func (o *TdeParameterOption) GetConfigSpecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpec, true
}

// SetConfigSpec sets field value.
func (o *TdeParameterOption) SetConfigSpec(v string) {
	o.ConfigSpec = v
}

// GetConfigFileName returns the ConfigFileName field value.
func (o *TdeParameterOption) GetConfigFileName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigFileName
}

// GetConfigFileNameOk returns a tuple with the ConfigFileName field value
// and a boolean to check if the value has been set.
func (o *TdeParameterOption) GetConfigFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigFileName, true
}

// SetConfigFileName sets field value.
func (o *TdeParameterOption) SetConfigFileName(v string) {
	o.ConfigFileName = v
}

// GetKey returns the Key field value.
func (o *TdeParameterOption) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TdeParameterOption) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *TdeParameterOption) SetKey(v string) {
	o.Key = v
}

// GetEnabledValues returns the EnabledValues field value if set, zero value otherwise.
func (o *TdeParameterOption) GetEnabledValues() []string {
	if o == nil || o.EnabledValues == nil {
		var ret []string
		return ret
	}
	return o.EnabledValues
}

// GetEnabledValuesOk returns a tuple with the EnabledValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TdeParameterOption) GetEnabledValuesOk() (*[]string, bool) {
	if o == nil || o.EnabledValues == nil {
		return nil, false
	}
	return &o.EnabledValues, true
}

// HasEnabledValues returns a boolean if a field has been set.
func (o *TdeParameterOption) HasEnabledValues() bool {
	return o != nil && o.EnabledValues != nil
}

// SetEnabledValues gets a reference to the given []string and assigns it to the EnabledValues field.
func (o *TdeParameterOption) SetEnabledValues(v []string) {
	o.EnabledValues = v
}

// GetDisabledValues returns the DisabledValues field value if set, zero value otherwise.
func (o *TdeParameterOption) GetDisabledValues() []string {
	if o == nil || o.DisabledValues == nil {
		var ret []string
		return ret
	}
	return o.DisabledValues
}

// GetDisabledValuesOk returns a tuple with the DisabledValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TdeParameterOption) GetDisabledValuesOk() (*[]string, bool) {
	if o == nil || o.DisabledValues == nil {
		return nil, false
	}
	return &o.DisabledValues, true
}

// HasDisabledValues returns a boolean if a field has been set.
func (o *TdeParameterOption) HasDisabledValues() bool {
	return o != nil && o.DisabledValues != nil
}

// SetDisabledValues gets a reference to the given []string and assigns it to the DisabledValues field.
func (o *TdeParameterOption) SetDisabledValues(v []string) {
	o.DisabledValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TdeParameterOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["configSpec"] = o.ConfigSpec
	toSerialize["configFileName"] = o.ConfigFileName
	toSerialize["key"] = o.Key
	if o.EnabledValues != nil {
		toSerialize["enabledValues"] = o.EnabledValues
	}
	if o.DisabledValues != nil {
		toSerialize["disabledValues"] = o.DisabledValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TdeParameterOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigSpec     *string  `json:"configSpec"`
		ConfigFileName *string  `json:"configFileName"`
		Key            *string  `json:"key"`
		EnabledValues  []string `json:"enabledValues,omitempty"`
		DisabledValues []string `json:"disabledValues,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ConfigSpec == nil {
		return fmt.Errorf("required field configSpec missing")
	}
	if all.ConfigFileName == nil {
		return fmt.Errorf("required field configFileName missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"configSpec", "configFileName", "key", "enabledValues", "disabledValues"})
	} else {
		return err
	}
	o.ConfigSpec = *all.ConfigSpec
	o.ConfigFileName = *all.ConfigFileName
	o.Key = *all.Key
	o.EnabledValues = all.EnabledValues
	o.DisabledValues = all.DisabledValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
