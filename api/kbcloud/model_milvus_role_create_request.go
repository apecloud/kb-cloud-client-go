// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MilvusRoleCreateRequest Milvus role create request.
type MilvusRoleCreateRequest struct {
	// Milvus role name.
	Name string `json:"name"`
	// Native Milvus privilege grants owned by this role.
	Grants []MilvusRoleGrant `json:"grants,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMilvusRoleCreateRequest instantiates a new MilvusRoleCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMilvusRoleCreateRequest(name string) *MilvusRoleCreateRequest {
	this := MilvusRoleCreateRequest{}
	this.Name = name
	return &this
}

// NewMilvusRoleCreateRequestWithDefaults instantiates a new MilvusRoleCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMilvusRoleCreateRequestWithDefaults() *MilvusRoleCreateRequest {
	this := MilvusRoleCreateRequest{}
	return &this
}

// GetName returns the Name field value.
func (o *MilvusRoleCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MilvusRoleCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MilvusRoleCreateRequest) SetName(v string) {
	o.Name = v
}

// GetGrants returns the Grants field value if set, zero value otherwise.
func (o *MilvusRoleCreateRequest) GetGrants() []MilvusRoleGrant {
	if o == nil || o.Grants == nil {
		var ret []MilvusRoleGrant
		return ret
	}
	return o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilvusRoleCreateRequest) GetGrantsOk() (*[]MilvusRoleGrant, bool) {
	if o == nil || o.Grants == nil {
		return nil, false
	}
	return &o.Grants, true
}

// HasGrants returns a boolean if a field has been set.
func (o *MilvusRoleCreateRequest) HasGrants() bool {
	return o != nil && o.Grants != nil
}

// SetGrants gets a reference to the given []MilvusRoleGrant and assigns it to the Grants field.
func (o *MilvusRoleCreateRequest) SetGrants(v []MilvusRoleGrant) {
	o.Grants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MilvusRoleCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.Grants != nil {
		toSerialize["grants"] = o.Grants
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MilvusRoleCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string           `json:"name"`
		Grants []MilvusRoleGrant `json:"grants,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "grants"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Grants = all.Grants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
