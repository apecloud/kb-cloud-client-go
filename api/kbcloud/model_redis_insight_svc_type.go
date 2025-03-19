// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisInsightSvcType string

// List of RedisInsightSvcType.
const (
	RedisInsightSvcTypeClusterIp    RedisInsightSvcType = "ClusterIP"
	RedisInsightSvcTypeLoadBalancer RedisInsightSvcType = "LoadBalancer"
	RedisInsightSvcTypeNodePort     RedisInsightSvcType = "NodePort"
)

var allowedRedisInsightSvcTypeEnumValues = []RedisInsightSvcType{
	RedisInsightSvcTypeClusterIp,
	RedisInsightSvcTypeLoadBalancer,
	RedisInsightSvcTypeNodePort,
}

// GetAllowedValues returns the list of possible values.
func (v *RedisInsightSvcType) GetAllowedValues() []RedisInsightSvcType {
	return allowedRedisInsightSvcTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *RedisInsightSvcType) UnmarshalJSON(src []byte) error {
	var value string
	err := common.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = RedisInsightSvcType(value)
	return nil
}

// NewRedisInsightSvcTypeFromValue returns a pointer to a valid RedisInsightSvcType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewRedisInsightSvcTypeFromValue(v string) (*RedisInsightSvcType, error) {
	ev := RedisInsightSvcType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RedisInsightSvcType: valid values are %v", v, allowedRedisInsightSvcTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v RedisInsightSvcType) IsValid() bool {
	for _, existing := range allowedRedisInsightSvcTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RedisInsightSvcType value.
func (v RedisInsightSvcType) Ptr() *RedisInsightSvcType {
	return &v
}
