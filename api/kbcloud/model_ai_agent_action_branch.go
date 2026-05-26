// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentActionBranch struct {
	Status *AiAgentActionBranchStatus `json:"status,omitempty"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	StatusDisplay *AiAgentStatusDisplay `json:"statusDisplay,omitempty"`
	ActionPlanId  *string               `json:"actionPlanId,omitempty"`
	ExecutionId   *string               `json:"executionId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentActionBranch instantiates a new AiAgentActionBranch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentActionBranch() *AiAgentActionBranch {
	this := AiAgentActionBranch{}
	return &this
}

// NewAiAgentActionBranchWithDefaults instantiates a new AiAgentActionBranch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentActionBranchWithDefaults() *AiAgentActionBranch {
	this := AiAgentActionBranch{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiAgentActionBranch) GetStatus() AiAgentActionBranchStatus {
	if o == nil || o.Status == nil {
		var ret AiAgentActionBranchStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionBranch) GetStatusOk() (*AiAgentActionBranchStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiAgentActionBranch) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AiAgentActionBranchStatus and assigns it to the Status field.
func (o *AiAgentActionBranch) SetStatus(v AiAgentActionBranchStatus) {
	o.Status = &v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *AiAgentActionBranch) GetStatusDisplay() AiAgentStatusDisplay {
	if o == nil || o.StatusDisplay == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionBranch) GetStatusDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *AiAgentActionBranch) HasStatusDisplay() bool {
	return o != nil && o.StatusDisplay != nil
}

// SetStatusDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the StatusDisplay field.
func (o *AiAgentActionBranch) SetStatusDisplay(v AiAgentStatusDisplay) {
	o.StatusDisplay = &v
}

// GetActionPlanId returns the ActionPlanId field value if set, zero value otherwise.
func (o *AiAgentActionBranch) GetActionPlanId() string {
	if o == nil || o.ActionPlanId == nil {
		var ret string
		return ret
	}
	return *o.ActionPlanId
}

// GetActionPlanIdOk returns a tuple with the ActionPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionBranch) GetActionPlanIdOk() (*string, bool) {
	if o == nil || o.ActionPlanId == nil {
		return nil, false
	}
	return o.ActionPlanId, true
}

// HasActionPlanId returns a boolean if a field has been set.
func (o *AiAgentActionBranch) HasActionPlanId() bool {
	return o != nil && o.ActionPlanId != nil
}

// SetActionPlanId gets a reference to the given string and assigns it to the ActionPlanId field.
func (o *AiAgentActionBranch) SetActionPlanId(v string) {
	o.ActionPlanId = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *AiAgentActionBranch) GetExecutionId() string {
	if o == nil || o.ExecutionId == nil {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentActionBranch) GetExecutionIdOk() (*string, bool) {
	if o == nil || o.ExecutionId == nil {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *AiAgentActionBranch) HasExecutionId() bool {
	return o != nil && o.ExecutionId != nil
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *AiAgentActionBranch) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentActionBranch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDisplay != nil {
		toSerialize["statusDisplay"] = o.StatusDisplay
	}
	if o.ActionPlanId != nil {
		toSerialize["actionPlanId"] = o.ActionPlanId
	}
	if o.ExecutionId != nil {
		toSerialize["executionId"] = o.ExecutionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentActionBranch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status        *AiAgentActionBranchStatus `json:"status,omitempty"`
		StatusDisplay *AiAgentStatusDisplay      `json:"statusDisplay,omitempty"`
		ActionPlanId  *string                    `json:"actionPlanId,omitempty"`
		ExecutionId   *string                    `json:"executionId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "statusDisplay", "actionPlanId", "executionId"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	if all.StatusDisplay != nil && all.StatusDisplay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusDisplay = all.StatusDisplay
	o.ActionPlanId = all.ActionPlanId
	o.ExecutionId = all.ExecutionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
