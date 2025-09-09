// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AccountListItem Cluster account information.
type AccountListItem struct {
	// Component type.
	Component *string `json:"component,omitempty"`
	// The name of user.
	Name string `json:"name"`
	// The password of user.
	Password *string `json:"password,omitempty"`
	// The user role name, should be one of [ROOT, SUPERUSER, BASICUSER].
	Role *AccountListRoleType `json:"role,omitempty"`
	// A list of privileges and their databases.
	PrivilegesList []PrivilegeListItem `json:"privilegesList,omitempty"`
	// extra information of account
	Extra map[string]interface{} `json:"extra,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAccountListItem instantiates a new AccountListItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAccountListItem(name string) *AccountListItem {
	this := AccountListItem{}
	this.Name = name
	return &this
}

// NewAccountListItemWithDefaults instantiates a new AccountListItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAccountListItemWithDefaults() *AccountListItem {
	this := AccountListItem{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *AccountListItem) GetComponent() string {
	if o == nil || o.Component == nil {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListItem) GetComponentOk() (*string, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *AccountListItem) HasComponent() bool {
	return o != nil && o.Component != nil
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *AccountListItem) SetComponent(v string) {
	o.Component = &v
}

// GetName returns the Name field value.
func (o *AccountListItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountListItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AccountListItem) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AccountListItem) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListItem) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AccountListItem) HasPassword() bool {
	return o != nil && o.Password != nil
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AccountListItem) SetPassword(v string) {
	o.Password = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AccountListItem) GetRole() AccountListRoleType {
	if o == nil || o.Role == nil {
		var ret AccountListRoleType
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListItem) GetRoleOk() (*AccountListRoleType, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AccountListItem) HasRole() bool {
	return o != nil && o.Role != nil
}

// SetRole gets a reference to the given AccountListRoleType and assigns it to the Role field.
func (o *AccountListItem) SetRole(v AccountListRoleType) {
	o.Role = &v
}

// GetPrivilegesList returns the PrivilegesList field value if set, zero value otherwise.
func (o *AccountListItem) GetPrivilegesList() []PrivilegeListItem {
	if o == nil || o.PrivilegesList == nil {
		var ret []PrivilegeListItem
		return ret
	}
	return o.PrivilegesList
}

// GetPrivilegesListOk returns a tuple with the PrivilegesList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListItem) GetPrivilegesListOk() (*[]PrivilegeListItem, bool) {
	if o == nil || o.PrivilegesList == nil {
		return nil, false
	}
	return &o.PrivilegesList, true
}

// HasPrivilegesList returns a boolean if a field has been set.
func (o *AccountListItem) HasPrivilegesList() bool {
	return o != nil && o.PrivilegesList != nil
}

// SetPrivilegesList gets a reference to the given []PrivilegeListItem and assigns it to the PrivilegesList field.
func (o *AccountListItem) SetPrivilegesList(v []PrivilegeListItem) {
	o.PrivilegesList = v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *AccountListItem) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountListItem) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return &o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *AccountListItem) HasExtra() bool {
	return o != nil && o.Extra != nil
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *AccountListItem) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AccountListItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Component != nil {
		toSerialize["component"] = o.Component
	}
	toSerialize["name"] = o.Name
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.PrivilegesList != nil {
		toSerialize["privilegesList"] = o.PrivilegesList
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AccountListItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component      *string                `json:"component,omitempty"`
		Name           *string                `json:"name"`
		Password       *string                `json:"password,omitempty"`
		Role           *AccountListRoleType   `json:"role,omitempty"`
		PrivilegesList []PrivilegeListItem    `json:"privilegesList,omitempty"`
		Extra          map[string]interface{} `json:"extra,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "name", "password", "role", "privilegesList", "extra"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Component = all.Component
	o.Name = *all.Name
	o.Password = all.Password
	if all.Role != nil && !all.Role.IsValid() {
		hasInvalidField = true
	} else {
		o.Role = all.Role
	}
	o.PrivilegesList = all.PrivilegesList
	o.Extra = all.Extra

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
