// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEvent struct {
	EventId        string              `json:"eventId"`
	RunId          string              `json:"runId"`
	ConversationId string              `json:"conversationId"`
	EventType      AiAgentEventType    `json:"eventType"`
	Status         *AiAgentEventStatus `json:"status,omitempty"`
	Timestamp      time.Time           `json:"timestamp"`
	Source         *AiAgentEventSource `json:"source,omitempty"`
	SchemaVersion  string              `json:"schemaVersion"`
	// Frontend-safe union payload for typed Agent events. For eventType=assistant_delta, messageId, partId, deltaKind, sequence, delta, and isFinal are required by contract; sequence is monotonically increasing per partId from 1 and is used by clients to order and deduplicate streamed deltas. Debug or hidden projections must not be sent on the ordinary user SSE stream.
	Payload *AiAgentEventPayload `json:"payload,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentEvent instantiates a new AiAgentEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentEvent(eventId string, runId string, conversationId string, eventType AiAgentEventType, timestamp time.Time, schemaVersion string) *AiAgentEvent {
	this := AiAgentEvent{}
	this.EventId = eventId
	this.RunId = runId
	this.ConversationId = conversationId
	this.EventType = eventType
	this.Timestamp = timestamp
	this.SchemaVersion = schemaVersion
	return &this
}

// NewAiAgentEventWithDefaults instantiates a new AiAgentEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentEventWithDefaults() *AiAgentEvent {
	this := AiAgentEvent{}
	return &this
}

// GetEventId returns the EventId field value.
func (o *AiAgentEvent) GetEventId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value.
func (o *AiAgentEvent) SetEventId(v string) {
	o.EventId = v
}

// GetRunId returns the RunId field value.
func (o *AiAgentEvent) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *AiAgentEvent) SetRunId(v string) {
	o.RunId = v
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentEvent) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentEvent) SetConversationId(v string) {
	o.ConversationId = v
}

// GetEventType returns the EventType field value.
func (o *AiAgentEvent) GetEventType() AiAgentEventType {
	if o == nil {
		var ret AiAgentEventType
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetEventTypeOk() (*AiAgentEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value.
func (o *AiAgentEvent) SetEventType(v AiAgentEventType) {
	o.EventType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiAgentEvent) GetStatus() AiAgentEventStatus {
	if o == nil || o.Status == nil {
		var ret AiAgentEventStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetStatusOk() (*AiAgentEventStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiAgentEvent) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AiAgentEventStatus and assigns it to the Status field.
func (o *AiAgentEvent) SetStatus(v AiAgentEventStatus) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value.
func (o *AiAgentEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *AiAgentEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AiAgentEvent) GetSource() AiAgentEventSource {
	if o == nil || o.Source == nil {
		var ret AiAgentEventSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetSourceOk() (*AiAgentEventSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AiAgentEvent) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given AiAgentEventSource and assigns it to the Source field.
func (o *AiAgentEvent) SetSource(v AiAgentEventSource) {
	o.Source = &v
}

// GetSchemaVersion returns the SchemaVersion field value.
func (o *AiAgentEvent) GetSchemaVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetSchemaVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value.
func (o *AiAgentEvent) SetSchemaVersion(v string) {
	o.SchemaVersion = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AiAgentEvent) GetPayload() AiAgentEventPayload {
	if o == nil || o.Payload == nil {
		var ret AiAgentEventPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvent) GetPayloadOk() (*AiAgentEventPayload, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AiAgentEvent) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given AiAgentEventPayload and assigns it to the Payload field.
func (o *AiAgentEvent) SetPayload(v AiAgentEventPayload) {
	o.Payload = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["eventId"] = o.EventId
	toSerialize["runId"] = o.RunId
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["eventType"] = o.EventType
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Timestamp.Nanosecond() == 0 {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	toSerialize["schemaVersion"] = o.SchemaVersion
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EventId        *string              `json:"eventId"`
		RunId          *string              `json:"runId"`
		ConversationId *string              `json:"conversationId"`
		EventType      *AiAgentEventType    `json:"eventType"`
		Status         *AiAgentEventStatus  `json:"status,omitempty"`
		Timestamp      *time.Time           `json:"timestamp"`
		Source         *AiAgentEventSource  `json:"source,omitempty"`
		SchemaVersion  *string              `json:"schemaVersion"`
		Payload        *AiAgentEventPayload `json:"payload,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EventId == nil {
		return fmt.Errorf("required field eventId missing")
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
	}
	if all.ConversationId == nil {
		return fmt.Errorf("required field conversationId missing")
	}
	if all.EventType == nil {
		return fmt.Errorf("required field eventType missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	if all.SchemaVersion == nil {
		return fmt.Errorf("required field schemaVersion missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"eventId", "runId", "conversationId", "eventType", "status", "timestamp", "source", "schemaVersion", "payload"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EventId = *all.EventId
	o.RunId = *all.RunId
	o.ConversationId = *all.ConversationId
	if !all.EventType.IsValid() {
		hasInvalidField = true
	} else {
		o.EventType = *all.EventType
	}
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Timestamp = *all.Timestamp
	if all.Source != nil && !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = all.Source
	}
	o.SchemaVersion = *all.SchemaVersion
	if all.Payload != nil && all.Payload.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Payload = all.Payload

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
