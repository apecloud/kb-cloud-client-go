// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type KoordinatorReservationResourceClass string

// List of KoordinatorReservationResourceClass.
const (
	KoordinatorReservationResourceClassDatabase KoordinatorReservationResourceClass = "database"
	KoordinatorReservationResourceClassBusiness KoordinatorReservationResourceClass = "business"
)

var allowedKoordinatorReservationResourceClassEnumValues = []KoordinatorReservationResourceClass{
	KoordinatorReservationResourceClassDatabase,
	KoordinatorReservationResourceClassBusiness,
}

// GetAllowedValues returns the list of possible values.
func (v *KoordinatorReservationResourceClass) GetAllowedValues() []KoordinatorReservationResourceClass {
	return allowedKoordinatorReservationResourceClassEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *KoordinatorReservationResourceClass) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = KoordinatorReservationResourceClass(value)
	return nil
}

// NewKoordinatorReservationResourceClassFromValue returns a pointer to a valid KoordinatorReservationResourceClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewKoordinatorReservationResourceClassFromValue(v string) (*KoordinatorReservationResourceClass, error) {
	ev := KoordinatorReservationResourceClass(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for KoordinatorReservationResourceClass: valid values are %v", v, allowedKoordinatorReservationResourceClassEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v KoordinatorReservationResourceClass) IsValid() bool {
	for _, existing := range allowedKoordinatorReservationResourceClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KoordinatorReservationResourceClass value.
func (v KoordinatorReservationResourceClass) Ptr() *KoordinatorReservationResourceClass {
	return &v
}
