// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentRootCause struct {
	Type       *AiAgentRootCauseType `json:"type,omitempty"`
	Text       *string               `json:"text,omitempty"`
	Confidence *AiAgentConfidence    `json:"confidence,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentRootCause instantiates a new AiAgentRootCause object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentRootCause() *AiAgentRootCause {
	this := AiAgentRootCause{}
	return &this
}

// NewAiAgentRootCauseWithDefaults instantiates a new AiAgentRootCause object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentRootCauseWithDefaults() *AiAgentRootCause {
	this := AiAgentRootCause{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AiAgentRootCause) GetType() AiAgentRootCauseType {
	if o == nil || o.Type == nil {
		var ret AiAgentRootCauseType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRootCause) GetTypeOk() (*AiAgentRootCauseType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AiAgentRootCause) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given AiAgentRootCauseType and assigns it to the Type field.
func (o *AiAgentRootCause) SetType(v AiAgentRootCauseType) {
	o.Type = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AiAgentRootCause) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRootCause) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AiAgentRootCause) HasText() bool {
	return o != nil && o.Text != nil
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AiAgentRootCause) SetText(v string) {
	o.Text = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *AiAgentRootCause) GetConfidence() AiAgentConfidence {
	if o == nil || o.Confidence == nil {
		var ret AiAgentConfidence
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRootCause) GetConfidenceOk() (*AiAgentConfidence, bool) {
	if o == nil || o.Confidence == nil {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *AiAgentRootCause) HasConfidence() bool {
	return o != nil && o.Confidence != nil
}

// SetConfidence gets a reference to the given AiAgentConfidence and assigns it to the Confidence field.
func (o *AiAgentRootCause) SetConfidence(v AiAgentConfidence) {
	o.Confidence = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentRootCause) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Confidence != nil {
		toSerialize["confidence"] = o.Confidence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentRootCause) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type       *AiAgentRootCauseType `json:"type,omitempty"`
		Text       *string               `json:"text,omitempty"`
		Confidence *AiAgentConfidence    `json:"confidence,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "text", "confidence"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Text = all.Text
	if all.Confidence != nil && !all.Confidence.IsValid() {
		hasInvalidField = true
	} else {
		o.Confidence = all.Confidence
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
