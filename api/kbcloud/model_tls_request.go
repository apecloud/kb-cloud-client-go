// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type TlsRequest struct {
	// Enable TLS or not
	Enable bool `json:"enable"`
	// TLS options
	Options *TlsOption `json:"options,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTlsRequest instantiates a new TlsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTlsRequest(enable bool) *TlsRequest {
	this := TlsRequest{}
	this.Enable = enable
	return &this
}

// NewTlsRequestWithDefaults instantiates a new TlsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTlsRequestWithDefaults() *TlsRequest {
	this := TlsRequest{}
	return &this
}

// GetEnable returns the Enable field value.
func (o *TlsRequest) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *TlsRequest) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value.
func (o *TlsRequest) SetEnable(v bool) {
	o.Enable = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *TlsRequest) GetOptions() TlsOption {
	if o == nil || o.Options == nil {
		var ret TlsOption
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsRequest) GetOptionsOk() (*TlsOption, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *TlsRequest) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given TlsOption and assigns it to the Options field.
func (o *TlsRequest) SetOptions(v TlsOption) {
	o.Options = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TlsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["enable"] = o.Enable
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TlsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enable  *bool      `json:"enable"`
		Options *TlsOption `json:"options,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Enable == nil {
		return fmt.Errorf("required field enable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"enable", "options"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enable = *all.Enable
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
