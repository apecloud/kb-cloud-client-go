// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineDefinitionDetail struct {
	DefinitionName *string `json:"definitionName,omitempty"`
	// definition supported available modes, empty means all modes are supported
	AvailableModes []string `json:"availableModes,omitempty"`
	// definition supported available versions, empty means all versions are supported
	AvailableVersions []EngineDefinitionVersion `json:"availableVersions,omitempty"`
	ExtraCfgs         []ExtraConfig             `json:"extraCfgs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineDefinitionDetail instantiates a new EngineDefinitionDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineDefinitionDetail() *EngineDefinitionDetail {
	this := EngineDefinitionDetail{}
	return &this
}

// NewEngineDefinitionDetailWithDefaults instantiates a new EngineDefinitionDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineDefinitionDetailWithDefaults() *EngineDefinitionDetail {
	this := EngineDefinitionDetail{}
	return &this
}

// GetDefinitionName returns the DefinitionName field value if set, zero value otherwise.
func (o *EngineDefinitionDetail) GetDefinitionName() string {
	if o == nil || o.DefinitionName == nil {
		var ret string
		return ret
	}
	return *o.DefinitionName
}

// GetDefinitionNameOk returns a tuple with the DefinitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineDefinitionDetail) GetDefinitionNameOk() (*string, bool) {
	if o == nil || o.DefinitionName == nil {
		return nil, false
	}
	return o.DefinitionName, true
}

// HasDefinitionName returns a boolean if a field has been set.
func (o *EngineDefinitionDetail) HasDefinitionName() bool {
	return o != nil && o.DefinitionName != nil
}

// SetDefinitionName gets a reference to the given string and assigns it to the DefinitionName field.
func (o *EngineDefinitionDetail) SetDefinitionName(v string) {
	o.DefinitionName = &v
}

// GetAvailableModes returns the AvailableModes field value if set, zero value otherwise.
func (o *EngineDefinitionDetail) GetAvailableModes() []string {
	if o == nil || o.AvailableModes == nil {
		var ret []string
		return ret
	}
	return o.AvailableModes
}

// GetAvailableModesOk returns a tuple with the AvailableModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineDefinitionDetail) GetAvailableModesOk() (*[]string, bool) {
	if o == nil || o.AvailableModes == nil {
		return nil, false
	}
	return &o.AvailableModes, true
}

// HasAvailableModes returns a boolean if a field has been set.
func (o *EngineDefinitionDetail) HasAvailableModes() bool {
	return o != nil && o.AvailableModes != nil
}

// SetAvailableModes gets a reference to the given []string and assigns it to the AvailableModes field.
func (o *EngineDefinitionDetail) SetAvailableModes(v []string) {
	o.AvailableModes = v
}

// GetAvailableVersions returns the AvailableVersions field value if set, zero value otherwise.
func (o *EngineDefinitionDetail) GetAvailableVersions() []EngineDefinitionVersion {
	if o == nil || o.AvailableVersions == nil {
		var ret []EngineDefinitionVersion
		return ret
	}
	return o.AvailableVersions
}

// GetAvailableVersionsOk returns a tuple with the AvailableVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineDefinitionDetail) GetAvailableVersionsOk() (*[]EngineDefinitionVersion, bool) {
	if o == nil || o.AvailableVersions == nil {
		return nil, false
	}
	return &o.AvailableVersions, true
}

// HasAvailableVersions returns a boolean if a field has been set.
func (o *EngineDefinitionDetail) HasAvailableVersions() bool {
	return o != nil && o.AvailableVersions != nil
}

// SetAvailableVersions gets a reference to the given []EngineDefinitionVersion and assigns it to the AvailableVersions field.
func (o *EngineDefinitionDetail) SetAvailableVersions(v []EngineDefinitionVersion) {
	o.AvailableVersions = v
}

// GetExtraCfgs returns the ExtraCfgs field value if set, zero value otherwise.
func (o *EngineDefinitionDetail) GetExtraCfgs() []ExtraConfig {
	if o == nil || o.ExtraCfgs == nil {
		var ret []ExtraConfig
		return ret
	}
	return o.ExtraCfgs
}

// GetExtraCfgsOk returns a tuple with the ExtraCfgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineDefinitionDetail) GetExtraCfgsOk() (*[]ExtraConfig, bool) {
	if o == nil || o.ExtraCfgs == nil {
		return nil, false
	}
	return &o.ExtraCfgs, true
}

// HasExtraCfgs returns a boolean if a field has been set.
func (o *EngineDefinitionDetail) HasExtraCfgs() bool {
	return o != nil && o.ExtraCfgs != nil
}

// SetExtraCfgs gets a reference to the given []ExtraConfig and assigns it to the ExtraCfgs field.
func (o *EngineDefinitionDetail) SetExtraCfgs(v []ExtraConfig) {
	o.ExtraCfgs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineDefinitionDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DefinitionName != nil {
		toSerialize["definitionName"] = o.DefinitionName
	}
	if o.AvailableModes != nil {
		toSerialize["availableModes"] = o.AvailableModes
	}
	if o.AvailableVersions != nil {
		toSerialize["availableVersions"] = o.AvailableVersions
	}
	if o.ExtraCfgs != nil {
		toSerialize["extraCfgs"] = o.ExtraCfgs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineDefinitionDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefinitionName    *string                   `json:"definitionName,omitempty"`
		AvailableModes    []string                  `json:"availableModes,omitempty"`
		AvailableVersions []EngineDefinitionVersion `json:"availableVersions,omitempty"`
		ExtraCfgs         []ExtraConfig             `json:"extraCfgs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"definitionName", "availableModes", "availableVersions", "extraCfgs"})
	} else {
		return err
	}
	o.DefinitionName = all.DefinitionName
	o.AvailableModes = all.AvailableModes
	o.AvailableVersions = all.AvailableVersions
	o.ExtraCfgs = all.ExtraCfgs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
