// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"
)

type URLCheck struct {
	Url           string  `json:"url"`
	IsConnectable *bool   `json:"isConnectable,omitempty"`
	Timeout       *int32  `json:"timeout,omitempty"`
	Error         *string `json:"error,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewURLCheck instantiates a new URLCheck object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewURLCheck(url string) *URLCheck {
	this := URLCheck{}
	this.Url = url
	var isConnectable bool = false
	this.IsConnectable = &isConnectable
	var timeout int32 = 5
	this.Timeout = &timeout
	return &this
}

// NewURLCheckWithDefaults instantiates a new URLCheck object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewURLCheckWithDefaults() *URLCheck {
	this := URLCheck{}
	var isConnectable bool = false
	this.IsConnectable = &isConnectable
	var timeout int32 = 5
	this.Timeout = &timeout
	return &this
}

// GetUrl returns the Url field value.
func (o *URLCheck) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *URLCheck) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *URLCheck) SetUrl(v string) {
	o.Url = v
}

// GetIsConnectable returns the IsConnectable field value if set, zero value otherwise.
func (o *URLCheck) GetIsConnectable() bool {
	if o == nil || o.IsConnectable == nil {
		var ret bool
		return ret
	}
	return *o.IsConnectable
}

// GetIsConnectableOk returns a tuple with the IsConnectable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLCheck) GetIsConnectableOk() (*bool, bool) {
	if o == nil || o.IsConnectable == nil {
		return nil, false
	}
	return o.IsConnectable, true
}

// HasIsConnectable returns a boolean if a field has been set.
func (o *URLCheck) HasIsConnectable() bool {
	return o != nil && o.IsConnectable != nil
}

// SetIsConnectable gets a reference to the given bool and assigns it to the IsConnectable field.
func (o *URLCheck) SetIsConnectable(v bool) {
	o.IsConnectable = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *URLCheck) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLCheck) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *URLCheck) HasTimeout() bool {
	return o != nil && o.Timeout != nil
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *URLCheck) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *URLCheck) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *URLCheck) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *URLCheck) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *URLCheck) SetError(v string) {
	o.Error = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o URLCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["url"] = o.Url
	if o.IsConnectable != nil {
		toSerialize["isConnectable"] = o.IsConnectable
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *URLCheck) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Url           *string `json:"url"`
		IsConnectable *bool   `json:"isConnectable,omitempty"`
		Timeout       *int32  `json:"timeout,omitempty"`
		Error         *string `json:"error,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"url", "isConnectable", "timeout", "error"})
	} else {
		return err
	}
	o.Url = *all.Url
	o.IsConnectable = all.IsConnectable
	o.Timeout = all.Timeout
	o.Error = all.Error

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
