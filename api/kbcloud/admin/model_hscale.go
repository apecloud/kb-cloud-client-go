// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Hscale struct {
	// the amount of pods
	Pods int32 `json:"pods"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHscale instantiates a new Hscale object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHscale(pods int32) *Hscale {
	this := Hscale{}
	this.Pods = pods
	return &this
}

// NewHscaleWithDefaults instantiates a new Hscale object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHscaleWithDefaults() *Hscale {
	this := Hscale{}
	return &this
}

// GetPods returns the Pods field value.
func (o *Hscale) GetPods() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Pods
}

// GetPodsOk returns a tuple with the Pods field value
// and a boolean to check if the value has been set.
func (o *Hscale) GetPodsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pods, true
}

// SetPods sets field value.
func (o *Hscale) SetPods(v int32) {
	o.Pods = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Hscale) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["pods"] = o.Pods

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Hscale) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Pods *int32 `json:"pods"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Pods == nil {
		return fmt.Errorf("required field pods missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"pods"})
	} else {
		return err
	}
	o.Pods = *all.Pods

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
