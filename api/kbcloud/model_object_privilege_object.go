// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeObject struct {
	Target ObjectPrivilegeTarget `json:"target"`
	// Display identity, including PostgreSQL routine input argument types. Use target for operations; identity is never accepted as SQL.
	Identity string `json:"identity"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeObject instantiates a new ObjectPrivilegeObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeObject(target ObjectPrivilegeTarget, identity string) *ObjectPrivilegeObject {
	this := ObjectPrivilegeObject{}
	this.Target = target
	this.Identity = identity
	return &this
}

// NewObjectPrivilegeObjectWithDefaults instantiates a new ObjectPrivilegeObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeObjectWithDefaults() *ObjectPrivilegeObject {
	this := ObjectPrivilegeObject{}
	return &this
}

// GetTarget returns the Target field value.
func (o *ObjectPrivilegeObject) GetTarget() ObjectPrivilegeTarget {
	if o == nil {
		var ret ObjectPrivilegeTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeObject) GetTargetOk() (*ObjectPrivilegeTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ObjectPrivilegeObject) SetTarget(v ObjectPrivilegeTarget) {
	o.Target = v
}

// GetIdentity returns the Identity field value.
func (o *ObjectPrivilegeObject) GetIdentity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeObject) GetIdentityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value.
func (o *ObjectPrivilegeObject) SetIdentity(v string) {
	o.Identity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["target"] = o.Target
	toSerialize["identity"] = o.Identity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target   *ObjectPrivilegeTarget `json:"target"`
		Identity *string                `json:"identity"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Identity == nil {
		return fmt.Errorf("required field identity missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"target", "identity"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target
	o.Identity = *all.Identity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
