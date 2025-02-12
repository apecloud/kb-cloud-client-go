// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsQueryHistory struct {
	// executed sql statements
	Sql *string `json:"sql,omitempty"`
	// sql executed massage
	ErrMsg *string `json:"err_msg,omitempty"`
	// sql executed time
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// sql executed duration
	Duration *int32 `json:"duration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsQueryHistory instantiates a new DmsQueryHistory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsQueryHistory() *DmsQueryHistory {
	this := DmsQueryHistory{}
	return &this
}

// NewDmsQueryHistoryWithDefaults instantiates a new DmsQueryHistory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsQueryHistoryWithDefaults() *DmsQueryHistory {
	this := DmsQueryHistory{}
	return &this
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetSql() string {
	if o == nil || o.Sql == nil {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetSqlOk() (*string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasSql() bool {
	return o != nil && o.Sql != nil
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *DmsQueryHistory) SetSql(v string) {
	o.Sql = &v
}

// GetErrMsg returns the ErrMsg field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetErrMsg() string {
	if o == nil || o.ErrMsg == nil {
		var ret string
		return ret
	}
	return *o.ErrMsg
}

// GetErrMsgOk returns a tuple with the ErrMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetErrMsgOk() (*string, bool) {
	if o == nil || o.ErrMsg == nil {
		return nil, false
	}
	return o.ErrMsg, true
}

// HasErrMsg returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasErrMsg() bool {
	return o != nil && o.ErrMsg != nil
}

// SetErrMsg gets a reference to the given string and assigns it to the ErrMsg field.
func (o *DmsQueryHistory) SetErrMsg(v string) {
	o.ErrMsg = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DmsQueryHistory) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *DmsQueryHistory) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsQueryHistory) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *DmsQueryHistory) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *DmsQueryHistory) SetDuration(v int32) {
	o.Duration = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsQueryHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.ErrMsg != nil {
		toSerialize["err_msg"] = o.ErrMsg
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *DmsQueryHistory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Sql       *string    `json:"sql,omitempty"`
		ErrMsg    *string    `json:"err_msg,omitempty"`
		CreatedAt *time.Time `json:"created_at,omitempty"`
		Duration  *int32     `json:"duration,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sql", "err_msg", "created_at", "duration"})
	} else {
		return err
	}
	o.Sql = all.Sql
	o.ErrMsg = all.ErrMsg
	o.CreatedAt = all.CreatedAt
	o.Duration = all.Duration

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
