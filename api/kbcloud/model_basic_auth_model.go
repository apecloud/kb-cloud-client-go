// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type BasicAuthModel struct {
	UserName *string `json:"userName,omitempty"`
	// if password is not provided when patch replication channel, it means not to modify the password
	Password common.NullableString `json:"password,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBasicAuthModel instantiates a new BasicAuthModel object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBasicAuthModel() *BasicAuthModel {
	this := BasicAuthModel{}
	return &this
}

// NewBasicAuthModelWithDefaults instantiates a new BasicAuthModel object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBasicAuthModelWithDefaults() *BasicAuthModel {
	this := BasicAuthModel{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *BasicAuthModel) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthModel) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *BasicAuthModel) HasUserName() bool {
	return o != nil && o.UserName != nil
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *BasicAuthModel) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BasicAuthModel) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *BasicAuthModel) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *BasicAuthModel) HasPassword() bool {
	return o != nil && o.Password.IsSet()
}

// SetPassword gets a reference to the given common.NullableString and assigns it to the Password field.
func (o *BasicAuthModel) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil.
func (o *BasicAuthModel) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil.
func (o *BasicAuthModel) UnsetPassword() {
	o.Password.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o BasicAuthModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BasicAuthModel) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		UserName *string               `json:"userName,omitempty"`
		Password common.NullableString `json:"password,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"userName", "password"})
	} else {
		return err
	}
	o.UserName = all.UserName
	o.Password = all.Password

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
