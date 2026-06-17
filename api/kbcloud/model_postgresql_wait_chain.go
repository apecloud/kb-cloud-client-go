// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type PostgresqlWaitChain struct {
	Nodes []PostgresqlWaitChainNode `json:"nodes"`
	Edges []PostgresqlWaitChainEdge `json:"edges"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlWaitChain instantiates a new PostgresqlWaitChain object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlWaitChain(nodes []PostgresqlWaitChainNode, edges []PostgresqlWaitChainEdge) *PostgresqlWaitChain {
	this := PostgresqlWaitChain{}
	this.Nodes = nodes
	this.Edges = edges
	return &this
}

// NewPostgresqlWaitChainWithDefaults instantiates a new PostgresqlWaitChain object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlWaitChainWithDefaults() *PostgresqlWaitChain {
	this := PostgresqlWaitChain{}
	return &this
}

// GetNodes returns the Nodes field value.
func (o *PostgresqlWaitChain) GetNodes() []PostgresqlWaitChainNode {
	if o == nil {
		var ret []PostgresqlWaitChainNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitChain) GetNodesOk() (*[]PostgresqlWaitChainNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value.
func (o *PostgresqlWaitChain) SetNodes(v []PostgresqlWaitChainNode) {
	o.Nodes = v
}

// GetEdges returns the Edges field value.
func (o *PostgresqlWaitChain) GetEdges() []PostgresqlWaitChainEdge {
	if o == nil {
		var ret []PostgresqlWaitChainEdge
		return ret
	}
	return o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitChain) GetEdgesOk() (*[]PostgresqlWaitChainEdge, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edges, true
}

// SetEdges sets field value.
func (o *PostgresqlWaitChain) SetEdges(v []PostgresqlWaitChainEdge) {
	o.Edges = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlWaitChain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["nodes"] = o.Nodes
	toSerialize["edges"] = o.Edges

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PostgresqlWaitChain) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Nodes *[]PostgresqlWaitChainNode `json:"nodes"`
		Edges *[]PostgresqlWaitChainEdge `json:"edges"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Nodes == nil {
		return fmt.Errorf("required field nodes missing")
	}
	if all.Edges == nil {
		return fmt.Errorf("required field edges missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodes", "edges"})
	} else {
		return err
	}
	o.Nodes = *all.Nodes
	o.Edges = *all.Edges

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
