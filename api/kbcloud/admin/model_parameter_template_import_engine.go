// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// ParameterTemplateImportEngine Target engine used to validate the component, major versions and configuration in the file.
type ParameterTemplateImportEngine string

// List of ParameterTemplateImportEngine.
const (
	ParameterTemplateImportEngineMongodb    ParameterTemplateImportEngine = "mongodb"
	ParameterTemplateImportEngineInfluxdb   ParameterTemplateImportEngine = "influxdb"
	ParameterTemplateImportEngineTdengine   ParameterTemplateImportEngine = "tdengine"
	ParameterTemplateImportEngineMssql      ParameterTemplateImportEngine = "mssql"
	ParameterTemplateImportEngineMysql      ParameterTemplateImportEngine = "mysql"
	ParameterTemplateImportEnginePostgresql ParameterTemplateImportEngine = "postgresql"
)

var allowedParameterTemplateImportEngineEnumValues = []ParameterTemplateImportEngine{
	ParameterTemplateImportEngineMongodb,
	ParameterTemplateImportEngineInfluxdb,
	ParameterTemplateImportEngineTdengine,
	ParameterTemplateImportEngineMssql,
	ParameterTemplateImportEngineMysql,
	ParameterTemplateImportEnginePostgresql,
}

// GetAllowedValues returns the list of possible values.
func (v *ParameterTemplateImportEngine) GetAllowedValues() []ParameterTemplateImportEngine {
	return allowedParameterTemplateImportEngineEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ParameterTemplateImportEngine) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ParameterTemplateImportEngine(value)
	return nil
}

// NewParameterTemplateImportEngineFromValue returns a pointer to a valid ParameterTemplateImportEngine
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewParameterTemplateImportEngineFromValue(v string) (*ParameterTemplateImportEngine, error) {
	ev := ParameterTemplateImportEngine(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ParameterTemplateImportEngine: valid values are %v", v, allowedParameterTemplateImportEngineEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ParameterTemplateImportEngine) IsValid() bool {
	for _, existing := range allowedParameterTemplateImportEngineEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ParameterTemplateImportEngine value.
func (v ParameterTemplateImportEngine) Ptr() *ParameterTemplateImportEngine {
	return &v
}
