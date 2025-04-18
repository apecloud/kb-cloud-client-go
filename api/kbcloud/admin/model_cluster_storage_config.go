// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ClusterStorageConfig struct {
	// the name of storage
	StorageName string `json:"storageName"`
	// the bucket name for the storage
	Bucket *string `json:"bucket,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterStorageConfig instantiates a new ClusterStorageConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterStorageConfig(storageName string) *ClusterStorageConfig {
	this := ClusterStorageConfig{}
	this.StorageName = storageName
	return &this
}

// NewClusterStorageConfigWithDefaults instantiates a new ClusterStorageConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterStorageConfigWithDefaults() *ClusterStorageConfig {
	this := ClusterStorageConfig{}
	return &this
}

// GetStorageName returns the StorageName field value.
func (o *ClusterStorageConfig) GetStorageName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageName
}

// GetStorageNameOk returns a tuple with the StorageName field value
// and a boolean to check if the value has been set.
func (o *ClusterStorageConfig) GetStorageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageName, true
}

// SetStorageName sets field value.
func (o *ClusterStorageConfig) SetStorageName(v string) {
	o.StorageName = v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *ClusterStorageConfig) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStorageConfig) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *ClusterStorageConfig) HasBucket() bool {
	return o != nil && o.Bucket != nil
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *ClusterStorageConfig) SetBucket(v string) {
	o.Bucket = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterStorageConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["storageName"] = o.StorageName
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterStorageConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StorageName *string `json:"storageName"`
		Bucket      *string `json:"bucket,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.StorageName == nil {
		return fmt.Errorf("required field storageName missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"storageName", "bucket"})
	} else {
		return err
	}
	o.StorageName = *all.StorageName
	o.Bucket = all.Bucket

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
