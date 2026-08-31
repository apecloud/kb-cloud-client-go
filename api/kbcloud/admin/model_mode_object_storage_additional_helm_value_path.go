// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeObjectStorageAdditionalHelmValuePath The path in helm values that some object storage config will use. If empty, the values will not be set.
type ModeObjectStorageAdditionalHelmValuePath struct {
	// the bucket name for the object storage
	Bucket string `json:"bucket"`
	// root path where cluster stores data in the bucket. This field is not user-provided.
	// It is always set to cluster id. If not set, it means the engine does not support specify a path.
	//
	Path *string `json:"path,omitempty"`
	// whether the object storage is using path style or virtual host style.
	// If not set, it means the engine does not need this option.
	//
	UsePathStyle *string `json:"usePathStyle,omitempty"`
	// whether the object storage endpoint uses TLS.
	// If not set, it means the engine does not need this option.
	//
	TlsEnabled *string `json:"tlsEnabled,omitempty"`
	// Secret name that stores the CA certificate for TLS object storage.
	// If not set, it means the engine does not need this option.
	//
	TlsCaCertSecret *string `json:"tlsCACertSecret,omitempty"`
	// Secret key that stores the CA certificate for TLS object storage.
	// If not set, it means the engine does not need this option.
	//
	TlsCaCertSecretKey *string `json:"tlsCACertSecretKey,omitempty"`
	// Region to use. If not set, it means the engine does not need this option.
	//
	Region *string `json:"region,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeObjectStorageAdditionalHelmValuePath instantiates a new ModeObjectStorageAdditionalHelmValuePath object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeObjectStorageAdditionalHelmValuePath(bucket string) *ModeObjectStorageAdditionalHelmValuePath {
	this := ModeObjectStorageAdditionalHelmValuePath{}
	this.Bucket = bucket
	return &this
}

// NewModeObjectStorageAdditionalHelmValuePathWithDefaults instantiates a new ModeObjectStorageAdditionalHelmValuePath object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeObjectStorageAdditionalHelmValuePathWithDefaults() *ModeObjectStorageAdditionalHelmValuePath {
	this := ModeObjectStorageAdditionalHelmValuePath{}
	return &this
}

// GetBucket returns the Bucket field value.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value.
func (o *ModeObjectStorageAdditionalHelmValuePath) SetBucket(v string) {
	o.Bucket = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ModeObjectStorageAdditionalHelmValuePath) SetPath(v string) {
	o.Path = &v
}

// GetUsePathStyle returns the UsePathStyle field value if set, zero value otherwise.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetUsePathStyle() string {
	if o == nil || o.UsePathStyle == nil {
		var ret string
		return ret
	}
	return *o.UsePathStyle
}

// GetUsePathStyleOk returns a tuple with the UsePathStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetUsePathStyleOk() (*string, bool) {
	if o == nil || o.UsePathStyle == nil {
		return nil, false
	}
	return o.UsePathStyle, true
}

// HasUsePathStyle returns a boolean if a field has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) HasUsePathStyle() bool {
	return o != nil && o.UsePathStyle != nil
}

// SetUsePathStyle gets a reference to the given string and assigns it to the UsePathStyle field.
func (o *ModeObjectStorageAdditionalHelmValuePath) SetUsePathStyle(v string) {
	o.UsePathStyle = &v
}

// GetTlsEnabled returns the TlsEnabled field value if set, zero value otherwise.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetTlsEnabled() string {
	if o == nil || o.TlsEnabled == nil {
		var ret string
		return ret
	}
	return *o.TlsEnabled
}

// GetTlsEnabledOk returns a tuple with the TlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetTlsEnabledOk() (*string, bool) {
	if o == nil || o.TlsEnabled == nil {
		return nil, false
	}
	return o.TlsEnabled, true
}

// HasTlsEnabled returns a boolean if a field has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) HasTlsEnabled() bool {
	return o != nil && o.TlsEnabled != nil
}

// SetTlsEnabled gets a reference to the given string and assigns it to the TlsEnabled field.
func (o *ModeObjectStorageAdditionalHelmValuePath) SetTlsEnabled(v string) {
	o.TlsEnabled = &v
}

// GetTlsCaCertSecret returns the TlsCaCertSecret field value if set, zero value otherwise.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetTlsCaCertSecret() string {
	if o == nil || o.TlsCaCertSecret == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCertSecret
}

// GetTlsCaCertSecretOk returns a tuple with the TlsCaCertSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetTlsCaCertSecretOk() (*string, bool) {
	if o == nil || o.TlsCaCertSecret == nil {
		return nil, false
	}
	return o.TlsCaCertSecret, true
}

// HasTlsCaCertSecret returns a boolean if a field has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) HasTlsCaCertSecret() bool {
	return o != nil && o.TlsCaCertSecret != nil
}

// SetTlsCaCertSecret gets a reference to the given string and assigns it to the TlsCaCertSecret field.
func (o *ModeObjectStorageAdditionalHelmValuePath) SetTlsCaCertSecret(v string) {
	o.TlsCaCertSecret = &v
}

// GetTlsCaCertSecretKey returns the TlsCaCertSecretKey field value if set, zero value otherwise.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetTlsCaCertSecretKey() string {
	if o == nil || o.TlsCaCertSecretKey == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCertSecretKey
}

// GetTlsCaCertSecretKeyOk returns a tuple with the TlsCaCertSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetTlsCaCertSecretKeyOk() (*string, bool) {
	if o == nil || o.TlsCaCertSecretKey == nil {
		return nil, false
	}
	return o.TlsCaCertSecretKey, true
}

// HasTlsCaCertSecretKey returns a boolean if a field has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) HasTlsCaCertSecretKey() bool {
	return o != nil && o.TlsCaCertSecretKey != nil
}

// SetTlsCaCertSecretKey gets a reference to the given string and assigns it to the TlsCaCertSecretKey field.
func (o *ModeObjectStorageAdditionalHelmValuePath) SetTlsCaCertSecretKey(v string) {
	o.TlsCaCertSecretKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ModeObjectStorageAdditionalHelmValuePath) SetRegion(v string) {
	o.Region = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeObjectStorageAdditionalHelmValuePath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["bucket"] = o.Bucket
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
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
func (o *ModeObjectStorageAdditionalHelmValuePath) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bucket             *string `json:"bucket"`
		Path               *string `json:"path,omitempty"`
		UsePathStyle       *string `json:"usePathStyle,omitempty"`
		TlsEnabled         *string `json:"tlsEnabled,omitempty"`
		TlsCaCertSecret    *string `json:"tlsCACertSecret,omitempty"`
		TlsCaCertSecretKey *string `json:"tlsCACertSecretKey,omitempty"`
		Region             *string `json:"region,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Bucket == nil {
		return fmt.Errorf("required field bucket missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"bucket", "path", "usePathStyle", "tlsEnabled", "tlsCACertSecret", "tlsCACertSecretKey", "region"})
	} else {
		return err
	}
	o.Bucket = *all.Bucket
	o.Path = all.Path
	o.UsePathStyle = all.UsePathStyle
	o.TlsEnabled = all.TlsEnabled
	o.TlsCaCertSecret = all.TlsCaCertSecret
	o.TlsCaCertSecretKey = all.TlsCaCertSecretKey
	o.Region = all.Region

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
