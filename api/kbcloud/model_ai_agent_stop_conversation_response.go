// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentStopConversationResponse struct {
	ConversationId string                    `json:"conversationId"`
	Status         AiAgentConversationStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentStopConversationResponse instantiates a new AiAgentStopConversationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentStopConversationResponse(conversationId string, status AiAgentConversationStatus) *AiAgentStopConversationResponse {
	this := AiAgentStopConversationResponse{}
	this.ConversationId = conversationId
	this.Status = status
	return &this
}

// NewAiAgentStopConversationResponseWithDefaults instantiates a new AiAgentStopConversationResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentStopConversationResponseWithDefaults() *AiAgentStopConversationResponse {
	this := AiAgentStopConversationResponse{}
	return &this
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentStopConversationResponse) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentStopConversationResponse) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentStopConversationResponse) SetConversationId(v string) {
	o.ConversationId = v
}

// GetStatus returns the Status field value.
func (o *AiAgentStopConversationResponse) GetStatus() AiAgentConversationStatus {
	if o == nil {
		var ret AiAgentConversationStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentStopConversationResponse) GetStatusOk() (*AiAgentConversationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentStopConversationResponse) SetStatus(v AiAgentConversationStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentStopConversationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentStopConversationResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConversationId *string                    `json:"conversationId"`
		Status         *AiAgentConversationStatus `json:"status"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ConversationId == nil {
		return fmt.Errorf("required field conversationId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"conversationId", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ConversationId = *all.ConversationId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
