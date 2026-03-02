// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterExecutionLogAggregateItem Aggregate item for grouped execution logs
type ClusterExecutionLogAggregateItem struct {
	GroupKey         *string  `json:"groupKey,omitempty"`
	Count            *int64   `json:"count,omitempty"`
	AvgExecutionTime *float64 `json:"avgExecutionTime,omitempty"`
	MaxExecutionTime *float64 `json:"maxExecutionTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterExecutionLogAggregateItem instantiates a new ClusterExecutionLogAggregateItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterExecutionLogAggregateItem() *ClusterExecutionLogAggregateItem {
	this := ClusterExecutionLogAggregateItem{}
	return &this
}

// NewClusterExecutionLogAggregateItemWithDefaults instantiates a new ClusterExecutionLogAggregateItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterExecutionLogAggregateItemWithDefaults() *ClusterExecutionLogAggregateItem {
	this := ClusterExecutionLogAggregateItem{}
	return &this
}

// GetGroupKey returns the GroupKey field value if set, zero value otherwise.
func (o *ClusterExecutionLogAggregateItem) GetGroupKey() string {
	if o == nil || o.GroupKey == nil {
		var ret string
		return ret
	}
	return *o.GroupKey
}

// GetGroupKeyOk returns a tuple with the GroupKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogAggregateItem) GetGroupKeyOk() (*string, bool) {
	if o == nil || o.GroupKey == nil {
		return nil, false
	}
	return o.GroupKey, true
}

// HasGroupKey returns a boolean if a field has been set.
func (o *ClusterExecutionLogAggregateItem) HasGroupKey() bool {
	return o != nil && o.GroupKey != nil
}

// SetGroupKey gets a reference to the given string and assigns it to the GroupKey field.
func (o *ClusterExecutionLogAggregateItem) SetGroupKey(v string) {
	o.GroupKey = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ClusterExecutionLogAggregateItem) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogAggregateItem) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ClusterExecutionLogAggregateItem) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ClusterExecutionLogAggregateItem) SetCount(v int64) {
	o.Count = &v
}

// GetAvgExecutionTime returns the AvgExecutionTime field value if set, zero value otherwise.
func (o *ClusterExecutionLogAggregateItem) GetAvgExecutionTime() float64 {
	if o == nil || o.AvgExecutionTime == nil {
		var ret float64
		return ret
	}
	return *o.AvgExecutionTime
}

// GetAvgExecutionTimeOk returns a tuple with the AvgExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogAggregateItem) GetAvgExecutionTimeOk() (*float64, bool) {
	if o == nil || o.AvgExecutionTime == nil {
		return nil, false
	}
	return o.AvgExecutionTime, true
}

// HasAvgExecutionTime returns a boolean if a field has been set.
func (o *ClusterExecutionLogAggregateItem) HasAvgExecutionTime() bool {
	return o != nil && o.AvgExecutionTime != nil
}

// SetAvgExecutionTime gets a reference to the given float64 and assigns it to the AvgExecutionTime field.
func (o *ClusterExecutionLogAggregateItem) SetAvgExecutionTime(v float64) {
	o.AvgExecutionTime = &v
}

// GetMaxExecutionTime returns the MaxExecutionTime field value if set, zero value otherwise.
func (o *ClusterExecutionLogAggregateItem) GetMaxExecutionTime() float64 {
	if o == nil || o.MaxExecutionTime == nil {
		var ret float64
		return ret
	}
	return *o.MaxExecutionTime
}

// GetMaxExecutionTimeOk returns a tuple with the MaxExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterExecutionLogAggregateItem) GetMaxExecutionTimeOk() (*float64, bool) {
	if o == nil || o.MaxExecutionTime == nil {
		return nil, false
	}
	return o.MaxExecutionTime, true
}

// HasMaxExecutionTime returns a boolean if a field has been set.
func (o *ClusterExecutionLogAggregateItem) HasMaxExecutionTime() bool {
	return o != nil && o.MaxExecutionTime != nil
}

// SetMaxExecutionTime gets a reference to the given float64 and assigns it to the MaxExecutionTime field.
func (o *ClusterExecutionLogAggregateItem) SetMaxExecutionTime(v float64) {
	o.MaxExecutionTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterExecutionLogAggregateItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.GroupKey != nil {
		toSerialize["groupKey"] = o.GroupKey
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.AvgExecutionTime != nil {
		toSerialize["avgExecutionTime"] = o.AvgExecutionTime
	}
	if o.MaxExecutionTime != nil {
		toSerialize["maxExecutionTime"] = o.MaxExecutionTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterExecutionLogAggregateItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupKey         *string  `json:"groupKey,omitempty"`
		Count            *int64   `json:"count,omitempty"`
		AvgExecutionTime *float64 `json:"avgExecutionTime,omitempty"`
		MaxExecutionTime *float64 `json:"maxExecutionTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"groupKey", "count", "avgExecutionTime", "maxExecutionTime"})
	} else {
		return err
	}
	o.GroupKey = all.GroupKey
	o.Count = all.Count
	o.AvgExecutionTime = all.AvgExecutionTime
	o.MaxExecutionTime = all.MaxExecutionTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
