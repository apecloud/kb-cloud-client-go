// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DmsServiceAccountStatus MinIO service account status.
type DmsServiceAccountStatus string

// List of DmsServiceAccountStatus.
const (
	DmsServiceAccountStatusEnabled  DmsServiceAccountStatus = "enabled"
	DmsServiceAccountStatusDisabled DmsServiceAccountStatus = "disabled"
)

var allowedDmsServiceAccountStatusEnumValues = []DmsServiceAccountStatus{
	DmsServiceAccountStatusEnabled,
	DmsServiceAccountStatusDisabled,
}

// GetAllowedValues returns the list of possible values.
func (v *DmsServiceAccountStatus) GetAllowedValues() []DmsServiceAccountStatus {
	return allowedDmsServiceAccountStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DmsServiceAccountStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DmsServiceAccountStatus(value)
	return nil
}

// NewDmsServiceAccountStatusFromValue returns a pointer to a valid DmsServiceAccountStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDmsServiceAccountStatusFromValue(v string) (*DmsServiceAccountStatus, error) {
	ev := DmsServiceAccountStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DmsServiceAccountStatus: valid values are %v", v, allowedDmsServiceAccountStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DmsServiceAccountStatus) IsValid() bool {
	for _, existing := range allowedDmsServiceAccountStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DmsServiceAccountStatus value.
func (v DmsServiceAccountStatus) Ptr() *DmsServiceAccountStatus {
	return &v
}
