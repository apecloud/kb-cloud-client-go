// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RbmqUser struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRbmqUser instantiates a new RbmqUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRbmqUser() *RbmqUser {
	this := RbmqUser{}
	return &this
}

// NewRbmqUserWithDefaults instantiates a new RbmqUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRbmqUserWithDefaults() *RbmqUser {
	this := RbmqUser{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *RbmqUser) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqUser) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *RbmqUser) HasUsername() bool {
	return o != nil && o.Username != nil
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *RbmqUser) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *RbmqUser) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqUser) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *RbmqUser) HasPassword() bool {
	return o != nil && o.Password != nil
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *RbmqUser) SetPassword(v string) {
	o.Password = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RbmqUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RbmqUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Username *string `json:"username,omitempty"`
		Password *string `json:"password,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"username", "password"})
	} else {
		return err
	}
	o.Username = all.Username
	o.Password = all.Password

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
