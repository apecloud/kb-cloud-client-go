// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiDataGatewayOperation struct {
	OperationId         *string                `json:"operationId,omitempty"`
	OrgName             *string                `json:"orgName,omitempty"`
	GatewayId           *string                `json:"gatewayId,omitempty"`
	AccessKeyId         *string                `json:"accessKeyId,omitempty"`
	DatasourceId        *string                `json:"datasourceId,omitempty"`
	PolicyId            *string                `json:"policyId,omitempty"`
	ToolName            *string                `json:"toolName,omitempty"`
	SqlType             *string                `json:"sqlType,omitempty"`
	PolicyDecision      *string                `json:"policyDecision,omitempty"`
	Status              *string                `json:"status,omitempty"`
	RiskLevel           *string                `json:"riskLevel,omitempty"`
	RowsReturned        *int64                 `json:"rowsReturned,omitempty"`
	DurationMs          *int64                 `json:"durationMs,omitempty"`
	MaskingApplied      *bool                  `json:"maskingApplied,omitempty"`
	MaskedColumns       []string               `json:"maskedColumns,omitempty"`
	ResultPreviewMasked map[string]interface{} `json:"resultPreviewMasked,omitempty"`
	ErrorCode           *string                `json:"errorCode,omitempty"`
	ErrorMessage        *string                `json:"errorMessage,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt           *time.Time             `json:"createdAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewayOperation instantiates a new AiDataGatewayOperation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewayOperation() *AiDataGatewayOperation {
	this := AiDataGatewayOperation{}
	return &this
}

// NewAiDataGatewayOperationWithDefaults instantiates a new AiDataGatewayOperation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiDataGatewayOperationWithDefaults() *AiDataGatewayOperation {
	this := AiDataGatewayOperation{}
	return &this
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetOperationId() string {
	if o == nil || o.OperationId == nil {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetOperationIdOk() (*string, bool) {
	if o == nil || o.OperationId == nil {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasOperationId() bool {
	return o != nil && o.OperationId != nil
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *AiDataGatewayOperation) SetOperationId(v string) {
	o.OperationId = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *AiDataGatewayOperation) SetOrgName(v string) {
	o.OrgName = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasGatewayId() bool {
	return o != nil && o.GatewayId != nil
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *AiDataGatewayOperation) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasAccessKeyId() bool {
	return o != nil && o.AccessKeyId != nil
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *AiDataGatewayOperation) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetDatasourceId returns the DatasourceId field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetDatasourceId() string {
	if o == nil || o.DatasourceId == nil {
		var ret string
		return ret
	}
	return *o.DatasourceId
}

// GetDatasourceIdOk returns a tuple with the DatasourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetDatasourceIdOk() (*string, bool) {
	if o == nil || o.DatasourceId == nil {
		return nil, false
	}
	return o.DatasourceId, true
}

// HasDatasourceId returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasDatasourceId() bool {
	return o != nil && o.DatasourceId != nil
}

// SetDatasourceId gets a reference to the given string and assigns it to the DatasourceId field.
func (o *AiDataGatewayOperation) SetDatasourceId(v string) {
	o.DatasourceId = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasPolicyId() bool {
	return o != nil && o.PolicyId != nil
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *AiDataGatewayOperation) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetToolName returns the ToolName field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetToolName() string {
	if o == nil || o.ToolName == nil {
		var ret string
		return ret
	}
	return *o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetToolNameOk() (*string, bool) {
	if o == nil || o.ToolName == nil {
		return nil, false
	}
	return o.ToolName, true
}

// HasToolName returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasToolName() bool {
	return o != nil && o.ToolName != nil
}

// SetToolName gets a reference to the given string and assigns it to the ToolName field.
func (o *AiDataGatewayOperation) SetToolName(v string) {
	o.ToolName = &v
}

// GetSqlType returns the SqlType field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetSqlType() string {
	if o == nil || o.SqlType == nil {
		var ret string
		return ret
	}
	return *o.SqlType
}

// GetSqlTypeOk returns a tuple with the SqlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetSqlTypeOk() (*string, bool) {
	if o == nil || o.SqlType == nil {
		return nil, false
	}
	return o.SqlType, true
}

// HasSqlType returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasSqlType() bool {
	return o != nil && o.SqlType != nil
}

// SetSqlType gets a reference to the given string and assigns it to the SqlType field.
func (o *AiDataGatewayOperation) SetSqlType(v string) {
	o.SqlType = &v
}

// GetPolicyDecision returns the PolicyDecision field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetPolicyDecision() string {
	if o == nil || o.PolicyDecision == nil {
		var ret string
		return ret
	}
	return *o.PolicyDecision
}

