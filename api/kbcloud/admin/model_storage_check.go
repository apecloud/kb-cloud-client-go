// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageCheck storageCheck is the schema for the storage connectivity check request
type StorageCheck struct {
	// Name of the storage provider
	StorageProvider string `json:"storageProvider"`
	// the parameters to check the storage
	Params map[string]string `json:"params"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageCheck instantiates a new StorageCheck object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageCheck(storageProvider string, params map[string]string) *StorageCheck {
	this := StorageCheck{}
	this.StorageProvider = storageProvider
	this.Params = params
	return &this
}

// NewStorageCheckWithDefaults instantiates a new StorageCheck object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageCheckWithDefaults() *StorageCheck {
	this := StorageCheck{}
	return &this
}

// GetStorageProvider returns the StorageProvider field value.
func (o *StorageCheck) GetStorageProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageProvider
}

// GetStorageProviderOk returns a tuple with the StorageProvider field value
// and a boolean to check if the value has been set.
func (o *StorageCheck) GetStorageProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageProvider, true
}

// SetStorageProvider sets field value.
func (o *StorageCheck) SetStorageProvider(v string) {
	o.StorageProvider = v
}

// GetParams returns the Params field value.
func (o *StorageCheck) GetParams() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *StorageCheck) GetParamsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value.
func (o *StorageCheck) SetParams(v map[string]string) {
	o.Params = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["storageProvider"] = o.StorageProvider
	toSerialize["params"] = o.Params

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageCheck) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StorageProvider *string            `json:"storageProvider"`
		Params          *map[string]string `json:"params"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.StorageProvider == nil {
		return fmt.Errorf("required field storageProvider missing")
	}
	if all.Params == nil {
		return fmt.Errorf("required field params missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"storageProvider", "params"})
	} else {
		return err
	}
	o.StorageProvider = *all.StorageProvider
	o.Params = *all.Params

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
