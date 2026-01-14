// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EncryptionKeyUsage The intended usage of the key.
type EncryptionKeyUsage string

// List of EncryptionKeyUsage.
const (
	EncryptionKeyUsageEncryption EncryptionKeyUsage = "encryption"
)

var allowedEncryptionKeyUsageEnumValues = []EncryptionKeyUsage{
	EncryptionKeyUsageEncryption,
}

// GetAllowedValues returns the list of possible values.
func (v *EncryptionKeyUsage) GetAllowedValues() []EncryptionKeyUsage {
	return allowedEncryptionKeyUsageEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EncryptionKeyUsage) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EncryptionKeyUsage(value)
	return nil
}

// NewEncryptionKeyUsageFromValue returns a pointer to a valid EncryptionKeyUsage
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEncryptionKeyUsageFromValue(v string) (*EncryptionKeyUsage, error) {
	ev := EncryptionKeyUsage(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EncryptionKeyUsage: valid values are %v", v, allowedEncryptionKeyUsageEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EncryptionKeyUsage) IsValid() bool {
	for _, existing := range allowedEncryptionKeyUsageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EncryptionKeyUsage value.
func (v EncryptionKeyUsage) Ptr() *EncryptionKeyUsage {
	return &v
}
