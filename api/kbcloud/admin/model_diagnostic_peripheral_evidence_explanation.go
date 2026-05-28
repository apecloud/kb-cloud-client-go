// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticPeripheralEvidenceExplanation One user-facing explanation card derived from K8s / Operator / Event evidence.
// It translates redacted raw evidence into impact and safe next actions, but does
// not claim database-internal causes such as slow SQL or lock waits.
type DiagnosticPeripheralEvidenceExplanation struct {
	// Product-facing category for one K8s / Operator / Event explanation.
	Category DiagnosticPeripheralEvidenceCategory `json:"category"`
	// One-sentence conclusion shown on the card.
	Title string `json:"title"`
	// How this external symptom can affect the database cluster.
	Impact string `json:"impact"`
	// Resource connected to one diagnostic check.
	AffectedObject DiagnosticRelatedResource `json:"affectedObject"`
	// Evidence source, such as Event, Pod, PVC, OpsRequest, Job, or OperatorLog.
	EvidenceSource string `json:"evidenceSource"`
	// Redacted evidence summary. Do not expose raw secrets or full sensitive logs.
	EvidenceSummary string `json:"evidenceSummary"`
	// Earliest occurrence represented by this explanation.
	FirstSeen time.Time `json:"firstSeen"`
	// Latest occurrence represented by this explanation.
	LastSeen time.Time `json:"lastSeen"`
	// Number of evidence entries represented by this explanation.
	Count int64 `json:"count"`
	// Likely layer indicated by the external evidence. This is not a database-internal root-cause claim.
	LikelyLayer DiagnosticPeripheralEvidenceLikelyLayer `json:"likelyLayer"`
	// Confidence of the explanation based only on the collected external evidence.
	Confidence  DiagnosticPeripheralEvidenceConfidence `json:"confidence"`
	NextActions []DiagnosticNextAction                 `json:"nextActions"`
	// Explicit limits of this evidence, for example what it cannot prove about database internals.
	CannotProve []string `json:"cannotProve"`
	// Report-internal pointers to raw evidence backing this explanation.
	RawEvidenceRefs []string `json:"rawEvidenceRefs"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticPeripheralEvidenceExplanation instantiates a new DiagnosticPeripheralEvidenceExplanation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticPeripheralEvidenceExplanation(category DiagnosticPeripheralEvidenceCategory, title string, impact string, affectedObject DiagnosticRelatedResource, evidenceSource string, evidenceSummary string, firstSeen time.Time, lastSeen time.Time, count int64, likelyLayer DiagnosticPeripheralEvidenceLikelyLayer, confidence DiagnosticPeripheralEvidenceConfidence, nextActions []DiagnosticNextAction, cannotProve []string, rawEvidenceRefs []string) *DiagnosticPeripheralEvidenceExplanation {
	this := DiagnosticPeripheralEvidenceExplanation{}
	this.Category = category
	this.Title = title
	this.Impact = impact
	this.AffectedObject = affectedObject
	this.EvidenceSource = evidenceSource
	this.EvidenceSummary = evidenceSummary
	this.FirstSeen = firstSeen
	this.LastSeen = lastSeen
	this.Count = count
	this.LikelyLayer = likelyLayer
	this.Confidence = confidence
	this.NextActions = nextActions
	this.CannotProve = cannotProve
	this.RawEvidenceRefs = rawEvidenceRefs
	return &this
}

// NewDiagnosticPeripheralEvidenceExplanationWithDefaults instantiates a new DiagnosticPeripheralEvidenceExplanation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticPeripheralEvidenceExplanationWithDefaults() *DiagnosticPeripheralEvidenceExplanation {
	this := DiagnosticPeripheralEvidenceExplanation{}
	return &this
}

// GetCategory returns the Category field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetCategory() DiagnosticPeripheralEvidenceCategory {
	if o == nil {
		var ret DiagnosticPeripheralEvidenceCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetCategoryOk() (*DiagnosticPeripheralEvidenceCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetCategory(v DiagnosticPeripheralEvidenceCategory) {
	o.Category = v
}

// GetTitle returns the Title field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetTitle(v string) {
	o.Title = v
}

// GetImpact returns the Impact field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetImpact() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Impact
}

// GetImpactOk returns a tuple with the Impact field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetImpactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Impact, true
}

// SetImpact sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetImpact(v string) {
	o.Impact = v
}

// GetAffectedObject returns the AffectedObject field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetAffectedObject() DiagnosticRelatedResource {
	if o == nil {
		var ret DiagnosticRelatedResource
		return ret
	}
	return o.AffectedObject
}

// GetAffectedObjectOk returns a tuple with the AffectedObject field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetAffectedObjectOk() (*DiagnosticRelatedResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AffectedObject, true
}

// SetAffectedObject sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetAffectedObject(v DiagnosticRelatedResource) {
	o.AffectedObject = v
}

// GetEvidenceSource returns the EvidenceSource field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetEvidenceSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvidenceSource
}

// GetEvidenceSourceOk returns a tuple with the EvidenceSource field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetEvidenceSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvidenceSource, true
}

// SetEvidenceSource sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetEvidenceSource(v string) {
	o.EvidenceSource = v
}

// GetEvidenceSummary returns the EvidenceSummary field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetEvidenceSummary() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EvidenceSummary
}

// GetEvidenceSummaryOk returns a tuple with the EvidenceSummary field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetEvidenceSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvidenceSummary, true
}

// SetEvidenceSummary sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetEvidenceSummary(v string) {
	o.EvidenceSummary = v
}

// GetFirstSeen returns the FirstSeen field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetFirstSeen() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetFirstSeenOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstSeen, true
}

// SetFirstSeen sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetFirstSeen(v time.Time) {
	o.FirstSeen = v
}

// GetLastSeen returns the LastSeen field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetLastSeen() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetLastSeenOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSeen, true
}

// SetLastSeen sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetLastSeen(v time.Time) {
	o.LastSeen = v
}

// GetCount returns the Count field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetCount(v int64) {
	o.Count = v
}

// GetLikelyLayer returns the LikelyLayer field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetLikelyLayer() DiagnosticPeripheralEvidenceLikelyLayer {
	if o == nil {
		var ret DiagnosticPeripheralEvidenceLikelyLayer
		return ret
	}
	return o.LikelyLayer
}

// GetLikelyLayerOk returns a tuple with the LikelyLayer field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetLikelyLayerOk() (*DiagnosticPeripheralEvidenceLikelyLayer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LikelyLayer, true
}

// SetLikelyLayer sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetLikelyLayer(v DiagnosticPeripheralEvidenceLikelyLayer) {
	o.LikelyLayer = v
}

// GetConfidence returns the Confidence field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetConfidence() DiagnosticPeripheralEvidenceConfidence {
	if o == nil {
		var ret DiagnosticPeripheralEvidenceConfidence
		return ret
	}
	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetConfidenceOk() (*DiagnosticPeripheralEvidenceConfidence, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetConfidence(v DiagnosticPeripheralEvidenceConfidence) {
	o.Confidence = v
}

// GetNextActions returns the NextActions field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetNextActions() []DiagnosticNextAction {
	if o == nil {
		var ret []DiagnosticNextAction
		return ret
	}
	return o.NextActions
}

// GetNextActionsOk returns a tuple with the NextActions field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetNextActionsOk() (*[]DiagnosticNextAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextActions, true
}

// SetNextActions sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetNextActions(v []DiagnosticNextAction) {
	o.NextActions = v
}

// GetCannotProve returns the CannotProve field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetCannotProve() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CannotProve
}

// GetCannotProveOk returns a tuple with the CannotProve field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetCannotProveOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CannotProve, true
}

// SetCannotProve sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetCannotProve(v []string) {
	o.CannotProve = v
}

// GetRawEvidenceRefs returns the RawEvidenceRefs field value.
func (o *DiagnosticPeripheralEvidenceExplanation) GetRawEvidenceRefs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RawEvidenceRefs
}

// GetRawEvidenceRefsOk returns a tuple with the RawEvidenceRefs field value
// and a boolean to check if the value has been set.
func (o *DiagnosticPeripheralEvidenceExplanation) GetRawEvidenceRefsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawEvidenceRefs, true
}

// SetRawEvidenceRefs sets field value.
func (o *DiagnosticPeripheralEvidenceExplanation) SetRawEvidenceRefs(v []string) {
	o.RawEvidenceRefs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticPeripheralEvidenceExplanation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["title"] = o.Title
	toSerialize["impact"] = o.Impact
	toSerialize["affectedObject"] = o.AffectedObject
	toSerialize["evidenceSource"] = o.EvidenceSource
	toSerialize["evidenceSummary"] = o.EvidenceSummary
	if o.FirstSeen.Nanosecond() == 0 {
		toSerialize["firstSeen"] = o.FirstSeen.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["firstSeen"] = o.FirstSeen.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.LastSeen.Nanosecond() == 0 {
		toSerialize["lastSeen"] = o.LastSeen.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["lastSeen"] = o.LastSeen.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["count"] = o.Count
	toSerialize["likelyLayer"] = o.LikelyLayer
	toSerialize["confidence"] = o.Confidence
	toSerialize["nextActions"] = o.NextActions
	toSerialize["cannotProve"] = o.CannotProve
	toSerialize["rawEvidenceRefs"] = o.RawEvidenceRefs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticPeripheralEvidenceExplanation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category        *DiagnosticPeripheralEvidenceCategory    `json:"category"`
		Title           *string                                  `json:"title"`
		Impact          *string                                  `json:"impact"`
		AffectedObject  *DiagnosticRelatedResource               `json:"affectedObject"`
		EvidenceSource  *string                                  `json:"evidenceSource"`
		EvidenceSummary *string                                  `json:"evidenceSummary"`
		FirstSeen       *time.Time                               `json:"firstSeen"`
		LastSeen        *time.Time                               `json:"lastSeen"`
		Count           *int64                                   `json:"count"`
		LikelyLayer     *DiagnosticPeripheralEvidenceLikelyLayer `json:"likelyLayer"`
		Confidence      *DiagnosticPeripheralEvidenceConfidence  `json:"confidence"`
		NextActions     *[]DiagnosticNextAction                  `json:"nextActions"`
		CannotProve     *[]string                                `json:"cannotProve"`
		RawEvidenceRefs *[]string                                `json:"rawEvidenceRefs"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Impact == nil {
		return fmt.Errorf("required field impact missing")
	}
	if all.AffectedObject == nil {
		return fmt.Errorf("required field affectedObject missing")
	}
	if all.EvidenceSource == nil {
		return fmt.Errorf("required field evidenceSource missing")
	}
	if all.EvidenceSummary == nil {
		return fmt.Errorf("required field evidenceSummary missing")
	}
	if all.FirstSeen == nil {
		return fmt.Errorf("required field firstSeen missing")
	}
	if all.LastSeen == nil {
		return fmt.Errorf("required field lastSeen missing")
	}
	if all.Count == nil {
		return fmt.Errorf("required field count missing")
	}
	if all.LikelyLayer == nil {
		return fmt.Errorf("required field likelyLayer missing")
	}
	if all.Confidence == nil {
		return fmt.Errorf("required field confidence missing")
	}
	if all.NextActions == nil {
		return fmt.Errorf("required field nextActions missing")
	}
	if all.CannotProve == nil {
		return fmt.Errorf("required field cannotProve missing")
	}
	if all.RawEvidenceRefs == nil {
		return fmt.Errorf("required field rawEvidenceRefs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"category", "title", "impact", "affectedObject", "evidenceSource", "evidenceSummary", "firstSeen", "lastSeen", "count", "likelyLayer", "confidence", "nextActions", "cannotProve", "rawEvidenceRefs"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Title = *all.Title
	o.Impact = *all.Impact
	if all.AffectedObject.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AffectedObject = *all.AffectedObject
	o.EvidenceSource = *all.EvidenceSource
	o.EvidenceSummary = *all.EvidenceSummary
	o.FirstSeen = *all.FirstSeen
	o.LastSeen = *all.LastSeen
	o.Count = *all.Count
	if !all.LikelyLayer.IsValid() {
		hasInvalidField = true
	} else {
		o.LikelyLayer = *all.LikelyLayer
	}
	if !all.Confidence.IsValid() {
		hasInvalidField = true
	} else {
		o.Confidence = *all.Confidence
	}
	o.NextActions = *all.NextActions
	o.CannotProve = *all.CannotProve
	o.RawEvidenceRefs = *all.RawEvidenceRefs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
