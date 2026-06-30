// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterStorageUsageHistoryPoint struct {
	Timestamp      string  `json:"timestamp"`
	UsedBytes      int64   `json:"usedBytes"`
	TotalBytes     int64   `json:"totalBytes"`
	AvailableBytes int64   `json:"availableBytes"`
	UsageRatio     float64 `json:"usageRatio"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterStorageUsageHistoryPoint instantiates a new ClusterStorageUsageHistoryPoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterStorageUsageHistoryPoint(timestamp string, usedBytes int64, totalBytes int64, availableBytes int64, usageRatio float64) *ClusterStorageUsageHistoryPoint {
	this := ClusterStorageUsageHistoryPoint{}
	this.Timestamp = timestamp
	this.UsedBytes = usedBytes
	this.TotalBytes = totalBytes
	this.AvailableBytes = availableBytes
	this.UsageRatio = usageRatio
	return &this
}

// NewClusterStorageUsageHistoryPointWithDefaults instantiates a new ClusterStorageUsageHistoryPoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterStorageUsageHistoryPointWithDefaults() *ClusterStorageUsageHistoryPoint {
	this := ClusterStorageUsageHistoryPoint{}
	return &this
}

// GetTimestamp returns the Timestamp field value.
func (o *ClusterStorageUsageHistoryPoint) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryPoint) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *ClusterStorageUsageHistoryPoint) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetUsedBytes returns the UsedBytes field value.
func (o *ClusterStorageUsageHistoryPoint) GetUsedBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UsedBytes
}

// GetUsedBytesOk returns a tuple with the UsedBytes field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryPoint) GetUsedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsedBytes, true
}

// SetUsedBytes sets field value.
func (o *ClusterStorageUsageHistoryPoint) SetUsedBytes(v int64) {
	o.UsedBytes = v
}

// GetTotalBytes returns the TotalBytes field value.
func (o *ClusterStorageUsageHistoryPoint) GetTotalBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryPoint) GetTotalBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalBytes, true
}

// SetTotalBytes sets field value.
func (o *ClusterStorageUsageHistoryPoint) SetTotalBytes(v int64) {
	o.TotalBytes = v
}

// GetAvailableBytes returns the AvailableBytes field value.
func (o *ClusterStorageUsageHistoryPoint) GetAvailableBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvailableBytes
}

// GetAvailableBytesOk returns a tuple with the AvailableBytes field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryPoint) GetAvailableBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableBytes, true
}

// SetAvailableBytes sets field value.
func (o *ClusterStorageUsageHistoryPoint) SetAvailableBytes(v int64) {
	o.AvailableBytes = v
}

// GetUsageRatio returns the UsageRatio field value.
func (o *ClusterStorageUsageHistoryPoint) GetUsageRatio() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.UsageRatio
}

// GetUsageRatioOk returns a tuple with the UsageRatio field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistoryPoint) GetUsageRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageRatio, true
}

// SetUsageRatio sets field value.
func (o *ClusterStorageUsageHistoryPoint) SetUsageRatio(v float64) {
	o.UsageRatio = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterStorageUsageHistoryPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["usedBytes"] = o.UsedBytes
	toSerialize["totalBytes"] = o.TotalBytes
	toSerialize["availableBytes"] = o.AvailableBytes
	toSerialize["usageRatio"] = o.UsageRatio

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterStorageUsageHistoryPoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Timestamp      *string  `json:"timestamp"`
		UsedBytes      *int64   `json:"usedBytes"`
		TotalBytes     *int64   `json:"totalBytes"`
		AvailableBytes *int64   `json:"availableBytes"`
		UsageRatio     *float64 `json:"usageRatio"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	if all.UsedBytes == nil {
		return fmt.Errorf("required field usedBytes missing")
	}
	if all.TotalBytes == nil {
		return fmt.Errorf("required field totalBytes missing")
	}
	if all.AvailableBytes == nil {
		return fmt.Errorf("required field availableBytes missing")
	}
	if all.UsageRatio == nil {
		return fmt.Errorf("required field usageRatio missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"timestamp", "usedBytes", "totalBytes", "availableBytes", "usageRatio"})
	} else {
		return err
	}
	o.Timestamp = *all.Timestamp
	o.UsedBytes = *all.UsedBytes
	o.TotalBytes = *all.TotalBytes
	o.AvailableBytes = *all.AvailableBytes
	o.UsageRatio = *all.UsageRatio

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
