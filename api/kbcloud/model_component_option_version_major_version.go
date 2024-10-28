// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION ComponentOptionVersionMajorVersion
type ComponentOptionVersionMajorVersion struct {
	// default major version.
	Default *string `json:"default,omitempty"`
	// ServiceVersion uses semver syntax(X Y.Z), such MySQL 5.7.4 and PG 14.8.0.
	// Currently supports [X, X.Y, X.Y.Z] formats to determine major versions of the engine from the serviceVersion
	//
	Rule string `json:"rule"`
	// NODESCRIPTION VersionMapping
	VersionMapping []ComponentOptionVersionMajorVersionVersionMappingItem `json:"versionMapping,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOptionVersionMajorVersion instantiates a new ComponentOptionVersionMajorVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOptionVersionMajorVersion(rule string) *ComponentOptionVersionMajorVersion {
	this := ComponentOptionVersionMajorVersion{}
	this.Rule = rule
	return &this
}

// NewComponentOptionVersionMajorVersionWithDefaults instantiates a new ComponentOptionVersionMajorVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOptionVersionMajorVersionWithDefaults() *ComponentOptionVersionMajorVersion {
	this := ComponentOptionVersionMajorVersion{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *ComponentOptionVersionMajorVersion) GetDefault() string {
	if o == nil || o.Default == nil {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersionMajorVersion) GetDefaultOk() (*string, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *ComponentOptionVersionMajorVersion) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *ComponentOptionVersionMajorVersion) SetDefault(v string) {
	o.Default = &v
}

// GetRule returns the Rule field value.
func (o *ComponentOptionVersionMajorVersion) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersionMajorVersion) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value.
func (o *ComponentOptionVersionMajorVersion) SetRule(v string) {
	o.Rule = v
}

// GetVersionMapping returns the VersionMapping field value if set, zero value otherwise.
func (o *ComponentOptionVersionMajorVersion) GetVersionMapping() []ComponentOptionVersionMajorVersionVersionMappingItem {
	if o == nil || o.VersionMapping == nil {
		var ret []ComponentOptionVersionMajorVersionVersionMappingItem
		return ret
	}
	return o.VersionMapping
}

// GetVersionMappingOk returns a tuple with the VersionMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersionMajorVersion) GetVersionMappingOk() (*[]ComponentOptionVersionMajorVersionVersionMappingItem, bool) {
	if o == nil || o.VersionMapping == nil {
		return nil, false
	}
	return &o.VersionMapping, true
}

// HasVersionMapping returns a boolean if a field has been set.
func (o *ComponentOptionVersionMajorVersion) HasVersionMapping() bool {
	return o != nil && o.VersionMapping != nil
}

// SetVersionMapping gets a reference to the given []ComponentOptionVersionMajorVersionVersionMappingItem and assigns it to the VersionMapping field.
func (o *ComponentOptionVersionMajorVersion) SetVersionMapping(v []ComponentOptionVersionMajorVersionVersionMappingItem) {
	o.VersionMapping = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOptionVersionMajorVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	toSerialize["rule"] = o.Rule
	if o.VersionMapping != nil {
		toSerialize["versionMapping"] = o.VersionMapping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOptionVersionMajorVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Default        *string                                                `json:"default,omitempty"`
		Rule           *string                                                `json:"rule"`
		VersionMapping []ComponentOptionVersionMajorVersionVersionMappingItem `json:"versionMapping,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Rule == nil {
		return fmt.Errorf("required field rule missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"default", "rule", "versionMapping"})
	} else {
		return err
	}
	o.Default = all.Default
	o.Rule = *all.Rule
	o.VersionMapping = all.VersionMapping

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
