// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticEvidence Evidence collected for one diagnostic finding.
type DiagnosticEvidence struct {
	// Evidence source, such as Pod, Event, Node, Component, version, or log.
	Source string `json:"source"`
	// Evidence collection time.
	CollectedAt time.Time `json:"collectedAt"`
	// Kubernetes or KBE object name.
	ObjectName *string `json:"objectName,omitempty"`
	// Kubernetes namespace, when applicable.
	Namespace *string `json:"namespace,omitempty"`
	// Object kind, when applicable.
	Kind *string `json:"kind,omitempty"`
	// Redacted evidence message or important field summary.
	Message *string `json:"message,omitempty"`
	// Redacted key-value evidence fields.
	Fields map[string]string `json:"fields,omitempty"`
	// Whether this evidence has passed redaction.
	Redacted bool `json:"redacted"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticEvidence instantiates a new DiagnosticEvidence object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticEvidence(source string, collectedAt time.Time, redacted bool) *DiagnosticEvidence {
	this := DiagnosticEvidence{}
	this.Source = source
	this.CollectedAt = collectedAt
	this.Redacted = redacted
	return &this
}

// NewDiagnosticEvidenceWithDefaults instantiates a new DiagnosticEvidence object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticEvidenceWithDefaults() *DiagnosticEvidence {
	this := DiagnosticEvidence{}
	return &this
}

// GetSource returns the Source field value.
func (o *DiagnosticEvidence) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEvidence) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DiagnosticEvidence) SetSource(v string) {
	o.Source = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticEvidence) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEvidence) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticEvidence) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// GetObjectName returns the ObjectName field value if set, zero value otherwise.
func (o *DiagnosticEvidence) GetObjectName() string {
	if o == nil || o.ObjectName == nil {
		var ret string
		return ret
	}
	return *o.ObjectName
}

// GetObjectNameOk returns a tuple with the ObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticEvidence) GetObjectNameOk() (*string, bool) {
	if o == nil || o.ObjectName == nil {
		return nil, false
	}
	return o.ObjectName, true
}

// HasObjectName returns a boolean if a field has been set.
func (o *DiagnosticEvidence) HasObjectName() bool {
	return o != nil && o.ObjectName != nil
}

// SetObjectName gets a reference to the given string and assigns it to the ObjectName field.
func (o *DiagnosticEvidence) SetObjectName(v string) {
	o.ObjectName = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *DiagnosticEvidence) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticEvidence) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *DiagnosticEvidence) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *DiagnosticEvidence) SetNamespace(v string) {
	o.Namespace = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *DiagnosticEvidence) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticEvidence) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *DiagnosticEvidence) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *DiagnosticEvidence) SetKind(v string) {
	o.Kind = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DiagnosticEvidence) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticEvidence) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DiagnosticEvidence) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *DiagnosticEvidence) SetMessage(v string) {
	o.Message = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *DiagnosticEvidence) GetFields() map[string]string {
	if o == nil || o.Fields == nil {
		var ret map[string]string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticEvidence) GetFieldsOk() (*map[string]string, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *DiagnosticEvidence) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given map[string]string and assigns it to the Fields field.
func (o *DiagnosticEvidence) SetFields(v map[string]string) {
	o.Fields = v
}

// GetRedacted returns the Redacted field value.
func (o *DiagnosticEvidence) GetRedacted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Redacted
}

// GetRedactedOk returns a tuple with the Redacted field value
// and a boolean to check if the value has been set.
func (o *DiagnosticEvidence) GetRedactedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redacted, true
}

// SetRedacted sets field value.
func (o *DiagnosticEvidence) SetRedacted(v bool) {
	o.Redacted = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticEvidence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["source"] = o.Source
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ObjectName != nil {
		toSerialize["objectName"] = o.ObjectName
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	toSerialize["redacted"] = o.Redacted

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticEvidence) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source      *string           `json:"source"`
		CollectedAt *time.Time        `json:"collectedAt"`
		ObjectName  *string           `json:"objectName,omitempty"`
		Namespace   *string           `json:"namespace,omitempty"`
		Kind        *string           `json:"kind,omitempty"`
		Message     *string           `json:"message,omitempty"`
		Fields      map[string]string `json:"fields,omitempty"`
		Redacted    *bool             `json:"redacted"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.Redacted == nil {
		return fmt.Errorf("required field redacted missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"source", "collectedAt", "objectName", "namespace", "kind", "message", "fields", "redacted"})
	} else {
		return err
	}
	o.Source = *all.Source
	o.CollectedAt = *all.CollectedAt
	o.ObjectName = all.ObjectName
	o.Namespace = all.Namespace
	o.Kind = all.Kind
	o.Message = all.Message
	o.Fields = all.Fields
	o.Redacted = *all.Redacted

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
