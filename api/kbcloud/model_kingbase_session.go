// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KingbaseSession struct {
	Pid                 int64   `json:"pid"`
	User                *string `json:"user,omitempty"`
	Database            *string `json:"database,omitempty"`
	ApplicationName     *string `json:"applicationName,omitempty"`
	ClientAddr          *string `json:"clientAddr,omitempty"`
	ClientPort          *int32  `json:"clientPort,omitempty"`
	State               string  `json:"state"`
	WaitEventType       *string `json:"waitEventType,omitempty"`
	WaitEvent           *string `json:"waitEvent,omitempty"`
	BackendStart        *string `json:"backendStart,omitempty"`
	QueryStart          *string `json:"queryStart,omitempty"`
	XactStart           *string `json:"xactStart,omitempty"`
	DurationSeconds     *int64  `json:"durationSeconds,omitempty"`
	XactDurationSeconds *int64  `json:"xactDurationSeconds,omitempty"`
	QueryDigest         string  `json:"queryDigest"`
	QuerySummary        string  `json:"querySummary"`
	BackendType         *string `json:"backendType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKingbaseSession instantiates a new KingbaseSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKingbaseSession(pid int64, state string, queryDigest string, querySummary string) *KingbaseSession {
	this := KingbaseSession{}
	this.Pid = pid
	this.State = state
	this.QueryDigest = queryDigest
	this.QuerySummary = querySummary
	return &this
}

// NewKingbaseSessionWithDefaults instantiates a new KingbaseSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKingbaseSessionWithDefaults() *KingbaseSession {
	this := KingbaseSession{}
	return &this
}

// GetPid returns the Pid field value.
func (o *KingbaseSession) GetPid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetPidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value.
func (o *KingbaseSession) SetPid(v int64) {
	o.Pid = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *KingbaseSession) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *KingbaseSession) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *KingbaseSession) SetUser(v string) {
	o.User = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *KingbaseSession) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *KingbaseSession) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *KingbaseSession) SetDatabase(v string) {
	o.Database = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *KingbaseSession) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *KingbaseSession) HasApplicationName() bool {
	return o != nil && o.ApplicationName != nil
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *KingbaseSession) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetClientAddr returns the ClientAddr field value if set, zero value otherwise.
func (o *KingbaseSession) GetClientAddr() string {
	if o == nil || o.ClientAddr == nil {
		var ret string
		return ret
	}
	return *o.ClientAddr
}

// GetClientAddrOk returns a tuple with the ClientAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetClientAddrOk() (*string, bool) {
	if o == nil || o.ClientAddr == nil {
		return nil, false
	}
	return o.ClientAddr, true
}

// HasClientAddr returns a boolean if a field has been set.
func (o *KingbaseSession) HasClientAddr() bool {
	return o != nil && o.ClientAddr != nil
}

// SetClientAddr gets a reference to the given string and assigns it to the ClientAddr field.
func (o *KingbaseSession) SetClientAddr(v string) {
	o.ClientAddr = &v
}

// GetClientPort returns the ClientPort field value if set, zero value otherwise.
func (o *KingbaseSession) GetClientPort() int32 {
	if o == nil || o.ClientPort == nil {
		var ret int32
		return ret
	}
	return *o.ClientPort
}

// GetClientPortOk returns a tuple with the ClientPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetClientPortOk() (*int32, bool) {
	if o == nil || o.ClientPort == nil {
		return nil, false
	}
	return o.ClientPort, true
}

// HasClientPort returns a boolean if a field has been set.
func (o *KingbaseSession) HasClientPort() bool {
	return o != nil && o.ClientPort != nil
}

// SetClientPort gets a reference to the given int32 and assigns it to the ClientPort field.
func (o *KingbaseSession) SetClientPort(v int32) {
	o.ClientPort = &v
}

// GetState returns the State field value.
func (o *KingbaseSession) GetState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *KingbaseSession) SetState(v string) {
	o.State = v
}

// GetWaitEventType returns the WaitEventType field value if set, zero value otherwise.
func (o *KingbaseSession) GetWaitEventType() string {
	if o == nil || o.WaitEventType == nil {
		var ret string
		return ret
	}
	return *o.WaitEventType
}

// GetWaitEventTypeOk returns a tuple with the WaitEventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetWaitEventTypeOk() (*string, bool) {
	if o == nil || o.WaitEventType == nil {
		return nil, false
	}
	return o.WaitEventType, true
}

// HasWaitEventType returns a boolean if a field has been set.
func (o *KingbaseSession) HasWaitEventType() bool {
	return o != nil && o.WaitEventType != nil
}

// SetWaitEventType gets a reference to the given string and assigns it to the WaitEventType field.
func (o *KingbaseSession) SetWaitEventType(v string) {
	o.WaitEventType = &v
}

// GetWaitEvent returns the WaitEvent field value if set, zero value otherwise.
func (o *KingbaseSession) GetWaitEvent() string {
	if o == nil || o.WaitEvent == nil {
		var ret string
		return ret
	}
	return *o.WaitEvent
}

// GetWaitEventOk returns a tuple with the WaitEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetWaitEventOk() (*string, bool) {
	if o == nil || o.WaitEvent == nil {
		return nil, false
	}
	return o.WaitEvent, true
}

// HasWaitEvent returns a boolean if a field has been set.
func (o *KingbaseSession) HasWaitEvent() bool {
	return o != nil && o.WaitEvent != nil
}

// SetWaitEvent gets a reference to the given string and assigns it to the WaitEvent field.
func (o *KingbaseSession) SetWaitEvent(v string) {
	o.WaitEvent = &v
}

// GetBackendStart returns the BackendStart field value if set, zero value otherwise.
func (o *KingbaseSession) GetBackendStart() string {
	if o == nil || o.BackendStart == nil {
		var ret string
		return ret
	}
	return *o.BackendStart
}

// GetBackendStartOk returns a tuple with the BackendStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetBackendStartOk() (*string, bool) {
	if o == nil || o.BackendStart == nil {
		return nil, false
	}
	return o.BackendStart, true
}

// HasBackendStart returns a boolean if a field has been set.
func (o *KingbaseSession) HasBackendStart() bool {
	return o != nil && o.BackendStart != nil
}

// SetBackendStart gets a reference to the given string and assigns it to the BackendStart field.
func (o *KingbaseSession) SetBackendStart(v string) {
	o.BackendStart = &v
}

// GetQueryStart returns the QueryStart field value if set, zero value otherwise.
func (o *KingbaseSession) GetQueryStart() string {
	if o == nil || o.QueryStart == nil {
		var ret string
		return ret
	}
	return *o.QueryStart
}

// GetQueryStartOk returns a tuple with the QueryStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetQueryStartOk() (*string, bool) {
	if o == nil || o.QueryStart == nil {
		return nil, false
	}
	return o.QueryStart, true
}

// HasQueryStart returns a boolean if a field has been set.
func (o *KingbaseSession) HasQueryStart() bool {
	return o != nil && o.QueryStart != nil
}

// SetQueryStart gets a reference to the given string and assigns it to the QueryStart field.
func (o *KingbaseSession) SetQueryStart(v string) {
	o.QueryStart = &v
}

// GetXactStart returns the XactStart field value if set, zero value otherwise.
func (o *KingbaseSession) GetXactStart() string {
	if o == nil || o.XactStart == nil {
		var ret string
		return ret
	}
	return *o.XactStart
}

// GetXactStartOk returns a tuple with the XactStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetXactStartOk() (*string, bool) {
	if o == nil || o.XactStart == nil {
		return nil, false
	}
	return o.XactStart, true
}

// HasXactStart returns a boolean if a field has been set.
func (o *KingbaseSession) HasXactStart() bool {
	return o != nil && o.XactStart != nil
}

// SetXactStart gets a reference to the given string and assigns it to the XactStart field.
func (o *KingbaseSession) SetXactStart(v string) {
	o.XactStart = &v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *KingbaseSession) GetDurationSeconds() int64 {
	if o == nil || o.DurationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetDurationSecondsOk() (*int64, bool) {
	if o == nil || o.DurationSeconds == nil {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *KingbaseSession) HasDurationSeconds() bool {
	return o != nil && o.DurationSeconds != nil
}

// SetDurationSeconds gets a reference to the given int64 and assigns it to the DurationSeconds field.
func (o *KingbaseSession) SetDurationSeconds(v int64) {
	o.DurationSeconds = &v
}

// GetXactDurationSeconds returns the XactDurationSeconds field value if set, zero value otherwise.
func (o *KingbaseSession) GetXactDurationSeconds() int64 {
	if o == nil || o.XactDurationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.XactDurationSeconds
}

// GetXactDurationSecondsOk returns a tuple with the XactDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetXactDurationSecondsOk() (*int64, bool) {
	if o == nil || o.XactDurationSeconds == nil {
		return nil, false
	}
	return o.XactDurationSeconds, true
}

// HasXactDurationSeconds returns a boolean if a field has been set.
func (o *KingbaseSession) HasXactDurationSeconds() bool {
	return o != nil && o.XactDurationSeconds != nil
}

// SetXactDurationSeconds gets a reference to the given int64 and assigns it to the XactDurationSeconds field.
func (o *KingbaseSession) SetXactDurationSeconds(v int64) {
	o.XactDurationSeconds = &v
}

// GetQueryDigest returns the QueryDigest field value.
func (o *KingbaseSession) GetQueryDigest() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueryDigest
}

// GetQueryDigestOk returns a tuple with the QueryDigest field value
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetQueryDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryDigest, true
}

// SetQueryDigest sets field value.
func (o *KingbaseSession) SetQueryDigest(v string) {
	o.QueryDigest = v
}

// GetQuerySummary returns the QuerySummary field value.
func (o *KingbaseSession) GetQuerySummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QuerySummary
}

// GetQuerySummaryOk returns a tuple with the QuerySummary field value
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetQuerySummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuerySummary, true
}

// SetQuerySummary sets field value.
func (o *KingbaseSession) SetQuerySummary(v string) {
	o.QuerySummary = v
}

// GetBackendType returns the BackendType field value if set, zero value otherwise.
func (o *KingbaseSession) GetBackendType() string {
	if o == nil || o.BackendType == nil {
		var ret string
		return ret
	}
	return *o.BackendType
}

// GetBackendTypeOk returns a tuple with the BackendType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseSession) GetBackendTypeOk() (*string, bool) {
	if o == nil || o.BackendType == nil {
		return nil, false
	}
	return o.BackendType, true
}

// HasBackendType returns a boolean if a field has been set.
func (o *KingbaseSession) HasBackendType() bool {
	return o != nil && o.BackendType != nil
}

// SetBackendType gets a reference to the given string and assigns it to the BackendType field.
func (o *KingbaseSession) SetBackendType(v string) {
	o.BackendType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o KingbaseSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pid"] = o.Pid
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.ApplicationName != nil {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if o.ClientAddr != nil {
		toSerialize["clientAddr"] = o.ClientAddr
	}
	if o.ClientPort != nil {
		toSerialize["clientPort"] = o.ClientPort
	}
	toSerialize["state"] = o.State
	if o.WaitEventType != nil {
		toSerialize["waitEventType"] = o.WaitEventType
	}
	if o.WaitEvent != nil {
		toSerialize["waitEvent"] = o.WaitEvent
	}
	if o.BackendStart != nil {
		toSerialize["backendStart"] = o.BackendStart
	}
	if o.QueryStart != nil {
		toSerialize["queryStart"] = o.QueryStart
	}
	if o.XactStart != nil {
		toSerialize["xactStart"] = o.XactStart
	}
	if o.DurationSeconds != nil {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}
	if o.XactDurationSeconds != nil {
		toSerialize["xactDurationSeconds"] = o.XactDurationSeconds
	}
	toSerialize["queryDigest"] = o.QueryDigest
	toSerialize["querySummary"] = o.QuerySummary
	if o.BackendType != nil {
		toSerialize["backendType"] = o.BackendType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KingbaseSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pid                 *int64  `json:"pid"`
		User                *string `json:"user,omitempty"`
		Database            *string `json:"database,omitempty"`
		ApplicationName     *string `json:"applicationName,omitempty"`
		ClientAddr          *string `json:"clientAddr,omitempty"`
		ClientPort          *int32  `json:"clientPort,omitempty"`
		State               *string `json:"state"`
		WaitEventType       *string `json:"waitEventType,omitempty"`
		WaitEvent           *string `json:"waitEvent,omitempty"`
		BackendStart        *string `json:"backendStart,omitempty"`
		QueryStart          *string `json:"queryStart,omitempty"`
		XactStart           *string `json:"xactStart,omitempty"`
		DurationSeconds     *int64  `json:"durationSeconds,omitempty"`
		XactDurationSeconds *int64  `json:"xactDurationSeconds,omitempty"`
		QueryDigest         *string `json:"queryDigest"`
		QuerySummary        *string `json:"querySummary"`
		BackendType         *string `json:"backendType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Pid == nil {
		return fmt.Errorf("required field pid missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.QueryDigest == nil {
		return fmt.Errorf("required field queryDigest missing")
	}
	if all.QuerySummary == nil {
		return fmt.Errorf("required field querySummary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pid", "user", "database", "applicationName", "clientAddr", "clientPort", "state", "waitEventType", "waitEvent", "backendStart", "queryStart", "xactStart", "durationSeconds", "xactDurationSeconds", "queryDigest", "querySummary", "backendType"})
	} else {
		return err
	}
	o.Pid = *all.Pid
	o.User = all.User
	o.Database = all.Database
	o.ApplicationName = all.ApplicationName
	o.ClientAddr = all.ClientAddr
	o.ClientPort = all.ClientPort
	o.State = *all.State
	o.WaitEventType = all.WaitEventType
	o.WaitEvent = all.WaitEvent
	o.BackendStart = all.BackendStart
	o.QueryStart = all.QueryStart
	o.XactStart = all.XactStart
	o.DurationSeconds = all.DurationSeconds
	o.XactDurationSeconds = all.XactDurationSeconds
	o.QueryDigest = *all.QueryDigest
	o.QuerySummary = *all.QuerySummary
	o.BackendType = all.BackendType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
