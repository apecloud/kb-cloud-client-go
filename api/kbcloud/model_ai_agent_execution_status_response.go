// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentExecutionStatusResponse struct {
	ExecutionId  *string                `json:"executionId,omitempty"`
	ActionPlanId string                 `json:"actionPlanId"`
	Status       AiAgentExecutionStatus `json:"status"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	StatusDisplay *AiAgentStatusDisplay  `json:"statusDisplay,omitempty"`
	StartedAt     *time.Time             `json:"startedAt,omitempty"`
	CompletedAt   *time.Time             `json:"completedAt,omitempty"`
	Steps         []AiAgentExecutionStep `json:"steps,omitempty"`
	UserMessage   *string                `json:"userMessage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentExecutionStatusResponse instantiates a new AiAgentExecutionStatusResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentExecutionStatusResponse(actionPlanId string, status AiAgentExecutionStatus) *AiAgentExecutionStatusResponse {
	this := AiAgentExecutionStatusResponse{}
	this.ActionPlanId = actionPlanId
	this.Status = status
	return &this
}

// NewAiAgentExecutionStatusResponseWithDefaults instantiates a new AiAgentExecutionStatusResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentExecutionStatusResponseWithDefaults() *AiAgentExecutionStatusResponse {
	this := AiAgentExecutionStatusResponse{}
	return &this
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *AiAgentExecutionStatusResponse) GetExecutionId() string {
	if o == nil || o.ExecutionId == nil {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentExecutionStatusResponse) GetExecutionIdOk() (*string, bool) {
	if o == nil || o.ExecutionId == nil {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *AiAgentExecutionStatusResponse) HasExecutionId() bool {
	return o != nil && o.ExecutionId != nil
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *AiAgentExecutionStatusResponse) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetActionPlanId returns the ActionPlanId field value.
func (o *AiAgentExecutionStatusResponse) GetActionPlanId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionPlanId
}

// GetActionPlanIdOk returns a tuple with the ActionPlanId field value
// and a boolean to check if the value has been set.
func (o *AiAgentExecutionStatusResponse) GetActionPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionPlanId, true
}

// SetActionPlanId sets field value.
func (o *AiAgentExecutionStatusResponse) SetActionPlanId(v string) {
	o.ActionPlanId = v
}

// GetStatus returns the Status field value.
func (o *AiAgentExecutionStatusResponse) GetStatus() AiAgentExecutionStatus {
	if o == nil {
		var ret AiAgentExecutionStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentExecutionStatusResponse) GetStatusOk() (*AiAgentExecutionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentExecutionStatusResponse) SetStatus(v AiAgentExecutionStatus) {
	o.Status = v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *AiAgentExecutionStatusResponse) GetStatusDisplay() AiAgentStatusDisplay {
	if o == nil || o.StatusDisplay == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentExecutionStatusResponse) GetStatusDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *AiAgentExecutionStatusResponse) HasStatusDisplay() bool {
	return o != nil && o.StatusDisplay != nil
}

// SetStatusDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the StatusDisplay field.
func (o *AiAgentExecutionStatusResponse) SetStatusDisplay(v AiAgentStatusDisplay) {
	o.StatusDisplay = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *AiAgentExecutionStatusResponse) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentExecutionStatusResponse) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *AiAgentExecutionStatusResponse) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *AiAgentExecutionStatusResponse) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *AiAgentExecutionStatusResponse) GetCompletedAt() time.Time {
	if o == nil || o.CompletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentExecutionStatusResponse) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || o.CompletedAt == nil {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *AiAgentExecutionStatusResponse) HasCompletedAt() bool {
	return o != nil && o.CompletedAt != nil
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *AiAgentExecutionStatusResponse) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *AiAgentExecutionStatusResponse) GetSteps() []AiAgentExecutionStep {
	if o == nil || o.Steps == nil {
		var ret []AiAgentExecutionStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentExecutionStatusResponse) GetStepsOk() (*[]AiAgentExecutionStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *AiAgentExecutionStatusResponse) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []AiAgentExecutionStep and assigns it to the Steps field.
func (o *AiAgentExecutionStatusResponse) SetSteps(v []AiAgentExecutionStep) {
	o.Steps = v
}

// GetUserMessage returns the UserMessage field value if set, zero value otherwise.
func (o *AiAgentExecutionStatusResponse) GetUserMessage() string {
	if o == nil || o.UserMessage == nil {
		var ret string
		return ret
	}
	return *o.UserMessage
}

// GetUserMessageOk returns a tuple with the UserMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentExecutionStatusResponse) GetUserMessageOk() (*string, bool) {
	if o == nil || o.UserMessage == nil {
		return nil, false
	}
	return o.UserMessage, true
}

// HasUserMessage returns a boolean if a field has been set.
func (o *AiAgentExecutionStatusResponse) HasUserMessage() bool {
	return o != nil && o.UserMessage != nil
}

// SetUserMessage gets a reference to the given string and assigns it to the UserMessage field.
func (o *AiAgentExecutionStatusResponse) SetUserMessage(v string) {
	o.UserMessage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentExecutionStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ExecutionId != nil {
		toSerialize["executionId"] = o.ExecutionId
	}
	toSerialize["actionPlanId"] = o.ActionPlanId
	toSerialize["status"] = o.Status
	if o.StatusDisplay != nil {
		toSerialize["statusDisplay"] = o.StatusDisplay
	}
	if o.StartedAt != nil {
		if o.StartedAt.Nanosecond() == 0 {
			toSerialize["startedAt"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startedAt"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CompletedAt != nil {
		if o.CompletedAt.Nanosecond() == 0 {
			toSerialize["completedAt"] = o.CompletedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["completedAt"] = o.CompletedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.UserMessage != nil {
		toSerialize["userMessage"] = o.UserMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentExecutionStatusResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExecutionId   *string                 `json:"executionId,omitempty"`
		ActionPlanId  *string                 `json:"actionPlanId"`
		Status        *AiAgentExecutionStatus `json:"status"`
		StatusDisplay *AiAgentStatusDisplay   `json:"statusDisplay,omitempty"`
		StartedAt     *time.Time              `json:"startedAt,omitempty"`
		CompletedAt   *time.Time              `json:"completedAt,omitempty"`
		Steps         []AiAgentExecutionStep  `json:"steps,omitempty"`
		UserMessage   *string                 `json:"userMessage,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ActionPlanId == nil {
		return fmt.Errorf("required field actionPlanId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"executionId", "actionPlanId", "status", "statusDisplay", "startedAt", "completedAt", "steps", "userMessage"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExecutionId = all.ExecutionId
	o.ActionPlanId = *all.ActionPlanId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if all.StatusDisplay != nil && all.StatusDisplay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusDisplay = all.StatusDisplay
	o.StartedAt = all.StartedAt
	o.CompletedAt = all.CompletedAt
	o.Steps = all.Steps
	o.UserMessage = all.UserMessage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
