// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// DiagnosticSQLDiagnosisPlanSummary PostgreSQL EXPLAIN plan summary for Console display.
type DiagnosticSQLDiagnosisPlanSummary struct {
	// Human-readable plan summary.
	Summary *string `json:"summary,omitempty"`
	// Top-level plan node type.
	TopNodeType *string `json:"topNodeType,omitempty"`
	// Tables or relations mentioned by the plan.
	Tables []string `json:"tables,omitempty"`
	// Flattened operator summary.
	Operators []DiagnosticSQLDiagnosisOperator `json:"operators,omitempty"`
	// Whether any summarized operator uses an index.
	UsesIndex *bool `json:"usesIndex,omitempty"`
	// Top-level estimated rows.
	EstimatedRows *float64 `json:"estimatedRows,omitempty"`
	// Top-level total cost.
	TotalCost *float64 `json:"totalCost,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSQLDiagnosisPlanSummary instantiates a new DiagnosticSQLDiagnosisPlanSummary object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSQLDiagnosisPlanSummary() *DiagnosticSQLDiagnosisPlanSummary {
	this := DiagnosticSQLDiagnosisPlanSummary{}
	return &this
}

// NewDiagnosticSQLDiagnosisPlanSummaryWithDefaults instantiates a new DiagnosticSQLDiagnosisPlanSummary object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSQLDiagnosisPlanSummaryWithDefaults() *DiagnosticSQLDiagnosisPlanSummary {
	this := DiagnosticSQLDiagnosisPlanSummary{}
	return &this
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *DiagnosticSQLDiagnosisPlanSummary) SetSummary(v string) {
	o.Summary = &v
}

// GetTopNodeType returns the TopNodeType field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetTopNodeType() string {
	if o == nil || o.TopNodeType == nil {
		var ret string
		return ret
	}
	return *o.TopNodeType
}

// GetTopNodeTypeOk returns a tuple with the TopNodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetTopNodeTypeOk() (*string, bool) {
	if o == nil || o.TopNodeType == nil {
		return nil, false
	}
	return o.TopNodeType, true
}

// HasTopNodeType returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) HasTopNodeType() bool {
	return o != nil && o.TopNodeType != nil
}

// SetTopNodeType gets a reference to the given string and assigns it to the TopNodeType field.
func (o *DiagnosticSQLDiagnosisPlanSummary) SetTopNodeType(v string) {
	o.TopNodeType = &v
}

// GetTables returns the Tables field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetTables() []string {
	if o == nil || o.Tables == nil {
		var ret []string
		return ret
	}
	return o.Tables
}

// GetTablesOk returns a tuple with the Tables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetTablesOk() (*[]string, bool) {
	if o == nil || o.Tables == nil {
		return nil, false
	}
	return &o.Tables, true
}

// HasTables returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) HasTables() bool {
	return o != nil && o.Tables != nil
}

// SetTables gets a reference to the given []string and assigns it to the Tables field.
func (o *DiagnosticSQLDiagnosisPlanSummary) SetTables(v []string) {
	o.Tables = v
}

// GetOperators returns the Operators field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetOperators() []DiagnosticSQLDiagnosisOperator {
	if o == nil || o.Operators == nil {
		var ret []DiagnosticSQLDiagnosisOperator
		return ret
	}
	return o.Operators
}

// GetOperatorsOk returns a tuple with the Operators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetOperatorsOk() (*[]DiagnosticSQLDiagnosisOperator, bool) {
	if o == nil || o.Operators == nil {
		return nil, false
	}
	return &o.Operators, true
}

// HasOperators returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) HasOperators() bool {
	return o != nil && o.Operators != nil
}

// SetOperators gets a reference to the given []DiagnosticSQLDiagnosisOperator and assigns it to the Operators field.
func (o *DiagnosticSQLDiagnosisPlanSummary) SetOperators(v []DiagnosticSQLDiagnosisOperator) {
	o.Operators = v
}

// GetUsesIndex returns the UsesIndex field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetUsesIndex() bool {
	if o == nil || o.UsesIndex == nil {
		var ret bool
		return ret
	}
	return *o.UsesIndex
}

// GetUsesIndexOk returns a tuple with the UsesIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetUsesIndexOk() (*bool, bool) {
	if o == nil || o.UsesIndex == nil {
		return nil, false
	}
	return o.UsesIndex, true
}

// HasUsesIndex returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) HasUsesIndex() bool {
	return o != nil && o.UsesIndex != nil
}

// SetUsesIndex gets a reference to the given bool and assigns it to the UsesIndex field.
func (o *DiagnosticSQLDiagnosisPlanSummary) SetUsesIndex(v bool) {
	o.UsesIndex = &v
}

// GetEstimatedRows returns the EstimatedRows field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetEstimatedRows() float64 {
	if o == nil || o.EstimatedRows == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedRows
}

// GetEstimatedRowsOk returns a tuple with the EstimatedRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetEstimatedRowsOk() (*float64, bool) {
	if o == nil || o.EstimatedRows == nil {
		return nil, false
	}
	return o.EstimatedRows, true
}

// HasEstimatedRows returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) HasEstimatedRows() bool {
	return o != nil && o.EstimatedRows != nil
}

// SetEstimatedRows gets a reference to the given float64 and assigns it to the EstimatedRows field.
func (o *DiagnosticSQLDiagnosisPlanSummary) SetEstimatedRows(v float64) {
	o.EstimatedRows = &v
}

// GetTotalCost returns the TotalCost field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetTotalCost() float64 {
	if o == nil || o.TotalCost == nil {
		var ret float64
		return ret
	}
	return *o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) GetTotalCostOk() (*float64, bool) {
	if o == nil || o.TotalCost == nil {
		return nil, false
	}
	return o.TotalCost, true
}

// HasTotalCost returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisPlanSummary) HasTotalCost() bool {
	return o != nil && o.TotalCost != nil
}

// SetTotalCost gets a reference to the given float64 and assigns it to the TotalCost field.
func (o *DiagnosticSQLDiagnosisPlanSummary) SetTotalCost(v float64) {
	o.TotalCost = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSQLDiagnosisPlanSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.TopNodeType != nil {
		toSerialize["topNodeType"] = o.TopNodeType
	}
	if o.Tables != nil {
		toSerialize["tables"] = o.Tables
	}
	if o.Operators != nil {
		toSerialize["operators"] = o.Operators
	}
	if o.UsesIndex != nil {
		toSerialize["usesIndex"] = o.UsesIndex
	}
	if o.EstimatedRows != nil {
		toSerialize["estimatedRows"] = o.EstimatedRows
	}
	if o.TotalCost != nil {
		toSerialize["totalCost"] = o.TotalCost
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSQLDiagnosisPlanSummary) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Summary       *string                          `json:"summary,omitempty"`
		TopNodeType   *string                          `json:"topNodeType,omitempty"`
		Tables        []string                         `json:"tables,omitempty"`
		Operators     []DiagnosticSQLDiagnosisOperator `json:"operators,omitempty"`
		UsesIndex     *bool                            `json:"usesIndex,omitempty"`
		EstimatedRows *float64                         `json:"estimatedRows,omitempty"`
		TotalCost     *float64                         `json:"totalCost,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"summary", "topNodeType", "tables", "operators", "usesIndex", "estimatedRows", "totalCost"})
	} else {
		return err
	}
	o.Summary = all.Summary
	o.TopNodeType = all.TopNodeType
	o.Tables = all.Tables
	o.Operators = all.Operators
	o.UsesIndex = all.UsesIndex
	o.EstimatedRows = all.EstimatedRows
	o.TotalCost = all.TotalCost

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
