// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlRealtimeSession struct {
	Pid      int64  `json:"pid"`
	User     string `json:"user"`
	Database string `json:"database"`
	// Masked client address.
	Client        string `json:"client"`
	Application   string `json:"application"`
	State         string `json:"state"`
	WaitEventType string `json:"waitEventType"`
	WaitEvent     string `json:"waitEvent"`
	// Backend start timestamp if available.
	BackendStart *string `json:"backendStart,omitempty"`
	// Query start timestamp if available.
	QueryStart *string `json:"queryStart,omitempty"`
	// Transaction start timestamp if available.
	XactStart           *string  `json:"xactStart,omitempty"`
	DurationSeconds     int64    `json:"durationSeconds"`
	XactDurationSeconds int64    `json:"xactDurationSeconds"`
	QueryDigest         string   `json:"queryDigest"`
	RedactedQuery       string   `json:"redactedQuery"`
	BackendType         string   `json:"backendType"`
	RiskTags            []string `json:"riskTags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlRealtimeSession instantiates a new PostgresqlRealtimeSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlRealtimeSession(pid int64, user string, database string, client string, application string, state string, waitEventType string, waitEvent string, durationSeconds int64, xactDurationSeconds int64, queryDigest string, redactedQuery string, backendType string, riskTags []string) *PostgresqlRealtimeSession {
	this := PostgresqlRealtimeSession{}
	this.Pid = pid
	this.User = user
	this.Database = database
	this.Client = client
	this.Application = application
	this.State = state
	this.WaitEventType = waitEventType
	this.WaitEvent = waitEvent
	this.DurationSeconds = durationSeconds
	this.XactDurationSeconds = xactDurationSeconds
	this.QueryDigest = queryDigest
	this.RedactedQuery = redactedQuery
	this.BackendType = backendType
	this.RiskTags = riskTags
	return &this
}

// NewPostgresqlRealtimeSessionWithDefaults instantiates a new PostgresqlRealtimeSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlRealtimeSessionWithDefaults() *PostgresqlRealtimeSession {
	this := PostgresqlRealtimeSession{}
	return &this
}

// GetPid returns the Pid field value.
func (o *PostgresqlRealtimeSession) GetPid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetPidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value.
func (o *PostgresqlRealtimeSession) SetPid(v int64) {
	o.Pid = v
}

// GetUser returns the User field value.
func (o *PostgresqlRealtimeSession) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *PostgresqlRealtimeSession) SetUser(v string) {
	o.User = v
}

// GetDatabase returns the Database field value.
func (o *PostgresqlRealtimeSession) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *PostgresqlRealtimeSession) SetDatabase(v string) {
	o.Database = v
}

// GetClient returns the Client field value.
func (o *PostgresqlRealtimeSession) GetClient() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Client
}

// GetClientOk returns a tuple with the Client field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetClientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Client, true
}

// SetClient sets field value.
func (o *PostgresqlRealtimeSession) SetClient(v string) {
	o.Client = v
}

// GetApplication returns the Application field value.
func (o *PostgresqlRealtimeSession) GetApplication() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetApplicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value.
func (o *PostgresqlRealtimeSession) SetApplication(v string) {
	o.Application = v
}

// GetState returns the State field value.
func (o *PostgresqlRealtimeSession) GetState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *PostgresqlRealtimeSession) SetState(v string) {
	o.State = v
}

// GetWaitEventType returns the WaitEventType field value.
func (o *PostgresqlRealtimeSession) GetWaitEventType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WaitEventType
}

// GetWaitEventTypeOk returns a tuple with the WaitEventType field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetWaitEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitEventType, true
}

// SetWaitEventType sets field value.
func (o *PostgresqlRealtimeSession) SetWaitEventType(v string) {
	o.WaitEventType = v
}

// GetWaitEvent returns the WaitEvent field value.
func (o *PostgresqlRealtimeSession) GetWaitEvent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WaitEvent
}

// GetWaitEventOk returns a tuple with the WaitEvent field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetWaitEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaitEvent, true
}

// SetWaitEvent sets field value.
func (o *PostgresqlRealtimeSession) SetWaitEvent(v string) {
	o.WaitEvent = v
}

// GetBackendStart returns the BackendStart field value if set, zero value otherwise.
func (o *PostgresqlRealtimeSession) GetBackendStart() string {
	if o == nil || o.BackendStart == nil {
		var ret string
		return ret
	}
	return *o.BackendStart
}

// GetBackendStartOk returns a tuple with the BackendStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetBackendStartOk() (*string, bool) {
	if o == nil || o.BackendStart == nil {
		return nil, false
	}
	return o.BackendStart, true
}

// HasBackendStart returns a boolean if a field has been set.
func (o *PostgresqlRealtimeSession) HasBackendStart() bool {
	return o != nil && o.BackendStart != nil
}

// SetBackendStart gets a reference to the given string and assigns it to the BackendStart field.
func (o *PostgresqlRealtimeSession) SetBackendStart(v string) {
	o.BackendStart = &v
}

// GetQueryStart returns the QueryStart field value if set, zero value otherwise.
func (o *PostgresqlRealtimeSession) GetQueryStart() string {
	if o == nil || o.QueryStart == nil {
		var ret string
		return ret
	}
	return *o.QueryStart
}

// GetQueryStartOk returns a tuple with the QueryStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetQueryStartOk() (*string, bool) {
	if o == nil || o.QueryStart == nil {
		return nil, false
	}
	return o.QueryStart, true
}

// HasQueryStart returns a boolean if a field has been set.
func (o *PostgresqlRealtimeSession) HasQueryStart() bool {
	return o != nil && o.QueryStart != nil
}

// SetQueryStart gets a reference to the given string and assigns it to the QueryStart field.
func (o *PostgresqlRealtimeSession) SetQueryStart(v string) {
	o.QueryStart = &v
}

// GetXactStart returns the XactStart field value if set, zero value otherwise.
func (o *PostgresqlRealtimeSession) GetXactStart() string {
	if o == nil || o.XactStart == nil {
		var ret string
		return ret
	}
	return *o.XactStart
}

// GetXactStartOk returns a tuple with the XactStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetXactStartOk() (*string, bool) {
	if o == nil || o.XactStart == nil {
		return nil, false
	}
	return o.XactStart, true
}

// HasXactStart returns a boolean if a field has been set.
func (o *PostgresqlRealtimeSession) HasXactStart() bool {
	return o != nil && o.XactStart != nil
}

// SetXactStart gets a reference to the given string and assigns it to the XactStart field.
func (o *PostgresqlRealtimeSession) SetXactStart(v string) {
	o.XactStart = &v
}

// GetDurationSeconds returns the DurationSeconds field value.
func (o *PostgresqlRealtimeSession) GetDurationSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationSeconds, true
}

// SetDurationSeconds sets field value.
func (o *PostgresqlRealtimeSession) SetDurationSeconds(v int64) {
	o.DurationSeconds = v
}

// GetXactDurationSeconds returns the XactDurationSeconds field value.
func (o *PostgresqlRealtimeSession) GetXactDurationSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.XactDurationSeconds
}

// GetXactDurationSecondsOk returns a tuple with the XactDurationSeconds field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetXactDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XactDurationSeconds, true
}

// SetXactDurationSeconds sets field value.
func (o *PostgresqlRealtimeSession) SetXactDurationSeconds(v int64) {
	o.XactDurationSeconds = v
}

// GetQueryDigest returns the QueryDigest field value.
func (o *PostgresqlRealtimeSession) GetQueryDigest() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryDigest
}

// GetQueryDigestOk returns a tuple with the QueryDigest field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetQueryDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryDigest, true
}

// SetQueryDigest sets field value.
func (o *PostgresqlRealtimeSession) SetQueryDigest(v string) {
	o.QueryDigest = v
}

// GetRedactedQuery returns the RedactedQuery field value.
func (o *PostgresqlRealtimeSession) GetRedactedQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RedactedQuery
}

// GetRedactedQueryOk returns a tuple with the RedactedQuery field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetRedactedQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedactedQuery, true
}

// SetRedactedQuery sets field value.
func (o *PostgresqlRealtimeSession) SetRedactedQuery(v string) {
	o.RedactedQuery = v
}

// GetBackendType returns the BackendType field value.
func (o *PostgresqlRealtimeSession) GetBackendType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BackendType
}

// GetBackendTypeOk returns a tuple with the BackendType field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetBackendTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackendType, true
}

// SetBackendType sets field value.
func (o *PostgresqlRealtimeSession) SetBackendType(v string) {
	o.BackendType = v
}

// GetRiskTags returns the RiskTags field value.
func (o *PostgresqlRealtimeSession) GetRiskTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RiskTags
}

// GetRiskTagsOk returns a tuple with the RiskTags field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSession) GetRiskTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskTags, true
}

// SetRiskTags sets field value.
func (o *PostgresqlRealtimeSession) SetRiskTags(v []string) {
	o.RiskTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlRealtimeSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pid"] = o.Pid
	toSerialize["user"] = o.User
	toSerialize["database"] = o.Database
	toSerialize["client"] = o.Client
	toSerialize["application"] = o.Application
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
	toSerialize["redactedQuery"] = o.RedactedQuery
	toSerialize["backendType"] = o.BackendType
	toSerialize["riskTags"] = o.RiskTags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlRealtimeSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pid                 *int64    `json:"pid"`
		User                *string   `json:"user"`
		Database            *string   `json:"database"`
		Client              *string   `json:"client"`
		Application         *string   `json:"application"`
		State               *string   `json:"state"`
		WaitEventType       *string   `json:"waitEventType"`
		WaitEvent           *string   `json:"waitEvent"`
		BackendStart        *string   `json:"backendStart,omitempty"`
		QueryStart          *string   `json:"queryStart,omitempty"`
		XactStart           *string   `json:"xactStart,omitempty"`
		DurationSeconds     *int64    `json:"durationSeconds"`
		XactDurationSeconds *int64    `json:"xactDurationSeconds"`
		QueryDigest         *string   `json:"queryDigest"`
		RedactedQuery       *string   `json:"redactedQuery"`
		BackendType         *string   `json:"backendType"`
		RiskTags            *[]string `json:"riskTags"`
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
	if all.Client == nil {
		return fmt.Errorf("required field client missing")
	}
	if all.Application == nil {
		return fmt.Errorf("required field application missing")
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
	if all.RedactedQuery == nil {
		return fmt.Errorf("required field redactedQuery missing")
	}
	if all.BackendType == nil {
		return fmt.Errorf("required field backendType missing")
	}
	if all.RiskTags == nil {
		return fmt.Errorf("required field riskTags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pid", "user", "database", "client", "application", "state", "waitEventType", "waitEvent", "backendStart", "queryStart", "xactStart", "durationSeconds", "xactDurationSeconds", "queryDigest", "redactedQuery", "backendType", "riskTags"})
	} else {
		return err
	}
	o.Pid = *all.Pid
	o.User = *all.User
	o.Database = *all.Database
	o.Client = *all.Client
	o.Application = *all.Application
	o.State = *all.State
	o.WaitEventType = *all.WaitEventType
	o.WaitEvent = *all.WaitEvent
	o.BackendStart = all.BackendStart
	o.QueryStart = all.QueryStart
	o.XactStart = all.XactStart
	o.DurationSeconds = *all.DurationSeconds
	o.XactDurationSeconds = *all.XactDurationSeconds
	o.QueryDigest = *all.QueryDigest
	o.RedactedQuery = *all.RedactedQuery
	o.BackendType = *all.BackendType
	o.RiskTags = *all.RiskTags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