// GetPolicyDecisionOk returns a tuple with the PolicyDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetPolicyDecisionOk() (*string, bool) {
	if o == nil || o.PolicyDecision == nil {
		return nil, false
	}
	return o.PolicyDecision, true
}

// HasPolicyDecision returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasPolicyDecision() bool {
	return o != nil && o.PolicyDecision != nil
}

// SetPolicyDecision gets a reference to the given string and assigns it to the PolicyDecision field.
func (o *AiDataGatewayOperation) SetPolicyDecision(v string) {
	o.PolicyDecision = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AiDataGatewayOperation) SetStatus(v string) {
	o.Status = &v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetRiskLevel() string {
	if o == nil || o.RiskLevel == nil {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetRiskLevelOk() (*string, bool) {
	if o == nil || o.RiskLevel == nil {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasRiskLevel() bool {
	return o != nil && o.RiskLevel != nil
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *AiDataGatewayOperation) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// GetRowsReturned returns the RowsReturned field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetRowsReturned() int64 {
	if o == nil || o.RowsReturned == nil {
		var ret int64
		return ret
	}
	return *o.RowsReturned
}

// GetRowsReturnedOk returns a tuple with the RowsReturned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetRowsReturnedOk() (*int64, bool) {
	if o == nil || o.RowsReturned == nil {
		return nil, false
	}
	return o.RowsReturned, true
}

// HasRowsReturned returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasRowsReturned() bool {
	return o != nil && o.RowsReturned != nil
}

// SetRowsReturned gets a reference to the given int64 and assigns it to the RowsReturned field.
func (o *AiDataGatewayOperation) SetRowsReturned(v int64) {
	o.RowsReturned = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetDurationMs() int64 {
	if o == nil || o.DurationMs == nil {
		var ret int64
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetDurationMsOk() (*int64, bool) {
	if o == nil || o.DurationMs == nil {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasDurationMs() bool {
	return o != nil && o.DurationMs != nil
}

// SetDurationMs gets a reference to the given int64 and assigns it to the DurationMs field.
func (o *AiDataGatewayOperation) SetDurationMs(v int64) {
	o.DurationMs = &v
}

// GetMaskingApplied returns the MaskingApplied field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetMaskingApplied() bool {
	if o == nil || o.MaskingApplied == nil {
		var ret bool
		return ret
	}
	return *o.MaskingApplied
}

// GetMaskingAppliedOk returns a tuple with the MaskingApplied field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetMaskingAppliedOk() (*bool, bool) {
	if o == nil || o.MaskingApplied == nil {
		return nil, false
	}
	return o.MaskingApplied, true
}

// HasMaskingApplied returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasMaskingApplied() bool {
	return o != nil && o.MaskingApplied != nil
}

// SetMaskingApplied gets a reference to the given bool and assigns it to the MaskingApplied field.
func (o *AiDataGatewayOperation) SetMaskingApplied(v bool) {
	o.MaskingApplied = &v
}

// GetMaskedColumns returns the MaskedColumns field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetMaskedColumns() []string {
	if o == nil || o.MaskedColumns == nil {
		var ret []string
		return ret
	}
	return o.MaskedColumns
}

// GetMaskedColumnsOk returns a tuple with the MaskedColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetMaskedColumnsOk() (*[]string, bool) {
	if o == nil || o.MaskedColumns == nil {
		return nil, false
	}
	return &o.MaskedColumns, true
}

// HasMaskedColumns returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasMaskedColumns() bool {
	return o != nil && o.MaskedColumns != nil
}

// SetMaskedColumns gets a reference to the given []string and assigns it to the MaskedColumns field.
func (o *AiDataGatewayOperation) SetMaskedColumns(v []string) {
	o.MaskedColumns = v
}

// GetResultPreviewMasked returns the ResultPreviewMasked field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetResultPreviewMasked() map[string]interface{} {
	if o == nil || o.ResultPreviewMasked == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ResultPreviewMasked
}

// GetResultPreviewMaskedOk returns a tuple with the ResultPreviewMasked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetResultPreviewMaskedOk() (*map[string]interface{}, bool) {
	if o == nil || o.ResultPreviewMasked == nil {
		return nil, false
	}
	return &o.ResultPreviewMasked, true
}

// HasResultPreviewMasked returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasResultPreviewMasked() bool {
	return o != nil && o.ResultPreviewMasked != nil
}

