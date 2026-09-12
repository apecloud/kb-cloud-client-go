// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLWindowComparisonItem struct {
	SqlId *string `json:"sqlId,omitempty"`
	// Delta sums after delivery deduplication. Calls and rows at or above 2^53 are unavailable because the storage aggregation uses float64. Summary contains only collector-redacted SQL; meanTimeMs is null for zero calls.
	Current PostgresqlSQLWindowItem `json:"current"`
	// Delta sums after delivery deduplication. Calls and rows at or above 2^53 are unavailable because the storage aggregation uses float64. Summary contains only collector-redacted SQL; meanTimeMs is null for zero calls.
	Baseline               PostgresqlSQLWindowItem `json:"baseline"`
	TotalTimeChangePercent common.NullableFloat64  `json:"totalTimeChangePercent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowComparisonItem instantiates a new PostgresqlSQLWindowComparisonItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowComparisonItem(current PostgresqlSQLWindowItem, baseline PostgresqlSQLWindowItem) *PostgresqlSQLWindowComparisonItem {
	this := PostgresqlSQLWindowComparisonItem{}
	this.Current = current
	this.Baseline = baseline
	return &this
}

// NewPostgresqlSQLWindowComparisonItemWithDefaults instantiates a new PostgresqlSQLWindowComparisonItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowComparisonItemWithDefaults() *PostgresqlSQLWindowComparisonItem {
	this := PostgresqlSQLWindowComparisonItem{}
	return &this
}

// GetSqlId returns the SqlId field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowComparisonItem) GetSqlId() string {
	if o == nil || o.SqlId == nil {
		var ret string
		return ret
	}
	return *o.SqlId
}

// GetSqlIdOk returns a tuple with the SqlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowComparisonItem) GetSqlIdOk() (*string, bool) {
	if o == nil || o.SqlId == nil {
		return nil, false
	}
	return o.SqlId, true
}

// HasSqlId returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowComparisonItem) HasSqlId() bool {
	return o != nil && o.SqlId != nil
}

// SetSqlId gets a reference to the given string and assigns it to the SqlId field.
func (o *PostgresqlSQLWindowComparisonItem) SetSqlId(v string) {
	o.SqlId = &v
}

// GetCurrent returns the Current field value.
func (o *PostgresqlSQLWindowComparisonItem) GetCurrent() PostgresqlSQLWindowItem {
	if o == nil {
		var ret PostgresqlSQLWindowItem
		return ret
	}
	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowComparisonItem) GetCurrentOk() (*PostgresqlSQLWindowItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value.
func (o *PostgresqlSQLWindowComparisonItem) SetCurrent(v PostgresqlSQLWindowItem) {
	o.Current = v
}

// GetBaseline returns the Baseline field value.
func (o *PostgresqlSQLWindowComparisonItem) GetBaseline() PostgresqlSQLWindowItem {
	if o == nil {
		var ret PostgresqlSQLWindowItem
		return ret
	}
	return o.Baseline
}

// GetBaselineOk returns a tuple with the Baseline field value
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowComparisonItem) GetBaselineOk() (*PostgresqlSQLWindowItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Baseline, true
}

// SetBaseline sets field value.
func (o *PostgresqlSQLWindowComparisonItem) SetBaseline(v PostgresqlSQLWindowItem) {
	o.Baseline = v
}

// GetTotalTimeChangePercent returns the TotalTimeChangePercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowComparisonItem) GetTotalTimeChangePercent() float64 {
	if o == nil || o.TotalTimeChangePercent.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TotalTimeChangePercent.Get()
}

// GetTotalTimeChangePercentOk returns a tuple with the TotalTimeChangePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowComparisonItem) GetTotalTimeChangePercentOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalTimeChangePercent.Get(), o.TotalTimeChangePercent.IsSet()
}

// HasTotalTimeChangePercent returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowComparisonItem) HasTotalTimeChangePercent() bool {
	return o != nil && o.TotalTimeChangePercent.IsSet()
}

// SetTotalTimeChangePercent gets a reference to the given common.NullableFloat64 and assigns it to the TotalTimeChangePercent field.
func (o *PostgresqlSQLWindowComparisonItem) SetTotalTimeChangePercent(v float64) {
	o.TotalTimeChangePercent.Set(&v)
}

// SetTotalTimeChangePercentNil sets the value for TotalTimeChangePercent to be an explicit nil.
func (o *PostgresqlSQLWindowComparisonItem) SetTotalTimeChangePercentNil() {
	o.TotalTimeChangePercent.Set(nil)
}

// UnsetTotalTimeChangePercent ensures that no value is present for TotalTimeChangePercent, not even an explicit nil.
func (o *PostgresqlSQLWindowComparisonItem) UnsetTotalTimeChangePercent() {
	o.TotalTimeChangePercent.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowComparisonItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SqlId != nil {
		toSerialize["sqlId"] = o.SqlId
	}
	toSerialize["current"] = o.Current
	toSerialize["baseline"] = o.Baseline
	if o.TotalTimeChangePercent.IsSet() {
		toSerialize["totalTimeChangePercent"] = o.TotalTimeChangePercent.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLWindowComparisonItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SqlId                  *string                  `json:"sqlId,omitempty"`
		Current                *PostgresqlSQLWindowItem `json:"current"`
		Baseline               *PostgresqlSQLWindowItem `json:"baseline"`
		TotalTimeChangePercent common.NullableFloat64   `json:"totalTimeChangePercent,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Current == nil {
		return fmt.Errorf("required field current missing")
	}
	if all.Baseline == nil {
		return fmt.Errorf("required field baseline missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sqlId", "current", "baseline", "totalTimeChangePercent"})
	} else {
		return err
	}

	hasInvalidField := false
	o.SqlId = all.SqlId
	if all.Current.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Current = *all.Current
	if all.Baseline.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Baseline = *all.Baseline
	o.TotalTimeChangePercent = all.TotalTimeChangePercent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
