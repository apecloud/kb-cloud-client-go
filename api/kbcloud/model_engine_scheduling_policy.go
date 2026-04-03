// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineSchedulingPolicy Engine-level default scheduling policy for a specific engine and engine mode.
type EngineSchedulingPolicy struct {
	// Id of engine scheduling policy.
	Id *string `json:"id,omitempty"`
	// * `HardAntiAffinity` - Strictly enforce pod anti-affinity across nodes. Pods will not be scheduled when the anti-affinity constraints cannot be satisfied.
	// * `SoftAntiAffinity` - Prefer to spread pods across nodes using anti-affinity, but allow scheduling on the same node when constraints cannot be fully satisfied.
	// * `Disabled` - Do not apply pod anti-affinity constraints on nodes.
	//
	SchedulingPolicyType SchedulingPolicyType `json:"schedulingPolicyType"`
	// Database engine that this policy applies to. Used together with `engineMode`.
	Engine string `json:"engine"`
	// Engine mode that this policy applies to, such as `standalone` or `replication`.
	EngineMode string `json:"engineMode"`
	// Human-readable description of the engine scheduling policy.
	Description *string `json:"description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineSchedulingPolicy instantiates a new EngineSchedulingPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineSchedulingPolicy(schedulingPolicyType SchedulingPolicyType, engine string, engineMode string) *EngineSchedulingPolicy {
	this := EngineSchedulingPolicy{}
	this.SchedulingPolicyType = schedulingPolicyType
	this.Engine = engine
	this.EngineMode = engineMode
	return &this
}

// NewEngineSchedulingPolicyWithDefaults instantiates a new EngineSchedulingPolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineSchedulingPolicyWithDefaults() *EngineSchedulingPolicy {
	this := EngineSchedulingPolicy{}
	var schedulingPolicyType SchedulingPolicyType = SchedulingPolicyTypeDisabled
	this.SchedulingPolicyType = schedulingPolicyType
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EngineSchedulingPolicy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineSchedulingPolicy) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EngineSchedulingPolicy) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EngineSchedulingPolicy) SetId(v string) {
	o.Id = &v
}

// GetSchedulingPolicyType returns the SchedulingPolicyType field value.
func (o *EngineSchedulingPolicy) GetSchedulingPolicyType() SchedulingPolicyType {
	if o == nil {
		var ret SchedulingPolicyType
		return ret
	}
	return o.SchedulingPolicyType
}

// GetSchedulingPolicyTypeOk returns a tuple with the SchedulingPolicyType field value
// and a boolean to check if the value has been set.
func (o *EngineSchedulingPolicy) GetSchedulingPolicyTypeOk() (*SchedulingPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchedulingPolicyType, true
}

// SetSchedulingPolicyType sets field value.
func (o *EngineSchedulingPolicy) SetSchedulingPolicyType(v SchedulingPolicyType) {
	o.SchedulingPolicyType = v
}

// GetEngine returns the Engine field value.
func (o *EngineSchedulingPolicy) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *EngineSchedulingPolicy) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *EngineSchedulingPolicy) SetEngine(v string) {
	o.Engine = v
}

// GetEngineMode returns the EngineMode field value.
func (o *EngineSchedulingPolicy) GetEngineMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineMode
}

// GetEngineModeOk returns a tuple with the EngineMode field value
// and a boolean to check if the value has been set.
func (o *EngineSchedulingPolicy) GetEngineModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineMode, true
}

// SetEngineMode sets field value.
func (o *EngineSchedulingPolicy) SetEngineMode(v string) {
	o.EngineMode = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EngineSchedulingPolicy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineSchedulingPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EngineSchedulingPolicy) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EngineSchedulingPolicy) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineSchedulingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["schedulingPolicyType"] = o.SchedulingPolicyType
	toSerialize["engine"] = o.Engine
	toSerialize["engineMode"] = o.EngineMode
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineSchedulingPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                   *string               `json:"id,omitempty"`
		SchedulingPolicyType *SchedulingPolicyType `json:"schedulingPolicyType"`
		Engine               *string               `json:"engine"`
		EngineMode           *string               `json:"engineMode"`
		Description          *string               `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SchedulingPolicyType == nil {
		return fmt.Errorf("required field schedulingPolicyType missing")
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.EngineMode == nil {
		return fmt.Errorf("required field engineMode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "schedulingPolicyType", "engine", "engineMode", "description"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	if !all.SchedulingPolicyType.IsValid() {
		hasInvalidField = true
	} else {
		o.SchedulingPolicyType = *all.SchedulingPolicyType
	}
	o.Engine = *all.Engine
	o.EngineMode = *all.EngineMode
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
