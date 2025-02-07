// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ParameterPropItem With the list of parameter properties and the configuration file name
type ParameterPropItem struct {
	// The list of parameters properties
	Props []ParameterProp `json:"props,omitempty"`
	// The name of the configuration file
	FileName *string `json:"fileName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterPropItem instantiates a new ParameterPropItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterPropItem() *ParameterPropItem {
	this := ParameterPropItem{}
	return &this
}

// NewParameterPropItemWithDefaults instantiates a new ParameterPropItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterPropItemWithDefaults() *ParameterPropItem {
	this := ParameterPropItem{}
	return &this
}

// GetProps returns the Props field value if set, zero value otherwise.
func (o *ParameterPropItem) GetProps() []ParameterProp {
	if o == nil || o.Props == nil {
		var ret []ParameterProp
		return ret
	}
	return o.Props
}

// GetPropsOk returns a tuple with the Props field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterPropItem) GetPropsOk() (*[]ParameterProp, bool) {
	if o == nil || o.Props == nil {
		return nil, false
	}
	return &o.Props, true
}

// HasProps returns a boolean if a field has been set.
func (o *ParameterPropItem) HasProps() bool {
	return o != nil && o.Props != nil
}

// SetProps gets a reference to the given []ParameterProp and assigns it to the Props field.
func (o *ParameterPropItem) SetProps(v []ParameterProp) {
	o.Props = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ParameterPropItem) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterPropItem) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ParameterPropItem) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ParameterPropItem) SetFileName(v string) {
	o.FileName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterPropItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Props != nil {
		toSerialize["props"] = o.Props
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
func (o *ParameterPropItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Props    []ParameterProp `json:"props,omitempty"`
		FileName *string         `json:"fileName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"props", "fileName"})
	} else {
		return err
	}
	o.Props = all.Props
	o.FileName = all.FileName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
