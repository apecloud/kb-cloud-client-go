// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HealthDiagnosisCapability struct {
	Module      HealthDiagnosisCapabilityModule `json:"module"`
	State       HealthDiagnosisCapabilityState  `json:"state"`
	PayloadKind *HealthDiagnosisPayloadKind     `json:"payloadKind,omitempty"`
	// Machine-readable reason for the capability state.
	ReasonCode string `json:"reasonCode"`
	// Human-readable capability summary.
	Summary *string `json:"summary,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHealthDiagnosisCapability instantiates a new HealthDiagnosisCapability object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHealthDiagnosisCapability(module HealthDiagnosisCapabilityModule, state HealthDiagnosisCapabilityState, reasonCode string) *HealthDiagnosisCapability {
	this := HealthDiagnosisCapability{}
	this.Module = module
	this.State = state
	this.ReasonCode = reasonCode
	return &this
}

// NewHealthDiagnosisCapabilityWithDefaults instantiates a new HealthDiagnosisCapability object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHealthDiagnosisCapabilityWithDefaults() *HealthDiagnosisCapability {
	this := HealthDiagnosisCapability{}
	return &this
}

// GetModule returns the Module field value.
func (o *HealthDiagnosisCapability) GetModule() HealthDiagnosisCapabilityModule {
	if o == nil {
		var ret HealthDiagnosisCapabilityModule
		return ret
	}
	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisCapability) GetModuleOk() (*HealthDiagnosisCapabilityModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value.
func (o *HealthDiagnosisCapability) SetModule(v HealthDiagnosisCapabilityModule) {
	o.Module = v
}

// GetState returns the State field value.
func (o *HealthDiagnosisCapability) GetState() HealthDiagnosisCapabilityState {
	if o == nil {
		var ret HealthDiagnosisCapabilityState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisCapability) GetStateOk() (*HealthDiagnosisCapabilityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *HealthDiagnosisCapability) SetState(v HealthDiagnosisCapabilityState) {
	o.State = v
}

// GetPayloadKind returns the PayloadKind field value if set, zero value otherwise.
func (o *HealthDiagnosisCapability) GetPayloadKind() HealthDiagnosisPayloadKind {
	if o == nil || o.PayloadKind == nil {
		var ret HealthDiagnosisPayloadKind
		return ret
	}
	return *o.PayloadKind
}

// GetPayloadKindOk returns a tuple with the PayloadKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisCapability) GetPayloadKindOk() (*HealthDiagnosisPayloadKind, bool) {
	if o == nil || o.PayloadKind == nil {
		return nil, false
	}
	return o.PayloadKind, true
}

// HasPayloadKind returns a boolean if a field has been set.
func (o *HealthDiagnosisCapability) HasPayloadKind() bool {
	return o != nil && o.PayloadKind != nil
}

// SetPayloadKind gets a reference to the given HealthDiagnosisPayloadKind and assigns it to the PayloadKind field.
func (o *HealthDiagnosisCapability) SetPayloadKind(v HealthDiagnosisPayloadKind) {
	o.PayloadKind = &v
}

// GetReasonCode returns the ReasonCode field value.
func (o *HealthDiagnosisCapability) GetReasonCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisCapability) GetReasonCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value.
func (o *HealthDiagnosisCapability) SetReasonCode(v string) {
	o.ReasonCode = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *HealthDiagnosisCapability) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisCapability) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *HealthDiagnosisCapability) HasSummary() bool {
	return o != nil && o.Summary != nil
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *HealthDiagnosisCapability) SetSummary(v string) {
	o.Summary = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HealthDiagnosisCapability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["module"] = o.Module
	toSerialize["state"] = o.State
	if o.PayloadKind != nil {
		toSerialize["payloadKind"] = o.PayloadKind
	}
	toSerialize["reasonCode"] = o.ReasonCode
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HealthDiagnosisCapability) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Module      *HealthDiagnosisCapabilityModule `json:"module"`
		State       *HealthDiagnosisCapabilityState  `json:"state"`
		PayloadKind *HealthDiagnosisPayloadKind      `json:"payloadKind,omitempty"`
		ReasonCode  *string                          `json:"reasonCode"`
		Summary     *string                          `json:"summary,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Module == nil {
		return fmt.Errorf("required field module missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.ReasonCode == nil {
		return fmt.Errorf("required field reasonCode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"module", "state", "payloadKind", "reasonCode", "summary"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Module.IsValid() {
		hasInvalidField = true
	} else {
		o.Module = *all.Module
	}
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	if all.PayloadKind != nil && !all.PayloadKind.IsValid() {
		hasInvalidField = true
	} else {
		o.PayloadKind = all.PayloadKind
	}
	o.ReasonCode = *all.ReasonCode
	o.Summary = all.Summary

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
