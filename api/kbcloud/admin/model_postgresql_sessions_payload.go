// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSessionsPayload struct {
	Stats    PostgresqlSessionsStats `json:"stats"`
	Sessions []PostgresqlSession     `json:"sessions"`
	// Redacted plain text context for copy actions.
	CopyContext string `json:"copyContext"`
	// Effective maximum number of sessions requested from pg_stat_activity.
	SessionLimit int64 `json:"sessionLimit"`
	// Effective threshold used for long_running tags and statistics.
	LongRunningThresholdSeconds int64 `json:"longRunningThresholdSeconds"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSessionsPayload instantiates a new PostgresqlSessionsPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSessionsPayload(stats PostgresqlSessionsStats, sessions []PostgresqlSession, copyContext string, sessionLimit int64, longRunningThresholdSeconds int64) *PostgresqlSessionsPayload {
	this := PostgresqlSessionsPayload{}
	this.Stats = stats
	this.Sessions = sessions
	this.CopyContext = copyContext
	this.SessionLimit = sessionLimit
	this.LongRunningThresholdSeconds = longRunningThresholdSeconds
	return &this
}

// NewPostgresqlSessionsPayloadWithDefaults instantiates a new PostgresqlSessionsPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSessionsPayloadWithDefaults() *PostgresqlSessionsPayload {
	this := PostgresqlSessionsPayload{}
	return &this
}

// GetStats returns the Stats field value.
func (o *PostgresqlSessionsPayload) GetStats() PostgresqlSessionsStats {
	if o == nil {
		var ret PostgresqlSessionsStats
		return ret
	}
	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsPayload) GetStatsOk() (*PostgresqlSessionsStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value.
func (o *PostgresqlSessionsPayload) SetStats(v PostgresqlSessionsStats) {
	o.Stats = v
}

// GetSessions returns the Sessions field value.
func (o *PostgresqlSessionsPayload) GetSessions() []PostgresqlSession {
	if o == nil {
		var ret []PostgresqlSession
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsPayload) GetSessionsOk() (*[]PostgresqlSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// SetSessions sets field value.
func (o *PostgresqlSessionsPayload) SetSessions(v []PostgresqlSession) {
	o.Sessions = v
}

// GetCopyContext returns the CopyContext field value.
func (o *PostgresqlSessionsPayload) GetCopyContext() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CopyContext
}

// GetCopyContextOk returns a tuple with the CopyContext field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsPayload) GetCopyContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyContext, true
}

// SetCopyContext sets field value.
func (o *PostgresqlSessionsPayload) SetCopyContext(v string) {
	o.CopyContext = v
}

// GetSessionLimit returns the SessionLimit field value.
func (o *PostgresqlSessionsPayload) GetSessionLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionLimit
}

// GetSessionLimitOk returns a tuple with the SessionLimit field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsPayload) GetSessionLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionLimit, true
}

// SetSessionLimit sets field value.
func (o *PostgresqlSessionsPayload) SetSessionLimit(v int64) {
	o.SessionLimit = v
}

// GetLongRunningThresholdSeconds returns the LongRunningThresholdSeconds field value.
func (o *PostgresqlSessionsPayload) GetLongRunningThresholdSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LongRunningThresholdSeconds
}

// GetLongRunningThresholdSecondsOk returns a tuple with the LongRunningThresholdSeconds field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsPayload) GetLongRunningThresholdSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LongRunningThresholdSeconds, true
}

// SetLongRunningThresholdSeconds sets field value.
func (o *PostgresqlSessionsPayload) SetLongRunningThresholdSeconds(v int64) {
	o.LongRunningThresholdSeconds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSessionsPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["stats"] = o.Stats
	toSerialize["sessions"] = o.Sessions
	toSerialize["copyContext"] = o.CopyContext
	toSerialize["sessionLimit"] = o.SessionLimit
	toSerialize["longRunningThresholdSeconds"] = o.LongRunningThresholdSeconds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSessionsPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Stats                       *PostgresqlSessionsStats `json:"stats"`
		Sessions                    *[]PostgresqlSession     `json:"sessions"`
		CopyContext                 *string                  `json:"copyContext"`
		SessionLimit                *int64                   `json:"sessionLimit"`
		LongRunningThresholdSeconds *int64                   `json:"longRunningThresholdSeconds"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Stats == nil {
		return fmt.Errorf("required field stats missing")
	}
	if all.Sessions == nil {
		return fmt.Errorf("required field sessions missing")
	}
	if all.CopyContext == nil {
		return fmt.Errorf("required field copyContext missing")
	}
	if all.SessionLimit == nil {
		return fmt.Errorf("required field sessionLimit missing")
	}
	if all.LongRunningThresholdSeconds == nil {
		return fmt.Errorf("required field longRunningThresholdSeconds missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"stats", "sessions", "copyContext", "sessionLimit", "longRunningThresholdSeconds"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Stats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Stats = *all.Stats
	o.Sessions = *all.Sessions
	o.CopyContext = *all.CopyContext
	o.SessionLimit = *all.SessionLimit
	o.LongRunningThresholdSeconds = *all.LongRunningThresholdSeconds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
