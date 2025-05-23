// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type SshNodeSpec struct {
	// IP address of the node
	Ip  string     `json:"ip"`
	Ssh *SshConfig `json:"ssh,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSshNodeSpec instantiates a new SshNodeSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSshNodeSpec(ip string) *SshNodeSpec {
	this := SshNodeSpec{}
	this.Ip = ip
	return &this
}

// NewSshNodeSpecWithDefaults instantiates a new SshNodeSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSshNodeSpecWithDefaults() *SshNodeSpec {
	this := SshNodeSpec{}
	return &this
}

// GetIp returns the Ip field value.
func (o *SshNodeSpec) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *SshNodeSpec) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value.
func (o *SshNodeSpec) SetIp(v string) {
	o.Ip = v
}

// GetSsh returns the Ssh field value if set, zero value otherwise.
func (o *SshNodeSpec) GetSsh() SshConfig {
	if o == nil || o.Ssh == nil {
		var ret SshConfig
		return ret
	}
	return *o.Ssh
}

// GetSshOk returns a tuple with the Ssh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshNodeSpec) GetSshOk() (*SshConfig, bool) {
	if o == nil || o.Ssh == nil {
		return nil, false
	}
	return o.Ssh, true
}

// HasSsh returns a boolean if a field has been set.
func (o *SshNodeSpec) HasSsh() bool {
	return o != nil && o.Ssh != nil
}

// SetSsh gets a reference to the given SshConfig and assigns it to the Ssh field.
func (o *SshNodeSpec) SetSsh(v SshConfig) {
	o.Ssh = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SshNodeSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["ip"] = o.Ip
	if o.Ssh != nil {
		toSerialize["ssh"] = o.Ssh
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SshNodeSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ip  *string    `json:"ip"`
		Ssh *SshConfig `json:"ssh,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Ip == nil {
		return fmt.Errorf("required field ip missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ip", "ssh"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Ip = *all.Ip
	if all.Ssh != nil && all.Ssh.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Ssh = all.Ssh

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
