// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package data

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ExpandPartitionRequest struct {
	NewPartitions int32 `json:"newPartitions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExpandPartitionRequest instantiates a new ExpandPartitionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExpandPartitionRequest(newPartitions int32) *ExpandPartitionRequest {
	this := ExpandPartitionRequest{}
	this.NewPartitions = newPartitions
	return &this
}

// NewExpandPartitionRequestWithDefaults instantiates a new ExpandPartitionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExpandPartitionRequestWithDefaults() *ExpandPartitionRequest {
	this := ExpandPartitionRequest{}
	return &this
}

// GetNewPartitions returns the NewPartitions field value.
func (o *ExpandPartitionRequest) GetNewPartitions() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.NewPartitions
}

// GetNewPartitionsOk returns a tuple with the NewPartitions field value
// and a boolean to check if the value has been set.
func (o *ExpandPartitionRequest) GetNewPartitionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPartitions, true
}

// SetNewPartitions sets field value.
func (o *ExpandPartitionRequest) SetNewPartitions(v int32) {
	o.NewPartitions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExpandPartitionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["newPartitions"] = o.NewPartitions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExpandPartitionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NewPartitions *int32 `json:"newPartitions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.NewPartitions == nil {
		return fmt.Errorf("required field newPartitions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"newPartitions"})
	} else {
		return err
	}
	o.NewPartitions = *all.NewPartitions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
