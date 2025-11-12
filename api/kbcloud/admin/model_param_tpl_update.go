// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ParamTplUpdate paramTplUpdate is the payload to update a parameter template
type ParamTplUpdate struct {
	// Specify parameters belongs to which spec. Currently, only one configuration file corresponds to a specName, so the configuration file name is omitted.
	SpecName *string `json:"specName,omitempty"`
	// Specify parameters list to be updated
	Parameters map[string]string `json:"parameters,omitempty"`
	// Specify the new name of the parameter template
	NewParamTplName *string `json:"newParamTplName,omitempty"`
	// The raw content of the configuration file
	RawContent *string `json:"rawContent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParamTplUpdate instantiates a new ParamTplUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParamTplUpdate() *ParamTplUpdate {
	this := ParamTplUpdate{}
	return &this
}

// NewParamTplUpdateWithDefaults instantiates a new ParamTplUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParamTplUpdateWithDefaults() *ParamTplUpdate {
	this := ParamTplUpdate{}
	return &this
}

// GetSpecName returns the SpecName field value if set, zero value otherwise.
func (o *ParamTplUpdate) GetSpecName() string {
	if o == nil || o.SpecName == nil {
		var ret string
		return ret
	}
	return *o.SpecName
}

// GetSpecNameOk returns a tuple with the SpecName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplUpdate) GetSpecNameOk() (*string, bool) {
	if o == nil || o.SpecName == nil {
		return nil, false
	}
	return o.SpecName, true
}

// HasSpecName returns a boolean if a field has been set.
func (o *ParamTplUpdate) HasSpecName() bool {
	return o != nil && o.SpecName != nil
}

// SetSpecName gets a reference to the given string and assigns it to the SpecName field.
func (o *ParamTplUpdate) SetSpecName(v string) {
	o.SpecName = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ParamTplUpdate) GetParameters() map[string]string {
	if o == nil || o.Parameters == nil {
		var ret map[string]string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplUpdate) GetParametersOk() (*map[string]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ParamTplUpdate) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *ParamTplUpdate) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetNewParamTplName returns the NewParamTplName field value if set, zero value otherwise.
func (o *ParamTplUpdate) GetNewParamTplName() string {
	if o == nil || o.NewParamTplName == nil {
		var ret string
		return ret
	}
	return *o.NewParamTplName
}

// GetNewParamTplNameOk returns a tuple with the NewParamTplName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplUpdate) GetNewParamTplNameOk() (*string, bool) {
	if o == nil || o.NewParamTplName == nil {
		return nil, false
	}
	return o.NewParamTplName, true
}

// HasNewParamTplName returns a boolean if a field has been set.
func (o *ParamTplUpdate) HasNewParamTplName() bool {
	return o != nil && o.NewParamTplName != nil
}

// SetNewParamTplName gets a reference to the given string and assigns it to the NewParamTplName field.
func (o *ParamTplUpdate) SetNewParamTplName(v string) {
	o.NewParamTplName = &v
}

// GetRawContent returns the RawContent field value if set, zero value otherwise.
func (o *ParamTplUpdate) GetRawContent() string {
	if o == nil || o.RawContent == nil {
		var ret string
		return ret
	}
	return *o.RawContent
}

// GetRawContentOk returns a tuple with the RawContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParamTplUpdate) GetRawContentOk() (*string, bool) {
	if o == nil || o.RawContent == nil {
		return nil, false
	}
	return o.RawContent, true
}

// HasRawContent returns a boolean if a field has been set.
func (o *ParamTplUpdate) HasRawContent() bool {
	return o != nil && o.RawContent != nil
}

// SetRawContent gets a reference to the given string and assigns it to the RawContent field.
func (o *ParamTplUpdate) SetRawContent(v string) {
	o.RawContent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParamTplUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SpecName != nil {
		toSerialize["specName"] = o.SpecName
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.NewParamTplName != nil {
		toSerialize["newParamTplName"] = o.NewParamTplName
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
func (o *ParamTplUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SpecName        *string           `json:"specName,omitempty"`
		Parameters      map[string]string `json:"parameters,omitempty"`
		NewParamTplName *string           `json:"newParamTplName,omitempty"`
		RawContent      *string           `json:"rawContent,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"specName", "parameters", "newParamTplName", "rawContent"})
	} else {
		return err
	}
	o.SpecName = all.SpecName
	o.Parameters = all.Parameters
	o.NewParamTplName = all.NewParamTplName
	o.RawContent = all.RawContent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
