// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanResult struct {
	OriginSql      *string                    `json:"originSQL,omitempty"`
	ExplainSql     common.NullableString      `json:"explainSQL,omitempty"`
	EngineType     *string                    `json:"engineType,omitempty"`
	Mode           *DmsExecutionPlanMode      `json:"mode,omitempty"`
	HasActualStats *bool                      `json:"hasActualStats,omitempty"`
	RawFormat      *DmsExecutionPlanRawFormat `json:"rawFormat,omitempty"`
	// Engine-native raw plan payload exposed for fallback display. SDKs model it as a free-form JSON object.
	RawPlan       map[string]interface{}    `json:"rawPlan,omitempty"`
	RootNodeIds   []string                  `json:"rootNodeIds,omitempty"`
	Nodes         []DmsExecutionPlanNode    `json:"nodes,omitempty"`
	Edges         []DmsExecutionPlanEdge    `json:"edges,omitempty"`
	Summary       *DmsExecutionPlanSummary  `json:"summary,omitempty"`
	Warnings      []DmsExecutionPlanWarning `json:"warnings,omitempty"`
	FallbackTable *DmsQueryResponse         `json:"fallbackTable,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExecutionPlanResult instantiates a new DmsExecutionPlanResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExecutionPlanResult() *DmsExecutionPlanResult {
	this := DmsExecutionPlanResult{}
	return &this
}

// NewDmsExecutionPlanResultWithDefaults instantiates a new DmsExecutionPlanResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExecutionPlanResultWithDefaults() *DmsExecutionPlanResult {
	this := DmsExecutionPlanResult{}
	return &this
}

// GetOriginSql returns the OriginSql field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetOriginSql() string {
	if o == nil || o.OriginSql == nil {
		var ret string
		return ret
	}
	return *o.OriginSql
}

// GetOriginSqlOk returns a tuple with the OriginSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetOriginSqlOk() (*string, bool) {
	if o == nil || o.OriginSql == nil {
		return nil, false
	}
	return o.OriginSql, true
}

// HasOriginSql returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasOriginSql() bool {
	return o != nil && o.OriginSql != nil
}

// SetOriginSql gets a reference to the given string and assigns it to the OriginSql field.
func (o *DmsExecutionPlanResult) SetOriginSql(v string) {
	o.OriginSql = &v
}

// GetExplainSql returns the ExplainSql field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanResult) GetExplainSql() string {
	if o == nil || o.ExplainSql.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExplainSql.Get()
}

// GetExplainSqlOk returns a tuple with the ExplainSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanResult) GetExplainSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExplainSql.Get(), o.ExplainSql.IsSet()
}

// HasExplainSql returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasExplainSql() bool {
	return o != nil && o.ExplainSql.IsSet()
}

// SetExplainSql gets a reference to the given common.NullableString and assigns it to the ExplainSql field.
func (o *DmsExecutionPlanResult) SetExplainSql(v string) {
	o.ExplainSql.Set(&v)
}

// SetExplainSqlNil sets the value for ExplainSql to be an explicit nil.
func (o *DmsExecutionPlanResult) SetExplainSqlNil() {
	o.ExplainSql.Set(nil)
}

// UnsetExplainSql ensures that no value is present for ExplainSql, not even an explicit nil.
func (o *DmsExecutionPlanResult) UnsetExplainSql() {
	o.ExplainSql.Unset()
}

// GetEngineType returns the EngineType field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetEngineType() string {
	if o == nil || o.EngineType == nil {
		var ret string
		return ret
	}
	return *o.EngineType
}

// GetEngineTypeOk returns a tuple with the EngineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetEngineTypeOk() (*string, bool) {
	if o == nil || o.EngineType == nil {
		return nil, false
	}
	return o.EngineType, true
}

// HasEngineType returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasEngineType() bool {
	return o != nil && o.EngineType != nil
}

// SetEngineType gets a reference to the given string and assigns it to the EngineType field.
func (o *DmsExecutionPlanResult) SetEngineType(v string) {
	o.EngineType = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetMode() DmsExecutionPlanMode {
	if o == nil || o.Mode == nil {
		var ret DmsExecutionPlanMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetModeOk() (*DmsExecutionPlanMode, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasMode() bool {
	return o != nil && o.Mode != nil
}

// SetMode gets a reference to the given DmsExecutionPlanMode and assigns it to the Mode field.
func (o *DmsExecutionPlanResult) SetMode(v DmsExecutionPlanMode) {
	o.Mode = &v
}

// GetHasActualStats returns the HasActualStats field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetHasActualStats() bool {
	if o == nil || o.HasActualStats == nil {
		var ret bool
		return ret
	}
	return *o.HasActualStats
}

// GetHasActualStatsOk returns a tuple with the HasActualStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetHasActualStatsOk() (*bool, bool) {
	if o == nil || o.HasActualStats == nil {
		return nil, false
	}
	return o.HasActualStats, true
}

// HasHasActualStats returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasHasActualStats() bool {
	return o != nil && o.HasActualStats != nil
}

// SetHasActualStats gets a reference to the given bool and assigns it to the HasActualStats field.
func (o *DmsExecutionPlanResult) SetHasActualStats(v bool) {
	o.HasActualStats = &v
}

// GetRawFormat returns the RawFormat field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetRawFormat() DmsExecutionPlanRawFormat {
	if o == nil || o.RawFormat == nil {
		var ret DmsExecutionPlanRawFormat
		return ret
	}
	return *o.RawFormat
}

// GetRawFormatOk returns a tuple with the RawFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetRawFormatOk() (*DmsExecutionPlanRawFormat, bool) {
	if o == nil || o.RawFormat == nil {
		return nil, false
	}
	return o.RawFormat, true
}

// HasRawFormat returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasRawFormat() bool {
	return o != nil && o.RawFormat != nil
}

// SetRawFormat gets a reference to the given DmsExecutionPlanRawFormat and assigns it to the RawFormat field.
func (o *DmsExecutionPlanResult) SetRawFormat(v DmsExecutionPlanRawFormat) {
	o.RawFormat = &v
}

// GetRawPlan returns the RawPlan field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetRawPlan() map[string]interface{} {
	if o == nil || o.RawPlan == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RawPlan
}

// GetRawPlanOk returns a tuple with the RawPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetRawPlanOk() (*map[string]interface{}, bool) {
	if o == nil || o.RawPlan == nil {
		return nil, false
	}
	return &o.RawPlan, true
}

// HasRawPlan returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasRawPlan() bool {
	return o != nil && o.RawPlan != nil
}

// SetRawPlan gets a reference to the given map[string]interface{} and assigns it to the RawPlan field.
func (o *DmsExecutionPlanResult) SetRawPlan(v map[string]interface{}) {
	o.RawPlan = v
}

// GetRootNodeIds returns the RootNodeIds field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetRootNodeIds() []string {
	if o == nil || o.RootNodeIds == nil {
		var ret []string
		return ret
	}
	return o.RootNodeIds
}

// GetRootNodeIdsOk returns a tuple with the RootNodeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetRootNodeIdsOk() (*[]string, bool) {
	if o == nil || o.RootNodeIds == nil {
		return nil, false
	}
	return &o.RootNodeIds, true
}

// HasRootNodeIds returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasRootNodeIds() bool {
	return o != nil && o.RootNodeIds != nil
}

// SetRootNodeIds gets a reference to the given []string and assigns it to the RootNodeIds field.
func (o *DmsExecutionPlanResult) SetRootNodeIds(v []string) {
	o.RootNodeIds = v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetNodes() []DmsExecutionPlanNode {
	if o == nil || o.Nodes == nil {
		var ret []DmsExecutionPlanNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetNodesOk() (*[]DmsExecutionPlanNode, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasNodes() bool {
	return o != nil && o.Nodes != nil
}

// SetNodes gets a reference to the given []DmsExecutionPlanNode and assigns it to the Nodes field.
func (o *DmsExecutionPlanResult) SetNodes(v []DmsExecutionPlanNode) {
	o.Nodes = v
}

// GetEdges returns the Edges field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetEdges() []DmsExecutionPlanEdge {
	if o == nil || o.Edges == nil {
		var ret []DmsExecutionPlanEdge
		return ret
	}
	return o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetEdgesOk() (*[]DmsExecutionPlanEdge, bool) {
	if o == nil || o.Edges == nil {
		return nil, false
	}
	return &o.Edges, true
}

// HasEdges returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasEdges() bool {
	return o != nil && o.Edges != nil
}

// SetEdges gets a reference to the given []DmsExecutionPlanEdge and assigns it to the Edges field.
func (o *DmsExecutionPlanResult) SetEdges(v []DmsExecutionPlanEdge) {
	o.Edges = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetSummary() DmsExecutionPlanSummary {
	if o == nil || o.Summary == nil {
		var ret DmsExecutionPlanSummary
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetSummaryOk() (*DmsExecutionPlanSummary, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given DmsExecutionPlanSummary and assigns it to the Summary field.
func (o *DmsExecutionPlanResult) SetSummary(v DmsExecutionPlanSummary) {
	o.Summary = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetWarnings() []DmsExecutionPlanWarning {
	if o == nil || o.Warnings == nil {
		var ret []DmsExecutionPlanWarning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetWarningsOk() (*[]DmsExecutionPlanWarning, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []DmsExecutionPlanWarning and assigns it to the Warnings field.
func (o *DmsExecutionPlanResult) SetWarnings(v []DmsExecutionPlanWarning) {
	o.Warnings = v
}

// GetFallbackTable returns the FallbackTable field value if set, zero value otherwise.
func (o *DmsExecutionPlanResult) GetFallbackTable() DmsQueryResponse {
	if o == nil || o.FallbackTable == nil {
		var ret DmsQueryResponse
		return ret
	}
	return *o.FallbackTable
}

// GetFallbackTableOk returns a tuple with the FallbackTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanResult) GetFallbackTableOk() (*DmsQueryResponse, bool) {
	if o == nil || o.FallbackTable == nil {
		return nil, false
	}
	return o.FallbackTable, true
}

// HasFallbackTable returns a boolean if a field has been set.
func (o *DmsExecutionPlanResult) HasFallbackTable() bool {
	return o != nil && o.FallbackTable != nil
}

// SetFallbackTable gets a reference to the given DmsQueryResponse and assigns it to the FallbackTable field.
func (o *DmsExecutionPlanResult) SetFallbackTable(v DmsQueryResponse) {
	o.FallbackTable = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExecutionPlanResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OriginSql != nil {
		toSerialize["originSQL"] = o.OriginSql
	}
	if o.ExplainSql.IsSet() {
		toSerialize["explainSQL"] = o.ExplainSql.Get()
	}
	if o.EngineType != nil {
		toSerialize["engineType"] = o.EngineType
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.HasActualStats != nil {
		toSerialize["hasActualStats"] = o.HasActualStats
	}
	if o.RawFormat != nil {
		toSerialize["rawFormat"] = o.RawFormat
	}
	if o.RawPlan != nil {
		toSerialize["rawPlan"] = o.RawPlan
	}
	if o.RootNodeIds != nil {
		toSerialize["rootNodeIds"] = o.RootNodeIds
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	if o.Edges != nil {
		toSerialize["edges"] = o.Edges
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.FallbackTable != nil {
		toSerialize["fallbackTable"] = o.FallbackTable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExecutionPlanResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OriginSql      *string                    `json:"originSQL,omitempty"`
		ExplainSql     common.NullableString      `json:"explainSQL,omitempty"`
		EngineType     *string                    `json:"engineType,omitempty"`
		Mode           *DmsExecutionPlanMode      `json:"mode,omitempty"`
		HasActualStats *bool                      `json:"hasActualStats,omitempty"`
		RawFormat      *DmsExecutionPlanRawFormat `json:"rawFormat,omitempty"`
		RawPlan        map[string]interface{}     `json:"rawPlan,omitempty"`
		RootNodeIds    []string                   `json:"rootNodeIds,omitempty"`
		Nodes          []DmsExecutionPlanNode     `json:"nodes,omitempty"`
		Edges          []DmsExecutionPlanEdge     `json:"edges,omitempty"`
		Summary        *DmsExecutionPlanSummary   `json:"summary,omitempty"`
		Warnings       []DmsExecutionPlanWarning  `json:"warnings,omitempty"`
		FallbackTable  *DmsQueryResponse          `json:"fallbackTable,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"originSQL", "explainSQL", "engineType", "mode", "hasActualStats", "rawFormat", "rawPlan", "rootNodeIds", "nodes", "edges", "summary", "warnings", "fallbackTable"})
	} else {
		return err
	}

	hasInvalidField := false
	o.OriginSql = all.OriginSql
	o.ExplainSql = all.ExplainSql
	o.EngineType = all.EngineType
	if all.Mode != nil && !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = all.Mode
	}
	o.HasActualStats = all.HasActualStats
	if all.RawFormat != nil && !all.RawFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.RawFormat = all.RawFormat
	}
	o.RawPlan = all.RawPlan
	o.RootNodeIds = all.RootNodeIds
	o.Nodes = all.Nodes
	o.Edges = all.Edges
	if all.Summary != nil && all.Summary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Summary = all.Summary
	o.Warnings = all.Warnings
	if all.FallbackTable != nil && all.FallbackTable.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FallbackTable = all.FallbackTable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
