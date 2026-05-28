// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticTimelineSource Source type for one diagnostic timeline item.
type DiagnosticTimelineSource string

// List of DiagnosticTimelineSource.
const (
	DiagnosticTimelineSourceEvent         DiagnosticTimelineSource = "event"
	DiagnosticTimelineSourceCondition     DiagnosticTimelineSource = "condition"
	DiagnosticTimelineSourceOpsrequest    DiagnosticTimelineSource = "opsrequest"
	DiagnosticTimelineSourceControllerLog DiagnosticTimelineSource = "controller_log"
	DiagnosticTimelineSourceMetric        DiagnosticTimelineSource = "metric"
)

var allowedDiagnosticTimelineSourceEnumValues = []DiagnosticTimelineSource{
	DiagnosticTimelineSourceEvent,
	DiagnosticTimelineSourceCondition,
	DiagnosticTimelineSourceOpsrequest,
	DiagnosticTimelineSourceControllerLog,
	DiagnosticTimelineSourceMetric,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticTimelineSource) GetAllowedValues() []DiagnosticTimelineSource {
	return allowedDiagnosticTimelineSourceEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticTimelineSource) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticTimelineSource(value)
	return nil
}

// NewDiagnosticTimelineSourceFromValue returns a pointer to a valid DiagnosticTimelineSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticTimelineSourceFromValue(v string) (*DiagnosticTimelineSource, error) {
	ev := DiagnosticTimelineSource(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticTimelineSource: valid values are %v", v, allowedDiagnosticTimelineSourceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticTimelineSource) IsValid() bool {
	for _, existing := range allowedDiagnosticTimelineSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticTimelineSource value.
func (v DiagnosticTimelineSource) Ptr() *DiagnosticTimelineSource {
	return &v
}
