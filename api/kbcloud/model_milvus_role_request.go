// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// MilvusRoleRequest Milvus role upsert request.
type MilvusRoleRequest struct {
	// Native Milvus privilege grants owned by this role. Existing grants are replaced.
	Grants []MilvusRoleGrant `json:"grants,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMilvusRoleRequest instantiates a new MilvusRoleRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMilvusRoleRequest() *MilvusRoleRequest {
	this := MilvusRoleRequest{}
	return &this
}

// NewMilvusRoleRequestWithDefaults instantiates a new MilvusRoleRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMilvusRoleRequestWithDefaults() *MilvusRoleRequest {
	this := MilvusRoleRequest{}
	return &this
}

// GetGrants returns the Grants field value if set, zero value otherwise.
func (o *MilvusRoleRequest) GetGrants() []MilvusRoleGrant {
	if o == nil || o.Grants == nil {
		var ret []MilvusRoleGrant
		return ret
	}
	return o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilvusRoleRequest) GetGrantsOk() (*[]MilvusRoleGrant, bool) {
	if o == nil || o.Grants == nil {
		return nil, false
	}
	return &o.Grants, true
}

// HasGrants returns a boolean if a field has been set.
func (o *MilvusRoleRequest) HasGrants() bool {
	return o != nil && o.Grants != nil
}

// SetGrants gets a reference to the given []MilvusRoleGrant and assigns it to the Grants field.
func (o *MilvusRoleRequest) SetGrants(v []MilvusRoleGrant) {
	o.Grants = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MilvusRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Grants != nil {
		toSerialize["grants"] = o.Grants
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MilvusRoleRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Grants []MilvusRoleGrant `json:"grants,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"grants"})
	} else {
		return err
	}
	o.Grants = all.Grants

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
