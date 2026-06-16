// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlLockAnalysis struct {
	SelectedSession  PostgresqlSession          `json:"selectedSession"`
	BlockingSessions []PostgresqlSession        `json:"blockingSessions"`
	BlockedSessions  []PostgresqlSession        `json:"blockedSessions"`
	LockRows         []PostgresqlLockRow        `json:"lockRows"`
	WaitChain        PostgresqlWaitChain        `json:"waitChain"`
	Deadlock         PostgresqlDeadlockEvidence `json:"deadlock"`
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
func NewPostgresqlLockAnalysis(selectedSession PostgresqlSession, blockingSessions []PostgresqlSession, blockedSessions []PostgresqlSession, lockRows []PostgresqlLockRow, waitChain PostgresqlWaitChain, deadlock PostgresqlDeadlockEvidence, cannotProve bool, warnings []string, capturedAt string) *PostgresqlLockAnalysis {
	this := PostgresqlLockAnalysis{}
	this.SelectedSession = selectedSession
	this.BlockingSessions = blockingSessions
	this.BlockedSessions = blockedSessions
	this.LockRows = lockRows
	this.WaitChain = waitChain
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

// GetBlockingSessions returns the BlockingSessions field value.
func (o *PostgresqlLockAnalysis) GetBlockingSessions() []PostgresqlSession {
	if o == nil {
		var ret []PostgresqlSession
		return ret
	}
	return o.BlockingSessions
}

// GetBlockingSessionsOk returns a tuple with the BlockingSessions field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetBlockingSessionsOk() (*[]PostgresqlSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockingSessions, true
}

// SetBlockingSessions sets field value.
func (o *PostgresqlLockAnalysis) SetBlockingSessions(v []PostgresqlSession) {
	o.BlockingSessions = v
}

// GetBlockedSessions returns the BlockedSessions field value.
func (o *PostgresqlLockAnalysis) GetBlockedSessions() []PostgresqlSession {
	if o == nil {
		var ret []PostgresqlSession
		return ret
	}
	return o.BlockedSessions
}

// GetBlockedSessionsOk returns a tuple with the BlockedSessions field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetBlockedSessionsOk() (*[]PostgresqlSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockedSessions, true
}

// SetBlockedSessions sets field value.
func (o *PostgresqlLockAnalysis) SetBlockedSessions(v []PostgresqlSession) {
	o.BlockedSessions = v
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

// GetWaitChain returns the WaitChain field value.
func (o *PostgresqlLockAnalysis) GetWaitChain() PostgresqlWaitChain {
	if o == nil {
		var ret PostgresqlWaitChain
		return ret
	}
	return o.WaitChain
}

// GetWaitChainOk returns a tuple with the WaitChain field value
// and a boolean to check if the value has been set.
func (o *PostgresqlLockAnalysis) GetWaitChainOk() (*PostgresqlWaitChain, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitChain, true
}

// SetWaitChain sets field value.
func (o *PostgresqlLockAnalysis) SetWaitChain(v PostgresqlWaitChain) {
	o.WaitChain = v
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
	toSerialize["blockingSessions"] = o.BlockingSessions
	toSerialize["blockedSessions"] = o.BlockedSessions
	toSerialize["lockRows"] = o.LockRows
	toSerialize["waitChain"] = o.WaitChain
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
		SelectedSession  *PostgresqlSession          `json:"selectedSession"`
		BlockingSessions *[]PostgresqlSession        `json:"blockingSessions"`
		BlockedSessions  *[]PostgresqlSession        `json:"blockedSessions"`
		LockRows         *[]PostgresqlLockRow        `json:"lockRows"`
		WaitChain        *PostgresqlWaitChain        `json:"waitChain"`
		Deadlock         *PostgresqlDeadlockEvidence `json:"deadlock"`
		CannotProve      *bool                       `json:"cannotProve"`
		Warnings         *[]string                   `json:"warnings"`
		CapturedAt       *string                     `json:"capturedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SelectedSession == nil {
		return fmt.Errorf("required field selectedSession missing")
	}
	if all.BlockingSessions == nil {
		return fmt.Errorf("required field blockingSessions missing")
	}
	if all.BlockedSessions == nil {
		return fmt.Errorf("required field blockedSessions missing")
	}
	if all.LockRows == nil {
		return fmt.Errorf("required field lockRows missing")
	}
	if all.WaitChain == nil {
		return fmt.Errorf("required field waitChain missing")
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
		common.DeleteKeys(additionalProperties, &[]string{"selectedSession", "blockingSessions", "blockedSessions", "lockRows", "waitChain", "deadlock", "cannotProve", "warnings", "capturedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.SelectedSession.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SelectedSession = *all.SelectedSession
	o.BlockingSessions = *all.BlockingSessions
	o.BlockedSessions = *all.BlockedSessions
	o.LockRows = *all.LockRows
	if all.WaitChain.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.WaitChain = *all.WaitChain
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
