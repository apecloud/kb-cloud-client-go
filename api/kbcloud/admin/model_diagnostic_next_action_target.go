// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticNextActionTarget Target surface this next action points to. Matches the Sprint 7 v0.1 product contract.
type DiagnosticNextActionTarget string

// List of DiagnosticNextActionTarget.
const (
	DiagnosticNextActionTargetClusterObject    DiagnosticNextActionTarget = "cluster_object"
	DiagnosticNextActionTargetDbPod            DiagnosticNextActionTarget = "db_pod"
	DiagnosticNextActionTargetComponent        DiagnosticNextActionTarget = "component"
	DiagnosticNextActionTargetInstanceset      DiagnosticNextActionTarget = "instanceset"
	DiagnosticNextActionTargetOpsRequest       DiagnosticNextActionTarget = "ops_request"
	DiagnosticNextActionTargetControllerLog    DiagnosticNextActionTarget = "controller_log"
	DiagnosticNextActionTargetEventOrCondition DiagnosticNextActionTarget = "event_or_condition"
	DiagnosticNextActionTargetMetricSource     DiagnosticNextActionTarget = "metric_source"
	DiagnosticNextActionTargetDataSourcePlan   DiagnosticNextActionTarget = "data_source_plan"
	DiagnosticNextActionTargetControlPlane     DiagnosticNextActionTarget = "control_plane"
)

var allowedDiagnosticNextActionTargetEnumValues = []DiagnosticNextActionTarget{
	DiagnosticNextActionTargetClusterObject,
	DiagnosticNextActionTargetDbPod,
	DiagnosticNextActionTargetComponent,
	DiagnosticNextActionTargetInstanceset,
	DiagnosticNextActionTargetOpsRequest,
	DiagnosticNextActionTargetControllerLog,
	DiagnosticNextActionTargetEventOrCondition,
	DiagnosticNextActionTargetMetricSource,
	DiagnosticNextActionTargetDataSourcePlan,
	DiagnosticNextActionTargetControlPlane,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticNextActionTarget) GetAllowedValues() []DiagnosticNextActionTarget {
	return allowedDiagnosticNextActionTargetEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticNextActionTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticNextActionTarget(value)
	return nil
}

// NewDiagnosticNextActionTargetFromValue returns a pointer to a valid DiagnosticNextActionTarget
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticNextActionTargetFromValue(v string) (*DiagnosticNextActionTarget, error) {
	ev := DiagnosticNextActionTarget(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticNextActionTarget: valid values are %v", v, allowedDiagnosticNextActionTargetEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticNextActionTarget) IsValid() bool {
	for _, existing := range allowedDiagnosticNextActionTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticNextActionTarget value.
func (v DiagnosticNextActionTarget) Ptr() *DiagnosticNextActionTarget {
	return &v
}
