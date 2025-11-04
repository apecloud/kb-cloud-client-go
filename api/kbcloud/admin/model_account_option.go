// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AccountOption struct {
	// number of super user accounts cloud create.
	// If not set, use default value.
	// If set to 0, it means not allowed to create super user account.
	//
	MaxSuperUserAccount      common.NullableInt32 `json:"maxSuperUserAccount,omitempty"`
	Enabled                  bool                 `json:"enabled"`
	Privileges               []string             `json:"privileges,omitempty"`
	AccountNamePattern       string               `json:"accountNamePattern"`
	Create                   bool                 `json:"create"`
	ResetPassword            bool                 `json:"resetPassword"`
	Delete                   bool                 `json:"delete"`
	DisplayRootAccount       *bool                `json:"displayRootAccount,omitempty"`
	ResetRootPassword        *bool                `json:"resetRootPassword,omitempty"`
	SupportMultipleComponent *bool                `json:"supportMultipleComponent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAccountOption instantiates a new AccountOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAccountOption(enabled bool, accountNamePattern string, create bool, resetPassword bool, delete bool) *AccountOption {
	this := AccountOption{}
	var maxSuperUserAccount int32 = 2
	this.MaxSuperUserAccount = *common.NewNullableInt32(&maxSuperUserAccount)
	this.Enabled = enabled
	this.AccountNamePattern = accountNamePattern
	this.Create = create
	this.ResetPassword = resetPassword
	this.Delete = delete
	var supportMultipleComponent bool = false
	this.SupportMultipleComponent = &supportMultipleComponent
	return &this
}

// NewAccountOptionWithDefaults instantiates a new AccountOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAccountOptionWithDefaults() *AccountOption {
	this := AccountOption{}
	var maxSuperUserAccount int32 = 2
	this.MaxSuperUserAccount = *common.NewNullableInt32(&maxSuperUserAccount)
	var supportMultipleComponent bool = false
	this.SupportMultipleComponent = &supportMultipleComponent
	return &this
}

// GetMaxSuperUserAccount returns the MaxSuperUserAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountOption) GetMaxSuperUserAccount() int32 {
	if o == nil || o.MaxSuperUserAccount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MaxSuperUserAccount.Get()
}

// GetMaxSuperUserAccountOk returns a tuple with the MaxSuperUserAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AccountOption) GetMaxSuperUserAccountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxSuperUserAccount.Get(), o.MaxSuperUserAccount.IsSet()
}

// HasMaxSuperUserAccount returns a boolean if a field has been set.
func (o *AccountOption) HasMaxSuperUserAccount() bool {
	return o != nil && o.MaxSuperUserAccount.IsSet()
}

// SetMaxSuperUserAccount gets a reference to the given common.NullableInt32 and assigns it to the MaxSuperUserAccount field.
func (o *AccountOption) SetMaxSuperUserAccount(v int32) {
	o.MaxSuperUserAccount.Set(&v)
}

// SetMaxSuperUserAccountNil sets the value for MaxSuperUserAccount to be an explicit nil.
func (o *AccountOption) SetMaxSuperUserAccountNil() {
	o.MaxSuperUserAccount.Set(nil)
}

// UnsetMaxSuperUserAccount ensures that no value is present for MaxSuperUserAccount, not even an explicit nil.
func (o *AccountOption) UnsetMaxSuperUserAccount() {
	o.MaxSuperUserAccount.Unset()
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

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *AccountOption) GetPrivileges() []string {
	if o == nil || o.Privileges == nil {
		var ret []string
		return ret
	}
	return o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetPrivilegesOk() (*[]string, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return &o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *AccountOption) HasPrivileges() bool {
	return o != nil && o.Privileges != nil
}

// SetPrivileges gets a reference to the given []string and assigns it to the Privileges field.
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

// GetDisplayRootAccount returns the DisplayRootAccount field value if set, zero value otherwise.
func (o *AccountOption) GetDisplayRootAccount() bool {
	if o == nil || o.DisplayRootAccount == nil {
		var ret bool
		return ret
	}
	return *o.DisplayRootAccount
}

// GetDisplayRootAccountOk returns a tuple with the DisplayRootAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetDisplayRootAccountOk() (*bool, bool) {
	if o == nil || o.DisplayRootAccount == nil {
		return nil, false
	}
	return o.DisplayRootAccount, true
}

// HasDisplayRootAccount returns a boolean if a field has been set.
func (o *AccountOption) HasDisplayRootAccount() bool {
	return o != nil && o.DisplayRootAccount != nil
}

// SetDisplayRootAccount gets a reference to the given bool and assigns it to the DisplayRootAccount field.
func (o *AccountOption) SetDisplayRootAccount(v bool) {
	o.DisplayRootAccount = &v
}

// GetResetRootPassword returns the ResetRootPassword field value if set, zero value otherwise.
func (o *AccountOption) GetResetRootPassword() bool {
	if o == nil || o.ResetRootPassword == nil {
		var ret bool
		return ret
	}
	return *o.ResetRootPassword
}

// GetResetRootPasswordOk returns a tuple with the ResetRootPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetResetRootPasswordOk() (*bool, bool) {
	if o == nil || o.ResetRootPassword == nil {
		return nil, false
	}
	return o.ResetRootPassword, true
}

// HasResetRootPassword returns a boolean if a field has been set.
func (o *AccountOption) HasResetRootPassword() bool {
	return o != nil && o.ResetRootPassword != nil
}

// SetResetRootPassword gets a reference to the given bool and assigns it to the ResetRootPassword field.
func (o *AccountOption) SetResetRootPassword(v bool) {
	o.ResetRootPassword = &v
}

// GetSupportMultipleComponent returns the SupportMultipleComponent field value if set, zero value otherwise.
func (o *AccountOption) GetSupportMultipleComponent() bool {
	if o == nil || o.SupportMultipleComponent == nil {
		var ret bool
		return ret
	}
	return *o.SupportMultipleComponent
}

// GetSupportMultipleComponentOk returns a tuple with the SupportMultipleComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountOption) GetSupportMultipleComponentOk() (*bool, bool) {
	if o == nil || o.SupportMultipleComponent == nil {
		return nil, false
	}
	return o.SupportMultipleComponent, true
}

// HasSupportMultipleComponent returns a boolean if a field has been set.
func (o *AccountOption) HasSupportMultipleComponent() bool {
	return o != nil && o.SupportMultipleComponent != nil
}

// SetSupportMultipleComponent gets a reference to the given bool and assigns it to the SupportMultipleComponent field.
func (o *AccountOption) SetSupportMultipleComponent(v bool) {
	o.SupportMultipleComponent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AccountOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.MaxSuperUserAccount.IsSet() {
		toSerialize["maxSuperUserAccount"] = o.MaxSuperUserAccount.Get()
	}
	toSerialize["enabled"] = o.Enabled
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	toSerialize["accountNamePattern"] = o.AccountNamePattern
	toSerialize["create"] = o.Create
	toSerialize["resetPassword"] = o.ResetPassword
	toSerialize["delete"] = o.Delete
	if o.DisplayRootAccount != nil {
		toSerialize["displayRootAccount"] = o.DisplayRootAccount
	}
	if o.ResetRootPassword != nil {
		toSerialize["resetRootPassword"] = o.ResetRootPassword
	}
	if o.SupportMultipleComponent != nil {
		toSerialize["supportMultipleComponent"] = o.SupportMultipleComponent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AccountOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxSuperUserAccount      common.NullableInt32 `json:"maxSuperUserAccount,omitempty"`
		Enabled                  *bool                `json:"enabled"`
		Privileges               []string             `json:"privileges,omitempty"`
		AccountNamePattern       *string              `json:"accountNamePattern"`
		Create                   *bool                `json:"create"`
		ResetPassword            *bool                `json:"resetPassword"`
		Delete                   *bool                `json:"delete"`
		DisplayRootAccount       *bool                `json:"displayRootAccount,omitempty"`
		ResetRootPassword        *bool                `json:"resetRootPassword,omitempty"`
		SupportMultipleComponent *bool                `json:"supportMultipleComponent,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
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
		common.DeleteKeys(additionalProperties, &[]string{"maxSuperUserAccount", "enabled", "privileges", "accountNamePattern", "create", "resetPassword", "delete", "displayRootAccount", "resetRootPassword", "supportMultipleComponent"})
	} else {
		return err
	}
	o.MaxSuperUserAccount = all.MaxSuperUserAccount
	o.Enabled = *all.Enabled
	o.Privileges = all.Privileges
	o.AccountNamePattern = *all.AccountNamePattern
	o.Create = *all.Create
	o.ResetPassword = *all.ResetPassword
	o.Delete = *all.Delete
	o.DisplayRootAccount = all.DisplayRootAccount
	o.ResetRootPassword = all.ResetRootPassword
	o.SupportMultipleComponent = all.SupportMultipleComponent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
