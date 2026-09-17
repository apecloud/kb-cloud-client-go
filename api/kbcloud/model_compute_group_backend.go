// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ComputeGroupBackend struct {
	// Opaque backend ID; do not convert to a JavaScript number.
	Id             *string `json:"id,omitempty"`
	Host           *string `json:"host,omitempty"`
	HeartbeatPort  *int64  `json:"heartbeatPort,omitempty"`
	Alive          *bool   `json:"alive,omitempty"`
	Decommissioned *bool   `json:"decommissioned,omitempty"`
	TabletCount    *int64  `json:"tabletCount,omitempty"`
	LastHeartbeat  *string `json:"lastHeartbeat,omitempty"`
	Version        *string `json:"version,omitempty"`
	ErrorMessage   *string `json:"errorMessage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComputeGroupBackend instantiates a new ComputeGroupBackend object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComputeGroupBackend() *ComputeGroupBackend {
	this := ComputeGroupBackend{}
	return &this
}

// NewComputeGroupBackendWithDefaults instantiates a new ComputeGroupBackend object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComputeGroupBackendWithDefaults() *ComputeGroupBackend {
	this := ComputeGroupBackend{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComputeGroupBackend) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupBackend) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComputeGroupBackend) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComputeGroupBackend) SetId(v string) {
	o.Id = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ComputeGroupBackend) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupBackend) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ComputeGroupBackend) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ComputeGroupBackend) SetHost(v string) {
	o.Host = &v
}

// GetHeartbeatPort returns the HeartbeatPort field value if set, zero value otherwise.
func (o *ComputeGroupBackend) GetHeartbeatPort() int64 {
	if o == nil || o.HeartbeatPort == nil {
		var ret int64
		return ret
	}
	return *o.HeartbeatPort
}

// GetHeartbeatPortOk returns a tuple with the HeartbeatPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupBackend) GetHeartbeatPortOk() (*int64, bool) {
	if o == nil || o.HeartbeatPort == nil {
		return nil, false
	}
	return o.HeartbeatPort, true
}

// HasHeartbeatPort returns a boolean if a field has been set.
func (o *ComputeGroupBackend) HasHeartbeatPort() bool {
	return o != nil && o.HeartbeatPort != nil
}

// SetHeartbeatPort gets a reference to the given int64 and assigns it to the HeartbeatPort field.
func (o *ComputeGroupBackend) SetHeartbeatPort(v int64) {
	o.HeartbeatPort = &v
}

// GetAlive returns the Alive field value if set, zero value otherwise.
func (o *ComputeGroupBackend) GetAlive() bool {
	if o == nil || o.Alive == nil {
		var ret bool
		return ret
	}
	return *o.Alive
}

// GetAliveOk returns a tuple with the Alive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupBackend) GetAliveOk() (*bool, bool) {
	if o == nil || o.Alive == nil {
		return nil, false
	}
	return o.Alive, true
}

// HasAlive returns a boolean if a field has been set.
func (o *ComputeGroupBackend) HasAlive() bool {
	return o != nil && o.Alive != nil
}

// SetAlive gets a reference to the given bool and assigns it to the Alive field.
func (o *ComputeGroupBackend) SetAlive(v bool) {
	o.Alive = &v
}

// GetDecommissioned returns the Decommissioned field value if set, zero value otherwise.
func (o *ComputeGroupBackend) GetDecommissioned() bool {
	if o == nil || o.Decommissioned == nil {
		var ret bool
		return ret
	}
	return *o.Decommissioned
}

// GetDecommissionedOk returns a tuple with the Decommissioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupBackend) GetDecommissionedOk() (*bool, bool) {
	if o == nil || o.Decommissioned == nil {
		return nil, false
	}
	return o.Decommissioned, true
}

// HasDecommissioned returns a boolean if a field has been set.
func (o *ComputeGroupBackend) HasDecommissioned() bool {
	return o != nil && o.Decommissioned != nil
}

// SetDecommissioned gets a reference to the given bool and assigns it to the Decommissioned field.
func (o *ComputeGroupBackend) SetDecommissioned(v bool) {
	o.Decommissioned = &v
}

// GetTabletCount returns the TabletCount field value if set, zero value otherwise.
func (o *ComputeGroupBackend) GetTabletCount() int64 {
	if o == nil || o.TabletCount == nil {
		var ret int64
		return ret
	}
	return *o.TabletCount
}

// GetTabletCountOk returns a tuple with the TabletCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupBackend) GetTabletCountOk() (*int64, bool) {
	if o == nil || o.TabletCount == nil {
		return nil, false
	}
	return o.TabletCount, true
}

// HasTabletCount returns a boolean if a field has been set.
func (o *ComputeGroupBackend) HasTabletCount() bool {
	return o != nil && o.TabletCount != nil
}

// SetTabletCount gets a reference to the given int64 and assigns it to the TabletCount field.
func (o *ComputeGroupBackend) SetTabletCount(v int64) {
	o.TabletCount = &v
}

// GetLastHeartbeat returns the LastHeartbeat field value if set, zero value otherwise.
func (o *ComputeGroupBackend) GetLastHeartbeat() string {
	if o == nil || o.LastHeartbeat == nil {
		var ret string
		return ret
	}
	return *o.LastHeartbeat
}

// GetLastHeartbeatOk returns a tuple with the LastHeartbeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupBackend) GetLastHeartbeatOk() (*string, bool) {
	if o == nil || o.LastHeartbeat == nil {
		return nil, false
	}
	return o.LastHeartbeat, true
}

// HasLastHeartbeat returns a boolean if a field has been set.
func (o *ComputeGroupBackend) HasLastHeartbeat() bool {
	return o != nil && o.LastHeartbeat != nil
}

// SetLastHeartbeat gets a reference to the given string and assigns it to the LastHeartbeat field.
func (o *ComputeGroupBackend) SetLastHeartbeat(v string) {
	o.LastHeartbeat = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ComputeGroupBackend) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupBackend) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ComputeGroupBackend) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ComputeGroupBackend) SetVersion(v string) {
	o.Version = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ComputeGroupBackend) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupBackend) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ComputeGroupBackend) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ComputeGroupBackend) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComputeGroupBackend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.HeartbeatPort != nil {
		toSerialize["heartbeatPort"] = o.HeartbeatPort
	}
	if o.Alive != nil {
		toSerialize["alive"] = o.Alive
	}
	if o.Decommissioned != nil {
		toSerialize["decommissioned"] = o.Decommissioned
	}
	if o.TabletCount != nil {
		toSerialize["tabletCount"] = o.TabletCount
	}
	if o.LastHeartbeat != nil {
		toSerialize["lastHeartbeat"] = o.LastHeartbeat
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComputeGroupBackend) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id             *string `json:"id,omitempty"`
		Host           *string `json:"host,omitempty"`
		HeartbeatPort  *int64  `json:"heartbeatPort,omitempty"`
		Alive          *bool   `json:"alive,omitempty"`
		Decommissioned *bool   `json:"decommissioned,omitempty"`
		TabletCount    *int64  `json:"tabletCount,omitempty"`
		LastHeartbeat  *string `json:"lastHeartbeat,omitempty"`
		Version        *string `json:"version,omitempty"`
		ErrorMessage   *string `json:"errorMessage,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "host", "heartbeatPort", "alive", "decommissioned", "tabletCount", "lastHeartbeat", "version", "errorMessage"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Host = all.Host
	o.HeartbeatPort = all.HeartbeatPort
	o.Alive = all.Alive
	o.Decommissioned = all.Decommissioned
	o.TabletCount = all.TabletCount
	o.LastHeartbeat = all.LastHeartbeat
	o.Version = all.Version
	o.ErrorMessage = all.ErrorMessage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
