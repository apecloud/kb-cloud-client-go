// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RedisKeyAnalysisReport One Redis keys analysis report row. Rows are persisted per shard. The query
// response appends one non-persisted aggregate row with isSummary=true.
type RedisKeyAnalysisReport struct {
	TaskId    *string `json:"taskId,omitempty"`
	ClusterId *string `json:"clusterId,omitempty"`
	ShardName *string `json:"shardName,omitempty"`
	// True for the aggregate row that is not persisted.
	IsSummary  *bool      `json:"isSummary,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Parameters for a Redis keys analysis task. topBigKey and topSlot are capped at 500.
	// numOfElemMin filters keys whose element count is below the threshold. bigKeySizeMin
	// filters keys whose memory usage is below the threshold (for example "10kb").
	//
	Parameters *RedisKeyAnalysisReportParameters `json:"parameters,omitempty"`
	// Key statistics by Redis data type: {type: {count, bytes}}.
	TypeStat map[string]interface{} `json:"typeStat,omitempty"`
	// Key statistics by memory size bucket: {bucket: {count, bytes}}.
	ByteLevelStat map[string]interface{} `json:"byteLevelStat,omitempty"`
	// Sum of counts across all byteLevelStat buckets.
	TotalNumKeys *int64 `json:"totalNumKeys,omitempty"`
	// Sum of bytes across all byteLevelStat buckets.
	TotlSizeKeys *string `json:"totlSizeKeys,omitempty"`
	// Key statistics by remaining TTL bucket: {bucket: {count, bytes}}.
	ExpireStat        map[string]interface{} `json:"expireStat,omitempty"`
	TopBigKeyByteStat []RedisKeyEntry        `json:"topBigKeyByteStat,omitempty"`
	TopBigKeyElemStat []RedisKeyEntry        `json:"topBigKeyElemStat,omitempty"`
	// Top cluster slots by memory usage: [{slot: bytes}]. Only populated for Redis Cluster mode.
	TopSlotKeyByteStat []map[string]int64 `json:"topSlotKeyByteStat,omitempty"`
	// Top cluster slots by key count: [{slot: count}]. Only populated for Redis Cluster mode.
	TopSlotKeyCountStat []map[string]int64 `json:"topSlotKeyCountStat,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeyAnalysisReport instantiates a new RedisKeyAnalysisReport object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeyAnalysisReport() *RedisKeyAnalysisReport {
	this := RedisKeyAnalysisReport{}
	return &this
}

// NewRedisKeyAnalysisReportWithDefaults instantiates a new RedisKeyAnalysisReport object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeyAnalysisReportWithDefaults() *RedisKeyAnalysisReport {
	this := RedisKeyAnalysisReport{}
	return &this
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasTaskId() bool {
	return o != nil && o.TaskId != nil
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *RedisKeyAnalysisReport) SetTaskId(v string) {
	o.TaskId = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasClusterId() bool {
	return o != nil && o.ClusterId != nil
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *RedisKeyAnalysisReport) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetShardName returns the ShardName field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetShardName() string {
	if o == nil || o.ShardName == nil {
		var ret string
		return ret
	}
	return *o.ShardName
}

// GetShardNameOk returns a tuple with the ShardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetShardNameOk() (*string, bool) {
	if o == nil || o.ShardName == nil {
		return nil, false
	}
	return o.ShardName, true
}

// HasShardName returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasShardName() bool {
	return o != nil && o.ShardName != nil
}

// SetShardName gets a reference to the given string and assigns it to the ShardName field.
func (o *RedisKeyAnalysisReport) SetShardName(v string) {
	o.ShardName = &v
}

// GetIsSummary returns the IsSummary field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetIsSummary() bool {
	if o == nil || o.IsSummary == nil {
		var ret bool
		return ret
	}
	return *o.IsSummary
}

// GetIsSummaryOk returns a tuple with the IsSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetIsSummaryOk() (*bool, bool) {
	if o == nil || o.IsSummary == nil {
		return nil, false
	}
	return o.IsSummary, true
}

// HasIsSummary returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasIsSummary() bool {
	return o != nil && o.IsSummary != nil
}

