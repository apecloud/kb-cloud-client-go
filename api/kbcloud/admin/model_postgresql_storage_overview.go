// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlStorageOverview struct {
	// Current total PVC capacity in bytes when backend-owned metrics are available.
	TotalBytes *int64 `json:"totalBytes,omitempty"`
	// Current PVC used bytes when backend-owned metrics are available.
	UsedBytes *int64 `json:"usedBytes,omitempty"`
	// Current available PVC bytes derived from totalBytes - usedBytes when both values are available.
	AvailableBytes *int64 `json:"availableBytes,omitempty"`
	// Current usedBytes / totalBytes ratio when both values are available.
	UsageRatio *float64 `json:"usageRatio,omitempty"`
	// Metrics collection timestamp in UTC.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	Source    string  `json:"source"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlStorageOverview instantiates a new PostgresqlStorageOverview object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlStorageOverview(source string) *PostgresqlStorageOverview {
	this := PostgresqlStorageOverview{}
	this.Source = source
	return &this
}

// NewPostgresqlStorageOverviewWithDefaults instantiates a new PostgresqlStorageOverview object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlStorageOverviewWithDefaults() *PostgresqlStorageOverview {
	this := PostgresqlStorageOverview{}
	return &this
}

// GetTotalBytes returns the TotalBytes field value if set, zero value otherwise.
func (o *PostgresqlStorageOverview) GetTotalBytes() int64 {
	if o == nil || o.TotalBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageOverview) GetTotalBytesOk() (*int64, bool) {
	if o == nil || o.TotalBytes == nil {
		return nil, false
	}
	return o.TotalBytes, true
}

// HasTotalBytes returns a boolean if a field has been set.
func (o *PostgresqlStorageOverview) HasTotalBytes() bool {
	return o != nil && o.TotalBytes != nil
}

// SetTotalBytes gets a reference to the given int64 and assigns it to the TotalBytes field.
func (o *PostgresqlStorageOverview) SetTotalBytes(v int64) {
	o.TotalBytes = &v
}

// GetUsedBytes returns the UsedBytes field value if set, zero value otherwise.
func (o *PostgresqlStorageOverview) GetUsedBytes() int64 {
	if o == nil || o.UsedBytes == nil {
		var ret int64
		return ret
	}
	return *o.UsedBytes
}

// GetUsedBytesOk returns a tuple with the UsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageOverview) GetUsedBytesOk() (*int64, bool) {
	if o == nil || o.UsedBytes == nil {
		return nil, false
	}
	return o.UsedBytes, true
}

// HasUsedBytes returns a boolean if a field has been set.
func (o *PostgresqlStorageOverview) HasUsedBytes() bool {
	return o != nil && o.UsedBytes != nil
}

// SetUsedBytes gets a reference to the given int64 and assigns it to the UsedBytes field.
func (o *PostgresqlStorageOverview) SetUsedBytes(v int64) {
	o.UsedBytes = &v
}

// GetAvailableBytes returns the AvailableBytes field value if set, zero value otherwise.
func (o *PostgresqlStorageOverview) GetAvailableBytes() int64 {
	if o == nil || o.AvailableBytes == nil {
		var ret int64
		return ret
	}
	return *o.AvailableBytes
}

// GetAvailableBytesOk returns a tuple with the AvailableBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageOverview) GetAvailableBytesOk() (*int64, bool) {
	if o == nil || o.AvailableBytes == nil {
		return nil, false
	}
	return o.AvailableBytes, true
}

// HasAvailableBytes returns a boolean if a field has been set.
func (o *PostgresqlStorageOverview) HasAvailableBytes() bool {
	return o != nil && o.AvailableBytes != nil
}

// SetAvailableBytes gets a reference to the given int64 and assigns it to the AvailableBytes field.
func (o *PostgresqlStorageOverview) SetAvailableBytes(v int64) {
	o.AvailableBytes = &v
}

// GetUsageRatio returns the UsageRatio field value if set, zero value otherwise.
func (o *PostgresqlStorageOverview) GetUsageRatio() float64 {
	if o == nil || o.UsageRatio == nil {
		var ret float64
		return ret
	}
	return *o.UsageRatio
}

// GetUsageRatioOk returns a tuple with the UsageRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageOverview) GetUsageRatioOk() (*float64, bool) {
	if o == nil || o.UsageRatio == nil {
		return nil, false
	}
	return o.UsageRatio, true
}

// HasUsageRatio returns a boolean if a field has been set.
func (o *PostgresqlStorageOverview) HasUsageRatio() bool {
	return o != nil && o.UsageRatio != nil
}

// SetUsageRatio gets a reference to the given float64 and assigns it to the UsageRatio field.
func (o *PostgresqlStorageOverview) SetUsageRatio(v float64) {
	o.UsageRatio = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PostgresqlStorageOverview) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageOverview) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PostgresqlStorageOverview) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PostgresqlStorageOverview) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetSource returns the Source field value.
func (o *PostgresqlStorageOverview) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PostgresqlStorageOverview) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *PostgresqlStorageOverview) SetSource(v string) {
	o.Source = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlStorageOverview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TotalBytes != nil {
		toSerialize["totalBytes"] = o.TotalBytes
	}
	if o.UsedBytes != nil {
		toSerialize["usedBytes"] = o.UsedBytes
	}
	if o.AvailableBytes != nil {
		toSerialize["availableBytes"] = o.AvailableBytes
	}
	if o.UsageRatio != nil {
		toSerialize["usageRatio"] = o.UsageRatio
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	toSerialize["source"] = o.Source

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlStorageOverview) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalBytes     *int64   `json:"totalBytes,omitempty"`
		UsedBytes      *int64   `json:"usedBytes,omitempty"`
		AvailableBytes *int64   `json:"availableBytes,omitempty"`
		UsageRatio     *float64 `json:"usageRatio,omitempty"`
		UpdatedAt      *string  `json:"updatedAt,omitempty"`
		Source         *string  `json:"source"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"totalBytes", "usedBytes", "availableBytes", "usageRatio", "updatedAt", "source"})
	} else {
		return err
	}
	o.TotalBytes = all.TotalBytes
	o.UsedBytes = all.UsedBytes
	o.AvailableBytes = all.AvailableBytes
	o.UsageRatio = all.UsageRatio
	o.UpdatedAt = all.UpdatedAt
	o.Source = *all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
