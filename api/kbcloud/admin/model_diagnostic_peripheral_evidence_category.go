// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticPeripheralEvidenceCategory Product-facing category for one K8s / Operator / Event explanation.
type DiagnosticPeripheralEvidenceCategory string

// List of DiagnosticPeripheralEvidenceCategory.
const (
	DiagnosticPeripheralEvidenceCategoryStorage       DiagnosticPeripheralEvidenceCategory = "storage"
	DiagnosticPeripheralEvidenceCategoryNetwork       DiagnosticPeripheralEvidenceCategory = "network"
	DiagnosticPeripheralEvidenceCategoryEntry         DiagnosticPeripheralEvidenceCategory = "entry"
	DiagnosticPeripheralEvidenceCategoryOps           DiagnosticPeripheralEvidenceCategory = "ops"
	DiagnosticPeripheralEvidenceCategoryBackupRestore DiagnosticPeripheralEvidenceCategory = "backup_restore"
	DiagnosticPeripheralEvidenceCategoryOperator      DiagnosticPeripheralEvidenceCategory = "operator"
	DiagnosticPeripheralEvidenceCategoryPodStartup    DiagnosticPeripheralEvidenceCategory = "pod_startup"
)

var allowedDiagnosticPeripheralEvidenceCategoryEnumValues = []DiagnosticPeripheralEvidenceCategory{
	DiagnosticPeripheralEvidenceCategoryStorage,
	DiagnosticPeripheralEvidenceCategoryNetwork,
	DiagnosticPeripheralEvidenceCategoryEntry,
	DiagnosticPeripheralEvidenceCategoryOps,
	DiagnosticPeripheralEvidenceCategoryBackupRestore,
	DiagnosticPeripheralEvidenceCategoryOperator,
	DiagnosticPeripheralEvidenceCategoryPodStartup,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticPeripheralEvidenceCategory) GetAllowedValues() []DiagnosticPeripheralEvidenceCategory {
	return allowedDiagnosticPeripheralEvidenceCategoryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticPeripheralEvidenceCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticPeripheralEvidenceCategory(value)
	return nil
}

// NewDiagnosticPeripheralEvidenceCategoryFromValue returns a pointer to a valid DiagnosticPeripheralEvidenceCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticPeripheralEvidenceCategoryFromValue(v string) (*DiagnosticPeripheralEvidenceCategory, error) {
	ev := DiagnosticPeripheralEvidenceCategory(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticPeripheralEvidenceCategory: valid values are %v", v, allowedDiagnosticPeripheralEvidenceCategoryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticPeripheralEvidenceCategory) IsValid() bool {
	for _, existing := range allowedDiagnosticPeripheralEvidenceCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticPeripheralEvidenceCategory value.
func (v DiagnosticPeripheralEvidenceCategory) Ptr() *DiagnosticPeripheralEvidenceCategory {
	return &v
}
