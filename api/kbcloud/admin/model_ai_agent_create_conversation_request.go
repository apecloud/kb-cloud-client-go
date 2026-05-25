// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentCreateConversationRequest struct {
	Scope *AiAgentScope `json:"scope,omitempty"`
	Title *string       `json:"title,omitempty"`
	Model *string       `json:"model,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentCreateConversationRequest instantiates a new AiAgentCreateConversationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentCreateConversationRequest() *AiAgentCreateConversationRequest {
	this := AiAgentCreateConversationRequest{}
	return &this
}

// NewAiAgentCreateConversationRequestWithDefaults instantiates a new AiAgentCreateConversationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentCreateConversationRequestWithDefaults() *AiAgentCreateConversationRequest {
	this := AiAgentCreateConversationRequest{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AiAgentCreateConversationRequest) GetScope() AiAgentScope {
	if o == nil || o.Scope == nil {
		var ret AiAgentScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentCreateConversationRequest) GetScopeOk() (*AiAgentScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AiAgentCreateConversationRequest) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given AiAgentScope and assigns it to the Scope field.
func (o *AiAgentCreateConversationRequest) SetScope(v AiAgentScope) {
	o.Scope = &v
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

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiAgentCreateConversationRequest) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentCreateConversationRequest) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiAgentCreateConversationRequest) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiAgentCreateConversationRequest) SetModel(v string) {
	o.Model = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentCreateConversationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentCreateConversationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Scope *AiAgentScope `json:"scope,omitempty"`
		Title *string       `json:"title,omitempty"`
		Model *string       `json:"model,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"scope", "title", "model"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Scope != nil && all.Scope.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Scope = all.Scope
	o.Title = all.Title
	o.Model = all.Model

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
