// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RegionUpdate The region that needs to be updated.

type RegionUpdate struct {
	// The Chinese name of the region.
	NameCn string `json:"nameCN"`
	// The English name of the region.
	NameEn string `json:"nameEN"`
	// The group of the region.
	Group string `json:"group"`
	// Whether the region is enabled.
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRegionUpdate instantiates a new RegionUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRegionUpdate(nameCn string, nameEn string, group string, enabled bool) *RegionUpdate {
	this := RegionUpdate{}
	this.NameCn = nameCn
	this.NameEn = nameEn
	this.Group = group
	this.Enabled = enabled
	return &this
}

// NewRegionUpdateWithDefaults instantiates a new RegionUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRegionUpdateWithDefaults() *RegionUpdate {
	this := RegionUpdate{}
	return &this
}

// GetNameCn returns the NameCn field value.
func (o *RegionUpdate) GetNameCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameCn
}

// GetNameCnOk returns a tuple with the NameCn field value
// and a boolean to check if the value has been set.
func (o *RegionUpdate) GetNameCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameCn, true
}

// SetNameCn sets field value.
func (o *RegionUpdate) SetNameCn(v string) {
	o.NameCn = v
}

// GetNameEn returns the NameEn field value.
func (o *RegionUpdate) GetNameEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value
// and a boolean to check if the value has been set.
func (o *RegionUpdate) GetNameEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameEn, true
}

// SetNameEn sets field value.
func (o *RegionUpdate) SetNameEn(v string) {
	o.NameEn = v
}

// GetGroup returns the Group field value.
func (o *RegionUpdate) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *RegionUpdate) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value.
func (o *RegionUpdate) SetGroup(v string) {
	o.Group = v
}

// GetEnabled returns the Enabled field value.
func (o *RegionUpdate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RegionUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *RegionUpdate) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RegionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["nameCN"] = o.NameCn
	toSerialize["nameEN"] = o.NameEn
	toSerialize["group"] = o.Group
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RegionUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NameCn  *string `json:"nameCN"`
		NameEn  *string `json:"nameEN"`
		Group   *string `json:"group"`
		Enabled *bool   `json:"enabled"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.NameCn == nil {
		return fmt.Errorf("required field nameCN missing")
	}
	if all.NameEn == nil {
		return fmt.Errorf("required field nameEN missing")
	}
	if all.Group == nil {
		return fmt.Errorf("required field group missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nameCN", "nameEN", "group", "enabled"})
	} else {
		return err
	}
	o.NameCn = *all.NameCn
	o.NameEn = *all.NameEn
	o.Group = *all.Group
	o.Enabled = *all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
