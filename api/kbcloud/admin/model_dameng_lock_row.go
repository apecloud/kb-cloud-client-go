// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengLockRow struct {
	// Transaction ID holding the lock.
	TrxId int64 `json:"trxId"`
	// Lock type.
	LockType *string `json:"lockType,omitempty"`
	// Lock mode.
	LockMode *string `json:"lockMode,omitempty"`
	// Session ID associated with this lock.
	SessId *int64 `json:"sessId,omitempty"`
	// Client IP address for the session.
	ClientIp *string `json:"clientIp,omitempty"`
	// SQL text for the session.
	SqlText *string `json:"sqlText,omitempty"`
	// User name for the session.
	UserName *string `json:"userName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengLockRow instantiates a new DamengLockRow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengLockRow(trxId int64) *DamengLockRow {
	this := DamengLockRow{}
	this.TrxId = trxId
	return &this
}

// NewDamengLockRowWithDefaults instantiates a new DamengLockRow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengLockRowWithDefaults() *DamengLockRow {
	this := DamengLockRow{}
	return &this
}

// GetTrxId returns the TrxId field value.
func (o *DamengLockRow) GetTrxId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TrxId
}

// GetTrxIdOk returns a tuple with the TrxId field value
// and a boolean to check if the value has been set.
func (o *DamengLockRow) GetTrxIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrxId, true
}

// SetTrxId sets field value.
func (o *DamengLockRow) SetTrxId(v int64) {
	o.TrxId = v
}

// GetLockType returns the LockType field value if set, zero value otherwise.
func (o *DamengLockRow) GetLockType() string {
	if o == nil || o.LockType == nil {
		var ret string
		return ret
	}
	return *o.LockType
}

// GetLockTypeOk returns a tuple with the LockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLockRow) GetLockTypeOk() (*string, bool) {
	if o == nil || o.LockType == nil {
		return nil, false
	}
	return o.LockType, true
}

// HasLockType returns a boolean if a field has been set.
func (o *DamengLockRow) HasLockType() bool {
	return o != nil && o.LockType != nil
}

// SetLockType gets a reference to the given string and assigns it to the LockType field.
func (o *DamengLockRow) SetLockType(v string) {
	o.LockType = &v
}

// GetLockMode returns the LockMode field value if set, zero value otherwise.
func (o *DamengLockRow) GetLockMode() string {
	if o == nil || o.LockMode == nil {
		var ret string
		return ret
	}
	return *o.LockMode
}

// GetLockModeOk returns a tuple with the LockMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLockRow) GetLockModeOk() (*string, bool) {
	if o == nil || o.LockMode == nil {
		return nil, false
	}
	return o.LockMode, true
}

// HasLockMode returns a boolean if a field has been set.
func (o *DamengLockRow) HasLockMode() bool {
	return o != nil && o.LockMode != nil
}

// SetLockMode gets a reference to the given string and assigns it to the LockMode field.
func (o *DamengLockRow) SetLockMode(v string) {
	o.LockMode = &v
}

// GetSessId returns the SessId field value if set, zero value otherwise.
func (o *DamengLockRow) GetSessId() int64 {
	if o == nil || o.SessId == nil {
		var ret int64
		return ret
	}
	return *o.SessId
}

// GetSessIdOk returns a tuple with the SessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLockRow) GetSessIdOk() (*int64, bool) {
	if o == nil || o.SessId == nil {
		return nil, false
	}
	return o.SessId, true
}

// HasSessId returns a boolean if a field has been set.
func (o *DamengLockRow) HasSessId() bool {
	return o != nil && o.SessId != nil
}

// SetSessId gets a reference to the given int64 and assigns it to the SessId field.
func (o *DamengLockRow) SetSessId(v int64) {
	o.SessId = &v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *DamengLockRow) GetClientIp() string {
	if o == nil || o.ClientIp == nil {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLockRow) GetClientIpOk() (*string, bool) {
	if o == nil || o.ClientIp == nil {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *DamengLockRow) HasClientIp() bool {
	return o != nil && o.ClientIp != nil
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *DamengLockRow) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetSqlText returns the SqlText field value if set, zero value otherwise.
func (o *DamengLockRow) GetSqlText() string {
	if o == nil || o.SqlText == nil {
		var ret string
		return ret
	}
	return *o.SqlText
}

// GetSqlTextOk returns a tuple with the SqlText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLockRow) GetSqlTextOk() (*string, bool) {
	if o == nil || o.SqlText == nil {
		return nil, false
	}
	return o.SqlText, true
}

// HasSqlText returns a boolean if a field has been set.
func (o *DamengLockRow) HasSqlText() bool {
	return o != nil && o.SqlText != nil
}

// SetSqlText gets a reference to the given string and assigns it to the SqlText field.
func (o *DamengLockRow) SetSqlText(v string) {
	o.SqlText = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *DamengLockRow) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamengLockRow) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *DamengLockRow) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *DamengLockRow) SetUserName(v string) {
	o.UserName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengLockRow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["trxId"] = o.TrxId
	if o.LockType != nil {
		toSerialize["lockType"] = o.LockType
	}
	if o.LockMode != nil {
		toSerialize["lockMode"] = o.LockMode
	}
	if o.SessId != nil {
		toSerialize["sessId"] = o.SessId
	}
	if o.ClientIp != nil {
		toSerialize["clientIp"] = o.ClientIp
	}
	if o.SqlText != nil {
		toSerialize["sqlText"] = o.SqlText
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengLockRow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TrxId    *int64  `json:"trxId"`
		LockType *string `json:"lockType,omitempty"`
		LockMode *string `json:"lockMode,omitempty"`
		SessId   *int64  `json:"sessId,omitempty"`
		ClientIp *string `json:"clientIp,omitempty"`
		SqlText  *string `json:"sqlText,omitempty"`
		UserName *string `json:"userName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TrxId == nil {
		return fmt.Errorf("required field trxId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"trxId", "lockType", "lockMode", "sessId", "clientIp", "sqlText", "userName"})
	} else {
		return err
	}
	o.TrxId = *all.TrxId
	o.LockType = all.LockType
	o.LockMode = all.LockMode
	o.SessId = all.SessId
	o.ClientIp = all.ClientIp
	o.SqlText = all.SqlText
	o.UserName = all.UserName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
