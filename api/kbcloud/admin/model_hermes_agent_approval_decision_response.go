// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HermesAgentApprovalDecisionResponse struct {
	RunId             string    `json:"runId"`
	ApprovalRequestId string    `json:"approvalRequestId"`
	Decision          string    `json:"decision"`
	CreatedAt         time.Time `json:"createdAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHermesAgentApprovalDecisionResponse instantiates a new HermesAgentApprovalDecisionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHermesAgentApprovalDecisionResponse(runId string, approvalRequestId string, decision string, createdAt time.Time) *HermesAgentApprovalDecisionResponse {
	this := HermesAgentApprovalDecisionResponse{}
	this.RunId = runId
	this.ApprovalRequestId = approvalRequestId
	this.Decision = decision
	this.CreatedAt = createdAt
	return &this
}

// NewHermesAgentApprovalDecisionResponseWithDefaults instantiates a new HermesAgentApprovalDecisionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHermesAgentApprovalDecisionResponseWithDefaults() *HermesAgentApprovalDecisionResponse {
	this := HermesAgentApprovalDecisionResponse{}
	return &this
}

// GetRunId returns the RunId field value.
func (o *HermesAgentApprovalDecisionResponse) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *HermesAgentApprovalDecisionResponse) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *HermesAgentApprovalDecisionResponse) SetRunId(v string) {
	o.RunId = v
}

// GetApprovalRequestId returns the ApprovalRequestId field value.
func (o *HermesAgentApprovalDecisionResponse) GetApprovalRequestId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApprovalRequestId
}

// GetApprovalRequestIdOk returns a tuple with the ApprovalRequestId field value
// and a boolean to check if the value has been set.
func (o *HermesAgentApprovalDecisionResponse) GetApprovalRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalRequestId, true
}

// SetApprovalRequestId sets field value.
func (o *HermesAgentApprovalDecisionResponse) SetApprovalRequestId(v string) {
	o.ApprovalRequestId = v
}

// GetDecision returns the Decision field value.
func (o *HermesAgentApprovalDecisionResponse) GetDecision() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *HermesAgentApprovalDecisionResponse) GetDecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value.
func (o *HermesAgentApprovalDecisionResponse) SetDecision(v string) {
	o.Decision = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *HermesAgentApprovalDecisionResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *HermesAgentApprovalDecisionResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *HermesAgentApprovalDecisionResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HermesAgentApprovalDecisionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["runId"] = o.RunId
	toSerialize["approvalRequestId"] = o.ApprovalRequestId
	toSerialize["decision"] = o.Decision
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
func (o *HermesAgentApprovalDecisionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RunId             *string    `json:"runId"`
		ApprovalRequestId *string    `json:"approvalRequestId"`
		Decision          *string    `json:"decision"`
		CreatedAt         *time.Time `json:"createdAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
	}
	if all.ApprovalRequestId == nil {
		return fmt.Errorf("required field approvalRequestId missing")
	}
	if all.Decision == nil {
		return fmt.Errorf("required field decision missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"runId", "approvalRequestId", "decision", "createdAt"})
	} else {
		return err
	}
	o.RunId = *all.RunId
	o.ApprovalRequestId = *all.ApprovalRequestId
	o.Decision = *all.Decision
	o.CreatedAt = *all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
