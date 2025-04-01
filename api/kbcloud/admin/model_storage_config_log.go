// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageConfigLog the storage config for log
type StorageConfigLog struct {
	// the name of storage
	StorageName string `json:"storageName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageConfigLog instantiates a new StorageConfigLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageConfigLog(storageName string) *StorageConfigLog {
	this := StorageConfigLog{}
	this.StorageName = storageName
	return &this
}

// NewStorageConfigLogWithDefaults instantiates a new StorageConfigLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageConfigLogWithDefaults() *StorageConfigLog {
	this := StorageConfigLog{}
	return &this
}

// GetStorageName returns the StorageName field value.
func (o *StorageConfigLog) GetStorageName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageName
}

// GetStorageNameOk returns a tuple with the StorageName field value
// and a boolean to check if the value has been set.
func (o *StorageConfigLog) GetStorageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageName, true
}

// SetStorageName sets field value.
func (o *StorageConfigLog) SetStorageName(v string) {
	o.StorageName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageConfigLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["storageName"] = o.StorageName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageConfigLog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StorageName *string `json:"storageName"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.StorageName == nil {
		return fmt.Errorf("required field storageName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"storageName"})
	} else {
		return err
	}
	o.StorageName = *all.StorageName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
