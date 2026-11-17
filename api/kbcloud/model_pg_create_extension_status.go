// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PgCreateExtensionStatus string

// List of PgCreateExtensionStatus.
const (
	PgCreateExtensionStatusInstalled        PgCreateExtensionStatus = "installed"
	PgCreateExtensionStatusAlreadyInstalled PgCreateExtensionStatus = "already_installed"
	PgCreateExtensionStatusScheduled        PgCreateExtensionStatus = "scheduled"
	PgCreateExtensionStatusFailed           PgCreateExtensionStatus = "failed"
)

var allowedPgCreateExtensionStatusEnumValues = []PgCreateExtensionStatus{
	PgCreateExtensionStatusInstalled,
	PgCreateExtensionStatusAlreadyInstalled,
	PgCreateExtensionStatusScheduled,
	PgCreateExtensionStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *PgCreateExtensionStatus) GetAllowedValues() []PgCreateExtensionStatus {
	return allowedPgCreateExtensionStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PgCreateExtensionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PgCreateExtensionStatus(value)
	return nil
}

// NewPgCreateExtensionStatusFromValue returns a pointer to a valid PgCreateExtensionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPgCreateExtensionStatusFromValue(v string) (*PgCreateExtensionStatus, error) {
	ev := PgCreateExtensionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PgCreateExtensionStatus: valid values are %v", v, allowedPgCreateExtensionStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PgCreateExtensionStatus) IsValid() bool {
	for _, existing := range allowedPgCreateExtensionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PgCreateExtensionStatus value.
func (v PgCreateExtensionStatus) Ptr() *PgCreateExtensionStatus {
	return &v
}
