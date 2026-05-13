// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentDiagnosisBranch struct {
	Status *AiAgentDiagnosisBranchStatus `json:"status,omitempty"`
	// Stable status code plus default user-facing copy. Frontend uses code for logic/tests and may override display text for i18n.
	StatusDisplay        *AiAgentStatusDisplay        `json:"statusDisplay,omitempty"`
	ReportId             *string                      `json:"reportId,omitempty"`
	EvidenceCompleteness *AiAgentEvidenceCompleteness `json:"evidenceCompleteness,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentDiagnosisBranch instantiates a new AiAgentDiagnosisBranch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentDiagnosisBranch() *AiAgentDiagnosisBranch {
	this := AiAgentDiagnosisBranch{}
	return &this
}

// NewAiAgentDiagnosisBranchWithDefaults instantiates a new AiAgentDiagnosisBranch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentDiagnosisBranchWithDefaults() *AiAgentDiagnosisBranch {
	this := AiAgentDiagnosisBranch{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiAgentDiagnosisBranch) GetStatus() AiAgentDiagnosisBranchStatus {
	if o == nil || o.Status == nil {
		var ret AiAgentDiagnosisBranchStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentDiagnosisBranch) GetStatusOk() (*AiAgentDiagnosisBranchStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiAgentDiagnosisBranch) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given AiAgentDiagnosisBranchStatus and assigns it to the Status field.
func (o *AiAgentDiagnosisBranch) SetStatus(v AiAgentDiagnosisBranchStatus) {
	o.Status = &v
}

// GetStatusDisplay returns the StatusDisplay field value if set, zero value otherwise.
func (o *AiAgentDiagnosisBranch) GetStatusDisplay() AiAgentStatusDisplay {
	if o == nil || o.StatusDisplay == nil {
		var ret AiAgentStatusDisplay
		return ret
	}
	return *o.StatusDisplay
}

// GetStatusDisplayOk returns a tuple with the StatusDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentDiagnosisBranch) GetStatusDisplayOk() (*AiAgentStatusDisplay, bool) {
	if o == nil || o.StatusDisplay == nil {
		return nil, false
	}
	return o.StatusDisplay, true
}

// HasStatusDisplay returns a boolean if a field has been set.
func (o *AiAgentDiagnosisBranch) HasStatusDisplay() bool {
	return o != nil && o.StatusDisplay != nil
}

// SetStatusDisplay gets a reference to the given AiAgentStatusDisplay and assigns it to the StatusDisplay field.
func (o *AiAgentDiagnosisBranch) SetStatusDisplay(v AiAgentStatusDisplay) {
	o.StatusDisplay = &v
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *AiAgentDiagnosisBranch) GetReportId() string {
	if o == nil || o.ReportId == nil {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentDiagnosisBranch) GetReportIdOk() (*string, bool) {
	if o == nil || o.ReportId == nil {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *AiAgentDiagnosisBranch) HasReportId() bool {
	return o != nil && o.ReportId != nil
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *AiAgentDiagnosisBranch) SetReportId(v string) {
	o.ReportId = &v
}

// GetEvidenceCompleteness returns the EvidenceCompleteness field value if set, zero value otherwise.
func (o *AiAgentDiagnosisBranch) GetEvidenceCompleteness() AiAgentEvidenceCompleteness {
	if o == nil || o.EvidenceCompleteness == nil {
		var ret AiAgentEvidenceCompleteness
		return ret
	}
	return *o.EvidenceCompleteness
}

// GetEvidenceCompletenessOk returns a tuple with the EvidenceCompleteness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentDiagnosisBranch) GetEvidenceCompletenessOk() (*AiAgentEvidenceCompleteness, bool) {
	if o == nil || o.EvidenceCompleteness == nil {
		return nil, false
	}
	return o.EvidenceCompleteness, true
}

// HasEvidenceCompleteness returns a boolean if a field has been set.
func (o *AiAgentDiagnosisBranch) HasEvidenceCompleteness() bool {
	return o != nil && o.EvidenceCompleteness != nil
}

// SetEvidenceCompleteness gets a reference to the given AiAgentEvidenceCompleteness and assigns it to the EvidenceCompleteness field.
func (o *AiAgentDiagnosisBranch) SetEvidenceCompleteness(v AiAgentEvidenceCompleteness) {
	o.EvidenceCompleteness = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentDiagnosisBranch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDisplay != nil {
		toSerialize["statusDisplay"] = o.StatusDisplay
	}
	if o.ReportId != nil {
		toSerialize["reportId"] = o.ReportId
	}
	if o.EvidenceCompleteness != nil {
		toSerialize["evidenceCompleteness"] = o.EvidenceCompleteness
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentDiagnosisBranch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status               *AiAgentDiagnosisBranchStatus `json:"status,omitempty"`
		StatusDisplay        *AiAgentStatusDisplay         `json:"statusDisplay,omitempty"`
		ReportId             *string                       `json:"reportId,omitempty"`
		EvidenceCompleteness *AiAgentEvidenceCompleteness  `json:"evidenceCompleteness,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "statusDisplay", "reportId", "evidenceCompleteness"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	if all.StatusDisplay != nil && all.StatusDisplay.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusDisplay = all.StatusDisplay
	o.ReportId = all.ReportId
	if all.EvidenceCompleteness != nil && !all.EvidenceCompleteness.IsValid() {
		hasInvalidField = true
	} else {
		o.EvidenceCompleteness = all.EvidenceCompleteness
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
