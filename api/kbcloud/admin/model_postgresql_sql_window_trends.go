// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlSQLWindowTrends struct {
	Window  *PostgresqlSQLWindow       `json:"window,omitempty"`
	Status  *PostgresqlSQLWindowStatus `json:"status,omitempty"`
	Reasons []string                   `json:"reasons,omitempty"`
	// Collection-level completeness before database, user or SQL filtering. expectedCollections is estimated from the recorded cadence; zero means cadence is unavailable when no coverage exists.
	Coverage *PostgresqlSQLWindowCoverage `json:"coverage,omitempty"`
	Points   []PostgresqlSQLWindowPoint   `json:"points,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowTrends instantiates a new PostgresqlSQLWindowTrends object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowTrends() *PostgresqlSQLWindowTrends {
	this := PostgresqlSQLWindowTrends{}
	return &this
}

// NewPostgresqlSQLWindowTrendsWithDefaults instantiates a new PostgresqlSQLWindowTrends object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowTrendsWithDefaults() *PostgresqlSQLWindowTrends {
	this := PostgresqlSQLWindowTrends{}
	return &this
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowTrends) GetWindow() PostgresqlSQLWindow {
	if o == nil || o.Window == nil {
		var ret PostgresqlSQLWindow
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowTrends) GetWindowOk() (*PostgresqlSQLWindow, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowTrends) HasWindow() bool {
	return o != nil && o.Window != nil
}

// SetWindow gets a reference to the given PostgresqlSQLWindow and assigns it to the Window field.
func (o *PostgresqlSQLWindowTrends) SetWindow(v PostgresqlSQLWindow) {
	o.Window = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowTrends) GetStatus() PostgresqlSQLWindowStatus {
	if o == nil || o.Status == nil {
		var ret PostgresqlSQLWindowStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowTrends) GetStatusOk() (*PostgresqlSQLWindowStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowTrends) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given PostgresqlSQLWindowStatus and assigns it to the Status field.
func (o *PostgresqlSQLWindowTrends) SetStatus(v PostgresqlSQLWindowStatus) {
	o.Status = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowTrends) GetReasons() []string {
	if o == nil || o.Reasons == nil {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowTrends) GetReasonsOk() (*[]string, bool) {
	if o == nil || o.Reasons == nil {
		return nil, false
	}
	return &o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowTrends) HasReasons() bool {
	return o != nil && o.Reasons != nil
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *PostgresqlSQLWindowTrends) SetReasons(v []string) {
	o.Reasons = v
}

// GetCoverage returns the Coverage field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowTrends) GetCoverage() PostgresqlSQLWindowCoverage {
	if o == nil || o.Coverage == nil {
		var ret PostgresqlSQLWindowCoverage
		return ret
	}
	return *o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowTrends) GetCoverageOk() (*PostgresqlSQLWindowCoverage, bool) {
	if o == nil || o.Coverage == nil {
		return nil, false
	}
	return o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowTrends) HasCoverage() bool {
	return o != nil && o.Coverage != nil
}

// SetCoverage gets a reference to the given PostgresqlSQLWindowCoverage and assigns it to the Coverage field.
func (o *PostgresqlSQLWindowTrends) SetCoverage(v PostgresqlSQLWindowCoverage) {
	o.Coverage = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowTrends) GetPoints() []PostgresqlSQLWindowPoint {
	if o == nil || o.Points == nil {
		var ret []PostgresqlSQLWindowPoint
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowTrends) GetPointsOk() (*[]PostgresqlSQLWindowPoint, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return &o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowTrends) HasPoints() bool {
	return o != nil && o.Points != nil
}

// SetPoints gets a reference to the given []PostgresqlSQLWindowPoint and assigns it to the Points field.
func (o *PostgresqlSQLWindowTrends) SetPoints(v []PostgresqlSQLWindowPoint) {
	o.Points = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowTrends) MarshalJSON() ([]byte, error) {
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
	if o.Reasons != nil {
		toSerialize["reasons"] = o.Reasons
	}
	if o.Coverage != nil {
		toSerialize["coverage"] = o.Coverage
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLWindowTrends) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Window   *PostgresqlSQLWindow         `json:"window,omitempty"`
		Status   *PostgresqlSQLWindowStatus   `json:"status,omitempty"`
		Reasons  []string                     `json:"reasons,omitempty"`
		Coverage *PostgresqlSQLWindowCoverage `json:"coverage,omitempty"`
		Points   []PostgresqlSQLWindowPoint   `json:"points,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"window", "status", "reasons", "coverage", "points"})
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
	o.Reasons = all.Reasons
	if all.Coverage != nil && all.Coverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Coverage = all.Coverage
	o.Points = all.Points

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
