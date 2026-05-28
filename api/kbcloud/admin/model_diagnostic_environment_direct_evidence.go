// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticEnvironmentDirectEvidence Sprint 28 direct K8s object evidence for the environment diagnosis page.
// This is a read-only object-state summary; it must not be rendered as an
// automatic fix or as proof of a database-internal root cause.
type DiagnosticEnvironmentDirectEvidence struct {
	// Kubernetes resource kind, such as Service, Endpoints, EndpointSlice, PVC, or PV.
	ResourceKind string `json:"resourceKind"`
	// Kubernetes resource name.
	ResourceName string `json:"resourceName"`
	// Kubernetes namespace. Empty for cluster-scoped resources such as PV.
	Namespace string `json:"namespace"`
	// Collection-result axis for the Sprint 24 K8s / Operator / Event explanation layer.
	// with_data means at least one external evidence explanation is available.
	// The remaining states explain why explanations are empty or incomplete without
	// inventing a database-internal conclusion.
	//
	State DiagnosticPeripheralEvidenceState `json:"state"`
	// Stable reason code for unavailable or degraded peripheral evidence.
	ReasonCode DiagnosticPeripheralEvidenceReasonCode `json:"reasonCode"`
	// How this object is related to the diagnosed cluster.
	RelationToCluster string `json:"relationToCluster"`
	// User-facing impact summary derived from safe object status fields.
	Impact string `json:"impact"`
	// Redacted object status summary; does not include Secret data, full endpoint lists, or CSI secret fields.
	EvidenceSummary string `json:"evidenceSummary"`
	// Collection timestamp.
	CollectedAt time.Time `json:"collectedAt"`
	// Report-internal pointers to raw evidence backing this object summary.
	RawEvidenceRefs []string `json:"rawEvidenceRefs"`
	// Safe next actions only.
	NextActions []DiagnosticNextAction `json:"nextActions"`
	// Explicit limits of this object evidence.
	CannotProve []string `json:"cannotProve"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticEnvironmentDirectEvidence instantiates a new DiagnosticEnvironmentDirectEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticEnvironmentDirectEvidence(resourceKind string, resourceName string, namespace string, state DiagnosticPeripheralEvidenceState, reasonCode DiagnosticPeripheralEvidenceReasonCode, relationToCluster string, impact string, evidenceSummary string, collectedAt time.Time, rawEvidenceRefs []string, nextActions []DiagnosticNextAction, cannotProve []string) *DiagnosticEnvironmentDirectEvidence {
	this := DiagnosticEnvironmentDirectEvidence{}
	this.ResourceKind = resourceKind
	this.ResourceName = resourceName
	this.Namespace = namespace
	this.State = state
	this.ReasonCode = reasonCode
	this.RelationToCluster = relationToCluster
	this.Impact = impact
	this.EvidenceSummary = evidenceSummary
	this.CollectedAt = collectedAt
	this.RawEvidenceRefs = rawEvidenceRefs
	this.NextActions = nextActions
	this.CannotProve = cannotProve
	return &this
}

// NewDiagnosticEnvironmentDirectEvidenceWithDefaults instantiates a new DiagnosticEnvironmentDirectEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticEnvironmentDirectEvidenceWithDefaults() *DiagnosticEnvironmentDirectEvidence {
	this := DiagnosticEnvironmentDirectEvidence{}
	return &this
}

// GetResourceKind returns the ResourceKind field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetResourceKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceKind
}

// GetResourceKindOk returns a tuple with the ResourceKind field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetResourceKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKind, true
}

// SetResourceKind sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetResourceKind(v string) {
	o.ResourceKind = v
}

// GetResourceName returns the ResourceName field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetResourceName(v string) {
	o.ResourceName = v
}

// GetNamespace returns the Namespace field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetNamespace(v string) {
	o.Namespace = v
}

// GetState returns the State field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetState() DiagnosticPeripheralEvidenceState {
	if o == nil {
		var ret DiagnosticPeripheralEvidenceState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetStateOk() (*DiagnosticPeripheralEvidenceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetState(v DiagnosticPeripheralEvidenceState) {
	o.State = v
}

// GetReasonCode returns the ReasonCode field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetReasonCode() DiagnosticPeripheralEvidenceReasonCode {
	if o == nil {
		var ret DiagnosticPeripheralEvidenceReasonCode
		return ret
	}
	return o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetReasonCodeOk() (*DiagnosticPeripheralEvidenceReasonCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonCode, true
}

// SetReasonCode sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetReasonCode(v DiagnosticPeripheralEvidenceReasonCode) {
	o.ReasonCode = v
}

// GetRelationToCluster returns the RelationToCluster field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetRelationToCluster() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RelationToCluster
}

// GetRelationToClusterOk returns a tuple with the RelationToCluster field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetRelationToClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationToCluster, true
}

// SetRelationToCluster sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetRelationToCluster(v string) {
	o.RelationToCluster = v
}

// GetImpact returns the Impact field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetImpact() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Impact
}

// GetImpactOk returns a tuple with the Impact field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetImpactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Impact, true
}

// SetImpact sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetImpact(v string) {
	o.Impact = v
}

// GetEvidenceSummary returns the EvidenceSummary field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetEvidenceSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvidenceSummary
}

// GetEvidenceSummaryOk returns a tuple with the EvidenceSummary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetEvidenceSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvidenceSummary, true
}

// SetEvidenceSummary sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetEvidenceSummary(v string) {
	o.EvidenceSummary = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// GetRawEvidenceRefs returns the RawEvidenceRefs field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetRawEvidenceRefs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RawEvidenceRefs
}

// GetRawEvidenceRefsOk returns a tuple with the RawEvidenceRefs field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetRawEvidenceRefsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawEvidenceRefs, true
}

// SetRawEvidenceRefs sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetRawEvidenceRefs(v []string) {
	o.RawEvidenceRefs = v
}

// GetNextActions returns the NextActions field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetNextActions() []DiagnosticNextAction {
	if o == nil {
		var ret []DiagnosticNextAction
		return ret
	}
	return o.NextActions
}

// GetNextActionsOk returns a tuple with the NextActions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetNextActionsOk() (*[]DiagnosticNextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextActions, true
}

// SetNextActions sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetNextActions(v []DiagnosticNextAction) {
	o.NextActions = v
}

// GetCannotProve returns the CannotProve field value.
func (o *DiagnosticEnvironmentDirectEvidence) GetCannotProve() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CannotProve
}

// GetCannotProveOk returns a tuple with the CannotProve field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEnvironmentDirectEvidence) GetCannotProveOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CannotProve, true
}

// SetCannotProve sets field value.
func (o *DiagnosticEnvironmentDirectEvidence) SetCannotProve(v []string) {
	o.CannotProve = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticEnvironmentDirectEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["resourceKind"] = o.ResourceKind
	toSerialize["resourceName"] = o.ResourceName
	toSerialize["namespace"] = o.Namespace
	toSerialize["state"] = o.State
	toSerialize["reasonCode"] = o.ReasonCode
	toSerialize["relationToCluster"] = o.RelationToCluster
	toSerialize["impact"] = o.Impact
	toSerialize["evidenceSummary"] = o.EvidenceSummary
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["rawEvidenceRefs"] = o.RawEvidenceRefs
	toSerialize["nextActions"] = o.NextActions
	toSerialize["cannotProve"] = o.CannotProve

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticEnvironmentDirectEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ResourceKind      *string                                 `json:"resourceKind"`
		ResourceName      *string                                 `json:"resourceName"`
		Namespace         *string                                 `json:"namespace"`
		State             *DiagnosticPeripheralEvidenceState      `json:"state"`
		ReasonCode        *DiagnosticPeripheralEvidenceReasonCode `json:"reasonCode"`
		RelationToCluster *string                                 `json:"relationToCluster"`
		Impact            *string                                 `json:"impact"`
		EvidenceSummary   *string                                 `json:"evidenceSummary"`
		CollectedAt       *time.Time                              `json:"collectedAt"`
		RawEvidenceRefs   *[]string                               `json:"rawEvidenceRefs"`
		NextActions       *[]DiagnosticNextAction                 `json:"nextActions"`
		CannotProve       *[]string                               `json:"cannotProve"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.ResourceKind == nil {
		return fmt.Errorf("required field resourceKind missing")
	}
	if all.ResourceName == nil {
		return fmt.Errorf("required field resourceName missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	if all.State == nil {
		return fmt.Errorf("required field state missing")
	}
	if all.ReasonCode == nil {
		return fmt.Errorf("required field reasonCode missing")
	}
	if all.RelationToCluster == nil {
		return fmt.Errorf("required field relationToCluster missing")
	}
	if all.Impact == nil {
		return fmt.Errorf("required field impact missing")
	}
	if all.EvidenceSummary == nil {
		return fmt.Errorf("required field evidenceSummary missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.RawEvidenceRefs == nil {
		return fmt.Errorf("required field rawEvidenceRefs missing")
	}
	if all.NextActions == nil {
		return fmt.Errorf("required field nextActions missing")
	}
	if all.CannotProve == nil {
		return fmt.Errorf("required field cannotProve missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"resourceKind", "resourceName", "namespace", "state", "reasonCode", "relationToCluster", "impact", "evidenceSummary", "collectedAt", "rawEvidenceRefs", "nextActions", "cannotProve"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ResourceKind = *all.ResourceKind
	o.ResourceName = *all.ResourceName
	o.Namespace = *all.Namespace
	if !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = *all.State
	}
	if !all.ReasonCode.IsValid() {
		hasInvalidField = true
	} else {
		o.ReasonCode = *all.ReasonCode
	}
	o.RelationToCluster = *all.RelationToCluster
	o.Impact = *all.Impact
	o.EvidenceSummary = *all.EvidenceSummary
	o.CollectedAt = *all.CollectedAt
	o.RawEvidenceRefs = *all.RawEvidenceRefs
	o.NextActions = *all.NextActions
	o.CannotProve = *all.CannotProve

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
