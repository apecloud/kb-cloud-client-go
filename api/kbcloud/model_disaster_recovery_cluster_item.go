// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DisasterRecoveryClusterItem DisasterRecovery cluster information
type DisasterRecoveryClusterItem struct {
	// Logical instance ID
	LogicalInstanceId *string `json:"logicalInstanceID,omitempty"`
	// KubeBlocks cluster information
	Cluster     *ClusterListItem `json:"cluster,omitempty"`
	TaskSummary *TaskSummary     `json:"taskSummary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDisasterRecoveryClusterItem instantiates a new DisasterRecoveryClusterItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDisasterRecoveryClusterItem() *DisasterRecoveryClusterItem {
	this := DisasterRecoveryClusterItem{}
	return &this
}

// NewDisasterRecoveryClusterItemWithDefaults instantiates a new DisasterRecoveryClusterItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDisasterRecoveryClusterItemWithDefaults() *DisasterRecoveryClusterItem {
	this := DisasterRecoveryClusterItem{}
	return &this
}

// GetLogicalInstanceId returns the LogicalInstanceId field value if set, zero value otherwise.
func (o *DisasterRecoveryClusterItem) GetLogicalInstanceId() string {
	if o == nil || o.LogicalInstanceId == nil {
		var ret string
		return ret
	}
	return *o.LogicalInstanceId
}

// GetLogicalInstanceIdOk returns a tuple with the LogicalInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryClusterItem) GetLogicalInstanceIdOk() (*string, bool) {
	if o == nil || o.LogicalInstanceId == nil {
		return nil, false
	}
	return o.LogicalInstanceId, true
}

// HasLogicalInstanceId returns a boolean if a field has been set.
func (o *DisasterRecoveryClusterItem) HasLogicalInstanceId() bool {
	return o != nil && o.LogicalInstanceId != nil
}

// SetLogicalInstanceId gets a reference to the given string and assigns it to the LogicalInstanceId field.
func (o *DisasterRecoveryClusterItem) SetLogicalInstanceId(v string) {
	o.LogicalInstanceId = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *DisasterRecoveryClusterItem) GetCluster() ClusterListItem {
	if o == nil || o.Cluster == nil {
		var ret ClusterListItem
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryClusterItem) GetClusterOk() (*ClusterListItem, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *DisasterRecoveryClusterItem) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given ClusterListItem and assigns it to the Cluster field.
func (o *DisasterRecoveryClusterItem) SetCluster(v ClusterListItem) {
	o.Cluster = &v
}

// GetTaskSummary returns the TaskSummary field value if set, zero value otherwise.
func (o *DisasterRecoveryClusterItem) GetTaskSummary() TaskSummary {
	if o == nil || o.TaskSummary == nil {
		var ret TaskSummary
		return ret
	}
	return *o.TaskSummary
}

// GetTaskSummaryOk returns a tuple with the TaskSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisasterRecoveryClusterItem) GetTaskSummaryOk() (*TaskSummary, bool) {
	if o == nil || o.TaskSummary == nil {
		return nil, false
	}
	return o.TaskSummary, true
}

// HasTaskSummary returns a boolean if a field has been set.
func (o *DisasterRecoveryClusterItem) HasTaskSummary() bool {
	return o != nil && o.TaskSummary != nil
}

// SetTaskSummary gets a reference to the given TaskSummary and assigns it to the TaskSummary field.
func (o *DisasterRecoveryClusterItem) SetTaskSummary(v TaskSummary) {
	o.TaskSummary = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DisasterRecoveryClusterItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.LogicalInstanceId != nil {
		toSerialize["logicalInstanceID"] = o.LogicalInstanceId
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.TaskSummary != nil {
		toSerialize["taskSummary"] = o.TaskSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DisasterRecoveryClusterItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LogicalInstanceId *string          `json:"logicalInstanceID,omitempty"`
		Cluster           *ClusterListItem `json:"cluster,omitempty"`
		TaskSummary       *TaskSummary     `json:"taskSummary,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"logicalInstanceID", "cluster", "taskSummary"})
	} else {
		return err
	}

	hasInvalidField := false
	o.LogicalInstanceId = all.LogicalInstanceId
	if all.Cluster != nil && all.Cluster.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cluster = all.Cluster
	if all.TaskSummary != nil && all.TaskSummary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TaskSummary = all.TaskSummary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
