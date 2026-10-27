// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoRBACUser struct {
	Db                  *string              `json:"db,omitempty"`
	Roles               []MongoRBACRoleRef   `json:"roles,omitempty"`
	InheritedRoles      []MongoRBACRoleRef   `json:"inheritedRoles,omitempty"`
	InheritedPrivileges []MongoRBACPrivilege `json:"inheritedPrivileges,omitempty"`
	User                *string              `json:"user,omitempty"`
	// Platform-protected account; mutations are rejected.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoRBACUser instantiates a new MongoRBACUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoRBACUser() *MongoRBACUser {
	this := MongoRBACUser{}
	return &this
}

// NewMongoRBACUserWithDefaults instantiates a new MongoRBACUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoRBACUserWithDefaults() *MongoRBACUser {
	this := MongoRBACUser{}
	return &this
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *MongoRBACUser) GetDb() string {
	if o == nil || o.Db == nil {
		var ret string
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACUser) GetDbOk() (*string, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *MongoRBACUser) HasDb() bool {
	return o != nil && o.Db != nil
}

// SetDb gets a reference to the given string and assigns it to the Db field.
func (o *MongoRBACUser) SetDb(v string) {
	o.Db = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *MongoRBACUser) GetRoles() []MongoRBACRoleRef {
	if o == nil || o.Roles == nil {
		var ret []MongoRBACRoleRef
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACUser) GetRolesOk() (*[]MongoRBACRoleRef, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *MongoRBACUser) HasRoles() bool {
	return o != nil && o.Roles != nil
}

// SetRoles gets a reference to the given []MongoRBACRoleRef and assigns it to the Roles field.
func (o *MongoRBACUser) SetRoles(v []MongoRBACRoleRef) {
	o.Roles = v
}

// GetInheritedRoles returns the InheritedRoles field value if set, zero value otherwise.
func (o *MongoRBACUser) GetInheritedRoles() []MongoRBACRoleRef {
	if o == nil || o.InheritedRoles == nil {
		var ret []MongoRBACRoleRef
		return ret
	}
	return o.InheritedRoles
}

// GetInheritedRolesOk returns a tuple with the InheritedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACUser) GetInheritedRolesOk() (*[]MongoRBACRoleRef, bool) {
	if o == nil || o.InheritedRoles == nil {
		return nil, false
	}
	return &o.InheritedRoles, true
}

// HasInheritedRoles returns a boolean if a field has been set.
func (o *MongoRBACUser) HasInheritedRoles() bool {
	return o != nil && o.InheritedRoles != nil
}

// SetInheritedRoles gets a reference to the given []MongoRBACRoleRef and assigns it to the InheritedRoles field.
func (o *MongoRBACUser) SetInheritedRoles(v []MongoRBACRoleRef) {
	o.InheritedRoles = v
}

// GetInheritedPrivileges returns the InheritedPrivileges field value if set, zero value otherwise.
func (o *MongoRBACUser) GetInheritedPrivileges() []MongoRBACPrivilege {
	if o == nil || o.InheritedPrivileges == nil {
		var ret []MongoRBACPrivilege
		return ret
	}
	return o.InheritedPrivileges
}

// GetInheritedPrivilegesOk returns a tuple with the InheritedPrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACUser) GetInheritedPrivilegesOk() (*[]MongoRBACPrivilege, bool) {
	if o == nil || o.InheritedPrivileges == nil {
		return nil, false
	}
	return &o.InheritedPrivileges, true
}

// HasInheritedPrivileges returns a boolean if a field has been set.
func (o *MongoRBACUser) HasInheritedPrivileges() bool {
	return o != nil && o.InheritedPrivileges != nil
}

// SetInheritedPrivileges gets a reference to the given []MongoRBACPrivilege and assigns it to the InheritedPrivileges field.
func (o *MongoRBACUser) SetInheritedPrivileges(v []MongoRBACPrivilege) {
	o.InheritedPrivileges = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MongoRBACUser) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACUser) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MongoRBACUser) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *MongoRBACUser) SetUser(v string) {
	o.User = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *MongoRBACUser) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACUser) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *MongoRBACUser) HasReadOnly() bool {
	return o != nil && o.ReadOnly != nil
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *MongoRBACUser) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoRBACUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Db != nil {
		toSerialize["db"] = o.Db
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
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoRBACUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Db                  *string              `json:"db,omitempty"`
		Roles               []MongoRBACRoleRef   `json:"roles,omitempty"`
		InheritedRoles      []MongoRBACRoleRef   `json:"inheritedRoles,omitempty"`
		InheritedPrivileges []MongoRBACPrivilege `json:"inheritedPrivileges,omitempty"`
		User                *string              `json:"user,omitempty"`
		ReadOnly            *bool                `json:"readOnly,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"db", "roles", "inheritedRoles", "inheritedPrivileges", "user", "readOnly"})
	} else {
		return err
	}
	o.Db = all.Db
	o.Roles = all.Roles
	o.InheritedRoles = all.InheritedRoles
	o.InheritedPrivileges = all.InheritedPrivileges
	o.User = all.User
	o.ReadOnly = all.ReadOnly

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
