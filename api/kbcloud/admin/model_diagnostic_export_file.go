// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticExportFile One file included in a redacted diagnostic package.
type DiagnosticExportFile struct {
	// Path inside the diagnostic package.
	Path string `json:"path"`
	// File type, such as report-json, manifest-json, yaml, event, or log.
	Type string `json:"type"`
	// Collection time for this file.
	CollectedAt time.Time `json:"collectedAt"`
	// Source that produced this file.
	Source string `json:"source"`
	// Whether the file has passed redaction.
	Redacted bool `json:"redacted"`
	// Collection result for one file included in a redacted diagnostic package.
	Result DiagnosticExportFileResult `json:"result"`
	// Failure, skipped, or not-applicable reason. Empty when result is success.
	Reason string `json:"reason"`
	// Redacted file size in bytes.
	SizeBytes int64 `json:"sizeBytes"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticExportFile instantiates a new DiagnosticExportFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticExportFile(path string, typeVar string, collectedAt time.Time, source string, redacted bool, result DiagnosticExportFileResult, reason string, sizeBytes int64) *DiagnosticExportFile {
	this := DiagnosticExportFile{}
	this.Path = path
	this.Type = typeVar
	this.CollectedAt = collectedAt
	this.Source = source
	this.Redacted = redacted
	this.Result = result
	this.Reason = reason
	this.SizeBytes = sizeBytes
	return &this
}

// NewDiagnosticExportFileWithDefaults instantiates a new DiagnosticExportFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticExportFileWithDefaults() *DiagnosticExportFile {
	this := DiagnosticExportFile{}
	return &this
}

// GetPath returns the Path field value.
func (o *DiagnosticExportFile) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportFile) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *DiagnosticExportFile) SetPath(v string) {
	o.Path = v
}

// GetType returns the Type field value.
func (o *DiagnosticExportFile) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportFile) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DiagnosticExportFile) SetType(v string) {
	o.Type = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *DiagnosticExportFile) GetCollectedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportFile) GetCollectedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *DiagnosticExportFile) SetCollectedAt(v time.Time) {
	o.CollectedAt = v
}

// GetSource returns the Source field value.
func (o *DiagnosticExportFile) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportFile) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *DiagnosticExportFile) SetSource(v string) {
	o.Source = v
}

// GetRedacted returns the Redacted field value.
func (o *DiagnosticExportFile) GetRedacted() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Redacted
}

// GetRedactedOk returns a tuple with the Redacted field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportFile) GetRedactedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redacted, true
}

// SetRedacted sets field value.
func (o *DiagnosticExportFile) SetRedacted(v bool) {
	o.Redacted = v
}

// GetResult returns the Result field value.
func (o *DiagnosticExportFile) GetResult() DiagnosticExportFileResult {
	if o == nil {
		var ret DiagnosticExportFileResult
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportFile) GetResultOk() (*DiagnosticExportFileResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value.
func (o *DiagnosticExportFile) SetResult(v DiagnosticExportFileResult) {
	o.Result = v
}

// GetReason returns the Reason field value.
func (o *DiagnosticExportFile) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportFile) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value.
func (o *DiagnosticExportFile) SetReason(v string) {
	o.Reason = v
}

// GetSizeBytes returns the SizeBytes field value.
func (o *DiagnosticExportFile) GetSizeBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SizeBytes
}

// GetSizeBytesOk returns a tuple with the SizeBytes field value
// and a boolean to check if the value has been set.
func (o *DiagnosticExportFile) GetSizeBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SizeBytes, true
}

// SetSizeBytes sets field value.
func (o *DiagnosticExportFile) SetSizeBytes(v int64) {
	o.SizeBytes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticExportFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["path"] = o.Path
	toSerialize["type"] = o.Type
	if o.CollectedAt.Nanosecond() == 0 {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["collectedAt"] = o.CollectedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["source"] = o.Source
	toSerialize["redacted"] = o.Redacted
	toSerialize["result"] = o.Result
	toSerialize["reason"] = o.Reason
	toSerialize["sizeBytes"] = o.SizeBytes

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticExportFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Path        *string                     `json:"path"`
		Type        *string                     `json:"type"`
		CollectedAt *time.Time                  `json:"collectedAt"`
		Source      *string                     `json:"source"`
		Redacted    *bool                       `json:"redacted"`
		Result      *DiagnosticExportFileResult `json:"result"`
		Reason      *string                     `json:"reason"`
		SizeBytes   *int64                      `json:"sizeBytes"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Redacted == nil {
		return fmt.Errorf("required field redacted missing")
	}
	if all.Result == nil {
		return fmt.Errorf("required field result missing")
	}
	if all.Reason == nil {
		return fmt.Errorf("required field reason missing")
	}
	if all.SizeBytes == nil {
		return fmt.Errorf("required field sizeBytes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"path", "type", "collectedAt", "source", "redacted", "result", "reason", "sizeBytes"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Path = *all.Path
	o.Type = *all.Type
	o.CollectedAt = *all.CollectedAt
	o.Source = *all.Source
	o.Redacted = *all.Redacted
	if !all.Result.IsValid() {
		hasInvalidField = true
	} else {
		o.Result = *all.Result
	}
	o.Reason = *all.Reason
	o.SizeBytes = *all.SizeBytes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
