// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentConfirmActionPlanResponse struct {
	ActionPlanId    string                          `json:"actionPlanId"`
	ApprovalStatus  AiAgentActionPlanApprovalStatus `json:"approvalStatus"`
	ExecutionId     *string                         `json:"executionId,omitempty"`
	ExecutionStatus AiAgentExecutionStatus          `json:"executionStatus"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentConfirmActionPlanResponse instantiates a new AiAgentConfirmActionPlanResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentConfirmActionPlanResponse(actionPlanId string, approvalStatus AiAgentActionPlanApprovalStatus, executionStatus AiAgentExecutionStatus) *AiAgentConfirmActionPlanResponse {
	this := AiAgentConfirmActionPlanResponse{}
	this.ActionPlanId = actionPlanId
	this.ApprovalStatus = approvalStatus
	this.ExecutionStatus = executionStatus
	return &this
}

// NewAiAgentConfirmActionPlanResponseWithDefaults instantiates a new AiAgentConfirmActionPlanResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentConfirmActionPlanResponseWithDefaults() *AiAgentConfirmActionPlanResponse {
	this := AiAgentConfirmActionPlanResponse{}
	return &this
}

// GetActionPlanId returns the ActionPlanId field value.
func (o *AiAgentConfirmActionPlanResponse) GetActionPlanId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionPlanId
}

// GetActionPlanIdOk returns a tuple with the ActionPlanId field value
// and a boolean to check if the value has been set.
func (o *AiAgentConfirmActionPlanResponse) GetActionPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionPlanId, true
}

// SetActionPlanId sets field value.
func (o *AiAgentConfirmActionPlanResponse) SetActionPlanId(v string) {
	o.ActionPlanId = v
}

// GetApprovalStatus returns the ApprovalStatus field value.
func (o *AiAgentConfirmActionPlanResponse) GetApprovalStatus() AiAgentActionPlanApprovalStatus {
	if o == nil {
		var ret AiAgentActionPlanApprovalStatus
		return ret
	}
	return o.ApprovalStatus
}

// GetApprovalStatusOk returns a tuple with the ApprovalStatus field value
// and a boolean to check if the value has been set.
func (o *AiAgentConfirmActionPlanResponse) GetApprovalStatusOk() (*AiAgentActionPlanApprovalStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalStatus, true
}

// SetApprovalStatus sets field value.
func (o *AiAgentConfirmActionPlanResponse) SetApprovalStatus(v AiAgentActionPlanApprovalStatus) {
	o.ApprovalStatus = v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *AiAgentConfirmActionPlanResponse) GetExecutionId() string {
	if o == nil || o.ExecutionId == nil {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentConfirmActionPlanResponse) GetExecutionIdOk() (*string, bool) {
	if o == nil || o.ExecutionId == nil {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *AiAgentConfirmActionPlanResponse) HasExecutionId() bool {
	return o != nil && o.ExecutionId != nil
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *AiAgentConfirmActionPlanResponse) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetExecutionStatus returns the ExecutionStatus field value.
func (o *AiAgentConfirmActionPlanResponse) GetExecutionStatus() AiAgentExecutionStatus {
	if o == nil {
		var ret AiAgentExecutionStatus
		return ret
	}
	return o.ExecutionStatus
}

// GetExecutionStatusOk returns a tuple with the ExecutionStatus field value
// and a boolean to check if the value has been set.
func (o *AiAgentConfirmActionPlanResponse) GetExecutionStatusOk() (*AiAgentExecutionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecutionStatus, true
}

// SetExecutionStatus sets field value.
func (o *AiAgentConfirmActionPlanResponse) SetExecutionStatus(v AiAgentExecutionStatus) {
	o.ExecutionStatus = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentConfirmActionPlanResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["actionPlanId"] = o.ActionPlanId
	toSerialize["approvalStatus"] = o.ApprovalStatus
	if o.ExecutionId != nil {
		toSerialize["executionId"] = o.ExecutionId
	}
	toSerialize["executionStatus"] = o.ExecutionStatus

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentConfirmActionPlanResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionPlanId    *string                          `json:"actionPlanId"`
		ApprovalStatus  *AiAgentActionPlanApprovalStatus `json:"approvalStatus"`
		ExecutionId     *string                          `json:"executionId,omitempty"`
		ExecutionStatus *AiAgentExecutionStatus          `json:"executionStatus"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ActionPlanId == nil {
		return fmt.Errorf("required field actionPlanId missing")
	}
	if all.ApprovalStatus == nil {
		return fmt.Errorf("required field approvalStatus missing")
	}
	if all.ExecutionStatus == nil {
		return fmt.Errorf("required field executionStatus missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"actionPlanId", "approvalStatus", "executionId", "executionStatus"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ActionPlanId = *all.ActionPlanId
	if !all.ApprovalStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ApprovalStatus = *all.ApprovalStatus
	}
	o.ExecutionId = all.ExecutionId
	if !all.ExecutionStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ExecutionStatus = *all.ExecutionStatus
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
