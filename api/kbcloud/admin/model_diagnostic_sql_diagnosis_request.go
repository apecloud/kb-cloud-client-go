// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSQLDiagnosisRequest Request a PostgreSQL read-only SQL diagnosis. The backend accepts only safe SELECT/EXPLAIN input and never runs the original SQL.
type DiagnosticSQLDiagnosisRequest struct {
	// PostgreSQL database name used for EXPLAIN.
	Database *string `json:"database,omitempty"`
	// Concrete SELECT SQL, parameterized SELECT SQL, or safe EXPLAIN SELECT SQL. Parameterized SQL such as $1 requires matching parameters before EXPLAIN.
	Sql string `json:"sql"`
	// Parameter values for PostgreSQL placeholders such as $1 and $2. Values are bound by ordinal position and are only returned to Console in redacted summaries.
	Parameters []DiagnosticSQLDiagnosisParameter `json:"parameters,omitempty"`
	// Optional source context for a SQL diagnosis request.
	Context *DiagnosticSQLDiagnosisContext `json:"context,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSQLDiagnosisRequest instantiates a new DiagnosticSQLDiagnosisRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSQLDiagnosisRequest(sql string) *DiagnosticSQLDiagnosisRequest {
	this := DiagnosticSQLDiagnosisRequest{}
	this.Sql = sql
	return &this
}

// NewDiagnosticSQLDiagnosisRequestWithDefaults instantiates a new DiagnosticSQLDiagnosisRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSQLDiagnosisRequestWithDefaults() *DiagnosticSQLDiagnosisRequest {
	this := DiagnosticSQLDiagnosisRequest{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisRequest) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisRequest) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisRequest) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *DiagnosticSQLDiagnosisRequest) SetDatabase(v string) {
	o.Database = &v
}

// GetSql returns the Sql field value.
func (o *DiagnosticSQLDiagnosisRequest) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisRequest) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value.
func (o *DiagnosticSQLDiagnosisRequest) SetSql(v string) {
	o.Sql = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisRequest) GetParameters() []DiagnosticSQLDiagnosisParameter {
	if o == nil || o.Parameters == nil {
		var ret []DiagnosticSQLDiagnosisParameter
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisRequest) GetParametersOk() (*[]DiagnosticSQLDiagnosisParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisRequest) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []DiagnosticSQLDiagnosisParameter and assigns it to the Parameters field.
func (o *DiagnosticSQLDiagnosisRequest) SetParameters(v []DiagnosticSQLDiagnosisParameter) {
	o.Parameters = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisRequest) GetContext() DiagnosticSQLDiagnosisContext {
	if o == nil || o.Context == nil {
		var ret DiagnosticSQLDiagnosisContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisRequest) GetContextOk() (*DiagnosticSQLDiagnosisContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisRequest) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given DiagnosticSQLDiagnosisContext and assigns it to the Context field.
func (o *DiagnosticSQLDiagnosisRequest) SetContext(v DiagnosticSQLDiagnosisContext) {
	o.Context = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSQLDiagnosisRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	toSerialize["sql"] = o.Sql
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSQLDiagnosisRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Database   *string                           `json:"database,omitempty"`
		Sql        *string                           `json:"sql"`
		Parameters []DiagnosticSQLDiagnosisParameter `json:"parameters,omitempty"`
		Context    *DiagnosticSQLDiagnosisContext    `json:"context,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Sql == nil {
		return fmt.Errorf("required field sql missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"database", "sql", "parameters", "context"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Database = all.Database
	o.Sql = *all.Sql
	o.Parameters = all.Parameters
	if all.Context != nil && all.Context.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Context = all.Context

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
