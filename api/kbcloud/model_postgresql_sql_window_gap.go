// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type PostgresqlSQLWindowGap struct {
	StartTime *string `json:"startTime,omitempty"`
	EndTime   *string `json:"endTime,omitempty"`
	// A collection status or read failure cause. collection_gap indicates a proven chain break or an unverifiable tail longer than three periods. Only proven missing predecessors increase expectedCollections.
	Reason *string `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowGap instantiates a new PostgresqlSQLWindowGap object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowGap() *PostgresqlSQLWindowGap {
	this := PostgresqlSQLWindowGap{}
	return &this
}

// NewPostgresqlSQLWindowGapWithDefaults instantiates a new PostgresqlSQLWindowGap object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowGapWithDefaults() *PostgresqlSQLWindowGap {
	this := PostgresqlSQLWindowGap{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowGap) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowGap) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowGap) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *PostgresqlSQLWindowGap) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowGap) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowGap) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowGap) HasEndTime() bool {
	return o != nil && o.EndTime != nil
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *PostgresqlSQLWindowGap) SetEndTime(v string) {
	o.EndTime = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowGap) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowGap) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowGap) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *PostgresqlSQLWindowGap) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowGap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLWindowGap) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StartTime *string `json:"startTime,omitempty"`
		EndTime   *string `json:"endTime,omitempty"`
		Reason    *string `json:"reason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"startTime", "endTime", "reason"})
	} else {
		return err
	}
	o.StartTime = all.StartTime
	o.EndTime = all.EndTime
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
