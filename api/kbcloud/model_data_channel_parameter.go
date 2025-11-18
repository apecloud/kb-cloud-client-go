// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DataChannelParameter struct {
	DisplayName  *string             `json:"displayName,omitempty"`
	Desc         *InternationalDesc  `json:"desc,omitempty"`
	Key          string              `json:"key"`
	DefaultValue *string             `json:"defaultValue,omitempty"`
	ValueType    *ParameterValueType `json:"valueType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataChannelParameter instantiates a new DataChannelParameter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataChannelParameter(key string) *DataChannelParameter {
	this := DataChannelParameter{}
	this.Key = key
	return &this
}

// NewDataChannelParameterWithDefaults instantiates a new DataChannelParameter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataChannelParameterWithDefaults() *DataChannelParameter {
	this := DataChannelParameter{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DataChannelParameter) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelParameter) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DataChannelParameter) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DataChannelParameter) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *DataChannelParameter) GetDesc() InternationalDesc {
	if o == nil || o.Desc == nil {
		var ret InternationalDesc
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelParameter) GetDescOk() (*InternationalDesc, bool) {
	if o == nil || o.Desc == nil {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *DataChannelParameter) HasDesc() bool {
	return o != nil && o.Desc != nil
}

// SetDesc gets a reference to the given InternationalDesc and assigns it to the Desc field.
func (o *DataChannelParameter) SetDesc(v InternationalDesc) {
	o.Desc = &v
}

// GetKey returns the Key field value.
func (o *DataChannelParameter) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DataChannelParameter) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *DataChannelParameter) SetKey(v string) {
	o.Key = v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *DataChannelParameter) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelParameter) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *DataChannelParameter) HasDefaultValue() bool {
	return o != nil && o.DefaultValue != nil
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *DataChannelParameter) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *DataChannelParameter) GetValueType() ParameterValueType {
	if o == nil || o.ValueType == nil {
		var ret ParameterValueType
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataChannelParameter) GetValueTypeOk() (*ParameterValueType, bool) {
	if o == nil || o.ValueType == nil {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *DataChannelParameter) HasValueType() bool {
	return o != nil && o.ValueType != nil
}

// SetValueType gets a reference to the given ParameterValueType and assigns it to the ValueType field.
func (o *DataChannelParameter) SetValueType(v ParameterValueType) {
	o.ValueType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataChannelParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Desc != nil {
		toSerialize["desc"] = o.Desc
	}
	toSerialize["key"] = o.Key
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.ValueType != nil {
		toSerialize["valueType"] = o.ValueType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataChannelParameter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName  *string             `json:"displayName,omitempty"`
		Desc         *InternationalDesc  `json:"desc,omitempty"`
		Key          *string             `json:"key"`
		DefaultValue *string             `json:"defaultValue,omitempty"`
		ValueType    *ParameterValueType `json:"valueType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"displayName", "desc", "key", "defaultValue", "valueType"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DisplayName = all.DisplayName
	if all.Desc != nil && all.Desc.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Desc = all.Desc
	o.Key = *all.Key
	o.DefaultValue = all.DefaultValue
	if all.ValueType != nil && !all.ValueType.IsValid() {
		hasInvalidField = true
	} else {
		o.ValueType = all.ValueType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
