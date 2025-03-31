// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DMSConsoleSvcType string

// List of DMSConsoleSvcType.
const (
	DMSConsoleSvcTypeClusterIp    DMSConsoleSvcType = "ClusterIP"
	DMSConsoleSvcTypeLoadBalancer DMSConsoleSvcType = "LoadBalancer"
	DMSConsoleSvcTypeNodePort     DMSConsoleSvcType = "NodePort"
)

var allowedDMSConsoleSvcTypeEnumValues = []DMSConsoleSvcType{
	DMSConsoleSvcTypeClusterIp,
	DMSConsoleSvcTypeLoadBalancer,
	DMSConsoleSvcTypeNodePort,
}

// GetAllowedValues returns the list of possible values.
func (v *DMSConsoleSvcType) GetAllowedValues() []DMSConsoleSvcType {
	return allowedDMSConsoleSvcTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DMSConsoleSvcType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DMSConsoleSvcType(value)
	return nil
}

// NewDMSConsoleSvcTypeFromValue returns a pointer to a valid DMSConsoleSvcType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDMSConsoleSvcTypeFromValue(v string) (*DMSConsoleSvcType, error) {
	ev := DMSConsoleSvcType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DMSConsoleSvcType: valid values are %v", v, allowedDMSConsoleSvcTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DMSConsoleSvcType) IsValid() bool {
	for _, existing := range allowedDMSConsoleSvcTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DMSConsoleSvcType value.
func (v DMSConsoleSvcType) Ptr() *DMSConsoleSvcType {
	return &v
}
