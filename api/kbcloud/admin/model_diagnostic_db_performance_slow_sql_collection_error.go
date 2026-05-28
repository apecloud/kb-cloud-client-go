// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceSlowSQLCollectionError Structured collection error detail when slowSqlState is source_present_collection_failed.
type DiagnosticDBPerformanceSlowSQLCollectionError struct {
	// Stable error code, such as connect_refused, permission_denied, query_timeout.
	Code string `json:"code"`
	// Human-readable error message passed through to Console.
	Message string `json:"message"`
	// Whether the controller indicates the collection error is retryable.
	Retryable bool `json:"retryable"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceSlowSQLCollectionError instantiates a new DiagnosticDBPerformanceSlowSQLCollectionError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceSlowSQLCollectionError(code string, message string, retryable bool) *DiagnosticDBPerformanceSlowSQLCollectionError {
	this := DiagnosticDBPerformanceSlowSQLCollectionError{}
	this.Code = code
	this.Message = message
	this.Retryable = retryable
	return &this
}

// NewDiagnosticDBPerformanceSlowSQLCollectionErrorWithDefaults instantiates a new DiagnosticDBPerformanceSlowSQLCollectionError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceSlowSQLCollectionErrorWithDefaults() *DiagnosticDBPerformanceSlowSQLCollectionError {
	this := DiagnosticDBPerformanceSlowSQLCollectionError{}
	return &this
}

// GetCode returns the Code field value.
func (o *DiagnosticDBPerformanceSlowSQLCollectionError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLCollectionError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *DiagnosticDBPerformanceSlowSQLCollectionError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value.
func (o *DiagnosticDBPerformanceSlowSQLCollectionError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLCollectionError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *DiagnosticDBPerformanceSlowSQLCollectionError) SetMessage(v string) {
	o.Message = v
}

// GetRetryable returns the Retryable field value.
func (o *DiagnosticDBPerformanceSlowSQLCollectionError) GetRetryable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceSlowSQLCollectionError) GetRetryableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retryable, true
}

// SetRetryable sets field value.
func (o *DiagnosticDBPerformanceSlowSQLCollectionError) SetRetryable(v bool) {
	o.Retryable = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceSlowSQLCollectionError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["retryable"] = o.Retryable

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceSlowSQLCollectionError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code      *string `json:"code"`
		Message   *string `json:"message"`
		Retryable *bool   `json:"retryable"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Code == nil {
		return fmt.Errorf("required field code missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Retryable == nil {
		return fmt.Errorf("required field retryable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"code", "message", "retryable"})
	} else {
		return err
	}
	o.Code = *all.Code
	o.Message = *all.Message
	o.Retryable = *all.Retryable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
