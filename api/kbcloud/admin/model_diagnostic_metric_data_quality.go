// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticMetricDataQuality Data quality status for one database performance metric.
type DiagnosticMetricDataQuality string

// List of DiagnosticMetricDataQuality.
const (
	DiagnosticMetricDataQualityComplete         DiagnosticMetricDataQuality = "complete"
	DiagnosticMetricDataQualityPartial          DiagnosticMetricDataQuality = "partial"
	DiagnosticMetricDataQualitySourceMissing    DiagnosticMetricDataQuality = "source_missing"
	DiagnosticMetricDataQualityCollectionFailed DiagnosticMetricDataQuality = "collection_failed"
	DiagnosticMetricDataQualityNotApplicable    DiagnosticMetricDataQuality = "not_applicable"
)

var allowedDiagnosticMetricDataQualityEnumValues = []DiagnosticMetricDataQuality{
	DiagnosticMetricDataQualityComplete,
	DiagnosticMetricDataQualityPartial,
	DiagnosticMetricDataQualitySourceMissing,
	DiagnosticMetricDataQualityCollectionFailed,
	DiagnosticMetricDataQualityNotApplicable,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticMetricDataQuality) GetAllowedValues() []DiagnosticMetricDataQuality {
	return allowedDiagnosticMetricDataQualityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticMetricDataQuality) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticMetricDataQuality(value)
	return nil
}

// NewDiagnosticMetricDataQualityFromValue returns a pointer to a valid DiagnosticMetricDataQuality
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticMetricDataQualityFromValue(v string) (*DiagnosticMetricDataQuality, error) {
	ev := DiagnosticMetricDataQuality(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticMetricDataQuality: valid values are %v", v, allowedDiagnosticMetricDataQualityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticMetricDataQuality) IsValid() bool {
	for _, existing := range allowedDiagnosticMetricDataQualityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticMetricDataQuality value.
func (v DiagnosticMetricDataQuality) Ptr() *DiagnosticMetricDataQuality {
	return &v
}
