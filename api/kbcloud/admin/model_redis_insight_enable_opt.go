// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisInsightEnableOpt struct {
	NetworkMode RedisInsightNetMode `json:"networkMode"`
	ServiceType RedisInsightSvcType `json:"serviceType"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisInsightEnableOpt instantiates a new RedisInsightEnableOpt object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisInsightEnableOpt(networkMode RedisInsightNetMode, serviceType RedisInsightSvcType) *RedisInsightEnableOpt {
	this := RedisInsightEnableOpt{}
	this.NetworkMode = networkMode
	this.ServiceType = serviceType
	return &this
}

// NewRedisInsightEnableOptWithDefaults instantiates a new RedisInsightEnableOpt object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisInsightEnableOptWithDefaults() *RedisInsightEnableOpt {
	this := RedisInsightEnableOpt{}
	return &this
}

// GetNetworkMode returns the NetworkMode field value.
func (o *RedisInsightEnableOpt) GetNetworkMode() RedisInsightNetMode {
	if o == nil {
		var ret RedisInsightNetMode
		return ret
	}
	return o.NetworkMode
}

// GetNetworkModeOk returns a tuple with the NetworkMode field value
// and a boolean to check if the value has been set.
func (o *RedisInsightEnableOpt) GetNetworkModeOk() (*RedisInsightNetMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkMode, true
}

// SetNetworkMode sets field value.
func (o *RedisInsightEnableOpt) SetNetworkMode(v RedisInsightNetMode) {
	o.NetworkMode = v
}

// GetServiceType returns the ServiceType field value.
func (o *RedisInsightEnableOpt) GetServiceType() RedisInsightSvcType {
	if o == nil {
		var ret RedisInsightSvcType
		return ret
	}
	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *RedisInsightEnableOpt) GetServiceTypeOk() (*RedisInsightSvcType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value.
func (o *RedisInsightEnableOpt) SetServiceType(v RedisInsightSvcType) {
	o.ServiceType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisInsightEnableOpt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["networkMode"] = o.NetworkMode
	toSerialize["serviceType"] = o.ServiceType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisInsightEnableOpt) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NetworkMode *RedisInsightNetMode `json:"networkMode"`
		ServiceType *RedisInsightSvcType `json:"serviceType"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.NetworkMode == nil {
		return fmt.Errorf("required field networkMode missing")
	}
	if all.ServiceType == nil {
		return fmt.Errorf("required field serviceType missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"networkMode", "serviceType"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.NetworkMode.IsValid() {
		hasInvalidField = true
	} else {
		o.NetworkMode = *all.NetworkMode
	}
	if !all.ServiceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ServiceType = *all.ServiceType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
