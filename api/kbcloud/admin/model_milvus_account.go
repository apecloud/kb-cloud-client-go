// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MilvusAccount Milvus account creation request.
type MilvusAccount struct {
	// Account name. Kept for compatibility with the generic account model.
	Name *string `json:"name,omitempty"`
	// Milvus username. When omitted, name is used.
	Username *string `json:"username,omitempty"`
	// Account password.
	Password string `json:"password"`
	// Generic account role. SUPERUSER maps to the Milvus admin role.
	Role *string `json:"role,omitempty"`
	// Milvus role names to assign to the user.
	Roles []string `json:"roles,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMilvusAccount instantiates a new MilvusAccount object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMilvusAccount(password string) *MilvusAccount {
	this := MilvusAccount{}
	this.Password = password
	return &this
}

// NewMilvusAccountWithDefaults instantiates a new MilvusAccount object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMilvusAccountWithDefaults() *MilvusAccount {
	this := MilvusAccount{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MilvusAccount) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilvusAccount) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MilvusAccount) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MilvusAccount) SetName(v string) {
	o.Name = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *MilvusAccount) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilvusAccount) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *MilvusAccount) HasUsername() bool {
	return o != nil && o.Username != nil
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *MilvusAccount) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value.
func (o *MilvusAccount) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *MilvusAccount) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *MilvusAccount) SetPassword(v string) {
	o.Password = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *MilvusAccount) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilvusAccount) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *MilvusAccount) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *MilvusAccount) SetRole(v string) {
	o.Role = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *MilvusAccount) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MilvusAccount) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *MilvusAccount) HasRoles() bool {
	return o != nil && o.Roles != nil
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *MilvusAccount) SetRoles(v []string) {
	o.Roles = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MilvusAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	toSerialize["password"] = o.Password
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
func (o *MilvusAccount) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string  `json:"name,omitempty"`
		Username *string  `json:"username,omitempty"`
		Password *string  `json:"password"`
		Role     *string  `json:"role,omitempty"`
		Roles    []string `json:"roles,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "username", "password", "role", "roles"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Username = all.Username
	o.Password = *all.Password
	o.Role = all.Role
	o.Roles = all.Roles

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
