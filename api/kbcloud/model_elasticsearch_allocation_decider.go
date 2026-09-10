// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchAllocationDecider struct {
	Decider     common.NullableString `json:"decider,omitempty"`
	Decision    common.NullableString `json:"decision,omitempty"`
	Explanation common.NullableString `json:"explanation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchAllocationDecider instantiates a new ElasticsearchAllocationDecider object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchAllocationDecider() *ElasticsearchAllocationDecider {
	this := ElasticsearchAllocationDecider{}
	return &this
}

// NewElasticsearchAllocationDeciderWithDefaults instantiates a new ElasticsearchAllocationDecider object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchAllocationDeciderWithDefaults() *ElasticsearchAllocationDecider {
	this := ElasticsearchAllocationDecider{}
	return &this
}

// GetDecider returns the Decider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationDecider) GetDecider() string {
	if o == nil || o.Decider.Get() == nil {
		var ret string
		return ret
	}
	return *o.Decider.Get()
}

// GetDeciderOk returns a tuple with the Decider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationDecider) GetDeciderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Decider.Get(), o.Decider.IsSet()
}

// HasDecider returns a boolean if a field has been set.
func (o *ElasticsearchAllocationDecider) HasDecider() bool {
	return o != nil && o.Decider.IsSet()
}

// SetDecider gets a reference to the given common.NullableString and assigns it to the Decider field.
func (o *ElasticsearchAllocationDecider) SetDecider(v string) {
	o.Decider.Set(&v)
}

// SetDeciderNil sets the value for Decider to be an explicit nil.
func (o *ElasticsearchAllocationDecider) SetDeciderNil() {
	o.Decider.Set(nil)
}

// UnsetDecider ensures that no value is present for Decider, not even an explicit nil.
func (o *ElasticsearchAllocationDecider) UnsetDecider() {
	o.Decider.Unset()
}

// GetDecision returns the Decision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationDecider) GetDecision() string {
	if o == nil || o.Decision.Get() == nil {
		var ret string
		return ret
	}
	return *o.Decision.Get()
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationDecider) GetDecisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Decision.Get(), o.Decision.IsSet()
}

// HasDecision returns a boolean if a field has been set.
func (o *ElasticsearchAllocationDecider) HasDecision() bool {
	return o != nil && o.Decision.IsSet()
}

// SetDecision gets a reference to the given common.NullableString and assigns it to the Decision field.
func (o *ElasticsearchAllocationDecider) SetDecision(v string) {
	o.Decision.Set(&v)
}

// SetDecisionNil sets the value for Decision to be an explicit nil.
func (o *ElasticsearchAllocationDecider) SetDecisionNil() {
	o.Decision.Set(nil)
}

// UnsetDecision ensures that no value is present for Decision, not even an explicit nil.
func (o *ElasticsearchAllocationDecider) UnsetDecision() {
	o.Decision.Unset()
}

// GetExplanation returns the Explanation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ElasticsearchAllocationDecider) GetExplanation() string {
	if o == nil || o.Explanation.Get() == nil {
		var ret string
		return ret
	}
	return *o.Explanation.Get()
}

// GetExplanationOk returns a tuple with the Explanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ElasticsearchAllocationDecider) GetExplanationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Explanation.Get(), o.Explanation.IsSet()
}

// HasExplanation returns a boolean if a field has been set.
func (o *ElasticsearchAllocationDecider) HasExplanation() bool {
	return o != nil && o.Explanation.IsSet()
}

// SetExplanation gets a reference to the given common.NullableString and assigns it to the Explanation field.
func (o *ElasticsearchAllocationDecider) SetExplanation(v string) {
	o.Explanation.Set(&v)
}

// SetExplanationNil sets the value for Explanation to be an explicit nil.
func (o *ElasticsearchAllocationDecider) SetExplanationNil() {
	o.Explanation.Set(nil)
}

// UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil.
func (o *ElasticsearchAllocationDecider) UnsetExplanation() {
	o.Explanation.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchAllocationDecider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Decider.IsSet() {
		toSerialize["decider"] = o.Decider.Get()
	}
	if o.Decision.IsSet() {
		toSerialize["decision"] = o.Decision.Get()
	}
	if o.Explanation.IsSet() {
		toSerialize["explanation"] = o.Explanation.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchAllocationDecider) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Decider     common.NullableString `json:"decider,omitempty"`
		Decision    common.NullableString `json:"decision,omitempty"`
		Explanation common.NullableString `json:"explanation,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"decider", "decision", "explanation"})
	} else {
		return err
	}
	o.Decider = all.Decider
	o.Decision = all.Decision
	o.Explanation = all.Explanation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
