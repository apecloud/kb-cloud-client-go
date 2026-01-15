// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsLicense OpsLicense is the payload to update a KubeBlocks cluster license
type OpsLicense struct {
	// component type
	Component *string `json:"component,omitempty"`
	// license ID
	LicenseId string `json:"licenseId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsLicense instantiates a new OpsLicense object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsLicense(licenseId string) *OpsLicense {
	this := OpsLicense{}
	this.LicenseId = licenseId
	return &this
}

// NewOpsLicenseWithDefaults instantiates a new OpsLicense object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsLicenseWithDefaults() *OpsLicense {
	this := OpsLicense{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *OpsLicense) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsLicense) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *OpsLicense) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *OpsLicense) SetComponent(v string) {
	o.Component = &v
}

// GetLicenseId returns the LicenseId field value.
func (o *OpsLicense) GetLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value
// and a boolean to check if the value has been set.
func (o *OpsLicense) GetLicenseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseId, true
}

// SetLicenseId sets field value.
func (o *OpsLicense) SetLicenseId(v string) {
	o.LicenseId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsLicense) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	toSerialize["licenseId"] = o.LicenseId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsLicense) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string `json:"component,omitempty"`
		LicenseId *string `json:"licenseId"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.LicenseId == nil {
		return fmt.Errorf("required field licenseId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "licenseId"})
	} else {
		return err
	}
	o.Component = all.Component
	o.LicenseId = *all.LicenseId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
