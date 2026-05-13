// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentRunControlResponse struct {
	RunId  string           `json:"runId"`
	Status AiAgentRunStatus `json:"status"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	StatusDisplay *AiAgentStatusDisplay `json:"statusDisplay,omitempty"`
	CancelledAt   *time.Time            `json:"cancelledAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentRunControlResponse instantiates a new AiAgentRunControlResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentRunControlResponse(runId string, status AiAgentRunStatus) *AiAgentRunControlResponse {
	this := AiAgentRunControlResponse{}
	this.RunId = runId
	this.Status = status
	return &this
}

// NewAiAgentRunControlResponseWithDefaults instantiates a new AiAgentRunControlResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentRunControlResponseWithDefaults() *AiAgentRunControlResponse {
	this := AiAgentRunControlResponse{}
	return &this
}

// GetRunId returns the RunId field value.
func (o *AiAgentRunControlResponse) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AiAgentRunControlResponse) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *AiAgentRunControlResponse) SetRunId(v string) {
	o.RunId = v
}

// GetStatus returns the Status field value.
func (o *AiAgentRunControlResponse) GetStatus() AiAgentRunStatus {
	if o == nil {
		var ret AiAgentRunStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentRunControlResponse) GetStatusOk() (*AiAgentRunStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentRunControlResponse) SetStatus(v AiAgentRunStatus) {
	o.Status = v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *AiAgentRunControlResponse) GetStatusDisplay() AiAgentStatusDisplay {
	if o == nil || o.StatusDisplay == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRunControlResponse) GetStatusDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *AiAgentRunControlResponse) HasStatusDisplay() bool {
	return o != nil && o.StatusDisplay != nil
}

// SetStatusDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the StatusDisplay field.
func (o *AiAgentRunControlResponse) SetStatusDisplay(v AiAgentStatusDisplay) {
	o.StatusDisplay = &v
}

// GetCancelledAt returns the CancelledAt field value if set, zero value otherwise.
func (o *AiAgentRunControlResponse) GetCancelledAt() time.Time {
	if o == nil || o.CancelledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CancelledAt
}

// GetCancelledAtOk returns a tuple with the CancelledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRunControlResponse) GetCancelledAtOk() (*time.Time, bool) {
	if o == nil || o.CancelledAt == nil {
		return nil, false
	}
	return o.CancelledAt, true
}

// HasCancelledAt returns a boolean if a field has been set.
func (o *AiAgentRunControlResponse) HasCancelledAt() bool {
	return o != nil && o.CancelledAt != nil
}

// SetCancelledAt gets a reference to the given time.Time and assigns it to the CancelledAt field.
func (o *AiAgentRunControlResponse) SetCancelledAt(v time.Time) {
	o.CancelledAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentRunControlResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["runId"] = o.RunId
	toSerialize["status"] = o.Status
	if o.StatusDisplay != nil {
		toSerialize["statusDisplay"] = o.StatusDisplay
	}
	if o.CancelledAt != nil {
		if o.CancelledAt.Nanosecond() == 0 {
			toSerialize["cancelledAt"] = o.CancelledAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["cancelledAt"] = o.CancelledAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentRunControlResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RunId         *string               `json:"runId"`
		Status        *AiAgentRunStatus     `json:"status"`
		StatusDisplay *AiAgentStatusDisplay `json:"statusDisplay,omitempty"`
		CancelledAt   *time.Time            `json:"cancelledAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"runId", "status", "statusDisplay", "cancelledAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RunId = *all.RunId
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if all.StatusDisplay != nil && all.StatusDisplay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusDisplay = all.StatusDisplay
	o.CancelledAt = all.CancelledAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
