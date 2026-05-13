// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEvidenceSummary struct {
	EvidenceId    string                `json:"evidenceId"`
	RequirementId string                `json:"requirementId"`
	Type          *string               `json:"type,omitempty"`
	Title         *string               `json:"title,omitempty"`
	Source        *string               `json:"source,omitempty"`
	ObjectRef     *AiAgentObjectRef     `json:"objectRef,omitempty"`
	Status        AiAgentEvidenceStatus `json:"status"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	StatusDisplay *AiAgentStatusDisplay      `json:"statusDisplay,omitempty"`
	Visibility    *AiAgentEvidenceVisibility `json:"visibility,omitempty"`
	MissingReason *AiAgentMissingReason      `json:"missingReason,omitempty"`
	PartialReason *AiAgentPartialReason      `json:"partialReason,omitempty"`
	Summary       string                     `json:"summary"`
	TimeRange     *AiAgentTimeRange          `json:"timeRange,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentEvidenceSummary instantiates a new AiAgentEvidenceSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentEvidenceSummary(evidenceId string, requirementId string, status AiAgentEvidenceStatus, summary string) *AiAgentEvidenceSummary {
	this := AiAgentEvidenceSummary{}
	this.EvidenceId = evidenceId
	this.RequirementId = requirementId
	this.Status = status
	this.Summary = summary
	return &this
}

// NewAiAgentEvidenceSummaryWithDefaults instantiates a new AiAgentEvidenceSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentEvidenceSummaryWithDefaults() *AiAgentEvidenceSummary {
	this := AiAgentEvidenceSummary{}
	return &this
}

// GetEvidenceId returns the EvidenceId field value.
func (o *AiAgentEvidenceSummary) GetEvidenceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvidenceId
}

// GetEvidenceIdOk returns a tuple with the EvidenceId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetEvidenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvidenceId, true
}

// SetEvidenceId sets field value.
func (o *AiAgentEvidenceSummary) SetEvidenceId(v string) {
	o.EvidenceId = v
}

// GetRequirementId returns the RequirementId field value.
func (o *AiAgentEvidenceSummary) GetRequirementId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RequirementId
}

// GetRequirementIdOk returns a tuple with the RequirementId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetRequirementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequirementId, true
}

// SetRequirementId sets field value.
func (o *AiAgentEvidenceSummary) SetRequirementId(v string) {
	o.RequirementId = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AiAgentEvidenceSummary) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AiAgentEvidenceSummary) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AiAgentEvidenceSummary) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiAgentEvidenceSummary) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiAgentEvidenceSummary) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiAgentEvidenceSummary) SetTitle(v string) {
	o.Title = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AiAgentEvidenceSummary) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AiAgentEvidenceSummary) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AiAgentEvidenceSummary) SetSource(v string) {
	o.Source = &v
}

// GetObjectRef returns the ObjectRef field value if set, zero value otherwise.
func (o *AiAgentEvidenceSummary) GetObjectRef() AiAgentObjectRef {
	if o == nil || o.ObjectRef == nil {
		var ret AiAgentObjectRef
		return ret
	}
	return *o.ObjectRef
}

// GetObjectRefOk returns a tuple with the ObjectRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetObjectRefOk() (*AiAgentObjectRef, bool) {
	if o == nil || o.ObjectRef == nil {
		return nil, false
	}
	return o.ObjectRef, true
}

// HasObjectRef returns a boolean if a field has been set.
func (o *AiAgentEvidenceSummary) HasObjectRef() bool {
	return o != nil && o.ObjectRef != nil
}

// SetObjectRef gets a reference to the given AiAgentObjectRef and assigns it to the ObjectRef field.
func (o *AiAgentEvidenceSummary) SetObjectRef(v AiAgentObjectRef) {
	o.ObjectRef = &v
}

