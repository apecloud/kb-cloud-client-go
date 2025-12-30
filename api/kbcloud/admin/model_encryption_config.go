// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// EncryptionConfig encryption config for cluster
type EncryptionConfig struct {
	// whether enable encryption for backup
	Enabled *bool `json:"enabled,omitempty"`
	// the key name used for encryption. If this field is empty, a new key will be generated.
	KeyName *string `json:"keyName,omitempty"`
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

// GetKeyName returns the KeyName field value if set, zero value otherwise.
func (o *EncryptionConfig) GetKeyName() string {
	if o == nil || o.KeyName == nil {
		var ret string
		return ret
	}
	return *o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionConfig) GetKeyNameOk() (*string, bool) {
	if o == nil || o.KeyName == nil {
		return nil, false
	}
	return o.KeyName, true
}

// HasKeyName returns a boolean if a field has been set.
func (o *EncryptionConfig) HasKeyName() bool {
	return o != nil && o.KeyName != nil
}

// SetKeyName gets a reference to the given string and assigns it to the KeyName field.
func (o *EncryptionConfig) SetKeyName(v string) {
	o.KeyName = &v
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
	if o.KeyName != nil {
		toSerialize["keyName"] = o.KeyName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EncryptionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool   `json:"enabled,omitempty"`
		KeyName *string `json:"keyName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "keyName"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.KeyName = all.KeyName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
