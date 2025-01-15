// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Account Cluster account information
type Account struct {
	// A list of privileges and their databases.
	PrivilegesList []PrivilegeListItem `json:"privilegesList,omitempty"`
	// Specify the name of component to be connected. If not specified, pick the first one.
	Component *string `json:"component,omitempty"`
	// Specify the name of instance to be connected.
	Instance *string `json:"instance,omitempty"`
	// Specify the name of user, which must be unique.
	Name string `json:"name"`
	// Specify the password of user. The default value is empty, which means a random password will be generated.
	Password *string `json:"password,omitempty"`
	// Role name should be one of [SUPERUSER, BASICUSER].
	Role AccountRoleType `json:"role"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAccount instantiates a new Account object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAccount(name string, role AccountRoleType) *Account {
	this := Account{}
	this.Name = name
	this.Role = role
	return &this
}

// NewAccountWithDefaults instantiates a new Account object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAccountWithDefaults() *Account {
	this := Account{}
	var role AccountRoleType = AccountRoleTypeSuperuser
	this.Role = role
	return &this
}

// GetPrivilegesList returns the PrivilegesList field value if set, zero value otherwise.
func (o *Account) GetPrivilegesList() []PrivilegeListItem {
	if o == nil || o.PrivilegesList == nil {
		var ret []PrivilegeListItem
		return ret
	}
	return o.PrivilegesList
}

// GetPrivilegesListOk returns a tuple with the PrivilegesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetPrivilegesListOk() (*[]PrivilegeListItem, bool) {
	if o == nil || o.PrivilegesList == nil {
		return nil, false
	}
	return &o.PrivilegesList, true
}

// HasPrivilegesList returns a boolean if a field has been set.
func (o *Account) HasPrivilegesList() bool {
	return o != nil && o.PrivilegesList != nil
}

// SetPrivilegesList gets a reference to the given []PrivilegeListItem and assigns it to the PrivilegesList field.
func (o *Account) SetPrivilegesList(v []PrivilegeListItem) {
	o.PrivilegesList = v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *Account) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *Account) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *Account) SetComponent(v string) {
	o.Component = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *Account) GetInstance() string {
	if o == nil || o.Instance == nil {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetInstanceOk() (*string, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *Account) HasInstance() bool {
	return o != nil && o.Instance != nil
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *Account) SetInstance(v string) {
	o.Instance = &v
}

// GetName returns the Name field value.
func (o *Account) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Account) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Account) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Account) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Account) HasPassword() bool {
	return o != nil && o.Password != nil
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Account) SetPassword(v string) {
	o.Password = &v
}

// GetRole returns the Role field value.
func (o *Account) GetRole() AccountRoleType {
	if o == nil {
		var ret AccountRoleType
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *Account) GetRoleOk() (*AccountRoleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value.
func (o *Account) SetRole(v AccountRoleType) {
	o.Role = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.PrivilegesList != nil {
		toSerialize["privilegesList"] = o.PrivilegesList
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}
	toSerialize["name"] = o.Name
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Account) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PrivilegesList []PrivilegeListItem `json:"privilegesList,omitempty"`
		Component      *string             `json:"component,omitempty"`
		Instance       *string             `json:"instance,omitempty"`
		Name           *string             `json:"name"`
		Password       *string             `json:"password,omitempty"`
		Role           *AccountRoleType    `json:"role"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Role == nil {
		return fmt.Errorf("required field role missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"privilegesList", "component", "instance", "name", "password", "role"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PrivilegesList = all.PrivilegesList
	o.Component = all.Component
	o.Instance = all.Instance
	o.Name = *all.Name
	o.Password = all.Password
	if !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = *all.Role
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
