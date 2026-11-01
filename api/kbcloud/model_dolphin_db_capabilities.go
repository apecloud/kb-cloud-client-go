// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBCapabilities struct {
	EngineVersion string                         `json:"engineVersion"`
	Privileges    []DolphinDBPrivilegeCapability `json:"privileges"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBCapabilities instantiates a new DolphinDBCapabilities object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBCapabilities(engineVersion string, privileges []DolphinDBPrivilegeCapability) *DolphinDBCapabilities {
	this := DolphinDBCapabilities{}
	this.EngineVersion = engineVersion
	this.Privileges = privileges
	return &this
}

// NewDolphinDBCapabilitiesWithDefaults instantiates a new DolphinDBCapabilities object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBCapabilitiesWithDefaults() *DolphinDBCapabilities {
	this := DolphinDBCapabilities{}
	return &this
}

// GetEngineVersion returns the EngineVersion field value.
func (o *DolphinDBCapabilities) GetEngineVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineVersion
}

// GetEngineVersionOk returns a tuple with the EngineVersion field value
// and a boolean to check if the value has been set.
func (o *DolphinDBCapabilities) GetEngineVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineVersion, true
}

// SetEngineVersion sets field value.
func (o *DolphinDBCapabilities) SetEngineVersion(v string) {
	o.EngineVersion = v
}

// GetPrivileges returns the Privileges field value.
func (o *DolphinDBCapabilities) GetPrivileges() []DolphinDBPrivilegeCapability {
	if o == nil {
		var ret []DolphinDBPrivilegeCapability
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *DolphinDBCapabilities) GetPrivilegesOk() (*[]DolphinDBPrivilegeCapability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// SetPrivileges sets field value.
func (o *DolphinDBCapabilities) SetPrivileges(v []DolphinDBPrivilegeCapability) {
	o.Privileges = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engineVersion"] = o.EngineVersion
	toSerialize["privileges"] = o.Privileges

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBCapabilities) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EngineVersion *string                         `json:"engineVersion"`
		Privileges    *[]DolphinDBPrivilegeCapability `json:"privileges"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.EngineVersion == nil {
		return fmt.Errorf("required field engineVersion missing")
	}
	if all.Privileges == nil {
		return fmt.Errorf("required field privileges missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engineVersion", "privileges"})
	} else {
		return err
	}
	o.EngineVersion = *all.EngineVersion
	o.Privileges = *all.Privileges

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
