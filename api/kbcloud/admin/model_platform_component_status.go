// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PlatformComponentStatus Platform component status
type PlatformComponentStatus string

// List of PlatformComponentStatus.
const (
	PlatformComponentStatusRunning PlatformComponentStatus = "Running"
	PlatformComponentStatusError   PlatformComponentStatus = "Error"
)

var allowedPlatformComponentStatusEnumValues = []PlatformComponentStatus{
	PlatformComponentStatusRunning,
	PlatformComponentStatusError,
}

// GetAllowedValues returns the list of possible values.
func (v *PlatformComponentStatus) GetAllowedValues() []PlatformComponentStatus {
	return allowedPlatformComponentStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PlatformComponentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PlatformComponentStatus(value)
	return nil
}

// NewPlatformComponentStatusFromValue returns a pointer to a valid PlatformComponentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPlatformComponentStatusFromValue(v string) (*PlatformComponentStatus, error) {
	ev := PlatformComponentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PlatformComponentStatus: valid values are %v", v, allowedPlatformComponentStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PlatformComponentStatus) IsValid() bool {
	for _, existing := range allowedPlatformComponentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlatformComponentStatus value.
func (v PlatformComponentStatus) Ptr() *PlatformComponentStatus {
	return &v
}