// SetIsSummary gets a reference to the given bool and assigns it to the IsSummary field.
func (o *RedisKeyAnalysisReport) SetIsSummary(v bool) {
	o.IsSummary = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasCreateTime() bool {
	return o != nil && o.CreateTime != nil
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *RedisKeyAnalysisReport) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetParameters() RedisKeyAnalysisReportParameters {
	if o == nil || o.Parameters == nil {
		var ret RedisKeyAnalysisReportParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetParametersOk() (*RedisKeyAnalysisReportParameters, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given RedisKeyAnalysisReportParameters and assigns it to the Parameters field.
func (o *RedisKeyAnalysisReport) SetParameters(v RedisKeyAnalysisReportParameters) {
	o.Parameters = &v
}

// GetTypeStat returns the TypeStat field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetTypeStat() map[string]interface{} {
	if o == nil || o.TypeStat == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.TypeStat
}

// GetTypeStatOk returns a tuple with the TypeStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetTypeStatOk() (*map[string]interface{}, bool) {
	if o == nil || o.TypeStat == nil {
		return nil, false
	}
	return &o.TypeStat, true
}

// HasTypeStat returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasTypeStat() bool {
	return o != nil && o.TypeStat != nil
}

// SetTypeStat gets a reference to the given map[string]interface{} and assigns it to the TypeStat field.
func (o *RedisKeyAnalysisReport) SetTypeStat(v map[string]interface{}) {
	o.TypeStat = v
}

// GetByteLevelStat returns the ByteLevelStat field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetByteLevelStat() map[string]interface{} {
	if o == nil || o.ByteLevelStat == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ByteLevelStat
}

// GetByteLevelStatOk returns a tuple with the ByteLevelStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetByteLevelStatOk() (*map[string]interface{}, bool) {
	if o == nil || o.ByteLevelStat == nil {
		return nil, false
	}
	return &o.ByteLevelStat, true
}

// HasByteLevelStat returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasByteLevelStat() bool {
	return o != nil && o.ByteLevelStat != nil
}

// SetByteLevelStat gets a reference to the given map[string]interface{} and assigns it to the ByteLevelStat field.
func (o *RedisKeyAnalysisReport) SetByteLevelStat(v map[string]interface{}) {
	o.ByteLevelStat = v
}

// GetTotalNumKeys returns the TotalNumKeys field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetTotalNumKeys() int64 {
	if o == nil || o.TotalNumKeys == nil {
		var ret int64
		return ret
	}
	return *o.TotalNumKeys
}

// GetTotalNumKeysOk returns a tuple with the TotalNumKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetTotalNumKeysOk() (*int64, bool) {
	if o == nil || o.TotalNumKeys == nil {
		return nil, false
	}
	return o.TotalNumKeys, true
}

// HasTotalNumKeys returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasTotalNumKeys() bool {
	return o != nil && o.TotalNumKeys != nil
}

// SetTotalNumKeys gets a reference to the given int64 and assigns it to the TotalNumKeys field.
func (o *RedisKeyAnalysisReport) SetTotalNumKeys(v int64) {
	o.TotalNumKeys = &v
}

// GetTotlSizeKeys returns the TotlSizeKeys field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetTotlSizeKeys() string {
	if o == nil || o.TotlSizeKeys == nil {
		var ret string
		return ret
	}
	return *o.TotlSizeKeys
}

// GetTotlSizeKeysOk returns a tuple with the TotlSizeKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetTotlSizeKeysOk() (*string, bool) {
	if o == nil || o.TotlSizeKeys == nil {
		return nil, false
	}
	return o.TotlSizeKeys, true
}

// HasTotlSizeKeys returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasTotlSizeKeys() bool {
	return o != nil && o.TotlSizeKeys != nil
}

// SetTotlSizeKeys gets a reference to the given string and assigns it to the TotlSizeKeys field.
func (o *RedisKeyAnalysisReport) SetTotlSizeKeys(v string) {
	o.TotlSizeKeys = &v
}

// GetExpireStat returns the ExpireStat field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetExpireStat() map[string]interface{} {
	if o == nil || o.ExpireStat == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExpireStat
}

// GetExpireStatOk returns a tuple with the ExpireStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetExpireStatOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExpireStat == nil {
		return nil, false
	}
	return &o.ExpireStat, true
}

// HasExpireStat returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasExpireStat() bool {
	return o != nil && o.ExpireStat != nil
}

// SetExpireStat gets a reference to the given map[string]interface{} and assigns it to the ExpireStat field.
func (o *RedisKeyAnalysisReport) SetExpireStat(v map[string]interface{}) {
	o.ExpireStat = v
}

