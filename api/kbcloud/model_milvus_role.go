// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MilvusRole Milvus role with native privilege grants.
type MilvusRole struct {
	// Milvus role name.
	Name string `json:"name"`
	// Native Milvus privilege grants owned by this role.
	Grants []MilvusRoleGrant `json:"grants"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMilvusRole instantiates a new MilvusRole object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMilvusRole(name string, grants []MilvusRoleGrant) *MilvusRole {
	this := MilvusRole{}
	this.Name = name
	this.Grants = grants
	return &this
}

// NewMilvusRoleWithDefaults instantiates a new MilvusRole object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMilvusRoleWithDefaults() *MilvusRole {
	this := MilvusRole{}
	return &this
}

// GetName returns the Name field value.
func (o *MilvusRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MilvusRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MilvusRole) SetName(v string) {
	o.Name = v
}

// GetGrants returns the Grants field value.
func (o *MilvusRole) GetGrants() []MilvusRoleGrant {
	if o == nil {
		var ret []MilvusRoleGrant
		return ret
	}
	return o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value
// and a boolean to check if the value has been set.
func (o *MilvusRole) GetGrantsOk() (*[]MilvusRoleGrant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grants, true
}

// SetGrants sets field value.
func (o *MilvusRole) SetGrants(v []MilvusRoleGrant) {
	o.Grants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MilvusRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["grants"] = o.Grants

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MilvusRole) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string            `json:"name"`
		Grants *[]MilvusRoleGrant `json:"grants"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Grants == nil {
		return fmt.Errorf("required field grants missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "grants"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Grants = *all.Grants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
