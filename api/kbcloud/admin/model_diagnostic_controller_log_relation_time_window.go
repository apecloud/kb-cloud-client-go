// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticControllerLogRelationTimeWindow Time range covered by one controller-log relation card. Together with
// lastOccurredAt, this is the recency contract for Console display.
type DiagnosticControllerLogRelationTimeWindow struct {
	// Earliest representative log timestamp or collection-window start.
	StartedAt time.Time `json:"startedAt"`
	// Latest representative log timestamp or collection-window end.
	EndedAt time.Time `json:"endedAt"`
	// Window duration in seconds. Must be positive when present.
	DurationSeconds int64 `json:"durationSeconds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticControllerLogRelationTimeWindow instantiates a new DiagnosticControllerLogRelationTimeWindow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticControllerLogRelationTimeWindow(startedAt time.Time, endedAt time.Time, durationSeconds int64) *DiagnosticControllerLogRelationTimeWindow {
	this := DiagnosticControllerLogRelationTimeWindow{}
	this.StartedAt = startedAt
	this.EndedAt = endedAt
	this.DurationSeconds = durationSeconds
	return &this
}

// NewDiagnosticControllerLogRelationTimeWindowWithDefaults instantiates a new DiagnosticControllerLogRelationTimeWindow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticControllerLogRelationTimeWindowWithDefaults() *DiagnosticControllerLogRelationTimeWindow {
	this := DiagnosticControllerLogRelationTimeWindow{}
	return &this
}

// GetStartedAt returns the StartedAt field value.
func (o *DiagnosticControllerLogRelationTimeWindow) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelationTimeWindow) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value.
func (o *DiagnosticControllerLogRelationTimeWindow) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetEndedAt returns the EndedAt field value.
func (o *DiagnosticControllerLogRelationTimeWindow) GetEndedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelationTimeWindow) GetEndedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndedAt, true
}

// SetEndedAt sets field value.
func (o *DiagnosticControllerLogRelationTimeWindow) SetEndedAt(v time.Time) {
	o.EndedAt = v
}

// GetDurationSeconds returns the DurationSeconds field value.
func (o *DiagnosticControllerLogRelationTimeWindow) GetDurationSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRelationTimeWindow) GetDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationSeconds, true
}

// SetDurationSeconds sets field value.
func (o *DiagnosticControllerLogRelationTimeWindow) SetDurationSeconds(v int64) {
	o.DurationSeconds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticControllerLogRelationTimeWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.StartedAt.Nanosecond() == 0 {
		toSerialize["startedAt"] = o.StartedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["startedAt"] = o.StartedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.EndedAt.Nanosecond() == 0 {
		toSerialize["endedAt"] = o.EndedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["endedAt"] = o.EndedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["durationSeconds"] = o.DurationSeconds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticControllerLogRelationTimeWindow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StartedAt       *time.Time `json:"startedAt"`
		EndedAt         *time.Time `json:"endedAt"`
		DurationSeconds *int64     `json:"durationSeconds"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.StartedAt == nil {
		return fmt.Errorf("required field startedAt missing")
	}
	if all.EndedAt == nil {
		return fmt.Errorf("required field endedAt missing")
	}
	if all.DurationSeconds == nil {
		return fmt.Errorf("required field durationSeconds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"startedAt", "endedAt", "durationSeconds"})
	} else {
		return err
	}
	o.StartedAt = *all.StartedAt
	o.EndedAt = *all.EndedAt
	o.DurationSeconds = *all.DurationSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
