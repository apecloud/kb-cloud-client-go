// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EsrallyTelemetry Rally runtime telemetry device
type EsrallyTelemetry string

// List of EsrallyTelemetry.
const (
	EsrallyTelemetryNodeStats                EsrallyTelemetry = "node-stats"
	EsrallyTelemetryRecoveryStats            EsrallyTelemetry = "recovery-stats"
	EsrallyTelemetryCcrStats                 EsrallyTelemetry = "ccr-stats"
	EsrallyTelemetrySegmentStats             EsrallyTelemetry = "segment-stats"
	EsrallyTelemetryTransformStats           EsrallyTelemetry = "transform-stats"
	EsrallyTelemetrySearchableSnapshotsStats EsrallyTelemetry = "searchable-snapshots-stats"
	EsrallyTelemetryShardStats               EsrallyTelemetry = "shard-stats"
	EsrallyTelemetryDataStreamStats          EsrallyTelemetry = "data-stream-stats"
	EsrallyTelemetryIngestPipelineStats      EsrallyTelemetry = "ingest-pipeline-stats"
	EsrallyTelemetryDiskUsageStats           EsrallyTelemetry = "disk-usage-stats"
	EsrallyTelemetryGeoipStats               EsrallyTelemetry = "geoip-stats"
)

var allowedEsrallyTelemetryEnumValues = []EsrallyTelemetry{
	EsrallyTelemetryNodeStats,
	EsrallyTelemetryRecoveryStats,
	EsrallyTelemetryCcrStats,
	EsrallyTelemetrySegmentStats,
	EsrallyTelemetryTransformStats,
	EsrallyTelemetrySearchableSnapshotsStats,
	EsrallyTelemetryShardStats,
	EsrallyTelemetryDataStreamStats,
	EsrallyTelemetryIngestPipelineStats,
	EsrallyTelemetryDiskUsageStats,
	EsrallyTelemetryGeoipStats,
}

// GetAllowedValues returns the list of possible values.
func (v *EsrallyTelemetry) GetAllowedValues() []EsrallyTelemetry {
	return allowedEsrallyTelemetryEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EsrallyTelemetry) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EsrallyTelemetry(value)
	return nil
}

// NewEsrallyTelemetryFromValue returns a pointer to a valid EsrallyTelemetry
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEsrallyTelemetryFromValue(v string) (*EsrallyTelemetry, error) {
	ev := EsrallyTelemetry(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EsrallyTelemetry: valid values are %v", v, allowedEsrallyTelemetryEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EsrallyTelemetry) IsValid() bool {
	for _, existing := range allowedEsrallyTelemetryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EsrallyTelemetry value.
func (v EsrallyTelemetry) Ptr() *EsrallyTelemetry {
	return &v
}
