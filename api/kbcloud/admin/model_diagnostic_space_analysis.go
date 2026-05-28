// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSpaceAnalysis Sprint 26 PostgreSQL readonly space analysis contract. Database Top objects
// and PVC/PV capacity evidence are intentionally separated so Console does
// not imply that a PVC capacity symptom proves one table or index is large.
type DiagnosticSpaceAnalysis struct {
	// Business-result axis for PostgreSQL readonly space analysis. It separates
	// "collector ran and found Top objects" from "collector ran but there was no
	// user object", "collector failed", and the compatibility state where no
	// supported data source is available.
	//
	State DiagnosticSpaceAnalysisState `json:"state"`
	// Stable reason code for unavailable or degraded space analysis.
	ReasonCode DiagnosticSpaceAnalysisReasonCode `json:"reasonCode"`
	// Data quality status for one database performance metric.
	DataQuality DiagnosticMetricDataQuality `json:"dataQuality"`
	// First-screen summary for the space analysis card.
	Summary string `json:"summary"`
	// Time when the space analysis result was assembled.
	CollectedAt time.Time `json:"collectedAt"`
	// Readonly collection thresholds and limits used by space analysis.
	Threshold DiagnosticSpaceAnalysisThreshold `json:"threshold"`
	// Total PostgreSQL database size observed by the readonly collector.
	TotalSize DiagnosticSpaceAnalysisTotalSize `json:"totalSize"`
	// Top PostgreSQL objects sorted by sizeBytes descending.
	TopObjects []DiagnosticSpaceAnalysisObject `json:"topObjects"`
	// PVC/PV capacity evidence shown in a separate storage section.
	StorageEvidence []DiagnosticSpaceAnalysisStorageEvidence `json:"storageEvidence"`
	// Explicit safety boundary for the space analysis collector.
	ReadonlyBoundary DiagnosticSpaceAnalysisReadonlyBoundary `json:"readonlyBoundary"`
	// Structured collection error detail when spaceAnalysisState is source_present_collection_failed.
	CollectionError *DiagnosticSpaceAnalysisCollectionError `json:"collectionError,omitempty"`
	// Safe actions only: open detail, copy redacted context, jump to evidence, or rerun diagnosis.
	NextActions []DiagnosticNextAction `json:"nextActions"`
	// Explicit limits of this evidence.
	CannotProve []string `json:"cannotProve"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSpaceAnalysis instantiates a new DiagnosticSpaceAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSpaceAnalysis(state DiagnosticSpaceAnalysisState, reasonCode DiagnosticSpaceAnalysisReasonCode, dataQuality DiagnosticMetricDataQuality, summary string, collectedAt time.Time, threshold DiagnosticSpaceAnalysisThreshold, totalSize DiagnosticSpaceAnalysisTotalSize, topObjects []DiagnosticSpaceAnalysisObject, storageEvidence []DiagnosticSpaceAnalysisStorageEvidence, readonlyBoundary DiagnosticSpaceAnalysisReadonlyBoundary, nextActions []DiagnosticNextAction, cannotProve []string) *DiagnosticSpaceAnalysis {
	this := DiagnosticSpaceAnalysis{}
	this.State = state
	this.ReasonCode = reasonCode
	this.DataQuality = dataQuality
	this.Summary = summary
	this.CollectedAt = collectedAt
	this.Threshold = threshold
	this.TotalSize = totalSize
	this.TopObjects = topObjects
	this.StorageEvidence = storageEvidence
	this.ReadonlyBoundary = readonlyBoundary
	this.NextActions = nextActions
	this.CannotProve = cannotProve
	return &this
}

// NewDiagnosticSpaceAnalysisWithDefaults instantiates a new DiagnosticSpaceAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSpaceAnalysisWithDefaults() *DiagnosticSpaceAnalysis {
	this := DiagnosticSpaceAnalysis{}
	return &this
}

// GetState returns the State field value.
func (o *DiagnosticSpaceAnalysis) GetState() DiagnosticSpaceAnalysisState {
	if o == nil {
		var ret DiagnosticSpaceAnalysisState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetStateOk() (*DiagnosticSpaceAnalysisState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *DiagnosticSpaceAnalysis) SetState(v DiagnosticSpaceAnalysisState) {
	o.State = v
}

// GetReasonCode returns the ReasonCode field value.
func (o *DiagnosticSpaceAnalysis) GetReasonCode() DiagnosticSpaceAnalysisReasonCode {
	if o == nil {
		var ret DiagnosticSpaceAnalysisReasonCode
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetReasonCodeOk() (*DiagnosticSpaceAnalysisReasonCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value.
func (o *DiagnosticSpaceAnalysis) SetReasonCode(v DiagnosticSpaceAnalysisReasonCode) {
	o.ReasonCode = v
}

// GetDataQuality returns the DataQuality field value.
func (o *DiagnosticSpaceAnalysis) GetDataQuality() DiagnosticMetricDataQuality {
	if o == nil {
		var ret DiagnosticMetricDataQuality
		return ret
	}
	return o.DataQuality
}

// GetDataQualityOk returns a tuple with the DataQuality field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetDataQualityOk() (*DiagnosticMetricDataQuality, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataQuality, true
}

// SetDataQuality sets field value.
func (o *DiagnosticSpaceAnalysis) SetDataQuality(v DiagnosticMetricDataQuality) {
	o.DataQuality = v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticSpaceAnalysis) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticSpaceAnalysis) SetSummary(v string) {
	o.Summary = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticSpaceAnalysis) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticSpaceAnalysis) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// GetThreshold returns the Threshold field value.
func (o *DiagnosticSpaceAnalysis) GetThreshold() DiagnosticSpaceAnalysisThreshold {
	if o == nil {
		var ret DiagnosticSpaceAnalysisThreshold
		return ret
	}
	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetThresholdOk() (*DiagnosticSpaceAnalysisThreshold, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value.
func (o *DiagnosticSpaceAnalysis) SetThreshold(v DiagnosticSpaceAnalysisThreshold) {
	o.Threshold = v
}

// GetTotalSize returns the TotalSize field value.
func (o *DiagnosticSpaceAnalysis) GetTotalSize() DiagnosticSpaceAnalysisTotalSize {
	if o == nil {
		var ret DiagnosticSpaceAnalysisTotalSize
		return ret
	}
	return o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetTotalSizeOk() (*DiagnosticSpaceAnalysisTotalSize, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSize, true
}

// SetTotalSize sets field value.
func (o *DiagnosticSpaceAnalysis) SetTotalSize(v DiagnosticSpaceAnalysisTotalSize) {
	o.TotalSize = v
}

// GetTopObjects returns the TopObjects field value.
func (o *DiagnosticSpaceAnalysis) GetTopObjects() []DiagnosticSpaceAnalysisObject {
	if o == nil {
		var ret []DiagnosticSpaceAnalysisObject
		return ret
	}
	return o.TopObjects
}

// GetTopObjectsOk returns a tuple with the TopObjects field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetTopObjectsOk() (*[]DiagnosticSpaceAnalysisObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopObjects, true
}

// SetTopObjects sets field value.
func (o *DiagnosticSpaceAnalysis) SetTopObjects(v []DiagnosticSpaceAnalysisObject) {
	o.TopObjects = v
}

// GetStorageEvidence returns the StorageEvidence field value.
func (o *DiagnosticSpaceAnalysis) GetStorageEvidence() []DiagnosticSpaceAnalysisStorageEvidence {
	if o == nil {
		var ret []DiagnosticSpaceAnalysisStorageEvidence
		return ret
	}
	return o.StorageEvidence
}

// GetStorageEvidenceOk returns a tuple with the StorageEvidence field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetStorageEvidenceOk() (*[]DiagnosticSpaceAnalysisStorageEvidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageEvidence, true
}

// SetStorageEvidence sets field value.
func (o *DiagnosticSpaceAnalysis) SetStorageEvidence(v []DiagnosticSpaceAnalysisStorageEvidence) {
	o.StorageEvidence = v
}

// GetReadonlyBoundary returns the ReadonlyBoundary field value.
func (o *DiagnosticSpaceAnalysis) GetReadonlyBoundary() DiagnosticSpaceAnalysisReadonlyBoundary {
	if o == nil {
		var ret DiagnosticSpaceAnalysisReadonlyBoundary
		return ret
	}
	return o.ReadonlyBoundary
}

// GetReadonlyBoundaryOk returns a tuple with the ReadonlyBoundary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetReadonlyBoundaryOk() (*DiagnosticSpaceAnalysisReadonlyBoundary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadonlyBoundary, true
}

// SetReadonlyBoundary sets field value.
func (o *DiagnosticSpaceAnalysis) SetReadonlyBoundary(v DiagnosticSpaceAnalysisReadonlyBoundary) {
	o.ReadonlyBoundary = v
}

// GetCollectionError returns the CollectionError field value if set, zero value otherwise.
func (o *DiagnosticSpaceAnalysis) GetCollectionError() DiagnosticSpaceAnalysisCollectionError {
	if o == nil || o.CollectionError == nil {
		var ret DiagnosticSpaceAnalysisCollectionError
		return ret
	}
	return *o.CollectionError
}

// GetCollectionErrorOk returns a tuple with the CollectionError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetCollectionErrorOk() (*DiagnosticSpaceAnalysisCollectionError, bool) {
	if o == nil || o.CollectionError == nil {
		return nil, false
	}
	return o.CollectionError, true
}

// HasCollectionError returns a boolean if a field has been set.
func (o *DiagnosticSpaceAnalysis) HasCollectionError() bool {
	return o != nil && o.CollectionError != nil
}

// SetCollectionError gets a reference to the given DiagnosticSpaceAnalysisCollectionError and assigns it to the CollectionError field.
func (o *DiagnosticSpaceAnalysis) SetCollectionError(v DiagnosticSpaceAnalysisCollectionError) {
	o.CollectionError = &v
}

// GetNextActions returns the NextActions field value.
func (o *DiagnosticSpaceAnalysis) GetNextActions() []DiagnosticNextAction {
	if o == nil {
		var ret []DiagnosticNextAction
		return ret
	}
	return o.NextActions
}

// GetNextActionsOk returns a tuple with the NextActions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetNextActionsOk() (*[]DiagnosticNextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextActions, true
}

// SetNextActions sets field value.
func (o *DiagnosticSpaceAnalysis) SetNextActions(v []DiagnosticNextAction) {
	o.NextActions = v
}

// GetCannotProve returns the CannotProve field value.
func (o *DiagnosticSpaceAnalysis) GetCannotProve() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CannotProve
}

// GetCannotProveOk returns a tuple with the CannotProve field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysis) GetCannotProveOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CannotProve, true
}

// SetCannotProve sets field value.
func (o *DiagnosticSpaceAnalysis) SetCannotProve(v []string) {
	o.CannotProve = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSpaceAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["state"] = o.State
	toSerialize["reasonCode"] = o.ReasonCode
	toSerialize["dataQuality"] = o.DataQuality
	toSerialize["summary"] = o.Summary
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["threshold"] = o.Threshold
	toSerialize["totalSize"] = o.TotalSize
	toSerialize["topObjects"] = o.TopObjects
	toSerialize["storageEvidence"] = o.StorageEvidence
	toSerialize["readonlyBoundary"] = o.ReadonlyBoundary
	if o.CollectionError != nil {
		toSerialize["collectionError"] = o.CollectionError
	}
	toSerialize["nextActions"] = o.NextActions
	toSerialize["cannotProve"] = o.CannotProve

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSpaceAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		State            *DiagnosticSpaceAnalysisState             `json:"state"`
		ReasonCode       *DiagnosticSpaceAnalysisReasonCode        `json:"reasonCode"`
		DataQuality      *DiagnosticMetricDataQuality              `json:"dataQuality"`
		Summary          *string                                   `json:"summary"`
		CollectedAt      *time.Time                                `json:"collectedAt"`
		Threshold        *DiagnosticSpaceAnalysisThreshold         `json:"threshold"`
		TotalSize        *DiagnosticSpaceAnalysisTotalSize         `json:"totalSize"`
		TopObjects       *[]DiagnosticSpaceAnalysisObject          `json:"topObjects"`
		StorageEvidence  *[]DiagnosticSpaceAnalysisStorageEvidence `json:"storageEvidence"`
		ReadonlyBoundary *DiagnosticSpaceAnalysisReadonlyBoundary  `json:"readonlyBoundary"`
		CollectionError  *DiagnosticSpaceAnalysisCollectionError   `json:"collectionError,omitempty"`
		NextActions      *[]DiagnosticNextAction                   `json:"nextActions"`
		CannotProve      *[]string                                 `json:"cannotProve"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.ReasonCode == nil {
		return fmt.Errorf("required field reasonCode missing")
	}
	if all.DataQuality == nil {
		return fmt.Errorf("required field dataQuality missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.Threshold == nil {
		return fmt.Errorf("required field threshold missing")
	}
	if all.TotalSize == nil {
		return fmt.Errorf("required field totalSize missing")
	}
	if all.TopObjects == nil {
		return fmt.Errorf("required field topObjects missing")
	}
	if all.StorageEvidence == nil {
		return fmt.Errorf("required field storageEvidence missing")
	}
	if all.ReadonlyBoundary == nil {
		return fmt.Errorf("required field readonlyBoundary missing")
	}
	if all.NextActions == nil {
		return fmt.Errorf("required field nextActions missing")
	}
	if all.CannotProve == nil {
		return fmt.Errorf("required field cannotProve missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"state", "reasonCode", "dataQuality", "summary", "collectedAt", "threshold", "totalSize", "topObjects", "storageEvidence", "readonlyBoundary", "collectionError", "nextActions", "cannotProve"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	if !all.ReasonCode.IsValid() {
		hasInvalidField = true
	} else {
		o.ReasonCode = *all.ReasonCode
	}
	if !all.DataQuality.IsValid() {
		hasInvalidField = true
	} else {
		o.DataQuality = *all.DataQuality
	}
	o.Summary = *all.Summary
	o.CollectedAt = *all.CollectedAt
	if all.Threshold.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Threshold = *all.Threshold
	if all.TotalSize.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TotalSize = *all.TotalSize
	o.TopObjects = *all.TopObjects
	o.StorageEvidence = *all.StorageEvidence
	if all.ReadonlyBoundary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReadonlyBoundary = *all.ReadonlyBoundary
	if all.CollectionError != nil && all.CollectionError.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CollectionError = all.CollectionError
	o.NextActions = *all.NextActions
	o.CannotProve = *all.CannotProve

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
