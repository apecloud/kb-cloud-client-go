// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// DiagnosticSQLDiagnosisOperator One PostgreSQL plan operator summarized from EXPLAIN output.
type DiagnosticSQLDiagnosisOperator struct {
	// PostgreSQL plan node type.
	NodeType *string `json:"nodeType,omitempty"`
	// Table or relation name when present.
	RelationName *string `json:"relationName,omitempty"`
	// Index name when this operator uses an index.
	IndexName *string `json:"indexName,omitempty"`
	// Estimated rows for this operator.
	EstimatedRows *float64 `json:"estimatedRows,omitempty"`
	// Estimated total cost for this operator.
	EstimatedCost *float64 `json:"estimatedCost,omitempty"`
	// Whether this operator appears to use an index.
	UsesIndex *bool `json:"usesIndex,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSQLDiagnosisOperator instantiates a new DiagnosticSQLDiagnosisOperator object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSQLDiagnosisOperator() *DiagnosticSQLDiagnosisOperator {
	this := DiagnosticSQLDiagnosisOperator{}
	return &this
}

// NewDiagnosticSQLDiagnosisOperatorWithDefaults instantiates a new DiagnosticSQLDiagnosisOperator object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSQLDiagnosisOperatorWithDefaults() *DiagnosticSQLDiagnosisOperator {
	this := DiagnosticSQLDiagnosisOperator{}
	return &this
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisOperator) GetNodeType() string {
	if o == nil || o.NodeType == nil {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisOperator) GetNodeTypeOk() (*string, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisOperator) HasNodeType() bool {
	return o != nil && o.NodeType != nil
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *DiagnosticSQLDiagnosisOperator) SetNodeType(v string) {
	o.NodeType = &v
}

// GetRelationName returns the RelationName field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisOperator) GetRelationName() string {
	if o == nil || o.RelationName == nil {
		var ret string
		return ret
	}
	return *o.RelationName
}

// GetRelationNameOk returns a tuple with the RelationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisOperator) GetRelationNameOk() (*string, bool) {
	if o == nil || o.RelationName == nil {
		return nil, false
	}
	return o.RelationName, true
}

// HasRelationName returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisOperator) HasRelationName() bool {
	return o != nil && o.RelationName != nil
}

// SetRelationName gets a reference to the given string and assigns it to the RelationName field.
func (o *DiagnosticSQLDiagnosisOperator) SetRelationName(v string) {
	o.RelationName = &v
}

// GetIndexName returns the IndexName field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisOperator) GetIndexName() string {
	if o == nil || o.IndexName == nil {
		var ret string
		return ret
	}
	return *o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisOperator) GetIndexNameOk() (*string, bool) {
	if o == nil || o.IndexName == nil {
		return nil, false
	}
	return o.IndexName, true
}

// HasIndexName returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisOperator) HasIndexName() bool {
	return o != nil && o.IndexName != nil
}

// SetIndexName gets a reference to the given string and assigns it to the IndexName field.
func (o *DiagnosticSQLDiagnosisOperator) SetIndexName(v string) {
	o.IndexName = &v
}

// GetEstimatedRows returns the EstimatedRows field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisOperator) GetEstimatedRows() float64 {
	if o == nil || o.EstimatedRows == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedRows
}

// GetEstimatedRowsOk returns a tuple with the EstimatedRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisOperator) GetEstimatedRowsOk() (*float64, bool) {
	if o == nil || o.EstimatedRows == nil {
		return nil, false
	}
	return o.EstimatedRows, true
}

// HasEstimatedRows returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisOperator) HasEstimatedRows() bool {
	return o != nil && o.EstimatedRows != nil
}

// SetEstimatedRows gets a reference to the given float64 and assigns it to the EstimatedRows field.
func (o *DiagnosticSQLDiagnosisOperator) SetEstimatedRows(v float64) {
	o.EstimatedRows = &v
}

// GetEstimatedCost returns the EstimatedCost field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisOperator) GetEstimatedCost() float64 {
	if o == nil || o.EstimatedCost == nil {
		var ret float64
		return ret
	}
	return *o.EstimatedCost
}

// GetEstimatedCostOk returns a tuple with the EstimatedCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisOperator) GetEstimatedCostOk() (*float64, bool) {
	if o == nil || o.EstimatedCost == nil {
		return nil, false
	}
	return o.EstimatedCost, true
}

// HasEstimatedCost returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisOperator) HasEstimatedCost() bool {
	return o != nil && o.EstimatedCost != nil
}

// SetEstimatedCost gets a reference to the given float64 and assigns it to the EstimatedCost field.
func (o *DiagnosticSQLDiagnosisOperator) SetEstimatedCost(v float64) {
	o.EstimatedCost = &v
}

// GetUsesIndex returns the UsesIndex field value if set, zero value otherwise.
func (o *DiagnosticSQLDiagnosisOperator) GetUsesIndex() bool {
	if o == nil || o.UsesIndex == nil {
		var ret bool
		return ret
	}
	return *o.UsesIndex
}

// GetUsesIndexOk returns a tuple with the UsesIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSQLDiagnosisOperator) GetUsesIndexOk() (*bool, bool) {
	if o == nil || o.UsesIndex == nil {
		return nil, false
	}
	return o.UsesIndex, true
}

// HasUsesIndex returns a boolean if a field has been set.
func (o *DiagnosticSQLDiagnosisOperator) HasUsesIndex() bool {
	return o != nil && o.UsesIndex != nil
}

// SetUsesIndex gets a reference to the given bool and assigns it to the UsesIndex field.
func (o *DiagnosticSQLDiagnosisOperator) SetUsesIndex(v bool) {
	o.UsesIndex = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSQLDiagnosisOperator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NodeType != nil {
		toSerialize["nodeType"] = o.NodeType
	}
	if o.RelationName != nil {
		toSerialize["relationName"] = o.RelationName
	}
	if o.IndexName != nil {
		toSerialize["indexName"] = o.IndexName
	}
	if o.EstimatedRows != nil {
		toSerialize["estimatedRows"] = o.EstimatedRows
	}
	if o.EstimatedCost != nil {
		toSerialize["estimatedCost"] = o.EstimatedCost
	}
	if o.UsesIndex != nil {
		toSerialize["usesIndex"] = o.UsesIndex
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSQLDiagnosisOperator) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodeType      *string  `json:"nodeType,omitempty"`
		RelationName  *string  `json:"relationName,omitempty"`
		IndexName     *string  `json:"indexName,omitempty"`
		EstimatedRows *float64 `json:"estimatedRows,omitempty"`
		EstimatedCost *float64 `json:"estimatedCost,omitempty"`
		UsesIndex     *bool    `json:"usesIndex,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodeType", "relationName", "indexName", "estimatedRows", "estimatedCost", "usesIndex"})
	} else {
		return err
	}
	o.NodeType = all.NodeType
	o.RelationName = all.RelationName
	o.IndexName = all.IndexName
	o.EstimatedRows = all.EstimatedRows
	o.EstimatedCost = all.EstimatedCost
	o.UsesIndex = all.UsesIndex

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
