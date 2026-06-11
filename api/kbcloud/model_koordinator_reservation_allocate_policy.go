// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KoordinatorReservationAllocatePolicy string

// List of KoordinatorReservationAllocatePolicy.
const (
	KoordinatorReservationAllocatePolicyRestricted KoordinatorReservationAllocatePolicy = "Restricted"
	KoordinatorReservationAllocatePolicyAligned    KoordinatorReservationAllocatePolicy = "Aligned"
)

var allowedKoordinatorReservationAllocatePolicyEnumValues = []KoordinatorReservationAllocatePolicy{
	KoordinatorReservationAllocatePolicyRestricted,
	KoordinatorReservationAllocatePolicyAligned,
}

// GetAllowedValues returns the list of possible values.
func (v *KoordinatorReservationAllocatePolicy) GetAllowedValues() []KoordinatorReservationAllocatePolicy {
	return allowedKoordinatorReservationAllocatePolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *KoordinatorReservationAllocatePolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = KoordinatorReservationAllocatePolicy(value)
	return nil
}

// NewKoordinatorReservationAllocatePolicyFromValue returns a pointer to a valid KoordinatorReservationAllocatePolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewKoordinatorReservationAllocatePolicyFromValue(v string) (*KoordinatorReservationAllocatePolicy, error) {
	ev := KoordinatorReservationAllocatePolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for KoordinatorReservationAllocatePolicy: valid values are %v", v, allowedKoordinatorReservationAllocatePolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v KoordinatorReservationAllocatePolicy) IsValid() bool {
	for _, existing := range allowedKoordinatorReservationAllocatePolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KoordinatorReservationAllocatePolicy value.
func (v KoordinatorReservationAllocatePolicy) Ptr() *KoordinatorReservationAllocatePolicy {
	return &v
}
