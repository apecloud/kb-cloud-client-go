// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanNode struct {
	Id           *string                         `json:"id,omitempty"`
	ParentIds    []string                        `json:"parentIds,omitempty"`
	ChildIds     []string                        `json:"childIds,omitempty"`
	Order        common.NullableInt64            `json:"order,omitempty"`
	Category     *DmsExecutionPlanNodeCategory   `json:"category,omitempty"`
	NodeType     *string                         `json:"nodeType,omitempty"`
	Label        *string                         `json:"label,omitempty"`
	RelationName common.NullableString           `json:"relationName,omitempty"`
	SchemaName   common.NullableString           `json:"schemaName,omitempty"`
	Alias        common.NullableString           `json:"alias,omitempty"`
	IndexName    common.NullableString           `json:"indexName,omitempty"`
	Cost         *DmsExecutionPlanNodeCost       `json:"cost,omitempty"`
	Rows         *DmsExecutionPlanNodeRows       `json:"rows,omitempty"`
	Timing       *DmsExecutionPlanNodeTiming     `json:"timing,omitempty"`
	Conditions   *DmsExecutionPlanNodeConditions `json:"conditions,omitempty"`
	Details      map[string]interface{}          `json:"details,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExecutionPlanNode instantiates a new DmsExecutionPlanNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExecutionPlanNode() *DmsExecutionPlanNode {
	this := DmsExecutionPlanNode{}
	return &this
}

// NewDmsExecutionPlanNodeWithDefaults instantiates a new DmsExecutionPlanNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExecutionPlanNodeWithDefaults() *DmsExecutionPlanNode {
	this := DmsExecutionPlanNode{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DmsExecutionPlanNode) SetId(v string) {
	o.Id = &v
}

// GetParentIds returns the ParentIds field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetParentIds() []string {
	if o == nil || o.ParentIds == nil {
		var ret []string
		return ret
	}
	return o.ParentIds
}

// GetParentIdsOk returns a tuple with the ParentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetParentIdsOk() (*[]string, bool) {
	if o == nil || o.ParentIds == nil {
		return nil, false
	}
	return &o.ParentIds, true
}

// HasParentIds returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasParentIds() bool {
	return o != nil && o.ParentIds != nil
}

// SetParentIds gets a reference to the given []string and assigns it to the ParentIds field.
func (o *DmsExecutionPlanNode) SetParentIds(v []string) {
	o.ParentIds = v
}

// GetChildIds returns the ChildIds field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetChildIds() []string {
	if o == nil || o.ChildIds == nil {
		var ret []string
		return ret
	}
	return o.ChildIds
}

// GetChildIdsOk returns a tuple with the ChildIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetChildIdsOk() (*[]string, bool) {
	if o == nil || o.ChildIds == nil {
		return nil, false
	}
	return &o.ChildIds, true
}

// HasChildIds returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasChildIds() bool {
	return o != nil && o.ChildIds != nil
}

// SetChildIds gets a reference to the given []string and assigns it to the ChildIds field.
func (o *DmsExecutionPlanNode) SetChildIds(v []string) {
	o.ChildIds = v
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNode) GetOrder() int64 {
	if o == nil || o.Order.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNode) GetOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasOrder() bool {
	return o != nil && o.Order.IsSet()
}

// SetOrder gets a reference to the given common.NullableInt64 and assigns it to the Order field.
func (o *DmsExecutionPlanNode) SetOrder(v int64) {
	o.Order.Set(&v)
}

// SetOrderNil sets the value for Order to be an explicit nil.
func (o *DmsExecutionPlanNode) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil.
func (o *DmsExecutionPlanNode) UnsetOrder() {
	o.Order.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetCategory() DmsExecutionPlanNodeCategory {
	if o == nil || o.Category == nil {
		var ret DmsExecutionPlanNodeCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetCategoryOk() (*DmsExecutionPlanNodeCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given DmsExecutionPlanNodeCategory and assigns it to the Category field.
func (o *DmsExecutionPlanNode) SetCategory(v DmsExecutionPlanNodeCategory) {
	o.Category = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetNodeType() string {
	if o == nil || o.NodeType == nil {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetNodeTypeOk() (*string, bool) {
	if o == nil || o.NodeType == nil {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasNodeType() bool {
	return o != nil && o.NodeType != nil
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *DmsExecutionPlanNode) SetNodeType(v string) {
	o.NodeType = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasLabel() bool {
	return o != nil && o.Label != nil
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DmsExecutionPlanNode) SetLabel(v string) {
	o.Label = &v
}

// GetRelationName returns the RelationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNode) GetRelationName() string {
	if o == nil || o.RelationName.Get() == nil {
		var ret string
		return ret
	}
	return *o.RelationName.Get()
}

// GetRelationNameOk returns a tuple with the RelationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNode) GetRelationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelationName.Get(), o.RelationName.IsSet()
}

// HasRelationName returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasRelationName() bool {
	return o != nil && o.RelationName.IsSet()
}

// SetRelationName gets a reference to the given common.NullableString and assigns it to the RelationName field.
func (o *DmsExecutionPlanNode) SetRelationName(v string) {
	o.RelationName.Set(&v)
}

// SetRelationNameNil sets the value for RelationName to be an explicit nil.
func (o *DmsExecutionPlanNode) SetRelationNameNil() {
	o.RelationName.Set(nil)
}

// UnsetRelationName ensures that no value is present for RelationName, not even an explicit nil.
func (o *DmsExecutionPlanNode) UnsetRelationName() {
	o.RelationName.Unset()
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNode) GetSchemaName() string {
	if o == nil || o.SchemaName.Get() == nil {
		var ret string
		return ret
	}
	return *o.SchemaName.Get()
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNode) GetSchemaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SchemaName.Get(), o.SchemaName.IsSet()
}

// HasSchemaName returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasSchemaName() bool {
	return o != nil && o.SchemaName.IsSet()
}

// SetSchemaName gets a reference to the given common.NullableString and assigns it to the SchemaName field.
func (o *DmsExecutionPlanNode) SetSchemaName(v string) {
	o.SchemaName.Set(&v)
}

// SetSchemaNameNil sets the value for SchemaName to be an explicit nil.
func (o *DmsExecutionPlanNode) SetSchemaNameNil() {
	o.SchemaName.Set(nil)
}

// UnsetSchemaName ensures that no value is present for SchemaName, not even an explicit nil.
func (o *DmsExecutionPlanNode) UnsetSchemaName() {
	o.SchemaName.Unset()
}

// GetAlias returns the Alias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNode) GetAlias() string {
	if o == nil || o.Alias.Get() == nil {
		var ret string
		return ret
	}
	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNode) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// HasAlias returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasAlias() bool {
	return o != nil && o.Alias.IsSet()
}

// SetAlias gets a reference to the given common.NullableString and assigns it to the Alias field.
func (o *DmsExecutionPlanNode) SetAlias(v string) {
	o.Alias.Set(&v)
}

// SetAliasNil sets the value for Alias to be an explicit nil.
func (o *DmsExecutionPlanNode) SetAliasNil() {
	o.Alias.Set(nil)
}

// UnsetAlias ensures that no value is present for Alias, not even an explicit nil.
func (o *DmsExecutionPlanNode) UnsetAlias() {
	o.Alias.Unset()
}

// GetIndexName returns the IndexName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanNode) GetIndexName() string {
	if o == nil || o.IndexName.Get() == nil {
		var ret string
		return ret
	}
	return *o.IndexName.Get()
}

// GetIndexNameOk returns a tuple with the IndexName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanNode) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexName.Get(), o.IndexName.IsSet()
}

// HasIndexName returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasIndexName() bool {
	return o != nil && o.IndexName.IsSet()
}

// SetIndexName gets a reference to the given common.NullableString and assigns it to the IndexName field.
func (o *DmsExecutionPlanNode) SetIndexName(v string) {
	o.IndexName.Set(&v)
}

// SetIndexNameNil sets the value for IndexName to be an explicit nil.
func (o *DmsExecutionPlanNode) SetIndexNameNil() {
	o.IndexName.Set(nil)
}

// UnsetIndexName ensures that no value is present for IndexName, not even an explicit nil.
func (o *DmsExecutionPlanNode) UnsetIndexName() {
	o.IndexName.Unset()
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetCost() DmsExecutionPlanNodeCost {
	if o == nil || o.Cost == nil {
		var ret DmsExecutionPlanNodeCost
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetCostOk() (*DmsExecutionPlanNodeCost, bool) {
	if o == nil || o.Cost == nil {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasCost() bool {
	return o != nil && o.Cost != nil
}

// SetCost gets a reference to the given DmsExecutionPlanNodeCost and assigns it to the Cost field.
func (o *DmsExecutionPlanNode) SetCost(v DmsExecutionPlanNodeCost) {
	o.Cost = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetRows() DmsExecutionPlanNodeRows {
	if o == nil || o.Rows == nil {
		var ret DmsExecutionPlanNodeRows
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetRowsOk() (*DmsExecutionPlanNodeRows, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given DmsExecutionPlanNodeRows and assigns it to the Rows field.
func (o *DmsExecutionPlanNode) SetRows(v DmsExecutionPlanNodeRows) {
	o.Rows = &v
}

// GetTiming returns the Timing field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetTiming() DmsExecutionPlanNodeTiming {
	if o == nil || o.Timing == nil {
		var ret DmsExecutionPlanNodeTiming
		return ret
	}
	return *o.Timing
}

// GetTimingOk returns a tuple with the Timing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetTimingOk() (*DmsExecutionPlanNodeTiming, bool) {
	if o == nil || o.Timing == nil {
		return nil, false
	}
	return o.Timing, true
}

// HasTiming returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasTiming() bool {
	return o != nil && o.Timing != nil
}

// SetTiming gets a reference to the given DmsExecutionPlanNodeTiming and assigns it to the Timing field.
func (o *DmsExecutionPlanNode) SetTiming(v DmsExecutionPlanNodeTiming) {
	o.Timing = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetConditions() DmsExecutionPlanNodeConditions {
	if o == nil || o.Conditions == nil {
		var ret DmsExecutionPlanNodeConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetConditionsOk() (*DmsExecutionPlanNodeConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasConditions() bool {
	return o != nil && o.Conditions != nil
}

// SetConditions gets a reference to the given DmsExecutionPlanNodeConditions and assigns it to the Conditions field.
func (o *DmsExecutionPlanNode) SetConditions(v DmsExecutionPlanNodeConditions) {
	o.Conditions = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *DmsExecutionPlanNode) GetDetails() map[string]interface{} {
	if o == nil || o.Details == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanNode) GetDetailsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return &o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *DmsExecutionPlanNode) HasDetails() bool {
	return o != nil && o.Details != nil
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *DmsExecutionPlanNode) SetDetails(v map[string]interface{}) {
	o.Details = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExecutionPlanNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ParentIds != nil {
		toSerialize["parentIds"] = o.ParentIds
	}
	if o.ChildIds != nil {
		toSerialize["childIds"] = o.ChildIds
	}
	if o.Order.IsSet() {
		toSerialize["order"] = o.Order.Get()
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.NodeType != nil {
		toSerialize["nodeType"] = o.NodeType
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.RelationName.IsSet() {
		toSerialize["relationName"] = o.RelationName.Get()
	}
	if o.SchemaName.IsSet() {
		toSerialize["schemaName"] = o.SchemaName.Get()
	}
	if o.Alias.IsSet() {
		toSerialize["alias"] = o.Alias.Get()
	}
	if o.IndexName.IsSet() {
		toSerialize["indexName"] = o.IndexName.Get()
	}
	if o.Cost != nil {
		toSerialize["cost"] = o.Cost
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.Timing != nil {
		toSerialize["timing"] = o.Timing
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExecutionPlanNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id           *string                         `json:"id,omitempty"`
		ParentIds    []string                        `json:"parentIds,omitempty"`
		ChildIds     []string                        `json:"childIds,omitempty"`
		Order        common.NullableInt64            `json:"order,omitempty"`
		Category     *DmsExecutionPlanNodeCategory   `json:"category,omitempty"`
		NodeType     *string                         `json:"nodeType,omitempty"`
		Label        *string                         `json:"label,omitempty"`
		RelationName common.NullableString           `json:"relationName,omitempty"`
		SchemaName   common.NullableString           `json:"schemaName,omitempty"`
		Alias        common.NullableString           `json:"alias,omitempty"`
		IndexName    common.NullableString           `json:"indexName,omitempty"`
		Cost         *DmsExecutionPlanNodeCost       `json:"cost,omitempty"`
		Rows         *DmsExecutionPlanNodeRows       `json:"rows,omitempty"`
		Timing       *DmsExecutionPlanNodeTiming     `json:"timing,omitempty"`
		Conditions   *DmsExecutionPlanNodeConditions `json:"conditions,omitempty"`
		Details      map[string]interface{}          `json:"details,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "parentIds", "childIds", "order", "category", "nodeType", "label", "relationName", "schemaName", "alias", "indexName", "cost", "rows", "timing", "conditions", "details"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.ParentIds = all.ParentIds
	o.ChildIds = all.ChildIds
	o.Order = all.Order
	if all.Category != nil && !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = all.Category
	}
	o.NodeType = all.NodeType
	o.Label = all.Label
	o.RelationName = all.RelationName
	o.SchemaName = all.SchemaName
	o.Alias = all.Alias
	o.IndexName = all.IndexName
	if all.Cost != nil && all.Cost.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cost = all.Cost
	if all.Rows != nil && all.Rows.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rows = all.Rows
	if all.Timing != nil && all.Timing.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Timing = all.Timing
	if all.Conditions != nil && all.Conditions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Conditions = all.Conditions
	o.Details = all.Details

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
