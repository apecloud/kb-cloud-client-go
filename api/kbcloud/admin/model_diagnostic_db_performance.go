// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformance First-screen database performance summary.
type DiagnosticDBPerformance struct {
	// One-sentence database key-metric conclusion for the first screen.
	Summary string `json:"summary"`
	// Metric key the user should inspect first. One of slowSql, longTransaction, lockWait, none, or collectionFailed.
	PrimaryMetric DiagnosticDBPerformancePrimaryMetric `json:"primaryMetric"`
	// One database performance metric with threshold and source quality.
	Cpu DiagnosticDBPerformanceMetric `json:"cpu"`
	// One database performance metric with threshold and source quality.
	Memory DiagnosticDBPerformanceMetric `json:"memory"`
	// One database performance metric with threshold and source quality.
	Disk DiagnosticDBPerformanceMetric `json:"disk"`
	// One database performance metric with threshold and source quality.
	Connections DiagnosticDBPerformanceMetric `json:"connections"`
	// Richer slow SQL performance metric. Inlines the generic metric fields and adds
	// Sprint 9 slow SQL specific result fields. Cross-field consistency between
	// slowSqlState, dataQuality, reasonCode and datasetSummary.totalCount is enforced
	// by backend loader Validate (see Sprint 9 v0.1 contract).
	//
	SlowSql DiagnosticDBPerformanceSlowSQLMetric `json:"slowSql"`
	// PostgreSQL long-transaction performance metric. It exposes a three-state
	// live collector contract (has data / no data / collection failed) plus the
	// source_missing compatibility state when collection is disabled or not
	// integrated. Full SQL text is never exposed; only digests are returned.
	//
	LongTransaction DiagnosticDBPerformanceLongTransactionMetric `json:"longTransaction"`
	// PostgreSQL lock-wait performance metric. It exposes a three-state live
	// collector contract (has data / no data / collection failed) plus the
	// source_missing compatibility state when collection is disabled or not
	// integrated. Full SQL text is never exposed; only digests are returned.
	//
	LockWait DiagnosticDBPerformanceLockWaitMetric `json:"lockWait"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformance instantiates a new DiagnosticDBPerformance object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformance(summary string, primaryMetric DiagnosticDBPerformancePrimaryMetric, cpu DiagnosticDBPerformanceMetric, memory DiagnosticDBPerformanceMetric, disk DiagnosticDBPerformanceMetric, connections DiagnosticDBPerformanceMetric, slowSql DiagnosticDBPerformanceSlowSQLMetric, longTransaction DiagnosticDBPerformanceLongTransactionMetric, lockWait DiagnosticDBPerformanceLockWaitMetric) *DiagnosticDBPerformance {
	this := DiagnosticDBPerformance{}
	this.Summary = summary
	this.PrimaryMetric = primaryMetric
	this.Cpu = cpu
	this.Memory = memory
	this.Disk = disk
	this.Connections = connections
	this.SlowSql = slowSql
	this.LongTransaction = longTransaction
	this.LockWait = lockWait
	return &this
}

// NewDiagnosticDBPerformanceWithDefaults instantiates a new DiagnosticDBPerformance object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceWithDefaults() *DiagnosticDBPerformance {
	this := DiagnosticDBPerformance{}
	return &this
}

// GetSummary returns the Summary field value.
func (o *DiagnosticDBPerformance) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformance) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *DiagnosticDBPerformance) SetSummary(v string) {
	o.Summary = v
}

// GetPrimaryMetric returns the PrimaryMetric field value.
func (o *DiagnosticDBPerformance) GetPrimaryMetric() DiagnosticDBPerformancePrimaryMetric {
	if o == nil {
		var ret DiagnosticDBPerformancePrimaryMetric
		return ret
	}
	return o.PrimaryMetric
}

// GetPrimaryMetricOk returns a tuple with the PrimaryMetric field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformance) GetPrimaryMetricOk() (*DiagnosticDBPerformancePrimaryMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryMetric, true
}

// SetPrimaryMetric sets field value.
func (o *DiagnosticDBPerformance) SetPrimaryMetric(v DiagnosticDBPerformancePrimaryMetric) {
	o.PrimaryMetric = v
}

// GetCpu returns the Cpu field value.
func (o *DiagnosticDBPerformance) GetCpu() DiagnosticDBPerformanceMetric {
	if o == nil {
		var ret DiagnosticDBPerformanceMetric
		return ret
	}
	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformance) GetCpuOk() (*DiagnosticDBPerformanceMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value.
func (o *DiagnosticDBPerformance) SetCpu(v DiagnosticDBPerformanceMetric) {
	o.Cpu = v
}

// GetMemory returns the Memory field value.
func (o *DiagnosticDBPerformance) GetMemory() DiagnosticDBPerformanceMetric {
	if o == nil {
		var ret DiagnosticDBPerformanceMetric
		return ret
	}
	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformance) GetMemoryOk() (*DiagnosticDBPerformanceMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value.
func (o *DiagnosticDBPerformance) SetMemory(v DiagnosticDBPerformanceMetric) {
	o.Memory = v
}

// GetDisk returns the Disk field value.
func (o *DiagnosticDBPerformance) GetDisk() DiagnosticDBPerformanceMetric {
	if o == nil {
		var ret DiagnosticDBPerformanceMetric
		return ret
	}
	return o.Disk
}

// GetDiskOk returns a tuple with the Disk field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformance) GetDiskOk() (*DiagnosticDBPerformanceMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disk, true
}

// SetDisk sets field value.
func (o *DiagnosticDBPerformance) SetDisk(v DiagnosticDBPerformanceMetric) {
	o.Disk = v
}

// GetConnections returns the Connections field value.
func (o *DiagnosticDBPerformance) GetConnections() DiagnosticDBPerformanceMetric {
	if o == nil {
		var ret DiagnosticDBPerformanceMetric
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformance) GetConnectionsOk() (*DiagnosticDBPerformanceMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value.
func (o *DiagnosticDBPerformance) SetConnections(v DiagnosticDBPerformanceMetric) {
	o.Connections = v
}

// GetSlowSql returns the SlowSql field value.
func (o *DiagnosticDBPerformance) GetSlowSql() DiagnosticDBPerformanceSlowSQLMetric {
	if o == nil {
		var ret DiagnosticDBPerformanceSlowSQLMetric
		return ret
	}
	return o.SlowSql
}

// GetSlowSqlOk returns a tuple with the SlowSql field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformance) GetSlowSqlOk() (*DiagnosticDBPerformanceSlowSQLMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowSql, true
}

// SetSlowSql sets field value.
func (o *DiagnosticDBPerformance) SetSlowSql(v DiagnosticDBPerformanceSlowSQLMetric) {
	o.SlowSql = v
}

// GetLongTransaction returns the LongTransaction field value.
func (o *DiagnosticDBPerformance) GetLongTransaction() DiagnosticDBPerformanceLongTransactionMetric {
	if o == nil {
		var ret DiagnosticDBPerformanceLongTransactionMetric
		return ret
	}
	return o.LongTransaction
}

// GetLongTransactionOk returns a tuple with the LongTransaction field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformance) GetLongTransactionOk() (*DiagnosticDBPerformanceLongTransactionMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LongTransaction, true
}

// SetLongTransaction sets field value.
func (o *DiagnosticDBPerformance) SetLongTransaction(v DiagnosticDBPerformanceLongTransactionMetric) {
	o.LongTransaction = v
}

// GetLockWait returns the LockWait field value.
func (o *DiagnosticDBPerformance) GetLockWait() DiagnosticDBPerformanceLockWaitMetric {
	if o == nil {
		var ret DiagnosticDBPerformanceLockWaitMetric
		return ret
	}
	return o.LockWait
}

// GetLockWaitOk returns a tuple with the LockWait field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformance) GetLockWaitOk() (*DiagnosticDBPerformanceLockWaitMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockWait, true
}

// SetLockWait sets field value.
func (o *DiagnosticDBPerformance) SetLockWait(v DiagnosticDBPerformanceLockWaitMetric) {
	o.LockWait = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["summary"] = o.Summary
	toSerialize["primaryMetric"] = o.PrimaryMetric
	toSerialize["cpu"] = o.Cpu
	toSerialize["memory"] = o.Memory
	toSerialize["disk"] = o.Disk
	toSerialize["connections"] = o.Connections
	toSerialize["slowSql"] = o.SlowSql
	toSerialize["longTransaction"] = o.LongTransaction
	toSerialize["lockWait"] = o.LockWait

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformance) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Summary         *string                                       `json:"summary"`
		PrimaryMetric   *DiagnosticDBPerformancePrimaryMetric         `json:"primaryMetric"`
		Cpu             *DiagnosticDBPerformanceMetric                `json:"cpu"`
		Memory          *DiagnosticDBPerformanceMetric                `json:"memory"`
		Disk            *DiagnosticDBPerformanceMetric                `json:"disk"`
		Connections     *DiagnosticDBPerformanceMetric                `json:"connections"`
		SlowSql         *DiagnosticDBPerformanceSlowSQLMetric         `json:"slowSql"`
		LongTransaction *DiagnosticDBPerformanceLongTransactionMetric `json:"longTransaction"`
		LockWait        *DiagnosticDBPerformanceLockWaitMetric        `json:"lockWait"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.PrimaryMetric == nil {
		return fmt.Errorf("required field primaryMetric missing")
	}
	if all.Cpu == nil {
		return fmt.Errorf("required field cpu missing")
	}
	if all.Memory == nil {
		return fmt.Errorf("required field memory missing")
	}
	if all.Disk == nil {
		return fmt.Errorf("required field disk missing")
	}
	if all.Connections == nil {
		return fmt.Errorf("required field connections missing")
	}
	if all.SlowSql == nil {
		return fmt.Errorf("required field slowSql missing")
	}
	if all.LongTransaction == nil {
		return fmt.Errorf("required field longTransaction missing")
	}
	if all.LockWait == nil {
		return fmt.Errorf("required field lockWait missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"summary", "primaryMetric", "cpu", "memory", "disk", "connections", "slowSql", "longTransaction", "lockWait"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Summary = *all.Summary
	if !all.PrimaryMetric.IsValid() {
		hasInvalidField = true
	} else {
		o.PrimaryMetric = *all.PrimaryMetric
	}
	if all.Cpu.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cpu = *all.Cpu
	if all.Memory.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Memory = *all.Memory
	if all.Disk.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Disk = *all.Disk
	if all.Connections.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Connections = *all.Connections
	if all.SlowSql.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SlowSql = *all.SlowSql
	if all.LongTransaction.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LongTransaction = *all.LongTransaction
	if all.LockWait.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LockWait = *all.LockWait

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
