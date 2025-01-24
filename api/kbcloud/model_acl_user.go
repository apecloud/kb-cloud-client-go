// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ACLUser struct {
	// The user role name, should be one of [ROOT, SUPERUSER, BASICUSER].
	Role AccountListRoleType `json:"role"`
	// Redis user privileges
	Privileges *RedisPrivilegeType `json:"privileges,omitempty"`
	// ACL user name
	Name string `json:"name"`
	// Whether the user is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Whether the user requires no password
	Nopass *bool `json:"nopass,omitempty"`
	// List of plain text passwords
	Passwords []string `json:"passwords,omitempty"`
	// List of hashed passwords
	PasswordHashes []string `json:"password_hashes,omitempty"`
	// List of passwords to be removed
	PasswordsToRemove []string `json:"passwords_to_remove,omitempty"`
	// List of password hashes to be removed
	PasswordHashesToRemove []string `json:"password_hashes_to_remove,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewACLUser instantiates a new ACLUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewACLUser(role AccountListRoleType, name string) *ACLUser {
	this := ACLUser{}
	this.Role = role
	this.Name = name
	return &this
}

// NewACLUserWithDefaults instantiates a new ACLUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewACLUserWithDefaults() *ACLUser {
	this := ACLUser{}
	return &this
}

// GetRole returns the Role field value.
func (o *ACLUser) GetRole() AccountListRoleType {
	if o == nil {
		var ret AccountListRoleType
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ACLUser) GetRoleOk() (*AccountListRoleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *ACLUser) SetRole(v AccountListRoleType) {
	o.Role = v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *ACLUser) GetPrivileges() RedisPrivilegeType {
	if o == nil || o.Privileges == nil {
		var ret RedisPrivilegeType
		return ret
	}
	return *o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLUser) GetPrivilegesOk() (*RedisPrivilegeType, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *ACLUser) HasPrivileges() bool {
	return o != nil && o.Privileges != nil
}

// SetPrivileges gets a reference to the given RedisPrivilegeType and assigns it to the Privileges field.
func (o *ACLUser) SetPrivileges(v RedisPrivilegeType) {
	o.Privileges = &v
}

// GetName returns the Name field value.
func (o *ACLUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ACLUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ACLUser) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ACLUser) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLUser) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ACLUser) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ACLUser) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNopass returns the Nopass field value if set, zero value otherwise.
func (o *ACLUser) GetNopass() bool {
	if o == nil || o.Nopass == nil {
		var ret bool
		return ret
	}
	return *o.Nopass
}

// GetNopassOk returns a tuple with the Nopass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLUser) GetNopassOk() (*bool, bool) {
	if o == nil || o.Nopass == nil {
		return nil, false
	}
	return o.Nopass, true
}

// HasNopass returns a boolean if a field has been set.
func (o *ACLUser) HasNopass() bool {
	return o != nil && o.Nopass != nil
}

// SetNopass gets a reference to the given bool and assigns it to the Nopass field.
func (o *ACLUser) SetNopass(v bool) {
	o.Nopass = &v
}

// GetPasswords returns the Passwords field value if set, zero value otherwise.
func (o *ACLUser) GetPasswords() []string {
	if o == nil || o.Passwords == nil {
		var ret []string
		return ret
	}
	return o.Passwords
}

// GetPasswordsOk returns a tuple with the Passwords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLUser) GetPasswordsOk() (*[]string, bool) {
	if o == nil || o.Passwords == nil {
		return nil, false
	}
	return &o.Passwords, true
}

// HasPasswords returns a boolean if a field has been set.
func (o *ACLUser) HasPasswords() bool {
	return o != nil && o.Passwords != nil
}

// SetPasswords gets a reference to the given []string and assigns it to the Passwords field.
func (o *ACLUser) SetPasswords(v []string) {
	o.Passwords = v
}

// GetPasswordHashes returns the PasswordHashes field value if set, zero value otherwise.
func (o *ACLUser) GetPasswordHashes() []string {
	if o == nil || o.PasswordHashes == nil {
		var ret []string
		return ret
	}
	return o.PasswordHashes
}

// GetPasswordHashesOk returns a tuple with the PasswordHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLUser) GetPasswordHashesOk() (*[]string, bool) {
	if o == nil || o.PasswordHashes == nil {
		return nil, false
	}
	return &o.PasswordHashes, true
}

// HasPasswordHashes returns a boolean if a field has been set.
func (o *ACLUser) HasPasswordHashes() bool {
	return o != nil && o.PasswordHashes != nil
}

// SetPasswordHashes gets a reference to the given []string and assigns it to the PasswordHashes field.
func (o *ACLUser) SetPasswordHashes(v []string) {
	o.PasswordHashes = v
}

// GetPasswordsToRemove returns the PasswordsToRemove field value if set, zero value otherwise.
func (o *ACLUser) GetPasswordsToRemove() []string {
	if o == nil || o.PasswordsToRemove == nil {
		var ret []string
		return ret
	}
	return o.PasswordsToRemove
}

// GetPasswordsToRemoveOk returns a tuple with the PasswordsToRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLUser) GetPasswordsToRemoveOk() (*[]string, bool) {
	if o == nil || o.PasswordsToRemove == nil {
		return nil, false
	}
	return &o.PasswordsToRemove, true
}

// HasPasswordsToRemove returns a boolean if a field has been set.
func (o *ACLUser) HasPasswordsToRemove() bool {
	return o != nil && o.PasswordsToRemove != nil
}

// SetPasswordsToRemove gets a reference to the given []string and assigns it to the PasswordsToRemove field.
func (o *ACLUser) SetPasswordsToRemove(v []string) {
	o.PasswordsToRemove = v
}

// GetPasswordHashesToRemove returns the PasswordHashesToRemove field value if set, zero value otherwise.
func (o *ACLUser) GetPasswordHashesToRemove() []string {
	if o == nil || o.PasswordHashesToRemove == nil {
		var ret []string
		return ret
	}
	return o.PasswordHashesToRemove
}

// GetPasswordHashesToRemoveOk returns a tuple with the PasswordHashesToRemove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLUser) GetPasswordHashesToRemoveOk() (*[]string, bool) {
	if o == nil || o.PasswordHashesToRemove == nil {
		return nil, false
	}
	return &o.PasswordHashesToRemove, true
}

// HasPasswordHashesToRemove returns a boolean if a field has been set.
func (o *ACLUser) HasPasswordHashesToRemove() bool {
	return o != nil && o.PasswordHashesToRemove != nil
}

// SetPasswordHashesToRemove gets a reference to the given []string and assigns it to the PasswordHashesToRemove field.
func (o *ACLUser) SetPasswordHashesToRemove(v []string) {
	o.PasswordHashesToRemove = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ACLUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["role"] = o.Role
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	toSerialize["name"] = o.Name
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Nopass != nil {
		toSerialize["nopass"] = o.Nopass
	}
	if o.Passwords != nil {
		toSerialize["passwords"] = o.Passwords
	}
	if o.PasswordHashes != nil {
		toSerialize["password_hashes"] = o.PasswordHashes
	}
	if o.PasswordsToRemove != nil {
		toSerialize["passwords_to_remove"] = o.PasswordsToRemove
	}
	if o.PasswordHashesToRemove != nil {
		toSerialize["password_hashes_to_remove"] = o.PasswordHashesToRemove
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ACLUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Role                   *AccountListRoleType `json:"role"`
		Privileges             *RedisPrivilegeType  `json:"privileges,omitempty"`
		Name                   *string              `json:"name"`
		Enabled                *bool                `json:"enabled,omitempty"`
		Nopass                 *bool                `json:"nopass,omitempty"`
		Passwords              []string             `json:"passwords,omitempty"`
		PasswordHashes         []string             `json:"password_hashes,omitempty"`
		PasswordsToRemove      []string             `json:"passwords_to_remove,omitempty"`
		PasswordHashesToRemove []string             `json:"password_hashes_to_remove,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"role", "privileges", "name", "enabled", "nopass", "passwords", "password_hashes", "passwords_to_remove", "password_hashes_to_remove"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = *all.Role
	}
	if all.Privileges != nil && !all.Privileges.IsValid() {
		hasInvalidField = true
	} else {
		o.Privileges = all.Privileges
	}
	o.Name = *all.Name
	o.Enabled = all.Enabled
	o.Nopass = all.Nopass
	o.Passwords = all.Passwords
	o.PasswordHashes = all.PasswordHashes
	o.PasswordsToRemove = all.PasswordsToRemove
	o.PasswordHashesToRemove = all.PasswordHashesToRemove

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
