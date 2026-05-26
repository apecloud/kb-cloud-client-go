// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type MongoShellEvaluateRequest struct {
	// mongosh input to evaluate
	Code string `json:"code"`
	// optional execution timeout in milliseconds
	TimeoutMs *int64 `json:"timeoutMs,omitempty"`
	// optional output cap in bytes
	MaxOutputBytes *int64 `json:"maxOutputBytes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoShellEvaluateRequest instantiates a new MongoShellEvaluateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoShellEvaluateRequest(code string) *MongoShellEvaluateRequest {
	this := MongoShellEvaluateRequest{}
	this.Code = code
	return &this
}

// NewMongoShellEvaluateRequestWithDefaults instantiates a new MongoShellEvaluateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoShellEvaluateRequestWithDefaults() *MongoShellEvaluateRequest {
	this := MongoShellEvaluateRequest{}
	return &this
}

// GetCode returns the Code field value.
func (o *MongoShellEvaluateRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *MongoShellEvaluateRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *MongoShellEvaluateRequest) SetCode(v string) {
	o.Code = v
}

// GetTimeoutMs returns the TimeoutMs field value if set, zero value otherwise.
func (o *MongoShellEvaluateRequest) GetTimeoutMs() int64 {
	if o == nil || o.TimeoutMs == nil {
		var ret int64
		return ret
	}
	return *o.TimeoutMs
}

// GetTimeoutMsOk returns a tuple with the TimeoutMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellEvaluateRequest) GetTimeoutMsOk() (*int64, bool) {
	if o == nil || o.TimeoutMs == nil {
		return nil, false
	}
	return o.TimeoutMs, true
}

// HasTimeoutMs returns a boolean if a field has been set.
func (o *MongoShellEvaluateRequest) HasTimeoutMs() bool {
	return o != nil && o.TimeoutMs != nil
}

// SetTimeoutMs gets a reference to the given int64 and assigns it to the TimeoutMs field.
func (o *MongoShellEvaluateRequest) SetTimeoutMs(v int64) {
	o.TimeoutMs = &v
}

// GetMaxOutputBytes returns the MaxOutputBytes field value if set, zero value otherwise.
func (o *MongoShellEvaluateRequest) GetMaxOutputBytes() int64 {
	if o == nil || o.MaxOutputBytes == nil {
		var ret int64
		return ret
	}
	return *o.MaxOutputBytes
}

// GetMaxOutputBytesOk returns a tuple with the MaxOutputBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellEvaluateRequest) GetMaxOutputBytesOk() (*int64, bool) {
	if o == nil || o.MaxOutputBytes == nil {
		return nil, false
	}
	return o.MaxOutputBytes, true
}

// HasMaxOutputBytes returns a boolean if a field has been set.
func (o *MongoShellEvaluateRequest) HasMaxOutputBytes() bool {
	return o != nil && o.MaxOutputBytes != nil
}

// SetMaxOutputBytes gets a reference to the given int64 and assigns it to the MaxOutputBytes field.
func (o *MongoShellEvaluateRequest) SetMaxOutputBytes(v int64) {
	o.MaxOutputBytes = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoShellEvaluateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	if o.TimeoutMs != nil {
		toSerialize["timeoutMs"] = o.TimeoutMs
	}
	if o.MaxOutputBytes != nil {
		toSerialize["maxOutputBytes"] = o.MaxOutputBytes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoShellEvaluateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code           *string `json:"code"`
		TimeoutMs      *int64  `json:"timeoutMs,omitempty"`
		MaxOutputBytes *int64  `json:"maxOutputBytes,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"code", "timeoutMs", "maxOutputBytes"})
	} else {
		return err
	}
	o.Code = *all.Code
	o.TimeoutMs = all.TimeoutMs
	o.MaxOutputBytes = all.MaxOutputBytes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
