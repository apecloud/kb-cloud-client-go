// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OutageRecord Records the outage events of a cluster for SLA calculation. It is designed to ensure that for any given cluster, there can be at most one active outage record at any time.
type OutageRecord struct {
	// The unique identifier for the outage record.
	Id *string `json:"id,omitempty"`
	// The id of cluster.
	ClusterId string `json:"clusterID"`
	// The Unix timestamp (in seconds) when the outage began (UTC).
	OutageStartTime *int64 `json:"outageStartTime,omitempty"`
	// The Unix timestamp (in seconds) when the outage was resolved (UTC). A null value indicates that the outage is still ongoing.
	OutageEndTime common.NullableInt64 `json:"outageEndTime,omitempty"`
	// The error message from the last failed probe.
	LastFailureReason common.NullableString `json:"lastFailureReason,omitempty"`
	// The number of consecutive failures in this outage event.
	FailureCount common.NullableInt32 `json:"failureCount,omitempty"`
	// The timestamp when the record was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The timestamp when the record was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOutageRecord instantiates a new OutageRecord object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOutageRecord(clusterId string) *OutageRecord {
	this := OutageRecord{}
	this.ClusterId = clusterId
	var failureCount int32 = 1
	this.FailureCount = *common.NewNullableInt32(&failureCount)
	return &this
}

// NewOutageRecordWithDefaults instantiates a new OutageRecord object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOutageRecordWithDefaults() *OutageRecord {
	this := OutageRecord{}
	var failureCount int32 = 1
	this.FailureCount = *common.NewNullableInt32(&failureCount)
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OutageRecord) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutageRecord) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OutageRecord) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OutageRecord) SetId(v string) {
	o.Id = &v
}

// GetClusterId returns the ClusterId field value.
func (o *OutageRecord) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *OutageRecord) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value.
func (o *OutageRecord) SetClusterId(v string) {
	o.ClusterId = v
}

// GetOutageStartTime returns the OutageStartTime field value if set, zero value otherwise.
func (o *OutageRecord) GetOutageStartTime() int64 {
	if o == nil || o.OutageStartTime == nil {
		var ret int64
		return ret
	}
	return *o.OutageStartTime
}

// GetOutageStartTimeOk returns a tuple with the OutageStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutageRecord) GetOutageStartTimeOk() (*int64, bool) {
	if o == nil || o.OutageStartTime == nil {
		return nil, false
	}
	return o.OutageStartTime, true
}

// HasOutageStartTime returns a boolean if a field has been set.
func (o *OutageRecord) HasOutageStartTime() bool {
	return o != nil && o.OutageStartTime != nil
}

// SetOutageStartTime gets a reference to the given int64 and assigns it to the OutageStartTime field.
func (o *OutageRecord) SetOutageStartTime(v int64) {
	o.OutageStartTime = &v
}

// GetOutageEndTime returns the OutageEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutageRecord) GetOutageEndTime() int64 {
	if o == nil || o.OutageEndTime.Get() == nil {
		var ret int64
		return ret
	}
	return *o.OutageEndTime.Get()
}

// GetOutageEndTimeOk returns a tuple with the OutageEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OutageRecord) GetOutageEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutageEndTime.Get(), o.OutageEndTime.IsSet()
}

// HasOutageEndTime returns a boolean if a field has been set.
func (o *OutageRecord) HasOutageEndTime() bool {
	return o != nil && o.OutageEndTime.IsSet()
}

// SetOutageEndTime gets a reference to the given common.NullableInt64 and assigns it to the OutageEndTime field.
func (o *OutageRecord) SetOutageEndTime(v int64) {
	o.OutageEndTime.Set(&v)
}

// SetOutageEndTimeNil sets the value for OutageEndTime to be an explicit nil.
func (o *OutageRecord) SetOutageEndTimeNil() {
	o.OutageEndTime.Set(nil)
}

// UnsetOutageEndTime ensures that no value is present for OutageEndTime, not even an explicit nil.
func (o *OutageRecord) UnsetOutageEndTime() {
	o.OutageEndTime.Unset()
}

