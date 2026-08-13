// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KingbaseSessionDetail struct {
	Session     KingbaseSession          `json:"session"`
	Source      KingbaseDiagnosticSource `json:"source"`
	Warnings    []string                 `json:"warnings"`
	CollectedAt string                   `json:"collectedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKingbaseSessionDetail instantiates a new KingbaseSessionDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKingbaseSessionDetail(session KingbaseSession, source KingbaseDiagnosticSource, warnings []string, collectedAt string) *KingbaseSessionDetail {
	this := KingbaseSessionDetail{}
	this.Session = session
	this.Source = source
	this.Warnings = warnings
	this.CollectedAt = collectedAt
	return &this
}

// NewKingbaseSessionDetailWithDefaults instantiates a new KingbaseSessionDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKingbaseSessionDetailWithDefaults() *KingbaseSessionDetail {
	this := KingbaseSessionDetail{}
	return &this
}

// GetSession returns the Session field value.
func (o *KingbaseSessionDetail) GetSession() KingbaseSession {
	if o == nil {
		var ret KingbaseSession
		return ret
	}
	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *KingbaseSessionDetail) GetSessionOk() (*KingbaseSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value.
func (o *KingbaseSessionDetail) SetSession(v KingbaseSession) {
	o.Session = v
}

// GetSource returns the Source field value.
func (o *KingbaseSessionDetail) GetSource() KingbaseDiagnosticSource {
	if o == nil {
		var ret KingbaseDiagnosticSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *KingbaseSessionDetail) GetSourceOk() (*KingbaseDiagnosticSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *KingbaseSessionDetail) SetSource(v KingbaseDiagnosticSource) {
	o.Source = v
}

// GetWarnings returns the Warnings field value.
func (o *KingbaseSessionDetail) GetWarnings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value
// and a boolean to check if the value has been set.
func (o *KingbaseSessionDetail) GetWarningsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// SetWarnings sets field value.
func (o *KingbaseSessionDetail) SetWarnings(v []string) {
	o.Warnings = v
}

// GetCollectedAt returns the CollectedAt field value.
func (o *KingbaseSessionDetail) GetCollectedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CollectedAt
}

// GetCollectedAtOk returns a tuple with the CollectedAt field value
// and a boolean to check if the value has been set.
func (o *KingbaseSessionDetail) GetCollectedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectedAt, true
}

// SetCollectedAt sets field value.
func (o *KingbaseSessionDetail) SetCollectedAt(v string) {
	o.CollectedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o KingbaseSessionDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["session"] = o.Session
	toSerialize["source"] = o.Source
	toSerialize["warnings"] = o.Warnings
	toSerialize["collectedAt"] = o.CollectedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *KingbaseSessionDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Session     *KingbaseSession          `json:"session"`
		Source      *KingbaseDiagnosticSource `json:"source"`
		Warnings    *[]string                 `json:"warnings"`
		CollectedAt *string                   `json:"collectedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Session == nil {
		return fmt.Errorf("required field session missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Warnings == nil {
		return fmt.Errorf("required field warnings missing")
	}
	if all.CollectedAt == nil {
		return fmt.Errorf("required field collectedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"session", "source", "warnings", "collectedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Session.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Session = *all.Session
	if all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = *all.Source
	o.Warnings = *all.Warnings
	o.CollectedAt = *all.CollectedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
