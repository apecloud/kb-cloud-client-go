// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlLockAnalysis struct {
	SelectedSession PostgresqlSession `json:"selectedSession"`
	// Sessions that directly block the selected session in the current snapshot.
	DirectBlockingSessions []PostgresqlSession `json:"directBlockingSessions"`
	// Sessions that are directly blocked by the selected session in the current snapshot.
	DirectBlockedSessions []PostgresqlSession `json:"directBlockedSessions"`
	LockRows              []PostgresqlLockRow `json:"lockRows"`
	// Wait graph related to the selected PID. It contains the selected session's ancestor and descendant wait edges that the backend can prove from the current snapshot; unrelated side branches are intentionally excluded.
	WaitGraph PostgresqlWaitGraph        `json:"waitGraph"`
	Deadlock  PostgresqlDeadlockEvidence `json:"deadlock"`
	// Whether the current snapshot cannot prove the lock relationship conclusively.
	CannotProve bool     `json:"cannotProve"`
	Warnings    []string `json:"warnings"`
	// Snapshot capture timestamp in UTC.
	CapturedAt string `json:"capturedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlLockAnalysis instantiates a new PostgresqlLockAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlLockAnalysis(selectedSession PostgresqlSession, directBlockingSessions []PostgresqlSession, directBlockedSessions []PostgresqlSession, lockRows []PostgresqlLockRow, waitGraph PostgresqlWaitGraph, deadlock PostgresqlDeadlockEvidence, cannotProve bool, warnings []string, capturedAt string) *PostgresqlLockAnalysis {
	this := PostgresqlLockAnalysis{}
	this.SelectedSession = selectedSession
	this.DirectBlockingSessions = directBlockingSessions
	this.DirectBlockedSessions = directBlockedSessions
	this.LockRows = lockRows
	this.WaitGraph = waitGraph
	this.Deadlock = deadlock
	this.CannotProve = cannotProve
	this.Warnings = warnings
	this.CapturedAt = capturedAt
	return &this
}

// NewPostgresqlLockAnalysisWithDefaults instantiates a new PostgresqlLockAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlLockAnalysisWithDefaults() *PostgresqlLockAnalysis {
	this := PostgresqlLockAnalysis{}
	return &this
}

// GetSelectedSession returns the SelectedSession field value.
func (o *PostgresqlLockAnalysis) GetSelectedSession() PostgresqlSession {
	if o == nil {
		var ret PostgresqlSession
		return ret
	}
	return o.SelectedSession
}

// GetSelectedSessionOk returns a tuple with the SelectedSession field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetSelectedSessionOk() (*PostgresqlSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedSession, true
}

// SetSelectedSession sets field value.
func (o *PostgresqlLockAnalysis) SetSelectedSession(v PostgresqlSession) {
	o.SelectedSession = v
}

// GetDirectBlockingSessions returns the DirectBlockingSessions field value.
func (o *PostgresqlLockAnalysis) GetDirectBlockingSessions() []PostgresqlSession {
	if o == nil {
		var ret []PostgresqlSession
		return ret
	}
	return o.DirectBlockingSessions
}

// GetDirectBlockingSessionsOk returns a tuple with the DirectBlockingSessions field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetDirectBlockingSessionsOk() (*[]PostgresqlSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectBlockingSessions, true
}

// SetDirectBlockingSessions sets field value.
func (o *PostgresqlLockAnalysis) SetDirectBlockingSessions(v []PostgresqlSession) {
	o.DirectBlockingSessions = v
}

// GetDirectBlockedSessions returns the DirectBlockedSessions field value.
func (o *PostgresqlLockAnalysis) GetDirectBlockedSessions() []PostgresqlSession {
	if o == nil {
		var ret []PostgresqlSession
		return ret
	}
	return o.DirectBlockedSessions
}

// GetDirectBlockedSessionsOk returns a tuple with the DirectBlockedSessions field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetDirectBlockedSessionsOk() (*[]PostgresqlSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectBlockedSessions, true
}

// SetDirectBlockedSessions sets field value.
func (o *PostgresqlLockAnalysis) SetDirectBlockedSessions(v []PostgresqlSession) {
	o.DirectBlockedSessions = v
}

// GetLockRows returns the LockRows field value.
func (o *PostgresqlLockAnalysis) GetLockRows() []PostgresqlLockRow {
	if o == nil {
		var ret []PostgresqlLockRow
		return ret
	}
	return o.LockRows
}

// GetLockRowsOk returns a tuple with the LockRows field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetLockRowsOk() (*[]PostgresqlLockRow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockRows, true
}

// SetLockRows sets field value.
func (o *PostgresqlLockAnalysis) SetLockRows(v []PostgresqlLockRow) {
	o.LockRows = v
}

// GetWaitGraph returns the WaitGraph field value.
func (o *PostgresqlLockAnalysis) GetWaitGraph() PostgresqlWaitGraph {
	if o == nil {
		var ret PostgresqlWaitGraph
		return ret
	}
	return o.WaitGraph
}

// GetWaitGraphOk returns a tuple with the WaitGraph field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetWaitGraphOk() (*PostgresqlWaitGraph, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitGraph, true
}

// SetWaitGraph sets field value.
func (o *PostgresqlLockAnalysis) SetWaitGraph(v PostgresqlWaitGraph) {
	o.WaitGraph = v
}

// GetDeadlock returns the Deadlock field value.
func (o *PostgresqlLockAnalysis) GetDeadlock() PostgresqlDeadlockEvidence {
	if o == nil {
		var ret PostgresqlDeadlockEvidence
		return ret
	}
	return o.Deadlock
}

// GetDeadlockOk returns a tuple with the Deadlock field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetDeadlockOk() (*PostgresqlDeadlockEvidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deadlock, true
}

// SetDeadlock sets field value.
func (o *PostgresqlLockAnalysis) SetDeadlock(v PostgresqlDeadlockEvidence) {
	o.Deadlock = v
}

// GetCannotProve returns the CannotProve field value.
func (o *PostgresqlLockAnalysis) GetCannotProve() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CannotProve
}

// GetCannotProveOk returns a tuple with the CannotProve field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetCannotProveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CannotProve, true
}

// SetCannotProve sets field value.
func (o *PostgresqlLockAnalysis) SetCannotProve(v bool) {
	o.CannotProve = v
}

// GetWarnings returns the Warnings field value.
func (o *PostgresqlLockAnalysis) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetWarningsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value.
func (o *PostgresqlLockAnalysis) SetWarnings(v []string) {
	o.Warnings = v
}

// GetCapturedAt returns the CapturedAt field value.
func (o *PostgresqlLockAnalysis) GetCapturedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CapturedAt
}

// GetCapturedAtOk returns a tuple with the CapturedAt field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetCapturedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAt, true
}

// SetCapturedAt sets field value.
func (o *PostgresqlLockAnalysis) SetCapturedAt(v string) {
	o.CapturedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlLockAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["selectedSession"] = o.SelectedSession
	toSerialize["directBlockingSessions"] = o.DirectBlockingSessions
	toSerialize["directBlockedSessions"] = o.DirectBlockedSessions
	toSerialize["lockRows"] = o.LockRows
	toSerialize["waitGraph"] = o.WaitGraph
	toSerialize["deadlock"] = o.Deadlock
	toSerialize["cannotProve"] = o.CannotProve
	toSerialize["warnings"] = o.Warnings
	toSerialize["capturedAt"] = o.CapturedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlLockAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SelectedSession        *PostgresqlSession          `json:"selectedSession"`
		DirectBlockingSessions *[]PostgresqlSession        `json:"directBlockingSessions"`
		DirectBlockedSessions  *[]PostgresqlSession        `json:"directBlockedSessions"`
		LockRows               *[]PostgresqlLockRow        `json:"lockRows"`
		WaitGraph              *PostgresqlWaitGraph        `json:"waitGraph"`
		Deadlock               *PostgresqlDeadlockEvidence `json:"deadlock"`
		CannotProve            *bool                       `json:"cannotProve"`
		Warnings               *[]string                   `json:"warnings"`
		CapturedAt             *string                     `json:"capturedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SelectedSession == nil {
		return fmt.Errorf("required field selectedSession missing")
	}
	if all.DirectBlockingSessions == nil {
		return fmt.Errorf("required field directBlockingSessions missing")
	}
	if all.DirectBlockedSessions == nil {
		return fmt.Errorf("required field directBlockedSessions missing")
	}
	if all.LockRows == nil {
		return fmt.Errorf("required field lockRows missing")
	}
	if all.WaitGraph == nil {
		return fmt.Errorf("required field waitGraph missing")
	}
	if all.Deadlock == nil {
		return fmt.Errorf("required field deadlock missing")
	}
	if all.CannotProve == nil {
		return fmt.Errorf("required field cannotProve missing")
	}
	if all.Warnings == nil {
		return fmt.Errorf("required field warnings missing")
	}
	if all.CapturedAt == nil {
		return fmt.Errorf("required field capturedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"selectedSession", "directBlockingSessions", "directBlockedSessions", "lockRows", "waitGraph", "deadlock", "cannotProve", "warnings", "capturedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.SelectedSession.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SelectedSession = *all.SelectedSession
	o.DirectBlockingSessions = *all.DirectBlockingSessions
	o.DirectBlockedSessions = *all.DirectBlockedSessions
	o.LockRows = *all.LockRows
	if all.WaitGraph.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.WaitGraph = *all.WaitGraph
	if all.Deadlock.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Deadlock = *all.Deadlock
	o.CannotProve = *all.CannotProve
	o.Warnings = *all.Warnings
	o.CapturedAt = *all.CapturedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
