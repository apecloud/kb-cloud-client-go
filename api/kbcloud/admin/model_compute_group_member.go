// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ComputeGroupMember struct {
	// Engine-native member identifier.
	Id    *string `json:"id,omitempty"`
	Host  *string `json:"host,omitempty"`
	Port  *int64  `json:"port,omitempty"`
	Alive *bool   `json:"alive,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComputeGroupMember instantiates a new ComputeGroupMember object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComputeGroupMember() *ComputeGroupMember {
	this := ComputeGroupMember{}
	return &this
}

// NewComputeGroupMemberWithDefaults instantiates a new ComputeGroupMember object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComputeGroupMemberWithDefaults() *ComputeGroupMember {
	this := ComputeGroupMember{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComputeGroupMember) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupMember) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComputeGroupMember) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComputeGroupMember) SetId(v string) {
	o.Id = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ComputeGroupMember) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupMember) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ComputeGroupMember) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ComputeGroupMember) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ComputeGroupMember) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupMember) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ComputeGroupMember) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *ComputeGroupMember) SetPort(v int64) {
	o.Port = &v
}

// GetAlive returns the Alive field value if set, zero value otherwise.
func (o *ComputeGroupMember) GetAlive() bool {
	if o == nil || o.Alive == nil {
		var ret bool
		return ret
	}
	return *o.Alive
}

// GetAliveOk returns a tuple with the Alive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeGroupMember) GetAliveOk() (*bool, bool) {
	if o == nil || o.Alive == nil {
		return nil, false
	}
	return o.Alive, true
}

// HasAlive returns a boolean if a field has been set.
func (o *ComputeGroupMember) HasAlive() bool {
	return o != nil && o.Alive != nil
}

// SetAlive gets a reference to the given bool and assigns it to the Alive field.
func (o *ComputeGroupMember) SetAlive(v bool) {
	o.Alive = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComputeGroupMember) MarshalJSON() ([]byte, error) {
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
	if o.Alive != nil {
		toSerialize["alive"] = o.Alive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComputeGroupMember) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id    *string `json:"id,omitempty"`
		Host  *string `json:"host,omitempty"`
		Port  *int64  `json:"port,omitempty"`
		Alive *bool   `json:"alive,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "host", "port", "alive"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Host = all.Host
	o.Port = all.Port
	o.Alive = all.Alive

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
