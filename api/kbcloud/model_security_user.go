// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SecurityUser struct {
	Roles    []string               `json:"roles,omitempty"`
	FullName *string                `json:"full_name,omitempty"`
	Email    *string                `json:"email,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Enabled  common.NullableBool    `json:"enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityUser instantiates a new SecurityUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityUser() *SecurityUser {
	this := SecurityUser{}
	return &this
}

// NewSecurityUserWithDefaults instantiates a new SecurityUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityUserWithDefaults() *SecurityUser {
	this := SecurityUser{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *SecurityUser) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUser) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return &o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *SecurityUser) HasRoles() bool {
	return o != nil && o.Roles != nil
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *SecurityUser) SetRoles(v []string) {
	o.Roles = v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *SecurityUser) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUser) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *SecurityUser) HasFullName() bool {
	return o != nil && o.FullName != nil
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *SecurityUser) SetFullName(v string) {
	o.FullName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SecurityUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SecurityUser) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SecurityUser) SetEmail(v string) {
	o.Email = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SecurityUser) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityUser) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SecurityUser) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SecurityUser) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityUser) GetEnabled() bool {
	if o == nil || o.Enabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityUser) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *SecurityUser) HasEnabled() bool {
	return o != nil && o.Enabled.IsSet()
}

// SetEnabled gets a reference to the given common.NullableBool and assigns it to the Enabled field.
func (o *SecurityUser) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}

// SetEnabledNil sets the value for Enabled to be an explicit nil.
func (o *SecurityUser) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil.
func (o *SecurityUser) UnsetEnabled() {
	o.Enabled.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
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
func (o *SecurityUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Roles    []string               `json:"roles,omitempty"`
		FullName *string                `json:"full_name,omitempty"`
		Email    *string                `json:"email,omitempty"`
		Metadata map[string]interface{} `json:"metadata,omitempty"`
		Enabled  common.NullableBool    `json:"enabled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"roles", "full_name", "email", "metadata", "enabled"})
	} else {
		return err
	}
	o.Roles = all.Roles
	o.FullName = all.FullName
	o.Email = all.Email
	o.Metadata = all.Metadata
	o.Enabled = all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
