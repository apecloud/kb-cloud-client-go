// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MssqlSession struct {
	SessionId int64 `json:"sessionId"`
	// Opaque SQL Server local start time, used to detect session or request reuse. Pass back unchanged; it is not a UTC timestamp.
	StartedAt          string `json:"startedAt"`
	Status             string `json:"status"`
	User               string `json:"user"`
	Host               string `json:"host"`
	Application        string `json:"application"`
	Client             string `json:"client"`
	Database           string `json:"database"`
	OpenTransactions   int64  `json:"openTransactions"`
	TransactionSeconds int64  `json:"transactionSeconds"`
	// Most recent batch on the connection, not necessarily the SQL that acquired its locks. Truncated to 8192 characters.
	LastSql           string         `json:"lastSql"`
	LastSqlTruncated  bool           `json:"lastSqlTruncated"`
	Requests          []MssqlRequest `json:"requests"`
	RequestsTruncated bool           `json:"requestsTruncated"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlSession instantiates a new MssqlSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlSession(sessionId int64, startedAt string, status string, user string, host string, application string, client string, database string, openTransactions int64, transactionSeconds int64, lastSql string, lastSqlTruncated bool, requests []MssqlRequest, requestsTruncated bool) *MssqlSession {
	this := MssqlSession{}
	this.SessionId = sessionId
	this.StartedAt = startedAt
	this.Status = status
	this.User = user
	this.Host = host
	this.Application = application
	this.Client = client
	this.Database = database
	this.OpenTransactions = openTransactions
	this.TransactionSeconds = transactionSeconds
	this.LastSql = lastSql
	this.LastSqlTruncated = lastSqlTruncated
	this.Requests = requests
	this.RequestsTruncated = requestsTruncated
	return &this
}

// NewMssqlSessionWithDefaults instantiates a new MssqlSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlSessionWithDefaults() *MssqlSession {
	this := MssqlSession{}
	return &this
}

// GetSessionId returns the SessionId field value.
func (o *MssqlSession) GetSessionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetSessionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *MssqlSession) SetSessionId(v int64) {
	o.SessionId = v
}

// GetStartedAt returns the StartedAt field value.
func (o *MssqlSession) GetStartedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetStartedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *MssqlSession) SetStartedAt(v string) {
	o.StartedAt = v
}

// GetStatus returns the Status field value.
func (o *MssqlSession) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *MssqlSession) SetStatus(v string) {
	o.Status = v
}

// GetUser returns the User field value.
func (o *MssqlSession) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *MssqlSession) SetUser(v string) {
	o.User = v
}

// GetHost returns the Host field value.
func (o *MssqlSession) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value.
func (o *MssqlSession) SetHost(v string) {
	o.Host = v
}

// GetApplication returns the Application field value.
func (o *MssqlSession) GetApplication() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetApplicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value.
func (o *MssqlSession) SetApplication(v string) {
	o.Application = v
}

// GetClient returns the Client field value.
func (o *MssqlSession) GetClient() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Client
}

// GetClientOk returns a tuple with the Client field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetClientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Client, true
}

// SetClient sets field value.
func (o *MssqlSession) SetClient(v string) {
	o.Client = v
}

// GetDatabase returns the Database field value.
func (o *MssqlSession) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *MssqlSession) SetDatabase(v string) {
	o.Database = v
}

// GetOpenTransactions returns the OpenTransactions field value.
func (o *MssqlSession) GetOpenTransactions() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OpenTransactions
}

// GetOpenTransactionsOk returns a tuple with the OpenTransactions field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetOpenTransactionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenTransactions, true
}

// SetOpenTransactions sets field value.
func (o *MssqlSession) SetOpenTransactions(v int64) {
	o.OpenTransactions = v
}

// GetTransactionSeconds returns the TransactionSeconds field value.
func (o *MssqlSession) GetTransactionSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TransactionSeconds
}

// GetTransactionSecondsOk returns a tuple with the TransactionSeconds field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetTransactionSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionSeconds, true
}

// SetTransactionSeconds sets field value.
func (o *MssqlSession) SetTransactionSeconds(v int64) {
	o.TransactionSeconds = v
}

// GetLastSql returns the LastSql field value.
func (o *MssqlSession) GetLastSql() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LastSql
}

// GetLastSqlOk returns a tuple with the LastSql field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetLastSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSql, true
}

// SetLastSql sets field value.
func (o *MssqlSession) SetLastSql(v string) {
	o.LastSql = v
}

// GetLastSqlTruncated returns the LastSqlTruncated field value.
func (o *MssqlSession) GetLastSqlTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.LastSqlTruncated
}

// GetLastSqlTruncatedOk returns a tuple with the LastSqlTruncated field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetLastSqlTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSqlTruncated, true
}

// SetLastSqlTruncated sets field value.
func (o *MssqlSession) SetLastSqlTruncated(v bool) {
	o.LastSqlTruncated = v
}

// GetRequests returns the Requests field value.
func (o *MssqlSession) GetRequests() []MssqlRequest {
	if o == nil {
		var ret []MssqlRequest
		return ret
	}
	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetRequestsOk() (*[]MssqlRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requests, true
}

// SetRequests sets field value.
func (o *MssqlSession) SetRequests(v []MssqlRequest) {
	o.Requests = v
}

// GetRequestsTruncated returns the RequestsTruncated field value.
func (o *MssqlSession) GetRequestsTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.RequestsTruncated
}

// GetRequestsTruncatedOk returns a tuple with the RequestsTruncated field value
// and a boolean to check if the value has been set.
func (o *MssqlSession) GetRequestsTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestsTruncated, true
}

// SetRequestsTruncated sets field value.
func (o *MssqlSession) SetRequestsTruncated(v bool) {
	o.RequestsTruncated = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["sessionId"] = o.SessionId
	toSerialize["startedAt"] = o.StartedAt
	toSerialize["status"] = o.Status
	toSerialize["user"] = o.User
	toSerialize["host"] = o.Host
	toSerialize["application"] = o.Application
	toSerialize["client"] = o.Client
	toSerialize["database"] = o.Database
	toSerialize["openTransactions"] = o.OpenTransactions
	toSerialize["transactionSeconds"] = o.TransactionSeconds
	toSerialize["lastSql"] = o.LastSql
	toSerialize["lastSqlTruncated"] = o.LastSqlTruncated
	toSerialize["requests"] = o.Requests
	toSerialize["requestsTruncated"] = o.RequestsTruncated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessionId          *int64          `json:"sessionId"`
		StartedAt          *string         `json:"startedAt"`
		Status             *string         `json:"status"`
		User               *string         `json:"user"`
		Host               *string         `json:"host"`
		Application        *string         `json:"application"`
		Client             *string         `json:"client"`
		Database           *string         `json:"database"`
		OpenTransactions   *int64          `json:"openTransactions"`
		TransactionSeconds *int64          `json:"transactionSeconds"`
		LastSql            *string         `json:"lastSql"`
		LastSqlTruncated   *bool           `json:"lastSqlTruncated"`
		Requests           *[]MssqlRequest `json:"requests"`
		RequestsTruncated  *bool           `json:"requestsTruncated"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field sessionId missing")
	}
	if all.StartedAt == nil {
		return fmt.Errorf("required field startedAt missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	if all.Host == nil {
		return fmt.Errorf("required field host missing")
	}
	if all.Application == nil {
		return fmt.Errorf("required field application missing")
	}
	if all.Client == nil {
		return fmt.Errorf("required field client missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.OpenTransactions == nil {
		return fmt.Errorf("required field openTransactions missing")
	}
	if all.TransactionSeconds == nil {
		return fmt.Errorf("required field transactionSeconds missing")
	}
	if all.LastSql == nil {
		return fmt.Errorf("required field lastSql missing")
	}
	if all.LastSqlTruncated == nil {
		return fmt.Errorf("required field lastSqlTruncated missing")
	}
	if all.Requests == nil {
		return fmt.Errorf("required field requests missing")
	}
	if all.RequestsTruncated == nil {
		return fmt.Errorf("required field requestsTruncated missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sessionId", "startedAt", "status", "user", "host", "application", "client", "database", "openTransactions", "transactionSeconds", "lastSql", "lastSqlTruncated", "requests", "requestsTruncated"})
	} else {
		return err
	}
	o.SessionId = *all.SessionId
	o.StartedAt = *all.StartedAt
	o.Status = *all.Status
	o.User = *all.User
	o.Host = *all.Host
	o.Application = *all.Application
	o.Client = *all.Client
	o.Database = *all.Database
	o.OpenTransactions = *all.OpenTransactions
	o.TransactionSeconds = *all.TransactionSeconds
	o.LastSql = *all.LastSql
	o.LastSqlTruncated = *all.LastSqlTruncated
	o.Requests = *all.Requests
	o.RequestsTruncated = *all.RequestsTruncated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
