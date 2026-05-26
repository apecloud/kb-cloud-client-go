// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentSendMessageRequest struct {
	Content  string         `json:"content"`
	Contexts []AiAgentScope `json:"contexts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentSendMessageRequest instantiates a new AiAgentSendMessageRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentSendMessageRequest(content string) *AiAgentSendMessageRequest {
	this := AiAgentSendMessageRequest{}
	this.Content = content
	return &this
}

// NewAiAgentSendMessageRequestWithDefaults instantiates a new AiAgentSendMessageRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentSendMessageRequestWithDefaults() *AiAgentSendMessageRequest {
	this := AiAgentSendMessageRequest{}
	return &this
}

// GetContent returns the Content field value.
func (o *AiAgentSendMessageRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AiAgentSendMessageRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *AiAgentSendMessageRequest) SetContent(v string) {
	o.Content = v
}

// GetContexts returns the Contexts field value if set, zero value otherwise.
func (o *AiAgentSendMessageRequest) GetContexts() []AiAgentScope {
	if o == nil || o.Contexts == nil {
		var ret []AiAgentScope
		return ret
	}
	return o.Contexts
}

// GetContextsOk returns a tuple with the Contexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSendMessageRequest) GetContextsOk() (*[]AiAgentScope, bool) {
	if o == nil || o.Contexts == nil {
		return nil, false
	}
	return &o.Contexts, true
}

// HasContexts returns a boolean if a field has been set.
func (o *AiAgentSendMessageRequest) HasContexts() bool {
	return o != nil && o.Contexts != nil
}

// SetContexts gets a reference to the given []AiAgentScope and assigns it to the Contexts field.
func (o *AiAgentSendMessageRequest) SetContexts(v []AiAgentScope) {
	o.Contexts = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentSendMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["content"] = o.Content
	if o.Contexts != nil {
		toSerialize["contexts"] = o.Contexts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentSendMessageRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Content  *string        `json:"content"`
		Contexts []AiAgentScope `json:"contexts,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"content", "contexts"})
	} else {
		return err
	}
	o.Content = *all.Content
	o.Contexts = all.Contexts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
