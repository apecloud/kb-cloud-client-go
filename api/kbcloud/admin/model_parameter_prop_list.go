// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ParameterPropList With the list of parameters and the configuration file name
type ParameterPropList struct {
	// The list of parameters properties
	Specs []ParameterPropItem `json:"specs,omitempty"`
	// The name of the configuration file
	FileName *string `json:"fileName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterPropList instantiates a new ParameterPropList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterPropList() *ParameterPropList {
	this := ParameterPropList{}
	return &this
}

// NewParameterPropListWithDefaults instantiates a new ParameterPropList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterPropListWithDefaults() *ParameterPropList {
	this := ParameterPropList{}
	return &this
}

// GetSpecs returns the Specs field value if set, zero value otherwise.
func (o *ParameterPropList) GetSpecs() []ParameterPropItem {
	if o == nil || o.Specs == nil {
		var ret []ParameterPropItem
		return ret
	}
	return o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterPropList) GetSpecsOk() (*[]ParameterPropItem, bool) {
	if o == nil || o.Specs == nil {
		return nil, false
	}
	return &o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *ParameterPropList) HasSpecs() bool {
	return o != nil && o.Specs != nil
}

// SetSpecs gets a reference to the given []ParameterPropItem and assigns it to the Specs field.
func (o *ParameterPropList) SetSpecs(v []ParameterPropItem) {
	o.Specs = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ParameterPropList) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterPropList) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ParameterPropList) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ParameterPropList) SetFileName(v string) {
	o.FileName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterPropList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Specs != nil {
		toSerialize["specs"] = o.Specs
	}
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterPropList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Specs    []ParameterPropItem `json:"specs,omitempty"`
		FileName *string             `json:"fileName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"specs", "fileName"})
	} else {
		return err
	}
	o.Specs = all.Specs
	o.FileName = all.FileName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
