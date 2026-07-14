// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type TestConnectResponse struct {
	// whether the datasource connection test passed
	Ok *bool `json:"ok,omitempty"`
	// diagnostic message returned by DMS
	Message *string `json:"message,omitempty"`
	// additional diagnostics returned by DMS
	ExtraInfo map[string]interface{} `json:"extraInfo,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTestConnectResponse instantiates a new TestConnectResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTestConnectResponse() *TestConnectResponse {
	this := TestConnectResponse{}
	return &this
}

// NewTestConnectResponseWithDefaults instantiates a new TestConnectResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTestConnectResponseWithDefaults() *TestConnectResponse {
	this := TestConnectResponse{}
	return &this
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *TestConnectResponse) GetOk() bool {
	if o == nil || o.Ok == nil {
		var ret bool
		return ret
	}
	return *o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestConnectResponse) GetOkOk() (*bool, bool) {
	if o == nil || o.Ok == nil {
		return nil, false
	}
	return o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *TestConnectResponse) HasOk() bool {
	return o != nil && o.Ok != nil
}

// SetOk gets a reference to the given bool and assigns it to the Ok field.
func (o *TestConnectResponse) SetOk(v bool) {
	o.Ok = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TestConnectResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestConnectResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TestConnectResponse) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TestConnectResponse) SetMessage(v string) {
	o.Message = &v
}

// GetExtraInfo returns the ExtraInfo field value if set, zero value otherwise.
func (o *TestConnectResponse) GetExtraInfo() map[string]interface{} {
	if o == nil || o.ExtraInfo == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraInfo
}

// GetExtraInfoOk returns a tuple with the ExtraInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestConnectResponse) GetExtraInfoOk() (*map[string]interface{}, bool) {
	if o == nil || o.ExtraInfo == nil {
		return nil, false
	}
	return &o.ExtraInfo, true
}

// HasExtraInfo returns a boolean if a field has been set.
func (o *TestConnectResponse) HasExtraInfo() bool {
	return o != nil && o.ExtraInfo != nil
}

// SetExtraInfo gets a reference to the given map[string]interface{} and assigns it to the ExtraInfo field.
func (o *TestConnectResponse) SetExtraInfo(v map[string]interface{}) {
	o.ExtraInfo = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TestConnectResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Ok != nil {
		toSerialize["ok"] = o.Ok
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ExtraInfo != nil {
		toSerialize["extraInfo"] = o.ExtraInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TestConnectResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ok        *bool                  `json:"ok,omitempty"`
		Message   *string                `json:"message,omitempty"`
		ExtraInfo map[string]interface{} `json:"extraInfo,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ok", "message", "extraInfo"})
	} else {
		return err
	}
	o.Ok = all.Ok
	o.Message = all.Message
	o.ExtraInfo = all.ExtraInfo

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
