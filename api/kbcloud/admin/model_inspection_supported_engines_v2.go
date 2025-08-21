// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InspectionSupportedEnginesV2 Specifies the supported engines for the inspection task.
type InspectionSupportedEnginesV2 string

// List of InspectionSupportedEnginesV2.
const (
	InspectionSupportedEnginesV2Node     InspectionSupportedEnginesV2 = "node"
	InspectionSupportedEnginesV2Mysql    InspectionSupportedEnginesV2 = "mysql"
	InspectionSupportedEnginesV2Redis    InspectionSupportedEnginesV2 = "redis"
	InspectionSupportedEnginesV2Damengdb InspectionSupportedEnginesV2 = "damengdb"
)

var allowedInspectionSupportedEnginesV2EnumValues = []InspectionSupportedEnginesV2{
	InspectionSupportedEnginesV2Node,
	InspectionSupportedEnginesV2Mysql,
	InspectionSupportedEnginesV2Redis,
	InspectionSupportedEnginesV2Damengdb,
}

// GetAllowedValues returns the list of possible values.
func (v *InspectionSupportedEnginesV2) GetAllowedValues() []InspectionSupportedEnginesV2 {
	return allowedInspectionSupportedEnginesV2EnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InspectionSupportedEnginesV2) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InspectionSupportedEnginesV2(value)
	return nil
}

// NewInspectionSupportedEnginesV2FromValue returns a pointer to a valid InspectionSupportedEnginesV2
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInspectionSupportedEnginesV2FromValue(v string) (*InspectionSupportedEnginesV2, error) {
	ev := InspectionSupportedEnginesV2(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InspectionSupportedEnginesV2: valid values are %v", v, allowedInspectionSupportedEnginesV2EnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InspectionSupportedEnginesV2) IsValid() bool {
	for _, existing := range allowedInspectionSupportedEnginesV2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InspectionSupportedEnginesV2 value.
func (v InspectionSupportedEnginesV2) Ptr() *InspectionSupportedEnginesV2 {
	return &v
}
