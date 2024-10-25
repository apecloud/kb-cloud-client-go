// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION AccountOption
type AccountOption struct {
	// NODESCRIPTION Enabled
	Enabled bool `json:"enabled"`
	// NODESCRIPTION Privileges
	Privileges []string `json:"privileges"`
	// NODESCRIPTION AccountNamePattern
	AccountNamePattern string `json:"accountNamePattern"`
	// NODESCRIPTION Create
	Create bool `json:"create"`
	// NODESCRIPTION ResetPassword
	ResetPassword bool `json:"resetPassword"`
	// NODESCRIPTION Delete
	Delete bool `json:"delete"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAccountOption instantiates a new AccountOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAccountOption(enabled bool, privileges []string, accountNamePattern string, create bool, resetPassword bool, delete bool) *AccountOption {
	this := AccountOption{}
	this.Enabled = enabled
	this.Privileges = privileges
	this.AccountNamePattern = accountNamePattern
	this.Create = create
	this.ResetPassword = resetPassword
	this.Delete = delete
	return &this
}

// NewAccountOptionWithDefaults instantiates a new AccountOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAccountOptionWithDefaults() *AccountOption {
	this := AccountOption{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *AccountOption) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AccountOption) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *AccountOption) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPrivileges returns the Privileges field value.
func (o *AccountOption) GetPrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value
// and a boolean to check if the value has been set.
func (o *AccountOption) GetPrivilegesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// SetPrivileges sets field value.
func (o *AccountOption) SetPrivileges(v []string) {
	o.Privileges = v
}

// GetAccountNamePattern returns the AccountNamePattern field value.
func (o *AccountOption) GetAccountNamePattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountNamePattern
}

// GetAccountNamePatternOk returns a tuple with the AccountNamePattern field value
// and a boolean to check if the value has been set.
func (o *AccountOption) GetAccountNamePatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNamePattern, true
}

// SetAccountNamePattern sets field value.
func (o *AccountOption) SetAccountNamePattern(v string) {
	o.AccountNamePattern = v
}

// GetCreate returns the Create field value.
func (o *AccountOption) GetCreate() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Create
}

// GetCreateOk returns a tuple with the Create field value
// and a boolean to check if the value has been set.
func (o *AccountOption) GetCreateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Create, true
}

// SetCreate sets field value.
func (o *AccountOption) SetCreate(v bool) {
	o.Create = v
}

// GetResetPassword returns the ResetPassword field value.
func (o *AccountOption) GetResetPassword() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ResetPassword
}

// GetResetPasswordOk returns a tuple with the ResetPassword field value
// and a boolean to check if the value has been set.
func (o *AccountOption) GetResetPasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResetPassword, true
}

// SetResetPassword sets field value.
func (o *AccountOption) SetResetPassword(v bool) {
	o.ResetPassword = v
}

// GetDelete returns the Delete field value.
func (o *AccountOption) GetDelete() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value
// and a boolean to check if the value has been set.
func (o *AccountOption) GetDeleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delete, true
}

// SetDelete sets field value.
func (o *AccountOption) SetDelete(v bool) {
	o.Delete = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AccountOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["privileges"] = o.Privileges
	toSerialize["accountNamePattern"] = o.AccountNamePattern
	toSerialize["create"] = o.Create
	toSerialize["resetPassword"] = o.ResetPassword
	toSerialize["delete"] = o.Delete

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AccountOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled            *bool     `json:"enabled"`
		Privileges         *[]string `json:"privileges"`
		AccountNamePattern *string   `json:"accountNamePattern"`
		Create             *bool     `json:"create"`
		ResetPassword      *bool     `json:"resetPassword"`
		Delete             *bool     `json:"delete"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Privileges == nil {
		return fmt.Errorf("required field privileges missing")
	}
	if all.AccountNamePattern == nil {
		return fmt.Errorf("required field accountNamePattern missing")
	}
	if all.Create == nil {
		return fmt.Errorf("required field create missing")
	}
	if all.ResetPassword == nil {
		return fmt.Errorf("required field resetPassword missing")
	}
	if all.Delete == nil {
		return fmt.Errorf("required field delete missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "privileges", "accountNamePattern", "create", "resetPassword", "delete"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.Privileges = *all.Privileges
	o.AccountNamePattern = *all.AccountNamePattern
	o.Create = *all.Create
	o.ResetPassword = *all.ResetPassword
	o.Delete = *all.Delete

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
