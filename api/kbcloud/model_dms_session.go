// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsSession struct {
	// session ID
	Id int64 `json:"id"`
	// user name
	User string `json:"user"`
	// tenant name
	Tenant common.NullableString `json:"tenant,omitempty"`
	// client host information
	Host string `json:"host"`
	// database name
	Db common.NullableString `json:"db,omitempty"`
	// current command
	Command common.NullableString `json:"command,omitempty"`
	// session duration time
	Time common.NullableInt64 `json:"time,omitempty"`
	// session state
	State common.NullableString `json:"state,omitempty"`
	// additional session information
	Info common.NullableString `json:"info,omitempty"`
	// server ip
	Ip common.NullableString `json:"ip,omitempty"`
	// server port
	Port common.NullableInt64 `json:"port,omitempty"`
	// whether the session is protected and don't allow to be killed
	Protected *bool `json:"protected,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsSession instantiates a new DmsSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsSession(id int64, user string, host string) *DmsSession {
	this := DmsSession{}
	this.Id = id
	this.User = user
	this.Host = host
	return &this
}

// NewDmsSessionWithDefaults instantiates a new DmsSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsSessionWithDefaults() *DmsSession {
	this := DmsSession{}
	return &this
}

// GetId returns the Id field value.
func (o *DmsSession) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DmsSession) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DmsSession) SetId(v int64) {
	o.Id = v
}

// GetUser returns the User field value.
func (o *DmsSession) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *DmsSession) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value.
func (o *DmsSession) SetUser(v string) {
	o.User = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsSession) GetTenant() string {
	if o == nil || o.Tenant.Get() == nil {
		var ret string
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsSession) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *DmsSession) HasTenant() bool {
	return o != nil && o.Tenant.IsSet()
}

// SetTenant gets a reference to the given common.NullableString and assigns it to the Tenant field.
func (o *DmsSession) SetTenant(v string) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil.
func (o *DmsSession) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil.
func (o *DmsSession) UnsetTenant() {
	o.Tenant.Unset()
}

// GetHost returns the Host field value.
func (o *DmsSession) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *DmsSession) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value.
func (o *DmsSession) SetHost(v string) {
	o.Host = v
}

// GetDb returns the Db field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsSession) GetDb() string {
	if o == nil || o.Db.Get() == nil {
		var ret string
		return ret
	}
	return *o.Db.Get()
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsSession) GetDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Db.Get(), o.Db.IsSet()
}

// HasDb returns a boolean if a field has been set.
func (o *DmsSession) HasDb() bool {
	return o != nil && o.Db.IsSet()
}

// SetDb gets a reference to the given common.NullableString and assigns it to the Db field.
func (o *DmsSession) SetDb(v string) {
	o.Db.Set(&v)
}

// SetDbNil sets the value for Db to be an explicit nil.
func (o *DmsSession) SetDbNil() {
	o.Db.Set(nil)
}

// UnsetDb ensures that no value is present for Db, not even an explicit nil.
func (o *DmsSession) UnsetDb() {
	o.Db.Unset()
}

// GetCommand returns the Command field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsSession) GetCommand() string {
	if o == nil || o.Command.Get() == nil {
		var ret string
		return ret
	}
	return *o.Command.Get()
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsSession) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Command.Get(), o.Command.IsSet()
}

// HasCommand returns a boolean if a field has been set.
func (o *DmsSession) HasCommand() bool {
	return o != nil && o.Command.IsSet()
}

// SetCommand gets a reference to the given common.NullableString and assigns it to the Command field.
func (o *DmsSession) SetCommand(v string) {
	o.Command.Set(&v)
}

// SetCommandNil sets the value for Command to be an explicit nil.
func (o *DmsSession) SetCommandNil() {
	o.Command.Set(nil)
}

// UnsetCommand ensures that no value is present for Command, not even an explicit nil.
func (o *DmsSession) UnsetCommand() {
	o.Command.Unset()
}

// GetTime returns the Time field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsSession) GetTime() int64 {
	if o == nil || o.Time.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Time.Get()
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsSession) GetTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Time.Get(), o.Time.IsSet()
}

// HasTime returns a boolean if a field has been set.
func (o *DmsSession) HasTime() bool {
	return o != nil && o.Time.IsSet()
}

// SetTime gets a reference to the given common.NullableInt64 and assigns it to the Time field.
func (o *DmsSession) SetTime(v int64) {
	o.Time.Set(&v)
}

// SetTimeNil sets the value for Time to be an explicit nil.
func (o *DmsSession) SetTimeNil() {
	o.Time.Set(nil)
}

// UnsetTime ensures that no value is present for Time, not even an explicit nil.
func (o *DmsSession) UnsetTime() {
	o.Time.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsSession) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsSession) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *DmsSession) HasState() bool {
	return o != nil && o.State.IsSet()
}

// SetState gets a reference to the given common.NullableString and assigns it to the State field.
func (o *DmsSession) SetState(v string) {
	o.State.Set(&v)
}

// SetStateNil sets the value for State to be an explicit nil.
func (o *DmsSession) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil.
func (o *DmsSession) UnsetState() {
	o.State.Unset()
}

// GetInfo returns the Info field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsSession) GetInfo() string {
	if o == nil || o.Info.Get() == nil {
		var ret string
		return ret
	}
	return *o.Info.Get()
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsSession) GetInfoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Info.Get(), o.Info.IsSet()
}

// HasInfo returns a boolean if a field has been set.
func (o *DmsSession) HasInfo() bool {
	return o != nil && o.Info.IsSet()
}

// SetInfo gets a reference to the given common.NullableString and assigns it to the Info field.
func (o *DmsSession) SetInfo(v string) {
	o.Info.Set(&v)
}

// SetInfoNil sets the value for Info to be an explicit nil.
func (o *DmsSession) SetInfoNil() {
	o.Info.Set(nil)
}

// UnsetInfo ensures that no value is present for Info, not even an explicit nil.
func (o *DmsSession) UnsetInfo() {
	o.Info.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsSession) GetIp() string {
	if o == nil || o.Ip.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsSession) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *DmsSession) HasIp() bool {
	return o != nil && o.Ip.IsSet()
}

// SetIp gets a reference to the given common.NullableString and assigns it to the Ip field.
func (o *DmsSession) SetIp(v string) {
	o.Ip.Set(&v)
}

// SetIpNil sets the value for Ip to be an explicit nil.
func (o *DmsSession) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil.
func (o *DmsSession) UnsetIp() {
	o.Ip.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsSession) GetPort() int64 {
	if o == nil || o.Port.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsSession) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *DmsSession) HasPort() bool {
	return o != nil && o.Port.IsSet()
}

// SetPort gets a reference to the given common.NullableInt64 and assigns it to the Port field.
func (o *DmsSession) SetPort(v int64) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil.
func (o *DmsSession) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil.
func (o *DmsSession) UnsetPort() {
	o.Port.Unset()
}

// GetProtected returns the Protected field value if set, zero value otherwise.
func (o *DmsSession) GetProtected() bool {
	if o == nil || o.Protected == nil {
		var ret bool
		return ret
	}
	return *o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetProtectedOk() (*bool, bool) {
	if o == nil || o.Protected == nil {
		return nil, false
	}
	return o.Protected, true
}

// HasProtected returns a boolean if a field has been set.
func (o *DmsSession) HasProtected() bool {
	return o != nil && o.Protected != nil
}

// SetProtected gets a reference to the given bool and assigns it to the Protected field.
func (o *DmsSession) SetProtected(v bool) {
	o.Protected = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["user"] = o.User
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	toSerialize["host"] = o.Host
	if o.Db.IsSet() {
		toSerialize["db"] = o.Db.Get()
	}
	if o.Command.IsSet() {
		toSerialize["command"] = o.Command.Get()
	}
	if o.Time.IsSet() {
		toSerialize["time"] = o.Time.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.Info.IsSet() {
		toSerialize["info"] = o.Info.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Protected != nil {
		toSerialize["protected"] = o.Protected
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *int64                `json:"id"`
		User      *string               `json:"user"`
		Tenant    common.NullableString `json:"tenant,omitempty"`
		Host      *string               `json:"host"`
		Db        common.NullableString `json:"db,omitempty"`
		Command   common.NullableString `json:"command,omitempty"`
		Time      common.NullableInt64  `json:"time,omitempty"`
		State     common.NullableString `json:"state,omitempty"`
		Info      common.NullableString `json:"info,omitempty"`
		Ip        common.NullableString `json:"ip,omitempty"`
		Port      common.NullableInt64  `json:"port,omitempty"`
		Protected *bool                 `json:"protected,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.User == nil {
		return fmt.Errorf("required field user missing")
	}
	if all.Host == nil {
		return fmt.Errorf("required field host missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "user", "tenant", "host", "db", "command", "time", "state", "info", "ip", "port", "protected"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.User = *all.User
	o.Tenant = all.Tenant
	o.Host = *all.Host
	o.Db = all.Db
	o.Command = all.Command
	o.Time = all.Time
	o.State = all.State
	o.Info = all.Info
	o.Ip = all.Ip
	o.Port = all.Port
	o.Protected = all.Protected

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
