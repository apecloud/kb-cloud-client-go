// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticSpaceAnalysisObjectType PostgreSQL relation type bucket used by space analysis.
type DiagnosticSpaceAnalysisObjectType string

// List of DiagnosticSpaceAnalysisObjectType.
const (
	DiagnosticSpaceAnalysisObjectTypeTable            DiagnosticSpaceAnalysisObjectType = "table"
	DiagnosticSpaceAnalysisObjectTypePartitionedTable DiagnosticSpaceAnalysisObjectType = "partitioned_table"
	DiagnosticSpaceAnalysisObjectTypeIndex            DiagnosticSpaceAnalysisObjectType = "index"
	DiagnosticSpaceAnalysisObjectTypeToast            DiagnosticSpaceAnalysisObjectType = "toast"
	DiagnosticSpaceAnalysisObjectTypeMaterializedView DiagnosticSpaceAnalysisObjectType = "materialized_view"
	DiagnosticSpaceAnalysisObjectTypeOther            DiagnosticSpaceAnalysisObjectType = "other"
)

var allowedDiagnosticSpaceAnalysisObjectTypeEnumValues = []DiagnosticSpaceAnalysisObjectType{
	DiagnosticSpaceAnalysisObjectTypeTable,
	DiagnosticSpaceAnalysisObjectTypePartitionedTable,
	DiagnosticSpaceAnalysisObjectTypeIndex,
	DiagnosticSpaceAnalysisObjectTypeToast,
	DiagnosticSpaceAnalysisObjectTypeMaterializedView,
	DiagnosticSpaceAnalysisObjectTypeOther,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticSpaceAnalysisObjectType) GetAllowedValues() []DiagnosticSpaceAnalysisObjectType {
	return allowedDiagnosticSpaceAnalysisObjectTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticSpaceAnalysisObjectType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticSpaceAnalysisObjectType(value)
	return nil
}

// NewDiagnosticSpaceAnalysisObjectTypeFromValue returns a pointer to a valid DiagnosticSpaceAnalysisObjectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticSpaceAnalysisObjectTypeFromValue(v string) (*DiagnosticSpaceAnalysisObjectType, error) {
	ev := DiagnosticSpaceAnalysisObjectType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticSpaceAnalysisObjectType: valid values are %v", v, allowedDiagnosticSpaceAnalysisObjectTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticSpaceAnalysisObjectType) IsValid() bool {
	for _, existing := range allowedDiagnosticSpaceAnalysisObjectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticSpaceAnalysisObjectType value.
func (v DiagnosticSpaceAnalysisObjectType) Ptr() *DiagnosticSpaceAnalysisObjectType {
	return &v
}
