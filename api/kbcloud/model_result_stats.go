// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION ResultStats
type ResultStats struct {
	// NODESCRIPTION ColumnsCount
	ColumnsCount *int32 `json:"columns_count,omitempty"`
	// NODESCRIPTION RowsCount
	RowsCount *int32 `json:"rows_count,omitempty"`
	// NODESCRIPTION RowsAffected
	RowsAffected *int32 `json:"rows_affected,omitempty"`
	// NODESCRIPTION QueryStartTime
	QueryStartTime *time.Time `json:"query_start_time,omitempty"`
	// NODESCRIPTION QueryFinishTime
	QueryFinishTime *time.Time `json:"query_finish_time,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResultStats instantiates a new ResultStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResultStats() *ResultStats {
	this := ResultStats{}
	return &this
}

// NewResultStatsWithDefaults instantiates a new ResultStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResultStatsWithDefaults() *ResultStats {
	this := ResultStats{}
	return &this
}

// GetColumnsCount returns the ColumnsCount field value if set, zero value otherwise.
func (o *ResultStats) GetColumnsCount() int32 {
	if o == nil || o.ColumnsCount == nil {
		var ret int32
		return ret
	}
	return *o.ColumnsCount
}

// GetColumnsCountOk returns a tuple with the ColumnsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultStats) GetColumnsCountOk() (*int32, bool) {
	if o == nil || o.ColumnsCount == nil {
		return nil, false
	}
	return o.ColumnsCount, true
}

// HasColumnsCount returns a boolean if a field has been set.
func (o *ResultStats) HasColumnsCount() bool {
	return o != nil && o.ColumnsCount != nil
}

// SetColumnsCount gets a reference to the given int32 and assigns it to the ColumnsCount field.
func (o *ResultStats) SetColumnsCount(v int32) {
	o.ColumnsCount = &v
}

// GetRowsCount returns the RowsCount field value if set, zero value otherwise.
func (o *ResultStats) GetRowsCount() int32 {
	if o == nil || o.RowsCount == nil {
		var ret int32
		return ret
	}
	return *o.RowsCount
}

// GetRowsCountOk returns a tuple with the RowsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultStats) GetRowsCountOk() (*int32, bool) {
	if o == nil || o.RowsCount == nil {
		return nil, false
	}
	return o.RowsCount, true
}

// HasRowsCount returns a boolean if a field has been set.
func (o *ResultStats) HasRowsCount() bool {
	return o != nil && o.RowsCount != nil
}

// SetRowsCount gets a reference to the given int32 and assigns it to the RowsCount field.
func (o *ResultStats) SetRowsCount(v int32) {
	o.RowsCount = &v
}

// GetRowsAffected returns the RowsAffected field value if set, zero value otherwise.
func (o *ResultStats) GetRowsAffected() int32 {
	if o == nil || o.RowsAffected == nil {
		var ret int32
		return ret
	}
	return *o.RowsAffected
}

// GetRowsAffectedOk returns a tuple with the RowsAffected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultStats) GetRowsAffectedOk() (*int32, bool) {
	if o == nil || o.RowsAffected == nil {
		return nil, false
	}
	return o.RowsAffected, true
}

// HasRowsAffected returns a boolean if a field has been set.
func (o *ResultStats) HasRowsAffected() bool {
	return o != nil && o.RowsAffected != nil
}

// SetRowsAffected gets a reference to the given int32 and assigns it to the RowsAffected field.
func (o *ResultStats) SetRowsAffected(v int32) {
	o.RowsAffected = &v
}

// GetQueryStartTime returns the QueryStartTime field value if set, zero value otherwise.
func (o *ResultStats) GetQueryStartTime() time.Time {
	if o == nil || o.QueryStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.QueryStartTime
}

// GetQueryStartTimeOk returns a tuple with the QueryStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultStats) GetQueryStartTimeOk() (*time.Time, bool) {
	if o == nil || o.QueryStartTime == nil {
		return nil, false
	}
	return o.QueryStartTime, true
}

// HasQueryStartTime returns a boolean if a field has been set.
func (o *ResultStats) HasQueryStartTime() bool {
	return o != nil && o.QueryStartTime != nil
}

// SetQueryStartTime gets a reference to the given time.Time and assigns it to the QueryStartTime field.
func (o *ResultStats) SetQueryStartTime(v time.Time) {
	o.QueryStartTime = &v
}

// GetQueryFinishTime returns the QueryFinishTime field value if set, zero value otherwise.
func (o *ResultStats) GetQueryFinishTime() time.Time {
	if o == nil || o.QueryFinishTime == nil {
		var ret time.Time
		return ret
	}
	return *o.QueryFinishTime
}

// GetQueryFinishTimeOk returns a tuple with the QueryFinishTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultStats) GetQueryFinishTimeOk() (*time.Time, bool) {
	if o == nil || o.QueryFinishTime == nil {
		return nil, false
	}
	return o.QueryFinishTime, true
}

// HasQueryFinishTime returns a boolean if a field has been set.
func (o *ResultStats) HasQueryFinishTime() bool {
	return o != nil && o.QueryFinishTime != nil
}

// SetQueryFinishTime gets a reference to the given time.Time and assigns it to the QueryFinishTime field.
func (o *ResultStats) SetQueryFinishTime(v time.Time) {
	o.QueryFinishTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResultStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ColumnsCount != nil {
		toSerialize["columns_count"] = o.ColumnsCount
	}
	if o.RowsCount != nil {
		toSerialize["rows_count"] = o.RowsCount
	}
	if o.RowsAffected != nil {
		toSerialize["rows_affected"] = o.RowsAffected
	}
	if o.QueryStartTime != nil {
		if o.QueryStartTime.Nanosecond() == 0 {
			toSerialize["query_start_time"] = o.QueryStartTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["query_start_time"] = o.QueryStartTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.QueryFinishTime != nil {
		if o.QueryFinishTime.Nanosecond() == 0 {
			toSerialize["query_finish_time"] = o.QueryFinishTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["query_finish_time"] = o.QueryFinishTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResultStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnsCount    *int32     `json:"columns_count,omitempty"`
		RowsCount       *int32     `json:"rows_count,omitempty"`
		RowsAffected    *int32     `json:"rows_affected,omitempty"`
		QueryStartTime  *time.Time `json:"query_start_time,omitempty"`
		QueryFinishTime *time.Time `json:"query_finish_time,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"columns_count", "rows_count", "rows_affected", "query_start_time", "query_finish_time"})
	} else {
		return err
	}
	o.ColumnsCount = all.ColumnsCount
	o.RowsCount = all.RowsCount
	o.RowsAffected = all.RowsAffected
	o.QueryStartTime = all.QueryStartTime
	o.QueryFinishTime = all.QueryFinishTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
