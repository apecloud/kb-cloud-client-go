// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ComponentOptionVersion struct {
	// component version name
	ComponentVersionName string                              `json:"componentVersionName"`
	MinorVersion         *ComponentOptionVersionMinorVersion `json:"minorVersion,omitempty"`
	MajorVersion         ComponentOptionVersionMajorVersion  `json:"majorVersion"`
	// Architecture support configuration for this component versions.
	// If not set, the component versions support all architectures.
	//
	Architecture NullableComponentVersionArchitecture `json:"architecture,omitempty"`
	// Upgrades that must not be offered or accepted for this component, keyed by the service version the cluster currently runs. Use it when a target version is not a drop-in replacement for a source version, for example when the source image ships modules the target image no longer carries. Upgrades matching no rule stay untouched.
	BlockedUpgradeFrom []ComponentOptionVersionBlockedUpgradeFromItem `json:"blockedUpgradeFrom,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComponentOptionVersion instantiates a new ComponentOptionVersion object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComponentOptionVersion(componentVersionName string, majorVersion ComponentOptionVersionMajorVersion) *ComponentOptionVersion {
	this := ComponentOptionVersion{}
	this.ComponentVersionName = componentVersionName
	this.MajorVersion = majorVersion
	return &this
}

// NewComponentOptionVersionWithDefaults instantiates a new ComponentOptionVersion object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComponentOptionVersionWithDefaults() *ComponentOptionVersion {
	this := ComponentOptionVersion{}
	return &this
}

// GetComponentVersionName returns the ComponentVersionName field value.
func (o *ComponentOptionVersion) GetComponentVersionName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ComponentVersionName
}

// GetComponentVersionNameOk returns a tuple with the ComponentVersionName field value
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersion) GetComponentVersionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentVersionName, true
}

// SetComponentVersionName sets field value.
func (o *ComponentOptionVersion) SetComponentVersionName(v string) {
	o.ComponentVersionName = v
}

// GetMinorVersion returns the MinorVersion field value if set, zero value otherwise.
func (o *ComponentOptionVersion) GetMinorVersion() ComponentOptionVersionMinorVersion {
	if o == nil || o.MinorVersion == nil {
		var ret ComponentOptionVersionMinorVersion
		return ret
	}
	return *o.MinorVersion
}

// GetMinorVersionOk returns a tuple with the MinorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersion) GetMinorVersionOk() (*ComponentOptionVersionMinorVersion, bool) {
	if o == nil || o.MinorVersion == nil {
		return nil, false
	}
	return o.MinorVersion, true
}

// HasMinorVersion returns a boolean if a field has been set.
func (o *ComponentOptionVersion) HasMinorVersion() bool {
	return o != nil && o.MinorVersion != nil
}

// SetMinorVersion gets a reference to the given ComponentOptionVersionMinorVersion and assigns it to the MinorVersion field.
func (o *ComponentOptionVersion) SetMinorVersion(v ComponentOptionVersionMinorVersion) {
	o.MinorVersion = &v
}

// GetMajorVersion returns the MajorVersion field value.
func (o *ComponentOptionVersion) GetMajorVersion() ComponentOptionVersionMajorVersion {
	if o == nil {
		var ret ComponentOptionVersionMajorVersion
		return ret
	}
	return o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersion) GetMajorVersionOk() (*ComponentOptionVersionMajorVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MajorVersion, true
}

// SetMajorVersion sets field value.
func (o *ComponentOptionVersion) SetMajorVersion(v ComponentOptionVersionMajorVersion) {
	o.MajorVersion = v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComponentOptionVersion) GetArchitecture() ComponentVersionArchitecture {
	if o == nil || o.Architecture.Get() == nil {
		var ret ComponentVersionArchitecture
		return ret
	}
	return *o.Architecture.Get()
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ComponentOptionVersion) GetArchitectureOk() (*ComponentVersionArchitecture, bool) {
	if o == nil {
		return nil, false
	}
	return o.Architecture.Get(), o.Architecture.IsSet()
}

// HasArchitecture returns a boolean if a field has been set.
func (o *ComponentOptionVersion) HasArchitecture() bool {
	return o != nil && o.Architecture.IsSet()
}

// SetArchitecture gets a reference to the given NullableComponentVersionArchitecture and assigns it to the Architecture field.
func (o *ComponentOptionVersion) SetArchitecture(v ComponentVersionArchitecture) {
	o.Architecture.Set(&v)
}

// SetArchitectureNil sets the value for Architecture to be an explicit nil.
func (o *ComponentOptionVersion) SetArchitectureNil() {
	o.Architecture.Set(nil)
}

// UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil.
func (o *ComponentOptionVersion) UnsetArchitecture() {
	o.Architecture.Unset()
}

// GetBlockedUpgradeFrom returns the BlockedUpgradeFrom field value if set, zero value otherwise.
func (o *ComponentOptionVersion) GetBlockedUpgradeFrom() []ComponentOptionVersionBlockedUpgradeFromItem {
	if o == nil || o.BlockedUpgradeFrom == nil {
		var ret []ComponentOptionVersionBlockedUpgradeFromItem
		return ret
	}
	return o.BlockedUpgradeFrom
}

// GetBlockedUpgradeFromOk returns a tuple with the BlockedUpgradeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentOptionVersion) GetBlockedUpgradeFromOk() (*[]ComponentOptionVersionBlockedUpgradeFromItem, bool) {
	if o == nil || o.BlockedUpgradeFrom == nil {
		return nil, false
	}
	return &o.BlockedUpgradeFrom, true
}

// HasBlockedUpgradeFrom returns a boolean if a field has been set.
func (o *ComponentOptionVersion) HasBlockedUpgradeFrom() bool {
	return o != nil && o.BlockedUpgradeFrom != nil
}

// SetBlockedUpgradeFrom gets a reference to the given []ComponentOptionVersionBlockedUpgradeFromItem and assigns it to the BlockedUpgradeFrom field.
func (o *ComponentOptionVersion) SetBlockedUpgradeFrom(v []ComponentOptionVersionBlockedUpgradeFromItem) {
	o.BlockedUpgradeFrom = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComponentOptionVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["componentVersionName"] = o.ComponentVersionName
	if o.MinorVersion != nil {
		toSerialize["minorVersion"] = o.MinorVersion
	}
	toSerialize["majorVersion"] = o.MajorVersion
	if o.Architecture.IsSet() {
		toSerialize["architecture"] = o.Architecture.Get()
	}
	if o.BlockedUpgradeFrom != nil {
		toSerialize["blockedUpgradeFrom"] = o.BlockedUpgradeFrom
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComponentOptionVersion) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentVersionName *string                                        `json:"componentVersionName"`
		MinorVersion         *ComponentOptionVersionMinorVersion            `json:"minorVersion,omitempty"`
		MajorVersion         *ComponentOptionVersionMajorVersion            `json:"majorVersion"`
		Architecture         NullableComponentVersionArchitecture           `json:"architecture,omitempty"`
		BlockedUpgradeFrom   []ComponentOptionVersionBlockedUpgradeFromItem `json:"blockedUpgradeFrom,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ComponentVersionName == nil {
		return fmt.Errorf("required field componentVersionName missing")
	}
	if all.MajorVersion == nil {
		return fmt.Errorf("required field majorVersion missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"componentVersionName", "minorVersion", "majorVersion", "architecture", "blockedUpgradeFrom"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ComponentVersionName = *all.ComponentVersionName
	if all.MinorVersion != nil && all.MinorVersion.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MinorVersion = all.MinorVersion
	if all.MajorVersion.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MajorVersion = *all.MajorVersion
	o.Architecture = all.Architecture
	o.BlockedUpgradeFrom = all.BlockedUpgradeFrom

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
