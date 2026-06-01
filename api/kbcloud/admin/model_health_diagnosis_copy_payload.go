// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type HealthDiagnosisCopyPayload struct {
	Format   HealthDiagnosisCopyPayloadFormat `json:"format"`
	Redacted bool                             `json:"redacted"`
	Content  string                           `json:"content"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHealthDiagnosisCopyPayload instantiates a new HealthDiagnosisCopyPayload object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHealthDiagnosisCopyPayload(format HealthDiagnosisCopyPayloadFormat, redacted bool, content string) *HealthDiagnosisCopyPayload {
	this := HealthDiagnosisCopyPayload{}
	this.Format = format
	this.Redacted = redacted
	this.Content = content
	return &this
}

// NewHealthDiagnosisCopyPayloadWithDefaults instantiates a new HealthDiagnosisCopyPayload object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHealthDiagnosisCopyPayloadWithDefaults() *HealthDiagnosisCopyPayload {
	this := HealthDiagnosisCopyPayload{}
	return &this
}

// GetFormat returns the Format field value.
func (o *HealthDiagnosisCopyPayload) GetFormat() HealthDiagnosisCopyPayloadFormat {
	if o == nil {
		var ret HealthDiagnosisCopyPayloadFormat
		return ret
	}
	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisCopyPayload) GetFormatOk() (*HealthDiagnosisCopyPayloadFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value.
func (o *HealthDiagnosisCopyPayload) SetFormat(v HealthDiagnosisCopyPayloadFormat) {
	o.Format = v
}

// GetRedacted returns the Redacted field value.
func (o *HealthDiagnosisCopyPayload) GetRedacted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Redacted
}

// GetRedactedOk returns a tuple with the Redacted field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisCopyPayload) GetRedactedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redacted, true
}

// SetRedacted sets field value.
func (o *HealthDiagnosisCopyPayload) SetRedacted(v bool) {
	o.Redacted = v
}

// GetContent returns the Content field value.
func (o *HealthDiagnosisCopyPayload) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *HealthDiagnosisCopyPayload) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *HealthDiagnosisCopyPayload) SetContent(v string) {
	o.Content = v
}

// MarshalJSON serializes the struct using spec logic.
func (o HealthDiagnosisCopyPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["format"] = o.Format
	toSerialize["redacted"] = o.Redacted
	toSerialize["content"] = o.Content

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HealthDiagnosisCopyPayload) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Format   *HealthDiagnosisCopyPayloadFormat `json:"format"`
		Redacted *bool                             `json:"redacted"`
		Content  *string                           `json:"content"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Format == nil {
		return fmt.Errorf("required field format missing")
	}
	if all.Redacted == nil {
		return fmt.Errorf("required field redacted missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"format", "redacted", "content"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Format.IsValid() {
		hasInvalidField = true
	} else {
		o.Format = *all.Format
	}
	o.Redacted = *all.Redacted
	o.Content = *all.Content

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
