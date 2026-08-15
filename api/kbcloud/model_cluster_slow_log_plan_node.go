// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterSlowLogPlanNode Best-effort visual execution plan node parsed from EXPLAIN.
type ClusterSlowLogPlanNode struct {
	Id            *string  `json:"id,omitempty"`
	ParentId      *string  `json:"parentId,omitempty"`
	Operation     *string  `json:"operation,omitempty"`
	Table         *string  `json:"table,omitempty"`
	AccessType    *string  `json:"accessType,omitempty"`
	Detail        *string  `json:"detail,omitempty"`
	CostStart     *float64 `json:"costStart,omitempty"`
	CostEnd       *float64 `json:"costEnd,omitempty"`
	Rows          *int64   `json:"rows,omitempty"`
	IssueCode     *string  `json:"issueCode,omitempty"`
	IssueSeverity *string  `json:"issueSeverity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterSlowLogPlanNode instantiates a new ClusterSlowLogPlanNode object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterSlowLogPlanNode() *ClusterSlowLogPlanNode {
	this := ClusterSlowLogPlanNode{}
	return &this
}

// NewClusterSlowLogPlanNodeWithDefaults instantiates a new ClusterSlowLogPlanNode object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterSlowLogPlanNodeWithDefaults() *ClusterSlowLogPlanNode {
	this := ClusterSlowLogPlanNode{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterSlowLogPlanNode) SetId(v string) {
	o.Id = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasParentId() bool {
	return o != nil && o.ParentId != nil
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *ClusterSlowLogPlanNode) SetParentId(v string) {
	o.ParentId = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasOperation() bool {
	return o != nil && o.Operation != nil
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ClusterSlowLogPlanNode) SetOperation(v string) {
	o.Operation = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *ClusterSlowLogPlanNode) SetTable(v string) {
	o.Table = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetAccessType() string {
	if o == nil || o.AccessType == nil {
		var ret string
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetAccessTypeOk() (*string, bool) {
	if o == nil || o.AccessType == nil {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasAccessType() bool {
	return o != nil && o.AccessType != nil
}

// SetAccessType gets a reference to the given string and assigns it to the AccessType field.
func (o *ClusterSlowLogPlanNode) SetAccessType(v string) {
	o.AccessType = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasDetail() bool {
	return o != nil && o.Detail != nil
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ClusterSlowLogPlanNode) SetDetail(v string) {
	o.Detail = &v
}

// GetCostStart returns the CostStart field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetCostStart() float64 {
	if o == nil || o.CostStart == nil {
		var ret float64
		return ret
	}
	return *o.CostStart
}

// GetCostStartOk returns a tuple with the CostStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetCostStartOk() (*float64, bool) {
	if o == nil || o.CostStart == nil {
		return nil, false
	}
	return o.CostStart, true
}

// HasCostStart returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasCostStart() bool {
	return o != nil && o.CostStart != nil
}

// SetCostStart gets a reference to the given float64 and assigns it to the CostStart field.
func (o *ClusterSlowLogPlanNode) SetCostStart(v float64) {
	o.CostStart = &v
}

// GetCostEnd returns the CostEnd field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetCostEnd() float64 {
	if o == nil || o.CostEnd == nil {
		var ret float64
		return ret
	}
	return *o.CostEnd
}

// GetCostEndOk returns a tuple with the CostEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetCostEndOk() (*float64, bool) {
	if o == nil || o.CostEnd == nil {
		return nil, false
	}
	return o.CostEnd, true
}

// HasCostEnd returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasCostEnd() bool {
	return o != nil && o.CostEnd != nil
}

// SetCostEnd gets a reference to the given float64 and assigns it to the CostEnd field.
func (o *ClusterSlowLogPlanNode) SetCostEnd(v float64) {
	o.CostEnd = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetRows() int64 {
	if o == nil || o.Rows == nil {
		var ret int64
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetRowsOk() (*int64, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given int64 and assigns it to the Rows field.
func (o *ClusterSlowLogPlanNode) SetRows(v int64) {
	o.Rows = &v
}

// GetIssueCode returns the IssueCode field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetIssueCode() string {
	if o == nil || o.IssueCode == nil {
		var ret string
		return ret
	}
	return *o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetIssueCodeOk() (*string, bool) {
	if o == nil || o.IssueCode == nil {
		return nil, false
	}
	return o.IssueCode, true
}

// HasIssueCode returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasIssueCode() bool {
	return o != nil && o.IssueCode != nil
}

// SetIssueCode gets a reference to the given string and assigns it to the IssueCode field.
func (o *ClusterSlowLogPlanNode) SetIssueCode(v string) {
	o.IssueCode = &v
}

// GetIssueSeverity returns the IssueSeverity field value if set, zero value otherwise.
func (o *ClusterSlowLogPlanNode) GetIssueSeverity() string {
	if o == nil || o.IssueSeverity == nil {
		var ret string
		return ret
	}
	return *o.IssueSeverity
}

// GetIssueSeverityOk returns a tuple with the IssueSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogPlanNode) GetIssueSeverityOk() (*string, bool) {
	if o == nil || o.IssueSeverity == nil {
		return nil, false
	}
	return o.IssueSeverity, true
}

// HasIssueSeverity returns a boolean if a field has been set.
func (o *ClusterSlowLogPlanNode) HasIssueSeverity() bool {
	return o != nil && o.IssueSeverity != nil
}

// SetIssueSeverity gets a reference to the given string and assigns it to the IssueSeverity field.
func (o *ClusterSlowLogPlanNode) SetIssueSeverity(v string) {
	o.IssueSeverity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterSlowLogPlanNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}
	if o.AccessType != nil {
		toSerialize["accessType"] = o.AccessType
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.CostStart != nil {
		toSerialize["costStart"] = o.CostStart
	}
	if o.CostEnd != nil {
		toSerialize["costEnd"] = o.CostEnd
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.IssueCode != nil {
		toSerialize["issueCode"] = o.IssueCode
	}
	if o.IssueSeverity != nil {
		toSerialize["issueSeverity"] = o.IssueSeverity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterSlowLogPlanNode) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *string  `json:"id,omitempty"`
		ParentId      *string  `json:"parentId,omitempty"`
		Operation     *string  `json:"operation,omitempty"`
		Table         *string  `json:"table,omitempty"`
		AccessType    *string  `json:"accessType,omitempty"`
		Detail        *string  `json:"detail,omitempty"`
		CostStart     *float64 `json:"costStart,omitempty"`
		CostEnd       *float64 `json:"costEnd,omitempty"`
		Rows          *int64   `json:"rows,omitempty"`
		IssueCode     *string  `json:"issueCode,omitempty"`
		IssueSeverity *string  `json:"issueSeverity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "parentId", "operation", "table", "accessType", "detail", "costStart", "costEnd", "rows", "issueCode", "issueSeverity"})
	} else {
		return err
	}
	o.Id = all.Id
	o.ParentId = all.ParentId
	o.Operation = all.Operation
	o.Table = all.Table
	o.AccessType = all.AccessType
	o.Detail = all.Detail
	o.CostStart = all.CostStart
	o.CostEnd = all.CostEnd
	o.Rows = all.Rows
	o.IssueCode = all.IssueCode
	o.IssueSeverity = all.IssueSeverity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
