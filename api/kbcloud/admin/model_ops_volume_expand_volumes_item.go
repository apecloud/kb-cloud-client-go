// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OpsVolumeExpandVolumesItem struct {
	// volume name
	Name string `json:"name"`
	// Storage size, the unit is Gi.
	Storage string `json:"storage"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsVolumeExpandVolumesItem instantiates a new OpsVolumeExpandVolumesItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsVolumeExpandVolumesItem(name string, storage string) *OpsVolumeExpandVolumesItem {
	this := OpsVolumeExpandVolumesItem{}
	this.Name = name
	this.Storage = storage
	return &this
}

// NewOpsVolumeExpandVolumesItemWithDefaults instantiates a new OpsVolumeExpandVolumesItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsVolumeExpandVolumesItemWithDefaults() *OpsVolumeExpandVolumesItem {
	this := OpsVolumeExpandVolumesItem{}
	return &this
}

// GetName returns the Name field value.
func (o *OpsVolumeExpandVolumesItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpsVolumeExpandVolumesItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OpsVolumeExpandVolumesItem) SetName(v string) {
	o.Name = v
}

// GetStorage returns the Storage field value.
func (o *OpsVolumeExpandVolumesItem) GetStorage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *OpsVolumeExpandVolumesItem) GetStorageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value.
func (o *OpsVolumeExpandVolumesItem) SetStorage(v string) {
	o.Storage = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsVolumeExpandVolumesItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["storage"] = o.Storage

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsVolumeExpandVolumesItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string `json:"name"`
		Storage *string `json:"storage"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Storage == nil {
		return fmt.Errorf("required field storage missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "storage"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Storage = *all.Storage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
