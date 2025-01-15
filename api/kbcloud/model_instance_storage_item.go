// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceStorageItem Instance storage information
type InstanceStorageItem struct {
	// Specify the name of storage, which must be unique.
	Name string `json:"name"`
	// Specify the size of storage, the unit is Gi.
	Size string `json:"size"`
	// The name of StorageClass in use
	StorageClass *string `json:"storageClass,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceStorageItem instantiates a new InstanceStorageItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceStorageItem(name string, size string) *InstanceStorageItem {
	this := InstanceStorageItem{}
	this.Name = name
	this.Size = size
	return &this
}

// NewInstanceStorageItemWithDefaults instantiates a new InstanceStorageItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceStorageItemWithDefaults() *InstanceStorageItem {
	this := InstanceStorageItem{}
	return &this
}

// GetName returns the Name field value.
func (o *InstanceStorageItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InstanceStorageItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *InstanceStorageItem) SetName(v string) {
	o.Name = v
}

// GetSize returns the Size field value.
func (o *InstanceStorageItem) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *InstanceStorageItem) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *InstanceStorageItem) SetSize(v string) {
	o.Size = v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *InstanceStorageItem) GetStorageClass() string {
	if o == nil || o.StorageClass == nil {
		var ret string
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceStorageItem) GetStorageClassOk() (*string, bool) {
	if o == nil || o.StorageClass == nil {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *InstanceStorageItem) HasStorageClass() bool {
	return o != nil && o.StorageClass != nil
}

// SetStorageClass gets a reference to the given string and assigns it to the StorageClass field.
func (o *InstanceStorageItem) SetStorageClass(v string) {
	o.StorageClass = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceStorageItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["size"] = o.Size
	if o.StorageClass != nil {
		toSerialize["storageClass"] = o.StorageClass
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceStorageItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string `json:"name"`
		Size         *string `json:"size"`
		StorageClass *string `json:"storageClass,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Size == nil {
		return fmt.Errorf("required field size missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "size", "storageClass"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Size = *all.Size
	o.StorageClass = all.StorageClass

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
