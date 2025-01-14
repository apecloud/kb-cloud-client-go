// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// WorkflowType workflow type
type WorkflowType string

// List of WorkflowType.
const (
	WorkflowTypeInstall                WorkflowType = "install"
	WorkflowTypeUpgradeKubeblocks      WorkflowType = "upgradeKubeblocks"
	WorkflowTypeUpgradeGemini          WorkflowType = "upgradeGemini"
	WorkflowTypeUninstall              WorkflowType = "uninstall"
	WorkflowTypeUpgradeVictoriaMetrics WorkflowType = "upgradeVictoriaMetrics"
	WorkflowTypeUpgradeLoki            WorkflowType = "upgradeLoki"
	WorkflowTypeUpgradeModule          WorkflowType = "upgradeModule"
)

var allowedWorkflowTypeEnumValues = []WorkflowType{
	WorkflowTypeInstall,
	WorkflowTypeUpgradeKubeblocks,
	WorkflowTypeUpgradeGemini,
	WorkflowTypeUninstall,
	WorkflowTypeUpgradeVictoriaMetrics,
	WorkflowTypeUpgradeLoki,
	WorkflowTypeUpgradeModule,
}

// GetAllowedValues returns the list of possible values.
func (v *WorkflowType) GetAllowedValues() []WorkflowType {
	return allowedWorkflowTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *WorkflowType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = WorkflowType(value)
	return nil
}

// NewWorkflowTypeFromValue returns a pointer to a valid WorkflowType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewWorkflowTypeFromValue(v string) (*WorkflowType, error) {
	ev := WorkflowType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for WorkflowType: valid values are %v", v, allowedWorkflowTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v WorkflowType) IsValid() bool {
	for _, existing := range allowedWorkflowTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkflowType value.
func (v WorkflowType) Ptr() *WorkflowType {
	return &v
}
