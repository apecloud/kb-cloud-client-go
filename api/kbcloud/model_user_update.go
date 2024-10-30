// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// UserUpdate User update 
type UserUpdate struct {
	// The display name of the user
	DisplayName *string `json:"displayName,omitempty"`
	// The email for the user
	Email *string `json:"email,omitempty"`
	// The phonenumber for the user.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The verification code for the user
	VerificationCode *string `json:"verificationCode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewUserUpdate instantiates a new UserUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserUpdate() *UserUpdate {
	this := UserUpdate{}
	return &this
}

// NewUserUpdateWithDefaults instantiates a new UserUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserUpdateWithDefaults() *UserUpdate {
	this := UserUpdate{}
	return &this
}
// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserUpdate) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}


// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserUpdate) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserUpdate) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserUpdate) SetEmail(v string) {
	o.Email = &v
}


// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UserUpdate) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UserUpdate) HasPhoneNumber() bool {
	return o != nil && o.PhoneNumber != nil
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UserUpdate) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}


// GetVerificationCode returns the VerificationCode field value if set, zero value otherwise.
func (o *UserUpdate) GetVerificationCode() string {
	if o == nil || o.VerificationCode == nil {
		var ret string
		return ret
	}
	return *o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetVerificationCodeOk() (*string, bool) {
	if o == nil || o.VerificationCode == nil {
		return nil, false
	}
	return o.VerificationCode, true
}

// HasVerificationCode returns a boolean if a field has been set.
func (o *UserUpdate) HasVerificationCode() bool {
	return o != nil && o.VerificationCode != nil
}

// SetVerificationCode gets a reference to the given string and assigns it to the VerificationCode field.
func (o *UserUpdate) SetVerificationCode(v string) {
	o.VerificationCode = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o UserUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.VerificationCode != nil {
		toSerialize["verificationCode"] = o.VerificationCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string `json:"displayName,omitempty"`
		Email *string `json:"email,omitempty"`
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		VerificationCode *string `json:"verificationCode,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "displayName", "email", "phoneNumber", "verificationCode",  })
	} else {
		return err
	}
	o.DisplayName = all.DisplayName
	o.Email = all.Email
	o.PhoneNumber = all.PhoneNumber
	o.VerificationCode = all.VerificationCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
