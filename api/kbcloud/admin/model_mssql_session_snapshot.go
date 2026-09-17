// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MssqlSessionSnapshot struct {
	CapturedAt string         `json:"capturedAt"`
	Sessions   []MssqlSession `json:"sessions"`
	Truncated  bool           `json:"truncated"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlSessionSnapshot instantiates a new MssqlSessionSnapshot object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlSessionSnapshot(capturedAt string, sessions []MssqlSession, truncated bool) *MssqlSessionSnapshot {
	this := MssqlSessionSnapshot{}
	this.CapturedAt = capturedAt
	this.Sessions = sessions
	this.Truncated = truncated
	return &this
}

// NewMssqlSessionSnapshotWithDefaults instantiates a new MssqlSessionSnapshot object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlSessionSnapshotWithDefaults() *MssqlSessionSnapshot {
	this := MssqlSessionSnapshot{}
	return &this
}

// GetCapturedAt returns the CapturedAt field value.
func (o *MssqlSessionSnapshot) GetCapturedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CapturedAt
}

// GetCapturedAtOk returns a tuple with the CapturedAt field value
// and a boolean to check if the value has been set.
func (o *MssqlSessionSnapshot) GetCapturedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAt, true
}

// SetCapturedAt sets field value.
func (o *MssqlSessionSnapshot) SetCapturedAt(v string) {
	o.CapturedAt = v
}

// GetSessions returns the Sessions field value.
func (o *MssqlSessionSnapshot) GetSessions() []MssqlSession {
	if o == nil {
		var ret []MssqlSession
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *MssqlSessionSnapshot) GetSessionsOk() (*[]MssqlSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// SetSessions sets field value.
func (o *MssqlSessionSnapshot) SetSessions(v []MssqlSession) {
	o.Sessions = v
}

// GetTruncated returns the Truncated field value.
func (o *MssqlSessionSnapshot) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *MssqlSessionSnapshot) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value.
func (o *MssqlSessionSnapshot) SetTruncated(v bool) {
	o.Truncated = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlSessionSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["capturedAt"] = o.CapturedAt
	toSerialize["sessions"] = o.Sessions
	toSerialize["truncated"] = o.Truncated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlSessionSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CapturedAt *string         `json:"capturedAt"`
		Sessions   *[]MssqlSession `json:"sessions"`
		Truncated  *bool           `json:"truncated"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CapturedAt == nil {
		return fmt.Errorf("required field capturedAt missing")
	}
	if all.Sessions == nil {
		return fmt.Errorf("required field sessions missing")
	}
	if all.Truncated == nil {
		return fmt.Errorf("required field truncated missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"capturedAt", "sessions", "truncated"})
	} else {
		return err
	}
	o.CapturedAt = *all.CapturedAt
	o.Sessions = *all.Sessions
	o.Truncated = *all.Truncated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
