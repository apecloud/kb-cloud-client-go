// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentExecutionStep struct {
	Label string `json:"label"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	Status AiAgentStatusDisplay `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentExecutionStep instantiates a new AiAgentExecutionStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentExecutionStep(label string, status AiAgentStatusDisplay) *AiAgentExecutionStep {
	this := AiAgentExecutionStep{}
	this.Label = label
	this.Status = status
	return &this
}

// NewAiAgentExecutionStepWithDefaults instantiates a new AiAgentExecutionStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentExecutionStepWithDefaults() *AiAgentExecutionStep {
	this := AiAgentExecutionStep{}
	return &this
}

// GetLabel returns the Label field value.
func (o *AiAgentExecutionStep) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AiAgentExecutionStep) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *AiAgentExecutionStep) SetLabel(v string) {
	o.Label = v
}

// GetStatus returns the Status field value.
func (o *AiAgentExecutionStep) GetStatus() AiAgentStatusDisplay {
	if o == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentExecutionStep) GetStatusOk() (*AiAgentStatusDisplay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentExecutionStep) SetStatus(v AiAgentStatusDisplay) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentExecutionStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["label"] = o.Label
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentExecutionStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label  *string               `json:"label"`
		Status *AiAgentStatusDisplay `json:"status"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"label", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Label = *all.Label
	if all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
