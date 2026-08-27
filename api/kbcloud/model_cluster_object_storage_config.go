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
	// whether the object storage endpoint uses TLS. External S3-compatible endpoints must include an explicit https:// scheme when this is true.
	TlsEnabled *bool `json:"tlsEnabled,omitempty"`
	// the Kubernetes secret name that stores the CA certificate for TLS object storage.
	TlsCaCertSecret *string `json:"tlsCACertSecret,omitempty"`
	// the key in tlsCACertSecret that stores the CA certificate.
	TlsCaCertSecretKey *string `json:"tlsCACertSecretKey,omitempty"`
	// region to use. If using a s3-compatible service that does not require a region (like minio), leave it blank.
	Region *string `json:"region,omitempty"`
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

// GetTlsEnabled returns the TlsEnabled field value if set, zero value otherwise.
func (o *ClusterObjectStorageConfig) GetTlsEnabled() bool {
	if o == nil || o.TlsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TlsEnabled
}

// GetTlsEnabledOk returns a tuple with the TlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterObjectStorageConfig) GetTlsEnabledOk() (*bool, bool) {
	if o == nil || o.TlsEnabled == nil {
		return nil, false
	}
	return o.TlsEnabled, true
}

// HasTlsEnabled returns a boolean if a field has been set.
func (o *ClusterObjectStorageConfig) HasTlsEnabled() bool {
	return o != nil && o.TlsEnabled != nil
}

// SetTlsEnabled gets a reference to the given bool and assigns it to the TlsEnabled field.
func (o *ClusterObjectStorageConfig) SetTlsEnabled(v bool) {
	o.TlsEnabled = &v
}

// GetTlsCaCertSecret returns the TlsCaCertSecret field value if set, zero value otherwise.
func (o *ClusterObjectStorageConfig) GetTlsCaCertSecret() string {
	if o == nil || o.TlsCaCertSecret == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCertSecret
}

// GetTlsCaCertSecretOk returns a tuple with the TlsCaCertSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterObjectStorageConfig) GetTlsCaCertSecretOk() (*string, bool) {
	if o == nil || o.TlsCaCertSecret == nil {
		return nil, false
	}
	return o.TlsCaCertSecret, true
}

// HasTlsCaCertSecret returns a boolean if a field has been set.
func (o *ClusterObjectStorageConfig) HasTlsCaCertSecret() bool {
	return o != nil && o.TlsCaCertSecret != nil
}

// SetTlsCaCertSecret gets a reference to the given string and assigns it to the TlsCaCertSecret field.
func (o *ClusterObjectStorageConfig) SetTlsCaCertSecret(v string) {
	o.TlsCaCertSecret = &v
}

// GetTlsCaCertSecretKey returns the TlsCaCertSecretKey field value if set, zero value otherwise.
func (o *ClusterObjectStorageConfig) GetTlsCaCertSecretKey() string {
	if o == nil || o.TlsCaCertSecretKey == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCertSecretKey
}

// GetTlsCaCertSecretKeyOk returns a tuple with the TlsCaCertSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterObjectStorageConfig) GetTlsCaCertSecretKeyOk() (*string, bool) {
	if o == nil || o.TlsCaCertSecretKey == nil {
		return nil, false
	}
	return o.TlsCaCertSecretKey, true
}

// HasTlsCaCertSecretKey returns a boolean if a field has been set.
func (o *ClusterObjectStorageConfig) HasTlsCaCertSecretKey() bool {
	return o != nil && o.TlsCaCertSecretKey != nil
}

// SetTlsCaCertSecretKey gets a reference to the given string and assigns it to the TlsCaCertSecretKey field.
func (o *ClusterObjectStorageConfig) SetTlsCaCertSecretKey(v string) {
	o.TlsCaCertSecretKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ClusterObjectStorageConfig) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterObjectStorageConfig) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ClusterObjectStorageConfig) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ClusterObjectStorageConfig) SetRegion(v string) {
	o.Region = &v
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
	if o.TlsEnabled != nil {
		toSerialize["tlsEnabled"] = o.TlsEnabled
	}
	if o.TlsCaCertSecret != nil {
		toSerialize["tlsCACertSecret"] = o.TlsCaCertSecret
	}
	if o.TlsCaCertSecretKey != nil {
		toSerialize["tlsCACertSecretKey"] = o.TlsCaCertSecretKey
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterObjectStorageConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ServiceRef         *ServiceRef `json:"serviceRef"`
		Bucket             *string     `json:"bucket"`
		UsePathStyle       *bool       `json:"usePathStyle,omitempty"`
		TlsEnabled         *bool       `json:"tlsEnabled,omitempty"`
		TlsCaCertSecret    *string     `json:"tlsCACertSecret,omitempty"`
		TlsCaCertSecretKey *string     `json:"tlsCACertSecretKey,omitempty"`
		Region             *string     `json:"region,omitempty"`
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
		common.DeleteKeys(additionalProperties, &[]string{"serviceRef", "bucket", "usePathStyle", "tlsEnabled", "tlsCACertSecret", "tlsCACertSecretKey", "region"})
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
	o.TlsEnabled = all.TlsEnabled
	o.TlsCaCertSecret = all.TlsCaCertSecret
	o.TlsCaCertSecretKey = all.TlsCaCertSecretKey
	o.Region = all.Region

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
