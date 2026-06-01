// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSessionsStats struct {
	Total             int64 `json:"total"`
	Active            int64 `json:"active"`
	Idle              int64 `json:"idle"`
	IdleInTransaction int64 `json:"idleInTransaction"`
	Waiting           int64 `json:"waiting"`
	LongRunning       int64 `json:"longRunning"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSessionsStats instantiates a new PostgresqlSessionsStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSessionsStats(total int64, active int64, idle int64, idleInTransaction int64, waiting int64, longRunning int64) *PostgresqlSessionsStats {
	this := PostgresqlSessionsStats{}
	this.Total = total
	this.Active = active
	this.Idle = idle
	this.IdleInTransaction = idleInTransaction
	this.Waiting = waiting
	this.LongRunning = longRunning
	return &this
}

// NewPostgresqlSessionsStatsWithDefaults instantiates a new PostgresqlSessionsStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSessionsStatsWithDefaults() *PostgresqlSessionsStats {
	this := PostgresqlSessionsStats{}
	return &this
}

// GetTotal returns the Total field value.
func (o *PostgresqlSessionsStats) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsStats) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value.
func (o *PostgresqlSessionsStats) SetTotal(v int64) {
	o.Total = v
}

// GetActive returns the Active field value.
func (o *PostgresqlSessionsStats) GetActive() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsStats) GetActiveOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value.
func (o *PostgresqlSessionsStats) SetActive(v int64) {
	o.Active = v
}

// GetIdle returns the Idle field value.
func (o *PostgresqlSessionsStats) GetIdle() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Idle
}

// GetIdleOk returns a tuple with the Idle field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsStats) GetIdleOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Idle, true
}

// SetIdle sets field value.
func (o *PostgresqlSessionsStats) SetIdle(v int64) {
	o.Idle = v
}

// GetIdleInTransaction returns the IdleInTransaction field value.
func (o *PostgresqlSessionsStats) GetIdleInTransaction() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.IdleInTransaction
}

// GetIdleInTransactionOk returns a tuple with the IdleInTransaction field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsStats) GetIdleInTransactionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdleInTransaction, true
}

// SetIdleInTransaction sets field value.
func (o *PostgresqlSessionsStats) SetIdleInTransaction(v int64) {
	o.IdleInTransaction = v
}

// GetWaiting returns the Waiting field value.
func (o *PostgresqlSessionsStats) GetWaiting() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Waiting
}

// GetWaitingOk returns a tuple with the Waiting field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsStats) GetWaitingOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Waiting, true
}

// SetWaiting sets field value.
func (o *PostgresqlSessionsStats) SetWaiting(v int64) {
	o.Waiting = v
}

// GetLongRunning returns the LongRunning field value.
func (o *PostgresqlSessionsStats) GetLongRunning() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LongRunning
}

// GetLongRunningOk returns a tuple with the LongRunning field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSessionsStats) GetLongRunningOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LongRunning, true
}

// SetLongRunning sets field value.
func (o *PostgresqlSessionsStats) SetLongRunning(v int64) {
	o.LongRunning = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSessionsStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["total"] = o.Total
	toSerialize["active"] = o.Active
	toSerialize["idle"] = o.Idle
	toSerialize["idleInTransaction"] = o.IdleInTransaction
	toSerialize["waiting"] = o.Waiting
	toSerialize["longRunning"] = o.LongRunning

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSessionsStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Total             *int64 `json:"total"`
		Active            *int64 `json:"active"`
		Idle              *int64 `json:"idle"`
		IdleInTransaction *int64 `json:"idleInTransaction"`
		Waiting           *int64 `json:"waiting"`
		LongRunning       *int64 `json:"longRunning"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Total == nil {
		return fmt.Errorf("required field total missing")
	}
	if all.Active == nil {
		return fmt.Errorf("required field active missing")
	}
	if all.Idle == nil {
		return fmt.Errorf("required field idle missing")
	}
	if all.IdleInTransaction == nil {
		return fmt.Errorf("required field idleInTransaction missing")
	}
	if all.Waiting == nil {
		return fmt.Errorf("required field waiting missing")
	}
	if all.LongRunning == nil {
		return fmt.Errorf("required field longRunning missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"total", "active", "idle", "idleInTransaction", "waiting", "longRunning"})
	} else {
		return err
	}
	o.Total = *all.Total
	o.Active = *all.Active
	o.Idle = *all.Idle
	o.IdleInTransaction = *all.IdleInTransaction
	o.Waiting = *all.Waiting
	o.LongRunning = *all.LongRunning

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
