// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentSendMessageResponse struct {
	MessageId    string               `json:"messageId"`
	TurnId       string               `json:"turnId"`
	Status       AiAgentTurnStatus    `json:"status"`
	EventsUrl    *string              `json:"eventsUrl,omitempty"`
	AgentMode    *AiAgentMode         `json:"agentMode,omitempty"`
	AgentProfile *AiAgentProfile      `json:"agent_profile,omitempty"`
	ContextState *AiAgentContextState `json:"contextState,omitempty"`
	// Safe display label for the run scope. It does not expose profile paths, skills paths, credentials, or endpoints.
	ScopeLabel *string `json:"scopeLabel,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentSendMessageResponse instantiates a new AiAgentSendMessageResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentSendMessageResponse(messageId string, turnId string, status AiAgentTurnStatus) *AiAgentSendMessageResponse {
	this := AiAgentSendMessageResponse{}
	this.MessageId = messageId
	this.TurnId = turnId
	this.Status = status
	return &this
}

// NewAiAgentSendMessageResponseWithDefaults instantiates a new AiAgentSendMessageResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentSendMessageResponseWithDefaults() *AiAgentSendMessageResponse {
	this := AiAgentSendMessageResponse{}
	return &this
}

// GetMessageId returns the MessageId field value.
func (o *AiAgentSendMessageResponse) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AiAgentSendMessageResponse) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value.
func (o *AiAgentSendMessageResponse) SetMessageId(v string) {
	o.MessageId = v
}

// GetTurnId returns the TurnId field value.
func (o *AiAgentSendMessageResponse) GetTurnId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TurnId
}

// GetTurnIdOk returns a tuple with the TurnId field value
// and a boolean to check if the value has been set.
func (o *AiAgentSendMessageResponse) GetTurnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TurnId, true
}

// SetTurnId sets field value.
func (o *AiAgentSendMessageResponse) SetTurnId(v string) {
	o.TurnId = v
}

// GetStatus returns the Status field value.
func (o *AiAgentSendMessageResponse) GetStatus() AiAgentTurnStatus {
	if o == nil {
		var ret AiAgentTurnStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentSendMessageResponse) GetStatusOk() (*AiAgentTurnStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentSendMessageResponse) SetStatus(v AiAgentTurnStatus) {
	o.Status = v
}

// GetEventsUrl returns the EventsUrl field value if set, zero value otherwise.
func (o *AiAgentSendMessageResponse) GetEventsUrl() string {
	if o == nil || o.EventsUrl == nil {
		var ret string
		return ret
	}
	return *o.EventsUrl
}

// GetEventsUrlOk returns a tuple with the EventsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSendMessageResponse) GetEventsUrlOk() (*string, bool) {
	if o == nil || o.EventsUrl == nil {
		return nil, false
	}
	return o.EventsUrl, true
}

// HasEventsUrl returns a boolean if a field has been set.
func (o *AiAgentSendMessageResponse) HasEventsUrl() bool {
	return o != nil && o.EventsUrl != nil
}

// SetEventsUrl gets a reference to the given string and assigns it to the EventsUrl field.
func (o *AiAgentSendMessageResponse) SetEventsUrl(v string) {
	o.EventsUrl = &v
}

// GetAgentMode returns the AgentMode field value if set, zero value otherwise.
func (o *AiAgentSendMessageResponse) GetAgentMode() AiAgentMode {
	if o == nil || o.AgentMode == nil {
		var ret AiAgentMode
		return ret
	}
	return *o.AgentMode
}

// GetAgentModeOk returns a tuple with the AgentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSendMessageResponse) GetAgentModeOk() (*AiAgentMode, bool) {
	if o == nil || o.AgentMode == nil {
		return nil, false
	}
	return o.AgentMode, true
}

// HasAgentMode returns a boolean if a field has been set.
func (o *AiAgentSendMessageResponse) HasAgentMode() bool {
	return o != nil && o.AgentMode != nil
}

// SetAgentMode gets a reference to the given AiAgentMode and assigns it to the AgentMode field.
func (o *AiAgentSendMessageResponse) SetAgentMode(v AiAgentMode) {
	o.AgentMode = &v
}

// GetAgentProfile returns the AgentProfile field value if set, zero value otherwise.
func (o *AiAgentSendMessageResponse) GetAgentProfile() AiAgentProfile {
	if o == nil || o.AgentProfile == nil {
		var ret AiAgentProfile
		return ret
	}
	return *o.AgentProfile
}

// GetAgentProfileOk returns a tuple with the AgentProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSendMessageResponse) GetAgentProfileOk() (*AiAgentProfile, bool) {
	if o == nil || o.AgentProfile == nil {
		return nil, false
	}
	return o.AgentProfile, true
}

// HasAgentProfile returns a boolean if a field has been set.
func (o *AiAgentSendMessageResponse) HasAgentProfile() bool {
	return o != nil && o.AgentProfile != nil
}

// SetAgentProfile gets a reference to the given AiAgentProfile and assigns it to the AgentProfile field.
func (o *AiAgentSendMessageResponse) SetAgentProfile(v AiAgentProfile) {
	o.AgentProfile = &v
}

// GetContextState returns the ContextState field value if set, zero value otherwise.
func (o *AiAgentSendMessageResponse) GetContextState() AiAgentContextState {
	if o == nil || o.ContextState == nil {
		var ret AiAgentContextState
		return ret
	}
	return *o.ContextState
}

// GetContextStateOk returns a tuple with the ContextState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSendMessageResponse) GetContextStateOk() (*AiAgentContextState, bool) {
	if o == nil || o.ContextState == nil {
		return nil, false
	}
	return o.ContextState, true
}

// HasContextState returns a boolean if a field has been set.
func (o *AiAgentSendMessageResponse) HasContextState() bool {
	return o != nil && o.ContextState != nil
}

// SetContextState gets a reference to the given AiAgentContextState and assigns it to the ContextState field.
func (o *AiAgentSendMessageResponse) SetContextState(v AiAgentContextState) {
	o.ContextState = &v
}

// GetScopeLabel returns the ScopeLabel field value if set, zero value otherwise.
func (o *AiAgentSendMessageResponse) GetScopeLabel() string {
	if o == nil || o.ScopeLabel == nil {
		var ret string
		return ret
	}
	return *o.ScopeLabel
}

// GetScopeLabelOk returns a tuple with the ScopeLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSendMessageResponse) GetScopeLabelOk() (*string, bool) {
	if o == nil || o.ScopeLabel == nil {
		return nil, false
	}
	return o.ScopeLabel, true
}

// HasScopeLabel returns a boolean if a field has been set.
func (o *AiAgentSendMessageResponse) HasScopeLabel() bool {
	return o != nil && o.ScopeLabel != nil
}

// SetScopeLabel gets a reference to the given string and assigns it to the ScopeLabel field.
func (o *AiAgentSendMessageResponse) SetScopeLabel(v string) {
	o.ScopeLabel = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentSendMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["messageId"] = o.MessageId
	toSerialize["turnId"] = o.TurnId
	toSerialize["status"] = o.Status
	if o.EventsUrl != nil {
		toSerialize["eventsUrl"] = o.EventsUrl
	}
	if o.AgentMode != nil {
		toSerialize["agentMode"] = o.AgentMode
	}
	if o.AgentProfile != nil {
		toSerialize["agent_profile"] = o.AgentProfile
	}
	if o.ContextState != nil {
		toSerialize["contextState"] = o.ContextState
	}
	if o.ScopeLabel != nil {
		toSerialize["scopeLabel"] = o.ScopeLabel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentSendMessageResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MessageId    *string              `json:"messageId"`
		TurnId       *string              `json:"turnId"`
		Status       *AiAgentTurnStatus   `json:"status"`
		EventsUrl    *string              `json:"eventsUrl,omitempty"`
		AgentMode    *AiAgentMode         `json:"agentMode,omitempty"`
		AgentProfile *AiAgentProfile      `json:"agent_profile,omitempty"`
		ContextState *AiAgentContextState `json:"contextState,omitempty"`
		ScopeLabel   *string              `json:"scopeLabel,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.MessageId == nil {
		return fmt.Errorf("required field messageId missing")
	}
	if all.TurnId == nil {
		return fmt.Errorf("required field turnId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"messageId", "turnId", "status", "eventsUrl", "agentMode", "agent_profile", "contextState", "scopeLabel"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MessageId = *all.MessageId
	o.TurnId = *all.TurnId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.EventsUrl = all.EventsUrl
	if all.AgentMode != nil && !all.AgentMode.IsValid() {
		hasInvalidField = true
	} else {
		o.AgentMode = all.AgentMode
	}
	if all.AgentProfile != nil && !all.AgentProfile.IsValid() {
		hasInvalidField = true
	} else {
		o.AgentProfile = all.AgentProfile
	}
	if all.ContextState != nil && !all.ContextState.IsValid() {
		hasInvalidField = true
	} else {
		o.ContextState = all.ContextState
	}
	o.ScopeLabel = all.ScopeLabel

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
