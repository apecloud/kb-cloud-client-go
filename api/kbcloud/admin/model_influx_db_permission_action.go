// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type InfluxDBPermissionAction string

// List of InfluxDBPermissionAction.
const (
	InfluxDBPermissionActionRead  InfluxDBPermissionAction = "read"
	InfluxDBPermissionActionWrite InfluxDBPermissionAction = "write"
)

var allowedInfluxDBPermissionActionEnumValues = []InfluxDBPermissionAction{
	InfluxDBPermissionActionRead,
	InfluxDBPermissionActionWrite,
}

// GetAllowedValues returns the list of possible values.
func (v *InfluxDBPermissionAction) GetAllowedValues() []InfluxDBPermissionAction {
	return allowedInfluxDBPermissionActionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InfluxDBPermissionAction) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InfluxDBPermissionAction(value)
	return nil
}

// NewInfluxDBPermissionActionFromValue returns a pointer to a valid InfluxDBPermissionAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInfluxDBPermissionActionFromValue(v string) (*InfluxDBPermissionAction, error) {
	ev := InfluxDBPermissionAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InfluxDBPermissionAction: valid values are %v", v, allowedInfluxDBPermissionActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InfluxDBPermissionAction) IsValid() bool {
	for _, existing := range allowedInfluxDBPermissionActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InfluxDBPermissionAction value.
func (v InfluxDBPermissionAction) Ptr() *InfluxDBPermissionAction {
	return &v
}
