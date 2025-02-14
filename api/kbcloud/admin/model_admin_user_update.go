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
	// The current email for the user who will be modified
	ModifiedUserEmail string `json:"modifiedUserEmail"`
	// The new username of the user who will be modified, it should be unique
	NewUserName string `json:"newUserName"`
	// The new email for the user who will be modified
	NewEmail string `json:"newEmail"`
	// The phonenumber for the user.
	NewPhoneNumber *string `json:"newPhoneNumber,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAdminUserUpdate instantiates a new AdminUserUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAdminUserUpdate(modifiedUserEmail string, newUserName string, newEmail string) *AdminUserUpdate {
	this := AdminUserUpdate{}
	this.ModifiedUserEmail = modifiedUserEmail
	this.NewUserName = newUserName
	this.NewEmail = newEmail
	return &this
}

// NewAdminUserUpdateWithDefaults instantiates a new AdminUserUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAdminUserUpdateWithDefaults() *AdminUserUpdate {
	this := AdminUserUpdate{}
	return &this
}

// GetModifiedUserEmail returns the ModifiedUserEmail field value.
func (o *AdminUserUpdate) GetModifiedUserEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedUserEmail
}

// GetModifiedUserEmailOk returns a tuple with the ModifiedUserEmail field value
// and a boolean to check if the value has been set.
func (o *AdminUserUpdate) GetModifiedUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedUserEmail, true
}

// SetModifiedUserEmail sets field value.
func (o *AdminUserUpdate) SetModifiedUserEmail(v string) {
	o.ModifiedUserEmail = v
}

// GetNewUserName returns the NewUserName field value.
func (o *AdminUserUpdate) GetNewUserName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NewUserName
}

// GetNewUserNameOk returns a tuple with the NewUserName field value
// and a boolean to check if the value has been set.
func (o *AdminUserUpdate) GetNewUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewUserName, true
}

// SetNewUserName sets field value.
func (o *AdminUserUpdate) SetNewUserName(v string) {
	o.NewUserName = v
}

// GetNewEmail returns the NewEmail field value.
func (o *AdminUserUpdate) GetNewEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NewEmail
}

// GetNewEmailOk returns a tuple with the NewEmail field value
// and a boolean to check if the value has been set.
func (o *AdminUserUpdate) GetNewEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewEmail, true
}

// SetNewEmail sets field value.
func (o *AdminUserUpdate) SetNewEmail(v string) {
	o.NewEmail = v
}

// GetNewPhoneNumber returns the NewPhoneNumber field value if set, zero value otherwise.
func (o *AdminUserUpdate) GetNewPhoneNumber() string {
	if o == nil || o.NewPhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.NewPhoneNumber
}

// GetNewPhoneNumberOk returns a tuple with the NewPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminUserUpdate) GetNewPhoneNumberOk() (*string, bool) {
	if o == nil || o.NewPhoneNumber == nil {
		return nil, false
	}
	return o.NewPhoneNumber, true
}

// HasNewPhoneNumber returns a boolean if a field has been set.
func (o *AdminUserUpdate) HasNewPhoneNumber() bool {
	return o != nil && o.NewPhoneNumber != nil
}

// SetNewPhoneNumber gets a reference to the given string and assigns it to the NewPhoneNumber field.
func (o *AdminUserUpdate) SetNewPhoneNumber(v string) {
	o.NewPhoneNumber = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AdminUserUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["modifiedUserEmail"] = o.ModifiedUserEmail
	toSerialize["newUserName"] = o.NewUserName
	toSerialize["newEmail"] = o.NewEmail
	if o.NewPhoneNumber != nil {
		toSerialize["newPhoneNumber"] = o.NewPhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AdminUserUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ModifiedUserEmail *string `json:"modifiedUserEmail"`
		NewUserName       *string `json:"newUserName"`
		NewEmail          *string `json:"newEmail"`
		NewPhoneNumber    *string `json:"newPhoneNumber,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ModifiedUserEmail == nil {
		return fmt.Errorf("required field modifiedUserEmail missing")
	}
	if all.NewUserName == nil {
		return fmt.Errorf("required field newUserName missing")
	}
	if all.NewEmail == nil {
		return fmt.Errorf("required field newEmail missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"modifiedUserEmail", "newUserName", "newEmail", "newPhoneNumber"})
	} else {
		return err
	}
	o.ModifiedUserEmail = *all.ModifiedUserEmail
	o.NewUserName = *all.NewUserName
	o.NewEmail = *all.NewEmail
	o.NewPhoneNumber = all.NewPhoneNumber

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
