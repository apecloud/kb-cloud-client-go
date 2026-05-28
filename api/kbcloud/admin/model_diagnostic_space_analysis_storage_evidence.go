// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSpaceAnalysisStorageEvidence PVC/PV capacity evidence kept separate from database object sizes.
type DiagnosticSpaceAnalysisStorageEvidence struct {
	// Kubernetes object kind. Sprint 26 emits PersistentVolumeClaim.
	Kind string `json:"kind"`
	// Kubernetes namespace.
	Namespace string `json:"namespace"`
	// PVC name.
	Name string `json:"name"`
	// PVC phase.
	Phase string `json:"phase"`
	// PVC storageClassName, when present.
	StorageClass *string `json:"storageClass,omitempty"`
	// Bound PV name from spec.volumeName, when present.
	PersistentVolumeName *string `json:"persistentVolumeName,omitempty"`
	// PVC declared or bound capacity in bytes. This is not database used size.
	CapacityBytes int64 `json:"capacityBytes"`
	// Human-readable PVC capacity.
	CapacityHuman string `json:"capacityHuman"`
	// Boundary statement explaining that PVC/PV capacity does not prove one database object is large.
	Message string `json:"message"`
	// Time when PVC/PV evidence was attached to space analysis.
	CollectedAt time.Time `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticSpaceAnalysisStorageEvidence instantiates a new DiagnosticSpaceAnalysisStorageEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticSpaceAnalysisStorageEvidence(kind string, namespace string, name string, phase string, capacityBytes int64, capacityHuman string, message string, collectedAt time.Time) *DiagnosticSpaceAnalysisStorageEvidence {
	this := DiagnosticSpaceAnalysisStorageEvidence{}
	this.Kind = kind
	this.Namespace = namespace
	this.Name = name
	this.Phase = phase
	this.CapacityBytes = capacityBytes
	this.CapacityHuman = capacityHuman
	this.Message = message
	this.CollectedAt = collectedAt
	return &this
}

// NewDiagnosticSpaceAnalysisStorageEvidenceWithDefaults instantiates a new DiagnosticSpaceAnalysisStorageEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticSpaceAnalysisStorageEvidenceWithDefaults() *DiagnosticSpaceAnalysisStorageEvidence {
	this := DiagnosticSpaceAnalysisStorageEvidence{}
	return &this
}

// GetKind returns the Kind field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) SetKind(v string) {
	o.Kind = v
}

// GetNamespace returns the Namespace field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) SetNamespace(v string) {
	o.Namespace = v
}

// GetName returns the Name field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) SetName(v string) {
	o.Name = v
}

// GetPhase returns the Phase field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) SetPhase(v string) {
	o.Phase = v
}

// GetStorageClass returns the StorageClass field value if set, zero value otherwise.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetStorageClass() string {
	if o == nil || o.StorageClass == nil {
		var ret string
		return ret
	}
	return *o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetStorageClassOk() (*string, bool) {
	if o == nil || o.StorageClass == nil {
		return nil, false
	}
	return o.StorageClass, true
}

// HasStorageClass returns a boolean if a field has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) HasStorageClass() bool {
	return o != nil && o.StorageClass != nil
}

// SetStorageClass gets a reference to the given string and assigns it to the StorageClass field.
func (o *DiagnosticSpaceAnalysisStorageEvidence) SetStorageClass(v string) {
	o.StorageClass = &v
}

// GetPersistentVolumeName returns the PersistentVolumeName field value if set, zero value otherwise.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetPersistentVolumeName() string {
	if o == nil || o.PersistentVolumeName == nil {
		var ret string
		return ret
	}
	return *o.PersistentVolumeName
}

// GetPersistentVolumeNameOk returns a tuple with the PersistentVolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetPersistentVolumeNameOk() (*string, bool) {
	if o == nil || o.PersistentVolumeName == nil {
		return nil, false
	}
	return o.PersistentVolumeName, true
}

// HasPersistentVolumeName returns a boolean if a field has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) HasPersistentVolumeName() bool {
	return o != nil && o.PersistentVolumeName != nil
}

// SetPersistentVolumeName gets a reference to the given string and assigns it to the PersistentVolumeName field.
func (o *DiagnosticSpaceAnalysisStorageEvidence) SetPersistentVolumeName(v string) {
	o.PersistentVolumeName = &v
}

// GetCapacityBytes returns the CapacityBytes field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetCapacityBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CapacityBytes
}

// GetCapacityBytesOk returns a tuple with the CapacityBytes field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetCapacityBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityBytes, true
}

// SetCapacityBytes sets field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) SetCapacityBytes(v int64) {
	o.CapacityBytes = v
}

// GetCapacityHuman returns the CapacityHuman field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetCapacityHuman() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CapacityHuman
}

// GetCapacityHumanOk returns a tuple with the CapacityHuman field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetCapacityHumanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityHuman, true
}

// SetCapacityHuman sets field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) SetCapacityHuman(v string) {
	o.CapacityHuman = v
}

// GetMessage returns the Message field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) SetMessage(v string) {
	o.Message = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSpaceAnalysisStorageEvidence) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticSpaceAnalysisStorageEvidence) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticSpaceAnalysisStorageEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["kind"] = o.Kind
	toSerialize["namespace"] = o.Namespace
	toSerialize["name"] = o.Name
	toSerialize["phase"] = o.Phase
	if o.StorageClass != nil {
		toSerialize["storageClass"] = o.StorageClass
	}
	if o.PersistentVolumeName != nil {
		toSerialize["persistentVolumeName"] = o.PersistentVolumeName
	}
	toSerialize["capacityBytes"] = o.CapacityBytes
	toSerialize["capacityHuman"] = o.CapacityHuman
	toSerialize["message"] = o.Message
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticSpaceAnalysisStorageEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Kind                 *string    `json:"kind"`
		Namespace            *string    `json:"namespace"`
		Name                 *string    `json:"name"`
		Phase                *string    `json:"phase"`
		StorageClass         *string    `json:"storageClass,omitempty"`
		PersistentVolumeName *string    `json:"persistentVolumeName,omitempty"`
		CapacityBytes        *int64     `json:"capacityBytes"`
		CapacityHuman        *string    `json:"capacityHuman"`
		Message              *string    `json:"message"`
		CollectedAt          *time.Time `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Phase == nil {
		return fmt.Errorf("required field phase missing")
	}
	if all.CapacityBytes == nil {
		return fmt.Errorf("required field capacityBytes missing")
	}
	if all.CapacityHuman == nil {
		return fmt.Errorf("required field capacityHuman missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"kind", "namespace", "name", "phase", "storageClass", "persistentVolumeName", "capacityBytes", "capacityHuman", "message", "collectedAt"})
	} else {
		return err
	}
	o.Kind = *all.Kind
	o.Namespace = *all.Namespace
	o.Name = *all.Name
	o.Phase = *all.Phase
	o.StorageClass = all.StorageClass
	o.PersistentVolumeName = all.PersistentVolumeName
	o.CapacityBytes = *all.CapacityBytes
	o.CapacityHuman = *all.CapacityHuman
	o.Message = *all.Message
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
