// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// UserCreate User create
type UserCreate struct {
	// The name of the user, is unique
	UserName string `json:"userName"`
	// The email for the user
	Email *string `json:"email,omitempty"`
	// The display name of the user
	DisplayName *string `json:"displayName,omitempty"`
	// The password for the user
	Password string `json:"password"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserCreate instantiates a new UserCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserCreate(userName string, password string) *UserCreate {
	this := UserCreate{}
	this.UserName = userName
	this.Password = password
	return &this
}

// NewUserCreateWithDefaults instantiates a new UserCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserCreateWithDefaults() *UserCreate {
	this := UserCreate{}
	return &this
}

// GetUserName returns the UserName field value.
func (o *UserCreate) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *UserCreate) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *UserCreate) SetUserName(v string) {
	o.UserName = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserCreate) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserCreate) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserCreate) SetEmail(v string) {
	o.Email = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserCreate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCreate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserCreate) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserCreate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPassword returns the Password field value.
func (o *UserCreate) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserCreate) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *UserCreate) SetPassword(v string) {
	o.Password = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["userName"] = o.UserName
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["password"] = o.Password

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserName    *string `json:"userName"`
		Email       *string `json:"email,omitempty"`
		DisplayName *string `json:"displayName,omitempty"`
		Password    *string `json:"password"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.UserName == nil {
		return fmt.Errorf("required field userName missing")
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"userName", "email", "displayName", "password"})
	} else {
		return err
	}
	o.UserName = *all.UserName
	o.Email = all.Email
	o.DisplayName = all.DisplayName
	o.Password = *all.Password

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
