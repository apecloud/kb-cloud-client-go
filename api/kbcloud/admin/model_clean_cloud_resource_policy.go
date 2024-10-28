// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// CleanCloudResourcePolicy The policy to clean cloud resources, either `Delete` or `Retain`
type CleanCloudResourcePolicy string

// List of CleanCloudResourcePolicy.
const (
	CLEANCLOUDRESOURCEPOLICY_DELETE CleanCloudResourcePolicy = "Delete"
	CLEANCLOUDRESOURCEPOLICY_RETAIN CleanCloudResourcePolicy = "Retain"
)

var allowedCleanCloudResourcePolicyEnumValues = []CleanCloudResourcePolicy{
	CLEANCLOUDRESOURCEPOLICY_DELETE,
	CLEANCLOUDRESOURCEPOLICY_RETAIN,
}

// GetAllowedValues returns the list of possible values.
func (v *CleanCloudResourcePolicy) GetAllowedValues() []CleanCloudResourcePolicy {
	return allowedCleanCloudResourcePolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CleanCloudResourcePolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CleanCloudResourcePolicy(value)
	return nil
}

// NewCleanCloudResourcePolicyFromValue returns a pointer to a valid CleanCloudResourcePolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCleanCloudResourcePolicyFromValue(v string) (*CleanCloudResourcePolicy, error) {
	ev := CleanCloudResourcePolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CleanCloudResourcePolicy: valid values are %v", v, allowedCleanCloudResourcePolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CleanCloudResourcePolicy) IsValid() bool {
	for _, existing := range allowedCleanCloudResourcePolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CleanCloudResourcePolicy value.
func (v CleanCloudResourcePolicy) Ptr() *CleanCloudResourcePolicy {
	return &v
}
