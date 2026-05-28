// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticEnvironmentDiagnosisLayerType K8s environment diagnosis display layer.
type DiagnosticEnvironmentDiagnosisLayerType string

// List of DiagnosticEnvironmentDiagnosisLayerType.
const (
	DiagnosticEnvironmentDiagnosisLayerTypePod              DiagnosticEnvironmentDiagnosisLayerType = "pod"
	DiagnosticEnvironmentDiagnosisLayerTypeStorage          DiagnosticEnvironmentDiagnosisLayerType = "storage"
	DiagnosticEnvironmentDiagnosisLayerTypeServiceEndpoint  DiagnosticEnvironmentDiagnosisLayerType = "service_endpoint"
	DiagnosticEnvironmentDiagnosisLayerTypeIngressLb        DiagnosticEnvironmentDiagnosisLayerType = "ingress_lb"
	DiagnosticEnvironmentDiagnosisLayerTypeOperator         DiagnosticEnvironmentDiagnosisLayerType = "operator"
	DiagnosticEnvironmentDiagnosisLayerTypeOpsrequest       DiagnosticEnvironmentDiagnosisLayerType = "opsrequest"
	DiagnosticEnvironmentDiagnosisLayerTypeBackupRestoreJob DiagnosticEnvironmentDiagnosisLayerType = "backup_restore_job"
	DiagnosticEnvironmentDiagnosisLayerTypeEvent            DiagnosticEnvironmentDiagnosisLayerType = "event"
	DiagnosticEnvironmentDiagnosisLayerTypeLog              DiagnosticEnvironmentDiagnosisLayerType = "log"
	DiagnosticEnvironmentDiagnosisLayerTypeOther            DiagnosticEnvironmentDiagnosisLayerType = "other"
)

var allowedDiagnosticEnvironmentDiagnosisLayerTypeEnumValues = []DiagnosticEnvironmentDiagnosisLayerType{
	DiagnosticEnvironmentDiagnosisLayerTypePod,
	DiagnosticEnvironmentDiagnosisLayerTypeStorage,
	DiagnosticEnvironmentDiagnosisLayerTypeServiceEndpoint,
	DiagnosticEnvironmentDiagnosisLayerTypeIngressLb,
	DiagnosticEnvironmentDiagnosisLayerTypeOperator,
	DiagnosticEnvironmentDiagnosisLayerTypeOpsrequest,
	DiagnosticEnvironmentDiagnosisLayerTypeBackupRestoreJob,
	DiagnosticEnvironmentDiagnosisLayerTypeEvent,
	DiagnosticEnvironmentDiagnosisLayerTypeLog,
	DiagnosticEnvironmentDiagnosisLayerTypeOther,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticEnvironmentDiagnosisLayerType) GetAllowedValues() []DiagnosticEnvironmentDiagnosisLayerType {
	return allowedDiagnosticEnvironmentDiagnosisLayerTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticEnvironmentDiagnosisLayerType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticEnvironmentDiagnosisLayerType(value)
	return nil
}

// NewDiagnosticEnvironmentDiagnosisLayerTypeFromValue returns a pointer to a valid DiagnosticEnvironmentDiagnosisLayerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticEnvironmentDiagnosisLayerTypeFromValue(v string) (*DiagnosticEnvironmentDiagnosisLayerType, error) {
	ev := DiagnosticEnvironmentDiagnosisLayerType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticEnvironmentDiagnosisLayerType: valid values are %v", v, allowedDiagnosticEnvironmentDiagnosisLayerTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticEnvironmentDiagnosisLayerType) IsValid() bool {
	for _, existing := range allowedDiagnosticEnvironmentDiagnosisLayerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticEnvironmentDiagnosisLayerType value.
func (v DiagnosticEnvironmentDiagnosisLayerType) Ptr() *DiagnosticEnvironmentDiagnosisLayerType {
	return &v
}
