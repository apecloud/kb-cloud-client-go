// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysisSQLPressureEvidence SQL pressure evidence using digest/summary only; no raw SQL text.
type DiagnosticCPUAnalysisSQLPressureEvidence struct {
	// Whether at least one active SQL pressure sample was observed.
	Observed bool `json:"observed"`
	// Active client query samples represented in topQueryDigests.
	ActiveQueryCount int64 `json:"activeQueryCount"`
	// Digest-based active query samples. These are not raw SQL texts.
	TopQueryDigests []string `json:"topQueryDigests"`
	// Redacted reason when SQL pressure evidence is unavailable.
	MissingReason string `json:"missingReason"`
	// Collection timestamp.
	CollectedAt time.Time `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticCPUAnalysisSQLPressureEvidence instantiates a new DiagnosticCPUAnalysisSQLPressureEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticCPUAnalysisSQLPressureEvidence(observed bool, activeQueryCount int64, topQueryDigests []string, missingReason string, collectedAt time.Time) *DiagnosticCPUAnalysisSQLPressureEvidence {
	this := DiagnosticCPUAnalysisSQLPressureEvidence{}
	this.Observed = observed
	this.ActiveQueryCount = activeQueryCount
	this.TopQueryDigests = topQueryDigests
	this.MissingReason = missingReason
	this.CollectedAt = collectedAt
	return &this
}

// NewDiagnosticCPUAnalysisSQLPressureEvidenceWithDefaults instantiates a new DiagnosticCPUAnalysisSQLPressureEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticCPUAnalysisSQLPressureEvidenceWithDefaults() *DiagnosticCPUAnalysisSQLPressureEvidence {
	this := DiagnosticCPUAnalysisSQLPressureEvidence{}
	return &this
}

// GetObserved returns the Observed field value.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) GetObserved() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Observed
}

// GetObservedOk returns a tuple with the Observed field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) GetObservedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Observed, true
}

// SetObserved sets field value.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) SetObserved(v bool) {
	o.Observed = v
}

// GetActiveQueryCount returns the ActiveQueryCount field value.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) GetActiveQueryCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ActiveQueryCount
}

// GetActiveQueryCountOk returns a tuple with the ActiveQueryCount field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) GetActiveQueryCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveQueryCount, true
}

// SetActiveQueryCount sets field value.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) SetActiveQueryCount(v int64) {
	o.ActiveQueryCount = v
}

// GetTopQueryDigests returns the TopQueryDigests field value.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) GetTopQueryDigests() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TopQueryDigests
}

// GetTopQueryDigestsOk returns a tuple with the TopQueryDigests field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) GetTopQueryDigestsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopQueryDigests, true
}

// SetTopQueryDigests sets field value.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) SetTopQueryDigests(v []string) {
	o.TopQueryDigests = v
}

// GetMissingReason returns the MissingReason field value.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) GetMissingReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MissingReason
}

// GetMissingReasonOk returns a tuple with the MissingReason field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) GetMissingReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MissingReason, true
}

// SetMissingReason sets field value.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) SetMissingReason(v string) {
	o.MissingReason = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticCPUAnalysisSQLPressureEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["observed"] = o.Observed
	toSerialize["activeQueryCount"] = o.ActiveQueryCount
	toSerialize["topQueryDigests"] = o.TopQueryDigests
	toSerialize["missingReason"] = o.MissingReason
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticCPUAnalysisSQLPressureEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Observed         *bool      `json:"observed"`
		ActiveQueryCount *int64     `json:"activeQueryCount"`
		TopQueryDigests  *[]string  `json:"topQueryDigests"`
		MissingReason    *string    `json:"missingReason"`
		CollectedAt      *time.Time `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Observed == nil {
		return fmt.Errorf("required field observed missing")
	}
	if all.ActiveQueryCount == nil {
		return fmt.Errorf("required field activeQueryCount missing")
	}
	if all.TopQueryDigests == nil {
		return fmt.Errorf("required field topQueryDigests missing")
	}
	if all.MissingReason == nil {
		return fmt.Errorf("required field missingReason missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"observed", "activeQueryCount", "topQueryDigests", "missingReason", "collectedAt"})
	} else {
		return err
	}
	o.Observed = *all.Observed
	o.ActiveQueryCount = *all.ActiveQueryCount
	o.TopQueryDigests = *all.TopQueryDigests
	o.MissingReason = *all.MissingReason
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
