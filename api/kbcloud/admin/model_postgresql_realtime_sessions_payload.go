// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlRealtimeSessionsPayload struct {
	CollectedAt      time.Time                       `json:"collectedAt"`
	Stats            PostgresqlRealtimeSessionStats  `json:"stats"`
	Sessions         []PostgresqlRealtimeSession     `json:"sessions"`
	ReadonlyBoundary HealthDiagnosisReadonlyBoundary `json:"readonlyBoundary"`
	CopyPayload      HealthDiagnosisCopyPayload      `json:"copyPayload"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlRealtimeSessionsPayload instantiates a new PostgresqlRealtimeSessionsPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlRealtimeSessionsPayload(collectedAt time.Time, stats PostgresqlRealtimeSessionStats, sessions []PostgresqlRealtimeSession, readonlyBoundary HealthDiagnosisReadonlyBoundary, copyPayload HealthDiagnosisCopyPayload) *PostgresqlRealtimeSessionsPayload {
	this := PostgresqlRealtimeSessionsPayload{}
	this.CollectedAt = collectedAt
	this.Stats = stats
	this.Sessions = sessions
	this.ReadonlyBoundary = readonlyBoundary
	this.CopyPayload = copyPayload
	return &this
}

// NewPostgresqlRealtimeSessionsPayloadWithDefaults instantiates a new PostgresqlRealtimeSessionsPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlRealtimeSessionsPayloadWithDefaults() *PostgresqlRealtimeSessionsPayload {
	this := PostgresqlRealtimeSessionsPayload{}
	return &this
}

// GetCollectedAt returns the CollectedAt field value.
func (o *PostgresqlRealtimeSessionsPayload) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsPayload) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *PostgresqlRealtimeSessionsPayload) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// GetStats returns the Stats field value.
func (o *PostgresqlRealtimeSessionsPayload) GetStats() PostgresqlRealtimeSessionStats {
	if o == nil {
		var ret PostgresqlRealtimeSessionStats
		return ret
	}
	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsPayload) GetStatsOk() (*PostgresqlRealtimeSessionStats, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value.
func (o *PostgresqlRealtimeSessionsPayload) SetStats(v PostgresqlRealtimeSessionStats) {
	o.Stats = v
}

// GetSessions returns the Sessions field value.
func (o *PostgresqlRealtimeSessionsPayload) GetSessions() []PostgresqlRealtimeSession {
	if o == nil {
		var ret []PostgresqlRealtimeSession
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsPayload) GetSessionsOk() (*[]PostgresqlRealtimeSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// SetSessions sets field value.
func (o *PostgresqlRealtimeSessionsPayload) SetSessions(v []PostgresqlRealtimeSession) {
	o.Sessions = v
}

// GetReadonlyBoundary returns the ReadonlyBoundary field value.
func (o *PostgresqlRealtimeSessionsPayload) GetReadonlyBoundary() HealthDiagnosisReadonlyBoundary {
	if o == nil {
		var ret HealthDiagnosisReadonlyBoundary
		return ret
	}
	return o.ReadonlyBoundary
}

// GetReadonlyBoundaryOk returns a tuple with the ReadonlyBoundary field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsPayload) GetReadonlyBoundaryOk() (*HealthDiagnosisReadonlyBoundary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadonlyBoundary, true
}

// SetReadonlyBoundary sets field value.
func (o *PostgresqlRealtimeSessionsPayload) SetReadonlyBoundary(v HealthDiagnosisReadonlyBoundary) {
	o.ReadonlyBoundary = v
}

// GetCopyPayload returns the CopyPayload field value.
func (o *PostgresqlRealtimeSessionsPayload) GetCopyPayload() HealthDiagnosisCopyPayload {
	if o == nil {
		var ret HealthDiagnosisCopyPayload
		return ret
	}
	return o.CopyPayload
}

// GetCopyPayloadOk returns a tuple with the CopyPayload field value
// and a boolean to check if the value has been set.
func (o *PostgresqlRealtimeSessionsPayload) GetCopyPayloadOk() (*HealthDiagnosisCopyPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyPayload, true
}

// SetCopyPayload sets field value.
func (o *PostgresqlRealtimeSessionsPayload) SetCopyPayload(v HealthDiagnosisCopyPayload) {
	o.CopyPayload = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlRealtimeSessionsPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["stats"] = o.Stats
	toSerialize["sessions"] = o.Sessions
	toSerialize["readonlyBoundary"] = o.ReadonlyBoundary
	toSerialize["copyPayload"] = o.CopyPayload

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlRealtimeSessionsPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CollectedAt      *time.Time                       `json:"collectedAt"`
		Stats            *PostgresqlRealtimeSessionStats  `json:"stats"`
		Sessions         *[]PostgresqlRealtimeSession     `json:"sessions"`
		ReadonlyBoundary *HealthDiagnosisReadonlyBoundary `json:"readonlyBoundary"`
		CopyPayload      *HealthDiagnosisCopyPayload      `json:"copyPayload"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.Stats == nil {
		return fmt.Errorf("required field stats missing")
	}
	if all.Sessions == nil {
		return fmt.Errorf("required field sessions missing")
	}
	if all.ReadonlyBoundary == nil {
		return fmt.Errorf("required field readonlyBoundary missing")
	}
	if all.CopyPayload == nil {
		return fmt.Errorf("required field copyPayload missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"collectedAt", "stats", "sessions", "readonlyBoundary", "copyPayload"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CollectedAt = *all.CollectedAt
	if all.Stats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Stats = *all.Stats
	o.Sessions = *all.Sessions
	if all.ReadonlyBoundary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReadonlyBoundary = *all.ReadonlyBoundary
	if all.CopyPayload.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CopyPayload = *all.CopyPayload

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
