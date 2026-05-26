// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentStartRunResponse struct {
	RunId          string           `json:"runId"`
	ConversationId string           `json:"conversationId"`
	Status         AiAgentRunStatus `json:"status"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	StatusDisplay *AiAgentStatusDisplay `json:"statusDisplay,omitempty"`
	EventsUrl     string                `json:"eventsUrl"`
	CreatedAt     time.Time             `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentStartRunResponse instantiates a new AiAgentStartRunResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentStartRunResponse(runId string, conversationId string, status AiAgentRunStatus, eventsUrl string, createdAt time.Time) *AiAgentStartRunResponse {
	this := AiAgentStartRunResponse{}
	this.RunId = runId
	this.ConversationId = conversationId
	this.Status = status
	this.EventsUrl = eventsUrl
	this.CreatedAt = createdAt
	return &this
}

// NewAiAgentStartRunResponseWithDefaults instantiates a new AiAgentStartRunResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentStartRunResponseWithDefaults() *AiAgentStartRunResponse {
	this := AiAgentStartRunResponse{}
	return &this
}

// GetRunId returns the RunId field value.
func (o *AiAgentStartRunResponse) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AiAgentStartRunResponse) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *AiAgentStartRunResponse) SetRunId(v string) {
	o.RunId = v
}

// GetConversationId returns the ConversationId field value.
func (o *AiAgentStartRunResponse) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentStartRunResponse) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value.
func (o *AiAgentStartRunResponse) SetConversationId(v string) {
	o.ConversationId = v
}

// GetStatus returns the Status field value.
func (o *AiAgentStartRunResponse) GetStatus() AiAgentRunStatus {
	if o == nil {
		var ret AiAgentRunStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentStartRunResponse) GetStatusOk() (*AiAgentRunStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentStartRunResponse) SetStatus(v AiAgentRunStatus) {
	o.Status = v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *AiAgentStartRunResponse) GetStatusDisplay() AiAgentStatusDisplay {
	if o == nil || o.StatusDisplay == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentStartRunResponse) GetStatusDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *AiAgentStartRunResponse) HasStatusDisplay() bool {
	return o != nil && o.StatusDisplay != nil
}

// SetStatusDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the StatusDisplay field.
func (o *AiAgentStartRunResponse) SetStatusDisplay(v AiAgentStatusDisplay) {
	o.StatusDisplay = &v
}

// GetEventsUrl returns the EventsUrl field value.
func (o *AiAgentStartRunResponse) GetEventsUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventsUrl
}

// GetEventsUrlOk returns a tuple with the EventsUrl field value
// and a boolean to check if the value has been set.
func (o *AiAgentStartRunResponse) GetEventsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventsUrl, true
}

// SetEventsUrl sets field value.
func (o *AiAgentStartRunResponse) SetEventsUrl(v string) {
	o.EventsUrl = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiAgentStartRunResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiAgentStartRunResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiAgentStartRunResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentStartRunResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["runId"] = o.RunId
	toSerialize["conversationId"] = o.ConversationId
	toSerialize["status"] = o.Status
	if o.StatusDisplay != nil {
		toSerialize["statusDisplay"] = o.StatusDisplay
	}
	toSerialize["eventsUrl"] = o.EventsUrl
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentStartRunResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RunId          *string               `json:"runId"`
		ConversationId *string               `json:"conversationId"`
		Status         *AiAgentRunStatus     `json:"status"`
		StatusDisplay  *AiAgentStatusDisplay `json:"statusDisplay,omitempty"`
		EventsUrl      *string               `json:"eventsUrl"`
		CreatedAt      *time.Time            `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
	}
	if all.ConversationId == nil {
		return fmt.Errorf("required field conversationId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.EventsUrl == nil {
		return fmt.Errorf("required field eventsUrl missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"runId", "conversationId", "status", "statusDisplay", "eventsUrl", "createdAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RunId = *all.RunId
	o.ConversationId = *all.ConversationId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if all.StatusDisplay != nil && all.StatusDisplay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusDisplay = all.StatusDisplay
	o.EventsUrl = *all.EventsUrl
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
