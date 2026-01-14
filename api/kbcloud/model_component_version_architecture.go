// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ComponentVersionArchitecture Architecture support configuration for this component versions.
// If not set, the component versions support all architectures.
type ComponentVersionArchitecture struct {
	// List of versions that support amd64 architecture.
	// Supports version patterns: exact version "8.0.33", major version "8.0", version range ">=8.0.30", etc.
	// If not set, all versions support amd64.
	//
	Amd64 common.NullableList[string] `json:"amd64,omitempty"`
	// List of versions that support arm64 architecture.
	// Supports version patterns: exact version "8.0.33", major version "8.0", version range ">=8.0.30", etc.
	// If not set, all versions support arm64.
	//
	Arm64 common.NullableList[string] `json:"arm64,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentVersionArchitecture instantiates a new ComponentVersionArchitecture object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentVersionArchitecture() *ComponentVersionArchitecture {
	this := ComponentVersionArchitecture{}
	return &this
}

// NewComponentVersionArchitectureWithDefaults instantiates a new ComponentVersionArchitecture object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentVersionArchitectureWithDefaults() *ComponentVersionArchitecture {
	this := ComponentVersionArchitecture{}
	return &this
}

// GetAmd64 returns the Amd64 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentVersionArchitecture) GetAmd64() []string {
	if o == nil || o.Amd64.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Amd64.Get()
}

// GetAmd64Ok returns a tuple with the Amd64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ComponentVersionArchitecture) GetAmd64Ok() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amd64.Get(), o.Amd64.IsSet()
}

// HasAmd64 returns a boolean if a field has been set.
func (o *ComponentVersionArchitecture) HasAmd64() bool {
	return o != nil && o.Amd64.IsSet()
}

// SetAmd64 gets a reference to the given common.NullableList[string] and assigns it to the Amd64 field.
func (o *ComponentVersionArchitecture) SetAmd64(v []string) {
	o.Amd64.Set(&v)
}

// SetAmd64Nil sets the value for Amd64 to be an explicit nil.
func (o *ComponentVersionArchitecture) SetAmd64Nil() {
	o.Amd64.Set(nil)
}

// UnsetAmd64 ensures that no value is present for Amd64, not even an explicit nil.
func (o *ComponentVersionArchitecture) UnsetAmd64() {
	o.Amd64.Unset()
}

// GetArm64 returns the Arm64 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentVersionArchitecture) GetArm64() []string {
	if o == nil || o.Arm64.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Arm64.Get()
}

// GetArm64Ok returns a tuple with the Arm64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ComponentVersionArchitecture) GetArm64Ok() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Arm64.Get(), o.Arm64.IsSet()
}

// HasArm64 returns a boolean if a field has been set.
func (o *ComponentVersionArchitecture) HasArm64() bool {
	return o != nil && o.Arm64.IsSet()
}

// SetArm64 gets a reference to the given common.NullableList[string] and assigns it to the Arm64 field.
func (o *ComponentVersionArchitecture) SetArm64(v []string) {
	o.Arm64.Set(&v)
}

// SetArm64Nil sets the value for Arm64 to be an explicit nil.
func (o *ComponentVersionArchitecture) SetArm64Nil() {
	o.Arm64.Set(nil)
}

// UnsetArm64 ensures that no value is present for Arm64, not even an explicit nil.
func (o *ComponentVersionArchitecture) UnsetArm64() {
	o.Arm64.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentVersionArchitecture) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Amd64.IsSet() {
		toSerialize["amd64"] = o.Amd64.Get()
	}
	if o.Arm64.IsSet() {
		toSerialize["arm64"] = o.Arm64.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentVersionArchitecture) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Amd64 common.NullableList[string] `json:"amd64,omitempty"`
		Arm64 common.NullableList[string] `json:"arm64,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"amd64", "arm64"})
	} else {
		return err
	}
	o.Amd64 = all.Amd64
	o.Arm64 = all.Arm64

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableComponentVersionArchitecture handles when a null is used for ComponentVersionArchitecture.
type NullableComponentVersionArchitecture struct {
	value *ComponentVersionArchitecture
	isSet bool
}

// Get returns the associated value.
func (v NullableComponentVersionArchitecture) Get() *ComponentVersionArchitecture {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableComponentVersionArchitecture) Set(val *ComponentVersionArchitecture) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableComponentVersionArchitecture) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableComponentVersionArchitecture) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableComponentVersionArchitecture initializes the struct as if Set has been called.
func NewNullableComponentVersionArchitecture(val *ComponentVersionArchitecture) *NullableComponentVersionArchitecture {
	return &NullableComponentVersionArchitecture{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableComponentVersionArchitecture) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableComponentVersionArchitecture) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return common.Unmarshal(src, &v.value)
}
