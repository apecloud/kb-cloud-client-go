// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// User Admin user info
type User struct {
	// The ID for the user
	Id string `json:"id"`
	// The name of the user, is unique
	UserName string `json:"userName"`
	// The display name of the user
	DisplayName *string `json:"displayName,omitempty"`
	// The email for the user
	Email *string `json:"email,omitempty"`
	// The phonenumber for the user.
	PhoneNumber *string   `json:"phoneNumber,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	// return true if the default admin user need to reset password
	IsDefaultPassword *bool `json:"isDefaultPassword,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUser instantiates a new User object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUser(id string, userName string, createdAt time.Time, updatedAt time.Time) *User {
	this := User{}
	this.Id = id
	this.UserName = userName
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewUserWithDefaults instantiates a new User object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value.
func (o *User) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *User) SetId(v string) {
	o.Id = v
}

// GetUserName returns the UserName field value.
func (o *User) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *User) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value.
func (o *User) SetUserName(v string) {
	o.UserName = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *User) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *User) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *User) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *User) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *User) HasPhoneNumber() bool {
	return o != nil && o.PhoneNumber != nil
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *User) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *User) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *User) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *User) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *User) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetIsDefaultPassword returns the IsDefaultPassword field value if set, zero value otherwise.
func (o *User) GetIsDefaultPassword() bool {
	if o == nil || o.IsDefaultPassword == nil {
		var ret bool
		return ret
	}
	return *o.IsDefaultPassword
}

// GetIsDefaultPasswordOk returns a tuple with the IsDefaultPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIsDefaultPasswordOk() (*bool, bool) {
	if o == nil || o.IsDefaultPassword == nil {
		return nil, false
	}
	return o.IsDefaultPassword, true
}

// HasIsDefaultPassword returns a boolean if a field has been set.
func (o *User) HasIsDefaultPassword() bool {
	return o != nil && o.IsDefaultPassword != nil
}

// SetIsDefaultPassword gets a reference to the given bool and assigns it to the IsDefaultPassword field.
func (o *User) SetIsDefaultPassword(v bool) {
	o.IsDefaultPassword = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["userName"] = o.UserName
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
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
	if o.IsDefaultPassword != nil {
		toSerialize["isDefaultPassword"] = o.IsDefaultPassword
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *User) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id                *string    `json:"id"`
		UserName          *string    `json:"userName"`
		DisplayName       *string    `json:"displayName,omitempty"`
		Email             *string    `json:"email,omitempty"`
		PhoneNumber       *string    `json:"phoneNumber,omitempty"`
		CreatedAt         *time.Time `json:"createdAt"`
		UpdatedAt         *time.Time `json:"updatedAt"`
		IsDefaultPassword *bool      `json:"isDefaultPassword,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "userName", "displayName", "email", "phoneNumber", "createdAt", "updatedAt", "isDefaultPassword"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.UserName = *all.UserName
	o.DisplayName = all.DisplayName
	o.Email = all.Email
	o.PhoneNumber = all.PhoneNumber
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt
	o.IsDefaultPassword = all.IsDefaultPassword

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
