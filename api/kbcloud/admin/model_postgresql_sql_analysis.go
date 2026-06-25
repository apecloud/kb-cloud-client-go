// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLAnalysis struct {
	// Data source. M1 uses pg_stat_statements only.
	Source string `json:"source"`
	// Source status. Expected values are available or unavailable.
	Status string `json:"status"`
	// Reason when status is unavailable, such as extension_disabled, permission_denied, or query_failed.
	UnavailableReason *string `json:"unavailableReason,omitempty"`
	// UTC timestamp reported by pg_stat_statements_info.stats_reset when available. Empty means the source did not expose this timestamp.
	StatsReset *string `json:"statsReset,omitempty"`
	// Backend collection timestamp in UTC. It is not a sampling-window end time.
	CollectedAt string `json:"collectedAt"`
	// Effective row limit after backend normalization.
	Limit int64 `json:"limit"`
	// Effective sort key. Expected values are totalTime, meanTime, maxTime, or calls.
	OrderBy string `json:"orderBy"`
	// SQL fingerprint ranking rows from pg_stat_statements. The list may be empty when no statement statistics exist.
	Items []PostgresqlSQLFingerprint `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLAnalysis instantiates a new PostgresqlSQLAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLAnalysis(source string, status string, collectedAt string, limit int64, orderBy string, items []PostgresqlSQLFingerprint) *PostgresqlSQLAnalysis {
	this := PostgresqlSQLAnalysis{}
	this.Source = source
	this.Status = status
	this.CollectedAt = collectedAt
	this.Limit = limit
	this.OrderBy = orderBy
	this.Items = items
	return &this
}

// NewPostgresqlSQLAnalysisWithDefaults instantiates a new PostgresqlSQLAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLAnalysisWithDefaults() *PostgresqlSQLAnalysis {
	this := PostgresqlSQLAnalysis{}
	return &this
}

// GetSource returns the Source field value.
func (o *PostgresqlSQLAnalysis) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLAnalysis) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *PostgresqlSQLAnalysis) SetSource(v string) {
	o.Source = v
}

// GetStatus returns the Status field value.
func (o *PostgresqlSQLAnalysis) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLAnalysis) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *PostgresqlSQLAnalysis) SetStatus(v string) {
	o.Status = v
}

// GetUnavailableReason returns the UnavailableReason field value if set, zero value otherwise.
func (o *PostgresqlSQLAnalysis) GetUnavailableReason() string {
	if o == nil || o.UnavailableReason == nil {
		var ret string
		return ret
	}
	return *o.UnavailableReason
}

// GetUnavailableReasonOk returns a tuple with the UnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLAnalysis) GetUnavailableReasonOk() (*string, bool) {
	if o == nil || o.UnavailableReason == nil {
		return nil, false
	}
	return o.UnavailableReason, true
}

// HasUnavailableReason returns a boolean if a field has been set.
func (o *PostgresqlSQLAnalysis) HasUnavailableReason() bool {
	return o != nil && o.UnavailableReason != nil
}

// SetUnavailableReason gets a reference to the given string and assigns it to the UnavailableReason field.
func (o *PostgresqlSQLAnalysis) SetUnavailableReason(v string) {
	o.UnavailableReason = &v
}

// GetStatsReset returns the StatsReset field value if set, zero value otherwise.
func (o *PostgresqlSQLAnalysis) GetStatsReset() string {
	if o == nil || o.StatsReset == nil {
		var ret string
		return ret
	}
	return *o.StatsReset
}

// GetStatsResetOk returns a tuple with the StatsReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLAnalysis) GetStatsResetOk() (*string, bool) {
	if o == nil || o.StatsReset == nil {
		return nil, false
	}
	return o.StatsReset, true
}

// HasStatsReset returns a boolean if a field has been set.
func (o *PostgresqlSQLAnalysis) HasStatsReset() bool {
	return o != nil && o.StatsReset != nil
}

// SetStatsReset gets a reference to the given string and assigns it to the StatsReset field.
func (o *PostgresqlSQLAnalysis) SetStatsReset(v string) {
	o.StatsReset = &v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *PostgresqlSQLAnalysis) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLAnalysis) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *PostgresqlSQLAnalysis) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// GetLimit returns the Limit field value.
func (o *PostgresqlSQLAnalysis) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLAnalysis) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *PostgresqlSQLAnalysis) SetLimit(v int64) {
	o.Limit = v
}

// GetOrderBy returns the OrderBy field value.
func (o *PostgresqlSQLAnalysis) GetOrderBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLAnalysis) GetOrderByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderBy, true
}

// SetOrderBy sets field value.
func (o *PostgresqlSQLAnalysis) SetOrderBy(v string) {
	o.OrderBy = v
}

// GetItems returns the Items field value.
func (o *PostgresqlSQLAnalysis) GetItems() []PostgresqlSQLFingerprint {
	if o == nil {
		var ret []PostgresqlSQLFingerprint
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLAnalysis) GetItemsOk() (*[]PostgresqlSQLFingerprint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *PostgresqlSQLAnalysis) SetItems(v []PostgresqlSQLFingerprint) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["source"] = o.Source
	toSerialize["status"] = o.Status
	if o.UnavailableReason != nil {
		toSerialize["unavailableReason"] = o.UnavailableReason
	}
	if o.StatsReset != nil {
		toSerialize["statsReset"] = o.StatsReset
	}
	toSerialize["collectedAt"] = o.CollectedAt
	toSerialize["limit"] = o.Limit
	toSerialize["orderBy"] = o.OrderBy
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source            *string                     `json:"source"`
		Status            *string                     `json:"status"`
		UnavailableReason *string                     `json:"unavailableReason,omitempty"`
		StatsReset        *string                     `json:"statsReset,omitempty"`
		CollectedAt       *string                     `json:"collectedAt"`
		Limit             *int64                      `json:"limit"`
		OrderBy           *string                     `json:"orderBy"`
		Items             *[]PostgresqlSQLFingerprint `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	if all.OrderBy == nil {
		return fmt.Errorf("required field orderBy missing")
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"source", "status", "unavailableReason", "statsReset", "collectedAt", "limit", "orderBy", "items"})
	} else {
		return err
	}
	o.Source = *all.Source
	o.Status = *all.Status
	o.UnavailableReason = all.UnavailableReason
	o.StatsReset = all.StatsReset
	o.CollectedAt = *all.CollectedAt
	o.Limit = *all.Limit
	o.OrderBy = *all.OrderBy
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
