// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RbmqAccountsListItemsItem struct {
	User           *RbmqUser       `json:"user,omitempty"`
	UserPermission *RbmqPermission `json:"userPermission,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRbmqAccountsListItemsItem instantiates a new RbmqAccountsListItemsItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRbmqAccountsListItemsItem() *RbmqAccountsListItemsItem {
	this := RbmqAccountsListItemsItem{}
	return &this
}

// NewRbmqAccountsListItemsItemWithDefaults instantiates a new RbmqAccountsListItemsItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRbmqAccountsListItemsItemWithDefaults() *RbmqAccountsListItemsItem {
	this := RbmqAccountsListItemsItem{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RbmqAccountsListItemsItem) GetUser() RbmqUser {
	if o == nil || o.User == nil {
		var ret RbmqUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqAccountsListItemsItem) GetUserOk() (*RbmqUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RbmqAccountsListItemsItem) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given RbmqUser and assigns it to the User field.
func (o *RbmqAccountsListItemsItem) SetUser(v RbmqUser) {
	o.User = &v
}

// GetUserPermission returns the UserPermission field value if set, zero value otherwise.
func (o *RbmqAccountsListItemsItem) GetUserPermission() RbmqPermission {
	if o == nil || o.UserPermission == nil {
		var ret RbmqPermission
		return ret
	}
	return *o.UserPermission
}

// GetUserPermissionOk returns a tuple with the UserPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqAccountsListItemsItem) GetUserPermissionOk() (*RbmqPermission, bool) {
	if o == nil || o.UserPermission == nil {
		return nil, false
	}
	return o.UserPermission, true
}

// HasUserPermission returns a boolean if a field has been set.
func (o *RbmqAccountsListItemsItem) HasUserPermission() bool {
	return o != nil && o.UserPermission != nil
}

// SetUserPermission gets a reference to the given RbmqPermission and assigns it to the UserPermission field.
func (o *RbmqAccountsListItemsItem) SetUserPermission(v RbmqPermission) {
	o.UserPermission = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RbmqAccountsListItemsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.UserPermission != nil {
		toSerialize["userPermission"] = o.UserPermission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RbmqAccountsListItemsItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		User           *RbmqUser       `json:"user,omitempty"`
		UserPermission *RbmqPermission `json:"userPermission,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"user", "userPermission"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.User != nil && all.User.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.User = all.User
	if all.UserPermission != nil && all.UserPermission.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UserPermission = all.UserPermission

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
