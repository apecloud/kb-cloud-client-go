// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterAlertSummary Cluster alert summary
type ClusterAlertSummary struct {
	// number of active alerts after cluster/status/severity filters
	ActiveCount *int64 `json:"activeCount,omitempty"`
	// active alert count grouped by severity
	SeverityCounts map[string]int64 `json:"severityCounts,omitempty"`
	// top alerts for the cluster after filters
	TopAlerts []AlertObject `json:"topAlerts,omitempty"`
	// response generation time
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterAlertSummary instantiates a new ClusterAlertSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterAlertSummary() *ClusterAlertSummary {
	this := ClusterAlertSummary{}
	return &this
}

// NewClusterAlertSummaryWithDefaults instantiates a new ClusterAlertSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterAlertSummaryWithDefaults() *ClusterAlertSummary {
	this := ClusterAlertSummary{}
	return &this
}

// GetActiveCount returns the ActiveCount field value if set, zero value otherwise.
func (o *ClusterAlertSummary) GetActiveCount() int64 {
	if o == nil || o.ActiveCount == nil {
		var ret int64
		return ret
	}
	return *o.ActiveCount
}

// GetActiveCountOk returns a tuple with the ActiveCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertSummary) GetActiveCountOk() (*int64, bool) {
	if o == nil || o.ActiveCount == nil {
		return nil, false
	}
	return o.ActiveCount, true
}

// HasActiveCount returns a boolean if a field has been set.
func (o *ClusterAlertSummary) HasActiveCount() bool {
	return o != nil && o.ActiveCount != nil
}

// SetActiveCount gets a reference to the given int64 and assigns it to the ActiveCount field.
func (o *ClusterAlertSummary) SetActiveCount(v int64) {
	o.ActiveCount = &v
}

// GetSeverityCounts returns the SeverityCounts field value if set, zero value otherwise.
func (o *ClusterAlertSummary) GetSeverityCounts() map[string]int64 {
	if o == nil || o.SeverityCounts == nil {
		var ret map[string]int64
		return ret
	}
	return o.SeverityCounts
}

// GetSeverityCountsOk returns a tuple with the SeverityCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertSummary) GetSeverityCountsOk() (*map[string]int64, bool) {
	if o == nil || o.SeverityCounts == nil {
		return nil, false
	}
	return &o.SeverityCounts, true
}

// HasSeverityCounts returns a boolean if a field has been set.
func (o *ClusterAlertSummary) HasSeverityCounts() bool {
	return o != nil && o.SeverityCounts != nil
}

// SetSeverityCounts gets a reference to the given map[string]int64 and assigns it to the SeverityCounts field.
func (o *ClusterAlertSummary) SetSeverityCounts(v map[string]int64) {
	o.SeverityCounts = v
}

// GetTopAlerts returns the TopAlerts field value if set, zero value otherwise.
func (o *ClusterAlertSummary) GetTopAlerts() []AlertObject {
	if o == nil || o.TopAlerts == nil {
		var ret []AlertObject
		return ret
	}
	return o.TopAlerts
}

// GetTopAlertsOk returns a tuple with the TopAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertSummary) GetTopAlertsOk() (*[]AlertObject, bool) {
	if o == nil || o.TopAlerts == nil {
		return nil, false
	}
	return &o.TopAlerts, true
}

// HasTopAlerts returns a boolean if a field has been set.
func (o *ClusterAlertSummary) HasTopAlerts() bool {
	return o != nil && o.TopAlerts != nil
}

// SetTopAlerts gets a reference to the given []AlertObject and assigns it to the TopAlerts field.
func (o *ClusterAlertSummary) SetTopAlerts(v []AlertObject) {
	o.TopAlerts = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ClusterAlertSummary) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAlertSummary) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ClusterAlertSummary) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ClusterAlertSummary) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterAlertSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ActiveCount != nil {
		toSerialize["activeCount"] = o.ActiveCount
	}
	if o.SeverityCounts != nil {
		toSerialize["severityCounts"] = o.SeverityCounts
	}
	if o.TopAlerts != nil {
		toSerialize["topAlerts"] = o.TopAlerts
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
func (o *ClusterAlertSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActiveCount    *int64           `json:"activeCount,omitempty"`
		SeverityCounts map[string]int64 `json:"severityCounts,omitempty"`
		TopAlerts      []AlertObject    `json:"topAlerts,omitempty"`
		UpdatedAt      *time.Time       `json:"updatedAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"activeCount", "severityCounts", "topAlerts", "updatedAt"})
	} else {
		return err
	}
	o.ActiveCount = all.ActiveCount
	o.SeverityCounts = all.SeverityCounts
	o.TopAlerts = all.TopAlerts
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
