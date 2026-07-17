// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengSessionDetail struct {
	// Dameng session ID from V$SESSIONS.
	SessId int64 `json:"sessId"`
	// Database user name.
	UserName *string `json:"userName,omitempty"`
	// Session state (ACTIVE, IDLE, etc.).
	State *string `json:"state,omitempty"`
	// Client IP address.
	ClientIp *string `json:"clientIp,omitempty"`
	// Client type.
	ClientType *string `json:"clientType,omitempty"`
	// Session creation time.
	CreateTime *string `json:"createTime,omitempty"`
	// Last SQL send time.
	LastSendTime *string `json:"lastSendTime,omitempty"`
	// Current or last SQL text from V$SESSIONS.
	SqlText *string `json:"sqlText,omitempty"`
	// Full SQL text via SF_GET_SESSION_SQL.
	FullSql *string `json:"fullSql,omitempty"`
	// Transaction ID.
	TrxId *string `json:"trxId,omitempty"`
	// Seconds since last SQL send time.
	TrxDurationSec *int64 `json:"trxDurationSec,omitempty"`
	// Application name.
	AppName *string `json:"appName,omitempty"`
	// Current schema.
	CurrSchema *string `json:"currSchema,omitempty"`
	// Auto-commit status (Y/N).
	AutoCommit *string `json:"autoCommit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengSessionDetail instantiates a new DamengSessionDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengSessionDetail(sessId int64) *DamengSessionDetail {
	this := DamengSessionDetail{}
	this.SessId = sessId
	return &this
}

// NewDamengSessionDetailWithDefaults instantiates a new DamengSessionDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengSessionDetailWithDefaults() *DamengSessionDetail {
	this := DamengSessionDetail{}
	return &this
}

// GetSessId returns the SessId field value.
func (o *DamengSessionDetail) GetSessId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessId
}

// GetSessIdOk returns a tuple with the SessId field value
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetSessIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessId, true
}

// SetSessId sets field value.
func (o *DamengSessionDetail) SetSessId(v int64) {
	o.SessId = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *DamengSessionDetail) SetUserName(v string) {
	o.UserName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DamengSessionDetail) SetState(v string) {
	o.State = &v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetClientIp() string {
	if o == nil || o.ClientIp == nil {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetClientIpOk() (*string, bool) {
	if o == nil || o.ClientIp == nil {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasClientIp() bool {
	return o != nil && o.ClientIp != nil
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *DamengSessionDetail) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetClientType() string {
	if o == nil || o.ClientType == nil {
		var ret string
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetClientTypeOk() (*string, bool) {
	if o == nil || o.ClientType == nil {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasClientType() bool {
	return o != nil && o.ClientType != nil
}

// SetClientType gets a reference to the given string and assigns it to the ClientType field.
func (o *DamengSessionDetail) SetClientType(v string) {
	o.ClientType = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetCreateTime() string {
	if o == nil || o.CreateTime == nil {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetCreateTimeOk() (*string, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasCreateTime() bool {
	return o != nil && o.CreateTime != nil
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *DamengSessionDetail) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetLastSendTime returns the LastSendTime field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetLastSendTime() string {
	if o == nil || o.LastSendTime == nil {
		var ret string
		return ret
	}
	return *o.LastSendTime
}

// GetLastSendTimeOk returns a tuple with the LastSendTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetLastSendTimeOk() (*string, bool) {
	if o == nil || o.LastSendTime == nil {
		return nil, false
	}
	return o.LastSendTime, true
}

// HasLastSendTime returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasLastSendTime() bool {
	return o != nil && o.LastSendTime != nil
}

// SetLastSendTime gets a reference to the given string and assigns it to the LastSendTime field.
func (o *DamengSessionDetail) SetLastSendTime(v string) {
	o.LastSendTime = &v
}

// GetSqlText returns the SqlText field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetSqlText() string {
	if o == nil || o.SqlText == nil {
		var ret string
		return ret
	}
	return *o.SqlText
}

// GetSqlTextOk returns a tuple with the SqlText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetSqlTextOk() (*string, bool) {
	if o == nil || o.SqlText == nil {
		return nil, false
	}
	return o.SqlText, true
}

// HasSqlText returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasSqlText() bool {
	return o != nil && o.SqlText != nil
}

// SetSqlText gets a reference to the given string and assigns it to the SqlText field.
func (o *DamengSessionDetail) SetSqlText(v string) {
	o.SqlText = &v
}

// GetFullSql returns the FullSql field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetFullSql() string {
	if o == nil || o.FullSql == nil {
		var ret string
		return ret
	}
	return *o.FullSql
}

// GetFullSqlOk returns a tuple with the FullSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetFullSqlOk() (*string, bool) {
	if o == nil || o.FullSql == nil {
		return nil, false
	}
	return o.FullSql, true
}

// HasFullSql returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasFullSql() bool {
	return o != nil && o.FullSql != nil
}

// SetFullSql gets a reference to the given string and assigns it to the FullSql field.
func (o *DamengSessionDetail) SetFullSql(v string) {
	o.FullSql = &v
}

// GetTrxId returns the TrxId field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetTrxId() string {
	if o == nil || o.TrxId == nil {
		var ret string
		return ret
	}
	return *o.TrxId
}

// GetTrxIdOk returns a tuple with the TrxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetTrxIdOk() (*string, bool) {
	if o == nil || o.TrxId == nil {
		return nil, false
	}
	return o.TrxId, true
}

// HasTrxId returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasTrxId() bool {
	return o != nil && o.TrxId != nil
}

// SetTrxId gets a reference to the given string and assigns it to the TrxId field.
func (o *DamengSessionDetail) SetTrxId(v string) {
	o.TrxId = &v
}

// GetTrxDurationSec returns the TrxDurationSec field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetTrxDurationSec() int64 {
	if o == nil || o.TrxDurationSec == nil {
		var ret int64
		return ret
	}
	return *o.TrxDurationSec
}

// GetTrxDurationSecOk returns a tuple with the TrxDurationSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetTrxDurationSecOk() (*int64, bool) {
	if o == nil || o.TrxDurationSec == nil {
		return nil, false
	}
	return o.TrxDurationSec, true
}

// HasTrxDurationSec returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasTrxDurationSec() bool {
	return o != nil && o.TrxDurationSec != nil
}

// SetTrxDurationSec gets a reference to the given int64 and assigns it to the TrxDurationSec field.
func (o *DamengSessionDetail) SetTrxDurationSec(v int64) {
	o.TrxDurationSec = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetAppName() string {
	if o == nil || o.AppName == nil {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetAppNameOk() (*string, bool) {
	if o == nil || o.AppName == nil {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasAppName() bool {
	return o != nil && o.AppName != nil
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *DamengSessionDetail) SetAppName(v string) {
	o.AppName = &v
}

// GetCurrSchema returns the CurrSchema field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetCurrSchema() string {
	if o == nil || o.CurrSchema == nil {
		var ret string
		return ret
	}
	return *o.CurrSchema
}

// GetCurrSchemaOk returns a tuple with the CurrSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetCurrSchemaOk() (*string, bool) {
	if o == nil || o.CurrSchema == nil {
		return nil, false
	}
	return o.CurrSchema, true
}

// HasCurrSchema returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasCurrSchema() bool {
	return o != nil && o.CurrSchema != nil
}

// SetCurrSchema gets a reference to the given string and assigns it to the CurrSchema field.
func (o *DamengSessionDetail) SetCurrSchema(v string) {
	o.CurrSchema = &v
}

// GetAutoCommit returns the AutoCommit field value if set, zero value otherwise.
func (o *DamengSessionDetail) GetAutoCommit() string {
	if o == nil || o.AutoCommit == nil {
		var ret string
		return ret
	}
	return *o.AutoCommit
}

// GetAutoCommitOk returns a tuple with the AutoCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionDetail) GetAutoCommitOk() (*string, bool) {
	if o == nil || o.AutoCommit == nil {
		return nil, false
	}
	return o.AutoCommit, true
}

// HasAutoCommit returns a boolean if a field has been set.
func (o *DamengSessionDetail) HasAutoCommit() bool {
	return o != nil && o.AutoCommit != nil
}

// SetAutoCommit gets a reference to the given string and assigns it to the AutoCommit field.
func (o *DamengSessionDetail) SetAutoCommit(v string) {
	o.AutoCommit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengSessionDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["sessId"] = o.SessId
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ClientIp != nil {
		toSerialize["clientIp"] = o.ClientIp
	}
	if o.ClientType != nil {
		toSerialize["clientType"] = o.ClientType
	}
	if o.CreateTime != nil {
		toSerialize["createTime"] = o.CreateTime
	}
	if o.LastSendTime != nil {
		toSerialize["lastSendTime"] = o.LastSendTime
	}
	if o.SqlText != nil {
		toSerialize["sqlText"] = o.SqlText
	}
	if o.FullSql != nil {
		toSerialize["fullSql"] = o.FullSql
	}
	if o.TrxId != nil {
		toSerialize["trxId"] = o.TrxId
	}
	if o.TrxDurationSec != nil {
		toSerialize["trxDurationSec"] = o.TrxDurationSec
	}
	if o.AppName != nil {
		toSerialize["appName"] = o.AppName
	}
	if o.CurrSchema != nil {
		toSerialize["currSchema"] = o.CurrSchema
	}
	if o.AutoCommit != nil {
		toSerialize["autoCommit"] = o.AutoCommit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengSessionDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessId         *int64  `json:"sessId"`
		UserName       *string `json:"userName,omitempty"`
		State          *string `json:"state,omitempty"`
		ClientIp       *string `json:"clientIp,omitempty"`
		ClientType     *string `json:"clientType,omitempty"`
		CreateTime     *string `json:"createTime,omitempty"`
		LastSendTime   *string `json:"lastSendTime,omitempty"`
		SqlText        *string `json:"sqlText,omitempty"`
		FullSql        *string `json:"fullSql,omitempty"`
		TrxId          *string `json:"trxId,omitempty"`
		TrxDurationSec *int64  `json:"trxDurationSec,omitempty"`
		AppName        *string `json:"appName,omitempty"`
		CurrSchema     *string `json:"currSchema,omitempty"`
		AutoCommit     *string `json:"autoCommit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SessId == nil {
		return fmt.Errorf("required field sessId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sessId", "userName", "state", "clientIp", "clientType", "createTime", "lastSendTime", "sqlText", "fullSql", "trxId", "trxDurationSec", "appName", "currSchema", "autoCommit"})
	} else {
		return err
	}
	o.SessId = *all.SessId
	o.UserName = all.UserName
	o.State = all.State
	o.ClientIp = all.ClientIp
	o.ClientType = all.ClientType
	o.CreateTime = all.CreateTime
	o.LastSendTime = all.LastSendTime
	o.SqlText = all.SqlText
	o.FullSql = all.FullSql
	o.TrxId = all.TrxId
	o.TrxDurationSec = all.TrxDurationSec
	o.AppName = all.AppName
	o.CurrSchema = all.CurrSchema
	o.AutoCommit = all.AutoCommit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
