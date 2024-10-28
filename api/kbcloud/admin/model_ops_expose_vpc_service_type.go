// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsExposeVPCServiceType Specifies the type of service for the KubeBlocks cluster.
type OpsExposeVPCServiceType string

// List of OpsExposeVPCServiceType.
const (
	OPSEXPOSEVPCSERVICETYPE_LOADBALANCER OpsExposeVPCServiceType = "LoadBalancer"
	OPSEXPOSEVPCSERVICETYPE_NODEPORT     OpsExposeVPCServiceType = "NodePort"
)

var allowedOpsExposeVPCServiceTypeEnumValues = []OpsExposeVPCServiceType{
	OPSEXPOSEVPCSERVICETYPE_LOADBALANCER,
	OPSEXPOSEVPCSERVICETYPE_NODEPORT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OpsExposeVPCServiceType) GetAllowedValues() []OpsExposeVPCServiceType {
	return allowedOpsExposeVPCServiceTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpsExposeVPCServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpsExposeVPCServiceType(value)
	return nil
}

// NewOpsExposeVPCServiceTypeFromValue returns a pointer to a valid OpsExposeVPCServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpsExposeVPCServiceTypeFromValue(v string) (*OpsExposeVPCServiceType, error) {
	ev := OpsExposeVPCServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpsExposeVPCServiceType: valid values are %v", v, allowedOpsExposeVPCServiceTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpsExposeVPCServiceType) IsValid() bool {
	for _, existing := range allowedOpsExposeVPCServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpsExposeVPCServiceType value.
func (v OpsExposeVPCServiceType) Ptr() *OpsExposeVPCServiceType {
	return &v
}
