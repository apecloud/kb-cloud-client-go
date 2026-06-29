// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterSlowLogDiagnosis Primary rule-based diagnosis card for a slow log EXPLAIN result
type ClusterSlowLogDiagnosis struct {
	// Diagnosis engine. Current value is rule_based.
	Engine *string `json:"engine,omitempty"`
	// Localized slow log diagnosis text with dynamic arguments
	Summary *ClusterSlowLogLocalizedText `json:"summary,omitempty"`
	// A single detected slow log diagnosis issue
	PrimaryIssue *ClusterSlowLogDiagnosisIssue `json:"primaryIssue,omitempty"`
	// Localized slow log diagnosis text with dynamic arguments
	OptimizationStrategy *ClusterSlowLogLocalizedText `json:"optimizationStrategy,omitempty"`
	// A rule-based recommendation for a slow log diagnosis
	RecommendedAction *ClusterSlowLogRecommendation `json:"recommendedAction,omitempty"`
	// Expected optimization benefit. Current values are unknown, low, medium, and high.
	ExpectedBenefit *string `json:"expectedBenefit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterSlowLogDiagnosis instantiates a new ClusterSlowLogDiagnosis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterSlowLogDiagnosis() *ClusterSlowLogDiagnosis {
	this := ClusterSlowLogDiagnosis{}
	return &this
}

// NewClusterSlowLogDiagnosisWithDefaults instantiates a new ClusterSlowLogDiagnosis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterSlowLogDiagnosisWithDefaults() *ClusterSlowLogDiagnosis {
	this := ClusterSlowLogDiagnosis{}
	return &this
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosis) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosis) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosis) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *ClusterSlowLogDiagnosis) SetEngine(v string) {
	o.Engine = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosis) GetSummary() ClusterSlowLogLocalizedText {
	if o == nil || o.Summary == nil {
		var ret ClusterSlowLogLocalizedText
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosis) GetSummaryOk() (*ClusterSlowLogLocalizedText, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosis) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given ClusterSlowLogLocalizedText and assigns it to the Summary field.
func (o *ClusterSlowLogDiagnosis) SetSummary(v ClusterSlowLogLocalizedText) {
	o.Summary = &v
}

// GetPrimaryIssue returns the PrimaryIssue field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosis) GetPrimaryIssue() ClusterSlowLogDiagnosisIssue {
	if o == nil || o.PrimaryIssue == nil {
		var ret ClusterSlowLogDiagnosisIssue
		return ret
	}
	return *o.PrimaryIssue
}

// GetPrimaryIssueOk returns a tuple with the PrimaryIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosis) GetPrimaryIssueOk() (*ClusterSlowLogDiagnosisIssue, bool) {
	if o == nil || o.PrimaryIssue == nil {
		return nil, false
	}
	return o.PrimaryIssue, true
}

// HasPrimaryIssue returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosis) HasPrimaryIssue() bool {
	return o != nil && o.PrimaryIssue != nil
}

// SetPrimaryIssue gets a reference to the given ClusterSlowLogDiagnosisIssue and assigns it to the PrimaryIssue field.
func (o *ClusterSlowLogDiagnosis) SetPrimaryIssue(v ClusterSlowLogDiagnosisIssue) {
	o.PrimaryIssue = &v
}

// GetOptimizationStrategy returns the OptimizationStrategy field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosis) GetOptimizationStrategy() ClusterSlowLogLocalizedText {
	if o == nil || o.OptimizationStrategy == nil {
		var ret ClusterSlowLogLocalizedText
		return ret
	}
	return *o.OptimizationStrategy
}

// GetOptimizationStrategyOk returns a tuple with the OptimizationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosis) GetOptimizationStrategyOk() (*ClusterSlowLogLocalizedText, bool) {
	if o == nil || o.OptimizationStrategy == nil {
		return nil, false
	}
	return o.OptimizationStrategy, true
}

// HasOptimizationStrategy returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosis) HasOptimizationStrategy() bool {
	return o != nil && o.OptimizationStrategy != nil
}

// SetOptimizationStrategy gets a reference to the given ClusterSlowLogLocalizedText and assigns it to the OptimizationStrategy field.
func (o *ClusterSlowLogDiagnosis) SetOptimizationStrategy(v ClusterSlowLogLocalizedText) {
	o.OptimizationStrategy = &v
}

// GetRecommendedAction returns the RecommendedAction field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosis) GetRecommendedAction() ClusterSlowLogRecommendation {
	if o == nil || o.RecommendedAction == nil {
		var ret ClusterSlowLogRecommendation
		return ret
	}
	return *o.RecommendedAction
}

// GetRecommendedActionOk returns a tuple with the RecommendedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosis) GetRecommendedActionOk() (*ClusterSlowLogRecommendation, bool) {
	if o == nil || o.RecommendedAction == nil {
		return nil, false
	}
	return o.RecommendedAction, true
}

// HasRecommendedAction returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosis) HasRecommendedAction() bool {
	return o != nil && o.RecommendedAction != nil
}

// SetRecommendedAction gets a reference to the given ClusterSlowLogRecommendation and assigns it to the RecommendedAction field.
func (o *ClusterSlowLogDiagnosis) SetRecommendedAction(v ClusterSlowLogRecommendation) {
	o.RecommendedAction = &v
}

// GetExpectedBenefit returns the ExpectedBenefit field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosis) GetExpectedBenefit() string {
	if o == nil || o.ExpectedBenefit == nil {
		var ret string
		return ret
	}
	return *o.ExpectedBenefit
}

// GetExpectedBenefitOk returns a tuple with the ExpectedBenefit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosis) GetExpectedBenefitOk() (*string, bool) {
	if o == nil || o.ExpectedBenefit == nil {
		return nil, false
	}
	return o.ExpectedBenefit, true
}

// HasExpectedBenefit returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosis) HasExpectedBenefit() bool {
	return o != nil && o.ExpectedBenefit != nil
}

// SetExpectedBenefit gets a reference to the given string and assigns it to the ExpectedBenefit field.
func (o *ClusterSlowLogDiagnosis) SetExpectedBenefit(v string) {
	o.ExpectedBenefit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterSlowLogDiagnosis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.PrimaryIssue != nil {
		toSerialize["primaryIssue"] = o.PrimaryIssue
	}
	if o.OptimizationStrategy != nil {
		toSerialize["optimizationStrategy"] = o.OptimizationStrategy
	}
	if o.RecommendedAction != nil {
		toSerialize["recommendedAction"] = o.RecommendedAction
	}
	if o.ExpectedBenefit != nil {
		toSerialize["expectedBenefit"] = o.ExpectedBenefit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterSlowLogDiagnosis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Engine               *string                       `json:"engine,omitempty"`
		Summary              *ClusterSlowLogLocalizedText  `json:"summary,omitempty"`
		PrimaryIssue         *ClusterSlowLogDiagnosisIssue `json:"primaryIssue,omitempty"`
		OptimizationStrategy *ClusterSlowLogLocalizedText  `json:"optimizationStrategy,omitempty"`
		RecommendedAction    *ClusterSlowLogRecommendation `json:"recommendedAction,omitempty"`
		ExpectedBenefit      *string                       `json:"expectedBenefit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engine", "summary", "primaryIssue", "optimizationStrategy", "recommendedAction", "expectedBenefit"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Engine = all.Engine
	if all.Summary != nil && all.Summary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Summary = all.Summary
	if all.PrimaryIssue != nil && all.PrimaryIssue.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PrimaryIssue = all.PrimaryIssue
	if all.OptimizationStrategy != nil && all.OptimizationStrategy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OptimizationStrategy = all.OptimizationStrategy
	if all.RecommendedAction != nil && all.RecommendedAction.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RecommendedAction = all.RecommendedAction
	o.ExpectedBenefit = all.ExpectedBenefit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
