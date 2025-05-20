// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TlsType TLS type
type TlsType string

// List of TlsType.
const (
	TlsTypeStandardTls  TlsType = "StandardTLS"
	TlsTypeJavaKeystore TlsType = "JavaKeystore"
)

var allowedTlsTypeEnumValues = []TlsType{
	TlsTypeStandardTls,
	TlsTypeJavaKeystore,
}

// GetAllowedValues returns the list of possible values.
func (v *TlsType) GetAllowedValues() []TlsType {
	return allowedTlsTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TlsType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TlsType(value)
	return nil
}

// NewTlsTypeFromValue returns a pointer to a valid TlsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTlsTypeFromValue(v string) (*TlsType, error) {
	ev := TlsType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TlsType: valid values are %v", v, allowedTlsTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TlsType) IsValid() bool {
	for _, existing := range allowedTlsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TlsType value.
func (v TlsType) Ptr() *TlsType {
	return &v
}
