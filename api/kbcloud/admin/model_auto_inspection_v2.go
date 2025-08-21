// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AutoInspectionV2 struct {
	Id *int32 `json:"id,omitempty"`
	// Specifies the type of the resource for the auto inspection.
	ResourceType AutoInspectionResourceTypeV2 `json:"resourceType"`
	ResourceId   *string                      `json:"resourceID,omitempty"`
	ResourceName string                       `json:"resourceName"`
	Creator      string                       `json:"creator"`
	Schedule     *string                      `json:"schedule,omitempty"`
	// Specifies the unit of time for the auto inspection schedule.
	RunEvery    *AutoInspectionRunUnit `json:"runEvery,omitempty"`
	DaysOfWeek  []int32                `json:"daysOfWeek,omitempty"`
	DaysOfMonth []int32                `json:"daysOfMonth,omitempty"`
	Hour        *int32                 `json:"hour,omitempty"`
	Minute      *int32                 `json:"minute,omitempty"`
	SavedDays   *int32                 `json:"savedDays,omitempty"`
	NextRunTime *time.Time             `json:"nextRunTime,omitempty"`
	Enabled     bool                   `json:"enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutoInspectionV2 instantiates a new AutoInspectionV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutoInspectionV2(resourceType AutoInspectionResourceTypeV2, resourceName string, creator string, enabled bool) *AutoInspectionV2 {
	this := AutoInspectionV2{}
	this.ResourceType = resourceType
	this.ResourceName = resourceName
	this.Creator = creator
	this.Enabled = enabled
	return &this
}

// NewAutoInspectionV2WithDefaults instantiates a new AutoInspectionV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutoInspectionV2WithDefaults() *AutoInspectionV2 {
	this := AutoInspectionV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutoInspectionV2) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutoInspectionV2) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AutoInspectionV2) SetId(v int32) {
	o.Id = &v
}

// GetResourceType returns the ResourceType field value.
func (o *AutoInspectionV2) GetResourceType() AutoInspectionResourceTypeV2 {
	if o == nil {
		var ret AutoInspectionResourceTypeV2
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetResourceTypeOk() (*AutoInspectionResourceTypeV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *AutoInspectionV2) SetResourceType(v AutoInspectionResourceTypeV2) {
	o.ResourceType = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *AutoInspectionV2) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *AutoInspectionV2) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *AutoInspectionV2) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceName returns the ResourceName field value.
func (o *AutoInspectionV2) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value.
func (o *AutoInspectionV2) SetResourceName(v string) {
	o.ResourceName = v
}

// GetCreator returns the Creator field value.
func (o *AutoInspectionV2) GetCreator() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetCreatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creator, true
}

// SetCreator sets field value.
func (o *AutoInspectionV2) SetCreator(v string) {
	o.Creator = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *AutoInspectionV2) GetSchedule() string {
	if o == nil || o.Schedule == nil {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetScheduleOk() (*string, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *AutoInspectionV2) HasSchedule() bool {
	return o != nil && o.Schedule != nil
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *AutoInspectionV2) SetSchedule(v string) {
	o.Schedule = &v
}

// GetRunEvery returns the RunEvery field value if set, zero value otherwise.
func (o *AutoInspectionV2) GetRunEvery() AutoInspectionRunUnit {
	if o == nil || o.RunEvery == nil {
		var ret AutoInspectionRunUnit
		return ret
	}
	return *o.RunEvery
}

// GetRunEveryOk returns a tuple with the RunEvery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetRunEveryOk() (*AutoInspectionRunUnit, bool) {
	if o == nil || o.RunEvery == nil {
		return nil, false
	}
	return o.RunEvery, true
}

// HasRunEvery returns a boolean if a field has been set.
func (o *AutoInspectionV2) HasRunEvery() bool {
	return o != nil && o.RunEvery != nil
}

// SetRunEvery gets a reference to the given AutoInspectionRunUnit and assigns it to the RunEvery field.
func (o *AutoInspectionV2) SetRunEvery(v AutoInspectionRunUnit) {
	o.RunEvery = &v
}

// GetDaysOfWeek returns the DaysOfWeek field value if set, zero value otherwise.
func (o *AutoInspectionV2) GetDaysOfWeek() []int32 {
	if o == nil || o.DaysOfWeek == nil {
		var ret []int32
		return ret
	}
	return o.DaysOfWeek
}

// GetDaysOfWeekOk returns a tuple with the DaysOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetDaysOfWeekOk() (*[]int32, bool) {
	if o == nil || o.DaysOfWeek == nil {
		return nil, false
	}
	return &o.DaysOfWeek, true
}

// HasDaysOfWeek returns a boolean if a field has been set.
func (o *AutoInspectionV2) HasDaysOfWeek() bool {
	return o != nil && o.DaysOfWeek != nil
}

// SetDaysOfWeek gets a reference to the given []int32 and assigns it to the DaysOfWeek field.
func (o *AutoInspectionV2) SetDaysOfWeek(v []int32) {
	o.DaysOfWeek = v
}

// GetDaysOfMonth returns the DaysOfMonth field value if set, zero value otherwise.
func (o *AutoInspectionV2) GetDaysOfMonth() []int32 {
	if o == nil || o.DaysOfMonth == nil {
		var ret []int32
		return ret
	}
	return o.DaysOfMonth
}

// GetDaysOfMonthOk returns a tuple with the DaysOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetDaysOfMonthOk() (*[]int32, bool) {
	if o == nil || o.DaysOfMonth == nil {
		return nil, false
	}
	return &o.DaysOfMonth, true
}

// HasDaysOfMonth returns a boolean if a field has been set.
func (o *AutoInspectionV2) HasDaysOfMonth() bool {
	return o != nil && o.DaysOfMonth != nil
}

// SetDaysOfMonth gets a reference to the given []int32 and assigns it to the DaysOfMonth field.
func (o *AutoInspectionV2) SetDaysOfMonth(v []int32) {
	o.DaysOfMonth = v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *AutoInspectionV2) GetHour() int32 {
	if o == nil || o.Hour == nil {
		var ret int32
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetHourOk() (*int32, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *AutoInspectionV2) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given int32 and assigns it to the Hour field.
func (o *AutoInspectionV2) SetHour(v int32) {
	o.Hour = &v
}

// GetMinute returns the Minute field value if set, zero value otherwise.
func (o *AutoInspectionV2) GetMinute() int32 {
	if o == nil || o.Minute == nil {
		var ret int32
		return ret
	}
	return *o.Minute
}

// GetMinuteOk returns a tuple with the Minute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetMinuteOk() (*int32, bool) {
	if o == nil || o.Minute == nil {
		return nil, false
	}
	return o.Minute, true
}

// HasMinute returns a boolean if a field has been set.
func (o *AutoInspectionV2) HasMinute() bool {
	return o != nil && o.Minute != nil
}

// SetMinute gets a reference to the given int32 and assigns it to the Minute field.
func (o *AutoInspectionV2) SetMinute(v int32) {
	o.Minute = &v
}

// GetSavedDays returns the SavedDays field value if set, zero value otherwise.
func (o *AutoInspectionV2) GetSavedDays() int32 {
	if o == nil || o.SavedDays == nil {
		var ret int32
		return ret
	}
	return *o.SavedDays
}

// GetSavedDaysOk returns a tuple with the SavedDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetSavedDaysOk() (*int32, bool) {
	if o == nil || o.SavedDays == nil {
		return nil, false
	}
	return o.SavedDays, true
}

// HasSavedDays returns a boolean if a field has been set.
func (o *AutoInspectionV2) HasSavedDays() bool {
	return o != nil && o.SavedDays != nil
}

// SetSavedDays gets a reference to the given int32 and assigns it to the SavedDays field.
func (o *AutoInspectionV2) SetSavedDays(v int32) {
	o.SavedDays = &v
}

// GetNextRunTime returns the NextRunTime field value if set, zero value otherwise.
func (o *AutoInspectionV2) GetNextRunTime() time.Time {
	if o == nil || o.NextRunTime == nil {
		var ret time.Time
		return ret
	}
	return *o.NextRunTime
}

// GetNextRunTimeOk returns a tuple with the NextRunTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetNextRunTimeOk() (*time.Time, bool) {
	if o == nil || o.NextRunTime == nil {
		return nil, false
	}
	return o.NextRunTime, true
}

// HasNextRunTime returns a boolean if a field has been set.
func (o *AutoInspectionV2) HasNextRunTime() bool {
	return o != nil && o.NextRunTime != nil
}

// SetNextRunTime gets a reference to the given time.Time and assigns it to the NextRunTime field.
func (o *AutoInspectionV2) SetNextRunTime(v time.Time) {
	o.NextRunTime = &v
}

// GetEnabled returns the Enabled field value.
func (o *AutoInspectionV2) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AutoInspectionV2) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *AutoInspectionV2) SetEnabled(v bool) {
	o.Enabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutoInspectionV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["resourceType"] = o.ResourceType
	if o.ResourceId != nil {
		toSerialize["resourceID"] = o.ResourceId
	}
	toSerialize["resourceName"] = o.ResourceName
	toSerialize["creator"] = o.Creator
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
	toSerialize["enabled"] = o.Enabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutoInspectionV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id           *int32                        `json:"id,omitempty"`
		ResourceType *AutoInspectionResourceTypeV2 `json:"resourceType"`
		ResourceId   *string                       `json:"resourceID,omitempty"`
		ResourceName *string                       `json:"resourceName"`
		Creator      *string                       `json:"creator"`
		Schedule     *string                       `json:"schedule,omitempty"`
		RunEvery     *AutoInspectionRunUnit        `json:"runEvery,omitempty"`
		DaysOfWeek   []int32                       `json:"daysOfWeek,omitempty"`
		DaysOfMonth  []int32                       `json:"daysOfMonth,omitempty"`
		Hour         *int32                        `json:"hour,omitempty"`
		Minute       *int32                        `json:"minute,omitempty"`
		SavedDays    *int32                        `json:"savedDays,omitempty"`
		NextRunTime  *time.Time                    `json:"nextRunTime,omitempty"`
		Enabled      *bool                         `json:"enabled"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resourceType missing")
	}
	if all.ResourceName == nil {
		return fmt.Errorf("required field resourceName missing")
	}
	if all.Creator == nil {
		return fmt.Errorf("required field creator missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "resourceType", "resourceID", "resourceName", "creator", "schedule", "runEvery", "daysOfWeek", "daysOfMonth", "hour", "minute", "savedDays", "nextRunTime", "enabled"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	if !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = *all.ResourceType
	}
	o.ResourceId = all.ResourceId
	o.ResourceName = *all.ResourceName
	o.Creator = *all.Creator
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
	o.Enabled = *all.Enabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
