// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// CertificateIssuer certificate issuer
type CertificateIssuer string

// List of CertificateIssuer.
const (
	CertificateIssuerApecloud CertificateIssuer = "Apecloud"
	CertificateIssuerUserSelf CertificateIssuer = "UserSelf"
)

var allowedCertificateIssuerEnumValues = []CertificateIssuer{
	CertificateIssuerApecloud,
	CertificateIssuerUserSelf,
}

// GetAllowedValues returns the list of possible values.
func (v *CertificateIssuer) GetAllowedValues() []CertificateIssuer {
	return allowedCertificateIssuerEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CertificateIssuer) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CertificateIssuer(value)
	return nil
}

// NewCertificateIssuerFromValue returns a pointer to a valid CertificateIssuer
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCertificateIssuerFromValue(v string) (*CertificateIssuer, error) {
	ev := CertificateIssuer(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CertificateIssuer: valid values are %v", v, allowedCertificateIssuerEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CertificateIssuer) IsValid() bool {
	for _, existing := range allowedCertificateIssuerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CertificateIssuer value.
func (v CertificateIssuer) Ptr() *CertificateIssuer {
	return &v
}
