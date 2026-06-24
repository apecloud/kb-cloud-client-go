// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type RedisWorkbenchExecuteRequest struct {
	Command          string `json:"command"`
	ConfirmDangerous *bool  `json:"confirmDangerous,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisWorkbenchExecuteRequest instantiates a new RedisWorkbenchExecuteRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisWorkbenchExecuteRequest(command string) *RedisWorkbenchExecuteRequest {
	this := RedisWorkbenchExecuteRequest{}
	this.Command = command
	return &this
}

// NewRedisWorkbenchExecuteRequestWithDefaults instantiates a new RedisWorkbenchExecuteRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisWorkbenchExecuteRequestWithDefaults() *RedisWorkbenchExecuteRequest {
	this := RedisWorkbenchExecuteRequest{}
	return &this
}

// GetCommand returns the Command field value.
func (o *RedisWorkbenchExecuteRequest) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *RedisWorkbenchExecuteRequest) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value.
func (o *RedisWorkbenchExecuteRequest) SetCommand(v string) {
	o.Command = v
}

// GetConfirmDangerous returns the ConfirmDangerous field value if set, zero value otherwise.
func (o *RedisWorkbenchExecuteRequest) GetConfirmDangerous() bool {
	if o == nil || o.ConfirmDangerous == nil {
		var ret bool
		return ret
	}
	return *o.ConfirmDangerous
}

// GetConfirmDangerousOk returns a tuple with the ConfirmDangerous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisWorkbenchExecuteRequest) GetConfirmDangerousOk() (*bool, bool) {
	if o == nil || o.ConfirmDangerous == nil {
		return nil, false
	}
	return o.ConfirmDangerous, true
}

// HasConfirmDangerous returns a boolean if a field has been set.
func (o *RedisWorkbenchExecuteRequest) HasConfirmDangerous() bool {
	return o != nil && o.ConfirmDangerous != nil
}

// SetConfirmDangerous gets a reference to the given bool and assigns it to the ConfirmDangerous field.
func (o *RedisWorkbenchExecuteRequest) SetConfirmDangerous(v bool) {
	o.ConfirmDangerous = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisWorkbenchExecuteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["command"] = o.Command
	if o.ConfirmDangerous != nil {
		toSerialize["confirmDangerous"] = o.ConfirmDangerous
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisWorkbenchExecuteRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Command          *string `json:"command"`
		ConfirmDangerous *bool   `json:"confirmDangerous,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Command == nil {
		return fmt.Errorf("required field command missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"command", "confirmDangerous"})
	} else {
		return err
	}
	o.Command = *all.Command
	o.ConfirmDangerous = all.ConfirmDangerous

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
