// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// SchedulingConfig Configuration of resource scheduling for this environment
type SchedulingConfig struct {
	// * `HardAntiAffinity` - Strictly enforce pod anti-affinity across nodes. Pods will not be scheduled when the anti-affinity constraints cannot be satisfied.
	// * `SoftAntiAffinity` - Prefer to spread pods across nodes using anti-affinity, but allow scheduling on the same node when constraints cannot be fully satisfied.
	// * `Disabled` - Do not apply pod anti-affinity constraints on nodes.
	// * `None` - Inherit the engine-level default scheduling policy.
	//
	ClusterSchedulingPolicy *ClusterSchedulingPolicy `json:"clusterSchedulingPolicy,omitempty"`
	// When creating a cluster, add the default tolerations from the bootstrap node to the pods
	TolerateDefaultTaints *TolerateDefaultTaints `json:"tolerateDefaultTaints,omitempty"`
	// Scheduler used by internal system component clusters, including MinIO, VictoriaMetrics, and Victoria Logs. The value is persisted on the environment. Defaults to the Kubernetes default scheduler regardless of whether Koordinator is enabled.
	SystemComponentSchedulerName            *string                              `json:"systemComponentSchedulerName,omitempty"`
	SystemComponentReservationResourceClass *KoordinatorReservationResourceClass `json:"systemComponentReservationResourceClass,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSchedulingConfig instantiates a new SchedulingConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSchedulingConfig() *SchedulingConfig {
	this := SchedulingConfig{}
	var clusterSchedulingPolicy ClusterSchedulingPolicy = ClusterSchedulingPolicySoftAntiAffinity
	this.ClusterSchedulingPolicy = &clusterSchedulingPolicy
	var systemComponentSchedulerName string = "default-scheduler"
	this.SystemComponentSchedulerName = &systemComponentSchedulerName
	return &this
}

// NewSchedulingConfigWithDefaults instantiates a new SchedulingConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSchedulingConfigWithDefaults() *SchedulingConfig {
	this := SchedulingConfig{}
	var clusterSchedulingPolicy ClusterSchedulingPolicy = ClusterSchedulingPolicySoftAntiAffinity
	this.ClusterSchedulingPolicy = &clusterSchedulingPolicy
	var systemComponentSchedulerName string = "default-scheduler"
	this.SystemComponentSchedulerName = &systemComponentSchedulerName
	return &this
}

// GetClusterSchedulingPolicy returns the ClusterSchedulingPolicy field value if set, zero value otherwise.
func (o *SchedulingConfig) GetClusterSchedulingPolicy() ClusterSchedulingPolicy {
	if o == nil || o.ClusterSchedulingPolicy == nil {
		var ret ClusterSchedulingPolicy
		return ret
	}
	return *o.ClusterSchedulingPolicy
}

// GetClusterSchedulingPolicyOk returns a tuple with the ClusterSchedulingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulingConfig) GetClusterSchedulingPolicyOk() (*ClusterSchedulingPolicy, bool) {
	if o == nil || o.ClusterSchedulingPolicy == nil {
		return nil, false
	}
	return o.ClusterSchedulingPolicy, true
}

// HasClusterSchedulingPolicy returns a boolean if a field has been set.
func (o *SchedulingConfig) HasClusterSchedulingPolicy() bool {
	return o != nil && o.ClusterSchedulingPolicy != nil
}

// SetClusterSchedulingPolicy gets a reference to the given ClusterSchedulingPolicy and assigns it to the ClusterSchedulingPolicy field.
func (o *SchedulingConfig) SetClusterSchedulingPolicy(v ClusterSchedulingPolicy) {
	o.ClusterSchedulingPolicy = &v
}

// GetTolerateDefaultTaints returns the TolerateDefaultTaints field value if set, zero value otherwise.
func (o *SchedulingConfig) GetTolerateDefaultTaints() TolerateDefaultTaints {
	if o == nil || o.TolerateDefaultTaints == nil {
		var ret TolerateDefaultTaints
		return ret
	}
	return *o.TolerateDefaultTaints
}

// GetTolerateDefaultTaintsOk returns a tuple with the TolerateDefaultTaints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulingConfig) GetTolerateDefaultTaintsOk() (*TolerateDefaultTaints, bool) {
	if o == nil || o.TolerateDefaultTaints == nil {
		return nil, false
	}
	return o.TolerateDefaultTaints, true
}

// HasTolerateDefaultTaints returns a boolean if a field has been set.
func (o *SchedulingConfig) HasTolerateDefaultTaints() bool {
	return o != nil && o.TolerateDefaultTaints != nil
}

// SetTolerateDefaultTaints gets a reference to the given TolerateDefaultTaints and assigns it to the TolerateDefaultTaints field.
func (o *SchedulingConfig) SetTolerateDefaultTaints(v TolerateDefaultTaints) {
	o.TolerateDefaultTaints = &v
}

// GetSystemComponentSchedulerName returns the SystemComponentSchedulerName field value if set, zero value otherwise.
func (o *SchedulingConfig) GetSystemComponentSchedulerName() string {
	if o == nil || o.SystemComponentSchedulerName == nil {
		var ret string
		return ret
	}
	return *o.SystemComponentSchedulerName
}

// GetSystemComponentSchedulerNameOk returns a tuple with the SystemComponentSchedulerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulingConfig) GetSystemComponentSchedulerNameOk() (*string, bool) {
	if o == nil || o.SystemComponentSchedulerName == nil {
		return nil, false
	}
	return o.SystemComponentSchedulerName, true
}

// HasSystemComponentSchedulerName returns a boolean if a field has been set.
func (o *SchedulingConfig) HasSystemComponentSchedulerName() bool {
	return o != nil && o.SystemComponentSchedulerName != nil
}

// SetSystemComponentSchedulerName gets a reference to the given string and assigns it to the SystemComponentSchedulerName field.
func (o *SchedulingConfig) SetSystemComponentSchedulerName(v string) {
	o.SystemComponentSchedulerName = &v
}

// GetSystemComponentReservationResourceClass returns the SystemComponentReservationResourceClass field value if set, zero value otherwise.
func (o *SchedulingConfig) GetSystemComponentReservationResourceClass() KoordinatorReservationResourceClass {
	if o == nil || o.SystemComponentReservationResourceClass == nil {
		var ret KoordinatorReservationResourceClass
		return ret
	}
	return *o.SystemComponentReservationResourceClass
}

// GetSystemComponentReservationResourceClassOk returns a tuple with the SystemComponentReservationResourceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulingConfig) GetSystemComponentReservationResourceClassOk() (*KoordinatorReservationResourceClass, bool) {
	if o == nil || o.SystemComponentReservationResourceClass == nil {
		return nil, false
	}
	return o.SystemComponentReservationResourceClass, true
}

// HasSystemComponentReservationResourceClass returns a boolean if a field has been set.
func (o *SchedulingConfig) HasSystemComponentReservationResourceClass() bool {
	return o != nil && o.SystemComponentReservationResourceClass != nil
}

// SetSystemComponentReservationResourceClass gets a reference to the given KoordinatorReservationResourceClass and assigns it to the SystemComponentReservationResourceClass field.
func (o *SchedulingConfig) SetSystemComponentReservationResourceClass(v KoordinatorReservationResourceClass) {
	o.SystemComponentReservationResourceClass = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SchedulingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ClusterSchedulingPolicy != nil {
		toSerialize["clusterSchedulingPolicy"] = o.ClusterSchedulingPolicy
	}
	if o.TolerateDefaultTaints != nil {
		toSerialize["tolerateDefaultTaints"] = o.TolerateDefaultTaints
	}
	if o.SystemComponentSchedulerName != nil {
		toSerialize["systemComponentSchedulerName"] = o.SystemComponentSchedulerName
	}
	if o.SystemComponentReservationResourceClass != nil {
		toSerialize["systemComponentReservationResourceClass"] = o.SystemComponentReservationResourceClass
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SchedulingConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClusterSchedulingPolicy                 *ClusterSchedulingPolicy             `json:"clusterSchedulingPolicy,omitempty"`
		TolerateDefaultTaints                   *TolerateDefaultTaints               `json:"tolerateDefaultTaints,omitempty"`
		SystemComponentSchedulerName            *string                              `json:"systemComponentSchedulerName,omitempty"`
		SystemComponentReservationResourceClass *KoordinatorReservationResourceClass `json:"systemComponentReservationResourceClass,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"clusterSchedulingPolicy", "tolerateDefaultTaints", "systemComponentSchedulerName", "systemComponentReservationResourceClass"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ClusterSchedulingPolicy != nil && !all.ClusterSchedulingPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.ClusterSchedulingPolicy = all.ClusterSchedulingPolicy
	}
	if all.TolerateDefaultTaints != nil && all.TolerateDefaultTaints.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TolerateDefaultTaints = all.TolerateDefaultTaints
	o.SystemComponentSchedulerName = all.SystemComponentSchedulerName
	if all.SystemComponentReservationResourceClass != nil && !all.SystemComponentReservationResourceClass.IsValid() {
		hasInvalidField = true
	} else {
		o.SystemComponentReservationResourceClass = all.SystemComponentReservationResourceClass
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
