// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PostgresqlSQLWindowResources Optional sums of SQL resource deltas. Missing values are unknown, never zero. Block counters are PostgreSQL blocks, WAL bytes are bytes, and timing fields are milliseconds. Legacy blk timing is the older PostgreSQL block I/O timing; it is not interchangeable with the PostgreSQL 17 split fields. Counters at or above 2^53 are unavailable to avoid storage aggregation precision loss.
type PostgresqlSQLWindowResources struct {
	SharedBlksHit        common.NullableFloat64 `json:"sharedBlksHit,omitempty"`
	SharedBlksRead       common.NullableFloat64 `json:"sharedBlksRead,omitempty"`
	SharedBlksDirtied    common.NullableFloat64 `json:"sharedBlksDirtied,omitempty"`
	SharedBlksWritten    common.NullableFloat64 `json:"sharedBlksWritten,omitempty"`
	LocalBlksHit         common.NullableFloat64 `json:"localBlksHit,omitempty"`
	LocalBlksRead        common.NullableFloat64 `json:"localBlksRead,omitempty"`
	LocalBlksDirtied     common.NullableFloat64 `json:"localBlksDirtied,omitempty"`
	LocalBlksWritten     common.NullableFloat64 `json:"localBlksWritten,omitempty"`
	TempBlksRead         common.NullableFloat64 `json:"tempBlksRead,omitempty"`
	TempBlksWritten      common.NullableFloat64 `json:"tempBlksWritten,omitempty"`
	Plans                common.NullableFloat64 `json:"plans,omitempty"`
	PlanningTimeMs       common.NullableFloat64 `json:"planningTimeMs,omitempty"`
	WalRecords           common.NullableFloat64 `json:"walRecords,omitempty"`
	WalFpi               common.NullableFloat64 `json:"walFpi,omitempty"`
	WalBytes             common.NullableFloat64 `json:"walBytes,omitempty"`
	SharedBlkReadTimeMs  common.NullableFloat64 `json:"sharedBlkReadTimeMs,omitempty"`
	SharedBlkWriteTimeMs common.NullableFloat64 `json:"sharedBlkWriteTimeMs,omitempty"`
	LocalBlkReadTimeMs   common.NullableFloat64 `json:"localBlkReadTimeMs,omitempty"`
	LocalBlkWriteTimeMs  common.NullableFloat64 `json:"localBlkWriteTimeMs,omitempty"`
	TempBlkReadTimeMs    common.NullableFloat64 `json:"tempBlkReadTimeMs,omitempty"`
	TempBlkWriteTimeMs   common.NullableFloat64 `json:"tempBlkWriteTimeMs,omitempty"`
	BlkReadTimeMs        common.NullableFloat64 `json:"blkReadTimeMs,omitempty"`
	BlkWriteTimeMs       common.NullableFloat64 `json:"blkWriteTimeMs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowResources instantiates a new PostgresqlSQLWindowResources object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowResources() *PostgresqlSQLWindowResources {
	this := PostgresqlSQLWindowResources{}
	return &this
}

// NewPostgresqlSQLWindowResourcesWithDefaults instantiates a new PostgresqlSQLWindowResources object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowResourcesWithDefaults() *PostgresqlSQLWindowResources {
	this := PostgresqlSQLWindowResources{}
	return &this
}

// GetSharedBlksHit returns the SharedBlksHit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetSharedBlksHit() float64 {
	if o == nil || o.SharedBlksHit.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SharedBlksHit.Get()
}

// GetSharedBlksHitOk returns a tuple with the SharedBlksHit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetSharedBlksHitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedBlksHit.Get(), o.SharedBlksHit.IsSet()
}

// HasSharedBlksHit returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasSharedBlksHit() bool {
	return o != nil && o.SharedBlksHit.IsSet()
}

// SetSharedBlksHit gets a reference to the given common.NullableFloat64 and assigns it to the SharedBlksHit field.
func (o *PostgresqlSQLWindowResources) SetSharedBlksHit(v float64) {
	o.SharedBlksHit.Set(&v)
}

// SetSharedBlksHitNil sets the value for SharedBlksHit to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetSharedBlksHitNil() {
	o.SharedBlksHit.Set(nil)
}

// UnsetSharedBlksHit ensures that no value is present for SharedBlksHit, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetSharedBlksHit() {
	o.SharedBlksHit.Unset()
}

// GetSharedBlksRead returns the SharedBlksRead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetSharedBlksRead() float64 {
	if o == nil || o.SharedBlksRead.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SharedBlksRead.Get()
}

// GetSharedBlksReadOk returns a tuple with the SharedBlksRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetSharedBlksReadOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedBlksRead.Get(), o.SharedBlksRead.IsSet()
}

// HasSharedBlksRead returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasSharedBlksRead() bool {
	return o != nil && o.SharedBlksRead.IsSet()
}

// SetSharedBlksRead gets a reference to the given common.NullableFloat64 and assigns it to the SharedBlksRead field.
func (o *PostgresqlSQLWindowResources) SetSharedBlksRead(v float64) {
	o.SharedBlksRead.Set(&v)
}

// SetSharedBlksReadNil sets the value for SharedBlksRead to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetSharedBlksReadNil() {
	o.SharedBlksRead.Set(nil)
}

// UnsetSharedBlksRead ensures that no value is present for SharedBlksRead, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetSharedBlksRead() {
	o.SharedBlksRead.Unset()
}

// GetSharedBlksDirtied returns the SharedBlksDirtied field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetSharedBlksDirtied() float64 {
	if o == nil || o.SharedBlksDirtied.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SharedBlksDirtied.Get()
}

// GetSharedBlksDirtiedOk returns a tuple with the SharedBlksDirtied field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetSharedBlksDirtiedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedBlksDirtied.Get(), o.SharedBlksDirtied.IsSet()
}

// HasSharedBlksDirtied returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasSharedBlksDirtied() bool {
	return o != nil && o.SharedBlksDirtied.IsSet()
}

// SetSharedBlksDirtied gets a reference to the given common.NullableFloat64 and assigns it to the SharedBlksDirtied field.
func (o *PostgresqlSQLWindowResources) SetSharedBlksDirtied(v float64) {
	o.SharedBlksDirtied.Set(&v)
}

// SetSharedBlksDirtiedNil sets the value for SharedBlksDirtied to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetSharedBlksDirtiedNil() {
	o.SharedBlksDirtied.Set(nil)
}

// UnsetSharedBlksDirtied ensures that no value is present for SharedBlksDirtied, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetSharedBlksDirtied() {
	o.SharedBlksDirtied.Unset()
}

// GetSharedBlksWritten returns the SharedBlksWritten field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetSharedBlksWritten() float64 {
	if o == nil || o.SharedBlksWritten.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SharedBlksWritten.Get()
}

// GetSharedBlksWrittenOk returns a tuple with the SharedBlksWritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetSharedBlksWrittenOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedBlksWritten.Get(), o.SharedBlksWritten.IsSet()
}

// HasSharedBlksWritten returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasSharedBlksWritten() bool {
	return o != nil && o.SharedBlksWritten.IsSet()
}

// SetSharedBlksWritten gets a reference to the given common.NullableFloat64 and assigns it to the SharedBlksWritten field.
func (o *PostgresqlSQLWindowResources) SetSharedBlksWritten(v float64) {
	o.SharedBlksWritten.Set(&v)
}

// SetSharedBlksWrittenNil sets the value for SharedBlksWritten to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetSharedBlksWrittenNil() {
	o.SharedBlksWritten.Set(nil)
}

// UnsetSharedBlksWritten ensures that no value is present for SharedBlksWritten, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetSharedBlksWritten() {
	o.SharedBlksWritten.Unset()
}

// GetLocalBlksHit returns the LocalBlksHit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetLocalBlksHit() float64 {
	if o == nil || o.LocalBlksHit.Get() == nil {
		var ret float64
		return ret
	}
	return *o.LocalBlksHit.Get()
}

// GetLocalBlksHitOk returns a tuple with the LocalBlksHit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetLocalBlksHitOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalBlksHit.Get(), o.LocalBlksHit.IsSet()
}

// HasLocalBlksHit returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasLocalBlksHit() bool {
	return o != nil && o.LocalBlksHit.IsSet()
}

// SetLocalBlksHit gets a reference to the given common.NullableFloat64 and assigns it to the LocalBlksHit field.
func (o *PostgresqlSQLWindowResources) SetLocalBlksHit(v float64) {
	o.LocalBlksHit.Set(&v)
}

// SetLocalBlksHitNil sets the value for LocalBlksHit to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetLocalBlksHitNil() {
	o.LocalBlksHit.Set(nil)
}

// UnsetLocalBlksHit ensures that no value is present for LocalBlksHit, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetLocalBlksHit() {
	o.LocalBlksHit.Unset()
}

// GetLocalBlksRead returns the LocalBlksRead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetLocalBlksRead() float64 {
	if o == nil || o.LocalBlksRead.Get() == nil {
		var ret float64
		return ret
	}
	return *o.LocalBlksRead.Get()
}

// GetLocalBlksReadOk returns a tuple with the LocalBlksRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetLocalBlksReadOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalBlksRead.Get(), o.LocalBlksRead.IsSet()
}

// HasLocalBlksRead returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasLocalBlksRead() bool {
	return o != nil && o.LocalBlksRead.IsSet()
}

// SetLocalBlksRead gets a reference to the given common.NullableFloat64 and assigns it to the LocalBlksRead field.
func (o *PostgresqlSQLWindowResources) SetLocalBlksRead(v float64) {
	o.LocalBlksRead.Set(&v)
}

// SetLocalBlksReadNil sets the value for LocalBlksRead to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetLocalBlksReadNil() {
	o.LocalBlksRead.Set(nil)
}

// UnsetLocalBlksRead ensures that no value is present for LocalBlksRead, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetLocalBlksRead() {
	o.LocalBlksRead.Unset()
}

// GetLocalBlksDirtied returns the LocalBlksDirtied field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetLocalBlksDirtied() float64 {
	if o == nil || o.LocalBlksDirtied.Get() == nil {
		var ret float64
		return ret
	}
	return *o.LocalBlksDirtied.Get()
}

// GetLocalBlksDirtiedOk returns a tuple with the LocalBlksDirtied field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetLocalBlksDirtiedOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalBlksDirtied.Get(), o.LocalBlksDirtied.IsSet()
}

// HasLocalBlksDirtied returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasLocalBlksDirtied() bool {
	return o != nil && o.LocalBlksDirtied.IsSet()
}

// SetLocalBlksDirtied gets a reference to the given common.NullableFloat64 and assigns it to the LocalBlksDirtied field.
func (o *PostgresqlSQLWindowResources) SetLocalBlksDirtied(v float64) {
	o.LocalBlksDirtied.Set(&v)
}

// SetLocalBlksDirtiedNil sets the value for LocalBlksDirtied to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetLocalBlksDirtiedNil() {
	o.LocalBlksDirtied.Set(nil)
}

// UnsetLocalBlksDirtied ensures that no value is present for LocalBlksDirtied, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetLocalBlksDirtied() {
	o.LocalBlksDirtied.Unset()
}

// GetLocalBlksWritten returns the LocalBlksWritten field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetLocalBlksWritten() float64 {
	if o == nil || o.LocalBlksWritten.Get() == nil {
		var ret float64
		return ret
	}
	return *o.LocalBlksWritten.Get()
}

// GetLocalBlksWrittenOk returns a tuple with the LocalBlksWritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetLocalBlksWrittenOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalBlksWritten.Get(), o.LocalBlksWritten.IsSet()
}

// HasLocalBlksWritten returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasLocalBlksWritten() bool {
	return o != nil && o.LocalBlksWritten.IsSet()
}

// SetLocalBlksWritten gets a reference to the given common.NullableFloat64 and assigns it to the LocalBlksWritten field.
func (o *PostgresqlSQLWindowResources) SetLocalBlksWritten(v float64) {
	o.LocalBlksWritten.Set(&v)
}

// SetLocalBlksWrittenNil sets the value for LocalBlksWritten to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetLocalBlksWrittenNil() {
	o.LocalBlksWritten.Set(nil)
}

// UnsetLocalBlksWritten ensures that no value is present for LocalBlksWritten, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetLocalBlksWritten() {
	o.LocalBlksWritten.Unset()
}

// GetTempBlksRead returns the TempBlksRead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetTempBlksRead() float64 {
	if o == nil || o.TempBlksRead.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TempBlksRead.Get()
}

// GetTempBlksReadOk returns a tuple with the TempBlksRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetTempBlksReadOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TempBlksRead.Get(), o.TempBlksRead.IsSet()
}

// HasTempBlksRead returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasTempBlksRead() bool {
	return o != nil && o.TempBlksRead.IsSet()
}

// SetTempBlksRead gets a reference to the given common.NullableFloat64 and assigns it to the TempBlksRead field.
func (o *PostgresqlSQLWindowResources) SetTempBlksRead(v float64) {
	o.TempBlksRead.Set(&v)
}

// SetTempBlksReadNil sets the value for TempBlksRead to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetTempBlksReadNil() {
	o.TempBlksRead.Set(nil)
}

// UnsetTempBlksRead ensures that no value is present for TempBlksRead, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetTempBlksRead() {
	o.TempBlksRead.Unset()
}

// GetTempBlksWritten returns the TempBlksWritten field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetTempBlksWritten() float64 {
	if o == nil || o.TempBlksWritten.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TempBlksWritten.Get()
}

// GetTempBlksWrittenOk returns a tuple with the TempBlksWritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetTempBlksWrittenOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TempBlksWritten.Get(), o.TempBlksWritten.IsSet()
}

// HasTempBlksWritten returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasTempBlksWritten() bool {
	return o != nil && o.TempBlksWritten.IsSet()
}

// SetTempBlksWritten gets a reference to the given common.NullableFloat64 and assigns it to the TempBlksWritten field.
func (o *PostgresqlSQLWindowResources) SetTempBlksWritten(v float64) {
	o.TempBlksWritten.Set(&v)
}

// SetTempBlksWrittenNil sets the value for TempBlksWritten to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetTempBlksWrittenNil() {
	o.TempBlksWritten.Set(nil)
}

// UnsetTempBlksWritten ensures that no value is present for TempBlksWritten, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetTempBlksWritten() {
	o.TempBlksWritten.Unset()
}

// GetPlans returns the Plans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetPlans() float64 {
	if o == nil || o.Plans.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Plans.Get()
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetPlansOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Plans.Get(), o.Plans.IsSet()
}

// HasPlans returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasPlans() bool {
	return o != nil && o.Plans.IsSet()
}

// SetPlans gets a reference to the given common.NullableFloat64 and assigns it to the Plans field.
func (o *PostgresqlSQLWindowResources) SetPlans(v float64) {
	o.Plans.Set(&v)
}

// SetPlansNil sets the value for Plans to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetPlansNil() {
	o.Plans.Set(nil)
}

// UnsetPlans ensures that no value is present for Plans, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetPlans() {
	o.Plans.Unset()
}

// GetPlanningTimeMs returns the PlanningTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetPlanningTimeMs() float64 {
	if o == nil || o.PlanningTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.PlanningTimeMs.Get()
}

// GetPlanningTimeMsOk returns a tuple with the PlanningTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetPlanningTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanningTimeMs.Get(), o.PlanningTimeMs.IsSet()
}

// HasPlanningTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasPlanningTimeMs() bool {
	return o != nil && o.PlanningTimeMs.IsSet()
}

// SetPlanningTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the PlanningTimeMs field.
func (o *PostgresqlSQLWindowResources) SetPlanningTimeMs(v float64) {
	o.PlanningTimeMs.Set(&v)
}

// SetPlanningTimeMsNil sets the value for PlanningTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetPlanningTimeMsNil() {
	o.PlanningTimeMs.Set(nil)
}

// UnsetPlanningTimeMs ensures that no value is present for PlanningTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetPlanningTimeMs() {
	o.PlanningTimeMs.Unset()
}

// GetWalRecords returns the WalRecords field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetWalRecords() float64 {
	if o == nil || o.WalRecords.Get() == nil {
		var ret float64
		return ret
	}
	return *o.WalRecords.Get()
}

// GetWalRecordsOk returns a tuple with the WalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetWalRecordsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalRecords.Get(), o.WalRecords.IsSet()
}

// HasWalRecords returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasWalRecords() bool {
	return o != nil && o.WalRecords.IsSet()
}

// SetWalRecords gets a reference to the given common.NullableFloat64 and assigns it to the WalRecords field.
func (o *PostgresqlSQLWindowResources) SetWalRecords(v float64) {
	o.WalRecords.Set(&v)
}

// SetWalRecordsNil sets the value for WalRecords to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetWalRecordsNil() {
	o.WalRecords.Set(nil)
}

// UnsetWalRecords ensures that no value is present for WalRecords, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetWalRecords() {
	o.WalRecords.Unset()
}

// GetWalFpi returns the WalFpi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetWalFpi() float64 {
	if o == nil || o.WalFpi.Get() == nil {
		var ret float64
		return ret
	}
	return *o.WalFpi.Get()
}

// GetWalFpiOk returns a tuple with the WalFpi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetWalFpiOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalFpi.Get(), o.WalFpi.IsSet()
}

// HasWalFpi returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasWalFpi() bool {
	return o != nil && o.WalFpi.IsSet()
}

// SetWalFpi gets a reference to the given common.NullableFloat64 and assigns it to the WalFpi field.
func (o *PostgresqlSQLWindowResources) SetWalFpi(v float64) {
	o.WalFpi.Set(&v)
}

// SetWalFpiNil sets the value for WalFpi to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetWalFpiNil() {
	o.WalFpi.Set(nil)
}

// UnsetWalFpi ensures that no value is present for WalFpi, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetWalFpi() {
	o.WalFpi.Unset()
}

// GetWalBytes returns the WalBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetWalBytes() float64 {
	if o == nil || o.WalBytes.Get() == nil {
		var ret float64
		return ret
	}
	return *o.WalBytes.Get()
}

// GetWalBytesOk returns a tuple with the WalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetWalBytesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalBytes.Get(), o.WalBytes.IsSet()
}

// HasWalBytes returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasWalBytes() bool {
	return o != nil && o.WalBytes.IsSet()
}

// SetWalBytes gets a reference to the given common.NullableFloat64 and assigns it to the WalBytes field.
func (o *PostgresqlSQLWindowResources) SetWalBytes(v float64) {
	o.WalBytes.Set(&v)
}

// SetWalBytesNil sets the value for WalBytes to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetWalBytesNil() {
	o.WalBytes.Set(nil)
}

// UnsetWalBytes ensures that no value is present for WalBytes, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetWalBytes() {
	o.WalBytes.Unset()
}

// GetSharedBlkReadTimeMs returns the SharedBlkReadTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetSharedBlkReadTimeMs() float64 {
	if o == nil || o.SharedBlkReadTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SharedBlkReadTimeMs.Get()
}

// GetSharedBlkReadTimeMsOk returns a tuple with the SharedBlkReadTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetSharedBlkReadTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedBlkReadTimeMs.Get(), o.SharedBlkReadTimeMs.IsSet()
}

// HasSharedBlkReadTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasSharedBlkReadTimeMs() bool {
	return o != nil && o.SharedBlkReadTimeMs.IsSet()
}

// SetSharedBlkReadTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the SharedBlkReadTimeMs field.
func (o *PostgresqlSQLWindowResources) SetSharedBlkReadTimeMs(v float64) {
	o.SharedBlkReadTimeMs.Set(&v)
}

// SetSharedBlkReadTimeMsNil sets the value for SharedBlkReadTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetSharedBlkReadTimeMsNil() {
	o.SharedBlkReadTimeMs.Set(nil)
}

// UnsetSharedBlkReadTimeMs ensures that no value is present for SharedBlkReadTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetSharedBlkReadTimeMs() {
	o.SharedBlkReadTimeMs.Unset()
}

// GetSharedBlkWriteTimeMs returns the SharedBlkWriteTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetSharedBlkWriteTimeMs() float64 {
	if o == nil || o.SharedBlkWriteTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.SharedBlkWriteTimeMs.Get()
}

// GetSharedBlkWriteTimeMsOk returns a tuple with the SharedBlkWriteTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetSharedBlkWriteTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharedBlkWriteTimeMs.Get(), o.SharedBlkWriteTimeMs.IsSet()
}

// HasSharedBlkWriteTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasSharedBlkWriteTimeMs() bool {
	return o != nil && o.SharedBlkWriteTimeMs.IsSet()
}

// SetSharedBlkWriteTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the SharedBlkWriteTimeMs field.
func (o *PostgresqlSQLWindowResources) SetSharedBlkWriteTimeMs(v float64) {
	o.SharedBlkWriteTimeMs.Set(&v)
}

// SetSharedBlkWriteTimeMsNil sets the value for SharedBlkWriteTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetSharedBlkWriteTimeMsNil() {
	o.SharedBlkWriteTimeMs.Set(nil)
}

// UnsetSharedBlkWriteTimeMs ensures that no value is present for SharedBlkWriteTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetSharedBlkWriteTimeMs() {
	o.SharedBlkWriteTimeMs.Unset()
}

// GetLocalBlkReadTimeMs returns the LocalBlkReadTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetLocalBlkReadTimeMs() float64 {
	if o == nil || o.LocalBlkReadTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.LocalBlkReadTimeMs.Get()
}

// GetLocalBlkReadTimeMsOk returns a tuple with the LocalBlkReadTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetLocalBlkReadTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalBlkReadTimeMs.Get(), o.LocalBlkReadTimeMs.IsSet()
}

// HasLocalBlkReadTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasLocalBlkReadTimeMs() bool {
	return o != nil && o.LocalBlkReadTimeMs.IsSet()
}

// SetLocalBlkReadTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the LocalBlkReadTimeMs field.
func (o *PostgresqlSQLWindowResources) SetLocalBlkReadTimeMs(v float64) {
	o.LocalBlkReadTimeMs.Set(&v)
}

// SetLocalBlkReadTimeMsNil sets the value for LocalBlkReadTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetLocalBlkReadTimeMsNil() {
	o.LocalBlkReadTimeMs.Set(nil)
}

// UnsetLocalBlkReadTimeMs ensures that no value is present for LocalBlkReadTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetLocalBlkReadTimeMs() {
	o.LocalBlkReadTimeMs.Unset()
}

// GetLocalBlkWriteTimeMs returns the LocalBlkWriteTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetLocalBlkWriteTimeMs() float64 {
	if o == nil || o.LocalBlkWriteTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.LocalBlkWriteTimeMs.Get()
}

// GetLocalBlkWriteTimeMsOk returns a tuple with the LocalBlkWriteTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetLocalBlkWriteTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalBlkWriteTimeMs.Get(), o.LocalBlkWriteTimeMs.IsSet()
}

// HasLocalBlkWriteTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasLocalBlkWriteTimeMs() bool {
	return o != nil && o.LocalBlkWriteTimeMs.IsSet()
}

// SetLocalBlkWriteTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the LocalBlkWriteTimeMs field.
func (o *PostgresqlSQLWindowResources) SetLocalBlkWriteTimeMs(v float64) {
	o.LocalBlkWriteTimeMs.Set(&v)
}

// SetLocalBlkWriteTimeMsNil sets the value for LocalBlkWriteTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetLocalBlkWriteTimeMsNil() {
	o.LocalBlkWriteTimeMs.Set(nil)
}

// UnsetLocalBlkWriteTimeMs ensures that no value is present for LocalBlkWriteTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetLocalBlkWriteTimeMs() {
	o.LocalBlkWriteTimeMs.Unset()
}

// GetTempBlkReadTimeMs returns the TempBlkReadTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetTempBlkReadTimeMs() float64 {
	if o == nil || o.TempBlkReadTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TempBlkReadTimeMs.Get()
}

// GetTempBlkReadTimeMsOk returns a tuple with the TempBlkReadTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetTempBlkReadTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TempBlkReadTimeMs.Get(), o.TempBlkReadTimeMs.IsSet()
}

// HasTempBlkReadTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasTempBlkReadTimeMs() bool {
	return o != nil && o.TempBlkReadTimeMs.IsSet()
}

// SetTempBlkReadTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the TempBlkReadTimeMs field.
func (o *PostgresqlSQLWindowResources) SetTempBlkReadTimeMs(v float64) {
	o.TempBlkReadTimeMs.Set(&v)
}

// SetTempBlkReadTimeMsNil sets the value for TempBlkReadTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetTempBlkReadTimeMsNil() {
	o.TempBlkReadTimeMs.Set(nil)
}

// UnsetTempBlkReadTimeMs ensures that no value is present for TempBlkReadTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetTempBlkReadTimeMs() {
	o.TempBlkReadTimeMs.Unset()
}

// GetTempBlkWriteTimeMs returns the TempBlkWriteTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetTempBlkWriteTimeMs() float64 {
	if o == nil || o.TempBlkWriteTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TempBlkWriteTimeMs.Get()
}

// GetTempBlkWriteTimeMsOk returns a tuple with the TempBlkWriteTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetTempBlkWriteTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TempBlkWriteTimeMs.Get(), o.TempBlkWriteTimeMs.IsSet()
}

// HasTempBlkWriteTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasTempBlkWriteTimeMs() bool {
	return o != nil && o.TempBlkWriteTimeMs.IsSet()
}

// SetTempBlkWriteTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the TempBlkWriteTimeMs field.
func (o *PostgresqlSQLWindowResources) SetTempBlkWriteTimeMs(v float64) {
	o.TempBlkWriteTimeMs.Set(&v)
}

// SetTempBlkWriteTimeMsNil sets the value for TempBlkWriteTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetTempBlkWriteTimeMsNil() {
	o.TempBlkWriteTimeMs.Set(nil)
}

// UnsetTempBlkWriteTimeMs ensures that no value is present for TempBlkWriteTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetTempBlkWriteTimeMs() {
	o.TempBlkWriteTimeMs.Unset()
}

// GetBlkReadTimeMs returns the BlkReadTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetBlkReadTimeMs() float64 {
	if o == nil || o.BlkReadTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.BlkReadTimeMs.Get()
}

// GetBlkReadTimeMsOk returns a tuple with the BlkReadTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetBlkReadTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlkReadTimeMs.Get(), o.BlkReadTimeMs.IsSet()
}

// HasBlkReadTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasBlkReadTimeMs() bool {
	return o != nil && o.BlkReadTimeMs.IsSet()
}

// SetBlkReadTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the BlkReadTimeMs field.
func (o *PostgresqlSQLWindowResources) SetBlkReadTimeMs(v float64) {
	o.BlkReadTimeMs.Set(&v)
}

// SetBlkReadTimeMsNil sets the value for BlkReadTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetBlkReadTimeMsNil() {
	o.BlkReadTimeMs.Set(nil)
}

// UnsetBlkReadTimeMs ensures that no value is present for BlkReadTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetBlkReadTimeMs() {
	o.BlkReadTimeMs.Unset()
}

// GetBlkWriteTimeMs returns the BlkWriteTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowResources) GetBlkWriteTimeMs() float64 {
	if o == nil || o.BlkWriteTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.BlkWriteTimeMs.Get()
}

// GetBlkWriteTimeMsOk returns a tuple with the BlkWriteTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowResources) GetBlkWriteTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlkWriteTimeMs.Get(), o.BlkWriteTimeMs.IsSet()
}

// HasBlkWriteTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowResources) HasBlkWriteTimeMs() bool {
	return o != nil && o.BlkWriteTimeMs.IsSet()
}

// SetBlkWriteTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the BlkWriteTimeMs field.
func (o *PostgresqlSQLWindowResources) SetBlkWriteTimeMs(v float64) {
	o.BlkWriteTimeMs.Set(&v)
}

// SetBlkWriteTimeMsNil sets the value for BlkWriteTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowResources) SetBlkWriteTimeMsNil() {
	o.BlkWriteTimeMs.Set(nil)
}

// UnsetBlkWriteTimeMs ensures that no value is present for BlkWriteTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowResources) UnsetBlkWriteTimeMs() {
	o.BlkWriteTimeMs.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SharedBlksHit.IsSet() {
		toSerialize["sharedBlksHit"] = o.SharedBlksHit.Get()
	}
	if o.SharedBlksRead.IsSet() {
		toSerialize["sharedBlksRead"] = o.SharedBlksRead.Get()
	}
	if o.SharedBlksDirtied.IsSet() {
		toSerialize["sharedBlksDirtied"] = o.SharedBlksDirtied.Get()
	}
	if o.SharedBlksWritten.IsSet() {
		toSerialize["sharedBlksWritten"] = o.SharedBlksWritten.Get()
	}
	if o.LocalBlksHit.IsSet() {
		toSerialize["localBlksHit"] = o.LocalBlksHit.Get()
	}
	if o.LocalBlksRead.IsSet() {
		toSerialize["localBlksRead"] = o.LocalBlksRead.Get()
	}
	if o.LocalBlksDirtied.IsSet() {
		toSerialize["localBlksDirtied"] = o.LocalBlksDirtied.Get()
	}
	if o.LocalBlksWritten.IsSet() {
		toSerialize["localBlksWritten"] = o.LocalBlksWritten.Get()
	}
	if o.TempBlksRead.IsSet() {
		toSerialize["tempBlksRead"] = o.TempBlksRead.Get()
	}
	if o.TempBlksWritten.IsSet() {
		toSerialize["tempBlksWritten"] = o.TempBlksWritten.Get()
	}
	if o.Plans.IsSet() {
		toSerialize["plans"] = o.Plans.Get()
	}
	if o.PlanningTimeMs.IsSet() {
		toSerialize["planningTimeMs"] = o.PlanningTimeMs.Get()
	}
	if o.WalRecords.IsSet() {
		toSerialize["walRecords"] = o.WalRecords.Get()
	}
	if o.WalFpi.IsSet() {
		toSerialize["walFpi"] = o.WalFpi.Get()
	}
	if o.WalBytes.IsSet() {
		toSerialize["walBytes"] = o.WalBytes.Get()
	}
	if o.SharedBlkReadTimeMs.IsSet() {
		toSerialize["sharedBlkReadTimeMs"] = o.SharedBlkReadTimeMs.Get()
	}
	if o.SharedBlkWriteTimeMs.IsSet() {
		toSerialize["sharedBlkWriteTimeMs"] = o.SharedBlkWriteTimeMs.Get()
	}
	if o.LocalBlkReadTimeMs.IsSet() {
		toSerialize["localBlkReadTimeMs"] = o.LocalBlkReadTimeMs.Get()
	}
	if o.LocalBlkWriteTimeMs.IsSet() {
		toSerialize["localBlkWriteTimeMs"] = o.LocalBlkWriteTimeMs.Get()
	}
	if o.TempBlkReadTimeMs.IsSet() {
		toSerialize["tempBlkReadTimeMs"] = o.TempBlkReadTimeMs.Get()
	}
	if o.TempBlkWriteTimeMs.IsSet() {
		toSerialize["tempBlkWriteTimeMs"] = o.TempBlkWriteTimeMs.Get()
	}
	if o.BlkReadTimeMs.IsSet() {
		toSerialize["blkReadTimeMs"] = o.BlkReadTimeMs.Get()
	}
	if o.BlkWriteTimeMs.IsSet() {
		toSerialize["blkWriteTimeMs"] = o.BlkWriteTimeMs.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLWindowResources) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SharedBlksHit        common.NullableFloat64 `json:"sharedBlksHit,omitempty"`
		SharedBlksRead       common.NullableFloat64 `json:"sharedBlksRead,omitempty"`
		SharedBlksDirtied    common.NullableFloat64 `json:"sharedBlksDirtied,omitempty"`
		SharedBlksWritten    common.NullableFloat64 `json:"sharedBlksWritten,omitempty"`
		LocalBlksHit         common.NullableFloat64 `json:"localBlksHit,omitempty"`
		LocalBlksRead        common.NullableFloat64 `json:"localBlksRead,omitempty"`
		LocalBlksDirtied     common.NullableFloat64 `json:"localBlksDirtied,omitempty"`
		LocalBlksWritten     common.NullableFloat64 `json:"localBlksWritten,omitempty"`
		TempBlksRead         common.NullableFloat64 `json:"tempBlksRead,omitempty"`
		TempBlksWritten      common.NullableFloat64 `json:"tempBlksWritten,omitempty"`
		Plans                common.NullableFloat64 `json:"plans,omitempty"`
		PlanningTimeMs       common.NullableFloat64 `json:"planningTimeMs,omitempty"`
		WalRecords           common.NullableFloat64 `json:"walRecords,omitempty"`
		WalFpi               common.NullableFloat64 `json:"walFpi,omitempty"`
		WalBytes             common.NullableFloat64 `json:"walBytes,omitempty"`
		SharedBlkReadTimeMs  common.NullableFloat64 `json:"sharedBlkReadTimeMs,omitempty"`
		SharedBlkWriteTimeMs common.NullableFloat64 `json:"sharedBlkWriteTimeMs,omitempty"`
		LocalBlkReadTimeMs   common.NullableFloat64 `json:"localBlkReadTimeMs,omitempty"`
		LocalBlkWriteTimeMs  common.NullableFloat64 `json:"localBlkWriteTimeMs,omitempty"`
		TempBlkReadTimeMs    common.NullableFloat64 `json:"tempBlkReadTimeMs,omitempty"`
		TempBlkWriteTimeMs   common.NullableFloat64 `json:"tempBlkWriteTimeMs,omitempty"`
		BlkReadTimeMs        common.NullableFloat64 `json:"blkReadTimeMs,omitempty"`
		BlkWriteTimeMs       common.NullableFloat64 `json:"blkWriteTimeMs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sharedBlksHit", "sharedBlksRead", "sharedBlksDirtied", "sharedBlksWritten", "localBlksHit", "localBlksRead", "localBlksDirtied", "localBlksWritten", "tempBlksRead", "tempBlksWritten", "plans", "planningTimeMs", "walRecords", "walFpi", "walBytes", "sharedBlkReadTimeMs", "sharedBlkWriteTimeMs", "localBlkReadTimeMs", "localBlkWriteTimeMs", "tempBlkReadTimeMs", "tempBlkWriteTimeMs", "blkReadTimeMs", "blkWriteTimeMs"})
	} else {
		return err
	}
	o.SharedBlksHit = all.SharedBlksHit
	o.SharedBlksRead = all.SharedBlksRead
	o.SharedBlksDirtied = all.SharedBlksDirtied
	o.SharedBlksWritten = all.SharedBlksWritten
	o.LocalBlksHit = all.LocalBlksHit
	o.LocalBlksRead = all.LocalBlksRead
	o.LocalBlksDirtied = all.LocalBlksDirtied
	o.LocalBlksWritten = all.LocalBlksWritten
	o.TempBlksRead = all.TempBlksRead
	o.TempBlksWritten = all.TempBlksWritten
	o.Plans = all.Plans
	o.PlanningTimeMs = all.PlanningTimeMs
	o.WalRecords = all.WalRecords
	o.WalFpi = all.WalFpi
	o.WalBytes = all.WalBytes
	o.SharedBlkReadTimeMs = all.SharedBlkReadTimeMs
	o.SharedBlkWriteTimeMs = all.SharedBlkWriteTimeMs
	o.LocalBlkReadTimeMs = all.LocalBlkReadTimeMs
	o.LocalBlkWriteTimeMs = all.LocalBlkWriteTimeMs
	o.TempBlkReadTimeMs = all.TempBlkReadTimeMs
	o.TempBlkWriteTimeMs = all.TempBlkWriteTimeMs
	o.BlkReadTimeMs = all.BlkReadTimeMs
	o.BlkWriteTimeMs = all.BlkWriteTimeMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