// GetStatus returns the Status field value.
func (o *AiAgentEvidenceSummary) GetStatus() AiAgentEvidenceStatus {
	if o == nil {
		var ret AiAgentEvidenceStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetStatusOk() (*AiAgentEvidenceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentEvidenceSummary) SetStatus(v AiAgentEvidenceStatus) {
	o.Status = v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *AiAgentEvidenceSummary) GetStatusDisplay() AiAgentStatusDisplay {
	if o == nil || o.StatusDisplay == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetStatusDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *AiAgentEvidenceSummary) HasStatusDisplay() bool {
	return o != nil && o.StatusDisplay != nil
}

// SetStatusDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the StatusDisplay field.
func (o *AiAgentEvidenceSummary) SetStatusDisplay(v AiAgentStatusDisplay) {
	o.StatusDisplay = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *AiAgentEvidenceSummary) GetVisibility() AiAgentEvidenceVisibility {
	if o == nil || o.Visibility == nil {
		var ret AiAgentEvidenceVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetVisibilityOk() (*AiAgentEvidenceVisibility, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *AiAgentEvidenceSummary) HasVisibility() bool {
	return o != nil && o.Visibility != nil
}

// SetVisibility gets a reference to the given AiAgentEvidenceVisibility and assigns it to the Visibility field.
func (o *AiAgentEvidenceSummary) SetVisibility(v AiAgentEvidenceVisibility) {
	o.Visibility = &v
}

// GetMissingReason returns the MissingReason field value if set, zero value otherwise.
func (o *AiAgentEvidenceSummary) GetMissingReason() AiAgentMissingReason {
	if o == nil || o.MissingReason == nil {
		var ret AiAgentMissingReason
		return ret
	}
	return *o.MissingReason
}

// GetMissingReasonOk returns a tuple with the MissingReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetMissingReasonOk() (*AiAgentMissingReason, bool) {
	if o == nil || o.MissingReason == nil {
		return nil, false
	}
	return o.MissingReason, true
}

// HasMissingReason returns a boolean if a field has been set.
func (o *AiAgentEvidenceSummary) HasMissingReason() bool {
	return o != nil && o.MissingReason != nil
}

// SetMissingReason gets a reference to the given AiAgentMissingReason and assigns it to the MissingReason field.
func (o *AiAgentEvidenceSummary) SetMissingReason(v AiAgentMissingReason) {
	o.MissingReason = &v
}

// GetPartialReason returns the PartialReason field value if set, zero value otherwise.
func (o *AiAgentEvidenceSummary) GetPartialReason() AiAgentPartialReason {
	if o == nil || o.PartialReason == nil {
		var ret AiAgentPartialReason
		return ret
	}
	return *o.PartialReason
}

// GetPartialReasonOk returns a tuple with the PartialReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetPartialReasonOk() (*AiAgentPartialReason, bool) {
	if o == nil || o.PartialReason == nil {
		return nil, false
	}
	return o.PartialReason, true
}

// HasPartialReason returns a boolean if a field has been set.
func (o *AiAgentEvidenceSummary) HasPartialReason() bool {
	return o != nil && o.PartialReason != nil
}

// SetPartialReason gets a reference to the given AiAgentPartialReason and assigns it to the PartialReason field.
func (o *AiAgentEvidenceSummary) SetPartialReason(v AiAgentPartialReason) {
	o.PartialReason = &v
}

// GetSummary returns the Summary field value.
func (o *AiAgentEvidenceSummary) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *AiAgentEvidenceSummary) SetSummary(v string) {
	o.Summary = v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *AiAgentEvidenceSummary) GetTimeRange() AiAgentTimeRange {
	if o == nil || o.TimeRange == nil {
		var ret AiAgentTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceSummary) GetTimeRangeOk() (*AiAgentTimeRange, bool) {
	if o == nil || o.TimeRange == nil {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *AiAgentEvidenceSummary) HasTimeRange() bool {
	return o != nil && o.TimeRange != nil
}

// SetTimeRange gets a reference to the given AiAgentTimeRange and assigns it to the TimeRange field.
func (o *AiAgentEvidenceSummary) SetTimeRange(v AiAgentTimeRange) {
	o.TimeRange = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentEvidenceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["evidenceId"] = o.EvidenceId
	toSerialize["requirementId"] = o.RequirementId
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.ObjectRef != nil {
		toSerialize["objectRef"] = o.ObjectRef
	}
	toSerialize["status"] = o.Status
	if o.StatusDisplay != nil {
		toSerialize["statusDisplay"] = o.StatusDisplay
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.MissingReason != nil {
		toSerialize["missingReason"] = o.MissingReason
	}
	if o.PartialReason != nil {
		toSerialize["partialReason"] = o.PartialReason
	}
	toSerialize["summary"] = o.Summary
	if o.TimeRange != nil {
		toSerialize["timeRange"] = o.TimeRange
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentEvidenceSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EvidenceId    *string                    `json:"evidenceId"`
		RequirementId *string                    `json:"requirementId"`
		Type          *string                    `json:"type,omitempty"`
		Title         *string                    `json:"title,omitempty"`
		Source        *string                    `json:"source,omitempty"`
		ObjectRef     *AiAgentObjectRef          `json:"objectRef,omitempty"`
		Status        *AiAgentEvidenceStatus     `json:"status"`
		StatusDisplay *AiAgentStatusDisplay      `json:"statusDisplay,omitempty"`
		Visibility    *AiAgentEvidenceVisibility `json:"visibility,omitempty"`
		MissingReason *AiAgentMissingReason      `json:"missingReason,omitempty"`
		PartialReason *AiAgentPartialReason      `json:"partialReason,omitempty"`
		Summary       *string                    `json:"summary"`
		TimeRange     *AiAgentTimeRange          `json:"timeRange,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EvidenceId == nil {
		return fmt.Errorf("required field evidenceId missing")
	}
	if all.RequirementId == nil {
		return fmt.Errorf("required field requirementId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"evidenceId", "requirementId", "type", "title", "source", "objectRef", "status", "statusDisplay", "visibility", "missingReason", "partialReason", "summary", "timeRange"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EvidenceId = *all.EvidenceId
	o.RequirementId = *all.RequirementId
	o.Type = all.Type
	o.Title = all.Title
	o.Source = all.Source
	if all.ObjectRef != nil && all.ObjectRef.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ObjectRef = all.ObjectRef
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if all.StatusDisplay != nil && all.StatusDisplay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusDisplay = all.StatusDisplay
	if all.Visibility != nil && !all.Visibility.IsValid() {
		hasInvalidField = true
	} else {
		o.Visibility = all.Visibility
	}
	if all.MissingReason != nil && !all.MissingReason.IsValid() {
		hasInvalidField = true
	} else {
		o.MissingReason = all.MissingReason
	}
	if all.PartialReason != nil && !all.PartialReason.IsValid() {
		hasInvalidField = true
	} else {
		o.PartialReason = all.PartialReason
	}
	o.Summary = *all.Summary
	if all.TimeRange != nil && all.TimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeRange = all.TimeRange

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
