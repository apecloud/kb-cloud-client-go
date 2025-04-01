// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OpsIoQuotasVolumesItem struct {
	// volume name
	Name string `json:"name"`
	// IO Quantity describes IOPS and BPS of a volume
	IoReserves *IoQuantity `json:"ioReserves,omitempty"`
	// IO Quantity describes IOPS and BPS of a volume
	IoLimits *IoQuantity `json:"ioLimits,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsIoQuotasVolumesItem instantiates a new OpsIoQuotasVolumesItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsIoQuotasVolumesItem(name string) *OpsIoQuotasVolumesItem {
	this := OpsIoQuotasVolumesItem{}
	this.Name = name
	return &this
}

// NewOpsIoQuotasVolumesItemWithDefaults instantiates a new OpsIoQuotasVolumesItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsIoQuotasVolumesItemWithDefaults() *OpsIoQuotasVolumesItem {
	this := OpsIoQuotasVolumesItem{}
	return &this
}

// GetName returns the Name field value.
func (o *OpsIoQuotasVolumesItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpsIoQuotasVolumesItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *OpsIoQuotasVolumesItem) SetName(v string) {
	o.Name = v
}

// GetIoReserves returns the IoReserves field value if set, zero value otherwise.
func (o *OpsIoQuotasVolumesItem) GetIoReserves() IoQuantity {
	if o == nil || o.IoReserves == nil {
		var ret IoQuantity
		return ret
	}
	return *o.IoReserves
}

// GetIoReservesOk returns a tuple with the IoReserves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsIoQuotasVolumesItem) GetIoReservesOk() (*IoQuantity, bool) {
	if o == nil || o.IoReserves == nil {
		return nil, false
	}
	return o.IoReserves, true
}

// HasIoReserves returns a boolean if a field has been set.
func (o *OpsIoQuotasVolumesItem) HasIoReserves() bool {
	return o != nil && o.IoReserves != nil
}

// SetIoReserves gets a reference to the given IoQuantity and assigns it to the IoReserves field.
func (o *OpsIoQuotasVolumesItem) SetIoReserves(v IoQuantity) {
	o.IoReserves = &v
}

// GetIoLimits returns the IoLimits field value if set, zero value otherwise.
func (o *OpsIoQuotasVolumesItem) GetIoLimits() IoQuantity {
	if o == nil || o.IoLimits == nil {
		var ret IoQuantity
		return ret
	}
	return *o.IoLimits
}

// GetIoLimitsOk returns a tuple with the IoLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsIoQuotasVolumesItem) GetIoLimitsOk() (*IoQuantity, bool) {
	if o == nil || o.IoLimits == nil {
		return nil, false
	}
	return o.IoLimits, true
}

// HasIoLimits returns a boolean if a field has been set.
func (o *OpsIoQuotasVolumesItem) HasIoLimits() bool {
	return o != nil && o.IoLimits != nil
}

// SetIoLimits gets a reference to the given IoQuantity and assigns it to the IoLimits field.
func (o *OpsIoQuotasVolumesItem) SetIoLimits(v IoQuantity) {
	o.IoLimits = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsIoQuotasVolumesItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.IoReserves != nil {
		toSerialize["ioReserves"] = o.IoReserves
	}
	if o.IoLimits != nil {
		toSerialize["ioLimits"] = o.IoLimits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsIoQuotasVolumesItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name       *string     `json:"name"`
		IoReserves *IoQuantity `json:"ioReserves,omitempty"`
		IoLimits   *IoQuantity `json:"ioLimits,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "ioReserves", "ioLimits"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.IoReserves != nil && all.IoReserves.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IoReserves = all.IoReserves
	if all.IoLimits != nil && all.IoLimits.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IoLimits = all.IoLimits

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
