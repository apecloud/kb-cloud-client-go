// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KingbaseLockRow struct {
	Pid                int64   `json:"pid"`
	LockType           string  `json:"lockType"`
	Mode               string  `json:"mode"`
	Granted            bool    `json:"granted"`
	RelationOid        *string `json:"relationOid,omitempty"`
	RelationName       *string `json:"relationName,omitempty"`
	DatabaseName       *string `json:"databaseName,omitempty"`
	TransactionId      *string `json:"transactionID,omitempty"`
	VirtualTransaction *string `json:"virtualTransaction,omitempty"`
	VirtualXid         *string `json:"virtualXID,omitempty"`
	Tuple              *string `json:"tuple,omitempty"`
	Page               *string `json:"page,omitempty"`
	ClassId            *string `json:"classID,omitempty"`
	ObjId              *string `json:"objID,omitempty"`
	ObjSubId           *string `json:"objSubID,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKingbaseLockRow instantiates a new KingbaseLockRow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKingbaseLockRow(pid int64, lockType string, mode string, granted bool) *KingbaseLockRow {
	this := KingbaseLockRow{}
	this.Pid = pid
	this.LockType = lockType
	this.Mode = mode
	this.Granted = granted
	return &this
}

// NewKingbaseLockRowWithDefaults instantiates a new KingbaseLockRow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKingbaseLockRowWithDefaults() *KingbaseLockRow {
	this := KingbaseLockRow{}
	return &this
}

// GetPid returns the Pid field value.
func (o *KingbaseLockRow) GetPid() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Pid
}

// GetPidOk returns a tuple with the Pid field value
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetPidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pid, true
}

// SetPid sets field value.
func (o *KingbaseLockRow) SetPid(v int64) {
	o.Pid = v
}

// GetLockType returns the LockType field value.
func (o *KingbaseLockRow) GetLockType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LockType
}

// GetLockTypeOk returns a tuple with the LockType field value
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetLockTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockType, true
}

// SetLockType sets field value.
func (o *KingbaseLockRow) SetLockType(v string) {
	o.LockType = v
}

// GetMode returns the Mode field value.
func (o *KingbaseLockRow) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *KingbaseLockRow) SetMode(v string) {
	o.Mode = v
}

// GetGranted returns the Granted field value.
func (o *KingbaseLockRow) GetGranted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Granted
}

// GetGrantedOk returns a tuple with the Granted field value
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetGrantedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granted, true
}

// SetGranted sets field value.
func (o *KingbaseLockRow) SetGranted(v bool) {
	o.Granted = v
}

// GetRelationOid returns the RelationOid field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetRelationOid() string {
	if o == nil || o.RelationOid == nil {
		var ret string
		return ret
	}
	return *o.RelationOid
}

// GetRelationOidOk returns a tuple with the RelationOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetRelationOidOk() (*string, bool) {
	if o == nil || o.RelationOid == nil {
		return nil, false
	}
	return o.RelationOid, true
}

// HasRelationOid returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasRelationOid() bool {
	return o != nil && o.RelationOid != nil
}

// SetRelationOid gets a reference to the given string and assigns it to the RelationOid field.
func (o *KingbaseLockRow) SetRelationOid(v string) {
	o.RelationOid = &v
}

// GetRelationName returns the RelationName field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetRelationName() string {
	if o == nil || o.RelationName == nil {
		var ret string
		return ret
	}
	return *o.RelationName
}

// GetRelationNameOk returns a tuple with the RelationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetRelationNameOk() (*string, bool) {
	if o == nil || o.RelationName == nil {
		return nil, false
	}
	return o.RelationName, true
}

// HasRelationName returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasRelationName() bool {
	return o != nil && o.RelationName != nil
}

// SetRelationName gets a reference to the given string and assigns it to the RelationName field.
func (o *KingbaseLockRow) SetRelationName(v string) {
	o.RelationName = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasDatabaseName() bool {
	return o != nil && o.DatabaseName != nil
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *KingbaseLockRow) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasTransactionId() bool {
	return o != nil && o.TransactionId != nil
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *KingbaseLockRow) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetVirtualTransaction returns the VirtualTransaction field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetVirtualTransaction() string {
	if o == nil || o.VirtualTransaction == nil {
		var ret string
		return ret
	}
	return *o.VirtualTransaction
}

// GetVirtualTransactionOk returns a tuple with the VirtualTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetVirtualTransactionOk() (*string, bool) {
	if o == nil || o.VirtualTransaction == nil {
		return nil, false
	}
	return o.VirtualTransaction, true
}

// HasVirtualTransaction returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasVirtualTransaction() bool {
	return o != nil && o.VirtualTransaction != nil
}

// SetVirtualTransaction gets a reference to the given string and assigns it to the VirtualTransaction field.
func (o *KingbaseLockRow) SetVirtualTransaction(v string) {
	o.VirtualTransaction = &v
}

// GetVirtualXid returns the VirtualXid field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetVirtualXid() string {
	if o == nil || o.VirtualXid == nil {
		var ret string
		return ret
	}
	return *o.VirtualXid
}

// GetVirtualXidOk returns a tuple with the VirtualXid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetVirtualXidOk() (*string, bool) {
	if o == nil || o.VirtualXid == nil {
		return nil, false
	}
	return o.VirtualXid, true
}

// HasVirtualXid returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasVirtualXid() bool {
	return o != nil && o.VirtualXid != nil
}

// SetVirtualXid gets a reference to the given string and assigns it to the VirtualXid field.
func (o *KingbaseLockRow) SetVirtualXid(v string) {
	o.VirtualXid = &v
}

// GetTuple returns the Tuple field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetTuple() string {
	if o == nil || o.Tuple == nil {
		var ret string
		return ret
	}
	return *o.Tuple
}

// GetTupleOk returns a tuple with the Tuple field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetTupleOk() (*string, bool) {
	if o == nil || o.Tuple == nil {
		return nil, false
	}
	return o.Tuple, true
}

// HasTuple returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasTuple() bool {
	return o != nil && o.Tuple != nil
}

// SetTuple gets a reference to the given string and assigns it to the Tuple field.
func (o *KingbaseLockRow) SetTuple(v string) {
	o.Tuple = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetPage() string {
	if o == nil || o.Page == nil {
		var ret string
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetPageOk() (*string, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given string and assigns it to the Page field.
func (o *KingbaseLockRow) SetPage(v string) {
	o.Page = &v
}

// GetClassId returns the ClassId field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetClassId() string {
	if o == nil || o.ClassId == nil {
		var ret string
		return ret
	}
	return *o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetClassIdOk() (*string, bool) {
	if o == nil || o.ClassId == nil {
		return nil, false
	}
	return o.ClassId, true
}

// HasClassId returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasClassId() bool {
	return o != nil && o.ClassId != nil
}

// SetClassId gets a reference to the given string and assigns it to the ClassId field.
func (o *KingbaseLockRow) SetClassId(v string) {
	o.ClassId = &v
}

// GetObjId returns the ObjId field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetObjId() string {
	if o == nil || o.ObjId == nil {
		var ret string
		return ret
	}
	return *o.ObjId
}

// GetObjIdOk returns a tuple with the ObjId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetObjIdOk() (*string, bool) {
	if o == nil || o.ObjId == nil {
		return nil, false
	}
	return o.ObjId, true
}

// HasObjId returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasObjId() bool {
	return o != nil && o.ObjId != nil
}

// SetObjId gets a reference to the given string and assigns it to the ObjId field.
func (o *KingbaseLockRow) SetObjId(v string) {
	o.ObjId = &v
}

// GetObjSubId returns the ObjSubId field value if set, zero value otherwise.
func (o *KingbaseLockRow) GetObjSubId() string {
	if o == nil || o.ObjSubId == nil {
		var ret string
		return ret
	}
	return *o.ObjSubId
}

// GetObjSubIdOk returns a tuple with the ObjSubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KingbaseLockRow) GetObjSubIdOk() (*string, bool) {
	if o == nil || o.ObjSubId == nil {
		return nil, false
	}
	return o.ObjSubId, true
}

// HasObjSubId returns a boolean if a field has been set.
func (o *KingbaseLockRow) HasObjSubId() bool {
	return o != nil && o.ObjSubId != nil
}

// SetObjSubId gets a reference to the given string and assigns it to the ObjSubId field.
func (o *KingbaseLockRow) SetObjSubId(v string) {
	o.ObjSubId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o KingbaseLockRow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pid"] = o.Pid
	toSerialize["lockType"] = o.LockType
	toSerialize["mode"] = o.Mode
	toSerialize["granted"] = o.Granted
	if o.RelationOid != nil {
		toSerialize["relationOid"] = o.RelationOid
	}
	if o.RelationName != nil {
		toSerialize["relationName"] = o.RelationName
	}
	if o.DatabaseName != nil {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if o.TransactionId != nil {
		toSerialize["transactionID"] = o.TransactionId
	}
	if o.VirtualTransaction != nil {
		toSerialize["virtualTransaction"] = o.VirtualTransaction
	}
	if o.VirtualXid != nil {
		toSerialize["virtualXID"] = o.VirtualXid
	}
	if o.Tuple != nil {
		toSerialize["tuple"] = o.Tuple
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.ClassId != nil {
		toSerialize["classID"] = o.ClassId
	}
	if o.ObjId != nil {
		toSerialize["objID"] = o.ObjId
	}
	if o.ObjSubId != nil {
		toSerialize["objSubID"] = o.ObjSubId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KingbaseLockRow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pid                *int64  `json:"pid"`
		LockType           *string `json:"lockType"`
		Mode               *string `json:"mode"`
		Granted            *bool   `json:"granted"`
		RelationOid        *string `json:"relationOid,omitempty"`
		RelationName       *string `json:"relationName,omitempty"`
		DatabaseName       *string `json:"databaseName,omitempty"`
		TransactionId      *string `json:"transactionID,omitempty"`
		VirtualTransaction *string `json:"virtualTransaction,omitempty"`
		VirtualXid         *string `json:"virtualXID,omitempty"`
		Tuple              *string `json:"tuple,omitempty"`
		Page               *string `json:"page,omitempty"`
		ClassId            *string `json:"classID,omitempty"`
		ObjId              *string `json:"objID,omitempty"`
		ObjSubId           *string `json:"objSubID,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Pid == nil {
		return fmt.Errorf("required field pid missing")
	}
	if all.LockType == nil {
		return fmt.Errorf("required field lockType missing")
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.Granted == nil {
		return fmt.Errorf("required field granted missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pid", "lockType", "mode", "granted", "relationOid", "relationName", "databaseName", "transactionID", "virtualTransaction", "virtualXID", "tuple", "page", "classID", "objID", "objSubID"})
	} else {
		return err
	}
	o.Pid = *all.Pid
	o.LockType = *all.LockType
	o.Mode = *all.Mode
	o.Granted = *all.Granted
	o.RelationOid = all.RelationOid
	o.RelationName = all.RelationName
	o.DatabaseName = all.DatabaseName
	o.TransactionId = all.TransactionId
	o.VirtualTransaction = all.VirtualTransaction
	o.VirtualXid = all.VirtualXid
	o.Tuple = all.Tuple
	o.Page = all.Page
	o.ClassId = all.ClassId
	o.ObjId = all.ObjId
	o.ObjSubId = all.ObjSubId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
