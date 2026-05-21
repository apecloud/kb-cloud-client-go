// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoShellEvaluateResponse struct {
	Result *MongoShellResult  `json:"result,omitempty"`
	Prints []MongoShellResult `json:"prints,omitempty"`
	// prompt after evaluation
	Prompt *string `json:"prompt,omitempty"`
	// whether output was truncated
	Truncated *bool `json:"truncated,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoShellEvaluateResponse instantiates a new MongoShellEvaluateResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoShellEvaluateResponse() *MongoShellEvaluateResponse {
	this := MongoShellEvaluateResponse{}
	return &this
}

// NewMongoShellEvaluateResponseWithDefaults instantiates a new MongoShellEvaluateResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoShellEvaluateResponseWithDefaults() *MongoShellEvaluateResponse {
	this := MongoShellEvaluateResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *MongoShellEvaluateResponse) GetResult() MongoShellResult {
	if o == nil || o.Result == nil {
		var ret MongoShellResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellEvaluateResponse) GetResultOk() (*MongoShellResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *MongoShellEvaluateResponse) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given MongoShellResult and assigns it to the Result field.
func (o *MongoShellEvaluateResponse) SetResult(v MongoShellResult) {
	o.Result = &v
}

// GetPrints returns the Prints field value if set, zero value otherwise.
func (o *MongoShellEvaluateResponse) GetPrints() []MongoShellResult {
	if o == nil || o.Prints == nil {
		var ret []MongoShellResult
		return ret
	}
	return o.Prints
}

// GetPrintsOk returns a tuple with the Prints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellEvaluateResponse) GetPrintsOk() (*[]MongoShellResult, bool) {
	if o == nil || o.Prints == nil {
		return nil, false
	}
	return &o.Prints, true
}

// HasPrints returns a boolean if a field has been set.
func (o *MongoShellEvaluateResponse) HasPrints() bool {
	return o != nil && o.Prints != nil
}

// SetPrints gets a reference to the given []MongoShellResult and assigns it to the Prints field.
func (o *MongoShellEvaluateResponse) SetPrints(v []MongoShellResult) {
	o.Prints = v
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *MongoShellEvaluateResponse) GetPrompt() string {
	if o == nil || o.Prompt == nil {
		var ret string
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellEvaluateResponse) GetPromptOk() (*string, bool) {
	if o == nil || o.Prompt == nil {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *MongoShellEvaluateResponse) HasPrompt() bool {
	return o != nil && o.Prompt != nil
}

// SetPrompt gets a reference to the given string and assigns it to the Prompt field.
func (o *MongoShellEvaluateResponse) SetPrompt(v string) {
	o.Prompt = &v
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *MongoShellEvaluateResponse) GetTruncated() bool {
	if o == nil || o.Truncated == nil {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellEvaluateResponse) GetTruncatedOk() (*bool, bool) {
	if o == nil || o.Truncated == nil {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *MongoShellEvaluateResponse) HasTruncated() bool {
	return o != nil && o.Truncated != nil
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *MongoShellEvaluateResponse) SetTruncated(v bool) {
	o.Truncated = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoShellEvaluateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Prints != nil {
		toSerialize["prints"] = o.Prints
	}
	if o.Prompt != nil {
		toSerialize["prompt"] = o.Prompt
	}
	if o.Truncated != nil {
		toSerialize["truncated"] = o.Truncated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoShellEvaluateResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Result    *MongoShellResult  `json:"result,omitempty"`
		Prints    []MongoShellResult `json:"prints,omitempty"`
		Prompt    *string            `json:"prompt,omitempty"`
		Truncated *bool              `json:"truncated,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"result", "prints", "prompt", "truncated"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Result != nil && all.Result.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Result = all.Result
	o.Prints = all.Prints
	o.Prompt = all.Prompt
	o.Truncated = all.Truncated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