// GetTopBigKeyByteStat returns the TopBigKeyByteStat field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetTopBigKeyByteStat() []RedisKeyEntry {
	if o == nil || o.TopBigKeyByteStat == nil {
		var ret []RedisKeyEntry
		return ret
	}
	return o.TopBigKeyByteStat
}

// GetTopBigKeyByteStatOk returns a tuple with the TopBigKeyByteStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetTopBigKeyByteStatOk() (*[]RedisKeyEntry, bool) {
	if o == nil || o.TopBigKeyByteStat == nil {
		return nil, false
	}
	return &o.TopBigKeyByteStat, true
}

// HasTopBigKeyByteStat returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasTopBigKeyByteStat() bool {
	return o != nil && o.TopBigKeyByteStat != nil
}

// SetTopBigKeyByteStat gets a reference to the given []RedisKeyEntry and assigns it to the TopBigKeyByteStat field.
func (o *RedisKeyAnalysisReport) SetTopBigKeyByteStat(v []RedisKeyEntry) {
	o.TopBigKeyByteStat = v
}

// GetTopBigKeyElemStat returns the TopBigKeyElemStat field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetTopBigKeyElemStat() []RedisKeyEntry {
	if o == nil || o.TopBigKeyElemStat == nil {
		var ret []RedisKeyEntry
		return ret
	}
	return o.TopBigKeyElemStat
}

// GetTopBigKeyElemStatOk returns a tuple with the TopBigKeyElemStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetTopBigKeyElemStatOk() (*[]RedisKeyEntry, bool) {
	if o == nil || o.TopBigKeyElemStat == nil {
		return nil, false
	}
	return &o.TopBigKeyElemStat, true
}

// HasTopBigKeyElemStat returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasTopBigKeyElemStat() bool {
	return o != nil && o.TopBigKeyElemStat != nil
}

// SetTopBigKeyElemStat gets a reference to the given []RedisKeyEntry and assigns it to the TopBigKeyElemStat field.
func (o *RedisKeyAnalysisReport) SetTopBigKeyElemStat(v []RedisKeyEntry) {
	o.TopBigKeyElemStat = v
}

// GetTopSlotKeyByteStat returns the TopSlotKeyByteStat field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetTopSlotKeyByteStat() []map[string]int64 {
	if o == nil || o.TopSlotKeyByteStat == nil {
		var ret []map[string]int64
		return ret
	}
	return o.TopSlotKeyByteStat
}

// GetTopSlotKeyByteStatOk returns a tuple with the TopSlotKeyByteStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetTopSlotKeyByteStatOk() (*[]map[string]int64, bool) {
	if o == nil || o.TopSlotKeyByteStat == nil {
		return nil, false
	}
	return &o.TopSlotKeyByteStat, true
}

// HasTopSlotKeyByteStat returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasTopSlotKeyByteStat() bool {
	return o != nil && o.TopSlotKeyByteStat != nil
}

// SetTopSlotKeyByteStat gets a reference to the given []map[string]int64 and assigns it to the TopSlotKeyByteStat field.
func (o *RedisKeyAnalysisReport) SetTopSlotKeyByteStat(v []map[string]int64) {
	o.TopSlotKeyByteStat = v
}

// GetTopSlotKeyCountStat returns the TopSlotKeyCountStat field value if set, zero value otherwise.
func (o *RedisKeyAnalysisReport) GetTopSlotKeyCountStat() []map[string]int64 {
	if o == nil || o.TopSlotKeyCountStat == nil {
		var ret []map[string]int64
		return ret
	}
	return o.TopSlotKeyCountStat
}

// GetTopSlotKeyCountStatOk returns a tuple with the TopSlotKeyCountStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyAnalysisReport) GetTopSlotKeyCountStatOk() (*[]map[string]int64, bool) {
	if o == nil || o.TopSlotKeyCountStat == nil {
		return nil, false
	}
	return &o.TopSlotKeyCountStat, true
}

// HasTopSlotKeyCountStat returns a boolean if a field has been set.
func (o *RedisKeyAnalysisReport) HasTopSlotKeyCountStat() bool {
	return o != nil && o.TopSlotKeyCountStat != nil
}

