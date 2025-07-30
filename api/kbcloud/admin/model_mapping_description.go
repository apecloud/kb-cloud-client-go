// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MappingDescription struct {
	ObjectExpressionTips *InternationalDesc `json:"objectExpressionTips,omitempty"`
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

// MarshalJSON serializes the struct using spec logic.
func (o MappingDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ObjectExpressionTips != nil {
		toSerialize["objectExpressionTips"] = o.ObjectExpressionTips
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
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"objectExpressionTips"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ObjectExpressionTips != nil && all.ObjectExpressionTips.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ObjectExpressionTips = all.ObjectExpressionTips

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
