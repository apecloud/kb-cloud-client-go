// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// SLA The SLA (Service Level Agreement) for a cluster.
type SLA struct {
	// KubeBlocks cluster information
	Cluster *ClusterListItem `json:"cluster,omitempty"`
	// A list of outage records.
	Detail []OutageRecord `json:"detail,omitempty"`
	// The SLA value for the cluster.
	Sla *float64 `json:"sla,omitempty"`
	// The total duration of the outage records.
	TotalDuration *int64 `json:"totalDuration,omitempty"`
	// The total downtime duration of the outage records.
	TotalDowntime *int64 `json:"totalDowntime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSLA instantiates a new SLA object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSLA() *SLA {
	this := SLA{}
	return &this
}

// NewSLAWithDefaults instantiates a new SLA object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSLAWithDefaults() *SLA {
	this := SLA{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *SLA) GetCluster() ClusterListItem {
	if o == nil || o.Cluster == nil {
		var ret ClusterListItem
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLA) GetClusterOk() (*ClusterListItem, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *SLA) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given ClusterListItem and assigns it to the Cluster field.
func (o *SLA) SetCluster(v ClusterListItem) {
	o.Cluster = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *SLA) GetDetail() []OutageRecord {
	if o == nil || o.Detail == nil {
		var ret []OutageRecord
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLA) GetDetailOk() (*[]OutageRecord, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return &o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *SLA) HasDetail() bool {
	return o != nil && o.Detail != nil
}

// SetDetail gets a reference to the given []OutageRecord and assigns it to the Detail field.
func (o *SLA) SetDetail(v []OutageRecord) {
	o.Detail = v
}

// GetSla returns the Sla field value if set, zero value otherwise.
func (o *SLA) GetSla() float64 {
	if o == nil || o.Sla == nil {
		var ret float64
		return ret
	}
	return *o.Sla
}

// GetSlaOk returns a tuple with the Sla field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLA) GetSlaOk() (*float64, bool) {
	if o == nil || o.Sla == nil {
		return nil, false
	}
	return o.Sla, true
}

// HasSla returns a boolean if a field has been set.
func (o *SLA) HasSla() bool {
	return o != nil && o.Sla != nil
}

// SetSla gets a reference to the given float64 and assigns it to the Sla field.
func (o *SLA) SetSla(v float64) {
	o.Sla = &v
}

// GetTotalDuration returns the TotalDuration field value if set, zero value otherwise.
func (o *SLA) GetTotalDuration() int64 {
	if o == nil || o.TotalDuration == nil {
		var ret int64
		return ret
	}
	return *o.TotalDuration
}

// GetTotalDurationOk returns a tuple with the TotalDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLA) GetTotalDurationOk() (*int64, bool) {
	if o == nil || o.TotalDuration == nil {
		return nil, false
	}
	return o.TotalDuration, true
}

// HasTotalDuration returns a boolean if a field has been set.
func (o *SLA) HasTotalDuration() bool {
	return o != nil && o.TotalDuration != nil
}

// SetTotalDuration gets a reference to the given int64 and assigns it to the TotalDuration field.
func (o *SLA) SetTotalDuration(v int64) {
	o.TotalDuration = &v
}

// GetTotalDowntime returns the TotalDowntime field value if set, zero value otherwise.
func (o *SLA) GetTotalDowntime() int64 {
	if o == nil || o.TotalDowntime == nil {
		var ret int64
		return ret
	}
	return *o.TotalDowntime
}

// GetTotalDowntimeOk returns a tuple with the TotalDowntime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLA) GetTotalDowntimeOk() (*int64, bool) {
	if o == nil || o.TotalDowntime == nil {
		return nil, false
	}
	return o.TotalDowntime, true
}

// HasTotalDowntime returns a boolean if a field has been set.
func (o *SLA) HasTotalDowntime() bool {
	return o != nil && o.TotalDowntime != nil
}

// SetTotalDowntime gets a reference to the given int64 and assigns it to the TotalDowntime field.
func (o *SLA) SetTotalDowntime(v int64) {
	o.TotalDowntime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SLA) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Sla != nil {
		toSerialize["sla"] = o.Sla
	}
	if o.TotalDuration != nil {
		toSerialize["totalDuration"] = o.TotalDuration
	}
	if o.TotalDowntime != nil {
		toSerialize["totalDowntime"] = o.TotalDowntime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SLA) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cluster       *ClusterListItem `json:"cluster,omitempty"`
		Detail        []OutageRecord   `json:"detail,omitempty"`
		Sla           *float64         `json:"sla,omitempty"`
		TotalDuration *int64           `json:"totalDuration,omitempty"`
		TotalDowntime *int64           `json:"totalDowntime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cluster", "detail", "sla", "totalDuration", "totalDowntime"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Cluster != nil && all.Cluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cluster = all.Cluster
	o.Detail = all.Detail
	o.Sla = all.Sla
	o.TotalDuration = all.TotalDuration
	o.TotalDowntime = all.TotalDowntime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
