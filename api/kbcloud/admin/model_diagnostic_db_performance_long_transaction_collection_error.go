// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// DiagnosticDBPerformanceLongTransactionCollectionError Structured collection error detail when longTransactionState is source_present_collection_failed.
type DiagnosticDBPerformanceLongTransactionCollectionError struct {
	// Stable error code, such as permission_denied, query_timeout, collection_failed, or unexpected_output.
	Code string `json:"code"`
	// Human-readable, redacted error message passed through to Console.
	Message string `json:"message"`
	// Whether the collection error is retryable.
	Retryable bool `json:"retryable"`
	// Name or redacted template of the readonly query that failed.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDBPerformanceLongTransactionCollectionError instantiates a new DiagnosticDBPerformanceLongTransactionCollectionError object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDBPerformanceLongTransactionCollectionError(code string, message string, retryable bool, query string) *DiagnosticDBPerformanceLongTransactionCollectionError {
	this := DiagnosticDBPerformanceLongTransactionCollectionError{}
	this.Code = code
	this.Message = message
	this.Retryable = retryable
	this.Query = query
	return &this
}

// NewDiagnosticDBPerformanceLongTransactionCollectionErrorWithDefaults instantiates a new DiagnosticDBPerformanceLongTransactionCollectionError object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDBPerformanceLongTransactionCollectionErrorWithDefaults() *DiagnosticDBPerformanceLongTransactionCollectionError {
	this := DiagnosticDBPerformanceLongTransactionCollectionError{}
	return &this
}

// GetCode returns the Code field value.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) SetMessage(v string) {
	o.Message = v
}

// GetRetryable returns the Retryable field value.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) GetRetryable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) GetRetryableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retryable, true
}

// SetRetryable sets field value.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) SetRetryable(v bool) {
	o.Retryable = v
}

// GetQuery returns the Query field value.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDBPerformanceLongTransactionCollectionError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["retryable"] = o.Retryable
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDBPerformanceLongTransactionCollectionError) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Code      *string `json:"code"`
		Message   *string `json:"message"`
		Retryable *bool   `json:"retryable"`
		Query     *string `json:"query"`
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
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"code", "message", "retryable", "query"})
	} else {
		return err
	}
	o.Code = *all.Code
	o.Message = *all.Message
	o.Retryable = *all.Retryable
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
