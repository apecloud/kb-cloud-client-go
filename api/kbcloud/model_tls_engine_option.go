// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TlsEngineOption struct {
	// whether the engine supports TLS/SSL configuration
	Supported bool `json:"supported"`
	// list of components that support TLS with version constraints. If not set, all components support TLS when supported is true
	Components []TlsComponentOption `json:"components,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTlsEngineOption instantiates a new TlsEngineOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTlsEngineOption(supported bool) *TlsEngineOption {
	this := TlsEngineOption{}
	this.Supported = supported
	return &this
}

// NewTlsEngineOptionWithDefaults instantiates a new TlsEngineOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTlsEngineOptionWithDefaults() *TlsEngineOption {
	this := TlsEngineOption{}
	return &this
}

// GetSupported returns the Supported field value.
func (o *TlsEngineOption) GetSupported() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value
// and a boolean to check if the value has been set.
func (o *TlsEngineOption) GetSupportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supported, true
}

// SetSupported sets field value.
func (o *TlsEngineOption) SetSupported(v bool) {
	o.Supported = v
}

// GetComponents returns the Components field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsEngineOption) GetComponents() []TlsComponentOption {
	if o == nil {
		var ret []TlsComponentOption
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TlsEngineOption) GetComponentsOk() (*[]TlsComponentOption, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *TlsEngineOption) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []TlsComponentOption and assigns it to the Components field.
func (o *TlsEngineOption) SetComponents(v []TlsComponentOption) {
	o.Components = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TlsEngineOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["supported"] = o.Supported
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TlsEngineOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Supported  *bool                `json:"supported"`
		Components []TlsComponentOption `json:"components,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Supported == nil {
		return fmt.Errorf("required field supported missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"supported", "components"})
	} else {
		return err
	}
	o.Supported = *all.Supported
	o.Components = all.Components

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
