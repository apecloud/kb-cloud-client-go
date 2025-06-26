// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AIConversationRequest struct {
	// Query type
	Type *AIChatType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAIConversationRequest instantiates a new AIConversationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAIConversationRequest() *AIConversationRequest {
	this := AIConversationRequest{}
	return &this
}

// NewAIConversationRequestWithDefaults instantiates a new AIConversationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAIConversationRequestWithDefaults() *AIConversationRequest {
	this := AIConversationRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AIConversationRequest) GetType() AIChatType {
	if o == nil || o.Type == nil {
		var ret AIChatType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIConversationRequest) GetTypeOk() (*AIChatType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AIConversationRequest) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given AIChatType and assigns it to the Type field.
func (o *AIConversationRequest) SetType(v AIChatType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AIConversationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AIConversationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type *AIChatType `json:"type,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
