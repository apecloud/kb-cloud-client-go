// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package data

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type Broker struct {
	Id *int32 `json:"id,omitempty"`
	Host *string `json:"host,omitempty"`
	Port *int32 `json:"port,omitempty"`
	LeaderPartitions []int32 `json:"leaderPartitions,omitempty"`
	FollowerPartitions []int32 `json:"followerPartitions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewBroker instantiates a new Broker object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBroker() *Broker {
	this := Broker{}
	return &this
}

// NewBrokerWithDefaults instantiates a new Broker object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBrokerWithDefaults() *Broker {
	this := Broker{}
	return &this
}
// GetId returns the Id field value if set, zero value otherwise.
func (o *Broker) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Broker) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Broker) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Broker) SetId(v int32) {
	o.Id = &v
}


// GetHost returns the Host field value if set, zero value otherwise.
func (o *Broker) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Broker) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Broker) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Broker) SetHost(v string) {
	o.Host = &v
}


// GetPort returns the Port field value if set, zero value otherwise.
func (o *Broker) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Broker) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Broker) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *Broker) SetPort(v int32) {
	o.Port = &v
}


// GetLeaderPartitions returns the LeaderPartitions field value if set, zero value otherwise.
func (o *Broker) GetLeaderPartitions() []int32 {
	if o == nil || o.LeaderPartitions == nil {
		var ret []int32
		return ret
	}
	return o.LeaderPartitions
}

// GetLeaderPartitionsOk returns a tuple with the LeaderPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Broker) GetLeaderPartitionsOk() (*[]int32, bool) {
	if o == nil || o.LeaderPartitions == nil {
		return nil, false
	}
	return &o.LeaderPartitions, true
}

// HasLeaderPartitions returns a boolean if a field has been set.
func (o *Broker) HasLeaderPartitions() bool {
	return o != nil && o.LeaderPartitions != nil
}

// SetLeaderPartitions gets a reference to the given []int32 and assigns it to the LeaderPartitions field.
func (o *Broker) SetLeaderPartitions(v []int32) {
	o.LeaderPartitions = v
}


// GetFollowerPartitions returns the FollowerPartitions field value if set, zero value otherwise.
func (o *Broker) GetFollowerPartitions() []int32 {
	if o == nil || o.FollowerPartitions == nil {
		var ret []int32
		return ret
	}
	return o.FollowerPartitions
}

// GetFollowerPartitionsOk returns a tuple with the FollowerPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Broker) GetFollowerPartitionsOk() (*[]int32, bool) {
	if o == nil || o.FollowerPartitions == nil {
		return nil, false
	}
	return &o.FollowerPartitions, true
}

// HasFollowerPartitions returns a boolean if a field has been set.
func (o *Broker) HasFollowerPartitions() bool {
	return o != nil && o.FollowerPartitions != nil
}

// SetFollowerPartitions gets a reference to the given []int32 and assigns it to the FollowerPartitions field.
func (o *Broker) SetFollowerPartitions(v []int32) {
	o.FollowerPartitions = v
}



// MarshalJSON serializes the struct using spec logic.
func (o Broker) MarshalJSON() ([]byte, error) {
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
	if o.LeaderPartitions != nil {
		toSerialize["leaderPartitions"] = o.LeaderPartitions
	}
	if o.FollowerPartitions != nil {
		toSerialize["followerPartitions"] = o.FollowerPartitions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Broker) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id *int32 `json:"id,omitempty"`
		Host *string `json:"host,omitempty"`
		Port *int32 `json:"port,omitempty"`
		LeaderPartitions []int32 `json:"leaderPartitions,omitempty"`
		FollowerPartitions []int32 `json:"followerPartitions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "id", "host", "port", "leaderPartitions", "followerPartitions",  })
	} else {
		return err
	}
	o.Id = all.Id
	o.Host = all.Host
	o.Port = all.Port
	o.LeaderPartitions = all.LeaderPartitions
	o.FollowerPartitions = all.FollowerPartitions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
