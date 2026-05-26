// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentCredentialUsageHint struct {
	// The intended usage of the environment credential.
	Usage EnvironmentCredentialUsage `json:"usage"`
	// Text hints describing the permissions the credential should have for this usage.
	PermissionTips []string `json:"permissionTips"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentCredentialUsageHint instantiates a new EnvironmentCredentialUsageHint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentCredentialUsageHint(usage EnvironmentCredentialUsage, permissionTips []string) *EnvironmentCredentialUsageHint {
	this := EnvironmentCredentialUsageHint{}
	this.Usage = usage
	this.PermissionTips = permissionTips
	return &this
}

// NewEnvironmentCredentialUsageHintWithDefaults instantiates a new EnvironmentCredentialUsageHint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentCredentialUsageHintWithDefaults() *EnvironmentCredentialUsageHint {
	this := EnvironmentCredentialUsageHint{}
	return &this
}

// GetUsage returns the Usage field value.
func (o *EnvironmentCredentialUsageHint) GetUsage() EnvironmentCredentialUsage {
	if o == nil {
		var ret EnvironmentCredentialUsage
		return ret
	}
	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialUsageHint) GetUsageOk() (*EnvironmentCredentialUsage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usage, true
}

// SetUsage sets field value.
func (o *EnvironmentCredentialUsageHint) SetUsage(v EnvironmentCredentialUsage) {
	o.Usage = v
}

// GetPermissionTips returns the PermissionTips field value.
func (o *EnvironmentCredentialUsageHint) GetPermissionTips() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PermissionTips
}

// GetPermissionTipsOk returns a tuple with the PermissionTips field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialUsageHint) GetPermissionTipsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionTips, true
}

// SetPermissionTips sets field value.
func (o *EnvironmentCredentialUsageHint) SetPermissionTips(v []string) {
	o.PermissionTips = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentCredentialUsageHint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["usage"] = o.Usage
	toSerialize["permissionTips"] = o.PermissionTips

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentCredentialUsageHint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Usage          *EnvironmentCredentialUsage `json:"usage"`
		PermissionTips *[]string                   `json:"permissionTips"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Usage == nil {
		return fmt.Errorf("required field usage missing")
	}
	if all.PermissionTips == nil {
		return fmt.Errorf("required field permissionTips missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"usage", "permissionTips"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Usage.IsValid() {
		hasInvalidField = true
	} else {
		o.Usage = *all.Usage
	}
	o.PermissionTips = *all.PermissionTips

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
