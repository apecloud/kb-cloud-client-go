// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceLockWaitSample One blocked session and its blocker, redacted for report display.
type DiagnosticDBPerformanceLockWaitSample struct {
	// PostgreSQL backend PID of the blocked session.
	BlockedPid int64 `json:"blockedPid"`
	// PostgreSQL backend PID of the blocking session.
	BlockingPid int64 `json:"blockingPid"`
	// Database user of the blocked session.
	BlockedUser *string `json:"blockedUser,omitempty"`
	// Database user of the blocking session.
	BlockingUser *string `json:"blockingUser,omitempty"`
	// Database name observed for the blocked session.
	DatabaseName *string `json:"databaseName,omitempty"`
	// How long the blocked session has been waiting, in seconds.
	WaitDurationSeconds int64 `json:"waitDurationSeconds"`
	// PostgreSQL wait_event_type for the blocked session.
	WaitEventType *string `json:"waitEventType,omitempty"`
	// PostgreSQL wait_event for the blocked session.
	WaitEvent *string `json:"waitEvent,omitempty"`
	// pg_locks.locktype for the waiting lock.
	LockType string `json:"lockType"`
	// pg_locks.mode for the waiting lock.
	LockMode *string `json:"lockMode,omitempty"`
	// Stable lock object identity, such as relation:<name>, transactionid:<id>, tuple:<relation>:<page>:<tuple>, or advisory:<classid>:<objid>:<objsubid>.
	LockObject string `json:"lockObject"`
	// Digest of the blocked SQL text. Full SQL text is not exposed.
	BlockedQueryDigest string `json:"blockedQueryDigest"`
	// Digest of the blocking SQL text. Full SQL text is not exposed.
	BlockingQueryDigest string `json:"blockingQueryDigest"`
	// Redacted summary of the blocked SQL, such as statement kind plus digest. Full SQL text and literal values are not exposed.
	BlockedQuerySummary string `json:"blockedQuerySummary"`
	// Redacted summary of the blocking SQL, such as statement kind plus digest. Full SQL text and literal values are not exposed.
	BlockingQuerySummary string `json:"blockingQuerySummary"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceLockWaitSample instantiates a new DiagnosticDBPerformanceLockWaitSample object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceLockWaitSample(blockedPid int64, blockingPid int64, waitDurationSeconds int64, lockType string, lockObject string, blockedQueryDigest string, blockingQueryDigest string, blockedQuerySummary string, blockingQuerySummary string) *DiagnosticDBPerformanceLockWaitSample {
	this := DiagnosticDBPerformanceLockWaitSample{}
	this.BlockedPid = blockedPid
	this.BlockingPid = blockingPid
	this.WaitDurationSeconds = waitDurationSeconds
	this.LockType = lockType
	this.LockObject = lockObject
	this.BlockedQueryDigest = blockedQueryDigest
	this.BlockingQueryDigest = blockingQueryDigest
	this.BlockedQuerySummary = blockedQuerySummary
	this.BlockingQuerySummary = blockingQuerySummary
	return &this
}

// NewDiagnosticDBPerformanceLockWaitSampleWithDefaults instantiates a new DiagnosticDBPerformanceLockWaitSample object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceLockWaitSampleWithDefaults() *DiagnosticDBPerformanceLockWaitSample {
	this := DiagnosticDBPerformanceLockWaitSample{}
	return &this
}

// GetBlockedPid returns the BlockedPid field value.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockedPid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BlockedPid
}

// GetBlockedPidOk returns a tuple with the BlockedPid field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockedPidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockedPid, true
}

// SetBlockedPid sets field value.
func (o *DiagnosticDBPerformanceLockWaitSample) SetBlockedPid(v int64) {
	o.BlockedPid = v
}

// GetBlockingPid returns the BlockingPid field value.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockingPid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BlockingPid
}

// GetBlockingPidOk returns a tuple with the BlockingPid field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockingPidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockingPid, true
}

// SetBlockingPid sets field value.
func (o *DiagnosticDBPerformanceLockWaitSample) SetBlockingPid(v int64) {
	o.BlockingPid = v
}

// GetBlockedUser returns the BlockedUser field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockedUser() string {
	if o == nil || o.BlockedUser == nil {
		var ret string
		return ret
	}
	return *o.BlockedUser
}

// GetBlockedUserOk returns a tuple with the BlockedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockedUserOk() (*string, bool) {
	if o == nil || o.BlockedUser == nil {
		return nil, false
	}
	return o.BlockedUser, true
}

// HasBlockedUser returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) HasBlockedUser() bool {
	return o != nil && o.BlockedUser != nil
}

// SetBlockedUser gets a reference to the given string and assigns it to the BlockedUser field.
func (o *DiagnosticDBPerformanceLockWaitSample) SetBlockedUser(v string) {
	o.BlockedUser = &v
}

// GetBlockingUser returns the BlockingUser field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockingUser() string {
	if o == nil || o.BlockingUser == nil {
		var ret string
		return ret
	}
	return *o.BlockingUser
}

// GetBlockingUserOk returns a tuple with the BlockingUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockingUserOk() (*string, bool) {
	if o == nil || o.BlockingUser == nil {
		return nil, false
	}
	return o.BlockingUser, true
}

// HasBlockingUser returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) HasBlockingUser() bool {
	return o != nil && o.BlockingUser != nil
}

// SetBlockingUser gets a reference to the given string and assigns it to the BlockingUser field.
func (o *DiagnosticDBPerformanceLockWaitSample) SetBlockingUser(v string) {
	o.BlockingUser = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitSample) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) HasDatabaseName() bool {
	return o != nil && o.DatabaseName != nil
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *DiagnosticDBPerformanceLockWaitSample) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetWaitDurationSeconds returns the WaitDurationSeconds field value.
func (o *DiagnosticDBPerformanceLockWaitSample) GetWaitDurationSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.WaitDurationSeconds
}

// GetWaitDurationSecondsOk returns a tuple with the WaitDurationSeconds field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetWaitDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitDurationSeconds, true
}

// SetWaitDurationSeconds sets field value.
func (o *DiagnosticDBPerformanceLockWaitSample) SetWaitDurationSeconds(v int64) {
	o.WaitDurationSeconds = v
}

// GetWaitEventType returns the WaitEventType field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitSample) GetWaitEventType() string {
	if o == nil || o.WaitEventType == nil {
		var ret string
		return ret
	}
	return *o.WaitEventType
}

// GetWaitEventTypeOk returns a tuple with the WaitEventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetWaitEventTypeOk() (*string, bool) {
	if o == nil || o.WaitEventType == nil {
		return nil, false
	}
	return o.WaitEventType, true
}

// HasWaitEventType returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) HasWaitEventType() bool {
	return o != nil && o.WaitEventType != nil
}

// SetWaitEventType gets a reference to the given string and assigns it to the WaitEventType field.
func (o *DiagnosticDBPerformanceLockWaitSample) SetWaitEventType(v string) {
	o.WaitEventType = &v
}

// GetWaitEvent returns the WaitEvent field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitSample) GetWaitEvent() string {
	if o == nil || o.WaitEvent == nil {
		var ret string
		return ret
	}
	return *o.WaitEvent
}

// GetWaitEventOk returns a tuple with the WaitEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetWaitEventOk() (*string, bool) {
	if o == nil || o.WaitEvent == nil {
		return nil, false
	}
	return o.WaitEvent, true
}

// HasWaitEvent returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) HasWaitEvent() bool {
	return o != nil && o.WaitEvent != nil
}

// SetWaitEvent gets a reference to the given string and assigns it to the WaitEvent field.
func (o *DiagnosticDBPerformanceLockWaitSample) SetWaitEvent(v string) {
	o.WaitEvent = &v
}

// GetLockType returns the LockType field value.
func (o *DiagnosticDBPerformanceLockWaitSample) GetLockType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LockType
}

// GetLockTypeOk returns a tuple with the LockType field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetLockTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockType, true
}

// SetLockType sets field value.
func (o *DiagnosticDBPerformanceLockWaitSample) SetLockType(v string) {
	o.LockType = v
}

// GetLockMode returns the LockMode field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLockWaitSample) GetLockMode() string {
	if o == nil || o.LockMode == nil {
		var ret string
		return ret
	}
	return *o.LockMode
}

// GetLockModeOk returns a tuple with the LockMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetLockModeOk() (*string, bool) {
	if o == nil || o.LockMode == nil {
		return nil, false
	}
	return o.LockMode, true
}

// HasLockMode returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) HasLockMode() bool {
	return o != nil && o.LockMode != nil
}

// SetLockMode gets a reference to the given string and assigns it to the LockMode field.
func (o *DiagnosticDBPerformanceLockWaitSample) SetLockMode(v string) {
	o.LockMode = &v
}

// GetLockObject returns the LockObject field value.
func (o *DiagnosticDBPerformanceLockWaitSample) GetLockObject() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LockObject
}

// GetLockObjectOk returns a tuple with the LockObject field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetLockObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockObject, true
}

// SetLockObject sets field value.
func (o *DiagnosticDBPerformanceLockWaitSample) SetLockObject(v string) {
	o.LockObject = v
}

// GetBlockedQueryDigest returns the BlockedQueryDigest field value.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockedQueryDigest() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BlockedQueryDigest
}

// GetBlockedQueryDigestOk returns a tuple with the BlockedQueryDigest field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockedQueryDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockedQueryDigest, true
}

// SetBlockedQueryDigest sets field value.
func (o *DiagnosticDBPerformanceLockWaitSample) SetBlockedQueryDigest(v string) {
	o.BlockedQueryDigest = v
}

// GetBlockingQueryDigest returns the BlockingQueryDigest field value.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockingQueryDigest() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BlockingQueryDigest
}

// GetBlockingQueryDigestOk returns a tuple with the BlockingQueryDigest field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockingQueryDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockingQueryDigest, true
}

// SetBlockingQueryDigest sets field value.
func (o *DiagnosticDBPerformanceLockWaitSample) SetBlockingQueryDigest(v string) {
	o.BlockingQueryDigest = v
}

// GetBlockedQuerySummary returns the BlockedQuerySummary field value.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockedQuerySummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BlockedQuerySummary
}

// GetBlockedQuerySummaryOk returns a tuple with the BlockedQuerySummary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockedQuerySummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockedQuerySummary, true
}

// SetBlockedQuerySummary sets field value.
func (o *DiagnosticDBPerformanceLockWaitSample) SetBlockedQuerySummary(v string) {
	o.BlockedQuerySummary = v
}

// GetBlockingQuerySummary returns the BlockingQuerySummary field value.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockingQuerySummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BlockingQuerySummary
}

// GetBlockingQuerySummaryOk returns a tuple with the BlockingQuerySummary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitSample) GetBlockingQuerySummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockingQuerySummary, true
}

// SetBlockingQuerySummary sets field value.
func (o *DiagnosticDBPerformanceLockWaitSample) SetBlockingQuerySummary(v string) {
	o.BlockingQuerySummary = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceLockWaitSample) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["blockedPid"] = o.BlockedPid
	toSerialize["blockingPid"] = o.BlockingPid
	if o.BlockedUser != nil {
		toSerialize["blockedUser"] = o.BlockedUser
	}
	if o.BlockingUser != nil {
		toSerialize["blockingUser"] = o.BlockingUser
	}
	if o.DatabaseName != nil {
		toSerialize["databaseName"] = o.DatabaseName
	}
	toSerialize["waitDurationSeconds"] = o.WaitDurationSeconds
	if o.WaitEventType != nil {
		toSerialize["waitEventType"] = o.WaitEventType
	}
	if o.WaitEvent != nil {
		toSerialize["waitEvent"] = o.WaitEvent
	}
	toSerialize["lockType"] = o.LockType
	if o.LockMode != nil {
		toSerialize["lockMode"] = o.LockMode
	}
	toSerialize["lockObject"] = o.LockObject
	toSerialize["blockedQueryDigest"] = o.BlockedQueryDigest
	toSerialize["blockingQueryDigest"] = o.BlockingQueryDigest
	toSerialize["blockedQuerySummary"] = o.BlockedQuerySummary
	toSerialize["blockingQuerySummary"] = o.BlockingQuerySummary

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceLockWaitSample) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BlockedPid           *int64  `json:"blockedPid"`
		BlockingPid          *int64  `json:"blockingPid"`
		BlockedUser          *string `json:"blockedUser,omitempty"`
		BlockingUser         *string `json:"blockingUser,omitempty"`
		DatabaseName         *string `json:"databaseName,omitempty"`
		WaitDurationSeconds  *int64  `json:"waitDurationSeconds"`
		WaitEventType        *string `json:"waitEventType,omitempty"`
		WaitEvent            *string `json:"waitEvent,omitempty"`
		LockType             *string `json:"lockType"`
		LockMode             *string `json:"lockMode,omitempty"`
		LockObject           *string `json:"lockObject"`
		BlockedQueryDigest   *string `json:"blockedQueryDigest"`
		BlockingQueryDigest  *string `json:"blockingQueryDigest"`
		BlockedQuerySummary  *string `json:"blockedQuerySummary"`
		BlockingQuerySummary *string `json:"blockingQuerySummary"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.BlockedPid == nil {
		return fmt.Errorf("required field blockedPid missing")
	}
	if all.BlockingPid == nil {
		return fmt.Errorf("required field blockingPid missing")
	}
	if all.WaitDurationSeconds == nil {
		return fmt.Errorf("required field waitDurationSeconds missing")
	}
	if all.LockType == nil {
		return fmt.Errorf("required field lockType missing")
	}
	if all.LockObject == nil {
		return fmt.Errorf("required field lockObject missing")
	}
	if all.BlockedQueryDigest == nil {
		return fmt.Errorf("required field blockedQueryDigest missing")
	}
	if all.BlockingQueryDigest == nil {
		return fmt.Errorf("required field blockingQueryDigest missing")
	}
	if all.BlockedQuerySummary == nil {
		return fmt.Errorf("required field blockedQuerySummary missing")
	}
	if all.BlockingQuerySummary == nil {
		return fmt.Errorf("required field blockingQuerySummary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"blockedPid", "blockingPid", "blockedUser", "blockingUser", "databaseName", "waitDurationSeconds", "waitEventType", "waitEvent", "lockType", "lockMode", "lockObject", "blockedQueryDigest", "blockingQueryDigest", "blockedQuerySummary", "blockingQuerySummary"})
	} else {
		return err
	}
	o.BlockedPid = *all.BlockedPid
	o.BlockingPid = *all.BlockingPid
	o.BlockedUser = all.BlockedUser
	o.BlockingUser = all.BlockingUser
	o.DatabaseName = all.DatabaseName
	o.WaitDurationSeconds = *all.WaitDurationSeconds
	o.WaitEventType = all.WaitEventType
	o.WaitEvent = all.WaitEvent
	o.LockType = *all.LockType
	o.LockMode = all.LockMode
	o.LockObject = *all.LockObject
	o.BlockedQueryDigest = *all.BlockedQueryDigest
	o.BlockingQueryDigest = *all.BlockingQueryDigest
	o.BlockedQuerySummary = *all.BlockedQuerySummary
	o.BlockingQuerySummary = *all.BlockingQuerySummary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
