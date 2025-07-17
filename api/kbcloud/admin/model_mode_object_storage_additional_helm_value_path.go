// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ModeObjectStorageAdditionalHelmValuePath The path in helm values that some object storage config will use.
type ModeObjectStorageAdditionalHelmValuePath struct {
	// the bucket name for the object storage
	Bucket string `json:"bucket"`
	// root path where cluster stores data in the bucket
	Path string `json:"path"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeObjectStorageAdditionalHelmValuePath instantiates a new ModeObjectStorageAdditionalHelmValuePath object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeObjectStorageAdditionalHelmValuePath(bucket string, path string) *ModeObjectStorageAdditionalHelmValuePath {
	this := ModeObjectStorageAdditionalHelmValuePath{}
	this.Bucket = bucket
	this.Path = path
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

// GetPath returns the Path field value.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ModeObjectStorageAdditionalHelmValuePath) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *ModeObjectStorageAdditionalHelmValuePath) SetPath(v string) {
	o.Path = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeObjectStorageAdditionalHelmValuePath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["bucket"] = o.Bucket
	toSerialize["path"] = o.Path

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeObjectStorageAdditionalHelmValuePath) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bucket *string `json:"bucket"`
		Path   *string `json:"path"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Bucket == nil {
		return fmt.Errorf("required field bucket missing")
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"bucket", "path"})
	} else {
		return err
	}
	o.Bucket = *all.Bucket
	o.Path = *all.Path

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
