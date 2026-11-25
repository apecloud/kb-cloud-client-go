// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// PostgresqlSQLWindowCoverage Collection-level completeness before database, user or SQL filtering. A complete window has continuous snapshot links and at least 95% time coverage. expectedCollections is the observed count plus estimated missing batches at detected breaks. It is for display only. Lookahead coverage is excluded from counts, time coverage and SQL totals.
type PostgresqlSQLWindowCoverage struct {
	// Digest of the verified delivery and coverage records for all requested windows.
	ObservationVersion *string `json:"observationVersion,omitempty"`
	// Percentage of requested time intersecting usable intervals.
	CoveragePercent *float64                    `json:"coveragePercent,omitempty"`
	Quality         *PostgresqlSQLWindowQuality `json:"quality,omitempty"`
	Gaps            []PostgresqlSQLWindowGap    `json:"gaps,omitempty"`
	Unchanged       *int64                      `json:"unchanged,omitempty"`
	Recreated       *int64                      `json:"recreated,omitempty"`
	// PostgreSQL major version at the latest observation.
	PgMajorVersion      *int64 `json:"pgMajorVersion,omitempty"`
	ExpectedCollections *int64 `json:"expectedCollections,omitempty"`
	ObservedCollections *int64 `json:"observedCollections,omitempty"`
	UsableCollections   *int64 `json:"usableCollections,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlSQLWindowCoverage instantiates a new PostgresqlSQLWindowCoverage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlSQLWindowCoverage() *PostgresqlSQLWindowCoverage {
	this := PostgresqlSQLWindowCoverage{}
	return &this
}

// NewPostgresqlSQLWindowCoverageWithDefaults instantiates a new PostgresqlSQLWindowCoverage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlSQLWindowCoverageWithDefaults() *PostgresqlSQLWindowCoverage {
	this := PostgresqlSQLWindowCoverage{}
	return &this
}

// GetObservationVersion returns the ObservationVersion field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowCoverage) GetObservationVersion() string {
	if o == nil || o.ObservationVersion == nil {
		var ret string
		return ret
	}
	return *o.ObservationVersion
}

// GetObservationVersionOk returns a tuple with the ObservationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowCoverage) GetObservationVersionOk() (*string, bool) {
	if o == nil || o.ObservationVersion == nil {
		return nil, false
	}
	return o.ObservationVersion, true
}

// HasObservationVersion returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowCoverage) HasObservationVersion() bool {
	return o != nil && o.ObservationVersion != nil
}

// SetObservationVersion gets a reference to the given string and assigns it to the ObservationVersion field.
func (o *PostgresqlSQLWindowCoverage) SetObservationVersion(v string) {
	o.ObservationVersion = &v
}

// GetCoveragePercent returns the CoveragePercent field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowCoverage) GetCoveragePercent() float64 {
	if o == nil || o.CoveragePercent == nil {
		var ret float64
		return ret
	}
	return *o.CoveragePercent
}

// GetCoveragePercentOk returns a tuple with the CoveragePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowCoverage) GetCoveragePercentOk() (*float64, bool) {
	if o == nil || o.CoveragePercent == nil {
		return nil, false
	}
	return o.CoveragePercent, true
}

// HasCoveragePercent returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowCoverage) HasCoveragePercent() bool {
	return o != nil && o.CoveragePercent != nil
}

// SetCoveragePercent gets a reference to the given float64 and assigns it to the CoveragePercent field.
func (o *PostgresqlSQLWindowCoverage) SetCoveragePercent(v float64) {
	o.CoveragePercent = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowCoverage) GetQuality() PostgresqlSQLWindowQuality {
	if o == nil || o.Quality == nil {
		var ret PostgresqlSQLWindowQuality
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowCoverage) GetQualityOk() (*PostgresqlSQLWindowQuality, bool) {
	if o == nil || o.Quality == nil {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowCoverage) HasQuality() bool {
	return o != nil && o.Quality != nil
}

// SetQuality gets a reference to the given PostgresqlSQLWindowQuality and assigns it to the Quality field.
func (o *PostgresqlSQLWindowCoverage) SetQuality(v PostgresqlSQLWindowQuality) {
	o.Quality = &v
}

// GetGaps returns the Gaps field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowCoverage) GetGaps() []PostgresqlSQLWindowGap {
	if o == nil || o.Gaps == nil {
		var ret []PostgresqlSQLWindowGap
		return ret
	}
	return o.Gaps
}

// GetGapsOk returns a tuple with the Gaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowCoverage) GetGapsOk() (*[]PostgresqlSQLWindowGap, bool) {
	if o == nil || o.Gaps == nil {
		return nil, false
	}
	return &o.Gaps, true
}

// HasGaps returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowCoverage) HasGaps() bool {
	return o != nil && o.Gaps != nil
}

// SetGaps gets a reference to the given []PostgresqlSQLWindowGap and assigns it to the Gaps field.
func (o *PostgresqlSQLWindowCoverage) SetGaps(v []PostgresqlSQLWindowGap) {
	o.Gaps = v
}

// GetUnchanged returns the Unchanged field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowCoverage) GetUnchanged() int64 {
	if o == nil || o.Unchanged == nil {
		var ret int64
		return ret
	}
	return *o.Unchanged
}

// GetUnchangedOk returns a tuple with the Unchanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowCoverage) GetUnchangedOk() (*int64, bool) {
	if o == nil || o.Unchanged == nil {
		return nil, false
	}
	return o.Unchanged, true
}

// HasUnchanged returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowCoverage) HasUnchanged() bool {
	return o != nil && o.Unchanged != nil
}

// SetUnchanged gets a reference to the given int64 and assigns it to the Unchanged field.
func (o *PostgresqlSQLWindowCoverage) SetUnchanged(v int64) {
	o.Unchanged = &v
}

// GetRecreated returns the Recreated field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowCoverage) GetRecreated() int64 {
	if o == nil || o.Recreated == nil {
		var ret int64
		return ret
	}
	return *o.Recreated
}

// GetRecreatedOk returns a tuple with the Recreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowCoverage) GetRecreatedOk() (*int64, bool) {
	if o == nil || o.Recreated == nil {
		return nil, false
	}
	return o.Recreated, true
}

// HasRecreated returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowCoverage) HasRecreated() bool {
	return o != nil && o.Recreated != nil
}

// SetRecreated gets a reference to the given int64 and assigns it to the Recreated field.
func (o *PostgresqlSQLWindowCoverage) SetRecreated(v int64) {
	o.Recreated = &v
}

// GetPgMajorVersion returns the PgMajorVersion field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowCoverage) GetPgMajorVersion() int64 {
	if o == nil || o.PgMajorVersion == nil {
		var ret int64
		return ret
	}
	return *o.PgMajorVersion
}

// GetPgMajorVersionOk returns a tuple with the PgMajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowCoverage) GetPgMajorVersionOk() (*int64, bool) {
	if o == nil || o.PgMajorVersion == nil {
		return nil, false
	}
	return o.PgMajorVersion, true
}

// HasPgMajorVersion returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowCoverage) HasPgMajorVersion() bool {
	return o != nil && o.PgMajorVersion != nil
}

// SetPgMajorVersion gets a reference to the given int64 and assigns it to the PgMajorVersion field.
func (o *PostgresqlSQLWindowCoverage) SetPgMajorVersion(v int64) {
	o.PgMajorVersion = &v
}

// GetExpectedCollections returns the ExpectedCollections field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowCoverage) GetExpectedCollections() int64 {
	if o == nil || o.ExpectedCollections == nil {
		var ret int64
		return ret
	}
	return *o.ExpectedCollections
}

// GetExpectedCollectionsOk returns a tuple with the ExpectedCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowCoverage) GetExpectedCollectionsOk() (*int64, bool) {
	if o == nil || o.ExpectedCollections == nil {
		return nil, false
	}
	return o.ExpectedCollections, true
}

// HasExpectedCollections returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowCoverage) HasExpectedCollections() bool {
	return o != nil && o.ExpectedCollections != nil
}

// SetExpectedCollections gets a reference to the given int64 and assigns it to the ExpectedCollections field.
func (o *PostgresqlSQLWindowCoverage) SetExpectedCollections(v int64) {
	o.ExpectedCollections = &v
}

// GetObservedCollections returns the ObservedCollections field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowCoverage) GetObservedCollections() int64 {
	if o == nil || o.ObservedCollections == nil {
		var ret int64
		return ret
	}
	return *o.ObservedCollections
}

// GetObservedCollectionsOk returns a tuple with the ObservedCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowCoverage) GetObservedCollectionsOk() (*int64, bool) {
	if o == nil || o.ObservedCollections == nil {
		return nil, false
	}
	return o.ObservedCollections, true
}

// HasObservedCollections returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowCoverage) HasObservedCollections() bool {
	return o != nil && o.ObservedCollections != nil
}

// SetObservedCollections gets a reference to the given int64 and assigns it to the ObservedCollections field.
func (o *PostgresqlSQLWindowCoverage) SetObservedCollections(v int64) {
	o.ObservedCollections = &v
}

// GetUsableCollections returns the UsableCollections field value if set, zero value otherwise.
func (o *PostgresqlSQLWindowCoverage) GetUsableCollections() int64 {
	if o == nil || o.UsableCollections == nil {
		var ret int64
		return ret
	}
	return *o.UsableCollections
}

// GetUsableCollectionsOk returns a tuple with the UsableCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlSQLWindowCoverage) GetUsableCollectionsOk() (*int64, bool) {
	if o == nil || o.UsableCollections == nil {
		return nil, false
	}
	return o.UsableCollections, true
}

// HasUsableCollections returns a boolean if a field has been set.
func (o *PostgresqlSQLWindowCoverage) HasUsableCollections() bool {
	return o != nil && o.UsableCollections != nil
}

// SetUsableCollections gets a reference to the given int64 and assigns it to the UsableCollections field.
func (o *PostgresqlSQLWindowCoverage) SetUsableCollections(v int64) {
	o.UsableCollections = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlSQLWindowCoverage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ObservationVersion != nil {
		toSerialize["observationVersion"] = o.ObservationVersion
	}
	if o.CoveragePercent != nil {
		toSerialize["coveragePercent"] = o.CoveragePercent
	}
	if o.Quality != nil {
		toSerialize["quality"] = o.Quality
	}
	if o.Gaps != nil {
		toSerialize["gaps"] = o.Gaps
	}
	if o.Unchanged != nil {
		toSerialize["unchanged"] = o.Unchanged
	}
	if o.Recreated != nil {
		toSerialize["recreated"] = o.Recreated
	}
	if o.PgMajorVersion != nil {
		toSerialize["pgMajorVersion"] = o.PgMajorVersion
	}
	if o.ExpectedCollections != nil {
		toSerialize["expectedCollections"] = o.ExpectedCollections
	}
	if o.ObservedCollections != nil {
		toSerialize["observedCollections"] = o.ObservedCollections
	}
	if o.UsableCollections != nil {
		toSerialize["usableCollections"] = o.UsableCollections
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlSQLWindowCoverage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ObservationVersion  *string                     `json:"observationVersion,omitempty"`
		CoveragePercent     *float64                    `json:"coveragePercent,omitempty"`
		Quality             *PostgresqlSQLWindowQuality `json:"quality,omitempty"`
		Gaps                []PostgresqlSQLWindowGap    `json:"gaps,omitempty"`
		Unchanged           *int64                      `json:"unchanged,omitempty"`
		Recreated           *int64                      `json:"recreated,omitempty"`
		PgMajorVersion      *int64                      `json:"pgMajorVersion,omitempty"`
		ExpectedCollections *int64                      `json:"expectedCollections,omitempty"`
		ObservedCollections *int64                      `json:"observedCollections,omitempty"`
		UsableCollections   *int64                      `json:"usableCollections,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"observationVersion", "coveragePercent", "quality", "gaps", "unchanged", "recreated", "pgMajorVersion", "expectedCollections", "observedCollections", "usableCollections"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ObservationVersion = all.ObservationVersion
	o.CoveragePercent = all.CoveragePercent
	if all.Quality != nil && !all.Quality.IsValid() {
		hasInvalidField = true
	} else {
		o.Quality = all.Quality
	}
	o.Gaps = all.Gaps
	o.Unchanged = all.Unchanged
	o.Recreated = all.Recreated
	o.PgMajorVersion = all.PgMajorVersion
	o.ExpectedCollections = all.ExpectedCollections
	o.ObservedCollections = all.ObservedCollections
	o.UsableCollections = all.UsableCollections

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
