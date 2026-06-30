// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PostgresqlWaitGraph Wait graph related to the selected PID. It contains the selected session's ancestor and descendant wait edges that the backend can prove from the current snapshot; unrelated side branches are intentionally excluded.
type PostgresqlWaitGraph struct {
	Nodes []PostgresqlWaitGraphNode `json:"nodes"`
	Edges []PostgresqlWaitGraphEdge `json:"edges"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPostgresqlWaitGraph instantiates a new PostgresqlWaitGraph object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPostgresqlWaitGraph(nodes []PostgresqlWaitGraphNode, edges []PostgresqlWaitGraphEdge) *PostgresqlWaitGraph {
	this := PostgresqlWaitGraph{}
	this.Nodes = nodes
	this.Edges = edges
	return &this
}

// NewPostgresqlWaitGraphWithDefaults instantiates a new PostgresqlWaitGraph object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPostgresqlWaitGraphWithDefaults() *PostgresqlWaitGraph {
	this := PostgresqlWaitGraph{}
	return &this
}

// GetNodes returns the Nodes field value.
func (o *PostgresqlWaitGraph) GetNodes() []PostgresqlWaitGraphNode {
	if o == nil {
		var ret []PostgresqlWaitGraphNode
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitGraph) GetNodesOk() (*[]PostgresqlWaitGraphNode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value.
func (o *PostgresqlWaitGraph) SetNodes(v []PostgresqlWaitGraphNode) {
	o.Nodes = v
}

// GetEdges returns the Edges field value.
func (o *PostgresqlWaitGraph) GetEdges() []PostgresqlWaitGraphEdge {
	if o == nil {
		var ret []PostgresqlWaitGraphEdge
		return ret
	}
	return o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value
// and a boolean to check if the value has been set.
func (o *PostgresqlWaitGraph) GetEdgesOk() (*[]PostgresqlWaitGraphEdge, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edges, true
}

// SetEdges sets field value.
func (o *PostgresqlWaitGraph) SetEdges(v []PostgresqlWaitGraphEdge) {
	o.Edges = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PostgresqlWaitGraph) MarshalJSON() ([]byte, error) {
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
func (o *PostgresqlWaitGraph) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Nodes *[]PostgresqlWaitGraphNode `json:"nodes"`
		Edges *[]PostgresqlWaitGraphEdge `json:"edges"`
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
