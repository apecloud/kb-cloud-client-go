// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DmsExecutionPlanWarning struct {
	Code     *string                          `json:"code,omitempty"`
	Severity *DmsExecutionPlanWarningSeverity `json:"severity,omitempty"`
	NodeId   common.NullableString            `json:"nodeId,omitempty"`
	Message  *string                          `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDmsExecutionPlanWarning instantiates a new DmsExecutionPlanWarning object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDmsExecutionPlanWarning() *DmsExecutionPlanWarning {
	this := DmsExecutionPlanWarning{}
	return &this
}

// NewDmsExecutionPlanWarningWithDefaults instantiates a new DmsExecutionPlanWarning object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDmsExecutionPlanWarningWithDefaults() *DmsExecutionPlanWarning {
	this := DmsExecutionPlanWarning{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DmsExecutionPlanWarning) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanWarning) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DmsExecutionPlanWarning) HasCode() bool {
	return o != nil && o.Code != nil
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DmsExecutionPlanWarning) SetCode(v string) {
	o.Code = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *DmsExecutionPlanWarning) GetSeverity() DmsExecutionPlanWarningSeverity {
	if o == nil || o.Severity == nil {
		var ret DmsExecutionPlanWarningSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanWarning) GetSeverityOk() (*DmsExecutionPlanWarningSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *DmsExecutionPlanWarning) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given DmsExecutionPlanWarningSeverity and assigns it to the Severity field.
func (o *DmsExecutionPlanWarning) SetSeverity(v DmsExecutionPlanWarningSeverity) {
	o.Severity = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DmsExecutionPlanWarning) GetNodeId() string {
	if o == nil || o.NodeId.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeId.Get()
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *DmsExecutionPlanWarning) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeId.Get(), o.NodeId.IsSet()
}

// HasNodeId returns a boolean if a field has been set.
func (o *DmsExecutionPlanWarning) HasNodeId() bool {
	return o != nil && o.NodeId.IsSet()
}

// SetNodeId gets a reference to the given common.NullableString and assigns it to the NodeId field.
func (o *DmsExecutionPlanWarning) SetNodeId(v string) {
	o.NodeId.Set(&v)
}

// SetNodeIdNil sets the value for NodeId to be an explicit nil.
func (o *DmsExecutionPlanWarning) SetNodeIdNil() {
	o.NodeId.Set(nil)
}

// UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil.
func (o *DmsExecutionPlanWarning) UnsetNodeId() {
	o.NodeId.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DmsExecutionPlanWarning) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DmsExecutionPlanWarning) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DmsExecutionPlanWarning) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DmsExecutionPlanWarning) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DmsExecutionPlanWarning) MarshalJSON() ([]byte, error) {
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
	if o.NodeId.IsSet() {
		toSerialize["nodeId"] = o.NodeId.Get()
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DmsExecutionPlanWarning) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code     *string                          `json:"code,omitempty"`
		Severity *DmsExecutionPlanWarningSeverity `json:"severity,omitempty"`
		NodeId   common.NullableString            `json:"nodeId,omitempty"`
		Message  *string                          `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"code", "severity", "nodeId", "message"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Code = all.Code
	if all.Severity != nil && !all.Severity.IsValid() {
		hasInvalidField = true
	} else {
		o.Severity = all.Severity
	}
	o.NodeId = all.NodeId
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
