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

// MarshalJSON serializes the struct using spec logic.
func (o LicenseOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["opsDefName"] = o.OpsDefName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LicenseOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component  *string `json:"component"`
		OpsDefName *string `json:"opsDefName"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.OpsDefName == nil {
		return fmt.Errorf("required field opsDefName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "opsDefName"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.OpsDefName = *all.OpsDefName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
