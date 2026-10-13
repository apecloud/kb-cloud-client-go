// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterLogRetentionConfig struct {
	// Retention days by log type. Omitted types inherit the environment policy.
	RetentionDaysByType map[string]int64 `json:"retentionDaysByType"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterLogRetentionConfig instantiates a new ClusterLogRetentionConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterLogRetentionConfig(retentionDaysByType map[string]int64) *ClusterLogRetentionConfig {
	this := ClusterLogRetentionConfig{}
	this.RetentionDaysByType = retentionDaysByType
	return &this
}

// NewClusterLogRetentionConfigWithDefaults instantiates a new ClusterLogRetentionConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterLogRetentionConfigWithDefaults() *ClusterLogRetentionConfig {
	this := ClusterLogRetentionConfig{}
	return &this
}

// GetRetentionDaysByType returns the RetentionDaysByType field value.
func (o *ClusterLogRetentionConfig) GetRetentionDaysByType() map[string]int64 {
	if o == nil {
		var ret map[string]int64
		return ret
	}
	return o.RetentionDaysByType
}

// GetRetentionDaysByTypeOk returns a tuple with the RetentionDaysByType field value
// and a boolean to check if the value has been set.
func (o *ClusterLogRetentionConfig) GetRetentionDaysByTypeOk() (*map[string]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionDaysByType, true
}

// SetRetentionDaysByType sets field value.
func (o *ClusterLogRetentionConfig) SetRetentionDaysByType(v map[string]int64) {
	o.RetentionDaysByType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterLogRetentionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["retentionDaysByType"] = o.RetentionDaysByType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterLogRetentionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RetentionDaysByType *map[string]int64 `json:"retentionDaysByType"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.RetentionDaysByType == nil {
		return fmt.Errorf("required field retentionDaysByType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"retentionDaysByType"})
	} else {
		return err
	}
	o.RetentionDaysByType = *all.RetentionDaysByType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
