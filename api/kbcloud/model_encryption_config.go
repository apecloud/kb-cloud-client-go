// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// EncryptionConfig encryption config for cluster
type EncryptionConfig struct {
	// whether enable enc
	Enabled *bool `json:"enabled,omitempty"`
	// the secret ref for encryption
	SecretKeyRef *string `json:"secretKeyRef,omitempty"`
	// the key name used for encryption
	Key *string `json:"key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEncryptionConfig instantiates a new EncryptionConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEncryptionConfig() *EncryptionConfig {
	this := EncryptionConfig{}
	return &this
}

// NewEncryptionConfigWithDefaults instantiates a new EncryptionConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEncryptionConfigWithDefaults() *EncryptionConfig {
	this := EncryptionConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EncryptionConfig) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EncryptionConfig) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EncryptionConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSecretKeyRef returns the SecretKeyRef field value if set, zero value otherwise.
func (o *EncryptionConfig) GetSecretKeyRef() string {
	if o == nil || o.SecretKeyRef == nil {
		var ret string
		return ret
	}
	return *o.SecretKeyRef
}

// GetSecretKeyRefOk returns a tuple with the SecretKeyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionConfig) GetSecretKeyRefOk() (*string, bool) {
	if o == nil || o.SecretKeyRef == nil {
		return nil, false
	}
	return o.SecretKeyRef, true
}

// HasSecretKeyRef returns a boolean if a field has been set.
func (o *EncryptionConfig) HasSecretKeyRef() bool {
	return o != nil && o.SecretKeyRef != nil
}

// SetSecretKeyRef gets a reference to the given string and assigns it to the SecretKeyRef field.
func (o *EncryptionConfig) SetSecretKeyRef(v string) {
	o.SecretKeyRef = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EncryptionConfig) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionConfig) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EncryptionConfig) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EncryptionConfig) SetKey(v string) {
	o.Key = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EncryptionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.SecretKeyRef != nil {
		toSerialize["secretKeyRef"] = o.SecretKeyRef
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EncryptionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled      *bool   `json:"enabled,omitempty"`
		SecretKeyRef *string `json:"secretKeyRef,omitempty"`
		Key          *string `json:"key,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "secretKeyRef", "key"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.SecretKeyRef = all.SecretKeyRef
	o.Key = all.Key

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
