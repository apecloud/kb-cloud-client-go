// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceSlowSQLSample One slow SQL statement digest sample shown in the Top-N list.
type DiagnosticDBPerformanceSlowSQLSample struct {
	// Stable normalized statement digest, such as a pg_stat_statements queryid string.
	Digest string `json:"digest"`
	// Optional redacted SQL text. Literals removed; safe to display.
	RedactedText *string `json:"redactedText,omitempty"`
	// Execution count within the time window.
	CallCount *int64 `json:"callCount,omitempty"`
	// Mean execution duration in milliseconds within the window.
	MeanDurationMs *float64 `json:"meanDurationMs,omitempty"`
	// Max execution duration in milliseconds within the window.
	MaxDurationMs *float64 `json:"maxDurationMs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceSlowSQLSample instantiates a new DiagnosticDBPerformanceSlowSQLSample object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceSlowSQLSample(digest string) *DiagnosticDBPerformanceSlowSQLSample {
	this := DiagnosticDBPerformanceSlowSQLSample{}
	this.Digest = digest
	return &this
}

// NewDiagnosticDBPerformanceSlowSQLSampleWithDefaults instantiates a new DiagnosticDBPerformanceSlowSQLSample object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceSlowSQLSampleWithDefaults() *DiagnosticDBPerformanceSlowSQLSample {
	this := DiagnosticDBPerformanceSlowSQLSample{}
	return &this
}

// GetDigest returns the Digest field value.
func (o *DiagnosticDBPerformanceSlowSQLSample) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLSample) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value.
func (o *DiagnosticDBPerformanceSlowSQLSample) SetDigest(v string) {
	o.Digest = v
}

// GetRedactedText returns the RedactedText field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLSample) GetRedactedText() string {
	if o == nil || o.RedactedText == nil {
		var ret string
		return ret
	}
	return *o.RedactedText
}

// GetRedactedTextOk returns a tuple with the RedactedText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLSample) GetRedactedTextOk() (*string, bool) {
	if o == nil || o.RedactedText == nil {
		return nil, false
	}
	return o.RedactedText, true
}

// HasRedactedText returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLSample) HasRedactedText() bool {
	return o != nil && o.RedactedText != nil
}

// SetRedactedText gets a reference to the given string and assigns it to the RedactedText field.
func (o *DiagnosticDBPerformanceSlowSQLSample) SetRedactedText(v string) {
	o.RedactedText = &v
}

// GetCallCount returns the CallCount field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLSample) GetCallCount() int64 {
	if o == nil || o.CallCount == nil {
		var ret int64
		return ret
	}
	return *o.CallCount
}

// GetCallCountOk returns a tuple with the CallCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLSample) GetCallCountOk() (*int64, bool) {
	if o == nil || o.CallCount == nil {
		return nil, false
	}
	return o.CallCount, true
}

// HasCallCount returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLSample) HasCallCount() bool {
	return o != nil && o.CallCount != nil
}

// SetCallCount gets a reference to the given int64 and assigns it to the CallCount field.
func (o *DiagnosticDBPerformanceSlowSQLSample) SetCallCount(v int64) {
	o.CallCount = &v
}

// GetMeanDurationMs returns the MeanDurationMs field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLSample) GetMeanDurationMs() float64 {
	if o == nil || o.MeanDurationMs == nil {
		var ret float64
		return ret
	}
	return *o.MeanDurationMs
}

// GetMeanDurationMsOk returns a tuple with the MeanDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLSample) GetMeanDurationMsOk() (*float64, bool) {
	if o == nil || o.MeanDurationMs == nil {
		return nil, false
	}
	return o.MeanDurationMs, true
}

// HasMeanDurationMs returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLSample) HasMeanDurationMs() bool {
	return o != nil && o.MeanDurationMs != nil
}

// SetMeanDurationMs gets a reference to the given float64 and assigns it to the MeanDurationMs field.
func (o *DiagnosticDBPerformanceSlowSQLSample) SetMeanDurationMs(v float64) {
	o.MeanDurationMs = &v
}

// GetMaxDurationMs returns the MaxDurationMs field value if set, zero value otherwise.
func (o *DiagnosticDBPerformanceSlowSQLSample) GetMaxDurationMs() float64 {
	if o == nil || o.MaxDurationMs == nil {
		var ret float64
		return ret
	}
	return *o.MaxDurationMs
}

// GetMaxDurationMsOk returns a tuple with the MaxDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLSample) GetMaxDurationMsOk() (*float64, bool) {
	if o == nil || o.MaxDurationMs == nil {
		return nil, false
	}
	return o.MaxDurationMs, true
}

// HasMaxDurationMs returns a boolean if a field has been set.
func (o *DiagnosticDBPerformanceSlowSQLSample) HasMaxDurationMs() bool {
	return o != nil && o.MaxDurationMs != nil
}

// SetMaxDurationMs gets a reference to the given float64 and assigns it to the MaxDurationMs field.
func (o *DiagnosticDBPerformanceSlowSQLSample) SetMaxDurationMs(v float64) {
	o.MaxDurationMs = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceSlowSQLSample) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["digest"] = o.Digest
	if o.RedactedText != nil {
		toSerialize["redactedText"] = o.RedactedText
	}
	if o.CallCount != nil {
		toSerialize["callCount"] = o.CallCount
	}
	if o.MeanDurationMs != nil {
		toSerialize["meanDurationMs"] = o.MeanDurationMs
	}
	if o.MaxDurationMs != nil {
		toSerialize["maxDurationMs"] = o.MaxDurationMs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceSlowSQLSample) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Digest         *string  `json:"digest"`
		RedactedText   *string  `json:"redactedText,omitempty"`
		CallCount      *int64   `json:"callCount,omitempty"`
		MeanDurationMs *float64 `json:"meanDurationMs,omitempty"`
		MaxDurationMs  *float64 `json:"maxDurationMs,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Digest == nil {
		return fmt.Errorf("required field digest missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"digest", "redactedText", "callCount", "meanDurationMs", "maxDurationMs"})
	} else {
		return err
	}
	o.Digest = *all.Digest
	o.RedactedText = all.RedactedText
	o.CallCount = all.CallCount
	o.MeanDurationMs = all.MeanDurationMs
	o.MaxDurationMs = all.MaxDurationMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
