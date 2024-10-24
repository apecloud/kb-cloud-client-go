// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// WorkflowType workflow type
type WorkflowType string

// List of WorkflowType.
const (
	WORKFLOWTYPE_INSTALL                WorkflowType = "install"
	WORKFLOWTYPE_UPGRADEKUBEBLOCKS      WorkflowType = "upgradeKubeblocks"
	WORKFLOWTYPE_UPGRADEGEMINI          WorkflowType = "upgradeGemini"
	WORKFLOWTYPE_UNINSTALL              WorkflowType = "uninstall"
	WORKFLOWTYPE_UPGRADEVICTORIAMETRICS WorkflowType = "upgradeVictoriaMetrics"
)

var allowedWorkflowTypeEnumValues = []WorkflowType{
	WORKFLOWTYPE_INSTALL,
	WORKFLOWTYPE_UPGRADEKUBEBLOCKS,
	WORKFLOWTYPE_UPGRADEGEMINI,
	WORKFLOWTYPE_UNINSTALL,
	WORKFLOWTYPE_UPGRADEVICTORIAMETRICS,
}

// GetAllowedValues reeturns the list of possible values.
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
