// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InspectionCriticality First-version criticality assumption for score weighting and red-item veto behavior. Missing legacy values are treated as medium.
type InspectionCriticality string

// List of InspectionCriticality.
const (
	InspectionCriticalityCritical InspectionCriticality = "critical"
	InspectionCriticalityHigh     InspectionCriticality = "high"
	InspectionCriticalityMedium   InspectionCriticality = "medium"
	InspectionCriticalityLow      InspectionCriticality = "low"
)

var allowedInspectionCriticalityEnumValues = []InspectionCriticality{
	InspectionCriticalityCritical,
	InspectionCriticalityHigh,
	InspectionCriticalityMedium,
	InspectionCriticalityLow,
}

// GetAllowedValues returns the list of possible values.
func (v *InspectionCriticality) GetAllowedValues() []InspectionCriticality {
	return allowedInspectionCriticalityEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *InspectionCriticality) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = InspectionCriticality(value)
	return nil
}

// NewInspectionCriticalityFromValue returns a pointer to a valid InspectionCriticality
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewInspectionCriticalityFromValue(v string) (*InspectionCriticality, error) {
	ev := InspectionCriticality(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for InspectionCriticality: valid values are %v", v, allowedInspectionCriticalityEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v InspectionCriticality) IsValid() bool {
	for _, existing := range allowedInspectionCriticalityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InspectionCriticality value.
func (v InspectionCriticality) Ptr() *InspectionCriticality {
	return &v
}
