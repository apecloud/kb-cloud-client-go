// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ClickhouseClusterHealthSummary struct {
	// Number of replicated table replicas collected.
	ReplicaCount *int64 `json:"replicaCount,omitempty"`
	// Number of nodes in system.clusters.
	NodeCount *int64 `json:"nodeCount,omitempty"`
	// Number of distinct cluster names.
	ClusterCount *int64 `json:"clusterCount,omitempty"`
	// Sum of total_replicas across replicated tables.
	TotalReplicas *int64 `json:"totalReplicas,omitempty"`
	// Sum of active_replicas across replicated tables.
	ActiveReplicas *int64 `json:"activeReplicas,omitempty"`
	// Replicas currently in readonly mode.
	ReadonlyReplicas *int64 `json:"readonlyReplicas,omitempty"`
	// Replicas whose Keeper session expired.
	SessionExpiredReplicas *int64 `json:"sessionExpiredReplicas,omitempty"`
	// Maximum absolute_delay in seconds.
	MaxAbsoluteDelaySeconds *float64 `json:"maxAbsoluteDelaySeconds,omitempty"`
	// Sum of replication queue sizes.
	TotalQueueSize *int64 `json:"totalQueueSize,omitempty"`
	// Replicas classified as warning.
	WarningReplicas *int64 `json:"warningReplicas,omitempty"`
	// Replicas classified as critical.
	CriticalReplicas *int64 `json:"criticalReplicas,omitempty"`
	// Nodes with connection errors, slowdowns or a non-zero estimated recovery time.
	UnhealthyNodes *int64 `json:"unhealthyNodes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClickhouseClusterHealthSummary instantiates a new ClickhouseClusterHealthSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClickhouseClusterHealthSummary() *ClickhouseClusterHealthSummary {
	this := ClickhouseClusterHealthSummary{}
	return &this
}

// NewClickhouseClusterHealthSummaryWithDefaults instantiates a new ClickhouseClusterHealthSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClickhouseClusterHealthSummaryWithDefaults() *ClickhouseClusterHealthSummary {
	this := ClickhouseClusterHealthSummary{}
	return &this
}

// GetReplicaCount returns the ReplicaCount field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetReplicaCount() int64 {
	if o == nil || o.ReplicaCount == nil {
		var ret int64
		return ret
	}
	return *o.ReplicaCount
}

// GetReplicaCountOk returns a tuple with the ReplicaCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetReplicaCountOk() (*int64, bool) {
	if o == nil || o.ReplicaCount == nil {
		return nil, false
	}
	return o.ReplicaCount, true
}

// HasReplicaCount returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasReplicaCount() bool {
	return o != nil && o.ReplicaCount != nil
}

// SetReplicaCount gets a reference to the given int64 and assigns it to the ReplicaCount field.
func (o *ClickhouseClusterHealthSummary) SetReplicaCount(v int64) {
	o.ReplicaCount = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetNodeCount() int64 {
	if o == nil || o.NodeCount == nil {
		var ret int64
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetNodeCountOk() (*int64, bool) {
	if o == nil || o.NodeCount == nil {
		return nil, false
	}
	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasNodeCount() bool {
	return o != nil && o.NodeCount != nil
}

// SetNodeCount gets a reference to the given int64 and assigns it to the NodeCount field.
func (o *ClickhouseClusterHealthSummary) SetNodeCount(v int64) {
	o.NodeCount = &v
}

// GetClusterCount returns the ClusterCount field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetClusterCount() int64 {
	if o == nil || o.ClusterCount == nil {
		var ret int64
		return ret
	}
	return *o.ClusterCount
}

// GetClusterCountOk returns a tuple with the ClusterCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetClusterCountOk() (*int64, bool) {
	if o == nil || o.ClusterCount == nil {
		return nil, false
	}
	return o.ClusterCount, true
}

// HasClusterCount returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasClusterCount() bool {
	return o != nil && o.ClusterCount != nil
}

// SetClusterCount gets a reference to the given int64 and assigns it to the ClusterCount field.
func (o *ClickhouseClusterHealthSummary) SetClusterCount(v int64) {
	o.ClusterCount = &v
}

// GetTotalReplicas returns the TotalReplicas field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetTotalReplicas() int64 {
	if o == nil || o.TotalReplicas == nil {
		var ret int64
		return ret
	}
	return *o.TotalReplicas
}

// GetTotalReplicasOk returns a tuple with the TotalReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetTotalReplicasOk() (*int64, bool) {
	if o == nil || o.TotalReplicas == nil {
		return nil, false
	}
	return o.TotalReplicas, true
}

// HasTotalReplicas returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasTotalReplicas() bool {
	return o != nil && o.TotalReplicas != nil
}

// SetTotalReplicas gets a reference to the given int64 and assigns it to the TotalReplicas field.
func (o *ClickhouseClusterHealthSummary) SetTotalReplicas(v int64) {
	o.TotalReplicas = &v
}

// GetActiveReplicas returns the ActiveReplicas field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetActiveReplicas() int64 {
	if o == nil || o.ActiveReplicas == nil {
		var ret int64
		return ret
	}
	return *o.ActiveReplicas
}

// GetActiveReplicasOk returns a tuple with the ActiveReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetActiveReplicasOk() (*int64, bool) {
	if o == nil || o.ActiveReplicas == nil {
		return nil, false
	}
	return o.ActiveReplicas, true
}

// HasActiveReplicas returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasActiveReplicas() bool {
	return o != nil && o.ActiveReplicas != nil
}

// SetActiveReplicas gets a reference to the given int64 and assigns it to the ActiveReplicas field.
func (o *ClickhouseClusterHealthSummary) SetActiveReplicas(v int64) {
	o.ActiveReplicas = &v
}

// GetReadonlyReplicas returns the ReadonlyReplicas field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetReadonlyReplicas() int64 {
	if o == nil || o.ReadonlyReplicas == nil {
		var ret int64
		return ret
	}
	return *o.ReadonlyReplicas
}

// GetReadonlyReplicasOk returns a tuple with the ReadonlyReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetReadonlyReplicasOk() (*int64, bool) {
	if o == nil || o.ReadonlyReplicas == nil {
		return nil, false
	}
	return o.ReadonlyReplicas, true
}

// HasReadonlyReplicas returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasReadonlyReplicas() bool {
	return o != nil && o.ReadonlyReplicas != nil
}

// SetReadonlyReplicas gets a reference to the given int64 and assigns it to the ReadonlyReplicas field.
func (o *ClickhouseClusterHealthSummary) SetReadonlyReplicas(v int64) {
	o.ReadonlyReplicas = &v
}

// GetSessionExpiredReplicas returns the SessionExpiredReplicas field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetSessionExpiredReplicas() int64 {
	if o == nil || o.SessionExpiredReplicas == nil {
		var ret int64
		return ret
	}
	return *o.SessionExpiredReplicas
}

// GetSessionExpiredReplicasOk returns a tuple with the SessionExpiredReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetSessionExpiredReplicasOk() (*int64, bool) {
	if o == nil || o.SessionExpiredReplicas == nil {
		return nil, false
	}
	return o.SessionExpiredReplicas, true
}

// HasSessionExpiredReplicas returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasSessionExpiredReplicas() bool {
	return o != nil && o.SessionExpiredReplicas != nil
}

// SetSessionExpiredReplicas gets a reference to the given int64 and assigns it to the SessionExpiredReplicas field.
func (o *ClickhouseClusterHealthSummary) SetSessionExpiredReplicas(v int64) {
	o.SessionExpiredReplicas = &v
}

// GetMaxAbsoluteDelaySeconds returns the MaxAbsoluteDelaySeconds field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetMaxAbsoluteDelaySeconds() float64 {
	if o == nil || o.MaxAbsoluteDelaySeconds == nil {
		var ret float64
		return ret
	}
	return *o.MaxAbsoluteDelaySeconds
}

// GetMaxAbsoluteDelaySecondsOk returns a tuple with the MaxAbsoluteDelaySeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetMaxAbsoluteDelaySecondsOk() (*float64, bool) {
	if o == nil || o.MaxAbsoluteDelaySeconds == nil {
		return nil, false
	}
	return o.MaxAbsoluteDelaySeconds, true
}

// HasMaxAbsoluteDelaySeconds returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasMaxAbsoluteDelaySeconds() bool {
	return o != nil && o.MaxAbsoluteDelaySeconds != nil
}

// SetMaxAbsoluteDelaySeconds gets a reference to the given float64 and assigns it to the MaxAbsoluteDelaySeconds field.
func (o *ClickhouseClusterHealthSummary) SetMaxAbsoluteDelaySeconds(v float64) {
	o.MaxAbsoluteDelaySeconds = &v
}

// GetTotalQueueSize returns the TotalQueueSize field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetTotalQueueSize() int64 {
	if o == nil || o.TotalQueueSize == nil {
		var ret int64
		return ret
	}
	return *o.TotalQueueSize
}

// GetTotalQueueSizeOk returns a tuple with the TotalQueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetTotalQueueSizeOk() (*int64, bool) {
	if o == nil || o.TotalQueueSize == nil {
		return nil, false
	}
	return o.TotalQueueSize, true
}

// HasTotalQueueSize returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasTotalQueueSize() bool {
	return o != nil && o.TotalQueueSize != nil
}

// SetTotalQueueSize gets a reference to the given int64 and assigns it to the TotalQueueSize field.
func (o *ClickhouseClusterHealthSummary) SetTotalQueueSize(v int64) {
	o.TotalQueueSize = &v
}

// GetWarningReplicas returns the WarningReplicas field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetWarningReplicas() int64 {
	if o == nil || o.WarningReplicas == nil {
		var ret int64
		return ret
	}
	return *o.WarningReplicas
}

// GetWarningReplicasOk returns a tuple with the WarningReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetWarningReplicasOk() (*int64, bool) {
	if o == nil || o.WarningReplicas == nil {
		return nil, false
	}
	return o.WarningReplicas, true
}

// HasWarningReplicas returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasWarningReplicas() bool {
	return o != nil && o.WarningReplicas != nil
}

// SetWarningReplicas gets a reference to the given int64 and assigns it to the WarningReplicas field.
func (o *ClickhouseClusterHealthSummary) SetWarningReplicas(v int64) {
	o.WarningReplicas = &v
}

// GetCriticalReplicas returns the CriticalReplicas field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetCriticalReplicas() int64 {
	if o == nil || o.CriticalReplicas == nil {
		var ret int64
		return ret
	}
	return *o.CriticalReplicas
}

// GetCriticalReplicasOk returns a tuple with the CriticalReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetCriticalReplicasOk() (*int64, bool) {
	if o == nil || o.CriticalReplicas == nil {
		return nil, false
	}
	return o.CriticalReplicas, true
}

// HasCriticalReplicas returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasCriticalReplicas() bool {
	return o != nil && o.CriticalReplicas != nil
}

// SetCriticalReplicas gets a reference to the given int64 and assigns it to the CriticalReplicas field.
func (o *ClickhouseClusterHealthSummary) SetCriticalReplicas(v int64) {
	o.CriticalReplicas = &v
}

// GetUnhealthyNodes returns the UnhealthyNodes field value if set, zero value otherwise.
func (o *ClickhouseClusterHealthSummary) GetUnhealthyNodes() int64 {
	if o == nil || o.UnhealthyNodes == nil {
		var ret int64
		return ret
	}
	return *o.UnhealthyNodes
}

// GetUnhealthyNodesOk returns a tuple with the UnhealthyNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClickhouseClusterHealthSummary) GetUnhealthyNodesOk() (*int64, bool) {
	if o == nil || o.UnhealthyNodes == nil {
		return nil, false
	}
	return o.UnhealthyNodes, true
}

// HasUnhealthyNodes returns a boolean if a field has been set.
func (o *ClickhouseClusterHealthSummary) HasUnhealthyNodes() bool {
	return o != nil && o.UnhealthyNodes != nil
}

// SetUnhealthyNodes gets a reference to the given int64 and assigns it to the UnhealthyNodes field.
func (o *ClickhouseClusterHealthSummary) SetUnhealthyNodes(v int64) {
	o.UnhealthyNodes = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClickhouseClusterHealthSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ReplicaCount != nil {
		toSerialize["replicaCount"] = o.ReplicaCount
	}
	if o.NodeCount != nil {
		toSerialize["nodeCount"] = o.NodeCount
	}
	if o.ClusterCount != nil {
		toSerialize["clusterCount"] = o.ClusterCount
	}
	if o.TotalReplicas != nil {
		toSerialize["totalReplicas"] = o.TotalReplicas
	}
	if o.ActiveReplicas != nil {
		toSerialize["activeReplicas"] = o.ActiveReplicas
	}
	if o.ReadonlyReplicas != nil {
		toSerialize["readonlyReplicas"] = o.ReadonlyReplicas
	}
	if o.SessionExpiredReplicas != nil {
		toSerialize["sessionExpiredReplicas"] = o.SessionExpiredReplicas
	}
	if o.MaxAbsoluteDelaySeconds != nil {
		toSerialize["maxAbsoluteDelaySeconds"] = o.MaxAbsoluteDelaySeconds
	}
	if o.TotalQueueSize != nil {
		toSerialize["totalQueueSize"] = o.TotalQueueSize
	}
	if o.WarningReplicas != nil {
		toSerialize["warningReplicas"] = o.WarningReplicas
	}
	if o.CriticalReplicas != nil {
		toSerialize["criticalReplicas"] = o.CriticalReplicas
	}
	if o.UnhealthyNodes != nil {
		toSerialize["unhealthyNodes"] = o.UnhealthyNodes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClickhouseClusterHealthSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ReplicaCount            *int64   `json:"replicaCount,omitempty"`
		NodeCount               *int64   `json:"nodeCount,omitempty"`
		ClusterCount            *int64   `json:"clusterCount,omitempty"`
		TotalReplicas           *int64   `json:"totalReplicas,omitempty"`
		ActiveReplicas          *int64   `json:"activeReplicas,omitempty"`
		ReadonlyReplicas        *int64   `json:"readonlyReplicas,omitempty"`
		SessionExpiredReplicas  *int64   `json:"sessionExpiredReplicas,omitempty"`
		MaxAbsoluteDelaySeconds *float64 `json:"maxAbsoluteDelaySeconds,omitempty"`
		TotalQueueSize          *int64   `json:"totalQueueSize,omitempty"`
		WarningReplicas         *int64   `json:"warningReplicas,omitempty"`
		CriticalReplicas        *int64   `json:"criticalReplicas,omitempty"`
		UnhealthyNodes          *int64   `json:"unhealthyNodes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"replicaCount", "nodeCount", "clusterCount", "totalReplicas", "activeReplicas", "readonlyReplicas", "sessionExpiredReplicas", "maxAbsoluteDelaySeconds", "totalQueueSize", "warningReplicas", "criticalReplicas", "unhealthyNodes"})
	} else {
		return err
	}
	o.ReplicaCount = all.ReplicaCount
	o.NodeCount = all.NodeCount
	o.ClusterCount = all.ClusterCount
	o.TotalReplicas = all.TotalReplicas
	o.ActiveReplicas = all.ActiveReplicas
	o.ReadonlyReplicas = all.ReadonlyReplicas
	o.SessionExpiredReplicas = all.SessionExpiredReplicas
	o.MaxAbsoluteDelaySeconds = all.MaxAbsoluteDelaySeconds
	o.TotalQueueSize = all.TotalQueueSize
	o.WarningReplicas = all.WarningReplicas
	o.CriticalReplicas = all.CriticalReplicas
	o.UnhealthyNodes = all.UnhealthyNodes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
