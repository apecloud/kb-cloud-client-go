// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RawBackupRepo backup repo info
type RawBackupRepo struct {
	// the backup repo name
	Name *string `json:"name,omitempty"`
	// default backup repo
	Default *bool `json:"default,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRawBackupRepo instantiates a new RawBackupRepo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRawBackupRepo() *RawBackupRepo {
	this := RawBackupRepo{}
	return &this
}

// NewRawBackupRepoWithDefaults instantiates a new RawBackupRepo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRawBackupRepoWithDefaults() *RawBackupRepo {
	this := RawBackupRepo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RawBackupRepo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawBackupRepo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RawBackupRepo) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RawBackupRepo) SetName(v string) {
	o.Name = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RawBackupRepo) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawBackupRepo) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RawBackupRepo) HasDefault() bool {
	return o != nil && o.Default != nil
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *RawBackupRepo) SetDefault(v bool) {
	o.Default = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RawBackupRepo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RawBackupRepo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string `json:"name,omitempty"`
		Default *bool   `json:"default,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "default"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Default = all.Default

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
