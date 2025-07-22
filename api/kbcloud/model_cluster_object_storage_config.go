// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterObjectStorageConfig Specify the object storage config for cluster like starrocks
type ClusterObjectStorageConfig struct {
	// defines a serviceRef that references service provided by other cluster or external service.
	// Only one of cluster or serviceDescriptor field should be set.
	//
	ServiceRef ServiceRef `json:"serviceRef"`
	// the bucket name for the object storage
	Bucket string `json:"bucket"`
	// whether the object storage is using path-style. If false, virtual host style will be used.
	UsePathStyle *bool `json:"usePathStyle,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterObjectStorageConfig instantiates a new ClusterObjectStorageConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterObjectStorageConfig(serviceRef ServiceRef, bucket string) *ClusterObjectStorageConfig {
	this := ClusterObjectStorageConfig{}
	this.ServiceRef = serviceRef
	this.Bucket = bucket
	return &this
}

// NewClusterObjectStorageConfigWithDefaults instantiates a new ClusterObjectStorageConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterObjectStorageConfigWithDefaults() *ClusterObjectStorageConfig {
	this := ClusterObjectStorageConfig{}
	return &this
}

// GetServiceRef returns the ServiceRef field value.
func (o *ClusterObjectStorageConfig) GetServiceRef() ServiceRef {
	if o == nil {
		var ret ServiceRef
		return ret
	}
	return o.ServiceRef
}

// GetServiceRefOk returns a tuple with the ServiceRef field value
// and a boolean to check if the value has been set.
func (o *ClusterObjectStorageConfig) GetServiceRefOk() (*ServiceRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceRef, true
}

// SetServiceRef sets field value.
func (o *ClusterObjectStorageConfig) SetServiceRef(v ServiceRef) {
	o.ServiceRef = v
}

// GetBucket returns the Bucket field value.
func (o *ClusterObjectStorageConfig) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *ClusterObjectStorageConfig) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value.
func (o *ClusterObjectStorageConfig) SetBucket(v string) {
	o.Bucket = v
}

// GetUsePathStyle returns the UsePathStyle field value if set, zero value otherwise.
func (o *ClusterObjectStorageConfig) GetUsePathStyle() bool {
	if o == nil || o.UsePathStyle == nil {
		var ret bool
		return ret
	}
	return *o.UsePathStyle
}

// GetUsePathStyleOk returns a tuple with the UsePathStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterObjectStorageConfig) GetUsePathStyleOk() (*bool, bool) {
	if o == nil || o.UsePathStyle == nil {
		return nil, false
	}
	return o.UsePathStyle, true
}

// HasUsePathStyle returns a boolean if a field has been set.
func (o *ClusterObjectStorageConfig) HasUsePathStyle() bool {
	return o != nil && o.UsePathStyle != nil
}

// SetUsePathStyle gets a reference to the given bool and assigns it to the UsePathStyle field.
func (o *ClusterObjectStorageConfig) SetUsePathStyle(v bool) {
	o.UsePathStyle = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterObjectStorageConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["serviceRef"] = o.ServiceRef
	toSerialize["bucket"] = o.Bucket
	if o.UsePathStyle != nil {
		toSerialize["usePathStyle"] = o.UsePathStyle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterObjectStorageConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ServiceRef   *ServiceRef `json:"serviceRef"`
		Bucket       *string     `json:"bucket"`
		UsePathStyle *bool       `json:"usePathStyle,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ServiceRef == nil {
		return fmt.Errorf("required field serviceRef missing")
	}
	if all.Bucket == nil {
		return fmt.Errorf("required field bucket missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"serviceRef", "bucket", "usePathStyle"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ServiceRef.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ServiceRef = *all.ServiceRef
	o.Bucket = *all.Bucket
	o.UsePathStyle = all.UsePathStyle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
