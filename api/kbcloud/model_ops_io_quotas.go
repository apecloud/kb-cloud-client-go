// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsIoQuotas IO Quotas describes IOPS and BPS for volumes of a KubeBlocks cluster
type OpsIoQuotas struct {
	Component string                   `json:"component"`
	Volumes   []OpsIoQuotasVolumesItem `json:"volumes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsIoQuotas instantiates a new OpsIoQuotas object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsIoQuotas(component string, volumes []OpsIoQuotasVolumesItem) *OpsIoQuotas {
	this := OpsIoQuotas{}
	this.Component = component
	this.Volumes = volumes
	return &this
}

// NewOpsIoQuotasWithDefaults instantiates a new OpsIoQuotas object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsIoQuotasWithDefaults() *OpsIoQuotas {
	this := OpsIoQuotas{}
	return &this
}

// GetComponent returns the Component field value.
func (o *OpsIoQuotas) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *OpsIoQuotas) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value.
func (o *OpsIoQuotas) SetComponent(v string) {
	o.Component = v
}

// GetVolumes returns the Volumes field value.
func (o *OpsIoQuotas) GetVolumes() []OpsIoQuotasVolumesItem {
	if o == nil {
		var ret []OpsIoQuotasVolumesItem
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value
// and a boolean to check if the value has been set.
func (o *OpsIoQuotas) GetVolumesOk() (*[]OpsIoQuotasVolumesItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Volumes, true
}

// SetVolumes sets field value.
func (o *OpsIoQuotas) SetVolumes(v []OpsIoQuotasVolumesItem) {
	o.Volumes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsIoQuotas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["component"] = o.Component
	toSerialize["volumes"] = o.Volumes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsIoQuotas) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Component *string                   `json:"component"`
		Volumes   *[]OpsIoQuotasVolumesItem `json:"volumes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Component == nil {
		return fmt.Errorf("required field component missing")
	}
	if all.Volumes == nil {
		return fmt.Errorf("required field volumes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"component", "volumes"})
	} else {
		return err
	}
	o.Component = *all.Component
	o.Volumes = *all.Volumes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
