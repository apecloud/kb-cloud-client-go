// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ESSecurityPasswordRequest struct {
	Password string `json:"password"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewESSecurityPasswordRequest instantiates a new ESSecurityPasswordRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewESSecurityPasswordRequest(password string) *ESSecurityPasswordRequest {
	this := ESSecurityPasswordRequest{}
	this.Password = password
	return &this
}

// NewESSecurityPasswordRequestWithDefaults instantiates a new ESSecurityPasswordRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewESSecurityPasswordRequestWithDefaults() *ESSecurityPasswordRequest {
	this := ESSecurityPasswordRequest{}
	return &this
}

// GetPassword returns the Password field value.
func (o *ESSecurityPasswordRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ESSecurityPasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value.
func (o *ESSecurityPasswordRequest) SetPassword(v string) {
	o.Password = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ESSecurityPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["password"] = o.Password

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ESSecurityPasswordRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Password *string `json:"password"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Password == nil {
		return fmt.Errorf("required field password missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"password"})
	} else {
		return err
	}
	o.Password = *all.Password

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
