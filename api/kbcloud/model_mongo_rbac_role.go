// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoRBACRole struct {
	Role                *string              `json:"role,omitempty"`
	Db                  *string              `json:"db,omitempty"`
	IsBuiltin           *bool                `json:"isBuiltin,omitempty"`
	Privileges          []MongoRBACPrivilege `json:"privileges,omitempty"`
	Roles               []MongoRBACRoleRef   `json:"roles,omitempty"`
	InheritedRoles      []MongoRBACRoleRef   `json:"inheritedRoles,omitempty"`
	InheritedPrivileges []MongoRBACPrivilege `json:"inheritedPrivileges,omitempty"`
	// Stale-form fingerprint. updateRole has no native CAS; external concurrent writes are not transactional.
	Revision *string `json:"revision,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoRBACRole instantiates a new MongoRBACRole object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoRBACRole() *MongoRBACRole {
	this := MongoRBACRole{}
	return &this
}

// NewMongoRBACRoleWithDefaults instantiates a new MongoRBACRole object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoRBACRoleWithDefaults() *MongoRBACRole {
	this := MongoRBACRole{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *MongoRBACRole) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACRole) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *MongoRBACRole) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *MongoRBACRole) SetRole(v string) {
	o.Role = &v
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *MongoRBACRole) GetDb() string {
	if o == nil || o.Db == nil {
		var ret string
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACRole) GetDbOk() (*string, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *MongoRBACRole) HasDb() bool {
	return o != nil && o.Db != nil
}

// SetDb gets a reference to the given string and assigns it to the Db field.
func (o *MongoRBACRole) SetDb(v string) {
	o.Db = &v
}

// GetIsBuiltin returns the IsBuiltin field value if set, zero value otherwise.
func (o *MongoRBACRole) GetIsBuiltin() bool {
	if o == nil || o.IsBuiltin == nil {
		var ret bool
		return ret
	}
	return *o.IsBuiltin
}

// GetIsBuiltinOk returns a tuple with the IsBuiltin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACRole) GetIsBuiltinOk() (*bool, bool) {
	if o == nil || o.IsBuiltin == nil {
		return nil, false
	}
	return o.IsBuiltin, true
}

// HasIsBuiltin returns a boolean if a field has been set.
func (o *MongoRBACRole) HasIsBuiltin() bool {
	return o != nil && o.IsBuiltin != nil
}

// SetIsBuiltin gets a reference to the given bool and assigns it to the IsBuiltin field.
func (o *MongoRBACRole) SetIsBuiltin(v bool) {
	o.IsBuiltin = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *MongoRBACRole) GetPrivileges() []MongoRBACPrivilege {
	if o == nil || o.Privileges == nil {
		var ret []MongoRBACPrivilege
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACRole) GetPrivilegesOk() (*[]MongoRBACPrivilege, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *MongoRBACRole) HasPrivileges() bool {
	return o != nil && o.Privileges != nil
}

// SetPrivileges gets a reference to the given []MongoRBACPrivilege and assigns it to the Privileges field.
func (o *MongoRBACRole) SetPrivileges(v []MongoRBACPrivilege) {
	o.Privileges = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *MongoRBACRole) GetRoles() []MongoRBACRoleRef {
	if o == nil || o.Roles == nil {
		var ret []MongoRBACRoleRef
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACRole) GetRolesOk() (*[]MongoRBACRoleRef, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *MongoRBACRole) HasRoles() bool {
	return o != nil && o.Roles != nil
}

// SetRoles gets a reference to the given []MongoRBACRoleRef and assigns it to the Roles field.
func (o *MongoRBACRole) SetRoles(v []MongoRBACRoleRef) {
	o.Roles = v
}

// GetInheritedRoles returns the InheritedRoles field value if set, zero value otherwise.
func (o *MongoRBACRole) GetInheritedRoles() []MongoRBACRoleRef {
	if o == nil || o.InheritedRoles == nil {
		var ret []MongoRBACRoleRef
		return ret
	}
	return o.InheritedRoles
}

// GetInheritedRolesOk returns a tuple with the InheritedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACRole) GetInheritedRolesOk() (*[]MongoRBACRoleRef, bool) {
	if o == nil || o.InheritedRoles == nil {
		return nil, false
	}
	return &o.InheritedRoles, true
}

// HasInheritedRoles returns a boolean if a field has been set.
func (o *MongoRBACRole) HasInheritedRoles() bool {
	return o != nil && o.InheritedRoles != nil
}

// SetInheritedRoles gets a reference to the given []MongoRBACRoleRef and assigns it to the InheritedRoles field.
func (o *MongoRBACRole) SetInheritedRoles(v []MongoRBACRoleRef) {
	o.InheritedRoles = v
}

// GetInheritedPrivileges returns the InheritedPrivileges field value if set, zero value otherwise.
func (o *MongoRBACRole) GetInheritedPrivileges() []MongoRBACPrivilege {
	if o == nil || o.InheritedPrivileges == nil {
		var ret []MongoRBACPrivilege
		return ret
	}
	return o.InheritedPrivileges
}

// GetInheritedPrivilegesOk returns a tuple with the InheritedPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACRole) GetInheritedPrivilegesOk() (*[]MongoRBACPrivilege, bool) {
	if o == nil || o.InheritedPrivileges == nil {
		return nil, false
	}
	return &o.InheritedPrivileges, true
}

// HasInheritedPrivileges returns a boolean if a field has been set.
func (o *MongoRBACRole) HasInheritedPrivileges() bool {
	return o != nil && o.InheritedPrivileges != nil
}

// SetInheritedPrivileges gets a reference to the given []MongoRBACPrivilege and assigns it to the InheritedPrivileges field.
func (o *MongoRBACRole) SetInheritedPrivileges(v []MongoRBACPrivilege) {
	o.InheritedPrivileges = v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *MongoRBACRole) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACRole) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *MongoRBACRole) HasRevision() bool {
	return o != nil && o.Revision != nil
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *MongoRBACRole) SetRevision(v string) {
	o.Revision = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoRBACRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Db != nil {
		toSerialize["db"] = o.Db
	}
	if o.IsBuiltin != nil {
		toSerialize["isBuiltin"] = o.IsBuiltin
	}
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.InheritedRoles != nil {
		toSerialize["inheritedRoles"] = o.InheritedRoles
	}
	if o.InheritedPrivileges != nil {
		toSerialize["inheritedPrivileges"] = o.InheritedPrivileges
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoRBACRole) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role                *string              `json:"role,omitempty"`
		Db                  *string              `json:"db,omitempty"`
		IsBuiltin           *bool                `json:"isBuiltin,omitempty"`
		Privileges          []MongoRBACPrivilege `json:"privileges,omitempty"`
		Roles               []MongoRBACRoleRef   `json:"roles,omitempty"`
		InheritedRoles      []MongoRBACRoleRef   `json:"inheritedRoles,omitempty"`
		InheritedPrivileges []MongoRBACPrivilege `json:"inheritedPrivileges,omitempty"`
		Revision            *string              `json:"revision,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "db", "isBuiltin", "privileges", "roles", "inheritedRoles", "inheritedPrivileges", "revision"})
	} else {
		return err
	}
	o.Role = all.Role
	o.Db = all.Db
	o.IsBuiltin = all.IsBuiltin
	o.Privileges = all.Privileges
	o.Roles = all.Roles
	o.InheritedRoles = all.InheritedRoles
	o.InheritedPrivileges = all.InheritedPrivileges
	o.Revision = all.Revision

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
