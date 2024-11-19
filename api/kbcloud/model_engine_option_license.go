// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

type EngineOptionLicense struct {
	// Indicate whether the current cluster requires users to input the license
	Required *bool `json:"required,omitempty"`
	// support to refer the cluster name with variable ${clusterName}
	SecretName string `json:"secretName"`
	// the license file name
	FileName   string          `json:"fileName"`
	Components []LicenseOption `json:"components,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEngineOptionLicense instantiates a new EngineOptionLicense object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEngineOptionLicense(secretName string, fileName string) *EngineOptionLicense {
	this := EngineOptionLicense{}
	this.SecretName = secretName
	this.FileName = fileName
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

// GetFileName returns the FileName field value.
func (o *EngineOptionLicense) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *EngineOptionLicense) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value.
func (o *EngineOptionLicense) SetFileName(v string) {
	o.FileName = v
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
	toSerialize["fileName"] = o.FileName
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
		Required   *bool           `json:"required,omitempty"`
		SecretName *string         `json:"secretName"`
		FileName   *string         `json:"fileName"`
		Components []LicenseOption `json:"components,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SecretName == nil {
		return fmt.Errorf("required field secretName missing")
	}
	if all.FileName == nil {
		return fmt.Errorf("required field fileName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"required", "secretName", "fileName", "components"})
	} else {
		return err
	}
	o.Required = all.Required
	o.SecretName = *all.SecretName
	o.FileName = *all.FileName
	o.Components = all.Components

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
