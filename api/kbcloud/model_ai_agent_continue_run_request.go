// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentContinueRunRequest struct {
	Message      AiAgentMessageInput `json:"message"`
	ContinueMode AiAgentContinueMode `json:"continueMode"`
	// Organization-scoped single-cluster Agent execution scope.
	ScopePatch *AiAgentResourceScope `json:"scopePatch,omitempty"`
	// Selected model name from the current organization's /llm/models response. Omit to keep the run default or prior selection.
	Model *string `json:"model,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentContinueRunRequest instantiates a new AiAgentContinueRunRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentContinueRunRequest(message AiAgentMessageInput, continueMode AiAgentContinueMode) *AiAgentContinueRunRequest {
	this := AiAgentContinueRunRequest{}
	this.Message = message
	this.ContinueMode = continueMode
	return &this
}

// NewAiAgentContinueRunRequestWithDefaults instantiates a new AiAgentContinueRunRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentContinueRunRequestWithDefaults() *AiAgentContinueRunRequest {
	this := AiAgentContinueRunRequest{}
	return &this
}

// GetMessage returns the Message field value.
func (o *AiAgentContinueRunRequest) GetMessage() AiAgentMessageInput {
	if o == nil {
		var ret AiAgentMessageInput
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunRequest) GetMessageOk() (*AiAgentMessageInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *AiAgentContinueRunRequest) SetMessage(v AiAgentMessageInput) {
	o.Message = v
}

// GetContinueMode returns the ContinueMode field value.
func (o *AiAgentContinueRunRequest) GetContinueMode() AiAgentContinueMode {
	if o == nil {
		var ret AiAgentContinueMode
		return ret
	}
	return o.ContinueMode
}

// GetContinueModeOk returns a tuple with the ContinueMode field value
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunRequest) GetContinueModeOk() (*AiAgentContinueMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContinueMode, true
}

// SetContinueMode sets field value.
func (o *AiAgentContinueRunRequest) SetContinueMode(v AiAgentContinueMode) {
	o.ContinueMode = v
}

// GetScopePatch returns the ScopePatch field value if set, zero value otherwise.
func (o *AiAgentContinueRunRequest) GetScopePatch() AiAgentResourceScope {
	if o == nil || o.ScopePatch == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return *o.ScopePatch
}

// GetScopePatchOk returns a tuple with the ScopePatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunRequest) GetScopePatchOk() (*AiAgentResourceScope, bool) {
	if o == nil || o.ScopePatch == nil {
		return nil, false
	}
	return o.ScopePatch, true
}

// HasScopePatch returns a boolean if a field has been set.
func (o *AiAgentContinueRunRequest) HasScopePatch() bool {
	return o != nil && o.ScopePatch != nil
}

// SetScopePatch gets a reference to the given AiAgentResourceScope and assigns it to the ScopePatch field.
func (o *AiAgentContinueRunRequest) SetScopePatch(v AiAgentResourceScope) {
	o.ScopePatch = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiAgentContinueRunRequest) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunRequest) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiAgentContinueRunRequest) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AiAgentContinueRunRequest) SetModel(v string) {
	o.Model = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentContinueRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["message"] = o.Message
	toSerialize["continueMode"] = o.ContinueMode
	if o.ScopePatch != nil {
		toSerialize["scopePatch"] = o.ScopePatch
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
func (o *AiAgentContinueRunRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message      *AiAgentMessageInput  `json:"message"`
		ContinueMode *AiAgentContinueMode  `json:"continueMode"`
		ScopePatch   *AiAgentResourceScope `json:"scopePatch,omitempty"`
		Model        *string               `json:"model,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.ContinueMode == nil {
		return fmt.Errorf("required field continueMode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"message", "continueMode", "scopePatch", "model"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Message.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Message = *all.Message
	if !all.ContinueMode.IsValid() {
		hasInvalidField = true
	} else {
		o.ContinueMode = *all.ContinueMode
	}
	if all.ScopePatch != nil && all.ScopePatch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScopePatch = all.ScopePatch
	o.Model = all.Model

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
