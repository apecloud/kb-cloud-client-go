// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PostgresqlSQLWindowPoint Whole observation deltas are assigned by interval end to this display bucket. Effective boundaries describe the actual observation intervals and may begin before the display bucket, but never outside the overall requested window. Missing SQL observations remain null, including empty buckets. Window-level partial quality is conservatively inherited by populated points.
type PostgresqlSQLWindowPoint struct {
	Window *PostgresqlSQLWindow       `json:"window,omitempty"`
	Status *PostgresqlSQLWindowStatus `json:"status,omitempty"`
	// Optional sums of SQL resource deltas. Missing values are unknown, never zero. Block counters are PostgreSQL blocks, WAL bytes are bytes, and timing fields are milliseconds. Legacy blk timing is the older PostgreSQL block I/O timing; it is not interchangeable with the PostgreSQL 17 split fields. Counters at or above 2^53 are unavailable to avoid storage aggregation precision loss.
	Resources *PostgresqlSQLWindowResources `json:"resources,omitempty"`
	// Resource field name to unavailable reason (unsupported, disabled, configuration_changed, invalid_value, not_collected, or partial). A resource sum is present only when every contributing deduplicated SQL observation has a valid value. This does not override collection-level partial coverage.
	UnavailableMetrics map[string]string      `json:"unavailableMetrics,omitempty"`
	Calls              common.NullableInt64   `json:"calls,omitempty"`
	TotalTimeMs        common.NullableFloat64 `json:"totalTimeMs,omitempty"`
	Rows               common.NullableInt64   `json:"rows,omitempty"`
	Reasons            []string               `json:"reasons,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowPoint instantiates a new PostgresqlSQLWindowPoint object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowPoint() *PostgresqlSQLWindowPoint {
	this := PostgresqlSQLWindowPoint{}
	return &this
}

// NewPostgresqlSQLWindowPointWithDefaults instantiates a new PostgresqlSQLWindowPoint object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowPointWithDefaults() *PostgresqlSQLWindowPoint {
	this := PostgresqlSQLWindowPoint{}
	return &this
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowPoint) GetWindow() PostgresqlSQLWindow {
	if o == nil || o.Window == nil {
		var ret PostgresqlSQLWindow
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowPoint) GetWindowOk() (*PostgresqlSQLWindow, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowPoint) HasWindow() bool {
	return o != nil && o.Window != nil
}

// SetWindow gets a reference to the given PostgresqlSQLWindow and assigns it to the Window field.
func (o *PostgresqlSQLWindowPoint) SetWindow(v PostgresqlSQLWindow) {
	o.Window = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowPoint) GetStatus() PostgresqlSQLWindowStatus {
	if o == nil || o.Status == nil {
		var ret PostgresqlSQLWindowStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowPoint) GetStatusOk() (*PostgresqlSQLWindowStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowPoint) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given PostgresqlSQLWindowStatus and assigns it to the Status field.
func (o *PostgresqlSQLWindowPoint) SetStatus(v PostgresqlSQLWindowStatus) {
	o.Status = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowPoint) GetResources() PostgresqlSQLWindowResources {
	if o == nil || o.Resources == nil {
		var ret PostgresqlSQLWindowResources
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowPoint) GetResourcesOk() (*PostgresqlSQLWindowResources, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowPoint) HasResources() bool {
	return o != nil && o.Resources != nil
}

// SetResources gets a reference to the given PostgresqlSQLWindowResources and assigns it to the Resources field.
func (o *PostgresqlSQLWindowPoint) SetResources(v PostgresqlSQLWindowResources) {
	o.Resources = &v
}

// GetUnavailableMetrics returns the UnavailableMetrics field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowPoint) GetUnavailableMetrics() map[string]string {
	if o == nil || o.UnavailableMetrics == nil {
		var ret map[string]string
		return ret
	}
	return o.UnavailableMetrics
}

// GetUnavailableMetricsOk returns a tuple with the UnavailableMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowPoint) GetUnavailableMetricsOk() (*map[string]string, bool) {
	if o == nil || o.UnavailableMetrics == nil {
		return nil, false
	}
	return &o.UnavailableMetrics, true
}

// HasUnavailableMetrics returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowPoint) HasUnavailableMetrics() bool {
	return o != nil && o.UnavailableMetrics != nil
}

// SetUnavailableMetrics gets a reference to the given map[string]string and assigns it to the UnavailableMetrics field.
func (o *PostgresqlSQLWindowPoint) SetUnavailableMetrics(v map[string]string) {
	o.UnavailableMetrics = v
}

// GetCalls returns the Calls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowPoint) GetCalls() int64 {
	if o == nil || o.Calls.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Calls.Get()
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowPoint) GetCallsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Calls.Get(), o.Calls.IsSet()
}

// HasCalls returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowPoint) HasCalls() bool {
	return o != nil && o.Calls.IsSet()
}

// SetCalls gets a reference to the given common.NullableInt64 and assigns it to the Calls field.
func (o *PostgresqlSQLWindowPoint) SetCalls(v int64) {
	o.Calls.Set(&v)
}

// SetCallsNil sets the value for Calls to be an explicit nil.
func (o *PostgresqlSQLWindowPoint) SetCallsNil() {
	o.Calls.Set(nil)
}

// UnsetCalls ensures that no value is present for Calls, not even an explicit nil.
func (o *PostgresqlSQLWindowPoint) UnsetCalls() {
	o.Calls.Unset()
}

// GetTotalTimeMs returns the TotalTimeMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowPoint) GetTotalTimeMs() float64 {
	if o == nil || o.TotalTimeMs.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TotalTimeMs.Get()
}

// GetTotalTimeMsOk returns a tuple with the TotalTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowPoint) GetTotalTimeMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalTimeMs.Get(), o.TotalTimeMs.IsSet()
}

// HasTotalTimeMs returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowPoint) HasTotalTimeMs() bool {
	return o != nil && o.TotalTimeMs.IsSet()
}

// SetTotalTimeMs gets a reference to the given common.NullableFloat64 and assigns it to the TotalTimeMs field.
func (o *PostgresqlSQLWindowPoint) SetTotalTimeMs(v float64) {
	o.TotalTimeMs.Set(&v)
}

// SetTotalTimeMsNil sets the value for TotalTimeMs to be an explicit nil.
func (o *PostgresqlSQLWindowPoint) SetTotalTimeMsNil() {
	o.TotalTimeMs.Set(nil)
}

// UnsetTotalTimeMs ensures that no value is present for TotalTimeMs, not even an explicit nil.
func (o *PostgresqlSQLWindowPoint) UnsetTotalTimeMs() {
	o.TotalTimeMs.Unset()
}

// GetRows returns the Rows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlSQLWindowPoint) GetRows() int64 {
	if o == nil || o.Rows.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Rows.Get()
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *PostgresqlSQLWindowPoint) GetRowsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rows.Get(), o.Rows.IsSet()
}

// HasRows returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowPoint) HasRows() bool {
	return o != nil && o.Rows.IsSet()
}

// SetRows gets a reference to the given common.NullableInt64 and assigns it to the Rows field.
func (o *PostgresqlSQLWindowPoint) SetRows(v int64) {
	o.Rows.Set(&v)
}

// SetRowsNil sets the value for Rows to be an explicit nil.
func (o *PostgresqlSQLWindowPoint) SetRowsNil() {
	o.Rows.Set(nil)
}

// UnsetRows ensures that no value is present for Rows, not even an explicit nil.
func (o *PostgresqlSQLWindowPoint) UnsetRows() {
	o.Rows.Unset()
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowPoint) GetReasons() []string {
	if o == nil || o.Reasons == nil {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowPoint) GetReasonsOk() (*[]string, bool) {
	if o == nil || o.Reasons == nil {
		return nil, false
	}
	return &o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowPoint) HasReasons() bool {
	return o != nil && o.Reasons != nil
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *PostgresqlSQLWindowPoint) SetReasons(v []string) {
	o.Reasons = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowPoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Window != nil {
		toSerialize["window"] = o.Window
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.UnavailableMetrics != nil {
		toSerialize["unavailableMetrics"] = o.UnavailableMetrics
	}
	if o.Calls.IsSet() {
		toSerialize["calls"] = o.Calls.Get()
	}
	if o.TotalTimeMs.IsSet() {
		toSerialize["totalTimeMs"] = o.TotalTimeMs.Get()
	}
	if o.Rows.IsSet() {
		toSerialize["rows"] = o.Rows.Get()
	}
	if o.Reasons != nil {
		toSerialize["reasons"] = o.Reasons
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLWindowPoint) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Window             *PostgresqlSQLWindow          `json:"window,omitempty"`
		Status             *PostgresqlSQLWindowStatus    `json:"status,omitempty"`
		Resources          *PostgresqlSQLWindowResources `json:"resources,omitempty"`
		UnavailableMetrics map[string]string             `json:"unavailableMetrics,omitempty"`
		Calls              common.NullableInt64          `json:"calls,omitempty"`
		TotalTimeMs        common.NullableFloat64        `json:"totalTimeMs,omitempty"`
		Rows               common.NullableInt64          `json:"rows,omitempty"`
		Reasons            []string                      `json:"reasons,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"window", "status", "resources", "unavailableMetrics", "calls", "totalTimeMs", "rows", "reasons"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Window != nil && all.Window.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Window = all.Window
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	if all.Resources != nil && all.Resources.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Resources = all.Resources
	o.UnavailableMetrics = all.UnavailableMetrics
	o.Calls = all.Calls
	o.TotalTimeMs = all.TotalTimeMs
	o.Rows = all.Rows
	o.Reasons = all.Reasons

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
