// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMessagePart struct {
	Type    AiAgentMessagePartType `json:"type"`
	Text    *string                `json:"text,omitempty"`
	Summary *string                `json:"summary,omitempty"`
	Payload map[string]interface{} `json:"payload,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentMessagePart instantiates a new AiAgentMessagePart object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentMessagePart(typeVar AiAgentMessagePartType) *AiAgentMessagePart {
	this := AiAgentMessagePart{}
	this.Type = typeVar
	return &this
}

// NewAiAgentMessagePartWithDefaults instantiates a new AiAgentMessagePart object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentMessagePartWithDefaults() *AiAgentMessagePart {
	this := AiAgentMessagePart{}
	return &this
}

// GetType returns the Type field value.
func (o *AiAgentMessagePart) GetType() AiAgentMessagePartType {
	if o == nil {
		var ret AiAgentMessagePartType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetTypeOk() (*AiAgentMessagePartType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiAgentMessagePart) SetType(v AiAgentMessagePartType) {
	o.Type = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasText() bool {
	return o != nil && o.Text != nil
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *AiAgentMessagePart) SetText(v string) {
	o.Text = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AiAgentMessagePart) SetSummary(v string) {
	o.Summary = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AiAgentMessagePart) GetPayload() map[string]interface{} {
	if o == nil || o.Payload == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessagePart) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AiAgentMessagePart) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *AiAgentMessagePart) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentMessagePart) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["type"] = o.Type
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentMessagePart) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type    *AiAgentMessagePartType `json:"type"`
		Text    *string                 `json:"text,omitempty"`
		Summary *string                 `json:"summary,omitempty"`
		Payload map[string]interface{}  `json:"payload,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type", "text", "summary", "payload"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Text = all.Text
	o.Summary = all.Summary
	o.Payload = all.Payload

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
