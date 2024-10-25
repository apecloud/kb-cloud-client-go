// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// NODESCRIPTION AutoInspection
type AutoInspection struct {
	// NODESCRIPTION Id
	Id *int32 `json:"id,omitempty"`
	// NODESCRIPTION OrgName
	OrgName *string `json:"orgName,omitempty"`
	// NODESCRIPTION UpdatedAt
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// NODESCRIPTION CreatedAt
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// NODESCRIPTION Creator
	Creator *string `json:"creator,omitempty"`
	// NODESCRIPTION Schedule
	Schedule *string `json:"schedule,omitempty"`
	// Specifies the unit of time for the auto inspection schedule.
	RunEvery *AutoInspectionRunUnit `json:"runEvery,omitempty"`
	// NODESCRIPTION DaysOfWeek
	DaysOfWeek []int32 `json:"daysOfWeek,omitempty"`
	// NODESCRIPTION DaysOfMonth
	DaysOfMonth []int32 `json:"daysOfMonth,omitempty"`
	// NODESCRIPTION Hour
	Hour *int32 `json:"hour,omitempty"`
	// NODESCRIPTION Minute
	Minute *int32 `json:"minute,omitempty"`
	// NODESCRIPTION SavedDays
	SavedDays *int32 `json:"savedDays,omitempty"`
	// NODESCRIPTION NextRunTime
	NextRunTime *time.Time `json:"nextRunTime,omitempty"`
	// NODESCRIPTION Enabled
	Enabled *bool `json:"enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutoInspection instantiates a new AutoInspection object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutoInspection() *AutoInspection {
	this := AutoInspection{}
	return &this
}

// NewAutoInspectionWithDefaults instantiates a new AutoInspection object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutoInspectionWithDefaults() *AutoInspection {
	this := AutoInspection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutoInspection) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutoInspection) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AutoInspection) SetId(v int32) {
	o.Id = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AutoInspection) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AutoInspection) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AutoInspection) SetOrgName(v string) {
	o.OrgName = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AutoInspection) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AutoInspection) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AutoInspection) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AutoInspection) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AutoInspection) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AutoInspection) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *AutoInspection) GetCreator() string {
	if o == nil || o.Creator == nil {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetCreatorOk() (*string, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *AutoInspection) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *AutoInspection) SetCreator(v string) {
	o.Creator = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *AutoInspection) GetSchedule() string {
	if o == nil || o.Schedule == nil {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetScheduleOk() (*string, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *AutoInspection) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *AutoInspection) SetSchedule(v string) {
	o.Schedule = &v
}

// GetRunEvery returns the RunEvery field value if set, zero value otherwise.
func (o *AutoInspection) GetRunEvery() AutoInspectionRunUnit {
	if o == nil || o.RunEvery == nil {
		var ret AutoInspectionRunUnit
		return ret
	}
	return *o.RunEvery
}

// GetRunEveryOk returns a tuple with the RunEvery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetRunEveryOk() (*AutoInspectionRunUnit, bool) {
	if o == nil || o.RunEvery == nil {
		return nil, false
	}
	return o.RunEvery, true
}

// HasRunEvery returns a boolean if a field has been set.
func (o *AutoInspection) HasRunEvery() bool {
	return o != nil && o.RunEvery != nil
}

// SetRunEvery gets a reference to the given AutoInspectionRunUnit and assigns it to the RunEvery field.
func (o *AutoInspection) SetRunEvery(v AutoInspectionRunUnit) {
	o.RunEvery = &v
}

// GetDaysOfWeek returns the DaysOfWeek field value if set, zero value otherwise.
func (o *AutoInspection) GetDaysOfWeek() []int32 {
	if o == nil || o.DaysOfWeek == nil {
		var ret []int32
		return ret
	}
	return o.DaysOfWeek
}

// GetDaysOfWeekOk returns a tuple with the DaysOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetDaysOfWeekOk() (*[]int32, bool) {
	if o == nil || o.DaysOfWeek == nil {
		return nil, false
	}
	return &o.DaysOfWeek, true
}

// HasDaysOfWeek returns a boolean if a field has been set.
func (o *AutoInspection) HasDaysOfWeek() bool {
	return o != nil && o.DaysOfWeek != nil
}

// SetDaysOfWeek gets a reference to the given []int32 and assigns it to the DaysOfWeek field.
func (o *AutoInspection) SetDaysOfWeek(v []int32) {
	o.DaysOfWeek = v
}

// GetDaysOfMonth returns the DaysOfMonth field value if set, zero value otherwise.
func (o *AutoInspection) GetDaysOfMonth() []int32 {
	if o == nil || o.DaysOfMonth == nil {
		var ret []int32
		return ret
	}
	return o.DaysOfMonth
}

// GetDaysOfMonthOk returns a tuple with the DaysOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetDaysOfMonthOk() (*[]int32, bool) {
	if o == nil || o.DaysOfMonth == nil {
		return nil, false
	}
	return &o.DaysOfMonth, true
}

// HasDaysOfMonth returns a boolean if a field has been set.
func (o *AutoInspection) HasDaysOfMonth() bool {
	return o != nil && o.DaysOfMonth != nil
}

// SetDaysOfMonth gets a reference to the given []int32 and assigns it to the DaysOfMonth field.
func (o *AutoInspection) SetDaysOfMonth(v []int32) {
	o.DaysOfMonth = v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *AutoInspection) GetHour() int32 {
	if o == nil || o.Hour == nil {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetHourOk() (*int32, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *AutoInspection) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *AutoInspection) SetHour(v int32) {
	o.Hour = &v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *AutoInspection) GetMinute() int32 {
	if o == nil || o.Minute == nil {
		var ret int32
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetMinuteOk() (*int32, bool) {
	if o == nil || o.Minute == nil {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *AutoInspection) HasMinute() bool {
	return o != nil && o.Minute != nil
}

// SetMinute gets a reference to the given int32 and assigns it to the Minute field.
func (o *AutoInspection) SetMinute(v int32) {
	o.Minute = &v
}

// GetSavedDays returns the SavedDays field value if set, zero value otherwise.
func (o *AutoInspection) GetSavedDays() int32 {
	if o == nil || o.SavedDays == nil {
		var ret int32
		return ret
	}
	return *o.SavedDays
}

// GetSavedDaysOk returns a tuple with the SavedDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetSavedDaysOk() (*int32, bool) {
	if o == nil || o.SavedDays == nil {
		return nil, false
	}
	return o.SavedDays, true
}

// HasSavedDays returns a boolean if a field has been set.
func (o *AutoInspection) HasSavedDays() bool {
	return o != nil && o.SavedDays != nil
}

// SetSavedDays gets a reference to the given int32 and assigns it to the SavedDays field.
func (o *AutoInspection) SetSavedDays(v int32) {
	o.SavedDays = &v
}

// GetNextRunTime returns the NextRunTime field value if set, zero value otherwise.
func (o *AutoInspection) GetNextRunTime() time.Time {
	if o == nil || o.NextRunTime == nil {
		var ret time.Time
		return ret
	}
	return *o.NextRunTime
}

// GetNextRunTimeOk returns a tuple with the NextRunTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetNextRunTimeOk() (*time.Time, bool) {
	if o == nil || o.NextRunTime == nil {
		return nil, false
	}
	return o.NextRunTime, true
}

// HasNextRunTime returns a boolean if a field has been set.
func (o *AutoInspection) HasNextRunTime() bool {
	return o != nil && o.NextRunTime != nil
}

// SetNextRunTime gets a reference to the given time.Time and assigns it to the NextRunTime field.
func (o *AutoInspection) SetNextRunTime(v time.Time) {
	o.NextRunTime = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AutoInspection) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspection) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AutoInspection) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AutoInspection) SetEnabled(v bool) {
	o.Enabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutoInspection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.RunEvery != nil {
		toSerialize["runEvery"] = o.RunEvery
	}
	if o.DaysOfWeek != nil {
		toSerialize["daysOfWeek"] = o.DaysOfWeek
	}
	if o.DaysOfMonth != nil {
		toSerialize["daysOfMonth"] = o.DaysOfMonth
	}
	if o.Hour != nil {
		toSerialize["hour"] = o.Hour
	}
	if o.Minute != nil {
		toSerialize["minute"] = o.Minute
	}
	if o.SavedDays != nil {
		toSerialize["savedDays"] = o.SavedDays
	}
	if o.NextRunTime != nil {
		if o.NextRunTime.Nanosecond() == 0 {
			toSerialize["nextRunTime"] = o.NextRunTime.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["nextRunTime"] = o.NextRunTime.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutoInspection) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id          *int32                 `json:"id,omitempty"`
		OrgName     *string                `json:"orgName,omitempty"`
		UpdatedAt   *time.Time             `json:"updatedAt,omitempty"`
		CreatedAt   *time.Time             `json:"createdAt,omitempty"`
		Creator     *string                `json:"creator,omitempty"`
		Schedule    *string                `json:"schedule,omitempty"`
		RunEvery    *AutoInspectionRunUnit `json:"runEvery,omitempty"`
		DaysOfWeek  []int32                `json:"daysOfWeek,omitempty"`
		DaysOfMonth []int32                `json:"daysOfMonth,omitempty"`
		Hour        *int32                 `json:"hour,omitempty"`
		Minute      *int32                 `json:"minute,omitempty"`
		SavedDays   *int32                 `json:"savedDays,omitempty"`
		NextRunTime *time.Time             `json:"nextRunTime,omitempty"`
		Enabled     *bool                  `json:"enabled,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "orgName", "updatedAt", "createdAt", "creator", "schedule", "runEvery", "daysOfWeek", "daysOfMonth", "hour", "minute", "savedDays", "nextRunTime", "enabled"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.OrgName = all.OrgName
	o.UpdatedAt = all.UpdatedAt
	o.CreatedAt = all.CreatedAt
	o.Creator = all.Creator
	o.Schedule = all.Schedule
	if all.RunEvery != nil && !all.RunEvery.IsValid() {
		hasInvalidField = true
	} else {
		o.RunEvery = all.RunEvery
	}
	o.DaysOfWeek = all.DaysOfWeek
	o.DaysOfMonth = all.DaysOfMonth
	o.Hour = all.Hour
	o.Minute = all.Minute
	o.SavedDays = all.SavedDays
	o.NextRunTime = all.NextRunTime
	o.Enabled = all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
