// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisSlowLogEntry struct {
	Id             *int64  `json:"id,omitempty"`
	Timestamp      *int64  `json:"timestamp,omitempty"`
	DurationMicros *int64  `json:"durationMicros,omitempty"`
	Command        *string `json:"command,omitempty"`
	ClientAddr     *string `json:"clientAddr,omitempty"`
	ClientName     *string `json:"clientName,omitempty"`
	NodeId         *string `json:"nodeId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisSlowLogEntry instantiates a new RedisSlowLogEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisSlowLogEntry() *RedisSlowLogEntry {
	this := RedisSlowLogEntry{}
	return &this
}

// NewRedisSlowLogEntryWithDefaults instantiates a new RedisSlowLogEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisSlowLogEntryWithDefaults() *RedisSlowLogEntry {
	this := RedisSlowLogEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RedisSlowLogEntry) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisSlowLogEntry) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RedisSlowLogEntry) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *RedisSlowLogEntry) SetId(v int64) {
	o.Id = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *RedisSlowLogEntry) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisSlowLogEntry) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *RedisSlowLogEntry) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *RedisSlowLogEntry) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetDurationMicros returns the DurationMicros field value if set, zero value otherwise.
func (o *RedisSlowLogEntry) GetDurationMicros() int64 {
	if o == nil || o.DurationMicros == nil {
		var ret int64
		return ret
	}
	return *o.DurationMicros
}

// GetDurationMicrosOk returns a tuple with the DurationMicros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisSlowLogEntry) GetDurationMicrosOk() (*int64, bool) {
	if o == nil || o.DurationMicros == nil {
		return nil, false
	}
	return o.DurationMicros, true
}

// HasDurationMicros returns a boolean if a field has been set.
func (o *RedisSlowLogEntry) HasDurationMicros() bool {
	return o != nil && o.DurationMicros != nil
}

// SetDurationMicros gets a reference to the given int64 and assigns it to the DurationMicros field.
func (o *RedisSlowLogEntry) SetDurationMicros(v int64) {
	o.DurationMicros = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *RedisSlowLogEntry) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisSlowLogEntry) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *RedisSlowLogEntry) HasCommand() bool {
	return o != nil && o.Command != nil
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *RedisSlowLogEntry) SetCommand(v string) {
	o.Command = &v
}

// GetClientAddr returns the ClientAddr field value if set, zero value otherwise.
func (o *RedisSlowLogEntry) GetClientAddr() string {
	if o == nil || o.ClientAddr == nil {
		var ret string
		return ret
	}
	return *o.ClientAddr
}

// GetClientAddrOk returns a tuple with the ClientAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisSlowLogEntry) GetClientAddrOk() (*string, bool) {
	if o == nil || o.ClientAddr == nil {
		return nil, false
	}
	return o.ClientAddr, true
}

// HasClientAddr returns a boolean if a field has been set.
func (o *RedisSlowLogEntry) HasClientAddr() bool {
	return o != nil && o.ClientAddr != nil
}

// SetClientAddr gets a reference to the given string and assigns it to the ClientAddr field.
func (o *RedisSlowLogEntry) SetClientAddr(v string) {
	o.ClientAddr = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *RedisSlowLogEntry) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisSlowLogEntry) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *RedisSlowLogEntry) HasClientName() bool {
	return o != nil && o.ClientName != nil
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *RedisSlowLogEntry) SetClientName(v string) {
	o.ClientName = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *RedisSlowLogEntry) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisSlowLogEntry) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *RedisSlowLogEntry) HasNodeId() bool {
	return o != nil && o.NodeId != nil
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *RedisSlowLogEntry) SetNodeId(v string) {
	o.NodeId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisSlowLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.DurationMicros != nil {
		toSerialize["durationMicros"] = o.DurationMicros
	}
	if o.Command != nil {
		toSerialize["command"] = o.Command
	}
	if o.ClientAddr != nil {
		toSerialize["clientAddr"] = o.ClientAddr
	}
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisSlowLogEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id             *int64  `json:"id,omitempty"`
		Timestamp      *int64  `json:"timestamp,omitempty"`
		DurationMicros *int64  `json:"durationMicros,omitempty"`
		Command        *string `json:"command,omitempty"`
		ClientAddr     *string `json:"clientAddr,omitempty"`
		ClientName     *string `json:"clientName,omitempty"`
		NodeId         *string `json:"nodeId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "timestamp", "durationMicros", "command", "clientAddr", "clientName", "nodeId"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Timestamp = all.Timestamp
	o.DurationMicros = all.DurationMicros
	o.Command = all.Command
	o.ClientAddr = all.ClientAddr
	o.ClientName = all.ClientName
	o.NodeId = all.NodeId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
