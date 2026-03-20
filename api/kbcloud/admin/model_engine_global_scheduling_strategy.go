// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EngineGlobalSchedulingStrategy Default scheduling strategy applied to all rules unless overridden.
type EngineGlobalSchedulingStrategy struct {
	// * `HardAntiAffinity` - Strictly enforced; pods will not be scheduled if constraints cannot be met.
	// * `SoftAntiAffinity` - Best-effort; the scheduler prefers to satisfy constraints but may place pods together if necessary.
	// * `Disabled` - No anti-affinity constraints applied.
	//
	GlobalSchedulingStrategy NullableEngineSchedulingStrategy `json:"globalSchedulingStrategy,omitempty"`
	// Database engine this global strategy applies to. Used to match strategies to clusters by engine.
	Engine *string `json:"engine,omitempty"`
	// Indicates if this global strategy is the default strategy applied when no specific rules match.
	Default *bool `json:"default,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineGlobalSchedulingStrategy instantiates a new EngineGlobalSchedulingStrategy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineGlobalSchedulingStrategy() *EngineGlobalSchedulingStrategy {
	this := EngineGlobalSchedulingStrategy{}
	var globalSchedulingStrategy EngineSchedulingStrategy = EngineSchedulingStrategyDisabled
	this.GlobalSchedulingStrategy = *NewNullableEngineSchedulingStrategy(&globalSchedulingStrategy)
	return &this
}

// NewEngineGlobalSchedulingStrategyWithDefaults instantiates a new EngineGlobalSchedulingStrategy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineGlobalSchedulingStrategyWithDefaults() *EngineGlobalSchedulingStrategy {
	this := EngineGlobalSchedulingStrategy{}
	var globalSchedulingStrategy EngineSchedulingStrategy = EngineSchedulingStrategyDisabled
	this.GlobalSchedulingStrategy = *NewNullableEngineSchedulingStrategy(&globalSchedulingStrategy)
	return &this
}

// GetGlobalSchedulingStrategy returns the GlobalSchedulingStrategy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineGlobalSchedulingStrategy) GetGlobalSchedulingStrategy() EngineSchedulingStrategy {
	if o == nil || o.GlobalSchedulingStrategy.Get() == nil {
		var ret EngineSchedulingStrategy
		return ret
	}
	return *o.GlobalSchedulingStrategy.Get()
}

// GetGlobalSchedulingStrategyOk returns a tuple with the GlobalSchedulingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineGlobalSchedulingStrategy) GetGlobalSchedulingStrategyOk() (*EngineSchedulingStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalSchedulingStrategy.Get(), o.GlobalSchedulingStrategy.IsSet()
}

// HasGlobalSchedulingStrategy returns a boolean if a field has been set.
func (o *EngineGlobalSchedulingStrategy) HasGlobalSchedulingStrategy() bool {
	return o != nil && o.GlobalSchedulingStrategy.IsSet()
}

// SetGlobalSchedulingStrategy gets a reference to the given NullableEngineSchedulingStrategy and assigns it to the GlobalSchedulingStrategy field.
func (o *EngineGlobalSchedulingStrategy) SetGlobalSchedulingStrategy(v EngineSchedulingStrategy) {
	o.GlobalSchedulingStrategy.Set(&v)
}

// SetGlobalSchedulingStrategyNil sets the value for GlobalSchedulingStrategy to be an explicit nil.
func (o *EngineGlobalSchedulingStrategy) SetGlobalSchedulingStrategyNil() {
	o.GlobalSchedulingStrategy.Set(nil)
}

// UnsetGlobalSchedulingStrategy ensures that no value is present for GlobalSchedulingStrategy, not even an explicit nil.
func (o *EngineGlobalSchedulingStrategy) UnsetGlobalSchedulingStrategy() {
	o.GlobalSchedulingStrategy.Unset()
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *EngineGlobalSchedulingStrategy) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineGlobalSchedulingStrategy) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *EngineGlobalSchedulingStrategy) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *EngineGlobalSchedulingStrategy) SetEngine(v string) {
	o.Engine = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *EngineGlobalSchedulingStrategy) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineGlobalSchedulingStrategy) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *EngineGlobalSchedulingStrategy) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *EngineGlobalSchedulingStrategy) SetDefault(v bool) {
	o.Default = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineGlobalSchedulingStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.GlobalSchedulingStrategy.IsSet() {
		toSerialize["globalSchedulingStrategy"] = o.GlobalSchedulingStrategy.Get()
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineGlobalSchedulingStrategy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GlobalSchedulingStrategy NullableEngineSchedulingStrategy `json:"globalSchedulingStrategy,omitempty"`
		Engine                   *string                          `json:"engine,omitempty"`
		Default                  *bool                            `json:"default,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"globalSchedulingStrategy", "engine", "default"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.GlobalSchedulingStrategy.Get() != nil && !all.GlobalSchedulingStrategy.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.GlobalSchedulingStrategy = all.GlobalSchedulingStrategy
	}
	o.Engine = all.Engine
	o.Default = all.Default

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
