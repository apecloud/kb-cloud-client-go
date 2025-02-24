// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DmsExportFormat the file format for export data
type DmsExportFormat string

// List of DmsExportFormat.
const (
	DmsExportFormatCsv  DmsExportFormat = "csv"
	DmsExportFormatJson DmsExportFormat = "json"
	DmsExportFormatXml  DmsExportFormat = "xml"
)

var allowedDmsExportFormatEnumValues = []DmsExportFormat{
	DmsExportFormatCsv,
	DmsExportFormatJson,
	DmsExportFormatXml,
}

// GetAllowedValues returns the list of possible values.
func (v *DmsExportFormat) GetAllowedValues() []DmsExportFormat {
	return allowedDmsExportFormatEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DmsExportFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DmsExportFormat(value)
	return nil
}

// NewDmsExportFormatFromValue returns a pointer to a valid DmsExportFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDmsExportFormatFromValue(v string) (*DmsExportFormat, error) {
	ev := DmsExportFormat(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DmsExportFormat: valid values are %v", v, allowedDmsExportFormatEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DmsExportFormat) IsValid() bool {
	for _, existing := range allowedDmsExportFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DmsExportFormat value.
func (v DmsExportFormat) Ptr() *DmsExportFormat {
	return &v
}
