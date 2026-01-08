// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ModeOptionValuesMappingsMappingsItem struct {
	Component *string `json:"component,omitempty"`
	SourceKey string  `json:"sourceKey"`
	// Can be any type (string, number, boolean, object, etc.)
	SourceValue map[string]interface{} `json:"sourceValue,omitempty"`
	TargetKey   string                 `json:"targetKey"`
	// Can be any type (string, number, boolean, object, etc.)
	TargetValue map[string]interface{} `json:"targetValue,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeOptionValuesMappingsMappingsItem instantiates a new ModeOptionValuesMappingsMappingsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeOptionValuesMappingsMappingsItem(sourceKey string, targetKey string) *ModeOptionValuesMappingsMappingsItem {
	this := ModeOptionValuesMappingsMappingsItem{}
	this.SourceKey = sourceKey
	this.TargetKey = targetKey
	return &this
}

// NewModeOptionValuesMappingsMappingsItemWithDefaults instantiates a new ModeOptionValuesMappingsMappingsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeOptionValuesMappingsMappingsItemWithDefaults() *ModeOptionValuesMappingsMappingsItem {
	this := ModeOptionValuesMappingsMappingsItem{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ModeOptionValuesMappingsMappingsItem) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOptionValuesMappingsMappingsItem) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ModeOptionValuesMappingsMappingsItem) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *ModeOptionValuesMappingsMappingsItem) SetComponent(v string) {
	o.Component = &v
}

// GetSourceKey returns the SourceKey field value.
func (o *ModeOptionValuesMappingsMappingsItem) GetSourceKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceKey
}

// GetSourceKeyOk returns a tuple with the SourceKey field value
// and a boolean to check if the value has been set.
func (o *ModeOptionValuesMappingsMappingsItem) GetSourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceKey, true
}

// SetSourceKey sets field value.
func (o *ModeOptionValuesMappingsMappingsItem) SetSourceKey(v string) {
	o.SourceKey = v
}

// GetSourceValue returns the SourceValue field value if set, zero value otherwise.
func (o *ModeOptionValuesMappingsMappingsItem) GetSourceValue() map[string]interface{} {
	if o == nil || o.SourceValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.SourceValue
}

// GetSourceValueOk returns a tuple with the SourceValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOptionValuesMappingsMappingsItem) GetSourceValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.SourceValue == nil {
		return nil, false
	}
	return &o.SourceValue, true
}

// HasSourceValue returns a boolean if a field has been set.
func (o *ModeOptionValuesMappingsMappingsItem) HasSourceValue() bool {
	return o != nil && o.SourceValue != nil
}

// SetSourceValue gets a reference to the given map[string]interface{} and assigns it to the SourceValue field.
func (o *ModeOptionValuesMappingsMappingsItem) SetSourceValue(v map[string]interface{}) {
	o.SourceValue = v
}

// GetTargetKey returns the TargetKey field value.
func (o *ModeOptionValuesMappingsMappingsItem) GetTargetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TargetKey
}

// GetTargetKeyOk returns a tuple with the TargetKey field value
// and a boolean to check if the value has been set.
func (o *ModeOptionValuesMappingsMappingsItem) GetTargetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetKey, true
}

// SetTargetKey sets field value.
func (o *ModeOptionValuesMappingsMappingsItem) SetTargetKey(v string) {
	o.TargetKey = v
}

// GetTargetValue returns the TargetValue field value if set, zero value otherwise.
func (o *ModeOptionValuesMappingsMappingsItem) GetTargetValue() map[string]interface{} {
	if o == nil || o.TargetValue == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TargetValue
}

// GetTargetValueOk returns a tuple with the TargetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOptionValuesMappingsMappingsItem) GetTargetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.TargetValue == nil {
		return nil, false
	}
	return &o.TargetValue, true
}

// HasTargetValue returns a boolean if a field has been set.
func (o *ModeOptionValuesMappingsMappingsItem) HasTargetValue() bool {
	return o != nil && o.TargetValue != nil
}

// SetTargetValue gets a reference to the given map[string]interface{} and assigns it to the TargetValue field.
func (o *ModeOptionValuesMappingsMappingsItem) SetTargetValue(v map[string]interface{}) {
	o.TargetValue = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeOptionValuesMappingsMappingsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	toSerialize["sourceKey"] = o.SourceKey
	if o.SourceValue != nil {
		toSerialize["sourceValue"] = o.SourceValue
	}
	toSerialize["targetKey"] = o.TargetKey
	if o.TargetValue != nil {
		toSerialize["targetValue"] = o.TargetValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeOptionValuesMappingsMappingsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component   *string                `json:"component,omitempty"`
		SourceKey   *string                `json:"sourceKey"`
		SourceValue map[string]interface{} `json:"sourceValue,omitempty"`
		TargetKey   *string                `json:"targetKey"`
		TargetValue map[string]interface{} `json:"targetValue,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SourceKey == nil {
		return fmt.Errorf("required field sourceKey missing")
	}
	if all.TargetKey == nil {
		return fmt.Errorf("required field targetKey missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "sourceKey", "sourceValue", "targetKey", "targetValue"})
	} else {
		return err
	}
	o.Component = all.Component
	o.SourceKey = *all.SourceKey
	o.SourceValue = all.SourceValue
	o.TargetKey = *all.TargetKey
	o.TargetValue = all.TargetValue

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
