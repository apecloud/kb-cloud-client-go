// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

// AdminUserCreate Admin user create
type AdminUserCreate struct {
	// The name of the user, is unique
	UserName string `json:"userName"`
	// The password of the admin user
	Password string `json:"password"`
	// The email for the user
	Email *string `json:"email,omitempty"`
	// The phonenumber for the user.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAdminUserCreate instantiates a new AdminUserCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAdminUserCreate(userName string, password string) *AdminUserCreate {
	this := AdminUserCreate{}
	this.UserName = userName
	this.Password = password
	return &this
}

// NewAdminUserCreateWithDefaults instantiates a new AdminUserCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAdminUserCreateWithDefaults() *AdminUserCreate {
	this := AdminUserCreate{}
	return &this
}

// GetUserName returns the UserName field value.
func (o *AdminUserCreate) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *AdminUserCreate) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *AdminUserCreate) SetUserName(v string) {
	o.UserName = v
}

// GetPassword returns the Password field value.
func (o *AdminUserCreate) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AdminUserCreate) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *AdminUserCreate) SetPassword(v string) {
	o.Password = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AdminUserCreate) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUserCreate) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AdminUserCreate) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AdminUserCreate) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *AdminUserCreate) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUserCreate) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AdminUserCreate) HasPhoneNumber() bool {
	return o != nil && o.PhoneNumber != nil
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *AdminUserCreate) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AdminUserCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["userName"] = o.UserName
	toSerialize["password"] = o.Password
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AdminUserCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserName    *string `json:"userName"`
		Password    *string `json:"password"`
		Email       *string `json:"email,omitempty"`
		PhoneNumber *string `json:"phoneNumber,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.UserName == nil {
		return fmt.Errorf("required field userName missing")
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"userName", "password", "email", "phoneNumber"})
	} else {
		return err
	}
	o.UserName = *all.UserName
	o.Password = *all.Password
	o.Email = all.Email
	o.PhoneNumber = all.PhoneNumber

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
