// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengSchemaItem struct {
	// Schema (owner) name.
	Name string `json:"name"`
	// Total segment size in bytes across all objects owned by this schema.
	SizeBytes int64 `json:"sizeBytes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengSchemaItem instantiates a new DamengSchemaItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengSchemaItem(name string, sizeBytes int64) *DamengSchemaItem {
	this := DamengSchemaItem{}
	this.Name = name
	this.SizeBytes = sizeBytes
	return &this
}

// NewDamengSchemaItemWithDefaults instantiates a new DamengSchemaItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengSchemaItemWithDefaults() *DamengSchemaItem {
	this := DamengSchemaItem{}
	return &this
}

// GetName returns the Name field value.
func (o *DamengSchemaItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DamengSchemaItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DamengSchemaItem) SetName(v string) {
	o.Name = v
}

// GetSizeBytes returns the SizeBytes field value.
func (o *DamengSchemaItem) GetSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value
// and a boolean to check if the value has been set.
func (o *DamengSchemaItem) GetSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeBytes, true
}

// SetSizeBytes sets field value.
func (o *DamengSchemaItem) SetSizeBytes(v int64) {
	o.SizeBytes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengSchemaItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["sizeBytes"] = o.SizeBytes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengSchemaItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string `json:"name"`
		SizeBytes *int64  `json:"sizeBytes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SizeBytes == nil {
		return fmt.Errorf("required field sizeBytes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "sizeBytes"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.SizeBytes = *all.SizeBytes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
