// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ModeOptionValuesMappings struct {
	CompatibleKbVersions []string                               `json:"compatibleKBVersions,omitempty"`
	Mappings             []ModeOptionValuesMappingsMappingsItem `json:"mappings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeOptionValuesMappings instantiates a new ModeOptionValuesMappings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeOptionValuesMappings() *ModeOptionValuesMappings {
	this := ModeOptionValuesMappings{}
	return &this
}

// NewModeOptionValuesMappingsWithDefaults instantiates a new ModeOptionValuesMappings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeOptionValuesMappingsWithDefaults() *ModeOptionValuesMappings {
	this := ModeOptionValuesMappings{}
	return &this
}

// GetCompatibleKbVersions returns the CompatibleKbVersions field value if set, zero value otherwise.
func (o *ModeOptionValuesMappings) GetCompatibleKbVersions() []string {
	if o == nil || o.CompatibleKbVersions == nil {
		var ret []string
		return ret
	}
	return o.CompatibleKbVersions
}

// GetCompatibleKbVersionsOk returns a tuple with the CompatibleKbVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOptionValuesMappings) GetCompatibleKbVersionsOk() (*[]string, bool) {
	if o == nil || o.CompatibleKbVersions == nil {
		return nil, false
	}
	return &o.CompatibleKbVersions, true
}

// HasCompatibleKbVersions returns a boolean if a field has been set.
func (o *ModeOptionValuesMappings) HasCompatibleKbVersions() bool {
	return o != nil && o.CompatibleKbVersions != nil
}

// SetCompatibleKbVersions gets a reference to the given []string and assigns it to the CompatibleKbVersions field.
func (o *ModeOptionValuesMappings) SetCompatibleKbVersions(v []string) {
	o.CompatibleKbVersions = v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *ModeOptionValuesMappings) GetMappings() []ModeOptionValuesMappingsMappingsItem {
	if o == nil || o.Mappings == nil {
		var ret []ModeOptionValuesMappingsMappingsItem
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOptionValuesMappings) GetMappingsOk() (*[]ModeOptionValuesMappingsMappingsItem, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *ModeOptionValuesMappings) HasMappings() bool {
	return o != nil && o.Mappings != nil
}

// SetMappings gets a reference to the given []ModeOptionValuesMappingsMappingsItem and assigns it to the Mappings field.
func (o *ModeOptionValuesMappings) SetMappings(v []ModeOptionValuesMappingsMappingsItem) {
	o.Mappings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeOptionValuesMappings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CompatibleKbVersions != nil {
		toSerialize["compatibleKBVersions"] = o.CompatibleKbVersions
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeOptionValuesMappings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompatibleKbVersions []string                               `json:"compatibleKBVersions,omitempty"`
		Mappings             []ModeOptionValuesMappingsMappingsItem `json:"mappings,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"compatibleKBVersions", "mappings"})
	} else {
		return err
	}
	o.CompatibleKbVersions = all.CompatibleKbVersions
	o.Mappings = all.Mappings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
