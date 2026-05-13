// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentRecommendation struct {
	RecommendationId string                    `json:"recommendationId"`
	Type             AiAgentRecommendationType `json:"type"`
	Title            string                    `json:"title"`
	Description      *string                   `json:"description,omitempty"`
	ActionPlanId     *string                   `json:"actionPlanId,omitempty"`
	RiskLevel        *string                   `json:"riskLevel,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentRecommendation instantiates a new AiAgentRecommendation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentRecommendation(recommendationId string, typeVar AiAgentRecommendationType, title string) *AiAgentRecommendation {
	this := AiAgentRecommendation{}
	this.RecommendationId = recommendationId
	this.Type = typeVar
	this.Title = title
	return &this
}

// NewAiAgentRecommendationWithDefaults instantiates a new AiAgentRecommendation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentRecommendationWithDefaults() *AiAgentRecommendation {
	this := AiAgentRecommendation{}
	return &this
}

// GetRecommendationId returns the RecommendationId field value.
func (o *AiAgentRecommendation) GetRecommendationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RecommendationId
}

// GetRecommendationIdOk returns a tuple with the RecommendationId field value
// and a boolean to check if the value has been set.
func (o *AiAgentRecommendation) GetRecommendationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationId, true
}

// SetRecommendationId sets field value.
func (o *AiAgentRecommendation) SetRecommendationId(v string) {
	o.RecommendationId = v
}

// GetType returns the Type field value.
func (o *AiAgentRecommendation) GetType() AiAgentRecommendationType {
	if o == nil {
		var ret AiAgentRecommendationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AiAgentRecommendation) GetTypeOk() (*AiAgentRecommendationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AiAgentRecommendation) SetType(v AiAgentRecommendationType) {
	o.Type = v
}

// GetTitle returns the Title field value.
func (o *AiAgentRecommendation) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AiAgentRecommendation) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *AiAgentRecommendation) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AiAgentRecommendation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRecommendation) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AiAgentRecommendation) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AiAgentRecommendation) SetDescription(v string) {
	o.Description = &v
}

// GetActionPlanId returns the ActionPlanId field value if set, zero value otherwise.
func (o *AiAgentRecommendation) GetActionPlanId() string {
	if o == nil || o.ActionPlanId == nil {
		var ret string
		return ret
	}
	return *o.ActionPlanId
}

// GetActionPlanIdOk returns a tuple with the ActionPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRecommendation) GetActionPlanIdOk() (*string, bool) {
	if o == nil || o.ActionPlanId == nil {
		return nil, false
	}
	return o.ActionPlanId, true
}

// HasActionPlanId returns a boolean if a field has been set.
func (o *AiAgentRecommendation) HasActionPlanId() bool {
	return o != nil && o.ActionPlanId != nil
}

// SetActionPlanId gets a reference to the given string and assigns it to the ActionPlanId field.
func (o *AiAgentRecommendation) SetActionPlanId(v string) {
	o.ActionPlanId = &v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *AiAgentRecommendation) GetRiskLevel() string {
	if o == nil || o.RiskLevel == nil {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentRecommendation) GetRiskLevelOk() (*string, bool) {
	if o == nil || o.RiskLevel == nil {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *AiAgentRecommendation) HasRiskLevel() bool {
	return o != nil && o.RiskLevel != nil
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *AiAgentRecommendation) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentRecommendation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["recommendationId"] = o.RecommendationId
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ActionPlanId != nil {
		toSerialize["actionPlanId"] = o.ActionPlanId
	}
	if o.RiskLevel != nil {
		toSerialize["riskLevel"] = o.RiskLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentRecommendation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RecommendationId *string                    `json:"recommendationId"`
		Type             *AiAgentRecommendationType `json:"type"`
		Title            *string                    `json:"title"`
		Description      *string                    `json:"description,omitempty"`
		ActionPlanId     *string                    `json:"actionPlanId,omitempty"`
		RiskLevel        *string                    `json:"riskLevel,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RecommendationId == nil {
		return fmt.Errorf("required field recommendationId missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"recommendationId", "type", "title", "description", "actionPlanId", "riskLevel"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RecommendationId = *all.RecommendationId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Title = *all.Title
	o.Description = all.Description
	o.ActionPlanId = all.ActionPlanId
	o.RiskLevel = all.RiskLevel

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
