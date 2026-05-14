// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentEvidenceDetail struct {
	EvidenceId    string                `json:"evidenceId"`
	RunId         string                `json:"runId"`
	RequirementId string                `json:"requirementId"`
	Type          *string               `json:"type,omitempty"`
	Status        AiAgentEvidenceStatus `json:"status"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	StatusDisplay   *AiAgentStatusDisplay                 `json:"statusDisplay,omitempty"`
	Visibility      AiAgentEvidenceVisibility             `json:"visibility"`
	MissingReason   *AiAgentMissingReason                 `json:"missingReason,omitempty"`
	PartialReason   *AiAgentPartialReason                 `json:"partialReason,omitempty"`
	MinimumFields   []string                              `json:"minimumFields,omitempty"`
	RedactionStatus *AiAgentEvidenceDetailRedactionStatus `json:"redactionStatus,omitempty"`
	Truncated       *bool                                 `json:"truncated,omitempty"`
	SensitiveTypes  []string                              `json:"sensitiveTypes,omitempty"`
	ObjectRef       *AiAgentObjectRef                     `json:"objectRef,omitempty"`
	TimeRange       *AiAgentTimeRange                     `json:"timeRange,omitempty"`
	Summary         *string                               `json:"summary,omitempty"`
	Snippets        []AiAgentEvidenceSnippet              `json:"snippets,omitempty"`
	Payload         map[string]interface{}                `json:"payload,omitempty"`
	SourceEventIds  []string                              `json:"sourceEventIds,omitempty"`
	RawAvailable    *bool                                 `json:"rawAvailable,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentEvidenceDetail instantiates a new AiAgentEvidenceDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentEvidenceDetail(evidenceId string, runId string, requirementId string, status AiAgentEvidenceStatus, visibility AiAgentEvidenceVisibility) *AiAgentEvidenceDetail {
	this := AiAgentEvidenceDetail{}
	this.EvidenceId = evidenceId
	this.RunId = runId
	this.RequirementId = requirementId
	this.Status = status
	this.Visibility = visibility
	var rawAvailable bool = false
	this.RawAvailable = &rawAvailable
	return &this
}

// NewAiAgentEvidenceDetailWithDefaults instantiates a new AiAgentEvidenceDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentEvidenceDetailWithDefaults() *AiAgentEvidenceDetail {
	this := AiAgentEvidenceDetail{}
	var rawAvailable bool = false
	this.RawAvailable = &rawAvailable
	return &this
}

// GetEvidenceId returns the EvidenceId field value.
func (o *AiAgentEvidenceDetail) GetEvidenceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvidenceId
}

// GetEvidenceIdOk returns a tuple with the EvidenceId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetEvidenceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvidenceId, true
}

// SetEvidenceId sets field value.
func (o *AiAgentEvidenceDetail) SetEvidenceId(v string) {
	o.EvidenceId = v
}

// GetRunId returns the RunId field value.
func (o *AiAgentEvidenceDetail) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *AiAgentEvidenceDetail) SetRunId(v string) {
	o.RunId = v
}

// GetRequirementId returns the RequirementId field value.
func (o *AiAgentEvidenceDetail) GetRequirementId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RequirementId
}

// GetRequirementIdOk returns a tuple with the RequirementId field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetRequirementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequirementId, true
}

// SetRequirementId sets field value.
func (o *AiAgentEvidenceDetail) SetRequirementId(v string) {
	o.RequirementId = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AiAgentEvidenceDetail) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value.