// SetTopSlotKeyCountStat gets a reference to the given []map[string]int64 and assigns it to the TopSlotKeyCountStat field.
func (o *RedisKeyAnalysisReport) SetTopSlotKeyCountStat(v []map[string]int64) {
	o.TopSlotKeyCountStat = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeyAnalysisReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}
	if o.ClusterId != nil {
		toSerialize["clusterId"] = o.ClusterId
	}
	if o.ShardName != nil {
		toSerialize["shardName"] = o.ShardName
	}
	if o.IsSummary != nil {
		toSerialize["isSummary"] = o.IsSummary
	}
	if o.CreateTime != nil {
		if o.CreateTime.Nanosecond() == 0 {
			toSerialize["createTime"] = o.CreateTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createTime"] = o.CreateTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.TypeStat != nil {
		toSerialize["typeStat"] = o.TypeStat
	}
	if o.ByteLevelStat != nil {
		toSerialize["byteLevelStat"] = o.ByteLevelStat
	}
	if o.TotalNumKeys != nil {
		toSerialize["totalNumKeys"] = o.TotalNumKeys
	}
	if o.TotlSizeKeys != nil {
		toSerialize["totlSizeKeys"] = o.TotlSizeKeys
	}
	if o.ExpireStat != nil {
		toSerialize["expireStat"] = o.ExpireStat
	}
	if o.TopBigKeyByteStat != nil {
		toSerialize["topBigKeyByteStat"] = o.TopBigKeyByteStat
	}
	if o.TopBigKeyElemStat != nil {
		toSerialize["topBigKeyElemStat"] = o.TopBigKeyElemStat
	}
	if o.TopSlotKeyByteStat != nil {
		toSerialize["topSlotKeyByteStat"] = o.TopSlotKeyByteStat
	}
	if o.TopSlotKeyCountStat != nil {
		toSerialize["topSlotKeyCountStat"] = o.TopSlotKeyCountStat
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisKeyAnalysisReport) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TaskId              *string                           `json:"taskId,omitempty"`
		ClusterId           *string                           `json:"clusterId,omitempty"`
		ShardName           *string                           `json:"shardName,omitempty"`
		IsSummary           *bool                             `json:"isSummary,omitempty"`
		CreateTime          *time.Time                        `json:"createTime,omitempty"`
		Parameters          *RedisKeyAnalysisReportParameters `json:"parameters,omitempty"`
		TypeStat            map[string]interface{}            `json:"typeStat,omitempty"`
		ByteLevelStat       map[string]interface{}            `json:"byteLevelStat,omitempty"`
		TotalNumKeys        *int64                            `json:"totalNumKeys,omitempty"`
		TotlSizeKeys        *string                           `json:"totlSizeKeys,omitempty"`
		ExpireStat          map[string]interface{}            `json:"expireStat,omitempty"`
		TopBigKeyByteStat   []RedisKeyEntry                   `json:"topBigKeyByteStat,omitempty"`
		TopBigKeyElemStat   []RedisKeyEntry                   `json:"topBigKeyElemStat,omitempty"`
		TopSlotKeyByteStat  []map[string]int64                `json:"topSlotKeyByteStat,omitempty"`
		TopSlotKeyCountStat []map[string]int64                `json:"topSlotKeyCountStat,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"taskId", "clusterId", "shardName", "isSummary", "createTime", "parameters", "typeStat", "byteLevelStat", "totalNumKeys", "totlSizeKeys", "expireStat", "topBigKeyByteStat", "topBigKeyElemStat", "topSlotKeyByteStat", "topSlotKeyCountStat"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TaskId = all.TaskId
	o.ClusterId = all.ClusterId
	o.ShardName = all.ShardName
	o.IsSummary = all.IsSummary
	o.CreateTime = all.CreateTime
	if all.Parameters != nil && all.Parameters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Parameters = all.Parameters
	o.TypeStat = all.TypeStat
	o.ByteLevelStat = all.ByteLevelStat
	o.TotalNumKeys = all.TotalNumKeys
	o.TotlSizeKeys = all.TotlSizeKeys
	o.ExpireStat = all.ExpireStat
	o.TopBigKeyByteStat = all.TopBigKeyByteStat
	o.TopBigKeyElemStat = all.TopBigKeyElemStat
	o.TopSlotKeyByteStat = all.TopSlotKeyByteStat
	o.TopSlotKeyCountStat = all.TopSlotKeyCountStat

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
