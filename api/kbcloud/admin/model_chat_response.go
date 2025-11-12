// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ChatResponse Chat message response
type ChatResponse struct {
	// Arbitrary data
	Data map[string]interface{} `json:"data"`
	// Message type
	Type ChatResponseType `json:"type"`
	// Message ID
	MessageId string `json:"messageID"`
	// Parent message ID
	ParentId string `json:"parentID"`
	// Session ID
	SessionId string `json:"sessionId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChatResponse instantiates a new ChatResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChatResponse(data map[string]interface{}, typeVar ChatResponseType, messageId string, parentId string, sessionId string) *ChatResponse {
	this := ChatResponse{}
	this.Data = data
	this.Type = typeVar
	this.MessageId = messageId
	this.ParentId = parentId
	this.SessionId = sessionId
	return &this
}

// NewChatResponseWithDefaults instantiates a new ChatResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChatResponseWithDefaults() *ChatResponse {
	this := ChatResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ChatResponse) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ChatResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetType returns the Type field value.
func (o *ChatResponse) GetType() ChatResponseType {
	if o == nil {
		var ret ChatResponseType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetTypeOk() (*ChatResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ChatResponse) SetType(v ChatResponseType) {
	o.Type = v
}

// GetMessageId returns the MessageId field value.
func (o *ChatResponse) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value.
func (o *ChatResponse) SetMessageId(v string) {
	o.MessageId = v
}

// GetParentId returns the ParentId field value.
func (o *ChatResponse) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value.
func (o *ChatResponse) SetParentId(v string) {
	o.ParentId = v
}

// GetSessionId returns the SessionId field value.
func (o *ChatResponse) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ChatResponse) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *ChatResponse) SetSessionId(v string) {
	o.SessionId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChatResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["type"] = o.Type
	toSerialize["messageID"] = o.MessageId
	toSerialize["parentID"] = o.ParentId
	toSerialize["sessionId"] = o.SessionId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChatResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data      *map[string]interface{} `json:"data"`
		Type      *ChatResponseType       `json:"type"`
		MessageId *string                 `json:"messageID"`
		ParentId  *string                 `json:"parentID"`
		SessionId *string                 `json:"sessionId"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.MessageId == nil {
		return fmt.Errorf("required field messageID missing")
	}
	if all.ParentId == nil {
		return fmt.Errorf("required field parentID missing")
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field sessionId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"data", "type", "messageID", "parentID", "sessionId"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.MessageId = *all.MessageId
	o.ParentId = *all.ParentId
	o.SessionId = *all.SessionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
