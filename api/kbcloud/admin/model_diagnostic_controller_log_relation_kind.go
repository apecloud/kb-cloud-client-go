// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticControllerLogRelationKind Product-facing relation bucket for one controller-log explanation card.
// This is separate from the legacy check-level relationStatus/riskDomain
// fields, which remain for Sprint 6/7 compatibility.
type DiagnosticControllerLogRelationKind string

// List of DiagnosticControllerLogRelationKind.
const (
	DiagnosticControllerLogRelationKindCurrentInstanceRelated DiagnosticControllerLogRelationKind = "current_instance_related"
	DiagnosticControllerLogRelationKindCurrentClusterRelated  DiagnosticControllerLogRelationKind = "current_cluster_related"
	DiagnosticControllerLogRelationKindControlPlaneRisk       DiagnosticControllerLogRelationKind = "control_plane_risk"
	DiagnosticControllerLogRelationKindUnknownRelation        DiagnosticControllerLogRelationKind = "unknown_relation"
	DiagnosticControllerLogRelationKindUnrelated              DiagnosticControllerLogRelationKind = "unrelated"
)

var allowedDiagnosticControllerLogRelationKindEnumValues = []DiagnosticControllerLogRelationKind{
	DiagnosticControllerLogRelationKindCurrentInstanceRelated,
	DiagnosticControllerLogRelationKindCurrentClusterRelated,
	DiagnosticControllerLogRelationKindControlPlaneRisk,
	DiagnosticControllerLogRelationKindUnknownRelation,
	DiagnosticControllerLogRelationKindUnrelated,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticControllerLogRelationKind) GetAllowedValues() []DiagnosticControllerLogRelationKind {
	return allowedDiagnosticControllerLogRelationKindEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticControllerLogRelationKind) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticControllerLogRelationKind(value)
	return nil
}

// NewDiagnosticControllerLogRelationKindFromValue returns a pointer to a valid DiagnosticControllerLogRelationKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticControllerLogRelationKindFromValue(v string) (*DiagnosticControllerLogRelationKind, error) {
	ev := DiagnosticControllerLogRelationKind(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticControllerLogRelationKind: valid values are %v", v, allowedDiagnosticControllerLogRelationKindEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticControllerLogRelationKind) IsValid() bool {
	for _, existing := range allowedDiagnosticControllerLogRelationKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticControllerLogRelationKind value.
func (v DiagnosticControllerLogRelationKind) Ptr() *DiagnosticControllerLogRelationKind {
	return &v
}
