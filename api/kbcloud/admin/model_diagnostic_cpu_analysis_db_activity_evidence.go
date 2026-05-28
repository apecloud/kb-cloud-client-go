// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysisDBActivityEvidence PostgreSQL session/activity evidence from readonly system views.
type DiagnosticCPUAnalysisDBActivityEvidence struct {
	// PostgreSQL database used for collection.
	DatabaseName string `json:"databaseName"`
	// Connections visible in pg_stat_activity, excluding collector backend.
	TotalConnections int64 `json:"totalConnections"`
	// Active client sessions, excluding collector backend.
	ActiveSessions int64 `json:"activeSessions"`
	// Sessions with a wait_event_type.
	WaitingSessions int64 `json:"waitingSessions"`
	// Background worker sessions observed in pg_stat_activity.
	BackgroundWorkers int64 `json:"backgroundWorkers"`
	// Redacted wait-event buckets by type/event/count.
	WaitEventSummary []string `json:"waitEventSummary"`
	// Collection timestamp.
	CollectedAt time.Time `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticCPUAnalysisDBActivityEvidence instantiates a new DiagnosticCPUAnalysisDBActivityEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticCPUAnalysisDBActivityEvidence(databaseName string, totalConnections int64, activeSessions int64, waitingSessions int64, backgroundWorkers int64, waitEventSummary []string, collectedAt time.Time) *DiagnosticCPUAnalysisDBActivityEvidence {
	this := DiagnosticCPUAnalysisDBActivityEvidence{}
	this.DatabaseName = databaseName
	this.TotalConnections = totalConnections
	this.ActiveSessions = activeSessions
	this.WaitingSessions = waitingSessions
	this.BackgroundWorkers = backgroundWorkers
	this.WaitEventSummary = waitEventSummary
	this.CollectedAt = collectedAt
	return &this
}

// NewDiagnosticCPUAnalysisDBActivityEvidenceWithDefaults instantiates a new DiagnosticCPUAnalysisDBActivityEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticCPUAnalysisDBActivityEvidenceWithDefaults() *DiagnosticCPUAnalysisDBActivityEvidence {
	this := DiagnosticCPUAnalysisDBActivityEvidence{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetTotalConnections returns the TotalConnections field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetTotalConnections() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalConnections
}

// GetTotalConnectionsOk returns a tuple with the TotalConnections field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetTotalConnectionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalConnections, true
}

// SetTotalConnections sets field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) SetTotalConnections(v int64) {
	o.TotalConnections = v
}

// GetActiveSessions returns the ActiveSessions field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetActiveSessions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ActiveSessions
}

// GetActiveSessionsOk returns a tuple with the ActiveSessions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetActiveSessionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveSessions, true
}

// SetActiveSessions sets field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) SetActiveSessions(v int64) {
	o.ActiveSessions = v
}

// GetWaitingSessions returns the WaitingSessions field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetWaitingSessions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.WaitingSessions
}

// GetWaitingSessionsOk returns a tuple with the WaitingSessions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetWaitingSessionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitingSessions, true
}

// SetWaitingSessions sets field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) SetWaitingSessions(v int64) {
	o.WaitingSessions = v
}

// GetBackgroundWorkers returns the BackgroundWorkers field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetBackgroundWorkers() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BackgroundWorkers
}

// GetBackgroundWorkersOk returns a tuple with the BackgroundWorkers field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetBackgroundWorkersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackgroundWorkers, true
}

// SetBackgroundWorkers sets field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) SetBackgroundWorkers(v int64) {
	o.BackgroundWorkers = v
}

// GetWaitEventSummary returns the WaitEventSummary field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetWaitEventSummary() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.WaitEventSummary
}

// GetWaitEventSummaryOk returns a tuple with the WaitEventSummary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetWaitEventSummaryOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitEventSummary, true
}

// SetWaitEventSummary sets field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) SetWaitEventSummary(v []string) {
	o.WaitEventSummary = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticCPUAnalysisDBActivityEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["databaseName"] = o.DatabaseName
	toSerialize["totalConnections"] = o.TotalConnections
	toSerialize["activeSessions"] = o.ActiveSessions
	toSerialize["waitingSessions"] = o.WaitingSessions
	toSerialize["backgroundWorkers"] = o.BackgroundWorkers
	toSerialize["waitEventSummary"] = o.WaitEventSummary
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticCPUAnalysisDBActivityEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatabaseName      *string    `json:"databaseName"`
		TotalConnections  *int64     `json:"totalConnections"`
		ActiveSessions    *int64     `json:"activeSessions"`
		WaitingSessions   *int64     `json:"waitingSessions"`
		BackgroundWorkers *int64     `json:"backgroundWorkers"`
		WaitEventSummary  *[]string  `json:"waitEventSummary"`
		CollectedAt       *time.Time `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DatabaseName == nil {
		return fmt.Errorf("required field databaseName missing")
	}
	if all.TotalConnections == nil {
		return fmt.Errorf("required field totalConnections missing")
	}
	if all.ActiveSessions == nil {
		return fmt.Errorf("required field activeSessions missing")
	}
	if all.WaitingSessions == nil {
		return fmt.Errorf("required field waitingSessions missing")
	}
	if all.BackgroundWorkers == nil {
		return fmt.Errorf("required field backgroundWorkers missing")
	}
	if all.WaitEventSummary == nil {
		return fmt.Errorf("required field waitEventSummary missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"databaseName", "totalConnections", "activeSessions", "waitingSessions", "backgroundWorkers", "waitEventSummary", "collectedAt"})
	} else {
		return err
	}
	o.DatabaseName = *all.DatabaseName
	o.TotalConnections = *all.TotalConnections
	o.ActiveSessions = *all.ActiveSessions
	o.WaitingSessions = *all.WaitingSessions
	o.BackgroundWorkers = *all.BackgroundWorkers
	o.WaitEventSummary = *all.WaitEventSummary
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
