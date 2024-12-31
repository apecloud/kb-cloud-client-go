// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EnvironmentStatusHistory EventEnvironmentState contains details for the current and previous state of the environment.
type EnvironmentStatusHistory struct {
	// The current state of the environment.
	CurrentState string `json:"CurrentState"`
	// The previous state of the environment.
	PrevState string `json:"PrevState"`
	// The reason for the state transition.
	Reason string `json:"Reason"`
	// The timestamp of the state transition.
	Timestamp time.Time `json:"Timestamp"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEnvironmentStatusHistory instantiates a new EnvironmentStatusHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEnvironmentStatusHistory(currentState string, prevState string, reason string, timestamp time.Time) *EnvironmentStatusHistory {
	this := EnvironmentStatusHistory{}
	this.CurrentState = currentState
	this.PrevState = prevState
	this.Reason = reason
	this.Timestamp = timestamp
	return &this
}

// NewEnvironmentStatusHistoryWithDefaults instantiates a new EnvironmentStatusHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEnvironmentStatusHistoryWithDefaults() *EnvironmentStatusHistory {
	this := EnvironmentStatusHistory{}
	return &this
}

// GetCurrentState returns the CurrentState field value.
func (o *EnvironmentStatusHistory) GetCurrentState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CurrentState
}

// GetCurrentStateOk returns a tuple with the CurrentState field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistory) GetCurrentStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentState, true
}

// SetCurrentState sets field value.
func (o *EnvironmentStatusHistory) SetCurrentState(v string) {
	o.CurrentState = v
}

// GetPrevState returns the PrevState field value.
func (o *EnvironmentStatusHistory) GetPrevState() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrevState
}

// GetPrevStateOk returns a tuple with the PrevState field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistory) GetPrevStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrevState, true
}

// SetPrevState sets field value.
func (o *EnvironmentStatusHistory) SetPrevState(v string) {
	o.PrevState = v
}

// GetReason returns the Reason field value.
func (o *EnvironmentStatusHistory) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistory) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *EnvironmentStatusHistory) SetReason(v string) {
	o.Reason = v
}

// GetTimestamp returns the Timestamp field value.
func (o *EnvironmentStatusHistory) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatusHistory) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *EnvironmentStatusHistory) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EnvironmentStatusHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["CurrentState"] = o.CurrentState
	toSerialize["PrevState"] = o.PrevState
	toSerialize["Reason"] = o.Reason
	if o.Timestamp.Nanosecond() == 0 {
		toSerialize["Timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["Timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EnvironmentStatusHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentState *string    `json:"CurrentState"`
		PrevState    *string    `json:"PrevState"`
		Reason       *string    `json:"Reason"`
		Timestamp    *time.Time `json:"Timestamp"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CurrentState == nil {
		return fmt.Errorf("required field CurrentState missing")
	}
	if all.PrevState == nil {
		return fmt.Errorf("required field PrevState missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field Reason missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field Timestamp missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"CurrentState", "PrevState", "Reason", "Timestamp"})
	} else {
		return err
	}
	o.CurrentState = *all.CurrentState
	o.PrevState = *all.PrevState
	o.Reason = *all.Reason
	o.Timestamp = *all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
