// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlRealtimeSessionsReport struct {
	Engine     string                                 `json:"engine"`
	Status     PostgresqlRealtimeSessionsReportStatus `json:"status"`
	CreatedAt  time.Time                              `json:"createdAt"`
	ReportType PostgresqlRealtimeSessionsReportType   `json:"reportType"`
	ReasonCode string                                 `json:"reasonCode"`
	Summary    string                                 `json:"summary"`
	Stats      PostgresqlRealtimeSessionStats         `json:"stats"`
	Sessions   []PostgresqlRealtimeSession            `json:"sessions"`
	// Redacted plain text context for copy actions.
	CopyContext string `json:"copyContext"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlRealtimeSessionsReport instantiates a new PostgresqlRealtimeSessionsReport object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlRealtimeSessionsReport(engine string, status PostgresqlRealtimeSessionsReportStatus, createdAt time.Time, reportType PostgresqlRealtimeSessionsReportType, reasonCode string, summary string, stats PostgresqlRealtimeSessionStats, sessions []PostgresqlRealtimeSession, copyContext string) *PostgresqlRealtimeSessionsReport {
	this := PostgresqlRealtimeSessionsReport{}
	this.Engine = engine
	this.Status = status
	this.CreatedAt = createdAt
	this.ReportType = reportType
	this.ReasonCode = reasonCode
	this.Summary = summary
	this.Stats = stats
	this.Sessions = sessions
	this.CopyContext = copyContext
	return &this
}

// NewPostgresqlRealtimeSessionsReportWithDefaults instantiates a new PostgresqlRealtimeSessionsReport object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlRealtimeSessionsReportWithDefaults() *PostgresqlRealtimeSessionsReport {
	this := PostgresqlRealtimeSessionsReport{}
	return &this
}

// GetEngine returns the Engine field value.
func (o *PostgresqlRealtimeSessionsReport) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsReport) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *PostgresqlRealtimeSessionsReport) SetEngine(v string) {
	o.Engine = v
}

// GetStatus returns the Status field value.
func (o *PostgresqlRealtimeSessionsReport) GetStatus() PostgresqlRealtimeSessionsReportStatus {
	if o == nil {
		var ret PostgresqlRealtimeSessionsReportStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsReport) GetStatusOk() (*PostgresqlRealtimeSessionsReportStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *PostgresqlRealtimeSessionsReport) SetStatus(v PostgresqlRealtimeSessionsReportStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *PostgresqlRealtimeSessionsReport) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsReport) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *PostgresqlRealtimeSessionsReport) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetReportType returns the ReportType field value.
func (o *PostgresqlRealtimeSessionsReport) GetReportType() PostgresqlRealtimeSessionsReportType {
	if o == nil {
		var ret PostgresqlRealtimeSessionsReportType
		return ret
	}
	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsReport) GetReportTypeOk() (*PostgresqlRealtimeSessionsReportType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value.
func (o *PostgresqlRealtimeSessionsReport) SetReportType(v PostgresqlRealtimeSessionsReportType) {
	o.ReportType = v
}

// GetReasonCode returns the ReasonCode field value.
func (o *PostgresqlRealtimeSessionsReport) GetReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsReport) GetReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value.
func (o *PostgresqlRealtimeSessionsReport) SetReasonCode(v string) {
	o.ReasonCode = v
}

// GetSummary returns the Summary field value.
func (o *PostgresqlRealtimeSessionsReport) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsReport) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *PostgresqlRealtimeSessionsReport) SetSummary(v string) {
	o.Summary = v
}

// GetStats returns the Stats field value.
func (o *PostgresqlRealtimeSessionsReport) GetStats() PostgresqlRealtimeSessionStats {
	if o == nil {
		var ret PostgresqlRealtimeSessionStats
		return ret
	}
	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsReport) GetStatsOk() (*PostgresqlRealtimeSessionStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value.
func (o *PostgresqlRealtimeSessionsReport) SetStats(v PostgresqlRealtimeSessionStats) {
	o.Stats = v
}

// GetSessions returns the Sessions field value.
func (o *PostgresqlRealtimeSessionsReport) GetSessions() []PostgresqlRealtimeSession {
	if o == nil {
		var ret []PostgresqlRealtimeSession
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsReport) GetSessionsOk() (*[]PostgresqlRealtimeSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// SetSessions sets field value.
func (o *PostgresqlRealtimeSessionsReport) SetSessions(v []PostgresqlRealtimeSession) {
	o.Sessions = v
}

// GetCopyContext returns the CopyContext field value.
func (o *PostgresqlRealtimeSessionsReport) GetCopyContext() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CopyContext
}

// GetCopyContextOk returns a tuple with the CopyContext field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsReport) GetCopyContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyContext, true
}

// SetCopyContext sets field value.
func (o *PostgresqlRealtimeSessionsReport) SetCopyContext(v string) {
	o.CopyContext = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlRealtimeSessionsReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engine"] = o.Engine
	toSerialize["status"] = o.Status
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["reportType"] = o.ReportType
	toSerialize["reasonCode"] = o.ReasonCode
	toSerialize["summary"] = o.Summary
	toSerialize["stats"] = o.Stats
	toSerialize["sessions"] = o.Sessions
	toSerialize["copyContext"] = o.CopyContext

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlRealtimeSessionsReport) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Engine      *string                                 `json:"engine"`
		Status      *PostgresqlRealtimeSessionsReportStatus `json:"status"`
		CreatedAt   *time.Time                              `json:"createdAt"`
		ReportType  *PostgresqlRealtimeSessionsReportType   `json:"reportType"`
		ReasonCode  *string                                 `json:"reasonCode"`
		Summary     *string                                 `json:"summary"`
		Stats       *PostgresqlRealtimeSessionStats         `json:"stats"`
		Sessions    *[]PostgresqlRealtimeSession            `json:"sessions"`
		CopyContext *string                                 `json:"copyContext"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.ReportType == nil {
		return fmt.Errorf("required field reportType missing")
	}
	if all.ReasonCode == nil {
		return fmt.Errorf("required field reasonCode missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.Stats == nil {
		return fmt.Errorf("required field stats missing")
	}
	if all.Sessions == nil {
		return fmt.Errorf("required field sessions missing")
	}
	if all.CopyContext == nil {
		return fmt.Errorf("required field copyContext missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engine", "status", "createdAt", "reportType", "reasonCode", "summary", "stats", "sessions", "copyContext"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Engine = *all.Engine
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.CreatedAt = *all.CreatedAt
	if !all.ReportType.IsValid() {
		hasInvalidField = true
	} else {
		o.ReportType = *all.ReportType
	}
	o.ReasonCode = *all.ReasonCode
	o.Summary = *all.Summary
	if all.Stats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Stats = *all.Stats
	o.Sessions = *all.Sessions
	o.CopyContext = *all.CopyContext

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
