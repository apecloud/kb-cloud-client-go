// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlPerformanceTrendRange struct {
	// Controlled query range label. Values are 1h, 6h, or 24h.
	Label           string `json:"label"`
	Start           string `json:"start"`
	End             string `json:"end"`
	DurationSeconds int64  `json:"durationSeconds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlPerformanceTrendRange instantiates a new PostgresqlPerformanceTrendRange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlPerformanceTrendRange(label string, start string, end string, durationSeconds int64) *PostgresqlPerformanceTrendRange {
	this := PostgresqlPerformanceTrendRange{}
	this.Label = label
	this.Start = start
	this.End = end
	this.DurationSeconds = durationSeconds
	return &this
}

// NewPostgresqlPerformanceTrendRangeWithDefaults instantiates a new PostgresqlPerformanceTrendRange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlPerformanceTrendRangeWithDefaults() *PostgresqlPerformanceTrendRange {
	this := PostgresqlPerformanceTrendRange{}
	return &this
}

// GetLabel returns the Label field value.
func (o *PostgresqlPerformanceTrendRange) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendRange) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value.
func (o *PostgresqlPerformanceTrendRange) SetLabel(v string) {
	o.Label = v
}

// GetStart returns the Start field value.
func (o *PostgresqlPerformanceTrendRange) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendRange) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value.
func (o *PostgresqlPerformanceTrendRange) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value.
func (o *PostgresqlPerformanceTrendRange) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendRange) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value.
func (o *PostgresqlPerformanceTrendRange) SetEnd(v string) {
	o.End = v
}

// GetDurationSeconds returns the DurationSeconds field value.
func (o *PostgresqlPerformanceTrendRange) GetDurationSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendRange) GetDurationSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationSeconds, true
}

// SetDurationSeconds sets field value.
func (o *PostgresqlPerformanceTrendRange) SetDurationSeconds(v int64) {
	o.DurationSeconds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlPerformanceTrendRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["label"] = o.Label
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End
	toSerialize["durationSeconds"] = o.DurationSeconds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlPerformanceTrendRange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Label           *string `json:"label"`
		Start           *string `json:"start"`
		End             *string `json:"end"`
		DurationSeconds *int64  `json:"durationSeconds"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Label == nil {
		return fmt.Errorf("required field label missing")
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
		common.DeleteKeys(additionalProperties, &[]string{"label", "start", "end", "durationSeconds"})
	} else {
		return err
	}
	o.Label = *all.Label
	o.Start = *all.Start
	o.End = *all.End
	o.DurationSeconds = *all.DurationSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