func (o *AiAgentEvidenceDetail) GetStatus() AiAgentEvidenceStatus {
	if o == nil {
		var ret AiAgentEvidenceStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetStatusOk() (*AiAgentEvidenceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *AiAgentEvidenceDetail) SetStatus(v AiAgentEvidenceStatus) {
	o.Status = v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetStatusDisplay() AiAgentStatusDisplay {
	if o == nil || o.StatusDisplay == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetStatusDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasStatusDisplay() bool {
	return o != nil && o.StatusDisplay != nil
}

// SetStatusDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the StatusDisplay field.
func (o *AiAgentEvidenceDetail) SetStatusDisplay(v AiAgentStatusDisplay) {
	o.StatusDisplay = &v
}

// GetVisibility returns the Visibility field value.
func (o *AiAgentEvidenceDetail) GetVisibility() AiAgentEvidenceVisibility {
	if o == nil {
		var ret AiAgentEvidenceVisibility
		return ret
	}
	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetVisibilityOk() (*AiAgentEvidenceVisibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value.
func (o *AiAgentEvidenceDetail) SetVisibility(v AiAgentEvidenceVisibility) {
	o.Visibility = v
}

// GetMissingReason returns the MissingReason field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetMissingReason() AiAgentMissingReason {
	if o == nil || o.MissingReason == nil {
		var ret AiAgentMissingReason
		return ret
	}
	return *o.MissingReason
}

// GetMissingReasonOk returns a tuple with the MissingReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetMissingReasonOk() (*AiAgentMissingReason, bool) {
	if o == nil || o.MissingReason == nil {
		return nil, false
	}
	return o.MissingReason, true
}

// HasMissingReason returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasMissingReason() bool {
	return o != nil && o.MissingReason != nil
}

// SetMissingReason gets a reference to the given AiAgentMissingReason and assigns it to the MissingReason field.
func (o *AiAgentEvidenceDetail) SetMissingReason(v AiAgentMissingReason) {
	o.MissingReason = &v
}

// GetPartialReason returns the PartialReason field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetPartialReason() AiAgentPartialReason {
	if o == nil || o.PartialReason == nil {
		var ret AiAgentPartialReason
		return ret
	}
	return *o.PartialReason
}

// GetPartialReasonOk returns a tuple with the PartialReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetPartialReasonOk() (*AiAgentPartialReason, bool) {
	if o == nil || o.PartialReason == nil {
		return nil, false
	}
	return o.PartialReason, true
}

// HasPartialReason returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasPartialReason() bool {
	return o != nil && o.PartialReason != nil
}

// SetPartialReason gets a reference to the given AiAgentPartialReason and assigns it to the PartialReason field.
func (o *AiAgentEvidenceDetail) SetPartialReason(v AiAgentPartialReason) {
	o.PartialReason = &v
}

// GetMinimumFields returns the MinimumFields field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetMinimumFields() []string {
	if o == nil || o.MinimumFields == nil {
		var ret []string
		return ret
	}
	return o.MinimumFields
}

// GetMinimumFieldsOk returns a tuple with the MinimumFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetMinimumFieldsOk() (*[]string, bool) {
	if o == nil || o.MinimumFields == nil {
		return nil, false
	}
	return &o.MinimumFields, true
}

// HasMinimumFields returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasMinimumFields() bool {
	return o != nil && o.MinimumFields != nil
}

// SetMinimumFields gets a reference to the given []string and assigns it to the MinimumFields field.
func (o *AiAgentEvidenceDetail) SetMinimumFields(v []string) {
	o.MinimumFields = v
}

// GetRedactionStatus returns the RedactionStatus field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetRedactionStatus() AiAgentEvidenceDetailRedactionStatus {
	if o == nil || o.RedactionStatus == nil {
		var ret AiAgentEvidenceDetailRedactionStatus
		return ret
	}
	return *o.RedactionStatus
}

// GetRedactionStatusOk returns a tuple with the RedactionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetRedactionStatusOk() (*AiAgentEvidenceDetailRedactionStatus, bool) {
	if o == nil || o.RedactionStatus == nil {
		return nil, false
	}
	return o.RedactionStatus, true
}

// HasRedactionStatus returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasRedactionStatus() bool {
	return o != nil && o.RedactionStatus != nil
}

// SetRedactionStatus gets a reference to the given AiAgentEvidenceDetailRedactionStatus and assigns it to the RedactionStatus field.
func (o *AiAgentEvidenceDetail) SetRedactionStatus(v AiAgentEvidenceDetailRedactionStatus) {
	o.RedactionStatus = &v
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetTruncated() bool {
	if o == nil || o.Truncated == nil {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetTruncatedOk() (*bool, bool) {
	if o == nil || o.Truncated == nil {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasTruncated() bool {
	return o != nil && o.Truncated != nil
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *AiAgentEvidenceDetail) SetTruncated(v bool) {
	o.Truncated = &v
}

// GetSensitiveTypes returns the SensitiveTypes field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetSensitiveTypes() []string {
	if o == nil || o.SensitiveTypes == nil {
		var ret []string
		return ret
	}
	return o.SensitiveTypes
}

// GetSensitiveTypesOk returns a tuple with the SensitiveTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetSensitiveTypesOk() (*[]string, bool) {
	if o == nil || o.SensitiveTypes == nil {
		return nil, false
	}
	return &o.SensitiveTypes, true
}

// HasSensitiveTypes returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasSensitiveTypes() bool {
	return o != nil && o.SensitiveTypes != nil
}

// SetSensitiveTypes gets a reference to the given []string and assigns it to the SensitiveTypes field.
func (o *AiAgentEvidenceDetail) SetSensitiveTypes(v []string) {
	o.SensitiveTypes = v
}

// GetObjectRef returns the ObjectRef field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetObjectRef() AiAgentObjectRef {
	if o == nil || o.ObjectRef == nil {
		var ret AiAgentObjectRef
		return ret
	}
	return *o.ObjectRef
}

// GetObjectRefOk returns a tuple with the ObjectRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetObjectRefOk() (*AiAgentObjectRef, bool) {
	if o == nil || o.ObjectRef == nil {
		return nil, false
	}
	return o.ObjectRef, true
}