// SetResultPreviewMasked gets a reference to the given map[string]interface{} and assigns it to the ResultPreviewMasked field.
func (o *AiDataGatewayOperation) SetResultPreviewMasked(v map[string]interface{}) {
	o.ResultPreviewMasked = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasErrorCode() bool {
	return o != nil && o.ErrorCode != nil
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *AiDataGatewayOperation) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AiDataGatewayOperation) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AiDataGatewayOperation) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AiDataGatewayOperation) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewayOperation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AiDataGatewayOperation) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AiDataGatewayOperation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewayOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OperationId != nil {
		toSerialize["operationId"] = o.OperationId
	}
	if o.OrgName != nil {
		toSerialize["orgName"] = o.OrgName
	}
	if o.GatewayId != nil {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if o.AccessKeyId != nil {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if o.DatasourceId != nil {
		toSerialize["datasourceId"] = o.DatasourceId
	}
	if o.PolicyId != nil {
		toSerialize["policyId"] = o.PolicyId
	}
	if o.ToolName != nil {
		toSerialize["toolName"] = o.ToolName
	}
	if o.SqlType != nil {
		toSerialize["sqlType"] = o.SqlType
	}
	if o.PolicyDecision != nil {
		toSerialize["policyDecision"] = o.PolicyDecision
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.RiskLevel != nil {
		toSerialize["riskLevel"] = o.RiskLevel
	}
	if o.RowsReturned != nil {
		toSerialize["rowsReturned"] = o.RowsReturned
	}
	if o.DurationMs != nil {
		toSerialize["durationMs"] = o.DurationMs
	}
	if o.MaskingApplied != nil {
		toSerialize["maskingApplied"] = o.MaskingApplied
	}
	if o.MaskedColumns != nil {
		toSerialize["maskedColumns"] = o.MaskedColumns
	}
	if o.ResultPreviewMasked != nil {
		toSerialize["resultPreviewMasked"] = o.ResultPreviewMasked
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewayOperation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OperationId         *string                `json:"operationId,omitempty"`
		OrgName             *string                `json:"orgName,omitempty"`
		GatewayId           *string                `json:"gatewayId,omitempty"`
		AccessKeyId         *string                `json:"accessKeyId,omitempty"`
		DatasourceId        *string                `json:"datasourceId,omitempty"`
		PolicyId            *string                `json:"policyId,omitempty"`
		ToolName            *string                `json:"toolName,omitempty"`
		SqlType             *string                `json:"sqlType,omitempty"`
		PolicyDecision      *string                `json:"policyDecision,omitempty"`
		Status              *string                `json:"status,omitempty"`
		RiskLevel           *string                `json:"riskLevel,omitempty"`
		RowsReturned        *int64                 `json:"rowsReturned,omitempty"`
		DurationMs          *int64                 `json:"durationMs,omitempty"`
		MaskingApplied      *bool                  `json:"maskingApplied,omitempty"`
		MaskedColumns       []string               `json:"maskedColumns,omitempty"`
		ResultPreviewMasked map[string]interface{} `json:"resultPreviewMasked,omitempty"`
		ErrorCode           *string                `json:"errorCode,omitempty"`
		ErrorMessage        *string                `json:"errorMessage,omitempty"`
		Metadata            map[string]interface{} `json:"metadata,omitempty"`
		CreatedAt           *time.Time             `json:"createdAt,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"operationId", "orgName", "gatewayId", "accessKeyId", "datasourceId", "policyId", "toolName", "sqlType", "policyDecision", "status", "riskLevel", "rowsReturned", "durationMs", "maskingApplied", "maskedColumns", "resultPreviewMasked", "errorCode", "errorMessage", "metadata", "createdAt"})
	} else {
		return err
	}
	o.OperationId = all.OperationId
	o.OrgName = all.OrgName
	o.GatewayId = all.GatewayId
	o.AccessKeyId = all.AccessKeyId
	o.DatasourceId = all.DatasourceId
	o.PolicyId = all.PolicyId
	o.ToolName = all.ToolName
	o.SqlType = all.SqlType
	o.PolicyDecision = all.PolicyDecision
	o.Status = all.Status
	o.RiskLevel = all.RiskLevel
	o.RowsReturned = all.RowsReturned
	o.DurationMs = all.DurationMs
	o.MaskingApplied = all.MaskingApplied
	o.MaskedColumns = all.MaskedColumns
	o.ResultPreviewMasked = all.ResultPreviewMasked
	o.ErrorCode = all.ErrorCode
	o.ErrorMessage = all.ErrorMessage
	o.Metadata = all.Metadata
	o.CreatedAt = all.CreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
