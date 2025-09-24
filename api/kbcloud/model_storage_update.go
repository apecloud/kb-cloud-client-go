// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageUpdate storageUpdate is the schema for storage update request
type StorageUpdate struct {
	// the parameters to create the storage
	Params map[string]string `json:"params,omitempty"`
	// the tags for the storage
	Tags        map[string]string `json:"tags,omitempty"`
	Engines     []string          `json:"engines,omitempty"`
	StorageName *string           `json:"storageName,omitempty"`
	EnvName     *string           `json:"envName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageUpdate instantiates a new StorageUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageUpdate() *StorageUpdate {
	this := StorageUpdate{}
	return &this
}

// NewStorageUpdateWithDefaults instantiates a new StorageUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageUpdateWithDefaults() *StorageUpdate {
	this := StorageUpdate{}
	return &this
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *StorageUpdate) GetParams() map[string]string {
	if o == nil || o.Params == nil {
		var ret map[string]string
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUpdate) GetParamsOk() (*map[string]string, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return &o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *StorageUpdate) HasParams() bool {
	return o != nil && o.Params != nil
}

// SetParams gets a reference to the given map[string]string and assigns it to the Params field.
func (o *StorageUpdate) SetParams(v map[string]string) {
	o.Params = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *StorageUpdate) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUpdate) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *StorageUpdate) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *StorageUpdate) SetTags(v map[string]string) {
	o.Tags = v
}

// GetEngines returns the Engines field value if set, zero value otherwise.
func (o *StorageUpdate) GetEngines() []string {
	if o == nil || o.Engines == nil {
		var ret []string
		return ret
	}
	return o.Engines
}

// GetEnginesOk returns a tuple with the Engines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUpdate) GetEnginesOk() (*[]string, bool) {
	if o == nil || o.Engines == nil {
		return nil, false
	}
	return &o.Engines, true
}

// HasEngines returns a boolean if a field has been set.
func (o *StorageUpdate) HasEngines() bool {
	return o != nil && o.Engines != nil
}

// SetEngines gets a reference to the given []string and assigns it to the Engines field.
func (o *StorageUpdate) SetEngines(v []string) {
	o.Engines = v
}

// GetStorageName returns the StorageName field value if set, zero value otherwise.
func (o *StorageUpdate) GetStorageName() string {
	if o == nil || o.StorageName == nil {
		var ret string
		return ret
	}
	return *o.StorageName
}

// GetStorageNameOk returns a tuple with the StorageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUpdate) GetStorageNameOk() (*string, bool) {
	if o == nil || o.StorageName == nil {
		return nil, false
	}
	return o.StorageName, true
}

// HasStorageName returns a boolean if a field has been set.
func (o *StorageUpdate) HasStorageName() bool {
	return o != nil && o.StorageName != nil
}

// SetStorageName gets a reference to the given string and assigns it to the StorageName field.
func (o *StorageUpdate) SetStorageName(v string) {
	o.StorageName = &v
}

// GetEnvName returns the EnvName field value if set, zero value otherwise.
func (o *StorageUpdate) GetEnvName() string {
	if o == nil || o.EnvName == nil {
		var ret string
		return ret
	}
	return *o.EnvName
}

// GetEnvNameOk returns a tuple with the EnvName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageUpdate) GetEnvNameOk() (*string, bool) {
	if o == nil || o.EnvName == nil {
		return nil, false
	}
	return o.EnvName, true
}

// HasEnvName returns a boolean if a field has been set.
func (o *StorageUpdate) HasEnvName() bool {
	return o != nil && o.EnvName != nil
}

// SetEnvName gets a reference to the given string and assigns it to the EnvName field.
func (o *StorageUpdate) SetEnvName(v string) {
	o.EnvName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Engines != nil {
		toSerialize["engines"] = o.Engines
	}
	if o.StorageName != nil {
		toSerialize["storageName"] = o.StorageName
	}
	if o.EnvName != nil {
		toSerialize["envName"] = o.EnvName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Params      map[string]string `json:"params,omitempty"`
		Tags        map[string]string `json:"tags,omitempty"`
		Engines     []string          `json:"engines,omitempty"`
		StorageName *string           `json:"storageName,omitempty"`
		EnvName     *string           `json:"envName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"params", "tags", "engines", "storageName", "envName"})
	} else {
		return err
	}
	o.Params = all.Params
	o.Tags = all.Tags
	o.Engines = all.Engines
	o.StorageName = all.StorageName
	o.EnvName = all.EnvName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
