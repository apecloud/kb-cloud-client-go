// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSpaceAnalysisThresholdSortBy Stable ordering key for topObjects.
type DiagnosticSpaceAnalysisThresholdSortBy string

// List of DiagnosticSpaceAnalysisThresholdSortBy.
const (
	DiagnosticSpaceAnalysisThresholdSortBySizeBytes DiagnosticSpaceAnalysisThresholdSortBy = "sizeBytes"
)

var allowedDiagnosticSpaceAnalysisThresholdSortByEnumValues = []DiagnosticSpaceAnalysisThresholdSortBy{
	DiagnosticSpaceAnalysisThresholdSortBySizeBytes,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticSpaceAnalysisThresholdSortBy) GetAllowedValues() []DiagnosticSpaceAnalysisThresholdSortBy {
	return allowedDiagnosticSpaceAnalysisThresholdSortByEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticSpaceAnalysisThresholdSortBy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticSpaceAnalysisThresholdSortBy(value)
	return nil
}

// NewDiagnosticSpaceAnalysisThresholdSortByFromValue returns a pointer to a valid DiagnosticSpaceAnalysisThresholdSortBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticSpaceAnalysisThresholdSortByFromValue(v string) (*DiagnosticSpaceAnalysisThresholdSortBy, error) {
	ev := DiagnosticSpaceAnalysisThresholdSortBy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticSpaceAnalysisThresholdSortBy: valid values are %v", v, allowedDiagnosticSpaceAnalysisThresholdSortByEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticSpaceAnalysisThresholdSortBy) IsValid() bool {
	for _, existing := range allowedDiagnosticSpaceAnalysisThresholdSortByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticSpaceAnalysisThresholdSortBy value.
func (v DiagnosticSpaceAnalysisThresholdSortBy) Ptr() *DiagnosticSpaceAnalysisThresholdSortBy {
	return &v
}
