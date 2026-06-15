// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSession struct {
	Pid             int64  `json:"pid"`
	User            string `json:"user"`
	Database        string `json:"database"`
	ApplicationName string `json:"applicationName"`
	// Client address.
	ClientAddr string `json:"clientAddr"`
	// Client port.
	ClientPort    string `json:"clientPort"`
	State         string `json:"state"`
	WaitEventType string `json:"waitEventType"`
	WaitEvent     string `json:"waitEvent"`
	// Backend start timestamp if available.
	BackendStart *string `json:"backendStart,omitempty"`
	// Query start timestamp if available.
	QueryStart *string `json:"queryStart,omitempty"`
	// Transaction start timestamp if available.
	XactStart           *string `json:"xactStart,omitempty"`
	DurationSeconds     int64   `json:"durationSeconds"`
	XactDurationSeconds int64   `json:"xactDurationSeconds"`
	QueryDigest         string  `json:"queryDigest"`
	QuerySummary        string  `json:"querySummary"`
	BackendType         string  `json:"backendType"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSession instantiates a new PostgresqlSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSession(pid int64, user string, database string, applicationName string, clientAddr string, clientPort string, state string, waitEventType string, waitEvent string, durationSeconds int64, xactDurationSeconds int64, queryDigest string, querySummary string, backendType string) *PostgresqlSession {
	this := PostgresqlSession{}
	this.Pid = pid
	this.User = user
	this.Database = database
	this.ApplicationName = applicationName
	this.ClientAddr = clientAddr
	this.ClientPort = clientPort
	this.State = state
	this.WaitEventType = waitEventType
	this.WaitEvent = waitEvent
	this.DurationSeconds = durationSeconds
	this.XactDurationSeconds = xactDurationSeconds
	this.QueryDigest = queryDigest
	this.QuerySummary = querySummary
	this.BackendType = backendType
	return &this
}

// NewPostgresqlSessionWithDefaults instantiates a new PostgresqlSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSessionWithDefaults() *PostgresqlSession {
	this := PostgresqlSession{}
	return &this
}

// GetPid returns the Pid field value.
func (o *PostgresqlSession) GetPid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetPidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value.
func (o *PostgresqlSession) SetPid(v int64) {
	o.Pid = v
}

// GetUser returns the User field value.
func (o *PostgresqlSession) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *PostgresqlSession) SetUser(v string) {
	o.User = v
}

// GetDatabase returns the Database field value.
func (o *PostgresqlSession) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *PostgresqlSession) SetDatabase(v string) {
	o.Database = v
}

// GetApplicationName returns the ApplicationName field value.
func (o *PostgresqlSession) GetApplicationName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetApplicationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationName, true
}

// SetApplicationName sets field value.
func (o *PostgresqlSession) SetApplicationName(v string) {
	o.ApplicationName = v
}

// GetClientAddr returns the ClientAddr field value.
func (o *PostgresqlSession) GetClientAddr() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientAddr
}

// GetClientAddrOk returns a tuple with the ClientAddr field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetClientAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientAddr, true
}

// SetClientAddr sets field value.
func (o *PostgresqlSession) SetClientAddr(v string) {
	o.ClientAddr = v
}

// GetClientPort returns the ClientPort field value.
func (o *PostgresqlSession) GetClientPort() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientPort
}

// GetClientPortOk returns a tuple with the ClientPort field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetClientPortOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientPort, true
}

// SetClientPort sets field value.
func (o *PostgresqlSession) SetClientPort(v string) {
	o.ClientPort = v
}

// GetState returns the State field value.
func (o *PostgresqlSession) GetState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *PostgresqlSession) SetState(v string) {
	o.State = v
}

// GetWaitEventType returns the WaitEventType field value.
func (o *PostgresqlSession) GetWaitEventType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WaitEventType
}

// GetWaitEventTypeOk returns a tuple with the WaitEventType field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetWaitEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitEventType, true
}

// SetWaitEventType sets field value.
func (o *PostgresqlSession) SetWaitEventType(v string) {
	o.WaitEventType = v
}

// GetWaitEvent returns the WaitEvent field value.
func (o *PostgresqlSession) GetWaitEvent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WaitEvent
}

// GetWaitEventOk returns a tuple with the WaitEvent field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetWaitEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitEvent, true
}

// SetWaitEvent sets field value.
func (o *PostgresqlSession) SetWaitEvent(v string) {
	o.WaitEvent = v
}

// GetBackendStart returns the BackendStart field value if set, zero value otherwise.
func (o *PostgresqlSession) GetBackendStart() string {
	if o == nil || o.BackendStart == nil {
		var ret string
		return ret
	}
	return *o.BackendStart
}

// GetBackendStartOk returns a tuple with the BackendStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetBackendStartOk() (*string, bool) {
	if o == nil || o.BackendStart == nil {
		return nil, false
	}
	return o.BackendStart, true
}

// HasBackendStart returns a boolean if a field has been set.
func (o *PostgresqlSession) HasBackendStart() bool {
	return o != nil && o.BackendStart != nil
}

// SetBackendStart gets a reference to the given string and assigns it to the BackendStart field.
func (o *PostgresqlSession) SetBackendStart(v string) {
	o.BackendStart = &v
}

// GetQueryStart returns the QueryStart field value if set, zero value otherwise.
func (o *PostgresqlSession) GetQueryStart() string {
	if o == nil || o.QueryStart == nil {
		var ret string
		return ret
	}
	return *o.QueryStart
}

// GetQueryStartOk returns a tuple with the QueryStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetQueryStartOk() (*string, bool) {
	if o == nil || o.QueryStart == nil {
		return nil, false
	}
	return o.QueryStart, true
}

// HasQueryStart returns a boolean if a field has been set.
func (o *PostgresqlSession) HasQueryStart() bool {
	return o != nil && o.QueryStart != nil
}

// SetQueryStart gets a reference to the given string and assigns it to the QueryStart field.
func (o *PostgresqlSession) SetQueryStart(v string) {
	o.QueryStart = &v
}

// GetXactStart returns the XactStart field value if set, zero value otherwise.
func (o *PostgresqlSession) GetXactStart() string {
	if o == nil || o.XactStart == nil {
		var ret string
		return ret
	}
	return *o.XactStart
}

// GetXactStartOk returns a tuple with the XactStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetXactStartOk() (*string, bool) {
	if o == nil || o.XactStart == nil {
		return nil, false
	}
	return o.XactStart, true
}

// HasXactStart returns a boolean if a field has been set.
func (o *PostgresqlSession) HasXactStart() bool {
	return o != nil && o.XactStart != nil
}

// SetXactStart gets a reference to the given string and assigns it to the XactStart field.
func (o *PostgresqlSession) SetXactStart(v string) {
	o.XactStart = &v
}

// GetDurationSeconds returns the DurationSeconds field value.
func (o *PostgresqlSession) GetDurationSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationSeconds, true
}

// SetDurationSeconds sets field value.
func (o *PostgresqlSession) SetDurationSeconds(v int64) {
	o.DurationSeconds = v
}

// GetXactDurationSeconds returns the XactDurationSeconds field value.
func (o *PostgresqlSession) GetXactDurationSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.XactDurationSeconds
}

// GetXactDurationSecondsOk returns a tuple with the XactDurationSeconds field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetXactDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XactDurationSeconds, true
}

// SetXactDurationSeconds sets field value.
func (o *PostgresqlSession) SetXactDurationSeconds(v int64) {
	o.XactDurationSeconds = v
}

// GetQueryDigest returns the QueryDigest field value.
func (o *PostgresqlSession) GetQueryDigest() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryDigest
}

// GetQueryDigestOk returns a tuple with the QueryDigest field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetQueryDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryDigest, true
}

// SetQueryDigest sets field value.
func (o *PostgresqlSession) SetQueryDigest(v string) {
	o.QueryDigest = v
}

// GetQuerySummary returns the QuerySummary field value.
func (o *PostgresqlSession) GetQuerySummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QuerySummary
}

// GetQuerySummaryOk returns a tuple with the QuerySummary field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetQuerySummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuerySummary, true
}

// SetQuerySummary sets field value.
func (o *PostgresqlSession) SetQuerySummary(v string) {
	o.QuerySummary = v
}

// GetBackendType returns the BackendType field value.
func (o *PostgresqlSession) GetBackendType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BackendType
}

// GetBackendTypeOk returns a tuple with the BackendType field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSession) GetBackendTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackendType, true
}

// SetBackendType sets field value.
func (o *PostgresqlSession) SetBackendType(v string) {
	o.BackendType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pid"] = o.Pid
	toSerialize["user"] = o.User
	toSerialize["database"] = o.Database
	toSerialize["applicationName"] = o.ApplicationName
	toSerialize["clientAddr"] = o.ClientAddr
	toSerialize["clientPort"] = o.ClientPort
	toSerialize["state"] = o.State
	toSerialize["waitEventType"] = o.WaitEventType
	toSerialize["waitEvent"] = o.WaitEvent
	if o.BackendStart != nil {
		toSerialize["backendStart"] = o.BackendStart
	}
	if o.QueryStart != nil {
		toSerialize["queryStart"] = o.QueryStart
	}
	if o.XactStart != nil {
		toSerialize["xactStart"] = o.XactStart
	}
	toSerialize["durationSeconds"] = o.DurationSeconds
	toSerialize["xactDurationSeconds"] = o.XactDurationSeconds
	toSerialize["queryDigest"] = o.QueryDigest
	toSerialize["querySummary"] = o.QuerySummary
	toSerialize["backendType"] = o.BackendType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pid                 *int64  `json:"pid"`
		User                *string `json:"user"`
		Database            *string `json:"database"`
		ApplicationName     *string `json:"applicationName"`
		ClientAddr          *string `json:"clientAddr"`
		ClientPort          *string `json:"clientPort"`
		State               *string `json:"state"`
		WaitEventType       *string `json:"waitEventType"`
		WaitEvent           *string `json:"waitEvent"`
		BackendStart        *string `json:"backendStart,omitempty"`
		QueryStart          *string `json:"queryStart,omitempty"`
		XactStart           *string `json:"xactStart,omitempty"`
		DurationSeconds     *int64  `json:"durationSeconds"`
		XactDurationSeconds *int64  `json:"xactDurationSeconds"`
		QueryDigest         *string `json:"queryDigest"`
		QuerySummary        *string `json:"querySummary"`
		BackendType         *string `json:"backendType"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Pid == nil {
		return fmt.Errorf("required field pid missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.ApplicationName == nil {
		return fmt.Errorf("required field applicationName missing")
	}
	if all.ClientAddr == nil {
		return fmt.Errorf("required field clientAddr missing")
	}
	if all.ClientPort == nil {
		return fmt.Errorf("required field clientPort missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.WaitEventType == nil {
		return fmt.Errorf("required field waitEventType missing")
	}
	if all.WaitEvent == nil {
		return fmt.Errorf("required field waitEvent missing")
	}
	if all.DurationSeconds == nil {
		return fmt.Errorf("required field durationSeconds missing")
	}
	if all.XactDurationSeconds == nil {
		return fmt.Errorf("required field xactDurationSeconds missing")
	}
	if all.QueryDigest == nil {
		return fmt.Errorf("required field queryDigest missing")
	}
	if all.QuerySummary == nil {
		return fmt.Errorf("required field querySummary missing")
	}
	if all.BackendType == nil {
		return fmt.Errorf("required field backendType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pid", "user", "database", "applicationName", "clientAddr", "clientPort", "state", "waitEventType", "waitEvent", "backendStart", "queryStart", "xactStart", "durationSeconds", "xactDurationSeconds", "queryDigest", "querySummary", "backendType"})
	} else {
		return err
	}
	o.Pid = *all.Pid
	o.User = *all.User
	o.Database = *all.Database
	o.ApplicationName = *all.ApplicationName
	o.ClientAddr = *all.ClientAddr
	o.ClientPort = *all.ClientPort
	o.State = *all.State
	o.WaitEventType = *all.WaitEventType
	o.WaitEvent = *all.WaitEvent
	o.BackendStart = all.BackendStart
	o.QueryStart = all.QueryStart
	o.XactStart = all.XactStart
	o.DurationSeconds = *all.DurationSeconds
	o.XactDurationSeconds = *all.XactDurationSeconds
	o.QueryDigest = *all.QueryDigest
	o.QuerySummary = *all.QuerySummary
	o.BackendType = *all.BackendType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
