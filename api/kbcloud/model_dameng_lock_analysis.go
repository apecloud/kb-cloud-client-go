// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DamengLockAnalysis struct {
	SelectedSession DamengSessionDetail `json:"selectedSession"`
	// Lock records for the selected session's transaction.
	Locks []DamengLockRow `json:"locks"`
	// Backend collection timestamp in UTC.
	CapturedAt string `json:"capturedAt"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDamengLockAnalysis instantiates a new DamengLockAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDamengLockAnalysis(selectedSession DamengSessionDetail, locks []DamengLockRow, capturedAt string) *DamengLockAnalysis {
	this := DamengLockAnalysis{}
	this.SelectedSession = selectedSession
	this.Locks = locks
	this.CapturedAt = capturedAt
	return &this
}

// NewDamengLockAnalysisWithDefaults instantiates a new DamengLockAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDamengLockAnalysisWithDefaults() *DamengLockAnalysis {
	this := DamengLockAnalysis{}
	return &this
}

// GetSelectedSession returns the SelectedSession field value.
func (o *DamengLockAnalysis) GetSelectedSession() DamengSessionDetail {
	if o == nil {
		var ret DamengSessionDetail
		return ret
	}
	return o.SelectedSession
}

// GetSelectedSessionOk returns a tuple with the SelectedSession field value
// and a boolean to check if the value has been set.
func (o *DamengLockAnalysis) GetSelectedSessionOk() (*DamengSessionDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedSession, true
}

// SetSelectedSession sets field value.
func (o *DamengLockAnalysis) SetSelectedSession(v DamengSessionDetail) {
	o.SelectedSession = v
}

// GetLocks returns the Locks field value.
func (o *DamengLockAnalysis) GetLocks() []DamengLockRow {
	if o == nil {
		var ret []DamengLockRow
		return ret
	}
	return o.Locks
}

// GetLocksOk returns a tuple with the Locks field value
// and a boolean to check if the value has been set.
func (o *DamengLockAnalysis) GetLocksOk() (*[]DamengLockRow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locks, true
}

// SetLocks sets field value.
func (o *DamengLockAnalysis) SetLocks(v []DamengLockRow) {
	o.Locks = v
}

// GetCapturedAt returns the CapturedAt field value.
func (o *DamengLockAnalysis) GetCapturedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CapturedAt
}

// GetCapturedAtOk returns a tuple with the CapturedAt field value
// and a boolean to check if the value has been set.
func (o *DamengLockAnalysis) GetCapturedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAt, true
}

// SetCapturedAt sets field value.
func (o *DamengLockAnalysis) SetCapturedAt(v string) {
	o.CapturedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DamengLockAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["selectedSession"] = o.SelectedSession
	toSerialize["locks"] = o.Locks
	toSerialize["capturedAt"] = o.CapturedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DamengLockAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SelectedSession *DamengSessionDetail `json:"selectedSession"`
		Locks           *[]DamengLockRow     `json:"locks"`
		CapturedAt      *string              `json:"capturedAt"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.SelectedSession == nil {
		return fmt.Errorf("required field selectedSession missing")
	}
	if all.Locks == nil {
		return fmt.Errorf("required field locks missing")
	}
	if all.CapturedAt == nil {
		return fmt.Errorf("required field capturedAt missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"selectedSession", "locks", "capturedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.SelectedSession.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SelectedSession = *all.SelectedSession
	o.Locks = *all.Locks
	o.CapturedAt = *all.CapturedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
