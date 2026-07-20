// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentModuleInstallRequest Request to check or install an environment module
type EnvironmentModuleInstallRequest struct {
	// When true, only run installation checks. When false, run the same checks and submit an installation task if all checks pass.
	DryRun bool `json:"dryRun"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentModuleInstallRequest instantiates a new EnvironmentModuleInstallRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentModuleInstallRequest(dryRun bool) *EnvironmentModuleInstallRequest {
	this := EnvironmentModuleInstallRequest{}
	this.DryRun = dryRun
	return &this
}

// NewEnvironmentModuleInstallRequestWithDefaults instantiates a new EnvironmentModuleInstallRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentModuleInstallRequestWithDefaults() *EnvironmentModuleInstallRequest {
	this := EnvironmentModuleInstallRequest{}
	return &this
}

// GetDryRun returns the DryRun field value.
func (o *EnvironmentModuleInstallRequest) GetDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value
// and a boolean to check if the value has been set.
func (o *EnvironmentModuleInstallRequest) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DryRun, true
}

// SetDryRun sets field value.
func (o *EnvironmentModuleInstallRequest) SetDryRun(v bool) {
	o.DryRun = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentModuleInstallRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["dryRun"] = o.DryRun

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentModuleInstallRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DryRun *bool `json:"dryRun"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DryRun == nil {
		return fmt.Errorf("required field dryRun missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"dryRun"})
	} else {
		return err
	}
	o.DryRun = *all.DryRun

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
