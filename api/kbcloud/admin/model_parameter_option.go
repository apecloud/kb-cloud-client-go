// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ParameterOption struct {
	Component string            `json:"component"`
	Configs   []ParameterConfig `json:"configs"`
	// deprecated
	Versions  []string `json:"versions"`
	ExportTpl bool     `json:"exportTpl"`
	// a alias with major version.
	Family string `json:"family"`
	// match the major version
	MajorVersion          *string              `json:"majorVersion,omitempty"`
	DefaultTplName        string               `json:"defaultTplName"`
	DefaultTplDescription LocalizedDescription `json:"defaultTplDescription"`
	DisableHa             *bool                `json:"disableHA,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewParameterOption instantiates a new ParameterOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewParameterOption(component string, configs []ParameterConfig, versions []string, exportTpl bool, family string, defaultTplName string, defaultTplDescription LocalizedDescription) *ParameterOption {
	this := ParameterOption{}
	this.Component = component
	this.Configs = configs
	this.Versions = versions
	this.ExportTpl = exportTpl
	this.Family = family
	this.DefaultTplName = defaultTplName
	this.DefaultTplDescription = defaultTplDescription
	var disableHa bool = false
	this.DisableHa = &disableHa
	return &this
}

// NewParameterOptionWithDefaults instantiates a new ParameterOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewParameterOptionWithDefaults() *ParameterOption {
	this := ParameterOption{}
	var disableHa bool = false
	this.DisableHa = &disableHa
	return &this
}

// GetComponent returns the Component field value.
func (o *ParameterOption) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *ParameterOption) SetComponent(v string) {
	o.Component = v
}

// GetConfigs returns the Configs field value.
func (o *ParameterOption) GetConfigs() []ParameterConfig {
	if o == nil {
		var ret []ParameterConfig
		return ret
	}
	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetConfigsOk() (*[]ParameterConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configs, true
}

// SetConfigs sets field value.
func (o *ParameterOption) SetConfigs(v []ParameterConfig) {
	o.Configs = v
}

// GetVersions returns the Versions field value.
func (o *ParameterOption) GetVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Versions, true
}

// SetVersions sets field value.
func (o *ParameterOption) SetVersions(v []string) {
	o.Versions = v
}

// GetExportTpl returns the ExportTpl field value.
func (o *ParameterOption) GetExportTpl() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ExportTpl
}

// GetExportTplOk returns a tuple with the ExportTpl field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetExportTplOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportTpl, true
}

// SetExportTpl sets field value.
func (o *ParameterOption) SetExportTpl(v bool) {
	o.ExportTpl = v
}

// GetFamily returns the Family field value.
func (o *ParameterOption) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value.
func (o *ParameterOption) SetFamily(v string) {
	o.Family = v
}

// GetMajorVersion returns the MajorVersion field value if set, zero value otherwise.
func (o *ParameterOption) GetMajorVersion() string {
	if o == nil || o.MajorVersion == nil {
		var ret string
		return ret
	}
	return *o.MajorVersion
}

// GetMajorVersionOk returns a tuple with the MajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetMajorVersionOk() (*string, bool) {
	if o == nil || o.MajorVersion == nil {
		return nil, false
	}
	return o.MajorVersion, true
}

// HasMajorVersion returns a boolean if a field has been set.
func (o *ParameterOption) HasMajorVersion() bool {
	return o != nil && o.MajorVersion != nil
}

// SetMajorVersion gets a reference to the given string and assigns it to the MajorVersion field.
func (o *ParameterOption) SetMajorVersion(v string) {
	o.MajorVersion = &v
}

// GetDefaultTplName returns the DefaultTplName field value.
func (o *ParameterOption) GetDefaultTplName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefaultTplName
}

// GetDefaultTplNameOk returns a tuple with the DefaultTplName field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetDefaultTplNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultTplName, true
}

// SetDefaultTplName sets field value.
func (o *ParameterOption) SetDefaultTplName(v string) {
	o.DefaultTplName = v
}

// GetDefaultTplDescription returns the DefaultTplDescription field value.
func (o *ParameterOption) GetDefaultTplDescription() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.DefaultTplDescription
}

// GetDefaultTplDescriptionOk returns a tuple with the DefaultTplDescription field value
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetDefaultTplDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultTplDescription, true
}

// SetDefaultTplDescription sets field value.
func (o *ParameterOption) SetDefaultTplDescription(v LocalizedDescription) {
	o.DefaultTplDescription = v
}

// GetDisableHa returns the DisableHa field value if set, zero value otherwise.
func (o *ParameterOption) GetDisableHa() bool {
	if o == nil || o.DisableHa == nil {
		var ret bool
		return ret
	}
	return *o.DisableHa
}

// GetDisableHaOk returns a tuple with the DisableHa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterOption) GetDisableHaOk() (*bool, bool) {
	if o == nil || o.DisableHa == nil {
		return nil, false
	}
	return o.DisableHa, true
}

// HasDisableHa returns a boolean if a field has been set.
func (o *ParameterOption) HasDisableHa() bool {
	return o != nil && o.DisableHa != nil
}

// SetDisableHa gets a reference to the given bool and assigns it to the DisableHa field.
func (o *ParameterOption) SetDisableHa(v bool) {
	o.DisableHa = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ParameterOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["configs"] = o.Configs
	toSerialize["versions"] = o.Versions
	toSerialize["exportTpl"] = o.ExportTpl
	toSerialize["family"] = o.Family
	if o.MajorVersion != nil {
		toSerialize["majorVersion"] = o.MajorVersion
	}
	toSerialize["defaultTplName"] = o.DefaultTplName
	toSerialize["defaultTplDescription"] = o.DefaultTplDescription
	if o.DisableHa != nil {
		toSerialize["disableHA"] = o.DisableHa
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ParameterOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component             *string               `json:"component"`
		Configs               *[]ParameterConfig    `json:"configs"`
		Versions              *[]string             `json:"versions"`
		ExportTpl             *bool                 `json:"exportTpl"`
		Family                *string               `json:"family"`
		MajorVersion          *string               `json:"majorVersion,omitempty"`
		DefaultTplName        *string               `json:"defaultTplName"`
		DefaultTplDescription *LocalizedDescription `json:"defaultTplDescription"`
		DisableHa             *bool                 `json:"disableHA,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Configs == nil {
		return fmt.Errorf("required field configs missing")
	}
	if all.Versions == nil {
		return fmt.Errorf("required field versions missing")
	}
	if all.ExportTpl == nil {
		return fmt.Errorf("required field exportTpl missing")
	}
	if all.Family == nil {
		return fmt.Errorf("required field family missing")
	}
	if all.DefaultTplName == nil {
		return fmt.Errorf("required field defaultTplName missing")
	}
	if all.DefaultTplDescription == nil {
		return fmt.Errorf("required field defaultTplDescription missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "configs", "versions", "exportTpl", "family", "majorVersion", "defaultTplName", "defaultTplDescription", "disableHA"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Component = *all.Component
	o.Configs = *all.Configs
	o.Versions = *all.Versions
	o.ExportTpl = *all.ExportTpl
	o.Family = *all.Family
	o.MajorVersion = all.MajorVersion
	o.DefaultTplName = *all.DefaultTplName
	if all.DefaultTplDescription.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DefaultTplDescription = *all.DefaultTplDescription
	o.DisableHa = all.DisableHa

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
