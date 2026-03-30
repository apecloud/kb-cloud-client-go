// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SelectedObjectOptionConvertToParametersItem struct {
	// The name of the parameter for selective backup, e.g., "tables"
	ParameterName string `json:"parameterName"`
	// The format of the value with the tree level index, e.g., "{0}", "{0}.{1}"
	ValueTemplate string `json:"valueTemplate"`
	// Whether to skip the current level when all of the parent level is selected all
	SkipWhenParentLevelsSelectedAll *bool `json:"skipWhenParentLevelsSelectedAll,omitempty"`
	// The index of the current level in the tree level order
	LevelIndex *int32 `json:"levelIndex,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSelectedObjectOptionConvertToParametersItem instantiates a new SelectedObjectOptionConvertToParametersItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSelectedObjectOptionConvertToParametersItem(parameterName string, valueTemplate string) *SelectedObjectOptionConvertToParametersItem {
	this := SelectedObjectOptionConvertToParametersItem{}
	this.ParameterName = parameterName
	this.ValueTemplate = valueTemplate
	return &this
}

// NewSelectedObjectOptionConvertToParametersItemWithDefaults instantiates a new SelectedObjectOptionConvertToParametersItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSelectedObjectOptionConvertToParametersItemWithDefaults() *SelectedObjectOptionConvertToParametersItem {
	this := SelectedObjectOptionConvertToParametersItem{}
	return &this
}

// GetParameterName returns the ParameterName field value.
func (o *SelectedObjectOptionConvertToParametersItem) GetParameterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value
// and a boolean to check if the value has been set.
func (o *SelectedObjectOptionConvertToParametersItem) GetParameterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParameterName, true
}

// SetParameterName sets field value.
func (o *SelectedObjectOptionConvertToParametersItem) SetParameterName(v string) {
	o.ParameterName = v
}

// GetValueTemplate returns the ValueTemplate field value.
func (o *SelectedObjectOptionConvertToParametersItem) GetValueTemplate() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ValueTemplate
}

// GetValueTemplateOk returns a tuple with the ValueTemplate field value
// and a boolean to check if the value has been set.
func (o *SelectedObjectOptionConvertToParametersItem) GetValueTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueTemplate, true
}

// SetValueTemplate sets field value.
func (o *SelectedObjectOptionConvertToParametersItem) SetValueTemplate(v string) {
	o.ValueTemplate = v
}

// GetSkipWhenParentLevelsSelectedAll returns the SkipWhenParentLevelsSelectedAll field value if set, zero value otherwise.
func (o *SelectedObjectOptionConvertToParametersItem) GetSkipWhenParentLevelsSelectedAll() bool {
	if o == nil || o.SkipWhenParentLevelsSelectedAll == nil {
		var ret bool
		return ret
	}
	return *o.SkipWhenParentLevelsSelectedAll
}

// GetSkipWhenParentLevelsSelectedAllOk returns a tuple with the SkipWhenParentLevelsSelectedAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedObjectOptionConvertToParametersItem) GetSkipWhenParentLevelsSelectedAllOk() (*bool, bool) {
	if o == nil || o.SkipWhenParentLevelsSelectedAll == nil {
		return nil, false
	}
	return o.SkipWhenParentLevelsSelectedAll, true
}

// HasSkipWhenParentLevelsSelectedAll returns a boolean if a field has been set.
func (o *SelectedObjectOptionConvertToParametersItem) HasSkipWhenParentLevelsSelectedAll() bool {
	return o != nil && o.SkipWhenParentLevelsSelectedAll != nil
}

// SetSkipWhenParentLevelsSelectedAll gets a reference to the given bool and assigns it to the SkipWhenParentLevelsSelectedAll field.
func (o *SelectedObjectOptionConvertToParametersItem) SetSkipWhenParentLevelsSelectedAll(v bool) {
	o.SkipWhenParentLevelsSelectedAll = &v
}

// GetLevelIndex returns the LevelIndex field value if set, zero value otherwise.
func (o *SelectedObjectOptionConvertToParametersItem) GetLevelIndex() int32 {
	if o == nil || o.LevelIndex == nil {
		var ret int32
		return ret
	}
	return *o.LevelIndex
}

// GetLevelIndexOk returns a tuple with the LevelIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedObjectOptionConvertToParametersItem) GetLevelIndexOk() (*int32, bool) {
	if o == nil || o.LevelIndex == nil {
		return nil, false
	}
	return o.LevelIndex, true
}

// HasLevelIndex returns a boolean if a field has been set.
func (o *SelectedObjectOptionConvertToParametersItem) HasLevelIndex() bool {
	return o != nil && o.LevelIndex != nil
}

// SetLevelIndex gets a reference to the given int32 and assigns it to the LevelIndex field.
func (o *SelectedObjectOptionConvertToParametersItem) SetLevelIndex(v int32) {
	o.LevelIndex = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SelectedObjectOptionConvertToParametersItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["parameterName"] = o.ParameterName
	toSerialize["valueTemplate"] = o.ValueTemplate
	if o.SkipWhenParentLevelsSelectedAll != nil {
		toSerialize["skipWhenParentLevelsSelectedAll"] = o.SkipWhenParentLevelsSelectedAll
	}
	if o.LevelIndex != nil {
		toSerialize["levelIndex"] = o.LevelIndex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SelectedObjectOptionConvertToParametersItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ParameterName                   *string `json:"parameterName"`
		ValueTemplate                   *string `json:"valueTemplate"`
		SkipWhenParentLevelsSelectedAll *bool   `json:"skipWhenParentLevelsSelectedAll,omitempty"`
		LevelIndex                      *int32  `json:"levelIndex,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ParameterName == nil {
		return fmt.Errorf("required field parameterName missing")
	}
	if all.ValueTemplate == nil {
		return fmt.Errorf("required field valueTemplate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"parameterName", "valueTemplate", "skipWhenParentLevelsSelectedAll", "levelIndex"})
	} else {
		return err
	}
	o.ParameterName = *all.ParameterName
	o.ValueTemplate = *all.ValueTemplate
	o.SkipWhenParentLevelsSelectedAll = all.SkipWhenParentLevelsSelectedAll
	o.LevelIndex = all.LevelIndex

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
