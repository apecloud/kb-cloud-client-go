// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type EnvironmentCredentialUpdate struct {
	// Description of the environment credential
	Description common.NullableString `json:"description,omitempty"`
	// Key of the environment credential entry, such as accessKeyId for AWS
	Key string `json:"key"`
	// Secret of the environment credential entry, such as secretAccessKey for AWS
	Secret string `json:"secret"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentCredentialUpdate instantiates a new EnvironmentCredentialUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentCredentialUpdate(key string, secret string) *EnvironmentCredentialUpdate {
	this := EnvironmentCredentialUpdate{}
	this.Key = key
	this.Secret = secret
	return &this
}

// NewEnvironmentCredentialUpdateWithDefaults instantiates a new EnvironmentCredentialUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentCredentialUpdateWithDefaults() *EnvironmentCredentialUpdate {
	this := EnvironmentCredentialUpdate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnvironmentCredentialUpdate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *EnvironmentCredentialUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *EnvironmentCredentialUpdate) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given common.NullableString and assigns it to the Description field.
func (o *EnvironmentCredentialUpdate) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *EnvironmentCredentialUpdate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *EnvironmentCredentialUpdate) UnsetDescription() {
	o.Description.Unset()
}

// GetKey returns the Key field value.
func (o *EnvironmentCredentialUpdate) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialUpdate) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *EnvironmentCredentialUpdate) SetKey(v string) {
	o.Key = v
}

// GetSecret returns the Secret field value.
func (o *EnvironmentCredentialUpdate) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *EnvironmentCredentialUpdate) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value.
func (o *EnvironmentCredentialUpdate) SetSecret(v string) {
	o.Secret = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentCredentialUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["key"] = o.Key
	toSerialize["secret"] = o.Secret

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentCredentialUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description common.NullableString `json:"description,omitempty"`
		Key         *string               `json:"key"`
		Secret      *string               `json:"secret"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Secret == nil {
		return fmt.Errorf("required field secret missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"description", "key", "secret"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Key = *all.Key
	o.Secret = *all.Secret

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
