// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

type AnalysisResult struct {
	Explanation *string `json:"explanation,omitempty"`
	Error       *string `json:"error,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnalysisResult instantiates a new AnalysisResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnalysisResult() *AnalysisResult {
	this := AnalysisResult{}
	return &this
}

// NewAnalysisResultWithDefaults instantiates a new AnalysisResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnalysisResultWithDefaults() *AnalysisResult {
	this := AnalysisResult{}
	return &this
}

// GetExplanation returns the Explanation field value if set, zero value otherwise.
func (o *AnalysisResult) GetExplanation() string {
	if o == nil || o.Explanation == nil {
		var ret string
		return ret
	}
	return *o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisResult) GetExplanationOk() (*string, bool) {
	if o == nil || o.Explanation == nil {
		return nil, false
	}
	return o.Explanation, true
}

// HasExplanation returns a boolean if a field has been set.
func (o *AnalysisResult) HasExplanation() bool {
	return o != nil && o.Explanation != nil
}

// SetExplanation gets a reference to the given string and assigns it to the Explanation field.
func (o *AnalysisResult) SetExplanation(v string) {
	o.Explanation = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AnalysisResult) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalysisResult) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AnalysisResult) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *AnalysisResult) SetError(v string) {
	o.Error = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnalysisResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Explanation != nil {
		toSerialize["explanation"] = o.Explanation
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnalysisResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Explanation *string `json:"explanation,omitempty"`
		Error       *string `json:"error,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"explanation", "error"})
	} else {
		return err
	}
	o.Explanation = all.Explanation
	o.Error = all.Error

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
