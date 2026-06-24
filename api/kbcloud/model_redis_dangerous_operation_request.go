// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisDangerousOperationRequest struct {
	ConfirmDangerous *bool `json:"confirmDangerous,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisDangerousOperationRequest instantiates a new RedisDangerousOperationRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisDangerousOperationRequest() *RedisDangerousOperationRequest {
	this := RedisDangerousOperationRequest{}
	return &this
}

// NewRedisDangerousOperationRequestWithDefaults instantiates a new RedisDangerousOperationRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisDangerousOperationRequestWithDefaults() *RedisDangerousOperationRequest {
	this := RedisDangerousOperationRequest{}
	return &this
}

// GetConfirmDangerous returns the ConfirmDangerous field value if set, zero value otherwise.
func (o *RedisDangerousOperationRequest) GetConfirmDangerous() bool {
	if o == nil || o.ConfirmDangerous == nil {
		var ret bool
		return ret
	}
	return *o.ConfirmDangerous
}

// GetConfirmDangerousOk returns a tuple with the ConfirmDangerous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisDangerousOperationRequest) GetConfirmDangerousOk() (*bool, bool) {
	if o == nil || o.ConfirmDangerous == nil {
		return nil, false
	}
	return o.ConfirmDangerous, true
}

// HasConfirmDangerous returns a boolean if a field has been set.
func (o *RedisDangerousOperationRequest) HasConfirmDangerous() bool {
	return o != nil && o.ConfirmDangerous != nil
}

// SetConfirmDangerous gets a reference to the given bool and assigns it to the ConfirmDangerous field.
func (o *RedisDangerousOperationRequest) SetConfirmDangerous(v bool) {
	o.ConfirmDangerous = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisDangerousOperationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ConfirmDangerous != nil {
		toSerialize["confirmDangerous"] = o.ConfirmDangerous
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisDangerousOperationRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfirmDangerous *bool `json:"confirmDangerous,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"confirmDangerous"})
	} else {
		return err
	}
	o.ConfirmDangerous = all.ConfirmDangerous

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
