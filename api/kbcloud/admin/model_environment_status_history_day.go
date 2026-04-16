// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentStatusHistoryDay Daily status history for the environment.
type EnvironmentStatusHistoryDay struct {
	// The date of this day's status history.
	Date string `json:"date"`
	// The primary status for this day based on priority.
	Status string `json:"status"`
	// Duration in seconds for each status on this day.
	StatusDuration map[string]float64 `json:"statusDuration"`
	// List of state transition events that occurred on this day.
	StatusHistory []EnvironmentStateEvent `json:"statusHistory"`
	// Indicates if the status was inherited from the previous day.
	Inherited bool `json:"inherited"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentStatusHistoryDay instantiates a new EnvironmentStatusHistoryDay object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentStatusHistoryDay(date string, status string, statusDuration map[string]float64, statusHistory []EnvironmentStateEvent, inherited bool) *EnvironmentStatusHistoryDay {
	this := EnvironmentStatusHistoryDay{}
	this.Date = date
	this.Status = status
	this.StatusDuration = statusDuration
	this.StatusHistory = statusHistory
	this.Inherited = inherited
	return &this
}

// NewEnvironmentStatusHistoryDayWithDefaults instantiates a new EnvironmentStatusHistoryDay object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentStatusHistoryDayWithDefaults() *EnvironmentStatusHistoryDay {
	this := EnvironmentStatusHistoryDay{}
	return &this
}

// GetDate returns the Date field value.
func (o *EnvironmentStatusHistoryDay) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistoryDay) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *EnvironmentStatusHistoryDay) SetDate(v string) {
	o.Date = v
}

// GetStatus returns the Status field value.
func (o *EnvironmentStatusHistoryDay) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistoryDay) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *EnvironmentStatusHistoryDay) SetStatus(v string) {
	o.Status = v
}

// GetStatusDuration returns the StatusDuration field value.
func (o *EnvironmentStatusHistoryDay) GetStatusDuration() map[string]float64 {
	if o == nil {
		var ret map[string]float64
		return ret
	}
	return o.StatusDuration
}

// GetStatusDurationOk returns a tuple with the StatusDuration field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistoryDay) GetStatusDurationOk() (*map[string]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusDuration, true
}

// SetStatusDuration sets field value.
func (o *EnvironmentStatusHistoryDay) SetStatusDuration(v map[string]float64) {
	o.StatusDuration = v
}

// GetStatusHistory returns the StatusHistory field value.
func (o *EnvironmentStatusHistoryDay) GetStatusHistory() []EnvironmentStateEvent {
	if o == nil {
		var ret []EnvironmentStateEvent
		return ret
	}
	return o.StatusHistory
}

// GetStatusHistoryOk returns a tuple with the StatusHistory field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistoryDay) GetStatusHistoryOk() (*[]EnvironmentStateEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusHistory, true
}

// SetStatusHistory sets field value.
func (o *EnvironmentStatusHistoryDay) SetStatusHistory(v []EnvironmentStateEvent) {
	o.StatusHistory = v
}

// GetInherited returns the Inherited field value.
func (o *EnvironmentStatusHistoryDay) GetInherited() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistoryDay) GetInheritedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inherited, true
}

// SetInherited sets field value.
func (o *EnvironmentStatusHistoryDay) SetInherited(v bool) {
	o.Inherited = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentStatusHistoryDay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["date"] = o.Date
	toSerialize["status"] = o.Status
	toSerialize["statusDuration"] = o.StatusDuration
	toSerialize["statusHistory"] = o.StatusHistory
	toSerialize["inherited"] = o.Inherited

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentStatusHistoryDay) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Date           *string                  `json:"date"`
		Status         *string                  `json:"status"`
		StatusDuration *map[string]float64      `json:"statusDuration"`
		StatusHistory  *[]EnvironmentStateEvent `json:"statusHistory"`
		Inherited      *bool                    `json:"inherited"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Date == nil {
		return fmt.Errorf("required field date missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.StatusDuration == nil {
		return fmt.Errorf("required field statusDuration missing")
	}
	if all.StatusHistory == nil {
		return fmt.Errorf("required field statusHistory missing")
	}
	if all.Inherited == nil {
		return fmt.Errorf("required field inherited missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"date", "status", "statusDuration", "statusHistory", "inherited"})
	} else {
		return err
	}
	o.Date = *all.Date
	o.Status = *all.Status
	o.StatusDuration = *all.StatusDuration
	o.StatusHistory = *all.StatusHistory
	o.Inherited = *all.Inherited

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
