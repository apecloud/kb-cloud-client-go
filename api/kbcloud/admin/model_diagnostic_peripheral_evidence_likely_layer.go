// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticPeripheralEvidenceLikelyLayer Likely layer indicated by the external evidence. This is not a database-internal root-cause claim.
type DiagnosticPeripheralEvidenceLikelyLayer string

// List of DiagnosticPeripheralEvidenceLikelyLayer.
const (
	DiagnosticPeripheralEvidenceLikelyLayerK8s             DiagnosticPeripheralEvidenceLikelyLayer = "k8s"
	DiagnosticPeripheralEvidenceLikelyLayerStorage         DiagnosticPeripheralEvidenceLikelyLayer = "storage"
	DiagnosticPeripheralEvidenceLikelyLayerNetwork         DiagnosticPeripheralEvidenceLikelyLayer = "network"
	DiagnosticPeripheralEvidenceLikelyLayerKubeblocks      DiagnosticPeripheralEvidenceLikelyLayer = "kubeblocks"
	DiagnosticPeripheralEvidenceLikelyLayerDatabaseStartup DiagnosticPeripheralEvidenceLikelyLayer = "database_startup"
)

var allowedDiagnosticPeripheralEvidenceLikelyLayerEnumValues = []DiagnosticPeripheralEvidenceLikelyLayer{
	DiagnosticPeripheralEvidenceLikelyLayerK8s,
	DiagnosticPeripheralEvidenceLikelyLayerStorage,
	DiagnosticPeripheralEvidenceLikelyLayerNetwork,
	DiagnosticPeripheralEvidenceLikelyLayerKubeblocks,
	DiagnosticPeripheralEvidenceLikelyLayerDatabaseStartup,
}

// GetAllowedValues returns the list of possible values.
func (v *DiagnosticPeripheralEvidenceLikelyLayer) GetAllowedValues() []DiagnosticPeripheralEvidenceLikelyLayer {
	return allowedDiagnosticPeripheralEvidenceLikelyLayerEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DiagnosticPeripheralEvidenceLikelyLayer) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DiagnosticPeripheralEvidenceLikelyLayer(value)
	return nil
}

// NewDiagnosticPeripheralEvidenceLikelyLayerFromValue returns a pointer to a valid DiagnosticPeripheralEvidenceLikelyLayer
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDiagnosticPeripheralEvidenceLikelyLayerFromValue(v string) (*DiagnosticPeripheralEvidenceLikelyLayer, error) {
	ev := DiagnosticPeripheralEvidenceLikelyLayer(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DiagnosticPeripheralEvidenceLikelyLayer: valid values are %v", v, allowedDiagnosticPeripheralEvidenceLikelyLayerEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DiagnosticPeripheralEvidenceLikelyLayer) IsValid() bool {
	for _, existing := range allowedDiagnosticPeripheralEvidenceLikelyLayerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DiagnosticPeripheralEvidenceLikelyLayer value.
func (v DiagnosticPeripheralEvidenceLikelyLayer) Ptr() *DiagnosticPeripheralEvidenceLikelyLayer {
	return &v
}
