// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// TlsName Name of the TLS certificates
type TlsName string

// List of TlsName.
const (
	TlsNameCa          TlsName = "ca"
	TlsNameCertificate TlsName = "certificate"
	TlsNamePrivateKey  TlsName = "privateKey"
	TlsNameKeystore    TlsName = "keystore"
	TlsNameTruststore  TlsName = "truststore"
)

var allowedTlsNameEnumValues = []TlsName{
	TlsNameCa,
	TlsNameCertificate,
	TlsNamePrivateKey,
	TlsNameKeystore,
	TlsNameTruststore,
}

// GetAllowedValues returns the list of possible values.
func (v *TlsName) GetAllowedValues() []TlsName {
	return allowedTlsNameEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *TlsName) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = TlsName(value)
	return nil
}

// NewTlsNameFromValue returns a pointer to a valid TlsName
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewTlsNameFromValue(v string) (*TlsName, error) {
	ev := TlsName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TlsName: valid values are %v", v, allowedTlsNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v TlsName) IsValid() bool {
	for _, existing := range allowedTlsNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TlsName value.
func (v TlsName) Ptr() *TlsName {
	return &v
}
