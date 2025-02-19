// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DmsSession struct {
	// session id
	Id *string `json:"id,omitempty"`
	// user name
	User *string `json:"user,omitempty"`
	// tenant name
	Tenant *string `json:"tenant,omitempty"`
	// host name
	Host *string `json:"host,omitempty"`
	// database name
	Db *string `json:"db,omitempty"`
	// command
	Command *string `json:"command,omitempty"`
	// time
	Time *string `json:"time,omitempty"`
	// state
	State *string `json:"state,omitempty"`
	// info
	Info *string `json:"info,omitempty"`
	// ip
	Ip *string `json:"ip,omitempty"`
	// port
	Port *string `json:"port,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsSession instantiates a new DmsSession object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsSession() *DmsSession {
	this := DmsSession{}
	return &this
}

// NewDmsSessionWithDefaults instantiates a new DmsSession object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsSessionWithDefaults() *DmsSession {
	this := DmsSession{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DmsSession) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DmsSession) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DmsSession) SetId(v string) {
	o.Id = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *DmsSession) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *DmsSession) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *DmsSession) SetUser(v string) {
	o.User = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *DmsSession) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *DmsSession) HasTenant() bool {
	return o != nil && o.Tenant != nil
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *DmsSession) SetTenant(v string) {
	o.Tenant = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *DmsSession) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *DmsSession) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *DmsSession) SetHost(v string) {
	o.Host = &v
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *DmsSession) GetDb() string {
	if o == nil || o.Db == nil {
		var ret string
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetDbOk() (*string, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *DmsSession) HasDb() bool {
	return o != nil && o.Db != nil
}

// SetDb gets a reference to the given string and assigns it to the Db field.
func (o *DmsSession) SetDb(v string) {
	o.Db = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *DmsSession) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *DmsSession) HasCommand() bool {
	return o != nil && o.Command != nil
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *DmsSession) SetCommand(v string) {
	o.Command = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *DmsSession) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *DmsSession) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *DmsSession) SetTime(v string) {
	o.Time = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DmsSession) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DmsSession) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *DmsSession) SetState(v string) {
	o.State = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *DmsSession) GetInfo() string {
	if o == nil || o.Info == nil {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetInfoOk() (*string, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *DmsSession) HasInfo() bool {
	return o != nil && o.Info != nil
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *DmsSession) SetInfo(v string) {
	o.Info = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *DmsSession) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *DmsSession) HasIp() bool {
	return o != nil && o.Ip != nil
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *DmsSession) SetIp(v string) {
	o.Ip = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DmsSession) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsSession) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DmsSession) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *DmsSession) SetPort(v string) {
	o.Port = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Db != nil {
		toSerialize["db"] = o.Db
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsSession) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id      *string `json:"id,omitempty"`
		User    *string `json:"user,omitempty"`
		Tenant  *string `json:"tenant,omitempty"`
		Host    *string `json:"host,omitempty"`
		Db      *string `json:"db,omitempty"`
		Command *string `json:"command,omitempty"`
		Time    *string `json:"time,omitempty"`
		State   *string `json:"state,omitempty"`
		Info    *string `json:"info,omitempty"`
		Ip      *string `json:"ip,omitempty"`
		Port    *string `json:"port,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "user", "tenant", "host", "db", "command", "time", "state", "info", "ip", "port"})
	} else {
		return err
	}
	o.Id = all.Id
	o.User = all.User
	o.Tenant = all.Tenant
	o.Host = all.Host
	o.Db = all.Db
	o.Command = all.Command
	o.Time = all.Time
	o.State = all.State
	o.Info = all.Info
	o.Ip = all.Ip
	o.Port = all.Port

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
