// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// StaticClusterComponentName Exact component name for the selected engine. Each component may appear only once.
type StaticClusterComponentName string

// List of StaticClusterComponentName.
const (
	StaticClusterComponentNameVmStorage StaticClusterComponentName = "vm-storage"
	StaticClusterComponentNameVmSelect  StaticClusterComponentName = "vm-select"
	StaticClusterComponentNameVmInsert  StaticClusterComponentName = "vm-insert"
	StaticClusterComponentNameVlstorage StaticClusterComponentName = "vlstorage"
	StaticClusterComponentNameVlselect  StaticClusterComponentName = "vlselect"
	StaticClusterComponentNameVlinsert  StaticClusterComponentName = "vlinsert"
	StaticClusterComponentNameMinio     StaticClusterComponentName = "minio"
)

var allowedStaticClusterComponentNameEnumValues = []StaticClusterComponentName{
	StaticClusterComponentNameVmStorage,
	StaticClusterComponentNameVmSelect,
	StaticClusterComponentNameVmInsert,
	StaticClusterComponentNameVlstorage,
	StaticClusterComponentNameVlselect,
	StaticClusterComponentNameVlinsert,
	StaticClusterComponentNameMinio,
}

// GetAllowedValues returns the list of possible values.
func (v *StaticClusterComponentName) GetAllowedValues() []StaticClusterComponentName {
	return allowedStaticClusterComponentNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *StaticClusterComponentName) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = StaticClusterComponentName(value)
	return nil
}

// NewStaticClusterComponentNameFromValue returns a pointer to a valid StaticClusterComponentName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewStaticClusterComponentNameFromValue(v string) (*StaticClusterComponentName, error) {
	ev := StaticClusterComponentName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for StaticClusterComponentName: valid values are %v", v, allowedStaticClusterComponentNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v StaticClusterComponentName) IsValid() bool {
	for _, existing := range allowedStaticClusterComponentNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StaticClusterComponentName value.
func (v StaticClusterComponentName) Ptr() *StaticClusterComponentName {
	return &v
}
