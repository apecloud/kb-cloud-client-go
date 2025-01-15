// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Iac Create environment by IaC
type Iac struct {
	// Role used to do AssumeRole action
	Role string `json:"role"`
	// Desired node size
	DesiredNodeSize *int32 `json:"desiredNodeSize,omitempty"`
	// Computing instance type
	InstanceType *string `json:"instanceType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIac instantiates a new Iac object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIac(role string) *Iac {
	this := Iac{}
	this.Role = role
	return &this
}

// NewIacWithDefaults instantiates a new Iac object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIacWithDefaults() *Iac {
	this := Iac{}
	return &this
}

// GetRole returns the Role field value.
func (o *Iac) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Iac) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *Iac) SetRole(v string) {
	o.Role = v
}

// GetDesiredNodeSize returns the DesiredNodeSize field value if set, zero value otherwise.
func (o *Iac) GetDesiredNodeSize() int32 {
	if o == nil || o.DesiredNodeSize == nil {
		var ret int32
		return ret
	}
	return *o.DesiredNodeSize
}

// GetDesiredNodeSizeOk returns a tuple with the DesiredNodeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iac) GetDesiredNodeSizeOk() (*int32, bool) {
	if o == nil || o.DesiredNodeSize == nil {
		return nil, false
	}
	return o.DesiredNodeSize, true
}

// HasDesiredNodeSize returns a boolean if a field has been set.
func (o *Iac) HasDesiredNodeSize() bool {
	return o != nil && o.DesiredNodeSize != nil
}

// SetDesiredNodeSize gets a reference to the given int32 and assigns it to the DesiredNodeSize field.
func (o *Iac) SetDesiredNodeSize(v int32) {
	o.DesiredNodeSize = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *Iac) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iac) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *Iac) HasInstanceType() bool {
	return o != nil && o.InstanceType != nil
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *Iac) SetInstanceType(v string) {
	o.InstanceType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Iac) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["role"] = o.Role
	if o.DesiredNodeSize != nil {
		toSerialize["desiredNodeSize"] = o.DesiredNodeSize
	}
	if o.InstanceType != nil {
		toSerialize["instanceType"] = o.InstanceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Iac) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role            *string `json:"role"`
		DesiredNodeSize *int32  `json:"desiredNodeSize,omitempty"`
		InstanceType    *string `json:"instanceType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "desiredNodeSize", "instanceType"})
	} else {
		return err
	}
	o.Role = *all.Role
	o.DesiredNodeSize = all.DesiredNodeSize
	o.InstanceType = all.InstanceType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
