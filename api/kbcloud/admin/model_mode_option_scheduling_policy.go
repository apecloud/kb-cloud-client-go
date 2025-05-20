// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type ModeOptionSchedulingPolicy struct {
	// when component names are specified in componentAntiAffinity, those components will be scheduled with anti-affinity rules applied to ensure they are spread across different nodes, especially when resource dispersion is enabled.
	//
	ComponentAntiAffinity []string `json:"componentAntiAffinity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeOptionSchedulingPolicy instantiates a new ModeOptionSchedulingPolicy object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeOptionSchedulingPolicy() *ModeOptionSchedulingPolicy {
	this := ModeOptionSchedulingPolicy{}
	return &this
}

// NewModeOptionSchedulingPolicyWithDefaults instantiates a new ModeOptionSchedulingPolicy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeOptionSchedulingPolicyWithDefaults() *ModeOptionSchedulingPolicy {
	this := ModeOptionSchedulingPolicy{}
	return &this
}

// GetComponentAntiAffinity returns the ComponentAntiAffinity field value if set, zero value otherwise.
func (o *ModeOptionSchedulingPolicy) GetComponentAntiAffinity() []string {
	if o == nil || o.ComponentAntiAffinity == nil {
		var ret []string
		return ret
	}
	return o.ComponentAntiAffinity
}

// GetComponentAntiAffinityOk returns a tuple with the ComponentAntiAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOptionSchedulingPolicy) GetComponentAntiAffinityOk() (*[]string, bool) {
	if o == nil || o.ComponentAntiAffinity == nil {
		return nil, false
	}
	return &o.ComponentAntiAffinity, true
}

// HasComponentAntiAffinity returns a boolean if a field has been set.
func (o *ModeOptionSchedulingPolicy) HasComponentAntiAffinity() bool {
	return o != nil && o.ComponentAntiAffinity != nil
}

// SetComponentAntiAffinity gets a reference to the given []string and assigns it to the ComponentAntiAffinity field.
func (o *ModeOptionSchedulingPolicy) SetComponentAntiAffinity(v []string) {
	o.ComponentAntiAffinity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeOptionSchedulingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ComponentAntiAffinity != nil {
		toSerialize["componentAntiAffinity"] = o.ComponentAntiAffinity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeOptionSchedulingPolicy) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentAntiAffinity []string `json:"componentAntiAffinity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"componentAntiAffinity"})
	} else {
		return err
	}
	o.ComponentAntiAffinity = all.ComponentAntiAffinity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
