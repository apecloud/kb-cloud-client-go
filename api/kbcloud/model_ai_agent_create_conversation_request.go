// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentCreateConversationRequest struct {
	Title *string `json:"title,omitempty"`
	// Organization-scoped single-cluster Agent execution scope.
	Scope AiAgentResourceScope `json:"scope"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentCreateConversationRequest instantiates a new AiAgentCreateConversationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentCreateConversationRequest(scope AiAgentResourceScope) *AiAgentCreateConversationRequest {
	this := AiAgentCreateConversationRequest{}
	this.Scope = scope
	return &this
}

// NewAiAgentCreateConversationRequestWithDefaults instantiates a new AiAgentCreateConversationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentCreateConversationRequestWithDefaults() *AiAgentCreateConversationRequest {
	this := AiAgentCreateConversationRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiAgentCreateConversationRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentCreateConversationRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiAgentCreateConversationRequest) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiAgentCreateConversationRequest) SetTitle(v string) {
	o.Title = &v
}

// GetScope returns the Scope field value.
func (o *AiAgentCreateConversationRequest) GetScope() AiAgentResourceScope {
	if o == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AiAgentCreateConversationRequest) GetScopeOk() (*AiAgentResourceScope, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value.
func (o *AiAgentCreateConversationRequest) SetScope(v AiAgentResourceScope) {
	o.Scope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentCreateConversationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	toSerialize["scope"] = o.Scope

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentCreateConversationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Title *string               `json:"title,omitempty"`
		Scope *AiAgentResourceScope `json:"scope"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Scope == nil {
		return fmt.Errorf("required field scope missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"title", "scope"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Title = all.Title
	if all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = *all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
