// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineTlsOption struct {
	// whether the engine supports TLS/SSL configuration
	Supported bool `json:"supported"`
	// Version patterns excluded from TLS when supported is true. If not set, all versions support TLS.
	// Each item is a regular expression (RE2 / Go regexp syntax) matched against the engine/service version string.
	// Example: ^8\.1\.3($|[.\-]) matches 8.1.3 and 8.1.3.* but not 8.1.30.
	//
	ExcludedVersions common.NullableList[string] `json:"excludedVersions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineTlsOption instantiates a new EngineTlsOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineTlsOption(supported bool) *EngineTlsOption {
	this := EngineTlsOption{}
	this.Supported = supported
	return &this
}

// NewEngineTlsOptionWithDefaults instantiates a new EngineTlsOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineTlsOptionWithDefaults() *EngineTlsOption {
	this := EngineTlsOption{}
	return &this
}

// GetSupported returns the Supported field value.
func (o *EngineTlsOption) GetSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *EngineTlsOption) GetSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supported, true
}

// SetSupported sets field value.
func (o *EngineTlsOption) SetSupported(v bool) {
	o.Supported = v
}

// GetExcludedVersions returns the ExcludedVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineTlsOption) GetExcludedVersions() []string {
	if o == nil || o.ExcludedVersions.Get() == nil {
		var ret []string
		return ret
	}
	return *o.ExcludedVersions.Get()
}

// GetExcludedVersionsOk returns a tuple with the ExcludedVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineTlsOption) GetExcludedVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExcludedVersions.Get(), o.ExcludedVersions.IsSet()
}

// HasExcludedVersions returns a boolean if a field has been set.
func (o *EngineTlsOption) HasExcludedVersions() bool {
	return o != nil && o.ExcludedVersions.IsSet()
}

// SetExcludedVersions gets a reference to the given common.NullableList[string] and assigns it to the ExcludedVersions field.
func (o *EngineTlsOption) SetExcludedVersions(v []string) {
	o.ExcludedVersions.Set(&v)
}

// SetExcludedVersionsNil sets the value for ExcludedVersions to be an explicit nil.
func (o *EngineTlsOption) SetExcludedVersionsNil() {
	o.ExcludedVersions.Set(nil)
}

// UnsetExcludedVersions ensures that no value is present for ExcludedVersions, not even an explicit nil.
func (o *EngineTlsOption) UnsetExcludedVersions() {
	o.ExcludedVersions.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineTlsOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["supported"] = o.Supported
	if o.ExcludedVersions.IsSet() {
		toSerialize["excludedVersions"] = o.ExcludedVersions.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineTlsOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Supported        *bool                       `json:"supported"`
		ExcludedVersions common.NullableList[string] `json:"excludedVersions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Supported == nil {
		return fmt.Errorf("required field supported missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"supported", "excludedVersions"})
	} else {
		return err
	}
	o.Supported = *all.Supported
	o.ExcludedVersions = all.ExcludedVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
