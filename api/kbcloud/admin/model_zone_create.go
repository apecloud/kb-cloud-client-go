// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// ZoneCreate The zone that needs to be created. 
type ZoneCreate struct {
	// The name of the zone.
	Name string `json:"name"`
	// The Chinese name of the zone.
	NameCn string `json:"nameCN"`
	// The English name of the zone.
	NameEn string `json:"nameEN"`
	// Whether the zone is enabled.
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewZoneCreate instantiates a new ZoneCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewZoneCreate(name string, nameCn string, nameEn string, enabled bool) *ZoneCreate {
	this := ZoneCreate{}
	this.Name = name
	this.NameCn = nameCn
	this.NameEn = nameEn
	this.Enabled = enabled
	return &this
}

// NewZoneCreateWithDefaults instantiates a new ZoneCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewZoneCreateWithDefaults() *ZoneCreate {
	this := ZoneCreate{}
	return &this
}
// GetName returns the Name field value.
func (o *ZoneCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ZoneCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ZoneCreate) SetName(v string) {
	o.Name = v
}


// GetNameCn returns the NameCn field value.
func (o *ZoneCreate) GetNameCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameCn
}

// GetNameCnOk returns a tuple with the NameCn field value
// and a boolean to check if the value has been set.
func (o *ZoneCreate) GetNameCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameCn, true
}

// SetNameCn sets field value.
func (o *ZoneCreate) SetNameCn(v string) {
	o.NameCn = v
}


// GetNameEn returns the NameEn field value.
func (o *ZoneCreate) GetNameEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value
// and a boolean to check if the value has been set.
func (o *ZoneCreate) GetNameEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameEn, true
}

// SetNameEn sets field value.
func (o *ZoneCreate) SetNameEn(v string) {
	o.NameEn = v
}


// GetEnabled returns the Enabled field value.
func (o *ZoneCreate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ZoneCreate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ZoneCreate) SetEnabled(v bool) {
	o.Enabled = v
}



// MarshalJSON serializes the struct using spec logic.
func (o ZoneCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["nameCN"] = o.NameCn
	toSerialize["nameEN"] = o.NameEn
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ZoneCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name"`
		NameCn *string `json:"nameCN"`
		NameEn *string `json:"nameEN"`
		Enabled *bool `json:"enabled"`
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
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "name", "nameCN", "nameEN", "enabled",  })
	} else {
		return err
	}
	o.Name = *all.Name
	o.NameCn = *all.NameCn
	o.NameEn = *all.NameEn
	o.Enabled = *all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
