// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

type NodeScaleSpec struct {
	// IP address of the node to be added
	Ip               string `json:"ip"`
	AvailabilityZone string `json:"availabilityZone"`
	// Whether mark the node serves as control plane
	MarkControlPlane bool `json:"markControlPlane"`
	// Whether mark the node serves as data plane
	MarkDataPlane bool       `json:"markDataPlane"`
	Ssh           *SshConfig `json:"ssh,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNodeScaleSpec instantiates a new NodeScaleSpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNodeScaleSpec(ip string, availabilityZone string, markControlPlane bool, markDataPlane bool) *NodeScaleSpec {
	this := NodeScaleSpec{}
	this.Ip = ip
	this.AvailabilityZone = availabilityZone
	this.MarkControlPlane = markControlPlane
	this.MarkDataPlane = markDataPlane
	return &this
}

// NewNodeScaleSpecWithDefaults instantiates a new NodeScaleSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodeScaleSpecWithDefaults() *NodeScaleSpec {
	this := NodeScaleSpec{}
	var markControlPlane bool = false
	this.MarkControlPlane = markControlPlane
	var markDataPlane bool = true
	this.MarkDataPlane = markDataPlane
	return &this
}

// GetIp returns the Ip field value.
func (o *NodeScaleSpec) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *NodeScaleSpec) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value.
func (o *NodeScaleSpec) SetIp(v string) {
	o.Ip = v
}

// GetAvailabilityZone returns the AvailabilityZone field value.
func (o *NodeScaleSpec) GetAvailabilityZone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value
// and a boolean to check if the value has been set.
func (o *NodeScaleSpec) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailabilityZone, true
}

// SetAvailabilityZone sets field value.
func (o *NodeScaleSpec) SetAvailabilityZone(v string) {
	o.AvailabilityZone = v
}

// GetMarkControlPlane returns the MarkControlPlane field value.
func (o *NodeScaleSpec) GetMarkControlPlane() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.MarkControlPlane
}

// GetMarkControlPlaneOk returns a tuple with the MarkControlPlane field value
// and a boolean to check if the value has been set.
func (o *NodeScaleSpec) GetMarkControlPlaneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarkControlPlane, true
}

// SetMarkControlPlane sets field value.
func (o *NodeScaleSpec) SetMarkControlPlane(v bool) {
	o.MarkControlPlane = v
}

// GetMarkDataPlane returns the MarkDataPlane field value.
func (o *NodeScaleSpec) GetMarkDataPlane() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.MarkDataPlane
}

// GetMarkDataPlaneOk returns a tuple with the MarkDataPlane field value
// and a boolean to check if the value has been set.
func (o *NodeScaleSpec) GetMarkDataPlaneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarkDataPlane, true
}

// SetMarkDataPlane sets field value.
func (o *NodeScaleSpec) SetMarkDataPlane(v bool) {
	o.MarkDataPlane = v
}

// GetSsh returns the Ssh field value if set, zero value otherwise.
func (o *NodeScaleSpec) GetSsh() SshConfig {
	if o == nil || o.Ssh == nil {
		var ret SshConfig
		return ret
	}
	return *o.Ssh
}

// GetSshOk returns a tuple with the Ssh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeScaleSpec) GetSshOk() (*SshConfig, bool) {
	if o == nil || o.Ssh == nil {
		return nil, false
	}
	return o.Ssh, true
}

// HasSsh returns a boolean if a field has been set.
func (o *NodeScaleSpec) HasSsh() bool {
	return o != nil && o.Ssh != nil
}

// SetSsh gets a reference to the given SshConfig and assigns it to the Ssh field.
func (o *NodeScaleSpec) SetSsh(v SshConfig) {
	o.Ssh = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NodeScaleSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["ip"] = o.Ip
	toSerialize["availabilityZone"] = o.AvailabilityZone
	toSerialize["markControlPlane"] = o.MarkControlPlane
	toSerialize["markDataPlane"] = o.MarkDataPlane
	if o.Ssh != nil {
		toSerialize["ssh"] = o.Ssh
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NodeScaleSpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ip               *string    `json:"ip"`
		AvailabilityZone *string    `json:"availabilityZone"`
		MarkControlPlane *bool      `json:"markControlPlane"`
		MarkDataPlane    *bool      `json:"markDataPlane"`
		Ssh              *SshConfig `json:"ssh,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Ip == nil {
		return fmt.Errorf("required field ip missing")
	}
	if all.AvailabilityZone == nil {
		return fmt.Errorf("required field availabilityZone missing")
	}
	if all.MarkControlPlane == nil {
		return fmt.Errorf("required field markControlPlane missing")
	}
	if all.MarkDataPlane == nil {
		return fmt.Errorf("required field markDataPlane missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ip", "availabilityZone", "markControlPlane", "markDataPlane", "ssh"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Ip = *all.Ip
	o.AvailabilityZone = *all.AvailabilityZone
	o.MarkControlPlane = *all.MarkControlPlane
	o.MarkDataPlane = *all.MarkDataPlane
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
