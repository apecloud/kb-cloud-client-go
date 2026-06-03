// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type EngineNetworkModeSupported struct {
	// supported network modes
	Supported []NetworkMode `json:"supported,omitempty"`
	// not supported versions, key is networkMode, value is array of versions
	NotSupportedVersions map[string][]string `json:"notSupportedVersions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineNetworkModeSupported instantiates a new EngineNetworkModeSupported object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineNetworkModeSupported() *EngineNetworkModeSupported {
	this := EngineNetworkModeSupported{}
	return &this
}

// NewEngineNetworkModeSupportedWithDefaults instantiates a new EngineNetworkModeSupported object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineNetworkModeSupportedWithDefaults() *EngineNetworkModeSupported {
	this := EngineNetworkModeSupported{}
	return &this
}

// GetSupported returns the Supported field value if set, zero value otherwise.
func (o *EngineNetworkModeSupported) GetSupported() []NetworkMode {
	if o == nil || o.Supported == nil {
		var ret []NetworkMode
		return ret
	}
	return o.Supported
}

// GetSupportedOk returns a tuple with the Supported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineNetworkModeSupported) GetSupportedOk() (*[]NetworkMode, bool) {
	if o == nil || o.Supported == nil {
		return nil, false
	}
	return &o.Supported, true
}

// HasSupported returns a boolean if a field has been set.
func (o *EngineNetworkModeSupported) HasSupported() bool {
	return o != nil && o.Supported != nil
}

// SetSupported gets a reference to the given []NetworkMode and assigns it to the Supported field.
func (o *EngineNetworkModeSupported) SetSupported(v []NetworkMode) {
	o.Supported = v
}

// GetNotSupportedVersions returns the NotSupportedVersions field value if set, zero value otherwise.
func (o *EngineNetworkModeSupported) GetNotSupportedVersions() map[string][]string {
	if o == nil || o.NotSupportedVersions == nil {
		var ret map[string][]string
		return ret
	}
	return o.NotSupportedVersions
}

// GetNotSupportedVersionsOk returns a tuple with the NotSupportedVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineNetworkModeSupported) GetNotSupportedVersionsOk() (*map[string][]string, bool) {
	if o == nil || o.NotSupportedVersions == nil {
		return nil, false
	}
	return &o.NotSupportedVersions, true
}

// HasNotSupportedVersions returns a boolean if a field has been set.
func (o *EngineNetworkModeSupported) HasNotSupportedVersions() bool {
	return o != nil && o.NotSupportedVersions != nil
}

// SetNotSupportedVersions gets a reference to the given map[string][]string and assigns it to the NotSupportedVersions field.
func (o *EngineNetworkModeSupported) SetNotSupportedVersions(v map[string][]string) {
	o.NotSupportedVersions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineNetworkModeSupported) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Supported != nil {
		toSerialize["supported"] = o.Supported
	}
	if o.NotSupportedVersions != nil {
		toSerialize["notSupportedVersions"] = o.NotSupportedVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineNetworkModeSupported) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Supported            []NetworkMode       `json:"supported,omitempty"`
		NotSupportedVersions map[string][]string `json:"notSupportedVersions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"supported", "notSupportedVersions"})
	} else {
		return err
	}
	o.Supported = all.Supported
	o.NotSupportedVersions = all.NotSupportedVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
