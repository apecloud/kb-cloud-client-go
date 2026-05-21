// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticTimelineTimeQuality Quality of the timestamp used for ordering a diagnostic timeline item.
type DiagnosticTimelineTimeQuality string

// List of DiagnosticTimelineTimeQuality.
const (
	DiagnosticTimelineTimeQualityPrecise      DiagnosticTimelineTimeQuality = "precise"
	DiagnosticTimelineTimeQualityObservedOnly DiagnosticTimelineTimeQuality = "observed_only"
)

var allowedDiagnosticTimelineTimeQualityEnumValues = []DiagnosticTimelineTimeQuality{
	DiagnosticTimelineTimeQualityPrecise,
	DiagnosticTimelineTimeQualityObservedOnly,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticTimelineTimeQuality) GetAllowedValues() []DiagnosticTimelineTimeQuality {
	return allowedDiagnosticTimelineTimeQualityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticTimelineTimeQuality) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticTimelineTimeQuality(value)
	return nil
}

// NewDiagnosticTimelineTimeQualityFromValue returns a pointer to a valid DiagnosticTimelineTimeQuality
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticTimelineTimeQualityFromValue(v string) (*DiagnosticTimelineTimeQuality, error) {
	ev := DiagnosticTimelineTimeQuality(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticTimelineTimeQuality: valid values are %v", v, allowedDiagnosticTimelineTimeQualityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticTimelineTimeQuality) IsValid() bool {
	for _, existing := range allowedDiagnosticTimelineTimeQualityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticTimelineTimeQuality value.
func (v DiagnosticTimelineTimeQuality) Ptr() *DiagnosticTimelineTimeQuality {
	return &v
}
