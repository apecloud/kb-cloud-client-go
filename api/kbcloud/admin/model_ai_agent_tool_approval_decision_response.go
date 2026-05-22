// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentToolApprovalDecisionResponse struct {
	ApprovalRequest AiAgentToolApprovalRequest `json:"approvalRequest"`
	Run             AiAgentRun                 `json:"run"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentToolApprovalDecisionResponse instantiates a new AiAgentToolApprovalDecisionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentToolApprovalDecisionResponse(approvalRequest AiAgentToolApprovalRequest, run AiAgentRun) *AiAgentToolApprovalDecisionResponse {
	this := AiAgentToolApprovalDecisionResponse{}
	this.ApprovalRequest = approvalRequest
	this.Run = run
	return &this
}

// NewAiAgentToolApprovalDecisionResponseWithDefaults instantiates a new AiAgentToolApprovalDecisionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentToolApprovalDecisionResponseWithDefaults() *AiAgentToolApprovalDecisionResponse {
	this := AiAgentToolApprovalDecisionResponse{}
	return &this
}

// GetApprovalRequest returns the ApprovalRequest field value.
func (o *AiAgentToolApprovalDecisionResponse) GetApprovalRequest() AiAgentToolApprovalRequest {
	if o == nil {
		var ret AiAgentToolApprovalRequest
		return ret
	}
	return o.ApprovalRequest
}

// GetApprovalRequestOk returns a tuple with the ApprovalRequest field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalDecisionResponse) GetApprovalRequestOk() (*AiAgentToolApprovalRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalRequest, true
}

// SetApprovalRequest sets field value.
func (o *AiAgentToolApprovalDecisionResponse) SetApprovalRequest(v AiAgentToolApprovalRequest) {
	o.ApprovalRequest = v
}

// GetRun returns the Run field value.
func (o *AiAgentToolApprovalDecisionResponse) GetRun() AiAgentRun {
	if o == nil {
		var ret AiAgentRun
		return ret
	}
	return o.Run
}

// GetRunOk returns a tuple with the Run field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalDecisionResponse) GetRunOk() (*AiAgentRun, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Run, true
}

// SetRun sets field value.
func (o *AiAgentToolApprovalDecisionResponse) SetRun(v AiAgentRun) {
	o.Run = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentToolApprovalDecisionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["approvalRequest"] = o.ApprovalRequest
	toSerialize["run"] = o.Run

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentToolApprovalDecisionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApprovalRequest *AiAgentToolApprovalRequest `json:"approvalRequest"`
		Run             *AiAgentRun                 `json:"run"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ApprovalRequest == nil {
		return fmt.Errorf("required field approvalRequest missing")
	}
	if all.Run == nil {
		return fmt.Errorf("required field run missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"approvalRequest", "run"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ApprovalRequest.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ApprovalRequest = *all.ApprovalRequest
	if all.Run.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Run = *all.Run

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
