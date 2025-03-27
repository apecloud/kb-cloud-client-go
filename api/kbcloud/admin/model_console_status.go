// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ConsoleStatus string

// List of ConsoleStatus.
const (
	ConsoleStatusDeployed        ConsoleStatus = "deployed"
	ConsoleStatusUndeployed      ConsoleStatus = "undeployed"
	ConsoleStatusUninstalled     ConsoleStatus = "uninstalled"
	ConsoleStatusSuperseded      ConsoleStatus = "superseded"
	ConsoleStatusFailed          ConsoleStatus = "failed"
	ConsoleStatusUninstalling    ConsoleStatus = "uninstalling"
	ConsoleStatusPendingInstall  ConsoleStatus = "pending-install"
	ConsoleStatusPendingUpgrade  ConsoleStatus = "pending-upgrade"
	ConsoleStatusPendingRollback ConsoleStatus = "pending-rollback"
	ConsoleStatusUnknown         ConsoleStatus = "unknown"
)

var allowedConsoleStatusEnumValues = []ConsoleStatus{
	ConsoleStatusDeployed,
	ConsoleStatusUndeployed,
	ConsoleStatusUninstalled,
	ConsoleStatusSuperseded,
	ConsoleStatusFailed,
	ConsoleStatusUninstalling,
	ConsoleStatusPendingInstall,
	ConsoleStatusPendingUpgrade,
	ConsoleStatusPendingRollback,
	ConsoleStatusUnknown,
}

// GetAllowedValues returns the list of possible values.
func (v *ConsoleStatus) GetAllowedValues() []ConsoleStatus {
	return allowedConsoleStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ConsoleStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ConsoleStatus(value)
	return nil
}

// NewConsoleStatusFromValue returns a pointer to a valid ConsoleStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewConsoleStatusFromValue(v string) (*ConsoleStatus, error) {
	ev := ConsoleStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConsoleStatus: valid values are %v", v, allowedConsoleStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ConsoleStatus) IsValid() bool {
	for _, existing := range allowedConsoleStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConsoleStatus value.
func (v ConsoleStatus) Ptr() *ConsoleStatus {
	return &v
}
