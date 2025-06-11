// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type NetworkModeOptionItem struct {
	Supported []NetworkMode `json:"supported"`
	// not supported versions, key is the networkMode, value is an array of versions which support major version or mirror version.
	NotSupportedVersions map[string][]string         `json:"notSupportedVersions,omitempty"`
	Modes                common.NullableList[string] `json:"modes,omitempty"`
	Versions             common.NullableList[string] `json:"versions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNetworkModeOptionItem instantiates a new NetworkModeOptionItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNetworkModeOptionItem(supported []NetworkMode) *NetworkModeOptionItem {
	this := NetworkModeOptionItem{}
	this.Supported = supported
	return &this
}

// NewNetworkModeOptionItemWithDefaults instantiates a new NetworkModeOptionItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNetworkModeOptionItemWithDefaults() *NetworkModeOptionItem {
	this := NetworkModeOptionItem{}
	return &this
}

// GetSupported returns the Supported field value.
func (o *NetworkModeOptionItem) GetSupported() []NetworkMode {
	if o == nil {
		var ret []NetworkMode
		return ret
	}
	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *NetworkModeOptionItem) GetSupportedOk() (*[]NetworkMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supported, true
}

// SetSupported sets field value.
func (o *NetworkModeOptionItem) SetSupported(v []NetworkMode) {
	o.Supported = v
}

// GetNotSupportedVersions returns the NotSupportedVersions field value if set, zero value otherwise.
func (o *NetworkModeOptionItem) GetNotSupportedVersions() map[string][]string {
	if o == nil || o.NotSupportedVersions == nil {
		var ret map[string][]string
		return ret
	}
	return o.NotSupportedVersions
}

// GetNotSupportedVersionsOk returns a tuple with the NotSupportedVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkModeOptionItem) GetNotSupportedVersionsOk() (*map[string][]string, bool) {
	if o == nil || o.NotSupportedVersions == nil {
		return nil, false
	}
	return &o.NotSupportedVersions, true
}

// HasNotSupportedVersions returns a boolean if a field has been set.
func (o *NetworkModeOptionItem) HasNotSupportedVersions() bool {
	return o != nil && o.NotSupportedVersions != nil
}

// SetNotSupportedVersions gets a reference to the given map[string][]string and assigns it to the NotSupportedVersions field.
func (o *NetworkModeOptionItem) SetNotSupportedVersions(v map[string][]string) {
	o.NotSupportedVersions = v
}

// GetModes returns the Modes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkModeOptionItem) GetModes() []string {
	if o == nil || o.Modes.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Modes.Get()
}

// GetModesOk returns a tuple with the Modes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NetworkModeOptionItem) GetModesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modes.Get(), o.Modes.IsSet()
}

// HasModes returns a boolean if a field has been set.
func (o *NetworkModeOptionItem) HasModes() bool {
	return o != nil && o.Modes.IsSet()
}

// SetModes gets a reference to the given common.NullableList[string] and assigns it to the Modes field.
func (o *NetworkModeOptionItem) SetModes(v []string) {
	o.Modes.Set(&v)
}

// SetModesNil sets the value for Modes to be an explicit nil.
func (o *NetworkModeOptionItem) SetModesNil() {
	o.Modes.Set(nil)
}

// UnsetModes ensures that no value is present for Modes, not even an explicit nil.
func (o *NetworkModeOptionItem) UnsetModes() {
	o.Modes.Unset()
}

// GetVersions returns the Versions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkModeOptionItem) GetVersions() []string {
	if o == nil || o.Versions.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Versions.Get()
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *NetworkModeOptionItem) GetVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Versions.Get(), o.Versions.IsSet()
}

// HasVersions returns a boolean if a field has been set.
func (o *NetworkModeOptionItem) HasVersions() bool {
	return o != nil && o.Versions.IsSet()
}

// SetVersions gets a reference to the given common.NullableList[string] and assigns it to the Versions field.
func (o *NetworkModeOptionItem) SetVersions(v []string) {
	o.Versions.Set(&v)
}

// SetVersionsNil sets the value for Versions to be an explicit nil.
func (o *NetworkModeOptionItem) SetVersionsNil() {
	o.Versions.Set(nil)
}

// UnsetVersions ensures that no value is present for Versions, not even an explicit nil.
func (o *NetworkModeOptionItem) UnsetVersions() {
	o.Versions.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o NetworkModeOptionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["supported"] = o.Supported
	if o.NotSupportedVersions != nil {
		toSerialize["notSupportedVersions"] = o.NotSupportedVersions
	}
	if o.Modes.IsSet() {
		toSerialize["modes"] = o.Modes.Get()
	}
	if o.Versions.IsSet() {
		toSerialize["versions"] = o.Versions.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NetworkModeOptionItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Supported            *[]NetworkMode              `json:"supported"`
		NotSupportedVersions map[string][]string         `json:"notSupportedVersions,omitempty"`
		Modes                common.NullableList[string] `json:"modes,omitempty"`
		Versions             common.NullableList[string] `json:"versions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Supported == nil {
		return fmt.Errorf("required field supported missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"supported", "notSupportedVersions", "modes", "versions"})
	} else {
		return err
	}
	o.Supported = *all.Supported
	o.NotSupportedVersions = all.NotSupportedVersions
	o.Modes = all.Modes
	o.Versions = all.Versions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
