// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// SchedulingConfig Configuration of resource scheduling for this environment
type SchedulingConfig struct {
	// Enable pod antiaffinity for cluster
	PodAntiAffinityEnabled *bool `json:"podAntiAffinityEnabled,omitempty"`
	// * `HardAntiAffinity` - Strictly enforced; pods will not be scheduled if constraints cannot be met.
	// * `SoftAntiAffinity` - Best-effort; the scheduler prefers to satisfy constraints but may place pods together if necessary.
	// * `Disabled` - No anti-affinity constraints applied.
	//
	EngineSchedulingStrategy *EngineSchedulingStrategy `json:"engineSchedulingStrategy,omitempty"`
	// When creating a cluster, add the default tolerations from the bootstrap node to the pods
	TolerateDefaultTaints *TolerateDefaultTaints `json:"tolerateDefaultTaints,omitempty"`
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
	var podAntiAffinityEnabled bool = true
	this.PodAntiAffinityEnabled = &podAntiAffinityEnabled
	var engineSchedulingStrategy EngineSchedulingStrategy = EngineSchedulingStrategyDisabled
	this.EngineSchedulingStrategy = &engineSchedulingStrategy
	return &this
}

// NewSchedulingConfigWithDefaults instantiates a new SchedulingConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSchedulingConfigWithDefaults() *SchedulingConfig {
	this := SchedulingConfig{}
	var podAntiAffinityEnabled bool = true
	this.PodAntiAffinityEnabled = &podAntiAffinityEnabled
	var engineSchedulingStrategy EngineSchedulingStrategy = EngineSchedulingStrategyDisabled
	this.EngineSchedulingStrategy = &engineSchedulingStrategy
	return &this
}

// GetPodAntiAffinityEnabled returns the PodAntiAffinityEnabled field value if set, zero value otherwise.
func (o *SchedulingConfig) GetPodAntiAffinityEnabled() bool {
	if o == nil || o.PodAntiAffinityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.PodAntiAffinityEnabled
}

// GetPodAntiAffinityEnabledOk returns a tuple with the PodAntiAffinityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulingConfig) GetPodAntiAffinityEnabledOk() (*bool, bool) {
	if o == nil || o.PodAntiAffinityEnabled == nil {
		return nil, false
	}
	return o.PodAntiAffinityEnabled, true
}

// HasPodAntiAffinityEnabled returns a boolean if a field has been set.
func (o *SchedulingConfig) HasPodAntiAffinityEnabled() bool {
	return o != nil && o.PodAntiAffinityEnabled != nil
}

// SetPodAntiAffinityEnabled gets a reference to the given bool and assigns it to the PodAntiAffinityEnabled field.
func (o *SchedulingConfig) SetPodAntiAffinityEnabled(v bool) {
	o.PodAntiAffinityEnabled = &v
}

// GetEngineSchedulingStrategy returns the EngineSchedulingStrategy field value if set, zero value otherwise.
func (o *SchedulingConfig) GetEngineSchedulingStrategy() EngineSchedulingStrategy {
	if o == nil || o.EngineSchedulingStrategy == nil {
		var ret EngineSchedulingStrategy
		return ret
	}
	return *o.EngineSchedulingStrategy
}

// GetEngineSchedulingStrategyOk returns a tuple with the EngineSchedulingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulingConfig) GetEngineSchedulingStrategyOk() (*EngineSchedulingStrategy, bool) {
	if o == nil || o.EngineSchedulingStrategy == nil {
		return nil, false
	}
	return o.EngineSchedulingStrategy, true
}

// HasEngineSchedulingStrategy returns a boolean if a field has been set.
func (o *SchedulingConfig) HasEngineSchedulingStrategy() bool {
	return o != nil && o.EngineSchedulingStrategy != nil
}

// SetEngineSchedulingStrategy gets a reference to the given EngineSchedulingStrategy and assigns it to the EngineSchedulingStrategy field.
func (o *SchedulingConfig) SetEngineSchedulingStrategy(v EngineSchedulingStrategy) {
	o.EngineSchedulingStrategy = &v
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

// MarshalJSON serializes the struct using spec logic.
func (o SchedulingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.PodAntiAffinityEnabled != nil {
		toSerialize["podAntiAffinityEnabled"] = o.PodAntiAffinityEnabled
	}
	if o.EngineSchedulingStrategy != nil {
		toSerialize["engineSchedulingStrategy"] = o.EngineSchedulingStrategy
	}
	if o.TolerateDefaultTaints != nil {
		toSerialize["tolerateDefaultTaints"] = o.TolerateDefaultTaints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SchedulingConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PodAntiAffinityEnabled   *bool                     `json:"podAntiAffinityEnabled,omitempty"`
		EngineSchedulingStrategy *EngineSchedulingStrategy `json:"engineSchedulingStrategy,omitempty"`
		TolerateDefaultTaints    *TolerateDefaultTaints    `json:"tolerateDefaultTaints,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"podAntiAffinityEnabled", "engineSchedulingStrategy", "tolerateDefaultTaints"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PodAntiAffinityEnabled = all.PodAntiAffinityEnabled
	if all.EngineSchedulingStrategy != nil && !all.EngineSchedulingStrategy.IsValid() {
		hasInvalidField = true
	} else {
		o.EngineSchedulingStrategy = all.EngineSchedulingStrategy
	}
	if all.TolerateDefaultTaints != nil && all.TolerateDefaultTaints.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TolerateDefaultTaints = all.TolerateDefaultTaints

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
