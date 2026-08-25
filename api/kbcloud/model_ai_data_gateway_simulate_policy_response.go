// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiDataGatewaySimulatePolicyResponse struct {
	Decision       *string  `json:"decision,omitempty"`
	PolicyId       *string  `json:"policyId,omitempty"`
	Reason         *string  `json:"reason,omitempty"`
	ReasonCode     *string  `json:"reasonCode,omitempty"`
	MaxRows        *int32   `json:"maxRows,omitempty"`
	TimeoutSeconds *int32   `json:"timeoutSeconds,omitempty"`
	MaskedColumns  []string `json:"maskedColumns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiDataGatewaySimulatePolicyResponse instantiates a new AiDataGatewaySimulatePolicyResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiDataGatewaySimulatePolicyResponse() *AiDataGatewaySimulatePolicyResponse {
	this := AiDataGatewaySimulatePolicyResponse{}
	return &this
}

// NewAiDataGatewaySimulatePolicyResponseWithDefaults instantiates a new AiDataGatewaySimulatePolicyResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiDataGatewaySimulatePolicyResponseWithDefaults() *AiDataGatewaySimulatePolicyResponse {
	this := AiDataGatewaySimulatePolicyResponse{}
	return &this
}

// GetDecision returns the Decision field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyResponse) GetDecision() string {
	if o == nil || o.Decision == nil {
		var ret string
		return ret
	}
	return *o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyResponse) GetDecisionOk() (*string, bool) {
	if o == nil || o.Decision == nil {
		return nil, false
	}
	return o.Decision, true
}

// HasDecision returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyResponse) HasDecision() bool {
	return o != nil && o.Decision != nil
}

// SetDecision gets a reference to the given string and assigns it to the Decision field.
func (o *AiDataGatewaySimulatePolicyResponse) SetDecision(v string) {
	o.Decision = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyResponse) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyResponse) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyResponse) HasPolicyId() bool {
	return o != nil && o.PolicyId != nil
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *AiDataGatewaySimulatePolicyResponse) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyResponse) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyResponse) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyResponse) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AiDataGatewaySimulatePolicyResponse) SetReason(v string) {
	o.Reason = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyResponse) GetReasonCode() string {
	if o == nil || o.ReasonCode == nil {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyResponse) GetReasonCodeOk() (*string, bool) {
	if o == nil || o.ReasonCode == nil {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyResponse) HasReasonCode() bool {
	return o != nil && o.ReasonCode != nil
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *AiDataGatewaySimulatePolicyResponse) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetMaxRows returns the MaxRows field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyResponse) GetMaxRows() int32 {
	if o == nil || o.MaxRows == nil {
		var ret int32
		return ret
	}
	return *o.MaxRows
}

// GetMaxRowsOk returns a tuple with the MaxRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyResponse) GetMaxRowsOk() (*int32, bool) {
	if o == nil || o.MaxRows == nil {
		return nil, false
	}
	return o.MaxRows, true
}

// HasMaxRows returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyResponse) HasMaxRows() bool {
	return o != nil && o.MaxRows != nil
}

// SetMaxRows gets a reference to the given int32 and assigns it to the MaxRows field.
func (o *AiDataGatewaySimulatePolicyResponse) SetMaxRows(v int32) {
	o.MaxRows = &v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyResponse) GetTimeoutSeconds() int32 {
	if o == nil || o.TimeoutSeconds == nil {
		var ret int32
		return ret
	}
	return *o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyResponse) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil || o.TimeoutSeconds == nil {
		return nil, false
	}
	return o.TimeoutSeconds, true
}

// HasTimeoutSeconds returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyResponse) HasTimeoutSeconds() bool {
	return o != nil && o.TimeoutSeconds != nil
}

// SetTimeoutSeconds gets a reference to the given int32 and assigns it to the TimeoutSeconds field.
func (o *AiDataGatewaySimulatePolicyResponse) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds = &v
}

// GetMaskedColumns returns the MaskedColumns field value if set, zero value otherwise.
func (o *AiDataGatewaySimulatePolicyResponse) GetMaskedColumns() []string {
	if o == nil || o.MaskedColumns == nil {
		var ret []string
		return ret
	}
	return o.MaskedColumns
}

// GetMaskedColumnsOk returns a tuple with the MaskedColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiDataGatewaySimulatePolicyResponse) GetMaskedColumnsOk() (*[]string, bool) {
	if o == nil || o.MaskedColumns == nil {
		return nil, false
	}
	return &o.MaskedColumns, true
}

// HasMaskedColumns returns a boolean if a field has been set.
func (o *AiDataGatewaySimulatePolicyResponse) HasMaskedColumns() bool {
	return o != nil && o.MaskedColumns != nil
}

// SetMaskedColumns gets a reference to the given []string and assigns it to the MaskedColumns field.
func (o *AiDataGatewaySimulatePolicyResponse) SetMaskedColumns(v []string) {
	o.MaskedColumns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiDataGatewaySimulatePolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Decision != nil {
		toSerialize["decision"] = o.Decision
	}
	if o.PolicyId != nil {
		toSerialize["policyId"] = o.PolicyId
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.ReasonCode != nil {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	if o.MaxRows != nil {
		toSerialize["maxRows"] = o.MaxRows
	}
	if o.TimeoutSeconds != nil {
		toSerialize["timeoutSeconds"] = o.TimeoutSeconds
	}
	if o.MaskedColumns != nil {
		toSerialize["maskedColumns"] = o.MaskedColumns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiDataGatewaySimulatePolicyResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Decision       *string  `json:"decision,omitempty"`
		PolicyId       *string  `json:"policyId,omitempty"`
		Reason         *string  `json:"reason,omitempty"`
		ReasonCode     *string  `json:"reasonCode,omitempty"`
		MaxRows        *int32   `json:"maxRows,omitempty"`
		TimeoutSeconds *int32   `json:"timeoutSeconds,omitempty"`
		MaskedColumns  []string `json:"maskedColumns,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"decision", "policyId", "reason", "reasonCode", "maxRows", "timeoutSeconds", "maskedColumns"})
	} else {
		return err
	}
	o.Decision = all.Decision
	o.PolicyId = all.PolicyId
	o.Reason = all.Reason
	o.ReasonCode = all.ReasonCode
	o.MaxRows = all.MaxRows
	o.TimeoutSeconds = all.TimeoutSeconds
	o.MaskedColumns = all.MaskedColumns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
