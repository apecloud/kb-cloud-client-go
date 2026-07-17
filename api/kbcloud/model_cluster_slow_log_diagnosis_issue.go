// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterSlowLogDiagnosisIssue A single detected slow log diagnosis issue
type ClusterSlowLogDiagnosisIssue struct {
	Code *string `json:"code,omitempty"`
	// Issue severity. Current values are info, low, medium, and high.
	Severity *string `json:"severity,omitempty"`
	// Evidence source used by the local rule
	Source   *string                       `json:"source,omitempty"`
	Message  *LocalizedDescription         `json:"message,omitempty"`
	Evidence []ClusterSlowLogLocalizedText `json:"evidence,omitempty"`
	NodeId   *string                       `json:"nodeId,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterSlowLogDiagnosisIssue instantiates a new ClusterSlowLogDiagnosisIssue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterSlowLogDiagnosisIssue() *ClusterSlowLogDiagnosisIssue {
	this := ClusterSlowLogDiagnosisIssue{}
	return &this
}

// NewClusterSlowLogDiagnosisIssueWithDefaults instantiates a new ClusterSlowLogDiagnosisIssue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterSlowLogDiagnosisIssueWithDefaults() *ClusterSlowLogDiagnosisIssue {
	this := ClusterSlowLogDiagnosisIssue{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosisIssue) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosisIssue) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosisIssue) HasCode() bool {
	return o != nil && o.Code != nil
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ClusterSlowLogDiagnosisIssue) SetCode(v string) {
	o.Code = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosisIssue) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosisIssue) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosisIssue) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ClusterSlowLogDiagnosisIssue) SetSeverity(v string) {
	o.Severity = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosisIssue) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosisIssue) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosisIssue) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ClusterSlowLogDiagnosisIssue) SetSource(v string) {
	o.Source = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosisIssue) GetMessage() LocalizedDescription {
	if o == nil || o.Message == nil {
		var ret LocalizedDescription
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosisIssue) GetMessageOk() (*LocalizedDescription, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosisIssue) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given LocalizedDescription and assigns it to the Message field.
func (o *ClusterSlowLogDiagnosisIssue) SetMessage(v LocalizedDescription) {
	o.Message = &v
}

// GetEvidence returns the Evidence field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosisIssue) GetEvidence() []ClusterSlowLogLocalizedText {
	if o == nil || o.Evidence == nil {
		var ret []ClusterSlowLogLocalizedText
		return ret
	}
	return o.Evidence
}

// GetEvidenceOk returns a tuple with the Evidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosisIssue) GetEvidenceOk() (*[]ClusterSlowLogLocalizedText, bool) {
	if o == nil || o.Evidence == nil {
		return nil, false
	}
	return &o.Evidence, true
}

// HasEvidence returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosisIssue) HasEvidence() bool {
	return o != nil && o.Evidence != nil
}

// SetEvidence gets a reference to the given []ClusterSlowLogLocalizedText and assigns it to the Evidence field.
func (o *ClusterSlowLogDiagnosisIssue) SetEvidence(v []ClusterSlowLogLocalizedText) {
	o.Evidence = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ClusterSlowLogDiagnosisIssue) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSlowLogDiagnosisIssue) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ClusterSlowLogDiagnosisIssue) HasNodeId() bool {
	return o != nil && o.NodeId != nil
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ClusterSlowLogDiagnosisIssue) SetNodeId(v string) {
	o.NodeId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterSlowLogDiagnosisIssue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Evidence != nil {
		toSerialize["evidence"] = o.Evidence
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterSlowLogDiagnosisIssue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code     *string                       `json:"code,omitempty"`
		Severity *string                       `json:"severity,omitempty"`
		Source   *string                       `json:"source,omitempty"`
		Message  *LocalizedDescription         `json:"message,omitempty"`
		Evidence []ClusterSlowLogLocalizedText `json:"evidence,omitempty"`
		NodeId   *string                       `json:"nodeId,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"code", "severity", "source", "message", "evidence", "nodeId"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Code = all.Code
	o.Severity = all.Severity
	o.Source = all.Source
	if all.Message != nil && all.Message.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Message = all.Message
	o.Evidence = all.Evidence
	o.NodeId = all.NodeId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
