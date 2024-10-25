// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterUpdate ClusterUpdate is the payload to update a KubeBlocks cluster

type ClusterUpdate struct {
	// The number of replicas, for standalone mode, the replicas is 1, for raftGroup mode, the default replicas is 3.
	Replicas *int32 `json:"replicas,omitempty"`
	// CPU cores.
	Cpu *float64 `json:"cpu,omitempty"`
	// Memory, the unit is Gi.
	Memory *float64 `json:"memory,omitempty"`
	// Storage size, the unit is Gi.
	Storage *float64 `json:"storage,omitempty"`
	// The termination policy of cluster.
	TerminationPolicy *ClusterTerminationPolicy `json:"terminationPolicy,omitempty"`
	// Specify whether the cluster enable monitoring.
	MonitorEnabled *bool `json:"monitorEnabled,omitempty"`
	// Specify whether the cluster can be accessed from within the VPC.
	VpcEndpointEnabled *bool `json:"vpcEndpointEnabled,omitempty"`
	// Specify whether the cluster can be accessed from the public internet.
	InternetEndpointEnabled *bool `json:"internetEndpointEnabled,omitempty"`
	// Items is the list of parameter template in the list
	ParamTpls []ParamTplsItem `json:"paramTpls,omitempty"`
	// Tolerations of cluster
	Tolerations *string `json:"tolerations,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterUpdate instantiates a new ClusterUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterUpdate() *ClusterUpdate {
	this := ClusterUpdate{}
	var terminationPolicy ClusterTerminationPolicy = CLUSTERTERMINATIONPOLICY_DELETE
	this.TerminationPolicy = &terminationPolicy
	return &this
}

// NewClusterUpdateWithDefaults instantiates a new ClusterUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterUpdateWithDefaults() *ClusterUpdate {
	this := ClusterUpdate{}
	var terminationPolicy ClusterTerminationPolicy = CLUSTERTERMINATIONPOLICY_DELETE
	this.TerminationPolicy = &terminationPolicy
	return &this
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ClusterUpdate) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ClusterUpdate) HasReplicas() bool {
	return o != nil && o.Replicas != nil
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *ClusterUpdate) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ClusterUpdate) GetCpu() float64 {
	if o == nil || o.Cpu == nil {
		var ret float64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetCpuOk() (*float64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ClusterUpdate) HasCpu() bool {
	return o != nil && o.Cpu != nil
}

// SetCpu gets a reference to the given float64 and assigns it to the Cpu field.
func (o *ClusterUpdate) SetCpu(v float64) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ClusterUpdate) GetMemory() float64 {
	if o == nil || o.Memory == nil {
		var ret float64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetMemoryOk() (*float64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ClusterUpdate) HasMemory() bool {
	return o != nil && o.Memory != nil
}

// SetMemory gets a reference to the given float64 and assigns it to the Memory field.
func (o *ClusterUpdate) SetMemory(v float64) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ClusterUpdate) GetStorage() float64 {
	if o == nil || o.Storage == nil {
		var ret float64
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetStorageOk() (*float64, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ClusterUpdate) HasStorage() bool {
	return o != nil && o.Storage != nil
}

// SetStorage gets a reference to the given float64 and assigns it to the Storage field.
func (o *ClusterUpdate) SetStorage(v float64) {
	o.Storage = &v
}

// GetTerminationPolicy returns the TerminationPolicy field value if set, zero value otherwise.
func (o *ClusterUpdate) GetTerminationPolicy() ClusterTerminationPolicy {
	if o == nil || o.TerminationPolicy == nil {
		var ret ClusterTerminationPolicy
		return ret
	}
	return *o.TerminationPolicy
}

// GetTerminationPolicyOk returns a tuple with the TerminationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetTerminationPolicyOk() (*ClusterTerminationPolicy, bool) {
	if o == nil || o.TerminationPolicy == nil {
		return nil, false
	}
	return o.TerminationPolicy, true
}

// HasTerminationPolicy returns a boolean if a field has been set.
func (o *ClusterUpdate) HasTerminationPolicy() bool {
	return o != nil && o.TerminationPolicy != nil
}

// SetTerminationPolicy gets a reference to the given ClusterTerminationPolicy and assigns it to the TerminationPolicy field.
func (o *ClusterUpdate) SetTerminationPolicy(v ClusterTerminationPolicy) {
	o.TerminationPolicy = &v
}

// GetMonitorEnabled returns the MonitorEnabled field value if set, zero value otherwise.
func (o *ClusterUpdate) GetMonitorEnabled() bool {
	if o == nil || o.MonitorEnabled == nil {
		var ret bool
		return ret
	}
	return *o.MonitorEnabled
}

// GetMonitorEnabledOk returns a tuple with the MonitorEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetMonitorEnabledOk() (*bool, bool) {
	if o == nil || o.MonitorEnabled == nil {
		return nil, false
	}
	return o.MonitorEnabled, true
}

// HasMonitorEnabled returns a boolean if a field has been set.
func (o *ClusterUpdate) HasMonitorEnabled() bool {
	return o != nil && o.MonitorEnabled != nil
}

// SetMonitorEnabled gets a reference to the given bool and assigns it to the MonitorEnabled field.
func (o *ClusterUpdate) SetMonitorEnabled(v bool) {
	o.MonitorEnabled = &v
}

// GetVpcEndpointEnabled returns the VpcEndpointEnabled field value if set, zero value otherwise.
func (o *ClusterUpdate) GetVpcEndpointEnabled() bool {
	if o == nil || o.VpcEndpointEnabled == nil {
		var ret bool
		return ret
	}
	return *o.VpcEndpointEnabled
}

// GetVpcEndpointEnabledOk returns a tuple with the VpcEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetVpcEndpointEnabledOk() (*bool, bool) {
	if o == nil || o.VpcEndpointEnabled == nil {
		return nil, false
	}
	return o.VpcEndpointEnabled, true
}

// HasVpcEndpointEnabled returns a boolean if a field has been set.
func (o *ClusterUpdate) HasVpcEndpointEnabled() bool {
	return o != nil && o.VpcEndpointEnabled != nil
}

// SetVpcEndpointEnabled gets a reference to the given bool and assigns it to the VpcEndpointEnabled field.
func (o *ClusterUpdate) SetVpcEndpointEnabled(v bool) {
	o.VpcEndpointEnabled = &v
}

// GetInternetEndpointEnabled returns the InternetEndpointEnabled field value if set, zero value otherwise.
func (o *ClusterUpdate) GetInternetEndpointEnabled() bool {
	if o == nil || o.InternetEndpointEnabled == nil {
		var ret bool
		return ret
	}
	return *o.InternetEndpointEnabled
}

// GetInternetEndpointEnabledOk returns a tuple with the InternetEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetInternetEndpointEnabledOk() (*bool, bool) {
	if o == nil || o.InternetEndpointEnabled == nil {
		return nil, false
	}
	return o.InternetEndpointEnabled, true
}

// HasInternetEndpointEnabled returns a boolean if a field has been set.
func (o *ClusterUpdate) HasInternetEndpointEnabled() bool {
	return o != nil && o.InternetEndpointEnabled != nil
}

// SetInternetEndpointEnabled gets a reference to the given bool and assigns it to the InternetEndpointEnabled field.
func (o *ClusterUpdate) SetInternetEndpointEnabled(v bool) {
	o.InternetEndpointEnabled = &v
}

// GetParamTpls returns the ParamTpls field value if set, zero value otherwise.
func (o *ClusterUpdate) GetParamTpls() []ParamTplsItem {
	if o == nil || o.ParamTpls == nil {
		var ret []ParamTplsItem
		return ret
	}
	return o.ParamTpls
}

// GetParamTplsOk returns a tuple with the ParamTpls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetParamTplsOk() (*[]ParamTplsItem, bool) {
	if o == nil || o.ParamTpls == nil {
		return nil, false
	}
	return &o.ParamTpls, true
}

// HasParamTpls returns a boolean if a field has been set.
func (o *ClusterUpdate) HasParamTpls() bool {
	return o != nil && o.ParamTpls != nil
}

// SetParamTpls gets a reference to the given []ParamTplsItem and assigns it to the ParamTpls field.
func (o *ClusterUpdate) SetParamTpls(v []ParamTplsItem) {
	o.ParamTpls = v
}

// GetTolerations returns the Tolerations field value if set, zero value otherwise.
func (o *ClusterUpdate) GetTolerations() string {
	if o == nil || o.Tolerations == nil {
		var ret string
		return ret
	}
	return *o.Tolerations
}

// GetTolerationsOk returns a tuple with the Tolerations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetTolerationsOk() (*string, bool) {
	if o == nil || o.Tolerations == nil {
		return nil, false
	}
	return o.Tolerations, true
}

// HasTolerations returns a boolean if a field has been set.
func (o *ClusterUpdate) HasTolerations() bool {
	return o != nil && o.Tolerations != nil
}

// SetTolerations gets a reference to the given string and assigns it to the Tolerations field.
func (o *ClusterUpdate) SetTolerations(v string) {
	o.Tolerations = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.TerminationPolicy != nil {
		toSerialize["terminationPolicy"] = o.TerminationPolicy
	}
	if o.MonitorEnabled != nil {
		toSerialize["monitorEnabled"] = o.MonitorEnabled
	}
	if o.VpcEndpointEnabled != nil {
		toSerialize["vpcEndpointEnabled"] = o.VpcEndpointEnabled
	}
	if o.InternetEndpointEnabled != nil {
		toSerialize["internetEndpointEnabled"] = o.InternetEndpointEnabled
	}
	if o.ParamTpls != nil {
		toSerialize["paramTpls"] = o.ParamTpls
	}
	if o.Tolerations != nil {
		toSerialize["tolerations"] = o.Tolerations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Replicas                *int32                    `json:"replicas,omitempty"`
		Cpu                     *float64                  `json:"cpu,omitempty"`
		Memory                  *float64                  `json:"memory,omitempty"`
		Storage                 *float64                  `json:"storage,omitempty"`
		TerminationPolicy       *ClusterTerminationPolicy `json:"terminationPolicy,omitempty"`
		MonitorEnabled          *bool                     `json:"monitorEnabled,omitempty"`
		VpcEndpointEnabled      *bool                     `json:"vpcEndpointEnabled,omitempty"`
		InternetEndpointEnabled *bool                     `json:"internetEndpointEnabled,omitempty"`
		ParamTpls               []ParamTplsItem           `json:"paramTpls,omitempty"`
		Tolerations             *string                   `json:"tolerations,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"replicas", "cpu", "memory", "storage", "terminationPolicy", "monitorEnabled", "vpcEndpointEnabled", "internetEndpointEnabled", "paramTpls", "tolerations"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Replicas = all.Replicas
	o.Cpu = all.Cpu
	o.Memory = all.Memory
	o.Storage = all.Storage
	if all.TerminationPolicy != nil && !all.TerminationPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.TerminationPolicy = all.TerminationPolicy
	}
	o.MonitorEnabled = all.MonitorEnabled
	o.VpcEndpointEnabled = all.VpcEndpointEnabled
	o.InternetEndpointEnabled = all.InternetEndpointEnabled
	o.ParamTpls = all.ParamTpls
	o.Tolerations = all.Tolerations

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
