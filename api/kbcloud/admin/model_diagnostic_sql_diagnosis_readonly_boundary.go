// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSQLDiagnosisReadonlyBoundary Safety boundary enforced by the backend for this SQL diagnosis.
type DiagnosticSQLDiagnosisReadonlyBoundary struct {
	// Normalized statement class accepted by the safety gate.
	StatementType string `json:"statementType"`
	// Whether the backend used an EXPLAIN-only execution path.
	OnlyExplain bool `json:"onlyExplain"`
	// Always false; EXPLAIN ANALYZE is forbidden.
	ExplainAnalyze bool `json:"explainAnalyze"`
	// Always false; the original SQL is never executed.
	OriginalSqlExecuted bool `json:"originalSQLExecuted"`
	// Backend EXPLAIN timeout in seconds.
	TimeoutSeconds int64 `json:"timeoutSeconds"`
	// Maximum accepted EXPLAIN payload size.
	ResultSizeLimitBytes int64 `json:"resultSizeLimitBytes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSQLDiagnosisReadonlyBoundary instantiates a new DiagnosticSQLDiagnosisReadonlyBoundary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSQLDiagnosisReadonlyBoundary(statementType string, onlyExplain bool, explainAnalyze bool, originalSqlExecuted bool, timeoutSeconds int64, resultSizeLimitBytes int64) *DiagnosticSQLDiagnosisReadonlyBoundary {
	this := DiagnosticSQLDiagnosisReadonlyBoundary{}
	this.StatementType = statementType
	this.OnlyExplain = onlyExplain
	this.ExplainAnalyze = explainAnalyze
	this.OriginalSqlExecuted = originalSqlExecuted
	this.TimeoutSeconds = timeoutSeconds
	this.ResultSizeLimitBytes = resultSizeLimitBytes
	return &this
}

// NewDiagnosticSQLDiagnosisReadonlyBoundaryWithDefaults instantiates a new DiagnosticSQLDiagnosisReadonlyBoundary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSQLDiagnosisReadonlyBoundaryWithDefaults() *DiagnosticSQLDiagnosisReadonlyBoundary {
	this := DiagnosticSQLDiagnosisReadonlyBoundary{}
	return &this
}

// GetStatementType returns the StatementType field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetStatementType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StatementType
}

// GetStatementTypeOk returns a tuple with the StatementType field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetStatementTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatementType, true
}

// SetStatementType sets field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) SetStatementType(v string) {
	o.StatementType = v
}

// GetOnlyExplain returns the OnlyExplain field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetOnlyExplain() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.OnlyExplain
}

// GetOnlyExplainOk returns a tuple with the OnlyExplain field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetOnlyExplainOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnlyExplain, true
}

// SetOnlyExplain sets field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) SetOnlyExplain(v bool) {
	o.OnlyExplain = v
}

// GetExplainAnalyze returns the ExplainAnalyze field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetExplainAnalyze() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.ExplainAnalyze
}

// GetExplainAnalyzeOk returns a tuple with the ExplainAnalyze field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetExplainAnalyzeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExplainAnalyze, true
}

// SetExplainAnalyze sets field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) SetExplainAnalyze(v bool) {
	o.ExplainAnalyze = v
}

// GetOriginalSqlExecuted returns the OriginalSqlExecuted field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetOriginalSqlExecuted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.OriginalSqlExecuted
}

// GetOriginalSqlExecutedOk returns a tuple with the OriginalSqlExecuted field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetOriginalSqlExecutedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalSqlExecuted, true
}

// SetOriginalSqlExecuted sets field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) SetOriginalSqlExecuted(v bool) {
	o.OriginalSqlExecuted = v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetTimeoutSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetTimeoutSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeoutSeconds, true
}

// SetTimeoutSeconds sets field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) SetTimeoutSeconds(v int64) {
	o.TimeoutSeconds = v
}

// GetResultSizeLimitBytes returns the ResultSizeLimitBytes field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetResultSizeLimitBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ResultSizeLimitBytes
}

// GetResultSizeLimitBytesOk returns a tuple with the ResultSizeLimitBytes field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) GetResultSizeLimitBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultSizeLimitBytes, true
}

// SetResultSizeLimitBytes sets field value.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) SetResultSizeLimitBytes(v int64) {
	o.ResultSizeLimitBytes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSQLDiagnosisReadonlyBoundary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["statementType"] = o.StatementType
	toSerialize["onlyExplain"] = o.OnlyExplain
	toSerialize["explainAnalyze"] = o.ExplainAnalyze
	toSerialize["originalSQLExecuted"] = o.OriginalSqlExecuted
	toSerialize["timeoutSeconds"] = o.TimeoutSeconds
	toSerialize["resultSizeLimitBytes"] = o.ResultSizeLimitBytes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSQLDiagnosisReadonlyBoundary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		StatementType        *string `json:"statementType"`
		OnlyExplain          *bool   `json:"onlyExplain"`
		ExplainAnalyze       *bool   `json:"explainAnalyze"`
		OriginalSqlExecuted  *bool   `json:"originalSQLExecuted"`
		TimeoutSeconds       *int64  `json:"timeoutSeconds"`
		ResultSizeLimitBytes *int64  `json:"resultSizeLimitBytes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.StatementType == nil {
		return fmt.Errorf("required field statementType missing")
	}
	if all.OnlyExplain == nil {
		return fmt.Errorf("required field onlyExplain missing")
	}
	if all.ExplainAnalyze == nil {
		return fmt.Errorf("required field explainAnalyze missing")
	}
	if all.OriginalSqlExecuted == nil {
		return fmt.Errorf("required field originalSQLExecuted missing")
	}
	if all.TimeoutSeconds == nil {
		return fmt.Errorf("required field timeoutSeconds missing")
	}
	if all.ResultSizeLimitBytes == nil {
		return fmt.Errorf("required field resultSizeLimitBytes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"statementType", "onlyExplain", "explainAnalyze", "originalSQLExecuted", "timeoutSeconds", "resultSizeLimitBytes"})
	} else {
		return err
	}
	o.StatementType = *all.StatementType
	o.OnlyExplain = *all.OnlyExplain
	o.ExplainAnalyze = *all.ExplainAnalyze
	o.OriginalSqlExecuted = *all.OriginalSqlExecuted
	o.TimeoutSeconds = *all.TimeoutSeconds
	o.ResultSizeLimitBytes = *all.ResultSizeLimitBytes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
