// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceLongTransactionSample One active PostgreSQL transaction that exceeded the long-transaction threshold.
type DiagnosticDBPerformanceLongTransactionSample struct {
	// PostgreSQL backend PID for the session holding the long transaction.
	Pid int64 `json:"pid"`
	// Database user of the session.
	UserName string `json:"userName"`
	// Database name observed for the session.
	DatabaseName string `json:"databaseName"`
	// PostgreSQL application_name for the session.
	ApplicationName *string `json:"applicationName,omitempty"`
	// Client address reported by pg_stat_activity, when available.
	ClientAddr *string `json:"clientAddr,omitempty"`
	// pg_stat_activity.xact_start for the active transaction.
	TransactionStart time.Time `json:"transactionStart"`
	// How long the transaction has been open, in seconds.
	TransactionDurationSeconds int64 `json:"transactionDurationSeconds"`
	// pg_stat_activity.state for the session.
	State string `json:"state"`
	// PostgreSQL wait_event_type for the session.
	WaitEventType *string `json:"waitEventType,omitempty"`
	// PostgreSQL wait_event for the session.
	WaitEvent *string `json:"waitEvent,omitempty"`
	// Digest of the current query text. Full SQL text is not exposed.
	QueryDigest string `json:"queryDigest"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceLongTransactionSample instantiates a new DiagnosticDBPerformanceLongTransactionSample object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceLongTransactionSample(pid int64, userName string, databaseName string, transactionStart time.Time, transactionDurationSeconds int64, state string, queryDigest string) *DiagnosticDBPerformanceLongTransactionSample {
	this := DiagnosticDBPerformanceLongTransactionSample{}
	this.Pid = pid
	this.UserName = userName
	this.DatabaseName = databaseName
	this.TransactionStart = transactionStart
	this.TransactionDurationSeconds = transactionDurationSeconds
	this.State = state
	this.QueryDigest = queryDigest
	return &this
}

// NewDiagnosticDBPerformanceLongTransactionSampleWithDefaults instantiates a new DiagnosticDBPerformanceLongTransactionSample object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceLongTransactionSampleWithDefaults() *DiagnosticDBPerformanceLongTransactionSample {
	this := DiagnosticDBPerformanceLongTransactionSample{}
	return &this
}

// GetPid returns the Pid field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetPid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetPidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetPid(v int64) {
	o.Pid = v
}

// GetUserName returns the UserName field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetUserName(v string) {
	o.UserName = v
}

// GetDatabaseName returns the DatabaseName field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) HasApplicationName() bool {
	return o != nil && o.ApplicationName != nil
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetClientAddr returns the ClientAddr field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetClientAddr() string {
	if o == nil || o.ClientAddr == nil {
		var ret string
		return ret
	}
	return *o.ClientAddr
}

// GetClientAddrOk returns a tuple with the ClientAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetClientAddrOk() (*string, bool) {
	if o == nil || o.ClientAddr == nil {
		return nil, false
	}
	return o.ClientAddr, true
}

// HasClientAddr returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) HasClientAddr() bool {
	return o != nil && o.ClientAddr != nil
}

// SetClientAddr gets a reference to the given string and assigns it to the ClientAddr field.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetClientAddr(v string) {
	o.ClientAddr = &v
}

// GetTransactionStart returns the TransactionStart field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetTransactionStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.TransactionStart
}

// GetTransactionStartOk returns a tuple with the TransactionStart field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetTransactionStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionStart, true
}

// SetTransactionStart sets field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetTransactionStart(v time.Time) {
	o.TransactionStart = v
}

// GetTransactionDurationSeconds returns the TransactionDurationSeconds field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetTransactionDurationSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TransactionDurationSeconds
}

// GetTransactionDurationSecondsOk returns a tuple with the TransactionDurationSeconds field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetTransactionDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionDurationSeconds, true
}

// SetTransactionDurationSeconds sets field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetTransactionDurationSeconds(v int64) {
	o.TransactionDurationSeconds = v
}

// GetState returns the State field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetState(v string) {
	o.State = v
}

// GetWaitEventType returns the WaitEventType field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetWaitEventType() string {
	if o == nil || o.WaitEventType == nil {
		var ret string
		return ret
	}
	return *o.WaitEventType
}

// GetWaitEventTypeOk returns a tuple with the WaitEventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetWaitEventTypeOk() (*string, bool) {
	if o == nil || o.WaitEventType == nil {
		return nil, false
	}
	return o.WaitEventType, true
}

// HasWaitEventType returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) HasWaitEventType() bool {
	return o != nil && o.WaitEventType != nil
}

// SetWaitEventType gets a reference to the given string and assigns it to the WaitEventType field.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetWaitEventType(v string) {
	o.WaitEventType = &v
}

// GetWaitEvent returns the WaitEvent field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetWaitEvent() string {
	if o == nil || o.WaitEvent == nil {
		var ret string
		return ret
	}
	return *o.WaitEvent
}

// GetWaitEventOk returns a tuple with the WaitEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetWaitEventOk() (*string, bool) {
	if o == nil || o.WaitEvent == nil {
		return nil, false
	}
	return o.WaitEvent, true
}

// HasWaitEvent returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) HasWaitEvent() bool {
	return o != nil && o.WaitEvent != nil
}

// SetWaitEvent gets a reference to the given string and assigns it to the WaitEvent field.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetWaitEvent(v string) {
	o.WaitEvent = &v
}

// GetQueryDigest returns the QueryDigest field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetQueryDigest() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryDigest
}

// GetQueryDigestOk returns a tuple with the QueryDigest field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionSample) GetQueryDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryDigest, true
}

// SetQueryDigest sets field value.
func (o *DiagnosticDBPerformanceLongTransactionSample) SetQueryDigest(v string) {
	o.QueryDigest = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceLongTransactionSample) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pid"] = o.Pid
	toSerialize["userName"] = o.UserName
	toSerialize["databaseName"] = o.DatabaseName
	if o.ApplicationName != nil {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if o.ClientAddr != nil {
		toSerialize["clientAddr"] = o.ClientAddr
	}
	if o.TransactionStart.Nanosecond() == 0 {
		toSerialize["transactionStart"] = o.TransactionStart.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["transactionStart"] = o.TransactionStart.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["transactionDurationSeconds"] = o.TransactionDurationSeconds
	toSerialize["state"] = o.State
	if o.WaitEventType != nil {
		toSerialize["waitEventType"] = o.WaitEventType
	}
	if o.WaitEvent != nil {
		toSerialize["waitEvent"] = o.WaitEvent
	}
	toSerialize["queryDigest"] = o.QueryDigest

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceLongTransactionSample) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pid                        *int64     `json:"pid"`
		UserName                   *string    `json:"userName"`
		DatabaseName               *string    `json:"databaseName"`
		ApplicationName            *string    `json:"applicationName,omitempty"`
		ClientAddr                 *string    `json:"clientAddr,omitempty"`
		TransactionStart           *time.Time `json:"transactionStart"`
		TransactionDurationSeconds *int64     `json:"transactionDurationSeconds"`
		State                      *string    `json:"state"`
		WaitEventType              *string    `json:"waitEventType,omitempty"`
		WaitEvent                  *string    `json:"waitEvent,omitempty"`
		QueryDigest                *string    `json:"queryDigest"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Pid == nil {
		return fmt.Errorf("required field pid missing")
	}
	if all.UserName == nil {
		return fmt.Errorf("required field userName missing")
	}
	if all.DatabaseName == nil {
		return fmt.Errorf("required field databaseName missing")
	}
	if all.TransactionStart == nil {
		return fmt.Errorf("required field transactionStart missing")
	}
	if all.TransactionDurationSeconds == nil {
		return fmt.Errorf("required field transactionDurationSeconds missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.QueryDigest == nil {
		return fmt.Errorf("required field queryDigest missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pid", "userName", "databaseName", "applicationName", "clientAddr", "transactionStart", "transactionDurationSeconds", "state", "waitEventType", "waitEvent", "queryDigest"})
	} else {
		return err
	}
	o.Pid = *all.Pid
	o.UserName = *all.UserName
	o.DatabaseName = *all.DatabaseName
	o.ApplicationName = all.ApplicationName
	o.ClientAddr = all.ClientAddr
	o.TransactionStart = *all.TransactionStart
	o.TransactionDurationSeconds = *all.TransactionDurationSeconds
	o.State = *all.State
	o.WaitEventType = all.WaitEventType
	o.WaitEvent = all.WaitEvent
	o.QueryDigest = *all.QueryDigest

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
