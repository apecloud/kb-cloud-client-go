// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type EngineServiceVersionsItem struct {
	Default             *bool    `json:"default,omitempty"`
	DefaultMinorVersion *string  `json:"defaultMinorVersion,omitempty"`
	MajorVersion        *string  `json:"majorVersion,omitempty"`
	MinorVersions       []string `json:"minorVersions,omitempty"`
	// Architecture support configuration for this component versions.
	// If not set, the component versions support all architectures.
	//
	Architecture NullableComponentVersionArchitecture `json:"architecture,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineServiceVersionsItem instantiates a new EngineServiceVersionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineServiceVersionsItem() *EngineServiceVersionsItem {
	this := EngineServiceVersionsItem{}
	return &this
}

// NewEngineServiceVersionsItemWithDefaults instantiates a new EngineServiceVersionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineServiceVersionsItemWithDefaults() *EngineServiceVersionsItem {
	this := EngineServiceVersionsItem{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *EngineServiceVersionsItem) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersionsItem) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *EngineServiceVersionsItem) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *EngineServiceVersionsItem) SetDefault(v bool) {
	o.Default = &v
}

// GetDefaultMinorVersion returns the DefaultMinorVersion field value if set, zero value otherwise.
func (o *EngineServiceVersionsItem) GetDefaultMinorVersion() string {
	if o == nil || o.DefaultMinorVersion == nil {
		var ret string
		return ret
	}
	return *o.DefaultMinorVersion
}

// GetDefaultMinorVersionOk returns a tuple with the DefaultMinorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersionsItem) GetDefaultMinorVersionOk() (*string, bool) {
	if o == nil || o.DefaultMinorVersion == nil {
		return nil, false
	}
	return o.DefaultMinorVersion, true
}

// HasDefaultMinorVersion returns a boolean if a field has been set.
func (o *EngineServiceVersionsItem) HasDefaultMinorVersion() bool {
	return o != nil && o.DefaultMinorVersion != nil
}

// SetDefaultMinorVersion gets a reference to the given string and assigns it to the DefaultMinorVersion field.
func (o *EngineServiceVersionsItem) SetDefaultMinorVersion(v string) {
	o.DefaultMinorVersion = &v
}

// GetMajorVersion returns the MajorVersion field value if set, zero value otherwise.
func (o *EngineServiceVersionsItem) GetMajorVersion() string {
	if o == nil || o.MajorVersion == nil {
		var ret string
		return ret
	}
	return *o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersionsItem) GetMajorVersionOk() (*string, bool) {
	if o == nil || o.MajorVersion == nil {
		return nil, false
	}
	return o.MajorVersion, true
}

// HasMajorVersion returns a boolean if a field has been set.
func (o *EngineServiceVersionsItem) HasMajorVersion() bool {
	return o != nil && o.MajorVersion != nil
}

// SetMajorVersion gets a reference to the given string and assigns it to the MajorVersion field.
func (o *EngineServiceVersionsItem) SetMajorVersion(v string) {
	o.MajorVersion = &v
}

// GetMinorVersions returns the MinorVersions field value if set, zero value otherwise.
func (o *EngineServiceVersionsItem) GetMinorVersions() []string {
	if o == nil || o.MinorVersions == nil {
		var ret []string
		return ret
	}
	return o.MinorVersions
}

// GetMinorVersionsOk returns a tuple with the MinorVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersionsItem) GetMinorVersionsOk() (*[]string, bool) {
	if o == nil || o.MinorVersions == nil {
		return nil, false
	}
	return &o.MinorVersions, true
}

// HasMinorVersions returns a boolean if a field has been set.
func (o *EngineServiceVersionsItem) HasMinorVersions() bool {
	return o != nil && o.MinorVersions != nil
}

// SetMinorVersions gets a reference to the given []string and assigns it to the MinorVersions field.
func (o *EngineServiceVersionsItem) SetMinorVersions(v []string) {
	o.MinorVersions = v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineServiceVersionsItem) GetArchitecture() ComponentVersionArchitecture {
	if o == nil || o.Architecture.Get() == nil {
		var ret ComponentVersionArchitecture
		return ret
	}
	return *o.Architecture.Get()
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EngineServiceVersionsItem) GetArchitectureOk() (*ComponentVersionArchitecture, bool) {
	if o == nil {
		return nil, false
	}
	return o.Architecture.Get(), o.Architecture.IsSet()
}

// HasArchitecture returns a boolean if a field has been set.
func (o *EngineServiceVersionsItem) HasArchitecture() bool {
	return o != nil && o.Architecture.IsSet()
}

// SetArchitecture gets a reference to the given NullableComponentVersionArchitecture and assigns it to the Architecture field.
func (o *EngineServiceVersionsItem) SetArchitecture(v ComponentVersionArchitecture) {
	o.Architecture.Set(&v)
}

// SetArchitectureNil sets the value for Architecture to be an explicit nil.
func (o *EngineServiceVersionsItem) SetArchitectureNil() {
	o.Architecture.Set(nil)
}

// UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil.
func (o *EngineServiceVersionsItem) UnsetArchitecture() {
	o.Architecture.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineServiceVersionsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.DefaultMinorVersion != nil {
		toSerialize["defaultMinorVersion"] = o.DefaultMinorVersion
	}
	if o.MajorVersion != nil {
		toSerialize["majorVersion"] = o.MajorVersion
	}
	if o.MinorVersions != nil {
		toSerialize["minorVersions"] = o.MinorVersions
	}
	if o.Architecture.IsSet() {
		toSerialize["architecture"] = o.Architecture.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineServiceVersionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default             *bool                                `json:"default,omitempty"`
		DefaultMinorVersion *string                              `json:"defaultMinorVersion,omitempty"`
		MajorVersion        *string                              `json:"majorVersion,omitempty"`
		MinorVersions       []string                             `json:"minorVersions,omitempty"`
		Architecture        NullableComponentVersionArchitecture `json:"architecture,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"default", "defaultMinorVersion", "majorVersion", "minorVersions", "architecture"})
	} else {
		return err
	}
	o.Default = all.Default
	o.DefaultMinorVersion = all.DefaultMinorVersion
	o.MajorVersion = all.MajorVersion
	o.MinorVersions = all.MinorVersions
	o.Architecture = all.Architecture

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
