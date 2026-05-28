// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticNextAction Structured next action for one diagnostic check or top issue.
type DiagnosticNextAction struct {
	// Short imperative title shown to the user, such as "查看关联对象和控制面日志".
	Title string `json:"title"`
	// Safe operation type for a diagnostic next action. Console uses this to decide whether to render a clickable action, a copy action, a rerun entry, or a manual suggestion.
	Type DiagnosticNextActionType `json:"type"`
	// Target surface this next action points to. Matches the Sprint 7 v0.1 product contract.
	Target DiagnosticNextActionTarget `json:"target"`
	// One-sentence rationale that explains why this action is suggested.
	Reason string `json:"reason"`
	// Suggested first responder for this action.
	OwnerHint string `json:"ownerHint"`
	// Optional report-internal evidence path or related resource identifier that the action points to.
	EvidencePath *string `json:"evidencePath,omitempty"`
	// Optional hint of the Console entry to open, such as cluster_console or log_drawer. Not enforced in Sprint 7.
	LinkType *string `json:"linkType,omitempty"`
	// Stable UI target identifier for clickable actions, such as dbPerformance.slowSql.detail or diagnostic.rerun.
	TargetId *string `json:"targetID,omitempty"`
	// Stable report evidence anchor for jump_to_evidence actions.
	EvidenceAnchor *string `json:"evidenceAnchor,omitempty"`
	// Safe, redacted text payload that Console may copy for DBA/SRE handoff. It must not contain passwords, tokens, cookies, kubeconfig, or full sensitive SQL text.
	CopyPayload *DiagnosticNextActionCopyPayload `json:"copyPayload,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticNextAction instantiates a new DiagnosticNextAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticNextAction(title string, typeVar DiagnosticNextActionType, target DiagnosticNextActionTarget, reason string, ownerHint string) *DiagnosticNextAction {
	this := DiagnosticNextAction{}
	this.Title = title
	this.Type = typeVar
	this.Target = target
	this.Reason = reason
	this.OwnerHint = ownerHint
	return &this
}

// NewDiagnosticNextActionWithDefaults instantiates a new DiagnosticNextAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticNextActionWithDefaults() *DiagnosticNextAction {
	this := DiagnosticNextAction{}
	return &this
}

// GetTitle returns the Title field value.
func (o *DiagnosticNextAction) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DiagnosticNextAction) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *DiagnosticNextAction) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value.
func (o *DiagnosticNextAction) GetType() DiagnosticNextActionType {
	if o == nil {
		var ret DiagnosticNextActionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DiagnosticNextAction) GetTypeOk() (*DiagnosticNextActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DiagnosticNextAction) SetType(v DiagnosticNextActionType) {
	o.Type = v
}

// GetTarget returns the Target field value.
func (o *DiagnosticNextAction) GetTarget() DiagnosticNextActionTarget {
	if o == nil {
		var ret DiagnosticNextActionTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *DiagnosticNextAction) GetTargetOk() (*DiagnosticNextActionTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *DiagnosticNextAction) SetTarget(v DiagnosticNextActionTarget) {
	o.Target = v
}

// GetReason returns the Reason field value.
func (o *DiagnosticNextAction) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DiagnosticNextAction) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *DiagnosticNextAction) SetReason(v string) {
	o.Reason = v
}

// GetOwnerHint returns the OwnerHint field value.
func (o *DiagnosticNextAction) GetOwnerHint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OwnerHint
}

// GetOwnerHintOk returns a tuple with the OwnerHint field value
// and a boolean to check if the value has been set.
func (o *DiagnosticNextAction) GetOwnerHintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerHint, true
}

// SetOwnerHint sets field value.
func (o *DiagnosticNextAction) SetOwnerHint(v string) {
	o.OwnerHint = v
}

// GetEvidencePath returns the EvidencePath field value if set, zero value otherwise.
func (o *DiagnosticNextAction) GetEvidencePath() string {
	if o == nil || o.EvidencePath == nil {
		var ret string
		return ret
	}
	return *o.EvidencePath
}

// GetEvidencePathOk returns a tuple with the EvidencePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticNextAction) GetEvidencePathOk() (*string, bool) {
	if o == nil || o.EvidencePath == nil {
		return nil, false
	}
	return o.EvidencePath, true
}

// HasEvidencePath returns a boolean if a field has been set.
func (o *DiagnosticNextAction) HasEvidencePath() bool {
	return o != nil && o.EvidencePath != nil
}

// SetEvidencePath gets a reference to the given string and assigns it to the EvidencePath field.
func (o *DiagnosticNextAction) SetEvidencePath(v string) {
	o.EvidencePath = &v
}

// GetLinkType returns the LinkType field value if set, zero value otherwise.
func (o *DiagnosticNextAction) GetLinkType() string {
	if o == nil || o.LinkType == nil {
		var ret string
		return ret
	}
	return *o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticNextAction) GetLinkTypeOk() (*string, bool) {
	if o == nil || o.LinkType == nil {
		return nil, false
	}
	return o.LinkType, true
}

// HasLinkType returns a boolean if a field has been set.
func (o *DiagnosticNextAction) HasLinkType() bool {
	return o != nil && o.LinkType != nil
}

// SetLinkType gets a reference to the given string and assigns it to the LinkType field.
func (o *DiagnosticNextAction) SetLinkType(v string) {
	o.LinkType = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *DiagnosticNextAction) GetTargetId() string {
	if o == nil || o.TargetId == nil {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticNextAction) GetTargetIdOk() (*string, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *DiagnosticNextAction) HasTargetId() bool {
	return o != nil && o.TargetId != nil
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *DiagnosticNextAction) SetTargetId(v string) {
	o.TargetId = &v
}

// GetEvidenceAnchor returns the EvidenceAnchor field value if set, zero value otherwise.
func (o *DiagnosticNextAction) GetEvidenceAnchor() string {
	if o == nil || o.EvidenceAnchor == nil {
		var ret string
		return ret
	}
	return *o.EvidenceAnchor
}

// GetEvidenceAnchorOk returns a tuple with the EvidenceAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticNextAction) GetEvidenceAnchorOk() (*string, bool) {
	if o == nil || o.EvidenceAnchor == nil {
		return nil, false
	}
	return o.EvidenceAnchor, true
}

// HasEvidenceAnchor returns a boolean if a field has been set.
func (o *DiagnosticNextAction) HasEvidenceAnchor() bool {
	return o != nil && o.EvidenceAnchor != nil
}

// SetEvidenceAnchor gets a reference to the given string and assigns it to the EvidenceAnchor field.
func (o *DiagnosticNextAction) SetEvidenceAnchor(v string) {
	o.EvidenceAnchor = &v
}

// GetCopyPayload returns the CopyPayload field value if set, zero value otherwise.
func (o *DiagnosticNextAction) GetCopyPayload() DiagnosticNextActionCopyPayload {
	if o == nil || o.CopyPayload == nil {
		var ret DiagnosticNextActionCopyPayload
		return ret
	}
	return *o.CopyPayload
}

// GetCopyPayloadOk returns a tuple with the CopyPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticNextAction) GetCopyPayloadOk() (*DiagnosticNextActionCopyPayload, bool) {
	if o == nil || o.CopyPayload == nil {
		return nil, false
	}
	return o.CopyPayload, true
}

// HasCopyPayload returns a boolean if a field has been set.
func (o *DiagnosticNextAction) HasCopyPayload() bool {
	return o != nil && o.CopyPayload != nil
}

// SetCopyPayload gets a reference to the given DiagnosticNextActionCopyPayload and assigns it to the CopyPayload field.
func (o *DiagnosticNextAction) SetCopyPayload(v DiagnosticNextActionCopyPayload) {
	o.CopyPayload = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticNextAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	toSerialize["target"] = o.Target
	toSerialize["reason"] = o.Reason
	toSerialize["ownerHint"] = o.OwnerHint
	if o.EvidencePath != nil {
		toSerialize["evidencePath"] = o.EvidencePath
	}
	if o.LinkType != nil {
		toSerialize["linkType"] = o.LinkType
	}
	if o.TargetId != nil {
		toSerialize["targetID"] = o.TargetId
	}
	if o.EvidenceAnchor != nil {
		toSerialize["evidenceAnchor"] = o.EvidenceAnchor
	}
	if o.CopyPayload != nil {
		toSerialize["copyPayload"] = o.CopyPayload
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticNextAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Title          *string                          `json:"title"`
		Type           *DiagnosticNextActionType        `json:"type"`
		Target         *DiagnosticNextActionTarget      `json:"target"`
		Reason         *string                          `json:"reason"`
		OwnerHint      *string                          `json:"ownerHint"`
		EvidencePath   *string                          `json:"evidencePath,omitempty"`
		LinkType       *string                          `json:"linkType,omitempty"`
		TargetId       *string                          `json:"targetID,omitempty"`
		EvidenceAnchor *string                          `json:"evidenceAnchor,omitempty"`
		CopyPayload    *DiagnosticNextActionCopyPayload `json:"copyPayload,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if all.OwnerHint == nil {
		return fmt.Errorf("required field ownerHint missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"title", "type", "target", "reason", "ownerHint", "evidencePath", "linkType", "targetID", "evidenceAnchor", "copyPayload"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Title = *all.Title
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if !all.Target.IsValid() {
		hasInvalidField = true
	} else {
		o.Target = *all.Target
	}
	o.Reason = *all.Reason
	o.OwnerHint = *all.OwnerHint
	o.EvidencePath = all.EvidencePath
	o.LinkType = all.LinkType
	o.TargetId = all.TargetId
	o.EvidenceAnchor = all.EvidenceAnchor
	if all.CopyPayload != nil && all.CopyPayload.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CopyPayload = all.CopyPayload

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
