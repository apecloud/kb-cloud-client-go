// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ParameterItem With the list of parameter properties and the configuration file name
type ParameterItem struct {
	// The list of parameters properties
	Props []ParameterProp `json:"props,omitempty"`
	// The additional properties of the parameter
	AdditionalProps []ParameterProp `json:"additionalProps,omitempty"`
	// The name of the configuration file
	FileName *string `json:"fileName,omitempty"`
	// The name of the config spec
	SpecName *string `json:"specName,omitempty"`
	// The raw content of the configuration file, return only when raw is true in the query.
	RawContent *string `json:"rawContent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterItem instantiates a new ParameterItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterItem() *ParameterItem {
	this := ParameterItem{}
	return &this
}

// NewParameterItemWithDefaults instantiates a new ParameterItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterItemWithDefaults() *ParameterItem {
	this := ParameterItem{}
	return &this
}

// GetProps returns the Props field value if set, zero value otherwise.
func (o *ParameterItem) GetProps() []ParameterProp {
	if o == nil || o.Props == nil {
		var ret []ParameterProp
		return ret
	}
	return o.Props
}

// GetPropsOk returns a tuple with the Props field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterItem) GetPropsOk() (*[]ParameterProp, bool) {
	if o == nil || o.Props == nil {
		return nil, false
	}
	return &o.Props, true
}

// HasProps returns a boolean if a field has been set.
func (o *ParameterItem) HasProps() bool {
	return o != nil && o.Props != nil
}

// SetProps gets a reference to the given []ParameterProp and assigns it to the Props field.
func (o *ParameterItem) SetProps(v []ParameterProp) {
	o.Props = v
}

// GetAdditionalProps returns the AdditionalProps field value if set, zero value otherwise.
func (o *ParameterItem) GetAdditionalProps() []ParameterProp {
	if o == nil || o.AdditionalProps == nil {
		var ret []ParameterProp
		return ret
	}
	return o.AdditionalProps
}

// GetAdditionalPropsOk returns a tuple with the AdditionalProps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterItem) GetAdditionalPropsOk() (*[]ParameterProp, bool) {
	if o == nil || o.AdditionalProps == nil {
		return nil, false
	}
	return &o.AdditionalProps, true
}

// HasAdditionalProps returns a boolean if a field has been set.
func (o *ParameterItem) HasAdditionalProps() bool {
	return o != nil && o.AdditionalProps != nil
}

// SetAdditionalProps gets a reference to the given []ParameterProp and assigns it to the AdditionalProps field.
func (o *ParameterItem) SetAdditionalProps(v []ParameterProp) {
	o.AdditionalProps = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ParameterItem) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterItem) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ParameterItem) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ParameterItem) SetFileName(v string) {
	o.FileName = &v
}

// GetSpecName returns the SpecName field value if set, zero value otherwise.
func (o *ParameterItem) GetSpecName() string {
	if o == nil || o.SpecName == nil {
		var ret string
		return ret
	}
	return *o.SpecName
}

// GetSpecNameOk returns a tuple with the SpecName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterItem) GetSpecNameOk() (*string, bool) {
	if o == nil || o.SpecName == nil {
		return nil, false
	}
	return o.SpecName, true
}

// HasSpecName returns a boolean if a field has been set.
func (o *ParameterItem) HasSpecName() bool {
	return o != nil && o.SpecName != nil
}

// SetSpecName gets a reference to the given string and assigns it to the SpecName field.
func (o *ParameterItem) SetSpecName(v string) {
	o.SpecName = &v
}

// GetRawContent returns the RawContent field value if set, zero value otherwise.
func (o *ParameterItem) GetRawContent() string {
	if o == nil || o.RawContent == nil {
		var ret string
		return ret
	}
	return *o.RawContent
}

// GetRawContentOk returns a tuple with the RawContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterItem) GetRawContentOk() (*string, bool) {
	if o == nil || o.RawContent == nil {
		return nil, false
	}
	return o.RawContent, true
}

// HasRawContent returns a boolean if a field has been set.
func (o *ParameterItem) HasRawContent() bool {
	return o != nil && o.RawContent != nil
}

// SetRawContent gets a reference to the given string and assigns it to the RawContent field.
func (o *ParameterItem) SetRawContent(v string) {
	o.RawContent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Props != nil {
		toSerialize["props"] = o.Props
	}
	if o.AdditionalProps != nil {
		toSerialize["additionalProps"] = o.AdditionalProps
	}
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}
	if o.SpecName != nil {
		toSerialize["specName"] = o.SpecName
	}
	if o.RawContent != nil {
		toSerialize["rawContent"] = o.RawContent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Props           []ParameterProp `json:"props,omitempty"`
		AdditionalProps []ParameterProp `json:"additionalProps,omitempty"`
		FileName        *string         `json:"fileName,omitempty"`
		SpecName        *string         `json:"specName,omitempty"`
		RawContent      *string         `json:"rawContent,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"props", "additionalProps", "fileName", "specName", "rawContent"})
	} else {
		return err
	}
	o.Props = all.Props
	o.AdditionalProps = all.AdditionalProps
	o.FileName = all.FileName
	o.SpecName = all.SpecName
	o.RawContent = all.RawContent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
