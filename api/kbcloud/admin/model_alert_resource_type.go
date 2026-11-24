// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AlertResourceType string

// List of AlertResourceType.
const (
	AlertResourceTypeCluster AlertResourceType = "cluster"
	AlertResourceTypeChannel AlertResourceType = "channel"
)

var allowedAlertResourceTypeEnumValues = []AlertResourceType{
	AlertResourceTypeCluster,
	AlertResourceTypeChannel,
}

// GetAllowedValues returns the list of possible values.
func (v *AlertResourceType) GetAllowedValues() []AlertResourceType {
	return allowedAlertResourceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AlertResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AlertResourceType(value)
	return nil
}

// NewAlertResourceTypeFromValue returns a pointer to a valid AlertResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAlertResourceTypeFromValue(v string) (*AlertResourceType, error) {
	ev := AlertResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AlertResourceType: valid values are %v", v, allowedAlertResourceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AlertResourceType) IsValid() bool {
	for _, existing := range allowedAlertResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertResourceType value.
func (v AlertResourceType) Ptr() *AlertResourceType {
	return &v
}
