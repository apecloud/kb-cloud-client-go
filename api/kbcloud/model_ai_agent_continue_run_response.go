// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentContinueRunResponse struct {
	RunId           string                `json:"runId"`
	MessageId       string                `json:"messageId"`
	MessageOutcome  AiAgentMessageOutcome `json:"messageOutcome"`
	NewRunCreated   bool                  `json:"newRunCreated"`
	ContinuedRunId  *string               `json:"continuedRunId,omitempty"`
	NewRunId        *string               `json:"newRunId,omitempty"`
	RetryOfRunId    *string               `json:"retryOfRunId,omitempty"`
	ForkedFromRunId *string               `json:"forkedFromRunId,omitempty"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	Status    AiAgentStatusDisplay `json:"status"`
	EventsUrl *string              `json:"eventsUrl,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentContinueRunResponse instantiates a new AiAgentContinueRunResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentContinueRunResponse(runId string, messageId string, messageOutcome AiAgentMessageOutcome, newRunCreated bool, status AiAgentStatusDisplay) *AiAgentContinueRunResponse {
	this := AiAgentContinueRunResponse{}
	this.RunId = runId
	this.MessageId = messageId
	this.MessageOutcome = messageOutcome
	this.NewRunCreated = newRunCreated
	this.Status = status
	return &this
}

// NewAiAgentContinueRunResponseWithDefaults instantiates a new AiAgentContinueRunResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentContinueRunResponseWithDefaults() *AiAgentContinueRunResponse {
	this := AiAgentContinueRunResponse{}
	return &this
}

// GetRunId returns the RunId field value.
func (o *AiAgentContinueRunResponse) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunResponse) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *AiAgentContinueRunResponse) SetRunId(v string) {
	o.RunId = v
}

// GetMessageId returns the MessageId field value.
func (o *AiAgentContinueRunResponse) GetMessageId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunResponse) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageId, true
}

// SetMessageId sets field value.
func (o *AiAgentContinueRunResponse) SetMessageId(v string) {
	o.MessageId = v
}

// GetMessageOutcome returns the MessageOutcome field value.
func (o *AiAgentContinueRunResponse) GetMessageOutcome() AiAgentMessageOutcome {
	if o == nil {
		var ret AiAgentMessageOutcome
		return ret
	}
	return o.MessageOutcome
}

// GetMessageOutcomeOk returns a tuple with the MessageOutcome field value
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunResponse) GetMessageOutcomeOk() (*AiAgentMessageOutcome, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageOutcome, true
}

// SetMessageOutcome sets field value.
func (o *AiAgentContinueRunResponse) SetMessageOutcome(v AiAgentMessageOutcome) {
	o.MessageOutcome = v
}

// GetNewRunCreated returns the NewRunCreated field value.
func (o *AiAgentContinueRunResponse) GetNewRunCreated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.NewRunCreated
}

// GetNewRunCreatedOk returns a tuple with the NewRunCreated field value
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunResponse) GetNewRunCreatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewRunCreated, true
}

// SetNewRunCreated sets field value.
func (o *AiAgentContinueRunResponse) SetNewRunCreated(v bool) {
	o.NewRunCreated = v
}

// GetContinuedRunId returns the ContinuedRunId field value if set, zero value otherwise.
func (o *AiAgentContinueRunResponse) GetContinuedRunId() string {
	if o == nil || o.ContinuedRunId == nil {
		var ret string
		return ret
	}
	return *o.ContinuedRunId
}

// GetContinuedRunIdOk returns a tuple with the ContinuedRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunResponse) GetContinuedRunIdOk() (*string, bool) {
	if o == nil || o.ContinuedRunId == nil {
		return nil, false
	}
	return o.ContinuedRunId, true
}

// HasContinuedRunId returns a boolean if a field has been set.
func (o *AiAgentContinueRunResponse) HasContinuedRunId() bool {
	return o != nil && o.ContinuedRunId != nil
}

// SetContinuedRunId gets a reference to the given string and assigns it to the ContinuedRunId field.
func (o *AiAgentContinueRunResponse) SetContinuedRunId(v string) {
	o.ContinuedRunId = &v
}

// GetNewRunId returns the NewRunId field value if set, zero value otherwise.
func (o *AiAgentContinueRunResponse) GetNewRunId() string {
	if o == nil || o.NewRunId == nil {
		var ret string
		return ret
	}
	return *o.NewRunId
}

// GetNewRunIdOk returns a tuple with the NewRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunResponse) GetNewRunIdOk() (*string, bool) {
	if o == nil || o.NewRunId == nil {
		return nil, false
	}
	return o.NewRunId, true
}

// HasNewRunId returns a boolean if a field has been set.
func (o *AiAgentContinueRunResponse) HasNewRunId() bool {
	return o != nil && o.NewRunId != nil
}

// SetNewRunId gets a reference to the given string and assigns it to the NewRunId field.
func (o *AiAgentContinueRunResponse) SetNewRunId(v string) {
	o.NewRunId = &v
}

// GetRetryOfRunId returns the RetryOfRunId field value if set, zero value otherwise.
func (o *AiAgentContinueRunResponse) GetRetryOfRunId() string {
	if o == nil || o.RetryOfRunId == nil {
		var ret string
		return ret
	}
	return *o.RetryOfRunId
}

