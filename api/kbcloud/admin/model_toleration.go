// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Toleration Toleration allows a pod to be scheduled on nodes with taints.
type Toleration struct {
	// The taint key that the toleration applies to.
	Key string `json:"key"`
	// The operator to use for matching the taint's key and value.
	Operator TolerationOperator `json:"operator"`
	// The taint value to match. Optional if operator is "Exists".
	Value *string `json:"value,omitempty"`
	// The effect of the taint that this toleration tolerates.
	Effect TolerationEffect `json:"effect"`
	// The duration for which the toleration is valid.
	TolerationSeconds *int32 `json:"tolerationSeconds,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewToleration instantiates a new Toleration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewToleration(key string, operator TolerationOperator, effect TolerationEffect) *Toleration {
	this := Toleration{}
	this.Key = key
	this.Operator = operator
	this.Effect = effect
	return &this
}

// NewTolerationWithDefaults instantiates a new Toleration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTolerationWithDefaults() *Toleration {
	this := Toleration{}
	return &this
}

// GetKey returns the Key field value.
func (o *Toleration) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Toleration) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *Toleration) SetKey(v string) {
	o.Key = v
}

// GetOperator returns the Operator field value.
func (o *Toleration) GetOperator() TolerationOperator {
	if o == nil {
		var ret TolerationOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *Toleration) GetOperatorOk() (*TolerationOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *Toleration) SetOperator(v TolerationOperator) {
	o.Operator = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Toleration) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toleration) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Toleration) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Toleration) SetValue(v string) {
	o.Value = &v
}

// GetEffect returns the Effect field value.
func (o *Toleration) GetEffect() TolerationEffect {
	if o == nil {
		var ret TolerationEffect
		return ret
	}
	return o.Effect
}

// GetEffectOk returns a tuple with the Effect field value
// and a boolean to check if the value has been set.
func (o *Toleration) GetEffectOk() (*TolerationEffect, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Effect, true
}

// SetEffect sets field value.
func (o *Toleration) SetEffect(v TolerationEffect) {
	o.Effect = v
}

// GetTolerationSeconds returns the TolerationSeconds field value if set, zero value otherwise.
func (o *Toleration) GetTolerationSeconds() int32 {
	if o == nil || o.TolerationSeconds == nil {
		var ret int32
		return ret
	}
	return *o.TolerationSeconds
}

// GetTolerationSecondsOk returns a tuple with the TolerationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Toleration) GetTolerationSecondsOk() (*int32, bool) {
	if o == nil || o.TolerationSeconds == nil {
		return nil, false
	}
	return o.TolerationSeconds, true
}

// HasTolerationSeconds returns a boolean if a field has been set.
func (o *Toleration) HasTolerationSeconds() bool {
	return o != nil && o.TolerationSeconds != nil
}

// SetTolerationSeconds gets a reference to the given int32 and assigns it to the TolerationSeconds field.
func (o *Toleration) SetTolerationSeconds(v int32) {
	o.TolerationSeconds = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Toleration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	toSerialize["operator"] = o.Operator
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	toSerialize["effect"] = o.Effect
	if o.TolerationSeconds != nil {
		toSerialize["tolerationSeconds"] = o.TolerationSeconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Toleration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key               *string             `json:"key"`
		Operator          *TolerationOperator `json:"operator"`
		Value             *string             `json:"value,omitempty"`
		Effect            *TolerationEffect   `json:"effect"`
		TolerationSeconds *int32              `json:"tolerationSeconds,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Effect == nil {
		return fmt.Errorf("required field effect missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key", "operator", "value", "effect", "tolerationSeconds"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Key = *all.Key
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	o.Value = all.Value
	if !all.Effect.IsValid() {
		hasInvalidField = true
	} else {
		o.Effect = *all.Effect
	}
	o.TolerationSeconds = all.TolerationSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
