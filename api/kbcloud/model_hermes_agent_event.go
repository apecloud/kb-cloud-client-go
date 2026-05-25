// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentEvent struct {
	// Cloud bridge replay cursor. This is also emitted as SSE id.
	EventId           string               `json:"eventId"`
	Type              HermesAgentEventType `json:"type"`
	Status            *string              `json:"status,omitempty"`
	Message           *string              `json:"message,omitempty"`
	RunId             string               `json:"runId"`
	HermesRunId       *string              `json:"hermesRunId,omitempty"`
	ConversationId    *string              `json:"conversationId,omitempty"`
	Timestamp         time.Time            `json:"timestamp"`
	ApprovalRequestId *string              `json:"approvalRequestId,omitempty"`
	ToolName          *string              `json:"toolName,omitempty"`
	ToolTarget        *string              `json:"toolTarget,omitempty"`
	ToolReason        *string              `json:"toolReason,omitempty"`
	ToolRisk          *string              `json:"toolRisk,omitempty"`
	ExpiresAt         *time.Time           `json:"expiresAt,omitempty"`
	FinalText         *string              `json:"finalText,omitempty"`
	ErrorMessage      *string              `json:"errorMessage,omitempty"`
	// Hermes raw event projection for debugging and parity checks. Frontend should prefer stable fields above.
	Payload map[string]interface{} `json:"payload,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHermesAgentEvent instantiates a new HermesAgentEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHermesAgentEvent(eventId string, typeVar HermesAgentEventType, runId string, timestamp time.Time) *HermesAgentEvent {
	this := HermesAgentEvent{}
	this.EventId = eventId
	this.Type = typeVar
	this.RunId = runId
	this.Timestamp = timestamp
	return &this
}

// NewHermesAgentEventWithDefaults instantiates a new HermesAgentEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHermesAgentEventWithDefaults() *HermesAgentEvent {
	this := HermesAgentEvent{}
	return &this
}

// GetEventId returns the EventId field value.
func (o *HermesAgentEvent) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value.
func (o *HermesAgentEvent) SetEventId(v string) {
	o.EventId = v
}

// GetType returns the Type field value.
func (o *HermesAgentEvent) GetType() HermesAgentEventType {
	if o == nil {
		var ret HermesAgentEventType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetTypeOk() (*HermesAgentEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *HermesAgentEvent) SetType(v HermesAgentEventType) {
	o.Type = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HermesAgentEvent) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HermesAgentEvent) SetMessage(v string) {
	o.Message = &v
}

// GetRunId returns the RunId field value.
func (o *HermesAgentEvent) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *HermesAgentEvent) SetRunId(v string) {
	o.RunId = v
}

// GetHermesRunId returns the HermesRunId field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetHermesRunId() string {
	if o == nil || o.HermesRunId == nil {
		var ret string
		return ret
	}
	return *o.HermesRunId
}

// GetHermesRunIdOk returns a tuple with the HermesRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetHermesRunIdOk() (*string, bool) {
	if o == nil || o.HermesRunId == nil {
		return nil, false
	}
	return o.HermesRunId, true
}

// HasHermesRunId returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasHermesRunId() bool {
	return o != nil && o.HermesRunId != nil
}

// SetHermesRunId gets a reference to the given string and assigns it to the HermesRunId field.
func (o *HermesAgentEvent) SetHermesRunId(v string) {
	o.HermesRunId = &v
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetConversationId() string {
	if o == nil || o.ConversationId == nil {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetConversationIdOk() (*string, bool) {
	if o == nil || o.ConversationId == nil {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasConversationId() bool {
	return o != nil && o.ConversationId != nil
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *HermesAgentEvent) SetConversationId(v string) {
	o.ConversationId = &v
}

// GetTimestamp returns the Timestamp field value.
func (o *HermesAgentEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *HermesAgentEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetApprovalRequestId returns the ApprovalRequestId field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetApprovalRequestId() string {
	if o == nil || o.ApprovalRequestId == nil {
		var ret string
		return ret
	}
	return *o.ApprovalRequestId
}

// GetApprovalRequestIdOk returns a tuple with the ApprovalRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetApprovalRequestIdOk() (*string, bool) {
	if o == nil || o.ApprovalRequestId == nil {
		return nil, false
	}
	return o.ApprovalRequestId, true
}

// HasApprovalRequestId returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasApprovalRequestId() bool {
	return o != nil && o.ApprovalRequestId != nil
}

// SetApprovalRequestId gets a reference to the given string and assigns it to the ApprovalRequestId field.
func (o *HermesAgentEvent) SetApprovalRequestId(v string) {
	o.ApprovalRequestId = &v
}

// GetToolName returns the ToolName field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetToolName() string {
	if o == nil || o.ToolName == nil {
		var ret string
		return ret
	}
	return *o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetToolNameOk() (*string, bool) {
	if o == nil || o.ToolName == nil {
		return nil, false
	}
	return o.ToolName, true
}

// HasToolName returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasToolName() bool {
	return o != nil && o.ToolName != nil
}

// SetToolName gets a reference to the given string and assigns it to the ToolName field.
func (o *HermesAgentEvent) SetToolName(v string) {
	o.ToolName = &v
}

// GetToolTarget returns the ToolTarget field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetToolTarget() string {
	if o == nil || o.ToolTarget == nil {
		var ret string
		return ret
	}
	return *o.ToolTarget
}

// GetToolTargetOk returns a tuple with the ToolTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetToolTargetOk() (*string, bool) {
	if o == nil || o.ToolTarget == nil {
		return nil, false
	}
	return o.ToolTarget, true
}

// HasToolTarget returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasToolTarget() bool {
	return o != nil && o.ToolTarget != nil
}

// SetToolTarget gets a reference to the given string and assigns it to the ToolTarget field.
func (o *HermesAgentEvent) SetToolTarget(v string) {
	o.ToolTarget = &v
}

// GetToolReason returns the ToolReason field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetToolReason() string {
	if o == nil || o.ToolReason == nil {
		var ret string
		return ret
	}
	return *o.ToolReason
}

// GetToolReasonOk returns a tuple with the ToolReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetToolReasonOk() (*string, bool) {
	if o == nil || o.ToolReason == nil {
		return nil, false
	}
	return o.ToolReason, true
}

// HasToolReason returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasToolReason() bool {
	return o != nil && o.ToolReason != nil
}

// SetToolReason gets a reference to the given string and assigns it to the ToolReason field.
func (o *HermesAgentEvent) SetToolReason(v string) {
	o.ToolReason = &v
}

// GetToolRisk returns the ToolRisk field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetToolRisk() string {
	if o == nil || o.ToolRisk == nil {
		var ret string
		return ret
	}
	return *o.ToolRisk
}

// GetToolRiskOk returns a tuple with the ToolRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetToolRiskOk() (*string, bool) {
	if o == nil || o.ToolRisk == nil {
		return nil, false
	}
	return o.ToolRisk, true
}

// HasToolRisk returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasToolRisk() bool {
	return o != nil && o.ToolRisk != nil
}

// SetToolRisk gets a reference to the given string and assigns it to the ToolRisk field.
func (o *HermesAgentEvent) SetToolRisk(v string) {
	o.ToolRisk = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasExpiresAt() bool {
	return o != nil && o.ExpiresAt != nil
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *HermesAgentEvent) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetFinalText returns the FinalText field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetFinalText() string {
	if o == nil || o.FinalText == nil {
		var ret string
		return ret
	}
	return *o.FinalText
}

// GetFinalTextOk returns a tuple with the FinalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetFinalTextOk() (*string, bool) {
	if o == nil || o.FinalText == nil {
		return nil, false
	}
	return o.FinalText, true
}

// HasFinalText returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasFinalText() bool {
	return o != nil && o.FinalText != nil
}

// SetFinalText gets a reference to the given string and assigns it to the FinalText field.
func (o *HermesAgentEvent) SetFinalText(v string) {
	o.FinalText = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *HermesAgentEvent) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *HermesAgentEvent) GetPayload() map[string]interface{} {
	if o == nil || o.Payload == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentEvent) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *HermesAgentEvent) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *HermesAgentEvent) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HermesAgentEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["eventId"] = o.EventId
	toSerialize["type"] = o.Type
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	toSerialize["runId"] = o.RunId
	if o.HermesRunId != nil {
		toSerialize["hermesRunId"] = o.HermesRunId
	}
	if o.ConversationId != nil {
		toSerialize["conversationId"] = o.ConversationId
	}
	if o.Timestamp.Nanosecond() == 0 {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ApprovalRequestId != nil {
		toSerialize["approvalRequestId"] = o.ApprovalRequestId
	}
	if o.ToolName != nil {
		toSerialize["toolName"] = o.ToolName
	}
	if o.ToolTarget != nil {
		toSerialize["toolTarget"] = o.ToolTarget
	}
	if o.ToolReason != nil {
		toSerialize["toolReason"] = o.ToolReason
	}
	if o.ToolRisk != nil {
		toSerialize["toolRisk"] = o.ToolRisk
	}
	if o.ExpiresAt != nil {
		if o.ExpiresAt.Nanosecond() == 0 {
			toSerialize["expiresAt"] = o.ExpiresAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["expiresAt"] = o.ExpiresAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.FinalText != nil {
		toSerialize["finalText"] = o.FinalText
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HermesAgentEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventId           *string                `json:"eventId"`
		Type              *HermesAgentEventType  `json:"type"`
		Status            *string                `json:"status,omitempty"`
		Message           *string                `json:"message,omitempty"`
		RunId             *string                `json:"runId"`
		HermesRunId       *string                `json:"hermesRunId,omitempty"`
		ConversationId    *string                `json:"conversationId,omitempty"`
		Timestamp         *time.Time             `json:"timestamp"`
		ApprovalRequestId *string                `json:"approvalRequestId,omitempty"`
		ToolName          *string                `json:"toolName,omitempty"`
		ToolTarget        *string                `json:"toolTarget,omitempty"`
		ToolReason        *string                `json:"toolReason,omitempty"`
		ToolRisk          *string                `json:"toolRisk,omitempty"`
		ExpiresAt         *time.Time             `json:"expiresAt,omitempty"`
		FinalText         *string                `json:"finalText,omitempty"`
		ErrorMessage      *string                `json:"errorMessage,omitempty"`
		Payload           map[string]interface{} `json:"payload,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EventId == nil {
		return fmt.Errorf("required field eventId missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"eventId", "type", "status", "message", "runId", "hermesRunId", "conversationId", "timestamp", "approvalRequestId", "toolName", "toolTarget", "toolReason", "toolRisk", "expiresAt", "finalText", "errorMessage", "payload"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EventId = *all.EventId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Status = all.Status
	o.Message = all.Message
	o.RunId = *all.RunId
	o.HermesRunId = all.HermesRunId
	o.ConversationId = all.ConversationId
	o.Timestamp = *all.Timestamp
	o.ApprovalRequestId = all.ApprovalRequestId
	o.ToolName = all.ToolName
	o.ToolTarget = all.ToolTarget
	o.ToolReason = all.ToolReason
	o.ToolRisk = all.ToolRisk
	o.ExpiresAt = all.ExpiresAt
	o.FinalText = all.FinalText
	o.ErrorMessage = all.ErrorMessage
	o.Payload = all.Payload

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
