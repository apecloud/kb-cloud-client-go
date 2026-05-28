// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticControllerLogRepresentativeLog Redacted representative controller log for one relation card.
type DiagnosticControllerLogRepresentativeLog struct {
	// Best known log occurrence time.
	OccurredAt time.Time `json:"occurredAt"`
	// Namespace of the controller log pod.
	Namespace string `json:"namespace"`
	// Pod name that produced the log.
	Pod string `json:"pod"`
	// Container name that produced the log.
	Container string `json:"container"`
	// Redacted representative log message.
	Message string `json:"message"`
	// Report-internal pointer to the rawEvidence entry backing this representative log.
	RawEvidencePath string `json:"rawEvidencePath"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticControllerLogRepresentativeLog instantiates a new DiagnosticControllerLogRepresentativeLog object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticControllerLogRepresentativeLog(occurredAt time.Time, namespace string, pod string, container string, message string, rawEvidencePath string) *DiagnosticControllerLogRepresentativeLog {
	this := DiagnosticControllerLogRepresentativeLog{}
	this.OccurredAt = occurredAt
	this.Namespace = namespace
	this.Pod = pod
	this.Container = container
	this.Message = message
	this.RawEvidencePath = rawEvidencePath
	return &this
}

// NewDiagnosticControllerLogRepresentativeLogWithDefaults instantiates a new DiagnosticControllerLogRepresentativeLog object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticControllerLogRepresentativeLogWithDefaults() *DiagnosticControllerLogRepresentativeLog {
	this := DiagnosticControllerLogRepresentativeLog{}
	return &this
}

// GetOccurredAt returns the OccurredAt field value.
func (o *DiagnosticControllerLogRepresentativeLog) GetOccurredAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.OccurredAt
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRepresentativeLog) GetOccurredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OccurredAt, true
}

// SetOccurredAt sets field value.
func (o *DiagnosticControllerLogRepresentativeLog) SetOccurredAt(v time.Time) {
	o.OccurredAt = v
}

// GetNamespace returns the Namespace field value.
func (o *DiagnosticControllerLogRepresentativeLog) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRepresentativeLog) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *DiagnosticControllerLogRepresentativeLog) SetNamespace(v string) {
	o.Namespace = v
}

// GetPod returns the Pod field value.
func (o *DiagnosticControllerLogRepresentativeLog) GetPod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Pod
}

// GetPodOk returns a tuple with the Pod field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRepresentativeLog) GetPodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pod, true
}

// SetPod sets field value.
func (o *DiagnosticControllerLogRepresentativeLog) SetPod(v string) {
	o.Pod = v
}

// GetContainer returns the Container field value.
func (o *DiagnosticControllerLogRepresentativeLog) GetContainer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRepresentativeLog) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value.
func (o *DiagnosticControllerLogRepresentativeLog) SetContainer(v string) {
	o.Container = v
}

// GetMessage returns the Message field value.
func (o *DiagnosticControllerLogRepresentativeLog) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRepresentativeLog) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *DiagnosticControllerLogRepresentativeLog) SetMessage(v string) {
	o.Message = v
}

// GetRawEvidencePath returns the RawEvidencePath field value.
func (o *DiagnosticControllerLogRepresentativeLog) GetRawEvidencePath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RawEvidencePath
}

// GetRawEvidencePathOk returns a tuple with the RawEvidencePath field value
// and a boolean to check if the value has been set.
func (o *DiagnosticControllerLogRepresentativeLog) GetRawEvidencePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawEvidencePath, true
}

// SetRawEvidencePath sets field value.
func (o *DiagnosticControllerLogRepresentativeLog) SetRawEvidencePath(v string) {
	o.RawEvidencePath = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticControllerLogRepresentativeLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.OccurredAt.Nanosecond() == 0 {
		toSerialize["occurredAt"] = o.OccurredAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["occurredAt"] = o.OccurredAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["namespace"] = o.Namespace
	toSerialize["pod"] = o.Pod
	toSerialize["container"] = o.Container
	toSerialize["message"] = o.Message
	toSerialize["rawEvidencePath"] = o.RawEvidencePath

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticControllerLogRepresentativeLog) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OccurredAt      *time.Time `json:"occurredAt"`
		Namespace       *string    `json:"namespace"`
		Pod             *string    `json:"pod"`
		Container       *string    `json:"container"`
		Message         *string    `json:"message"`
		RawEvidencePath *string    `json:"rawEvidencePath"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.OccurredAt == nil {
		return fmt.Errorf("required field occurredAt missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	if all.Pod == nil {
		return fmt.Errorf("required field pod missing")
	}
	if all.Container == nil {
		return fmt.Errorf("required field container missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.RawEvidencePath == nil {
		return fmt.Errorf("required field rawEvidencePath missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"occurredAt", "namespace", "pod", "container", "message", "rawEvidencePath"})
	} else {
		return err
	}
	o.OccurredAt = *all.OccurredAt
	o.Namespace = *all.Namespace
	o.Pod = *all.Pod
	o.Container = *all.Container
	o.Message = *all.Message
	o.RawEvidencePath = *all.RawEvidencePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
