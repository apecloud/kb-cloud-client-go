// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type NodeScaleRequest struct {
	// List of masters (jumpserver for ssh), at least one node is needed
	Masters []SshNodeSpec `json:"masters,omitempty"`
	// List of nodes to be added, support 1-20 nodes per request
	Nodes      []NodeScaleSpec `json:"nodes"`
	DefaultSsh *SshConfig      `json:"defaultSSH,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNodeScaleRequest instantiates a new NodeScaleRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNodeScaleRequest(nodes []NodeScaleSpec) *NodeScaleRequest {
	this := NodeScaleRequest{}
	this.Nodes = nodes
	return &this
}

// NewNodeScaleRequestWithDefaults instantiates a new NodeScaleRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNodeScaleRequestWithDefaults() *NodeScaleRequest {
	this := NodeScaleRequest{}
	return &this
}

// GetMasters returns the Masters field value if set, zero value otherwise.
func (o *NodeScaleRequest) GetMasters() []SshNodeSpec {
	if o == nil || o.Masters == nil {
		var ret []SshNodeSpec
		return ret
	}
	return o.Masters
}

// GetMastersOk returns a tuple with the Masters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeScaleRequest) GetMastersOk() (*[]SshNodeSpec, bool) {
	if o == nil || o.Masters == nil {
		return nil, false
	}
	return &o.Masters, true
}

// HasMasters returns a boolean if a field has been set.
func (o *NodeScaleRequest) HasMasters() bool {
	return o != nil && o.Masters != nil
}

// SetMasters gets a reference to the given []SshNodeSpec and assigns it to the Masters field.
func (o *NodeScaleRequest) SetMasters(v []SshNodeSpec) {
	o.Masters = v
}

// GetNodes returns the Nodes field value.
func (o *NodeScaleRequest) GetNodes() []NodeScaleSpec {
	if o == nil {
		var ret []NodeScaleSpec
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *NodeScaleRequest) GetNodesOk() (*[]NodeScaleSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value.
func (o *NodeScaleRequest) SetNodes(v []NodeScaleSpec) {
	o.Nodes = v
}

// GetDefaultSsh returns the DefaultSsh field value if set, zero value otherwise.
func (o *NodeScaleRequest) GetDefaultSsh() SshConfig {
	if o == nil || o.DefaultSsh == nil {
		var ret SshConfig
		return ret
	}
	return *o.DefaultSsh
}

// GetDefaultSshOk returns a tuple with the DefaultSsh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeScaleRequest) GetDefaultSshOk() (*SshConfig, bool) {
	if o == nil || o.DefaultSsh == nil {
		return nil, false
	}
	return o.DefaultSsh, true
}

// HasDefaultSsh returns a boolean if a field has been set.
func (o *NodeScaleRequest) HasDefaultSsh() bool {
	return o != nil && o.DefaultSsh != nil
}

// SetDefaultSsh gets a reference to the given SshConfig and assigns it to the DefaultSsh field.
func (o *NodeScaleRequest) SetDefaultSsh(v SshConfig) {
	o.DefaultSsh = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NodeScaleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Masters != nil {
		toSerialize["masters"] = o.Masters
	}
	toSerialize["nodes"] = o.Nodes
	if o.DefaultSsh != nil {
		toSerialize["defaultSSH"] = o.DefaultSsh
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NodeScaleRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Masters    []SshNodeSpec    `json:"masters,omitempty"`
		Nodes      *[]NodeScaleSpec `json:"nodes"`
		DefaultSsh *SshConfig       `json:"defaultSSH,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Nodes == nil {
		return fmt.Errorf("required field nodes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"masters", "nodes", "defaultSSH"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Masters = all.Masters
	o.Nodes = *all.Nodes
	if all.DefaultSsh != nil && all.DefaultSsh.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DefaultSsh = all.DefaultSsh

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
