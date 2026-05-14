// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentStage struct {
	StageId     *string `json:"stageId,omitempty"`
	Label       *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	State *AiAgentStatusDisplay `json:"state,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentStage instantiates a new AiAgentStage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentStage() *AiAgentStage {
	this := AiAgentStage{}
	return &this
}

// NewAiAgentStageWithDefaults instantiates a new AiAgentStage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentStageWithDefaults() *AiAgentStage {
	this := AiAgentStage{}
	return &this
}

// GetStageId returns the StageId field value if set, zero value otherwise.
func (o *AiAgentStage) GetStageId() string {
	if o == nil || o.StageId == nil {
		var ret string
		return ret
	}
	return *o.StageId
}

// GetStageIdOk returns a tuple with the StageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentStage) GetStageIdOk() (*string, bool) {
	if o == nil || o.StageId == nil {
		return nil, false
	}
	return o.StageId, true
}

// HasStageId returns a boolean if a field has been set.
func (o *AiAgentStage) HasStageId() bool {
	return o != nil && o.StageId != nil
}

// SetStageId gets a reference to the given string and assigns it to the StageId field.
func (o *AiAgentStage) SetStageId(v string) {
	o.StageId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AiAgentStage) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentStage) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AiAgentStage) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AiAgentStage) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AiAgentStage) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentStage) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AiAgentStage) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AiAgentStage) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AiAgentStage) GetState() AiAgentStatusDisplay {
	if o == nil || o.State == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentStage) GetStateOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AiAgentStage) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given AiAgentStatusDisplay and assigns it to the State field.
func (o *AiAgentStage) SetState(v AiAgentStatusDisplay) {
	o.State = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentStage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.StageId != nil {
		toSerialize["stageId"] = o.StageId
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentStage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StageId     *string               `json:"stageId,omitempty"`
		Label       *string               `json:"label,omitempty"`
		Description *string               `json:"description,omitempty"`
		State       *AiAgentStatusDisplay `json:"state,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"stageId", "label", "description", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.StageId = all.StageId
	o.Label = all.Label
	o.Description = all.Description
	if all.State != nil && all.State.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.State = all.State

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
