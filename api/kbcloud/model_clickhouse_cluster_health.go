// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClickhouseClusterHealth struct {
	// Collection time in RFC3339 UTC.
	CollectedAt string `json:"collectedAt"`
	// Collection status. success when every source was collected, partial when topology or replicas could only be partially read.
	Status string `json:"status"`
	// Derived health status. normal, warning or critical.
	OverallStatus *string `json:"overallStatus,omitempty"`
	// Detected deployment mode (standalone or cluster).
	ClusterMode *string                         `json:"clusterMode,omitempty"`
	Summary     *ClickhouseClusterHealthSummary `json:"summary,omitempty"`
	// Replicated tables ordered by risk, absolute delay and queue size.
	Replicas []ClickhouseReplicaHealth `json:"replicas"`
	// Shard and replica topology from system.clusters on the connected node.
	Nodes []ClickhouseClusterNode `json:"nodes"`
	// Per-source collection status.
	Sources []ClickhouseSource `json:"sources"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseClusterHealth instantiates a new ClickhouseClusterHealth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseClusterHealth(collectedAt string, status string, replicas []ClickhouseReplicaHealth, nodes []ClickhouseClusterNode, sources []ClickhouseSource) *ClickhouseClusterHealth {
	this := ClickhouseClusterHealth{}
	this.CollectedAt = collectedAt
	this.Status = status
	this.Replicas = replicas
	this.Nodes = nodes
	this.Sources = sources
	return &this
}

// NewClickhouseClusterHealthWithDefaults instantiates a new ClickhouseClusterHealth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseClusterHealthWithDefaults() *ClickhouseClusterHealth {
	this := ClickhouseClusterHealth{}
	return &this
}

// GetCollectedAt returns the CollectedAt field value.
func (o *ClickhouseClusterHealth) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealth) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *ClickhouseClusterHealth) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// GetStatus returns the Status field value.
func (o *ClickhouseClusterHealth) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealth) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ClickhouseClusterHealth) SetStatus(v string) {
	o.Status = v
}

// GetOverallStatus returns the OverallStatus field value if set, zero value otherwise.
func (o *ClickhouseClusterHealth) GetOverallStatus() string {
	if o == nil || o.OverallStatus == nil {
		var ret string
		return ret
	}
	return *o.OverallStatus
}

// GetOverallStatusOk returns a tuple with the OverallStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealth) GetOverallStatusOk() (*string, bool) {
	if o == nil || o.OverallStatus == nil {
		return nil, false
	}
	return o.OverallStatus, true
}

// HasOverallStatus returns a boolean if a field has been set.
func (o *ClickhouseClusterHealth) HasOverallStatus() bool {
	return o != nil && o.OverallStatus != nil
}

// SetOverallStatus gets a reference to the given string and assigns it to the OverallStatus field.
func (o *ClickhouseClusterHealth) SetOverallStatus(v string) {
	o.OverallStatus = &v
}

// GetClusterMode returns the ClusterMode field value if set, zero value otherwise.
func (o *ClickhouseClusterHealth) GetClusterMode() string {
	if o == nil || o.ClusterMode == nil {
		var ret string
		return ret
	}
	return *o.ClusterMode
}

// GetClusterModeOk returns a tuple with the ClusterMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealth) GetClusterModeOk() (*string, bool) {
	if o == nil || o.ClusterMode == nil {
		return nil, false
	}
	return o.ClusterMode, true
}

// HasClusterMode returns a boolean if a field has been set.
func (o *ClickhouseClusterHealth) HasClusterMode() bool {
	return o != nil && o.ClusterMode != nil
}

// SetClusterMode gets a reference to the given string and assigns it to the ClusterMode field.
func (o *ClickhouseClusterHealth) SetClusterMode(v string) {
	o.ClusterMode = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ClickhouseClusterHealth) GetSummary() ClickhouseClusterHealthSummary {
	if o == nil || o.Summary == nil {
		var ret ClickhouseClusterHealthSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealth) GetSummaryOk() (*ClickhouseClusterHealthSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ClickhouseClusterHealth) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given ClickhouseClusterHealthSummary and assigns it to the Summary field.
func (o *ClickhouseClusterHealth) SetSummary(v ClickhouseClusterHealthSummary) {
	o.Summary = &v
}

// GetReplicas returns the Replicas field value.
func (o *ClickhouseClusterHealth) GetReplicas() []ClickhouseReplicaHealth {
	if o == nil {
		var ret []ClickhouseReplicaHealth
		return ret
	}
	return o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealth) GetReplicasOk() (*[]ClickhouseReplicaHealth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replicas, true
}

// SetReplicas sets field value.
func (o *ClickhouseClusterHealth) SetReplicas(v []ClickhouseReplicaHealth) {
	o.Replicas = v
}

// GetNodes returns the Nodes field value.
func (o *ClickhouseClusterHealth) GetNodes() []ClickhouseClusterNode {
	if o == nil {
		var ret []ClickhouseClusterNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealth) GetNodesOk() (*[]ClickhouseClusterNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value.
func (o *ClickhouseClusterHealth) SetNodes(v []ClickhouseClusterNode) {
	o.Nodes = v
}

// GetSources returns the Sources field value.
func (o *ClickhouseClusterHealth) GetSources() []ClickhouseSource {
	if o == nil {
		var ret []ClickhouseSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealth) GetSourcesOk() (*[]ClickhouseSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *ClickhouseClusterHealth) SetSources(v []ClickhouseSource) {
	o.Sources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseClusterHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["collectedAt"] = o.CollectedAt
	toSerialize["status"] = o.Status
	if o.OverallStatus != nil {
		toSerialize["overallStatus"] = o.OverallStatus
	}
	if o.ClusterMode != nil {
		toSerialize["clusterMode"] = o.ClusterMode
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	toSerialize["replicas"] = o.Replicas
	toSerialize["nodes"] = o.Nodes
	toSerialize["sources"] = o.Sources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseClusterHealth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CollectedAt   *string                         `json:"collectedAt"`
		Status        *string                         `json:"status"`
		OverallStatus *string                         `json:"overallStatus,omitempty"`
		ClusterMode   *string                         `json:"clusterMode,omitempty"`
		Summary       *ClickhouseClusterHealthSummary `json:"summary,omitempty"`
		Replicas      *[]ClickhouseReplicaHealth      `json:"replicas"`
		Nodes         *[]ClickhouseClusterNode        `json:"nodes"`
		Sources       *[]ClickhouseSource             `json:"sources"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Replicas == nil {
		return fmt.Errorf("required field replicas missing")
	}
	if all.Nodes == nil {
		return fmt.Errorf("required field nodes missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"collectedAt", "status", "overallStatus", "clusterMode", "summary", "replicas", "nodes", "sources"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CollectedAt = *all.CollectedAt
	o.Status = *all.Status
	o.OverallStatus = all.OverallStatus
	o.ClusterMode = all.ClusterMode
	if all.Summary != nil && all.Summary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Summary = all.Summary
	o.Replicas = *all.Replicas
	o.Nodes = *all.Nodes
	o.Sources = *all.Sources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
