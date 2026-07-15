// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PutSecurityUserRequest struct {
	Password     *string                `json:"password,omitempty"`
	PasswordHash *string                `json:"password_hash,omitempty"`
	Roles        []string               `json:"roles"`
	FullName     *string                `json:"full_name,omitempty"`
	Email        *string                `json:"email,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Enabled      common.NullableBool    `json:"enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPutSecurityUserRequest instantiates a new PutSecurityUserRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPutSecurityUserRequest(roles []string) *PutSecurityUserRequest {
	this := PutSecurityUserRequest{}
	this.Roles = roles
	return &this
}

// NewPutSecurityUserRequestWithDefaults instantiates a new PutSecurityUserRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPutSecurityUserRequestWithDefaults() *PutSecurityUserRequest {
	this := PutSecurityUserRequest{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PutSecurityUserRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutSecurityUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PutSecurityUserRequest) HasPassword() bool {
	return o != nil && o.Password != nil
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PutSecurityUserRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise.
func (o *PutSecurityUserRequest) GetPasswordHash() string {
	if o == nil || o.PasswordHash == nil {
		var ret string
		return ret
	}
	return *o.PasswordHash
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutSecurityUserRequest) GetPasswordHashOk() (*string, bool) {
	if o == nil || o.PasswordHash == nil {
		return nil, false
	}
	return o.PasswordHash, true
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *PutSecurityUserRequest) HasPasswordHash() bool {
	return o != nil && o.PasswordHash != nil
}

// SetPasswordHash gets a reference to the given string and assigns it to the PasswordHash field.
func (o *PutSecurityUserRequest) SetPasswordHash(v string) {
	o.PasswordHash = &v
}

// GetRoles returns the Roles field value.
func (o *PutSecurityUserRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *PutSecurityUserRequest) GetRolesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value.
func (o *PutSecurityUserRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *PutSecurityUserRequest) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutSecurityUserRequest) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *PutSecurityUserRequest) HasFullName() bool {
	return o != nil && o.FullName != nil
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *PutSecurityUserRequest) SetFullName(v string) {
	o.FullName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PutSecurityUserRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutSecurityUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PutSecurityUserRequest) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PutSecurityUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PutSecurityUserRequest) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutSecurityUserRequest) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PutSecurityUserRequest) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PutSecurityUserRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutSecurityUserRequest) GetEnabled() bool {
	if o == nil || o.Enabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PutSecurityUserRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *PutSecurityUserRequest) HasEnabled() bool {
	return o != nil && o.Enabled.IsSet()
}

// SetEnabled gets a reference to the given common.NullableBool and assigns it to the Enabled field.
func (o *PutSecurityUserRequest) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil.
func (o *PutSecurityUserRequest) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil.
func (o *PutSecurityUserRequest) UnsetEnabled() {
	o.Enabled.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o PutSecurityUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PasswordHash != nil {
		toSerialize["password_hash"] = o.PasswordHash
	}
	toSerialize["roles"] = o.Roles
	if o.FullName != nil {
		toSerialize["full_name"] = o.FullName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PutSecurityUserRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Password     *string                `json:"password,omitempty"`
		PasswordHash *string                `json:"password_hash,omitempty"`
		Roles        *[]string              `json:"roles"`
		FullName     *string                `json:"full_name,omitempty"`
		Email        *string                `json:"email,omitempty"`
		Metadata     map[string]interface{} `json:"metadata,omitempty"`
		Enabled      common.NullableBool    `json:"enabled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Roles == nil {
		return fmt.Errorf("required field roles missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"password", "password_hash", "roles", "full_name", "email", "metadata", "enabled"})
	} else {
		return err
	}
	o.Password = all.Password
	o.PasswordHash = all.PasswordHash
	o.Roles = *all.Roles
	o.FullName = all.FullName
	o.Email = all.Email
	o.Metadata = all.Metadata
	o.Enabled = all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
