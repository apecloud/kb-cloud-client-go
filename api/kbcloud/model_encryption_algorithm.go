// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// EncryptionAlgorithm the encryption algorithm
type EncryptionAlgorithm string

// List of EncryptionAlgorithm.
const (
	EncryptionAlgorithmAes128Cfb EncryptionAlgorithm = "AES-128-CFB"
	EncryptionAlgorithmAes192Cfb EncryptionAlgorithm = "AES-192-CFB"
	EncryptionAlgorithmAes256Cfb EncryptionAlgorithm = "AES-256-CFB"
)

var allowedEncryptionAlgorithmEnumValues = []EncryptionAlgorithm{
	EncryptionAlgorithmAes128Cfb,
	EncryptionAlgorithmAes192Cfb,
	EncryptionAlgorithmAes256Cfb,
}

// GetAllowedValues returns the list of possible values.
func (v *EncryptionAlgorithm) GetAllowedValues() []EncryptionAlgorithm {
	return allowedEncryptionAlgorithmEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *EncryptionAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = EncryptionAlgorithm(value)
	return nil
}

// NewEncryptionAlgorithmFromValue returns a pointer to a valid EncryptionAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewEncryptionAlgorithmFromValue(v string) (*EncryptionAlgorithm, error) {
	ev := EncryptionAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EncryptionAlgorithm: valid values are %v", v, allowedEncryptionAlgorithmEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v EncryptionAlgorithm) IsValid() bool {
	for _, existing := range allowedEncryptionAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EncryptionAlgorithm value.
func (v EncryptionAlgorithm) Ptr() *EncryptionAlgorithm {
	return &v
}
