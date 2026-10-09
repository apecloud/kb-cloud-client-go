// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AlertClusterSilence Pause notifications while continuing to record alerts. Set enabled to false to cancel. Time windows include their start and exclude their end.
type AlertClusterSilence struct {
	Enabled bool `json:"enabled"`
	// Alert names to silence. Empty or omitted means all cluster alerts.
	Rules []string `json:"rules,omitempty"`
	// Required when enabled. Once uses absolute timestamps; weekly uses weekdays and local times in timeZone.
	Mode *AlertClusterSilenceMode `json:"mode,omitempty"`
	// Required for once. Absolute start of the silence window.
	StartsAt *time.Time `json:"startsAt,omitempty"`
	// Required for once. Must be later than startsAt.
	EndsAt *time.Time `json:"endsAt,omitempty"`
	// Required for weekly. Sunday is 0. Overnight windows belong to the weekday on which they start.
	Weekdays []int32 `json:"weekdays,omitempty"`
	// Required for weekly, in HH:mm format.
	StartTime *string `json:"startTime,omitempty"`
	// Required for weekly, in HH:mm format. Earlier than startTime means the following day; equal times are invalid.
	EndTime *string `json:"endTime,omitempty"`
	// Required for weekly. IANA timezone, e.g. Asia/Shanghai or UTC.
	TimeZone *string `json:"timeZone,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAlertClusterSilence instantiates a new AlertClusterSilence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAlertClusterSilence(enabled bool) *AlertClusterSilence {
	this := AlertClusterSilence{}
	this.Enabled = enabled
	return &this
}

// NewAlertClusterSilenceWithDefaults instantiates a new AlertClusterSilence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAlertClusterSilenceWithDefaults() *AlertClusterSilence {
	this := AlertClusterSilence{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *AlertClusterSilence) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AlertClusterSilence) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *AlertClusterSilence) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *AlertClusterSilence) GetRules() []string {
	if o == nil || o.Rules == nil {
		var ret []string
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertClusterSilence) GetRulesOk() (*[]string, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *AlertClusterSilence) HasRules() bool {
	return o != nil && o.Rules != nil
}

// SetRules gets a reference to the given []string and assigns it to the Rules field.
func (o *AlertClusterSilence) SetRules(v []string) {
	o.Rules = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *AlertClusterSilence) GetMode() AlertClusterSilenceMode {
	if o == nil || o.Mode == nil {
		var ret AlertClusterSilenceMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertClusterSilence) GetModeOk() (*AlertClusterSilenceMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *AlertClusterSilence) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given AlertClusterSilenceMode and assigns it to the Mode field.
func (o *AlertClusterSilence) SetMode(v AlertClusterSilenceMode) {
	o.Mode = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *AlertClusterSilence) GetStartsAt() time.Time {
	if o == nil || o.StartsAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertClusterSilence) GetStartsAtOk() (*time.Time, bool) {
	if o == nil || o.StartsAt == nil {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *AlertClusterSilence) HasStartsAt() bool {
	return o != nil && o.StartsAt != nil
}

// SetStartsAt gets a reference to the given time.Time and assigns it to the StartsAt field.
func (o *AlertClusterSilence) SetStartsAt(v time.Time) {
	o.StartsAt = &v
}

// GetEndsAt returns the EndsAt field value if set, zero value otherwise.
func (o *AlertClusterSilence) GetEndsAt() time.Time {
	if o == nil || o.EndsAt == nil {
		var ret time.Time
		return ret
	}
	return *o.EndsAt
}

// GetEndsAtOk returns a tuple with the EndsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertClusterSilence) GetEndsAtOk() (*time.Time, bool) {
	if o == nil || o.EndsAt == nil {
		return nil, false
	}
	return o.EndsAt, true
}

// HasEndsAt returns a boolean if a field has been set.
func (o *AlertClusterSilence) HasEndsAt() bool {
	return o != nil && o.EndsAt != nil
}

// SetEndsAt gets a reference to the given time.Time and assigns it to the EndsAt field.
func (o *AlertClusterSilence) SetEndsAt(v time.Time) {
	o.EndsAt = &v
}

// GetWeekdays returns the Weekdays field value if set, zero value otherwise.
func (o *AlertClusterSilence) GetWeekdays() []int32 {
	if o == nil || o.Weekdays == nil {
		var ret []int32
		return ret
	}
	return o.Weekdays
}

// GetWeekdaysOk returns a tuple with the Weekdays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertClusterSilence) GetWeekdaysOk() (*[]int32, bool) {
	if o == nil || o.Weekdays == nil {
		return nil, false
	}
	return &o.Weekdays, true
}

// HasWeekdays returns a boolean if a field has been set.
func (o *AlertClusterSilence) HasWeekdays() bool {
	return o != nil && o.Weekdays != nil
}

// SetWeekdays gets a reference to the given []int32 and assigns it to the Weekdays field.
func (o *AlertClusterSilence) SetWeekdays(v []int32) {
	o.Weekdays = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AlertClusterSilence) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertClusterSilence) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AlertClusterSilence) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *AlertClusterSilence) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AlertClusterSilence) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertClusterSilence) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AlertClusterSilence) HasEndTime() bool {
	return o != nil && o.EndTime != nil
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *AlertClusterSilence) SetEndTime(v string) {
	o.EndTime = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *AlertClusterSilence) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertClusterSilence) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *AlertClusterSilence) HasTimeZone() bool {
	return o != nil && o.TimeZone != nil
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *AlertClusterSilence) SetTimeZone(v string) {
	o.TimeZone = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AlertClusterSilence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.StartsAt != nil {
		if o.StartsAt.Nanosecond() == 0 {
			toSerialize["startsAt"] = o.StartsAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["startsAt"] = o.StartsAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.EndsAt != nil {
		if o.EndsAt.Nanosecond() == 0 {
			toSerialize["endsAt"] = o.EndsAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["endsAt"] = o.EndsAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Weekdays != nil {
		toSerialize["weekdays"] = o.Weekdays
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AlertClusterSilence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled   *bool                    `json:"enabled"`
		Rules     []string                 `json:"rules,omitempty"`
		Mode      *AlertClusterSilenceMode `json:"mode,omitempty"`
		StartsAt  *time.Time               `json:"startsAt,omitempty"`
		EndsAt    *time.Time               `json:"endsAt,omitempty"`
		Weekdays  []int32                  `json:"weekdays,omitempty"`
		StartTime *string                  `json:"startTime,omitempty"`
		EndTime   *string                  `json:"endTime,omitempty"`
		TimeZone  *string                  `json:"timeZone,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enabled", "rules", "mode", "startsAt", "endsAt", "weekdays", "startTime", "endTime", "timeZone"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = *all.Enabled
	o.Rules = all.Rules
	if all.Mode != nil && !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = all.Mode
	}
	o.StartsAt = all.StartsAt
	o.EndsAt = all.EndsAt
	o.Weekdays = all.Weekdays
	o.StartTime = all.StartTime
	o.EndTime = all.EndTime
	o.TimeZone = all.TimeZone

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
