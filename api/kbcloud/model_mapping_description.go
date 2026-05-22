// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MappingDescription struct {
	ObjectExpressionTips *InternationalDesc `json:"objectExpressionTips,omitempty"`
	ObjectIncludeTips    *InternationalDesc `json:"objectIncludeTips,omitempty"`
	ObjectExcludeTips    *InternationalDesc `json:"objectExcludeTips,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMappingDescription instantiates a new MappingDescription object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMappingDescription() *MappingDescription {
	this := MappingDescription{}
	return &this
}

// NewMappingDescriptionWithDefaults instantiates a new MappingDescription object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMappingDescriptionWithDefaults() *MappingDescription {
	this := MappingDescription{}
	return &this
}

// GetObjectExpressionTips returns the ObjectExpressionTips field value if set, zero value otherwise.
func (o *MappingDescription) GetObjectExpressionTips() InternationalDesc {
	if o == nil || o.ObjectExpressionTips == nil {
		var ret InternationalDesc
		return ret
	}
	return *o.ObjectExpressionTips
}

// GetObjectExpressionTipsOk returns a tuple with the ObjectExpressionTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingDescription) GetObjectExpressionTipsOk() (*InternationalDesc, bool) {
	if o == nil || o.ObjectExpressionTips == nil {
		return nil, false
	}
	return o.ObjectExpressionTips, true
}

// HasObjectExpressionTips returns a boolean if a field has been set.
func (o *MappingDescription) HasObjectExpressionTips() bool {
	return o != nil && o.ObjectExpressionTips != nil
}

// SetObjectExpressionTips gets a reference to the given InternationalDesc and assigns it to the ObjectExpressionTips field.
func (o *MappingDescription) SetObjectExpressionTips(v InternationalDesc) {
	o.ObjectExpressionTips = &v
}

// GetObjectIncludeTips returns the ObjectIncludeTips field value if set, zero value otherwise.
func (o *MappingDescription) GetObjectIncludeTips() InternationalDesc {
	if o == nil || o.ObjectIncludeTips == nil {
		var ret InternationalDesc
		return ret
	}
	return *o.ObjectIncludeTips
}

// GetObjectIncludeTipsOk returns a tuple with the ObjectIncludeTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingDescription) GetObjectIncludeTipsOk() (*InternationalDesc, bool) {
	if o == nil || o.ObjectIncludeTips == nil {
		return nil, false
	}
	return o.ObjectIncludeTips, true
}

// HasObjectIncludeTips returns a boolean if a field has been set.
func (o *MappingDescription) HasObjectIncludeTips() bool {
	return o != nil && o.ObjectIncludeTips != nil
}

// SetObjectIncludeTips gets a reference to the given InternationalDesc and assigns it to the ObjectIncludeTips field.
func (o *MappingDescription) SetObjectIncludeTips(v InternationalDesc) {
	o.ObjectIncludeTips = &v
}

// GetObjectExcludeTips returns the ObjectExcludeTips field value if set, zero value otherwise.
func (o *MappingDescription) GetObjectExcludeTips() InternationalDesc {
	if o == nil || o.ObjectExcludeTips == nil {
		var ret InternationalDesc
		return ret
	}
	return *o.ObjectExcludeTips
}

// GetObjectExcludeTipsOk returns a tuple with the ObjectExcludeTips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MappingDescription) GetObjectExcludeTipsOk() (*InternationalDesc, bool) {
	if o == nil || o.ObjectExcludeTips == nil {
		return nil, false
	}
	return o.ObjectExcludeTips, true
}

// HasObjectExcludeTips returns a boolean if a field has been set.
func (o *MappingDescription) HasObjectExcludeTips() bool {
	return o != nil && o.ObjectExcludeTips != nil
}

// SetObjectExcludeTips gets a reference to the given InternationalDesc and assigns it to the ObjectExcludeTips field.
func (o *MappingDescription) SetObjectExcludeTips(v InternationalDesc) {
	o.ObjectExcludeTips = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MappingDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ObjectExpressionTips != nil {
		toSerialize["objectExpressionTips"] = o.ObjectExpressionTips
	}
	if o.ObjectIncludeTips != nil {
		toSerialize["objectIncludeTips"] = o.ObjectIncludeTips
	}
	if o.ObjectExcludeTips != nil {
		toSerialize["objectExcludeTips"] = o.ObjectExcludeTips
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MappingDescription) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ObjectExpressionTips *InternationalDesc `json:"objectExpressionTips,omitempty"`
		ObjectIncludeTips    *InternationalDesc `json:"objectIncludeTips,omitempty"`
		ObjectExcludeTips    *InternationalDesc `json:"objectExcludeTips,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"objectExpressionTips", "objectIncludeTips", "objectExcludeTips"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ObjectExpressionTips != nil && all.ObjectExpressionTips.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ObjectExpressionTips = all.ObjectExpressionTips
	if all.ObjectIncludeTips != nil && all.ObjectIncludeTips.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ObjectIncludeTips = all.ObjectIncludeTips
	if all.ObjectExcludeTips != nil && all.ObjectExcludeTips.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ObjectExcludeTips = all.ObjectExcludeTips

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
