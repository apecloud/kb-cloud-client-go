// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ZoneUpdate The zone that needs to be updated.

type ZoneUpdate struct {
	// The Chinese name of the zone.
	NameCn string `json:"nameCN"`
	// The English name of the zone.
	NameEn string `json:"nameEN"`
	// Whether the zone is enabled.
	Enabled bool `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewZoneUpdate instantiates a new ZoneUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewZoneUpdate(nameCn string, nameEn string, enabled bool) *ZoneUpdate {
	this := ZoneUpdate{}
	this.NameCn = nameCn
	this.NameEn = nameEn
	this.Enabled = enabled
	return &this
}

// NewZoneUpdateWithDefaults instantiates a new ZoneUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewZoneUpdateWithDefaults() *ZoneUpdate {
	this := ZoneUpdate{}
	return &this
}

// GetNameCn returns the NameCn field value.
func (o *ZoneUpdate) GetNameCn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameCn
}

// GetNameCnOk returns a tuple with the NameCn field value
// and a boolean to check if the value has been set.
func (o *ZoneUpdate) GetNameCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameCn, true
}

// SetNameCn sets field value.
func (o *ZoneUpdate) SetNameCn(v string) {
	o.NameCn = v
}

// GetNameEn returns the NameEn field value.
func (o *ZoneUpdate) GetNameEn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.NameEn
}

// GetNameEnOk returns a tuple with the NameEn field value
// and a boolean to check if the value has been set.
func (o *ZoneUpdate) GetNameEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameEn, true
}

// SetNameEn sets field value.
func (o *ZoneUpdate) SetNameEn(v string) {
	o.NameEn = v
}

// GetEnabled returns the Enabled field value.
func (o *ZoneUpdate) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ZoneUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ZoneUpdate) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ZoneUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["nameCN"] = o.NameCn
	toSerialize["nameEN"] = o.NameEn
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ZoneUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NameCn  *string `json:"nameCN"`
		NameEn  *string `json:"nameEN"`
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
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nameCN", "nameEN", "enabled"})
	} else {
		return err
	}
	o.NameCn = *all.NameCn
	o.NameEn = *all.NameEn
	o.Enabled = *all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
