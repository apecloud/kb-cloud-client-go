// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MssqlLockSnapshot Bounded live DMV observations. Edges are SQL Server reported blocking identifiers. Collection is not atomic; graph cycles are not historical deadlock evidence.
type MssqlLockSnapshot struct {
	CapturedAt string         `json:"capturedAt"`
	Sessions   []MssqlSession `json:"sessions"`
	Waits      []MssqlWait    `json:"waits"`
	Locks      []MssqlLock    `json:"locks"`
	Truncated  bool           `json:"truncated"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMssqlLockSnapshot instantiates a new MssqlLockSnapshot object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMssqlLockSnapshot(capturedAt string, sessions []MssqlSession, waits []MssqlWait, locks []MssqlLock, truncated bool) *MssqlLockSnapshot {
	this := MssqlLockSnapshot{}
	this.CapturedAt = capturedAt
	this.Sessions = sessions
	this.Waits = waits
	this.Locks = locks
	this.Truncated = truncated
	return &this
}

// NewMssqlLockSnapshotWithDefaults instantiates a new MssqlLockSnapshot object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMssqlLockSnapshotWithDefaults() *MssqlLockSnapshot {
	this := MssqlLockSnapshot{}
	return &this
}

// GetCapturedAt returns the CapturedAt field value.
func (o *MssqlLockSnapshot) GetCapturedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CapturedAt
}

// GetCapturedAtOk returns a tuple with the CapturedAt field value
// and a boolean to check if the value has been set.
func (o *MssqlLockSnapshot) GetCapturedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapturedAt, true
}

// SetCapturedAt sets field value.
func (o *MssqlLockSnapshot) SetCapturedAt(v string) {
	o.CapturedAt = v
}

// GetSessions returns the Sessions field value.
func (o *MssqlLockSnapshot) GetSessions() []MssqlSession {
	if o == nil {
		var ret []MssqlSession
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *MssqlLockSnapshot) GetSessionsOk() (*[]MssqlSession, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// SetSessions sets field value.
func (o *MssqlLockSnapshot) SetSessions(v []MssqlSession) {
	o.Sessions = v
}

// GetWaits returns the Waits field value.
func (o *MssqlLockSnapshot) GetWaits() []MssqlWait {
	if o == nil {
		var ret []MssqlWait
		return ret
	}
	return o.Waits
}

// GetWaitsOk returns a tuple with the Waits field value
// and a boolean to check if the value has been set.
func (o *MssqlLockSnapshot) GetWaitsOk() (*[]MssqlWait, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Waits, true
}

// SetWaits sets field value.
func (o *MssqlLockSnapshot) SetWaits(v []MssqlWait) {
	o.Waits = v
}

// GetLocks returns the Locks field value.
func (o *MssqlLockSnapshot) GetLocks() []MssqlLock {
	if o == nil {
		var ret []MssqlLock
		return ret
	}
	return o.Locks
}

// GetLocksOk returns a tuple with the Locks field value
// and a boolean to check if the value has been set.
func (o *MssqlLockSnapshot) GetLocksOk() (*[]MssqlLock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locks, true
}

// SetLocks sets field value.
func (o *MssqlLockSnapshot) SetLocks(v []MssqlLock) {
	o.Locks = v
}

// GetTruncated returns the Truncated field value.
func (o *MssqlLockSnapshot) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *MssqlLockSnapshot) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value.
func (o *MssqlLockSnapshot) SetTruncated(v bool) {
	o.Truncated = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MssqlLockSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["capturedAt"] = o.CapturedAt
	toSerialize["sessions"] = o.Sessions
	toSerialize["waits"] = o.Waits
	toSerialize["locks"] = o.Locks
	toSerialize["truncated"] = o.Truncated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MssqlLockSnapshot) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CapturedAt *string         `json:"capturedAt"`
		Sessions   *[]MssqlSession `json:"sessions"`
		Waits      *[]MssqlWait    `json:"waits"`
		Locks      *[]MssqlLock    `json:"locks"`
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
	if all.Waits == nil {
		return fmt.Errorf("required field waits missing")
	}
	if all.Locks == nil {
		return fmt.Errorf("required field locks missing")
	}
	if all.Truncated == nil {
		return fmt.Errorf("required field truncated missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"capturedAt", "sessions", "waits", "locks", "truncated"})
	} else {
		return err
	}
	o.CapturedAt = *all.CapturedAt
	o.Sessions = *all.Sessions
	o.Waits = *all.Waits
	o.Locks = *all.Locks
	o.Truncated = *all.Truncated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
