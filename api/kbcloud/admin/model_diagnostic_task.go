// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticTask Diagnostic task summary.
type DiagnosticTask struct {
	// Diagnostic task ID.
	TaskId string `json:"taskId"`
	// Organization name.
	OrgName string `json:"orgName"`
	// Cluster name.
	ClusterName string `json:"clusterName"`
	// Diagnostic task status.
	Status DiagnosticTaskStatus `json:"status"`
	// Progress percentage.
	Progress int64 `json:"progress"`
	// Failure or cancellation reason.
	Reason *string `json:"reason,omitempty"`
	// Creation time.
	CreatedAt time.Time `json:"createdAt"`
	// Last update time.
	UpdatedAt time.Time `json:"updatedAt"`
	// Start time.
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// Finish time.
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// Layered diagnostic report.
	Report *DiagnosticReport `json:"report,omitempty"`
	// Manifest for a redacted diagnostic package.
	ExportManifest *DiagnosticExportManifest `json:"exportManifest,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticTask instantiates a new DiagnosticTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticTask(taskId string, orgName string, clusterName string, status DiagnosticTaskStatus, progress int64, createdAt time.Time, updatedAt time.Time) *DiagnosticTask {
	this := DiagnosticTask{}
	this.TaskId = taskId
	this.OrgName = orgName
	this.ClusterName = clusterName
	this.Status = status
	this.Progress = progress
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewDiagnosticTaskWithDefaults instantiates a new DiagnosticTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticTaskWithDefaults() *DiagnosticTask {
	this := DiagnosticTask{}
	return &this
}

// GetTaskId returns the TaskId field value.
func (o *DiagnosticTask) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *DiagnosticTask) SetTaskId(v string) {
	o.TaskId = v
}

// GetOrgName returns the OrgName field value.
func (o *DiagnosticTask) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value.
func (o *DiagnosticTask) SetOrgName(v string) {
	o.OrgName = v
}

// GetClusterName returns the ClusterName field value.
func (o *DiagnosticTask) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value.
func (o *DiagnosticTask) SetClusterName(v string) {
	o.ClusterName = v
}

// GetStatus returns the Status field value.
func (o *DiagnosticTask) GetStatus() DiagnosticTaskStatus {
	if o == nil {
		var ret DiagnosticTaskStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetStatusOk() (*DiagnosticTaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DiagnosticTask) SetStatus(v DiagnosticTaskStatus) {
	o.Status = v
}

// GetProgress returns the Progress field value.
func (o *DiagnosticTask) GetProgress() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetProgressOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value.
func (o *DiagnosticTask) SetProgress(v int64) {
	o.Progress = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DiagnosticTask) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DiagnosticTask) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DiagnosticTask) SetReason(v string) {
	o.Reason = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *DiagnosticTask) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *DiagnosticTask) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *DiagnosticTask) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *DiagnosticTask) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *DiagnosticTask) GetStartedAt() time.Time {
	if o == nil || o.StartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DiagnosticTask) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *DiagnosticTask) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *DiagnosticTask) GetFinishedAt() time.Time {
	if o == nil || o.FinishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *DiagnosticTask) HasFinishedAt() bool {
	return o != nil && o.FinishedAt != nil
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *DiagnosticTask) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *DiagnosticTask) GetReport() DiagnosticReport {
	if o == nil || o.Report == nil {
		var ret DiagnosticReport
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetReportOk() (*DiagnosticReport, bool) {
	if o == nil || o.Report == nil {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *DiagnosticTask) HasReport() bool {
	return o != nil && o.Report != nil
}

// SetReport gets a reference to the given DiagnosticReport and assigns it to the Report field.
func (o *DiagnosticTask) SetReport(v DiagnosticReport) {
	o.Report = &v
}

// GetExportManifest returns the ExportManifest field value if set, zero value otherwise.
func (o *DiagnosticTask) GetExportManifest() DiagnosticExportManifest {
	if o == nil || o.ExportManifest == nil {
		var ret DiagnosticExportManifest
		return ret
	}
	return *o.ExportManifest
}

// GetExportManifestOk returns a tuple with the ExportManifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTask) GetExportManifestOk() (*DiagnosticExportManifest, bool) {
	if o == nil || o.ExportManifest == nil {
		return nil, false
	}
	return o.ExportManifest, true
}

// HasExportManifest returns a boolean if a field has been set.
func (o *DiagnosticTask) HasExportManifest() bool {
	return o != nil && o.ExportManifest != nil
}

// SetExportManifest gets a reference to the given DiagnosticExportManifest and assigns it to the ExportManifest field.
func (o *DiagnosticTask) SetExportManifest(v DiagnosticExportManifest) {
	o.ExportManifest = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId
	toSerialize["orgName"] = o.OrgName
	toSerialize["clusterName"] = o.ClusterName
	toSerialize["status"] = o.Status
	toSerialize["progress"] = o.Progress
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.StartedAt != nil {
		if o.StartedAt.Nanosecond() == 0 {
			toSerialize["startedAt"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startedAt"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.FinishedAt != nil {
		if o.FinishedAt.Nanosecond() == 0 {
			toSerialize["finishedAt"] = o.FinishedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["finishedAt"] = o.FinishedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Report != nil {
		toSerialize["report"] = o.Report
	}
	if o.ExportManifest != nil {
		toSerialize["exportManifest"] = o.ExportManifest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId         *string                   `json:"taskId"`
		OrgName        *string                   `json:"orgName"`
		ClusterName    *string                   `json:"clusterName"`
		Status         *DiagnosticTaskStatus     `json:"status"`
		Progress       *int64                    `json:"progress"`
		Reason         *string                   `json:"reason,omitempty"`
		CreatedAt      *time.Time                `json:"createdAt"`
		UpdatedAt      *time.Time                `json:"updatedAt"`
		StartedAt      *time.Time                `json:"startedAt,omitempty"`
		FinishedAt     *time.Time                `json:"finishedAt,omitempty"`
		Report         *DiagnosticReport         `json:"report,omitempty"`
		ExportManifest *DiagnosticExportManifest `json:"exportManifest,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	if all.OrgName == nil {
		return fmt.Errorf("required field orgName missing")
	}
	if all.ClusterName == nil {
		return fmt.Errorf("required field clusterName missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Progress == nil {
		return fmt.Errorf("required field progress missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskId", "orgName", "clusterName", "status", "progress", "reason", "createdAt", "updatedAt", "startedAt", "finishedAt", "report", "exportManifest"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TaskId = *all.TaskId
	o.OrgName = *all.OrgName
	o.ClusterName = *all.ClusterName
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Progress = *all.Progress
	o.Reason = all.Reason
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt
	o.StartedAt = all.StartedAt
	o.FinishedAt = all.FinishedAt
	if all.Report != nil && all.Report.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Report = all.Report
	if all.ExportManifest != nil && all.ExportManifest.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExportManifest = all.ExportManifest

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
