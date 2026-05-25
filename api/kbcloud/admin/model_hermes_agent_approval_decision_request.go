// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type HermesAgentApprovalDecisionRequest struct {
	Decision *HermesAgentApprovalDecision `json:"decision,omitempty"`
	Reason   *string                      `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHermesAgentApprovalDecisionRequest instantiates a new HermesAgentApprovalDecisionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHermesAgentApprovalDecisionRequest() *HermesAgentApprovalDecisionRequest {
	this := HermesAgentApprovalDecisionRequest{}
	return &this
}

// NewHermesAgentApprovalDecisionRequestWithDefaults instantiates a new HermesAgentApprovalDecisionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHermesAgentApprovalDecisionRequestWithDefaults() *HermesAgentApprovalDecisionRequest {
	this := HermesAgentApprovalDecisionRequest{}
	return &this
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *HermesAgentApprovalDecisionRequest) GetDecision() HermesAgentApprovalDecision {
	if o == nil || o.Decision == nil {
		var ret HermesAgentApprovalDecision
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentApprovalDecisionRequest) GetDecisionOk() (*HermesAgentApprovalDecision, bool) {
	if o == nil || o.Decision == nil {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *HermesAgentApprovalDecisionRequest) HasDecision() bool {
	return o != nil && o.Decision != nil
}

// SetDecision gets a reference to the given HermesAgentApprovalDecision and assigns it to the Decision field.
func (o *HermesAgentApprovalDecisionRequest) SetDecision(v HermesAgentApprovalDecision) {
	o.Decision = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *HermesAgentApprovalDecisionRequest) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HermesAgentApprovalDecisionRequest) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *HermesAgentApprovalDecisionRequest) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *HermesAgentApprovalDecisionRequest) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HermesAgentApprovalDecisionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Decision != nil {
		toSerialize["decision"] = o.Decision
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HermesAgentApprovalDecisionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Decision *HermesAgentApprovalDecision `json:"decision,omitempty"`
		Reason   *string                      `json:"reason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"decision", "reason"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Decision != nil && !all.Decision.IsValid() {
		hasInvalidField = true
	} else {
		o.Decision = all.Decision
	}
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
