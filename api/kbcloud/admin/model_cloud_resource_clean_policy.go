// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// CloudResourceCleanPolicy The policy to clean cloud resources, either `Delete` or `Retain`
type CloudResourceCleanPolicy string

// List of CloudResourceCleanPolicy.
const (
	CLOUDRESOURCECLEANPOLICY_DELETE CloudResourceCleanPolicy = "Delete"
	CLOUDRESOURCECLEANPOLICY_RETAIN CloudResourceCleanPolicy = "Retain"
)

var allowedCloudResourceCleanPolicyEnumValues = []CloudResourceCleanPolicy{
	CLOUDRESOURCECLEANPOLICY_DELETE,
	CLOUDRESOURCECLEANPOLICY_RETAIN,
}

// GetAllowedValues returns the list of possible values.
func (v *CloudResourceCleanPolicy) GetAllowedValues() []CloudResourceCleanPolicy {
	return allowedCloudResourceCleanPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CloudResourceCleanPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CloudResourceCleanPolicy(value)
	return nil
}

// NewCloudResourceCleanPolicyFromValue returns a pointer to a valid CloudResourceCleanPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCloudResourceCleanPolicyFromValue(v string) (*CloudResourceCleanPolicy, error) {
	ev := CloudResourceCleanPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CloudResourceCleanPolicy: valid values are %v", v, allowedCloudResourceCleanPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CloudResourceCleanPolicy) IsValid() bool {
	for _, existing := range allowedCloudResourceCleanPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudResourceCleanPolicy value.
func (v CloudResourceCleanPolicy) Ptr() *CloudResourceCleanPolicy {
	return &v
}
