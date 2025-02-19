// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AdminUserUpdate Admin user update
type AdminUserUpdate struct {
	// The name of the user, is unique
	UserName string `json:"userName"`
	// The phonenumber for the user.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The email for the user
	Email string `json:"email"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAdminUserUpdate instantiates a new AdminUserUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAdminUserUpdate(userName string, email string) *AdminUserUpdate {
	this := AdminUserUpdate{}
	this.UserName = userName
	this.Email = email
	return &this
}

// NewAdminUserUpdateWithDefaults instantiates a new AdminUserUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAdminUserUpdateWithDefaults() *AdminUserUpdate {
	this := AdminUserUpdate{}
	return &this
}

// GetUserName returns the UserName field value.
func (o *AdminUserUpdate) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *AdminUserUpdate) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *AdminUserUpdate) SetUserName(v string) {
	o.UserName = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *AdminUserUpdate) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUserUpdate) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AdminUserUpdate) HasPhoneNumber() bool {
	return o != nil && o.PhoneNumber != nil
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *AdminUserUpdate) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value.
func (o *AdminUserUpdate) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AdminUserUpdate) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *AdminUserUpdate) SetEmail(v string) {
	o.Email = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AdminUserUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["userName"] = o.UserName
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AdminUserUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserName    *string `json:"userName"`
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		Email       *string `json:"email"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.UserName == nil {
		return fmt.Errorf("required field userName missing")
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"userName", "phoneNumber", "email"})
	} else {
		return err
	}
	o.UserName = *all.UserName
	o.PhoneNumber = all.PhoneNumber
	o.Email = *all.Email

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
