// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentSuggestion struct {
	SuggestionId string  `json:"suggestionId"`
	Label        string  `json:"label"`
	Prompt       string  `json:"prompt"`
	PlaybookId   *string `json:"playbookId,omitempty"`
	// Organization-scoped single-cluster Agent execution scope.
	ScopePatch *AiAgentResourceScope `json:"scopePatch,omitempty"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	DisabledReason     *AiAgentStatusDisplay `json:"disabledReason,omitempty"`
	DisabledReasonCode *string               `json:"disabledReasonCode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentSuggestion instantiates a new AiAgentSuggestion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentSuggestion(suggestionId string, label string, prompt string) *AiAgentSuggestion {
	this := AiAgentSuggestion{}
	this.SuggestionId = suggestionId
	this.Label = label
	this.Prompt = prompt
	return &this
}

// NewAiAgentSuggestionWithDefaults instantiates a new AiAgentSuggestion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentSuggestionWithDefaults() *AiAgentSuggestion {
	this := AiAgentSuggestion{}
	return &this
}

// GetSuggestionId returns the SuggestionId field value.
func (o *AiAgentSuggestion) GetSuggestionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SuggestionId
}

// GetSuggestionIdOk returns a tuple with the SuggestionId field value
// and a boolean to check if the value has been set.
func (o *AiAgentSuggestion) GetSuggestionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuggestionId, true
}

// SetSuggestionId sets field value.
func (o *AiAgentSuggestion) SetSuggestionId(v string) {
	o.SuggestionId = v
}

// GetLabel returns the Label field value.
func (o *AiAgentSuggestion) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AiAgentSuggestion) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *AiAgentSuggestion) SetLabel(v string) {
	o.Label = v
}

// GetPrompt returns the Prompt field value.
func (o *AiAgentSuggestion) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *AiAgentSuggestion) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value.
func (o *AiAgentSuggestion) SetPrompt(v string) {
	o.Prompt = v
}

// GetPlaybookId returns the PlaybookId field value if set, zero value otherwise.
func (o *AiAgentSuggestion) GetPlaybookId() string {
	if o == nil || o.PlaybookId == nil {
		var ret string
		return ret
	}
	return *o.PlaybookId
}

// GetPlaybookIdOk returns a tuple with the PlaybookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSuggestion) GetPlaybookIdOk() (*string, bool) {
	if o == nil || o.PlaybookId == nil {
		return nil, false
	}
	return o.PlaybookId, true
}

// HasPlaybookId returns a boolean if a field has been set.
func (o *AiAgentSuggestion) HasPlaybookId() bool {
	return o != nil && o.PlaybookId != nil
}

// SetPlaybookId gets a reference to the given string and assigns it to the PlaybookId field.
func (o *AiAgentSuggestion) SetPlaybookId(v string) {
	o.PlaybookId = &v
}

// GetScopePatch returns the ScopePatch field value if set, zero value otherwise.
func (o *AiAgentSuggestion) GetScopePatch() AiAgentResourceScope {
	if o == nil || o.ScopePatch == nil {
		var ret AiAgentResourceScope
		return ret
	}
	return *o.ScopePatch
}

// GetScopePatchOk returns a tuple with the ScopePatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSuggestion) GetScopePatchOk() (*AiAgentResourceScope, bool) {
	if o == nil || o.ScopePatch == nil {
		return nil, false
	}
	return o.ScopePatch, true
}

// HasScopePatch returns a boolean if a field has been set.
func (o *AiAgentSuggestion) HasScopePatch() bool {
	return o != nil && o.ScopePatch != nil
}

// SetScopePatch gets a reference to the given AiAgentResourceScope and assigns it to the ScopePatch field.
func (o *AiAgentSuggestion) SetScopePatch(v AiAgentResourceScope) {
	o.ScopePatch = &v
}

// GetDisabledReason returns the DisabledReason field value if set, zero value otherwise.
func (o *AiAgentSuggestion) GetDisabledReason() AiAgentStatusDisplay {
	if o == nil || o.DisabledReason == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.DisabledReason
}

// GetDisabledReasonOk returns a tuple with the DisabledReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSuggestion) GetDisabledReasonOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.DisabledReason == nil {
		return nil, false
	}
	return o.DisabledReason, true
}

// HasDisabledReason returns a boolean if a field has been set.
func (o *AiAgentSuggestion) HasDisabledReason() bool {
	return o != nil && o.DisabledReason != nil
}

// SetDisabledReason gets a reference to the given AiAgentStatusDisplay and assigns it to the DisabledReason field.
func (o *AiAgentSuggestion) SetDisabledReason(v AiAgentStatusDisplay) {
	o.DisabledReason = &v
}

// GetDisabledReasonCode returns the DisabledReasonCode field value if set, zero value otherwise.
func (o *AiAgentSuggestion) GetDisabledReasonCode() string {
	if o == nil || o.DisabledReasonCode == nil {
		var ret string
		return ret
	}
	return *o.DisabledReasonCode
}

// GetDisabledReasonCodeOk returns a tuple with the DisabledReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentSuggestion) GetDisabledReasonCodeOk() (*string, bool) {
	if o == nil || o.DisabledReasonCode == nil {
		return nil, false
	}
	return o.DisabledReasonCode, true
}

// HasDisabledReasonCode returns a boolean if a field has been set.
func (o *AiAgentSuggestion) HasDisabledReasonCode() bool {
	return o != nil && o.DisabledReasonCode != nil
}

// SetDisabledReasonCode gets a reference to the given string and assigns it to the DisabledReasonCode field.
func (o *AiAgentSuggestion) SetDisabledReasonCode(v string) {
	o.DisabledReasonCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentSuggestion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["suggestionId"] = o.SuggestionId
	toSerialize["label"] = o.Label
	toSerialize["prompt"] = o.Prompt
	if o.PlaybookId != nil {
		toSerialize["playbookId"] = o.PlaybookId
	}
	if o.ScopePatch != nil {
		toSerialize["scopePatch"] = o.ScopePatch
	}
	if o.DisabledReason != nil {
		toSerialize["disabledReason"] = o.DisabledReason
	}
	if o.DisabledReasonCode != nil {
		toSerialize["disabledReasonCode"] = o.DisabledReasonCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentSuggestion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SuggestionId       *string               `json:"suggestionId"`
		Label              *string               `json:"label"`
		Prompt             *string               `json:"prompt"`
		PlaybookId         *string               `json:"playbookId,omitempty"`
		ScopePatch         *AiAgentResourceScope `json:"scopePatch,omitempty"`
		DisabledReason     *AiAgentStatusDisplay `json:"disabledReason,omitempty"`
		DisabledReasonCode *string               `json:"disabledReasonCode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SuggestionId == nil {
		return fmt.Errorf("required field suggestionId missing")
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Prompt == nil {
		return fmt.Errorf("required field prompt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"suggestionId", "label", "prompt", "playbookId", "scopePatch", "disabledReason", "disabledReasonCode"})
	} else {
		return err
	}

	hasInvalidField := false
	o.SuggestionId = *all.SuggestionId
	o.Label = *all.Label
	o.Prompt = *all.Prompt
	o.PlaybookId = all.PlaybookId
	if all.ScopePatch != nil && all.ScopePatch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ScopePatch = all.ScopePatch
	if all.DisabledReason != nil && all.DisabledReason.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DisabledReason = all.DisabledReason
	o.DisabledReasonCode = all.DisabledReasonCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
