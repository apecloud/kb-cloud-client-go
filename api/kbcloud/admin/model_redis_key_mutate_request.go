// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisKeyMutateRequest struct {
	Key          string  `json:"key"`
	ExpectedType *string `json:"expectedType,omitempty"`
	// Redis key mutation operation, such as set, hset, hdel, lset, lpush, rpush, sadd, srem, zadd, zrem, xadd, xtrim, json_set, json_del, expire, persist, rename, or delete
	Operation string                 `json:"operation"`
	Payload   map[string]interface{} `json:"payload,omitempty"`
	Confirm   *bool                  `json:"confirm,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeyMutateRequest instantiates a new RedisKeyMutateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeyMutateRequest(key string, operation string) *RedisKeyMutateRequest {
	this := RedisKeyMutateRequest{}
	this.Key = key
	this.Operation = operation
	return &this
}

// NewRedisKeyMutateRequestWithDefaults instantiates a new RedisKeyMutateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeyMutateRequestWithDefaults() *RedisKeyMutateRequest {
	this := RedisKeyMutateRequest{}
	return &this
}

// GetKey returns the Key field value.
func (o *RedisKeyMutateRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *RedisKeyMutateRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *RedisKeyMutateRequest) SetKey(v string) {
	o.Key = v
}

// GetExpectedType returns the ExpectedType field value if set, zero value otherwise.
func (o *RedisKeyMutateRequest) GetExpectedType() string {
	if o == nil || o.ExpectedType == nil {
		var ret string
		return ret
	}
	return *o.ExpectedType
}

// GetExpectedTypeOk returns a tuple with the ExpectedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyMutateRequest) GetExpectedTypeOk() (*string, bool) {
	if o == nil || o.ExpectedType == nil {
		return nil, false
	}
	return o.ExpectedType, true
}

// HasExpectedType returns a boolean if a field has been set.
func (o *RedisKeyMutateRequest) HasExpectedType() bool {
	return o != nil && o.ExpectedType != nil
}

// SetExpectedType gets a reference to the given string and assigns it to the ExpectedType field.
func (o *RedisKeyMutateRequest) SetExpectedType(v string) {
	o.ExpectedType = &v
}

// GetOperation returns the Operation field value.
func (o *RedisKeyMutateRequest) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *RedisKeyMutateRequest) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value.
func (o *RedisKeyMutateRequest) SetOperation(v string) {
	o.Operation = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *RedisKeyMutateRequest) GetPayload() map[string]interface{} {
	if o == nil || o.Payload == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyMutateRequest) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *RedisKeyMutateRequest) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *RedisKeyMutateRequest) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetConfirm returns the Confirm field value if set, zero value otherwise.
func (o *RedisKeyMutateRequest) GetConfirm() bool {
	if o == nil || o.Confirm == nil {
		var ret bool
		return ret
	}
	return *o.Confirm
}

// GetConfirmOk returns a tuple with the Confirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyMutateRequest) GetConfirmOk() (*bool, bool) {
	if o == nil || o.Confirm == nil {
		return nil, false
	}
	return o.Confirm, true
}

// HasConfirm returns a boolean if a field has been set.
func (o *RedisKeyMutateRequest) HasConfirm() bool {
	return o != nil && o.Confirm != nil
}

// SetConfirm gets a reference to the given bool and assigns it to the Confirm field.
func (o *RedisKeyMutateRequest) SetConfirm(v bool) {
	o.Confirm = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeyMutateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	if o.ExpectedType != nil {
		toSerialize["expectedType"] = o.ExpectedType
	}
	toSerialize["operation"] = o.Operation
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if o.Confirm != nil {
		toSerialize["confirm"] = o.Confirm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisKeyMutateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key          *string                `json:"key"`
		ExpectedType *string                `json:"expectedType,omitempty"`
		Operation    *string                `json:"operation"`
		Payload      map[string]interface{} `json:"payload,omitempty"`
		Confirm      *bool                  `json:"confirm,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Operation == nil {
		return fmt.Errorf("required field operation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"key", "expectedType", "operation", "payload", "confirm"})
	} else {
		return err
	}
	o.Key = *all.Key
	o.ExpectedType = all.ExpectedType
	o.Operation = *all.Operation
	o.Payload = all.Payload
	o.Confirm = all.Confirm

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
