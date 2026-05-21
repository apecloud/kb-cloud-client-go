// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceLockWaitDatasetSummary Summary of the lock-wait sample set.
type DiagnosticDBPerformanceLockWaitDatasetSummary struct {
	// Number of blocked sessions found by the readonly lock-wait query.
	TotalCount int64 `json:"totalCount"`
	// Top blocked sessions sorted by wait duration.
	BlockedSessions []DiagnosticDBPerformanceLockWaitSample `json:"blockedSessions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceLockWaitDatasetSummary instantiates a new DiagnosticDBPerformanceLockWaitDatasetSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceLockWaitDatasetSummary(totalCount int64, blockedSessions []DiagnosticDBPerformanceLockWaitSample) *DiagnosticDBPerformanceLockWaitDatasetSummary {
	this := DiagnosticDBPerformanceLockWaitDatasetSummary{}
	this.TotalCount = totalCount
	this.BlockedSessions = blockedSessions
	return &this
}

// NewDiagnosticDBPerformanceLockWaitDatasetSummaryWithDefaults instantiates a new DiagnosticDBPerformanceLockWaitDatasetSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceLockWaitDatasetSummaryWithDefaults() *DiagnosticDBPerformanceLockWaitDatasetSummary {
	this := DiagnosticDBPerformanceLockWaitDatasetSummary{}
	return &this
}

// GetTotalCount returns the TotalCount field value.
func (o *DiagnosticDBPerformanceLockWaitDatasetSummary) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitDatasetSummary) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value.
func (o *DiagnosticDBPerformanceLockWaitDatasetSummary) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetBlockedSessions returns the BlockedSessions field value.
func (o *DiagnosticDBPerformanceLockWaitDatasetSummary) GetBlockedSessions() []DiagnosticDBPerformanceLockWaitSample {
	if o == nil {
		var ret []DiagnosticDBPerformanceLockWaitSample
		return ret
	}
	return o.BlockedSessions
}

// GetBlockedSessionsOk returns a tuple with the BlockedSessions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLockWaitDatasetSummary) GetBlockedSessionsOk() (*[]DiagnosticDBPerformanceLockWaitSample, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockedSessions, true
}

// SetBlockedSessions sets field value.
func (o *DiagnosticDBPerformanceLockWaitDatasetSummary) SetBlockedSessions(v []DiagnosticDBPerformanceLockWaitSample) {
	o.BlockedSessions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceLockWaitDatasetSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["totalCount"] = o.TotalCount
	toSerialize["blockedSessions"] = o.BlockedSessions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceLockWaitDatasetSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalCount      *int64                                   `json:"totalCount"`
		BlockedSessions *[]DiagnosticDBPerformanceLockWaitSample `json:"blockedSessions"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.TotalCount == nil {
		return fmt.Errorf("required field totalCount missing")
	}
	if all.BlockedSessions == nil {
		return fmt.Errorf("required field blockedSessions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"totalCount", "blockedSessions"})
	} else {
		return err
	}
	o.TotalCount = *all.TotalCount
	o.BlockedSessions = *all.BlockedSessions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