// HasObjectRef returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasObjectRef() bool {
	return o != nil && o.ObjectRef != nil
}

// SetObjectRef gets a reference to the given AiAgentObjectRef and assigns it to the ObjectRef field.
func (o *AiAgentEvidenceDetail) SetObjectRef(v AiAgentObjectRef) {
	o.ObjectRef = &v
}

// GetTimeRange returns the TimeRange field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetTimeRange() AiAgentTimeRange {
	if o == nil || o.TimeRange == nil {
		var ret AiAgentTimeRange
		return ret
	}
	return *o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetTimeRangeOk() (*AiAgentTimeRange, bool) {
	if o == nil || o.TimeRange == nil {
		return nil, false
	}
	return o.TimeRange, true
}

// HasTimeRange returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasTimeRange() bool {
	return o != nil && o.TimeRange != nil
}

// SetTimeRange gets a reference to the given AiAgentTimeRange and assigns it to the TimeRange field.
func (o *AiAgentEvidenceDetail) SetTimeRange(v AiAgentTimeRange) {
	o.TimeRange = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AiAgentEvidenceDetail) SetSummary(v string) {
	o.Summary = &v
}

// GetSnippets returns the Snippets field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetSnippets() []AiAgentEvidenceSnippet {
	if o == nil || o.Snippets == nil {
		var ret []AiAgentEvidenceSnippet
		return ret
	}
	return o.Snippets
}

// GetSnippetsOk returns a tuple with the Snippets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetSnippetsOk() (*[]AiAgentEvidenceSnippet, bool) {
	if o == nil || o.Snippets == nil {
		return nil, false
	}
	return &o.Snippets, true
}

// HasSnippets returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasSnippets() bool {
	return o != nil && o.Snippets != nil
}

// SetSnippets gets a reference to the given []AiAgentEvidenceSnippet and assigns it to the Snippets field.
func (o *AiAgentEvidenceDetail) SetSnippets(v []AiAgentEvidenceSnippet) {
	o.Snippets = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetPayload() map[string]interface{} {
	if o == nil || o.Payload == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *AiAgentEvidenceDetail) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetSourceEventIds returns the SourceEventIds field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetSourceEventIds() []string {
	if o == nil || o.SourceEventIds == nil {
		var ret []string
		return ret
	}
	return o.SourceEventIds
}

// GetSourceEventIdsOk returns a tuple with the SourceEventIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetSourceEventIdsOk() (*[]string, bool) {
	if o == nil || o.SourceEventIds == nil {
		return nil, false
	}
	return &o.SourceEventIds, true
}

// HasSourceEventIds returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasSourceEventIds() bool {
	return o != nil && o.SourceEventIds != nil
}

// SetSourceEventIds gets a reference to the given []string and assigns it to the SourceEventIds field.
func (o *AiAgentEvidenceDetail) SetSourceEventIds(v []string) {
	o.SourceEventIds = v
}

// GetRawAvailable returns the RawAvailable field value if set, zero value otherwise.
func (o *AiAgentEvidenceDetail) GetRawAvailable() bool {
	if o == nil || o.RawAvailable == nil {
		var ret bool
		return ret
	}
	return *o.RawAvailable
}

// GetRawAvailableOk returns a tuple with the RawAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceDetail) GetRawAvailableOk() (*bool, bool) {
	if o == nil || o.RawAvailable == nil {
		return nil, false
	}
	return o.RawAvailable, true
}

// HasRawAvailable returns a boolean if a field has been set.
func (o *AiAgentEvidenceDetail) HasRawAvailable() bool {
	return o != nil && o.RawAvailable != nil
}

