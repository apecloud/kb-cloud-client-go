// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MssqlDatabaseSpace struct {
	Name       string `json:"name"`
	State      string `json:"state"`
	Accessible bool   `json:"accessible"`
	DataBytes  int64  `json:"dataBytes"`
	LogBytes   int64  `json:"logBytes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlDatabaseSpace instantiates a new MssqlDatabaseSpace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlDatabaseSpace(name string, state string, accessible bool, dataBytes int64, logBytes int64) *MssqlDatabaseSpace {
	this := MssqlDatabaseSpace{}
	this.Name = name
	this.State = state
	this.Accessible = accessible
	this.DataBytes = dataBytes
	this.LogBytes = logBytes
	return &this
}

// NewMssqlDatabaseSpaceWithDefaults instantiates a new MssqlDatabaseSpace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlDatabaseSpaceWithDefaults() *MssqlDatabaseSpace {
	this := MssqlDatabaseSpace{}
	return &this
}

// GetName returns the Name field value.
func (o *MssqlDatabaseSpace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MssqlDatabaseSpace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MssqlDatabaseSpace) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value.
func (o *MssqlDatabaseSpace) GetState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *MssqlDatabaseSpace) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *MssqlDatabaseSpace) SetState(v string) {
	o.State = v
}

// GetAccessible returns the Accessible field value.
func (o *MssqlDatabaseSpace) GetAccessible() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Accessible
}

// GetAccessibleOk returns a tuple with the Accessible field value
// and a boolean to check if the value has been set.
func (o *MssqlDatabaseSpace) GetAccessibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accessible, true
}

// SetAccessible sets field value.
func (o *MssqlDatabaseSpace) SetAccessible(v bool) {
	o.Accessible = v
}

// GetDataBytes returns the DataBytes field value.
func (o *MssqlDatabaseSpace) GetDataBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DataBytes
}

// GetDataBytesOk returns a tuple with the DataBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlDatabaseSpace) GetDataBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataBytes, true
}

// SetDataBytes sets field value.
func (o *MssqlDatabaseSpace) SetDataBytes(v int64) {
	o.DataBytes = v
}

// GetLogBytes returns the LogBytes field value.
func (o *MssqlDatabaseSpace) GetLogBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LogBytes
}

// GetLogBytesOk returns a tuple with the LogBytes field value
// and a boolean to check if the value has been set.
func (o *MssqlDatabaseSpace) GetLogBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogBytes, true
}

// SetLogBytes sets field value.
func (o *MssqlDatabaseSpace) SetLogBytes(v int64) {
	o.LogBytes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlDatabaseSpace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["accessible"] = o.Accessible
	toSerialize["dataBytes"] = o.DataBytes
	toSerialize["logBytes"] = o.LogBytes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlDatabaseSpace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string `json:"name"`
		State      *string `json:"state"`
		Accessible *bool   `json:"accessible"`
		DataBytes  *int64  `json:"dataBytes"`
		LogBytes   *int64  `json:"logBytes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.Accessible == nil {
		return fmt.Errorf("required field accessible missing")
	}
	if all.DataBytes == nil {
		return fmt.Errorf("required field dataBytes missing")
	}
	if all.LogBytes == nil {
		return fmt.Errorf("required field logBytes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "state", "accessible", "dataBytes", "logBytes"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.State = *all.State
	o.Accessible = *all.Accessible
	o.DataBytes = *all.DataBytes
	o.LogBytes = *all.LogBytes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
