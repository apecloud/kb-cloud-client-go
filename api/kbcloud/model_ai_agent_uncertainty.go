// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentUncertainty struct {
	RequirementId *string `json:"requirementId,omitempty"`
	Reason        *string `json:"reason,omitempty"`
	DisplayText   *string `json:"displayText,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentUncertainty instantiates a new AiAgentUncertainty object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentUncertainty() *AiAgentUncertainty {
	this := AiAgentUncertainty{}
	return &this
}

// NewAiAgentUncertaintyWithDefaults instantiates a new AiAgentUncertainty object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentUncertaintyWithDefaults() *AiAgentUncertainty {
	this := AiAgentUncertainty{}
	return &this
}

// GetRequirementId returns the RequirementId field value if set, zero value otherwise.
func (o *AiAgentUncertainty) GetRequirementId() string {
	if o == nil || o.RequirementId == nil {
		var ret string
		return ret
	}
	return *o.RequirementId
}

// GetRequirementIdOk returns a tuple with the RequirementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentUncertainty) GetRequirementIdOk() (*string, bool) {
	if o == nil || o.RequirementId == nil {
		return nil, false
	}
	return o.RequirementId, true
}

// HasRequirementId returns a boolean if a field has been set.
func (o *AiAgentUncertainty) HasRequirementId() bool {
	return o != nil && o.RequirementId != nil
}

// SetRequirementId gets a reference to the given string and assigns it to the RequirementId field.
func (o *AiAgentUncertainty) SetRequirementId(v string) {
	o.RequirementId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AiAgentUncertainty) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentUncertainty) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AiAgentUncertainty) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AiAgentUncertainty) SetReason(v string) {
	o.Reason = &v
}

// GetDisplayText returns the DisplayText field value if set, zero value otherwise.
func (o *AiAgentUncertainty) GetDisplayText() string {
	if o == nil || o.DisplayText == nil {
		var ret string
		return ret
	}
	return *o.DisplayText
}

// GetDisplayTextOk returns a tuple with the DisplayText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentUncertainty) GetDisplayTextOk() (*string, bool) {
	if o == nil || o.DisplayText == nil {
		return nil, false
	}
	return o.DisplayText, true
}

// HasDisplayText returns a boolean if a field has been set.
func (o *AiAgentUncertainty) HasDisplayText() bool {
	return o != nil && o.DisplayText != nil
}

// SetDisplayText gets a reference to the given string and assigns it to the DisplayText field.
func (o *AiAgentUncertainty) SetDisplayText(v string) {
	o.DisplayText = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentUncertainty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.RequirementId != nil {
		toSerialize["requirementId"] = o.RequirementId
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.DisplayText != nil {
		toSerialize["displayText"] = o.DisplayText
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentUncertainty) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RequirementId *string `json:"requirementId,omitempty"`
		Reason        *string `json:"reason,omitempty"`
		DisplayText   *string `json:"displayText,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"requirementId", "reason", "displayText"})
	} else {
		return err
	}
	o.RequirementId = all.RequirementId
	o.Reason = all.Reason
	o.DisplayText = all.DisplayText

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
