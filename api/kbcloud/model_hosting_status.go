// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// HostingStatus Hosting status (Hostable, Non-hostable, Hosted). When hosting_status is Hosted, cluster_info will be returned
type HostingStatus string

// List of HostingStatus.
const (
	HostingStatusHostable    HostingStatus = "Hostable"
	HostingStatusNonHostable HostingStatus = "Non-hostable"
	HostingStatusHosted      HostingStatus = "Hosted"
)

var allowedHostingStatusEnumValues = []HostingStatus{
	HostingStatusHostable,
	HostingStatusNonHostable,
	HostingStatusHosted,
}

// GetAllowedValues returns the list of possible values.
func (v *HostingStatus) GetAllowedValues() []HostingStatus {
	return allowedHostingStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *HostingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = HostingStatus(value)
	return nil
}

// NewHostingStatusFromValue returns a pointer to a valid HostingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewHostingStatusFromValue(v string) (*HostingStatus, error) {
	ev := HostingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for HostingStatus: valid values are %v", v, allowedHostingStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v HostingStatus) IsValid() bool {
	for _, existing := range allowedHostingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostingStatus value.
func (v HostingStatus) Ptr() *HostingStatus {
	return &v
}
