// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchNodeAllocation struct {
	Node               string                 `json:"node"`
	Host               *string                `json:"host,omitempty"`
	Ip                 *string                `json:"ip,omitempty"`
	Shards             common.NullableInt64   `json:"shards"`
	DiskIndicesBytes   common.NullableInt64   `json:"diskIndicesBytes"`
	DiskUsedBytes      common.NullableInt64   `json:"diskUsedBytes"`
	DiskAvailableBytes common.NullableInt64   `json:"diskAvailableBytes"`
	DiskTotalBytes     common.NullableInt64   `json:"diskTotalBytes"`
	DiskPercent        common.NullableFloat64 `json:"diskPercent"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchNodeAllocation instantiates a new ElasticsearchNodeAllocation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchNodeAllocation(node string, shards common.NullableInt64, diskIndicesBytes common.NullableInt64, diskUsedBytes common.NullableInt64, diskAvailableBytes common.NullableInt64, diskTotalBytes common.NullableInt64, diskPercent common.NullableFloat64) *ElasticsearchNodeAllocation {
	this := ElasticsearchNodeAllocation{}
	this.Node = node
	this.Shards = shards
	this.DiskIndicesBytes = diskIndicesBytes
	this.DiskUsedBytes = diskUsedBytes
	this.DiskAvailableBytes = diskAvailableBytes
	this.DiskTotalBytes = diskTotalBytes
	this.DiskPercent = diskPercent
	return &this
}

// NewElasticsearchNodeAllocationWithDefaults instantiates a new ElasticsearchNodeAllocation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchNodeAllocationWithDefaults() *ElasticsearchNodeAllocation {
	this := ElasticsearchNodeAllocation{}
	return &this
}

// GetNode returns the Node field value.
func (o *ElasticsearchNodeAllocation) GetNode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchNodeAllocation) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value.
func (o *ElasticsearchNodeAllocation) SetNode(v string) {
	o.Node = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ElasticsearchNodeAllocation) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchNodeAllocation) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ElasticsearchNodeAllocation) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ElasticsearchNodeAllocation) SetHost(v string) {
	o.Host = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ElasticsearchNodeAllocation) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchNodeAllocation) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ElasticsearchNodeAllocation) HasIp() bool {
	return o != nil && o.Ip != nil
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *ElasticsearchNodeAllocation) SetIp(v string) {
	o.Ip = &v
}

// GetShards returns the Shards field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchNodeAllocation) GetShards() int64 {
	if o == nil || o.Shards.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Shards.Get()
}

// GetShardsOk returns a tuple with the Shards field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchNodeAllocation) GetShardsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shards.Get(), o.Shards.IsSet()
}

// SetShards sets field value.
func (o *ElasticsearchNodeAllocation) SetShards(v int64) {
	o.Shards.Set(&v)
}

// GetDiskIndicesBytes returns the DiskIndicesBytes field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchNodeAllocation) GetDiskIndicesBytes() int64 {
	if o == nil || o.DiskIndicesBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DiskIndicesBytes.Get()
}

// GetDiskIndicesBytesOk returns a tuple with the DiskIndicesBytes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchNodeAllocation) GetDiskIndicesBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskIndicesBytes.Get(), o.DiskIndicesBytes.IsSet()
}

// SetDiskIndicesBytes sets field value.
func (o *ElasticsearchNodeAllocation) SetDiskIndicesBytes(v int64) {
	o.DiskIndicesBytes.Set(&v)
}

// GetDiskUsedBytes returns the DiskUsedBytes field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchNodeAllocation) GetDiskUsedBytes() int64 {
	if o == nil || o.DiskUsedBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DiskUsedBytes.Get()
}

// GetDiskUsedBytesOk returns a tuple with the DiskUsedBytes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchNodeAllocation) GetDiskUsedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskUsedBytes.Get(), o.DiskUsedBytes.IsSet()
}

// SetDiskUsedBytes sets field value.
func (o *ElasticsearchNodeAllocation) SetDiskUsedBytes(v int64) {
	o.DiskUsedBytes.Set(&v)
}

// GetDiskAvailableBytes returns the DiskAvailableBytes field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchNodeAllocation) GetDiskAvailableBytes() int64 {
	if o == nil || o.DiskAvailableBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DiskAvailableBytes.Get()
}

// GetDiskAvailableBytesOk returns a tuple with the DiskAvailableBytes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchNodeAllocation) GetDiskAvailableBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskAvailableBytes.Get(), o.DiskAvailableBytes.IsSet()
}

// SetDiskAvailableBytes sets field value.
func (o *ElasticsearchNodeAllocation) SetDiskAvailableBytes(v int64) {
	o.DiskAvailableBytes.Set(&v)
}

// GetDiskTotalBytes returns the DiskTotalBytes field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ElasticsearchNodeAllocation) GetDiskTotalBytes() int64 {
	if o == nil || o.DiskTotalBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DiskTotalBytes.Get()
}

// GetDiskTotalBytesOk returns a tuple with the DiskTotalBytes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchNodeAllocation) GetDiskTotalBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskTotalBytes.Get(), o.DiskTotalBytes.IsSet()
}

// SetDiskTotalBytes sets field value.
func (o *ElasticsearchNodeAllocation) SetDiskTotalBytes(v int64) {
	o.DiskTotalBytes.Set(&v)
}

// GetDiskPercent returns the DiskPercent field value.
// If the value is explicit nil, the zero value for float64 will be returned.
func (o *ElasticsearchNodeAllocation) GetDiskPercent() float64 {
	if o == nil || o.DiskPercent.Get() == nil {
		var ret float64
		return ret
	}
	return *o.DiskPercent.Get()
}

// GetDiskPercentOk returns a tuple with the DiskPercent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchNodeAllocation) GetDiskPercentOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskPercent.Get(), o.DiskPercent.IsSet()
}

// SetDiskPercent sets field value.
func (o *ElasticsearchNodeAllocation) SetDiskPercent(v float64) {
	o.DiskPercent.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchNodeAllocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["node"] = o.Node
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	toSerialize["shards"] = o.Shards.Get()
	toSerialize["diskIndicesBytes"] = o.DiskIndicesBytes.Get()
	toSerialize["diskUsedBytes"] = o.DiskUsedBytes.Get()
	toSerialize["diskAvailableBytes"] = o.DiskAvailableBytes.Get()
	toSerialize["diskTotalBytes"] = o.DiskTotalBytes.Get()
	toSerialize["diskPercent"] = o.DiskPercent.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchNodeAllocation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Node               *string                `json:"node"`
		Host               *string                `json:"host,omitempty"`
		Ip                 *string                `json:"ip,omitempty"`
		Shards             common.NullableInt64   `json:"shards"`
		DiskIndicesBytes   common.NullableInt64   `json:"diskIndicesBytes"`
		DiskUsedBytes      common.NullableInt64   `json:"diskUsedBytes"`
		DiskAvailableBytes common.NullableInt64   `json:"diskAvailableBytes"`
		DiskTotalBytes     common.NullableInt64   `json:"diskTotalBytes"`
		DiskPercent        common.NullableFloat64 `json:"diskPercent"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Node == nil {
		return fmt.Errorf("required field node missing")
	}
	if !all.Shards.IsSet() {
		return fmt.Errorf("required field shards missing")
	}
	if !all.DiskIndicesBytes.IsSet() {
		return fmt.Errorf("required field diskIndicesBytes missing")
	}
	if !all.DiskUsedBytes.IsSet() {
		return fmt.Errorf("required field diskUsedBytes missing")
	}
	if !all.DiskAvailableBytes.IsSet() {
		return fmt.Errorf("required field diskAvailableBytes missing")
	}
	if !all.DiskTotalBytes.IsSet() {
		return fmt.Errorf("required field diskTotalBytes missing")
	}
	if !all.DiskPercent.IsSet() {
		return fmt.Errorf("required field diskPercent missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"node", "host", "ip", "shards", "diskIndicesBytes", "diskUsedBytes", "diskAvailableBytes", "diskTotalBytes", "diskPercent"})
	} else {
		return err
	}
	o.Node = *all.Node
	o.Host = all.Host
	o.Ip = all.Ip
	o.Shards = all.Shards
	o.DiskIndicesBytes = all.DiskIndicesBytes
	o.DiskUsedBytes = all.DiskUsedBytes
	o.DiskAvailableBytes = all.DiskAvailableBytes
	o.DiskTotalBytes = all.DiskTotalBytes
	o.DiskPercent = all.DiskPercent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
