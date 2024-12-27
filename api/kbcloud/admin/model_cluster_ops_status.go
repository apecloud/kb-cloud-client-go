// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type Cluster_opsStatus string

// List of Cluster_opsStatus.
const (
	Cluster_opsStatusPending    Cluster_opsStatus = "Pending"
	Cluster_opsStatusCreating   Cluster_opsStatus = "Creating"
	Cluster_opsStatusRunning    Cluster_opsStatus = "Running"
	Cluster_opsStatusCancelling Cluster_opsStatus = "Cancelling"
	Cluster_opsStatusSucceed    Cluster_opsStatus = "Succeed"
	Cluster_opsStatusCancelled  Cluster_opsStatus = "Cancelled"
	Cluster_opsStatusFailed     Cluster_opsStatus = "Failed"
	Cluster_opsStatusAborted    Cluster_opsStatus = "Aborted"
)

var allowedCluster_opsStatusEnumValues = []Cluster_opsStatus{
	Cluster_opsStatusPending,
	Cluster_opsStatusCreating,
	Cluster_opsStatusRunning,
	Cluster_opsStatusCancelling,
	Cluster_opsStatusSucceed,
	Cluster_opsStatusCancelled,
	Cluster_opsStatusFailed,
	Cluster_opsStatusAborted,
}

// GetAllowedValues returns the list of possible values.
func (v *Cluster_opsStatus) GetAllowedValues() []Cluster_opsStatus {
	return allowedCluster_opsStatusEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *Cluster_opsStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = Cluster_opsStatus(value)
	return nil
}

// NewCluster_opsStatusFromValue returns a pointer to a valid Cluster_opsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCluster_opsStatusFromValue(v string) (*Cluster_opsStatus, error) {
	ev := Cluster_opsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Cluster_opsStatus: valid values are %v", v, allowedCluster_opsStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v Cluster_opsStatus) IsValid() bool {
	for _, existing := range allowedCluster_opsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Cluster_opsStatus value.
func (v Cluster_opsStatus) Ptr() *Cluster_opsStatus {
	return &v
}
