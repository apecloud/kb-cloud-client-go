// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type DolphinDBNodeHealth struct {
	Name *string `json:"name,omitempty"`
	Host *string `json:"host,omitempty"`
	Port *int32  `json:"port,omitempty"`
	// 0 data, 1 native Agent, 2 controller, 3 standalone, 4 compute
	Mode          *int32                `json:"mode,omitempty"`
	IsLeader      *bool                 `json:"isLeader,omitempty"`
	Process       *DolphinDBHealthCheck `json:"process,omitempty"`
	PortReachable *DolphinDBHealthCheck `json:"portReachable,omitempty"`
	Registered    *DolphinDBHealthCheck `json:"registered,omitempty"`
	Online        *DolphinDBHealthCheck `json:"online,omitempty"`
	Initialized   *DolphinDBHealthCheck `json:"initialized,omitempty"`
	Queryable     *DolphinDBHealthCheck `json:"queryable,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBNodeHealth instantiates a new DolphinDBNodeHealth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBNodeHealth() *DolphinDBNodeHealth {
	this := DolphinDBNodeHealth{}
	return &this
}

// NewDolphinDBNodeHealthWithDefaults instantiates a new DolphinDBNodeHealth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBNodeHealthWithDefaults() *DolphinDBNodeHealth {
	this := DolphinDBNodeHealth{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DolphinDBNodeHealth) SetName(v string) {
	o.Name = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *DolphinDBNodeHealth) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *DolphinDBNodeHealth) SetPort(v int32) {
	o.Port = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetMode() int32 {
	if o == nil || o.Mode == nil {
		var ret int32
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetModeOk() (*int32, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given int32 and assigns it to the Mode field.
func (o *DolphinDBNodeHealth) SetMode(v int32) {
	o.Mode = &v
}

// GetIsLeader returns the IsLeader field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetIsLeader() bool {
	if o == nil || o.IsLeader == nil {
		var ret bool
		return ret
	}
	return *o.IsLeader
}

// GetIsLeaderOk returns a tuple with the IsLeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetIsLeaderOk() (*bool, bool) {
	if o == nil || o.IsLeader == nil {
		return nil, false
	}
	return o.IsLeader, true
}

// HasIsLeader returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasIsLeader() bool {
	return o != nil && o.IsLeader != nil
}

// SetIsLeader gets a reference to the given bool and assigns it to the IsLeader field.
func (o *DolphinDBNodeHealth) SetIsLeader(v bool) {
	o.IsLeader = &v
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetProcess() DolphinDBHealthCheck {
	if o == nil || o.Process == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetProcessOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasProcess() bool {
	return o != nil && o.Process != nil
}

// SetProcess gets a reference to the given DolphinDBHealthCheck and assigns it to the Process field.
func (o *DolphinDBNodeHealth) SetProcess(v DolphinDBHealthCheck) {
	o.Process = &v
}

// GetPortReachable returns the PortReachable field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetPortReachable() DolphinDBHealthCheck {
	if o == nil || o.PortReachable == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.PortReachable
}

// GetPortReachableOk returns a tuple with the PortReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetPortReachableOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.PortReachable == nil {
		return nil, false
	}
	return o.PortReachable, true
}

// HasPortReachable returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasPortReachable() bool {
	return o != nil && o.PortReachable != nil
}

// SetPortReachable gets a reference to the given DolphinDBHealthCheck and assigns it to the PortReachable field.
func (o *DolphinDBNodeHealth) SetPortReachable(v DolphinDBHealthCheck) {
	o.PortReachable = &v
}

// GetRegistered returns the Registered field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetRegistered() DolphinDBHealthCheck {
	if o == nil || o.Registered == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.Registered
}

// GetRegisteredOk returns a tuple with the Registered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetRegisteredOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.Registered == nil {
		return nil, false
	}
	return o.Registered, true
}

// HasRegistered returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasRegistered() bool {
	return o != nil && o.Registered != nil
}

// SetRegistered gets a reference to the given DolphinDBHealthCheck and assigns it to the Registered field.
func (o *DolphinDBNodeHealth) SetRegistered(v DolphinDBHealthCheck) {
	o.Registered = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetOnline() DolphinDBHealthCheck {
	if o == nil || o.Online == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetOnlineOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.Online == nil {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasOnline() bool {
	return o != nil && o.Online != nil
}

// SetOnline gets a reference to the given DolphinDBHealthCheck and assigns it to the Online field.
func (o *DolphinDBNodeHealth) SetOnline(v DolphinDBHealthCheck) {
	o.Online = &v
}

// GetInitialized returns the Initialized field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetInitialized() DolphinDBHealthCheck {
	if o == nil || o.Initialized == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.Initialized
}

// GetInitializedOk returns a tuple with the Initialized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetInitializedOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.Initialized == nil {
		return nil, false
	}
	return o.Initialized, true
}

// HasInitialized returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasInitialized() bool {
	return o != nil && o.Initialized != nil
}

// SetInitialized gets a reference to the given DolphinDBHealthCheck and assigns it to the Initialized field.
func (o *DolphinDBNodeHealth) SetInitialized(v DolphinDBHealthCheck) {
	o.Initialized = &v
}

// GetQueryable returns the Queryable field value if set, zero value otherwise.
func (o *DolphinDBNodeHealth) GetQueryable() DolphinDBHealthCheck {
	if o == nil || o.Queryable == nil {
		var ret DolphinDBHealthCheck
		return ret
	}
	return *o.Queryable
}

// GetQueryableOk returns a tuple with the Queryable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBNodeHealth) GetQueryableOk() (*DolphinDBHealthCheck, bool) {
	if o == nil || o.Queryable == nil {
		return nil, false
	}
	return o.Queryable, true
}

// HasQueryable returns a boolean if a field has been set.
func (o *DolphinDBNodeHealth) HasQueryable() bool {
	return o != nil && o.Queryable != nil
}

// SetQueryable gets a reference to the given DolphinDBHealthCheck and assigns it to the Queryable field.
func (o *DolphinDBNodeHealth) SetQueryable(v DolphinDBHealthCheck) {
	o.Queryable = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBNodeHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.IsLeader != nil {
		toSerialize["isLeader"] = o.IsLeader
	}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}
	if o.PortReachable != nil {
		toSerialize["portReachable"] = o.PortReachable
	}
	if o.Registered != nil {
		toSerialize["registered"] = o.Registered
	}
	if o.Online != nil {
		toSerialize["online"] = o.Online
	}
	if o.Initialized != nil {
		toSerialize["initialized"] = o.Initialized
	}
	if o.Queryable != nil {
		toSerialize["queryable"] = o.Queryable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBNodeHealth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name          *string               `json:"name,omitempty"`
		Host          *string               `json:"host,omitempty"`
		Port          *int32                `json:"port,omitempty"`
		Mode          *int32                `json:"mode,omitempty"`
		IsLeader      *bool                 `json:"isLeader,omitempty"`
		Process       *DolphinDBHealthCheck `json:"process,omitempty"`
		PortReachable *DolphinDBHealthCheck `json:"portReachable,omitempty"`
		Registered    *DolphinDBHealthCheck `json:"registered,omitempty"`
		Online        *DolphinDBHealthCheck `json:"online,omitempty"`
		Initialized   *DolphinDBHealthCheck `json:"initialized,omitempty"`
		Queryable     *DolphinDBHealthCheck `json:"queryable,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "host", "port", "mode", "isLeader", "process", "portReachable", "registered", "online", "initialized", "queryable"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Host = all.Host
	o.Port = all.Port
	o.Mode = all.Mode
	o.IsLeader = all.IsLeader
	if all.Process != nil && all.Process.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Process = all.Process
	if all.PortReachable != nil && all.PortReachable.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PortReachable = all.PortReachable
	if all.Registered != nil && all.Registered.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Registered = all.Registered
	if all.Online != nil && all.Online.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Online = all.Online
	if all.Initialized != nil && all.Initialized.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Initialized = all.Initialized
	if all.Queryable != nil && all.Queryable.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Queryable = all.Queryable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
