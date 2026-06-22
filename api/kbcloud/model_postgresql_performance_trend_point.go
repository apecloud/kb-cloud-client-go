// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlPerformanceTrendPoint struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlPerformanceTrendPoint instantiates a new PostgresqlPerformanceTrendPoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlPerformanceTrendPoint(timestamp string, value float64) *PostgresqlPerformanceTrendPoint {
	this := PostgresqlPerformanceTrendPoint{}
	this.Timestamp = timestamp
	this.Value = value
	return &this
}

// NewPostgresqlPerformanceTrendPointWithDefaults instantiates a new PostgresqlPerformanceTrendPoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlPerformanceTrendPointWithDefaults() *PostgresqlPerformanceTrendPoint {
	this := PostgresqlPerformanceTrendPoint{}
	return &this
}

// GetTimestamp returns the Timestamp field value.
func (o *PostgresqlPerformanceTrendPoint) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendPoint) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *PostgresqlPerformanceTrendPoint) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetValue returns the Value field value.
func (o *PostgresqlPerformanceTrendPoint) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PostgresqlPerformanceTrendPoint) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *PostgresqlPerformanceTrendPoint) SetValue(v float64) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlPerformanceTrendPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlPerformanceTrendPoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Timestamp *string  `json:"timestamp"`
		Value     *float64 `json:"value"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"timestamp", "value"})
	} else {
		return err
	}
	o.Timestamp = *all.Timestamp
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
