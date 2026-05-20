// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoShellPromptResponse struct {
	// current mongosh prompt
	Prompt *string `json:"prompt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoShellPromptResponse instantiates a new MongoShellPromptResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoShellPromptResponse() *MongoShellPromptResponse {
	this := MongoShellPromptResponse{}
	return &this
}

// NewMongoShellPromptResponseWithDefaults instantiates a new MongoShellPromptResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoShellPromptResponseWithDefaults() *MongoShellPromptResponse {
	this := MongoShellPromptResponse{}
	return &this
}

// GetPrompt returns the Prompt field value if set, zero value otherwise.
func (o *MongoShellPromptResponse) GetPrompt() string {
	if o == nil || o.Prompt == nil {
		var ret string
		return ret
	}
	return *o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellPromptResponse) GetPromptOk() (*string, bool) {
	if o == nil || o.Prompt == nil {
		return nil, false
	}
	return o.Prompt, true
}

// HasPrompt returns a boolean if a field has been set.
func (o *MongoShellPromptResponse) HasPrompt() bool {
	return o != nil && o.Prompt != nil
}

// SetPrompt gets a reference to the given string and assigns it to the Prompt field.
func (o *MongoShellPromptResponse) SetPrompt(v string) {
	o.Prompt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoShellPromptResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Prompt != nil {
		toSerialize["prompt"] = o.Prompt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoShellPromptResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Prompt *string `json:"prompt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"prompt"})
	} else {
		return err
	}
	o.Prompt = all.Prompt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
