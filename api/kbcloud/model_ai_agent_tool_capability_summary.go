// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentToolCapabilitySummary struct {
	AgentToolId string            `json:"agentToolId"`
	Title       string            `json:"title"`
	Enabled     bool              `json:"enabled"`
	Visibility  *string           `json:"visibility,omitempty"`
	CostClass   *AiAgentCostClass `json:"costClass,omitempty"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	UnavailableReason *AiAgentStatusDisplay `json:"unavailableReason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentToolCapabilitySummary instantiates a new AiAgentToolCapabilitySummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentToolCapabilitySummary(agentToolId string, title string, enabled bool) *AiAgentToolCapabilitySummary {
	this := AiAgentToolCapabilitySummary{}
	this.AgentToolId = agentToolId
	this.Title = title
	this.Enabled = enabled
	return &this
}

// NewAiAgentToolCapabilitySummaryWithDefaults instantiates a new AiAgentToolCapabilitySummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentToolCapabilitySummaryWithDefaults() *AiAgentToolCapabilitySummary {
	this := AiAgentToolCapabilitySummary{}
	return &this
}

// GetAgentToolId returns the AgentToolId field value.
func (o *AiAgentToolCapabilitySummary) GetAgentToolId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AgentToolId
}

// GetAgentToolIdOk returns a tuple with the AgentToolId field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolCapabilitySummary) GetAgentToolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentToolId, true
}

// SetAgentToolId sets field value.
func (o *AiAgentToolCapabilitySummary) SetAgentToolId(v string) {
	o.AgentToolId = v
}

// GetTitle returns the Title field value.
func (o *AiAgentToolCapabilitySummary) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolCapabilitySummary) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *AiAgentToolCapabilitySummary) SetTitle(v string) {
	o.Title = v
}

// GetEnabled returns the Enabled field value.
func (o *AiAgentToolCapabilitySummary) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AiAgentToolCapabilitySummary) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *AiAgentToolCapabilitySummary) SetEnabled(v bool) {
	o.Enabled = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *AiAgentToolCapabilitySummary) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolCapabilitySummary) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *AiAgentToolCapabilitySummary) HasVisibility() bool {
	return o != nil && o.Visibility != nil
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *AiAgentToolCapabilitySummary) SetVisibility(v string) {
	o.Visibility = &v
}

// GetCostClass returns the CostClass field value if set, zero value otherwise.
func (o *AiAgentToolCapabilitySummary) GetCostClass() AiAgentCostClass {
	if o == nil || o.CostClass == nil {
		var ret AiAgentCostClass
		return ret
	}
	return *o.CostClass
}

// GetCostClassOk returns a tuple with the CostClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolCapabilitySummary) GetCostClassOk() (*AiAgentCostClass, bool) {
	if o == nil || o.CostClass == nil {
		return nil, false
	}
	return o.CostClass, true
}

// HasCostClass returns a boolean if a field has been set.
func (o *AiAgentToolCapabilitySummary) HasCostClass() bool {
	return o != nil && o.CostClass != nil
}

// SetCostClass gets a reference to the given AiAgentCostClass and assigns it to the CostClass field.
func (o *AiAgentToolCapabilitySummary) SetCostClass(v AiAgentCostClass) {
	o.CostClass = &v
}

// GetUnavailableReason returns the UnavailableReason field value if set, zero value otherwise.
func (o *AiAgentToolCapabilitySummary) GetUnavailableReason() AiAgentStatusDisplay {
	if o == nil || o.UnavailableReason == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.UnavailableReason
}

// GetUnavailableReasonOk returns a tuple with the UnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentToolCapabilitySummary) GetUnavailableReasonOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.UnavailableReason == nil {
		return nil, false
	}
	return o.UnavailableReason, true
}

// HasUnavailableReason returns a boolean if a field has been set.
func (o *AiAgentToolCapabilitySummary) HasUnavailableReason() bool {
	return o != nil && o.UnavailableReason != nil
}

// SetUnavailableReason gets a reference to the given AiAgentStatusDisplay and assigns it to the UnavailableReason field.
func (o *AiAgentToolCapabilitySummary) SetUnavailableReason(v AiAgentStatusDisplay) {
	o.UnavailableReason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentToolCapabilitySummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["agentToolId"] = o.AgentToolId
	toSerialize["title"] = o.Title
	toSerialize["enabled"] = o.Enabled
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.CostClass != nil {
		toSerialize["costClass"] = o.CostClass
	}
	if o.UnavailableReason != nil {
		toSerialize["unavailableReason"] = o.UnavailableReason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentToolCapabilitySummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentToolId       *string               `json:"agentToolId"`
		Title             *string               `json:"title"`
		Enabled           *bool                 `json:"enabled"`
		Visibility        *string               `json:"visibility,omitempty"`
		CostClass         *AiAgentCostClass     `json:"costClass,omitempty"`
		UnavailableReason *AiAgentStatusDisplay `json:"unavailableReason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.AgentToolId == nil {
		return fmt.Errorf("required field agentToolId missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"agentToolId", "title", "enabled", "visibility", "costClass", "unavailableReason"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AgentToolId = *all.AgentToolId
	o.Title = *all.Title
	o.Enabled = *all.Enabled
	o.Visibility = all.Visibility
	if all.CostClass != nil && !all.CostClass.IsValid() {
		hasInvalidField = true
	} else {
		o.CostClass = all.CostClass
	}
	if all.UnavailableReason != nil && all.UnavailableReason.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UnavailableReason = all.UnavailableReason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
