// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ComponentOptionVersionMajorVersionVersionMappingItem Configure the mapping relationship with the main component's major versions.
type ComponentOptionVersionMajorVersionVersionMappingItem struct {
	// major version of the main component
	MainComponentMajorVersion *string `json:"mainComponentMajorVersion,omitempty"`
	// major versions of the current component
	Versions []string `json:"versions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOptionVersionMajorVersionVersionMappingItem instantiates a new ComponentOptionVersionMajorVersionVersionMappingItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOptionVersionMajorVersionVersionMappingItem() *ComponentOptionVersionMajorVersionVersionMappingItem {
	this := ComponentOptionVersionMajorVersionVersionMappingItem{}
	return &this
}

// NewComponentOptionVersionMajorVersionVersionMappingItemWithDefaults instantiates a new ComponentOptionVersionMajorVersionVersionMappingItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOptionVersionMajorVersionVersionMappingItemWithDefaults() *ComponentOptionVersionMajorVersionVersionMappingItem {
	this := ComponentOptionVersionMajorVersionVersionMappingItem{}
	return &this
}

// GetMainComponentMajorVersion returns the MainComponentMajorVersion field value if set, zero value otherwise.
func (o *ComponentOptionVersionMajorVersionVersionMappingItem) GetMainComponentMajorVersion() string {
	if o == nil || o.MainComponentMajorVersion == nil {
		var ret string
		return ret
	}
	return *o.MainComponentMajorVersion
}

// GetMainComponentMajorVersionOk returns a tuple with the MainComponentMajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersionMajorVersionVersionMappingItem) GetMainComponentMajorVersionOk() (*string, bool) {
	if o == nil || o.MainComponentMajorVersion == nil {
		return nil, false
	}
	return o.MainComponentMajorVersion, true
}

// HasMainComponentMajorVersion returns a boolean if a field has been set.
func (o *ComponentOptionVersionMajorVersionVersionMappingItem) HasMainComponentMajorVersion() bool {
	return o != nil && o.MainComponentMajorVersion != nil
}

// SetMainComponentMajorVersion gets a reference to the given string and assigns it to the MainComponentMajorVersion field.
func (o *ComponentOptionVersionMajorVersionVersionMappingItem) SetMainComponentMajorVersion(v string) {
	o.MainComponentMajorVersion = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ComponentOptionVersionMajorVersionVersionMappingItem) GetVersions() []string {
	if o == nil || o.Versions == nil {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersionMajorVersionVersionMappingItem) GetVersionsOk() (*[]string, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return &o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ComponentOptionVersionMajorVersionVersionMappingItem) HasVersions() bool {
	return o != nil && o.Versions != nil
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *ComponentOptionVersionMajorVersionVersionMappingItem) SetVersions(v []string) {
	o.Versions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOptionVersionMajorVersionVersionMappingItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.MainComponentMajorVersion != nil {
		toSerialize["mainComponentMajorVersion"] = o.MainComponentMajorVersion
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOptionVersionMajorVersionVersionMappingItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MainComponentMajorVersion *string  `json:"mainComponentMajorVersion,omitempty"`
		Versions                  []string `json:"versions,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"mainComponentMajorVersion", "versions"})
	} else {
		return err
	}
	o.MainComponentMajorVersion = all.MainComponentMajorVersion
	o.Versions = all.Versions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
