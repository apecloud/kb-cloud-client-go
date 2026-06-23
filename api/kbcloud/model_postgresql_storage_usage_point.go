// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlStorageUsagePoint struct {
	Timestamp  string   `json:"timestamp"`
	UsedBytes  *int64   `json:"usedBytes,omitempty"`
	TotalBytes *int64   `json:"totalBytes,omitempty"`
	UsageRatio *float64 `json:"usageRatio,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlStorageUsagePoint instantiates a new PostgresqlStorageUsagePoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlStorageUsagePoint(timestamp string) *PostgresqlStorageUsagePoint {
	this := PostgresqlStorageUsagePoint{}
	this.Timestamp = timestamp
	return &this
}

// NewPostgresqlStorageUsagePointWithDefaults instantiates a new PostgresqlStorageUsagePoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlStorageUsagePointWithDefaults() *PostgresqlStorageUsagePoint {
	this := PostgresqlStorageUsagePoint{}
	return &this
}

// GetTimestamp returns the Timestamp field value.
func (o *PostgresqlStorageUsagePoint) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageUsagePoint) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *PostgresqlStorageUsagePoint) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetUsedBytes returns the UsedBytes field value if set, zero value otherwise.
func (o *PostgresqlStorageUsagePoint) GetUsedBytes() int64 {
	if o == nil || o.UsedBytes == nil {
		var ret int64
		return ret
	}
	return *o.UsedBytes
}

// GetUsedBytesOk returns a tuple with the UsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageUsagePoint) GetUsedBytesOk() (*int64, bool) {
	if o == nil || o.UsedBytes == nil {
		return nil, false
	}
	return o.UsedBytes, true
}

// HasUsedBytes returns a boolean if a field has been set.
func (o *PostgresqlStorageUsagePoint) HasUsedBytes() bool {
	return o != nil && o.UsedBytes != nil
}

// SetUsedBytes gets a reference to the given int64 and assigns it to the UsedBytes field.
func (o *PostgresqlStorageUsagePoint) SetUsedBytes(v int64) {
	o.UsedBytes = &v
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *PostgresqlStorageUsagePoint) GetTotalBytes() int64 {
	if o == nil || o.TotalBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageUsagePoint) GetTotalBytesOk() (*int64, bool) {
	if o == nil || o.TotalBytes == nil {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *PostgresqlStorageUsagePoint) HasTotalBytes() bool {
	return o != nil && o.TotalBytes != nil
}

// SetTotalBytes gets a reference to the given int64 and assigns it to the TotalBytes field.
func (o *PostgresqlStorageUsagePoint) SetTotalBytes(v int64) {
	o.TotalBytes = &v
}

// GetUsageRatio returns the UsageRatio field value if set, zero value otherwise.
func (o *PostgresqlStorageUsagePoint) GetUsageRatio() float64 {
	if o == nil || o.UsageRatio == nil {
		var ret float64
		return ret
	}
	return *o.UsageRatio
}

// GetUsageRatioOk returns a tuple with the UsageRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageUsagePoint) GetUsageRatioOk() (*float64, bool) {
	if o == nil || o.UsageRatio == nil {
		return nil, false
	}
	return o.UsageRatio, true
}

// HasUsageRatio returns a boolean if a field has been set.
func (o *PostgresqlStorageUsagePoint) HasUsageRatio() bool {
	return o != nil && o.UsageRatio != nil
}

// SetUsageRatio gets a reference to the given float64 and assigns it to the UsageRatio field.
func (o *PostgresqlStorageUsagePoint) SetUsageRatio(v float64) {
	o.UsageRatio = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlStorageUsagePoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["timestamp"] = o.Timestamp
	if o.UsedBytes != nil {
		toSerialize["usedBytes"] = o.UsedBytes
	}
	if o.TotalBytes != nil {
		toSerialize["totalBytes"] = o.TotalBytes
	}
	if o.UsageRatio != nil {
		toSerialize["usageRatio"] = o.UsageRatio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlStorageUsagePoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Timestamp  *string  `json:"timestamp"`
		UsedBytes  *int64   `json:"usedBytes,omitempty"`
		TotalBytes *int64   `json:"totalBytes,omitempty"`
		UsageRatio *float64 `json:"usageRatio,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"timestamp", "usedBytes", "totalBytes", "usageRatio"})
	} else {
		return err
	}
	o.Timestamp = *all.Timestamp
	o.UsedBytes = all.UsedBytes
	o.TotalBytes = all.TotalBytes
	o.UsageRatio = all.UsageRatio

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
