// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterJarConfig struct {
	Supported     bool    `json:"supported"`
	Reason        *string `json:"reason,omitempty"`
	Generation    int64   `json:"generation"`
	MaxFileBytes  int64   `json:"maxFileBytes"`
	MaxTotalBytes int64   `json:"maxTotalBytes"`
	UsedBytes     int64   `json:"usedBytes"`
	TaskId        *string `json:"taskId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterJarConfig instantiates a new ClusterJarConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterJarConfig(supported bool, generation int64, maxFileBytes int64, maxTotalBytes int64, usedBytes int64) *ClusterJarConfig {
	this := ClusterJarConfig{}
	this.Supported = supported
	this.Generation = generation
	this.MaxFileBytes = maxFileBytes
	this.MaxTotalBytes = maxTotalBytes
	this.UsedBytes = usedBytes
	return &this
}

// NewClusterJarConfigWithDefaults instantiates a new ClusterJarConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterJarConfigWithDefaults() *ClusterJarConfig {
	this := ClusterJarConfig{}
	return &this
}

// GetSupported returns the Supported field value.
func (o *ClusterJarConfig) GetSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *ClusterJarConfig) GetSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supported, true
}

// SetSupported sets field value.
func (o *ClusterJarConfig) SetSupported(v bool) {
	o.Supported = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ClusterJarConfig) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterJarConfig) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ClusterJarConfig) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ClusterJarConfig) SetReason(v string) {
	o.Reason = &v
}

// GetGeneration returns the Generation field value.
func (o *ClusterJarConfig) GetGeneration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value
// and a boolean to check if the value has been set.
func (o *ClusterJarConfig) GetGenerationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Generation, true
}

// SetGeneration sets field value.
func (o *ClusterJarConfig) SetGeneration(v int64) {
	o.Generation = v
}

// GetMaxFileBytes returns the MaxFileBytes field value.
func (o *ClusterJarConfig) GetMaxFileBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxFileBytes
}

// GetMaxFileBytesOk returns a tuple with the MaxFileBytes field value
// and a boolean to check if the value has been set.
func (o *ClusterJarConfig) GetMaxFileBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxFileBytes, true
}

// SetMaxFileBytes sets field value.
func (o *ClusterJarConfig) SetMaxFileBytes(v int64) {
	o.MaxFileBytes = v
}

// GetMaxTotalBytes returns the MaxTotalBytes field value.
func (o *ClusterJarConfig) GetMaxTotalBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxTotalBytes
}

// GetMaxTotalBytesOk returns a tuple with the MaxTotalBytes field value
// and a boolean to check if the value has been set.
func (o *ClusterJarConfig) GetMaxTotalBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxTotalBytes, true
}

// SetMaxTotalBytes sets field value.
func (o *ClusterJarConfig) SetMaxTotalBytes(v int64) {
	o.MaxTotalBytes = v
}

// GetUsedBytes returns the UsedBytes field value.
func (o *ClusterJarConfig) GetUsedBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UsedBytes
}

// GetUsedBytesOk returns a tuple with the UsedBytes field value
// and a boolean to check if the value has been set.
func (o *ClusterJarConfig) GetUsedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedBytes, true
}

// SetUsedBytes sets field value.
func (o *ClusterJarConfig) SetUsedBytes(v int64) {
	o.UsedBytes = v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *ClusterJarConfig) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterJarConfig) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *ClusterJarConfig) HasTaskId() bool {
	return o != nil && o.TaskId != nil
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *ClusterJarConfig) SetTaskId(v string) {
	o.TaskId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterJarConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["supported"] = o.Supported
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["generation"] = o.Generation
	toSerialize["maxFileBytes"] = o.MaxFileBytes
	toSerialize["maxTotalBytes"] = o.MaxTotalBytes
	toSerialize["usedBytes"] = o.UsedBytes
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterJarConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Supported     *bool   `json:"supported"`
		Reason        *string `json:"reason,omitempty"`
		Generation    *int64  `json:"generation"`
		MaxFileBytes  *int64  `json:"maxFileBytes"`
		MaxTotalBytes *int64  `json:"maxTotalBytes"`
		UsedBytes     *int64  `json:"usedBytes"`
		TaskId        *string `json:"taskId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Supported == nil {
		return fmt.Errorf("required field supported missing")
	}
	if all.Generation == nil {
		return fmt.Errorf("required field generation missing")
	}
	if all.MaxFileBytes == nil {
		return fmt.Errorf("required field maxFileBytes missing")
	}
	if all.MaxTotalBytes == nil {
		return fmt.Errorf("required field maxTotalBytes missing")
	}
	if all.UsedBytes == nil {
		return fmt.Errorf("required field usedBytes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"supported", "reason", "generation", "maxFileBytes", "maxTotalBytes", "usedBytes", "taskId"})
	} else {
		return err
	}
	o.Supported = *all.Supported
	o.Reason = all.Reason
	o.Generation = *all.Generation
	o.MaxFileBytes = *all.MaxFileBytes
	o.MaxTotalBytes = *all.MaxTotalBytes
	o.UsedBytes = *all.UsedBytes
	o.TaskId = all.TaskId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
