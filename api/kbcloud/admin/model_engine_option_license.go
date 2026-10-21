// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EngineOptionLicense struct {
	// Indicate whether the current cluster requires users to input the license
	Required *bool `json:"required,omitempty"`
	// support to refer the cluster name with variable ${clusterName}
	SecretName string `json:"secretName"`
	// the license file name
	FileName *string `json:"fileName,omitempty"`
	// Allowed engine license types. Omit to retain unrestricted license types.
	AllowedTypes []string `json:"allowedTypes,omitempty"`
	// Engine versions that support license installation. Entries match an exact service version or its major version according to the main component version rule. Omit or leave empty for unrestricted versions.
	Versions []string `json:"versions,omitempty"`
	// Effective cluster network modes that support license installation. Omit or leave empty for unrestricted network modes.
	NetworkModes []NetworkMode `json:"networkModes,omitempty"`
	// Explicit license file upload slots. Each item maps one UI upload field to the file name written into the license Secret.
	Files []EngineOptionLicenseFile `json:"files,omitempty"`
	// set the custom ops to update license for different components
	Components []LicenseOption `json:"components,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineOptionLicense instantiates a new EngineOptionLicense object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineOptionLicense(secretName string) *EngineOptionLicense {
	this := EngineOptionLicense{}
	this.SecretName = secretName
	return &this
}

// NewEngineOptionLicenseWithDefaults instantiates a new EngineOptionLicense object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEngineOptionLicenseWithDefaults() *EngineOptionLicense {
	this := EngineOptionLicense{}
	return &this
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *EngineOptionLicense) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionLicense) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *EngineOptionLicense) HasRequired() bool {
	return o != nil && o.Required != nil
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *EngineOptionLicense) SetRequired(v bool) {
	o.Required = &v
}

// GetSecretName returns the SecretName field value.
func (o *EngineOptionLicense) GetSecretName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SecretName
}

// GetSecretNameOk returns a tuple with the SecretName field value
// and a boolean to check if the value has been set.
func (o *EngineOptionLicense) GetSecretNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretName, true
}

// SetSecretName sets field value.
func (o *EngineOptionLicense) SetSecretName(v string) {
	o.SecretName = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *EngineOptionLicense) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionLicense) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *EngineOptionLicense) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *EngineOptionLicense) SetFileName(v string) {
	o.FileName = &v
}

// GetAllowedTypes returns the AllowedTypes field value if set, zero value otherwise.
func (o *EngineOptionLicense) GetAllowedTypes() []string {
	if o == nil || o.AllowedTypes == nil {
		var ret []string
		return ret
	}
	return o.AllowedTypes
}

// GetAllowedTypesOk returns a tuple with the AllowedTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionLicense) GetAllowedTypesOk() (*[]string, bool) {
	if o == nil || o.AllowedTypes == nil {
		return nil, false
	}
	return &o.AllowedTypes, true
}

// HasAllowedTypes returns a boolean if a field has been set.
func (o *EngineOptionLicense) HasAllowedTypes() bool {
	return o != nil && o.AllowedTypes != nil
}

// SetAllowedTypes gets a reference to the given []string and assigns it to the AllowedTypes field.
func (o *EngineOptionLicense) SetAllowedTypes(v []string) {
	o.AllowedTypes = v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *EngineOptionLicense) GetVersions() []string {
	if o == nil || o.Versions == nil {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionLicense) GetVersionsOk() (*[]string, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return &o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *EngineOptionLicense) HasVersions() bool {
	return o != nil && o.Versions != nil
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *EngineOptionLicense) SetVersions(v []string) {
	o.Versions = v
}

// GetNetworkModes returns the NetworkModes field value if set, zero value otherwise.
func (o *EngineOptionLicense) GetNetworkModes() []NetworkMode {
	if o == nil || o.NetworkModes == nil {
		var ret []NetworkMode
		return ret
	}
	return o.NetworkModes
}

// GetNetworkModesOk returns a tuple with the NetworkModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionLicense) GetNetworkModesOk() (*[]NetworkMode, bool) {
	if o == nil || o.NetworkModes == nil {
		return nil, false
	}
	return &o.NetworkModes, true
}

// HasNetworkModes returns a boolean if a field has been set.
func (o *EngineOptionLicense) HasNetworkModes() bool {
	return o != nil && o.NetworkModes != nil
}

// SetNetworkModes gets a reference to the given []NetworkMode and assigns it to the NetworkModes field.
func (o *EngineOptionLicense) SetNetworkModes(v []NetworkMode) {
	o.NetworkModes = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *EngineOptionLicense) GetFiles() []EngineOptionLicenseFile {
	if o == nil || o.Files == nil {
		var ret []EngineOptionLicenseFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionLicense) GetFilesOk() (*[]EngineOptionLicenseFile, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return &o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *EngineOptionLicense) HasFiles() bool {
	return o != nil && o.Files != nil
}

// SetFiles gets a reference to the given []EngineOptionLicenseFile and assigns it to the Files field.
func (o *EngineOptionLicense) SetFiles(v []EngineOptionLicenseFile) {
	o.Files = v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *EngineOptionLicense) GetComponents() []LicenseOption {
	if o == nil || o.Components == nil {
		var ret []LicenseOption
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineOptionLicense) GetComponentsOk() (*[]LicenseOption, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *EngineOptionLicense) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []LicenseOption and assigns it to the Components field.
func (o *EngineOptionLicense) SetComponents(v []LicenseOption) {
	o.Components = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EngineOptionLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	toSerialize["secretName"] = o.SecretName
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}
	if o.AllowedTypes != nil {
		toSerialize["allowedTypes"] = o.AllowedTypes
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}
	if o.NetworkModes != nil {
		toSerialize["networkModes"] = o.NetworkModes
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EngineOptionLicense) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Required     *bool                     `json:"required,omitempty"`
		SecretName   *string                   `json:"secretName"`
		FileName     *string                   `json:"fileName,omitempty"`
		AllowedTypes []string                  `json:"allowedTypes,omitempty"`
		Versions     []string                  `json:"versions,omitempty"`
		NetworkModes []NetworkMode             `json:"networkModes,omitempty"`
		Files        []EngineOptionLicenseFile `json:"files,omitempty"`
		Components   []LicenseOption           `json:"components,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SecretName == nil {
		return fmt.Errorf("required field secretName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"required", "secretName", "fileName", "allowedTypes", "versions", "networkModes", "files", "components"})
	} else {
		return err
	}
	o.Required = all.Required
	o.SecretName = *all.SecretName
	o.FileName = all.FileName
	o.AllowedTypes = all.AllowedTypes
	o.Versions = all.Versions
	o.NetworkModes = all.NetworkModes
	o.Files = all.Files
	o.Components = all.Components

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
