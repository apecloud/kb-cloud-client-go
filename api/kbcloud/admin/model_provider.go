// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Provider The cloud provider that the environment is running on.
type Provider struct {
	// The name of the cloud provider.
	Name string `json:"name"`
	// The Chinese name of the cloud provider.
	NameCn string `json:"nameCN"`
	// The English name of the cloud provider.
	NameEn string `json:"nameEN"`
	// The logo of the cloud provider.
	Logo string `json:"logo"`
	// Whether the cloud provider is enabled.
	Enabled bool `json:"enabled"`
	// Whether the cloud provider supports ARN.
	SupportArn bool `json:"supportARN"`
	// The number of environments that the cloud provider has.
	EnvironmentCount int32 `json:"environmentCount"`
	// The number of regions that the cloud provider has.
	RegionCount int32 `json:"regionCount"`
	// The number of zones that the cloud provider has.
	ZoneCount int32 `json:"zoneCount"`
	// The time when the cloud provider was created.
	CreatedAt time.Time `json:"createdAt"`
	// The time when the cloud provider was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProvider instantiates a new Provider object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProvider(name string, nameCn string, nameEn string, logo string, enabled bool, supportArn bool, environmentCount int32, regionCount int32, zoneCount int32, createdAt time.Time, updatedAt time.Time) *Provider {
	this := Provider{}
	this.Name = name
	this.NameCn = nameCn
	this.NameEn = nameEn
	this.Logo = logo
	this.Enabled = enabled
	this.SupportArn = supportArn
	this.EnvironmentCount = environmentCount
	this.RegionCount = regionCount
	this.ZoneCount = zoneCount
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewProviderWithDefaults instantiates a new Provider object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProviderWithDefaults() *Provider {
	this := Provider{}
	return &this
}

// GetName returns the Name field value.
func (o *Provider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Provider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Provider) SetName(v string) {
	o.Name = v
}

// GetNameCn returns the NameCn field value.
func (o *Provider) GetNameCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameCn
}

// GetNameCnOk returns a tuple with the NameCn field value
// and a boolean to check if the value has been set.
func (o *Provider) GetNameCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameCn, true
}

// SetNameCn sets field value.
func (o *Provider) SetNameCn(v string) {
	o.NameCn = v
}

// GetNameEn returns the NameEn field value.
func (o *Provider) GetNameEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value
// and a boolean to check if the value has been set.
func (o *Provider) GetNameEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameEn, true
}

// SetNameEn sets field value.
func (o *Provider) SetNameEn(v string) {
	o.NameEn = v
}

// GetLogo returns the Logo field value.
func (o *Provider) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *Provider) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value.
func (o *Provider) SetLogo(v string) {
	o.Logo = v
}

// GetEnabled returns the Enabled field value.
func (o *Provider) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Provider) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *Provider) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSupportArn returns the SupportArn field value.
func (o *Provider) GetSupportArn() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.SupportArn
}

// GetSupportArnOk returns a tuple with the SupportArn field value
// and a boolean to check if the value has been set.
func (o *Provider) GetSupportArnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportArn, true
}

// SetSupportArn sets field value.
func (o *Provider) SetSupportArn(v bool) {
	o.SupportArn = v
}

// GetEnvironmentCount returns the EnvironmentCount field value.
func (o *Provider) GetEnvironmentCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.EnvironmentCount
}

// GetEnvironmentCountOk returns a tuple with the EnvironmentCount field value
// and a boolean to check if the value has been set.
func (o *Provider) GetEnvironmentCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentCount, true
}

// SetEnvironmentCount sets field value.
func (o *Provider) SetEnvironmentCount(v int32) {
	o.EnvironmentCount = v
}

// GetRegionCount returns the RegionCount field value.
func (o *Provider) GetRegionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.RegionCount
}

// GetRegionCountOk returns a tuple with the RegionCount field value
// and a boolean to check if the value has been set.
func (o *Provider) GetRegionCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionCount, true
}

// SetRegionCount sets field value.
func (o *Provider) SetRegionCount(v int32) {
	o.RegionCount = v
}

// GetZoneCount returns the ZoneCount field value.
func (o *Provider) GetZoneCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ZoneCount
}

// GetZoneCountOk returns a tuple with the ZoneCount field value
// and a boolean to check if the value has been set.
func (o *Provider) GetZoneCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneCount, true
}

// SetZoneCount sets field value.
func (o *Provider) SetZoneCount(v int32) {
	o.ZoneCount = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *Provider) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Provider) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Provider) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *Provider) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Provider) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *Provider) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Provider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["nameCN"] = o.NameCn
	toSerialize["nameEN"] = o.NameEn
	toSerialize["logo"] = o.Logo
	toSerialize["enabled"] = o.Enabled
	toSerialize["supportARN"] = o.SupportArn
	toSerialize["environmentCount"] = o.EnvironmentCount
	toSerialize["regionCount"] = o.RegionCount
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
func (o *Provider) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name             *string    `json:"name"`
		NameCn           *string    `json:"nameCN"`
		NameEn           *string    `json:"nameEN"`
		Logo             *string    `json:"logo"`
		Enabled          *bool      `json:"enabled"`
		SupportArn       *bool      `json:"supportARN"`
		EnvironmentCount *int32     `json:"environmentCount"`
		RegionCount      *int32     `json:"regionCount"`
		ZoneCount        *int32     `json:"zoneCount"`
		CreatedAt        *time.Time `json:"createdAt"`
		UpdatedAt        *time.Time `json:"updatedAt"`
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
	if all.Logo == nil {
		return fmt.Errorf("required field logo missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.SupportArn == nil {
		return fmt.Errorf("required field supportARN missing")
	}
	if all.EnvironmentCount == nil {
		return fmt.Errorf("required field environmentCount missing")
	}
	if all.RegionCount == nil {
		return fmt.Errorf("required field regionCount missing")
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
		common.DeleteKeys(additionalProperties, &[]string{"name", "nameCN", "nameEN", "logo", "enabled", "supportARN", "environmentCount", "regionCount", "zoneCount", "createdAt", "updatedAt"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.NameCn = *all.NameCn
	o.NameEn = *all.NameEn
	o.Logo = *all.Logo
	o.Enabled = *all.Enabled
	o.SupportArn = *all.SupportArn
	o.EnvironmentCount = *all.EnvironmentCount
	o.RegionCount = *all.RegionCount
	o.ZoneCount = *all.ZoneCount
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
