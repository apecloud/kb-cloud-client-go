// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type LicenseOption struct {
	Component  string `json:"component"`
	OpsDefName string `json:"opsDefName"`
	// Optional custom Ops parameter receiving a JSON object of candidate file SHA256 digests.
	FileDigestParameter *string `json:"fileDigestParameter,omitempty"`
	// Create a standard restart dependent on successful license installation.
	RestartAfterUpdate *bool `json:"restartAfterUpdate,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLicenseOption instantiates a new LicenseOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLicenseOption(component string, opsDefName string) *LicenseOption {
	this := LicenseOption{}
	this.Component = component
	this.OpsDefName = opsDefName
	return &this
}

// NewLicenseOptionWithDefaults instantiates a new LicenseOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLicenseOptionWithDefaults() *LicenseOption {
	this := LicenseOption{}
	return &this
}

// GetComponent returns the Component field value.
func (o *LicenseOption) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *LicenseOption) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *LicenseOption) SetComponent(v string) {
	o.Component = v
}

// GetOpsDefName returns the OpsDefName field value.
func (o *LicenseOption) GetOpsDefName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OpsDefName
}

// GetOpsDefNameOk returns a tuple with the OpsDefName field value
// and a boolean to check if the value has been set.
func (o *LicenseOption) GetOpsDefNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpsDefName, true
}

// SetOpsDefName sets field value.
func (o *LicenseOption) SetOpsDefName(v string) {
	o.OpsDefName = v
}

// GetFileDigestParameter returns the FileDigestParameter field value if set, zero value otherwise.
func (o *LicenseOption) GetFileDigestParameter() string {
	if o == nil || o.FileDigestParameter == nil {
		var ret string
		return ret
	}
	return *o.FileDigestParameter
}

// GetFileDigestParameterOk returns a tuple with the FileDigestParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseOption) GetFileDigestParameterOk() (*string, bool) {
	if o == nil || o.FileDigestParameter == nil {
		return nil, false
	}
	return o.FileDigestParameter, true
}

// HasFileDigestParameter returns a boolean if a field has been set.
func (o *LicenseOption) HasFileDigestParameter() bool {
	return o != nil && o.FileDigestParameter != nil
}

// SetFileDigestParameter gets a reference to the given string and assigns it to the FileDigestParameter field.
func (o *LicenseOption) SetFileDigestParameter(v string) {
	o.FileDigestParameter = &v
}

// GetRestartAfterUpdate returns the RestartAfterUpdate field value if set, zero value otherwise.
func (o *LicenseOption) GetRestartAfterUpdate() bool {
	if o == nil || o.RestartAfterUpdate == nil {
		var ret bool
		return ret
	}
	return *o.RestartAfterUpdate
}

// GetRestartAfterUpdateOk returns a tuple with the RestartAfterUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseOption) GetRestartAfterUpdateOk() (*bool, bool) {
	if o == nil || o.RestartAfterUpdate == nil {
		return nil, false
	}
	return o.RestartAfterUpdate, true
}

// HasRestartAfterUpdate returns a boolean if a field has been set.
func (o *LicenseOption) HasRestartAfterUpdate() bool {
	return o != nil && o.RestartAfterUpdate != nil
}

// SetRestartAfterUpdate gets a reference to the given bool and assigns it to the RestartAfterUpdate field.
func (o *LicenseOption) SetRestartAfterUpdate(v bool) {
	o.RestartAfterUpdate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LicenseOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["opsDefName"] = o.OpsDefName
	if o.FileDigestParameter != nil {
		toSerialize["fileDigestParameter"] = o.FileDigestParameter
	}
	if o.RestartAfterUpdate != nil {
		toSerialize["restartAfterUpdate"] = o.RestartAfterUpdate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LicenseOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component           *string `json:"component"`
		OpsDefName          *string `json:"opsDefName"`
		FileDigestParameter *string `json:"fileDigestParameter,omitempty"`
		RestartAfterUpdate  *bool   `json:"restartAfterUpdate,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.OpsDefName == nil {
		return fmt.Errorf("required field opsDefName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "opsDefName", "fileDigestParameter", "restartAfterUpdate"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.OpsDefName = *all.OpsDefName
	o.FileDigestParameter = all.FileDigestParameter
	o.RestartAfterUpdate = all.RestartAfterUpdate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
