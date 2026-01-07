// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SqlAuditParameterOption Parameter configuration for querying and controlling (required when type is parameter). When type is parameter, both query and control (reconfigure) are inferred from this configuration.
type SqlAuditParameterOption struct {
	// Configuration spec name
	ConfigSpec string `json:"configSpec"`
	// Configuration file name
	ConfigFileName string `json:"configFileName"`
	// Parameter key to query
	Key string `json:"key"`
	// List of parameter values that indicate SQL audit is enabled (case-insensitive comparison)
	EnabledValues []string `json:"enabledValues,omitempty"`
	// List of parameter values that indicate SQL audit is disabled (case-insensitive comparison)
	DisabledValues []string `json:"disabledValues,omitempty"`
	// If true, process parameter value from left to right, supporting subtraction (e.g., pgaudit.log with -read). Each part modifies the current state. Final state determines if audit is enabled. Default is false.
	ProcessOrder *bool `json:"processOrder,omitempty"`
	// Special value that enables all classes when processOrder is true (e.g., "all" for pgaudit.log). If this value is encountered, it enables all audit classes. Optional, only used when processOrder is true.
	AllValue *string `json:"allValue,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSqlAuditParameterOption instantiates a new SqlAuditParameterOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSqlAuditParameterOption(configSpec string, configFileName string, key string) *SqlAuditParameterOption {
	this := SqlAuditParameterOption{}
	this.ConfigSpec = configSpec
	this.ConfigFileName = configFileName
	this.Key = key
	var processOrder bool = false
	this.ProcessOrder = &processOrder
	return &this
}

// NewSqlAuditParameterOptionWithDefaults instantiates a new SqlAuditParameterOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSqlAuditParameterOptionWithDefaults() *SqlAuditParameterOption {
	this := SqlAuditParameterOption{}
	var processOrder bool = false
	this.ProcessOrder = &processOrder
	return &this
}

// GetConfigSpec returns the ConfigSpec field value.
func (o *SqlAuditParameterOption) GetConfigSpec() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigSpec
}

// GetConfigSpecOk returns a tuple with the ConfigSpec field value
// and a boolean to check if the value has been set.
func (o *SqlAuditParameterOption) GetConfigSpecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigSpec, true
}

// SetConfigSpec sets field value.
func (o *SqlAuditParameterOption) SetConfigSpec(v string) {
	o.ConfigSpec = v
}

// GetConfigFileName returns the ConfigFileName field value.
func (o *SqlAuditParameterOption) GetConfigFileName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigFileName
}

// GetConfigFileNameOk returns a tuple with the ConfigFileName field value
// and a boolean to check if the value has been set.
func (o *SqlAuditParameterOption) GetConfigFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigFileName, true
}

// SetConfigFileName sets field value.
func (o *SqlAuditParameterOption) SetConfigFileName(v string) {
	o.ConfigFileName = v
}

// GetKey returns the Key field value.
func (o *SqlAuditParameterOption) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SqlAuditParameterOption) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *SqlAuditParameterOption) SetKey(v string) {
	o.Key = v
}

// GetEnabledValues returns the EnabledValues field value if set, zero value otherwise.
func (o *SqlAuditParameterOption) GetEnabledValues() []string {
	if o == nil || o.EnabledValues == nil {
		var ret []string
		return ret
	}
	return o.EnabledValues
}

// GetEnabledValuesOk returns a tuple with the EnabledValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditParameterOption) GetEnabledValuesOk() (*[]string, bool) {
	if o == nil || o.EnabledValues == nil {
		return nil, false
	}
	return &o.EnabledValues, true
}

// HasEnabledValues returns a boolean if a field has been set.
func (o *SqlAuditParameterOption) HasEnabledValues() bool {
	return o != nil && o.EnabledValues != nil
}

// SetEnabledValues gets a reference to the given []string and assigns it to the EnabledValues field.
func (o *SqlAuditParameterOption) SetEnabledValues(v []string) {
	o.EnabledValues = v
}

// GetDisabledValues returns the DisabledValues field value if set, zero value otherwise.
func (o *SqlAuditParameterOption) GetDisabledValues() []string {
	if o == nil || o.DisabledValues == nil {
		var ret []string
		return ret
	}
	return o.DisabledValues
}

// GetDisabledValuesOk returns a tuple with the DisabledValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditParameterOption) GetDisabledValuesOk() (*[]string, bool) {
	if o == nil || o.DisabledValues == nil {
		return nil, false
	}
	return &o.DisabledValues, true
}

// HasDisabledValues returns a boolean if a field has been set.
func (o *SqlAuditParameterOption) HasDisabledValues() bool {
	return o != nil && o.DisabledValues != nil
}

// SetDisabledValues gets a reference to the given []string and assigns it to the DisabledValues field.
func (o *SqlAuditParameterOption) SetDisabledValues(v []string) {
	o.DisabledValues = v
}

// GetProcessOrder returns the ProcessOrder field value if set, zero value otherwise.
func (o *SqlAuditParameterOption) GetProcessOrder() bool {
	if o == nil || o.ProcessOrder == nil {
		var ret bool
		return ret
	}
	return *o.ProcessOrder
}

// GetProcessOrderOk returns a tuple with the ProcessOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditParameterOption) GetProcessOrderOk() (*bool, bool) {
	if o == nil || o.ProcessOrder == nil {
		return nil, false
	}
	return o.ProcessOrder, true
}

// HasProcessOrder returns a boolean if a field has been set.
func (o *SqlAuditParameterOption) HasProcessOrder() bool {
	return o != nil && o.ProcessOrder != nil
}

// SetProcessOrder gets a reference to the given bool and assigns it to the ProcessOrder field.
func (o *SqlAuditParameterOption) SetProcessOrder(v bool) {
	o.ProcessOrder = &v
}

// GetAllValue returns the AllValue field value if set, zero value otherwise.
func (o *SqlAuditParameterOption) GetAllValue() string {
	if o == nil || o.AllValue == nil {
		var ret string
		return ret
	}
	return *o.AllValue
}

// GetAllValueOk returns a tuple with the AllValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlAuditParameterOption) GetAllValueOk() (*string, bool) {
	if o == nil || o.AllValue == nil {
		return nil, false
	}
	return o.AllValue, true
}

// HasAllValue returns a boolean if a field has been set.
func (o *SqlAuditParameterOption) HasAllValue() bool {
	return o != nil && o.AllValue != nil
}

// SetAllValue gets a reference to the given string and assigns it to the AllValue field.
func (o *SqlAuditParameterOption) SetAllValue(v string) {
	o.AllValue = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SqlAuditParameterOption) MarshalJSON() ([]byte, error) {
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
	if o.ProcessOrder != nil {
		toSerialize["processOrder"] = o.ProcessOrder
	}
	if o.AllValue != nil {
		toSerialize["allValue"] = o.AllValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SqlAuditParameterOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigSpec     *string  `json:"configSpec"`
		ConfigFileName *string  `json:"configFileName"`
		Key            *string  `json:"key"`
		EnabledValues  []string `json:"enabledValues,omitempty"`
		DisabledValues []string `json:"disabledValues,omitempty"`
		ProcessOrder   *bool    `json:"processOrder,omitempty"`
		AllValue       *string  `json:"allValue,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"configSpec", "configFileName", "key", "enabledValues", "disabledValues", "processOrder", "allValue"})
	} else {
		return err
	}
	o.ConfigSpec = *all.ConfigSpec
	o.ConfigFileName = *all.ConfigFileName
	o.Key = *all.Key
	o.EnabledValues = all.EnabledValues
	o.DisabledValues = all.DisabledValues
	o.ProcessOrder = all.ProcessOrder
	o.AllValue = all.AllValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
