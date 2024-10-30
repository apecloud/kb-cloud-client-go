// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



type ManifestType string

// List of ManifestType.
const (
	ManifestTypeService ManifestType = "service"
	ManifestTypeConfigmap ManifestType = "configmap"
	ManifestTypeSecret ManifestType = "secret"
	ManifestTypeEvent ManifestType = "event"
	ManifestTypeOpsrequest ManifestType = "opsrequest"
	ManifestTypeConfigconstraint ManifestType = "configconstraint"
)

var allowedManifestTypeEnumValues = []ManifestType{
	ManifestTypeService,
	ManifestTypeConfigmap,
	ManifestTypeSecret,
	ManifestTypeEvent,
	ManifestTypeOpsrequest,
	ManifestTypeConfigconstraint,
}

// GetAllowedValues returns the list of possible values.
func (v *ManifestType) GetAllowedValues() []ManifestType {
	return allowedManifestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ManifestType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ManifestType(value)
	return nil
}

// NewManifestTypeFromValue returns a pointer to a valid ManifestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewManifestTypeFromValue(v string) (*ManifestType, error) {
	ev := ManifestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ManifestType: valid values are %v", v, allowedManifestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ManifestType) IsValid() bool {
	for _, existing := range allowedManifestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManifestType value.
func (v ManifestType) Ptr() *ManifestType {
	return &v
}
