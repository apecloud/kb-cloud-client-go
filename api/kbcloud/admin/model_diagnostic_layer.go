// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticLayer Diagnostic report layer.
type DiagnosticLayer string

// List of DiagnosticLayer.
const (
	DiagnosticLayerPlatform DiagnosticLayer = "platform"
	DiagnosticLayerKbeKb    DiagnosticLayer = "kbe-kb"
	DiagnosticLayerDatabase DiagnosticLayer = "database"
)

var allowedDiagnosticLayerEnumValues = []DiagnosticLayer{
	DiagnosticLayerPlatform,
	DiagnosticLayerKbeKb,
	DiagnosticLayerDatabase,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticLayer) GetAllowedValues() []DiagnosticLayer {
	return allowedDiagnosticLayerEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticLayer) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticLayer(value)
	return nil
}

// NewDiagnosticLayerFromValue returns a pointer to a valid DiagnosticLayer
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticLayerFromValue(v string) (*DiagnosticLayer, error) {
	ev := DiagnosticLayer(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticLayer: valid values are %v", v, allowedDiagnosticLayerEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticLayer) IsValid() bool {
	for _, existing := range allowedDiagnosticLayerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticLayer value.
func (v DiagnosticLayer) Ptr() *DiagnosticLayer {
	return &v
}
