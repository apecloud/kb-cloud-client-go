// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// EncryptionConfig encryption config for cluster
type EncryptionConfig struct {
	// whether enable enc
	Enabled *bool `json:"enabled,omitempty"`
	// the master key for encryption
	MasterKey *string `json:"masterKey,omitempty"`
	// the data key for encryption
	DataKey *string `json:"dataKey,omitempty"`
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

// GetMasterKey returns the MasterKey field value if set, zero value otherwise.
func (o *EncryptionConfig) GetMasterKey() string {
	if o == nil || o.MasterKey == nil {
		var ret string
		return ret
	}
	return *o.MasterKey
}

// GetMasterKeyOk returns a tuple with the MasterKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionConfig) GetMasterKeyOk() (*string, bool) {
	if o == nil || o.MasterKey == nil {
		return nil, false
	}
	return o.MasterKey, true
}

// HasMasterKey returns a boolean if a field has been set.
func (o *EncryptionConfig) HasMasterKey() bool {
	return o != nil && o.MasterKey != nil
}

// SetMasterKey gets a reference to the given string and assigns it to the MasterKey field.
func (o *EncryptionConfig) SetMasterKey(v string) {
	o.MasterKey = &v
}

// GetDataKey returns the DataKey field value if set, zero value otherwise.
func (o *EncryptionConfig) GetDataKey() string {
	if o == nil || o.DataKey == nil {
		var ret string
		return ret
	}
	return *o.DataKey
}

// GetDataKeyOk returns a tuple with the DataKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionConfig) GetDataKeyOk() (*string, bool) {
	if o == nil || o.DataKey == nil {
		return nil, false
	}
	return o.DataKey, true
}

// HasDataKey returns a boolean if a field has been set.
func (o *EncryptionConfig) HasDataKey() bool {
	return o != nil && o.DataKey != nil
}

// SetDataKey gets a reference to the given string and assigns it to the DataKey field.
func (o *EncryptionConfig) SetDataKey(v string) {
	o.DataKey = &v
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
	if o.MasterKey != nil {
		toSerialize["masterKey"] = o.MasterKey
	}
	if o.DataKey != nil {
		toSerialize["dataKey"] = o.DataKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EncryptionConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled   *bool   `json:"enabled,omitempty"`
		MasterKey *string `json:"masterKey,omitempty"`
		DataKey   *string `json:"dataKey,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "masterKey", "dataKey"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.MasterKey = all.MasterKey
	o.DataKey = all.DataKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
