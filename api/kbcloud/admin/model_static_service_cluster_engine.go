// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StaticServiceClusterEngine Engine whose retained clusters should be listed. Candidates are discovered from the kubeconfig and existing cluster records. The caller must select one; the server does not choose a default.
type StaticServiceClusterEngine string

// List of StaticServiceClusterEngine.
const (
	StaticServiceClusterEngineMinio           StaticServiceClusterEngine = "minio"
	StaticServiceClusterEngineVictoriaMetrics StaticServiceClusterEngine = "victoria-metrics"
	StaticServiceClusterEngineVictoriaLogs    StaticServiceClusterEngine = "victoria-logs"
)

var allowedStaticServiceClusterEngineEnumValues = []StaticServiceClusterEngine{
	StaticServiceClusterEngineMinio,
	StaticServiceClusterEngineVictoriaMetrics,
	StaticServiceClusterEngineVictoriaLogs,
}

// GetAllowedValues returns the list of possible values.
func (v *StaticServiceClusterEngine) GetAllowedValues() []StaticServiceClusterEngine {
	return allowedStaticServiceClusterEngineEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StaticServiceClusterEngine) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StaticServiceClusterEngine(value)
	return nil
}

// NewStaticServiceClusterEngineFromValue returns a pointer to a valid StaticServiceClusterEngine
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStaticServiceClusterEngineFromValue(v string) (*StaticServiceClusterEngine, error) {
	ev := StaticServiceClusterEngine(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StaticServiceClusterEngine: valid values are %v", v, allowedStaticServiceClusterEngineEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StaticServiceClusterEngine) IsValid() bool {
	for _, existing := range allowedStaticServiceClusterEngineEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StaticServiceClusterEngine value.
func (v StaticServiceClusterEngine) Ptr() *StaticServiceClusterEngine {
	return &v
}
