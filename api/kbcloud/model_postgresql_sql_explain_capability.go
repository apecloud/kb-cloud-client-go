// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLExplainCapability struct {
	// Capability status for explicit EXPLAIN. Expected values are not_requested, available, unavailable, and unsupported.
	Status string `json:"status"`
	// Reason when status is unavailable or unsupported, such as no_explainable_sql_sample or non_select.
	UnavailableReason *string `json:"unavailableReason,omitempty"`
	// Server-side sample source when available. Raw SQL is not returned.
	SampleSource *string `json:"sampleSource,omitempty"`
	// Sample collection timestamp when available.
	SampleCollectedAt *string `json:"sampleCollectedAt,omitempty"`
	// Redacted sample SQL summary when available. Full raw SQL is intentionally not returned.
	SampleSqlSummary *string `json:"sampleSQLSummary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLExplainCapability instantiates a new PostgresqlSQLExplainCapability object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLExplainCapability(status string) *PostgresqlSQLExplainCapability {
	this := PostgresqlSQLExplainCapability{}
	this.Status = status
	return &this
}

// NewPostgresqlSQLExplainCapabilityWithDefaults instantiates a new PostgresqlSQLExplainCapability object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLExplainCapabilityWithDefaults() *PostgresqlSQLExplainCapability {
	this := PostgresqlSQLExplainCapability{}
	return &this
}

// GetStatus returns the Status field value.
func (o *PostgresqlSQLExplainCapability) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLExplainCapability) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *PostgresqlSQLExplainCapability) SetStatus(v string) {
	o.Status = v
}

// GetUnavailableReason returns the UnavailableReason field value if set, zero value otherwise.
func (o *PostgresqlSQLExplainCapability) GetUnavailableReason() string {
	if o == nil || o.UnavailableReason == nil {
		var ret string
		return ret
	}
	return *o.UnavailableReason
}

// GetUnavailableReasonOk returns a tuple with the UnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLExplainCapability) GetUnavailableReasonOk() (*string, bool) {
	if o == nil || o.UnavailableReason == nil {
		return nil, false
	}
	return o.UnavailableReason, true
}

// HasUnavailableReason returns a boolean if a field has been set.
func (o *PostgresqlSQLExplainCapability) HasUnavailableReason() bool {
	return o != nil && o.UnavailableReason != nil
}

// SetUnavailableReason gets a reference to the given string and assigns it to the UnavailableReason field.
func (o *PostgresqlSQLExplainCapability) SetUnavailableReason(v string) {
	o.UnavailableReason = &v
}

// GetSampleSource returns the SampleSource field value if set, zero value otherwise.
func (o *PostgresqlSQLExplainCapability) GetSampleSource() string {
	if o == nil || o.SampleSource == nil {
		var ret string
		return ret
	}
	return *o.SampleSource
}

// GetSampleSourceOk returns a tuple with the SampleSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLExplainCapability) GetSampleSourceOk() (*string, bool) {
	if o == nil || o.SampleSource == nil {
		return nil, false
	}
	return o.SampleSource, true
}

// HasSampleSource returns a boolean if a field has been set.
func (o *PostgresqlSQLExplainCapability) HasSampleSource() bool {
	return o != nil && o.SampleSource != nil
}

// SetSampleSource gets a reference to the given string and assigns it to the SampleSource field.
func (o *PostgresqlSQLExplainCapability) SetSampleSource(v string) {
	o.SampleSource = &v
}

// GetSampleCollectedAt returns the SampleCollectedAt field value if set, zero value otherwise.
func (o *PostgresqlSQLExplainCapability) GetSampleCollectedAt() string {
	if o == nil || o.SampleCollectedAt == nil {
		var ret string
		return ret
	}
	return *o.SampleCollectedAt
}

// GetSampleCollectedAtOk returns a tuple with the SampleCollectedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLExplainCapability) GetSampleCollectedAtOk() (*string, bool) {
	if o == nil || o.SampleCollectedAt == nil {
		return nil, false
	}
	return o.SampleCollectedAt, true
}

// HasSampleCollectedAt returns a boolean if a field has been set.
func (o *PostgresqlSQLExplainCapability) HasSampleCollectedAt() bool {
	return o != nil && o.SampleCollectedAt != nil
}

// SetSampleCollectedAt gets a reference to the given string and assigns it to the SampleCollectedAt field.
func (o *PostgresqlSQLExplainCapability) SetSampleCollectedAt(v string) {
	o.SampleCollectedAt = &v
}

// GetSampleSqlSummary returns the SampleSqlSummary field value if set, zero value otherwise.
func (o *PostgresqlSQLExplainCapability) GetSampleSqlSummary() string {
	if o == nil || o.SampleSqlSummary == nil {
		var ret string
		return ret
	}
	return *o.SampleSqlSummary
}

// GetSampleSqlSummaryOk returns a tuple with the SampleSqlSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLExplainCapability) GetSampleSqlSummaryOk() (*string, bool) {
	if o == nil || o.SampleSqlSummary == nil {
		return nil, false
	}
	return o.SampleSqlSummary, true
}

// HasSampleSqlSummary returns a boolean if a field has been set.
func (o *PostgresqlSQLExplainCapability) HasSampleSqlSummary() bool {
	return o != nil && o.SampleSqlSummary != nil
}

// SetSampleSqlSummary gets a reference to the given string and assigns it to the SampleSqlSummary field.
func (o *PostgresqlSQLExplainCapability) SetSampleSqlSummary(v string) {
	o.SampleSqlSummary = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLExplainCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status
	if o.UnavailableReason != nil {
		toSerialize["unavailableReason"] = o.UnavailableReason
	}
	if o.SampleSource != nil {
		toSerialize["sampleSource"] = o.SampleSource
	}
	if o.SampleCollectedAt != nil {
		toSerialize["sampleCollectedAt"] = o.SampleCollectedAt
	}
	if o.SampleSqlSummary != nil {
		toSerialize["sampleSQLSummary"] = o.SampleSqlSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLExplainCapability) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status            *string `json:"status"`
		UnavailableReason *string `json:"unavailableReason,omitempty"`
		SampleSource      *string `json:"sampleSource,omitempty"`
		SampleCollectedAt *string `json:"sampleCollectedAt,omitempty"`
		SampleSqlSummary  *string `json:"sampleSQLSummary,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "unavailableReason", "sampleSource", "sampleCollectedAt", "sampleSQLSummary"})
	} else {
		return err
	}
	o.Status = *all.Status
	o.UnavailableReason = all.UnavailableReason
	o.SampleSource = all.SampleSource
	o.SampleCollectedAt = all.SampleCollectedAt
	o.SampleSqlSummary = all.SampleSqlSummary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
