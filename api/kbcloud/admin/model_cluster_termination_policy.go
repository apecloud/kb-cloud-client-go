// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterTerminationPolicy The termination policy of cluster.
type ClusterTerminationPolicy string

// List of ClusterTerminationPolicy.
const (
	CLUSTERTERMINATIONPOLICY_DONOTTERMINATE ClusterTerminationPolicy = "DoNotTerminate"
	CLUSTERTERMINATIONPOLICY_HALT           ClusterTerminationPolicy = "Halt"
	CLUSTERTERMINATIONPOLICY_DELETE         ClusterTerminationPolicy = "Delete"
	CLUSTERTERMINATIONPOLICY_WIPEOUT        ClusterTerminationPolicy = "WipeOut"
)

var allowedClusterTerminationPolicyEnumValues = []ClusterTerminationPolicy{
	CLUSTERTERMINATIONPOLICY_DONOTTERMINATE,
	CLUSTERTERMINATIONPOLICY_HALT,
	CLUSTERTERMINATIONPOLICY_DELETE,
	CLUSTERTERMINATIONPOLICY_WIPEOUT,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ClusterTerminationPolicy) GetAllowedValues() []ClusterTerminationPolicy {
	return allowedClusterTerminationPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClusterTerminationPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClusterTerminationPolicy(value)
	return nil
}

// NewClusterTerminationPolicyFromValue returns a pointer to a valid ClusterTerminationPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterTerminationPolicyFromValue(v string) (*ClusterTerminationPolicy, error) {
	ev := ClusterTerminationPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClusterTerminationPolicy: valid values are %v", v, allowedClusterTerminationPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterTerminationPolicy) IsValid() bool {
	for _, existing := range allowedClusterTerminationPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterTerminationPolicy value.
func (v ClusterTerminationPolicy) Ptr() *ClusterTerminationPolicy {
	return &v
}
