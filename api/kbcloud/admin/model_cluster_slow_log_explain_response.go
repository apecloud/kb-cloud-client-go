// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterSlowLogExplainResponse Explain response for a selected slow log template sample
type ClusterSlowLogExplainResponse struct {
	// Query template identifier
	TemplateId *string `json:"templateId,omitempty"`
	// Concrete sample SQL used for EXPLAIN
	Sql *string `json:"sql,omitempty"`
	// Statement type accepted by the backend safety check
	SqlType *string `json:"sqlType,omitempty"`
	// Database selected from the slow log sample
	Database *string `json:"database,omitempty"`
	// Normalized SQL template from oteld
	NormalizedQuery *string `json:"normalizedQuery,omitempty"`
	// Timestamp of the selected sample in epoch nanoseconds
	SampleTimestamp *int64 `json:"sampleTimestamp,omitempty"`
	// Execution time of the selected sample in seconds
	ExecutionTime *float64 `json:"executionTime,omitempty"`
	// Cluster execution log item represents a single log entry
	Sample        *ClusterExecutionLogItem `json:"sample,omitempty"`
	ExplainResult *DmsQueryResponse        `json:"explainResult,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterSlowLogExplainResponse instantiates a new ClusterSlowLogExplainResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterSlowLogExplainResponse() *ClusterSlowLogExplainResponse {
	this := ClusterSlowLogExplainResponse{}
	return &this
}

// NewClusterSlowLogExplainResponseWithDefaults instantiates a new ClusterSlowLogExplainResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterSlowLogExplainResponseWithDefaults() *ClusterSlowLogExplainResponse {
	this := ClusterSlowLogExplainResponse{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *ClusterSlowLogExplainResponse) GetTemplateId() string {
	if o == nil || o.TemplateId == nil {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogExplainResponse) GetTemplateIdOk() (*string, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ClusterSlowLogExplainResponse) HasTemplateId() bool {
	return o != nil && o.TemplateId != nil
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *ClusterSlowLogExplainResponse) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *ClusterSlowLogExplainResponse) GetSql() string {
	if o == nil || o.Sql == nil {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogExplainResponse) GetSqlOk() (*string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *ClusterSlowLogExplainResponse) HasSql() bool {
	return o != nil && o.Sql != nil
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *ClusterSlowLogExplainResponse) SetSql(v string) {
	o.Sql = &v
}

// GetSqlType returns the SqlType field value if set, zero value otherwise.
func (o *ClusterSlowLogExplainResponse) GetSqlType() string {
	if o == nil || o.SqlType == nil {
		var ret string
		return ret
	}
	return *o.SqlType
}

// GetSqlTypeOk returns a tuple with the SqlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogExplainResponse) GetSqlTypeOk() (*string, bool) {
	if o == nil || o.SqlType == nil {
		return nil, false
	}
	return o.SqlType, true
}

// HasSqlType returns a boolean if a field has been set.
func (o *ClusterSlowLogExplainResponse) HasSqlType() bool {
	return o != nil && o.SqlType != nil
}

// SetSqlType gets a reference to the given string and assigns it to the SqlType field.
func (o *ClusterSlowLogExplainResponse) SetSqlType(v string) {
	o.SqlType = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ClusterSlowLogExplainResponse) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogExplainResponse) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ClusterSlowLogExplainResponse) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ClusterSlowLogExplainResponse) SetDatabase(v string) {
	o.Database = &v
}

// GetNormalizedQuery returns the NormalizedQuery field value if set, zero value otherwise.
func (o *ClusterSlowLogExplainResponse) GetNormalizedQuery() string {
	if o == nil || o.NormalizedQuery == nil {
		var ret string
		return ret
	}
	return *o.NormalizedQuery
}

// GetNormalizedQueryOk returns a tuple with the NormalizedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogExplainResponse) GetNormalizedQueryOk() (*string, bool) {
	if o == nil || o.NormalizedQuery == nil {
		return nil, false
	}
	return o.NormalizedQuery, true
}

// HasNormalizedQuery returns a boolean if a field has been set.
func (o *ClusterSlowLogExplainResponse) HasNormalizedQuery() bool {
	return o != nil && o.NormalizedQuery != nil
}

// SetNormalizedQuery gets a reference to the given string and assigns it to the NormalizedQuery field.
func (o *ClusterSlowLogExplainResponse) SetNormalizedQuery(v string) {
	o.NormalizedQuery = &v
}

// GetSampleTimestamp returns the SampleTimestamp field value if set, zero value otherwise.
func (o *ClusterSlowLogExplainResponse) GetSampleTimestamp() int64 {
	if o == nil || o.SampleTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.SampleTimestamp
}

// GetSampleTimestampOk returns a tuple with the SampleTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogExplainResponse) GetSampleTimestampOk() (*int64, bool) {
	if o == nil || o.SampleTimestamp == nil {
		return nil, false
	}
	return o.SampleTimestamp, true
}

// HasSampleTimestamp returns a boolean if a field has been set.
func (o *ClusterSlowLogExplainResponse) HasSampleTimestamp() bool {
	return o != nil && o.SampleTimestamp != nil
}

// SetSampleTimestamp gets a reference to the given int64 and assigns it to the SampleTimestamp field.
func (o *ClusterSlowLogExplainResponse) SetSampleTimestamp(v int64) {
	o.SampleTimestamp = &v
}

// GetExecutionTime returns the ExecutionTime field value if set, zero value otherwise.
func (o *ClusterSlowLogExplainResponse) GetExecutionTime() float64 {
	if o == nil || o.ExecutionTime == nil {
		var ret float64
		return ret
	}
	return *o.ExecutionTime
}

// GetExecutionTimeOk returns a tuple with the ExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogExplainResponse) GetExecutionTimeOk() (*float64, bool) {
	if o == nil || o.ExecutionTime == nil {
		return nil, false
	}
	return o.ExecutionTime, true
}

// HasExecutionTime returns a boolean if a field has been set.
func (o *ClusterSlowLogExplainResponse) HasExecutionTime() bool {
	return o != nil && o.ExecutionTime != nil
}

// SetExecutionTime gets a reference to the given float64 and assigns it to the ExecutionTime field.
func (o *ClusterSlowLogExplainResponse) SetExecutionTime(v float64) {
	o.ExecutionTime = &v
}

// GetSample returns the Sample field value if set, zero value otherwise.
func (o *ClusterSlowLogExplainResponse) GetSample() ClusterExecutionLogItem {
	if o == nil || o.Sample == nil {
		var ret ClusterExecutionLogItem
		return ret
	}
	return *o.Sample
}

// GetSampleOk returns a tuple with the Sample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogExplainResponse) GetSampleOk() (*ClusterExecutionLogItem, bool) {
	if o == nil || o.Sample == nil {
		return nil, false
	}
	return o.Sample, true
}

// HasSample returns a boolean if a field has been set.
func (o *ClusterSlowLogExplainResponse) HasSample() bool {
	return o != nil && o.Sample != nil
}

// SetSample gets a reference to the given ClusterExecutionLogItem and assigns it to the Sample field.
func (o *ClusterSlowLogExplainResponse) SetSample(v ClusterExecutionLogItem) {
	o.Sample = &v
}

// GetExplainResult returns the ExplainResult field value if set, zero value otherwise.
func (o *ClusterSlowLogExplainResponse) GetExplainResult() DmsQueryResponse {
	if o == nil || o.ExplainResult == nil {
		var ret DmsQueryResponse
		return ret
	}
	return *o.ExplainResult
}

// GetExplainResultOk returns a tuple with the ExplainResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogExplainResponse) GetExplainResultOk() (*DmsQueryResponse, bool) {
	if o == nil || o.ExplainResult == nil {
		return nil, false
	}
	return o.ExplainResult, true
}

// HasExplainResult returns a boolean if a field has been set.
func (o *ClusterSlowLogExplainResponse) HasExplainResult() bool {
	return o != nil && o.ExplainResult != nil
}

// SetExplainResult gets a reference to the given DmsQueryResponse and assigns it to the ExplainResult field.
func (o *ClusterSlowLogExplainResponse) SetExplainResult(v DmsQueryResponse) {
	o.ExplainResult = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterSlowLogExplainResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TemplateId != nil {
		toSerialize["templateId"] = o.TemplateId
	}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	if o.SqlType != nil {
		toSerialize["sqlType"] = o.SqlType
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.NormalizedQuery != nil {
		toSerialize["normalizedQuery"] = o.NormalizedQuery
	}
	if o.SampleTimestamp != nil {
		toSerialize["sampleTimestamp"] = o.SampleTimestamp
	}
	if o.ExecutionTime != nil {
		toSerialize["executionTime"] = o.ExecutionTime
	}
	if o.Sample != nil {
		toSerialize["sample"] = o.Sample
	}
	if o.ExplainResult != nil {
		toSerialize["explainResult"] = o.ExplainResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterSlowLogExplainResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TemplateId      *string                  `json:"templateId,omitempty"`
		Sql             *string                  `json:"sql,omitempty"`
		SqlType         *string                  `json:"sqlType,omitempty"`
		Database        *string                  `json:"database,omitempty"`
		NormalizedQuery *string                  `json:"normalizedQuery,omitempty"`
		SampleTimestamp *int64                   `json:"sampleTimestamp,omitempty"`
		ExecutionTime   *float64                 `json:"executionTime,omitempty"`
		Sample          *ClusterExecutionLogItem `json:"sample,omitempty"`
		ExplainResult   *DmsQueryResponse        `json:"explainResult,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"templateId", "sql", "sqlType", "database", "normalizedQuery", "sampleTimestamp", "executionTime", "sample", "explainResult"})
	} else {
		return err
	}

	hasInvalidField := false
	o.TemplateId = all.TemplateId
	o.Sql = all.Sql
	o.SqlType = all.SqlType
	o.Database = all.Database
	o.NormalizedQuery = all.NormalizedQuery
	o.SampleTimestamp = all.SampleTimestamp
	o.ExecutionTime = all.ExecutionTime
	if all.Sample != nil && all.Sample.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sample = all.Sample
	if all.ExplainResult != nil && all.ExplainResult.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ExplainResult = all.ExplainResult

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
