// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type BrokerNode struct {
	Id      *int32  `json:"id,omitempty"`
	Host    *string `json:"host,omitempty"`
	Port    *int32  `json:"port,omitempty"`
	LogSize *int64  `json:"logSize,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBrokerNode instantiates a new BrokerNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBrokerNode() *BrokerNode {
	this := BrokerNode{}
	return &this
}

// NewBrokerNodeWithDefaults instantiates a new BrokerNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBrokerNodeWithDefaults() *BrokerNode {
	this := BrokerNode{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrokerNode) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerNode) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrokerNode) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BrokerNode) SetId(v int32) {
	o.Id = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *BrokerNode) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerNode) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *BrokerNode) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *BrokerNode) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *BrokerNode) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerNode) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *BrokerNode) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *BrokerNode) SetPort(v int32) {
	o.Port = &v
}

// GetLogSize returns the LogSize field value if set, zero value otherwise.
func (o *BrokerNode) GetLogSize() int64 {
	if o == nil || o.LogSize == nil {
		var ret int64
		return ret
	}
	return *o.LogSize
}

// GetLogSizeOk returns a tuple with the LogSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerNode) GetLogSizeOk() (*int64, bool) {
	if o == nil || o.LogSize == nil {
		return nil, false
	}
	return o.LogSize, true
}

// HasLogSize returns a boolean if a field has been set.
func (o *BrokerNode) HasLogSize() bool {
	return o != nil && o.LogSize != nil
}

// SetLogSize gets a reference to the given int64 and assigns it to the LogSize field.
func (o *BrokerNode) SetLogSize(v int64) {
	o.LogSize = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BrokerNode) MarshalJSON() ([]byte, error) {
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
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.LogSize != nil {
		toSerialize["logSize"] = o.LogSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BrokerNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id      *int32  `json:"id,omitempty"`
		Host    *string `json:"host,omitempty"`
		Port    *int32  `json:"port,omitempty"`
		LogSize *int64  `json:"logSize,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "host", "port", "logSize"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Host = all.Host
	o.Port = all.Port
	o.LogSize = all.LogSize

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
