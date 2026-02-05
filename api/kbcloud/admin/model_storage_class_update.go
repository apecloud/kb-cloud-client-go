// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StorageClassUpdate StorageClassUpdate provides the fields for updating a specific storage class.
type StorageClassUpdate struct {
	// A user-friendly name for the storage class.
	DisplayName string `json:"displayName"`
	// A detailed description of the storage class.
	Description string `json:"description"`
	// Indicates if the storage class is enabled.
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStorageClassUpdate instantiates a new StorageClassUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStorageClassUpdate(displayName string, description string, enabled bool) *StorageClassUpdate {
	this := StorageClassUpdate{}
	this.DisplayName = displayName
	this.Description = description
	this.Enabled = enabled
	return &this
}

// NewStorageClassUpdateWithDefaults instantiates a new StorageClassUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStorageClassUpdateWithDefaults() *StorageClassUpdate {
	this := StorageClassUpdate{}
	return &this
}

// GetDisplayName returns the DisplayName field value.
func (o *StorageClassUpdate) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *StorageClassUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *StorageClassUpdate) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDescription returns the Description field value.
func (o *StorageClassUpdate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *StorageClassUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *StorageClassUpdate) SetDescription(v string) {
	o.Description = v
}

// GetEnabled returns the Enabled field value.
func (o *StorageClassUpdate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *StorageClassUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *StorageClassUpdate) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o StorageClassUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["description"] = o.Description
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StorageClassUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string `json:"displayName"`
		Description *string `json:"description"`
		Enabled     *bool   `json:"enabled"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field displayName missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"displayName", "description", "enabled"})
	} else {
		return err
	}
	o.DisplayName = *all.DisplayName
	o.Description = *all.Description
	o.Enabled = *all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
