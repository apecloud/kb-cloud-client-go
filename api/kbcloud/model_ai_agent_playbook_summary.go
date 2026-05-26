// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentPlaybookSummary struct {
	PlaybookId           string   `json:"playbookId"`
	Title                string   `json:"title"`
	Description          *string  `json:"description,omitempty"`
	Enabled              bool     `json:"enabled"`
	RequiredCapabilities []string `json:"requiredCapabilities,omitempty"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	DisabledReason     *AiAgentStatusDisplay `json:"disabledReason,omitempty"`
	DisabledReasonCode *string               `json:"disabledReasonCode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentPlaybookSummary instantiates a new AiAgentPlaybookSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentPlaybookSummary(playbookId string, title string, enabled bool) *AiAgentPlaybookSummary {
	this := AiAgentPlaybookSummary{}
	this.PlaybookId = playbookId
	this.Title = title
	this.Enabled = enabled
	return &this
}

// NewAiAgentPlaybookSummaryWithDefaults instantiates a new AiAgentPlaybookSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentPlaybookSummaryWithDefaults() *AiAgentPlaybookSummary {
	this := AiAgentPlaybookSummary{}
	return &this
}

// GetPlaybookId returns the PlaybookId field value.
func (o *AiAgentPlaybookSummary) GetPlaybookId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PlaybookId
}

// GetPlaybookIdOk returns a tuple with the PlaybookId field value
// and a boolean to check if the value has been set.
func (o *AiAgentPlaybookSummary) GetPlaybookIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlaybookId, true
}

// SetPlaybookId sets field value.
func (o *AiAgentPlaybookSummary) SetPlaybookId(v string) {
	o.PlaybookId = v
}

// GetTitle returns the Title field value.
func (o *AiAgentPlaybookSummary) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiAgentPlaybookSummary) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *AiAgentPlaybookSummary) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AiAgentPlaybookSummary) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentPlaybookSummary) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AiAgentPlaybookSummary) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AiAgentPlaybookSummary) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value.
func (o *AiAgentPlaybookSummary) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AiAgentPlaybookSummary) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *AiAgentPlaybookSummary) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRequiredCapabilities returns the RequiredCapabilities field value if set, zero value otherwise.
func (o *AiAgentPlaybookSummary) GetRequiredCapabilities() []string {
	if o == nil || o.RequiredCapabilities == nil {
		var ret []string
		return ret
	}
	return o.RequiredCapabilities
}

// GetRequiredCapabilitiesOk returns a tuple with the RequiredCapabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentPlaybookSummary) GetRequiredCapabilitiesOk() (*[]string, bool) {
	if o == nil || o.RequiredCapabilities == nil {
		return nil, false
	}
	return &o.RequiredCapabilities, true
}

// HasRequiredCapabilities returns a boolean if a field has been set.
func (o *AiAgentPlaybookSummary) HasRequiredCapabilities() bool {
	return o != nil && o.RequiredCapabilities != nil
}

// SetRequiredCapabilities gets a reference to the given []string and assigns it to the RequiredCapabilities field.
func (o *AiAgentPlaybookSummary) SetRequiredCapabilities(v []string) {
	o.RequiredCapabilities = v
}

// GetDisabledReason returns the DisabledReason field value if set, zero value otherwise.
func (o *AiAgentPlaybookSummary) GetDisabledReason() AiAgentStatusDisplay {
	if o == nil || o.DisabledReason == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.DisabledReason
}

// GetDisabledReasonOk returns a tuple with the DisabledReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentPlaybookSummary) GetDisabledReasonOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.DisabledReason == nil {
		return nil, false
	}
	return o.DisabledReason, true
}

// HasDisabledReason returns a boolean if a field has been set.
func (o *AiAgentPlaybookSummary) HasDisabledReason() bool {
	return o != nil && o.DisabledReason != nil
}

// SetDisabledReason gets a reference to the given AiAgentStatusDisplay and assigns it to the DisabledReason field.
func (o *AiAgentPlaybookSummary) SetDisabledReason(v AiAgentStatusDisplay) {
	o.DisabledReason = &v
}

// GetDisabledReasonCode returns the DisabledReasonCode field value if set, zero value otherwise.
func (o *AiAgentPlaybookSummary) GetDisabledReasonCode() string {
	if o == nil || o.DisabledReasonCode == nil {
		var ret string
		return ret
	}
	return *o.DisabledReasonCode
}

// GetDisabledReasonCodeOk returns a tuple with the DisabledReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentPlaybookSummary) GetDisabledReasonCodeOk() (*string, bool) {
	if o == nil || o.DisabledReasonCode == nil {
		return nil, false
	}
	return o.DisabledReasonCode, true
}

// HasDisabledReasonCode returns a boolean if a field has been set.
func (o *AiAgentPlaybookSummary) HasDisabledReasonCode() bool {
	return o != nil && o.DisabledReasonCode != nil
}

// SetDisabledReasonCode gets a reference to the given string and assigns it to the DisabledReasonCode field.
func (o *AiAgentPlaybookSummary) SetDisabledReasonCode(v string) {
	o.DisabledReasonCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentPlaybookSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["playbookId"] = o.PlaybookId
	toSerialize["title"] = o.Title
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if o.RequiredCapabilities != nil {
		toSerialize["requiredCapabilities"] = o.RequiredCapabilities
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
func (o *AiAgentPlaybookSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PlaybookId           *string               `json:"playbookId"`
		Title                *string               `json:"title"`
		Description          *string               `json:"description,omitempty"`
		Enabled              *bool                 `json:"enabled"`
		RequiredCapabilities []string              `json:"requiredCapabilities,omitempty"`
		DisabledReason       *AiAgentStatusDisplay `json:"disabledReason,omitempty"`
		DisabledReasonCode   *string               `json:"disabledReasonCode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.PlaybookId == nil {
		return fmt.Errorf("required field playbookId missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"playbookId", "title", "description", "enabled", "requiredCapabilities", "disabledReason", "disabledReasonCode"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PlaybookId = *all.PlaybookId
	o.Title = *all.Title
	o.Description = all.Description
	o.Enabled = *all.Enabled
	o.RequiredCapabilities = all.RequiredCapabilities
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
