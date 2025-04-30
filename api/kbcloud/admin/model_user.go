// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// User User info
type User struct {
	// created time of the user
	CreatedAt time.Time `json:"createdAt"`
	// The name of the user, is unique
	UserName string `json:"userName"`
	// The display name of the user
	DisplayName string `json:"displayName"`
	// The email for the user
	Email string `json:"email"`
	// The ID for the user
	Id string `json:"id"`
	// The phonenumber for the user.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// updated time of the user
	UpdatedAt time.Time `json:"updatedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUser instantiates a new User object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUser(createdAt time.Time, userName string, displayName string, email string, id string, updatedAt time.Time) *User {
	this := User{}
	this.CreatedAt = createdAt
	this.UserName = userName
	this.DisplayName = displayName
	this.Email = email
	this.Id = id
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

// GetDisplayName returns the DisplayName field value.
func (o *User) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *User) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *User) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEmail returns the Email field value.
func (o *User) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value.
func (o *User) SetEmail(v string) {
	o.Email = v
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

// MarshalJSON serializes the struct using spec logic.
func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["userName"] = o.UserName
	toSerialize["displayName"] = o.DisplayName
	toSerialize["email"] = o.Email
	toSerialize["id"] = o.Id
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *User) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time `json:"createdAt"`
		UserName    *string    `json:"userName"`
		DisplayName *string    `json:"displayName"`
		Email       *string    `json:"email"`
		Id          *string    `json:"id"`
		PhoneNumber *string    `json:"phoneNumber,omitempty"`
		UpdatedAt   *time.Time `json:"updatedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UserName == nil {
		return fmt.Errorf("required field userName missing")
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field displayName missing")
	}
	if all.Email == nil {
		return fmt.Errorf("required field email missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"createdAt", "userName", "displayName", "email", "id", "phoneNumber", "updatedAt"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.UserName = *all.UserName
	o.DisplayName = *all.DisplayName
	o.Email = *all.Email
	o.Id = *all.Id
	o.PhoneNumber = all.PhoneNumber
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
