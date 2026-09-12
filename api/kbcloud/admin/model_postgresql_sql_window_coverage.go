// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// PostgresqlSQLWindowCoverage Collection-level completeness before database, user or SQL filtering. expectedCollections is estimated from the recorded cadence; zero means cadence is unavailable when no coverage exists.
type PostgresqlSQLWindowCoverage struct {
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
		ExpectedCollections *int64 `json:"expectedCollections,omitempty"`
		ObservedCollections *int64 `json:"observedCollections,omitempty"`
		UsableCollections   *int64 `json:"usableCollections,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"expectedCollections", "observedCollections", "usableCollections"})
	} else {
		return err
	}
	o.ExpectedCollections = all.ExpectedCollections
	o.ObservedCollections = all.ObservedCollections
	o.UsableCollections = all.UsableCollections

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
