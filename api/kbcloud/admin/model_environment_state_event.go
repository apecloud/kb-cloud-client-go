// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentStateEvent EventEnvironmentState contains details for the current and previous state of the environment.
type EnvironmentStateEvent struct {
	// The current state of the environment.
	CurrentState string `json:"currentState"`
	// The previous state of the environment.
	PrevState string `json:"prevState"`
	// The reason for the state transition.
	Reason string `json:"reason"`
	// The timestamp of the state transition.
	Timestamp time.Time `json:"timestamp"`
	// The duration of the state in seconds.
	Duration *int64 `json:"duration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentStateEvent instantiates a new EnvironmentStateEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentStateEvent(currentState string, prevState string, reason string, timestamp time.Time) *EnvironmentStateEvent {
	this := EnvironmentStateEvent{}
	this.CurrentState = currentState
	this.PrevState = prevState
	this.Reason = reason
	this.Timestamp = timestamp
	return &this
}

// NewEnvironmentStateEventWithDefaults instantiates a new EnvironmentStateEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentStateEventWithDefaults() *EnvironmentStateEvent {
	this := EnvironmentStateEvent{}
	return &this
}

// GetCurrentState returns the CurrentState field value.
func (o *EnvironmentStateEvent) GetCurrentState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CurrentState
}

// GetCurrentStateOk returns a tuple with the CurrentState field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStateEvent) GetCurrentStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentState, true
}

// SetCurrentState sets field value.
func (o *EnvironmentStateEvent) SetCurrentState(v string) {
	o.CurrentState = v
}

// GetPrevState returns the PrevState field value.
func (o *EnvironmentStateEvent) GetPrevState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrevState
}

// GetPrevStateOk returns a tuple with the PrevState field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStateEvent) GetPrevStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrevState, true
}

// SetPrevState sets field value.
func (o *EnvironmentStateEvent) SetPrevState(v string) {
	o.PrevState = v
}

// GetReason returns the Reason field value.
func (o *EnvironmentStateEvent) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStateEvent) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *EnvironmentStateEvent) SetReason(v string) {
	o.Reason = v
}

// GetTimestamp returns the Timestamp field value.
func (o *EnvironmentStateEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStateEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *EnvironmentStateEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *EnvironmentStateEvent) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentStateEvent) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *EnvironmentStateEvent) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *EnvironmentStateEvent) SetDuration(v int64) {
	o.Duration = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentStateEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["currentState"] = o.CurrentState
	toSerialize["prevState"] = o.PrevState
	toSerialize["reason"] = o.Reason
	if o.Timestamp.Nanosecond() == 0 {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentStateEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentState *string    `json:"currentState"`
		PrevState    *string    `json:"prevState"`
		Reason       *string    `json:"reason"`
		Timestamp    *time.Time `json:"timestamp"`
		Duration     *int64     `json:"duration,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CurrentState == nil {
		return fmt.Errorf("required field currentState missing")
	}
	if all.PrevState == nil {
		return fmt.Errorf("required field prevState missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"currentState", "prevState", "reason", "timestamp", "duration"})
	} else {
		return err
	}
	o.CurrentState = *all.CurrentState
	o.PrevState = *all.PrevState
	o.Reason = *all.Reason
	o.Timestamp = *all.Timestamp
	o.Duration = all.Duration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
