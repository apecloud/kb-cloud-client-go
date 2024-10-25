// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Region The region that the cloud provider has.

type Region struct {
	// The name of the region.
	Name string `json:"name"`
	// The Chinese name of the region.
	NameCn string `json:"nameCN"`
	// The English name of the region.
	NameEn string `json:"nameEN"`
	// The name of the cloud provider.
	Provider string `json:"provider"`
	// The group of the region.
	Group string `json:"group"`
	// Whether the region is enabled.
	Enabled bool `json:"enabled"`
	// The number of zones that the region has.
	ZoneCount int32 `json:"zoneCount"`
	// The time when the region was created.
	CreatedAt time.Time `json:"createdAt"`
	// The time when the region was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRegion instantiates a new Region object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRegion(name string, nameCn string, nameEn string, provider string, group string, enabled bool, zoneCount int32, createdAt time.Time, updatedAt time.Time) *Region {
	this := Region{}
	this.Name = name
	this.NameCn = nameCn
	this.NameEn = nameEn
	this.Provider = provider
	this.Group = group
	this.Enabled = enabled
	this.ZoneCount = zoneCount
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewRegionWithDefaults instantiates a new Region object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRegionWithDefaults() *Region {
	this := Region{}
	return &this
}

// GetName returns the Name field value.
func (o *Region) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Region) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Region) SetName(v string) {
	o.Name = v
}

// GetNameCn returns the NameCn field value.
func (o *Region) GetNameCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameCn
}

// GetNameCnOk returns a tuple with the NameCn field value
// and a boolean to check if the value has been set.
func (o *Region) GetNameCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameCn, true
}

// SetNameCn sets field value.
func (o *Region) SetNameCn(v string) {
	o.NameCn = v
}

// GetNameEn returns the NameEn field value.
func (o *Region) GetNameEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value
// and a boolean to check if the value has been set.
func (o *Region) GetNameEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameEn, true
}

// SetNameEn sets field value.
func (o *Region) SetNameEn(v string) {
	o.NameEn = v
}

// GetProvider returns the Provider field value.
func (o *Region) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *Region) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *Region) SetProvider(v string) {
	o.Provider = v
}

// GetGroup returns the Group field value.
func (o *Region) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *Region) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value.
func (o *Region) SetGroup(v string) {
	o.Group = v
}

// GetEnabled returns the Enabled field value.
func (o *Region) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Region) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *Region) SetEnabled(v bool) {
	o.Enabled = v
}

// GetZoneCount returns the ZoneCount field value.
func (o *Region) GetZoneCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ZoneCount
}

// GetZoneCountOk returns a tuple with the ZoneCount field value
// and a boolean to check if the value has been set.
func (o *Region) GetZoneCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneCount, true
}

// SetZoneCount sets field value.
func (o *Region) SetZoneCount(v int32) {
	o.ZoneCount = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Region) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Region) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Region) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *Region) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Region) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *Region) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Region) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["nameCN"] = o.NameCn
	toSerialize["nameEN"] = o.NameEn
	toSerialize["provider"] = o.Provider
	toSerialize["group"] = o.Group
	toSerialize["enabled"] = o.Enabled
	toSerialize["zoneCount"] = o.ZoneCount
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Region) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string    `json:"name"`
		NameCn    *string    `json:"nameCN"`
		NameEn    *string    `json:"nameEN"`
		Provider  *string    `json:"provider"`
		Group     *string    `json:"group"`
		Enabled   *bool      `json:"enabled"`
		ZoneCount *int32     `json:"zoneCount"`
		CreatedAt *time.Time `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
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
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.Group == nil {
		return fmt.Errorf("required field group missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.ZoneCount == nil {
		return fmt.Errorf("required field zoneCount missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "nameCN", "nameEN", "provider", "group", "enabled", "zoneCount", "createdAt", "updatedAt"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.NameCn = *all.NameCn
	o.NameEn = *all.NameEn
	o.Provider = *all.Provider
	o.Group = *all.Group
	o.Enabled = *all.Enabled
	o.ZoneCount = *all.ZoneCount
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
