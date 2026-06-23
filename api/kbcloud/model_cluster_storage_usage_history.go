// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterStorageUsageHistory Cluster storage usage history collected from backend-owned fixed metrics.
type ClusterStorageUsageHistory struct {
	// Requested history range. Supported values are 24h and 7d.
	TimeRange string `json:"timeRange"`
	// History query start timestamp in UTC.
	Start string `json:"start"`
	// History query end timestamp in UTC.
	End    string                            `json:"end"`
	Points []ClusterStorageUsageHistoryPoint `json:"points"`
	// Metrics source used to build the history.
	Source string `json:"source"`
	// Backend collection timestamp in UTC.
	CollectedAt string `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterStorageUsageHistory instantiates a new ClusterStorageUsageHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterStorageUsageHistory(timeRange string, start string, end string, points []ClusterStorageUsageHistoryPoint, source string, collectedAt string) *ClusterStorageUsageHistory {
	this := ClusterStorageUsageHistory{}
	this.TimeRange = timeRange
	this.Start = start
	this.End = end
	this.Points = points
	this.Source = source
	this.CollectedAt = collectedAt
	return &this
}

// NewClusterStorageUsageHistoryWithDefaults instantiates a new ClusterStorageUsageHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterStorageUsageHistoryWithDefaults() *ClusterStorageUsageHistory {
	this := ClusterStorageUsageHistory{}
	return &this
}

// GetTimeRange returns the TimeRange field value.
func (o *ClusterStorageUsageHistory) GetTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistory) GetTimeRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value.
func (o *ClusterStorageUsageHistory) SetTimeRange(v string) {
	o.TimeRange = v
}

// GetStart returns the Start field value.
func (o *ClusterStorageUsageHistory) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistory) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *ClusterStorageUsageHistory) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value.
func (o *ClusterStorageUsageHistory) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistory) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *ClusterStorageUsageHistory) SetEnd(v string) {
	o.End = v
}

// GetPoints returns the Points field value.
func (o *ClusterStorageUsageHistory) GetPoints() []ClusterStorageUsageHistoryPoint {
	if o == nil {
		var ret []ClusterStorageUsageHistoryPoint
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistory) GetPointsOk() (*[]ClusterStorageUsageHistoryPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Points, true
}

// SetPoints sets field value.
func (o *ClusterStorageUsageHistory) SetPoints(v []ClusterStorageUsageHistoryPoint) {
	o.Points = v
}

// GetSource returns the Source field value.
func (o *ClusterStorageUsageHistory) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistory) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *ClusterStorageUsageHistory) SetSource(v string) {
	o.Source = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *ClusterStorageUsageHistory) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageUsageHistory) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *ClusterStorageUsageHistory) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterStorageUsageHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["timeRange"] = o.TimeRange
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	toSerialize["points"] = o.Points
	toSerialize["source"] = o.Source
	toSerialize["collectedAt"] = o.CollectedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterStorageUsageHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TimeRange   *string                            `json:"timeRange"`
		Start       *string                            `json:"start"`
		End         *string                            `json:"end"`
		Points      *[]ClusterStorageUsageHistoryPoint `json:"points"`
		Source      *string                            `json:"source"`
		CollectedAt *string                            `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TimeRange == nil {
		return fmt.Errorf("required field timeRange missing")
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.Points == nil {
		return fmt.Errorf("required field points missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"timeRange", "start", "end", "points", "source", "collectedAt"})
	} else {
		return err
	}
	o.TimeRange = *all.TimeRange
	o.Start = *all.Start
	o.End = *all.End
	o.Points = *all.Points
	o.Source = *all.Source
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
