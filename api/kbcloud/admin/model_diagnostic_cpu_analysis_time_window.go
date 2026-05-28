// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysisTimeWindow Time window represented by the CPU analysis evidence.
type DiagnosticCPUAnalysisTimeWindow struct {
	// Start time of the evidence window.
	Start time.Time `json:"start"`
	// End time of the evidence window.
	End time.Time `json:"end"`
	// Window length in seconds.
	DurationSeconds int64 `json:"durationSeconds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticCPUAnalysisTimeWindow instantiates a new DiagnosticCPUAnalysisTimeWindow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticCPUAnalysisTimeWindow(start time.Time, end time.Time, durationSeconds int64) *DiagnosticCPUAnalysisTimeWindow {
	this := DiagnosticCPUAnalysisTimeWindow{}
	this.Start = start
	this.End = end
	this.DurationSeconds = durationSeconds
	return &this
}

// NewDiagnosticCPUAnalysisTimeWindowWithDefaults instantiates a new DiagnosticCPUAnalysisTimeWindow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticCPUAnalysisTimeWindowWithDefaults() *DiagnosticCPUAnalysisTimeWindow {
	this := DiagnosticCPUAnalysisTimeWindow{}
	return &this
}

// GetStart returns the Start field value.
func (o *DiagnosticCPUAnalysisTimeWindow) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisTimeWindow) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *DiagnosticCPUAnalysisTimeWindow) SetStart(v time.Time) {
	o.Start = v
}

// GetEnd returns the End field value.
func (o *DiagnosticCPUAnalysisTimeWindow) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisTimeWindow) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *DiagnosticCPUAnalysisTimeWindow) SetEnd(v time.Time) {
	o.End = v
}

// GetDurationSeconds returns the DurationSeconds field value.
func (o *DiagnosticCPUAnalysisTimeWindow) GetDurationSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisTimeWindow) GetDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationSeconds, true
}

// SetDurationSeconds sets field value.
func (o *DiagnosticCPUAnalysisTimeWindow) SetDurationSeconds(v int64) {
	o.DurationSeconds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticCPUAnalysisTimeWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Start.Nanosecond() == 0 {
		toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["start"] = o.Start.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.End.Nanosecond() == 0 {
		toSerialize["end"] = o.End.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["end"] = o.End.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["durationSeconds"] = o.DurationSeconds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticCPUAnalysisTimeWindow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Start           *time.Time `json:"start"`
		End             *time.Time `json:"end"`
		DurationSeconds *int64     `json:"durationSeconds"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Start == nil {
		return fmt.Errorf("required field start missing")
	}
	if all.End == nil {
		return fmt.Errorf("required field end missing")
	}
	if all.DurationSeconds == nil {
		return fmt.Errorf("required field durationSeconds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"start", "end", "durationSeconds"})
	} else {
		return err
	}
	o.Start = *all.Start
	o.End = *all.End
	o.DurationSeconds = *all.DurationSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
