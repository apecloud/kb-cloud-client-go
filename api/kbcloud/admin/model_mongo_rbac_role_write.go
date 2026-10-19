// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MongoRBACRoleWrite struct {
	Role             string               `json:"role"`
	Db               string               `json:"db"`
	Privileges       []MongoRBACPrivilege `json:"privileges"`
	Roles            []MongoRBACRoleRef   `json:"roles"`
	ExpectedRevision *string              `json:"expectedRevision,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoRBACRoleWrite instantiates a new MongoRBACRoleWrite object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoRBACRoleWrite(role string, db string, privileges []MongoRBACPrivilege, roles []MongoRBACRoleRef) *MongoRBACRoleWrite {
	this := MongoRBACRoleWrite{}
	this.Role = role
	this.Db = db
	this.Privileges = privileges
	this.Roles = roles
	return &this
}

// NewMongoRBACRoleWriteWithDefaults instantiates a new MongoRBACRoleWrite object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoRBACRoleWriteWithDefaults() *MongoRBACRoleWrite {
	this := MongoRBACRoleWrite{}
	return &this
}

// GetRole returns the Role field value.
func (o *MongoRBACRoleWrite) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *MongoRBACRoleWrite) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *MongoRBACRoleWrite) SetRole(v string) {
	o.Role = v
}

// GetDb returns the Db field value.
func (o *MongoRBACRoleWrite) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *MongoRBACRoleWrite) GetDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value.
func (o *MongoRBACRoleWrite) SetDb(v string) {
	o.Db = v
}

// GetPrivileges returns the Privileges field value.
func (o *MongoRBACRoleWrite) GetPrivileges() []MongoRBACPrivilege {
	if o == nil {
		var ret []MongoRBACPrivilege
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *MongoRBACRoleWrite) GetPrivilegesOk() (*[]MongoRBACPrivilege, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// SetPrivileges sets field value.
func (o *MongoRBACRoleWrite) SetPrivileges(v []MongoRBACPrivilege) {
	o.Privileges = v
}

// GetRoles returns the Roles field value.
func (o *MongoRBACRoleWrite) GetRoles() []MongoRBACRoleRef {
	if o == nil {
		var ret []MongoRBACRoleRef
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *MongoRBACRoleWrite) GetRolesOk() (*[]MongoRBACRoleRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value.
func (o *MongoRBACRoleWrite) SetRoles(v []MongoRBACRoleRef) {
	o.Roles = v
}

// GetExpectedRevision returns the ExpectedRevision field value if set, zero value otherwise.
func (o *MongoRBACRoleWrite) GetExpectedRevision() string {
	if o == nil || o.ExpectedRevision == nil {
		var ret string
		return ret
	}
	return *o.ExpectedRevision
}

// GetExpectedRevisionOk returns a tuple with the ExpectedRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACRoleWrite) GetExpectedRevisionOk() (*string, bool) {
	if o == nil || o.ExpectedRevision == nil {
		return nil, false
	}
	return o.ExpectedRevision, true
}

// HasExpectedRevision returns a boolean if a field has been set.
func (o *MongoRBACRoleWrite) HasExpectedRevision() bool {
	return o != nil && o.ExpectedRevision != nil
}

// SetExpectedRevision gets a reference to the given string and assigns it to the ExpectedRevision field.
func (o *MongoRBACRoleWrite) SetExpectedRevision(v string) {
	o.ExpectedRevision = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoRBACRoleWrite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["role"] = o.Role
	toSerialize["db"] = o.Db
	toSerialize["privileges"] = o.Privileges
	toSerialize["roles"] = o.Roles
	if o.ExpectedRevision != nil {
		toSerialize["expectedRevision"] = o.ExpectedRevision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoRBACRoleWrite) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role             *string               `json:"role"`
		Db               *string               `json:"db"`
		Privileges       *[]MongoRBACPrivilege `json:"privileges"`
		Roles            *[]MongoRBACRoleRef   `json:"roles"`
		ExpectedRevision *string               `json:"expectedRevision,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.Db == nil {
		return fmt.Errorf("required field db missing")
	}
	if all.Privileges == nil {
		return fmt.Errorf("required field privileges missing")
	}
	if all.Roles == nil {
		return fmt.Errorf("required field roles missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "db", "privileges", "roles", "expectedRevision"})
	} else {
		return err
	}
	o.Role = *all.Role
	o.Db = *all.Db
	o.Privileges = *all.Privileges
	o.Roles = *all.Roles
	o.ExpectedRevision = all.ExpectedRevision

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
