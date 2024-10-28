// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// FeaturePreReleaseType Indicates the maturity level of a feature
type FeaturePreReleaseType string

// List of FeaturePreReleaseType.
const (
	FEATUREPRERELEASETYPE_ALPHA      FeaturePreReleaseType = "ALPHA"
	FEATUREPRERELEASETYPE_BETA       FeaturePreReleaseType = "BETA"
	FEATUREPRERELEASETYPE_           FeaturePreReleaseType = ""
	FEATUREPRERELEASETYPE_DEPRECATED FeaturePreReleaseType = "DEPRECATED"
)

var allowedFeaturePreReleaseTypeEnumValues = []FeaturePreReleaseType{
	FEATUREPRERELEASETYPE_ALPHA,
	FEATUREPRERELEASETYPE_BETA,
	FEATUREPRERELEASETYPE_,
	FEATUREPRERELEASETYPE_DEPRECATED,
}

// GetAllowedValues returns the list of possible values.
func (v *FeaturePreReleaseType) GetAllowedValues() []FeaturePreReleaseType {
	return allowedFeaturePreReleaseTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *FeaturePreReleaseType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = FeaturePreReleaseType(value)
	return nil
}

// NewFeaturePreReleaseTypeFromValue returns a pointer to a valid FeaturePreReleaseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewFeaturePreReleaseTypeFromValue(v string) (*FeaturePreReleaseType, error) {
	ev := FeaturePreReleaseType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for FeaturePreReleaseType: valid values are %v", v, allowedFeaturePreReleaseTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v FeaturePreReleaseType) IsValid() bool {
	for _, existing := range allowedFeaturePreReleaseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeaturePreReleaseType value.
func (v FeaturePreReleaseType) Ptr() *FeaturePreReleaseType {
	return &v
}
