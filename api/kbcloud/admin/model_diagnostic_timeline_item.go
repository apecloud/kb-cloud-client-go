// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticTimelineItem One time-ordered diagnostic event used to explain what happened.
type DiagnosticTimelineItem struct {
	// Best known time when the event happened.
	OccurredAt time.Time `json:"occurredAt"`
	// Time when this diagnostic run observed the event.
	ObservedAt time.Time `json:"observedAt"`
	// Quality of the timestamp used for ordering a diagnostic timeline item.
	TimeQuality DiagnosticTimelineTimeQuality `json:"timeQuality"`
	// Source type for one diagnostic timeline item.
	Source DiagnosticTimelineSource `json:"source"`
	// Related resource reference.
	ResourceRef string `json:"resourceRef"`
	// Related diagnostic check ID.
	CheckId string `json:"checkId"`
	// User-facing severity and priority bucket for one diagnostic check.
	Severity DiagnosticCheckSeverity `json:"severity"`
	// One-sentence timeline summary.
	Summary string `json:"summary"`
	// Optional structured reason code from the source.
	ReasonCode *string `json:"reasonCode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticTimelineItem instantiates a new DiagnosticTimelineItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticTimelineItem(occurredAt time.Time, observedAt time.Time, timeQuality DiagnosticTimelineTimeQuality, source DiagnosticTimelineSource, resourceRef string, checkId string, severity DiagnosticCheckSeverity, summary string) *DiagnosticTimelineItem {
	this := DiagnosticTimelineItem{}
	this.OccurredAt = occurredAt
	this.ObservedAt = observedAt
	this.TimeQuality = timeQuality
	this.Source = source
	this.ResourceRef = resourceRef
	this.CheckId = checkId
	this.Severity = severity
	this.Summary = summary
	return &this
}

// NewDiagnosticTimelineItemWithDefaults instantiates a new DiagnosticTimelineItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticTimelineItemWithDefaults() *DiagnosticTimelineItem {
	this := DiagnosticTimelineItem{}
	return &this
}

// GetOccurredAt returns the OccurredAt field value.
func (o *DiagnosticTimelineItem) GetOccurredAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.OccurredAt
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTimelineItem) GetOccurredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OccurredAt, true
}

// SetOccurredAt sets field value.
func (o *DiagnosticTimelineItem) SetOccurredAt(v time.Time) {
	o.OccurredAt = v
}

// GetObservedAt returns the ObservedAt field value.
func (o *DiagnosticTimelineItem) GetObservedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTimelineItem) GetObservedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObservedAt, true
}

// SetObservedAt sets field value.
func (o *DiagnosticTimelineItem) SetObservedAt(v time.Time) {
	o.ObservedAt = v
}

// GetTimeQuality returns the TimeQuality field value.
func (o *DiagnosticTimelineItem) GetTimeQuality() DiagnosticTimelineTimeQuality {
	if o == nil {
		var ret DiagnosticTimelineTimeQuality
		return ret
	}
	return o.TimeQuality
}

// GetTimeQualityOk returns a tuple with the TimeQuality field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTimelineItem) GetTimeQualityOk() (*DiagnosticTimelineTimeQuality, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeQuality, true
}

// SetTimeQuality sets field value.
func (o *DiagnosticTimelineItem) SetTimeQuality(v DiagnosticTimelineTimeQuality) {
	o.TimeQuality = v
}

// GetSource returns the Source field value.
func (o *DiagnosticTimelineItem) GetSource() DiagnosticTimelineSource {
	if o == nil {
		var ret DiagnosticTimelineSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTimelineItem) GetSourceOk() (*DiagnosticTimelineSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DiagnosticTimelineItem) SetSource(v DiagnosticTimelineSource) {
	o.Source = v
}

// GetResourceRef returns the ResourceRef field value.
func (o *DiagnosticTimelineItem) GetResourceRef() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceRef
}

// GetResourceRefOk returns a tuple with the ResourceRef field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTimelineItem) GetResourceRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceRef, true
}

// SetResourceRef sets field value.
func (o *DiagnosticTimelineItem) SetResourceRef(v string) {
	o.ResourceRef = v
}

// GetCheckId returns the CheckId field value.
func (o *DiagnosticTimelineItem) GetCheckId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CheckId
}

// GetCheckIdOk returns a tuple with the CheckId field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTimelineItem) GetCheckIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckId, true
}

// SetCheckId sets field value.
func (o *DiagnosticTimelineItem) SetCheckId(v string) {
	o.CheckId = v
}

// GetSeverity returns the Severity field value.
func (o *DiagnosticTimelineItem) GetSeverity() DiagnosticCheckSeverity {
	if o == nil {
		var ret DiagnosticCheckSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTimelineItem) GetSeverityOk() (*DiagnosticCheckSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value.
func (o *DiagnosticTimelineItem) SetSeverity(v DiagnosticCheckSeverity) {
	o.Severity = v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticTimelineItem) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTimelineItem) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticTimelineItem) SetSummary(v string) {
	o.Summary = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *DiagnosticTimelineItem) GetReasonCode() string {
	if o == nil || o.ReasonCode == nil {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTimelineItem) GetReasonCodeOk() (*string, bool) {
	if o == nil || o.ReasonCode == nil {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *DiagnosticTimelineItem) HasReasonCode() bool {
	return o != nil && o.ReasonCode != nil
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *DiagnosticTimelineItem) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticTimelineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OccurredAt.Nanosecond() == 0 {
		toSerialize["occurredAt"] = o.OccurredAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["occurredAt"] = o.OccurredAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ObservedAt.Nanosecond() == 0 {
		toSerialize["observedAt"] = o.ObservedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["observedAt"] = o.ObservedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["timeQuality"] = o.TimeQuality
	toSerialize["source"] = o.Source
	toSerialize["resourceRef"] = o.ResourceRef
	toSerialize["checkId"] = o.CheckId
	toSerialize["severity"] = o.Severity
	toSerialize["summary"] = o.Summary
	if o.ReasonCode != nil {
		toSerialize["reasonCode"] = o.ReasonCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticTimelineItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OccurredAt  *time.Time                     `json:"occurredAt"`
		ObservedAt  *time.Time                     `json:"observedAt"`
		TimeQuality *DiagnosticTimelineTimeQuality `json:"timeQuality"`
		Source      *DiagnosticTimelineSource      `json:"source"`
		ResourceRef *string                        `json:"resourceRef"`
		CheckId     *string                        `json:"checkId"`
		Severity    *DiagnosticCheckSeverity       `json:"severity"`
		Summary     *string                        `json:"summary"`
		ReasonCode  *string                        `json:"reasonCode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OccurredAt == nil {
		return fmt.Errorf("required field occurredAt missing")
	}
	if all.ObservedAt == nil {
		return fmt.Errorf("required field observedAt missing")
	}
	if all.TimeQuality == nil {
		return fmt.Errorf("required field timeQuality missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.ResourceRef == nil {
		return fmt.Errorf("required field resourceRef missing")
	}
	if all.CheckId == nil {
		return fmt.Errorf("required field checkId missing")
	}
	if all.Severity == nil {
		return fmt.Errorf("required field severity missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"occurredAt", "observedAt", "timeQuality", "source", "resourceRef", "checkId", "severity", "summary", "reasonCode"})
	} else {
		return err
	}

	hasInvalidField := false
	o.OccurredAt = *all.OccurredAt
	o.ObservedAt = *all.ObservedAt
	if !all.TimeQuality.IsValid() {
		hasInvalidField = true
	} else {
		o.TimeQuality = *all.TimeQuality
	}
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.ResourceRef = *all.ResourceRef
	o.CheckId = *all.CheckId
	if !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = *all.Severity
	}
	o.Summary = *all.Summary
	o.ReasonCode = all.ReasonCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
