// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EsrallyDataProfile Generated Elasticsearch dataset shape
type EsrallyDataProfile string

// List of EsrallyDataProfile.
const (
	EsrallyDataProfileLogs        EsrallyDataProfile = "logs"
	EsrallyDataProfileMetrics     EsrallyDataProfile = "metrics"
	EsrallyDataProfileHttpLogs    EsrallyDataProfile = "http_logs"
	EsrallyDataProfileMetricbeat  EsrallyDataProfile = "metricbeat"
	EsrallyDataProfileGeonames    EsrallyDataProfile = "geonames"
	EsrallyDataProfileNycTaxis    EsrallyDataProfile = "nyc_taxis"
	EsrallyDataProfileNoaa        EsrallyDataProfile = "noaa"
	EsrallyDataProfileNested      EsrallyDataProfile = "nested"
	EsrallyDataProfilePmc         EsrallyDataProfile = "pmc"
	EsrallyDataProfileSo          EsrallyDataProfile = "so"
	EsrallyDataProfileDenseVector EsrallyDataProfile = "dense_vector"
)

var allowedEsrallyDataProfileEnumValues = []EsrallyDataProfile{
	EsrallyDataProfileLogs,
	EsrallyDataProfileMetrics,
	EsrallyDataProfileHttpLogs,
	EsrallyDataProfileMetricbeat,
	EsrallyDataProfileGeonames,
	EsrallyDataProfileNycTaxis,
	EsrallyDataProfileNoaa,
	EsrallyDataProfileNested,
	EsrallyDataProfilePmc,
	EsrallyDataProfileSo,
	EsrallyDataProfileDenseVector,
}

// GetAllowedValues returns the list of possible values.
func (v *EsrallyDataProfile) GetAllowedValues() []EsrallyDataProfile {
	return allowedEsrallyDataProfileEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EsrallyDataProfile) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EsrallyDataProfile(value)
	return nil
}

// NewEsrallyDataProfileFromValue returns a pointer to a valid EsrallyDataProfile
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEsrallyDataProfileFromValue(v string) (*EsrallyDataProfile, error) {
	ev := EsrallyDataProfile(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EsrallyDataProfile: valid values are %v", v, allowedEsrallyDataProfileEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EsrallyDataProfile) IsValid() bool {
	for _, existing := range allowedEsrallyDataProfileEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EsrallyDataProfile value.
func (v EsrallyDataProfile) Ptr() *EsrallyDataProfile {
	return &v
}
