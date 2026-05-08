// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ClusterSchedulingPolicy * `HardAntiAffinity` - Strictly enforce pod anti-affinity across nodes. Pods will not be scheduled when the anti-affinity constraints cannot be satisfied.
// * `SoftAntiAffinity` - Prefer to spread pods across nodes using anti-affinity, but allow scheduling on the same node when constraints cannot be fully satisfied.
// * `Disabled` - Do not apply pod anti-affinity constraints on nodes.
// * `None` - Inherit the engine-level default scheduling policy.
type ClusterSchedulingPolicy string

// List of ClusterSchedulingPolicy.
const (
	ClusterSchedulingPolicyHardAntiAffinity ClusterSchedulingPolicy = "HardAntiAffinity"
	ClusterSchedulingPolicySoftAntiAffinity ClusterSchedulingPolicy = "SoftAntiAffinity"
	ClusterSchedulingPolicyDisabled         ClusterSchedulingPolicy = "Disabled"
	ClusterSchedulingPolicyNon              ClusterSchedulingPolicy = "None"
)

var allowedClusterSchedulingPolicyEnumValues = []ClusterSchedulingPolicy{
	ClusterSchedulingPolicyHardAntiAffinity,
	ClusterSchedulingPolicySoftAntiAffinity,
	ClusterSchedulingPolicyDisabled,
	ClusterSchedulingPolicyNon,
}

// GetAllowedValues returns the list of possible values.
func (v *ClusterSchedulingPolicy) GetAllowedValues() []ClusterSchedulingPolicy {
	return allowedClusterSchedulingPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ClusterSchedulingPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ClusterSchedulingPolicy(value)
	return nil
}

// NewClusterSchedulingPolicyFromValue returns a pointer to a valid ClusterSchedulingPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewClusterSchedulingPolicyFromValue(v string) (*ClusterSchedulingPolicy, error) {
	ev := ClusterSchedulingPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ClusterSchedulingPolicy: valid values are %v", v, allowedClusterSchedulingPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ClusterSchedulingPolicy) IsValid() bool {
	for _, existing := range allowedClusterSchedulingPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterSchedulingPolicy value.
func (v ClusterSchedulingPolicy) Ptr() *ClusterSchedulingPolicy {
	return &v
}
