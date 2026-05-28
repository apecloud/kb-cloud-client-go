// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticNextActionCopyPayload Safe, redacted text payload that Console may copy for DBA/SRE handoff. It must not contain passwords, tokens, cookies, kubeconfig, or full sensitive SQL text.
type DiagnosticNextActionCopyPayload struct {
	// Payload format. Sprint 19 only uses plain_text.
	Format *DiagnosticNextActionCopyPayloadFormat `json:"format,omitempty"`
	// Whether the payload has been redacted before returning to Console.
	Redacted bool `json:"redacted"`
	// Redacted copy-ready diagnostic context.
	Text string `json:"text"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticNextActionCopyPayload instantiates a new DiagnosticNextActionCopyPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticNextActionCopyPayload(redacted bool, text string) *DiagnosticNextActionCopyPayload {
	this := DiagnosticNextActionCopyPayload{}
	this.Redacted = redacted
	this.Text = text
	return &this
}

// NewDiagnosticNextActionCopyPayloadWithDefaults instantiates a new DiagnosticNextActionCopyPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticNextActionCopyPayloadWithDefaults() *DiagnosticNextActionCopyPayload {
	this := DiagnosticNextActionCopyPayload{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *DiagnosticNextActionCopyPayload) GetFormat() DiagnosticNextActionCopyPayloadFormat {
	if o == nil || o.Format == nil {
		var ret DiagnosticNextActionCopyPayloadFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticNextActionCopyPayload) GetFormatOk() (*DiagnosticNextActionCopyPayloadFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *DiagnosticNextActionCopyPayload) HasFormat() bool {
	return o != nil && o.Format != nil
}

// SetFormat gets a reference to the given DiagnosticNextActionCopyPayloadFormat and assigns it to the Format field.
func (o *DiagnosticNextActionCopyPayload) SetFormat(v DiagnosticNextActionCopyPayloadFormat) {
	o.Format = &v
}

// GetRedacted returns the Redacted field value.
func (o *DiagnosticNextActionCopyPayload) GetRedacted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Redacted
}

// GetRedactedOk returns a tuple with the Redacted field value
// and a boolean to check if the value has been set.
func (o *DiagnosticNextActionCopyPayload) GetRedactedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redacted, true
}

// SetRedacted sets field value.
func (o *DiagnosticNextActionCopyPayload) SetRedacted(v bool) {
	o.Redacted = v
}

// GetText returns the Text field value.
func (o *DiagnosticNextActionCopyPayload) GetText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *DiagnosticNextActionCopyPayload) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value.
func (o *DiagnosticNextActionCopyPayload) SetText(v string) {
	o.Text = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticNextActionCopyPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	toSerialize["redacted"] = o.Redacted
	toSerialize["text"] = o.Text

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticNextActionCopyPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Format   *DiagnosticNextActionCopyPayloadFormat `json:"format,omitempty"`
		Redacted *bool                                  `json:"redacted"`
		Text     *string                                `json:"text"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Redacted == nil {
		return fmt.Errorf("required field redacted missing")
	}
	if all.Text == nil {
		return fmt.Errorf("required field text missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"format", "redacted", "text"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Format != nil && !all.Format.IsValid() {
		hasInvalidField = true
	} else {
		o.Format = all.Format
	}
	o.Redacted = *all.Redacted
	o.Text = *all.Text

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