// GetRetryOfRunIdOk returns a tuple with the RetryOfRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunResponse) GetRetryOfRunIdOk() (*string, bool) {
	if o == nil || o.RetryOfRunId == nil {
		return nil, false
	}
	return o.RetryOfRunId, true
}

// HasRetryOfRunId returns a boolean if a field has been set.
func (o *AiAgentContinueRunResponse) HasRetryOfRunId() bool {
	return o != nil && o.RetryOfRunId != nil
}

// SetRetryOfRunId gets a reference to the given string and assigns it to the RetryOfRunId field.
func (o *AiAgentContinueRunResponse) SetRetryOfRunId(v string) {
	o.RetryOfRunId = &v
}

// GetForkedFromRunId returns the ForkedFromRunId field value if set, zero value otherwise.
func (o *AiAgentContinueRunResponse) GetForkedFromRunId() string {
	if o == nil || o.ForkedFromRunId == nil {
		var ret string
		return ret
	}
	return *o.ForkedFromRunId
}

// GetForkedFromRunIdOk returns a tuple with the ForkedFromRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunResponse) GetForkedFromRunIdOk() (*string, bool) {
	if o == nil || o.ForkedFromRunId == nil {
		return nil, false
	}
	return o.ForkedFromRunId, true
}

// HasForkedFromRunId returns a boolean if a field has been set.
func (o *AiAgentContinueRunResponse) HasForkedFromRunId() bool {
	return o != nil && o.ForkedFromRunId != nil
}

// SetForkedFromRunId gets a reference to the given string and assigns it to the ForkedFromRunId field.
func (o *AiAgentContinueRunResponse) SetForkedFromRunId(v string) {
	o.ForkedFromRunId = &v
}

// GetStatus returns the Status field value.
func (o *AiAgentContinueRunResponse) GetStatus() AiAgentStatusDisplay {
	if o == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunResponse) GetStatusOk() (*AiAgentStatusDisplay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentContinueRunResponse) SetStatus(v AiAgentStatusDisplay) {
	o.Status = v
}

// GetEventsUrl returns the EventsUrl field value if set, zero value otherwise.
func (o *AiAgentContinueRunResponse) GetEventsUrl() string {
	if o == nil || o.EventsUrl == nil {
		var ret string
		return ret
	}
	return *o.EventsUrl
}

// GetEventsUrlOk returns a tuple with the EventsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentContinueRunResponse) GetEventsUrlOk() (*string, bool) {
	if o == nil || o.EventsUrl == nil {
		return nil, false
	}
	return o.EventsUrl, true
}

// HasEventsUrl returns a boolean if a field has been set.
func (o *AiAgentContinueRunResponse) HasEventsUrl() bool {
	return o != nil && o.EventsUrl != nil
}

// SetEventsUrl gets a reference to the given string and assigns it to the EventsUrl field.
func (o *AiAgentContinueRunResponse) SetEventsUrl(v string) {
	o.EventsUrl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentContinueRunResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["runId"] = o.RunId
	toSerialize["messageId"] = o.MessageId
	toSerialize["messageOutcome"] = o.MessageOutcome
	toSerialize["newRunCreated"] = o.NewRunCreated
	if o.ContinuedRunId != nil {
		toSerialize["continuedRunId"] = o.ContinuedRunId
	}
	if o.NewRunId != nil {
		toSerialize["newRunId"] = o.NewRunId
	}
	if o.RetryOfRunId != nil {
		toSerialize["retryOfRunId"] = o.RetryOfRunId
	}
	if o.ForkedFromRunId != nil {
		toSerialize["forkedFromRunId"] = o.ForkedFromRunId
	}
	toSerialize["status"] = o.Status
	if o.EventsUrl != nil {
		toSerialize["eventsUrl"] = o.EventsUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentContinueRunResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RunId           *string                `json:"runId"`
		MessageId       *string                `json:"messageId"`
		MessageOutcome  *AiAgentMessageOutcome `json:"messageOutcome"`
		NewRunCreated   *bool                  `json:"newRunCreated"`
		ContinuedRunId  *string                `json:"continuedRunId,omitempty"`
		NewRunId        *string                `json:"newRunId,omitempty"`
		RetryOfRunId    *string                `json:"retryOfRunId,omitempty"`
		ForkedFromRunId *string                `json:"forkedFromRunId,omitempty"`
		Status          *AiAgentStatusDisplay  `json:"status"`
		EventsUrl       *string                `json:"eventsUrl,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
	}
	if all.MessageId == nil {
		return fmt.Errorf("required field messageId missing")
	}
	if all.MessageOutcome == nil {
		return fmt.Errorf("required field messageOutcome missing")
	}
	if all.NewRunCreated == nil {
		return fmt.Errorf("required field newRunCreated missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"runId", "messageId", "messageOutcome", "newRunCreated", "continuedRunId", "newRunId", "retryOfRunId", "forkedFromRunId", "status", "eventsUrl"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RunId = *all.RunId
	o.MessageId = *all.MessageId
	if !all.MessageOutcome.IsValid() {
		hasInvalidField = true
	} else {
		o.MessageOutcome = *all.MessageOutcome
	}
	o.NewRunCreated = *all.NewRunCreated
	o.ContinuedRunId = all.ContinuedRunId
	o.NewRunId = all.NewRunId
	o.RetryOfRunId = all.RetryOfRunId
	o.ForkedFromRunId = all.ForkedFromRunId
	if all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = *all.Status
	o.EventsUrl = all.EventsUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
