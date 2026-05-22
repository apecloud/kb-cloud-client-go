// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentToolApprovalDecisionRequest struct {
	ToolCallId      string  `json:"toolCallId"`
	ArgsHash        string  `json:"argsHash"`
	ApprovalVersion int64   `json:"approvalVersion"`
	Feedback        *string `json:"feedback,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentToolApprovalDecisionRequest instantiates a new AiAgentToolApprovalDecisionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentToolApprovalDecisionRequest(toolCallId string, argsHash string, approvalVersion int64) *AiAgentToolApprovalDecisionRequest {
	this := AiAgentToolApprovalDecisionRequest{}
	this.ToolCallId = toolCallId
	this.ArgsHash = argsHash
	this.ApprovalVersion = approvalVersion
	return &this
}

// NewAiAgentToolApprovalDecisionRequestWithDefaults instantiates a new AiAgentToolApprovalDecisionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentToolApprovalDecisionRequestWithDefaults() *AiAgentToolApprovalDecisionRequest {
	this := AiAgentToolApprovalDecisionRequest{}
	return &this
}

// GetToolCallId returns the ToolCallId field value.
func (o *AiAgentToolApprovalDecisionRequest) GetToolCallId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ToolCallId
}

// GetToolCallIdOk returns a tuple with the ToolCallId field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalDecisionRequest) GetToolCallIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToolCallId, true
}

// SetToolCallId sets field value.
func (o *AiAgentToolApprovalDecisionRequest) SetToolCallId(v string) {
	o.ToolCallId = v
}

// GetArgsHash returns the ArgsHash field value.
func (o *AiAgentToolApprovalDecisionRequest) GetArgsHash() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArgsHash
}

// GetArgsHashOk returns a tuple with the ArgsHash field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalDecisionRequest) GetArgsHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArgsHash, true
}

// SetArgsHash sets field value.
func (o *AiAgentToolApprovalDecisionRequest) SetArgsHash(v string) {
	o.ArgsHash = v
}

// GetApprovalVersion returns the ApprovalVersion field value.
func (o *AiAgentToolApprovalDecisionRequest) GetApprovalVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ApprovalVersion
}

// GetApprovalVersionOk returns a tuple with the ApprovalVersion field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalDecisionRequest) GetApprovalVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApprovalVersion, true
}

// SetApprovalVersion sets field value.
func (o *AiAgentToolApprovalDecisionRequest) SetApprovalVersion(v int64) {
	o.ApprovalVersion = v
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *AiAgentToolApprovalDecisionRequest) GetFeedback() string {
	if o == nil || o.Feedback == nil {
		var ret string
		return ret
	}
	return *o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolApprovalDecisionRequest) GetFeedbackOk() (*string, bool) {
	if o == nil || o.Feedback == nil {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *AiAgentToolApprovalDecisionRequest) HasFeedback() bool {
	return o != nil && o.Feedback != nil
}

// SetFeedback gets a reference to the given string and assigns it to the Feedback field.
func (o *AiAgentToolApprovalDecisionRequest) SetFeedback(v string) {
	o.Feedback = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentToolApprovalDecisionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["toolCallId"] = o.ToolCallId
	toSerialize["argsHash"] = o.ArgsHash
	toSerialize["approvalVersion"] = o.ApprovalVersion
	if o.Feedback != nil {
		toSerialize["feedback"] = o.Feedback
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentToolApprovalDecisionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ToolCallId      *string `json:"toolCallId"`
		ArgsHash        *string `json:"argsHash"`
		ApprovalVersion *int64  `json:"approvalVersion"`
		Feedback        *string `json:"feedback,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ToolCallId == nil {
		return fmt.Errorf("required field toolCallId missing")
	}
	if all.ArgsHash == nil {
		return fmt.Errorf("required field argsHash missing")
	}
	if all.ApprovalVersion == nil {
		return fmt.Errorf("required field approvalVersion missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"toolCallId", "argsHash", "approvalVersion", "feedback"})
	} else {
		return err
	}
	o.ToolCallId = *all.ToolCallId
	o.ArgsHash = *all.ArgsHash
	o.ApprovalVersion = *all.ApprovalVersion
	o.Feedback = all.Feedback

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