// SetRawAvailable gets a reference to the given bool and assigns it to the RawAvailable field.
func (o *AiAgentEvidenceDetail) SetRawAvailable(v bool) {
	o.RawAvailable = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentEvidenceDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["evidenceId"] = o.EvidenceId
	toSerialize["runId"] = o.RunId
	toSerialize["requirementId"] = o.RequirementId
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["status"] = o.Status
	if o.StatusDisplay != nil {
		toSerialize["statusDisplay"] = o.StatusDisplay
	}
	toSerialize["visibility"] = o.Visibility
	if o.MissingReason != nil {
		toSerialize["missingReason"] = o.MissingReason
	}
	if o.PartialReason != nil {
		toSerialize["partialReason"] = o.PartialReason
	}
	if o.MinimumFields != nil {
		toSerialize["minimumFields"] = o.MinimumFields
	}
	if o.RedactionStatus != nil {
		toSerialize["redactionStatus"] = o.RedactionStatus
	}
	if o.Truncated != nil {
		toSerialize["truncated"] = o.Truncated
	}
	if o.SensitiveTypes != nil {
		toSerialize["sensitiveTypes"] = o.SensitiveTypes
	}
	if o.ObjectRef != nil {
		toSerialize["objectRef"] = o.ObjectRef
	}
	if o.TimeRange != nil {
		toSerialize["timeRange"] = o.TimeRange
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Snippets != nil {
		toSerialize["snippets"] = o.Snippets
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.SourceEventIds != nil {
		toSerialize["sourceEventIds"] = o.SourceEventIds
	}
	if o.RawAvailable != nil {
		toSerialize["rawAvailable"] = o.RawAvailable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentEvidenceDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EvidenceId      *string                               `json:"evidenceId"`
		RunId           *string                               `json:"runId"`
		RequirementId   *string                               `json:"requirementId"`
		Type            *string                               `json:"type,omitempty"`
		Status          *AiAgentEvidenceStatus                `json:"status"`
		StatusDisplay   *AiAgentStatusDisplay                 `json:"statusDisplay,omitempty"`
		Visibility      *AiAgentEvidenceVisibility            `json:"visibility"`
		MissingReason   *AiAgentMissingReason                 `json:"missingReason,omitempty"`
		PartialReason   *AiAgentPartialReason                 `json:"partialReason,omitempty"`
		MinimumFields   []string                              `json:"minimumFields,omitempty"`
		RedactionStatus *AiAgentEvidenceDetailRedactionStatus `json:"redactionStatus,omitempty"`
		Truncated       *bool                                 `json:"truncated,omitempty"`
		SensitiveTypes  []string                              `json:"sensitiveTypes,omitempty"`
		ObjectRef       *AiAgentObjectRef                     `json:"objectRef,omitempty"`
		TimeRange       *AiAgentTimeRange                     `json:"timeRange,omitempty"`
		Summary         *string                               `json:"summary,omitempty"`
		Snippets        []AiAgentEvidenceSnippet              `json:"snippets,omitempty"`
		Payload         map[string]interface{}                `json:"payload,omitempty"`
		SourceEventIds  []string                              `json:"sourceEventIds,omitempty"`
		RawAvailable    *bool                                 `json:"rawAvailable,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EvidenceId == nil {
		return fmt.Errorf("required field evidenceId missing")
	}
	if all.RunId == nil {
		return fmt.Errorf("required field runId missing")
	}
	if all.RequirementId == nil {
		return fmt.Errorf("required field requirementId missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Visibility == nil {
		return fmt.Errorf("required field visibility missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"evidenceId", "runId", "requirementId", "type", "status", "statusDisplay", "visibility", "missingReason", "partialReason", "minimumFields", "redactionStatus", "truncated", "sensitiveTypes", "objectRef", "timeRange", "summary", "snippets", "payload", "sourceEventIds", "rawAvailable"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EvidenceId = *all.EvidenceId
	o.RunId = *all.RunId
	o.RequirementId = *all.RequirementId
	o.Type = all.Type
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	if all.StatusDisplay != nil && all.StatusDisplay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusDisplay = all.StatusDisplay
	if !all.Visibility.IsValid() {
		hasInvalidField = true
	} else {
		o.Visibility = *all.Visibility
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
	o.MinimumFields = all.MinimumFields
	if all.RedactionStatus != nil && !all.RedactionStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.RedactionStatus = all.RedactionStatus
	}
	o.Truncated = all.Truncated
	o.SensitiveTypes = all.SensitiveTypes
	if all.ObjectRef != nil && all.ObjectRef.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ObjectRef = all.ObjectRef
	if all.TimeRange != nil && all.TimeRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeRange = all.TimeRange
	o.Summary = all.Summary
	o.Snippets = all.Snippets
	o.Payload = all.Payload
	o.SourceEventIds = all.SourceEventIds
	o.RawAvailable = all.RawAvailable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
