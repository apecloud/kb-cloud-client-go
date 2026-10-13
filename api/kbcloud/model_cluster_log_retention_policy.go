// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterLogRetentionPolicy Effective cluster log retention and its cluster-level overrides.
type ClusterLogRetentionPolicy struct {
	// Environment log retention days.
	Effective EnvironmentLogRetentionConfig `json:"effective"`
	// Cluster log retention overrides. Omitted fields inherit from the environment.
	Overrides ClusterLogRetentionConfig `json:"overrides"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterLogRetentionPolicy instantiates a new ClusterLogRetentionPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterLogRetentionPolicy(effective EnvironmentLogRetentionConfig, overrides ClusterLogRetentionConfig) *ClusterLogRetentionPolicy {
	this := ClusterLogRetentionPolicy{}
	this.Effective = effective
	this.Overrides = overrides
	return &this
}

// NewClusterLogRetentionPolicyWithDefaults instantiates a new ClusterLogRetentionPolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterLogRetentionPolicyWithDefaults() *ClusterLogRetentionPolicy {
	this := ClusterLogRetentionPolicy{}
	return &this
}

// GetEffective returns the Effective field value.
func (o *ClusterLogRetentionPolicy) GetEffective() EnvironmentLogRetentionConfig {
	if o == nil {
		var ret EnvironmentLogRetentionConfig
		return ret
	}
	return o.Effective
}

// GetEffectiveOk returns a tuple with the Effective field value
// and a boolean to check if the value has been set.
func (o *ClusterLogRetentionPolicy) GetEffectiveOk() (*EnvironmentLogRetentionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Effective, true
}

// SetEffective sets field value.
func (o *ClusterLogRetentionPolicy) SetEffective(v EnvironmentLogRetentionConfig) {
	o.Effective = v
}

// GetOverrides returns the Overrides field value.
func (o *ClusterLogRetentionPolicy) GetOverrides() ClusterLogRetentionConfig {
	if o == nil {
		var ret ClusterLogRetentionConfig
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value
// and a boolean to check if the value has been set.
func (o *ClusterLogRetentionPolicy) GetOverridesOk() (*ClusterLogRetentionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overrides, true
}

// SetOverrides sets field value.
func (o *ClusterLogRetentionPolicy) SetOverrides(v ClusterLogRetentionConfig) {
	o.Overrides = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterLogRetentionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["effective"] = o.Effective
	toSerialize["overrides"] = o.Overrides

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterLogRetentionPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Effective *EnvironmentLogRetentionConfig `json:"effective"`
		Overrides *ClusterLogRetentionConfig     `json:"overrides"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Effective == nil {
		return fmt.Errorf("required field effective missing")
	}
	if all.Overrides == nil {
		return fmt.Errorf("required field overrides missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"effective", "overrides"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Effective.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Effective = *all.Effective
	if all.Overrides.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Overrides = *all.Overrides

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
