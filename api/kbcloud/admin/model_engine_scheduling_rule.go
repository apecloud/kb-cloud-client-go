// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineSchedulingRule struct {
	// Id of engine scheduling rule.
	Id *string `json:"id,omitempty"`
	// Rule name, used to reference this rule when applying scheduling policies to a cluster.
	Name string `json:"name"`
	// Database engine type this rule applies to. Used to match rules to clusters by engine.
	Engine string `json:"engine"`
	// Cluster mode this rule applies to. Used to match rules to clusters by mode.
	EngineMode string `json:"engineMode"`
	// Selects the two mutually exclusive targets that this rule applies to. The array length must be exactly 2.
	TargetSelector []Target `json:"targetSelector"`
	// * `HardAntiAffinity` - Strictly enforce pod anti-affinity across nodes. Pods will not be scheduled when the anti-affinity constraints cannot be satisfied.
	// * `SoftAntiAffinity` - Prefer to spread pods across nodes using anti-affinity, but allow scheduling on the same node when constraints cannot be fully satisfied.
	// * `Disabled` - Do not apply pod anti-affinity constraints on nodes.
	//
	SchedulingPolicyType SchedulingPolicyType `json:"schedulingPolicyType"`
	// Optional description of the scheduling rule.
	Description *string `json:"description,omitempty"`
	// Indicates whether this rule is used as the fallback when no other rule matches.
	Default common.NullableBool `json:"default,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineSchedulingRule instantiates a new EngineSchedulingRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineSchedulingRule(name string, engine string, engineMode string, targetSelector []Target, schedulingPolicyType SchedulingPolicyType) *EngineSchedulingRule {
	this := EngineSchedulingRule{}
	this.Name = name
	this.Engine = engine
	this.EngineMode = engineMode
	this.TargetSelector = targetSelector
	this.SchedulingPolicyType = schedulingPolicyType
	return &this
}

// NewEngineSchedulingRuleWithDefaults instantiates a new EngineSchedulingRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineSchedulingRuleWithDefaults() *EngineSchedulingRule {
	this := EngineSchedulingRule{}
	var schedulingPolicyType SchedulingPolicyType = SchedulingPolicyTypeDisabled
	this.SchedulingPolicyType = schedulingPolicyType
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EngineSchedulingRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineSchedulingRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EngineSchedulingRule) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EngineSchedulingRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value.
func (o *EngineSchedulingRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EngineSchedulingRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EngineSchedulingRule) SetName(v string) {
	o.Name = v
}

// GetEngine returns the Engine field value.
func (o *EngineSchedulingRule) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *EngineSchedulingRule) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *EngineSchedulingRule) SetEngine(v string) {
	o.Engine = v
}

// GetEngineMode returns the EngineMode field value.
func (o *EngineSchedulingRule) GetEngineMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineMode
}

// GetEngineModeOk returns a tuple with the EngineMode field value
// and a boolean to check if the value has been set.
func (o *EngineSchedulingRule) GetEngineModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineMode, true
}

// SetEngineMode sets field value.
func (o *EngineSchedulingRule) SetEngineMode(v string) {
	o.EngineMode = v
}

// GetTargetSelector returns the TargetSelector field value.
func (o *EngineSchedulingRule) GetTargetSelector() []Target {
	if o == nil {
		var ret []Target
		return ret
	}
	return o.TargetSelector
}

// GetTargetSelectorOk returns a tuple with the TargetSelector field value
// and a boolean to check if the value has been set.
func (o *EngineSchedulingRule) GetTargetSelectorOk() (*[]Target, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetSelector, true
}

// SetTargetSelector sets field value.
func (o *EngineSchedulingRule) SetTargetSelector(v []Target) {
	o.TargetSelector = v
}

// GetSchedulingPolicyType returns the SchedulingPolicyType field value.
func (o *EngineSchedulingRule) GetSchedulingPolicyType() SchedulingPolicyType {
	if o == nil {
		var ret SchedulingPolicyType
		return ret
	}
	return o.SchedulingPolicyType
}

// GetSchedulingPolicyTypeOk returns a tuple with the SchedulingPolicyType field value
// and a boolean to check if the value has been set.
func (o *EngineSchedulingRule) GetSchedulingPolicyTypeOk() (*SchedulingPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchedulingPolicyType, true
}

// SetSchedulingPolicyType sets field value.
func (o *EngineSchedulingRule) SetSchedulingPolicyType(v SchedulingPolicyType) {
	o.SchedulingPolicyType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EngineSchedulingRule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineSchedulingRule) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EngineSchedulingRule) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EngineSchedulingRule) SetDescription(v string) {
	o.Description = &v
}

// GetDefault returns the Default field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineSchedulingRule) GetDefault() bool {
	if o == nil || o.Default.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Default.Get()
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineSchedulingRule) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Default.Get(), o.Default.IsSet()
}

// HasDefault returns a boolean if a field has been set.
func (o *EngineSchedulingRule) HasDefault() bool {
	return o != nil && o.Default.IsSet()
}

// SetDefault gets a reference to the given common.NullableBool and assigns it to the Default field.
func (o *EngineSchedulingRule) SetDefault(v bool) {
	o.Default.Set(&v)
}

// SetDefaultNil sets the value for Default to be an explicit nil.
func (o *EngineSchedulingRule) SetDefaultNil() {
	o.Default.Set(nil)
}

// UnsetDefault ensures that no value is present for Default, not even an explicit nil.
func (o *EngineSchedulingRule) UnsetDefault() {
	o.Default.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineSchedulingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["engine"] = o.Engine
	toSerialize["engineMode"] = o.EngineMode
	toSerialize["targetSelector"] = o.TargetSelector
	toSerialize["schedulingPolicyType"] = o.SchedulingPolicyType
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Default.IsSet() {
		toSerialize["default"] = o.Default.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineSchedulingRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                   *string               `json:"id,omitempty"`
		Name                 *string               `json:"name"`
		Engine               *string               `json:"engine"`
		EngineMode           *string               `json:"engineMode"`
		TargetSelector       *[]Target             `json:"targetSelector"`
		SchedulingPolicyType *SchedulingPolicyType `json:"schedulingPolicyType"`
		Description          *string               `json:"description,omitempty"`
		Default              common.NullableBool   `json:"default,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.EngineMode == nil {
		return fmt.Errorf("required field engineMode missing")
	}
	if all.TargetSelector == nil {
		return fmt.Errorf("required field targetSelector missing")
	}
	if all.SchedulingPolicyType == nil {
		return fmt.Errorf("required field schedulingPolicyType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "engine", "engineMode", "targetSelector", "schedulingPolicyType", "description", "default"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = *all.Name
	o.Engine = *all.Engine
	o.EngineMode = *all.EngineMode
	o.TargetSelector = *all.TargetSelector
	if !all.SchedulingPolicyType.IsValid() {
		hasInvalidField = true
	} else {
		o.SchedulingPolicyType = *all.SchedulingPolicyType
	}
	o.Description = all.Description
	o.Default = all.Default

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
