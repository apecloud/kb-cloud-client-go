// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisInsightEnableOpt struct {
	NetworkMode *RedisInsightNetMode `json:"networkMode,omitempty"`
	ServiceType *RedisInsightSvcType `json:"serviceType,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisInsightEnableOpt instantiates a new RedisInsightEnableOpt object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisInsightEnableOpt() *RedisInsightEnableOpt {
	this := RedisInsightEnableOpt{}
	return &this
}

// NewRedisInsightEnableOptWithDefaults instantiates a new RedisInsightEnableOpt object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisInsightEnableOptWithDefaults() *RedisInsightEnableOpt {
	this := RedisInsightEnableOpt{}
	return &this
}

// GetNetworkMode returns the NetworkMode field value if set, zero value otherwise.
func (o *RedisInsightEnableOpt) GetNetworkMode() RedisInsightNetMode {
	if o == nil || o.NetworkMode == nil {
		var ret RedisInsightNetMode
		return ret
	}
	return *o.NetworkMode
}

// GetNetworkModeOk returns a tuple with the NetworkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisInsightEnableOpt) GetNetworkModeOk() (*RedisInsightNetMode, bool) {
	if o == nil || o.NetworkMode == nil {
		return nil, false
	}
	return o.NetworkMode, true
}

// HasNetworkMode returns a boolean if a field has been set.
func (o *RedisInsightEnableOpt) HasNetworkMode() bool {
	return o != nil && o.NetworkMode != nil
}

// SetNetworkMode gets a reference to the given RedisInsightNetMode and assigns it to the NetworkMode field.
func (o *RedisInsightEnableOpt) SetNetworkMode(v RedisInsightNetMode) {
	o.NetworkMode = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *RedisInsightEnableOpt) GetServiceType() RedisInsightSvcType {
	if o == nil || o.ServiceType == nil {
		var ret RedisInsightSvcType
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisInsightEnableOpt) GetServiceTypeOk() (*RedisInsightSvcType, bool) {
	if o == nil || o.ServiceType == nil {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *RedisInsightEnableOpt) HasServiceType() bool {
	return o != nil && o.ServiceType != nil
}

// SetServiceType gets a reference to the given RedisInsightSvcType and assigns it to the ServiceType field.
func (o *RedisInsightEnableOpt) SetServiceType(v RedisInsightSvcType) {
	o.ServiceType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisInsightEnableOpt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NetworkMode != nil {
		toSerialize["networkMode"] = o.NetworkMode
	}
	if o.ServiceType != nil {
		toSerialize["serviceType"] = o.ServiceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisInsightEnableOpt) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NetworkMode *RedisInsightNetMode `json:"networkMode,omitempty"`
		ServiceType *RedisInsightSvcType `json:"serviceType,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"networkMode", "serviceType"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.NetworkMode != nil && !all.NetworkMode.IsValid() {
		hasInvalidField = true
	} else {
		o.NetworkMode = all.NetworkMode
	}
	if all.ServiceType != nil && !all.ServiceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ServiceType = all.ServiceType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
