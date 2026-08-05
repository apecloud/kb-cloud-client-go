// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// MilvusAccountRoles Milvus account role update request.
type MilvusAccountRoles struct {
	// Generic account role. SUPERUSER maps to the Milvus admin role.
	Role *string `json:"role,omitempty"`
	// Milvus role names to assign to the user. Existing roles are replaced.
	Roles []string `json:"roles,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMilvusAccountRoles instantiates a new MilvusAccountRoles object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMilvusAccountRoles() *MilvusAccountRoles {
	this := MilvusAccountRoles{}
	return &this
}

// NewMilvusAccountRolesWithDefaults instantiates a new MilvusAccountRoles object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMilvusAccountRolesWithDefaults() *MilvusAccountRoles {
	this := MilvusAccountRoles{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *MilvusAccountRoles) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilvusAccountRoles) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *MilvusAccountRoles) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *MilvusAccountRoles) SetRole(v string) {
	o.Role = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *MilvusAccountRoles) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilvusAccountRoles) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *MilvusAccountRoles) HasRoles() bool {
	return o != nil && o.Roles != nil
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *MilvusAccountRoles) SetRoles(v []string) {
	o.Roles = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MilvusAccountRoles) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MilvusAccountRoles) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role  *string  `json:"role,omitempty"`
		Roles []string `json:"roles,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "roles"})
	} else {
		return err
	}
	o.Role = all.Role
	o.Roles = all.Roles

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
