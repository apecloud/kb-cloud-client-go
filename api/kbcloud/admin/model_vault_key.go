// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type VaultKey struct {
	Engine *string                `json:"engine,omitempty"`
	Key    *string                `json:"key,omitempty"`
	Data   map[string]interface{} `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewVaultKey instantiates a new VaultKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewVaultKey() *VaultKey {
	this := VaultKey{}
	return &this
}

// NewVaultKeyWithDefaults instantiates a new VaultKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewVaultKeyWithDefaults() *VaultKey {
	this := VaultKey{}
	return &this
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *VaultKey) GetEngine() string {
	if o == nil || o.Engine == nil {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultKey) GetEngineOk() (*string, bool) {
	if o == nil || o.Engine == nil {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *VaultKey) HasEngine() bool {
	return o != nil && o.Engine != nil
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *VaultKey) SetEngine(v string) {
	o.Engine = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VaultKey) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultKey) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VaultKey) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VaultKey) SetKey(v string) {
	o.Key = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VaultKey) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultKey) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VaultKey) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *VaultKey) SetData(v map[string]interface{}) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o VaultKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Engine != nil {
		toSerialize["engine"] = o.Engine
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *VaultKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Engine *string                `json:"engine,omitempty"`
		Key    *string                `json:"key,omitempty"`
		Data   map[string]interface{} `json:"data,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engine", "key", "data"})
	} else {
		return err
	}
	o.Engine = all.Engine
	o.Key = all.Key
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
