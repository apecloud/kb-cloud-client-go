// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLWindow struct {
	StartTime          string  `json:"startTime"`
	EndTime            string  `json:"endTime"`
	EffectiveStartTime *string `json:"effectiveStartTime,omitempty"`
	EffectiveEndTime   *string `json:"effectiveEndTime,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindow instantiates a new PostgresqlSQLWindow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindow(startTime string, endTime string) *PostgresqlSQLWindow {
	this := PostgresqlSQLWindow{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewPostgresqlSQLWindowWithDefaults instantiates a new PostgresqlSQLWindow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowWithDefaults() *PostgresqlSQLWindow {
	this := PostgresqlSQLWindow{}
	return &this
}

// GetStartTime returns the StartTime field value.
func (o *PostgresqlSQLWindow) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindow) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value.
func (o *PostgresqlSQLWindow) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value.
func (o *PostgresqlSQLWindow) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindow) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value.
func (o *PostgresqlSQLWindow) SetEndTime(v string) {
	o.EndTime = v
}

// GetEffectiveStartTime returns the EffectiveStartTime field value if set, zero value otherwise.
func (o *PostgresqlSQLWindow) GetEffectiveStartTime() string {
	if o == nil || o.EffectiveStartTime == nil {
		var ret string
		return ret
	}
	return *o.EffectiveStartTime
}

// GetEffectiveStartTimeOk returns a tuple with the EffectiveStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindow) GetEffectiveStartTimeOk() (*string, bool) {
	if o == nil || o.EffectiveStartTime == nil {
		return nil, false
	}
	return o.EffectiveStartTime, true
}

// HasEffectiveStartTime returns a boolean if a field has been set.
func (o *PostgresqlSQLWindow) HasEffectiveStartTime() bool {
	return o != nil && o.EffectiveStartTime != nil
}

// SetEffectiveStartTime gets a reference to the given string and assigns it to the EffectiveStartTime field.
func (o *PostgresqlSQLWindow) SetEffectiveStartTime(v string) {
	o.EffectiveStartTime = &v
}

// GetEffectiveEndTime returns the EffectiveEndTime field value if set, zero value otherwise.
func (o *PostgresqlSQLWindow) GetEffectiveEndTime() string {
	if o == nil || o.EffectiveEndTime == nil {
		var ret string
		return ret
	}
	return *o.EffectiveEndTime
}

// GetEffectiveEndTimeOk returns a tuple with the EffectiveEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindow) GetEffectiveEndTimeOk() (*string, bool) {
	if o == nil || o.EffectiveEndTime == nil {
		return nil, false
	}
	return o.EffectiveEndTime, true
}

// HasEffectiveEndTime returns a boolean if a field has been set.
func (o *PostgresqlSQLWindow) HasEffectiveEndTime() bool {
	return o != nil && o.EffectiveEndTime != nil
}

// SetEffectiveEndTime gets a reference to the given string and assigns it to the EffectiveEndTime field.
func (o *PostgresqlSQLWindow) SetEffectiveEndTime(v string) {
	o.EffectiveEndTime = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	if o.EffectiveStartTime != nil {
		toSerialize["effectiveStartTime"] = o.EffectiveStartTime
	}
	if o.EffectiveEndTime != nil {
		toSerialize["effectiveEndTime"] = o.EffectiveEndTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLWindow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StartTime          *string `json:"startTime"`
		EndTime            *string `json:"endTime"`
		EffectiveStartTime *string `json:"effectiveStartTime,omitempty"`
		EffectiveEndTime   *string `json:"effectiveEndTime,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.StartTime == nil {
		return fmt.Errorf("required field startTime missing")
	}
	if all.EndTime == nil {
		return fmt.Errorf("required field endTime missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"startTime", "endTime", "effectiveStartTime", "effectiveEndTime"})
	} else {
		return err
	}
	o.StartTime = *all.StartTime
	o.EndTime = *all.EndTime
	o.EffectiveStartTime = all.EffectiveStartTime
	o.EffectiveEndTime = all.EffectiveEndTime

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
