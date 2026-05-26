// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentMessageInput struct {
	Role    AiAgentMessageRole `json:"role"`
	Content string             `json:"content"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentMessageInput instantiates a new AiAgentMessageInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentMessageInput(role AiAgentMessageRole, content string) *AiAgentMessageInput {
	this := AiAgentMessageInput{}
	this.Role = role
	this.Content = content
	return &this
}

// NewAiAgentMessageInputWithDefaults instantiates a new AiAgentMessageInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentMessageInputWithDefaults() *AiAgentMessageInput {
	this := AiAgentMessageInput{}
	return &this
}

// GetRole returns the Role field value.
func (o *AiAgentMessageInput) GetRole() AiAgentMessageRole {
	if o == nil {
		var ret AiAgentMessageRole
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *AiAgentMessageInput) GetRoleOk() (*AiAgentMessageRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *AiAgentMessageInput) SetRole(v AiAgentMessageRole) {
	o.Role = v
}

// GetContent returns the Content field value.
func (o *AiAgentMessageInput) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *AiAgentMessageInput) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *AiAgentMessageInput) SetContent(v string) {
	o.Content = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentMessageInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["role"] = o.Role
	toSerialize["content"] = o.Content

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentMessageInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role    *AiAgentMessageRole `json:"role"`
		Content *string             `json:"content"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "content"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = *all.Role
	}
	o.Content = *all.Content

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
