// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InspectionThresholdDirection Direction used to interpret warnThreshold and critThreshold. asc means larger values are worse, desc means smaller values are worse, boolean means the check expression decides.
type InspectionThresholdDirection string

// List of InspectionThresholdDirection.
const (
	InspectionThresholdDirectionAsc     InspectionThresholdDirection = "asc"
	InspectionThresholdDirectionDesc    InspectionThresholdDirection = "desc"
	InspectionThresholdDirectionBoolean InspectionThresholdDirection = "boolean"
)

var allowedInspectionThresholdDirectionEnumValues = []InspectionThresholdDirection{
	InspectionThresholdDirectionAsc,
	InspectionThresholdDirectionDesc,
	InspectionThresholdDirectionBoolean,
}

// GetAllowedValues returns the list of possible values.
func (v *InspectionThresholdDirection) GetAllowedValues() []InspectionThresholdDirection {
	return allowedInspectionThresholdDirectionEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InspectionThresholdDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InspectionThresholdDirection(value)
	return nil
}

// NewInspectionThresholdDirectionFromValue returns a pointer to a valid InspectionThresholdDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInspectionThresholdDirectionFromValue(v string) (*InspectionThresholdDirection, error) {
	ev := InspectionThresholdDirection(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InspectionThresholdDirection: valid values are %v", v, allowedInspectionThresholdDirectionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InspectionThresholdDirection) IsValid() bool {
	for _, existing := range allowedInspectionThresholdDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InspectionThresholdDirection value.
func (v InspectionThresholdDirection) Ptr() *InspectionThresholdDirection {
	return &v
}
