// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ImportBackup import backuprecords from storage
type ImportBackup struct {
	// the name of bucket
	Bucket *string `json:"bucket,omitempty"`
	// the path to backup's kubeblocks-backup.json
	Path *string `json:"path,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewImportBackup instantiates a new ImportBackup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewImportBackup() *ImportBackup {
	this := ImportBackup{}
	return &this
}

// NewImportBackupWithDefaults instantiates a new ImportBackup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewImportBackupWithDefaults() *ImportBackup {
	this := ImportBackup{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *ImportBackup) GetBucket() string {
	if o == nil || o.Bucket == nil {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBackup) GetBucketOk() (*string, bool) {
	if o == nil || o.Bucket == nil {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *ImportBackup) HasBucket() bool {
	return o != nil && o.Bucket != nil
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *ImportBackup) SetBucket(v string) {
	o.Bucket = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ImportBackup) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportBackup) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ImportBackup) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ImportBackup) SetPath(v string) {
	o.Path = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ImportBackup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Bucket != nil {
		toSerialize["bucket"] = o.Bucket
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ImportBackup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bucket *string `json:"bucket,omitempty"`
		Path   *string `json:"path,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"bucket", "path"})
	} else {
		return err
	}
	o.Bucket = all.Bucket
	o.Path = all.Path

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
