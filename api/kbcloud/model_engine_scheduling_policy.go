// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineSchedulingPolicy Defines scheduling constraints for cluster pods, including affinity, anti-affinity, and topology spread rules.
type EngineSchedulingPolicy struct {
	// Default scheduling strategy applied to all rules unless overridden.
	GlobalSchedulingStrategy *EngineGlobalSchedulingStrategy `json:"globalSchedulingStrategy,omitempty"`
	// Scheduling rules that define pod placement constraints scoped by engine, mode, and target.
	SchedulingRules []EngineSchedulingRule `json:"schedulingRules,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineSchedulingPolicy instantiates a new EngineSchedulingPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineSchedulingPolicy() *EngineSchedulingPolicy {
	this := EngineSchedulingPolicy{}
	return &this
}

// NewEngineSchedulingPolicyWithDefaults instantiates a new EngineSchedulingPolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineSchedulingPolicyWithDefaults() *EngineSchedulingPolicy {
	this := EngineSchedulingPolicy{}
	return &this
}

// GetGlobalSchedulingStrategy returns the GlobalSchedulingStrategy field value if set, zero value otherwise.
func (o *EngineSchedulingPolicy) GetGlobalSchedulingStrategy() EngineGlobalSchedulingStrategy {
	if o == nil || o.GlobalSchedulingStrategy == nil {
		var ret EngineGlobalSchedulingStrategy
		return ret
	}
	return *o.GlobalSchedulingStrategy
}

// GetGlobalSchedulingStrategyOk returns a tuple with the GlobalSchedulingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineSchedulingPolicy) GetGlobalSchedulingStrategyOk() (*EngineGlobalSchedulingStrategy, bool) {
	if o == nil || o.GlobalSchedulingStrategy == nil {
		return nil, false
	}
	return o.GlobalSchedulingStrategy, true
}

// HasGlobalSchedulingStrategy returns a boolean if a field has been set.
func (o *EngineSchedulingPolicy) HasGlobalSchedulingStrategy() bool {
	return o != nil && o.GlobalSchedulingStrategy != nil
}

// SetGlobalSchedulingStrategy gets a reference to the given EngineGlobalSchedulingStrategy and assigns it to the GlobalSchedulingStrategy field.
func (o *EngineSchedulingPolicy) SetGlobalSchedulingStrategy(v EngineGlobalSchedulingStrategy) {
	o.GlobalSchedulingStrategy = &v
}

// GetSchedulingRules returns the SchedulingRules field value if set, zero value otherwise.
func (o *EngineSchedulingPolicy) GetSchedulingRules() []EngineSchedulingRule {
	if o == nil || o.SchedulingRules == nil {
		var ret []EngineSchedulingRule
		return ret
	}
	return o.SchedulingRules
}

// GetSchedulingRulesOk returns a tuple with the SchedulingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineSchedulingPolicy) GetSchedulingRulesOk() (*[]EngineSchedulingRule, bool) {
	if o == nil || o.SchedulingRules == nil {
		return nil, false
	}
	return &o.SchedulingRules, true
}

// HasSchedulingRules returns a boolean if a field has been set.
func (o *EngineSchedulingPolicy) HasSchedulingRules() bool {
	return o != nil && o.SchedulingRules != nil
}

// SetSchedulingRules gets a reference to the given []EngineSchedulingRule and assigns it to the SchedulingRules field.
func (o *EngineSchedulingPolicy) SetSchedulingRules(v []EngineSchedulingRule) {
	o.SchedulingRules = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineSchedulingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.GlobalSchedulingStrategy != nil {
		toSerialize["globalSchedulingStrategy"] = o.GlobalSchedulingStrategy
	}
	if o.SchedulingRules != nil {
		toSerialize["schedulingRules"] = o.SchedulingRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineSchedulingPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GlobalSchedulingStrategy *EngineGlobalSchedulingStrategy `json:"globalSchedulingStrategy,omitempty"`
		SchedulingRules          []EngineSchedulingRule          `json:"schedulingRules,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"globalSchedulingStrategy", "schedulingRules"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.GlobalSchedulingStrategy != nil && all.GlobalSchedulingStrategy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GlobalSchedulingStrategy = all.GlobalSchedulingStrategy
	o.SchedulingRules = all.SchedulingRules

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
