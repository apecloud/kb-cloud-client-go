// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DiagnosticsSessionsReport struct {
	Engine             string                          `json:"engine"`
	Status             DiagnosticsSessionsReportStatus `json:"status"`
	CreatedAt          time.Time                       `json:"createdAt"`
	ReportType         DiagnosticsSessionsReportType   `json:"reportType"`
	ReasonCode         string                          `json:"reasonCode"`
	Summary            string                          `json:"summary"`
	PostgresqlSessions PostgresqlSessionsPayload       `json:"postgresqlSessions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticsSessionsReport instantiates a new DiagnosticsSessionsReport object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticsSessionsReport(engine string, status DiagnosticsSessionsReportStatus, createdAt time.Time, reportType DiagnosticsSessionsReportType, reasonCode string, summary string, postgresqlSessions PostgresqlSessionsPayload) *DiagnosticsSessionsReport {
	this := DiagnosticsSessionsReport{}
	this.Engine = engine
	this.Status = status
	this.CreatedAt = createdAt
	this.ReportType = reportType
	this.ReasonCode = reasonCode
	this.Summary = summary
	this.PostgresqlSessions = postgresqlSessions
	return &this
}

// NewDiagnosticsSessionsReportWithDefaults instantiates a new DiagnosticsSessionsReport object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticsSessionsReportWithDefaults() *DiagnosticsSessionsReport {
	this := DiagnosticsSessionsReport{}
	return &this
}

// GetEngine returns the Engine field value.
func (o *DiagnosticsSessionsReport) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *DiagnosticsSessionsReport) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *DiagnosticsSessionsReport) SetEngine(v string) {
	o.Engine = v
}

// GetStatus returns the Status field value.
func (o *DiagnosticsSessionsReport) GetStatus() DiagnosticsSessionsReportStatus {
	if o == nil {
		var ret DiagnosticsSessionsReportStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DiagnosticsSessionsReport) GetStatusOk() (*DiagnosticsSessionsReportStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DiagnosticsSessionsReport) SetStatus(v DiagnosticsSessionsReportStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DiagnosticsSessionsReport) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticsSessionsReport) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DiagnosticsSessionsReport) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetReportType returns the ReportType field value.
func (o *DiagnosticsSessionsReport) GetReportType() DiagnosticsSessionsReportType {
	if o == nil {
		var ret DiagnosticsSessionsReportType
		return ret
	}
	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *DiagnosticsSessionsReport) GetReportTypeOk() (*DiagnosticsSessionsReportType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value.
func (o *DiagnosticsSessionsReport) SetReportType(v DiagnosticsSessionsReportType) {
	o.ReportType = v
}

// GetReasonCode returns the ReasonCode field value.
func (o *DiagnosticsSessionsReport) GetReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *DiagnosticsSessionsReport) GetReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value.
func (o *DiagnosticsSessionsReport) SetReasonCode(v string) {
	o.ReasonCode = v
}

// GetSummary returns the Summary field value.
func (o *DiagnosticsSessionsReport) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticsSessionsReport) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticsSessionsReport) SetSummary(v string) {
	o.Summary = v
}

// GetPostgresqlSessions returns the PostgresqlSessions field value.
func (o *DiagnosticsSessionsReport) GetPostgresqlSessions() PostgresqlSessionsPayload {
	if o == nil {
		var ret PostgresqlSessionsPayload
		return ret
	}
	return o.PostgresqlSessions
}

// GetPostgresqlSessionsOk returns a tuple with the PostgresqlSessions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticsSessionsReport) GetPostgresqlSessionsOk() (*PostgresqlSessionsPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostgresqlSessions, true
}

// SetPostgresqlSessions sets field value.
func (o *DiagnosticsSessionsReport) SetPostgresqlSessions(v PostgresqlSessionsPayload) {
	o.PostgresqlSessions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticsSessionsReport) MarshalJSON() ([]byte, error) {
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
	toSerialize["postgresqlSessions"] = o.PostgresqlSessions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticsSessionsReport) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Engine             *string                          `json:"engine"`
		Status             *DiagnosticsSessionsReportStatus `json:"status"`
		CreatedAt          *time.Time                       `json:"createdAt"`
		ReportType         *DiagnosticsSessionsReportType   `json:"reportType"`
		ReasonCode         *string                          `json:"reasonCode"`
		Summary            *string                          `json:"summary"`
		PostgresqlSessions *PostgresqlSessionsPayload       `json:"postgresqlSessions"`
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
	if all.PostgresqlSessions == nil {
		return fmt.Errorf("required field postgresqlSessions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engine", "status", "createdAt", "reportType", "reasonCode", "summary", "postgresqlSessions"})
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
	if all.PostgresqlSessions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PostgresqlSessions = *all.PostgresqlSessions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
