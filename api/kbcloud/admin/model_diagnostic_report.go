// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticReport Layered diagnostic report.
type DiagnosticReport struct {
	// Diagnostic task ID.
	TaskId string `json:"taskId"`
	// Report generation time.
	GeneratedAt time.Time           `json:"generatedAt"`
	Findings    []DiagnosticFinding `json:"findings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticReport instantiates a new DiagnosticReport object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticReport(taskId string, generatedAt time.Time, findings []DiagnosticFinding) *DiagnosticReport {
	this := DiagnosticReport{}
	this.TaskId = taskId
	this.GeneratedAt = generatedAt
	this.Findings = findings
	return &this
}

// NewDiagnosticReportWithDefaults instantiates a new DiagnosticReport object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticReportWithDefaults() *DiagnosticReport {
	this := DiagnosticReport{}
	return &this
}

// GetTaskId returns the TaskId field value.
func (o *DiagnosticReport) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value.
func (o *DiagnosticReport) SetTaskId(v string) {
	o.TaskId = v
}

// GetGeneratedAt returns the GeneratedAt field value.
func (o *DiagnosticReport) GetGeneratedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.GeneratedAt
}

// GetGeneratedAtOk returns a tuple with the GeneratedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetGeneratedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GeneratedAt, true
}

// SetGeneratedAt sets field value.
func (o *DiagnosticReport) SetGeneratedAt(v time.Time) {
	o.GeneratedAt = v
}

// GetFindings returns the Findings field value.
func (o *DiagnosticReport) GetFindings() []DiagnosticFinding {
	if o == nil {
		var ret []DiagnosticFinding
		return ret
	}
	return o.Findings
}

// GetFindingsOk returns a tuple with the Findings field value
// and a boolean to check if the value has been set.
func (o *DiagnosticReport) GetFindingsOk() (*[]DiagnosticFinding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Findings, true
}

// SetFindings sets field value.
func (o *DiagnosticReport) SetFindings(v []DiagnosticFinding) {
	o.Findings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["taskId"] = o.TaskId
	if o.GeneratedAt.Nanosecond() == 0 {
		toSerialize["generatedAt"] = o.GeneratedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["generatedAt"] = o.GeneratedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["findings"] = o.Findings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticReport) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId      *string              `json:"taskId"`
		GeneratedAt *time.Time           `json:"generatedAt"`
		Findings    *[]DiagnosticFinding `json:"findings"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TaskId == nil {
		return fmt.Errorf("required field taskId missing")
	}
	if all.GeneratedAt == nil {
		return fmt.Errorf("required field generatedAt missing")
	}
	if all.Findings == nil {
		return fmt.Errorf("required field findings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskId", "generatedAt", "findings"})
	} else {
		return err
	}
	o.TaskId = *all.TaskId
	o.GeneratedAt = *all.GeneratedAt
	o.Findings = *all.Findings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
