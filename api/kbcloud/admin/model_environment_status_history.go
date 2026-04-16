// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentStatusHistory Environment status history with SLA calculation.
type EnvironmentStatusHistory struct {
	// The SLA percentage (availability) for the environment.
	Sla string `json:"sla"`
	// Daily status history data.
	Days []EnvironmentStatusHistoryDay `json:"days"`
	// Total duration in seconds for each status across all days.
	StatusDuration map[string]float64 `json:"statusDuration"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentStatusHistory instantiates a new EnvironmentStatusHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentStatusHistory(sla string, days []EnvironmentStatusHistoryDay, statusDuration map[string]float64) *EnvironmentStatusHistory {
	this := EnvironmentStatusHistory{}
	this.Sla = sla
	this.Days = days
	this.StatusDuration = statusDuration
	return &this
}

// NewEnvironmentStatusHistoryWithDefaults instantiates a new EnvironmentStatusHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentStatusHistoryWithDefaults() *EnvironmentStatusHistory {
	this := EnvironmentStatusHistory{}
	return &this
}

// GetSla returns the Sla field value.
func (o *EnvironmentStatusHistory) GetSla() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sla
}

// GetSlaOk returns a tuple with the Sla field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistory) GetSlaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sla, true
}

// SetSla sets field value.
func (o *EnvironmentStatusHistory) SetSla(v string) {
	o.Sla = v
}

// GetDays returns the Days field value.
func (o *EnvironmentStatusHistory) GetDays() []EnvironmentStatusHistoryDay {
	if o == nil {
		var ret []EnvironmentStatusHistoryDay
		return ret
	}
	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistory) GetDaysOk() (*[]EnvironmentStatusHistoryDay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Days, true
}

// SetDays sets field value.
func (o *EnvironmentStatusHistory) SetDays(v []EnvironmentStatusHistoryDay) {
	o.Days = v
}

// GetStatusDuration returns the StatusDuration field value.
func (o *EnvironmentStatusHistory) GetStatusDuration() map[string]float64 {
	if o == nil {
		var ret map[string]float64
		return ret
	}
	return o.StatusDuration
}

// GetStatusDurationOk returns a tuple with the StatusDuration field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistory) GetStatusDurationOk() (*map[string]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDuration, true
}

// SetStatusDuration sets field value.
func (o *EnvironmentStatusHistory) SetStatusDuration(v map[string]float64) {
	o.StatusDuration = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentStatusHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["sla"] = o.Sla
	toSerialize["days"] = o.Days
	toSerialize["statusDuration"] = o.StatusDuration

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentStatusHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Sla            *string                        `json:"sla"`
		Days           *[]EnvironmentStatusHistoryDay `json:"days"`
		StatusDuration *map[string]float64            `json:"statusDuration"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Sla == nil {
		return fmt.Errorf("required field sla missing")
	}
	if all.Days == nil {
		return fmt.Errorf("required field days missing")
	}
	if all.StatusDuration == nil {
		return fmt.Errorf("required field statusDuration missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sla", "days", "statusDuration"})
	} else {
		return err
	}
	o.Sla = *all.Sla
	o.Days = *all.Days
	o.StatusDuration = *all.StatusDuration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
