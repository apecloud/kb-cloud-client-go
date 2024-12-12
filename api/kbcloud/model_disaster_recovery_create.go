// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

// DisasterRecoveryCreate DisasterRecoveryCreate is the description to create a DisasterRecovery cluster
type DisasterRecoveryCreate struct {
	// KubeBlocks cluster information
	Cluster Cluster `json:"cluster"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDisasterRecoveryCreate instantiates a new DisasterRecoveryCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDisasterRecoveryCreate(cluster Cluster) *DisasterRecoveryCreate {
	this := DisasterRecoveryCreate{}
	this.Cluster = cluster
	return &this
}

// NewDisasterRecoveryCreateWithDefaults instantiates a new DisasterRecoveryCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDisasterRecoveryCreateWithDefaults() *DisasterRecoveryCreate {
	this := DisasterRecoveryCreate{}
	return &this
}

// GetCluster returns the Cluster field value.
func (o *DisasterRecoveryCreate) GetCluster() Cluster {
	if o == nil {
		var ret Cluster
		return ret
	}
	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryCreate) GetClusterOk() (*Cluster, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value.
func (o *DisasterRecoveryCreate) SetCluster(v Cluster) {
	o.Cluster = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DisasterRecoveryCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cluster"] = o.Cluster

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DisasterRecoveryCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cluster *Cluster `json:"cluster"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Cluster == nil {
		return fmt.Errorf("required field cluster missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cluster"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Cluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cluster = *all.Cluster

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
