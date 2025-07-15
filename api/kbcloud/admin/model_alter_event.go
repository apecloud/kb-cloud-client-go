// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// AlterEvent Alert event type
type AlterEvent string

// List of AlterEvent.
const (
	AlterEventFailover           AlterEvent = "failover"
	AlterEventSwitchover         AlterEvent = "switchover"
	AlterEventRestart            AlterEvent = "restart"
	AlterEventUpgrade            AlterEvent = "upgrade"
	AlterEventVolumeExpand       AlterEvent = "volume expand"
	AlterEventVscale             AlterEvent = "vscale"
	AlterEventHscale             AlterEvent = "hscale"
	AlterEventReconfigure        AlterEvent = "reconfigure"
	AlterEventBackup             AlterEvent = "backup"
	AlterEventNetworkUnavailable AlterEvent = "network unavailable"
)

var allowedAlterEventEnumValues = []AlterEvent{
	AlterEventFailover,
	AlterEventSwitchover,
	AlterEventRestart,
	AlterEventUpgrade,
	AlterEventVolumeExpand,
	AlterEventVscale,
	AlterEventHscale,
	AlterEventReconfigure,
	AlterEventBackup,
	AlterEventNetworkUnavailable,
}

// GetAllowedValues returns the list of possible values.
func (v *AlterEvent) GetAllowedValues() []AlterEvent {
	return allowedAlterEventEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlterEvent) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlterEvent(value)
	return nil
}

// NewAlterEventFromValue returns a pointer to a valid AlterEvent
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlterEventFromValue(v string) (*AlterEvent, error) {
	ev := AlterEvent(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlterEvent: valid values are %v", v, allowedAlterEventEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlterEvent) IsValid() bool {
	for _, existing := range allowedAlterEventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlterEvent value.
func (v AlterEvent) Ptr() *AlterEvent {
	return &v
}
