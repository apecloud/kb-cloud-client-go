// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentCreateConversationRequest struct {
	Title            *string `json:"title,omitempty"`
	Model            *string `json:"model,omitempty"`
	EntryClusterName *string `json:"entryClusterName,omitempty"`
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

// GetEntryClusterName returns the EntryClusterName field value if set, zero value otherwise.
func (o *AiAgentCreateConversationRequest) GetEntryClusterName() string {
	if o == nil || o.EntryClusterName == nil {
		var ret string
		return ret
	}
	return *o.EntryClusterName
}

// GetEntryClusterNameOk returns a tuple with the EntryClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentCreateConversationRequest) GetEntryClusterNameOk() (*string, bool) {
	if o == nil || o.EntryClusterName == nil {
		return nil, false
	}
	return o.EntryClusterName, true
}

// HasEntryClusterName returns a boolean if a field has been set.
func (o *AiAgentCreateConversationRequest) HasEntryClusterName() bool {
	return o != nil && o.EntryClusterName != nil
}

// SetEntryClusterName gets a reference to the given string and assigns it to the EntryClusterName field.
func (o *AiAgentCreateConversationRequest) SetEntryClusterName(v string) {
	o.EntryClusterName = &v
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
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.EntryClusterName != nil {
		toSerialize["entryClusterName"] = o.EntryClusterName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentCreateConversationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Title            *string `json:"title,omitempty"`
		Model            *string `json:"model,omitempty"`
		EntryClusterName *string `json:"entryClusterName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"title", "model", "entryClusterName"})
	} else {
		return err
	}
	o.Title = all.Title
	o.Model = all.Model
	o.EntryClusterName = all.EntryClusterName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
