// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KingbaseLockAnalysis struct {
	SelectedSession KingbaseSession          `json:"selectedSession"`
	LockRows        []KingbaseLockRow        `json:"lockRows"`
	CannotProve     bool                     `json:"cannotProve"`
	Warnings        []string                 `json:"warnings"`
	Source          KingbaseDiagnosticSource `json:"source"`
	CapturedAt      string                   `json:"capturedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKingbaseLockAnalysis instantiates a new KingbaseLockAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKingbaseLockAnalysis(selectedSession KingbaseSession, lockRows []KingbaseLockRow, cannotProve bool, warnings []string, source KingbaseDiagnosticSource, capturedAt string) *KingbaseLockAnalysis {
	this := KingbaseLockAnalysis{}
	this.SelectedSession = selectedSession
	this.LockRows = lockRows
	this.CannotProve = cannotProve
	this.Warnings = warnings
	this.Source = source
	this.CapturedAt = capturedAt
	return &this
}

// NewKingbaseLockAnalysisWithDefaults instantiates a new KingbaseLockAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKingbaseLockAnalysisWithDefaults() *KingbaseLockAnalysis {
	this := KingbaseLockAnalysis{}
	return &this
}

// GetSelectedSession returns the SelectedSession field value.
func (o *KingbaseLockAnalysis) GetSelectedSession() KingbaseSession {
	if o == nil {
		var ret KingbaseSession
		return ret
	}
	return o.SelectedSession
}

// GetSelectedSessionOk returns a tuple with the SelectedSession field value
// and a boolean to check if the value has been set.
func (o *KingbaseLockAnalysis) GetSelectedSessionOk() (*KingbaseSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedSession, true
}

// SetSelectedSession sets field value.
func (o *KingbaseLockAnalysis) SetSelectedSession(v KingbaseSession) {
	o.SelectedSession = v
}

// GetLockRows returns the LockRows field value.
func (o *KingbaseLockAnalysis) GetLockRows() []KingbaseLockRow {
	if o == nil {
		var ret []KingbaseLockRow
		return ret
	}
	return o.LockRows
}

// GetLockRowsOk returns a tuple with the LockRows field value
// and a boolean to check if the value has been set.
func (o *KingbaseLockAnalysis) GetLockRowsOk() (*[]KingbaseLockRow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockRows, true
}

// SetLockRows sets field value.
func (o *KingbaseLockAnalysis) SetLockRows(v []KingbaseLockRow) {
	o.LockRows = v
}

// GetCannotProve returns the CannotProve field value.
func (o *KingbaseLockAnalysis) GetCannotProve() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CannotProve
}

// GetCannotProveOk returns a tuple with the CannotProve field value
// and a boolean to check if the value has been set.
func (o *KingbaseLockAnalysis) GetCannotProveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CannotProve, true
}

// SetCannotProve sets field value.
func (o *KingbaseLockAnalysis) SetCannotProve(v bool) {
	o.CannotProve = v
}

// GetWarnings returns the Warnings field value.
func (o *KingbaseLockAnalysis) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *KingbaseLockAnalysis) GetWarningsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value.
func (o *KingbaseLockAnalysis) SetWarnings(v []string) {
	o.Warnings = v
}

// GetSource returns the Source field value.
func (o *KingbaseLockAnalysis) GetSource() KingbaseDiagnosticSource {
	if o == nil {
		var ret KingbaseDiagnosticSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *KingbaseLockAnalysis) GetSourceOk() (*KingbaseDiagnosticSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *KingbaseLockAnalysis) SetSource(v KingbaseDiagnosticSource) {
	o.Source = v
}

// GetCapturedAt returns the CapturedAt field value.
func (o *KingbaseLockAnalysis) GetCapturedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CapturedAt
}

// GetCapturedAtOk returns a tuple with the CapturedAt field value
// and a boolean to check if the value has been set.
func (o *KingbaseLockAnalysis) GetCapturedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAt, true
}

// SetCapturedAt sets field value.
func (o *KingbaseLockAnalysis) SetCapturedAt(v string) {
	o.CapturedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o KingbaseLockAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["selectedSession"] = o.SelectedSession
	toSerialize["lockRows"] = o.LockRows
	toSerialize["cannotProve"] = o.CannotProve
	toSerialize["warnings"] = o.Warnings
	toSerialize["source"] = o.Source
	toSerialize["capturedAt"] = o.CapturedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KingbaseLockAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SelectedSession *KingbaseSession          `json:"selectedSession"`
		LockRows        *[]KingbaseLockRow        `json:"lockRows"`
		CannotProve     *bool                     `json:"cannotProve"`
		Warnings        *[]string                 `json:"warnings"`
		Source          *KingbaseDiagnosticSource `json:"source"`
		CapturedAt      *string                   `json:"capturedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SelectedSession == nil {
		return fmt.Errorf("required field selectedSession missing")
	}
	if all.LockRows == nil {
		return fmt.Errorf("required field lockRows missing")
	}
	if all.CannotProve == nil {
		return fmt.Errorf("required field cannotProve missing")
	}
	if all.Warnings == nil {
		return fmt.Errorf("required field warnings missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.CapturedAt == nil {
		return fmt.Errorf("required field capturedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"selectedSession", "lockRows", "cannotProve", "warnings", "source", "capturedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.SelectedSession.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SelectedSession = *all.SelectedSession
	o.LockRows = *all.LockRows
	o.CannotProve = *all.CannotProve
	o.Warnings = *all.Warnings
	if all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = *all.Source
	o.CapturedAt = *all.CapturedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
