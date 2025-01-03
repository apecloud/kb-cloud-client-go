// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type CdcNetworkType string

// List of CdcNetworkType.
const (
	CdcNetworkTypeClusterIp    CdcNetworkType = "ClusterIP"
	CdcNetworkTypeNodePort     CdcNetworkType = "NodePort"
	CdcNetworkTypeLoadBalancer CdcNetworkType = "LoadBalancer"
	CdcNetworkTypeFixedPodIp   CdcNetworkType = "FixedPodIP"
)

var allowedCdcNetworkTypeEnumValues = []CdcNetworkType{
	CdcNetworkTypeClusterIp,
	CdcNetworkTypeNodePort,
	CdcNetworkTypeLoadBalancer,
	CdcNetworkTypeFixedPodIp,
}

// GetAllowedValues returns the list of possible values.
func (v *CdcNetworkType) GetAllowedValues() []CdcNetworkType {
	return allowedCdcNetworkTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CdcNetworkType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CdcNetworkType(value)
	return nil
}

// NewCdcNetworkTypeFromValue returns a pointer to a valid CdcNetworkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCdcNetworkTypeFromValue(v string) (*CdcNetworkType, error) {
	ev := CdcNetworkType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CdcNetworkType: valid values are %v", v, allowedCdcNetworkTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CdcNetworkType) IsValid() bool {
	for _, existing := range allowedCdcNetworkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CdcNetworkType value.
func (v CdcNetworkType) Ptr() *CdcNetworkType {
	return &v
}
