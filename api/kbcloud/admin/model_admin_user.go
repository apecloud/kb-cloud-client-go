// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AdminUser Admin user info
// NODESCRIPTION AdminUser
//
// Deprecated: This model is deprecated.
type AdminUser struct {
	// The name of the user, is unique
	UserName string `json:"userName"`
	// The email for the user
	Email *string `json:"email,omitempty"`
	// The phonenumber for the user.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// NODESCRIPTION CreatedAt
	CreatedAt time.Time `json:"createdAt"`
	// NODESCRIPTION UpdatedAt
	UpdatedAt time.Time `json:"updatedAt"`
	// return true if the default admin user need to reset password
	IsDefaultPassword bool `json:"isDefaultPassword"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAdminUser instantiates a new AdminUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAdminUser(userName string, createdAt time.Time, updatedAt time.Time, isDefaultPassword bool) *AdminUser {
	this := AdminUser{}
	this.UserName = userName
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.IsDefaultPassword = isDefaultPassword
	return &this
}

// NewAdminUserWithDefaults instantiates a new AdminUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAdminUserWithDefaults() *AdminUser {
	this := AdminUser{}
	return &this
}

// GetUserName returns the UserName field value.
func (o *AdminUser) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *AdminUser) SetUserName(v string) {
	o.UserName = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AdminUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AdminUser) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AdminUser) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *AdminUser) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUser) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AdminUser) HasPhoneNumber() bool {
	return o != nil && o.PhoneNumber != nil
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *AdminUser) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AdminUser) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AdminUser) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AdminUser) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AdminUser) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetIsDefaultPassword returns the IsDefaultPassword field value.
func (o *AdminUser) GetIsDefaultPassword() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDefaultPassword
}

// GetIsDefaultPasswordOk returns a tuple with the IsDefaultPassword field value
// and a boolean to check if the value has been set.
func (o *AdminUser) GetIsDefaultPasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefaultPassword, true
}

// SetIsDefaultPassword sets field value.
func (o *AdminUser) SetIsDefaultPassword(v bool) {
	o.IsDefaultPassword = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AdminUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["userName"] = o.UserName
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["isDefaultPassword"] = o.IsDefaultPassword

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AdminUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserName          *string    `json:"userName"`
		Email             *string    `json:"email,omitempty"`
		PhoneNumber       *string    `json:"phoneNumber,omitempty"`
		CreatedAt         *time.Time `json:"createdAt"`
		UpdatedAt         *time.Time `json:"updatedAt"`
		IsDefaultPassword *bool      `json:"isDefaultPassword"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.UserName == nil {
		return fmt.Errorf("required field userName missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	if all.IsDefaultPassword == nil {
		return fmt.Errorf("required field isDefaultPassword missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"userName", "email", "phoneNumber", "createdAt", "updatedAt", "isDefaultPassword"})
	} else {
		return err
	}
	o.UserName = *all.UserName
	o.Email = all.Email
	o.PhoneNumber = all.PhoneNumber
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt
	o.IsDefaultPassword = *all.IsDefaultPassword

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
