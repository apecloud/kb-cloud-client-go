// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type CreateTopicRequest struct {
	Name              string `json:"name"`
	Partitions        int32  `json:"partitions"`
	ReplicationFactor int32  `json:"replicationFactor"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateTopicRequest instantiates a new CreateTopicRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateTopicRequest(name string, partitions int32, replicationFactor int32) *CreateTopicRequest {
	this := CreateTopicRequest{}
	this.Name = name
	this.Partitions = partitions
	this.ReplicationFactor = replicationFactor
	return &this
}

// NewCreateTopicRequestWithDefaults instantiates a new CreateTopicRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateTopicRequestWithDefaults() *CreateTopicRequest {
	this := CreateTopicRequest{}
	return &this
}

// GetName returns the Name field value.
func (o *CreateTopicRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTopicRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateTopicRequest) SetName(v string) {
	o.Name = v
}

// GetPartitions returns the Partitions field value.
func (o *CreateTopicRequest) GetPartitions() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value
// and a boolean to check if the value has been set.
func (o *CreateTopicRequest) GetPartitionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Partitions, true
}

// SetPartitions sets field value.
func (o *CreateTopicRequest) SetPartitions(v int32) {
	o.Partitions = v
}

// GetReplicationFactor returns the ReplicationFactor field value.
func (o *CreateTopicRequest) GetReplicationFactor() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ReplicationFactor
}

// GetReplicationFactorOk returns a tuple with the ReplicationFactor field value
// and a boolean to check if the value has been set.
func (o *CreateTopicRequest) GetReplicationFactorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplicationFactor, true
}

// SetReplicationFactor sets field value.
func (o *CreateTopicRequest) SetReplicationFactor(v int32) {
	o.ReplicationFactor = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateTopicRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["partitions"] = o.Partitions
	toSerialize["replicationFactor"] = o.ReplicationFactor

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateTopicRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string `json:"name"`
		Partitions        *int32  `json:"partitions"`
		ReplicationFactor *int32  `json:"replicationFactor"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Partitions == nil {
		return fmt.Errorf("required field partitions missing")
	}
	if all.ReplicationFactor == nil {
		return fmt.Errorf("required field replicationFactor missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "partitions", "replicationFactor"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Partitions = *all.Partitions
	o.ReplicationFactor = *all.ReplicationFactor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
