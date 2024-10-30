// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// Zone The zone that the region has. 
type Zone struct {
	// The name of the zone.
	Name string `json:"name"`
	// The Chinese name of the zone.
	NameCn string `json:"nameCN"`
	// The English name of the zone.
	NameEn string `json:"nameEN"`
	// Whether the zone is enabled.
	Enabled bool `json:"enabled"`
	// The name of the region.
	Region string `json:"region"`
	// The time when the zone was created.
	CreatedAt time.Time `json:"createdAt"`
	// The time when the zone was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewZone instantiates a new Zone object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewZone(name string, nameCn string, nameEn string, enabled bool, region string, createdAt time.Time, updatedAt time.Time) *Zone {
	this := Zone{}
	this.Name = name
	this.NameCn = nameCn
	this.NameEn = nameEn
	this.Enabled = enabled
	this.Region = region
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewZoneWithDefaults instantiates a new Zone object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewZoneWithDefaults() *Zone {
	this := Zone{}
	return &this
}
// GetName returns the Name field value.
func (o *Zone) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Zone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Zone) SetName(v string) {
	o.Name = v
}


// GetNameCn returns the NameCn field value.
func (o *Zone) GetNameCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameCn
}

// GetNameCnOk returns a tuple with the NameCn field value
// and a boolean to check if the value has been set.
func (o *Zone) GetNameCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameCn, true
}

// SetNameCn sets field value.
func (o *Zone) SetNameCn(v string) {
	o.NameCn = v
}


// GetNameEn returns the NameEn field value.
func (o *Zone) GetNameEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value
// and a boolean to check if the value has been set.
func (o *Zone) GetNameEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameEn, true
}

// SetNameEn sets field value.
func (o *Zone) SetNameEn(v string) {
	o.NameEn = v
}


// GetEnabled returns the Enabled field value.
func (o *Zone) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Zone) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *Zone) SetEnabled(v bool) {
	o.Enabled = v
}


// GetRegion returns the Region field value.
func (o *Zone) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Zone) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *Zone) SetRegion(v string) {
	o.Region = v
}


// GetCreatedAt returns the CreatedAt field value.
func (o *Zone) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Zone) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *Zone) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}


// GetUpdatedAt returns the UpdatedAt field value.
func (o *Zone) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Zone) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *Zone) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}



// MarshalJSON serializes the struct using spec logic.
func (o Zone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["nameCN"] = o.NameCn
	toSerialize["nameEN"] = o.NameEn
	toSerialize["enabled"] = o.Enabled
	toSerialize["region"] = o.Region
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
func (o *Zone) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name"`
		NameCn *string `json:"nameCN"`
		NameEn *string `json:"nameEN"`
		Enabled *bool `json:"enabled"`
		Region *string `json:"region"`
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
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field createdAt missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updatedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "name", "nameCN", "nameEN", "enabled", "region", "createdAt", "updatedAt",  })
	} else {
		return err
	}
	o.Name = *all.Name
	o.NameCn = *all.NameCn
	o.NameEn = *all.NameEn
	o.Enabled = *all.Enabled
	o.Region = *all.Region
	o.CreatedAt = *all.CreatedAt
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
