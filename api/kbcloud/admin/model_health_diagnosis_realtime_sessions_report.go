// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HealthDiagnosisRealtimeSessionsReport struct {
	Engine                 string                      `json:"engine"`
	DiagnosticCapabilities []HealthDiagnosisCapability `json:"diagnosticCapabilities"`
	Module                 HealthDiagnosisModule       `json:"module"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHealthDiagnosisRealtimeSessionsReport instantiates a new HealthDiagnosisRealtimeSessionsReport object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHealthDiagnosisRealtimeSessionsReport(engine string, diagnosticCapabilities []HealthDiagnosisCapability, module HealthDiagnosisModule) *HealthDiagnosisRealtimeSessionsReport {
	this := HealthDiagnosisRealtimeSessionsReport{}
	this.Engine = engine
	this.DiagnosticCapabilities = diagnosticCapabilities
	this.Module = module
	return &this
}

// NewHealthDiagnosisRealtimeSessionsReportWithDefaults instantiates a new HealthDiagnosisRealtimeSessionsReport object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHealthDiagnosisRealtimeSessionsReportWithDefaults() *HealthDiagnosisRealtimeSessionsReport {
	this := HealthDiagnosisRealtimeSessionsReport{}
	return &this
}

// GetEngine returns the Engine field value.
func (o *HealthDiagnosisRealtimeSessionsReport) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisRealtimeSessionsReport) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value.
func (o *HealthDiagnosisRealtimeSessionsReport) SetEngine(v string) {
	o.Engine = v
}

// GetDiagnosticCapabilities returns the DiagnosticCapabilities field value.
func (o *HealthDiagnosisRealtimeSessionsReport) GetDiagnosticCapabilities() []HealthDiagnosisCapability {
	if o == nil {
		var ret []HealthDiagnosisCapability
		return ret
	}
	return o.DiagnosticCapabilities
}

// GetDiagnosticCapabilitiesOk returns a tuple with the DiagnosticCapabilities field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisRealtimeSessionsReport) GetDiagnosticCapabilitiesOk() (*[]HealthDiagnosisCapability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiagnosticCapabilities, true
}

// SetDiagnosticCapabilities sets field value.
func (o *HealthDiagnosisRealtimeSessionsReport) SetDiagnosticCapabilities(v []HealthDiagnosisCapability) {
	o.DiagnosticCapabilities = v
}

// GetModule returns the Module field value.
func (o *HealthDiagnosisRealtimeSessionsReport) GetModule() HealthDiagnosisModule {
	if o == nil {
		var ret HealthDiagnosisModule
		return ret
	}
	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisRealtimeSessionsReport) GetModuleOk() (*HealthDiagnosisModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value.
func (o *HealthDiagnosisRealtimeSessionsReport) SetModule(v HealthDiagnosisModule) {
	o.Module = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HealthDiagnosisRealtimeSessionsReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["engine"] = o.Engine
	toSerialize["diagnosticCapabilities"] = o.DiagnosticCapabilities
	toSerialize["module"] = o.Module

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HealthDiagnosisRealtimeSessionsReport) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Engine                 *string                      `json:"engine"`
		DiagnosticCapabilities *[]HealthDiagnosisCapability `json:"diagnosticCapabilities"`
		Module                 *HealthDiagnosisModule       `json:"module"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Engine == nil {
		return fmt.Errorf("required field engine missing")
	}
	if all.DiagnosticCapabilities == nil {
		return fmt.Errorf("required field diagnosticCapabilities missing")
	}
	if all.Module == nil {
		return fmt.Errorf("required field module missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"engine", "diagnosticCapabilities", "module"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Engine = *all.Engine
	o.DiagnosticCapabilities = *all.DiagnosticCapabilities
	if all.Module.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Module = *all.Module

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
