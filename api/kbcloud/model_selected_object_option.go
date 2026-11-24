// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type SelectedObjectOption struct {
	// Strategy for converting the selected tree to parameters. levelMapping maps every configured level and applies skipWhenParentLevelsSelectedAll. wholeObject merges duplicate branches and maps only whole selected nodes (selectedAll or leaves), omitting covered descendants. Defaults to levelMapping when omitted.
	ConversionStrategy *SelectedObjectConversionStrategy `json:"conversionStrategy,omitempty"`
	Tree               *SelectedObjectOptionTree         `json:"tree,omitempty"`
	// Parameters for specifying which objects to include in the selective backup
	ConvertToParameters []SelectedObjectOptionConvertToParametersItem `json:"convertToParameters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSelectedObjectOption instantiates a new SelectedObjectOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSelectedObjectOption() *SelectedObjectOption {
	this := SelectedObjectOption{}
	var conversionStrategy SelectedObjectConversionStrategy = SelectedObjectConversionStrategyLevelMapping
	this.ConversionStrategy = &conversionStrategy
	return &this
}

// NewSelectedObjectOptionWithDefaults instantiates a new SelectedObjectOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSelectedObjectOptionWithDefaults() *SelectedObjectOption {
	this := SelectedObjectOption{}
	var conversionStrategy SelectedObjectConversionStrategy = SelectedObjectConversionStrategyLevelMapping
	this.ConversionStrategy = &conversionStrategy
	return &this
}

// GetConversionStrategy returns the ConversionStrategy field value if set, zero value otherwise.
func (o *SelectedObjectOption) GetConversionStrategy() SelectedObjectConversionStrategy {
	if o == nil || o.ConversionStrategy == nil {
		var ret SelectedObjectConversionStrategy
		return ret
	}
	return *o.ConversionStrategy
}

// GetConversionStrategyOk returns a tuple with the ConversionStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedObjectOption) GetConversionStrategyOk() (*SelectedObjectConversionStrategy, bool) {
	if o == nil || o.ConversionStrategy == nil {
		return nil, false
	}
	return o.ConversionStrategy, true
}

// HasConversionStrategy returns a boolean if a field has been set.
func (o *SelectedObjectOption) HasConversionStrategy() bool {
	return o != nil && o.ConversionStrategy != nil
}

// SetConversionStrategy gets a reference to the given SelectedObjectConversionStrategy and assigns it to the ConversionStrategy field.
func (o *SelectedObjectOption) SetConversionStrategy(v SelectedObjectConversionStrategy) {
	o.ConversionStrategy = &v
}

// GetTree returns the Tree field value if set, zero value otherwise.
func (o *SelectedObjectOption) GetTree() SelectedObjectOptionTree {
	if o == nil || o.Tree == nil {
		var ret SelectedObjectOptionTree
		return ret
	}
	return *o.Tree
}

// GetTreeOk returns a tuple with the Tree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedObjectOption) GetTreeOk() (*SelectedObjectOptionTree, bool) {
	if o == nil || o.Tree == nil {
		return nil, false
	}
	return o.Tree, true
}

// HasTree returns a boolean if a field has been set.
func (o *SelectedObjectOption) HasTree() bool {
	return o != nil && o.Tree != nil
}

// SetTree gets a reference to the given SelectedObjectOptionTree and assigns it to the Tree field.
func (o *SelectedObjectOption) SetTree(v SelectedObjectOptionTree) {
	o.Tree = &v
}

// GetConvertToParameters returns the ConvertToParameters field value if set, zero value otherwise.
func (o *SelectedObjectOption) GetConvertToParameters() []SelectedObjectOptionConvertToParametersItem {
	if o == nil || o.ConvertToParameters == nil {
		var ret []SelectedObjectOptionConvertToParametersItem
		return ret
	}
	return o.ConvertToParameters
}

// GetConvertToParametersOk returns a tuple with the ConvertToParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedObjectOption) GetConvertToParametersOk() (*[]SelectedObjectOptionConvertToParametersItem, bool) {
	if o == nil || o.ConvertToParameters == nil {
		return nil, false
	}
	return &o.ConvertToParameters, true
}

// HasConvertToParameters returns a boolean if a field has been set.
func (o *SelectedObjectOption) HasConvertToParameters() bool {
	return o != nil && o.ConvertToParameters != nil
}

// SetConvertToParameters gets a reference to the given []SelectedObjectOptionConvertToParametersItem and assigns it to the ConvertToParameters field.
func (o *SelectedObjectOption) SetConvertToParameters(v []SelectedObjectOptionConvertToParametersItem) {
	o.ConvertToParameters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SelectedObjectOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ConversionStrategy != nil {
		toSerialize["conversionStrategy"] = o.ConversionStrategy
	}
	if o.Tree != nil {
		toSerialize["tree"] = o.Tree
	}
	if o.ConvertToParameters != nil {
		toSerialize["convertToParameters"] = o.ConvertToParameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SelectedObjectOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConversionStrategy  *SelectedObjectConversionStrategy             `json:"conversionStrategy,omitempty"`
		Tree                *SelectedObjectOptionTree                     `json:"tree,omitempty"`
		ConvertToParameters []SelectedObjectOptionConvertToParametersItem `json:"convertToParameters,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"conversionStrategy", "tree", "convertToParameters"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ConversionStrategy != nil && !all.ConversionStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.ConversionStrategy = all.ConversionStrategy
	}
	if all.Tree != nil && all.Tree.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tree = all.Tree
	o.ConvertToParameters = all.ConvertToParameters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
