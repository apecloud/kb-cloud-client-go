// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengSessionListItem struct {
	// Dameng session ID from V$SESSIONS.
	SessId int64 `json:"sessId"`
	// Database user name.
	UserName *string `json:"userName,omitempty"`
	// Session state (ACTIVE, IDLE, etc.).
	State *string `json:"state,omitempty"`
	// Client IP address.
	ClientIp *string `json:"clientIp,omitempty"`
	// Transaction ID.
	TrxId *string `json:"trxId,omitempty"`
	// Whether this session has locks from V$LOCK.
	HasLock bool `json:"hasLock"`
	// Current SQL text from V$SESSIONS.
	SqlText *string `json:"sqlText,omitempty"`
	// Current schema (database context).
	CurrSchema *string `json:"currSchema,omitempty"`
	// Application name.
	AppName *string `json:"appName,omitempty"`
	// Session duration in seconds.
	DurationSec *int64 `json:"durationSec,omitempty"`
	// Client type.
	ClientType *string `json:"clientType,omitempty"`
	// Auto-commit status.
	AutoCommit *string `json:"autoCommit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengSessionListItem instantiates a new DamengSessionListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengSessionListItem(sessId int64, hasLock bool) *DamengSessionListItem {
	this := DamengSessionListItem{}
	this.SessId = sessId
	this.HasLock = hasLock
	return &this
}

// NewDamengSessionListItemWithDefaults instantiates a new DamengSessionListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengSessionListItemWithDefaults() *DamengSessionListItem {
	this := DamengSessionListItem{}
	return &this
}

// GetSessId returns the SessId field value.
func (o *DamengSessionListItem) GetSessId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessId
}

// GetSessIdOk returns a tuple with the SessId field value
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetSessIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessId, true
}

// SetSessId sets field value.
func (o *DamengSessionListItem) SetSessId(v int64) {
	o.SessId = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *DamengSessionListItem) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *DamengSessionListItem) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *DamengSessionListItem) SetUserName(v string) {
	o.UserName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DamengSessionListItem) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DamengSessionListItem) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DamengSessionListItem) SetState(v string) {
	o.State = &v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *DamengSessionListItem) GetClientIp() string {
	if o == nil || o.ClientIp == nil {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetClientIpOk() (*string, bool) {
	if o == nil || o.ClientIp == nil {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *DamengSessionListItem) HasClientIp() bool {
	return o != nil && o.ClientIp != nil
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *DamengSessionListItem) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetTrxId returns the TrxId field value if set, zero value otherwise.
func (o *DamengSessionListItem) GetTrxId() string {
	if o == nil || o.TrxId == nil {
		var ret string
		return ret
	}
	return *o.TrxId
}

// GetTrxIdOk returns a tuple with the TrxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetTrxIdOk() (*string, bool) {
	if o == nil || o.TrxId == nil {
		return nil, false
	}
	return o.TrxId, true
}

// HasTrxId returns a boolean if a field has been set.
func (o *DamengSessionListItem) HasTrxId() bool {
	return o != nil && o.TrxId != nil
}

// SetTrxId gets a reference to the given string and assigns it to the TrxId field.
func (o *DamengSessionListItem) SetTrxId(v string) {
	o.TrxId = &v
}

// GetHasLock returns the HasLock field value.
func (o *DamengSessionListItem) GetHasLock() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasLock
}

// GetHasLockOk returns a tuple with the HasLock field value
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetHasLockOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasLock, true
}

// SetHasLock sets field value.
func (o *DamengSessionListItem) SetHasLock(v bool) {
	o.HasLock = v
}

// GetSqlText returns the SqlText field value if set, zero value otherwise.
func (o *DamengSessionListItem) GetSqlText() string {
	if o == nil || o.SqlText == nil {
		var ret string
		return ret
	}
	return *o.SqlText
}

// GetSqlTextOk returns a tuple with the SqlText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetSqlTextOk() (*string, bool) {
	if o == nil || o.SqlText == nil {
		return nil, false
	}
	return o.SqlText, true
}

// HasSqlText returns a boolean if a field has been set.
func (o *DamengSessionListItem) HasSqlText() bool {
	return o != nil && o.SqlText != nil
}

// SetSqlText gets a reference to the given string and assigns it to the SqlText field.
func (o *DamengSessionListItem) SetSqlText(v string) {
	o.SqlText = &v
}

// GetCurrSchema returns the CurrSchema field value if set, zero value otherwise.
func (o *DamengSessionListItem) GetCurrSchema() string {
	if o == nil || o.CurrSchema == nil {
		var ret string
		return ret
	}
	return *o.CurrSchema
}

// GetCurrSchemaOk returns a tuple with the CurrSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetCurrSchemaOk() (*string, bool) {
	if o == nil || o.CurrSchema == nil {
		return nil, false
	}
	return o.CurrSchema, true
}

// HasCurrSchema returns a boolean if a field has been set.
func (o *DamengSessionListItem) HasCurrSchema() bool {
	return o != nil && o.CurrSchema != nil
}

// SetCurrSchema gets a reference to the given string and assigns it to the CurrSchema field.
func (o *DamengSessionListItem) SetCurrSchema(v string) {
	o.CurrSchema = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *DamengSessionListItem) GetAppName() string {
	if o == nil || o.AppName == nil {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetAppNameOk() (*string, bool) {
	if o == nil || o.AppName == nil {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *DamengSessionListItem) HasAppName() bool {
	return o != nil && o.AppName != nil
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *DamengSessionListItem) SetAppName(v string) {
	o.AppName = &v
}

// GetDurationSec returns the DurationSec field value if set, zero value otherwise.
func (o *DamengSessionListItem) GetDurationSec() int64 {
	if o == nil || o.DurationSec == nil {
		var ret int64
		return ret
	}
	return *o.DurationSec
}

// GetDurationSecOk returns a tuple with the DurationSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetDurationSecOk() (*int64, bool) {
	if o == nil || o.DurationSec == nil {
		return nil, false
	}
	return o.DurationSec, true
}

// HasDurationSec returns a boolean if a field has been set.
func (o *DamengSessionListItem) HasDurationSec() bool {
	return o != nil && o.DurationSec != nil
}

// SetDurationSec gets a reference to the given int64 and assigns it to the DurationSec field.
func (o *DamengSessionListItem) SetDurationSec(v int64) {
	o.DurationSec = &v
}

// GetClientType returns the ClientType field value if set, zero value otherwise.
func (o *DamengSessionListItem) GetClientType() string {
	if o == nil || o.ClientType == nil {
		var ret string
		return ret
	}
	return *o.ClientType
}

// GetClientTypeOk returns a tuple with the ClientType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetClientTypeOk() (*string, bool) {
	if o == nil || o.ClientType == nil {
		return nil, false
	}
	return o.ClientType, true
}

// HasClientType returns a boolean if a field has been set.
func (o *DamengSessionListItem) HasClientType() bool {
	return o != nil && o.ClientType != nil
}

// SetClientType gets a reference to the given string and assigns it to the ClientType field.
func (o *DamengSessionListItem) SetClientType(v string) {
	o.ClientType = &v
}

// GetAutoCommit returns the AutoCommit field value if set, zero value otherwise.
func (o *DamengSessionListItem) GetAutoCommit() string {
	if o == nil || o.AutoCommit == nil {
		var ret string
		return ret
	}
	return *o.AutoCommit
}

// GetAutoCommitOk returns a tuple with the AutoCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengSessionListItem) GetAutoCommitOk() (*string, bool) {
	if o == nil || o.AutoCommit == nil {
		return nil, false
	}
	return o.AutoCommit, true
}

// HasAutoCommit returns a boolean if a field has been set.
func (o *DamengSessionListItem) HasAutoCommit() bool {
	return o != nil && o.AutoCommit != nil
}

// SetAutoCommit gets a reference to the given string and assigns it to the AutoCommit field.
func (o *DamengSessionListItem) SetAutoCommit(v string) {
	o.AutoCommit = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengSessionListItem) MarshalJSON() ([]byte, error) {
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
	if o.TrxId != nil {
		toSerialize["trxId"] = o.TrxId
	}
	toSerialize["hasLock"] = o.HasLock
	if o.SqlText != nil {
		toSerialize["sqlText"] = o.SqlText
	}
	if o.CurrSchema != nil {
		toSerialize["currSchema"] = o.CurrSchema
	}
	if o.AppName != nil {
		toSerialize["appName"] = o.AppName
	}
	if o.DurationSec != nil {
		toSerialize["durationSec"] = o.DurationSec
	}
	if o.ClientType != nil {
		toSerialize["clientType"] = o.ClientType
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
func (o *DamengSessionListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SessId      *int64  `json:"sessId"`
		UserName    *string `json:"userName,omitempty"`
		State       *string `json:"state,omitempty"`
		ClientIp    *string `json:"clientIp,omitempty"`
		TrxId       *string `json:"trxId,omitempty"`
		HasLock     *bool   `json:"hasLock"`
		SqlText     *string `json:"sqlText,omitempty"`
		CurrSchema  *string `json:"currSchema,omitempty"`
		AppName     *string `json:"appName,omitempty"`
		DurationSec *int64  `json:"durationSec,omitempty"`
		ClientType  *string `json:"clientType,omitempty"`
		AutoCommit  *string `json:"autoCommit,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SessId == nil {
		return fmt.Errorf("required field sessId missing")
	}
	if all.HasLock == nil {
		return fmt.Errorf("required field hasLock missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sessId", "userName", "state", "clientIp", "trxId", "hasLock", "sqlText", "currSchema", "appName", "durationSec", "clientType", "autoCommit"})
	} else {
		return err
	}
	o.SessId = *all.SessId
	o.UserName = all.UserName
	o.State = all.State
	o.ClientIp = all.ClientIp
	o.TrxId = all.TrxId
	o.HasLock = *all.HasLock
	o.SqlText = all.SqlText
	o.CurrSchema = all.CurrSchema
	o.AppName = all.AppName
	o.DurationSec = all.DurationSec
	o.ClientType = all.ClientType
	o.AutoCommit = all.AutoCommit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
