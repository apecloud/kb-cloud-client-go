// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SshConfigCheckResult struct {
	Ip        string  `json:"ip"`
	Reachable bool    `json:"reachable"`
	Message   *string `json:"message,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSshConfigCheckResult instantiates a new SshConfigCheckResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSshConfigCheckResult(ip string, reachable bool) *SshConfigCheckResult {
	this := SshConfigCheckResult{}
	this.Ip = ip
	this.Reachable = reachable
	return &this
}

// NewSshConfigCheckResultWithDefaults instantiates a new SshConfigCheckResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSshConfigCheckResultWithDefaults() *SshConfigCheckResult {
	this := SshConfigCheckResult{}
	return &this
}

// GetIp returns the Ip field value.
func (o *SshConfigCheckResult) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *SshConfigCheckResult) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value.
func (o *SshConfigCheckResult) SetIp(v string) {
	o.Ip = v
}

// GetReachable returns the Reachable field value.
func (o *SshConfigCheckResult) GetReachable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Reachable
}

// GetReachableOk returns a tuple with the Reachable field value
// and a boolean to check if the value has been set.
func (o *SshConfigCheckResult) GetReachableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reachable, true
}

// SetReachable sets field value.
func (o *SshConfigCheckResult) SetReachable(v bool) {
	o.Reachable = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SshConfigCheckResult) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshConfigCheckResult) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SshConfigCheckResult) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SshConfigCheckResult) SetMessage(v string) {
	o.Message = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SshConfigCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["ip"] = o.Ip
	toSerialize["reachable"] = o.Reachable
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SshConfigCheckResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ip        *string `json:"ip"`
		Reachable *bool   `json:"reachable"`
		Message   *string `json:"message,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Ip == nil {
		return fmt.Errorf("required field ip missing")
	}
	if all.Reachable == nil {
		return fmt.Errorf("required field reachable missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ip", "reachable", "message"})
	} else {
		return err
	}
	o.Ip = *all.Ip
	o.Reachable = *all.Reachable
	o.Message = all.Message

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
