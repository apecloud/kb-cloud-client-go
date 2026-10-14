// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ObjectPrivilegeSnapshot struct {
	Target                    ObjectPrivilegeTarget `json:"target"`
	SupportedPrivileges       []string              `json:"supportedPrivileges"`
	SupportedColumnPrivileges []string              `json:"supportedColumnPrivileges"`
	// Native grants to this account at this exact object scope. Does not include inherited grants or prove executable access.
	DirectPrivileges []ObjectPrivilegeGrant `json:"directPrivileges"`
	// Native PostgreSQL checks when available. False for PostgreSQL large objects before 18 and MySQL (active roles depend on the session). False or omitted effectivePrivileges never means no access.
	EffectivePrivilegesKnown bool `json:"effectivePrivilegesKnown"`
	// PostgreSQL native has_*_privilege checks, including PUBLIC and inherited permissions. Other prerequisites, RLS and session SET ROLE can affect actual execution.
	EffectivePrivileges []ObjectPrivilegeGrant `json:"effectivePrivileges,omitempty"`
	// Native behavior and scope overlap explanations, including preset overlap and partial revokes.
	Warnings []string               `json:"warnings,omitempty"`
	Subject  ObjectPrivilegeSubject `json:"subject"`
	// Versioned namespaced JSON extensions. Response keys unknown to a client may be ignored. Change keys must be registered and strictly decoded by the current engine; PG/MySQL currently accept no change extensions.
	Extensions         map[string]interface{}             `json:"extensions,omitempty"`
	PrerequisiteChecks []ObjectPrivilegePrerequisiteCheck `json:"prerequisiteChecks,omitempty"`
	Evidence           []ObjectPrivilegeEvidence          `json:"evidence,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObjectPrivilegeSnapshot instantiates a new ObjectPrivilegeSnapshot object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObjectPrivilegeSnapshot(target ObjectPrivilegeTarget, supportedPrivileges []string, supportedColumnPrivileges []string, directPrivileges []ObjectPrivilegeGrant, effectivePrivilegesKnown bool, subject ObjectPrivilegeSubject) *ObjectPrivilegeSnapshot {
	this := ObjectPrivilegeSnapshot{}
	this.Target = target
	this.SupportedPrivileges = supportedPrivileges
	this.SupportedColumnPrivileges = supportedColumnPrivileges
	this.DirectPrivileges = directPrivileges
	this.EffectivePrivilegesKnown = effectivePrivilegesKnown
	this.Subject = subject
	return &this
}

// NewObjectPrivilegeSnapshotWithDefaults instantiates a new ObjectPrivilegeSnapshot object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObjectPrivilegeSnapshotWithDefaults() *ObjectPrivilegeSnapshot {
	this := ObjectPrivilegeSnapshot{}
	return &this
}

// GetTarget returns the Target field value.
func (o *ObjectPrivilegeSnapshot) GetTarget() ObjectPrivilegeTarget {
	if o == nil {
		var ret ObjectPrivilegeTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetTargetOk() (*ObjectPrivilegeTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *ObjectPrivilegeSnapshot) SetTarget(v ObjectPrivilegeTarget) {
	o.Target = v
}

// GetSupportedPrivileges returns the SupportedPrivileges field value.
func (o *ObjectPrivilegeSnapshot) GetSupportedPrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedPrivileges
}

// GetSupportedPrivilegesOk returns a tuple with the SupportedPrivileges field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetSupportedPrivilegesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedPrivileges, true
}

// SetSupportedPrivileges sets field value.
func (o *ObjectPrivilegeSnapshot) SetSupportedPrivileges(v []string) {
	o.SupportedPrivileges = v
}

// GetSupportedColumnPrivileges returns the SupportedColumnPrivileges field value.
func (o *ObjectPrivilegeSnapshot) GetSupportedColumnPrivileges() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SupportedColumnPrivileges
}

// GetSupportedColumnPrivilegesOk returns a tuple with the SupportedColumnPrivileges field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetSupportedColumnPrivilegesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedColumnPrivileges, true
}

// SetSupportedColumnPrivileges sets field value.
func (o *ObjectPrivilegeSnapshot) SetSupportedColumnPrivileges(v []string) {
	o.SupportedColumnPrivileges = v
}

// GetDirectPrivileges returns the DirectPrivileges field value.
func (o *ObjectPrivilegeSnapshot) GetDirectPrivileges() []ObjectPrivilegeGrant {
	if o == nil {
		var ret []ObjectPrivilegeGrant
		return ret
	}
	return o.DirectPrivileges
}

// GetDirectPrivilegesOk returns a tuple with the DirectPrivileges field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetDirectPrivilegesOk() (*[]ObjectPrivilegeGrant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DirectPrivileges, true
}

// SetDirectPrivileges sets field value.
func (o *ObjectPrivilegeSnapshot) SetDirectPrivileges(v []ObjectPrivilegeGrant) {
	o.DirectPrivileges = v
}

// GetEffectivePrivilegesKnown returns the EffectivePrivilegesKnown field value.
func (o *ObjectPrivilegeSnapshot) GetEffectivePrivilegesKnown() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.EffectivePrivilegesKnown
}

// GetEffectivePrivilegesKnownOk returns a tuple with the EffectivePrivilegesKnown field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetEffectivePrivilegesKnownOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectivePrivilegesKnown, true
}

// SetEffectivePrivilegesKnown sets field value.
func (o *ObjectPrivilegeSnapshot) SetEffectivePrivilegesKnown(v bool) {
	o.EffectivePrivilegesKnown = v
}

// GetEffectivePrivileges returns the EffectivePrivileges field value if set, zero value otherwise.
func (o *ObjectPrivilegeSnapshot) GetEffectivePrivileges() []ObjectPrivilegeGrant {
	if o == nil || o.EffectivePrivileges == nil {
		var ret []ObjectPrivilegeGrant
		return ret
	}
	return o.EffectivePrivileges
}

// GetEffectivePrivilegesOk returns a tuple with the EffectivePrivileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetEffectivePrivilegesOk() (*[]ObjectPrivilegeGrant, bool) {
	if o == nil || o.EffectivePrivileges == nil {
		return nil, false
	}
	return &o.EffectivePrivileges, true
}

// HasEffectivePrivileges returns a boolean if a field has been set.
func (o *ObjectPrivilegeSnapshot) HasEffectivePrivileges() bool {
	return o != nil && o.EffectivePrivileges != nil
}

// SetEffectivePrivileges gets a reference to the given []ObjectPrivilegeGrant and assigns it to the EffectivePrivileges field.
func (o *ObjectPrivilegeSnapshot) SetEffectivePrivileges(v []ObjectPrivilegeGrant) {
	o.EffectivePrivileges = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ObjectPrivilegeSnapshot) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetWarningsOk() (*[]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ObjectPrivilegeSnapshot) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *ObjectPrivilegeSnapshot) SetWarnings(v []string) {
	o.Warnings = v
}

// GetSubject returns the Subject field value.
func (o *ObjectPrivilegeSnapshot) GetSubject() ObjectPrivilegeSubject {
	if o == nil {
		var ret ObjectPrivilegeSubject
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetSubjectOk() (*ObjectPrivilegeSubject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value.
func (o *ObjectPrivilegeSnapshot) SetSubject(v ObjectPrivilegeSubject) {
	o.Subject = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *ObjectPrivilegeSnapshot) GetExtensions() map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetExtensionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return &o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *ObjectPrivilegeSnapshot) HasExtensions() bool {
	return o != nil && o.Extensions != nil
}

// SetExtensions gets a reference to the given map[string]interface{} and assigns it to the Extensions field.
func (o *ObjectPrivilegeSnapshot) SetExtensions(v map[string]interface{}) {
	o.Extensions = v
}

// GetPrerequisiteChecks returns the PrerequisiteChecks field value if set, zero value otherwise.
func (o *ObjectPrivilegeSnapshot) GetPrerequisiteChecks() []ObjectPrivilegePrerequisiteCheck {
	if o == nil || o.PrerequisiteChecks == nil {
		var ret []ObjectPrivilegePrerequisiteCheck
		return ret
	}
	return o.PrerequisiteChecks
}

// GetPrerequisiteChecksOk returns a tuple with the PrerequisiteChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetPrerequisiteChecksOk() (*[]ObjectPrivilegePrerequisiteCheck, bool) {
	if o == nil || o.PrerequisiteChecks == nil {
		return nil, false
	}
	return &o.PrerequisiteChecks, true
}

// HasPrerequisiteChecks returns a boolean if a field has been set.
func (o *ObjectPrivilegeSnapshot) HasPrerequisiteChecks() bool {
	return o != nil && o.PrerequisiteChecks != nil
}

// SetPrerequisiteChecks gets a reference to the given []ObjectPrivilegePrerequisiteCheck and assigns it to the PrerequisiteChecks field.
func (o *ObjectPrivilegeSnapshot) SetPrerequisiteChecks(v []ObjectPrivilegePrerequisiteCheck) {
	o.PrerequisiteChecks = v
}

// GetEvidence returns the Evidence field value if set, zero value otherwise.
func (o *ObjectPrivilegeSnapshot) GetEvidence() []ObjectPrivilegeEvidence {
	if o == nil || o.Evidence == nil {
		var ret []ObjectPrivilegeEvidence
		return ret
	}
	return o.Evidence
}

// GetEvidenceOk returns a tuple with the Evidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPrivilegeSnapshot) GetEvidenceOk() (*[]ObjectPrivilegeEvidence, bool) {
	if o == nil || o.Evidence == nil {
		return nil, false
	}
	return &o.Evidence, true
}

// HasEvidence returns a boolean if a field has been set.
func (o *ObjectPrivilegeSnapshot) HasEvidence() bool {
	return o != nil && o.Evidence != nil
}

// SetEvidence gets a reference to the given []ObjectPrivilegeEvidence and assigns it to the Evidence field.
func (o *ObjectPrivilegeSnapshot) SetEvidence(v []ObjectPrivilegeEvidence) {
	o.Evidence = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObjectPrivilegeSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["target"] = o.Target
	toSerialize["supportedPrivileges"] = o.SupportedPrivileges
	toSerialize["supportedColumnPrivileges"] = o.SupportedColumnPrivileges
	toSerialize["directPrivileges"] = o.DirectPrivileges
	toSerialize["effectivePrivilegesKnown"] = o.EffectivePrivilegesKnown
	if o.EffectivePrivileges != nil {
		toSerialize["effectivePrivileges"] = o.EffectivePrivileges
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	toSerialize["subject"] = o.Subject
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.PrerequisiteChecks != nil {
		toSerialize["prerequisiteChecks"] = o.PrerequisiteChecks
	}
	if o.Evidence != nil {
		toSerialize["evidence"] = o.Evidence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObjectPrivilegeSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Target                    *ObjectPrivilegeTarget             `json:"target"`
		SupportedPrivileges       *[]string                          `json:"supportedPrivileges"`
		SupportedColumnPrivileges *[]string                          `json:"supportedColumnPrivileges"`
		DirectPrivileges          *[]ObjectPrivilegeGrant            `json:"directPrivileges"`
		EffectivePrivilegesKnown  *bool                              `json:"effectivePrivilegesKnown"`
		EffectivePrivileges       []ObjectPrivilegeGrant             `json:"effectivePrivileges,omitempty"`
		Warnings                  []string                           `json:"warnings,omitempty"`
		Subject                   *ObjectPrivilegeSubject            `json:"subject"`
		Extensions                map[string]interface{}             `json:"extensions,omitempty"`
		PrerequisiteChecks        []ObjectPrivilegePrerequisiteCheck `json:"prerequisiteChecks,omitempty"`
		Evidence                  []ObjectPrivilegeEvidence          `json:"evidence,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.SupportedPrivileges == nil {
		return fmt.Errorf("required field supportedPrivileges missing")
	}
	if all.SupportedColumnPrivileges == nil {
		return fmt.Errorf("required field supportedColumnPrivileges missing")
	}
	if all.DirectPrivileges == nil {
		return fmt.Errorf("required field directPrivileges missing")
	}
	if all.EffectivePrivilegesKnown == nil {
		return fmt.Errorf("required field effectivePrivilegesKnown missing")
	}
	if all.Subject == nil {
		return fmt.Errorf("required field subject missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"target", "supportedPrivileges", "supportedColumnPrivileges", "directPrivileges", "effectivePrivilegesKnown", "effectivePrivileges", "warnings", "subject", "extensions", "prerequisiteChecks", "evidence"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target
	o.SupportedPrivileges = *all.SupportedPrivileges
	o.SupportedColumnPrivileges = *all.SupportedColumnPrivileges
	o.DirectPrivileges = *all.DirectPrivileges
	o.EffectivePrivilegesKnown = *all.EffectivePrivilegesKnown
	o.EffectivePrivileges = all.EffectivePrivileges
	o.Warnings = all.Warnings
	if all.Subject.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Subject = *all.Subject
	o.Extensions = all.Extensions
	o.PrerequisiteChecks = all.PrerequisiteChecks
	o.Evidence = all.Evidence

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
