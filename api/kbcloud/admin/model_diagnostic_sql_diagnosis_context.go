// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// DiagnosticSQLDiagnosisContext Optional source context for a SQL diagnosis request.
type DiagnosticSQLDiagnosisContext struct {
	// Where the SQL came from, such as manual_input, slow_sql_detail, or controlled_fixture.
	Source *string `json:"source,omitempty"`
	// Diagnostic task ID that provided the SQL, when applicable.
	TaskId *string `json:"taskId,omitempty"`
	// Diagnostic metric that suggested the SQL, such as slowSql.
	Metric *string `json:"metric,omitempty"`
	// SQL digest or fingerprint. This is safe to show and copy.
	Digest *string `json:"digest,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSQLDiagnosisContext instantiates a new DiagnosticSQLDiagnosisContext object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSQLDiagnosisContext() *DiagnosticSQLDiagnosisContext {
	this := DiagnosticSQLDiagnosisContext{}
	return &this
}

// NewDiagnosticSQLDiagnosisContextWithDefaults instantiates a new DiagnosticSQLDiagnosisContext object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSQLDiagnosisContextWithDefaults() *DiagnosticSQLDiagnosisContext {
	this := DiagnosticSQLDiagnosisContext{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisContext) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisContext) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisContext) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *DiagnosticSQLDiagnosisContext) SetSource(v string) {
	o.Source = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisContext) GetTaskId() string {
	if o == nil || o.TaskId == nil {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisContext) GetTaskIdOk() (*string, bool) {
	if o == nil || o.TaskId == nil {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisContext) HasTaskId() bool {
	return o != nil && o.TaskId != nil
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *DiagnosticSQLDiagnosisContext) SetTaskId(v string) {
	o.TaskId = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisContext) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisContext) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisContext) HasMetric() bool {
	return o != nil && o.Metric != nil
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *DiagnosticSQLDiagnosisContext) SetMetric(v string) {
	o.Metric = &v
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisContext) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisContext) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisContext) HasDigest() bool {
	return o != nil && o.Digest != nil
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *DiagnosticSQLDiagnosisContext) SetDigest(v string) {
	o.Digest = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSQLDiagnosisContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.TaskId != nil {
		toSerialize["taskId"] = o.TaskId
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSQLDiagnosisContext) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source *string `json:"source,omitempty"`
		TaskId *string `json:"taskId,omitempty"`
		Metric *string `json:"metric,omitempty"`
		Digest *string `json:"digest,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"source", "taskId", "metric", "digest"})
	} else {
		return err
	}
	o.Source = all.Source
	o.TaskId = all.TaskId
	o.Metric = all.Metric
	o.Digest = all.Digest

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