// GetLastFailureReason returns the LastFailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutageRecord) GetLastFailureReason() string {
	if o == nil || o.LastFailureReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastFailureReason.Get()
}

// GetLastFailureReasonOk returns a tuple with the LastFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OutageRecord) GetLastFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastFailureReason.Get(), o.LastFailureReason.IsSet()
}

// HasLastFailureReason returns a boolean if a field has been set.
func (o *OutageRecord) HasLastFailureReason() bool {
	return o != nil && o.LastFailureReason.IsSet()
}

// SetLastFailureReason gets a reference to the given common.NullableString and assigns it to the LastFailureReason field.
func (o *OutageRecord) SetLastFailureReason(v string) {
	o.LastFailureReason.Set(&v)
}

// SetLastFailureReasonNil sets the value for LastFailureReason to be an explicit nil.
func (o *OutageRecord) SetLastFailureReasonNil() {
	o.LastFailureReason.Set(nil)
}

// UnsetLastFailureReason ensures that no value is present for LastFailureReason, not even an explicit nil.
func (o *OutageRecord) UnsetLastFailureReason() {
	o.LastFailureReason.Unset()
}

// GetFailureCount returns the FailureCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutageRecord) GetFailureCount() int32 {
	if o == nil || o.FailureCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.FailureCount.Get()
}

// GetFailureCountOk returns a tuple with the FailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OutageRecord) GetFailureCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureCount.Get(), o.FailureCount.IsSet()
}

// HasFailureCount returns a boolean if a field has been set.
func (o *OutageRecord) HasFailureCount() bool {
	return o != nil && o.FailureCount.IsSet()
}

// SetFailureCount gets a reference to the given common.NullableInt32 and assigns it to the FailureCount field.
func (o *OutageRecord) SetFailureCount(v int32) {
	o.FailureCount.Set(&v)
}

// SetFailureCountNil sets the value for FailureCount to be an explicit nil.
func (o *OutageRecord) SetFailureCountNil() {
	o.FailureCount.Set(nil)
}

// UnsetFailureCount ensures that no value is present for FailureCount, not even an explicit nil.
func (o *OutageRecord) UnsetFailureCount() {
	o.FailureCount.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OutageRecord) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutageRecord) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OutageRecord) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OutageRecord) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OutageRecord) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutageRecord) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OutageRecord) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OutageRecord) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OutageRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["clusterID"] = o.ClusterId
	if o.OutageStartTime != nil {
		toSerialize["outageStartTime"] = o.OutageStartTime
	}
	if o.OutageEndTime.IsSet() {
		toSerialize["outageEndTime"] = o.OutageEndTime.Get()
	}
	if o.LastFailureReason.IsSet() {
		toSerialize["lastFailureReason"] = o.LastFailureReason.Get()
	}
	if o.FailureCount.IsSet() {
		toSerialize["failureCount"] = o.FailureCount.Get()
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OutageRecord) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                *string               `json:"id,omitempty"`
		ClusterId         *string               `json:"clusterID"`
		OutageStartTime   *int64                `json:"outageStartTime,omitempty"`
		OutageEndTime     common.NullableInt64  `json:"outageEndTime,omitempty"`
		LastFailureReason common.NullableString `json:"lastFailureReason,omitempty"`
		FailureCount      common.NullableInt32  `json:"failureCount,omitempty"`
		CreatedAt         *time.Time            `json:"createdAt,omitempty"`
		UpdatedAt         *time.Time            `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ClusterId == nil {
		return fmt.Errorf("required field clusterID missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "clusterID", "outageStartTime", "outageEndTime", "lastFailureReason", "failureCount", "createdAt", "updatedAt"})
	} else {
		return err
	}
	o.Id = all.Id
	o.ClusterId = *all.ClusterId
	o.OutageStartTime = all.OutageStartTime
	o.OutageEndTime = all.OutageEndTime
	o.LastFailureReason = all.LastFailureReason
	o.FailureCount = all.FailureCount
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
