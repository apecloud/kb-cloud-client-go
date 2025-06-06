// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Password Admin user password
type Password struct {
	// The new password for the admin user.
	NewPassword string `json:"newPassword"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPassword instantiates a new Password object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPassword(newPassword string) *Password {
	this := Password{}
	this.NewPassword = newPassword
	return &this
}

// NewPasswordWithDefaults instantiates a new Password object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPasswordWithDefaults() *Password {
	this := Password{}
	return &this
}

// GetNewPassword returns the NewPassword field value.
func (o *Password) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *Password) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value.
func (o *Password) SetNewPassword(v string) {
	o.NewPassword = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Password) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["newPassword"] = o.NewPassword

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Password) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NewPassword *string `json:"newPassword"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.NewPassword == nil {
		return fmt.Errorf("required field newPassword missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"newPassword"})
	} else {
		return err
	}
	o.NewPassword = *all.NewPassword

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
