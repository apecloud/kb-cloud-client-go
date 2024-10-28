// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RegionCreate The region that needs to be created.
type RegionCreate struct {
	// The name of the region.
	Name string `json:"name"`
	// The Chinese name of the region.
	NameCn string `json:"nameCN"`
	// The English name of the region.
	NameEn string `json:"nameEN"`
	// Whether the region is enabled.
	Enabled bool `json:"enabled"`
	// The group of the region.
	Group string `json:"group"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRegionCreate instantiates a new RegionCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRegionCreate(name string, nameCn string, nameEn string, enabled bool, group string) *RegionCreate {
	this := RegionCreate{}
	this.Name = name
	this.NameCn = nameCn
	this.NameEn = nameEn
	this.Enabled = enabled
	this.Group = group
	return &this
}

// NewRegionCreateWithDefaults instantiates a new RegionCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRegionCreateWithDefaults() *RegionCreate {
	this := RegionCreate{}
	return &this
}

// GetName returns the Name field value.
func (o *RegionCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RegionCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RegionCreate) SetName(v string) {
	o.Name = v
}

// GetNameCn returns the NameCn field value.
func (o *RegionCreate) GetNameCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameCn
}

// GetNameCnOk returns a tuple with the NameCn field value
// and a boolean to check if the value has been set.
func (o *RegionCreate) GetNameCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameCn, true
}

// SetNameCn sets field value.
func (o *RegionCreate) SetNameCn(v string) {
	o.NameCn = v
}

// GetNameEn returns the NameEn field value.
func (o *RegionCreate) GetNameEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value
// and a boolean to check if the value has been set.
func (o *RegionCreate) GetNameEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameEn, true
}

// SetNameEn sets field value.
func (o *RegionCreate) SetNameEn(v string) {
	o.NameEn = v
}

// GetEnabled returns the Enabled field value.
func (o *RegionCreate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RegionCreate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *RegionCreate) SetEnabled(v bool) {
	o.Enabled = v
}

// GetGroup returns the Group field value.
func (o *RegionCreate) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *RegionCreate) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value.
func (o *RegionCreate) SetGroup(v string) {
	o.Group = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RegionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["nameCN"] = o.NameCn
	toSerialize["nameEN"] = o.NameEn
	toSerialize["enabled"] = o.Enabled
	toSerialize["group"] = o.Group

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RegionCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name    *string `json:"name"`
		NameCn  *string `json:"nameCN"`
		NameEn  *string `json:"nameEN"`
		Enabled *bool   `json:"enabled"`
		Group   *string `json:"group"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.NameCn == nil {
		return fmt.Errorf("required field nameCN missing")
	}
	if all.NameEn == nil {
		return fmt.Errorf("required field nameEN missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Group == nil {
		return fmt.Errorf("required field group missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "nameCN", "nameEN", "enabled", "group"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.NameCn = *all.NameCn
	o.NameEn = *all.NameEn
	o.Enabled = *all.Enabled
	o.Group = *all.Group

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
