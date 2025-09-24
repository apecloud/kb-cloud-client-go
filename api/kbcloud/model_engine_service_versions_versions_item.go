// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineServiceVersionsVersionsItem struct {
	Default             *bool    `json:"default,omitempty"`
	DefaultMinorVersion *string  `json:"defaultMinorVersion,omitempty"`
	MajorVersion        *string  `json:"majorVersion,omitempty"`
	MinorVersions       []string `json:"minorVersions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineServiceVersionsVersionsItem instantiates a new EngineServiceVersionsVersionsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineServiceVersionsVersionsItem() *EngineServiceVersionsVersionsItem {
	this := EngineServiceVersionsVersionsItem{}
	return &this
}

// NewEngineServiceVersionsVersionsItemWithDefaults instantiates a new EngineServiceVersionsVersionsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineServiceVersionsVersionsItemWithDefaults() *EngineServiceVersionsVersionsItem {
	this := EngineServiceVersionsVersionsItem{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *EngineServiceVersionsVersionsItem) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersionsVersionsItem) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *EngineServiceVersionsVersionsItem) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *EngineServiceVersionsVersionsItem) SetDefault(v bool) {
	o.Default = &v
}

// GetDefaultMinorVersion returns the DefaultMinorVersion field value if set, zero value otherwise.
func (o *EngineServiceVersionsVersionsItem) GetDefaultMinorVersion() string {
	if o == nil || o.DefaultMinorVersion == nil {
		var ret string
		return ret
	}
	return *o.DefaultMinorVersion
}

// GetDefaultMinorVersionOk returns a tuple with the DefaultMinorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersionsVersionsItem) GetDefaultMinorVersionOk() (*string, bool) {
	if o == nil || o.DefaultMinorVersion == nil {
		return nil, false
	}
	return o.DefaultMinorVersion, true
}

// HasDefaultMinorVersion returns a boolean if a field has been set.
func (o *EngineServiceVersionsVersionsItem) HasDefaultMinorVersion() bool {
	return o != nil && o.DefaultMinorVersion != nil
}

// SetDefaultMinorVersion gets a reference to the given string and assigns it to the DefaultMinorVersion field.
func (o *EngineServiceVersionsVersionsItem) SetDefaultMinorVersion(v string) {
	o.DefaultMinorVersion = &v
}

// GetMajorVersion returns the MajorVersion field value if set, zero value otherwise.
func (o *EngineServiceVersionsVersionsItem) GetMajorVersion() string {
	if o == nil || o.MajorVersion == nil {
		var ret string
		return ret
	}
	return *o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersionsVersionsItem) GetMajorVersionOk() (*string, bool) {
	if o == nil || o.MajorVersion == nil {
		return nil, false
	}
	return o.MajorVersion, true
}

// HasMajorVersion returns a boolean if a field has been set.
func (o *EngineServiceVersionsVersionsItem) HasMajorVersion() bool {
	return o != nil && o.MajorVersion != nil
}

// SetMajorVersion gets a reference to the given string and assigns it to the MajorVersion field.
func (o *EngineServiceVersionsVersionsItem) SetMajorVersion(v string) {
	o.MajorVersion = &v
}

// GetMinorVersions returns the MinorVersions field value if set, zero value otherwise.
func (o *EngineServiceVersionsVersionsItem) GetMinorVersions() []string {
	if o == nil || o.MinorVersions == nil {
		var ret []string
		return ret
	}
	return o.MinorVersions
}

// GetMinorVersionsOk returns a tuple with the MinorVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineServiceVersionsVersionsItem) GetMinorVersionsOk() (*[]string, bool) {
	if o == nil || o.MinorVersions == nil {
		return nil, false
	}
	return &o.MinorVersions, true
}

// HasMinorVersions returns a boolean if a field has been set.
func (o *EngineServiceVersionsVersionsItem) HasMinorVersions() bool {
	return o != nil && o.MinorVersions != nil
}

// SetMinorVersions gets a reference to the given []string and assigns it to the MinorVersions field.
func (o *EngineServiceVersionsVersionsItem) SetMinorVersions(v []string) {
	o.MinorVersions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineServiceVersionsVersionsItem) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineServiceVersionsVersionsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default             *bool    `json:"default,omitempty"`
		DefaultMinorVersion *string  `json:"defaultMinorVersion,omitempty"`
		MajorVersion        *string  `json:"majorVersion,omitempty"`
		MinorVersions       []string `json:"minorVersions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"default", "defaultMinorVersion", "majorVersion", "minorVersions"})
	} else {
		return err
	}
	o.Default = all.Default
	o.DefaultMinorVersion = all.DefaultMinorVersion
	o.MajorVersion = all.MajorVersion
	o.MinorVersions = all.MinorVersions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
