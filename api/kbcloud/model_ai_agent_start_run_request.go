// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentStartRunRequest struct {
	UserMessage string `json:"userMessage"`
	// Organization-scoped single-cluster Agent execution scope.
	Scope      AiAgentResourceScope `json:"scope"`
	PlaybookId *string              `json:"playbookId,omitempty"`
	Stream     *bool                `json:"stream,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentStartRunRequest instantiates a new AiAgentStartRunRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentStartRunRequest(userMessage string, scope AiAgentResourceScope) *AiAgentStartRunRequest {
	this := AiAgentStartRunRequest{}
	this.UserMessage = userMessage
	this.Scope = scope
	var stream bool = true
	this.Stream = &stream
	return &this
}

// NewAiAgentStartRunRequestWithDefaults instantiates a new AiAgentStartRunRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentStartRunRequestWithDefaults() *AiAgentStartRunRequest {
	this := AiAgentStartRunRequest{}
	var stream bool = true
	this.Stream = &stream
	return &this
}

// GetUserMessage returns the UserMessage field value.
func (o *AiAgentStartRunRequest) GetUserMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserMessage
}

// GetUserMessageOk returns a tuple with the UserMessage field value
// and a boolean to check if the value has been set.
func (o *AiAgentStartRunRequest) GetUserMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMessage, true
}

// SetUserMessage sets field value.
func (o *AiAgentStartRunRequest) SetUserMessage(v string) {
	o.UserMessage = v
}

// GetScope returns the Scope field value.
func (o *AiAgentStartRunRequest) GetScope() AiAgentResourceScope {
	if o == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AiAgentStartRunRequest) GetScopeOk() (*AiAgentResourceScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *AiAgentStartRunRequest) SetScope(v AiAgentResourceScope) {
	o.Scope = v
}

// GetPlaybookId returns the PlaybookId field value if set, zero value otherwise.
func (o *AiAgentStartRunRequest) GetPlaybookId() string {
	if o == nil || o.PlaybookId == nil {
		var ret string
		return ret
	}
	return *o.PlaybookId
}

// GetPlaybookIdOk returns a tuple with the PlaybookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentStartRunRequest) GetPlaybookIdOk() (*string, bool) {
	if o == nil || o.PlaybookId == nil {
		return nil, false
	}
	return o.PlaybookId, true
}

// HasPlaybookId returns a boolean if a field has been set.
func (o *AiAgentStartRunRequest) HasPlaybookId() bool {
	return o != nil && o.PlaybookId != nil
}

// SetPlaybookId gets a reference to the given string and assigns it to the PlaybookId field.
func (o *AiAgentStartRunRequest) SetPlaybookId(v string) {
	o.PlaybookId = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *AiAgentStartRunRequest) GetStream() bool {
	if o == nil || o.Stream == nil {
		var ret bool
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentStartRunRequest) GetStreamOk() (*bool, bool) {
	if o == nil || o.Stream == nil {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *AiAgentStartRunRequest) HasStream() bool {
	return o != nil && o.Stream != nil
}

// SetStream gets a reference to the given bool and assigns it to the Stream field.
func (o *AiAgentStartRunRequest) SetStream(v bool) {
	o.Stream = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentStartRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["userMessage"] = o.UserMessage
	toSerialize["scope"] = o.Scope
	if o.PlaybookId != nil {
		toSerialize["playbookId"] = o.PlaybookId
	}
	if o.Stream != nil {
		toSerialize["stream"] = o.Stream
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentStartRunRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserMessage *string               `json:"userMessage"`
		Scope       *AiAgentResourceScope `json:"scope"`
		PlaybookId  *string               `json:"playbookId,omitempty"`
		Stream      *bool                 `json:"stream,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.UserMessage == nil {
		return fmt.Errorf("required field userMessage missing")
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"userMessage", "scope", "playbookId", "stream"})
	} else {
		return err
	}

	hasInvalidField := false
	o.UserMessage = *all.UserMessage
	if all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = *all.Scope
	o.PlaybookId = all.PlaybookId
	o.Stream = all.Stream

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
