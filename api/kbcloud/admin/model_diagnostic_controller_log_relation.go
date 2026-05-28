// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticControllerLogRelation Controller-log explanation card consumed by Console before raw evidence.
// It groups redacted raw controller logs by product relation, error type,
// recency and next action.
type DiagnosticControllerLogRelation struct {
	// Product-facing relation bucket for one controller-log explanation card.
	// This is separate from the legacy check-level relationStatus/riskDomain
	// fields, which remain for Sprint 6/7 compatibility.
	//
	Relation         DiagnosticControllerLogRelationKind `json:"relation"`
	RelatedResources []DiagnosticRelatedResource         `json:"relatedResources"`
	// Time range covered by one controller-log relation card. Together with
	// lastOccurredAt, this is the recency contract for Console display.
	//
	TimeWindow DiagnosticControllerLogRelationTimeWindow `json:"timeWindow"`
	// Latest known occurrence time for this relation card.
	LastOccurredAt time.Time `json:"lastOccurredAt"`
	// Stable snake_case grouping hint, such as reconcile_failed, object_not_found, permission_denied, or unknown_error.
	ErrorType string `json:"errorType"`
	// Number of redacted controller log entries represented by this card.
	Count              int64                                      `json:"count"`
	RepresentativeLogs []DiagnosticControllerLogRepresentativeLog `json:"representativeLogs"`
	// Structured next action for one diagnostic check or top issue.
	NextAction DiagnosticNextAction `json:"nextAction"`
	// Report-internal pointers to rawEvidence entries backing this relation card.
	RawEvidenceRefs []string `json:"rawEvidenceRefs"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticControllerLogRelation instantiates a new DiagnosticControllerLogRelation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticControllerLogRelation(relation DiagnosticControllerLogRelationKind, relatedResources []DiagnosticRelatedResource, timeWindow DiagnosticControllerLogRelationTimeWindow, lastOccurredAt time.Time, errorType string, count int64, representativeLogs []DiagnosticControllerLogRepresentativeLog, nextAction DiagnosticNextAction, rawEvidenceRefs []string) *DiagnosticControllerLogRelation {
	this := DiagnosticControllerLogRelation{}
	this.Relation = relation
	this.RelatedResources = relatedResources
	this.TimeWindow = timeWindow
	this.LastOccurredAt = lastOccurredAt
	this.ErrorType = errorType
	this.Count = count
	this.RepresentativeLogs = representativeLogs
	this.NextAction = nextAction
	this.RawEvidenceRefs = rawEvidenceRefs
	return &this
}

// NewDiagnosticControllerLogRelationWithDefaults instantiates a new DiagnosticControllerLogRelation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticControllerLogRelationWithDefaults() *DiagnosticControllerLogRelation {
	this := DiagnosticControllerLogRelation{}
	return &this
}

// GetRelation returns the Relation field value.
func (o *DiagnosticControllerLogRelation) GetRelation() DiagnosticControllerLogRelationKind {
	if o == nil {
		var ret DiagnosticControllerLogRelationKind
		return ret
	}
	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelation) GetRelationOk() (*DiagnosticControllerLogRelationKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value.
func (o *DiagnosticControllerLogRelation) SetRelation(v DiagnosticControllerLogRelationKind) {
	o.Relation = v
}

// GetRelatedResources returns the RelatedResources field value.
func (o *DiagnosticControllerLogRelation) GetRelatedResources() []DiagnosticRelatedResource {
	if o == nil {
		var ret []DiagnosticRelatedResource
		return ret
	}
	return o.RelatedResources
}

// GetRelatedResourcesOk returns a tuple with the RelatedResources field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelation) GetRelatedResourcesOk() (*[]DiagnosticRelatedResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelatedResources, true
}

// SetRelatedResources sets field value.
func (o *DiagnosticControllerLogRelation) SetRelatedResources(v []DiagnosticRelatedResource) {
	o.RelatedResources = v
}

// GetTimeWindow returns the TimeWindow field value.
func (o *DiagnosticControllerLogRelation) GetTimeWindow() DiagnosticControllerLogRelationTimeWindow {
	if o == nil {
		var ret DiagnosticControllerLogRelationTimeWindow
		return ret
	}
	return o.TimeWindow
}

// GetTimeWindowOk returns a tuple with the TimeWindow field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelation) GetTimeWindowOk() (*DiagnosticControllerLogRelationTimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeWindow, true
}

// SetTimeWindow sets field value.
func (o *DiagnosticControllerLogRelation) SetTimeWindow(v DiagnosticControllerLogRelationTimeWindow) {
	o.TimeWindow = v
}

// GetLastOccurredAt returns the LastOccurredAt field value.
func (o *DiagnosticControllerLogRelation) GetLastOccurredAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastOccurredAt
}

// GetLastOccurredAtOk returns a tuple with the LastOccurredAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelation) GetLastOccurredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastOccurredAt, true
}

// SetLastOccurredAt sets field value.
func (o *DiagnosticControllerLogRelation) SetLastOccurredAt(v time.Time) {
	o.LastOccurredAt = v
}

// GetErrorType returns the ErrorType field value.
func (o *DiagnosticControllerLogRelation) GetErrorType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelation) GetErrorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorType, true
}

// SetErrorType sets field value.
func (o *DiagnosticControllerLogRelation) SetErrorType(v string) {
	o.ErrorType = v
}

// GetCount returns the Count field value.
func (o *DiagnosticControllerLogRelation) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelation) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *DiagnosticControllerLogRelation) SetCount(v int64) {
	o.Count = v
}

// GetRepresentativeLogs returns the RepresentativeLogs field value.
func (o *DiagnosticControllerLogRelation) GetRepresentativeLogs() []DiagnosticControllerLogRepresentativeLog {
	if o == nil {
		var ret []DiagnosticControllerLogRepresentativeLog
		return ret
	}
	return o.RepresentativeLogs
}

// GetRepresentativeLogsOk returns a tuple with the RepresentativeLogs field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelation) GetRepresentativeLogsOk() (*[]DiagnosticControllerLogRepresentativeLog, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepresentativeLogs, true
}

// SetRepresentativeLogs sets field value.
func (o *DiagnosticControllerLogRelation) SetRepresentativeLogs(v []DiagnosticControllerLogRepresentativeLog) {
	o.RepresentativeLogs = v
}

// GetNextAction returns the NextAction field value.
func (o *DiagnosticControllerLogRelation) GetNextAction() DiagnosticNextAction {
	if o == nil {
		var ret DiagnosticNextAction
		return ret
	}
	return o.NextAction
}

// GetNextActionOk returns a tuple with the NextAction field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelation) GetNextActionOk() (*DiagnosticNextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextAction, true
}

// SetNextAction sets field value.
func (o *DiagnosticControllerLogRelation) SetNextAction(v DiagnosticNextAction) {
	o.NextAction = v
}

// GetRawEvidenceRefs returns the RawEvidenceRefs field value.
func (o *DiagnosticControllerLogRelation) GetRawEvidenceRefs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RawEvidenceRefs
}

// GetRawEvidenceRefsOk returns a tuple with the RawEvidenceRefs field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelation) GetRawEvidenceRefsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawEvidenceRefs, true
}

// SetRawEvidenceRefs sets field value.
func (o *DiagnosticControllerLogRelation) SetRawEvidenceRefs(v []string) {
	o.RawEvidenceRefs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticControllerLogRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["relation"] = o.Relation
	toSerialize["relatedResources"] = o.RelatedResources
	toSerialize["timeWindow"] = o.TimeWindow
	if o.LastOccurredAt.Nanosecond() == 0 {
		toSerialize["lastOccurredAt"] = o.LastOccurredAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["lastOccurredAt"] = o.LastOccurredAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["errorType"] = o.ErrorType
	toSerialize["count"] = o.Count
	toSerialize["representativeLogs"] = o.RepresentativeLogs
	toSerialize["nextAction"] = o.NextAction
	toSerialize["rawEvidenceRefs"] = o.RawEvidenceRefs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticControllerLogRelation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Relation           *DiagnosticControllerLogRelationKind        `json:"relation"`
		RelatedResources   *[]DiagnosticRelatedResource                `json:"relatedResources"`
		TimeWindow         *DiagnosticControllerLogRelationTimeWindow  `json:"timeWindow"`
		LastOccurredAt     *time.Time                                  `json:"lastOccurredAt"`
		ErrorType          *string                                     `json:"errorType"`
		Count              *int64                                      `json:"count"`
		RepresentativeLogs *[]DiagnosticControllerLogRepresentativeLog `json:"representativeLogs"`
		NextAction         *DiagnosticNextAction                       `json:"nextAction"`
		RawEvidenceRefs    *[]string                                   `json:"rawEvidenceRefs"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Relation == nil {
		return fmt.Errorf("required field relation missing")
	}
	if all.RelatedResources == nil {
		return fmt.Errorf("required field relatedResources missing")
	}
	if all.TimeWindow == nil {
		return fmt.Errorf("required field timeWindow missing")
	}
	if all.LastOccurredAt == nil {
		return fmt.Errorf("required field lastOccurredAt missing")
	}
	if all.ErrorType == nil {
		return fmt.Errorf("required field errorType missing")
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.RepresentativeLogs == nil {
		return fmt.Errorf("required field representativeLogs missing")
	}
	if all.NextAction == nil {
		return fmt.Errorf("required field nextAction missing")
	}
	if all.RawEvidenceRefs == nil {
		return fmt.Errorf("required field rawEvidenceRefs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"relation", "relatedResources", "timeWindow", "lastOccurredAt", "errorType", "count", "representativeLogs", "nextAction", "rawEvidenceRefs"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Relation.IsValid() {
		hasInvalidField = true
	} else {
		o.Relation = *all.Relation
	}
	o.RelatedResources = *all.RelatedResources
	if all.TimeWindow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TimeWindow = *all.TimeWindow
	o.LastOccurredAt = *all.LastOccurredAt
	o.ErrorType = *all.ErrorType
	o.Count = *all.Count
	o.RepresentativeLogs = *all.RepresentativeLogs
	if all.NextAction.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NextAction = *all.NextAction
	o.RawEvidenceRefs = *all.RawEvidenceRefs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
