// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// IpPoolProvider KBE provider used for Pod IP pool discovery and explicit selection.
type IpPoolProvider string

// List of IpPoolProvider.
const (
	IpPoolProviderSpiderpool IpPoolProvider = "spiderpool"
)

var allowedIpPoolProviderEnumValues = []IpPoolProvider{
	IpPoolProviderSpiderpool,
}

// GetAllowedValues returns the list of possible values.
func (v *IpPoolProvider) GetAllowedValues() []IpPoolProvider {
	return allowedIpPoolProviderEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *IpPoolProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = IpPoolProvider(value)
	return nil
}

// NewIpPoolProviderFromValue returns a pointer to a valid IpPoolProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewIpPoolProviderFromValue(v string) (*IpPoolProvider, error) {
	ev := IpPoolProvider(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for IpPoolProvider: valid values are %v", v, allowedIpPoolProviderEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v IpPoolProvider) IsValid() bool {
	for _, existing := range allowedIpPoolProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IpPoolProvider value.
func (v IpPoolProvider) Ptr() *IpPoolProvider {
	return &v
}
