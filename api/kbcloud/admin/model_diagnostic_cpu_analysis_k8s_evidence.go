// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticCPUAnalysisK8sEvidence Kubernetes resource/event evidence kept separate from DB activity and SQL pressure evidence.
type DiagnosticCPUAnalysisK8sEvidence struct {
	// Evidence source, such as PodMetrics, PodSpec, Event, or ResourceSpec.
	Source string `json:"source"`
	// Kubernetes namespace.
	Namespace string `json:"namespace"`
	// Kubernetes object name.
	ObjectName string `json:"objectName"`
	// Redacted reason or category.
	Reason string `json:"reason"`
	// Redacted evidence summary.
	Message string `json:"message"`
	// Collection timestamp.
	CollectedAt time.Time `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticCPUAnalysisK8sEvidence instantiates a new DiagnosticCPUAnalysisK8sEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticCPUAnalysisK8sEvidence(source string, namespace string, objectName string, reason string, message string, collectedAt time.Time) *DiagnosticCPUAnalysisK8sEvidence {
	this := DiagnosticCPUAnalysisK8sEvidence{}
	this.Source = source
	this.Namespace = namespace
	this.ObjectName = objectName
	this.Reason = reason
	this.Message = message
	this.CollectedAt = collectedAt
	return &this
}

// NewDiagnosticCPUAnalysisK8sEvidenceWithDefaults instantiates a new DiagnosticCPUAnalysisK8sEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticCPUAnalysisK8sEvidenceWithDefaults() *DiagnosticCPUAnalysisK8sEvidence {
	this := DiagnosticCPUAnalysisK8sEvidence{}
	return &this
}

// GetSource returns the Source field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) SetSource(v string) {
	o.Source = v
}

// GetNamespace returns the Namespace field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) SetNamespace(v string) {
	o.Namespace = v
}

// GetObjectName returns the ObjectName field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetObjectName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetObjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectName, true
}

// SetObjectName sets field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) SetObjectName(v string) {
	o.ObjectName = v
}

// GetReason returns the Reason field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) SetReason(v string) {
	o.Reason = v
}

// GetMessage returns the Message field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) SetMessage(v string) {
	o.Message = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticCPUAnalysisK8sEvidence) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticCPUAnalysisK8sEvidence) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticCPUAnalysisK8sEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["source"] = o.Source
	toSerialize["namespace"] = o.Namespace
	toSerialize["objectName"] = o.ObjectName
	toSerialize["reason"] = o.Reason
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
func (o *DiagnosticCPUAnalysisK8sEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source      *string    `json:"source"`
		Namespace   *string    `json:"namespace"`
		ObjectName  *string    `json:"objectName"`
		Reason      *string    `json:"reason"`
		Message     *string    `json:"message"`
		CollectedAt *time.Time `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	if all.ObjectName == nil {
		return fmt.Errorf("required field objectName missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"source", "namespace", "objectName", "reason", "message", "collectedAt"})
	} else {
		return err
	}
	o.Source = *all.Source
	o.Namespace = *all.Namespace
	o.ObjectName = *all.ObjectName
	o.Reason = *all.Reason
	o.Message = *all.Message
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
