// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InspectionTaskBatchCreate struct {
	// Organizations to inspect. Every organization must be accessible to the caller.
	OrgNames []string `json:"orgNames"`
	// Engine names to inspect within each selected organization. Omit or pass an empty array to inspect all engines. Organizations without matching non-stopped clusters are skipped.
	Engines []string `json:"engines,omitempty"`
	// Start of the inspection window. If either bound is omitted, defaults to the preceding 24 hours.
	TimeRangeStart *time.Time `json:"timeRangeStart,omitempty"`
	TimeRangeEnd   *time.Time `json:"timeRangeEnd,omitempty"`
	// Report retention in days. Zero uses the existing default retention policy.
	SavedDays *int64 `json:"savedDays,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInspectionTaskBatchCreate instantiates a new InspectionTaskBatchCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInspectionTaskBatchCreate(orgNames []string) *InspectionTaskBatchCreate {
	this := InspectionTaskBatchCreate{}
	this.OrgNames = orgNames
	return &this
}

// NewInspectionTaskBatchCreateWithDefaults instantiates a new InspectionTaskBatchCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInspectionTaskBatchCreateWithDefaults() *InspectionTaskBatchCreate {
	this := InspectionTaskBatchCreate{}
	return &this
}

// GetOrgNames returns the OrgNames field value.
func (o *InspectionTaskBatchCreate) GetOrgNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OrgNames
}

// GetOrgNamesOk returns a tuple with the OrgNames field value
// and a boolean to check if the value has been set.
func (o *InspectionTaskBatchCreate) GetOrgNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgNames, true
}

// SetOrgNames sets field value.
func (o *InspectionTaskBatchCreate) SetOrgNames(v []string) {
	o.OrgNames = v
}

// GetEngines returns the Engines field value if set, zero value otherwise.
func (o *InspectionTaskBatchCreate) GetEngines() []string {
	if o == nil || o.Engines == nil {
		var ret []string
		return ret
	}
	return o.Engines
}

// GetEnginesOk returns a tuple with the Engines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskBatchCreate) GetEnginesOk() (*[]string, bool) {
	if o == nil || o.Engines == nil {
		return nil, false
	}
	return &o.Engines, true
}

// HasEngines returns a boolean if a field has been set.
func (o *InspectionTaskBatchCreate) HasEngines() bool {
	return o != nil && o.Engines != nil
}

// SetEngines gets a reference to the given []string and assigns it to the Engines field.
func (o *InspectionTaskBatchCreate) SetEngines(v []string) {
	o.Engines = v
}

// GetTimeRangeStart returns the TimeRangeStart field value if set, zero value otherwise.
func (o *InspectionTaskBatchCreate) GetTimeRangeStart() time.Time {
	if o == nil || o.TimeRangeStart == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeRangeStart
}

// GetTimeRangeStartOk returns a tuple with the TimeRangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskBatchCreate) GetTimeRangeStartOk() (*time.Time, bool) {
	if o == nil || o.TimeRangeStart == nil {
		return nil, false
	}
	return o.TimeRangeStart, true
}

// HasTimeRangeStart returns a boolean if a field has been set.
func (o *InspectionTaskBatchCreate) HasTimeRangeStart() bool {
	return o != nil && o.TimeRangeStart != nil
}

// SetTimeRangeStart gets a reference to the given time.Time and assigns it to the TimeRangeStart field.
func (o *InspectionTaskBatchCreate) SetTimeRangeStart(v time.Time) {
	o.TimeRangeStart = &v
}

// GetTimeRangeEnd returns the TimeRangeEnd field value if set, zero value otherwise.
func (o *InspectionTaskBatchCreate) GetTimeRangeEnd() time.Time {
	if o == nil || o.TimeRangeEnd == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeRangeEnd
}

// GetTimeRangeEndOk returns a tuple with the TimeRangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskBatchCreate) GetTimeRangeEndOk() (*time.Time, bool) {
	if o == nil || o.TimeRangeEnd == nil {
		return nil, false
	}
	return o.TimeRangeEnd, true
}

// HasTimeRangeEnd returns a boolean if a field has been set.
func (o *InspectionTaskBatchCreate) HasTimeRangeEnd() bool {
	return o != nil && o.TimeRangeEnd != nil
}

// SetTimeRangeEnd gets a reference to the given time.Time and assigns it to the TimeRangeEnd field.
func (o *InspectionTaskBatchCreate) SetTimeRangeEnd(v time.Time) {
	o.TimeRangeEnd = &v
}

// GetSavedDays returns the SavedDays field value if set, zero value otherwise.
func (o *InspectionTaskBatchCreate) GetSavedDays() int64 {
	if o == nil || o.SavedDays == nil {
		var ret int64
		return ret
	}
	return *o.SavedDays
}

// GetSavedDaysOk returns a tuple with the SavedDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InspectionTaskBatchCreate) GetSavedDaysOk() (*int64, bool) {
	if o == nil || o.SavedDays == nil {
		return nil, false
	}
	return o.SavedDays, true
}

// HasSavedDays returns a boolean if a field has been set.
func (o *InspectionTaskBatchCreate) HasSavedDays() bool {
	return o != nil && o.SavedDays != nil
}

// SetSavedDays gets a reference to the given int64 and assigns it to the SavedDays field.
func (o *InspectionTaskBatchCreate) SetSavedDays(v int64) {
	o.SavedDays = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InspectionTaskBatchCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["orgNames"] = o.OrgNames
	if o.Engines != nil {
		toSerialize["engines"] = o.Engines
	}
	if o.TimeRangeStart != nil {
		if o.TimeRangeStart.Nanosecond() == 0 {
			toSerialize["timeRangeStart"] = o.TimeRangeStart.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timeRangeStart"] = o.TimeRangeStart.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.TimeRangeEnd != nil {
		if o.TimeRangeEnd.Nanosecond() == 0 {
			toSerialize["timeRangeEnd"] = o.TimeRangeEnd.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timeRangeEnd"] = o.TimeRangeEnd.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.SavedDays != nil {
		toSerialize["savedDays"] = o.SavedDays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InspectionTaskBatchCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgNames       *[]string  `json:"orgNames"`
		Engines        []string   `json:"engines,omitempty"`
		TimeRangeStart *time.Time `json:"timeRangeStart,omitempty"`
		TimeRangeEnd   *time.Time `json:"timeRangeEnd,omitempty"`
		SavedDays      *int64     `json:"savedDays,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OrgNames == nil {
		return fmt.Errorf("required field orgNames missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"orgNames", "engines", "timeRangeStart", "timeRangeEnd", "savedDays"})
	} else {
		return err
	}
	o.OrgNames = *all.OrgNames
	o.Engines = all.Engines
	o.TimeRangeStart = all.TimeRangeStart
	o.TimeRangeEnd = all.TimeRangeEnd
	o.SavedDays = all.SavedDays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
