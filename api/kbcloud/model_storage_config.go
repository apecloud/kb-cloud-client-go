// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type StorageConfig struct {
	// the name of storage
	StorageName *string `json:"storageName,omitempty"`
	// the bucket name for the storage
	BucketName *string `json:"bucketName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageConfig instantiates a new StorageConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageConfig() *StorageConfig {
	this := StorageConfig{}
	return &this
}

// NewStorageConfigWithDefaults instantiates a new StorageConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageConfigWithDefaults() *StorageConfig {
	this := StorageConfig{}
	return &this
}

// GetStorageName returns the StorageName field value if set, zero value otherwise.
func (o *StorageConfig) GetStorageName() string {
	if o == nil || o.StorageName == nil {
		var ret string
		return ret
	}
	return *o.StorageName
}

// GetStorageNameOk returns a tuple with the StorageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageConfig) GetStorageNameOk() (*string, bool) {
	if o == nil || o.StorageName == nil {
		return nil, false
	}
	return o.StorageName, true
}

// HasStorageName returns a boolean if a field has been set.
func (o *StorageConfig) HasStorageName() bool {
	return o != nil && o.StorageName != nil
}

// SetStorageName gets a reference to the given string and assigns it to the StorageName field.
func (o *StorageConfig) SetStorageName(v string) {
	o.StorageName = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *StorageConfig) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageConfig) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *StorageConfig) HasBucketName() bool {
	return o != nil && o.BucketName != nil
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *StorageConfig) SetBucketName(v string) {
	o.BucketName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.StorageName != nil {
		toSerialize["storageName"] = o.StorageName
	}
	if o.BucketName != nil {
		toSerialize["bucketName"] = o.BucketName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StorageName *string `json:"storageName,omitempty"`
		BucketName  *string `json:"bucketName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"storageName", "bucketName"})
	} else {
		return err
	}
	o.StorageName = all.StorageName
	o.BucketName = all.BucketName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
