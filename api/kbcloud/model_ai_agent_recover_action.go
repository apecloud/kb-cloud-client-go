// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentRecoverAction struct {
	ActionType AiAgentRecoverActionType `json:"actionType"`
	Label      string                   `json:"label"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	Status       *AiAgentStatusDisplay `json:"status,omitempty"`
	ContinueMode *AiAgentContinueMode  `json:"continueMode,omitempty"`
	// Organization-scoped single-cluster Agent execution scope.
	ScopePatch *AiAgentResourceScope `json:"scopePatch,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentRecoverAction instantiates a new AiAgentRecoverAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentRecoverAction(actionType AiAgentRecoverActionType, label string) *AiAgentRecoverAction {
	this := AiAgentRecoverAction{}
	this.ActionType = actionType
	this.Label = label
	return &this
}

// NewAiAgentRecoverActionWithDefaults instantiates a new AiAgentRecoverAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentRecoverActionWithDefaults() *AiAgentRecoverAction {
	this := AiAgentRecoverAction{}
	return &this
}

// GetActionType returns the ActionType field value.
func (o *AiAgentRecoverAction) GetActionType() AiAgentRecoverActionType {
	if o == nil {
		var ret AiAgentRecoverActionType
		return ret
	}
	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *AiAgentRecoverAction) GetActionTypeOk() (*AiAgentRecoverActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value.
func (o *AiAgentRecoverAction) SetActionType(v AiAgentRecoverActionType) {
	o.ActionType = v
}

// GetLabel returns the Label field value.
func (o *AiAgentRecoverAction) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AiAgentRecoverAction) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *AiAgentRecoverAction) SetLabel(v string) {
	o.Label = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiAgentRecoverAction) GetStatus() AiAgentStatusDisplay {
	if o == nil || o.Status == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRecoverAction) GetStatusOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiAgentRecoverAction) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AiAgentStatusDisplay and assigns it to the Status field.
func (o *AiAgentRecoverAction) SetStatus(v AiAgentStatusDisplay) {
	o.Status = &v
}

// GetContinueMode returns the ContinueMode field value if set, zero value otherwise.
func (o *AiAgentRecoverAction) GetContinueMode() AiAgentContinueMode {
	if o == nil || o.ContinueMode == nil {
		var ret AiAgentContinueMode
		return ret
	}
	return *o.ContinueMode
}

// GetContinueModeOk returns a tuple with the ContinueMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRecoverAction) GetContinueModeOk() (*AiAgentContinueMode, bool) {
	if o == nil || o.ContinueMode == nil {
		return nil, false
	}
	return o.ContinueMode, true
}

// HasContinueMode returns a boolean if a field has been set.
func (o *AiAgentRecoverAction) HasContinueMode() bool {
	return o != nil && o.ContinueMode != nil
}

// SetContinueMode gets a reference to the given AiAgentContinueMode and assigns it to the ContinueMode field.
func (o *AiAgentRecoverAction) SetContinueMode(v AiAgentContinueMode) {
	o.ContinueMode = &v
}

// GetScopePatch returns the ScopePatch field value if set, zero value otherwise.
func (o *AiAgentRecoverAction) GetScopePatch() AiAgentResourceScope {
	if o == nil || o.ScopePatch == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return *o.ScopePatch
}

// GetScopePatchOk returns a tuple with the ScopePatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRecoverAction) GetScopePatchOk() (*AiAgentResourceScope, bool) {
	if o == nil || o.ScopePatch == nil {
		return nil, false
	}
	return o.ScopePatch, true
}

// HasScopePatch returns a boolean if a field has been set.
func (o *AiAgentRecoverAction) HasScopePatch() bool {
	return o != nil && o.ScopePatch != nil
}

// SetScopePatch gets a reference to the given AiAgentResourceScope and assigns it to the ScopePatch field.
func (o *AiAgentRecoverAction) SetScopePatch(v AiAgentResourceScope) {
	o.ScopePatch = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentRecoverAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["actionType"] = o.ActionType
	toSerialize["label"] = o.Label
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ContinueMode != nil {
		toSerialize["continueMode"] = o.ContinueMode
	}
	if o.ScopePatch != nil {
		toSerialize["scopePatch"] = o.ScopePatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentRecoverAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionType   *AiAgentRecoverActionType `json:"actionType"`
		Label        *string                   `json:"label"`
		Status       *AiAgentStatusDisplay     `json:"status,omitempty"`
		ContinueMode *AiAgentContinueMode      `json:"continueMode,omitempty"`
		ScopePatch   *AiAgentResourceScope     `json:"scopePatch,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ActionType == nil {
		return fmt.Errorf("required field actionType missing")
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"actionType", "label", "status", "continueMode", "scopePatch"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.ActionType.IsValid() {
		hasInvalidField = true
	} else {
		o.ActionType = *all.ActionType
	}
	o.Label = *all.Label
	if all.Status != nil && all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = all.Status
	if all.ContinueMode != nil && !all.ContinueMode.IsValid() {
		hasInvalidField = true
	} else {
		o.ContinueMode = all.ContinueMode
	}
	if all.ScopePatch != nil && all.ScopePatch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScopePatch = all.ScopePatch

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
