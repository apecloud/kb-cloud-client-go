// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Environment_storageConfigLog the storage config for log
type Environment_storageConfigLog struct {
	// the name of storage
	StorageName string `json:"storageName"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironment_storageConfigLog instantiates a new Environment_storageConfigLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironment_storageConfigLog(storageName string) *Environment_storageConfigLog {
	this := Environment_storageConfigLog{}
	this.StorageName = storageName
	return &this
}

// NewEnvironment_storageConfigLogWithDefaults instantiates a new Environment_storageConfigLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironment_storageConfigLogWithDefaults() *Environment_storageConfigLog {
	this := Environment_storageConfigLog{}
	return &this
}

// GetStorageName returns the StorageName field value.
func (o *Environment_storageConfigLog) GetStorageName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageName
}

// GetStorageNameOk returns a tuple with the StorageName field value
// and a boolean to check if the value has been set.
func (o *Environment_storageConfigLog) GetStorageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageName, true
}

// SetStorageName sets field value.
func (o *Environment_storageConfigLog) SetStorageName(v string) {
	o.StorageName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Environment_storageConfigLog) MarshalJSON() ([]byte, error) {
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
func (o *Environment_storageConfigLog) UnmarshalJSON(bytes []byte) (err error) {
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
