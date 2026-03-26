// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentCredentialEntry struct {
	// Key of the credential entry
	Key string `json:"key"`
	// Value of the credential entry
	Value string `json:"value"`
	// Whether the credential entry is secret
	Secret bool `json:"secret"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentCredentialEntry instantiates a new EnvironmentCredentialEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentCredentialEntry(key string, value string, secret bool) *EnvironmentCredentialEntry {
	this := EnvironmentCredentialEntry{}
	this.Key = key
	this.Value = value
	this.Secret = secret
	return &this
}

// NewEnvironmentCredentialEntryWithDefaults instantiates a new EnvironmentCredentialEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentCredentialEntryWithDefaults() *EnvironmentCredentialEntry {
	this := EnvironmentCredentialEntry{}
	return &this
}

// GetKey returns the Key field value.
func (o *EnvironmentCredentialEntry) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialEntry) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *EnvironmentCredentialEntry) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value.
func (o *EnvironmentCredentialEntry) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialEntry) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *EnvironmentCredentialEntry) SetValue(v string) {
	o.Value = v
}

// GetSecret returns the Secret field value.
func (o *EnvironmentCredentialEntry) GetSecret() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialEntry) GetSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value.
func (o *EnvironmentCredentialEntry) SetSecret(v bool) {
	o.Secret = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentCredentialEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	toSerialize["secret"] = o.Secret

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentCredentialEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key    *string `json:"key"`
		Value  *string `json:"value"`
		Secret *bool   `json:"secret"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	if all.Secret == nil {
		return fmt.Errorf("required field secret missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key", "value", "secret"})
	} else {
		return err
	}
	o.Key = *all.Key
	o.Value = *all.Value
	o.Secret = *all.Secret

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
