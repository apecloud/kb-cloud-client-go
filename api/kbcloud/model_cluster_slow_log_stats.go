// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterSlowLogStats Summary statistics for slow logs
type ClusterSlowLogStats struct {
	TotalSlowLogs    *int64   `json:"totalSlowLogs,omitempty"`
	AvgExecutionTime *float64 `json:"avgExecutionTime,omitempty"`
	P95Latency       *float64 `json:"p95Latency,omitempty"`
	ActiveDBs        *int64   `json:"activeDBs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterSlowLogStats instantiates a new ClusterSlowLogStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterSlowLogStats() *ClusterSlowLogStats {
	this := ClusterSlowLogStats{}
	return &this
}

// NewClusterSlowLogStatsWithDefaults instantiates a new ClusterSlowLogStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterSlowLogStatsWithDefaults() *ClusterSlowLogStats {
	this := ClusterSlowLogStats{}
	return &this
}

// GetTotalSlowLogs returns the TotalSlowLogs field value if set, zero value otherwise.
func (o *ClusterSlowLogStats) GetTotalSlowLogs() int64 {
	if o == nil || o.TotalSlowLogs == nil {
		var ret int64
		return ret
	}
	return *o.TotalSlowLogs
}

// GetTotalSlowLogsOk returns a tuple with the TotalSlowLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogStats) GetTotalSlowLogsOk() (*int64, bool) {
	if o == nil || o.TotalSlowLogs == nil {
		return nil, false
	}
	return o.TotalSlowLogs, true
}

// HasTotalSlowLogs returns a boolean if a field has been set.
func (o *ClusterSlowLogStats) HasTotalSlowLogs() bool {
	return o != nil && o.TotalSlowLogs != nil
}

// SetTotalSlowLogs gets a reference to the given int64 and assigns it to the TotalSlowLogs field.
func (o *ClusterSlowLogStats) SetTotalSlowLogs(v int64) {
	o.TotalSlowLogs = &v
}

// GetAvgExecutionTime returns the AvgExecutionTime field value if set, zero value otherwise.
func (o *ClusterSlowLogStats) GetAvgExecutionTime() float64 {
	if o == nil || o.AvgExecutionTime == nil {
		var ret float64
		return ret
	}
	return *o.AvgExecutionTime
}

// GetAvgExecutionTimeOk returns a tuple with the AvgExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogStats) GetAvgExecutionTimeOk() (*float64, bool) {
	if o == nil || o.AvgExecutionTime == nil {
		return nil, false
	}
	return o.AvgExecutionTime, true
}

// HasAvgExecutionTime returns a boolean if a field has been set.
func (o *ClusterSlowLogStats) HasAvgExecutionTime() bool {
	return o != nil && o.AvgExecutionTime != nil
}

// SetAvgExecutionTime gets a reference to the given float64 and assigns it to the AvgExecutionTime field.
func (o *ClusterSlowLogStats) SetAvgExecutionTime(v float64) {
	o.AvgExecutionTime = &v
}

// GetP95Latency returns the P95Latency field value if set, zero value otherwise.
func (o *ClusterSlowLogStats) GetP95Latency() float64 {
	if o == nil || o.P95Latency == nil {
		var ret float64
		return ret
	}
	return *o.P95Latency
}

// GetP95LatencyOk returns a tuple with the P95Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogStats) GetP95LatencyOk() (*float64, bool) {
	if o == nil || o.P95Latency == nil {
		return nil, false
	}
	return o.P95Latency, true
}

// HasP95Latency returns a boolean if a field has been set.
func (o *ClusterSlowLogStats) HasP95Latency() bool {
	return o != nil && o.P95Latency != nil
}

// SetP95Latency gets a reference to the given float64 and assigns it to the P95Latency field.
func (o *ClusterSlowLogStats) SetP95Latency(v float64) {
	o.P95Latency = &v
}

// GetActiveDBs returns the ActiveDBs field value if set, zero value otherwise.
func (o *ClusterSlowLogStats) GetActiveDBs() int64 {
	if o == nil || o.ActiveDBs == nil {
		var ret int64
		return ret
	}
	return *o.ActiveDBs
}

// GetActiveDBsOk returns a tuple with the ActiveDBs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogStats) GetActiveDBsOk() (*int64, bool) {
	if o == nil || o.ActiveDBs == nil {
		return nil, false
	}
	return o.ActiveDBs, true
}

// HasActiveDBs returns a boolean if a field has been set.
func (o *ClusterSlowLogStats) HasActiveDBs() bool {
	return o != nil && o.ActiveDBs != nil
}

// SetActiveDBs gets a reference to the given int64 and assigns it to the ActiveDBs field.
func (o *ClusterSlowLogStats) SetActiveDBs(v int64) {
	o.ActiveDBs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterSlowLogStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TotalSlowLogs != nil {
		toSerialize["totalSlowLogs"] = o.TotalSlowLogs
	}
	if o.AvgExecutionTime != nil {
		toSerialize["avgExecutionTime"] = o.AvgExecutionTime
	}
	if o.P95Latency != nil {
		toSerialize["p95Latency"] = o.P95Latency
	}
	if o.ActiveDBs != nil {
		toSerialize["activeDBs"] = o.ActiveDBs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterSlowLogStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalSlowLogs    *int64   `json:"totalSlowLogs,omitempty"`
		AvgExecutionTime *float64 `json:"avgExecutionTime,omitempty"`
		P95Latency       *float64 `json:"p95Latency,omitempty"`
		ActiveDBs        *int64   `json:"activeDBs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"totalSlowLogs", "avgExecutionTime", "p95Latency", "activeDBs"})
	} else {
		return err
	}
	o.TotalSlowLogs = all.TotalSlowLogs
	o.AvgExecutionTime = all.AvgExecutionTime
	o.P95Latency = all.P95Latency
	o.ActiveDBs = all.ActiveDBs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
