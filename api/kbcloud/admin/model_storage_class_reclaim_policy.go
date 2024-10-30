// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// StorageClassReclaimPolicy reclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
type StorageClassReclaimPolicy string

// List of StorageClassReclaimPolicy.
const (
	StorageClassReclaimPolicyDelete StorageClassReclaimPolicy = "Delete"
	StorageClassReclaimPolicyRetain StorageClassReclaimPolicy = "Retain"
	StorageClassReclaimPolicyRecycle StorageClassReclaimPolicy = "Recycle"
)

var allowedStorageClassReclaimPolicyEnumValues = []StorageClassReclaimPolicy{
	StorageClassReclaimPolicyDelete,
	StorageClassReclaimPolicyRetain,
	StorageClassReclaimPolicyRecycle,
}

// GetAllowedValues returns the list of possible values.
func (v *StorageClassReclaimPolicy) GetAllowedValues() []StorageClassReclaimPolicy {
	return allowedStorageClassReclaimPolicyEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StorageClassReclaimPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StorageClassReclaimPolicy(value)
	return nil
}

// NewStorageClassReclaimPolicyFromValue returns a pointer to a valid StorageClassReclaimPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStorageClassReclaimPolicyFromValue(v string) (*StorageClassReclaimPolicy, error) {
	ev := StorageClassReclaimPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StorageClassReclaimPolicy: valid values are %v", v, allowedStorageClassReclaimPolicyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StorageClassReclaimPolicy) IsValid() bool {
	for _, existing := range allowedStorageClassReclaimPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageClassReclaimPolicy value.
func (v StorageClassReclaimPolicy) Ptr() *StorageClassReclaimPolicy {
	return &v
}
