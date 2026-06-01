// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HealthDiagnosisModule struct {
	Module                     HealthDiagnosisCapabilityModule    `json:"module"`
	State                      HealthDiagnosisModuleState         `json:"state"`
	PayloadKind                HealthDiagnosisPayloadKind         `json:"payloadKind"`
	Summary                    string                             `json:"summary"`
	ReasonCode                 *string                            `json:"reasonCode,omitempty"`
	ReadonlyBoundary           HealthDiagnosisReadonlyBoundary    `json:"readonlyBoundary"`
	PostgresqlRealtimeSessions *PostgresqlRealtimeSessionsPayload `json:"postgresqlRealtimeSessions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHealthDiagnosisModule instantiates a new HealthDiagnosisModule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHealthDiagnosisModule(module HealthDiagnosisCapabilityModule, state HealthDiagnosisModuleState, payloadKind HealthDiagnosisPayloadKind, summary string, readonlyBoundary HealthDiagnosisReadonlyBoundary) *HealthDiagnosisModule {
	this := HealthDiagnosisModule{}
	this.Module = module
	this.State = state
	this.PayloadKind = payloadKind
	this.Summary = summary
	this.ReadonlyBoundary = readonlyBoundary
	return &this
}

// NewHealthDiagnosisModuleWithDefaults instantiates a new HealthDiagnosisModule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHealthDiagnosisModuleWithDefaults() *HealthDiagnosisModule {
	this := HealthDiagnosisModule{}
	return &this
}

// GetModule returns the Module field value.
func (o *HealthDiagnosisModule) GetModule() HealthDiagnosisCapabilityModule {
	if o == nil {
		var ret HealthDiagnosisCapabilityModule
		return ret
	}
	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisModule) GetModuleOk() (*HealthDiagnosisCapabilityModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value.
func (o *HealthDiagnosisModule) SetModule(v HealthDiagnosisCapabilityModule) {
	o.Module = v
}

// GetState returns the State field value.
func (o *HealthDiagnosisModule) GetState() HealthDiagnosisModuleState {
	if o == nil {
		var ret HealthDiagnosisModuleState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisModule) GetStateOk() (*HealthDiagnosisModuleState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *HealthDiagnosisModule) SetState(v HealthDiagnosisModuleState) {
	o.State = v
}

// GetPayloadKind returns the PayloadKind field value.
func (o *HealthDiagnosisModule) GetPayloadKind() HealthDiagnosisPayloadKind {
	if o == nil {
		var ret HealthDiagnosisPayloadKind
		return ret
	}
	return o.PayloadKind
}

// GetPayloadKindOk returns a tuple with the PayloadKind field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisModule) GetPayloadKindOk() (*HealthDiagnosisPayloadKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadKind, true
}

// SetPayloadKind sets field value.
func (o *HealthDiagnosisModule) SetPayloadKind(v HealthDiagnosisPayloadKind) {
	o.PayloadKind = v
}

// GetSummary returns the Summary field value.
func (o *HealthDiagnosisModule) GetSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisModule) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Summary, true
}

// SetSummary sets field value.
func (o *HealthDiagnosisModule) SetSummary(v string) {
	o.Summary = v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *HealthDiagnosisModule) GetReasonCode() string {
	if o == nil || o.ReasonCode == nil {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisModule) GetReasonCodeOk() (*string, bool) {
	if o == nil || o.ReasonCode == nil {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *HealthDiagnosisModule) HasReasonCode() bool {
	return o != nil && o.ReasonCode != nil
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *HealthDiagnosisModule) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetReadonlyBoundary returns the ReadonlyBoundary field value.
func (o *HealthDiagnosisModule) GetReadonlyBoundary() HealthDiagnosisReadonlyBoundary {
	if o == nil {
		var ret HealthDiagnosisReadonlyBoundary
		return ret
	}
	return o.ReadonlyBoundary
}

// GetReadonlyBoundaryOk returns a tuple with the ReadonlyBoundary field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisModule) GetReadonlyBoundaryOk() (*HealthDiagnosisReadonlyBoundary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadonlyBoundary, true
}

// SetReadonlyBoundary sets field value.
func (o *HealthDiagnosisModule) SetReadonlyBoundary(v HealthDiagnosisReadonlyBoundary) {
	o.ReadonlyBoundary = v
}

// GetPostgresqlRealtimeSessions returns the PostgresqlRealtimeSessions field value if set, zero value otherwise.
func (o *HealthDiagnosisModule) GetPostgresqlRealtimeSessions() PostgresqlRealtimeSessionsPayload {
	if o == nil || o.PostgresqlRealtimeSessions == nil {
		var ret PostgresqlRealtimeSessionsPayload
		return ret
	}
	return *o.PostgresqlRealtimeSessions
}

// GetPostgresqlRealtimeSessionsOk returns a tuple with the PostgresqlRealtimeSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisModule) GetPostgresqlRealtimeSessionsOk() (*PostgresqlRealtimeSessionsPayload, bool) {
	if o == nil || o.PostgresqlRealtimeSessions == nil {
		return nil, false
	}
	return o.PostgresqlRealtimeSessions, true
}

// HasPostgresqlRealtimeSessions returns a boolean if a field has been set.
func (o *HealthDiagnosisModule) HasPostgresqlRealtimeSessions() bool {
	return o != nil && o.PostgresqlRealtimeSessions != nil
}

// SetPostgresqlRealtimeSessions gets a reference to the given PostgresqlRealtimeSessionsPayload and assigns it to the PostgresqlRealtimeSessions field.
func (o *HealthDiagnosisModule) SetPostgresqlRealtimeSessions(v PostgresqlRealtimeSessionsPayload) {
	o.PostgresqlRealtimeSessions = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HealthDiagnosisModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["module"] = o.Module
	toSerialize["state"] = o.State
	toSerialize["payloadKind"] = o.PayloadKind
	toSerialize["summary"] = o.Summary
	if o.ReasonCode != nil {
		toSerialize["reasonCode"] = o.ReasonCode
	}
	toSerialize["readonlyBoundary"] = o.ReadonlyBoundary
	if o.PostgresqlRealtimeSessions != nil {
		toSerialize["postgresqlRealtimeSessions"] = o.PostgresqlRealtimeSessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HealthDiagnosisModule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Module                     *HealthDiagnosisCapabilityModule   `json:"module"`
		State                      *HealthDiagnosisModuleState        `json:"state"`
		PayloadKind                *HealthDiagnosisPayloadKind        `json:"payloadKind"`
		Summary                    *string                            `json:"summary"`
		ReasonCode                 *string                            `json:"reasonCode,omitempty"`
		ReadonlyBoundary           *HealthDiagnosisReadonlyBoundary   `json:"readonlyBoundary"`
		PostgresqlRealtimeSessions *PostgresqlRealtimeSessionsPayload `json:"postgresqlRealtimeSessions,omitempty"`
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
	if all.PayloadKind == nil {
		return fmt.Errorf("required field payloadKind missing")
	}
	if all.Summary == nil {
		return fmt.Errorf("required field summary missing")
	}
	if all.ReadonlyBoundary == nil {
		return fmt.Errorf("required field readonlyBoundary missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"module", "state", "payloadKind", "summary", "reasonCode", "readonlyBoundary", "postgresqlRealtimeSessions"})
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
	if !all.PayloadKind.IsValid() {
		hasInvalidField = true
	} else {
		o.PayloadKind = *all.PayloadKind
	}
	o.Summary = *all.Summary
	o.ReasonCode = all.ReasonCode
	if all.ReadonlyBoundary.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReadonlyBoundary = *all.ReadonlyBoundary
	if all.PostgresqlRealtimeSessions != nil && all.PostgresqlRealtimeSessions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PostgresqlRealtimeSessions = all.PostgresqlRealtimeSessions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
