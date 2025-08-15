// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PreCheckStatus string

// List of PreCheckStatus.
const (
	PreCheckStatusRunning PreCheckStatus = "running"
	PreCheckStatusSucceed PreCheckStatus = "succeed"
	PreCheckStatusFailed  PreCheckStatus = "failed"
)

var allowedPreCheckStatusEnumValues = []PreCheckStatus{
	PreCheckStatusRunning,
	PreCheckStatusSucceed,
	PreCheckStatusFailed,
}

// GetAllowedValues returns the list of possible values.
func (v *PreCheckStatus) GetAllowedValues() []PreCheckStatus {
	return allowedPreCheckStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PreCheckStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PreCheckStatus(value)
	return nil
}

// NewPreCheckStatusFromValue returns a pointer to a valid PreCheckStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPreCheckStatusFromValue(v string) (*PreCheckStatus, error) {
	ev := PreCheckStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PreCheckStatus: valid values are %v", v, allowedPreCheckStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PreCheckStatus) IsValid() bool {
	for _, existing := range allowedPreCheckStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PreCheckStatus value.
func (v PreCheckStatus) Ptr() *PreCheckStatus {
	return &v
}
