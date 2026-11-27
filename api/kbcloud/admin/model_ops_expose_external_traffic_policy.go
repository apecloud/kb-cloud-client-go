// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// OpsExposeExternalTrafficPolicy Updates existing LoadBalancer or NodePort exposure endpoints on KubeBlocks 1.0. At least one matching endpoint must already exist. Omit when creating an exposure, which uses the platform parameter external_traffic_policy. Existing endpoints keep their policy when this field is omitted.
type OpsExposeExternalTrafficPolicy string

// List of OpsExposeExternalTrafficPolicy.
const (
	OpsExposeExternalTrafficPolicyLocal   OpsExposeExternalTrafficPolicy = "Local"
	OpsExposeExternalTrafficPolicyCluster OpsExposeExternalTrafficPolicy = "Cluster"
)

var allowedOpsExposeExternalTrafficPolicyEnumValues = []OpsExposeExternalTrafficPolicy{
	OpsExposeExternalTrafficPolicyLocal,
	OpsExposeExternalTrafficPolicyCluster,
}

// GetAllowedValues returns the list of possible values.
func (v *OpsExposeExternalTrafficPolicy) GetAllowedValues() []OpsExposeExternalTrafficPolicy {
	return allowedOpsExposeExternalTrafficPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OpsExposeExternalTrafficPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OpsExposeExternalTrafficPolicy(value)
	return nil
}

// NewOpsExposeExternalTrafficPolicyFromValue returns a pointer to a valid OpsExposeExternalTrafficPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOpsExposeExternalTrafficPolicyFromValue(v string) (*OpsExposeExternalTrafficPolicy, error) {
	ev := OpsExposeExternalTrafficPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OpsExposeExternalTrafficPolicy: valid values are %v", v, allowedOpsExposeExternalTrafficPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OpsExposeExternalTrafficPolicy) IsValid() bool {
	for _, existing := range allowedOpsExposeExternalTrafficPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OpsExposeExternalTrafficPolicy value.
func (v OpsExposeExternalTrafficPolicy) Ptr() *OpsExposeExternalTrafficPolicy {
	return &v
}

// NullableOpsExposeExternalTrafficPolicy handles when a null is used for OpsExposeExternalTrafficPolicy.
type NullableOpsExposeExternalTrafficPolicy struct {
	value *OpsExposeExternalTrafficPolicy
	isSet bool
}

// Get returns the associated value.
func (v NullableOpsExposeExternalTrafficPolicy) Get() *OpsExposeExternalTrafficPolicy {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableOpsExposeExternalTrafficPolicy) Set(val *OpsExposeExternalTrafficPolicy) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableOpsExposeExternalTrafficPolicy) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableOpsExposeExternalTrafficPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOpsExposeExternalTrafficPolicy initializes the struct as if Set has been called.
func NewNullableOpsExposeExternalTrafficPolicy(val *OpsExposeExternalTrafficPolicy) *NullableOpsExposeExternalTrafficPolicy {
	return &NullableOpsExposeExternalTrafficPolicy{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableOpsExposeExternalTrafficPolicy) MarshalJSON() ([]byte, error) {
	return common.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableOpsExposeExternalTrafficPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return common.Unmarshal(src, &v.value)
}
