// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type OccupationType string

// List of OccupationType.
const (
	OccupationTypeGeneral   OccupationType = "general"
	OccupationTypeExclusive OccupationType = "exclusive"
)

var allowedOccupationTypeEnumValues = []OccupationType{
	OccupationTypeGeneral,
	OccupationTypeExclusive,
}

// GetAllowedValues returns the list of possible values.
func (v *OccupationType) GetAllowedValues() []OccupationType {
	return allowedOccupationTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OccupationType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OccupationType(value)
	return nil
}

// NewOccupationTypeFromValue returns a pointer to a valid OccupationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOccupationTypeFromValue(v string) (*OccupationType, error) {
	ev := OccupationType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OccupationType: valid values are %v", v, allowedOccupationTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OccupationType) IsValid() bool {
	for _, existing := range allowedOccupationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OccupationType value.
func (v OccupationType) Ptr() *OccupationType {
	return &v
}
