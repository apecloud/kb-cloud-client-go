// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package kbcloud

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)


 
type TlsRequest struct {
	// Enable TLS or not
	Enable *bool `json:"enable,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewTlsRequest instantiates a new TlsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTlsRequest() *TlsRequest {
	this := TlsRequest{}
	return &this
}

// NewTlsRequestWithDefaults instantiates a new TlsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTlsRequestWithDefaults() *TlsRequest {
	this := TlsRequest{}
	return &this
}
// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *TlsRequest) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsRequest) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *TlsRequest) HasEnable() bool {
	return o != nil && o.Enable != nil
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *TlsRequest) SetEnable(v bool) {
	o.Enable = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o TlsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TlsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enable *bool `json:"enable,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "enable",  })
	} else {
		return err
	}
	o.Enable = all.Enable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
